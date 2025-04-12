package utils

import (
	"encoding/json"
	"log"
	"net/http"
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

func GetJWTFromCookie(w http.ResponseWriter, r *http.Request) string {
	cookie, err := r.Cookie("auth_token")
	if err != nil {
		http.Error(w, "auth_token cookie missing", http.StatusUnauthorized)
		return ""
	}

	return cookie.Value
}
