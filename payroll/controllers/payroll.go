package controllers

import (
	"fmt"
	"os"
	"io"
	"strings"
	"regexp"
	"time"
	"strconv"
	
	"payroll/utils"
	"payroll/models"
)

func DisplayPayroll(date string, employeePayroll *[]models.Payroll) {
	year, month, err := utils.ParseMonthYear(date)
	attendance := make(map[string][]string)

	if err != nil {
		fmt.Fprintln(os.Stderr, "Error: ", err)
		return
	}

	next, close, err := utils.ReadFile("data/clock_inout/month.txt")
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error: opening file error at DisplayPayroll", err)
		return
	}
	defer close()

	isCaptured := false
	for {
		line, err := next()
		if err == io.EOF {
			break
		}

		if err != nil {
			fmt.Fprintln(os.Stderr, "Error: reading file content at DisplayPayroll - all", err)
			return
		}
		
		line = strings.TrimSpace(line)

		if line == "===" {
			isCaptured = false

			line2Date, err := next()
			if err == io.EOF {
				break
			}
			
			if err != nil {
				fmt.Fprintln(os.Stderr, "Error: reading file content at DisplayPayroll - date", err)
				return
			}

			t, err := utils.ParseDateString(strings.TrimSpace(line2Date))
			if err != nil {
				fmt.Fprintln(os.Stderr, "Invalid date: ", err)
				return
			}

			if t.Year() == year && int(t.Month()) == month {
				isCaptured = true
				continue
			}
		}

		if isCaptured {
			columns := strings.Split(line, " | ")

			if len(columns) < 3 {
				continue
			}

			employeedId := columns[0]
			in := columns[1]
			out := columns[2]

			attTime := in + " | " + out
			attendance[employeedId] = append(attendance[employeedId], attTime)
		}
	}

	employees := make(map[string]*models.Employee)
	getEmployees(employees)

	for id, logs := range attendance {
		var employee *models.Employee
		employee = employees[id]

		hourlyRate := employee.Hourly_Rate

		var totalHoursWorked, totalRegWorked, totalOT, grossIncome float64

		//fmt.Printf("%v\n", employee)
		for _, log := range logs {
			hours, err := calculateHoursWorked(log)
			if err != nil {
				fmt.Fprintln(os.Stderr, "Error: calculating hours worked", err)
				continue
			}

			if hours > 8 {
				overTime := hours - 8.0
				totalOT += overTime
			}

			totalHoursWorked += hours
			
			//fmt.Printf("\t%s: %2f - overtime: %.2f \n", log, hours, overTime)
		}

		totalRegWorked = totalHoursWorked - totalOT
		grossIncome = (totalRegWorked * hourlyRate) + ((totalOT * hourlyRate) * 1.5)

		*employeePayroll = append(*employeePayroll, models.Payroll{
			Employee: employee,
			TotalHoursWorked: totalHoursWorked,
			TotalOverTime: totalOT,
			GrossIncome: grossIncome,
		})
	}
}

func calculateHoursWorked(log string) (float64, error) {
	log = strings.TrimSpace(log)
	log = regexp.MustCompile(`\s*\|\s*`).ReplaceAllString(log, " | ")

	re := regexp.MustCompile(`IN:\s*(\d{2}:\d{2}:\d{2})\s*\|\s*OUT:\s*(\d{2}:\d{2}:\d{2})`)
	matches := re.FindStringSubmatch(log)

	if len(matches) != 3 {
		return 0, fmt.Errorf("invalid format: could not parse IN and OUT times")
	}

	inTimeStr := matches[1]
	outTimeStr := matches[2]

	layout := "15:04:05"
	inTime, err := time.Parse(layout, inTimeStr)
	if err != nil {
		return 0, fmt.Errorf("invalid IN time: %v", err)
	}
	outTime, err := time.Parse(layout, outTimeStr)
	if err != nil {
		return 0, fmt.Errorf("invalid OUT time: %v", err)
	}

	if outTime.Before(inTime) {
		outTime = outTime.Add(24 * time.Hour)
	}

	duration := outTime.Sub(inTime)
	hours := duration.Hours()

	return hours, nil
}

func getEmployees(employees map[string]*models.Employee) error {
	next, close, err := utils.ReadFile("data/employees.txt")
	if err != nil {
		return fmt.Errorf("reading file employees: %s", err)
	}
	defer close()

	for {
		line, err := next()
		if err == io.EOF {
			break
		}

		if err != nil {
			return fmt.Errorf("reading file content error: %s", err)
		}

		line = strings.TrimSpace(line)
		line = regexp.MustCompile(`\s*\|\s*`).ReplaceAllString(line, " | ")

		columns := strings.Split(line, " | ")
		if len(columns) < 6 {
			fmt.Fprintln(os.Stderr, "broken employee record")
			continue
		}

		employeeId := strings.TrimSpace(columns[1]);

		hourlyRate, err := strconv.ParseFloat(strings.TrimSpace(columns[4]), 64)
		if err != nil {
			fmt.Fprintln(os.Stderr, "unable to convert string to float", err)
			continue
		}

		t, err := time.Parse("2006-01-02T15:04:05Z", strings.TrimSpace(columns[5]))
		if err != nil {
			fmt.Fprintln(os.Stderr, "unable to parse string to date", err)
			continue
		}

		employees[employeeId] = &models.Employee{
			Id: 0, // I've set it to 0 since this is not necessary in this case
			EmployeeId: employeeId,
			Name: strings.TrimSpace(columns[2]),
			Position: strings.TrimSpace(columns[3]),
			Hourly_Rate: hourlyRate,
			Timestamp: t,
		}
	}

	return nil
}