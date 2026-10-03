package jwt

import "os"

var SecretKey = []byte(os.Getenv("JWT_SECRET_KEY"))