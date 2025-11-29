package models

import (
	"time"
)

type Employee struct {
	Id int
	EmployeeId string
	Name string
	Position string
	Hourly_Rate float64
	Timestamp time.Time
}