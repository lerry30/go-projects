package auth

import (
	"os"
)

type contextKey string

const ContextKeyUser contextKey = "user"

// secret key — load this from an env variable
var SecretKey = []byte(os.Getenv("JWT_SECRET_KEY"))
