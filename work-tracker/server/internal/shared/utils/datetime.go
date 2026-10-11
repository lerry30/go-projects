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