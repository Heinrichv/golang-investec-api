package controllers

import (
	"encoding/json"
	"log"
	"net/http"

	services "../services"
	mux "github.com/gorilla/mux"
)

// GetAutenticationToken - Get authentication token for account holder
func GetAutenticationToken(w http.ResponseWriter, r *http.Request) {
	client := r.Header.Get("X-Client-ID")
	secret := r.Header.Get("X-Client-Secret")

	token := services.GetAuthenticationToken(client, secret)

	log.Println(token)
	json.NewEncoder(w).Encode(token)
}

// InitiateAuthRoutes - Initialize routes for account controller
func InitiateAuthRoutes(GlobalRouter *mux.Router) *mux.Router {
	GlobalRouter.HandleFunc("/api/v1/Authenticate", GetAutenticationToken).Methods("GET")

	return GlobalRouter
}
