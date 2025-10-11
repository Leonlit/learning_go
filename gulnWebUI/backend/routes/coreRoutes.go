package routes

import (
	handlers "gulnManagement/gulnWebUI/handlers/core"
	"gulnManagement/gulnWebUI/middlewares"

	"github.com/gorilla/mux"
)

func RegisterCoreRoutes(router *mux.Router) {
	subRoute := router.PathPrefix("/core").Subrouter()
	subRoute.Use(middlewares.AuthenticateJWT)

	subRoute.HandleFunc("/projects-hosts-counts", handlers.GetUserProjectAndHostCounts).Methods("GET")
}
