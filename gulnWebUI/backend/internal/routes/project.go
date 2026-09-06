package routes

import (
	handlers "gulnManagement/gulnWebUI/internal/handler"
	"gulnManagement/gulnWebUI/internal/middlewares"

	"github.com/gorilla/mux"
)

func RegisterProjectRoutes(
	router *mux.Router,
	projectHandler *handlers.ProjectHandler,
	assessmentHandler *handlers.AssessmentHandler,
) {
	subRoute := router.PathPrefix("/projects").Subrouter()
	subRoute.Use(middlewares.AuthenticateJWT)

	subRoute.HandleFunc("/count", projectHandler.GetProjectCount).Methods("GET")
	subRoute.HandleFunc("/new", projectHandler.CreateNewProject).Methods("POST")
	subRoute.HandleFunc("/list/{page}", projectHandler.GetProjectsList).Methods("GET")
	subRoute.HandleFunc("/{projectUUID}/info", projectHandler.GetProjectInfo).Methods("GET")
	//subRoute.HandleFunc("/upload/{projectUUID}", projectHandler.UploadProjectScan).Methods("POST")

	assessmentRoute := subRoute.PathPrefix("/{projectUUID}/assessments").Subrouter()
	assessmentRoute.HandleFunc("/count", assessmentHandler.GetAssessmentCount).Methods("GET")
}
