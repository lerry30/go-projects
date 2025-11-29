package views

import (
	"fmt"
	"bufio"
	"strconv"
	"strings"

	"payroll/utils"
	"payroll/models"
	"payroll/controllers"
)

func AddEmployee(scanner *bufio.Scanner) {
	var employee models.Employee

	for {
		utils.ClearScreen()
		fmt.Print("\n")
		fmt.Println("╔════════════════════════════════════════════════╗")
		fmt.Println("║                ADD NEW EMPLOYEE                ║")
		fmt.Println("╚════════════════════════════════════════════════╝")
		fmt.Println("Enter 'cancel' to abort")

		if employee.Name == "" {
			name, abort := utils.Prompt("\nEnter Full Name: ", scanner)
			if abort {
				break
			}

			if name == "" {
				continue
			}
			employee.Name = name
		} else {
			fmt.Println("\nName: ", employee.Name)
		}

		if employee.Position == "" {
			position, abort := utils.Prompt("Enter Position: ", scanner)
			if abort {
				break
			}

			if position == "" {
				continue
			}
			employee.Position = position
		} else {
			fmt.Println("Position: ", employee.Position)
		}

		if employee.Hourly_Rate == 0 {
			input, abort := utils.Prompt("Enter Hourly Rate ($/hr): ", scanner)
			if abort {
				break
			}

			rate, err := strconv.ParseFloat(input, 64)
			if err != nil || rate <= 0 {
				continue
			}
			employee.Hourly_Rate = rate
		} else {
			fmt.Println("Hourly Rate: ", employee.Hourly_Rate)
		}

		controllers.AddEmployee(&employee)
		fmt.Println("Employee added successfully 🎉")

		utils.WaitToPressEnter(scanner)
		break
	}
}

func DisplayEmployees(scanner *bufio.Scanner) {
	var employees []models.Employee
	controllers.GetEmployees(&employees)

	fmt.Print("\n")
	fmt.Println("╔════════════════════════════════════════════════╗")
	fmt.Println("║                    EMPLOYEES                   ║")
	fmt.Println("╚════════════════════════════════════════════════╝")

	for _, employee := range employees {
		fmt.Print("\n")
		fmt.Println("ID: ", employee.EmployeeId)
		fmt.Println("Name: ", employee.Name)
		fmt.Println("Position: ", employee.Position)
		fmt.Println("Hourly Rate: ", employee.Hourly_Rate)
		fmt.Println("Date Started: ", employee.Timestamp.Format("Jan 02, 2006 3:04 PM"))
		fmt.Println(strings.Repeat("-", 60))
		fmt.Print("\n\n")
	}

	utils.WaitToPressEnter(scanner)
}

/*
	
*/
func CalcPayroll(scanner *bufio.Scanner) {

	for {
		utils.ClearScreen()
		fmt.Print("\n")
		fmt.Println("╔════════════════════════════════════════════════╗")
		fmt.Println("║                     PAYROLL                    ║")
		fmt.Println("╚════════════════════════════════════════════════╝")
		fmt.Println("Enter 'cancel' to abort")

		dateInput, abort := utils.Prompt("Enter date(MM/YYYY): ", scanner)
		if abort {
			break
		}

		if dateInput == "" {
			continue
		}

		utils.WaitToPressEnter(scanner)
	}
}