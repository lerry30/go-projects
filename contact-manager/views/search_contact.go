package views

import (
	"fmt"
	"bufio"
	"strings"

	"contact-manager/utils"
	"contact-manager/models"
)

func SearchContact(book *models.ContactBook, scanner *bufio.Scanner) {
	
	for {
		utils.ClearScreen()
		fmt.Println("╔════════════════════════════════════════════════╗")
		fmt.Println("║                 SEARCH CONTACT                 ║")
		fmt.Println("╚════════════════════════════════════════════════╝")
		fmt.Println("\nEnter 'cancel' to abort")

		fmt.Print("\nEnter name, email, or phone number: ")
		
		scanner.Scan()
		input := strings.ToLower(strings.TrimSpace(scanner.Text()))

		if input == "" {
			continue
		}

		if input == "cancel" {
			break
		}

		fmt.Println()

		found := false
		for _, c := range book.Contacts {
			if strings.Contains(strings.ToLower(c.Name), input) ||
				strings.Contains(strings.ToLower(c.Email), input) ||
				strings.Contains(strings.ToLower(c.Phone), input) {
				
				found = true
					
				fmt.Printf("ID: %d | Name: %s | Email: %s | Phone: %s", 
					c.ID, c.Name, c.Email, c.Phone)
			}
		}

		if !found {
			fmt.Printf("%s not found\n", input)
		}

		fmt.Println()
		utils.WaitToPressEnter(scanner)
	}
}