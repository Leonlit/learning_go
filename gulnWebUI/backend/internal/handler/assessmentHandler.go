package handler

import (
	"gulnManagement/gulnWebUI/internal/service"
	"gulnManagement/gulnWebUI/internal/utils"
	"net/http"
)

type AssessmentHandler struct {
	assessmentService *service.AssessmentService
}

func NewAssessmentHandler(assessmentService *service.AssessmentService) *AssessmentHandler {
	return &AssessmentHandler{assessmentService: assessmentService}
}

func (s *AssessmentHandler) GetAssessmentCount(w http.ResponseWriter, r *http.Request) {

	ctx := r.Context()

	userUUID, ok := ctx.Value("UserUUID").(string)
	if !ok {
		utils.SendJSONResponse(w,
			utils.Unauthorized("Unauthorized", nil),
			http.StatusUnauthorized,
		)
		return
	}

	assessmentCounts, err := s.assessmentService.GetAssessmentCount(ctx, userUUID)
	if err != nil {
		utils.SendJSONResponse(w, err, err.Status)
		return
	}

	utils.SendJSONResponse(w, map[string]int{
		"count": assessmentCounts,
	}, http.StatusOK)
}
