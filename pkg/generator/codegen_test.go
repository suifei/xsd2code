package generator

import (
	"strings"
	"testing"

	"github.com/suifei/xsd2code/pkg/types"
)

func TestGenerateValidationCode(t *testing.T) {
	generator := NewCodeGenerator("test", "")
	generator.SetJSONCompatible(true)
	generator.SetIncludeComments(true)
	// Add test types
	generator.SetGoTypes([]types.GoType{
		{
			Name: "TestType",
			Fields: []types.GoField{
				{
					Name:       "RequiredField",
					Type:       "string",
					IsOptional: false,
				},
				{
					Name:       "OptionalField",
					Type:       "*string",
					IsOptional: true,
				},
				{
					Name:      "ArrayField",
					Type:      "[]string",
					IsArray:   true,
					MinOccurs: 1,
					MaxOccurs: 5,
				},
			},
		},
	})

	code := generator.GenerateValidationCode()

	// Check that code contains expected elements
	expectedContents := []string{
		"type Validator interface",
		"Validate() error",
		"func (v *TestType) Validate() error",
		"RequiredField is required",
		"must have at least 1 elements",
		"must have at most 5 elements",
		"func validateDateTime",
		"func validatePattern",
		"func validateIntRange",
	}

	for _, expected := range expectedContents {
		if !strings.Contains(code, expected) {
			t.Errorf("Generated validation code does not contain expected content: %s", expected)
		}
	}
}

func TestGenerateTestCode(t *testing.T) {
	generator := NewCodeGenerator("test", "")

	generator.SetGoTypes([]types.GoType{
		{
			Name: "TestType",
			Fields: []types.GoField{
				{
					Name: "StringField",
					Type: "string",
				},
				{
					Name: "IntField",
					Type: "int",
				},
				{
					Name:       "OptionalField",
					Type:       "*string",
					IsOptional: true,
				}},
		},
	})

	code := generator.GenerateTestCode()

	expectedContents := []string{
		"func TestTestTypeXMLMarshaling(t *testing.T)",
		"xml.Marshal(original)",
		"xml.Unmarshal(xmlData, &unmarshaled)",
		"original.Validate()",
		"func TestTestTypeValidation(t *testing.T)",
		"func BenchmarkTestTypeMarshaling(b *testing.B)",
		"func stringPtr(s string) *string",
		"func intPtr(i int) *int",
	}

	for _, expected := range expectedContents {
		if !strings.Contains(code, expected) {
			t.Errorf("Generated test code does not contain expected content: %s", expected)
		}
	}
}

func TestGenerateSampleValue(t *testing.T) {
	generator := NewCodeGenerator("test", "")
	var builder strings.Builder

	testCases := []struct {
		typeName string
		expected string
	}{
		{"string", "\"test_value\""},
		{"int", "42"},
		{"float64", "3.14"},
		{"bool", "true"},
		{"time.Time", "time.Now()"},
		{"CustomType", "CustomType{}"},
	}

	for _, tc := range testCases {
		builder.Reset()
		generator.generateSampleValue(&builder, tc.typeName)
		result := builder.String()

		if result != tc.expected {
			t.Errorf("generateSampleValue(%s) = %s, expected %s", tc.typeName, result, tc.expected)
		}
	}
}

func TestGeneratePointerValue(t *testing.T) {
	generator := NewCodeGenerator("test", "")
	var builder strings.Builder

	testCases := []struct {
		typeName string
		expected string
	}{
		{"string", "stringPtr(\"test_value\")"},
		{"int", "intPtr(42)"},
		{"float64", "floatPtr(3.14)"},
		{"bool", "boolPtr(true)"},
		{"time.Time", "timePtr(time.Now())"},
		{"CustomType", "&CustomType{}"},
	}

	for _, tc := range testCases {
		builder.Reset()
		generator.generatePointerValue(&builder, tc.typeName)
		result := builder.String()

		if result != tc.expected {
			t.Errorf("generatePointerValue(%s) = %s, expected %s", tc.typeName, result, tc.expected)
		}
	}
}
