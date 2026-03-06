package utils

import (
	"encoding/json"
	"log"
	"net"
	"net/http"
)

type ErrorResponse struct {
	Error   string `json:"error"`
	Message string `json:"message"`
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

func SendJSONResponse(w http.ResponseWriter, payload interface{}, status int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	if err := json.NewEncoder(w).Encode(payload); err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
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
