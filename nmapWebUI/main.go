package main

import (
	"fmt"
	"log"
	"net/http"
	"nmapManagement/nmapWebUI/routes"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "/login", http.StatusFound) // 302 Redirect
	}).Methods("GET")

	// Register routes
	routes.RegisterAuthRoutes(router)
	routes.RegisterScanRoutes(router)
	routes.RegisterHostRoutes(router)

	fmt.Println("Using port 8080")
	fmt.Println("Server running on: http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
