package handler

import (
	"encoding/json"
	"fmt"
	"gulnManagement/gulnWebUI/internal/dto"
	"gulnManagement/gulnWebUI/internal/service"
	"gulnManagement/gulnWebUI/internal/utils"
	"math"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type TeamMemberHandler struct {
	teamMemberService *service.TeamMemberService
}

func NewTeamMemberHandler(teamMemberService *service.TeamMemberService) *TeamMemberHandler {
	return &TeamMemberHandler{teamMemberService: teamMemberService}
}

func (h *TeamMemberHandler) GetTeamMemberCount(w http.ResponseWriter, r *http.Request) {

	ctx := r.Context()

	userUUID, ok := ctx.Value("UserUUID").(string)
	if !ok {
		utils.SendJSONResponse(w,
			utils.Unauthorized("Unauthorized", nil),
			http.StatusUnauthorized,
		)
		return
	}

	teamMemberCount, err := h.teamMemberService.GetTeamMemberCount(ctx, userUUID)
	if err != nil {
		utils.SendJSONResponse(w, err, err.Status)
		return
	}

	utils.SendJSONResponse(w, map[string]int{
		"count": teamMemberCount,
	}, http.StatusOK)
}

func (h *TeamMemberHandler) GetTeamMemberList(w http.ResponseWriter, r *http.Request) {

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

	teamMemberCount, countErr := h.teamMemberService.GetTeamMemberCount(ctx, userUUID)
	if err != nil {
		utils.SendJSONResponse(w, err, countErr.Status)
		return
	}
	maxPage := int(math.Ceil(float64(teamMemberCount) / 10))
	fmt.Println(maxPage, page)
	if page > maxPage {
		utils.SendJSONResponse(w,
			utils.BadRequest("INVALID_PAGE", "Page does not exist"),
			http.StatusBadRequest,
		)
		return
	}

	teamMembers, appErr := h.teamMemberService.GetTeamMemberList(ctx, page)
	if appErr != nil {
		utils.SendJSONResponse(w, appErr, appErr.Status)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(teamMembers)
}

func (h *TeamMemberHandler) AddNewTeamMember(w http.ResponseWriter, r *http.Request) {

	var req dto.AddTeamMemberRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	teamMembersID, err := h.teamMemberService.AddNewTeamMember(r.Context(), req)
	if err != nil {
		utils.SendJSONResponse(w, err, err.Status)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"teamMemberID": teamMembersID,
	})
}

func (h *TeamMemberHandler) GetTeamMemberInfo(w http.ResponseWriter, r *http.Request) {

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
	teamMemberUUID := vars["teamMemberUUID"]

	if teamMemberUUID == "" {
		utils.SendJSONResponse(w,
			utils.BadRequest("INVALID_PROJECT_UUID", "Team Member UUID is required"),
			http.StatusBadRequest,
		)
		return
	}

	teamMember, appErr := h.teamMemberService.GetTeamMemberInfo(ctx, userUUID, teamMemberUUID)
	if appErr != nil {
		utils.SendJSONResponse(w, appErr, appErr.Status)
		return
	}

	utils.SendJSONResponse(w, teamMember, http.StatusOK)
}
