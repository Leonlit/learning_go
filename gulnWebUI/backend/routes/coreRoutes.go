package routes

import (
	handlers "gulnManagement/gulnWebUI/handlers/core"
	"gulnManagement/gulnWebUI/middlewares"
	scanner "gulnManagement/gulnWebUI/scanner"
	"net/http"

	"github.com/gorilla/mux"
)

func RegisterCoreRoutes(router *mux.Router) {
	subRoute := router.PathPrefix("/core").Subrouter()
	subRoute.Use(middlewares.AuthenticateJWT)

	subRoute.HandleFunc("/projects-hosts-counts", handlers.GetUserProjectAndHostCounts).Methods("GET")
	subRoute.HandleFunc("/test-vuln-scan", func(w http.ResponseWriter, r *http.Request) {
		scanner.GetPortVulns("af92c60c-4a28-44ab-9f72-35a2c7b6801f", "ba5ea793-fb2c-4c0a-aa0e-fd4926434364")
	}).Methods("GET")
}
