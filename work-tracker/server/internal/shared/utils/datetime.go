package utils

import (
	"fmt"
	"time"
)

func Now() string {
	return time.Now().Format("2006-01-02T15:04:05Z")
}

func DateNow() string {
	return time.Now().Format("2006-01-02")
}

func To24HourFormat(t time.Time) string {
	return t.Format("15:04:05")
}

func ParseStrTo24HourFormat(str string) (time.Time, error) {
	// enforce HH:MM:SS exactly; "1:00:00" would otherwise parse
	if len(str) != 8 {
		return time.Time{}, fmt.Errorf("invalid time %q, expected HH:MM:SS", str)
	}
	
	t, err := time.Parse("15:04:05", str)
	if err != nil {
		return time.Time{}, fmt.Errorf("invalid time %q, expected HH:MM:SS", str)
	}
	return t, nil
}