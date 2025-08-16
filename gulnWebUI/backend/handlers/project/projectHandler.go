package handlers

import (
	"encoding/json"
	"fmt"
	"gulnManagement/gulnWebUI/databases"
	"gulnManagement/gulnWebUI/handlers/parser"
	"log"
	"net/http"
	"path/filepath"
	"strconv"

	"github.com/gorilla/mux"
)

func CreateNewProjects(w http.ResponseWriter, r *http.Request) {

	projectName := r.FormValue("projectName")
	fmt.Println(projectName)
	if projectName == "" {
		http.Error(w, "Project name is required", http.StatusBadRequest)
		return
	}

	userUUID := r.Context().Value("UserUUID").(string)

	projects := databases.CreateNewProject(userUUID, projectName)
	if projects == "Error" {
		http.Error(w, "Error creating project", http.StatusInternalServerError)
		log.Println("CreateNewProject error:", projects)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"projectID": projects,
	})
}

func GetProjectsList(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	pageStr := vars["page"]

	// Convert to integer
	page, err := strconv.Atoi(pageStr)
	if err != nil || page < 1 {
		http.Error(w, "Invalid page number", http.StatusBadRequest)
		return
	}

	userUUID := r.Context().Value("UserUUID").(string)

	projects, err := databases.GetProjectList(userUUID, page)
	if err != nil {
		http.Error(w, "Error fetching project list", http.StatusInternalServerError)
		log.Println("GetProjectsList error:", err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(projects)

}

func GetProjectInfo(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	projectUUID := vars["projectUUID"]
	userUUID := r.Context().Value("UserUUID").(string)

	projects, err := databases.GetProjectInfo(userUUID, projectUUID)
	if err != nil {
		http.Error(w, "Error fetching project info", http.StatusInternalServerError)
		log.Println("GetProjectInfo error:", err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(projects)

}

func GetProjectScan(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	projectUUID := vars["projectUUID"]
	pageStr := vars["page"]

	// Convert to integer
	page, err := strconv.Atoi(pageStr)
	if err != nil || page < 1 {
		http.Error(w, "Invalid page number", http.StatusBadRequest)
		return
	}

	projects, err := databases.GetProjectScan(projectUUID, page)
	if err != nil {
		http.Error(w, "Error fetching project scan info", http.StatusInternalServerError)
		log.Println("GetProjectScan error:", err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(projects)

}

func UploadProjectScan(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	projectUUID := vars["projectUUID"]

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

	nmapRun, err := parser.ParseNmap(file)

	if err != nil {
		http.Error(w, "Error Parsing Nmap results", http.StatusBadRequest)
		return
	}

	savedScanUUID := databases.SaveScanResultsToDatabase(projectUUID, scanName, nmapRun)
	if !savedScanUUID {
		http.Error(w, "Error Saving Nmap results", http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, "File uploaded successfully: %s", file)
}
