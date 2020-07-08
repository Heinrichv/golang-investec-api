package middleware

import (
	"net/http"
	"strings"

	services "../services"
)

// OAuthMiddleware - Verifies that auth token is valid
func OAuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		scheme := GetAuthScheme(r)

		if scheme == "" {
			next.ServeHTTP(w, r)
		} else {
			err := services.TokenValid(r)

			if err == nil {
				next.ServeHTTP(w, r)
			} else {
				http.Error(w, "Unauthorized: Invalid Token Received", http.StatusUnauthorized)
			}
		}
	})
}

func GetAuthScheme(r *http.Request) string {
	bearToken := r.Header.Get("Authorization")

	if bearToken == "" {
		return ""
	}

	strArr := strings.Split(bearToken, " ")

	if len(strArr) == 2 {
		return strArr[0]
	}

	return ""
}
