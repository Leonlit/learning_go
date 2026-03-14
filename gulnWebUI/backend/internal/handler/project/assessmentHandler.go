package handler

import (
	"encoding/json"
	databases "gulnManagement/gulnWebUI/internal/repository"
	"log"
	"net/http"
)

func GetAssessmentCount(w http.ResponseWriter, r *http.Request) {

	userUUID := r.Context().Value("UserUUID").(string)

	assessmentCounts, err := databases.GetProjectCount(userUUID)
	if err != nil {
		http.Error(w, "Error fetching project list", http.StatusInternalServerError)
		log.Println("GetProjectsList error:", err)
		return
	}

	res := map[string]int{
		"count": assessmentCounts,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(res)
}
