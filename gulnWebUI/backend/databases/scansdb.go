package databases

import (
	"log"
	"time"
)

type Scan struct {
	ScanUUID         string    `json:"scan_uuid"`
	UserUUID         string    `json:"user_uuid"`
	ScanStartTime    time.Time `json:"scan_start_time"`
	ScanFinishedTime time.Time `json:"scan_finish_time"`
	HostsUp          int       `json:"hosts_up"`
	HostsDown        int       `json:"hosts_down"`
	TotalHosts       int       `json:"total_hosts"`
}

func GetScanList(userUUID string, page int) ([]Scan, error) {
	offset := (page - 1) * 10
	query := `
		SELECT * FROM scans WHERE user_uuid = $1 LIMIT 10 OFFSET $2
	`
	rows, err := DBObj.Query(query, userUUID, offset)
	if err != nil {
		log.Println("Query error:", err)
		return nil, err
	}
	defer rows.Close()
	var scans []Scan
	for rows.Next() {
		var scan Scan
		if err := rows.Scan(&scan.ScanUUID, &scan.UserUUID, &scan.ScanStartTime, &scan.ScanFinishedTime, &scan.HostsUp, &scan.HostsDown, &scan.TotalHosts); err != nil {
			return nil, err
		}
		scans = append(scans, scan)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return scans, nil
}
