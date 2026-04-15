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
