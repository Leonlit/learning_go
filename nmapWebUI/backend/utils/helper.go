package utils

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"
)

type Response struct {
	Message string `json:"message"`
	Status  int    `json:"status"`
}

func SendJSONResponse(w http.ResponseWriter, message string, status int) {
	// Create the response object
	response := Response{
		Message: message,
		Status:  status,
	}

	// Set the response header to indicate that the content is JSON
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	// Marshal the response object into JSON
	err := json.NewEncoder(w).Encode(response)
	if err != nil {
		log.Printf("Error encoding response: %v", err)
		http.Error(w, "Unable to encode response", http.StatusInternalServerError)
		return
	}
}

func GetJWTAuthHeaderString(w http.ResponseWriter, r *http.Request) string {
	authHeader := r.Header.Get("Authorization")
	if authHeader == "" {
		http.Error(w, "Authorization header missing", http.StatusUnauthorized)
		return ""
	}

	// Token is passed as "Bearer <token>"
	tokenString := strings.TrimPrefix(authHeader, "Bearer ")
	return tokenString
}
