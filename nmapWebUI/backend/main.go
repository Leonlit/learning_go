package main

import (
	"fmt"
	"log"
	"net/http"
	"nmapManagement/nmapWebUI/databases"
	"nmapManagement/nmapWebUI/routes"
	"nmapManagement/nmapWebUI/utils"
	"strconv"

	"github.com/gorilla/mux"
)

func enableCORS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "http://localhost:5173") // Vite uses 5173
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}
		next.ServeHTTP(w, r)
	})
}

func main() {
	dbPort, err := strconv.Atoi(utils.LoadEnv("DB_PORT"))
	if err != nil {
		log.Fatalf("Invalid DB port number: %v", err)
	}

	dbConfig := databases.DBConfig{
		Host:     utils.LoadEnv("DB_HOST"),
		Port:     dbPort,
		User:     utils.LoadEnv("DB_USER"),
		Password: utils.LoadEnv("DB_PASS"),
		DBName:   utils.LoadEnv("DB_NAME"),
		SSLMode:  "disable", // Use "require" for production
	}

	db, err := databases.NewDB(dbConfig)
	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
	}
	defer db.Close()

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
