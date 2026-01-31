package routes

import (
	handlers "gulnManagement/gulnWebUI/internal/handler/project"
	"gulnManagement/gulnWebUI/internal/middlewares"

	"github.com/gorilla/mux"
)

func RegisterProjectRoutes(router *mux.Router) {
	subRoute := router.PathPrefix("/projects").Subrouter()
	subRoute.Use(middlewares.AuthenticateJWT)

	subRoute.HandleFunc("/new", handlers.CreateNewProjects).Methods("POST")
	subRoute.HandleFunc("/list/{page}", handlers.GetProjectsList).Methods("GET")
	subRoute.HandleFunc("/info/{projectUUID}", handlers.GetProjectInfo).Methods("GET")
	subRoute.HandleFunc("/upload/{projectUUID}", handlers.UploadProjectScan).Methods("POST")
}
