package views

import (
	"fmt"
	"bufio"

	"contact-manager/utils"
	"contact-manager/models"
)

func SearchContact(book *models.ContactBook, scanner *bufio.Scanner) {
	for {
		utils.ClearScreen()
		fmt.Println("╔════════════════════════════════════════════════╗")
		fmt.Println("║                 SEARCH CONTACT                 ║")
		fmt.Println("╚════════════════════════════════════════════════╝")
		fmt.Println("\n Type 'cancel' to abort")

		

	}
}