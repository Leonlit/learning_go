package loader

import (
	"errors"
	"fmt"
	"gulnManagement/webScanner/internal/models"
	"os"
	"path/filepath"
	"strings"

	"gopkg.in/yaml.v3"
)

func LoadRulesFromYAML(dir string) ([]models.Rule, error) {
	var rules []models.Rule

	entries, err := os.ReadDir(dir)
	if err != nil {
		return nil, err
	}

	for _, entry := range entries {
		if entry.IsDir() {
			continue
		}

		if !strings.HasSuffix(entry.Name(), ".yaml") {
			continue
		}

		path := filepath.Join(dir, entry.Name())
		data, err := os.ReadFile(path)
		if err != nil {
			return nil, fmt.Errorf("read %s: %w", path, err)
		}

		var rule models.Rule
		if err := yaml.Unmarshal(data, &rule); err != nil {
			return nil, fmt.Errorf("parse %s: %w", path, err)
		}

		if err := validateRule(rule, path); err != nil {
			return nil, err
		}

		rules = append(rules, rule)
	}

	if len(rules) == 0 {
		return nil, errors.New("no rules loaded")
	}

	return rules, nil
}

func validateRule(r models.Rule, path string) error {
	if r.ID == "" {
		return fmt.Errorf("%s: missing id", path)
	}
	if r.Header == "" {
		return fmt.Errorf("%s: missing header", path)
	}
	if len(r.Checks) == 0 {
		return fmt.Errorf("%s: no checks defined", path)
	}
	return nil
}
