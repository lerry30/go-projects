package utils

import (
	"fmt"
	"bufio"
)

func WaitToPressEnter(scanner *bufio.Scanner) {
	fmt.Println("\nPress enter to continue...")
	scanner.Scan()
}

func ClearScreen() {
	fmt.Print("\033[H\033[2J")
}