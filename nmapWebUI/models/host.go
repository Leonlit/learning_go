package models

type Host struct {
	ID        int    `json:"id"`
	ScanID    int    `json:"scan_id"`
	IPAddress string `json:"ip_address"`
	Status    string `json:"status"`
}
