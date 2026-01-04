package views

import (
	"fmt"
	"bufio"
	"strings"
	"strconv"
	"os"

	"contact-manager/utils"
	"contact-manager/models"
)

func DeleteContact(book *models.ContactBook, scanner *bufio.Scanner) {
	isDisplay := false

	for {
		utils.ClearScreen()
		fmt.Println("╔════════════════════════════════════════════════╗")
		fmt.Println("║                 DELETE CONTACT                 ║")
		fmt.Println("╚════════════════════════════════════════════════╝")
		fmt.Println()
		fmt.Println("[1] View All Contacts")
		fmt.Println("[2] Delete Contact")
		fmt.Println("[3] Exit")

		if isDisplay {
			fmt.Println()
			book.DisplayContacts()
			fmt.Println()
		}
		
		fmt.Print("\nEnter your choice [1-3]: ")
		scanner.Scan()
		input := strings.TrimSpace(scanner.Text())

		if input == "1" {
			isDisplay = true
		} else if input == "2" {
			fmt.Print("Enter contact ID to remove: ")

			scanner.Scan()
			strId := strings.TrimSpace(scanner.Text())

			id, err := strconv.Atoi(strId)
			if err != nil {
				fmt.Fprintf(os.Stderr, "ID not found: %s", strId)
				continue
			}

			for i, c := range book.Contacts {
				if c.ID == id {
					book.Contacts = append(book.Contacts[:i], book.Contacts[i+1:]...)
					book.Save()
					
					fmt.Println("Contact successfully removed")
					break
				}
			}
		} else if input == "3" {
			return
		} else {
			fmt.Println("Undefined input. Please try again.")
		}
	}
}