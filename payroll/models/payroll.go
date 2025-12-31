package models

type Payroll struct {
	Employee *Employee
	TotalHoursWorked float64
	TotalOverTime float64
	GrossIncome float64
}