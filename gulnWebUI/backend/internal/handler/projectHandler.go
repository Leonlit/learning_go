package handler

import (
	"encoding/json"
	"fmt"
	"gulnManagement/gulnWebUI/internal/service"
	"gulnManagement/gulnWebUI/internal/utils"
	"math"
	"net/http"
	"path/filepath"
	"strconv"

	"github.com/gorilla/mux"
)

type ProjectHandler struct {
	projectService *service.ProjectService
}

func NewProjectHandler(projectService *service.ProjectService) *ProjectHandler {
	return &ProjectHandler{projectService: projectService}
}

func (h *ProjectHandler) GetProjectCount(w http.ResponseWriter, r *http.Request) {

	ctx := r.Context()

	userUUID, ok := ctx.Value("UserUUID").(string)
	if !ok {
		utils.SendJSONResponse(w,
			utils.Unauthorized("Unauthorized", nil),
			http.StatusUnauthorized,
		)
		return
	}

	projectCount, err := h.projectService.GetProjectCount(ctx, userUUID)
	if err != nil {
		utils.SendJSONResponse(w, err, err.Status)
		return
	}

	utils.SendJSONResponse(w, map[string]int{
		"count": projectCount,
	}, http.StatusOK)
}

func (h *ProjectHandler) CreateNewProject(w http.ResponseWriter, r *http.Request) {

	ctx := r.Context()

	userUUID, ok := ctx.Value("UserUUID").(string)
	if !ok {
		utils.SendJSONResponse(w,
			utils.Unauthorized("Unauthorized", nil),
			http.StatusUnauthorized,
		)
		return
	}

	projectName := r.FormValue("projectName")
	if projectName == "" {
		utils.SendJSONResponse(w,
			utils.BadRequest("INVALID_PROJECT_NAME", "Project name is required"),
			http.StatusBadRequest,
		)
		return
	}

	projectsID, err := h.projectService.CreateNewProject(r.Context(), userUUID, projectName)
	if err != nil {
		utils.SendJSONResponse(w, err, err.Status)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"projectID": projectsID,
	})
}

func (h *ProjectHandler) GetProjectsList(w http.ResponseWriter, r *http.Request) {

	ctx := r.Context()

	userUUID, ok := ctx.Value("UserUUID").(string)
	if !ok {
		utils.SendJSONResponse(w,
			utils.Unauthorized("Unauthorized", nil),
			http.StatusUnauthorized,
		)
		return
	}

	vars := mux.Vars(r)
	pageStr := vars["page"]

	// Convert to integer
	page, err := strconv.Atoi(pageStr)
	if err != nil || page < 1 {
		utils.SendJSONResponse(w,
			utils.BadRequest("INVALID_PAGE_NUMBER", "Page number must be a number"),
			http.StatusBadRequest,
		)
		return
	}

	page -= 1

	projectCount, countErr := h.projectService.GetProjectCount(ctx, userUUID)
	if err != nil {
		utils.SendJSONResponse(w, err, countErr.Status)
		return
	}
	maxPage := int(math.Ceil(float64(projectCount) / 10))
	fmt.Println(maxPage, page)
	if page > maxPage {
		utils.SendJSONResponse(w,
			utils.BadRequest("INVALID_PAGE", "Page does not exist"),
			http.StatusBadRequest,
		)
		return
	}

	projects, appErr := h.projectService.GetProjectsList(ctx, userUUID, page)
	if appErr != nil {
		utils.SendJSONResponse(w, appErr, appErr.Status)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(projects)
}

func (h *ProjectHandler) GetProjectInfo(w http.ResponseWriter, r *http.Request) {

	ctx := r.Context()

	userUUID, ok := ctx.Value("UserUUID").(string)
	if !ok {
		utils.SendJSONResponse(w,
			utils.Unauthorized("Unauthorized", nil),
			http.StatusUnauthorized,
		)
		return
	}

	vars := mux.Vars(r)
	projectUUID := vars["projectUUID"]

	if projectUUID == "" {
		utils.SendJSONResponse(w,
			utils.BadRequest("INVALID_PROJECT_UUID", "Project UUID is required"),
			http.StatusBadRequest,
		)
		return
	}

	project, appErr := h.projectService.GetProjectInfo(ctx, userUUID, projectUUID)
	if appErr != nil {
		utils.SendJSONResponse(w, appErr, appErr.Status)
		return
	}

	utils.SendJSONResponse(w, project, http.StatusOK)
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
