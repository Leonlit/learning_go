package scanner

type CpeResponse struct {
	TotalResults    int             `json:"totalResults"`
	NVDVersion      string          `json:"version"`
	Vulnerabilities []Vulnerability `json:"vulnerabilities"`
}

type Vulnerability struct {
	TotalResults    int                 `json:"totalResults"`
	Version         string              `json:"version"`
	Vulnerabilities []VulnerabilityItem `json:"vulnerabilities"`
}

type VulnerabilityItem struct {
	CVE CVE `json:"cve"`
}

type CVE struct {
	ID             string          `json:"id"`
	Published      string          `json:"published"`
	LastModified   string          `json:"lastModified"`
	Descriptions   []LangValue     `json:"descriptions"`
	Metrics        Metrics         `json:"metrics"`
	Weaknesses     []Weakness      `json:"weaknesses"`
	Configurations []Configuration `json:"configurations"`
	References     []Reference     `json:"references"`
}

type LangValue struct {
	Lang  string `json:"lang"`
	Value string `json:"value"`
}

type Metrics struct {
	CvssMetricV30 []CvssMetricV3 `json:"cvssMetricV30"`
	CvssMetricV31 []CvssMetricV3 `json:"cvssMetricV31"`
	CvssMetricV40 []CvssMetricV4 `json:"cvssMetricV40"`
	CvssMetricV2  []CvssMetricV2 `json:"cvssMetricV2"`
}

type CvssMetricV4 struct {
	Source              string     `json:"source"`
	Type                string     `json:"type"`
	CvssData            CvssDataV4 `json:"cvssData"`
	ExploitabilityScore float64    `json:"exploitabilityScore"`
	ImpactScore         float64    `json:"impactScore"`
}

type CvssDataV4 struct {
	Version                           string  `json:"version"`
	VectorString                      string  `json:"vectorString"`
	BaseScore                         float64 `json:"baseScore"`
	BaseSeverity                      string  `json:"baseSeverity"`
	AttackVector                      string  `json:"attackVector"`
	AttackComplexity                  string  `json:"attackComplexity"`
	AttackRequirements                string  `json:"attackRequirements"`
	PrivilegesRequired                string  `json:"privilegesRequired"`
	UserInteraction                   string  `json:"userInteraction"`
	Scope                             string  `json:"scope"`
	ConfidentialityImpact             string  `json:"confidentialityImpact"`
	IntegrityImpact                   string  `json:"integrityImpact"`
	AvailabilityImpact                string  `json:"availabilityImpact"`
	VulnConfidentialityImpact         string  `json:"vulnConfidentialityImpact"`
	VulnIntegrityImpact               string  `json:"vulnIntegrityImpact"`
	VulnAvailabilityImpact            string  `json:"vulnAvailabilityImpact"`
	SubConfidentialityImpact          string  `json:"subConfidentialityImpact"`
	SubIntegrityImpact                string  `json:"subIntegrityImpact"`
	SubAvailabilityImpact             string  `json:"subAvailabilityImpact"`
	ExploitMaturity                   string  `json:"exploitMaturity"`
	ConfidentialityRequirement        string  `json:"confidentialityRequirement"`
	IntegrityRequirement              string  `json:"integrityRequirement"`
	AvailabilityRequirement           string  `json:"availabilityRequirement"`
	ModifiedAttackVector              string  `json:"modifiedAttackVector"`
	ModifiedAttackComplexity          string  `json:"modifiedAttackComplexity"`
	ModifiedAttackRequirements        string  `json:"modifiedAttackRequirements"`
	ModifiedPrivilegesRequired        string  `json:"modifiedPrivilegesRequired"`
	ModifiedUserInteraction           string  `json:"modifiedUserInteraction"`
	ModifiedVulnConfidentialityImpact string  `json:"modifiedVulnConfidentialityImpact"`
	ModifiedVulnIntegrityImpact       string  `json:"modifiedVulnIntegrityImpact"`
	ModifiedVulnAvailabilityImpact    string  `json:"modifiedVulnAvailabilityImpact"`
	ModifiedSubConfidentialityImpact  string  `json:"modifiedSubConfidentialityImpact"`
	ModifiedSubIntegrityImpact        string  `json:"modifiedSubIntegrityImpact"`
	ModifiedSubAvailabilityImpact     string  `json:"modifiedSubAvailabilityImpact"`
	Safety                            string  `json:"Safety"`
	Automatable                       string  `json:"Automatable"`
	Recovery                          string  `json:"Recovery"`
	ValueDensity                      string  `json:"valueDensity"`
	VulnerabilityResponseEffort       string  `json:"vulnerabilityResponseEffort"`
	ProviderUrgency                   string  `json:"providerUrgency"`
}

type CvssMetricV3 struct {
	Source              string     `json:"source"`
	Type                string     `json:"type"`
	CvssData            CvssDataV3 `json:"cvssData"`
	ExploitabilityScore float64    `json:"exploitabilityScore"`
	ImpactScore         float64    `json:"impactScore"`
}

type CvssDataV3 struct {
	Version               string  `json:"version"`
	VectorString          string  `json:"vectorString"`
	BaseScore             float64 `json:"baseScore"`
	BaseSeverity          string  `json:"baseSeverity"`
	AttackVector          string  `json:"attackVector"`
	AttackComplexity      string  `json:"attackComplexity"`
	PrivilegesRequired    string  `json:"privilegesRequired"`
	UserInteraction       string  `json:"userInteraction"`
	Scope                 string  `json:"scope"`
	ConfidentialityImpact string  `json:"confidentialityImpact"`
	IntegrityImpact       string  `json:"integrityImpact"`
	AvailabilityImpact    string  `json:"availabilityImpact"`
}

type CvssMetricV2 struct {
	Source                  string     `json:"source"`
	Type                    string     `json:"type"`
	CvssData                CvssDataV2 `json:"cvssData"`
	BaseSeverity            string     `json:"baseSeverity"`
	ExploitabilityScore     float64    `json:"exploitabilityScore"`
	ImpactScore             float64    `json:"impactScore"`
	AcInsufInfo             bool       `json:"acInsufInfo"`
	ObtainAllPrivilege      bool       `json:"obtainAllPrivilege"`
	ObtainUserPrivilege     bool       `json:"obtainUserPrivilege"`
	ObtainOtherPrivilege    bool       `json:"obtainOtherPrivilege"`
	UserInteractionRequired bool       `json:"userInteractionRequired"`
}

type CvssDataV2 struct {
	Version               string  `json:"version"`
	VectorString          string  `json:"vectorString"`
	BaseScore             float64 `json:"baseScore"`
	AccessVector          string  `json:"accessVector"`
	AccessComplexity      string  `json:"accessComplexity"`
	Authentication        string  `json:"authentication"`
	ConfidentialityImpact string  `json:"confidentialityImpact"`
	IntegrityImpact       string  `json:"integrityImpact"`
	AvailabilityImpact    string  `json:"availabilityImpact"`
}

type Weakness struct {
	Source      string      `json:"source"`
	Type        string      `json:"type"`
	Description []LangValue `json:"description"`
}

type Configuration struct {
	Nodes []Node `json:"nodes"`
}

type Node struct {
	Operator string     `json:"operator"`
	Negate   bool       `json:"negate"`
	CpeMatch []CpeMatch `json:"cpeMatch"`
}

type CpeMatch struct {
	Vulnerable          bool   `json:"vulnerable"`
	Criteria            string `json:"criteria"`
	VersionEndIncluding string `json:"versionEndIncluding"`
	MatchCriteriaID     string `json:"matchCriteriaId"`
}

type Reference struct {
	URL    string   `json:"url"`
	Source string   `json:"source"`
	Tags   []string `json:"tags"`
}
