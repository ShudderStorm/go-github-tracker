package oauth

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"sync"
	"time"
)

const DefaultStateTTL time.Duration = 5 * time.Minute

type StateValidator interface {
	Store(string, time.Duration) error
	Validate(string) (bool, error)
}

type StateHolder struct {
	stateExpiration map[string]time.Time
	mu              sync.RWMutex
}

func DefaultStateHolder() *StateHolder {
	return &StateHolder{
		stateExpiration: make(map[string]time.Time),
	}
}

func (h *StateHolder) Store(state string, ttl time.Duration) error {
	h.mu.Lock()
	defer h.mu.Unlock()

	h.stateExpiration[state] = time.Now().Add(ttl)
	return nil
}

func (h *StateHolder) Validate(state string) (bool, error) {
	h.mu.RLock()
	defer h.mu.RUnlock()

	expiration, exists := h.stateExpiration[state]
	if !exists || !expiration.After(time.Now()) {
		return false, nil
	}

	return true, nil
}

func generateState() (string, error) {
	bytes := make([]byte, 24)
	_, err := rand.Read(bytes)
	if err != nil {
		return "", fmt.Errorf("state generating error: %w", err)
	}

	return base64.URLEncoding.EncodeToString(bytes), nil
}
