package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/trepudox/golang-client-credentials-server/internal/auth"
)

const (
	port = "8080"
)

func main() {
	mux := http.NewServeMux()

	cache := auth.NewTokenCache()
	repository := auth.NewTokenRepository()
	service := auth.NewTokenService(repository, cache)
	handler := auth.NewTokenHandler(service)

	mux.HandleFunc("POST /token/generate", handler.GenerateTokenHandler)
	mux.HandleFunc("POST /token/refresh", handler.RefreshTokenHandler)
	mux.HandleFunc("POST /token/introspect", handler.IntrospectTokenHandler)

	log.Printf("golang-authorization-server listening to the %s port", port)
	http.ListenAndServe(fmt.Sprintf("localhost:%s", port), mux)
}
