package timelogs

import (
	"time"
)

type TimeLogsEntity struct {
	ID int32 `db:"id"`
	EmployeeID int32 `db:"employee_id"`
	Login time.Time `db:"login"`
	FirstBreakOut time.Time `db:"first_break_out"`
	FirstBreakIn time.Time `db:"first_break_in"`
	LunchOut time.Time `db:"lunch_out"`
	LunchIn time.Time `db:"lunch_in"`
	SecondBreakOut time.Time `db:"second_break_out"`
	SecondBreakIn time.Time `db:"second_break_in"`
	Logout time.Time `db:"logout"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}

type TimeLogsColumns struct {
	ID string
	EmployeeID string
	Login string
	FirstBreakOut string
	FirstBreakIn string
	LunchOut string
	LunchIn string
	SecondBreakOut string
	SecondBreakIn string
	Logout string
	CreatedAt string
	UpdatedAt string
}

var TableName = "time_logs"

var TimeLogsCols = TimeLogsColumns{
	ID: "id",
	EmployeeID: "employee_id",
	Login: "login",
	FirstBreakOut: "first_break_out",
	FirstBreakIn: "first_break_in",
	LunchOut: "lunch_out",
	LunchIn: "lunch_in",
	SecondBreakOut: "second_break_out",
	SecondBreakIn: "second_break_in",
	Logout: "logout",
	CreatedAt: "created_at",
	UpdatedAt: "updated_at",
}