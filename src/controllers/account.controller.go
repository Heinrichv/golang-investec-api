package controllers

import (
	"encoding/json"
	"net/http"

	mux "github.com/gorilla/mux"
)

// GetAccounts - Get a list of the account holders accounts
func GetAccounts(w http.ResponseWriter, r *http.Request) {

	json.NewEncoder(w).Encode(nil)
}

// InitiateAccountRoutes - Initialize routes for account controller
func InitiateAccountRoutes(GlobalRouter *mux.Router) *mux.Router {
	GlobalRouter.HandleFunc("/api/v1/Accounts", GetAutenticationToken).Methods("GET")

	return GlobalRouter
}
