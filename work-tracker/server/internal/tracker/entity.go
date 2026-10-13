package tracker

import (
	"time"
)

type TrackerEntity struct {
	ID int32 `db:"id"`
	EmployeeID int32 `db:"employee_id"`
	SpreadSheetID string `db:"spread_sheet_id"`
	Year int16 `db:"year"`
	Month int16 `db:"month"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}

type TrackerColumns struct {
	ID string
	EmployeeID string
	SpreadSheetID string
	Year string
	Month string
	CreatedAt string
	UpdatedAt string
}

var TableName = "tracker"

var TrackerCols = TrackerColumns{
	ID: "id",
	EmployeeID: "employee_id",
	SpreadSheetID: "spread_sheet_id",
	Year: "year",
	Month: "month",
	CreatedAt: "created_at",
	UpdatedAt: "updated_at",
}