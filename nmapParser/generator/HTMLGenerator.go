package generator

import (
    "html/template"
	"os"
	"nmapParser/parser"
)

// ParseXML reads an XML file and unmarshals it into a Document struct.
func HTMLGenerator(scan *parser.NmapRun) error{
    tmpl, err := template.ParseFiles("template.html")
	if err != nil {
		panic(err)
	}

	file, err := os.Create("../outputs/report.html")
	if err != nil {
		return err
	}
	defer file.Close()

	// Execute template with struct data and write to file
	err = tmpl.Execute(file, scan)
	if err != nil {
		return err
	}

	return nil
}