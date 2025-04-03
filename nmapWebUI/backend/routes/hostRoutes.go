package routes

import (
	"nmapManagement/nmapWebUI/handlers"

	"github.com/gorilla/mux"
)

func RegisterHostRoutes(router *mux.Router) {
	router.HandleFunc("/hosts", handlers.GetHosts).Methods("GET")
	router.HandleFunc("/hosts/{id}", handlers.GetHostByID).Methods("GET")
}
