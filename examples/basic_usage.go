package main

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"log"

	"github.com/suifei/xsd2code/pkg/xsdparser"
)

func main() {
	// Example 1: Basic XSD parsing
	fmt.Println("=== Example 1: Basic XSD Parsing ===")
	basicExample()

	// Example 2: JSON compatibility
	fmt.Println("\n=== Example 2: JSON Compatibility ===")
	jsonExample()

	// Example 3: Debug mode
	fmt.Println("\n=== Example 3: Debug Mode ===")
	debugExample()
}

func basicExample() {
	// Parse XSD file
	parser := xsdparser.NewUnifiedXSDParser("../test/TC6_XML_V10_B.xsd", "output/basic.go", "plcopen")

	if err := parser.Parse(); err != nil {
		log.Printf("Failed to parse XSD: %v", err)
		return
	}

	if err := parser.GenerateGoCode(); err != nil {
		log.Printf("Failed to generate Go code: %v", err)
		return
	}

	fmt.Println("Successfully generated basic Go structs")
}

func jsonExample() {
	// Parse XSD file with JSON compatibility
	parser := xsdparser.NewUnifiedXSDParser("../test/TC6_XML_V10_B.xsd", "output/json.go", "plcopen")
	parser.SetJSONCompatible(true)

	if err := parser.Parse(); err != nil {
		log.Printf("Failed to parse XSD: %v", err)
		return
	}

	if err := parser.GenerateGoCode(); err != nil {
		log.Printf("Failed to generate Go code: %v", err)
		return
	}

	fmt.Println("Successfully generated JSON-compatible Go structs")
}

func debugExample() {
	// Parse XSD file with debug mode
	parser := xsdparser.NewUnifiedXSDParser("../test/TC6_XML_V10_B.xsd", "output/debug.go", "plcopen")
	parser.SetDebugMode(true)
	parser.SetJSONCompatible(true)

	if err := parser.Parse(); err != nil {
		log.Printf("Failed to parse XSD: %v", err)
		return
	}

	// Get parsed types for inspection
	goTypes := parser.GetGoTypes()
	fmt.Printf("Parsed %d Go types:\n", len(goTypes))
	for _, goType := range goTypes {
		fmt.Printf("- %s (%d fields)\n", goType.Name, len(goType.Fields))
	}

	if err := parser.GenerateGoCode(); err != nil {
		log.Printf("Failed to generate Go code: %v", err)
		return
	}

	fmt.Println("Successfully generated Go structs with debug output")
}

// Example of using generated structs (this would be in a separate file after generation)
func usageExample() {
	// This is an example of how the generated structs would be used

	// Create a struct instance (example structure)
	type Project struct {
		XMLName    xml.Name `xml:"http://www.plcopen.org/xml/tc6.xsd project" json:"-"`
		FileHeader struct {
			CompanyName      string `xml:"companyName,attr" json:"company_name"`
			ProductName      string `xml:"productName,attr" json:"product_name"`
			ProductVersion   string `xml:"productVersion,attr" json:"product_version"`
			CreationDateTime string `xml:"creationDateTime,attr" json:"creation_date_time"`
		} `xml:"fileHeader" json:"file_header"`
	}

	// Example XML data
	xmlData := `<?xml version="1.0" encoding="UTF-8"?>
	<project xmlns="http://www.plcopen.org/xml/tc6.xsd">
		<fileHeader companyName="Example Corp" productName="Test Product" 
		           productVersion="1.0" creationDateTime="2024-01-01T00:00:00"/>
	</project>`

	// Parse XML
	var project Project
	if err := xml.Unmarshal([]byte(xmlData), &project); err != nil {
		log.Printf("Failed to unmarshal XML: %v", err)
		return
	}

	fmt.Printf("Parsed project: %+v\n", project)

	// Convert to JSON
	jsonData, err := json.MarshalIndent(project, "", "  ")
	if err != nil {
		log.Printf("Failed to marshal JSON: %v", err)
		return
	}

	fmt.Printf("JSON representation:\n%s\n", jsonData)
}
