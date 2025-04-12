package middlewares

import (
	"context"
	"log"
	"net/http"
	"nmapManagement/nmapWebUI/handlers"
	"nmapManagement/nmapWebUI/utils"
)

// Middleware function to protect routes that require a valid JWT
func AuthenticateJWT(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Get token from the Authorization header

		tokenString := utils.GetJWTFromCookie(w, r)
		log.Println(tokenString)

		if tokenString == "" {
			return
		}
		// Parse and validate the JWT token
		claims, err := handlers.ParseJWT(tokenString)
		if err != nil {
			http.Error(w, "Invalid token", http.StatusUnauthorized)
			return
		}

		ctx := context.WithValue(r.Context(), "UserUUID", claims.UserUUID)
		r = r.WithContext(ctx)

		// Proceed with the handler
		next.ServeHTTP(w, r)
	})
}
