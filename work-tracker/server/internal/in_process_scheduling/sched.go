package in_process_scheduling

import (
	"context"
	"fmt"
	"time"
	
	"tracker/internal/platform/logger"
	"tracker/internal/webhook"
)

type UsersWithSchedules struct {
	schedulesRepo SchedulesRepository
	schedulers map[string]*Scheduler
	loc *time.Location
	UserID chan int32
}

func NewUsersWithSchedules(schedulesRepo SchedulesRepository, loc *time.Location, buf int) *UsersWithSchedules {
	dbUsersSched, err := schedulesRepo.Users()
	if err != nil {
		logger.LogError("failed to get (users) data for setting up scheduler", err)
		return &UsersWithSchedules{}
	}

	schedulers := make(map[string]*Scheduler)
	for i := 0; i < len((*dbUsersSched)); i++ {
		user, err := PopulateUserData(&(*dbUsersSched)[i])
		if err != nil {
			logger.LogError("failed to create new user", err)
			continue
		}

		scheduler := NewScheduler(loc, user)
		callback := webhook.NewWebhookAppscript(user.FirstName, user.LastName, user.Endpoint, user.Status, user.SpreadSheetID)
		_ = scheduler.Create(callback)

		key := fmt.Sprintf("user-%d", user.UserID)
		schedulers[key] = scheduler
	}

	return &UsersWithSchedules{
		schedulesRepo: schedulesRepo,
		schedulers: schedulers,
		loc: loc,
		UserID: make(chan int32, buf),
	}
}

func (u *UsersWithSchedules) StartAll() {
	for _, s := range u.schedulers {
		s.Start()
	}
}

func (u *UsersWithSchedules) StopAll(ctx context.Context) error {
	for _, s := range u.schedulers {
		if err := s.Stop(ctx); err != nil {
			return err
		}
	}
	return nil
}