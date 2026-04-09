package service

import (
	"context"
	"database/sql"
	"errors"
	"gulnManagement/gulnWebUI/internal/dto"
	"gulnManagement/gulnWebUI/internal/repository"
	"gulnManagement/gulnWebUI/internal/utils"
)

type TeamMemberService struct {
	teamMemberRepo *repository.TeamMemberRepository
}

func NewTeamMemberService(teamMemberRepo *repository.TeamMemberRepository) *TeamMemberService {
	return &TeamMemberService{teamMemberRepo: teamMemberRepo}
}

func (s *TeamMemberService) GetTeamMemberCount(ctx context.Context, userUUID string) (int, *utils.AppError) {
	projectCounts, err := s.teamMemberRepo.GetTeamMemberCount(ctx)
	if err != nil {
		return -1, utils.InternalError("Error fetching project count", err)
	}

	return projectCounts, nil
}

func (s *TeamMemberService) GetTeamMemberList(ctx context.Context, page int) ([]repository.TeamMember, *utils.AppError) {

	projects, err := s.teamMemberRepo.GetTeamMemberList(ctx, page)
	if err != nil {
		return []repository.TeamMember{}, utils.InternalError("Error fetching project list", err)
	}
	return projects, nil
}

func (s *TeamMemberService) AddNewTeamMember(ctx context.Context, data dto.AddTeamMemberRequest) (string, *utils.AppError) {

	projectUUID, err := s.teamMemberRepo.AddNewTeamMember(ctx, data)
	if err != nil {
		return "", utils.InternalError("Error creating new project", err)
	}

	return projectUUID, nil
}

func (s *TeamMemberService) GetTeamMemberInfo(ctx context.Context, userUUID, projectUUID string) (repository.TeamMember, *utils.AppError) {

	project, err := s.teamMemberRepo.GetTeamMemberInfo(ctx, userUUID, projectUUID)

	if err != nil {

		if errors.Is(err, sql.ErrNoRows) {
			return repository.TeamMember{}, utils.BadRequest(
				"PROJECT_NOT_FOUND",
				"Project not found",
			)
		}

		return repository.TeamMember{}, utils.InternalError(
			"Error fetching project info",
			err,
		)
	}

	return project, nil
}
