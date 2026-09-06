package handler

//TODO: create generate stats APIs for the users like projects
// and assessment they're involved in

import (
	"gulnManagement/gulnWebUI/internal/service"
	"gulnManagement/gulnWebUI/internal/utils"
	"net/http"
)

type StatsHandler struct {
	statsService *service.StatsService
}

func NewStatsHandler(statsService *service.StatsService) *StatsHandler {
	return &StatsHandler{statsService: statsService}
}

func (h *StatsHandler) GetOwnAssessmentsCount(w http.ResponseWriter, r *http.Request) {

	ctx := r.Context()

	userUUID, ok := ctx.Value("UserUUID").(string)
	if !ok {
		utils.SendJSONResponse(w,
			utils.Unauthorized("Unauthorized", nil),
			http.StatusUnauthorized,
		)
		return
	}

	ownAssessmentCount, err := h.statsService.GetOwnAssessmentCount(ctx, userUUID)
	if err != nil {
		utils.SendJSONResponse(w, err, err.Status)
		return
	}

	utils.SendJSONResponse(w, map[string]int{
		"count": ownAssessmentCount,
	}, http.StatusOK)
}
