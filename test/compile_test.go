package plcopen

import (
	"testing"
)

// Test that the generated code compiles without duplicate type errors
func TestGeneratedCodeCompiles(t *testing.T) {
	// Create instances of the generated types to ensure they compile
	var project Project
	var dataType DataType
	var value Value
	// Test that we can create different scaling types
	var fbdScaling ProjectContentHeaderCoordinateInfoFbdScaling
	var ldScaling ProjectContentHeaderCoordinateInfoLdScaling
	var sfcScaling ProjectContentHeaderCoordinateInfoSfcScaling

	// Verify types are not nil (this ensures they exist and compile)
	if &project == nil || &dataType == nil || &value == nil {
		t.Error("Basic types should not be nil")
	}

	if &fbdScaling == nil || &ldScaling == nil || &sfcScaling == nil {
		t.Error("Context-specific scaling types should not be nil")
	}

	t.Log("All generated types compile successfully without duplicate type errors")
}
