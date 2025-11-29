package utils

import (
	"fmt"
	"bufio"
	"strings"
)

func Prompt(message string, scanner *bufio.Scanner) (string, bool) {
	fmt.Printf("%s ", message)
	scanner.Scan()
	input := strings.TrimSpace(scanner.Text())
	if input == "cancel" {
		return "", true
	}

	return input, false
}