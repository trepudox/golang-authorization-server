package auth

import "sync"

// TODO: Logica pra limpar os tokens do cache

// "Cache" dos tokens atuais
//
//	normal - clientId: token
//  reverse - tokenValue: clientId

type TokenCache struct {
	mu                     sync.Mutex
	currentSessions        map[string]Token
	reverseCurrentSessions map[string]string
}

func NewTokenCache() *TokenCache {
	return &TokenCache{
		currentSessions:        map[string]Token{},
		reverseCurrentSessions: map[string]string{},
	}
}

func (tc *TokenCache) Get(key string) (Token, bool, error) {
	tc.mu.Lock()
	defer tc.mu.Unlock()

	val, ok := tc.currentSessions[key]
	return val, ok, nil
}

func (tc *TokenCache) GetReverse(key string) (string, bool, error) {
	tc.mu.Lock()
	defer tc.mu.Unlock()

	val, ok := tc.reverseCurrentSessions[key]
	return val, ok, nil
}

func (tc *TokenCache) Put(key string, value Token) error {
	tc.mu.Lock()
	defer tc.mu.Unlock()

	tc.currentSessions[key] = value
	tc.reverseCurrentSessions[value.AccessToken] = key
	return nil
}
