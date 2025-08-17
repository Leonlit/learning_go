package databases

import (
	"log"
	"time"
)

type Project struct {
	ProjectUUID    string    `json:"project_uuid"`
	ProjectName    string    `json:"project_name"`
	ProjectCreated time.Time `json:"project_created"`
}

type Host struct {
	HostUUID string `json:"host_uuid"`
	IPAddr   string `json:"ip_address"`
	AddrType string `json:"addr_type"`
	Hostname string `json:"hostname"`
	Status   string `json:"status"`
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

func GetProjectScan(projectUUID string, page int) ([]Scan, error) {
	offset := (page - 1) * 10
	query := `
		SELECT scan_uuid, scan_name, scan_start_time, scan_finish_time, 
		hosts_up, hosts_down, total_hosts FROM scans
		WHERE project_uuid = $1
		LIMIT 10 OFFSET $2
	`

	rows, err := DBObj.Query(query, projectUUID, offset)
	if err != nil {
		log.Println("Query error:", err)
		return nil, err
	}
	defer rows.Close()
	var scans []Scan
	for rows.Next() {
		var scan Scan
		if err := rows.Scan(&scan.ScanUUID, &scan.ScanName, &scan.ScanStartTime, &scan.ScanFinishedTime, &scan.HostsUp, &scan.HostsDown, &scan.TotalHosts); err != nil {
			return nil, err
		}
		scans = append(scans, scan)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return scans, nil
}

func GetProjectScanHosts(projectUUID, scanUUID string, page int) ([]Host, error) {
	offset := (page - 1) * 10
	query := `
		SELECT h.host_uuid, 
		h.ip_address, 
		h.addr_type, 
		h.hostname, 
		h.status
		FROM hosts h
		JOIN scans s ON h.scan_uuid = s.scan_uuid
		WHERE s.project_uuid = $1
		AND s.scan_uuid = $2
		LIMIT 10 OFFSET $3;
	`

	rows, err := DBObj.Query(query, projectUUID, scanUUID, offset)
	if err != nil {
		log.Println("Query error:", err)
		return nil, err
	}
	defer rows.Close()
	var hosts []Host
	for rows.Next() {
		var host Host
		if err := rows.Scan(&host.HostUUID, &host.IPAddr, &host.AddrType, &host.Hostname, &host.Status); err != nil {
			return nil, err
		}
		hosts = append(hosts, host)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return hosts, nil
}
