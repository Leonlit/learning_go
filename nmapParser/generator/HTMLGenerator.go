package generator

import (
	"fmt"
	"html/template"
	"nmapManagement/nmapParser/parser"
	"os"
	"path/filepath"
	"runtime"
)

type HostWithPortCount struct {
	parser.Host
	PortCount int
}

// ParseXML reads an XML file and unmarshals it into a Document struct.
func HTMLGenerator(scan *parser.NmapRun) error {

	_, filename, _, _ := runtime.Caller(0)
	dir := filepath.Dir(filename)

	templateDirPath := dir + "/template/"
	templatePath := templateDirPath + "all.tmpl"

	outputDirPath := dir + "/outputs/"
	outputPath := outputDirPath + "report_all.html"

	tmpl, err := template.ParseFiles(templatePath)
	if err != nil {
		panic(err)
	}

	var hostsWithPortCount []HostWithPortCount
	for _, host := range scan.TargetHosts {
		hostsWithPortCount = append(hostsWithPortCount, HostWithPortCount{
			Host:      host,
			PortCount: len(host.Ports),
		})
	}

	err = os.MkdirAll(outputDirPath, 0755)

	if err != nil {
		fmt.Println("Error creating directory")
		return err
	}

	file, err := os.Create(outputPath)
	if err != nil {
		return err
	}
	defer file.Close()

	// Execute template with struct data and write to file
	err = tmpl.Execute(file, struct {
		RunStats    parser.RunStats
		TargetHosts []HostWithPortCount
	}{
		RunStats:    scan.RunStats,
		TargetHosts: hostsWithPortCount,
	})
	if err != nil {
		return err
	}

	return nil
}
