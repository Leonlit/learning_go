package repository

import (
	"context"
	"database/sql"
	"gulnManagement/gulnWebUI/internal/dto"
	"log"
)

type TeamMemberRepository struct {
	db *sql.DB
}

func NewTeamMemberRepository(db *sql.DB) *TeamMemberRepository {
	return &TeamMemberRepository{db: db}
}

type TeamMember struct {
	TeamMemberUUID *string `json:"team_member_uuid"`
	TeamMemberName *string `json:"team_member_name"`
}

func (r *TeamMemberRepository) GetTeamMemberCount(ctx context.Context) (int, error) {
	query := `
		SELECT COUNT(uuid) FROM project_team_members
	`

	var count int

	err := r.db.QueryRowContext(ctx, query).Scan(&count)
	if err != nil {
		log.Println("Query error:", err)
		return 0, err
	}

	return count, nil
}

func (r *TeamMemberRepository) GetTeamMemberList(ctx context.Context, page int) ([]TeamMember, error) {
	offset := page * 10
	query := `
		SELECT uuid, name, department, role FROM project_team_members LIMIT 10 OFFSET $1
	`
	rows, err := r.db.QueryContext(ctx, query, offset)
	if err != nil {
		log.Println("Query error:", err)
		return []TeamMember{}, err
	}
	defer rows.Close()

	projects := []TeamMember{}

	for rows.Next() {
		var teamMember TeamMember
		if err := rows.Scan(&teamMember.TeamMemberUUID, &teamMember.TeamMemberName); err != nil {
			return []TeamMember{}, err
		}
		projects = append(projects, teamMember)
	}

	if err := rows.Err(); err != nil {
		return []TeamMember{}, err
	}

	return projects, nil
}

func (r *TeamMemberRepository) AddNewTeamMember(ctx context.Context, data dto.AddTeamMemberRequest) (string, error) {
	var TeamMemberUUID string
	query := `
		INSERT INTO projects (uuid, name, department, role)
		VALUES (uuid_generate_v4(), $1, $2, $3)
		RETURNING uuid
	`

	err := r.db.QueryRowContext(ctx, query,
		data.Name,
		data.Department,
		data.Role,
	).Scan(&TeamMemberUUID)

	if err != nil {
		log.Println("Query error:", err)
		return "", err
	}
	return TeamMemberUUID, nil
}

func (r *TeamMemberRepository) GetTeamMemberInfo(
	ctx context.Context,
	userUUID, TeamMemberUUID string,
) (TeamMember, error) {

	query := `
		SELECT uuid, project_name, created_time
		FROM projects
		WHERE uuid = $1 AND person_in_charge_uuid = $2
	`

	var teamMember TeamMember

	err := r.db.QueryRowContext(ctx, query, TeamMemberUUID, userUUID).Scan(
		&teamMember.TeamMemberUUID,
		&teamMember.TeamMemberName,
	)

	if err != nil {
		return TeamMember{}, err
	}

	return teamMember, nil
}
