package auth

import "time"

// --------- Responses ---------

type TokenResponse struct {
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"`
	ExpiresIn   int64  `json:"expires_in"`
}

func NewTokenResponse(token Token) TokenResponse {
	return TokenResponse{
		AccessToken: token.AccessToken,
		TokenType:   "Bearer",
		ExpiresIn:   token.ExpiresIn,
	}
}

type TokenIntrospectionResponse struct {
	Active    bool   `json:"active"`
	ClientID  string `json:"client_id,omitempty"`
	TokenType string `json:"token_type,omitempty"`
	IssuedAt  int64  `json:"iat,omitempty"`
	ExpiresAt int64  `json:"exp,omitempty"`
}

func NewTokenIntrospectionResponse(ti TokenIntrospection) TokenIntrospectionResponse {
	return TokenIntrospectionResponse{
		Active:    ti.Active,
		ClientID:  ti.ClientID,
		TokenType: ti.TokenType,
		IssuedAt:  ti.IssuedAt,
		ExpiresAt: ti.ExpiresAt,
	}
}

type ErrorResponse struct {
	Message   string `json:"message"`
	Timestamp string `json:"timestamp"`
}

func NewErrorResponse(message string) ErrorResponse {
	return ErrorResponse{
		Message:   message,
		Timestamp: time.Now().Format("2006-01-02T15:04:05.000-07"),
	}
}
