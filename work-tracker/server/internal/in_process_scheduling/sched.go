package in_process_scheduling

import (
	"context"
	"fmt"
	"time"

	"tracker/internal/platform/logger"
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
		logger.LogError("WARNING: failed to get appscripts running on db", err)
		return &UsersWithSchedules{}
	}

	schedulers := make(map[string]*Scheduler)
	for i := 0; i < len((*dbUsersSched)); i++ {
		user := PopulateUserData(&(*dbUsersSched)[i])

		scheduler := NewScheduler(loc, user)
		_ = scheduler.Create()

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

func PopulateUserData(rawUserData *UsersScheduleEntity) *User {
	var user User
	user.UserID = rawUserData.UserID
	user.FirstName = rawUserData.FirstName
	user.LastName = rawUserData.LastName
	user.Username = rawUserData.Username
	user.AppScriptID = rawUserData.AppScriptID
	user.Endpoint = rawUserData.Endpoint
	user.Status = rawUserData.Status
	user.TrackerID = rawUserData.TrackerID
	user.SpreadSheetID = rawUserData.SpreadSheetID
	user.Year = rawUserData.Year
	user.Month = rawUserData.Month
	user.TimeLogsID = rawUserData.TimeLogsID

	timeLogs := []time.Time{
		rawUserData.Login,
		rawUserData.FirstBreakOut,
		rawUserData.FirstBreakIn,
		rawUserData.LunchOut,
		rawUserData.LunchIn,
		rawUserData.SecondBreakOut,
		rawUserData.SecondBreakIn,
		rawUserData.Logout,
	}
	user.CreateSchedules(timeLogs)
	
	return &user
}