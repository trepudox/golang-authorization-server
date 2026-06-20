package auth

import (
	"encoding/json"
	"errors"
	"net/http"
)

type TokenHandler struct {
	service *TokenService
}

func NewTokenHandler(service *TokenService) *TokenHandler {
	return &TokenHandler{
		service: service,
	}
}

func (th *TokenHandler) GenerateToken(w http.ResponseWriter, r *http.Request) {
	if r.Header.Get("Content-Type") != "application/x-www-form-urlencoded" {
		sendResponse(w, http.StatusUnsupportedMediaType, nil)
		return
	}

	if err := r.ParseForm(); err != nil {
		sendResponse(w, 500, NewErrorResponse(err.Error()))
		return
	}

	clientId := r.PostForm.Get("client_id")
	clientSecret := r.PostForm.Get("client_secret")

	if clientId == "" || clientSecret == "" {
		sendResponse(w, 400, NewErrorResponse("the request is missing credentials"))
		return
	}

	tkn, err := th.service.GenerateToken(clientId, clientSecret)
	if err != nil {
		if errors.Is(err, ErrInvalidCredentials) {
			sendResponse(w, 401, nil)
		} else {
			sendResponse(w, 500, NewErrorResponse(err.Error()))
		}

		return
	}

	sendResponse(w, 200, NewTokenResponse(tkn))
}

func (th *TokenHandler) IntrospectToken(w http.ResponseWriter, r *http.Request) {
	clientId, clientSecret, ok := r.BasicAuth()
	if !ok || clientId == "" || clientSecret == "" {
		sendResponse(w, 401, nil)
		return
	}

	if r.Header.Get("Content-Type") != "application/x-www-form-urlencoded" {
		sendResponse(w, 415, nil)
		return
	}

	if err := r.ParseForm(); err != nil {
		sendResponse(w, 500, NewErrorResponse(err.Error()))
		return
	}

	sentToken := r.PostForm.Get("token")

	tknIntrospection, err := th.service.IntrospectToken(clientId, clientSecret, sentToken)
	if err != nil {
		if errors.Is(err, ErrInvalidCredentials) {
			sendResponse(w, 401, nil)
		} else {
			sendResponse(w, 500, NewErrorResponse(err.Error()))
		}

		return
	}

	if sentToken == "" {
		sendResponse(w, 400, NewErrorResponse("the request is missing the introspection token"))
		return
	}

	sendResponse(w, 200, NewTokenIntrospectionResponse(tknIntrospection))
}

func sendResponse(w http.ResponseWriter, status int, data any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	if data != nil {
		json.NewEncoder(w).Encode(data)
	}
}
