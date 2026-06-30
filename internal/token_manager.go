package internal

import "sync"

// TokenManager holds the current OAuth access token and allows thread-safe rotation while requests are in flight.
type TokenManager struct {
	mu    sync.RWMutex
	token string
}

func NewTokenManager(token string) *TokenManager {
	return &TokenManager{token: token}
}

func (m *TokenManager) Get() string {
	m.mu.RLock()
	defer m.mu.RUnlock()
	return m.token
}

func (m *TokenManager) Update(token string) {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.token = token
}
