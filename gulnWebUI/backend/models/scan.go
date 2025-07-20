package models

type Scan struct {
	ID       int    `json:"id"`
	ScanTime string `json:"scan_time"`
	HostsUp  int    `json:"hosts_up"`
}
