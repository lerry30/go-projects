package in_process_scheduling

import (
	"time"
	"fmt"
)

type User struct {
	UserID int32 `db:"user_id"`
	FirstName string `db:"first_name"`
	LastName string `db:"last_name"`
	Username string `db:"username"`
	
	AppScriptID int32 `db:"appscript_id"`
	Endpoint string `db:"endpoint"`
	Status string `db:"status"`

	TrackerID int32 `db:"tracker_id"`
	SpreadSheetID string `db:"spread_sheet_id"`
	Year int16 `db:"year"`
	Month int16 `db:"month"`

	TimeLogsID int32 `db:"time_logs_id"`
	SchedTimeLogs map[string][]time.Time
}

func (u *User) Spec(weekday string, index int) string {
	slot := u.SchedTimeLogs[weekday][index]
	return fmt.Sprintf("%d %d * * %d", slot.Minute(), slot.Hour(), int(slot.Weekday()))
}