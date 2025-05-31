package xsdparser

import (
	"testing"

	"github.com/suifei/xsd2code/pkg/types"
)

func TestConvertElementWithInlineComplexType(t *testing.T) {
	parser := NewXSDParser("", "", "test")

	// Create test element with inline complex type
	element := types.XSDElement{
		Name:      "testElement",
		Type:      "",
		MinOccurs: "1",
		MaxOccurs: "1",
		ComplexType: &types.XSDComplexType{
			Sequence: &types.XSDSequence{
				Elements: []types.XSDElement{
					{
						Name: "childElement",
						Type: "string",
					},
				},
			},
			Attributes: []types.XSDAttribute{
				{
					Name: "testAttr",
					Type: "string",
					Use:  "required",
				},
			},
		},
	}

	// Convert element
	field, err := parser.convertElement(element)
	if err != nil {
		t.Fatalf("Failed to convert element: %v", err)
	}

	// Verify field properties
	if field.Name != "TestElement" {
		t.Errorf("Expected field name 'TestElement', got '%s'", field.Name)
	}

	if field.Type != "TestElement" {
		t.Errorf("Expected field type 'TestElement', got '%s'", field.Type)
	}

	// Verify inline type was created
	if len(parser.goTypes) != 1 {
		t.Errorf("Expected 1 inline type to be created, got %d", len(parser.goTypes))
	}

	inlineType := parser.goTypes[0]
	if inlineType.Name != "TestElement" {
		t.Errorf("Expected inline type name 'TestElement', got '%s'", inlineType.Name)
	}

	// Should have 2 fields: 1 element + 1 attribute
	if len(inlineType.Fields) != 2 {
		t.Errorf("Expected 2 fields in inline type, got %d", len(inlineType.Fields))
	}
}

func TestConvertElementWithInlineSimpleType(t *testing.T) {
	parser := NewXSDParser("", "", "test")

	element := types.XSDElement{
		Name:      "testElement",
		Type:      "",
		MinOccurs: "1",
		MaxOccurs: "1",
		SimpleType: &types.XSDSimpleType{
			Name: "testSimpleType",
		},
	}

	field, err := parser.convertElement(element)
	if err != nil {
		t.Fatalf("Failed to convert element: %v", err)
	}

	// Should default to string for inline simple types
	if field.Type != "string" {
		t.Errorf("Expected field type 'string' for inline simple type, got '%s'", field.Type)
	}
}

func TestConvertElementOptionalArray(t *testing.T) {
	parser := NewXSDParser("", "", "test")

	element := types.XSDElement{
		Name:      "testElement",
		Type:      "string",
		MinOccurs: "0",
		MaxOccurs: "unbounded",
	}

	field, err := parser.convertElement(element)
	if err != nil {
		t.Fatalf("Failed to convert element: %v", err)
	}

	// Should be array type
	if field.Type != "[]string" {
		t.Errorf("Expected field type '[]string', got '%s'", field.Type)
	}

	if !field.IsArray {
		t.Error("Expected field to be marked as array")
	}
}

func TestConvertElementOptionalPointer(t *testing.T) {
	parser := NewXSDParser("", "", "test")

	element := types.XSDElement{
		Name:      "testElement",
		Type:      "string",
		MinOccurs: "0",
		MaxOccurs: "1",
	}

	field, err := parser.convertElement(element)
	if err != nil {
		t.Fatalf("Failed to convert element: %v", err)
	}

	// Should be pointer type
	if field.Type != "*string" {
		t.Errorf("Expected field type '*string', got '%s'", field.Type)
	}

	if !field.IsOptional {
		t.Error("Expected field to be marked as optional")
	}
}
