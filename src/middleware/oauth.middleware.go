package middleware

import (
	"net/http"

	services "../services"
)

// OAuthMiddleware - Verifies that auth token is valid
func OAuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		err := services.TokenValid(r)

		if err == nil {
			next.ServeHTTP(w, r)
		} else {
			http.Error(w, "Unauthorized: Invalid_Token", http.StatusUnauthorized)
		}
	})
}
