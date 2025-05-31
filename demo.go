package main

import (
	"fmt"
	"log"
	"os"

	"xsd2code/pkg/generator"
	"xsd2code/pkg/validator"
	"xsd2code/pkg/xsdparser"
)

func main() {
	// Test with real-world XSD file
	xsdFile := "test/TC6_XML_V10_B.xsd"

	fmt.Printf("Testing with real-world XSD file: %s\n", xsdFile)

	// Parse the XSD file
	parser := xsdparser.NewParser()
	schema, err := parser.ParseFile(xsdFile)
	if err != nil {
		log.Fatalf("Failed to parse XSD file: %v", err)
	}

	fmt.Printf("✓ Successfully parsed XSD schema\n")
	fmt.Printf("  - Target namespace: %s\n", schema.TargetNamespace)
	fmt.Printf("  - Elements found: %d\n", len(schema.Elements))
	fmt.Printf("  - Complex types found: %d\n", len(schema.ComplexTypes))
	fmt.Printf("  - Simple types found: %d\n", len(schema.SimpleTypes))

	// Generate Go code
	codeGen := generator.NewCodeGenerator("github.com/test/plcopen", "plcopen")
	goCode, err := codeGen.GenerateGoCode(schema)
	if err != nil {
		log.Fatalf("Failed to generate Go code: %v", err)
	}

	fmt.Printf("✓ Successfully generated Go code (%d bytes)\n", len(goCode))

	// Write generated code to file
	outputFile := "test/generated_plcopen.go"
	err = os.WriteFile(outputFile, []byte(goCode), 0644)
	if err != nil {
		log.Fatalf("Failed to write generated code: %v", err)
	}

	fmt.Printf("✓ Generated code written to: %s\n", outputFile)

	// Test validation (optional - validate a subset of the schema)
	validator := validator.NewValidator()
	isValid, err := validator.ValidateXSDStructure(schema)
	if err != nil {
		fmt.Printf("⚠ Validation warning: %v\n", err)
	} else if isValid {
		fmt.Printf("✓ XSD structure validation passed\n")
	} else {
		fmt.Printf("⚠ XSD structure validation failed\n")
	}

	// Show some sample generated types
	fmt.Printf("\n=== Sample Generated Types ===\n")
	showSampleTypes(goCode)
}

func showSampleTypes(code string) {
	lines := splitLines(code)
	inType := false
	typeLines := []string{}
	typeCount := 0

	for _, line := range lines {
		if len(line) > 0 && line[0] != '\t' && line[0] != ' ' {
			if inType && len(typeLines) > 0 {
				// Print the accumulated type
				for _, typeLine := range typeLines {
					fmt.Println(typeLine)
				}
				fmt.Println()
				typeCount++
				if typeCount >= 3 { // Show only first 3 types
					break
				}
				typeLines = []string{}
			}
			inType = false
		}

		if len(line) > 5 && line[:5] == "type " {
			inType = true
			typeLines = []string{line}
		} else if inType {
			typeLines = append(typeLines, line)
			if line == "}" {
				inType = false
			}
		}
	}
}

func splitLines(s string) []string {
	var lines []string
	start := 0
	for i, c := range s {
		if c == '\n' {
			lines = append(lines, s[start:i])
			start = i + 1
		}
	}
	if start < len(s) {
		lines = append(lines, s[start:])
	}
	return lines
}
