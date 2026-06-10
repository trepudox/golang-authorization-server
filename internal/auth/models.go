package auth

type Token struct {
	AccessToken string
	ExpiresIn   int
}

type ClientCredentials struct {
	ClientId     string
	ClientSecret string
}
