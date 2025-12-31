package models

type Logs struct {
	InOut string
	HoursWorked float64
	OverTime float64
}

type Attendance struct {
	EmployeeId string
	Logs []Logs
}