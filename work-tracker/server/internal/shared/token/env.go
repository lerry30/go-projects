package token

import "os"

var SecretKey = []byte(os.Getenv("JWT_SECRET_KEY"))