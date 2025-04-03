package parser

type NmapRun struct {
	RunStats    RunStats `xml:"runstats"`
	TargetHosts []Host   `xml:"host"`
}

type Host struct {
	HostIP   Address    `xml:"address"`
	Hostname []HostName `xml:"hostnames>hostname"`
	Ports    []HostPort `xml:"ports>port"`
}

// Host IP Address for current Host
type Address struct {
	Addr     string `xml:"addr,attr"`
	AddrType string `xml:"addrtype,attr"`
}

// Hostnames for current Host
type HostName struct {
	Name string `xml:"name,attr"`
	Type string `xml:"type,attr"`
}

// Ports for current Host
type HostPort struct {
	Protocol string       `xml:"protocol,attr"`
	ID       int          `xml:"portid,attr"`
	State    PortState    `xml:"state"`
	Service  PortService  `xml:"service"`
	Scripts  []PortScript `xml:"script"`
}

type PortState struct {
	State  string `xml:"state,attr"`
	Reason string `xml:"reason,attr"`
}

type PortService struct {
	ServiceName        string `xml:"name,attr"`
	ServiceProduct     string `xml:"product,attr"`
	ServiceVersion     string `xml:"version,attr"`
	ServiceFingerPrint string `xml:"servicefp,attr"`
	ServiceCPE         []CPE  `xml:"cpe"`
}

type CPE struct {
	Value string `xml:",chardata"`
}

type PortScript struct {
	Id     string `xml:"id,attr"`
	Output string `xml:"output,attr"`
	Elems  []Elem `xml:"elem"`
}

type Elem struct {
	Key   string `xml:"key,attr"`
	Value string `xml:",chardata"`
}

// Scanning availablity status
type RunStats struct {
	Hosts Hosts `xml:"hosts"`
}

// hosts stats
type Hosts struct {
	Up    int `xml:"up,attr"`
	Down  int `xml:"down,attr"`
	Total int `xml:"total,attr"`
}
