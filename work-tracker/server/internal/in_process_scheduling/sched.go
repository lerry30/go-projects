package in_process_scheduling

import (
	"time"

	"tracker/internal/platform/logger"
)

func Sched(schedulesRepo SchedulesRepository) []UserSched {
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

	usersSched := make([]UserSched, len((*dbUsersSched)))
	for userIndex, sched := range (*dbUsersSched) {
		timeLogs := []time.Time{
			sched.Login,
			sched.FirstBreakOut,
			sched.FirstBreakIn,
			sched.LunchOut,
			sched.LunchIn,
			sched.SecondBreakOut,
			sched.SecondBreakIn,
			sched.Logout,
		}

		/*for _, tm := range timeLogs {
			schedDay := tm.AddDate(dateRef.Year(), int(dateRef.Month()), dateRef.Day())
		}*/
		schedDays := make(map[string][]time.Time, len(refs))
		for _, wd := range refs {
			seq := make([]time.Time, len(timeLogs))
			for i, tl := range timeLogs {
				nDate := time.Date(wd.Year(), wd.Month(), wd.Day(), tl.Hour(), tl.Minute(), 0, 0, time.Local)
				if i > 0 && !nDate.After(seq[i-1]) {
					// There is a filtration on the TimeLogs endpoint so I don't worry about the extension of next other day
					nDate = nDate.AddDate(0, 0, 1) // Add a day
				}
				seq[i] = nDate
			}
			schedDays[wd.Weekday().String()] = seq
		}

		usersSched[userIndex].SchedTimeLogs = schedDays

		usersSched[userIndex].UserID = sched.UserID
		usersSched[userIndex].FirstName = sched.FirstName
		usersSched[userIndex].LastName = sched.LastName
		usersSched[userIndex].Username = sched.Username
		usersSched[userIndex].AppScriptID = sched.AppScriptID
		usersSched[userIndex].Endpoint = sched.Endpoint
		usersSched[userIndex].Status = sched.Status
		usersSched[userIndex].TrackerID = sched.TrackerID
		usersSched[userIndex].SpreadSheetID = sched.SpreadSheetID
		usersSched[userIndex].Year = sched.Year
		usersSched[userIndex].Month = sched.Month
		usersSched[userIndex].TimeLogsID = sched.TimeLogsID
	}

	return usersSched
}