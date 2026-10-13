package appscripts

import (
	"time"
)

type AppScriptsEntity struct {
	ID int32 `db:"id"`
	EmployeeID int32 `db:"employee_id"`
	Endpoint string `db:"appscript_endpoint"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}

// Table column names
type AppScriptsColumns struct {
	ID string
	EmployeeID string
	Endpoint string
	CreatedAt string
	UpdatedAt string
}

const TableName = "appscripts"

var AppScriptsCols = AppScriptsColumns{
	ID: "id",
	EmployeeID: "employee_id",
	Endpoint: "appscript_endpoint",
	CreatedAt: "created_at",
	UpdatedAt: "updated_at",
}