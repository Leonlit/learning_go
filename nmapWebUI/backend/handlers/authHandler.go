package handlers

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"nmapManagement/nmapWebUI/databases"
	"nmapManagement/nmapWebUI/utils"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type Claims struct {
	Username string `json:"username"`
	jwt.RegisteredClaims
}

var secretKey = utils.LoadEnv("JWT_SECRET_KEY")

func hashPassword(password string) (string, error) {
	// Generate a hashed password with bcrypt
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		log.Printf("failed to hash password: %v", err)
		return "", fmt.Errorf("failed to hash password: %v", err)
	}
	return string(hashedPassword), nil
}

// Handle login operation
func LoginHandler(w http.ResponseWriter, r *http.Request) {
	var user map[string]string

	// Parse the login credentials from the request body
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		log.Println("Invalid request body!")
		log.Println(err)
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	username := user["username"]
	password := user["password"]

	fmt.Println("Loging in user")

	// Validate user credentials (e.g., check against a database)
	if !databases.VerifyUserCredentials(username, password) {
		http.Error(w, "Invalid credentials", http.StatusUnauthorized)
		return
	}

	// Generate JWT token
	token, err := generateJWT(username)
	if err != nil {
		log.Println("Error generating token!")
		log.Println(err)
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
	utils.SendJSONResponse(w, "User valid", http.StatusOK)
}

func RegisterUserHandler(w http.ResponseWriter, r *http.Request) {
	var user map[string]string

	// Parse the login credentials from the request body
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		log.Println("Invalid request body!")
		log.Println(err)
		http.Error(w, "Invalid request body!", http.StatusBadRequest)
		return
	}

	username := user["username"]
	password := user["password"]
	repeatPassword := user["repeatPassword"]

	if repeatPassword != password {
		http.Error(w, "Different password used!", http.StatusBadRequest)
		return
	}

	userExists, err := databases.CheckUsernameExists(username)

	if err != nil {
		http.Error(w, "Unexpected Error!", http.StatusInternalServerError)
		return
	}

	if userExists {
		http.Error(w, "Invalid request body!", http.StatusBadRequest)
		return
	}

	passwordHash, err := hashPassword(password)

	if err != nil {
		http.Error(w, "Unexpected Error!", http.StatusInternalServerError)
		return
	}

	created, err := databases.CreateNewUser(username, passwordHash)

	if err != nil {
		http.Error(w, "Unexpected Error!", http.StatusInternalServerError)
		return
	}

	if created == 0 {
		log.Println("No DB entry created! Did not register new user!")
		http.Error(w, "Unexpected Error!", http.StatusInternalServerError)
		return
	}

	utils.SendJSONResponse(w, "User registered successfully", http.StatusCreated)
}

func generateJWT(username string) (string, error) {
	claims := Claims{
		Username: username,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    "nmap-management",
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(8 * time.Hour)), // Expiry set to 8 hour
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(secretKey))
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
		log.Println("Error parsing JWT!")
		log.Println(err)
		return nil, err
	}

	// If token is valid, return the claims
	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		return claims, nil
	}

	return nil, errors.New("invalid token")
}

func AuthMe(w http.ResponseWriter, r *http.Request) {
	// If we're here, middleware already validated the token
	w.WriteHeader(http.StatusOK)
}
