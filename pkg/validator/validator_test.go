package validator

import (
	"path/filepath"
	"testing"

	"github.com/suifei/xsd2code/pkg/xsdparser"
)

func TestXSDValidator_BasicValidation(t *testing.T) {
	// Parse test schema
	schemaPath := filepath.Join("testdata", "simple.xsd")
	parser := xsdparser.NewXSDParser(schemaPath, "", "test")

	if err := parser.Parse(); err != nil {
		t.Fatalf("Failed to parse schema: %v", err)
	}

	schema := parser.GetSchema()
	validator := NewXSDValidator(schema)

	tests := []struct {
		name      string
		xmlFile   string
		expectErr bool
	}{
		{
			name:      "valid XML",
			xmlFile:   "valid.xml",
			expectErr: false,
		},
		{
			name:      "missing required attribute",
			xmlFile:   "missing_required_attr.xml",
			expectErr: true,
		},
		{
			name:      "missing required element",
			xmlFile:   "missing_required_element.xml",
			expectErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			xmlPath := filepath.Join("testdata", tt.xmlFile)
			err := validator.ValidateXML(xmlPath)

			if tt.expectErr && err == nil {
				t.Errorf("Expected error but got none")
			}
			if !tt.expectErr && err != nil {
				t.Errorf("Unexpected error: %v", err)
			}
		})
	}
}

func TestXSDValidator_ValidationReport(t *testing.T) {
	// Parse test schema
	schemaPath := filepath.Join("testdata", "simple.xsd")
	parser := xsdparser.NewXSDParser(schemaPath, "", "test")

	if err := parser.Parse(); err != nil {
		t.Fatalf("Failed to parse schema: %v", err)
	}

	schema := parser.GetSchema()
	validator := NewXSDValidator(schema)

	xmlPath := filepath.Join("testdata", "valid.xml")
	report, err := validator.GenerateValidationReport(xmlPath)

	if err != nil {
		t.Fatalf("Failed to generate validation report: %v", err)
	}

	if report == nil {
		t.Fatal("Report is nil")
	}

	if report.XMLPath != xmlPath {
		t.Errorf("Expected XMLPath %s, got %s", xmlPath, report.XMLPath)
	}

	// Test report string representation
	reportStr := report.String()
	if reportStr == "" {
		t.Error("Report string is empty")
	}
}

func TestValidationReport_String(t *testing.T) {
	report := &ValidationReport{
		XMLPath: "test.xml",
		IsValid: false,
		Errors: []ValidationError{
			{
				Message: "Missing required element",
				Line:    5,
				Column:  10,
				Element: "age",
			},
		},
		Warnings: []ValidationWarning{
			{
				Message: "Optional element missing",
				Line:    7,
				Column:  15,
				Element: "email",
			},
		},
	}

	str := report.String()

	if str == "" {
		t.Error("Report string is empty")
	}

	// Check if report contains expected information
	expectedContents := []string{
		"test.xml",
		"Valid: false",
		"Missing required element",
		"Optional element missing",
		"Line: 5, Column: 10",
		"Element: age",
	}

	for _, expected := range expectedContents {
		if !contains(str, expected) {
			t.Errorf("Report string missing expected content: %s", expected)
		}
	}
}

// Helper function to check if string contains substring
func contains(s, substr string) bool {
	return len(s) >= len(substr) && (s == substr ||
		len(s) > len(substr) && (s[:len(substr)] == substr ||
			s[len(s)-len(substr):] == substr ||
			indexOfSubstring(s, substr) >= 0))
}

func indexOfSubstring(s, substr string) int {
	for i := 0; i <= len(s)-len(substr); i++ {
		if s[i:i+len(substr)] == substr {
			return i
		}
	}
	return -1
}
