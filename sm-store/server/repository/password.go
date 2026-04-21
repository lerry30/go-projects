package repository

import (
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
	// Cost factor - bcrypt.DefaultCost = 10; range is 4–31 (higher = slower = more secure)
	// Higher cost for more security (e.g., sensitive systems)
	// hash, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

// password := "mysecretpassword"
// hash, err := HashPassword(password)
// if err != nil {
// 	panic(err)
// }
// fmt.Println("Hash:", hash)

// match := CheckPasswordHash(password, hash)
// fmt.Println("Match (correct):", match) // true

// Verify wrong password
// match = CheckPasswordHash("wrongpassword", hash)
// fmt.Println("Match (wrong):", match) // false
