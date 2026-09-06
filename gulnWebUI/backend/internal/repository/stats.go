package repository

import (
	"context"
	"database/sql"
	"log"
)

type StatsRepository struct {
	db *sql.DB
}

func NewStatsRepository(db *sql.DB) *StatsRepository {
	return &StatsRepository{db: db}
}

func (r *StatsRepository) GetOwnAssessmentsCount(ctx context.Context, userUUID string) (int, error) {
	query := `
		SELECT COUNT(assessment_uuid)
		FROM assessment_teams
		WHERE user_uuid = $1
	`

	var count int

	err := r.db.QueryRowContext(ctx, query, userUUID).Scan(&count)
	if err != nil {
		log.Println("Query error:", err)
		return 0, err
	}

	return count, nil
}
