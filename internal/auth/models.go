package auth

type Token struct {
	AccessToken string
	ClientID    string
	IssuedAt    int64
	ExpiresIn   int64
	ExpiresAt   int64
}

type TokenIntrospection struct {
	Active    bool
	ClientID  string
	TokenType string
	IssuedAt  int64
	ExpiresAt int64
}

func NewActiveTokenIntrospection(clientID string, issuedAt int64, expiresAt int64) TokenIntrospection {
	return TokenIntrospection{
		Active:    true,
		ClientID:  clientID,
		TokenType: "Bearer",
		IssuedAt:  issuedAt,
		ExpiresAt: expiresAt,
	}
}

func NewInactiveTokenIntrospection() TokenIntrospection {
	return TokenIntrospection{
		Active:    false,
		ClientID:  "",
		TokenType: "",
		IssuedAt:  0,
		ExpiresAt: 0,
	}
}

type ClientCredentials struct {
	ClientId     string
	ClientSecret string
}
