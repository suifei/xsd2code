package main

import (
	"fmt"
	"log"

	"github.com/suifei/xsd2code/pkg/generator"
	"github.com/suifei/xsd2code/pkg/types"
	"github.com/suifei/xsd2code/pkg/validator"
	"github.com/suifei/xsd2code/pkg/xsdparser"
)

func main() {
	// Test with real-world XSD file
	xsdFile := "test/TC6_XML_V10_B.xsd"

	fmt.Printf("Testing with real-world XSD file: %s\n", xsdFile) // Parse the XSD file
	parser := xsdparser.NewXSDParser(xsdFile, "", "test")
	parser.SetJSONCompatible(true) // Enable JSON compatibility
	parser.SetDebugMode(true)      // Enable debug mode
	err := parser.Parse()
	if err != nil {
		log.Fatalf("Failed to parse XSD file: %v", err)
	}

	schema := parser.GetSchema()

	fmt.Printf("✓ Successfully parsed XSD schema\n")
	fmt.Printf("  - Target namespace: %s\n", schema.TargetNamespace)
	fmt.Printf("  - Elements found: %d\n", len(schema.Elements))
	fmt.Printf("  - Complex types found: %d\n", len(schema.ComplexTypes))
	fmt.Printf("  - Simple types found: %d\n", len(schema.SimpleTypes))
	// Generate Go code
	codeGen := generator.NewCodeGenerator("plcopen", "test/generated_plcopen.go")
	codeGen.SetGoTypes(parser.GetGoTypes())
	codeGen.SetJSONCompatible(true)
	err = codeGen.Generate()
	if err != nil {
		log.Fatalf("Failed to generate Go code: %v", err)
	}

	fmt.Printf("✓ Successfully generated Go code file\n")
	// Test validation (optional - validate a subset of the schema)
	_ = validator.NewXSDValidator(schema) // Just create validator to show it works
	fmt.Printf("✓ XSD validator created successfully\n")
	// Show some sample Go types found
	fmt.Printf("\n=== Sample Go Types Found ===\n")
	showSampleGoTypes(parser.GetGoTypes())
}

func showSampleGoTypes(goTypes []types.GoType) {
	if len(goTypes) == 0 {
		fmt.Println("No Go types found")
		return
	}

	// Show first 3 types as examples
	maxTypes := 3
	if len(goTypes) < maxTypes {
		maxTypes = len(goTypes)
	}

	for i := 0; i < maxTypes; i++ {
		goType := goTypes[i]
		fmt.Printf("Type: %s\n", goType.Name)
		if goType.IsEnum {
			fmt.Printf("  Kind: enum (base: %s)\n", goType.BaseType)
			fmt.Printf("  Constants (%d):\n", len(goType.Constants))
			maxConstants := 3
			if len(goType.Constants) < maxConstants {
				maxConstants = len(goType.Constants)
			}
			for j := 0; j < maxConstants; j++ {
				constant := goType.Constants[j]
				fmt.Printf("    - %s = %s\n", constant.Name, constant.Value)
			}
			if len(goType.Constants) > maxConstants {
				fmt.Printf("    ... and %d more constants\n", len(goType.Constants)-maxConstants)
			}
		} else {
			fmt.Printf("  Kind: struct\n")
			if len(goType.Fields) > 0 {
				fmt.Printf("  Fields (%d):\n", len(goType.Fields))
				maxFields := 5
				if len(goType.Fields) < maxFields {
					maxFields = len(goType.Fields)
				}
				for j := 0; j < maxFields; j++ {
					field := goType.Fields[j]
					fmt.Printf("    - %s: %s", field.Name, field.Type)
					if field.XMLTag != "" {
						fmt.Printf(" (xml:\"%s\")", field.XMLTag)
					}
					if field.IsAttribute {
						fmt.Printf(" [attr]")
					}
					if field.IsOptional {
						fmt.Printf(" [optional]")
					}
					if field.IsArray {
						fmt.Printf(" [array]")
					}
					fmt.Println()
				}
				if len(goType.Fields) > maxFields {
					fmt.Printf("    ... and %d more fields\n", len(goType.Fields)-maxFields)
				}
			}
		}
		fmt.Println()
	}

	if len(goTypes) > maxTypes {
		fmt.Printf("... and %d more types\n", len(goTypes)-maxTypes)
	}
}
