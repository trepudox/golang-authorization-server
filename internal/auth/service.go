package auth

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"time"
)

const defaultTokenTtl int64 = 600

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

func (ts *TokenService) createToken(clientId string) (Token, error) {
	accessToken, err := ts.generateOpaqueToken()
	if err != nil {
		return Token{}, fmt.Errorf("it was not possible to generate the token: %s", err)
	}

	issuedAt := time.Now().UTC().Unix()

	return Token{
		AccessToken: accessToken,
		ClientID:    clientId,
		IssuedAt:    issuedAt,
		ExpiresIn:   defaultTokenTtl,
		ExpiresAt:   issuedAt + defaultTokenTtl,
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
		// TODO: log err
		return Token{}, ErrInternalServerError
	}

	if !validCredentials {
		return Token{}, ErrInvalidCredentials
	}

	currToken, ok, err := ts.cache.Get(clientId)
	if err != nil {
		// TODO: log err
		return Token{}, ErrInternalServerError
	}

	if ok {
		return currToken, nil
	}

	token, err := ts.createToken(clientId)
	if err != nil {
		// TODO: log err
		return Token{}, ErrInternalServerError
	}

	if err = ts.cache.Put(clientId, token); err != nil {
		// TODO: log err
		return Token{}, ErrInternalServerError
	}

	return token, nil
}

func (ts *TokenService) IntrospectToken(clientId, clientSecret, token string) (TokenIntrospection, error) {
	validCredentials, err := ts.areCredentialsValid(clientId, clientSecret)
	if err != nil {
		// TODO: log err
		return TokenIntrospection{}, ErrInternalServerError
	}

	if !validCredentials {
		return TokenIntrospection{}, ErrInvalidCredentials
	}

	bearerTknClientId, ok, err := ts.cache.GetReverse(token)
	if err != nil {
		// TODO: log err
		return TokenIntrospection{}, ErrInternalServerError
	}

	if !ok {
		return NewInactiveTokenIntrospection(), nil
	}

	bearerTkn, ok, err := ts.cache.Get(bearerTknClientId)
	if err != nil {
		// TODO: log err
		return TokenIntrospection{}, ErrInternalServerError
	}

	if !ok {
		return NewInactiveTokenIntrospection(), nil
	}

	return NewActiveTokenIntrospection(bearerTkn.ClientID, bearerTkn.IssuedAt, bearerTkn.ExpiresAt), err
}
