package models

type PortMinimalInfo struct {
	HostUUID           *string `json:"host_uuid"`
	IPAddress          *string `json:"ip_address"`
	PortUUID           *string `json:"port_uuid"`
	PortNumber         *int64  `json:"port_number"`
	PortServiceName    *string `json:"service_name"`
	PortServiceProduct *string `json:"service_product"`
	PortServiceVersion *string `json:"service_version"`
	PortServiceCPE     *string `json:"service_cpe"`
}
