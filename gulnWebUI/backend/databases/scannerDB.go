package databases

import (
	"encoding/json"
	"fmt"
	"gulnManagement/gulnWebUI/models"
	"log"
	"time"
)

func GetHostPortsInfo(projectUUID, scanUUID string) ([]models.PortMinimalInfo, error) {
	query := `
		SELECT 
			h.host_uuid,
			h.ip_address,
			p.port_uuid, 
			p.port_number,
			sv.service_name,
			sv.service_product,
			sv.service_version,
			sv.service_cpe
		FROM hosts h
		JOIN scans sn ON h.scan_uuid = sn.scan_uuid
			LEFT JOIN ports p ON h.host_uuid = p.host_uuid
			LEFT JOIN services sv ON p.port_uuid = sv.port_uuid
        WHERE sn.project_uuid = $1
        	AND sn.scan_uuid = $2
	`

	rows, err := DBObj.Query(query, projectUUID, scanUUID)
	if err != nil {
		log.Println("Query error:", err)
		return nil, err
	}
	defer rows.Close()
	var portsMinimalInfo []models.PortMinimalInfo
	for rows.Next() {
		var portInfo models.PortMinimalInfo
		if err := rows.Scan(
			&portInfo.HostUUID,
			&portInfo.IPAddress,
			&portInfo.PortUUID,
			&portInfo.PortNumber,
			&portInfo.PortServiceName,
			&portInfo.PortServiceProduct,
			&portInfo.PortServiceVersion,
			&portInfo.PortServiceCPE,
		); err != nil {
			return nil, err
		}
		portsMinimalInfo = append(portsMinimalInfo, portInfo)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return portsMinimalInfo, nil
}

func SaveNVDResults(cpeRes *models.CpeResponse, portUUID string) error {
	for _, item := range cpeRes.Vulnerabilities {
		cve := item.CVE

		version, severity, score, vector := pickCVSS(cve)

		fmt.Println(version, severity, score, vector)

		desc := englishDescription(cve.Descriptions)

		published, _ := time.Parse(time.RFC3339, cve.Published)
		modified, _ := time.Parse(time.RFC3339, cve.LastModified)

		rawJSON, err := marshalCVE(cve)
		if err != nil {
			return err
		}

		_, err = DBObj.Exec(`
            INSERT INTO vulnerabilitiesCVE (
                vulnerability_uuid,
                port_uuid,
                cve_id,
                published,
                last_modified,
                description,
                severity,
                version,
                base_score,
                vector_string,
                nvd_json
            ) VALUES (
                uuid_generate_v4(),
                $1, $2, $3, $4, $5,
                $6, $7, $8, $9, $10
            )
            ON CONFLICT (cve_id, port_uuid) DO NOTHING
        `,
			portUUID,
			cve.ID,
			published,
			modified,
			desc,
			severity,
			version,
			score,
			vector,
			rawJSON,
		)

		if err != nil {
			return err
		}
	}
	return nil
}

func pickCVSS(cve models.CVE) (version, severity string, score float64, vector string) {
	if len(cve.Metrics.CvssMetricV40) > 0 {
		d := cve.Metrics.CvssMetricV40[0].CvssData
		return "4.0", d.BaseSeverity, d.BaseScore, d.VectorString
	}
	if len(cve.Metrics.CvssMetricV31) > 0 {
		d := cve.Metrics.CvssMetricV31[0].CvssData
		return "3.1", d.BaseSeverity, d.BaseScore, d.VectorString
	}
	if len(cve.Metrics.CvssMetricV30) > 0 {
		d := cve.Metrics.CvssMetricV30[0].CvssData
		return "3.0", d.BaseSeverity, d.BaseScore, d.VectorString
	}
	if len(cve.Metrics.CvssMetricV2) > 0 {
		d := cve.Metrics.CvssMetricV2[0].CvssData
		return "2.0", cve.Metrics.CvssMetricV2[0].BaseSeverity, d.BaseScore, d.VectorString
	}
	return "", "", 0, ""
}

func englishDescription(desc []models.LangValue) string {
	for _, d := range desc {
		if d.Lang == "en" {
			return d.Value
		}
	}
	return ""
}

func marshalCVE(cve models.CVE) ([]byte, error) {
	return json.Marshal(cve)
}
