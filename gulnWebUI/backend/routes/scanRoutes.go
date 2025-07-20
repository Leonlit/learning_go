package routes

import (
	"gulnManagement/gulnWebUI/handlers"
	"gulnManagement/gulnWebUI/middlewares"

	"github.com/gorilla/mux"
)

func RegisterScanRoutes(router *mux.Router) {
	subRoute := router.PathPrefix("/scans").Subrouter()
	subRoute.Use(middlewares.AuthenticateJWT)

	subRoute.HandleFunc("/list/{page}", handlers.GetScansList).Methods("GET")
	subRoute.HandleFunc("/{id}", handlers.GetScanByID).Methods("GET")
	subRoute.HandleFunc("/upload", handlers.UploadScan).Methods("POST")
}
