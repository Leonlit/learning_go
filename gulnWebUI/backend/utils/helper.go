package utils

import (
	"encoding/json"
	"log"
	"net"
	"net/http"
)

type Response struct {
	Message string `json:"message"`
	Status  int    `json:"status"`
}

func CheckAddressValid(addr string) bool {
	// Try parsing as an IP address first
	ip := net.ParseIP(addr)
	if ip != nil {
		log.Printf("'%s' is a valid IP Address.\n", addr)
		if ip.To4() != nil {
			log.Println("  - It's an IPv4 address.")
			return true
		} else {
			log.Println("  - It's an IPv6 address.")
			return true
		}
	}

	// If not an IP, try looking it up as a hostname
	hosts, err := net.LookupHost(addr)
	if err == nil && len(hosts) > 0 {
		log.Printf("'%s' is a valid Hostname (resolves to: %v).\n", addr, hosts)
		return true
	}

	// If both fail, it's neither
	log.Printf("'%s' is neither a valid IP address nor a resolvable hostname.\n", addr)
	return false
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
