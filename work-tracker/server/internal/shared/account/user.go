package account

import (
	"fmt"
	"net/http"
	"strconv"

	"tracker/internal/shared/token"
	"tracker/internal/shared/token/jwt"
)

// CurrentUser is a helper
func CurrentUser(r *http.Request) (*token.Claims, bool) {
	claims, ok := r.Context().Value(jwt.ContextKeyUser).(*token.Claims)
	return claims, ok
}

func CurrentUserID(r *http.Request) (int64, error) {
	idStr, ok := CurrentUser(r)
	if !ok {
		return 0, fmt.Errorf("user doesn't exist")
	}

	id, err := strconv.ParseInt(idStr.UserID, 10, 32)
	if err != nil {
		return 0, fmt.Errorf("invalid user")
	}

	return id, nil
}