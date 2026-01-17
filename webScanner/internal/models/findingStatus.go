package models

import (
	"fmt"
	"strings"
)

type Severity int

const (
	CRITICAL Severity = iota + 1
	HIGH
	MEDIUM
	LOW
	NONE
)

func (s Severity) String() string {
	return [...]string{"Critical", "High", "Medium", "Low", "None"}[s-1]
}

func (s Severity) EnumIndex() int {
	return int(s)
}

type Finding struct {
	RuleID    string
	Header    string
	Type      string
	Severity  Severity
	Statement string
	Evidence  []string
}

func NewFinding(
	ruleID, header, checkType, severityStr, statementStr string,
	evidence []string,
) (Finding, error) {

	severity, err := ParseSeverity(severityStr)
	if err != nil {
		return Finding{}, err
	}

	return Finding{
		RuleID:    ruleID,
		Header:    header,
		Type:      checkType,
		Severity:  severity,
		Statement: statementStr,
		Evidence:  evidence,
	}, nil
}

func ParseSeverity(s string) (Severity, error) {
	switch strings.ToLower(s) {
	case "critical":
		return CRITICAL, nil
	case "high":
		return HIGH, nil
	case "medium":
		return MEDIUM, nil
	case "low":
		return LOW, nil
	case "none":
		return NONE, nil
	default:
		return NONE, fmt.Errorf("unknown severity: %s", s)
	}
}
