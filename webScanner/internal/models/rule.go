package models

type Rule struct {
	ID          string            `yaml:"id"`
	Header      string            `yaml:"header"`
	Category    string            `yaml:"category"`
	Required    bool              `yaml:"required"`
	Severity    map[string]string `yaml:"severity"`
	Description string            `yaml:"description"`
	Statements  map[string]string `yaml:"statements"`
	Checks      []Check           `yaml:"checks"`
	Remediation map[string]string `yaml:"remediation"`
	References  []string          `yaml:"references"`
}

type Check struct {
	Type     string `yaml:"type"`
	Name     string `yaml:"name,omitempty"`
	Min      *int   `yaml:"min,omitempty"`
	Optional bool   `yaml:"optional,omitempty"`

	MissingStatement   string `yaml:"missing-statement,omitempty"`
	MisconfigStatement string `yaml:"misconfig-statement,omitempty"`
}

const (
	CheckPresent   = "present"
	CheckDirective = "directive"
)

type Remediation struct {
	When   string `yaml:"when"`
	Action string `yaml:"action"`
}
