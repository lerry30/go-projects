package jwt

import "time"

type contextKey string
const ContextKeyUser contextKey = "user"

// 
var Store = &RevokedStore{
	revoked: make(map[string]time.Time),
}