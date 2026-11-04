package in_process_scheduling

import (
	"fmt"
	"time"
	"context"

	"tracker/internal/platform/logger"
	"tracker/internal/webhook"

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
				continue // channel close
			}

			dbUserSched, err := usersWithScheds.schedulesRepo.User(userID)
			if err != nil {
				logger.LogError("failed to get (user) data for setting up scheduler", err)
				continue
			}

			userKey := fmt.Sprintf("user-%d", userID)
			if _, ok := usersWithScheds.schedulers[userKey]; ok {
				ctxScheduler, cancel := context.WithTimeout(ctx, 5*time.Second)
				defer cancel()
				usersWithScheds.schedulers[userKey].Stop(ctxScheduler)
				delete(usersWithScheds.schedulers, userKey)
			}

			user, err := PopulateUserData(dbUserSched)
			if err != nil {
				logger.LogError("failed to create new user(runtime)", err)
				continue
			}

			scheduler := NewScheduler(usersWithScheds.loc, user)
			scheduler.Start()
			callback := webhook.NewWebhookAppscript(user.FirstName, user.LastName, user.Endpoint, user.Status, user.SpreadSheetID)
			_ = scheduler.Create(callback)

			usersWithScheds.schedulers[userKey] = scheduler

			fmt.Printf("%v\n", usersWithScheds.schedulers)
		case <-ctx.Done():
			return
		}
	}
}