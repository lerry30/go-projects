package repository

import (
	"fmt"
	"regexp"
)

// validateIdentifier rejects anything that isn't a simple snake_case/alphanumeric name.
// This prevents SQL injection via table/column names since those can't be parameterized.
var validIdentifier = regexp.MustCompile(`^[a-zA-Z_][a-zA-Z0-9_]*$`)

func ValidateIdentifier(name string) error {
	if !validIdentifier.MatchString(name) {
		return fmt.Errorf("%q is not a valid SQL identifier", name)
	}
	return nil
}
