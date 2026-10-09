package dashboard

import (
	"time"
)

type EmployeeTimeLogsEntity struct {
	UserID int32 `db:"user_id"`
	FirstName string `db:"first_name"`
	LastName string `db:"last_name"`
	Username string `db:"username"`
	UserCreatedAt time.Time `db:"user_created_at"`
	UserUpdatedAt time.Time `db:"user_updated_at"`
	TimeID int32 `db:"time_id"`
	Login time.Time `db:"time_login"`
	FirstBreakOut time.Time `db:"time_first_break_out"`
	FirstBreakIN time.Time `db:"time_first_break_in"`
	LunchOut time.Time `db:"time_lunch_out"`
	LunchIn time.Time `db:"time_lunch_in"`
	SecondBreakOut time.Time `db:"time_second_break_out"`
	SecondBreakIn time.Time `db:"time_second_break_in"`
	TimeCreatedAt time.Time `db:"time_created_at"`
	TimeUpdatedAt time.Time `db:"time_updated_at"`
	Logout time.Time `db:"time_logout"`
	TrackerID int32 `db:"tracker_id"`
	TrackerSpreadSheetID string `db:"tracker_spread_sheet_id"`
	TrackerYear int16 `db:"tracker_year"`
	TrackerMonth int16 `db:"tracker_month"`
	TrackerCreatedAt time.Time `db:"tracker_created_at"`
	TrackerUpdatedAt time.Time `db:"tracker_updated_at"`
	SecretID int32 `db:"secret_id"`
	SecretKey string `db:"secret_key"`
	SecretCreatedAt time.Time `db:"secret_created_at"`
	SecretUpdatedAt time.Time `db:"secret_updated_at"`
}

var queryAllStmt = `
	SELECT
		-- user entry
		u.id AS user_id,
		u.first_name,
		u.last_name,
		u.username,
		u.created_at AS user_created_at,
		u.updated_at AS user_updated_at,

		-- time logs entry
		tm.id                AS time_id,
		tm.login             AS time_login,
		tm.first_break_out   AS time_first_break_out,
		tm.first_break_in    AS time_first_break_in,
		tm.lunch_out         AS time_lunch_out,
		tm.lunch_in          AS time_lunch_in,
		tm.second_break_out  AS time_second_break_out,
		tm.second_break_in   AS time_second_break_in,
		tm.logout            AS time_logout,
		tm.created_at        AS time_created_at,
		tm.updated_at        AS time_updated_at,

		-- tracker (monthly spreadsheet)
		tr.id                AS tracker_id,
		tr.spread_sheet_id   AS tracker_spread_sheet_id,
		tr.year              AS tracker_year,
		tr.month             AS tracker_month,
		tr.created_at        AS tracker_created_at,
		tr.updated_at        AS tracker_updated_at,

		-- secrets
		s.id                 AS secret_id,
		s.secret_key         AS secret_key,
		s.created_at         AS secret_created_at,
		s.updated_at         AS secret_updated_at
	FROM users AS u
	LEFT JOIN time_logs AS tm
		ON tm.employee_id = u.id
	LEFT JOIN tracker AS tr
		ON tr.employee_id = u.id
	LEFT JOIN secrets AS s
		ON s.employee_id = u.id
	WHERE u.id = @user_id;
`