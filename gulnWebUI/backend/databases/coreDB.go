package databases

import (
	"log"
)

type ProjectAndHostCounts struct {
	ProjectCount *int `json:"project_count"`
	HostCount    *int `json:"host_count"`
}

func GetUserProjectAndHostCounts(userUUID string) (ProjectAndHostCounts, error) {
	query := `
		SELECT
			COUNT(DISTINCT p.project_uuid) AS project_count,
			COUNT(DISTINCT h.host_uuid) AS host_count
		FROM projects p
		LEFT JOIN scans s ON p.project_uuid = s.project_uuid
		LEFT JOIN hosts h ON s.scan_uuid = h.scan_uuid
		WHERE p.user_uuid = $1;
	`
	var result ProjectAndHostCounts
	err := DBObj.QueryRow(query, userUUID).Scan(&result.ProjectCount, &result.HostCount)
	if err != nil {
		log.Println("Query error:", err)
		return ProjectAndHostCounts{}, err
	}

	return result, nil
}
