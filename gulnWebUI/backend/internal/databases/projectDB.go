package databases

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

func GetProjectScanUniqueHostsCount(userUUID, projectUUID string) (ScanAndHostsCount, error) {
	query := `
		SELECT 
			COUNT(DISTINCT sc.scan_uuid) AS scan_count,
			COUNT(DISTINCT h.ip_address) AS unique_hosts
		FROM scans sc
		JOIN projects p ON sc.project_uuid = p.project_uuid
		LEFT JOIN hosts h ON sc.scan_uuid = h.scan_uuid
		WHERE p.project_uuid = $1
		AND p.user_uuid = $2;
	`

	var count ScanAndHostsCount
	err := DBObj.QueryRow(query, projectUUID, userUUID).Scan(
		&count.ScanCount,
		&count.HostsCount,
	)
	if err != nil {
		return ScanAndHostsCount{}, err // return empty project + error
	}

	return count, nil
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

func GetProjectScanHostInfo(projectUUID, scanUUID, hostUUID string) ([]PortInfo, error) {
	query := `
		SELECT 
			p.port_uuid, 
			p.port_number, 
			p.protocol, 
			p.state, 
			p.reason,
			sv.service_uuid,
			sv.service_name,
			sv.service_product,
			sv.service_version
		FROM hosts h
		JOIN scans sn ON h.scan_uuid = sn.scan_uuid
			LEFT JOIN ports p ON h.host_uuid = p.host_uuid
        	LEFT JOIN services sv ON p.port_uuid = sv.port_uuid
        WHERE sn.project_uuid = $1
        	AND sn.scan_uuid = $2
        	AND h.host_uuid = $3;
	`

	rows, err := DBObj.Query(query, projectUUID, scanUUID, hostUUID)
	if err != nil {
		log.Println("Query error:", err)
		return nil, err
	}
	defer rows.Close()
	var portsInfo []PortInfo
	for rows.Next() {
		var portInfo PortInfo
		if err := rows.Scan(&portInfo.PortUUID, &portInfo.PortNumber, &portInfo.PortProtocol,
			&portInfo.PortState, &portInfo.PortReason, &portInfo.PortServiceUUID,
			&portInfo.PortServiceName, &portInfo.PortServiceProduct, &portInfo.PortServiceVersion); err != nil {
			return nil, err
		}
		portsInfo = append(portsInfo, portInfo)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return portsInfo, nil
}

func GetProjectScanHostPortInfo(projectUUID, scanUUID, hostUUID, portUUID string) (CleanPortDetails, error) {
	query := `
		SELECT 
			p.port_uuid, 
			p.port_number, 
			p.protocol, 
			p.state, 
			p.reason,
			sv.service_uuid,
			sv.service_name,
			sv.service_product,
			sv.service_version,
			sv.service_fp,
			sv.service_cpe,
			sc.script_uuid,
			sc.script_id,
			sc.script_output
		FROM hosts h
		JOIN scans sn ON h.scan_uuid = sn.scan_uuid
		LEFT JOIN ports p ON h.host_uuid = p.host_uuid
		LEFT JOIN services sv ON p.port_uuid = sv.port_uuid
		LEFT JOIN scripts sc ON p.port_uuid = sc.port_uuid
		WHERE sn.project_uuid = $1
		AND sn.scan_uuid = $2
		AND h.host_uuid = $3
		AND p.port_uuid = $4;
	`

	rows, err := DBObj.Query(query, projectUUID, scanUUID, hostUUID, portUUID)
	if err != nil {
		log.Println("Query error:", err)
		return CleanPortDetails{}, err
	}
	defer rows.Close()
	var portDetails []PortDetails
	for rows.Next() {
		var portDetail PortDetails
		if err := rows.Scan(&portDetail.PortUUID, &portDetail.PortNumber, &portDetail.PortProtocol,
			&portDetail.PortState, &portDetail.PortReason, &portDetail.PortServiceUUID,
			&portDetail.PortServiceName, &portDetail.PortServiceProduct, &portDetail.PortServiceVersion,
			&portDetail.ServiceFP, &portDetail.ServiceCPE, &portDetail.ScriptUUID, &portDetail.ScriptID,
			&portDetail.ScriptOutput); err != nil {
			return CleanPortDetails{}, err
		}
		portDetails = append(portDetails, portDetail)
	}

	if err := rows.Err(); err != nil {
		return CleanPortDetails{}, err
	}

	cleanedPortDetails := CleanPortDetails{
		PortUUID:           portDetails[0].PortUUID,
		PortNumber:         portDetails[0].PortNumber,
		PortProtocol:       portDetails[0].PortProtocol,
		PortState:          portDetails[0].PortState,
		PortReason:         portDetails[0].PortReason,
		PortServiceUUID:    portDetails[0].PortServiceUUID,
		PortServiceName:    portDetails[0].PortServiceName,
		PortServiceProduct: portDetails[0].PortServiceProduct,
		PortServiceVersion: portDetails[0].PortServiceVersion,
		ServiceFP:          portDetails[0].ServiceFP,
		ServiceCPE:         portDetails[0].ServiceCPE,
	}

	var scripts []Scripts
	for _, pd := range portDetails {
		if pd.ScriptUUID != nil { // only add if script exists
			scripts = append(scripts, Scripts{
				ScriptUUID:   pd.ScriptUUID,
				ScriptID:     pd.ScriptID,
				ScriptOutput: pd.ScriptOutput,
			})
		}
	}

	cleanedPortDetails.Scripts = scripts

	return cleanedPortDetails, nil
}
