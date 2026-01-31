package repository

import (
	"log"
	"time"
)

type Project struct {
	ProjectUUID    *string    `json:"project_uuid"`
	ProjectName    *string    `json:"project_name"`
	ProjectCreated *time.Time `json:"project_created"`
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

func GetProjectList(userUUID string, page int) ([]Project, error) {
	offset := (page - 1) * 10
	query := `
		SELECT project_uuid, project_name, project_created FROM projects WHERE user_uuid = $1 LIMIT 10 OFFSET $2
	`
	rows, err := DBObj.Query(query, userUUID, offset)
	if err != nil {
		log.Println("Query error:", err)
		return nil, err
	}
	defer rows.Close()
	var projects []Project
	for rows.Next() {
		var project Project
		if err := rows.Scan(&project.ProjectUUID, &project.ProjectName, &project.ProjectCreated); err != nil {
			return nil, err
		}
		projects = append(projects, project)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return projects, nil
}

func CreateNewProject(userUUID, projectName string) string {
	var projectUUID string
	query := `
		INSERT INTO projects (project_uuid, user_uuid, project_name, project_created)
		VALUES (uuid_generate_v4(), $1, $2, $3)
		RETURNING project_uuid
	`

	err := DBObj.QueryRow(query,
		userUUID,
		projectName,
		time.Now(),
	).Scan(&projectUUID)
	if err != nil {
		log.Fatal(err)
		return "Error"
	}
	return projectUUID
}

func GetProjectInfo(userUUID, projectUUID string) (Project, error) {
	query := `
		SELECT project_uuid, project_name, project_created
		FROM projects
		WHERE project_uuid = $1 AND user_uuid = $2
	`

	var project Project
	err := DBObj.QueryRow(query, projectUUID, userUUID).Scan(
		&project.ProjectUUID,
		&project.ProjectName,
		&project.ProjectCreated,
	)
	if err != nil {
		return Project{}, err // return empty project + error
	}

	return project, nil
}
