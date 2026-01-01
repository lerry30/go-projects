package views

import (
	"fmt"
	"bufio"
	"os"
	"strings"

	"payroll/utils"
)

func VerifyCred(scanner *bufio.Scanner) bool {
	utils.ClearScreen()
	fmt.Println("╔════════════════════════════════════════════════╗")
	fmt.Println("║                    LOG IN                      ║")
	fmt.Println("╚════════════════════════════════════════════════╝")
	fmt.Println("\nEnter 'cancel' to abort")

	password, abort := utils.Prompt("Enter password: ", scanner)
	if abort || password == "" {
		return false
	}

	next, close, err := utils.ReadFile("data/admin.txt")
	if err != nil {
		fmt.Fprintln(os.Stderr, "Admin credential unavailable", err)
		return false
	}
	defer close()

	adminCred, err := next()
	if err != nil {
		fmt.Fprintln(os.Stderr, "Admin credential unavailable", err)
		return false
	}

	adminCred = strings.TrimSpace(adminCred)

	if password == adminCred {
		return true
	}

	return false
}