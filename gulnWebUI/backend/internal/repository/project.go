package repository

import (
	"context"
	"database/sql"
	"log"
	"time"
)

type ProjectRepository struct {
	db *sql.DB
}

func NewProjectRepository(db *sql.DB) *ProjectRepository {
	return &ProjectRepository{db: db}
}

type Project struct {
	ProjectUUID      *string    `json:"project_uuid"`
	ProjectName      *string    `json:"project_name"`
	ProjectCreatedAt *time.Time `json:"project_created_at"`
	ProjectUpdatedAt *time.Time `json:"project_updated_at"`
}

type ScanAndHostsCount struct {
	ScanCount  *int `json:"scan_count"`
	HostsCount *int `json:"hosts_count"`
}

type Host struct {
	HostUUID *string `json:"host_uuid"`
	IPAddr   *string `json:"ip_address"`
	AddrType *string `json:"addr_type"`
	Hostname *string `json:"hostname"`
	Status   *string `json:"status"`
}

type PortInfo struct {
	PortUUID           *string `json:"port_uuid"`
	PortNumber         *int64  `json:"port_number"`
	PortProtocol       *string `json:"protocol"`
	PortState          *string `json:"state"`
	PortReason         *string `json:"reason"`
	PortServiceUUID    *string `json:"service_uuid"`
	PortServiceName    *string `json:"service_name"`
	PortServiceProduct *string `json:"service_product"`
	PortServiceVersion *string `json:"service_version"`
}

type PortDetails struct {
	PortUUID           *string `json:"port_uuid"`
	PortNumber         *int64  `json:"port_number"`
	PortProtocol       *string `json:"protocol"`
	PortState          *string `json:"state"`
	PortReason         *string `json:"reason"`
	PortServiceUUID    *string `json:"service_uuid"`
	PortServiceName    *string `json:"service_name"`
	PortServiceProduct *string `json:"service_product"`
	PortServiceVersion *string `json:"service_version"`
	ServiceFP          *string `json:"service_fp"`
	ServiceCPE         *string `json:"service_cpe"`
	ScriptUUID         *string `json:"script_uuid"`
	ScriptID           *string `json:"script_id"`
	ScriptOutput       *string `json:"script_output"`
}

type CleanPortDetails struct {
	PortUUID           *string   `json:"port_uuid"`
	PortNumber         *int64    `json:"port_number"`
	PortProtocol       *string   `json:"protocol"`
	PortState          *string   `json:"state"`
	PortReason         *string   `json:"reason"`
	PortServiceUUID    *string   `json:"service_uuid"`
	PortServiceName    *string   `json:"service_name"`
	PortServiceProduct *string   `json:"service_product"`
	PortServiceVersion *string   `json:"service_version"`
	ServiceFP          *string   `json:"service_fp"`
	ServiceCPE         *string   `json:"service_cpe"`
	Scripts            []Scripts `json:"scripts"`
}

type Scripts struct {
	ScriptUUID   *string `json:"script_uuid"`
	ScriptID     *string `json:"script_id"`
	ScriptOutput *string `json:"script_output"`
}

func (r *ProjectRepository) GetProjectCount(ctx context.Context, userUUID string) (int, error) {
	query := `
		SELECT COUNT(uuid)
		FROM projects
		WHERE person_in_charge_uuid = $1
	`

	var count int

	err := r.db.QueryRowContext(ctx, query, userUUID).Scan(&count)
	if err != nil {
		log.Println("Query error:", err)
		return 0, err
	}

	return count, nil
}

func (r *ProjectRepository) GetProjectList(ctx context.Context, userUUID string, page int) ([]Project, error) {
	offset := page * 10
	query := `
		SELECT uuid, name, created_at, updated_at FROM projects WHERE person_in_charge_uuid = $1 LIMIT 10 OFFSET $2
	`
	rows, err := r.db.QueryContext(ctx, query, userUUID, offset)
	if err != nil {
		log.Println("Query error:", err)
		return []Project{}, err
	}
	defer rows.Close()

	projects := []Project{}

	for rows.Next() {
		var project Project
		if err := rows.Scan(&project.ProjectUUID, &project.ProjectName, &project.ProjectCreatedAt, &project.ProjectUpdatedAt); err != nil {
			return []Project{}, err
		}
		projects = append(projects, project)
	}

	if err := rows.Err(); err != nil {
		return []Project{}, err
	}

	return projects, nil
}

func (r *ProjectRepository) CreateNewProject(ctx context.Context, userUUID, projectName string) (string, error) {
	var projectUUID string
	query := `
		INSERT INTO projects (uuid, person_in_charge_uuid, name)
		VALUES (uuid_generate_v4(), $1, $2)
		RETURNING uuid
	`

	err := r.db.QueryRowContext(ctx, query,
		userUUID,
		projectName,
	).Scan(&projectUUID)

	if err != nil {
		log.Println("Query error:", err)
		return "", err
	}
	return projectUUID, nil
}

func (r *ProjectRepository) GetProjectInfo(
	ctx context.Context,
	userUUID, projectUUID string,
) (Project, error) {

	query := `
		SELECT uuid, name, created_at, updated_at
		FROM projects
		WHERE uuid = $1 AND person_in_charge_uuid = $2
	`

	var project Project

	err := r.db.QueryRowContext(ctx, query, projectUUID, userUUID).Scan(
		&project.ProjectUUID,
		&project.ProjectName,
		&project.ProjectCreatedAt,
		&project.ProjectUpdatedAt,
	)

	if err != nil {
		return Project{}, err
	}

	return project, nil
}
