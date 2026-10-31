package in_process_scheduling

import (
	"time"
)

type UsersScheduleEntity struct {
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
	Login time.Time `db:"login"`
	FirstBreakOut time.Time `db:"first_break_out"`
	FirstBreakIn time.Time `db:"first_break_in"`
	LunchOut time.Time `db:"lunch_out"`
	LunchIn time.Time `db:"lunch_in"`
	SecondBreakOut time.Time `db:"second_break_out"`
	SecondBreakIn time.Time `db:"second_break_in"`
	Logout time.Time `db:"logout"`
}

// Table column names
type AppScriptsColumns struct {
	ID string
	EmployeeID string
	Endpoint string
	Status string
	CreatedAt string
	UpdatedAt string
}

type UsersColumns struct {
	ID string
	FirstName string
	LastName string
	Username string
	Password string
	CreatedAt string
	UpdatedAt string
}

//-----------------------------

var AppScriptsCols = AppScriptsColumns{
	ID: "id",
	EmployeeID: "employee_id",
	Endpoint: "appscript_endpoint",
	Status: "status",
	CreatedAt: "created_at",
	UpdatedAt: "updated_at",
}

var UserCols = UsersColumns{
	ID: "id",
	FirstName: "first_name",
	LastName: "last_name",
	Username: "username",
	Password: "password",
	CreatedAt: "created_at",
	UpdatedAt: "updated_at",
}