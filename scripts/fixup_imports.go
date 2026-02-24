//go:build ignore
// +build ignore

package main

import (
	"bytes"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	// Define the replacements
	replacements := []struct {
		old string
		new string
	}{
		// imports
		{`"time"`, `"github.com/moov-io/fednow20022/pkg/fednow"`},
		// Date
		{`type ISODate time.Time`, ``},
		{` ISODate `, ` fednow.ISODate `},
		{` *ISODate `, ` *fednow.ISODate `},
		// DateTime
		{`type ISODateTime time.Time`, ``},
		{` ISODateTime `, ` fednow.ISODateTime `},
		{` *ISODateTime `, ` *fednow.ISODateTime `},

		// admi.002
		{"Admi00200101 Admi00200101 `xml:\",any\"`", "Admi00200101 Admi00200101 `xml:\"admi.002.001.01\"`"},
	}

	// Walk through the directory
	err := filepath.WalkDir("gen", func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		// Only process regular files with .go extension
		if !d.IsDir() && strings.HasSuffix(strings.ToLower(path), ".go") {
			if err := replaceInFile(path, replacements); err != nil {
				fmt.Printf("Error processing %s: %v\n", path, err)
			}
		}
		return nil
	})

	if err != nil {
		fmt.Printf("Error walking directory: %v\n", err)
		os.Exit(1)
	}

	// Define xmlns for each schema
	type schemaInfo struct {
		ns    string
		field string
		tag   string
	}
	xmlnsMap := map[string]string{
		"gen/admi_002_001_01/models.go":                           "urn:iso:std:iso:20022:tech:xsd:admi.002.001.01",
		"gen/admi_004_001_02/models.go":                           "urn:iso:std:iso:20022:tech:xsd:admi.004.001.02",
		"gen/admi_006_001_01/models.go":                           "urn:iso:std:iso:20022:tech:xsd:admi.006.001.01",
		"gen/admi_007_001_01/models.go":                           "urn:iso:std:iso:20022:tech:xsd:admi.007.001.01",
		"gen/admi_011_001_01/models.go":                           "urn:iso:std:iso:20022:tech:xsd:admi.011.001.01",
		"gen/admi_998_001_02/models.go":                           "urn:swift:xsd:admi.998.001.02",
		"gen/sup_FedNowParticipantFile_admi_998_001_02/models.go": "urn:fed:xsd:admi.998.001.02",
		"gen/pacs_002_001_10/models.go":                           "urn:iso:std:iso:20022:tech:xsd:pacs.002.001.10",
		"gen/pacs_004_001_10/models.go":                           "urn:iso:std:iso:20022:tech:xsd:pacs.004.001.10",
		"gen/pacs_008_001_08/models.go":                           "urn:iso:std:iso:20022:tech:xsd:pacs.008.001.08",
		"gen/pacs_009_001_08/models.go":                           "urn:iso:std:iso:20022:tech:xsd:pacs.009.001.08",
		"gen/pacs_028_001_03/models.go":                           "urn:iso:std:iso:20022:tech:xsd:pacs.028.001.03",
		"gen/pain_013_001_07/models.go":                           "urn:iso:std:iso:20022:tech:xsd:pain.013.001.07",
		"gen/pain_014_001_07/models.go":                           "urn:iso:std:iso:20022:tech:xsd:pain.014.001.07",
		"gen/camt_026_001_07/models.go":                           "urn:iso:std:iso:20022:tech:xsd:camt.026.001.07",
		"gen/camt_028_001_09/models.go":                           "urn:iso:std:iso:20022:tech:xsd:camt.028.001.09",
		"gen/camt_029_001_09/models.go":                           "urn:iso:std:iso:20022:tech:xsd:camt.029.001.09",
		"gen/camt_052_001_08/models.go":                           "urn:iso:std:iso:20022:tech:xsd:camt.052.001.08",
		"gen/camt_054_001_08/models.go":                           "urn:iso:std:iso:20022:tech:xsd:camt.054.001.08",
		"gen/camt_055_001_09/models.go":                           "urn:iso:std:iso:20022:tech:xsd:camt.055.001.09",
		"gen/camt_056_001_08/models.go":                           "urn:iso:std:iso:20022:tech:xsd:camt.056.001.08",
		"gen/camt_060_001_05/models.go":                           "urn:iso:std:iso:20022:tech:xsd:camt.060.001.05",
	}

	// Create marshal.go files for custom MarshalXML
	for path, ns := range xmlnsMap {
		// Extract schema from path, e.g., gen/admi_002_001_01/models.go -> admi_002_001_01
		parts := strings.Split(path, "/")
		if len(parts) < 3 {
			continue
		}
		schema := parts[1]
		modelsPath := filepath.Join("gen", schema, "models.go")

		// Read models.go to find the field name in Document
		data, err := os.ReadFile(modelsPath)
		if err != nil {
			fmt.Printf("Error reading %s: %v\n", modelsPath, err)
			continue
		}
		content := string(data)

		// Find the Document struct
		docStart := strings.Index(content, "type Document struct {")
		if docStart == -1 {
			fmt.Printf("No Document struct in %s\n", modelsPath)
			continue
		}
		docEnd := strings.Index(content[docStart:], "}")
		if docEnd == -1 {
			continue
		}
		docContent := content[docStart : docStart+docEnd+1]

		// Find the field, assuming one field besides XMLName
		lines := strings.Split(docContent, "\n")
		var fieldName, tag string
		for _, line := range lines {
			line = strings.TrimSpace(line)
			if strings.HasPrefix(line, "XMLName ") {
				continue
			}
			if strings.Contains(line, "`xml:") {
				// e.g., Admi00200101 Admi00200101 `xml:"admi.002.001.01"`
				parts := strings.Fields(line)
				if len(parts) >= 3 {
					fieldName = parts[0]
					tagPart := strings.Split(line, "`xml:\"")[1]
					if strings.Contains(tagPart, "\"") {
						tag = strings.Split(tagPart, "\"`")[0]
					} else {
						// for ,any
						tag = strings.Split(tagPart, "`")[0]
					}
				}
				break
			}
		}
		if fieldName == "" {
			fmt.Printf("No field found in Document for %s\n", modelsPath)
			continue
		}

		marshalPath := filepath.Join("gen", schema, "marshal.go")
		encodeInner := fmt.Sprintf("e.Encode(d.%s)", fieldName)
		if tag != ",any" {
			encodeInner = fmt.Sprintf("e.EncodeElement(d.%s, xml.StartElement{Name: xml.Name{Local: \"%s\"}})", fieldName, tag)
		}
		marshalContent := fmt.Sprintf(`// Code generated by fixup_imports.go; DO NOT EDIT.

package %s

import (
	"encoding/xml"
)

func (d Document) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Name = xml.Name{Local: "Document"}
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "xmlns"}, Value: "%s"})
	e.EncodeToken(start)
	%s
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}
`, schema, ns, encodeInner)
		if err := os.WriteFile(marshalPath, []byte(marshalContent), 0644); err != nil {
			fmt.Printf("Error writing %s: %v\n", marshalPath, err)
		} else {
			fmt.Printf("Created %s\n", marshalPath)
		}
	}

	fmt.Println("Text replacement completed.")
}

func replaceInFile(path string, replacements []struct{ old, new string }) error {
	// Read the file
	data, err := os.ReadFile(path)
	if err != nil {
		return err
	}

	// Perform replacements
	modified := data
	changed := false
	for _, r := range replacements {
		if bytes.Contains(modified, []byte(r.old)) {
			modified = bytes.ReplaceAll(modified, []byte(r.old), []byte(r.new))
			changed = true
		}
	}

	// Write back to the file only if changes were made
	if changed {
		err = os.WriteFile(path, modified, 0644)
		if err != nil {
			return err
		}
		fmt.Printf("Updated %s\n", path)
	}
	return nil
}
