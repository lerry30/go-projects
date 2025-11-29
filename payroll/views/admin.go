package views

import (
	"fmt"
	"bufio"

	"payroll/utils"
)

func AdminMenu(scanner *bufio.Scanner) {
	for {
		utils.ClearScreen()
		fmt.Println("╔════════════════════════════════════════════════╗")
		fmt.Println("║          PAYROLL MANAGEMENT SYSTEM v1.0        ║")
		fmt.Println("╚════════════════════════════════════════════════╝")
		fmt.Println("\n[1] Add Employee")
		fmt.Println("[2] View All Employees")
		fmt.Println("[3] Calculate Payroll")
		fmt.Println("[4] Generate Report")
		fmt.Println("[5] Exit")
		fmt.Print("\nEnter your choice [1-5]: ")

		scanner.Scan()
		choice := scanner.Text()

		switch choice {
		case "1":
			utils.ClearScreen()
			AddEmployee(scanner)
		case "2":
			utils.ClearScreen()
			DisplayEmployees(scanner)
		case "3":
			utils.ClearScreen()
			CalcPayroll(scanner)
		case "5":
			fmt.Println("\nBye bye! 👋")
			return
		default:
			fmt.Println("\nUndefined input. Please try again.")
			utils.WaitToPressEnter(scanner)		
		}
	}
}