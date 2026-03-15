package routes

import (
	handlers "gulnManagement/gulnWebUI/internal/handler"
	"gulnManagement/gulnWebUI/internal/middlewares"

	"github.com/gorilla/mux"
)

func RegisterAssessmentRoutes(router *mux.Router, AssessmentHandler *handlers.AssessmentHandler) {
	subRoute := router.PathPrefix("/assessments").Subrouter()
	subRoute.Use(middlewares.AuthenticateJWT)

	subRoute.HandleFunc("/count", AssessmentHandler.GetAssessmentCount).Methods("GET")
}
