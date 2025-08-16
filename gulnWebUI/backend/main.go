package main

import (
	"gulnManagement/gulnWebUI/databases"
	"gulnManagement/gulnWebUI/logs"
	"gulnManagement/gulnWebUI/routes"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func enableCORS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "http://localhost:5173")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Cookie")
		w.Header().Set("Access-Control-Allow-Credentials", "true")

		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func main() {
	logs.InitLogs()
	databases.InitDB()

	router := mux.NewRouter()

	corsRouter := enableCORS(router)

	// Register routes
	routes.RegisterAuthRoutes(router)
	routes.RegisterScanRoutes(router)
	routes.RegisterHostRoutes(router)
	routes.RegisterProjectRoutes(router)

	log.Println("Using port 8080")
	log.Println("Server running on: http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", corsRouter))
}
