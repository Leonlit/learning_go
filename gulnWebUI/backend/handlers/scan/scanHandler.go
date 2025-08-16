package handlers

import (
	"encoding/json"
	"gulnManagement/gulnWebUI/databases"
	"log"
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
