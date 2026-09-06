package service

import (
	"context"
	"gulnManagement/gulnWebUI/internal/repository"
	"gulnManagement/gulnWebUI/internal/utils"
)

type StatsService struct {
	statsRepo *repository.StatsRepository
}

func NewStatsService(statsRepo *repository.StatsRepository) *StatsService {
	return &StatsService{statsRepo: statsRepo}
}

func (s *StatsService) GetOwnAssessmentCount(ctx context.Context, userUUID string) (int, *utils.AppError) {
	projectCounts, err := s.statsRepo.GetOwnAssessmentsCount(ctx, userUUID)
	if err != nil {
		return -1, utils.InternalError("Error fetching project count", err)
	}

	return projectCounts, nil
}
