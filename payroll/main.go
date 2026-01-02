package main

import (
	"fmt"
	"math/rand"
	"time"
	"os"
	"bufio"
	"strings"

	"payroll/utils"
	"payroll/views"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	scanner := bufio.NewScanner(os.Stdin)

	for {
		utils.ClearScreen()
		fmt.Println("╔════════════════════════════════════════════════╗")
		fmt.Println("║          PAYROLL MANAGEMENT SYSTEM v1.0        ║")
		fmt.Println("╚════════════════════════════════════════════════╝")
		fmt.Println("\n[1] Clock In / Clock Out")
		fmt.Println("[2] Admin Dashboard")
		fmt.Println("[3] Exit")
		fmt.Print("\nEnter your choice [1-3]: ")

		scanner.Scan()
		choice := strings.TrimSpace(scanner.Text())

		switch choice {
		case "1":
			utils.ClearScreen()
			views.ClockInOut(scanner)
		case "2":
			utils.ClearScreen()
			// ADMIN AUTHENTICATION
			isAdmin := views.VerifyCred(scanner)

			if isAdmin {
				views.AdminMenu(scanner)
			} else {
				fmt.Println("Invalid credential.")
				utils.WaitToPressEnter(scanner)	
			}
		case "3":
			fmt.Println("\nBye bye! 👋")
			return
		default:
			fmt.Println("\nUndefined input. Please try again.")
			utils.WaitToPressEnter(scanner)	
		}
	}
}