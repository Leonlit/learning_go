package routes

import (
	"gulnManagement/gulnWebUI/internal/handler"
	"gulnManagement/gulnWebUI/internal/middlewares"

	"github.com/gorilla/mux"
)

func RegisterAuthRoutes(
	router *mux.Router,
	authHandler *handler.AuthHandler,
) {
	router.HandleFunc("/login", authHandler.LoginAuthHandler).Methods("POST")
	router.HandleFunc("/register", authHandler.RegisterUserHandler).Methods("POST")
	router.HandleFunc("/logout", authHandler.LogoutHandler).Methods("POST")

	subRoute := router.PathPrefix("/auth").Subrouter()
	subRoute.Use(middlewares.AuthenticateJWT)
	subRoute.HandleFunc("/me", authHandler.AuthMe).Methods("GET")
}
