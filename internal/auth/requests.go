package auth

import "time"

// --------- Responses ---------

type TokenResponse struct {
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"`
	ExpiresIn   int    `json:"expires_in"`
}

func NewTokenResponse(accessToken string, expiresIn int) TokenResponse {
	return TokenResponse{
		AccessToken: accessToken,
		TokenType:   "Bearer",
		ExpiresIn:   expiresIn,
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
