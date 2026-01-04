package main

import (
	"fmt"
	"bufio"
	"os"
	"strings"

	"contact-manager/utils"
	"contact-manager/models"
	"contact-manager/views"
)

func main() {
	var book models.ContactBook
	book.Load()
	scanner := bufio.NewScanner(os.Stdin)

	for {
		utils.ClearScreen()
		fmt.Println("╔════════════════════════════════════════════════╗")
		fmt.Println("║                 CONTACT MANAGER                ║")
		fmt.Println("╚════════════════════════════════════════════════╝")
		fmt.Println()
		fmt.Println("[1] Add Contact")
		fmt.Println("[2] View All Contacts")
		fmt.Println("[3] Find Contact")
		fmt.Println("[4] Delete Contact")
		fmt.Println("[5] Exit")
		fmt.Print("\nEnter your choice [1-5]: ")

		scanner.Scan()
		input := strings.TrimSpace(scanner.Text())

		switch input {
		case "1":
			views.AddContact(&book, scanner)
		case "2":
			fmt.Println()
			book.DisplayContacts()
		case "3":
			views.SearchContact(&book, scanner)
		case "4":
			views.DeleteContact(&book, scanner)
		case "5":
			fmt.Println("\nBye bye! 👋")
			return
		default:
			fmt.Println("Undefined input. Please try again.")
		}

		utils.WaitToPressEnter(scanner)
	}
}