package handlers

import (
	"encoding/json"
	"gulnManagement/gulnWebUI/databases"
	"log"
	"net/http"
)

func GetUserProjectAndHostCounts(w http.ResponseWriter, r *http.Request) {

	userUUID := r.Context().Value("UserUUID").(string)

	infos, err := databases.GetUserProjectAndHostCounts(userUUID)

	if err != nil {
		http.Error(w, "Error fetching project and host counts", http.StatusInternalServerError)
		log.Println("GetProjectScan error:", err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(infos)
}
