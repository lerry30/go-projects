package views

import (
	"fmt"
	"bufio"
	"strings"
	"os"
	
	"payroll/utils"
	"payroll/controllers"
	"payroll/models"
)

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

		var employeePayroll []models.Payroll
		if err := controllers.DisplayPayroll(dateInput, &employeePayroll); err != nil {
			fmt.Fprintln(os.Stderr, "Error:", err)
			utils.WaitToPressEnter(scanner)
			continue
		}

		//fmt.Printf("|    ID    |      Name    |     Position    |     Hours     |       OT        |        Gross Income     |\n")

		// Loop over each employee's payroll info
		for _, e := range employeePayroll {
			id := e.Employee.EmployeeId
			name := e.Employee.Name
			position := e.Employee.Position
			totalHoursWorked := e.TotalHoursWorked
			totalOT := e.TotalOverTime
			grossIncome := e.GrossIncome

			//fmt.Printf("|%-6s|%-6s|%-10s|%7.2f|%6.2f|%14.2f\n", id, name, position, totalHoursWorked, totalOT, grossIncome)

			fmt.Println()
			fmt.Printf("ID: %s\n", id)
			fmt.Printf("Name: %s\n", name)
			fmt.Printf("Position: %s\n", position)
			fmt.Printf("Hours: %.2f\n", totalHoursWorked)
			fmt.Printf("OT: %.2f\n", totalOT)
			fmt.Printf("Gross Income: $ %.2f\n", grossIncome)
			fmt.Println()

			fmt.Println(strings.Repeat("=", 60))
		}

		utils.WaitToPressEnter(scanner)
	}
}