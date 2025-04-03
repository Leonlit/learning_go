package routes

import (
	"nmapManagement/nmapWebUI/handlers"

	"github.com/gorilla/mux"
)

func RegisterAuthRoutes(router *mux.Router) {
	router.HandleFunc("/login", handlers.GetLoginPage).Methods("GET")
	router.HandleFunc("/register", handlers.GetRegistraterPage).Methods("GET")
	router.HandleFunc("/login", handlers.VerifyUserLogin).Methods("POST")
	router.HandleFunc("/register", handlers.RegisterUser).Methods("POST")
}
