package in_process_scheduling

import (
	"time"

	"tracker/internal/platform/logger"
)

type UsersWithSchedules struct {
	users []User
}

func NewUsersWithSchedules(schedulesRepo SchedulesRepository) *UsersWithSchedules {
	dbUsersSched, err := schedulesRepo.RunningAppScripts()
	if err != nil {
		logger.LogError("WARNING: failed to get appscripts running on db", err)
	}

	//now := time.Now()
	// time.UTC - is the primary global time standard that does not change for seasons or time zones.
	dateRef := time.Date(2026, time.October, 26, 11, 45, 0, 0, time.Local) // Monday
	refs := []time.Time{
		dateRef.AddDate(0, 0, 0), // Monday
		dateRef.AddDate(0, 0, 1), // Tuesday
		dateRef.AddDate(0, 0, 2), // Wednesday
		dateRef.AddDate(0, 0, 3), // Thursday
		dateRef.AddDate(0, 0, 4), // Friday
	}

	users := make([]User, len((*dbUsersSched)))
	for userIndex, user := range (*dbUsersSched) {
		timeLogs := []time.Time{
			user.Login,
			user.FirstBreakOut,
			user.FirstBreakIn,
			user.LunchOut,
			user.LunchIn,
			user.SecondBreakOut,
			user.SecondBreakIn,
			user.Logout,
		}

		schedDays := make(map[string][]time.Time, len(refs))
		for _, wd := range refs {
			seq := make([]time.Time, len(timeLogs))
			for i, tl := range timeLogs {
				nDate := time.Date(wd.Year(), wd.Month(), wd.Day(), tl.Hour(), tl.Minute(), 0, 0, time.Local)
				if i > 0 && !nDate.After(seq[i-1]) {
					// There is a filtration on the TimeLogs endpoint so I don't worry about the expansion of more than a day, that is why adding a single day is enough in this case.
					nDate = nDate.AddDate(0, 0, 1) // Add a day
				}
				seq[i] = nDate
			}
			schedDays[wd.Weekday().String()] = seq
		}

		users[userIndex].SchedTimeLogs = schedDays

		users[userIndex].UserID = user.UserID
		users[userIndex].FirstName = user.FirstName
		users[userIndex].LastName = user.LastName
		users[userIndex].Username = user.Username
		users[userIndex].AppScriptID = user.AppScriptID
		users[userIndex].Endpoint = user.Endpoint
		users[userIndex].Status = user.Status
		users[userIndex].TrackerID = user.TrackerID
		users[userIndex].SpreadSheetID = user.SpreadSheetID
		users[userIndex].Year = user.Year
		users[userIndex].Month = user.Month
		users[userIndex].TimeLogsID = user.TimeLogsID
	}

	return &UsersWithSchedules{
		users: users,
	}
}