package routes

import (
	"nmapManagement/nmapWebUI/handlers"

	"github.com/gorilla/mux"
)

func RegisterScanRoutes(router *mux.Router) {
	router.HandleFunc("/scans", handlers.GetScans).Methods("GET")
	router.HandleFunc("/scans/{id}", handlers.GetScanByID).Methods("GET")
	router.HandleFunc("/scans/upload", handlers.UploadScan).Methods("POST")
}
