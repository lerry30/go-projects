package in_process_scheduling

import (
	"time"
	"fmt"
)

type User struct {
	UserID int32
	FirstName string
	LastName string
	Username string
	
	AppScriptID int32
	Endpoint string
	Status string

	TrackerID int32
	SpreadSheetID string
	Year int16
	Month int16

	TimeLogsID int32
	SchedTimeLogs map[string][]time.Time
}

func (u *User) Spec(weekday string, index int) string {
	slot := u.SchedTimeLogs[weekday][index]
	return fmt.Sprintf("%d %d * * %d", slot.Minute(), slot.Hour(), int(slot.Weekday()))
}

func (u *User) IsReadyForSchedule() bool {
	return u.UserID != 0 &&
		u.FirstName != "" &&
		u.LastName != "" &&
		u.Username != "" &&
		u.AppScriptID != 0 &&
		u.Endpoint != "" &&
		u.Status != "" &&
		u.TrackerID != 0 &&
		u.SpreadSheetID != "" &&
		u.Year != 0 &&
		u.Month != 0 &&
		u.TimeLogsID != 0
}

func (u *User) CreateSchedules(timeLogs []time.Time) {
	if !u.IsReadyForSchedule() {
		fmt.Println("Insufficient user data, not ready to create schedule!")
		return
	}

	// time.UTC - is the primary global time standard that does not change for seasons or time zones.
	dateRef := time.Date(2026, time.October, 26, 11, 45, 0, 0, time.Local) // Monday
	refs := []time.Time{
		dateRef.AddDate(0, 0, 0), // Monday
		dateRef.AddDate(0, 0, 1), // Tuesday
		dateRef.AddDate(0, 0, 2), // Wednesday
		dateRef.AddDate(0, 0, 3), // Thursday
		dateRef.AddDate(0, 0, 4), // Friday
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

	u.SchedTimeLogs = schedDays
}