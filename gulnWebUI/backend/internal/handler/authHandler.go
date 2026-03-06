package handler

import (
	"encoding/json"
	"errors"
	"gulnManagement/gulnWebUI/internal/auth"
	"gulnManagement/gulnWebUI/internal/repository"
	service "gulnManagement/gulnWebUI/internal/service"
	"gulnManagement/gulnWebUI/internal/utils"
	"log"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type AuthHandler struct {
	authService *service.AuthService
}

func NewAuthHandler(authService *service.AuthService) *AuthHandler {
	return &AuthHandler{authService: authService}
}

func (h *AuthHandler) LoginAuthHandler(w http.ResponseWriter, r *http.Request) {
	var req struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	token, err := h.authService.Login(r.Context(), req.Username, req.Password)
	if errors.Is(err, service.ErrInvalidCredentials) {
		http.Error(w, "Invalid credentials", http.StatusUnauthorized)
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:     "auth_token",
		Value:    token,
		HttpOnly: true,
		Secure:   false,
		Path:     "/",
		SameSite: http.SameSiteLaxMode,
	})

	utils.SendJSONResponse(w, "User valid", http.StatusOK)
}

func (h *AuthHandler) RegisterUserHandler(w http.ResponseWriter, r *http.Request) {

	var req struct {
		Username       string `json:"username"`
		Password       string `json:"password"`
		RepeatPassword string `json:"repeatPassword"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	err := h.authService.Register(r.Context(), req.Username, req.Password, req.RepeatPassword)
	if errors.Is(err, service.ErrInvalidCredentials) {
		http.Error(w, "Invalid credentials", http.StatusBadRequest)
		return
	}

	utils.SendJSONResponse(w, "User registered successfully", http.StatusCreated)
}

func ParseJWT(tokenString string) (*auth.Claims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &auth.Claims{}, func(token *jwt.Token) (interface{}, error) {
		// Make sure the token method is what we expect (HS256)
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return []byte(repository.JWTSecretKey), nil
	})

	if err != nil {
		log.Println("Error parsing JWT!")
		log.Println(err)
		return nil, err
	}

	// If token is valid, return the claims
	if claims, ok := token.Claims.(*auth.Claims); ok && token.Valid {
		return claims, nil
	}

	return nil, errors.New("invalid token")
}

func (h *AuthHandler) AuthMe(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

func (h *AuthHandler) LogoutHandler(w http.ResponseWriter, r *http.Request) {
	// Remove the JWT by setting the cookie with the same name and an expired date
	http.SetCookie(w, &http.Cookie{
		Name:     "auth_token",         // Same name as the cookie you set
		Value:    "",                   // Empty value
		Expires:  time.Unix(0, 0),      // Expire date in the past
		HttpOnly: true,                 // Make it HttpOnly for security
		Secure:   false,                // Set to true for HTTPS (secure flag)
		Path:     "/",                  // Same path as where it was originally set
		SameSite: http.SameSiteLaxMode, // Ensure it works cross-site if needed
	})

	// Optionally, you can also redirect or send a response back to the client
	http.Redirect(w, r, "/login", http.StatusFound)
}
