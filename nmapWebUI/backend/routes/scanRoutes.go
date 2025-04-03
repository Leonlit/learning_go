package routes

import (
	"nmapManagement/nmapWebUI/handlers"
	"nmapManagement/nmapWebUI/middlewares"

	"github.com/gorilla/mux"
)

func RegisterScanRoutes(router *mux.Router) {
	subRoute := router.PathPrefix("/scans").Subrouter()
	subRoute.Use(middlewares.AuthenticateJWT)

	subRoute.HandleFunc("/", handlers.GetScans).Methods("GET")
	subRoute.HandleFunc("/{id}", handlers.GetScanByID).Methods("GET")
	subRoute.HandleFunc("/upload", handlers.UploadScan).Methods("POST")
}
