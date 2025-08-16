package routes

import (
	handlers "gulnManagement/gulnWebUI/handlers/scan"
	"gulnManagement/gulnWebUI/middlewares"

	"github.com/gorilla/mux"
)

func RegisterScanRoutes(router *mux.Router) {
	subRoute := router.PathPrefix("/scans").Subrouter()
	subRoute.Use(middlewares.AuthenticateJWT)

	subRoute.HandleFunc("/list/{page}", handlers.GetScansList).Methods("GET")
	subRoute.HandleFunc("/{id}", handlers.GetScanByID).Methods("GET")
}
