package main

import (
	"gulnManagement/gulnWebUI/internal/handler"
	"gulnManagement/gulnWebUI/internal/logs"
	"gulnManagement/gulnWebUI/internal/repository"
	"gulnManagement/gulnWebUI/internal/routes"
	"gulnManagement/gulnWebUI/internal/service"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func enableCORS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "http://localhost:5173")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Cookie")
		w.Header().Set("Access-Control-Allow-Credentials", "true")

		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func main() {
	logs.InitLogs()
	repository.InitDB()

	router := mux.NewRouter()
	corsRouter := enableCORS(router)

	// --- build dependencies ---
	userRepo := repository.NewUserRepository(repository.DBObj)
	projectRepo := repository.NewProjectRepository(repository.DBObj)
	assessmentRepo := repository.NewAssessmentRepository(repository.DBObj)
	teamMemberRepo := repository.NewTeamMemberRepository(repository.DBObj)

	authService := service.NewAuthService(userRepo)
	authHandler := handler.NewAuthHandler(authService)

	projectService := service.NewProjectService(projectRepo)
	projectHandler := handler.NewProjectHandler(projectService)

	assessmentService := service.NewAssessmentService(assessmentRepo)
	assessmentHandler := handler.NewAssessmentHandler(assessmentService)

	teamMemberService := service.NewTeamMemberService(teamMemberRepo)
	teamMemberHandler := handler.NewTeamMemberHandler(teamMemberService)

	// API prefix
	apiRouter := router.PathPrefix("/api").Subrouter()

	// --- register routes ---
	routes.RegisterAuthRoutes(apiRouter, authHandler)
	routes.RegisterProjectRoutes(apiRouter, projectHandler)
	routes.RegisterAssessmentRoutes(apiRouter, assessmentHandler)
	routes.RegisterTeamRoutes(apiRouter, teamMemberHandler)

	log.Println("Server running on: http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", corsRouter))
}
