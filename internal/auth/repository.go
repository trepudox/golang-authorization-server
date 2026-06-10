package auth

import "sync"

// "Banco de dados" das client_credentials
//
//	clientId - clientSecret

type TokenRepository struct {
	mu             sync.Mutex
	allowedClients map[string]ClientCredentials
}

func NewTokenRepository() *TokenRepository {
	return &TokenRepository{
		allowedClients: map[string]ClientCredentials{
			"testId":  {ClientId: "testId", ClientSecret: "testSecret"},
			"testId2": {ClientId: "testSecret2", ClientSecret: "testSecret2"},
		},
	}
}

func (tr *TokenRepository) GetClientCredentialsByClientId(clientId string) (ClientCredentials, bool, error) {
	tr.mu.Lock()
	defer tr.mu.Unlock()

	value, ok := tr.allowedClients[clientId]
	return value, ok, nil
}
