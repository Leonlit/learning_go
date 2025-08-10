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

func GetScansList(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	pageStr := vars["page"]

	// Convert to integer
	page, err := strconv.Atoi(pageStr)
	if err != nil || page < 1 {
		http.Error(w, "Invalid page number", http.StatusBadRequest)
		return
	}

	userUUID := r.Context().Value("UserUUID").(string)

	scans, err := databases.GetScanList(userUUID, page)
	if err != nil {
		http.Error(w, "Error fetching scans", http.StatusInternalServerError)
		log.Println("GetScanList error:", err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(scans)

}

func GetScanByID(w http.ResponseWriter, r *http.Request) {
	// Fetch scan by ID
}

func UploadScan(w http.ResponseWriter, r *http.Request) {
	userUUID := r.Context().Value("UserUUID").(string)

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

	saved := databases.SaveScanResultsToDatabase(userUUID, nmapRun)
	if !saved {
		http.Error(w, "Error Saving Nmap results", http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, "File uploaded successfully: %s", file)
}
