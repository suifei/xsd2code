package testpkg

// Generated validation functions

import (
	"fmt"
	"regexp"
	"strings"
	"time"
)

// Generated types for testing
type CollapsedStringType string
type ExactLengthCodeType string 
type PercentageType float64
type PreservedStringType string
type TestDocument struct {
	Version string
	Id      string
	Created *time.Time
}

// Validator interface for all generated types
type Validator interface {
	Validate() error
}

// Validate validates the CollapsedStringType struct
func (v *CollapsedStringType) Validate() error {
	return nil
}

// Validate validates the ExactLengthCodeType struct
func (v *ExactLengthCodeType) Validate() error {
	return nil
}

// Validate validates the PercentageType struct
func (v *PercentageType) Validate() error {
	return nil
}

// Validate validates the PreservedStringType struct
func (v *PreservedStringType) Validate() error {
	return nil
}

// Validate validates the TestDocument struct
func (v *TestDocument) Validate() error {
	if v.Version == "" {
		return fmt.Errorf("TestDocument.Version is required")
	}
	if v.Id == "" {
		return fmt.Errorf("TestDocument.Id is required")
	}
	if v.Created != nil {
		if err := validateDateTime(*v.Created); err != nil {
			return fmt.Errorf("TestDocument.Created: %v", err)
		}
	}
	return nil
}

// Helper validation functions

func validateDateTime(dt time.Time) error {
	if dt.IsZero() {
		return fmt.Errorf("invalid datetime")
	}
	return nil
}

func validatePattern(value, pattern string) error {
	matched, err := regexp.MatchString(pattern, value)
	if err != nil {
		return fmt.Errorf("invalid pattern: %v", err)
	}
	if !matched {
		return fmt.Errorf("value does not match pattern %s", pattern)
	}
	return nil
}

func validateIntRange(value, min, max int) error {
	if value < min || value > max {
		return fmt.Errorf("value %d is out of range [%d, %d]", value, min, max)
	}
	return nil
}

// applyWhiteSpaceProcessing applies XSD whiteSpace facet processing
func applyWhiteSpaceProcessing(value, whiteSpaceAction string) string {
	switch whiteSpaceAction {
	case "replace":
		// Replace tab, newline, and carriage return with space
		value = strings.ReplaceAll(value, "\t", " ")
		value = strings.ReplaceAll(value, "\n", " ")
		value = strings.ReplaceAll(value, "\r", " ")
		return value
	case "collapse":
		// First apply replace processing
		value = strings.ReplaceAll(value, "\t", " ")
		value = strings.ReplaceAll(value, "\n", " ")
		value = strings.ReplaceAll(value, "\r", " ")
		// Then collapse sequences of spaces and trim
		value = regexp.MustCompile(`\s+`).ReplaceAllString(value, " ")
		value = strings.TrimSpace(value)
		return value
	case "preserve":
		fallthrough
	default:
		// Preserve all whitespace as-is
		return value
	}
}

func validateFixedValue(value, expectedValue string) error {
	if value != expectedValue {
		return fmt.Errorf("value '%s' does not match fixed value '%s'", value, expectedValue)
	}
	return nil
}
