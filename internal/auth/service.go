package auth

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
)

const defaultTokenTtl int = 600

type TokenService struct {
	repository *TokenRepository
	cache      *TokenCache
}

func NewTokenService(repository *TokenRepository, cache *TokenCache) *TokenService {
	return &TokenService{
		repository: repository,
		cache:      cache,
	}
}

func (ts *TokenService) createToken() (Token, error) {
	accessToken, err := ts.generateOpaqueToken()
	if err != nil {
		return Token{}, fmt.Errorf("it was not possible to generate the token: %s", err)
	}

	return Token{
		AccessToken: accessToken,
		ExpiresIn:   defaultTokenTtl,
	}, nil
}

func (ts *TokenService) generateOpaqueToken() (string, error) {
	bytes := make([]byte, 32)
	_, err := rand.Read(bytes)
	if err != nil {
		return "", err
	}

	return hex.EncodeToString(bytes), nil
}

func (ts *TokenService) areCredentialsValid(clientId, clientSecret string) (bool, error) {
	credentials, ok, err := ts.repository.GetClientCredentialsByClientId(clientId)
	if err != nil {
		return false, err
	}

	if !ok {
		return false, nil
	}

	return credentials.ClientSecret == clientSecret, nil
}

func (ts *TokenService) GenerateToken(clientId, clientSecret string) (Token, error) {
	validCredentials, err := ts.areCredentialsValid(clientId, clientSecret)
	if err != nil {
		return Token{}, fmt.Errorf("it was not possible to validate the credentials: %s", err)
	}

	if !validCredentials {
		return Token{}, ErrInvalidCredentials
	}

	currToken, ok, err := ts.cache.Get(clientId)
	if err != nil {
		return Token{}, err
	}

	if ok {
		return currToken, nil
	}

	token, err := ts.createToken()
	if err != nil {
		return Token{}, err
	}

	if err = ts.cache.Put(clientId, token); err != nil {
		return Token{}, err
	}

	return token, nil
}
