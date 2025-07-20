package routes

import (
	"gulnManagement/gulnWebUI/handlers"
	"gulnManagement/gulnWebUI/middlewares"

	"github.com/gorilla/mux"
)

func RegisterAuthRoutes(router *mux.Router) {
	router.HandleFunc("/login", handlers.LoginHandler).Methods("POST")
	router.HandleFunc("/register", handlers.RegisterUserHandler).Methods("POST")
	router.HandleFunc("/logout", handlers.HandleLogout).Methods("POST")

	subRoute := router.PathPrefix("/auth").Subrouter()
	subRoute.Use(middlewares.AuthenticateJWT)
	subRoute.HandleFunc("/me", handlers.AuthMe).Methods("GET")
}
