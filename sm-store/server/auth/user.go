package auth

import (
	"net/http"
)

// GetCurrentUser is a helper
func GetCurrentUser(r *http.Request) (*Claims, bool) {
	claims, ok := r.Context().Value(ContextKeyUser).(*Claims)
	return claims, ok
}
