package handlers

import (
	"net/http"
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

	/* // Example: Return page number in response
	json.NewEncoder(w).Encode(map[string]interface{}{
		"message": "Scans page retrieved successfully",
		"page":    page,
	}) */

}

func GetScanByID(w http.ResponseWriter, r *http.Request) {
	// Fetch scan by ID
}

func UploadScan(w http.ResponseWriter, r *http.Request) {
	// Insert new scan into DB
}
