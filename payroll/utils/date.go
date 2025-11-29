package utils

import (
	"fmt"
	"time"
)

// ParseDateString tries common formats and returns time.Time + error
func ParseDateString(s string) (time.Time, error) {
    formats := []string{
        time.RFC3339,
        time.RFC3339Nano,
        "2006-01-02 15:04:05",
        "2006-01-02T15:04:05",
        "2006-01-02 15:04:05Z",
        "2006-01-02",
        "02/01/2006",
        "01/02/2006", // US format
    }

    for _, layout := range formats {
        if t, err := time.Parse(layout, s); err == nil {
            return t, nil
        }
    }
    return time.Time{}, fmt.Errorf("unable to parse date: %s", s)
}

// usage
/*
	t, err := ParseDateString("2025-04-15T10:30:00Z")
	if err != nil {
		log.Fatal(err)
	}
	if t.Before(time.Now()) {
		fmt.Println("This date has already passed!")
	}
*/