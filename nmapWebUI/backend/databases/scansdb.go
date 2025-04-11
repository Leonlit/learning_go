package databases

import (
	"log"
	"time"
)

type Scan struct {
	ID         int       `json:"id"`
	UserID     int       `json:"user_id"`
	ScanTime   time.Time `json:"scan_time"`
	HostsUp    int       `json:"hosts_up"`
	HostsDown  int       `json:"hosts_down"`
	TotalHosts int       `json:"total_hosts"`
}

func GetScanList(userID string, page int) ([]Scan, error) {
	offset := (page - 1) * 10
	query := `
		SELECT * FROM scans WHERE user_id = $1 LIMIT 10 OFFSET $2
	`
	rows, err := DBObj.Query(query, userID, offset)
	if err != nil {
		log.Println("Query error:", err)
		return nil, err
	}
	defer rows.Close()
	var scans []Scan
	for rows.Next() {
		var scan Scan
		if err := rows.Scan(&scan.ID, &scan.UserID, &scan.ScanTime, &scan.HostsUp, &scan.HostsDown, &scan.TotalHosts); err != nil {
			return nil, err
		}
		scans = append(scans, scan)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return scans, nil
}
