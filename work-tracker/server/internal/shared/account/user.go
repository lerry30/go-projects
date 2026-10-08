package account

import (
	"fmt"
	"net/http"
	"strconv"

	"tracker/internal/shared/token"
	"tracker/internal/shared/token/jwt"
)

// GetCurrentUser is a helper
func GetCurrentUser(r *http.Request) (*token.Claims, bool) {
	claims, ok := r.Context().Value(jwt.ContextKeyUser).(*token.Claims)
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