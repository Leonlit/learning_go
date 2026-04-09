package dto

type AddTeamMemberRequest struct {
	Name       string `json:"name"`
	Department string `json:"department"`
	Role       string `json:"role"`
}
