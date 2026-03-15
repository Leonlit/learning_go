package service

import (
	"context"
	"gulnManagement/gulnWebUI/internal/repository"
	"gulnManagement/gulnWebUI/internal/utils"
)

type AssessmentService struct {
	assessmentRepo *repository.AssessmentRepository
}

func NewAssessmentService(assessmentRepo *repository.AssessmentRepository) *AssessmentService {
	return &AssessmentService{assessmentRepo: assessmentRepo}
}

func (s *AssessmentService) GetAssessmentCount(ctx context.Context, userUUID string) (int, *utils.AppError) {
	projectCounts, err := s.assessmentRepo.GetAssessmentCount(ctx, userUUID)
	if err != nil {
		return -1, utils.InternalError("Error fetching project count", err)
	}

	return projectCounts, nil
}
