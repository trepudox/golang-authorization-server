package auth

import "errors"

var (
	ErrInvalidCredentials = errors.New("credentials are not valid")
)