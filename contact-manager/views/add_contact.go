package views

import (
	"fmt"
	"bufio"
	"strings"

	"contact-manager/models"
	"contact-manager/utils"
)

func AddContact(book *models.ContactBook, scanner *bufio.Scanner) {
	var contact models.Contact
	var input string

	for {
		utils.ClearScreen()
		fmt.Println("╔════════════════════════════════════════════════╗")
		fmt.Println("║                   ADD CONTACT                  ║")
		fmt.Println("╚════════════════════════════════════════════════╝")
		fmt.Println("\n Type 'cancel' to abort")

		fmt.Print("Enter Name: ")
		scanner.Scan()
		input = strings.TrimSpace(scanner.Text())

		if input == "cancel" {
			break
		}

		contact.Name = input

		fmt.Println()

		fmt.Print("Enter Email: ")
		scanner.Scan()
		input = strings.TrimSpace(scanner.Text())

		if input == "cancel" {
			break
		}

		contact.Email = input

		fmt.Println()

		fmt.Print("Enter Phone: ")
		scanner.Scan()
		input = strings.TrimSpace(scanner.Text())

		if input == "cancel" {
			break
		}

		contact.Phone = input

		id := book.GetLastInsertedId() + 1
		contact.ID = id

		book.AddContact(contact)
		book.Save()

		fmt.Println("New contact successfully inserted.")
		break
	}
}