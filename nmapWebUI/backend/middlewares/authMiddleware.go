package middlewares

import (
	"context"
	"net/http"
	"nmapManagement/nmapWebUI/handlers"
	"strings"
)

// Middleware function to protect routes that require a valid JWT
func AuthenticateJWT(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Get token from the Authorization header
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			http.Error(w, "Authorization header missing", http.StatusUnauthorized)
			return
		}

		// Token is passed as "Bearer <token>"
		tokenString := strings.TrimPrefix(authHeader, "Bearer ")

		// Parse and validate the JWT token
		claims, err := handlers.ParseJWT(tokenString)
		if err != nil {
			http.Error(w, "Invalid token", http.StatusUnauthorized)
			return
		}

		ctx := context.WithValue(r.Context(), "username", claims.Username)
		r = r.WithContext(ctx)

		// Proceed with the handler
		next.ServeHTTP(w, r)
	})
}
