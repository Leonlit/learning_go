package scanner

import (
	"encoding/json"
	"fmt"
	"gulnManagement/gulnWebUI/databases"
	"gulnManagement/gulnWebUI/models"
	"io"
	"log"
	"net/http"
	"strings"
	"time"
)

// https://services.nvd.nist.gov/rest/json/cves/2.0?cpeName=cpe:2.3:a:boa:boa:0.93.15&resultsPerPage=10
const CPE_SEARCH_API_START string = "https://services.nvd.nist.gov/rest/json/cves/2.0?cpeName=cpe:2.3:"
const CPE_SEARCH_API_END string = "&resultsPerPage=10"

// https://nvd.nist.gov/vuln/search#/nvd/home?cpeFilterMode=cpe&cpeName=cpe:2.3:a:boa:boa:0.94.13:*:*:*:*:*:*:*&resultType=records
const CPE_MULTI_LIST_LINK_START string = "https://nvd.nist.gov/vuln/search#/nvd/home?cpeFilterMode=cpe&cpeName=cpe:2.2:"
const CPE_MULTI_LIST_LINK_END string = ":*:*:*:*:*:*:*&resultType=records"

func GetPortVulns(projectUUID string, scanUUID string) {

	portsInfoArr, err := databases.GetHostPortsInfo(projectUUID, scanUUID)
	if err != nil {
		log.Println("GetProjectsList error:", err)
		return
	}

	for _, portInfo := range portsInfoArr {
		fmt.Printf("scanning %s with port %d\n", *portInfo.IPAddress, *portInfo.PortNumber)
		fmt.Printf("Running with: %s version %s\n\n", *portInfo.PortServiceName, *portInfo.PortServiceVersion)
		nvdScanner(portInfo)
	}
}

func nvdScanner(portInfo models.PortMinimalInfo) {
	cpeRes := getNVDVulns(*portInfo.PortServiceName, *portInfo.PortServiceCPE, *portInfo.PortServiceVersion)
	saveNVDVulns(cpeRes, *portInfo.PortUUID)
}

func getNVDVulns(productName, portCPE, productVersion string) *models.CpeResponse {
	c := &http.Client{
		Timeout: 10 * time.Second,
	}
	if portCPE != "" {
		cleanedPortCPE := strings.TrimPrefix(portCPE, "cpe:/")
		fmt.Println(portCPE)
		fmt.Println(cleanedPortCPE)
		url := CPE_SEARCH_API_START + cleanedPortCPE + CPE_SEARCH_API_END
		fmt.Println(url)

		res, err := c.Get(url)
		if err != nil {
			fmt.Println(err)
		}

		if res.StatusCode != http.StatusOK {
			log.Printf("Error: Got %d status code for %s", res.StatusCode, url)
		}

		body, err := io.ReadAll(res.Body)
		if err != nil {
			log.Println("Read failed:", err)
		}

		// Parse JSON into struct
		var cpeResp models.CpeResponse
		err = json.Unmarshal(body, &cpeResp)
		if err != nil {
			log.Println("JSON parse error:", err)
		}

		return &cpeResp
	} else {
		return &models.CpeResponse{}
	}
}

func saveNVDVulns(cpeRes *models.CpeResponse, portUUID string) error {
	err := databases.SaveNVDResults(cpeRes, portUUID)
	if err != nil {
		log.Println("Error: Unable to save NVD results,", err)
		return err
	}
	return nil
}
