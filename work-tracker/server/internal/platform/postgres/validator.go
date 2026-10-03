package postgres

import (
	"fmt"
	"regexp"

	"github.com/jackc/pgx/v5"
)

/*
IDENTIFIERS
- Table names
- Column name
- ...
*/

const maxIdentifierLen = 63 // Postgres NAMEDATALEN-1; longer names get silently truncated

// validateIdentifier rejects anything that isn't a simple snake_case/alphanumeric name.
// This prevents SQL injection via table/column names since those can't be parameterized.
var validIdentifier = regexp.MustCompile(`^[a-z_][a-z0-9_]*$`)

func ValidateIdentifier(name string) error {
	if name == "" || len(name) > maxIdentifierLen {
		return fmt.Errorf("identifier must be 1-%d bytes, got %d", maxIdentifierLen, len(name))
	}
	if !validIdentifier.MatchString(name) {
		return fmt.Errorf("%q is not a valid SQL identifier", name)
	}
	return nil
}

// QuoteIdentifier validates each part and returns a safely quoted identifier.
// Pass one part for a column/table, or two for schema + table.
// We have to treat this something like this "name; drop table name;" as just 
// a string by wrapping it with quotes.
func QuoteIdentifier(parts ...string) (string, error) {
	if len(parts) == 0 {
		return "", fmt.Errorf("no identifier given")
	}
	for _, p := range parts {
		if err := ValidateIdentifier(p); err != nil {
			return "", err
		}
	}
	return pgx.Identifier(parts).Sanitize(), nil
}