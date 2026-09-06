package routes

import (
	handlers "gulnManagement/gulnWebUI/internal/handler"
	"gulnManagement/gulnWebUI/internal/middlewares"

	"github.com/gorilla/mux"
)

func RegisterStatsRoutes(
	router *mux.Router,
	statsHandler *handlers.StatsHandler,
) {
	subRoute := router.PathPrefix("/stats").Subrouter()
	subRoute.Use(middlewares.AuthenticateJWT)

	subRoute.HandleFunc("/assessments/count", statsHandler.GetOwnAssessmentsCount).Methods("GET")
}
