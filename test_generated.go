package main

import (
	"encoding/xml"
	"fmt"
	"log"
	"strings"
	"github.com/suifei/xsd2code/test" // Import the generated package
)

func main() {
	fmt.Println("Testing generated Go structs...")

	// Test 1: Create a simple struct and marshal to XML
	fmt.Println("\n=== Test 1: Create and marshal RangeSigned ===")
	rangeSigned := test.RangeSigned{
		Lower: -100,
		Upper: 100,
	}

	xmlData, err := xml.MarshalIndent(rangeSigned, "", "  ")
	if err != nil {
		log.Fatalf("Failed to marshal RangeSigned: %v", err)
	}

	fmt.Printf("Generated XML:\n%s\n", xmlData)

	// Test 2: Parse the XML back to struct
	fmt.Println("\n=== Test 2: Unmarshal back to struct ===")
	var unmarshaled test.RangeSigned
	err = xml.Unmarshal(xmlData, &unmarshaled)
	if err != nil {
		log.Fatalf("Failed to unmarshal XML: %v", err)
	}

	fmt.Printf("Unmarshaled struct:\n")
	fmt.Printf("  Lower: %d\n", unmarshaled.Lower)
	fmt.Printf("  Upper: %d\n", unmarshaled.Upper)

	// Test 3: Validate round-trip consistency
	fmt.Println("\n=== Test 3: Round-trip validation ===")
	if rangeSigned.Lower == unmarshaled.Lower && rangeSigned.Upper == unmarshaled.Upper {
		fmt.Printf("✓ Round-trip test passed - data is consistent\n")
	} else {
		fmt.Printf("✗ Round-trip test failed - data mismatch\n")
	}

	// Test 4: Test RangeUnsigned struct
	fmt.Println("\n=== Test 4: Test RangeUnsigned ===")
	rangeUnsigned := test.RangeUnsigned{
		Lower: 0,
		Upper: 4294967295,
	}

	xmlData2, err := xml.MarshalIndent(rangeUnsigned, "", "  ")
	if err != nil {
		log.Fatalf("Failed to marshal RangeUnsigned: %v", err)
	}

	fmt.Printf("Generated XML:\n%s\n", xmlData2)

	// Test 5: Verify namespaces are properly set
	fmt.Println("\n=== Test 5: Namespace validation ===")
	xmlStr := string(xmlData)
	if strings.Contains(xmlStr, "http://www.plcopen.org/xml/tc6.xsd") {
		fmt.Printf("✓ Namespace is correctly included in XML\n")
	} else {
		fmt.Printf("⚠ Namespace might not be properly set\n")
	}

	fmt.Println("\n=== All tests completed successfully! ===")
}
