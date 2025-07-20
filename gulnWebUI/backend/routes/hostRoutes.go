package routes

import (
	"gulnManagement/gulnWebUI/handlers"
	"gulnManagement/gulnWebUI/middlewares"

	"github.com/gorilla/mux"
)

func RegisterHostRoutes(router *mux.Router) {
	subRoute := router.PathPrefix("/hosts").Subrouter()
	subRoute.Use(middlewares.AuthenticateJWT)

	subRoute.HandleFunc("/", handlers.GetHosts).Methods("GET")
	subRoute.HandleFunc("/{id}", handlers.GetHostByID).Methods("GET")
}
