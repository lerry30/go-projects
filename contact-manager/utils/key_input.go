package utils

/*
import (
	"fmt"
	"os"
	
	"golang.org/x/term"
)

func KeyStrokes() (func () (byte, error), func () error, error) {
	fd := int(os.Stdin.Fd())

	oldState, err := term.MakeRaw(fd)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to initialize term: %w", err)
	}

	restore := func() error {
		return term.Restore(fd, oldState)
	}

	buf := make([]byte, 1)
	getKey := func() (byte, error) {
		n, err := os.Stdin.Read(buf) // waiting for key press
		if err != nil {
			return 0, fmt.Errorf("failed to get key: %w", err)
		}

		if n > 0 {
			key := buf[0]

			// Ctrl+C
			if key == 3 {
				return 0, fmt.Errorf("terminated")
			}

			return key, nil
		}

		return 0, nil
	}

	return getKey, restore, nil
}
	*/



// views
/*
func SearchContact(book *models.ContactBook) {
	// init key press
	getKey, restore, err := utils.KeyStrokes()
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error:", err)
		return
	}

	defer restore() // restore the terminal to normal state

	utils.ClearScreen()
	fmt.Println("╔════════════════════════════════════════════════╗")
	fmt.Println("║                 SEARCH CONTACT                 ║")
	fmt.Println("╚════════════════════════════════════════════════╝")
	fmt.Println("\nCTRL+C to abort")

	fmt.Print("Enter name, email, or phone number: ")

	var input string
	for {
		// key press
		key, err := getKey() // return: (byte, error)
		if err != nil {
			fmt.Fprintln(os.Stderr, "Error:", err)
			break
		}

		if key == 127 { // backspace
			if len(input) == 0 {
				continue
			}

			input = input[:len(input)-1]
			continue
		}

		if key == 13 { // enter
			break
		}

		input = input + string(key)
	}
}
	*/