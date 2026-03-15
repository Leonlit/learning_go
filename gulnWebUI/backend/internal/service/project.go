package service

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"gulnManagement/gulnWebUI/internal/repository"
	"gulnManagement/gulnWebUI/internal/utils"
	"net/http"
	"path/filepath"

	"github.com/gorilla/mux"
)

type ProjectService struct {
	projectRepo *repository.ProjectRepository
}

func NewProjectService(projectRepo *repository.ProjectRepository) *ProjectService {
	return &ProjectService{projectRepo: projectRepo}
}

func (s *ProjectService) GetProjectCount(ctx context.Context, userUUID string) (int, *utils.AppError) {
	projectCounts, err := s.projectRepo.GetProjectCount(ctx, userUUID)
	if err != nil {
		return -1, utils.InternalError("Error fetching project count", err)
	}

	return projectCounts, nil
}

func (s *ProjectService) CreateNewProject(ctx context.Context, userUUID, projectName string) (string, *utils.AppError) {

	projectUUID, err := s.projectRepo.CreateNewProject(ctx, userUUID, projectName)
	if err != nil {
		return "", utils.InternalError("Error creating new project", err)
	}

	return projectUUID, nil
}

func (s *ProjectService) GetProjectsList(ctx context.Context, userUUID string, page int) ([]repository.Project, *utils.AppError) {

	projects, err := s.projectRepo.GetProjectList(ctx, userUUID, page)
	if err != nil {
		return []repository.Project{}, utils.InternalError("Error fetching project list", err)
	}
	return projects, nil
}

func (s *ProjectService) GetProjectInfo(ctx context.Context, userUUID, projectUUID string) (repository.Project, *utils.AppError) {

	project, err := s.projectRepo.GetProjectInfo(ctx, userUUID, projectUUID)

	if err != nil {

		if errors.Is(err, sql.ErrNoRows) {
			return repository.Project{}, utils.BadRequest(
				"PROJECT_NOT_FOUND",
				"Project not found",
			)
		}

		return repository.Project{}, utils.InternalError(
			"Error fetching project info",
			err,
		)
	}

	return project, nil
}

func UploadProjectScan(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	projectUUID := vars["projectUUID"]

	fmt.Println(projectUUID)

	scanName := r.FormValue("scanName")
	fmt.Println(scanName)
	if scanName == "" {
		http.Error(w, "Project name is required", http.StatusBadRequest)
		return
	}

	r.Body = http.MaxBytesReader(w, r.Body, 10<<20) // 10 MB max, 10,485,760 bytes

	// Parse multipart form
	if err := r.ParseMultipartForm(10 << 20); err != nil {
		http.Error(w, "File too big or bad request", http.StatusBadRequest)
		return
	}

	// Get uploaded file
	file, handler, err := r.FormFile("file")
	if err != nil {
		http.Error(w, "Error retrieving file", http.StatusBadRequest)
		return
	}
	defer file.Close()

	// Validate file type, XML
	ext := filepath.Ext(handler.Filename)
	if ext != ".xml" {
		http.Error(w, "Invalid file type", http.StatusBadRequest)
		return
	}

	/* nmapRun, err := nmapParser.ParseNmap(file)

	if err != nil {
		http.Error(w, "Error Parsing Nmap results", http.StatusBadRequest)
		return
	}

	savedScanUUID := databases.SaveScanResultsToDatabase(projectUUID, scanName, nmapRun)
	scanner.GetPortVulns(projectUUID, savedScanUUID)
	if savedScanUUID == "" {
		http.Error(w, "Error Saving Nmap results", http.StatusInternalServerError)
		return
	} */

	fmt.Fprintf(w, "File uploaded successfully: %s", file)
}
