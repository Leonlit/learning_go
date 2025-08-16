package databases

import (
	"gulnManagement/gulnWebUI/handlers/parser"
	"log"
	"strings"
	"time"
)

func SaveScanResultsToDatabase(userUUID string, nmapStruct *parser.NmapRun) bool {

	var scanUUID string
	query := `
		INSERT INTO scans (scan_uuid, user_uuid, scan_start_time, scan_finish_time, hosts_up, hosts_down, total_hosts)
		VALUES (uuid_generate_v4(), $1, $2, $3, $4, $5, $6)
		RETURNING scan_uuid
	`

	err := DBObj.QueryRow(query,
		userUUID,
		time.Unix(nmapStruct.StartTime, 0),
		time.Unix(nmapStruct.RunStats.Finished.Time, 0),
		nmapStruct.RunStats.Hosts.Up,
		nmapStruct.RunStats.Hosts.Down,
		nmapStruct.RunStats.Hosts.Total,
	).Scan(&scanUUID)
	if err != nil {
		log.Fatal(err)
		return false
	}

	hostQuery := `
		INSERT INTO hosts (host_uuid, scan_uuid, ip_address, addr_type, hostname)
		VALUES (uuid_generate_v4(), $1, $2, $3, $4)
		RETURNING host_uuid
	`

	portQuery := `
		INSERT INTO ports (port_uuid, host_uuid, port_number, protocol, state, reason)
		VALUES (uuid_generate_v4(), $1, $2, $3, $4, $5)
		RETURNING port_uuid
	`

	serviceQuery := `
		INSERT INTO services (service_uuid, port_uuid, service_name, service_product, service_version, service_fp, service_cpe)
		VALUES (uuid_generate_v4(), $1, $2, $3, $4, $5, $6)
	`

	scriptQuery := `
		INSERT INTO scripts (script_uuid, port_uuid, script_id, script_output)
		VALUES (uuid_generate_v4(), $1, $2, $3)
		RETURNING script_uuid
	`

	scriptElemQuery := `
		INSERT INTO scripts_element (script_element_uuid, script_uuid, elem_key, elem_value)
		VALUES (uuid_generate_v4(), $1, $2, $3)
	`

	// Saving hosts data
	for _, host := range nmapStruct.TargetHosts {
		var hostUUID string

		var hostnames []string
		for _, hn := range host.Hostname {
			hostnames = append(hostnames, hn.Name+"("+hn.Type+")")
		}
		hostnameStr := strings.Join(hostnames, ", ")

		err := DBObj.QueryRow(hostQuery,
			scanUUID,
			host.HostIP.Addr,
			host.HostIP.AddrType,
			hostnameStr,
		).Scan(&hostUUID)

		if err != nil {
			log.Fatal(err)
			continue
		}
		// Saving host's ports data
		for _, port := range host.Ports {
			var portUUID string
			err := DBObj.QueryRow(portQuery,
				hostUUID,
				port.ID,
				port.Protocol,
				port.State.State,
				port.State.Reason,
			).Scan(&portUUID)

			if err != nil {
				log.Fatal(err)
				continue
			}

			// Saving service data
			serviceSavedResult := DBObj.QueryRow(serviceQuery,
				portUUID,
				port.Service.ServiceName,
				port.Service.ServiceProduct,
				port.Service.ServiceVersion,
				port.Service.ServiceFingerPrint,
				port.Service.ServiceCPE.Value,
			)

			if serviceSavedResult.Err() != nil {
				log.Fatal(serviceSavedResult.Err().Error())
				continue
			}

			// Saving scripts data
			for _, script := range port.Scripts {
				var scriptUUID string
				err = DBObj.QueryRow(scriptQuery,
					portUUID,
					script.Id,
					script.Output,
				).Scan(&scriptUUID)

				if err != nil {
					log.Fatal(err)
					continue
				}

				// Saving script elements
				for _, elem := range script.Elems {
					elemSavedResults := DBObj.QueryRow(scriptElemQuery,
						scriptUUID,
						elem.Key,
						elem.Value,
					)

					if elemSavedResults.Err() != nil {
						log.Fatal(elemSavedResults.Err().Error())
						continue
					}
				}
			}

		}
	}
	return true
}
