package auth

import (
	"fmt"
	"net/http"
	"strconv"
)

// GetCurrentUser is a helper
func GetCurrentUser(r *http.Request) (*Claims, bool) {
	claims, ok := r.Context().Value(ContextKeyUser).(*Claims)
	return claims, ok
}

func GetCurrentUserID(r *http.Request) (int64, error) {
	idStr, ok := GetCurrentUser(r)
	if !ok {
		return 0, fmt.Errorf("user doesn't exist")
	}

	id, err := strconv.ParseInt(idStr.UserID, 10, 32)
	if err != nil {
		return 0, fmt.Errorf("invalid user")
	}

	return id, nil
}
