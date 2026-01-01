package views

import (
	"fmt"
	"bufio"
	"os"
	
	"payroll/utils"
	"payroll/controllers"
	"payroll/models"
)

func GenReport(scanner *bufio.Scanner) {
	for {
		utils.ClearScreen()
		fmt.Println("╔════════════════════════════════════════════════╗")
		fmt.Println("║             GENERATE PAYROLL REPORT            ║")
		fmt.Println("╚════════════════════════════════════════════════╝")
		fmt.Println("\nEnter 'cancel' to abort")

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

		controllers.GenerateReport(dateInput, &employeePayroll)

		utils.WaitToPressEnter(scanner)
	}
}