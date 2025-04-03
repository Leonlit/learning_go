package routes

import (
	"nmapManagement/nmapWebUI/handlers"
	"nmapManagement/nmapWebUI/middlewares"

	"github.com/gorilla/mux"
)

func RegisterHostRoutes(router *mux.Router) {
	subRoute := router.PathPrefix("/hosts").Subrouter()
	subRoute.Use(middlewares.AuthenticateJWT)

	subRoute.HandleFunc("/", handlers.GetHosts).Methods("GET")
	subRoute.HandleFunc("/{id}", handlers.GetHostByID).Methods("GET")
}
