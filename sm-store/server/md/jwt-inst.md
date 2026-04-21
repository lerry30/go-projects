Perfect! That gives me a clear picture. Here's what you need:

- **REST API** setup (Go backend + React frontend)
- **Explicit logout** → you need **token revocation**
- **Single app for now** → HS256 (simple secret key) is fine, RS256 is overkill until you split services
- **Kubernetes homelab** → you'll eventually want Redis for revocation, but we'll start with **in-memory** so you don't need extra infra yet

---

## Complete JWT Implementation in Go

### 1. Install the dependency

```bash
go get github.com/golang-jwt/jwt/v5
```

---

### 2. Project structure

```
/auth
  jwt.go        ← create/verify tokens
  middleware.go ← protect routes
  store.go      ← token revocation (in-memory for now)
```

---

### 3. `auth/jwt.go` — Create & Verify Tokens

This is the equivalent of what you did in Node.js with `jsonwebtoken`:

```go
package auth

import (
    "fmt"
    "time"

    "github.com/golang-jwt/jwt/v5"
)

// Your secret key — load this from an env variable in real use
var SecretKey = []byte("your-secret-key")

// Claims is the data you store inside the token (like jwt.sign payload in Node.js)
type Claims struct {
    UserID string `json:"user_id"`
    Email  string `json:"email"`
    jwt.RegisteredClaims
}

// GenerateToken is like jwt.sign() in Node.js
func GenerateToken(userID string, email string) (string, error) {
    claims := Claims{
        UserID: userID,
        Email:  email,
        RegisteredClaims: jwt.RegisteredClaims{
            ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)), // expires in 24hrs
            IssuedAt:  jwt.NewNumericDate(time.Now()),
            ID:        userID + fmt.Sprintf("%d", time.Now().UnixNano()), // unique token ID (jti)
        },
    }

    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
    return token.SignedString(SecretKey)
}

// VerifyToken is like jwt.verify() in Node.js
func VerifyToken(tokenString string) (*Claims, error) {
    claims := &Claims{}

    token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
        // Prevent "alg: none" attack — always check the signing method
        if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
            return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
        }
        return SecretKey, nil
    })

    if err != nil || !token.Valid {
        return nil, fmt.Errorf("invalid token")
    }

    return claims, nil
}
```

---

### 4. `auth/store.go` — Token Revocation (for Logout)

This is what makes logout actually work. When a user logs out, you **blacklist** their token so it can't be reused even if it hasn't expired yet.

```go
package auth

import (
    "sync"
    "time"
)

// In-memory store — simple for now, swap with Redis later on Kubernetes
type RevokedStore struct {
    mu      sync.RWMutex
    revoked map[string]time.Time // jti → expiry time
}

var Store = &RevokedStore{
    revoked: make(map[string]time.Time),
}

// Revoke adds the token's jti to the blacklist (called on logout)
func (s *RevokedStore) Revoke(jti string, expiry time.Time) {
    s.mu.Lock()
    defer s.mu.Unlock()
    s.revoked[jti] = expiry
}

// IsRevoked checks if the token has been logged out
func (s *RevokedStore) IsRevoked(jti string) bool {
    s.mu.RLock()
    defer s.mu.RUnlock()
    expiry, exists := s.revoked[jti]
    if !exists {
        return false
    }
    // Auto-clean: if the token is already expired, it doesn't matter
    if time.Now().After(expiry) {
        return false
    }
    return true
}
```

---

### 5. `auth/middleware.go` — Protect Your Routes

This is like `express-jwt` or your custom `verifyToken` middleware in Node.js/Express:

```go
package auth

import (
    "context"
    "net/http"
    "strings"
)

type contextKey string
const ContextKeyUser contextKey = "user"

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

// GetCurrentUser is a helper — use this in your handlers like req.user in Express
func GetCurrentUser(r *http.Request) (*Claims, bool) {
    claims, ok := r.Context().Value(ContextKeyUser).(*Claims)
    return claims, ok
}
```

---

### 6. Wiring it all together — `main.go`

```go
package main

import (
    "encoding/json"
    "net/http"
    "yourapp/auth"
)

func main() {
    mux := http.NewServeMux()

    // Public routes (no token needed)
    mux.HandleFunc("POST /login", loginHandler)

    // Protected routes — wrap with auth.Middleware
    mux.Handle("GET /profile", auth.Middleware(http.HandlerFunc(profileHandler)))
    mux.Handle("POST /logout", auth.Middleware(http.HandlerFunc(logoutHandler)))

    http.ListenAndServe(":8080", mux)
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
    // After verifying password from DB...
    token, err := auth.GenerateToken("user-123", "user@email.com")
    if err != nil {
        http.Error(w, "could not generate token", http.StatusInternalServerError)
        return
    }
    json.NewEncoder(w).Encode(map[string]string{"token": token})
}

func profileHandler(w http.ResponseWriter, r *http.Request) {
    user, _ := auth.GetCurrentUser(r) // like req.user in Express
    json.NewEncoder(w).Encode(map[string]string{
        "user_id": user.UserID,
        "email":   user.Email,
    })
}

func logoutHandler(w http.ResponseWriter, r *http.Request) {
    user, _ := auth.GetCurrentUser(r)
    // Blacklist this token so it can't be reused
    auth.Store.Revoke(user.ID, user.ExpiresAt.Time)
    w.WriteHeader(http.StatusOK)
    json.NewEncoder(w).Encode(map[string]string{"message": "logged out"})
}
```

---

## Node.js → Go Cheat Sheet

| Node.js / Express | Go equivalent |
|---|---|
| `jwt.sign(payload, secret)` | `auth.GenerateToken(userID, email)` |
| `jwt.verify(token, secret)` | `auth.VerifyToken(tokenString)` |
| `req.user` | `auth.GetCurrentUser(r)` |
| `verifyToken` middleware | `auth.Middleware` |
| Express `app.use()` | `mux.Handle()` wrapping |

---

## When to Upgrade to Redis

Your in-memory store works fine for now but has one limitation — **if your server restarts, the revocation list is wiped**, meaning logged-out tokens briefly become valid again. When you're ready on Kubernetes, just swap `store.go` with a Redis-backed version and nothing else in your code changes.