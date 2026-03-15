package repository

import (
	"context"
	"database/sql"
	"log"
	"time"
)

type AssessmentRepository struct {
	db *sql.DB
}

func NewAssessmentRepository(db *sql.DB) *AssessmentRepository {
	return &AssessmentRepository{db: db}
}

type Assessment struct {
	AssessmentUUID    *string    `json:"project_uuid"`
	AssessmentName    *string    `json:"project_name"`
	AssessmentCreated *time.Time `json:"project_created"`
}

func (r *AssessmentRepository) GetAssessmentCount(ctx context.Context, userUUID string) (int, error) {
	query := `
		SELECT COUNT(uuid)
		FROM assessments
		WHERE created_by = $1
	`

	var count int

	err := r.db.QueryRowContext(ctx, query, userUUID).Scan(&count)
	if err != nil {
		log.Println("Query error:", err)
		return 0, err
	}

	return count, nil
}
