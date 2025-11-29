package controllers

import (
	"fmt"
	"time"
	"os"
	"io"
	"strings"
	"strconv"

	"payroll/models"
	"payroll/utils"
)

func AddEmployee(employee *models.Employee) {
	employee.EmployeeId = utils.GenerateRandomId()
	employee.Timestamp = time.Now()
	row := fmt.Sprintf("%s | %s | %s | %0.2f | %s",
		employee.EmployeeId,
		employee.Name, 
		employee.Position, 
		employee.Hourly_Rate,
		employee.Timestamp.Format("2006-01-02T15:04:05Z"))
	utils.LogFile("data/employees.txt", row, true)
}

func GetEmployees(employees *[]models.Employee)  {
	next, close, err := utils.ReadFile("data/employees.txt")
	if err != nil {
		fmt.Fprintln(os.Stderr, "Read file error: ", err)
		return
	}
	defer close()

	for {
		row, err := next()
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Fprintln(os.Stderr, "Scan file error: ", err)
			break
		}

		// parse
		columns := strings.Split(row, " | ")
		if len(columns) < 6 {
			continue
		}

		id, err := strconv.Atoi(strings.TrimSpace(columns[0]))
		if err != nil {
			fmt.Fprintln(os.Stderr, "string to int error", err)
			continue
		}

		employeeId := strings.TrimSpace(columns[1])
		name := strings.TrimSpace(columns[2])
		position := strings.TrimSpace(columns[3])

		rate, err := strconv.ParseFloat(strings.TrimSpace(columns[4]), 64)
		if err != nil {
			fmt.Fprintln(os.Stderr, "string to float error", err)
			continue
		}

		t, err := time.Parse("2006-01-02T15:04:05Z", strings.TrimSpace(columns[5]))
		if err != nil {
			fmt.Fprintln(os.Stderr, "String to date error: GetEmployees", err)
			continue
		}

		*employees = append(
			*employees,
			models.Employee{
				Id: id,
				EmployeeId: employeeId,
				Name: name,
				Position: position,
				Hourly_Rate: rate,
				Timestamp: t,
			})
	}
}