package headers

import (
	"gulnManagement/webScanner/internal/models"
	"log"
	"strconv"
	"strings"
)

func CheckHeader(rule models.Rule, values []string) []models.Finding {
	var findings []models.Finding

	for _, check := range rule.Checks {

		switch check.Type {

		case models.CheckPresent:
			if len(values) == 0 {
				newFinding, err := models.NewFinding(
					rule.ID,
					rule.Header,
					models.CheckPresent,
					rule.Severity[check.MissingStatement],
					rule.Statements[check.MissingStatement],
					values,
				)
				if err != nil {
					log.Println(err)
				}
				findings = append(findings, newFinding)
			}

		case models.CheckDirective:

			if len(values) == 0 {
				// Header missing entirely → directive checks don't apply
				continue
			}

			// Normalize directive name
			directiveName := strings.ToLower(check.Name)

			// Parse *all* header values
			directives := make(map[string]int)
			for _, v := range values {
				for k, val := range ParseDirectives(v) {
					directives[k] = val
				}
			}

			val, exists := directives[directiveName]

			if !exists {
				if !check.Optional && check.MissingStatement != "" {

					f, err := models.NewFinding(
						rule.ID,
						rule.Header,
						directiveName,
						rule.Severity[check.MissingStatement],
						rule.Statements[check.MissingStatement],
						values,
					)
					if err == nil {
						findings = append(findings, f)
					}
				}
				continue
			}

			if check.Min != nil && val < *check.Min && check.MisconfigStatement != "" {

				f, err := models.NewFinding(
					rule.ID,
					rule.Header,
					directiveName,
					rule.Severity[check.MisconfigStatement],
					rule.Statements[check.MisconfigStatement],
					values,
				)
				if err == nil {
					findings = append(findings, f)
				}
			}

		}
	}

	return findings
}

func ParseDirectives(headerValue string) map[string]int {
	directives := make(map[string]int)

	parts := strings.Split(headerValue, ";")
	for _, part := range parts {
		part = strings.TrimSpace(part)

		if part == "" {
			continue
		}

		// key=value or flag
		kv := strings.SplitN(part, "=", 2)
		key := strings.ToLower(strings.TrimSpace(kv[0]))

		if key == "" {
			continue
		}

		// flag directive (e.g. includeSubDomains)
		if len(kv) == 1 {
			directives[key] = 1
			continue
		}

		// key=value directive
		valueStr := strings.TrimSpace(kv[1])
		if valueStr == "" {
			continue
		}

		value, err := strconv.Atoi(valueStr)
		if err != nil {
			// Non-integer values are ignored for numeric checks
			continue
		}

		directives[key] = value
	}

	return directives
}
