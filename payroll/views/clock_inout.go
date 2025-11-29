package views

import (
	"fmt"
	"bufio"
	"strings"

	"payroll/utils"
	"payroll/controllers"
)

func ClockInOut(scanner *bufio.Scanner) {
	for {
		utils.ClearScreen()
		fmt.Println("╔════════════════════════════════════════════════╗")
		fmt.Println("║               CLOCK IN / CLOCK OUT             ║")
		fmt.Println("╚════════════════════════════════════════════════╝")
		fmt.Print("\nEnter your ID to clock in/out (or 'exit' to exit): ")

		scanner.Scan()
		input := strings.TrimSpace(scanner.Text())

		if strings.ToLower(input) == "exit" {
			return
		}

		controllers.ClockInOutEmployee(input)
		utils.WaitToPressEnter(scanner)
	}
}