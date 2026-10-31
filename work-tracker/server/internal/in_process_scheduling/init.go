package in_process_scheduling

import (
	"fmt"
	"time"
	"context"

	"tracker/internal/platform/logger"

	"github.com/jackc/pgx/v5/pgxpool"
)

func Init(db *pgxpool.Pool, loc *time.Location) *UsersWithSchedules {
	// Retrived data from the database
	schedulesService := NewSchedulesService(db)
	usersWithScheds := NewUsersWithSchedules(schedulesService, loc, 200)

	return usersWithScheds
}

func UpdateSchedules(ctx context.Context, usersWithScheds *UsersWithSchedules) {
	for {
		select {
		case userID, ok := <-usersWithScheds.UserID:
			if !ok {
				return // channel close
			}

			dbUserSched, err := usersWithScheds.schedulesRepo.User(userID)
			if err != nil {
				logger.LogError("WARNING: failed to get appscripts running on db", err)
				return
			}

			userKey := fmt.Sprintf("user-%d", userID)
			if _, ok := usersWithScheds.schedulers[userKey]; ok {
				ctxScheduler, cancel := context.WithTimeout(ctx, 5*time.Second)
				defer cancel()
				usersWithScheds.schedulers[userKey].Stop(ctxScheduler)
				delete(usersWithScheds.schedulers, userKey)
			}

			user := PopulateUserData(dbUserSched)

			scheduler := NewScheduler(usersWithScheds.loc, user)
			_ = scheduler.Create()

			usersWithScheds.schedulers[userKey] = scheduler
		case <-ctx.Done():
			return
		}
	}
}