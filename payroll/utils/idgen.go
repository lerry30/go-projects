package utils

import (
	"math/rand"
)

func GenerateRandomId() string {
	letters := make([]byte, 3)
	for i := range letters {
		letters[i] = byte(rand.Intn(26) + 65)
	}

	digits := make([]byte, 4)
	for i := range digits {
		digits[i] = byte(rand.Intn(10) + 48)
	}

	return string(letters) + "-" + string(digits)
}