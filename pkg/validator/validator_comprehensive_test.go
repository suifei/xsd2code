package validator

import (
	"encoding/xml"
	"testing"

	"github.com/suifei/xsd2code/pkg/types"
)

func TestValidateElementStructure(t *testing.T) {
	// Create test schema
	schema := &types.XSDSchema{
		Elements: []types.XSDElement{
			{
				Name: "testElement",
				ComplexType: &types.XSDComplexType{
					Sequence: &types.XSDSequence{
						Elements: []types.XSDElement{
							{
								Name:      "childElement",
								Type:      "string",
								MinOccurs: "1",
								MaxOccurs: "1",
							},
						},
					},
					Attributes: []types.XSDAttribute{
						{
							Name: "requiredAttr",
							Type: "string",
							Use:  "required",
						},
					},
				},
			},
		},
	}

	validator := NewXSDValidator(schema)

	testCases := []struct {
		name        string
		xmlContent  string
		expectError bool
		errorMsg    string
	}{
		{
			name: "valid XML with required attribute",
			xmlContent: `<testElement requiredAttr="value">
				<childElement>test</childElement>
			</testElement>`,
			expectError: false,
		},
		{
			name: "missing required attribute",
			xmlContent: `<testElement>
				<childElement>test</childElement>
			</testElement>`,
			expectError: true,
			errorMsg:    "required attribute",
		},
		{
			name: "missing required child element",
			xmlContent: `<testElement requiredAttr="value">
			</testElement>`,
			expectError: true,
			errorMsg:    "missing required element",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			err := validator.ValidateXMLContent([]byte(tc.xmlContent))

			if tc.expectError {
				if err == nil {
					t.Error("Expected validation error but got none")
				} else if tc.errorMsg != "" && !contains(err.Error(), tc.errorMsg) {
					t.Errorf("Expected error to contain '%s', got: %s", tc.errorMsg, err.Error())
				}
			} else {
				if err != nil {
					t.Errorf("Expected no validation error but got: %v", err)
				}
			}
		})
	}
}

func TestFindElementDefinition(t *testing.T) {
	schema := &types.XSDSchema{
		Elements: []types.XSDElement{
			{Name: "rootElement"},
			{Name: "anotherElement"},
		},
		ComplexTypes: []types.XSDComplexType{
			{
				Name: "TestComplexType",
				Sequence: &types.XSDSequence{
					Elements: []types.XSDElement{
						{Name: "nestedElement"},
					},
				},
			},
		},
	}

	validator := NewXSDValidator(schema)

	// Test finding top-level elements
	element := validator.findElementDefinition("rootElement")
	if element == nil {
		t.Error("Expected to find rootElement")
	}

	// Test finding nested elements
	element = validator.findElementDefinition("nestedElement")
	if element == nil {
		t.Error("Expected to find nestedElement in complex type")
	}

	// Test element not found
	element = validator.findElementDefinition("nonExistentElement")
	if element != nil {
		t.Error("Expected not to find nonExistentElement")
	}
}

func TestValidateSequence(t *testing.T) {
	validator := NewXSDValidator(&types.XSDSchema{})
	ctx := &ValidationContext{
		errors:   make([]ValidationError, 0),
		warnings: make([]ValidationWarning, 0),
		line:     1,
		column:   1,
	}

	sequence := &types.XSDSequence{
		Elements: []types.XSDElement{
			{Name: "first", MinOccurs: "1", MaxOccurs: "1"},
			{Name: "second", MinOccurs: "0", MaxOccurs: "1"},
			{Name: "third", MinOccurs: "1", MaxOccurs: "1"},
		},
	}

	// Test valid sequence
	children := []XMLElement{
		{XMLName: newXMLName("first")},
		{XMLName: newXMLName("third")}, // second is optional
	}

	err := validator.validateSequence(children, sequence, ctx)
	if err != nil {
		t.Errorf("Expected no error for valid sequence, got: %v", err)
	}

	// Test invalid sequence (missing required element)
	ctx.errors = make([]ValidationError, 0) // Reset errors
	children = []XMLElement{
		{XMLName: newXMLName("first")},
		// missing required "third" element
	}

	err = validator.validateSequence(children, sequence, ctx)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	if len(ctx.errors) == 0 {
		t.Error("Expected validation error for missing required element")
	}
}

func TestValidateChoice(t *testing.T) {
	validator := NewXSDValidator(&types.XSDSchema{})
	ctx := &ValidationContext{
		errors:   make([]ValidationError, 0),
		warnings: make([]ValidationWarning, 0),
		line:     1,
		column:   1,
	}

	choice := &types.XSDChoice{
		Elements: []types.XSDElement{
			{Name: "option1"},
			{Name: "option2"},
			{Name: "option3"},
		},
	}

	// Test valid choice
	children := []XMLElement{
		{XMLName: newXMLName("option2")},
	}

	err := validator.validateChoice(children, choice, ctx)
	if err != nil {
		t.Errorf("Expected no error for valid choice, got: %v", err)
	}

	// Test invalid choice (element not in choice)
	ctx.errors = make([]ValidationError, 0) // Reset errors
	children = []XMLElement{
		{XMLName: newXMLName("invalidOption")},
	}

	err = validator.validateChoice(children, choice, ctx)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	if len(ctx.errors) == 0 {
		t.Error("Expected validation error for invalid choice element")
	}
}

func TestValidateAll(t *testing.T) {
	validator := NewXSDValidator(&types.XSDSchema{})
	ctx := &ValidationContext{
		errors:   make([]ValidationError, 0),
		warnings: make([]ValidationWarning, 0),
		line:     1,
		column:   1,
	}

	all := &types.XSDAll{
		Elements: []types.XSDElement{
			{Name: "element1", MinOccurs: "1", MaxOccurs: "1"},
			{Name: "element2", MinOccurs: "0", MaxOccurs: "1"},
			{Name: "element3", MinOccurs: "1", MaxOccurs: "1"},
		},
	}

	// Test valid all (different order)
	children := []XMLElement{
		{XMLName: newXMLName("element3")},
		{XMLName: newXMLName("element1")},
		// element2 is optional
	}

	err := validator.validateAll(children, all, ctx)
	if err != nil {
		t.Errorf("Expected no error for valid all, got: %v", err)
	}

	// Test duplicate element in all
	ctx.errors = make([]ValidationError, 0) // Reset errors
	children = []XMLElement{
		{XMLName: newXMLName("element1")},
		{XMLName: newXMLName("element1")}, // duplicate
	}

	err = validator.validateAll(children, all, ctx)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	if len(ctx.errors) == 0 {
		t.Error("Expected validation error for duplicate element in all")
	}
}

func TestIsValidType(t *testing.T) {
	schema := &types.XSDSchema{
		ComplexTypes: []types.XSDComplexType{
			{Name: "CustomComplexType"},
		},
		SimpleTypes: []types.XSDSimpleType{
			{Name: "CustomSimpleType"},
		},
	}

	validator := NewXSDValidator(schema)

	testCases := []struct {
		typeName string
		expected bool
	}{
		{"string", true},
		{"xs:string", true},
		{"xsd:int", true},
		{"CustomComplexType", true},
		{"CustomSimpleType", true},
		{"NonExistentType", false},
	}

	for _, tc := range testCases {
		result := validator.isValidType(tc.typeName)
		if result != tc.expected {
			t.Errorf("isValidType(%s) = %v, expected %v", tc.typeName, result, tc.expected)
		}
	}
}

// Helper functions for tests

func newXMLName(local string) xml.Name {
	return xml.Name{Local: local}
}
