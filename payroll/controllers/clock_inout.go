package controllers

import (
	"fmt"
	"strings"
	"os"
	"io"
	"time"
	"log"

	"payroll/utils"	
)

func ClockInOutEmployee(id string) bool {
	id = strings.TrimSpace(id)

	next, close, err := utils.ReadFile("data/employees.txt")
	if err != nil {
		fmt.Fprintln(os.Stderr, "clock in/out - opening file error: employees", err)
		return false
	}
	defer close()

	isEmpExists := false
	for {
		row, err := next()
		if err == io.EOF {
			break
		}

		if err != nil {
			fmt.Fprintln(os.Stderr, "clock in/out - reading file error", err)
			break
		}

		// parse
		columns := strings.Split(row, " | ")
		if len(columns) < 6 {
			continue
		}

		if strings.TrimSpace(columns[1]) == id {
			isEmpExists = true
			break
		}
	}

	if isEmpExists {
		nextFInOutDay, closeFInOutDay, err := utils.ReadFile("data/clock_inout/day.txt")
		if err != nil {
			fmt.Fprintln(os.Stderr, "clock in/out - opening file error: clock_inout", err)
			return false
		}
		defer closeFInOutDay()

		line, err := nextFInOutDay()
		punchDate := strings.TrimSpace(line) // timestamp
		if err == nil || punchDate != "" {
			// is the date today
			t, err := utils.ParseDateString(punchDate)
			if err != nil {
				log.Fatal(err)
			}

			if t.Before(time.Now()) {
				moveAttTimeToMonth() // transfer content
				punchDate = attachDate() // insert date to file and return that date
			}
		} else {
			punchDate = attachDate()
		}

		var kept []string
		var employeeRec []string
		for {
			line, err := nextFInOutDay()
			if err == io.EOF {
				break
			}

			if err != nil {
				fmt.Fprintln(os.Stderr, "clock in/out - reading file error: clock_inout", err)
				break
			}

			// parse
			columns := strings.Split(line, " | ")
			if len(columns) < 3 {
				continue
			}

			if strings.TrimSpace(columns[0]) == id {
				employeeRec = columns
				continue
			}

			kept = append(kept, line)
		}

		timeNow := time.Now().Format("15:04:05")
		if len(employeeRec) > 0 {
			employeeRec[2] = "OUT: " + timeNow
		} else {
			employeeRec = append(employeeRec, id)
			employeeRec = append(employeeRec, "IN: " + timeNow)
			employeeRec = append(employeeRec, "OUT: -")
		}

		sEmpRec := strings.Join(employeeRec, " | ")
		kept = append(kept, sEmpRec)
		records := strings.Join(kept, "\n")
		content := punchDate + "\n" + records
		utils.FileContentOverride("data/clock_inout/day.txt", content)
		return true
	}
	return false
}

func moveAttTimeToMonth() {
	// get the attendance time in/out in day.txt to append in the month.txt
	next, close, err := utils.ReadFile("data/clock_inout/day.txt")
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error reading file: day on MoveAttTimeToMonth", err)
		return
	}
	defer close()

	attTime := "==="
	for {
		line, err := next()
		if err == io.EOF {
			break
		}

		if err != nil {
			fmt.Fprintln(os.Stderr, "Error reading file content", err)
			break
		}

		attTime = attTime + "\n" + line
	}

	utils.LogFile("data/clock_inout/month.txt", attTime, false)
}

func attachDate() string {
	now := time.Now().Format("2006-01-02T15:04:05Z")
	utils.FileContentOverride("data/clock_inout/day.txt", now)
	return now
}