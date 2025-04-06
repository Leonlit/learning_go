package handlers

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"nmapManagement/nmapWebUI/databases"
	"nmapManagement/nmapWebUI/utils"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type Claims struct {
	Username string `json:"username"`
	jwt.RegisteredClaims
}

var secretKey = utils.LoadEnv("JWT_SECRET_KEY")

// Handle login operation
func LoginHandler(w http.ResponseWriter, r *http.Request) {
	var user map[string]string

	// Parse the login credentials from the request body
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	username := user["username"]
	password := user["password"]

	fmt.Print("Loging in user")

	// Validate user credentials (e.g., check against a database)
	if !databases.VerifyUserCredentials(username, password) {
		http.Error(w, "Invalid credentials", http.StatusUnauthorized)
		return
	}

	// Generate JWT token
	token, err := GenerateJWT(username)
	if err != nil {
		http.Error(w, "Error generating token", http.StatusInternalServerError)
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:     "auth_token",
		Value:    token,
		HttpOnly: true,
		Secure:   false, // Only for HTTPS //TODO: Turn on cookies Secure flag during prods
		Path:     "/",
		SameSite: http.SameSiteStrictMode,
	})
}

func RegisterUserHandler(w http.ResponseWriter, r *http.Request) {
	// register user
}

func GenerateJWT(username string) (string, error) {
	claims := Claims{
		Username: username,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    "nmap-management",
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(8 * time.Hour)), // Expiry set to 8 hour
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(secretKey)
}

func ParseJWT(tokenString string) (*Claims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		// Make sure the token method is what we expect (HS256)
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return secretKey, nil
	})

	if err != nil {
		return nil, err
	}

	// If token is valid, return the claims
	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		return claims, nil
	}

	return nil, errors.New("invalid token")
}
