package utils

import (
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