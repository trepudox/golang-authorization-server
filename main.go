package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/trepudox/golang-authorization-server/internal/auth"
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

	mux.HandleFunc("POST /token/generate", handler.GenerateToken)
	mux.HandleFunc("POST /token/introspect", handler.IntrospectToken)

	log.Printf("golang-authorization-server listening to the %s port", port)
	http.ListenAndServe(fmt.Sprintf("localhost:%s", port), mux)
}
