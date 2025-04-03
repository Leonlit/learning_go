package handlers

import (
	"encoding/json"
	"net/http"
	"nmapManagement/nmapWebUI/models"
)

func GetHosts(w http.ResponseWriter, r *http.Request) {
	hosts := []models.Host{} // Replace with DB call
	json.NewEncoder(w).Encode(hosts)
}

func GetHostByID(w http.ResponseWriter, r *http.Request) {
	// Fetch host by ID
}
