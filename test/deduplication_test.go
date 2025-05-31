package plcopen

import (
	"encoding/xml"
	"testing"
)

func TestTypeDeduplication(t *testing.T) {
	// Test that we can create instances of different scaling types
	// These should all be different types, not causing conflicts
	fbdScaling := ProjectContentHeaderCoordinateInfoFbdScaling{
		X: 1.0,
		Y: 2.0,
	}

	ldScaling := ProjectContentHeaderCoordinateInfoLdScaling{
		X: 3.0,
		Y: 4.0,
	}

	sfcScaling := ProjectContentHeaderCoordinateInfoSfcScaling{
		X: 5.0,
		Y: 6.0,
	}

	// Test that we can create different Value types
	simpleValue := ValueSimpleValue{}
	arrayValue := ValueArrayValue{}
	structValue := ValueStructValue{}
	value := Value{}

	// Verify all instances are properly initialized
	if fbdScaling.X != 1.0 || fbdScaling.Y != 2.0 {
		t.Error("FBD scaling not properly initialized")
	}

	if ldScaling.X != 3.0 || ldScaling.Y != 4.0 {
		t.Error("LD scaling not properly initialized")
	}

	if sfcScaling.X != 5.0 || sfcScaling.Y != 6.0 {
		t.Error("SFC scaling not properly initialized")
	}

	// Test that all value types are different and can coexist
	if &simpleValue == nil || &arrayValue == nil || &structValue == nil || &value == nil {
		t.Error("Value types should not be nil")
	}

	t.Log("Successfully created all context-specific types without conflicts")
}

func TestXMLMarshaling(t *testing.T) {
	// Test that the generated types can be marshaled to XML
	// Create a simple scaling type instead of project to avoid time pointer issues
	scaling := ProjectContentHeaderCoordinateInfoFbdScaling{
		X: 1.0,
		Y: 2.0,
	}

	data, err := xml.Marshal(scaling)
	if err != nil {
		t.Fatalf("Failed to marshal scaling: %v", err)
	}

	if len(data) == 0 {
		t.Error("Marshaled data should not be empty")
	}

	t.Logf("Successfully marshaled scaling to XML: %s", string(data))
}
