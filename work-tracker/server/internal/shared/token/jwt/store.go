package jwt

import (
	"sync"
	"time"
)

type RevokedStore struct {
	mu sync.Mutex
	revoked map[string]time.Time
}

func (s *RevokedStore) Revoke(jti string, expiry time.Time) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.revoked[jti] = expiry
}

func (s *RevokedStore) IsRevoked(jti string) bool {
	s.mu.Lock()
	defer s.mu.Unlock()

	expiry, exists := s.revoked[jti]
	if !exists {
		return false
	}

	if time.Now().After(expiry) {
		return false
	}

	return true
}