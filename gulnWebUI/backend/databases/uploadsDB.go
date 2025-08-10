package databases

import (
	"gulnManagement/gulnWebUI/handlers/parser"
	"log"
	"time"
)

func SaveScanResultsToDatabase(userUUID string, nmapStruct *parser.NmapRun) bool {
	query := `
		INSERT INTO scans (scan_uuid, user_uuid, scan_start_time, scan_finish_time, hosts_up, hosts_down, total_hosts)
		VALUES (uuid_generate_v4(), $1, $2, $3, $4, $5, $6)
	`

	_, err := DBObj.Exec(query,
		userUUID,
		time.Unix(nmapStruct.StartTime, 0),
		time.Unix(nmapStruct.RunStats.Finished.Time, 0),
		nmapStruct.RunStats.Hosts.Up,
		nmapStruct.RunStats.Hosts.Down,
		nmapStruct.RunStats.Hosts.Total,
	)
	if err != nil {
		log.Fatal(err)
		return false
	}
	return true
}
