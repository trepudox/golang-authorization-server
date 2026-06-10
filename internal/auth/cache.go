package auth

import "sync"

// "Cache" dos tokens atuais
//
//	clientId - token

type TokenCache struct {
	mu              sync.Mutex
	currentSessions map[string]Token
}

func NewTokenCache() *TokenCache {
	return &TokenCache{
		currentSessions: map[string]Token{},
	}
}

func (tc *TokenCache) Get(key string) (Token, bool, error) {
	tc.mu.Lock()
	defer tc.mu.Unlock()

	val, ok := tc.currentSessions[key]
	return val, ok, nil
}

func (tc *TokenCache) Put(key string, value Token) error {
	tc.mu.Lock()
	defer tc.mu.Unlock()

	tc.currentSessions[key] = value
	return nil
}
