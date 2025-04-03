package main

import (
	"fmt"
	"log"
	"net/http"
	"nmapManagement/nmapWebUI/routes"

	"github.com/gorilla/mux"
)

func enableCORS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "http://localhost:5173") // Vite uses 5173
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		next.ServeHTTP(w, r)
	})
}

func main() {
	router := mux.NewRouter()

	corsRouter := enableCORS(router)

	// Register routes
	routes.RegisterAuthRoutes(router)
	routes.RegisterScanRoutes(router)
	routes.RegisterHostRoutes(router)

	fmt.Println("Using port 8080")
	fmt.Println("Server running on: http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", corsRouter))
}
