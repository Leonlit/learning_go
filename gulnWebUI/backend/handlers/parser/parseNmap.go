package parser

import (
    "encoding/xml"
    "io"
)

// ParseXML reads an XML file and unmarshals it into a Document struct.
func ParseNmap(reader io.Reader) (*NmapRun, error) {
    var doc NmapRun
    decoder := xml.NewDecoder(reader)
    err := decoder.Decode(&doc)
    if err != nil {
        return nil, err
    }
    return &doc, nil
}