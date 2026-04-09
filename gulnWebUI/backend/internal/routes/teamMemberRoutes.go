package routes

import (
	handlers "gulnManagement/gulnWebUI/internal/handler"
	"gulnManagement/gulnWebUI/internal/middlewares"

	"github.com/gorilla/mux"
)

func RegisterTeamRoutes(router *mux.Router, teamMemberHandler *handlers.TeamMemberHandler) {
	subRoute := router.PathPrefix("/team-member").Subrouter()
	subRoute.Use(middlewares.AuthenticateJWT)

	subRoute.HandleFunc("/count", teamMemberHandler.GetTeamMemberCount).Methods("GET")
	subRoute.HandleFunc("/new", teamMemberHandler.AddNewTeamMember).Methods("POST")
	subRoute.HandleFunc("/list/{page}", teamMemberHandler.GetTeamMemberList).Methods("GET")
	//subRoute.HandleFunc("/info/{projectUUID}", projectHandler.GetProjectInfo).Methods("GET")
	//subRoute.HandleFunc("/upload/{projectUUID}", projectHandler.UploadProjectScan).Methods("POST")
}
