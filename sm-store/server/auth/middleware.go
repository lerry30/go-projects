package auth

import (
	"context"
	"net/http"
	"strings"
)

func Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Get token from "Authorization: Bearer <token>" header
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
			http.Error(w, "missing token", http.StatusUnauthorized)
			return
		}

		tokenString := strings.TrimPrefix(authHeader, "Bearer ")

		// Verify the token signature + expiry
		claims, err := VerifyToken(tokenString)
		if err != nil {
			http.Error(w, "invalid token", http.StatusUnauthorized)
			return
		}

		// Check if user has logged out
		if Store.IsRevoked(claims.ID) {
			http.Error(w, "token revoked, please log in again", http.StatusUnauthorized)
			return
		}

		// Attach user info to request context (like req.user in Express)
		ctx := context.WithValue(r.Context(), ContextKeyUser, claims)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
