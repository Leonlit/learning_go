package routes

import (
	handlers "gulnManagement/gulnWebUI/handlers/project"
	"gulnManagement/gulnWebUI/middlewares"

	"github.com/gorilla/mux"
)

func RegisterProjectRoutes(router *mux.Router) {
	subRoute := router.PathPrefix("/projects").Subrouter()
	subRoute.Use(middlewares.AuthenticateJWT)

	subRoute.HandleFunc("/new", handlers.CreateNewProjects).Methods("POST")
	subRoute.HandleFunc("/list/{page}", handlers.GetProjectsList).Methods("GET")
	subRoute.HandleFunc("/info/{projectUUID}", handlers.GetProjectInfo).Methods("GET")
	subRoute.HandleFunc("/info/scans/{projectUUID}/{page}", handlers.GetProjectScan).Methods("GET")
	subRoute.HandleFunc("/scans/info/{projectUUID}/{scanUUID}/{page}", handlers.GetProjectScanInfo).Methods("GET")
	subRoute.HandleFunc("/scans/host/info/{projectUUID}/{scanUUID}/{hostUUID}", handlers.GetProjectScanHostInfo).Methods("GET")
	subRoute.HandleFunc("/upload/{projectUUID}", handlers.UploadProjectScan).Methods("POST")
}
