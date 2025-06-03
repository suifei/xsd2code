package generated

// Generated test code

import (
	"encoding/xml"
	"testing"
	"time"
)

func TestDataTypeStringXMLMarshaling(t *testing.T) {
	original := &DataTypeString{
		Length: uintPtr(42),
	}

	// Test marshaling
	xmlData, err := xml.Marshal(original)
	if err != nil {
		t.Fatalf("Failed to marshal XML: %v", err)
	}

	// Test unmarshaling
	var unmarshaled DataTypeString
	err = xml.Unmarshal(xmlData, &unmarshaled)
	if err != nil {
		t.Fatalf("Failed to unmarshal XML: %v", err)
	}

	// Test validation
	if err := original.Validate(); err != nil {
		t.Errorf("Validation failed: %v", err)
	}
}

func TestDataTypeStringValidation(t *testing.T) {
	// Test valid case
	valid := &DataTypeString{
	}
	if err := valid.Validate(); err != nil {
		t.Errorf("Valid object should not have validation errors: %v", err)
	}

}

func TestDataTypeWstringXMLMarshaling(t *testing.T) {
	original := &DataTypeWstring{
		Length: uintPtr(42),
	}

	// Test marshaling
	xmlData, err := xml.Marshal(original)
	if err != nil {
		t.Fatalf("Failed to marshal XML: %v", err)
	}

	// Test unmarshaling
	var unmarshaled DataTypeWstring
	err = xml.Unmarshal(xmlData, &unmarshaled)
	if err != nil {
		t.Fatalf("Failed to unmarshal XML: %v", err)
	}

	// Test validation
	if err := original.Validate(); err != nil {
		t.Errorf("Validation failed: %v", err)
	}
}

func TestDataTypeWstringValidation(t *testing.T) {
	// Test valid case
	valid := &DataTypeWstring{
	}
	if err := valid.Validate(); err != nil {
		t.Errorf("Valid object should not have validation errors: %v", err)
	}

}

func TestDataTypeArrayXMLMarshaling(t *testing.T) {
	original := &DataTypeArray{
		Dimension: []RangeSigned{RangeSigned{}},
		BaseType: DataType{},
	}

	// Test marshaling
	xmlData, err := xml.Marshal(original)
	if err != nil {
		t.Fatalf("Failed to marshal XML: %v", err)
	}

	// Test unmarshaling
	var unmarshaled DataTypeArray
	err = xml.Unmarshal(xmlData, &unmarshaled)
	if err != nil {
		t.Fatalf("Failed to unmarshal XML: %v", err)
	}

	// Test validation
	if err := original.Validate(); err != nil {
		t.Errorf("Validation failed: %v", err)
	}
}

func TestDataTypeArrayValidation(t *testing.T) {
	// Test valid case
	valid := &DataTypeArray{
		Dimension: []RangeSigned{RangeSigned{}},
		BaseType: DataType{},
	}
	if err := valid.Validate(); err != nil {
		t.Errorf("Valid object should not have validation errors: %v", err)
	}

	// Test missing required field: Dimension
	invalidDimension := &DataTypeArray{}
	if err := invalidDimension.Validate(); err == nil {
		t.Error("Missing required field Dimension should cause validation error")
	}

	// Test missing required field: BaseType
	invalidBaseType := &DataTypeArray{}
	if err := invalidBaseType.Validate(); err == nil {
		t.Error("Missing required field BaseType should cause validation error")
	}

}

func TestDataTypeDerivedXMLMarshaling(t *testing.T) {
	original := &DataTypeDerived{
		Name: "test_value",
	}

	// Test marshaling
	xmlData, err := xml.Marshal(original)
	if err != nil {
		t.Fatalf("Failed to marshal XML: %v", err)
	}

	// Test unmarshaling
	var unmarshaled DataTypeDerived
	err = xml.Unmarshal(xmlData, &unmarshaled)
	if err != nil {
		t.Fatalf("Failed to unmarshal XML: %v", err)
	}

	// Test validation
	if err := original.Validate(); err != nil {
		t.Errorf("Validation failed: %v", err)
	}
}

func TestDataTypeDerivedValidation(t *testing.T) {
	// Test valid case
	valid := &DataTypeDerived{
		Name: "test_value",
	}
	if err := valid.Validate(); err != nil {
		t.Errorf("Valid object should not have validation errors: %v", err)
	}

	// Test missing required field: Name
	invalidName := &DataTypeDerived{}
	if err := invalidName.Validate(); err == nil {
		t.Error("Missing required field Name should cause validation error")
	}

}

func TestDataTypeEnumValuesValueXMLMarshaling(t *testing.T) {
	original := &DataTypeEnumValuesValue{
		Name: "test_value",
		Value: stringPtr("test_value"),
	}

	// Test marshaling
	xmlData, err := xml.Marshal(original)
	if err != nil {
		t.Fatalf("Failed to marshal XML: %v", err)
	}

	// Test unmarshaling
	var unmarshaled DataTypeEnumValuesValue
	err = xml.Unmarshal(xmlData, &unmarshaled)
	if err != nil {
		t.Fatalf("Failed to unmarshal XML: %v", err)
	}

	// Test validation
	if err := original.Validate(); err != nil {
		t.Errorf("Validation failed: %v", err)
	}
}

func TestDataTypeEnumValuesValueValidation(t *testing.T) {
	// Test valid case
	valid := &DataTypeEnumValuesValue{
		Name: "test_value",
	}
	if err := valid.Validate(); err != nil {
		t.Errorf("Valid object should not have validation errors: %v", err)
	}

	// Test missing required field: Name
	invalidName := &DataTypeEnumValuesValue{}
	if err := invalidName.Validate(); err == nil {
		t.Error("Missing required field Name should cause validation error")
	}

}

func TestDataTypeEnumValuesXMLMarshaling(t *testing.T) {
	original := &DataTypeEnumValues{
		Value: DataTypeEnumValuesValue{},
	}

	// Test marshaling
	xmlData, err := xml.Marshal(original)
	if err != nil {
		t.Fatalf("Failed to marshal XML: %v", err)
	}

	// Test unmarshaling
	var unmarshaled DataTypeEnumValues
	err = xml.Unmarshal(xmlData, &unmarshaled)
	if err != nil {
		t.Fatalf("Failed to unmarshal XML: %v", err)
	}

	// Test validation
	if err := original.Validate(); err != nil {
		t.Errorf("Validation failed: %v", err)
	}
}

func TestDataTypeEnumValuesValidation(t *testing.T) {
	// Test valid case
	valid := &DataTypeEnumValues{
		Value: DataTypeEnumValuesValue{},
	}
	if err := valid.Validate(); err != nil {
		t.Errorf("Valid object should not have validation errors: %v", err)
	}

	// Test missing required field: Value
	invalidValue := &DataTypeEnumValues{}
	if err := invalidValue.Validate(); err == nil {
		t.Error("Missing required field Value should cause validation error")
	}

}

func TestDataTypeEnumXMLMarshaling(t *testing.T) {
	original := &DataTypeEnum{
		Values: DataTypeEnumValues{},
		BaseType: &DataType{},
	}

	// Test marshaling
	xmlData, err := xml.Marshal(original)
	if err != nil {
		t.Fatalf("Failed to marshal XML: %v", err)
	}

	// Test unmarshaling
	var unmarshaled DataTypeEnum
	err = xml.Unmarshal(xmlData, &unmarshaled)
	if err != nil {
		t.Fatalf("Failed to unmarshal XML: %v", err)
	}

	// Test validation
	if err := original.Validate(); err != nil {
		t.Errorf("Validation failed: %v", err)
	}
}

func TestDataTypeEnumValidation(t *testing.T) {
	// Test valid case
	valid := &DataTypeEnum{
		Values: DataTypeEnumValues{},
	}
	if err := valid.Validate(); err != nil {
		t.Errorf("Valid object should not have validation errors: %v", err)
	}

	// Test missing required field: Values
	invalidValues := &DataTypeEnum{}
	if err := invalidValues.Validate(); err == nil {
		t.Error("Missing required field Values should cause validation error")
	}

}

func TestDataTypeSubrangeSignedXMLMarshaling(t *testing.T) {
	original := &DataTypeSubrangeSigned{
		Range: RangeSigned{},
		BaseType: DataType{},
	}

	// Test marshaling
	xmlData, err := xml.Marshal(original)
	if err != nil {
		t.Fatalf("Failed to marshal XML: %v", err)
	}

	// Test unmarshaling
	var unmarshaled DataTypeSubrangeSigned
	err = xml.Unmarshal(xmlData, &unmarshaled)
	if err != nil {
		t.Fatalf("Failed to unmarshal XML: %v", err)
	}

	// Test validation
	if err := original.Validate(); err != nil {
		t.Errorf("Validation failed: %v", err)
	}
}

func TestDataTypeSubrangeSignedValidation(t *testing.T) {
	// Test valid case
	valid := &DataTypeSubrangeSigned{
		Range: RangeSigned{},
		BaseType: DataType{},
	}
	if err := valid.Validate(); err != nil {
		t.Errorf("Valid object should not have validation errors: %v", err)
	}

	// Test missing required field: Range
	invalidRange := &DataTypeSubrangeSigned{}
	if err := invalidRange.Validate(); err == nil {
		t.Error("Missing required field Range should cause validation error")
	}

	// Test missing required field: BaseType
	invalidBaseType := &DataTypeSubrangeSigned{}
	if err := invalidBaseType.Validate(); err == nil {
		t.Error("Missing required field BaseType should cause validation error")
	}

}

func TestDataTypeSubrangeUnsignedXMLMarshaling(t *testing.T) {
	original := &DataTypeSubrangeUnsigned{
		Range: RangeUnsigned{},
		BaseType: DataType{},
	}

	// Test marshaling
	xmlData, err := xml.Marshal(original)
	if err != nil {
		t.Fatalf("Failed to marshal XML: %v", err)
	}

	// Test unmarshaling
	var unmarshaled DataTypeSubrangeUnsigned
	err = xml.Unmarshal(xmlData, &unmarshaled)
	if err != nil {
		t.Fatalf("Failed to unmarshal XML: %v", err)
	}

	// Test validation
	if err := original.Validate(); err != nil {
		t.Errorf("Validation failed: %v", err)
	}
}

func TestDataTypeSubrangeUnsignedValidation(t *testing.T) {
	// Test valid case
	valid := &DataTypeSubrangeUnsigned{
		Range: RangeUnsigned{},
		BaseType: DataType{},
	}
	if err := valid.Validate(); err != nil {
		t.Errorf("Valid object should not have validation errors: %v", err)
	}

	// Test missing required field: Range
	invalidRange := &DataTypeSubrangeUnsigned{}
	if err := invalidRange.Validate(); err == nil {
		t.Error("Missing required field Range should cause validation error")
	}

	// Test missing required field: BaseType
	invalidBaseType := &DataTypeSubrangeUnsigned{}
	if err := invalidBaseType.Validate(); err == nil {
		t.Error("Missing required field BaseType should cause validation error")
	}

}

func TestDataTypePointerXMLMarshaling(t *testing.T) {
	original := &DataTypePointer{
		BaseType: DataType{},
	}

	// Test marshaling
	xmlData, err := xml.Marshal(original)
	if err != nil {
		t.Fatalf("Failed to marshal XML: %v", err)
	}

	// Test unmarshaling
	var unmarshaled DataTypePointer
	err = xml.Unmarshal(xmlData, &unmarshaled)
	if err != nil {
		t.Fatalf("Failed to unmarshal XML: %v", err)
	}

	// Test validation
	if err := original.Validate(); err != nil {
		t.Errorf("Validation failed: %v", err)
	}
}

func TestDataTypePointerValidation(t *testing.T) {
	// Test valid case
	valid := &DataTypePointer{
		BaseType: DataType{},
	}
	if err := valid.Validate(); err != nil {
		t.Errorf("Valid object should not have validation errors: %v", err)
	}

	// Test missing required field: BaseType
	invalidBaseType := &DataTypePointer{}
	if err := invalidBaseType.Validate(); err == nil {
		t.Error("Missing required field BaseType should cause validation error")
	}

}

func TestDataTypeXMLMarshaling(t *testing.T) {
	original := &DataType{
		BOOL: true,
		BYTE: uint8(42),
		WORD: uint16(42),
		DWORD: 42,
		LWORD: 42,
		SINT: int8(42),
		INT: int16(42),
		DINT: 42,
		LINT: 42,
		USINT: uint8(42),
		UINT: uint16(42),
		UDINT: 42,
		ULINT: 42,
		REAL: 3.14,
		LREAL: 3.14,
		TIME: durationPtr(time.Second),
		DATE: timePtr(time.Now()),
		DT: timePtr(time.Now()),
		TOD: timePtr(time.Now()),
		String: &DataTypeString{},
		Wstring: &DataTypeWstring{},
		Array: &DataTypeArray{},
		Derived: &DataTypeDerived{},
		Enum: &DataTypeEnum{},
		Struct: &VarListPlain{},
		SubrangeSigned: &DataTypeSubrangeSigned{},
		SubrangeUnsigned: &DataTypeSubrangeUnsigned{},
		Pointer: &DataTypePointer{},
	}

	// Test marshaling
	xmlData, err := xml.Marshal(original)
	if err != nil {
		t.Fatalf("Failed to marshal XML: %v", err)
	}

	// Test unmarshaling
	var unmarshaled DataType
	err = xml.Unmarshal(xmlData, &unmarshaled)
	if err != nil {
		t.Fatalf("Failed to unmarshal XML: %v", err)
	}

	// Test validation
	if err := original.Validate(); err != nil {
		t.Errorf("Validation failed: %v", err)
	}
}

func TestDataTypeValidation(t *testing.T) {
	// Test valid case
	valid := &DataType{
	}
	if err := valid.Validate(); err != nil {
		t.Errorf("Valid object should not have validation errors: %v", err)
	}

}

func TestRangeSignedXMLMarshaling(t *testing.T) {
	original := &RangeSigned{
		Lower: 42,
		Upper: 42,
	}

	// Test marshaling
	xmlData, err := xml.Marshal(original)
	if err != nil {
		t.Fatalf("Failed to marshal XML: %v", err)
	}

	// Test unmarshaling
	var unmarshaled RangeSigned
	err = xml.Unmarshal(xmlData, &unmarshaled)
	if err != nil {
		t.Fatalf("Failed to unmarshal XML: %v", err)
	}

	// Test validation
	if err := original.Validate(); err != nil {
		t.Errorf("Validation failed: %v", err)
	}
}

func TestRangeSignedValidation(t *testing.T) {
	// Test valid case
	valid := &RangeSigned{
		Lower: 42,
		Upper: 42,
	}
	if err := valid.Validate(); err != nil {
		t.Errorf("Valid object should not have validation errors: %v", err)
	}

	// Test missing required field: Lower
	invalidLower := &RangeSigned{}
	if err := invalidLower.Validate(); err == nil {
		t.Error("Missing required field Lower should cause validation error")
	}

	// Test missing required field: Upper
	invalidUpper := &RangeSigned{}
	if err := invalidUpper.Validate(); err == nil {
		t.Error("Missing required field Upper should cause validation error")
	}

}

func TestRangeUnsignedXMLMarshaling(t *testing.T) {
	original := &RangeUnsigned{
		Lower: 42,
		Upper: 42,
	}

	// Test marshaling
	xmlData, err := xml.Marshal(original)
	if err != nil {
		t.Fatalf("Failed to marshal XML: %v", err)
	}

	// Test unmarshaling
	var unmarshaled RangeUnsigned
	err = xml.Unmarshal(xmlData, &unmarshaled)
	if err != nil {
		t.Fatalf("Failed to unmarshal XML: %v", err)
	}

	// Test validation
	if err := original.Validate(); err != nil {
		t.Errorf("Validation failed: %v", err)
	}
}

func TestRangeUnsignedValidation(t *testing.T) {
	// Test valid case
	valid := &RangeUnsigned{
		Lower: 42,
		Upper: 42,
	}
	if err := valid.Validate(); err != nil {
		t.Errorf("Valid object should not have validation errors: %v", err)
	}

	// Test missing required field: Lower
	invalidLower := &RangeUnsigned{}
	if err := invalidLower.Validate(); err == nil {
		t.Error("Missing required field Lower should cause validation error")
	}

	// Test missing required field: Upper
	invalidUpper := &RangeUnsigned{}
	if err := invalidUpper.Validate(); err == nil {
		t.Error("Missing required field Upper should cause validation error")
	}

}

func TestValueSimpleValueXMLMarshaling(t *testing.T) {
	original := &ValueSimpleValue{
		Value: stringPtr("test_value"),
	}

	// Test marshaling
	xmlData, err := xml.Marshal(original)
	if err != nil {
		t.Fatalf("Failed to marshal XML: %v", err)
	}

	// Test unmarshaling
	var unmarshaled ValueSimpleValue
	err = xml.Unmarshal(xmlData, &unmarshaled)
	if err != nil {
		t.Fatalf("Failed to unmarshal XML: %v", err)
	}

	// Test validation
	if err := original.Validate(); err != nil {
		t.Errorf("Validation failed: %v", err)
	}
}

func TestValueSimpleValueValidation(t *testing.T) {
	// Test valid case
	valid := &ValueSimpleValue{
	}
	if err := valid.Validate(); err != nil {
		t.Errorf("Valid object should not have validation errors: %v", err)
	}

}

func TestValueArrayValueValueXMLMarshaling(t *testing.T) {
	original := &ValueArrayValueValue{
	}

	// Test marshaling
	xmlData, err := xml.Marshal(original)
	if err != nil {
		t.Fatalf("Failed to marshal XML: %v", err)
	}

	// Test unmarshaling
	var unmarshaled ValueArrayValueValue
	err = xml.Unmarshal(xmlData, &unmarshaled)
	if err != nil {
		t.Fatalf("Failed to unmarshal XML: %v", err)
	}

	// Test validation
	if err := original.Validate(); err != nil {
		t.Errorf("Validation failed: %v", err)
	}
}

func TestValueArrayValueValueValidation(t *testing.T) {
	// Test valid case
	valid := &ValueArrayValueValue{
	}
	if err := valid.Validate(); err != nil {
		t.Errorf("Valid object should not have validation errors: %v", err)
	}

}

func TestValueArrayValueXMLMarshaling(t *testing.T) {
	original := &ValueArrayValue{
		Value: ValueArrayValueValue{},
	}

	// Test marshaling
	xmlData, err := xml.Marshal(original)
	if err != nil {
		t.Fatalf("Failed to marshal XML: %v", err)
	}

	// Test unmarshaling
	var unmarshaled ValueArrayValue
	err = xml.Unmarshal(xmlData, &unmarshaled)
	if err != nil {
		t.Fatalf("Failed to unmarshal XML: %v", err)
	}

	// Test validation
	if err := original.Validate(); err != nil {
		t.Errorf("Validation failed: %v", err)
	}
}

func TestValueArrayValueValidation(t *testing.T) {
	// Test valid case
	valid := &ValueArrayValue{
		Value: ValueArrayValueValue{},
	}
	if err := valid.Validate(); err != nil {
		t.Errorf("Valid object should not have validation errors: %v", err)
	}

	// Test missing required field: Value
	invalidValue := &ValueArrayValue{}
	if err := invalidValue.Validate(); err == nil {
		t.Error("Missing required field Value should cause validation error")
	}

}

func TestValueStructValueValueXMLMarshaling(t *testing.T) {
	original := &ValueStructValueValue{
	}

	// Test marshaling
	xmlData, err := xml.Marshal(original)
	if err != nil {
		t.Fatalf("Failed to marshal XML: %v", err)
	}

	// Test unmarshaling
	var unmarshaled ValueStructValueValue
	err = xml.Unmarshal(xmlData, &unmarshaled)
	if err != nil {
		t.Fatalf("Failed to unmarshal XML: %v", err)
	}

	// Test validation
	if err := original.Validate(); err != nil {
		t.Errorf("Validation failed: %v", err)
	}
}

func TestValueStructValueValueValidation(t *testing.T) {
	// Test valid case
	valid := &ValueStructValueValue{
	}
	if err := valid.Validate(); err != nil {
		t.Errorf("Valid object should not have validation errors: %v", err)
	}

}

func TestValueStructValueXMLMarshaling(t *testing.T) {
	original := &ValueStructValue{
		Value: ValueStructValueValue{},
	}

	// Test marshaling
	xmlData, err := xml.Marshal(original)
	if err != nil {
		t.Fatalf("Failed to marshal XML: %v", err)
	}

	// Test unmarshaling
	var unmarshaled ValueStructValue
	err = xml.Unmarshal(xmlData, &unmarshaled)
	if err != nil {
		t.Fatalf("Failed to unmarshal XML: %v", err)
	}

	// Test validation
	if err := original.Validate(); err != nil {
		t.Errorf("Validation failed: %v", err)
	}
}

func TestValueStructValueValidation(t *testing.T) {
	// Test valid case
	valid := &ValueStructValue{
		Value: ValueStructValueValue{},
	}
	if err := valid.Validate(); err != nil {
		t.Errorf("Valid object should not have validation errors: %v", err)
	}

	// Test missing required field: Value
	invalidValue := &ValueStructValue{}
	if err := invalidValue.Validate(); err == nil {
		t.Error("Missing required field Value should cause validation error")
	}

}

func TestValueXMLMarshaling(t *testing.T) {
	original := &Value{
		SimpleValue: &ValueSimpleValue{},
		ArrayValue: &ValueArrayValue{},
		StructValue: &ValueStructValue{},
	}

	// Test marshaling
	xmlData, err := xml.Marshal(original)
	if err != nil {
		t.Fatalf("Failed to marshal XML: %v", err)
	}

	// Test unmarshaling
	var unmarshaled Value
	err = xml.Unmarshal(xmlData, &unmarshaled)
	if err != nil {
		t.Fatalf("Failed to unmarshal XML: %v", err)
	}

	// Test validation
	if err := original.Validate(); err != nil {
		t.Errorf("Validation failed: %v", err)
	}
}

func TestValueValidation(t *testing.T) {
	// Test valid case
	valid := &Value{
	}
	if err := valid.Validate(); err != nil {
		t.Errorf("Valid object should not have validation errors: %v", err)
	}

}

func TestBodyFBDCommentXMLMarshaling(t *testing.T) {
	original := &BodyFBDComment{
		Position: Position{},
		Content: FormattedText{},
		Documentation: &FormattedText{},
		LocalId: 42,
		Height: 3.14,
		Width: 3.14,
	}

	// Test marshaling
	xmlData, err := xml.Marshal(original)
	if err != nil {
		t.Fatalf("Failed to marshal XML: %v", err)
	}

	// Test unmarshaling
	var unmarshaled BodyFBDComment
	err = xml.Unmarshal(xmlData, &unmarshaled)
	if err != nil {
		t.Fatalf("Failed to unmarshal XML: %v", err)
	}

	// Test validation
	if err := original.Validate(); err != nil {
		t.Errorf("Validation failed: %v", err)
	}
}

func TestBodyFBDCommentValidation(t *testing.T) {
	// Test valid case
	valid := &BodyFBDComment{
		Position: Position{},
		Content: FormattedText{},
		LocalId: 42,
		Height: 3.14,
		Width: 3.14,
	}
	if err := valid.Validate(); err != nil {
		t.Errorf("Valid object should not have validation errors: %v", err)
	}

	// Test missing required field: Position
	invalidPosition := &BodyFBDComment{}
	if err := invalidPosition.Validate(); err == nil {
		t.Error("Missing required field Position should cause validation error")
	}

	// Test missing required field: Content
	invalidContent := &BodyFBDComment{}
	if err := invalidContent.Validate(); err == nil {
		t.Error("Missing required field Content should cause validation error")
	}

	// Test missing required field: LocalId
	invalidLocalId := &BodyFBDComment{}
	if err := invalidLocalId.Validate(); err == nil {
		t.Error("Missing required field LocalId should cause validation error")
	}

	// Test missing required field: Height
	invalidHeight := &BodyFBDComment{}
	if err := invalidHeight.Validate(); err == nil {
		t.Error("Missing required field Height should cause validation error")
	}

	// Test missing required field: Width
	invalidWidth := &BodyFBDComment{}
	if err := invalidWidth.Validate(); err == nil {
		t.Error("Missing required field Width should cause validation error")
	}

}

func TestBodyFBDErrorXMLMarshaling(t *testing.T) {
	original := &BodyFBDError{
		Position: Position{},
		Content: FormattedText{},
		Documentation: &FormattedText{},
		LocalId: 42,
		Height: 3.14,
		Width: 3.14,
	}

	// Test marshaling
	xmlData, err := xml.Marshal(original)
	if err != nil {
		t.Fatalf("Failed to marshal XML: %v", err)
	}

	// Test unmarshaling
	var unmarshaled BodyFBDError
	err = xml.Unmarshal(xmlData, &unmarshaled)
	if err != nil {
		t.Fatalf("Failed to unmarshal XML: %v", err)
	}

	// Test validation
	if err := original.Validate(); err != nil {
		t.Errorf("Validation failed: %v", err)
	}
}

func TestBodyFBDErrorValidation(t *testing.T) {
	// Test valid case
	valid := &BodyFBDError{
		Position: Position{},
		Content: FormattedText{},
		LocalId: 42,
		Height: 3.14,
		Width: 3.14,
	}
	if err := valid.Validate(); err != nil {
		t.Errorf("Valid object should not have validation errors: %v", err)
	}

	// Test missing required field: Position
	invalidPosition := &BodyFBDError{}
	if err := invalidPosition.Validate(); err == nil {
		t.Error("Missing required field Position should cause validation error")
	}

	// Test missing required field: Content
	invalidContent := &BodyFBDError{}
	if err := invalidContent.Validate(); err == nil {
		t.Error("Missing required field Content should cause validation error")
	}

	// Test missing required field: LocalId
	invalidLocalId := &BodyFBDError{}
	if err := invalidLocalId.Validate(); err == nil {
		t.Error("Missing required field LocalId should cause validation error")
	}

	// Test missing required field: Height
	invalidHeight := &BodyFBDError{}
	if err := invalidHeight.Validate(); err == nil {
		t.Error("Missing required field Height should cause validation error")
	}

	// Test missing required field: Width
	invalidWidth := &BodyFBDError{}
	if err := invalidWidth.Validate(); err == nil {
		t.Error("Missing required field Width should cause validation error")
	}

}

func TestBodyFBDConnectorXMLMarshaling(t *testing.T) {
	original := &BodyFBDConnector{
		Position: Position{},
		ConnectionPointIn: &ConnectionPointIn{},
		Documentation: &FormattedText{},
		Name: "test_value",
		LocalId: 42,
		Height: floatPtr(3.14),
		Width: floatPtr(3.14),
	}

	// Test marshaling
	xmlData, err := xml.Marshal(original)
	if err != nil {
		t.Fatalf("Failed to marshal XML: %v", err)
	}

	// Test unmarshaling
	var unmarshaled BodyFBDConnector
	err = xml.Unmarshal(xmlData, &unmarshaled)
	if err != nil {
		t.Fatalf("Failed to unmarshal XML: %v", err)
	}

	// Test validation
	if err := original.Validate(); err != nil {
		t.Errorf("Validation failed: %v", err)
	}
}

func TestBodyFBDConnectorValidation(t *testing.T) {
	// Test valid case
	valid := &BodyFBDConnector{
		Position: Position{},
		Name: "test_value",
		LocalId: 42,
	}
	if err := valid.Validate(); err != nil {
		t.Errorf("Valid object should not have validation errors: %v", err)
	}

	// Test missing required field: Position
	invalidPosition := &BodyFBDConnector{}
	if err := invalidPosition.Validate(); err == nil {
		t.Error("Missing required field Position should cause validation error")
	}

	// Test missing required field: Name
	invalidName := &BodyFBDConnector{}
	if err := invalidName.Validate(); err == nil {
		t.Error("Missing required field Name should cause validation error")
	}

	// Test missing required field: LocalId
	invalidLocalId := &BodyFBDConnector{}
	if err := invalidLocalId.Validate(); err == nil {
		t.Error("Missing required field LocalId should cause validation error")
	}

}

func TestBodyFBDContinuationXMLMarshaling(t *testing.T) {
	original := &BodyFBDContinuation{
		Position: Position{},
		ConnectionPointOut: &ConnectionPointOut{},
		Documentation: &FormattedText{},
		Name: "test_value",
		LocalId: 42,
		Height: floatPtr(3.14),
		Width: floatPtr(3.14),
	}

	// Test marshaling
	xmlData, err := xml.Marshal(original)
	if err != nil {
		t.Fatalf("Failed to marshal XML: %v", err)
	}

	// Test unmarshaling
	var unmarshaled BodyFBDContinuation
	err = xml.Unmarshal(xmlData, &unmarshaled)
	if err != nil {
		t.Fatalf("Failed to unmarshal XML: %v", err)
	}

	// Test validation
	if err := original.Validate(); err != nil {
		t.Errorf("Validation failed: %v", err)
	}
}

func TestBodyFBDContinuationValidation(t *testing.T) {
	// Test valid case
	valid := &BodyFBDContinuation{
		Position: Position{},
		Name: "test_value",
		LocalId: 42,
	}
	if err := valid.Validate(); err != nil {
		t.Errorf("Valid object should not have validation errors: %v", err)
	}

	// Test missing required field: Position
	invalidPosition := &BodyFBDContinuation{}
	if err := invalidPosition.Validate(); err == nil {
		t.Error("Missing required field Position should cause validation error")
	}

	// Test missing required field: Name
	invalidName := &BodyFBDContinuation{}
	if err := invalidName.Validate(); err == nil {
		t.Error("Missing required field Name should cause validation error")
	}

	// Test missing required field: LocalId
	invalidLocalId := &BodyFBDContinuation{}
	if err := invalidLocalId.Validate(); err == nil {
		t.Error("Missing required field LocalId should cause validation error")
	}

}

func TestBodyFBDActionBlockActionReferenceXMLMarshaling(t *testing.T) {
	original := &BodyFBDActionBlockActionReference{
		Name: "test_value",
	}

	// Test marshaling
	xmlData, err := xml.Marshal(original)
	if err != nil {
		t.Fatalf("Failed to marshal XML: %v", err)
	}

	// Test unmarshaling
	var unmarshaled BodyFBDActionBlockActionReference
	err = xml.Unmarshal(xmlData, &unmarshaled)
	if err != nil {
		t.Fatalf("Failed to unmarshal XML: %v", err)
	}

	// Test validation
	if err := original.Validate(); err != nil {
		t.Errorf("Validation failed: %v", err)
	}
}

func TestBodyFBDActionBlockActionReferenceValidation(t *testing.T) {
	// Test valid case
	valid := &BodyFBDActionBlockActionReference{
		Name: "test_value",
	}
	if err := valid.Validate(); err != nil {
		t.Errorf("Valid object should not have validation errors: %v", err)
	}

	// Test missing required field: Name
	invalidName := &BodyFBDActionBlockActionReference{}
	if err := invalidName.Validate(); err == nil {
		t.Error("Missing required field Name should cause validation error")
	}

}

func TestBodyFBDActionBlockActionXMLMarshaling(t *testing.T) {
	original := &BodyFBDActionBlockAction{
		Reference: &BodyFBDActionBlockActionReference{},
		Inline: &Body{},
		Documentation: &FormattedText{},
		Qualifier: stringPtr("test_value"),
		Duration: stringPtr("test_value"),
		Indicator: stringPtr("test_value"),
	}

	// Test marshaling
	xmlData, err := xml.Marshal(original)
	if err != nil {
		t.Fatalf("Failed to marshal XML: %v", err)
	}

	// Test unmarshaling
	var unmarshaled BodyFBDActionBlockAction
	err = xml.Unmarshal(xmlData, &unmarshaled)
	if err != nil {
		t.Fatalf("Failed to unmarshal XML: %v", err)
	}

	// Test validation
	if err := original.Validate(); err != nil {
		t.Errorf("Validation failed: %v", err)
	}
}

func TestBodyFBDActionBlockActionValidation(t *testing.T) {
	// Test valid case
	valid := &BodyFBDActionBlockAction{
	}
	if err := valid.Validate(); err != nil {
		t.Errorf("Valid object should not have validation errors: %v", err)
	}

}

func TestBodyFBDActionBlockXMLMarshaling(t *testing.T) {
	original := &BodyFBDActionBlock{
		Position: Position{},
		ConnectionPointIn: &ConnectionPointIn{},
		ConnectionPointOut: &ConnectionPointOut{},
		Action: []BodyFBDActionBlockAction{BodyFBDActionBlockAction{}},
		Documentation: &FormattedText{},
		LocalId: 42,
		Height: floatPtr(3.14),
		Width: floatPtr(3.14),
		Negated: boolPtr(true),
	}

	// Test marshaling
	xmlData, err := xml.Marshal(original)
	if err != nil {
		t.Fatalf("Failed to marshal XML: %v", err)
	}

	// Test unmarshaling
	var unmarshaled BodyFBDActionBlock
	err = xml.Unmarshal(xmlData, &unmarshaled)
	if err != nil {
		t.Fatalf("Failed to unmarshal XML: %v", err)
	}

	// Test validation
	if err := original.Validate(); err != nil {
		t.Errorf("Validation failed: %v", err)
	}
}

func TestBodyFBDActionBlockValidation(t *testing.T) {
	// Test valid case
	valid := &BodyFBDActionBlock{
		Position: Position{},
		LocalId: 42,
	}
	if err := valid.Validate(); err != nil {
		t.Errorf("Valid object should not have validation errors: %v", err)
	}

	// Test missing required field: Position
	invalidPosition := &BodyFBDActionBlock{}
	if err := invalidPosition.Validate(); err == nil {
		t.Error("Missing required field Position should cause validation error")
	}

	// Test missing required field: LocalId
	invalidLocalId := &BodyFBDActionBlock{}
	if err := invalidLocalId.Validate(); err == nil {
		t.Error("Missing required field LocalId should cause validation error")
	}

}

func TestBodyFBDBlockInputVariablesVariableXMLMarshaling(t *testing.T) {
	original := &BodyFBDBlockInputVariablesVariable{
		ConnectionPointIn: ConnectionPointIn{},
		Documentation: &FormattedText{},
		FormalParameter: "test_value",
		Negated: boolPtr(true),
		Edge: func() *EdgeModifierType { v := EdgeModifierType(""); return &v }(),
		Storage: func() *StorageModifierType { v := StorageModifierType(""); return &v }(),
		Hidden: boolPtr(true),
	}

	// Test marshaling
	xmlData, err := xml.Marshal(original)
	if err != nil {
		t.Fatalf("Failed to marshal XML: %v", err)
	}

	// Test unmarshaling
	var unmarshaled BodyFBDBlockInputVariablesVariable
	err = xml.Unmarshal(xmlData, &unmarshaled)
	if err != nil {
		t.Fatalf("Failed to unmarshal XML: %v", err)
	}

	// Test validation
	if err := original.Validate(); err != nil {
		t.Errorf("Validation failed: %v", err)
	}
}

func TestBodyFBDBlockInputVariablesVariableValidation(t *testing.T) {
	// Test valid case
	valid := &BodyFBDBlockInputVariablesVariable{
		ConnectionPointIn: ConnectionPointIn{},
		FormalParameter: "test_value",
	}
	if err := valid.Validate(); err != nil {
		t.Errorf("Valid object should not have validation errors: %v", err)
	}

	// Test missing required field: ConnectionPointIn
	invalidConnectionPointIn := &BodyFBDBlockInputVariablesVariable{}
	if err := invalidConnectionPointIn.Validate(); err == nil {
		t.Error("Missing required field ConnectionPointIn should cause validation error")
	}

	// Test missing required field: FormalParameter
	invalidFormalParameter := &BodyFBDBlockInputVariablesVariable{}
	if err := invalidFormalParameter.Validate(); err == nil {
		t.Error("Missing required field FormalParameter should cause validation error")
	}

}

func TestBodyFBDBlockInputVariablesXMLMarshaling(t *testing.T) {
	original := &BodyFBDBlockInputVariables{
		Variable: []BodyFBDBlockInputVariablesVariable{BodyFBDBlockInputVariablesVariable{}},
	}

	// Test marshaling
	xmlData, err := xml.Marshal(original)
	if err != nil {
		t.Fatalf("Failed to marshal XML: %v", err)
	}

	// Test unmarshaling
	var unmarshaled BodyFBDBlockInputVariables
	err = xml.Unmarshal(xmlData, &unmarshaled)
	if err != nil {
		t.Fatalf("Failed to unmarshal XML: %v", err)
	}

	// Test validation
	if err := original.Validate(); err != nil {
		t.Errorf("Validation failed: %v", err)
	}
}

func TestBodyFBDBlockInputVariablesValidation(t *testing.T) {
	// Test valid case
	valid := &BodyFBDBlockInputVariables{
	}
	if err := valid.Validate(); err != nil {
		t.Errorf("Valid object should not have validation errors: %v", err)
	}

}

func TestBodyFBDBlockInOutVariablesVariableXMLMarshaling(t *testing.T) {
	original := &BodyFBDBlockInOutVariablesVariable{
		ConnectionPointIn: &ConnectionPointIn{},
		ConnectionPointOut: &ConnectionPointOut{},
		Documentation: &FormattedText{},
		FormalParameter: "test_value",
		Negated: boolPtr(true),
		Edge: func() *EdgeModifierType { v := EdgeModifierType(""); return &v }(),
		Storage: func() *StorageModifierType { v := StorageModifierType(""); return &v }(),
		Hidden: boolPtr(true),
	}

	// Test marshaling
	xmlData, err := xml.Marshal(original)
	if err != nil {
		t.Fatalf("Failed to marshal XML: %v", err)
	}

	// Test unmarshaling
	var unmarshaled BodyFBDBlockInOutVariablesVariable
	err = xml.Unmarshal(xmlData, &unmarshaled)
	if err != nil {
		t.Fatalf("Failed to unmarshal XML: %v", err)
	}

	// Test validation
	if err := original.Validate(); err != nil {
		t.Errorf("Validation failed: %v", err)
	}
}

func TestBodyFBDBlockInOutVariablesVariableValidation(t *testing.T) {
	// Test valid case
	valid := &BodyFBDBlockInOutVariablesVariable{
		FormalParameter: "test_value",
	}
	if err := valid.Validate(); err != nil {
		t.Errorf("Valid object should not have validation errors: %v", err)
	}

	// Test missing required field: FormalParameter
	invalidFormalParameter := &BodyFBDBlockInOutVariablesVariable{}
	if err := invalidFormalParameter.Validate(); err == nil {
		t.Error("Missing required field FormalParameter should cause validation error")
	}

}

func TestBodyFBDBlockInOutVariablesXMLMarshaling(t *testing.T) {
	original := &BodyFBDBlockInOutVariables{
		Variable: []BodyFBDBlockInOutVariablesVariable{BodyFBDBlockInOutVariablesVariable{}},
	}

	// Test marshaling
	xmlData, err := xml.Marshal(original)
	if err != nil {
		t.Fatalf("Failed to marshal XML: %v", err)
	}

	// Test unmarshaling
	var unmarshaled BodyFBDBlockInOutVariables
	err = xml.Unmarshal(xmlData, &unmarshaled)
	if err != nil {
		t.Fatalf("Failed to unmarshal XML: %v", err)
	}

	// Test validation
	if err := original.Validate(); err != nil {
		t.Errorf("Validation failed: %v", err)
	}
}

func TestBodyFBDBlockInOutVariablesValidation(t *testing.T) {
	// Test valid case
	valid := &BodyFBDBlockInOutVariables{
	}
	if err := valid.Validate(); err != nil {
		t.Errorf("Valid object should not have validation errors: %v", err)
	}

}

func TestBodyFBDBlockOutputVariablesVariableXMLMarshaling(t *testing.T) {
	original := &BodyFBDBlockOutputVariablesVariable{
		ConnectionPointOut: &ConnectionPointOut{},
		Documentation: &FormattedText{},
		FormalParameter: "test_value",
		Negated: boolPtr(true),
		Edge: func() *EdgeModifierType { v := EdgeModifierType(""); return &v }(),
		Storage: func() *StorageModifierType { v := StorageModifierType(""); return &v }(),
		Hidden: boolPtr(true),
	}

	// Test marshaling
	xmlData, err := xml.Marshal(original)
	if err != nil {
		t.Fatalf("Failed to marshal XML: %v", err)
	}

	// Test unmarshaling
	var unmarshaled BodyFBDBlockOutputVariablesVariable
	err = xml.Unmarshal(xmlData, &unmarshaled)
	if err != nil {
		t.Fatalf("Failed to unmarshal XML: %v", err)
	}

	// Test validation
	if err := original.Validate(); err != nil {
		t.Errorf("Validation failed: %v", err)
	}
}

func TestBodyFBDBlockOutputVariablesVariableValidation(t *testing.T) {
	// Test valid case
	valid := &BodyFBDBlockOutputVariablesVariable{
		FormalParameter: "test_value",
	}
	if err := valid.Validate(); err != nil {
		t.Errorf("Valid object should not have validation errors: %v", err)
	}

	// Test missing required field: FormalParameter
	invalidFormalParameter := &BodyFBDBlockOutputVariablesVariable{}
	if err := invalidFormalParameter.Validate(); err == nil {
		t.Error("Missing required field FormalParameter should cause validation error")
	}

}

func TestBodyFBDBlockOutputVariablesXMLMarshaling(t *testing.T) {
	original := &BodyFBDBlockOutputVariables{
		Variable: []BodyFBDBlockOutputVariablesVariable{BodyFBDBlockOutputVariablesVariable{}},
	}

	// Test marshaling
	xmlData, err := xml.Marshal(original)
	if err != nil {
		t.Fatalf("Failed to marshal XML: %v", err)
	}

	// Test unmarshaling
	var unmarshaled BodyFBDBlockOutputVariables
	err = xml.Unmarshal(xmlData, &unmarshaled)
	if err != nil {
		t.Fatalf("Failed to unmarshal XML: %v", err)
	}

	// Test validation
	if err := original.Validate(); err != nil {
		t.Errorf("Validation failed: %v", err)
	}
}

func TestBodyFBDBlockOutputVariablesValidation(t *testing.T) {
	// Test valid case
	valid := &BodyFBDBlockOutputVariables{
	}
	if err := valid.Validate(); err != nil {
		t.Errorf("Valid object should not have validation errors: %v", err)
	}

}

func TestBodyFBDBlockXMLMarshaling(t *testing.T) {
	original := &BodyFBDBlock{
		Position: Position{},
		InputVariables: BodyFBDBlockInputVariables{},
		InOutVariables: BodyFBDBlockInOutVariables{},
		OutputVariables: BodyFBDBlockOutputVariables{},
		Documentation: &FormattedText{},
		LocalId: 42,
		Width: floatPtr(3.14),
		Height: floatPtr(3.14),
		TypeName: "test_value",
		InstanceName: stringPtr("test_value"),
		ExecutionOrderId: uintPtr(42),
	}

	// Test marshaling
	xmlData, err := xml.Marshal(original)
	if err != nil {
		t.Fatalf("Failed to marshal XML: %v", err)
	}

	// Test unmarshaling
	var unmarshaled BodyFBDBlock
	err = xml.Unmarshal(xmlData, &unmarshaled)
	if err != nil {
		t.Fatalf("Failed to unmarshal XML: %v", err)
	}

	// Test validation
	if err := original.Validate(); err != nil {
		t.Errorf("Validation failed: %v", err)
	}
}

func TestBodyFBDBlockValidation(t *testing.T) {
	// Test valid case
	valid := &BodyFBDBlock{
		Position: Position{},
		InputVariables: BodyFBDBlockInputVariables{},
		InOutVariables: BodyFBDBlockInOutVariables{},
		OutputVariables: BodyFBDBlockOutputVariables{},
		LocalId: 42,
		TypeName: "test_value",
	}
	if err := valid.Validate(); err != nil {
		t.Errorf("Valid object should not have validation errors: %v", err)
	}

	// Test missing required field: Position
	invalidPosition := &BodyFBDBlock{}
	if err := invalidPosition.Validate(); err == nil {
		t.Error("Missing required field Position should cause validation error")
	}

	// Test missing required field: InputVariables
	invalidInputVariables := &BodyFBDBlock{}
	if err := invalidInputVariables.Validate(); err == nil {
		t.Error("Missing required field InputVariables should cause validation error")
	}

	// Test missing required field: InOutVariables
	invalidInOutVariables := &BodyFBDBlock{}
	if err := invalidInOutVariables.Validate(); err == nil {
		t.Error("Missing required field InOutVariables should cause validation error")
	}

	// Test missing required field: OutputVariables
	invalidOutputVariables := &BodyFBDBlock{}
	if err := invalidOutputVariables.Validate(); err == nil {
		t.Error("Missing required field OutputVariables should cause validation error")
	}

	// Test missing required field: LocalId
	invalidLocalId := &BodyFBDBlock{}
	if err := invalidLocalId.Validate(); err == nil {
		t.Error("Missing required field LocalId should cause validation error")
	}

	// Test missing required field: TypeName
	invalidTypeName := &BodyFBDBlock{}
	if err := invalidTypeName.Validate(); err == nil {
		t.Error("Missing required field TypeName should cause validation error")
	}

}

func TestBodyFBDInVariableXMLMarshaling(t *testing.T) {
	original := &BodyFBDInVariable{
		Position: Position{},
		ConnectionPointOut: &ConnectionPointOut{},
		Expression: "test_value",
		Documentation: &FormattedText{},
		LocalId: 42,
		Height: floatPtr(3.14),
		Width: floatPtr(3.14),
		ExecutionOrderId: uintPtr(42),
		Negated: boolPtr(true),
		Edge: func() *EdgeModifierType { v := EdgeModifierType(""); return &v }(),
		Storage: func() *StorageModifierType { v := StorageModifierType(""); return &v }(),
	}

	// Test marshaling
	xmlData, err := xml.Marshal(original)
	if err != nil {
		t.Fatalf("Failed to marshal XML: %v", err)
	}

	// Test unmarshaling
	var unmarshaled BodyFBDInVariable
	err = xml.Unmarshal(xmlData, &unmarshaled)
	if err != nil {
		t.Fatalf("Failed to unmarshal XML: %v", err)
	}

	// Test validation
	if err := original.Validate(); err != nil {
		t.Errorf("Validation failed: %v", err)
	}
}

func TestBodyFBDInVariableValidation(t *testing.T) {
	// Test valid case
	valid := &BodyFBDInVariable{
		Position: Position{},
		Expression: "test_value",
		LocalId: 42,
	}
	if err := valid.Validate(); err != nil {
		t.Errorf("Valid object should not have validation errors: %v", err)
	}

	// Test missing required field: Position
	invalidPosition := &BodyFBDInVariable{}
	if err := invalidPosition.Validate(); err == nil {
		t.Error("Missing required field Position should cause validation error")
	}

	// Test missing required field: Expression
	invalidExpression := &BodyFBDInVariable{}
	if err := invalidExpression.Validate(); err == nil {
		t.Error("Missing required field Expression should cause validation error")
	}

	// Test missing required field: LocalId
	invalidLocalId := &BodyFBDInVariable{}
	if err := invalidLocalId.Validate(); err == nil {
		t.Error("Missing required field LocalId should cause validation error")
	}

}

func TestBodyFBDOutVariableXMLMarshaling(t *testing.T) {
	original := &BodyFBDOutVariable{
		Position: Position{},
		ConnectionPointIn: &ConnectionPointIn{},
		Expression: "test_value",
		Documentation: &FormattedText{},
		LocalId: 42,
		Height: floatPtr(3.14),
		Width: floatPtr(3.14),
		ExecutionOrderId: uintPtr(42),
		Negated: boolPtr(true),
		Edge: func() *EdgeModifierType { v := EdgeModifierType(""); return &v }(),
		Storage: func() *StorageModifierType { v := StorageModifierType(""); return &v }(),
	}

	// Test marshaling
	xmlData, err := xml.Marshal(original)
	if err != nil {
		t.Fatalf("Failed to marshal XML: %v", err)
	}

	// Test unmarshaling
	var unmarshaled BodyFBDOutVariable
	err = xml.Unmarshal(xmlData, &unmarshaled)
	if err != nil {
		t.Fatalf("Failed to unmarshal XML: %v", err)
	}

	// Test validation
	if err := original.Validate(); err != nil {
		t.Errorf("Validation failed: %v", err)
	}
}

func TestBodyFBDOutVariableValidation(t *testing.T) {
	// Test valid case
	valid := &BodyFBDOutVariable{
		Position: Position{},
		Expression: "test_value",
		LocalId: 42,
	}
	if err := valid.Validate(); err != nil {
		t.Errorf("Valid object should not have validation errors: %v", err)
	}

	// Test missing required field: Position
	invalidPosition := &BodyFBDOutVariable{}
	if err := invalidPosition.Validate(); err == nil {
		t.Error("Missing required field Position should cause validation error")
	}

	// Test missing required field: Expression
	invalidExpression := &BodyFBDOutVariable{}
	if err := invalidExpression.Validate(); err == nil {
		t.Error("Missing required field Expression should cause validation error")
	}

	// Test missing required field: LocalId
	invalidLocalId := &BodyFBDOutVariable{}
	if err := invalidLocalId.Validate(); err == nil {
		t.Error("Missing required field LocalId should cause validation error")
	}

}

func TestBodyFBDInOutVariableXMLMarshaling(t *testing.T) {
	original := &BodyFBDInOutVariable{
		Position: Position{},
		ConnectionPointIn: &ConnectionPointIn{},
		ConnectionPointOut: &ConnectionPointOut{},
		Expression: "test_value",
		Documentation: &FormattedText{},
		LocalId: 42,
		Height: floatPtr(3.14),
		Width: floatPtr(3.14),
		ExecutionOrderId: uintPtr(42),
		NegatedIn: boolPtr(true),
		EdgeIn: func() *EdgeModifierType { v := EdgeModifierType(""); return &v }(),
		StorageIn: func() *StorageModifierType { v := StorageModifierType(""); return &v }(),
		NegatedOut: boolPtr(true),
		EdgeOut: func() *EdgeModifierType { v := EdgeModifierType(""); return &v }(),
		StorageOut: func() *StorageModifierType { v := StorageModifierType(""); return &v }(),
	}

	// Test marshaling
	xmlData, err := xml.Marshal(original)
	if err != nil {
		t.Fatalf("Failed to marshal XML: %v", err)
	}

	// Test unmarshaling
	var unmarshaled BodyFBDInOutVariable
	err = xml.Unmarshal(xmlData, &unmarshaled)
	if err != nil {
		t.Fatalf("Failed to unmarshal XML: %v", err)
	}

	// Test validation
	if err := original.Validate(); err != nil {
		t.Errorf("Validation failed: %v", err)
	}
}

func TestBodyFBDInOutVariableValidation(t *testing.T) {
	// Test valid case
	valid := &BodyFBDInOutVariable{
		Position: Position{},
		Expression: "test_value",
		LocalId: 42,
	}
	if err := valid.Validate(); err != nil {
		t.Errorf("Valid object should not have validation errors: %v", err)
	}

	// Test missing required field: Position
	invalidPosition := &BodyFBDInOutVariable{}
	if err := invalidPosition.Validate(); err == nil {
		t.Error("Missing required field Position should cause validation error")
	}

	// Test missing required field: Expression
	invalidExpression := &BodyFBDInOutVariable{}
	if err := invalidExpression.Validate(); err == nil {
		t.Error("Missing required field Expression should cause validation error")
	}

	// Test missing required field: LocalId
	invalidLocalId := &BodyFBDInOutVariable{}
	if err := invalidLocalId.Validate(); err == nil {
		t.Error("Missing required field LocalId should cause validation error")
	}

}

func TestBodyFBDLabelXMLMarshaling(t *testing.T) {
	original := &BodyFBDLabel{
		Position: Position{},
		Documentation: &FormattedText{},
		LocalId: 42,
		Height: floatPtr(3.14),
		Width: floatPtr(3.14),
		Label: "test_value",
		ExecutionOrderId: uintPtr(42),
	}

	// Test marshaling
	xmlData, err := xml.Marshal(original)
	if err != nil {
		t.Fatalf("Failed to marshal XML: %v", err)
	}

	// Test unmarshaling
	var unmarshaled BodyFBDLabel
	err = xml.Unmarshal(xmlData, &unmarshaled)
	if err != nil {
		t.Fatalf("Failed to unmarshal XML: %v", err)
	}

	// Test validation
	if err := original.Validate(); err != nil {
		t.Errorf("Validation failed: %v", err)
	}
}

func TestBodyFBDLabelValidation(t *testing.T) {
	// Test valid case
	valid := &BodyFBDLabel{
		Position: Position{},
		LocalId: 42,
		Label: "test_value",
	}
	if err := valid.Validate(); err != nil {
		t.Errorf("Valid object should not have validation errors: %v", err)
	}

	// Test missing required field: Position
	invalidPosition := &BodyFBDLabel{}
	if err := invalidPosition.Validate(); err == nil {
		t.Error("Missing required field Position should cause validation error")
	}

	// Test missing required field: LocalId
	invalidLocalId := &BodyFBDLabel{}
	if err := invalidLocalId.Validate(); err == nil {
		t.Error("Missing required field LocalId should cause validation error")
	}

	// Test missing required field: Label
	invalidLabel := &BodyFBDLabel{}
	if err := invalidLabel.Validate(); err == nil {
		t.Error("Missing required field Label should cause validation error")
	}

}

func TestBodyFBDJumpXMLMarshaling(t *testing.T) {
	original := &BodyFBDJump{
		Position: Position{},
		ConnectionPointIn: &ConnectionPointIn{},
		Documentation: &FormattedText{},
		LocalId: 42,
		Height: floatPtr(3.14),
		Width: floatPtr(3.14),
		Label: "test_value",
		ExecutionOrderId: uintPtr(42),
	}

	// Test marshaling
	xmlData, err := xml.Marshal(original)
	if err != nil {
		t.Fatalf("Failed to marshal XML: %v", err)
	}

	// Test unmarshaling
	var unmarshaled BodyFBDJump
	err = xml.Unmarshal(xmlData, &unmarshaled)
	if err != nil {
		t.Fatalf("Failed to unmarshal XML: %v", err)
	}

	// Test validation
	if err := original.Validate(); err != nil {
		t.Errorf("Validation failed: %v", err)
	}
}

func TestBodyFBDJumpValidation(t *testing.T) {
	// Test valid case
	valid := &BodyFBDJump{
		Position: Position{},
		LocalId: 42,
		Label: "test_value",
	}
	if err := valid.Validate(); err != nil {
		t.Errorf("Valid object should not have validation errors: %v", err)
	}

	// Test missing required field: Position
	invalidPosition := &BodyFBDJump{}
	if err := invalidPosition.Validate(); err == nil {
		t.Error("Missing required field Position should cause validation error")
	}

	// Test missing required field: LocalId
	invalidLocalId := &BodyFBDJump{}
	if err := invalidLocalId.Validate(); err == nil {
		t.Error("Missing required field LocalId should cause validation error")
	}

	// Test missing required field: Label
	invalidLabel := &BodyFBDJump{}
	if err := invalidLabel.Validate(); err == nil {
		t.Error("Missing required field Label should cause validation error")
	}

}

func TestBodyFBDReturnXMLMarshaling(t *testing.T) {
	original := &BodyFBDReturn{
		Position: Position{},
		ConnectionPointIn: &ConnectionPointIn{},
		Documentation: &FormattedText{},
		LocalId: 42,
		Height: floatPtr(3.14),
		Width: floatPtr(3.14),
		ExecutionOrderId: uintPtr(42),
	}

	// Test marshaling
	xmlData, err := xml.Marshal(original)
	if err != nil {
		t.Fatalf("Failed to marshal XML: %v", err)
	}

	// Test unmarshaling
	var unmarshaled BodyFBDReturn
	err = xml.Unmarshal(xmlData, &unmarshaled)
	if err != nil {
		t.Fatalf("Failed to unmarshal XML: %v", err)
	}

	// Test validation
	if err := original.Validate(); err != nil {
		t.Errorf("Validation failed: %v", err)
	}
}

func TestBodyFBDReturnValidation(t *testing.T) {
	// Test valid case
	valid := &BodyFBDReturn{
		Position: Position{},
		LocalId: 42,
	}
	if err := valid.Validate(); err != nil {
		t.Errorf("Valid object should not have validation errors: %v", err)
	}

	// Test missing required field: Position
	invalidPosition := &BodyFBDReturn{}
	if err := invalidPosition.Validate(); err == nil {
		t.Error("Missing required field Position should cause validation error")
	}

	// Test missing required field: LocalId
	invalidLocalId := &BodyFBDReturn{}
	if err := invalidLocalId.Validate(); err == nil {
		t.Error("Missing required field LocalId should cause validation error")
	}

}

func TestBodyFBDXMLMarshaling(t *testing.T) {
	original := &BodyFBD{
		Comment: &BodyFBDComment{},
		Error: &BodyFBDError{},
		Connector: &BodyFBDConnector{},
		Continuation: &BodyFBDContinuation{},
		ActionBlock: &BodyFBDActionBlock{},
		Block: &BodyFBDBlock{},
		InVariable: &BodyFBDInVariable{},
		OutVariable: &BodyFBDOutVariable{},
		InOutVariable: &BodyFBDInOutVariable{},
		Label: &BodyFBDLabel{},
		Jump: &BodyFBDJump{},
		Return: &BodyFBDReturn{},
	}

	// Test marshaling
	xmlData, err := xml.Marshal(original)
	if err != nil {
		t.Fatalf("Failed to marshal XML: %v", err)
	}

	// Test unmarshaling
	var unmarshaled BodyFBD
	err = xml.Unmarshal(xmlData, &unmarshaled)
	if err != nil {
		t.Fatalf("Failed to unmarshal XML: %v", err)
	}

	// Test validation
	if err := original.Validate(); err != nil {
		t.Errorf("Validation failed: %v", err)
	}
}

func TestBodyFBDValidation(t *testing.T) {
	// Test valid case
	valid := &BodyFBD{
	}
	if err := valid.Validate(); err != nil {
		t.Errorf("Valid object should not have validation errors: %v", err)
	}

}

func TestBodyLDCommentXMLMarshaling(t *testing.T) {
	original := &BodyLDComment{
		Position: Position{},
		Content: FormattedText{},
		Documentation: &FormattedText{},
		LocalId: 42,
		Height: 3.14,
		Width: 3.14,
	}

	// Test marshaling
	xmlData, err := xml.Marshal(original)
	if err != nil {
		t.Fatalf("Failed to marshal XML: %v", err)
	}

	// Test unmarshaling
	var unmarshaled BodyLDComment
	err = xml.Unmarshal(xmlData, &unmarshaled)
	if err != nil {
		t.Fatalf("Failed to unmarshal XML: %v", err)
	}

	// Test validation
	if err := original.Validate(); err != nil {
		t.Errorf("Validation failed: %v", err)
	}
}

func TestBodyLDCommentValidation(t *testing.T) {
	// Test valid case
	valid := &BodyLDComment{
		Position: Position{},
		Content: FormattedText{},
		LocalId: 42,
		Height: 3.14,
		Width: 3.14,
	}
	if err := valid.Validate(); err != nil {
		t.Errorf("Valid object should not have validation errors: %v", err)
	}

	// Test missing required field: Position
	invalidPosition := &BodyLDComment{}
	if err := invalidPosition.Validate(); err == nil {
		t.Error("Missing required field Position should cause validation error")
	}

	// Test missing required field: Content
	invalidContent := &BodyLDComment{}
	if err := invalidContent.Validate(); err == nil {
		t.Error("Missing required field Content should cause validation error")
	}

	// Test missing required field: LocalId
	invalidLocalId := &BodyLDComment{}
	if err := invalidLocalId.Validate(); err == nil {
		t.Error("Missing required field LocalId should cause validation error")
	}

	// Test missing required field: Height
	invalidHeight := &BodyLDComment{}
	if err := invalidHeight.Validate(); err == nil {
		t.Error("Missing required field Height should cause validation error")
	}

	// Test missing required field: Width
	invalidWidth := &BodyLDComment{}
	if err := invalidWidth.Validate(); err == nil {
		t.Error("Missing required field Width should cause validation error")
	}

}

func TestBodyLDErrorXMLMarshaling(t *testing.T) {
	original := &BodyLDError{
		Position: Position{},
		Content: FormattedText{},
		Documentation: &FormattedText{},
		LocalId: 42,
		Height: 3.14,
		Width: 3.14,
	}

	// Test marshaling
	xmlData, err := xml.Marshal(original)
	if err != nil {
		t.Fatalf("Failed to marshal XML: %v", err)
	}

	// Test unmarshaling
	var unmarshaled BodyLDError
	err = xml.Unmarshal(xmlData, &unmarshaled)
	if err != nil {
		t.Fatalf("Failed to unmarshal XML: %v", err)
	}

	// Test validation
	if err := original.Validate(); err != nil {
		t.Errorf("Validation failed: %v", err)
	}
}

func TestBodyLDErrorValidation(t *testing.T) {
	// Test valid case
	valid := &BodyLDError{
		Position: Position{},
		Content: FormattedText{},
		LocalId: 42,
		Height: 3.14,
		Width: 3.14,
	}
	if err := valid.Validate(); err != nil {
		t.Errorf("Valid object should not have validation errors: %v", err)
	}

	// Test missing required field: Position
	invalidPosition := &BodyLDError{}
	if err := invalidPosition.Validate(); err == nil {
		t.Error("Missing required field Position should cause validation error")
	}

	// Test missing required field: Content
	invalidContent := &BodyLDError{}
	if err := invalidContent.Validate(); err == nil {
		t.Error("Missing required field Content should cause validation error")
	}

	// Test missing required field: LocalId
	invalidLocalId := &BodyLDError{}
	if err := invalidLocalId.Validate(); err == nil {
		t.Error("Missing required field LocalId should cause validation error")
	}

	// Test missing required field: Height
	invalidHeight := &BodyLDError{}
	if err := invalidHeight.Validate(); err == nil {
		t.Error("Missing required field Height should cause validation error")
	}

	// Test missing required field: Width
	invalidWidth := &BodyLDError{}
	if err := invalidWidth.Validate(); err == nil {
		t.Error("Missing required field Width should cause validation error")
	}

}

func TestBodyLDConnectorXMLMarshaling(t *testing.T) {
	original := &BodyLDConnector{
		Position: Position{},
		ConnectionPointIn: &ConnectionPointIn{},
		Documentation: &FormattedText{},
		Name: "test_value",
		LocalId: 42,
		Height: floatPtr(3.14),
		Width: floatPtr(3.14),
	}

	// Test marshaling
	xmlData, err := xml.Marshal(original)
	if err != nil {
		t.Fatalf("Failed to marshal XML: %v", err)
	}

	// Test unmarshaling
	var unmarshaled BodyLDConnector
	err = xml.Unmarshal(xmlData, &unmarshaled)
	if err != nil {
		t.Fatalf("Failed to unmarshal XML: %v", err)
	}

	// Test validation
	if err := original.Validate(); err != nil {
		t.Errorf("Validation failed: %v", err)
	}
}

func TestBodyLDConnectorValidation(t *testing.T) {
	// Test valid case
	valid := &BodyLDConnector{
		Position: Position{},
		Name: "test_value",
		LocalId: 42,
	}
	if err := valid.Validate(); err != nil {
		t.Errorf("Valid object should not have validation errors: %v", err)
	}

	// Test missing required field: Position
	invalidPosition := &BodyLDConnector{}
	if err := invalidPosition.Validate(); err == nil {
		t.Error("Missing required field Position should cause validation error")
	}

	// Test missing required field: Name
	invalidName := &BodyLDConnector{}
	if err := invalidName.Validate(); err == nil {
		t.Error("Missing required field Name should cause validation error")
	}

	// Test missing required field: LocalId
	invalidLocalId := &BodyLDConnector{}
	if err := invalidLocalId.Validate(); err == nil {
		t.Error("Missing required field LocalId should cause validation error")
	}

}

func TestBodyLDContinuationXMLMarshaling(t *testing.T) {
	original := &BodyLDContinuation{
		Position: Position{},
		ConnectionPointOut: &ConnectionPointOut{},
		Documentation: &FormattedText{},
		Name: "test_value",
		LocalId: 42,
		Height: floatPtr(3.14),
		Width: floatPtr(3.14),
	}

	// Test marshaling
	xmlData, err := xml.Marshal(original)
	if err != nil {
		t.Fatalf("Failed to marshal XML: %v", err)
	}

	// Test unmarshaling
	var unmarshaled BodyLDContinuation
	err = xml.Unmarshal(xmlData, &unmarshaled)
	if err != nil {
		t.Fatalf("Failed to unmarshal XML: %v", err)
	}

	// Test validation
	if err := original.Validate(); err != nil {
		t.Errorf("Validation failed: %v", err)
	}
}

func TestBodyLDContinuationValidation(t *testing.T) {
	// Test valid case
	valid := &BodyLDContinuation{
		Position: Position{},
		Name: "test_value",
		LocalId: 42,
	}
	if err := valid.Validate(); err != nil {
		t.Errorf("Valid object should not have validation errors: %v", err)
	}

	// Test missing required field: Position
	invalidPosition := &BodyLDContinuation{}
	if err := invalidPosition.Validate(); err == nil {
		t.Error("Missing required field Position should cause validation error")
	}

	// Test missing required field: Name
	invalidName := &BodyLDContinuation{}
	if err := invalidName.Validate(); err == nil {
		t.Error("Missing required field Name should cause validation error")
	}

	// Test missing required field: LocalId
	invalidLocalId := &BodyLDContinuation{}
	if err := invalidLocalId.Validate(); err == nil {
		t.Error("Missing required field LocalId should cause validation error")
	}

}

func TestBodyLDActionBlockActionReferenceXMLMarshaling(t *testing.T) {
	original := &BodyLDActionBlockActionReference{
		Name: "test_value",
	}

	// Test marshaling
	xmlData, err := xml.Marshal(original)
	if err != nil {
		t.Fatalf("Failed to marshal XML: %v", err)
	}

	// Test unmarshaling
	var unmarshaled BodyLDActionBlockActionReference
	err = xml.Unmarshal(xmlData, &unmarshaled)
	if err != nil {
		t.Fatalf("Failed to unmarshal XML: %v", err)
	}

	// Test validation
	if err := original.Validate(); err != nil {
		t.Errorf("Validation failed: %v", err)
	}
}

func TestBodyLDActionBlockActionReferenceValidation(t *testing.T) {
	// Test valid case
	valid := &BodyLDActionBlockActionReference{
		Name: "test_value",
	}
	if err := valid.Validate(); err != nil {
		t.Errorf("Valid object should not have validation errors: %v", err)
	}

	// Test missing required field: Name
	invalidName := &BodyLDActionBlockActionReference{}
	if err := invalidName.Validate(); err == nil {
		t.Error("Missing required field Name should cause validation error")
	}

}

func TestBodyLDActionBlockActionXMLMarshaling(t *testing.T) {
	original := &BodyLDActionBlockAction{
		Reference: &BodyLDActionBlockActionReference{},
		Inline: &Body{},
		Documentation: &FormattedText{},
		Qualifier: stringPtr("test_value"),
		Duration: stringPtr("test_value"),
		Indicator: stringPtr("test_value"),
	}

	// Test marshaling
	xmlData, err := xml.Marshal(original)
	if err != nil {
		t.Fatalf("Failed to marshal XML: %v", err)
	}

	// Test unmarshaling
	var unmarshaled BodyLDActionBlockAction
	err = xml.Unmarshal(xmlData, &unmarshaled)
	if err != nil {
		t.Fatalf("Failed to unmarshal XML: %v", err)
	}

	// Test validation
	if err := original.Validate(); err != nil {
		t.Errorf("Validation failed: %v", err)
	}
}

func TestBodyLDActionBlockActionValidation(t *testing.T) {
	// Test valid case
	valid := &BodyLDActionBlockAction{
	}
	if err := valid.Validate(); err != nil {
		t.Errorf("Valid object should not have validation errors: %v", err)
	}

}

func TestBodyLDActionBlockXMLMarshaling(t *testing.T) {
	original := &BodyLDActionBlock{
		Position: Position{},
		ConnectionPointIn: &ConnectionPointIn{},
		ConnectionPointOut: &ConnectionPointOut{},
		Action: []BodyLDActionBlockAction{BodyLDActionBlockAction{}},
		Documentation: &FormattedText{},
		LocalId: 42,
		Height: floatPtr(3.14),
		Width: floatPtr(3.14),
		Negated: boolPtr(true),
	}

	// Test marshaling
	xmlData, err := xml.Marshal(original)
	if err != nil {
		t.Fatalf("Failed to marshal XML: %v", err)
	}

	// Test unmarshaling
	var unmarshaled BodyLDActionBlock
	err = xml.Unmarshal(xmlData, &unmarshaled)
	if err != nil {
		t.Fatalf("Failed to unmarshal XML: %v", err)
	}

	// Test validation
	if err := original.Validate(); err != nil {
		t.Errorf("Validation failed: %v", err)
	}
}

func TestBodyLDActionBlockValidation(t *testing.T) {
	// Test valid case
	valid := &BodyLDActionBlock{
		Position: Position{},
		LocalId: 42,
	}
	if err := valid.Validate(); err != nil {
		t.Errorf("Valid object should not have validation errors: %v", err)
	}

	// Test missing required field: Position
	invalidPosition := &BodyLDActionBlock{}
	if err := invalidPosition.Validate(); err == nil {
		t.Error("Missing required field Position should cause validation error")
	}

	// Test missing required field: LocalId
	invalidLocalId := &BodyLDActionBlock{}
	if err := invalidLocalId.Validate(); err == nil {
		t.Error("Missing required field LocalId should cause validation error")
	}

}

func TestBodyLDBlockInputVariablesVariableXMLMarshaling(t *testing.T) {
	original := &BodyLDBlockInputVariablesVariable{
		ConnectionPointIn: ConnectionPointIn{},
		Documentation: &FormattedText{},
		FormalParameter: "test_value",
		Negated: boolPtr(true),
		Edge: func() *EdgeModifierType { v := EdgeModifierType(""); return &v }(),
		Storage: func() *StorageModifierType { v := StorageModifierType(""); return &v }(),
		Hidden: boolPtr(true),
	}

	// Test marshaling
	xmlData, err := xml.Marshal(original)
	if err != nil {
		t.Fatalf("Failed to marshal XML: %v", err)
	}

	// Test unmarshaling
	var unmarshaled BodyLDBlockInputVariablesVariable
	err = xml.Unmarshal(xmlData, &unmarshaled)
	if err != nil {
		t.Fatalf("Failed to unmarshal XML: %v", err)
	}

	// Test validation
	if err := original.Validate(); err != nil {
		t.Errorf("Validation failed: %v", err)
	}
}

func TestBodyLDBlockInputVariablesVariableValidation(t *testing.T) {
	// Test valid case
	valid := &BodyLDBlockInputVariablesVariable{
		ConnectionPointIn: ConnectionPointIn{},
		FormalParameter: "test_value",
	}
	if err := valid.Validate(); err != nil {
		t.Errorf("Valid object should not have validation errors: %v", err)
	}

	// Test missing required field: ConnectionPointIn
	invalidConnectionPointIn := &BodyLDBlockInputVariablesVariable{}
	if err := invalidConnectionPointIn.Validate(); err == nil {
		t.Error("Missing required field ConnectionPointIn should cause validation error")
	}

	// Test missing required field: FormalParameter
	invalidFormalParameter := &BodyLDBlockInputVariablesVariable{}
	if err := invalidFormalParameter.Validate(); err == nil {
		t.Error("Missing required field FormalParameter should cause validation error")
	}

}

func TestBodyLDBlockInputVariablesXMLMarshaling(t *testing.T) {
	original := &BodyLDBlockInputVariables{
		Variable: []BodyLDBlockInputVariablesVariable{BodyLDBlockInputVariablesVariable{}},
	}

	// Test marshaling
	xmlData, err := xml.Marshal(original)
	if err != nil {
		t.Fatalf("Failed to marshal XML: %v", err)
	}

	// Test unmarshaling
	var unmarshaled BodyLDBlockInputVariables
	err = xml.Unmarshal(xmlData, &unmarshaled)
	if err != nil {
		t.Fatalf("Failed to unmarshal XML: %v", err)
	}

	// Test validation
	if err := original.Validate(); err != nil {
		t.Errorf("Validation failed: %v", err)
	}
}

func TestBodyLDBlockInputVariablesValidation(t *testing.T) {
	// Test valid case
	valid := &BodyLDBlockInputVariables{
	}
	if err := valid.Validate(); err != nil {
		t.Errorf("Valid object should not have validation errors: %v", err)
	}

}

func TestBodyLDBlockInOutVariablesVariableXMLMarshaling(t *testing.T) {
	original := &BodyLDBlockInOutVariablesVariable{
		ConnectionPointIn: &ConnectionPointIn{},
		ConnectionPointOut: &ConnectionPointOut{},
		Documentation: &FormattedText{},
		FormalParameter: "test_value",
		Negated: boolPtr(true),
		Edge: func() *EdgeModifierType { v := EdgeModifierType(""); return &v }(),
		Storage: func() *StorageModifierType { v := StorageModifierType(""); return &v }(),
		Hidden: boolPtr(true),
	}

	// Test marshaling
	xmlData, err := xml.Marshal(original)
	if err != nil {
		t.Fatalf("Failed to marshal XML: %v", err)
	}

	// Test unmarshaling
	var unmarshaled BodyLDBlockInOutVariablesVariable
	err = xml.Unmarshal(xmlData, &unmarshaled)
	if err != nil {
		t.Fatalf("Failed to unmarshal XML: %v", err)
	}

	// Test validation
	if err := original.Validate(); err != nil {
		t.Errorf("Validation failed: %v", err)
	}
}

func TestBodyLDBlockInOutVariablesVariableValidation(t *testing.T) {
	// Test valid case
	valid := &BodyLDBlockInOutVariablesVariable{
		FormalParameter: "test_value",
	}
	if err := valid.Validate(); err != nil {
		t.Errorf("Valid object should not have validation errors: %v", err)
	}

	// Test missing required field: FormalParameter
	invalidFormalParameter := &BodyLDBlockInOutVariablesVariable{}
	if err := invalidFormalParameter.Validate(); err == nil {
		t.Error("Missing required field FormalParameter should cause validation error")
	}

}

func TestBodyLDBlockInOutVariablesXMLMarshaling(t *testing.T) {
	original := &BodyLDBlockInOutVariables{
		Variable: []BodyLDBlockInOutVariablesVariable{BodyLDBlockInOutVariablesVariable{}},
	}

	// Test marshaling
	xmlData, err := xml.Marshal(original)
	if err != nil {
		t.Fatalf("Failed to marshal XML: %v", err)
	}

	// Test unmarshaling
	var unmarshaled BodyLDBlockInOutVariables
	err = xml.Unmarshal(xmlData, &unmarshaled)
	if err != nil {
		t.Fatalf("Failed to unmarshal XML: %v", err)
	}

	// Test validation
	if err := original.Validate(); err != nil {
		t.Errorf("Validation failed: %v", err)
	}
}

func TestBodyLDBlockInOutVariablesValidation(t *testing.T) {
	// Test valid case
	valid := &BodyLDBlockInOutVariables{
	}
	if err := valid.Validate(); err != nil {
		t.Errorf("Valid object should not have validation errors: %v", err)
	}

}

func TestBodyLDBlockOutputVariablesVariableXMLMarshaling(t *testing.T) {
	original := &BodyLDBlockOutputVariablesVariable{
		ConnectionPointOut: &ConnectionPointOut{},
		Documentation: &FormattedText{},
		FormalParameter: "test_value",
		Negated: boolPtr(true),
		Edge: func() *EdgeModifierType { v := EdgeModifierType(""); return &v }(),
		Storage: func() *StorageModifierType { v := StorageModifierType(""); return &v }(),
		Hidden: boolPtr(true),
	}

	// Test marshaling
	xmlData, err := xml.Marshal(original)
	if err != nil {
		t.Fatalf("Failed to marshal XML: %v", err)
	}

	// Test unmarshaling
	var unmarshaled BodyLDBlockOutputVariablesVariable
	err = xml.Unmarshal(xmlData, &unmarshaled)
	if err != nil {
		t.Fatalf("Failed to unmarshal XML: %v", err)
	}

	// Test validation
	if err := original.Validate(); err != nil {
		t.Errorf("Validation failed: %v", err)
	}
}

func TestBodyLDBlockOutputVariablesVariableValidation(t *testing.T) {
	// Test valid case
	valid := &BodyLDBlockOutputVariablesVariable{
		FormalParameter: "test_value",
	}
	if err := valid.Validate(); err != nil {
		t.Errorf("Valid object should not have validation errors: %v", err)
	}

	// Test missing required field: FormalParameter
	invalidFormalParameter := &BodyLDBlockOutputVariablesVariable{}
	if err := invalidFormalParameter.Validate(); err == nil {
		t.Error("Missing required field FormalParameter should cause validation error")
	}

}

func TestBodyLDBlockOutputVariablesXMLMarshaling(t *testing.T) {
	original := &BodyLDBlockOutputVariables{
		Variable: []BodyLDBlockOutputVariablesVariable{BodyLDBlockOutputVariablesVariable{}},
	}

	// Test marshaling
	xmlData, err := xml.Marshal(original)
	if err != nil {
		t.Fatalf("Failed to marshal XML: %v", err)
	}

	// Test unmarshaling
	var unmarshaled BodyLDBlockOutputVariables
	err = xml.Unmarshal(xmlData, &unmarshaled)
	if err != nil {
		t.Fatalf("Failed to unmarshal XML: %v", err)
	}

	// Test validation
	if err := original.Validate(); err != nil {
		t.Errorf("Validation failed: %v", err)
	}
}

func TestBodyLDBlockOutputVariablesValidation(t *testing.T) {
	// Test valid case
	valid := &BodyLDBlockOutputVariables{
	}
	if err := valid.Validate(); err != nil {
		t.Errorf("Valid object should not have validation errors: %v", err)
	}

}

func TestBodyLDBlockXMLMarshaling(t *testing.T) {
	original := &BodyLDBlock{
		Position: Position{},
		InputVariables: BodyLDBlockInputVariables{},
		InOutVariables: BodyLDBlockInOutVariables{},
		OutputVariables: BodyLDBlockOutputVariables{},
		Documentation: &FormattedText{},
		LocalId: 42,
		Width: floatPtr(3.14),
		Height: floatPtr(3.14),
		TypeName: "test_value",
		InstanceName: stringPtr("test_value"),
		ExecutionOrderId: uintPtr(42),
	}

	// Test marshaling
	xmlData, err := xml.Marshal(original)
	if err != nil {
		t.Fatalf("Failed to marshal XML: %v", err)
	}

	// Test unmarshaling
	var unmarshaled BodyLDBlock
	err = xml.Unmarshal(xmlData, &unmarshaled)
	if err != nil {
		t.Fatalf("Failed to unmarshal XML: %v", err)
	}

	// Test validation
	if err := original.Validate(); err != nil {
		t.Errorf("Validation failed: %v", err)
	}
}

func TestBodyLDBlockValidation(t *testing.T) {
	// Test valid case
	valid := &BodyLDBlock{
		Position: Position{},
		InputVariables: BodyLDBlockInputVariables{},
		InOutVariables: BodyLDBlockInOutVariables{},
		OutputVariables: BodyLDBlockOutputVariables{},
		LocalId: 42,
		TypeName: "test_value",
	}
	if err := valid.Validate(); err != nil {
		t.Errorf("Valid object should not have validation errors: %v", err)
	}

	// Test missing required field: Position
	invalidPosition := &BodyLDBlock{}
	if err := invalidPosition.Validate(); err == nil {
		t.Error("Missing required field Position should cause validation error")
	}

	// Test missing required field: InputVariables
	invalidInputVariables := &BodyLDBlock{}
	if err := invalidInputVariables.Validate(); err == nil {
		t.Error("Missing required field InputVariables should cause validation error")
	}

	// Test missing required field: InOutVariables
	invalidInOutVariables := &BodyLDBlock{}
	if err := invalidInOutVariables.Validate(); err == nil {
		t.Error("Missing required field InOutVariables should cause validation error")
	}

	// Test missing required field: OutputVariables
	invalidOutputVariables := &BodyLDBlock{}
	if err := invalidOutputVariables.Validate(); err == nil {
		t.Error("Missing required field OutputVariables should cause validation error")
	}

	// Test missing required field: LocalId
	invalidLocalId := &BodyLDBlock{}
	if err := invalidLocalId.Validate(); err == nil {
		t.Error("Missing required field LocalId should cause validation error")
	}

	// Test missing required field: TypeName
	invalidTypeName := &BodyLDBlock{}
	if err := invalidTypeName.Validate(); err == nil {
		t.Error("Missing required field TypeName should cause validation error")
	}

}

func TestBodyLDInVariableXMLMarshaling(t *testing.T) {
	original := &BodyLDInVariable{
		Position: Position{},
		ConnectionPointOut: &ConnectionPointOut{},
		Expression: "test_value",
		Documentation: &FormattedText{},
		LocalId: 42,
		Height: floatPtr(3.14),
		Width: floatPtr(3.14),
		ExecutionOrderId: uintPtr(42),
		Negated: boolPtr(true),
		Edge: func() *EdgeModifierType { v := EdgeModifierType(""); return &v }(),
		Storage: func() *StorageModifierType { v := StorageModifierType(""); return &v }(),
	}

	// Test marshaling
	xmlData, err := xml.Marshal(original)
	if err != nil {
		t.Fatalf("Failed to marshal XML: %v", err)
	}

	// Test unmarshaling
	var unmarshaled BodyLDInVariable
	err = xml.Unmarshal(xmlData, &unmarshaled)
	if err != nil {
		t.Fatalf("Failed to unmarshal XML: %v", err)
	}

	// Test validation
	if err := original.Validate(); err != nil {
		t.Errorf("Validation failed: %v", err)
	}
}

func TestBodyLDInVariableValidation(t *testing.T) {
	// Test valid case
	valid := &BodyLDInVariable{
		Position: Position{},
		Expression: "test_value",
		LocalId: 42,
	}
	if err := valid.Validate(); err != nil {
		t.Errorf("Valid object should not have validation errors: %v", err)
	}

	// Test missing required field: Position
	invalidPosition := &BodyLDInVariable{}
	if err := invalidPosition.Validate(); err == nil {
		t.Error("Missing required field Position should cause validation error")
	}

	// Test missing required field: Expression
	invalidExpression := &BodyLDInVariable{}
	if err := invalidExpression.Validate(); err == nil {
		t.Error("Missing required field Expression should cause validation error")
	}

	// Test missing required field: LocalId
	invalidLocalId := &BodyLDInVariable{}
	if err := invalidLocalId.Validate(); err == nil {
		t.Error("Missing required field LocalId should cause validation error")
	}

}

func TestBodyLDOutVariableXMLMarshaling(t *testing.T) {
	original := &BodyLDOutVariable{
		Position: Position{},
		ConnectionPointIn: &ConnectionPointIn{},
		Expression: "test_value",
		Documentation: &FormattedText{},
		LocalId: 42,
		Height: floatPtr(3.14),
		Width: floatPtr(3.14),
		ExecutionOrderId: uintPtr(42),
		Negated: boolPtr(true),
		Edge: func() *EdgeModifierType { v := EdgeModifierType(""); return &v }(),
		Storage: func() *StorageModifierType { v := StorageModifierType(""); return &v }(),
	}

	// Test marshaling
	xmlData, err := xml.Marshal(original)
	if err != nil {
		t.Fatalf("Failed to marshal XML: %v", err)
	}

	// Test unmarshaling
	var unmarshaled BodyLDOutVariable
	err = xml.Unmarshal(xmlData, &unmarshaled)
	if err != nil {
		t.Fatalf("Failed to unmarshal XML: %v", err)
	}

	// Test validation
	if err := original.Validate(); err != nil {
		t.Errorf("Validation failed: %v", err)
	}
}

func TestBodyLDOutVariableValidation(t *testing.T) {
	// Test valid case
	valid := &BodyLDOutVariable{
		Position: Position{},
		Expression: "test_value",
		LocalId: 42,
	}
	if err := valid.Validate(); err != nil {
		t.Errorf("Valid object should not have validation errors: %v", err)
	}

	// Test missing required field: Position
	invalidPosition := &BodyLDOutVariable{}
	if err := invalidPosition.Validate(); err == nil {
		t.Error("Missing required field Position should cause validation error")
	}

	// Test missing required field: Expression
	invalidExpression := &BodyLDOutVariable{}
	if err := invalidExpression.Validate(); err == nil {
		t.Error("Missing required field Expression should cause validation error")
	}

	// Test missing required field: LocalId
	invalidLocalId := &BodyLDOutVariable{}
	if err := invalidLocalId.Validate(); err == nil {
		t.Error("Missing required field LocalId should cause validation error")
	}

}

func TestBodyLDInOutVariableXMLMarshaling(t *testing.T) {
	original := &BodyLDInOutVariable{
		Position: Position{},
		ConnectionPointIn: &ConnectionPointIn{},
		ConnectionPointOut: &ConnectionPointOut{},
		Expression: "test_value",
		Documentation: &FormattedText{},
		LocalId: 42,
		Height: floatPtr(3.14),
		Width: floatPtr(3.14),
		ExecutionOrderId: uintPtr(42),
		NegatedIn: boolPtr(true),
		EdgeIn: func() *EdgeModifierType { v := EdgeModifierType(""); return &v }(),
		StorageIn: func() *StorageModifierType { v := StorageModifierType(""); return &v }(),
		NegatedOut: boolPtr(true),
		EdgeOut: func() *EdgeModifierType { v := EdgeModifierType(""); return &v }(),
		StorageOut: func() *StorageModifierType { v := StorageModifierType(""); return &v }(),
	}

	// Test marshaling
	xmlData, err := xml.Marshal(original)
	if err != nil {
		t.Fatalf("Failed to marshal XML: %v", err)
	}

	// Test unmarshaling
	var unmarshaled BodyLDInOutVariable
	err = xml.Unmarshal(xmlData, &unmarshaled)
	if err != nil {
		t.Fatalf("Failed to unmarshal XML: %v", err)
	}

	// Test validation
	if err := original.Validate(); err != nil {
		t.Errorf("Validation failed: %v", err)
	}
}

func TestBodyLDInOutVariableValidation(t *testing.T) {
	// Test valid case
	valid := &BodyLDInOutVariable{
		Position: Position{},
		Expression: "test_value",
		LocalId: 42,
	}
	if err := valid.Validate(); err != nil {
		t.Errorf("Valid object should not have validation errors: %v", err)
	}

	// Test missing required field: Position
	invalidPosition := &BodyLDInOutVariable{}
	if err := invalidPosition.Validate(); err == nil {
		t.Error("Missing required field Position should cause validation error")
	}

	// Test missing required field: Expression
	invalidExpression := &BodyLDInOutVariable{}
	if err := invalidExpression.Validate(); err == nil {
		t.Error("Missing required field Expression should cause validation error")
	}

	// Test missing required field: LocalId
	invalidLocalId := &BodyLDInOutVariable{}
	if err := invalidLocalId.Validate(); err == nil {
		t.Error("Missing required field LocalId should cause validation error")
	}

}

func TestBodyLDLabelXMLMarshaling(t *testing.T) {
	original := &BodyLDLabel{
		Position: Position{},
		Documentation: &FormattedText{},
		LocalId: 42,
		Height: floatPtr(3.14),
		Width: floatPtr(3.14),
		Label: "test_value",
		ExecutionOrderId: uintPtr(42),
	}

	// Test marshaling
	xmlData, err := xml.Marshal(original)
	if err != nil {
		t.Fatalf("Failed to marshal XML: %v", err)
	}

	// Test unmarshaling
	var unmarshaled BodyLDLabel
	err = xml.Unmarshal(xmlData, &unmarshaled)
	if err != nil {
		t.Fatalf("Failed to unmarshal XML: %v", err)
	}

	// Test validation
	if err := original.Validate(); err != nil {
		t.Errorf("Validation failed: %v", err)
	}
}

func TestBodyLDLabelValidation(t *testing.T) {
	// Test valid case
	valid := &BodyLDLabel{
		Position: Position{},
		LocalId: 42,
		Label: "test_value",
	}
	if err := valid.Validate(); err != nil {
		t.Errorf("Valid object should not have validation errors: %v", err)
	}

	// Test missing required field: Position
	invalidPosition := &BodyLDLabel{}
	if err := invalidPosition.Validate(); err == nil {
		t.Error("Missing required field Position should cause validation error")
	}

	// Test missing required field: LocalId
	invalidLocalId := &BodyLDLabel{}
	if err := invalidLocalId.Validate(); err == nil {
		t.Error("Missing required field LocalId should cause validation error")
	}

	// Test missing required field: Label
	invalidLabel := &BodyLDLabel{}
	if err := invalidLabel.Validate(); err == nil {
		t.Error("Missing required field Label should cause validation error")
	}

}

func TestBodyLDJumpXMLMarshaling(t *testing.T) {
	original := &BodyLDJump{
		Position: Position{},
		ConnectionPointIn: &ConnectionPointIn{},
		Documentation: &FormattedText{},
		LocalId: 42,
		Height: floatPtr(3.14),
		Width: floatPtr(3.14),
		Label: "test_value",
		ExecutionOrderId: uintPtr(42),
	}

	// Test marshaling
	xmlData, err := xml.Marshal(original)
	if err != nil {
		t.Fatalf("Failed to marshal XML: %v", err)
	}

	// Test unmarshaling
	var unmarshaled BodyLDJump
	err = xml.Unmarshal(xmlData, &unmarshaled)
	if err != nil {
		t.Fatalf("Failed to unmarshal XML: %v", err)
	}

	// Test validation
	if err := original.Validate(); err != nil {
		t.Errorf("Validation failed: %v", err)
	}
}

func TestBodyLDJumpValidation(t *testing.T) {
	// Test valid case
	valid := &BodyLDJump{
		Position: Position{},
		LocalId: 42,
		Label: "test_value",
	}
	if err := valid.Validate(); err != nil {
		t.Errorf("Valid object should not have validation errors: %v", err)
	}

	// Test missing required field: Position
	invalidPosition := &BodyLDJump{}
	if err := invalidPosition.Validate(); err == nil {
		t.Error("Missing required field Position should cause validation error")
	}

	// Test missing required field: LocalId
	invalidLocalId := &BodyLDJump{}
	if err := invalidLocalId.Validate(); err == nil {
		t.Error("Missing required field LocalId should cause validation error")
	}

	// Test missing required field: Label
	invalidLabel := &BodyLDJump{}
	if err := invalidLabel.Validate(); err == nil {
		t.Error("Missing required field Label should cause validation error")
	}

}

func TestBodyLDReturnXMLMarshaling(t *testing.T) {
	original := &BodyLDReturn{
		Position: Position{},
		ConnectionPointIn: &ConnectionPointIn{},
		Documentation: &FormattedText{},
		LocalId: 42,
		Height: floatPtr(3.14),
		Width: floatPtr(3.14),
		ExecutionOrderId: uintPtr(42),
	}

	// Test marshaling
	xmlData, err := xml.Marshal(original)
	if err != nil {
		t.Fatalf("Failed to marshal XML: %v", err)
	}

	// Test unmarshaling
	var unmarshaled BodyLDReturn
	err = xml.Unmarshal(xmlData, &unmarshaled)
	if err != nil {
		t.Fatalf("Failed to unmarshal XML: %v", err)
	}

	// Test validation
	if err := original.Validate(); err != nil {
		t.Errorf("Validation failed: %v", err)
	}
}

func TestBodyLDReturnValidation(t *testing.T) {
	// Test valid case
	valid := &BodyLDReturn{
		Position: Position{},
		LocalId: 42,
	}
	if err := valid.Validate(); err != nil {
		t.Errorf("Valid object should not have validation errors: %v", err)
	}

	// Test missing required field: Position
	invalidPosition := &BodyLDReturn{}
	if err := invalidPosition.Validate(); err == nil {
		t.Error("Missing required field Position should cause validation error")
	}

	// Test missing required field: LocalId
	invalidLocalId := &BodyLDReturn{}
	if err := invalidLocalId.Validate(); err == nil {
		t.Error("Missing required field LocalId should cause validation error")
	}

}

func TestBodyLDLeftPowerRailConnectionPointOutXMLMarshaling(t *testing.T) {
	original := &BodyLDLeftPowerRailConnectionPointOut{
	}

	// Test marshaling
	xmlData, err := xml.Marshal(original)
	if err != nil {
		t.Fatalf("Failed to marshal XML: %v", err)
	}

	// Test unmarshaling
	var unmarshaled BodyLDLeftPowerRailConnectionPointOut
	err = xml.Unmarshal(xmlData, &unmarshaled)
	if err != nil {
		t.Fatalf("Failed to unmarshal XML: %v", err)
	}

	// Test validation
	if err := original.Validate(); err != nil {
		t.Errorf("Validation failed: %v", err)
	}
}

func TestBodyLDLeftPowerRailConnectionPointOutValidation(t *testing.T) {
	// Test valid case
	valid := &BodyLDLeftPowerRailConnectionPointOut{
	}
	if err := valid.Validate(); err != nil {
		t.Errorf("Valid object should not have validation errors: %v", err)
	}

}

func TestBodyLDLeftPowerRailXMLMarshaling(t *testing.T) {
	original := &BodyLDLeftPowerRail{
		Position: Position{},
		ConnectionPointOut: []BodyLDLeftPowerRailConnectionPointOut{BodyLDLeftPowerRailConnectionPointOut{}},
		Documentation: &FormattedText{},
		LocalId: 42,
		Height: floatPtr(3.14),
		Width: floatPtr(3.14),
	}

	// Test marshaling
	xmlData, err := xml.Marshal(original)
	if err != nil {
		t.Fatalf("Failed to marshal XML: %v", err)
	}

	// Test unmarshaling
	var unmarshaled BodyLDLeftPowerRail
	err = xml.Unmarshal(xmlData, &unmarshaled)
	if err != nil {
		t.Fatalf("Failed to unmarshal XML: %v", err)
	}

	// Test validation
	if err := original.Validate(); err != nil {
		t.Errorf("Validation failed: %v", err)
	}
}

func TestBodyLDLeftPowerRailValidation(t *testing.T) {
	// Test valid case
	valid := &BodyLDLeftPowerRail{
		Position: Position{},
		LocalId: 42,
	}
	if err := valid.Validate(); err != nil {
		t.Errorf("Valid object should not have validation errors: %v", err)
	}

	// Test missing required field: Position
	invalidPosition := &BodyLDLeftPowerRail{}
	if err := invalidPosition.Validate(); err == nil {
		t.Error("Missing required field Position should cause validation error")
	}

	// Test missing required field: LocalId
	invalidLocalId := &BodyLDLeftPowerRail{}
	if err := invalidLocalId.Validate(); err == nil {
		t.Error("Missing required field LocalId should cause validation error")
	}

}

func TestBodyLDRightPowerRailXMLMarshaling(t *testing.T) {
	original := &BodyLDRightPowerRail{
		Position: Position{},
		ConnectionPointIn: []ConnectionPointIn{ConnectionPointIn{}},
		Documentation: &FormattedText{},
		LocalId: 42,
		Height: floatPtr(3.14),
		Width: floatPtr(3.14),
	}

	// Test marshaling
	xmlData, err := xml.Marshal(original)
	if err != nil {
		t.Fatalf("Failed to marshal XML: %v", err)
	}

	// Test unmarshaling
	var unmarshaled BodyLDRightPowerRail
	err = xml.Unmarshal(xmlData, &unmarshaled)
	if err != nil {
		t.Fatalf("Failed to unmarshal XML: %v", err)
	}

	// Test validation
	if err := original.Validate(); err != nil {
		t.Errorf("Validation failed: %v", err)
	}
}

func TestBodyLDRightPowerRailValidation(t *testing.T) {
	// Test valid case
	valid := &BodyLDRightPowerRail{
		Position: Position{},
		LocalId: 42,
	}
	if err := valid.Validate(); err != nil {
		t.Errorf("Valid object should not have validation errors: %v", err)
	}

	// Test missing required field: Position
	invalidPosition := &BodyLDRightPowerRail{}
	if err := invalidPosition.Validate(); err == nil {
		t.Error("Missing required field Position should cause validation error")
	}

	// Test missing required field: LocalId
	invalidLocalId := &BodyLDRightPowerRail{}
	if err := invalidLocalId.Validate(); err == nil {
		t.Error("Missing required field LocalId should cause validation error")
	}

}

func TestBodyLDCoilXMLMarshaling(t *testing.T) {
	original := &BodyLDCoil{
		Position: Position{},
		ConnectionPointIn: &ConnectionPointIn{},
		ConnectionPointOut: &ConnectionPointOut{},
		Variable: "test_value",
		Documentation: &FormattedText{},
		LocalId: 42,
		Height: floatPtr(3.14),
		Width: floatPtr(3.14),
		ExecutionOrderId: uintPtr(42),
		Negated: boolPtr(true),
		Edge: func() *EdgeModifierType { v := EdgeModifierType(""); return &v }(),
		Storage: func() *StorageModifierType { v := StorageModifierType(""); return &v }(),
	}

	// Test marshaling
	xmlData, err := xml.Marshal(original)
	if err != nil {
		t.Fatalf("Failed to marshal XML: %v", err)
	}

	// Test unmarshaling
	var unmarshaled BodyLDCoil
	err = xml.Unmarshal(xmlData, &unmarshaled)
	if err != nil {
		t.Fatalf("Failed to unmarshal XML: %v", err)
	}

	// Test validation
	if err := original.Validate(); err != nil {
		t.Errorf("Validation failed: %v", err)
	}
}

func TestBodyLDCoilValidation(t *testing.T) {
	// Test valid case
	valid := &BodyLDCoil{
		Position: Position{},
		Variable: "test_value",
		LocalId: 42,
	}
	if err := valid.Validate(); err != nil {
		t.Errorf("Valid object should not have validation errors: %v", err)
	}

	// Test missing required field: Position
	invalidPosition := &BodyLDCoil{}
	if err := invalidPosition.Validate(); err == nil {
		t.Error("Missing required field Position should cause validation error")
	}

	// Test missing required field: Variable
	invalidVariable := &BodyLDCoil{}
	if err := invalidVariable.Validate(); err == nil {
		t.Error("Missing required field Variable should cause validation error")
	}

	// Test missing required field: LocalId
	invalidLocalId := &BodyLDCoil{}
	if err := invalidLocalId.Validate(); err == nil {
		t.Error("Missing required field LocalId should cause validation error")
	}

}

func TestBodyLDContactXMLMarshaling(t *testing.T) {
	original := &BodyLDContact{
		Position: Position{},
		ConnectionPointIn: &ConnectionPointIn{},
		ConnectionPointOut: &ConnectionPointOut{},
		Variable: "test_value",
		Documentation: &FormattedText{},
		LocalId: 42,
		Height: floatPtr(3.14),
		Width: floatPtr(3.14),
		ExecutionOrderId: uintPtr(42),
		Negated: boolPtr(true),
		Edge: func() *EdgeModifierType { v := EdgeModifierType(""); return &v }(),
		Storage: func() *StorageModifierType { v := StorageModifierType(""); return &v }(),
	}

	// Test marshaling
	xmlData, err := xml.Marshal(original)
	if err != nil {
		t.Fatalf("Failed to marshal XML: %v", err)
	}

	// Test unmarshaling
	var unmarshaled BodyLDContact
	err = xml.Unmarshal(xmlData, &unmarshaled)
	if err != nil {
		t.Fatalf("Failed to unmarshal XML: %v", err)
	}

	// Test validation
	if err := original.Validate(); err != nil {
		t.Errorf("Validation failed: %v", err)
	}
}

func TestBodyLDContactValidation(t *testing.T) {
	// Test valid case
	valid := &BodyLDContact{
		Position: Position{},
		Variable: "test_value",
		LocalId: 42,
	}
	if err := valid.Validate(); err != nil {
		t.Errorf("Valid object should not have validation errors: %v", err)
	}

	// Test missing required field: Position
	invalidPosition := &BodyLDContact{}
	if err := invalidPosition.Validate(); err == nil {
		t.Error("Missing required field Position should cause validation error")
	}

	// Test missing required field: Variable
	invalidVariable := &BodyLDContact{}
	if err := invalidVariable.Validate(); err == nil {
		t.Error("Missing required field Variable should cause validation error")
	}

	// Test missing required field: LocalId
	invalidLocalId := &BodyLDContact{}
	if err := invalidLocalId.Validate(); err == nil {
		t.Error("Missing required field LocalId should cause validation error")
	}

}

func TestBodyLDXMLMarshaling(t *testing.T) {
	original := &BodyLD{
		Comment: &BodyLDComment{},
		Error: &BodyLDError{},
		Connector: &BodyLDConnector{},
		Continuation: &BodyLDContinuation{},
		ActionBlock: &BodyLDActionBlock{},
		Block: &BodyLDBlock{},
		InVariable: &BodyLDInVariable{},
		OutVariable: &BodyLDOutVariable{},
		InOutVariable: &BodyLDInOutVariable{},
		Label: &BodyLDLabel{},
		Jump: &BodyLDJump{},
		Return: &BodyLDReturn{},
		LeftPowerRail: &BodyLDLeftPowerRail{},
		RightPowerRail: &BodyLDRightPowerRail{},
		Coil: &BodyLDCoil{},
		Contact: &BodyLDContact{},
	}

	// Test marshaling
	xmlData, err := xml.Marshal(original)
	if err != nil {
		t.Fatalf("Failed to marshal XML: %v", err)
	}

	// Test unmarshaling
	var unmarshaled BodyLD
	err = xml.Unmarshal(xmlData, &unmarshaled)
	if err != nil {
		t.Fatalf("Failed to unmarshal XML: %v", err)
	}

	// Test validation
	if err := original.Validate(); err != nil {
		t.Errorf("Validation failed: %v", err)
	}
}

func TestBodyLDValidation(t *testing.T) {
	// Test valid case
	valid := &BodyLD{
	}
	if err := valid.Validate(); err != nil {
		t.Errorf("Valid object should not have validation errors: %v", err)
	}

}

func TestBodySFCCommentXMLMarshaling(t *testing.T) {
	original := &BodySFCComment{
		Position: Position{},
		Content: FormattedText{},
		Documentation: &FormattedText{},
		LocalId: 42,
		Height: 3.14,
		Width: 3.14,
	}

	// Test marshaling
	xmlData, err := xml.Marshal(original)
	if err != nil {
		t.Fatalf("Failed to marshal XML: %v", err)
	}

	// Test unmarshaling
	var unmarshaled BodySFCComment
	err = xml.Unmarshal(xmlData, &unmarshaled)
	if err != nil {
		t.Fatalf("Failed to unmarshal XML: %v", err)
	}

	// Test validation
	if err := original.Validate(); err != nil {
		t.Errorf("Validation failed: %v", err)
	}
}

func TestBodySFCCommentValidation(t *testing.T) {
	// Test valid case
	valid := &BodySFCComment{
		Position: Position{},
		Content: FormattedText{},
		LocalId: 42,
		Height: 3.14,
		Width: 3.14,
	}
	if err := valid.Validate(); err != nil {
		t.Errorf("Valid object should not have validation errors: %v", err)
	}

	// Test missing required field: Position
	invalidPosition := &BodySFCComment{}
	if err := invalidPosition.Validate(); err == nil {
		t.Error("Missing required field Position should cause validation error")
	}

	// Test missing required field: Content
	invalidContent := &BodySFCComment{}
	if err := invalidContent.Validate(); err == nil {
		t.Error("Missing required field Content should cause validation error")
	}

	// Test missing required field: LocalId
	invalidLocalId := &BodySFCComment{}
	if err := invalidLocalId.Validate(); err == nil {
		t.Error("Missing required field LocalId should cause validation error")
	}

	// Test missing required field: Height
	invalidHeight := &BodySFCComment{}
	if err := invalidHeight.Validate(); err == nil {
		t.Error("Missing required field Height should cause validation error")
	}

	// Test missing required field: Width
	invalidWidth := &BodySFCComment{}
	if err := invalidWidth.Validate(); err == nil {
		t.Error("Missing required field Width should cause validation error")
	}

}

func TestBodySFCErrorXMLMarshaling(t *testing.T) {
	original := &BodySFCError{
		Position: Position{},
		Content: FormattedText{},
		Documentation: &FormattedText{},
		LocalId: 42,
		Height: 3.14,
		Width: 3.14,
	}

	// Test marshaling
	xmlData, err := xml.Marshal(original)
	if err != nil {
		t.Fatalf("Failed to marshal XML: %v", err)
	}

	// Test unmarshaling
	var unmarshaled BodySFCError
	err = xml.Unmarshal(xmlData, &unmarshaled)
	if err != nil {
		t.Fatalf("Failed to unmarshal XML: %v", err)
	}

	// Test validation
	if err := original.Validate(); err != nil {
		t.Errorf("Validation failed: %v", err)
	}
}

func TestBodySFCErrorValidation(t *testing.T) {
	// Test valid case
	valid := &BodySFCError{
		Position: Position{},
		Content: FormattedText{},
		LocalId: 42,
		Height: 3.14,
		Width: 3.14,
	}
	if err := valid.Validate(); err != nil {
		t.Errorf("Valid object should not have validation errors: %v", err)
	}

	// Test missing required field: Position
	invalidPosition := &BodySFCError{}
	if err := invalidPosition.Validate(); err == nil {
		t.Error("Missing required field Position should cause validation error")
	}

	// Test missing required field: Content
	invalidContent := &BodySFCError{}
	if err := invalidContent.Validate(); err == nil {
		t.Error("Missing required field Content should cause validation error")
	}

	// Test missing required field: LocalId
	invalidLocalId := &BodySFCError{}
	if err := invalidLocalId.Validate(); err == nil {
		t.Error("Missing required field LocalId should cause validation error")
	}

	// Test missing required field: Height
	invalidHeight := &BodySFCError{}
	if err := invalidHeight.Validate(); err == nil {
		t.Error("Missing required field Height should cause validation error")
	}

	// Test missing required field: Width
	invalidWidth := &BodySFCError{}
	if err := invalidWidth.Validate(); err == nil {
		t.Error("Missing required field Width should cause validation error")
	}

}

func TestBodySFCConnectorXMLMarshaling(t *testing.T) {
	original := &BodySFCConnector{
		Position: Position{},
		ConnectionPointIn: &ConnectionPointIn{},
		Documentation: &FormattedText{},
		Name: "test_value",
		LocalId: 42,
		Height: floatPtr(3.14),
		Width: floatPtr(3.14),
	}

	// Test marshaling
	xmlData, err := xml.Marshal(original)
	if err != nil {
		t.Fatalf("Failed to marshal XML: %v", err)
	}

	// Test unmarshaling
	var unmarshaled BodySFCConnector
	err = xml.Unmarshal(xmlData, &unmarshaled)
	if err != nil {
		t.Fatalf("Failed to unmarshal XML: %v", err)
	}

	// Test validation
	if err := original.Validate(); err != nil {
		t.Errorf("Validation failed: %v", err)
	}
}

func TestBodySFCConnectorValidation(t *testing.T) {
	// Test valid case
	valid := &BodySFCConnector{
		Position: Position{},
		Name: "test_value",
		LocalId: 42,
	}
	if err := valid.Validate(); err != nil {
		t.Errorf("Valid object should not have validation errors: %v", err)
	}

	// Test missing required field: Position
	invalidPosition := &BodySFCConnector{}
	if err := invalidPosition.Validate(); err == nil {
		t.Error("Missing required field Position should cause validation error")
	}

	// Test missing required field: Name
	invalidName := &BodySFCConnector{}
	if err := invalidName.Validate(); err == nil {
		t.Error("Missing required field Name should cause validation error")
	}

	// Test missing required field: LocalId
	invalidLocalId := &BodySFCConnector{}
	if err := invalidLocalId.Validate(); err == nil {
		t.Error("Missing required field LocalId should cause validation error")
	}

}

func TestBodySFCContinuationXMLMarshaling(t *testing.T) {
	original := &BodySFCContinuation{
		Position: Position{},
		ConnectionPointOut: &ConnectionPointOut{},
		Documentation: &FormattedText{},
		Name: "test_value",
		LocalId: 42,
		Height: floatPtr(3.14),
		Width: floatPtr(3.14),
	}

	// Test marshaling
	xmlData, err := xml.Marshal(original)
	if err != nil {
		t.Fatalf("Failed to marshal XML: %v", err)
	}

	// Test unmarshaling
	var unmarshaled BodySFCContinuation
	err = xml.Unmarshal(xmlData, &unmarshaled)
	if err != nil {
		t.Fatalf("Failed to unmarshal XML: %v", err)
	}

	// Test validation
	if err := original.Validate(); err != nil {
		t.Errorf("Validation failed: %v", err)
	}
}

func TestBodySFCContinuationValidation(t *testing.T) {
	// Test valid case
	valid := &BodySFCContinuation{
		Position: Position{},
		Name: "test_value",
		LocalId: 42,
	}
	if err := valid.Validate(); err != nil {
		t.Errorf("Valid object should not have validation errors: %v", err)
	}

	// Test missing required field: Position
	invalidPosition := &BodySFCContinuation{}
	if err := invalidPosition.Validate(); err == nil {
		t.Error("Missing required field Position should cause validation error")
	}

	// Test missing required field: Name
	invalidName := &BodySFCContinuation{}
	if err := invalidName.Validate(); err == nil {
		t.Error("Missing required field Name should cause validation error")
	}

	// Test missing required field: LocalId
	invalidLocalId := &BodySFCContinuation{}
	if err := invalidLocalId.Validate(); err == nil {
		t.Error("Missing required field LocalId should cause validation error")
	}

}

func TestBodySFCActionBlockActionReferenceXMLMarshaling(t *testing.T) {
	original := &BodySFCActionBlockActionReference{
		Name: "test_value",
	}

	// Test marshaling
	xmlData, err := xml.Marshal(original)
	if err != nil {
		t.Fatalf("Failed to marshal XML: %v", err)
	}

	// Test unmarshaling
	var unmarshaled BodySFCActionBlockActionReference
	err = xml.Unmarshal(xmlData, &unmarshaled)
	if err != nil {
		t.Fatalf("Failed to unmarshal XML: %v", err)
	}

	// Test validation
	if err := original.Validate(); err != nil {
		t.Errorf("Validation failed: %v", err)
	}
}

func TestBodySFCActionBlockActionReferenceValidation(t *testing.T) {
	// Test valid case
	valid := &BodySFCActionBlockActionReference{
		Name: "test_value",
	}
	if err := valid.Validate(); err != nil {
		t.Errorf("Valid object should not have validation errors: %v", err)
	}

	// Test missing required field: Name
	invalidName := &BodySFCActionBlockActionReference{}
	if err := invalidName.Validate(); err == nil {
		t.Error("Missing required field Name should cause validation error")
	}

}

func TestBodySFCActionBlockActionXMLMarshaling(t *testing.T) {
	original := &BodySFCActionBlockAction{
		Reference: &BodySFCActionBlockActionReference{},
		Inline: &Body{},
		Documentation: &FormattedText{},
		Qualifier: stringPtr("test_value"),
		Duration: stringPtr("test_value"),
		Indicator: stringPtr("test_value"),
	}

	// Test marshaling
	xmlData, err := xml.Marshal(original)
	if err != nil {
		t.Fatalf("Failed to marshal XML: %v", err)
	}

	// Test unmarshaling
	var unmarshaled BodySFCActionBlockAction
	err = xml.Unmarshal(xmlData, &unmarshaled)
	if err != nil {
		t.Fatalf("Failed to unmarshal XML: %v", err)
	}

	// Test validation
	if err := original.Validate(); err != nil {
		t.Errorf("Validation failed: %v", err)
	}
}

func TestBodySFCActionBlockActionValidation(t *testing.T) {
	// Test valid case
	valid := &BodySFCActionBlockAction{
	}
	if err := valid.Validate(); err != nil {
		t.Errorf("Valid object should not have validation errors: %v", err)
	}

}

func TestBodySFCActionBlockXMLMarshaling(t *testing.T) {
	original := &BodySFCActionBlock{
		Position: Position{},
		ConnectionPointIn: &ConnectionPointIn{},
		ConnectionPointOut: &ConnectionPointOut{},
		Action: []BodySFCActionBlockAction{BodySFCActionBlockAction{}},
		Documentation: &FormattedText{},
		LocalId: 42,
		Height: floatPtr(3.14),
		Width: floatPtr(3.14),
		Negated: boolPtr(true),
	}

	// Test marshaling
	xmlData, err := xml.Marshal(original)
	if err != nil {
		t.Fatalf("Failed to marshal XML: %v", err)
	}

	// Test unmarshaling
	var unmarshaled BodySFCActionBlock
	err = xml.Unmarshal(xmlData, &unmarshaled)
	if err != nil {
		t.Fatalf("Failed to unmarshal XML: %v", err)
	}

	// Test validation
	if err := original.Validate(); err != nil {
		t.Errorf("Validation failed: %v", err)
	}
}

func TestBodySFCActionBlockValidation(t *testing.T) {
	// Test valid case
	valid := &BodySFCActionBlock{
		Position: Position{},
		LocalId: 42,
	}
	if err := valid.Validate(); err != nil {
		t.Errorf("Valid object should not have validation errors: %v", err)
	}

	// Test missing required field: Position
	invalidPosition := &BodySFCActionBlock{}
	if err := invalidPosition.Validate(); err == nil {
		t.Error("Missing required field Position should cause validation error")
	}

	// Test missing required field: LocalId
	invalidLocalId := &BodySFCActionBlock{}
	if err := invalidLocalId.Validate(); err == nil {
		t.Error("Missing required field LocalId should cause validation error")
	}

}

func TestBodySFCBlockInputVariablesVariableXMLMarshaling(t *testing.T) {
	original := &BodySFCBlockInputVariablesVariable{
		ConnectionPointIn: ConnectionPointIn{},
		Documentation: &FormattedText{},
		FormalParameter: "test_value",
		Negated: boolPtr(true),
		Edge: func() *EdgeModifierType { v := EdgeModifierType(""); return &v }(),
		Storage: func() *StorageModifierType { v := StorageModifierType(""); return &v }(),
		Hidden: boolPtr(true),
	}

	// Test marshaling
	xmlData, err := xml.Marshal(original)
	if err != nil {
		t.Fatalf("Failed to marshal XML: %v", err)
	}

	// Test unmarshaling
	var unmarshaled BodySFCBlockInputVariablesVariable
	err = xml.Unmarshal(xmlData, &unmarshaled)
	if err != nil {
		t.Fatalf("Failed to unmarshal XML: %v", err)
	}

	// Test validation
	if err := original.Validate(); err != nil {
		t.Errorf("Validation failed: %v", err)
	}
}

func TestBodySFCBlockInputVariablesVariableValidation(t *testing.T) {
	// Test valid case
	valid := &BodySFCBlockInputVariablesVariable{
		ConnectionPointIn: ConnectionPointIn{},
		FormalParameter: "test_value",
	}
	if err := valid.Validate(); err != nil {
		t.Errorf("Valid object should not have validation errors: %v", err)
	}

	// Test missing required field: ConnectionPointIn
	invalidConnectionPointIn := &BodySFCBlockInputVariablesVariable{}
	if err := invalidConnectionPointIn.Validate(); err == nil {
		t.Error("Missing required field ConnectionPointIn should cause validation error")
	}

	// Test missing required field: FormalParameter
	invalidFormalParameter := &BodySFCBlockInputVariablesVariable{}
	if err := invalidFormalParameter.Validate(); err == nil {
		t.Error("Missing required field FormalParameter should cause validation error")
	}

}

func TestBodySFCBlockInputVariablesXMLMarshaling(t *testing.T) {
	original := &BodySFCBlockInputVariables{
		Variable: []BodySFCBlockInputVariablesVariable{BodySFCBlockInputVariablesVariable{}},
	}

	// Test marshaling
	xmlData, err := xml.Marshal(original)
	if err != nil {
		t.Fatalf("Failed to marshal XML: %v", err)
	}

	// Test unmarshaling
	var unmarshaled BodySFCBlockInputVariables
	err = xml.Unmarshal(xmlData, &unmarshaled)
	if err != nil {
		t.Fatalf("Failed to unmarshal XML: %v", err)
	}

	// Test validation
	if err := original.Validate(); err != nil {
		t.Errorf("Validation failed: %v", err)
	}
}

func TestBodySFCBlockInputVariablesValidation(t *testing.T) {
	// Test valid case
	valid := &BodySFCBlockInputVariables{
	}
	if err := valid.Validate(); err != nil {
		t.Errorf("Valid object should not have validation errors: %v", err)
	}

}

func TestBodySFCBlockInOutVariablesVariableXMLMarshaling(t *testing.T) {
	original := &BodySFCBlockInOutVariablesVariable{
		ConnectionPointIn: &ConnectionPointIn{},
		ConnectionPointOut: &ConnectionPointOut{},
		Documentation: &FormattedText{},
		FormalParameter: "test_value",
		Negated: boolPtr(true),
		Edge: func() *EdgeModifierType { v := EdgeModifierType(""); return &v }(),
		Storage: func() *StorageModifierType { v := StorageModifierType(""); return &v }(),
		Hidden: boolPtr(true),
	}

	// Test marshaling
	xmlData, err := xml.Marshal(original)
	if err != nil {
		t.Fatalf("Failed to marshal XML: %v", err)
	}

	// Test unmarshaling
	var unmarshaled BodySFCBlockInOutVariablesVariable
	err = xml.Unmarshal(xmlData, &unmarshaled)
	if err != nil {
		t.Fatalf("Failed to unmarshal XML: %v", err)
	}

	// Test validation
	if err := original.Validate(); err != nil {
		t.Errorf("Validation failed: %v", err)
	}
}

func TestBodySFCBlockInOutVariablesVariableValidation(t *testing.T) {
	// Test valid case
	valid := &BodySFCBlockInOutVariablesVariable{
		FormalParameter: "test_value",
	}
	if err := valid.Validate(); err != nil {
		t.Errorf("Valid object should not have validation errors: %v", err)
	}

	// Test missing required field: FormalParameter
	invalidFormalParameter := &BodySFCBlockInOutVariablesVariable{}
	if err := invalidFormalParameter.Validate(); err == nil {
		t.Error("Missing required field FormalParameter should cause validation error")
	}

}

func TestBodySFCBlockInOutVariablesXMLMarshaling(t *testing.T) {
	original := &BodySFCBlockInOutVariables{
		Variable: []BodySFCBlockInOutVariablesVariable{BodySFCBlockInOutVariablesVariable{}},
	}

	// Test marshaling
	xmlData, err := xml.Marshal(original)
	if err != nil {
		t.Fatalf("Failed to marshal XML: %v", err)
	}

	// Test unmarshaling
	var unmarshaled BodySFCBlockInOutVariables
	err = xml.Unmarshal(xmlData, &unmarshaled)
	if err != nil {
		t.Fatalf("Failed to unmarshal XML: %v", err)
	}

	// Test validation
	if err := original.Validate(); err != nil {
		t.Errorf("Validation failed: %v", err)
	}
}

func TestBodySFCBlockInOutVariablesValidation(t *testing.T) {
	// Test valid case
	valid := &BodySFCBlockInOutVariables{
	}
	if err := valid.Validate(); err != nil {
		t.Errorf("Valid object should not have validation errors: %v", err)
	}

}

func TestBodySFCBlockOutputVariablesVariableXMLMarshaling(t *testing.T) {
	original := &BodySFCBlockOutputVariablesVariable{
		ConnectionPointOut: &ConnectionPointOut{},
		Documentation: &FormattedText{},
		FormalParameter: "test_value",
		Negated: boolPtr(true),
		Edge: func() *EdgeModifierType { v := EdgeModifierType(""); return &v }(),
		Storage: func() *StorageModifierType { v := StorageModifierType(""); return &v }(),
		Hidden: boolPtr(true),
	}

	// Test marshaling
	xmlData, err := xml.Marshal(original)
	if err != nil {
		t.Fatalf("Failed to marshal XML: %v", err)
	}

	// Test unmarshaling
	var unmarshaled BodySFCBlockOutputVariablesVariable
	err = xml.Unmarshal(xmlData, &unmarshaled)
	if err != nil {
		t.Fatalf("Failed to unmarshal XML: %v", err)
	}

	// Test validation
	if err := original.Validate(); err != nil {
		t.Errorf("Validation failed: %v", err)
	}
}

func TestBodySFCBlockOutputVariablesVariableValidation(t *testing.T) {
	// Test valid case
	valid := &BodySFCBlockOutputVariablesVariable{
		FormalParameter: "test_value",
	}
	if err := valid.Validate(); err != nil {
		t.Errorf("Valid object should not have validation errors: %v", err)
	}

	// Test missing required field: FormalParameter
	invalidFormalParameter := &BodySFCBlockOutputVariablesVariable{}
	if err := invalidFormalParameter.Validate(); err == nil {
		t.Error("Missing required field FormalParameter should cause validation error")
	}

}

func TestBodySFCBlockOutputVariablesXMLMarshaling(t *testing.T) {
	original := &BodySFCBlockOutputVariables{
		Variable: []BodySFCBlockOutputVariablesVariable{BodySFCBlockOutputVariablesVariable{}},
	}

	// Test marshaling
	xmlData, err := xml.Marshal(original)
	if err != nil {
		t.Fatalf("Failed to marshal XML: %v", err)
	}

	// Test unmarshaling
	var unmarshaled BodySFCBlockOutputVariables
	err = xml.Unmarshal(xmlData, &unmarshaled)
	if err != nil {
		t.Fatalf("Failed to unmarshal XML: %v", err)
	}

	// Test validation
	if err := original.Validate(); err != nil {
		t.Errorf("Validation failed: %v", err)
	}
}

func TestBodySFCBlockOutputVariablesValidation(t *testing.T) {
	// Test valid case
	valid := &BodySFCBlockOutputVariables{
	}
	if err := valid.Validate(); err != nil {
		t.Errorf("Valid object should not have validation errors: %v", err)
	}

}

func TestBodySFCBlockXMLMarshaling(t *testing.T) {
	original := &BodySFCBlock{
		Position: Position{},
		InputVariables: BodySFCBlockInputVariables{},
		InOutVariables: BodySFCBlockInOutVariables{},
		OutputVariables: BodySFCBlockOutputVariables{},
		Documentation: &FormattedText{},
		LocalId: 42,
		Width: floatPtr(3.14),
		Height: floatPtr(3.14),
		TypeName: "test_value",
		InstanceName: stringPtr("test_value"),
		ExecutionOrderId: uintPtr(42),
	}

	// Test marshaling
	xmlData, err := xml.Marshal(original)
	if err != nil {
		t.Fatalf("Failed to marshal XML: %v", err)
	}

	// Test unmarshaling
	var unmarshaled BodySFCBlock
	err = xml.Unmarshal(xmlData, &unmarshaled)
	if err != nil {
		t.Fatalf("Failed to unmarshal XML: %v", err)
	}

	// Test validation
	if err := original.Validate(); err != nil {
		t.Errorf("Validation failed: %v", err)
	}
}

func TestBodySFCBlockValidation(t *testing.T) {
	// Test valid case
	valid := &BodySFCBlock{
		Position: Position{},
		InputVariables: BodySFCBlockInputVariables{},
		InOutVariables: BodySFCBlockInOutVariables{},
		OutputVariables: BodySFCBlockOutputVariables{},
		LocalId: 42,
		TypeName: "test_value",
	}
	if err := valid.Validate(); err != nil {
		t.Errorf("Valid object should not have validation errors: %v", err)
	}

	// Test missing required field: Position
	invalidPosition := &BodySFCBlock{}
	if err := invalidPosition.Validate(); err == nil {
		t.Error("Missing required field Position should cause validation error")
	}

	// Test missing required field: InputVariables
	invalidInputVariables := &BodySFCBlock{}
	if err := invalidInputVariables.Validate(); err == nil {
		t.Error("Missing required field InputVariables should cause validation error")
	}

	// Test missing required field: InOutVariables
	invalidInOutVariables := &BodySFCBlock{}
	if err := invalidInOutVariables.Validate(); err == nil {
		t.Error("Missing required field InOutVariables should cause validation error")
	}

	// Test missing required field: OutputVariables
	invalidOutputVariables := &BodySFCBlock{}
	if err := invalidOutputVariables.Validate(); err == nil {
		t.Error("Missing required field OutputVariables should cause validation error")
	}

	// Test missing required field: LocalId
	invalidLocalId := &BodySFCBlock{}
	if err := invalidLocalId.Validate(); err == nil {
		t.Error("Missing required field LocalId should cause validation error")
	}

	// Test missing required field: TypeName
	invalidTypeName := &BodySFCBlock{}
	if err := invalidTypeName.Validate(); err == nil {
		t.Error("Missing required field TypeName should cause validation error")
	}

}

func TestBodySFCInVariableXMLMarshaling(t *testing.T) {
	original := &BodySFCInVariable{
		Position: Position{},
		ConnectionPointOut: &ConnectionPointOut{},
		Expression: "test_value",
		Documentation: &FormattedText{},
		LocalId: 42,
		Height: floatPtr(3.14),
		Width: floatPtr(3.14),
		ExecutionOrderId: uintPtr(42),
		Negated: boolPtr(true),
		Edge: func() *EdgeModifierType { v := EdgeModifierType(""); return &v }(),
		Storage: func() *StorageModifierType { v := StorageModifierType(""); return &v }(),
	}

	// Test marshaling
	xmlData, err := xml.Marshal(original)
	if err != nil {
		t.Fatalf("Failed to marshal XML: %v", err)
	}

	// Test unmarshaling
	var unmarshaled BodySFCInVariable
	err = xml.Unmarshal(xmlData, &unmarshaled)
	if err != nil {
		t.Fatalf("Failed to unmarshal XML: %v", err)
	}

	// Test validation
	if err := original.Validate(); err != nil {
		t.Errorf("Validation failed: %v", err)
	}
}

func TestBodySFCInVariableValidation(t *testing.T) {
	// Test valid case
	valid := &BodySFCInVariable{
		Position: Position{},
		Expression: "test_value",
		LocalId: 42,
	}
	if err := valid.Validate(); err != nil {
		t.Errorf("Valid object should not have validation errors: %v", err)
	}

	// Test missing required field: Position
	invalidPosition := &BodySFCInVariable{}
	if err := invalidPosition.Validate(); err == nil {
		t.Error("Missing required field Position should cause validation error")
	}

	// Test missing required field: Expression
	invalidExpression := &BodySFCInVariable{}
	if err := invalidExpression.Validate(); err == nil {
		t.Error("Missing required field Expression should cause validation error")
	}

	// Test missing required field: LocalId
	invalidLocalId := &BodySFCInVariable{}
	if err := invalidLocalId.Validate(); err == nil {
		t.Error("Missing required field LocalId should cause validation error")
	}

}

func TestBodySFCOutVariableXMLMarshaling(t *testing.T) {
	original := &BodySFCOutVariable{
		Position: Position{},
		ConnectionPointIn: &ConnectionPointIn{},
		Expression: "test_value",
		Documentation: &FormattedText{},
		LocalId: 42,
		Height: floatPtr(3.14),
		Width: floatPtr(3.14),
		ExecutionOrderId: uintPtr(42),
		Negated: boolPtr(true),
		Edge: func() *EdgeModifierType { v := EdgeModifierType(""); return &v }(),
		Storage: func() *StorageModifierType { v := StorageModifierType(""); return &v }(),
	}

	// Test marshaling
	xmlData, err := xml.Marshal(original)
	if err != nil {
		t.Fatalf("Failed to marshal XML: %v", err)
	}

	// Test unmarshaling
	var unmarshaled BodySFCOutVariable
	err = xml.Unmarshal(xmlData, &unmarshaled)
	if err != nil {
		t.Fatalf("Failed to unmarshal XML: %v", err)
	}

	// Test validation
	if err := original.Validate(); err != nil {
		t.Errorf("Validation failed: %v", err)
	}
}

func TestBodySFCOutVariableValidation(t *testing.T) {
	// Test valid case
	valid := &BodySFCOutVariable{
		Position: Position{},
		Expression: "test_value",
		LocalId: 42,
	}
	if err := valid.Validate(); err != nil {
		t.Errorf("Valid object should not have validation errors: %v", err)
	}

	// Test missing required field: Position
	invalidPosition := &BodySFCOutVariable{}
	if err := invalidPosition.Validate(); err == nil {
		t.Error("Missing required field Position should cause validation error")
	}

	// Test missing required field: Expression
	invalidExpression := &BodySFCOutVariable{}
	if err := invalidExpression.Validate(); err == nil {
		t.Error("Missing required field Expression should cause validation error")
	}

	// Test missing required field: LocalId
	invalidLocalId := &BodySFCOutVariable{}
	if err := invalidLocalId.Validate(); err == nil {
		t.Error("Missing required field LocalId should cause validation error")
	}

}

func TestBodySFCInOutVariableXMLMarshaling(t *testing.T) {
	original := &BodySFCInOutVariable{
		Position: Position{},
		ConnectionPointIn: &ConnectionPointIn{},
		ConnectionPointOut: &ConnectionPointOut{},
		Expression: "test_value",
		Documentation: &FormattedText{},
		LocalId: 42,
		Height: floatPtr(3.14),
		Width: floatPtr(3.14),
		ExecutionOrderId: uintPtr(42),
		NegatedIn: boolPtr(true),
		EdgeIn: func() *EdgeModifierType { v := EdgeModifierType(""); return &v }(),
		StorageIn: func() *StorageModifierType { v := StorageModifierType(""); return &v }(),
		NegatedOut: boolPtr(true),
		EdgeOut: func() *EdgeModifierType { v := EdgeModifierType(""); return &v }(),
		StorageOut: func() *StorageModifierType { v := StorageModifierType(""); return &v }(),
	}

	// Test marshaling
	xmlData, err := xml.Marshal(original)
	if err != nil {
		t.Fatalf("Failed to marshal XML: %v", err)
	}

	// Test unmarshaling
	var unmarshaled BodySFCInOutVariable
	err = xml.Unmarshal(xmlData, &unmarshaled)
	if err != nil {
		t.Fatalf("Failed to unmarshal XML: %v", err)
	}

	// Test validation
	if err := original.Validate(); err != nil {
		t.Errorf("Validation failed: %v", err)
	}
}

func TestBodySFCInOutVariableValidation(t *testing.T) {
	// Test valid case
	valid := &BodySFCInOutVariable{
		Position: Position{},
		Expression: "test_value",
		LocalId: 42,
	}
	if err := valid.Validate(); err != nil {
		t.Errorf("Valid object should not have validation errors: %v", err)
	}

	// Test missing required field: Position
	invalidPosition := &BodySFCInOutVariable{}
	if err := invalidPosition.Validate(); err == nil {
		t.Error("Missing required field Position should cause validation error")
	}

	// Test missing required field: Expression
	invalidExpression := &BodySFCInOutVariable{}
	if err := invalidExpression.Validate(); err == nil {
		t.Error("Missing required field Expression should cause validation error")
	}

	// Test missing required field: LocalId
	invalidLocalId := &BodySFCInOutVariable{}
	if err := invalidLocalId.Validate(); err == nil {
		t.Error("Missing required field LocalId should cause validation error")
	}

}

func TestBodySFCLabelXMLMarshaling(t *testing.T) {
	original := &BodySFCLabel{
		Position: Position{},
		Documentation: &FormattedText{},
		LocalId: 42,
		Height: floatPtr(3.14),
		Width: floatPtr(3.14),
		Label: "test_value",
		ExecutionOrderId: uintPtr(42),
	}

	// Test marshaling
	xmlData, err := xml.Marshal(original)
	if err != nil {
		t.Fatalf("Failed to marshal XML: %v", err)
	}

	// Test unmarshaling
	var unmarshaled BodySFCLabel
	err = xml.Unmarshal(xmlData, &unmarshaled)
	if err != nil {
		t.Fatalf("Failed to unmarshal XML: %v", err)
	}

	// Test validation
	if err := original.Validate(); err != nil {
		t.Errorf("Validation failed: %v", err)
	}
}

func TestBodySFCLabelValidation(t *testing.T) {
	// Test valid case
	valid := &BodySFCLabel{
		Position: Position{},
		LocalId: 42,
		Label: "test_value",
	}
	if err := valid.Validate(); err != nil {
		t.Errorf("Valid object should not have validation errors: %v", err)
	}

	// Test missing required field: Position
	invalidPosition := &BodySFCLabel{}
	if err := invalidPosition.Validate(); err == nil {
		t.Error("Missing required field Position should cause validation error")
	}

	// Test missing required field: LocalId
	invalidLocalId := &BodySFCLabel{}
	if err := invalidLocalId.Validate(); err == nil {
		t.Error("Missing required field LocalId should cause validation error")
	}

	// Test missing required field: Label
	invalidLabel := &BodySFCLabel{}
	if err := invalidLabel.Validate(); err == nil {
		t.Error("Missing required field Label should cause validation error")
	}

}

func TestBodySFCJumpXMLMarshaling(t *testing.T) {
	original := &BodySFCJump{
		Position: Position{},
		ConnectionPointIn: &ConnectionPointIn{},
		Documentation: &FormattedText{},
		LocalId: 42,
		Height: floatPtr(3.14),
		Width: floatPtr(3.14),
		Label: "test_value",
		ExecutionOrderId: uintPtr(42),
	}

	// Test marshaling
	xmlData, err := xml.Marshal(original)
	if err != nil {
		t.Fatalf("Failed to marshal XML: %v", err)
	}

	// Test unmarshaling
	var unmarshaled BodySFCJump
	err = xml.Unmarshal(xmlData, &unmarshaled)
	if err != nil {
		t.Fatalf("Failed to unmarshal XML: %v", err)
	}

	// Test validation
	if err := original.Validate(); err != nil {
		t.Errorf("Validation failed: %v", err)
	}
}

func TestBodySFCJumpValidation(t *testing.T) {
	// Test valid case
	valid := &BodySFCJump{
		Position: Position{},
		LocalId: 42,
		Label: "test_value",
	}
	if err := valid.Validate(); err != nil {
		t.Errorf("Valid object should not have validation errors: %v", err)
	}

	// Test missing required field: Position
	invalidPosition := &BodySFCJump{}
	if err := invalidPosition.Validate(); err == nil {
		t.Error("Missing required field Position should cause validation error")
	}

	// Test missing required field: LocalId
	invalidLocalId := &BodySFCJump{}
	if err := invalidLocalId.Validate(); err == nil {
		t.Error("Missing required field LocalId should cause validation error")
	}

	// Test missing required field: Label
	invalidLabel := &BodySFCJump{}
	if err := invalidLabel.Validate(); err == nil {
		t.Error("Missing required field Label should cause validation error")
	}

}

func TestBodySFCReturnXMLMarshaling(t *testing.T) {
	original := &BodySFCReturn{
		Position: Position{},
		ConnectionPointIn: &ConnectionPointIn{},
		Documentation: &FormattedText{},
		LocalId: 42,
		Height: floatPtr(3.14),
		Width: floatPtr(3.14),
		ExecutionOrderId: uintPtr(42),
	}

	// Test marshaling
	xmlData, err := xml.Marshal(original)
	if err != nil {
		t.Fatalf("Failed to marshal XML: %v", err)
	}

	// Test unmarshaling
	var unmarshaled BodySFCReturn
	err = xml.Unmarshal(xmlData, &unmarshaled)
	if err != nil {
		t.Fatalf("Failed to unmarshal XML: %v", err)
	}

	// Test validation
	if err := original.Validate(); err != nil {
		t.Errorf("Validation failed: %v", err)
	}
}

func TestBodySFCReturnValidation(t *testing.T) {
	// Test valid case
	valid := &BodySFCReturn{
		Position: Position{},
		LocalId: 42,
	}
	if err := valid.Validate(); err != nil {
		t.Errorf("Valid object should not have validation errors: %v", err)
	}

	// Test missing required field: Position
	invalidPosition := &BodySFCReturn{}
	if err := invalidPosition.Validate(); err == nil {
		t.Error("Missing required field Position should cause validation error")
	}

	// Test missing required field: LocalId
	invalidLocalId := &BodySFCReturn{}
	if err := invalidLocalId.Validate(); err == nil {
		t.Error("Missing required field LocalId should cause validation error")
	}

}

func TestBodySFCLeftPowerRailConnectionPointOutXMLMarshaling(t *testing.T) {
	original := &BodySFCLeftPowerRailConnectionPointOut{
	}

	// Test marshaling
	xmlData, err := xml.Marshal(original)
	if err != nil {
		t.Fatalf("Failed to marshal XML: %v", err)
	}

	// Test unmarshaling
	var unmarshaled BodySFCLeftPowerRailConnectionPointOut
	err = xml.Unmarshal(xmlData, &unmarshaled)
	if err != nil {
		t.Fatalf("Failed to unmarshal XML: %v", err)
	}

	// Test validation
	if err := original.Validate(); err != nil {
		t.Errorf("Validation failed: %v", err)
	}
}

func TestBodySFCLeftPowerRailConnectionPointOutValidation(t *testing.T) {
	// Test valid case
	valid := &BodySFCLeftPowerRailConnectionPointOut{
	}
	if err := valid.Validate(); err != nil {
		t.Errorf("Valid object should not have validation errors: %v", err)
	}

}

func TestBodySFCLeftPowerRailXMLMarshaling(t *testing.T) {
	original := &BodySFCLeftPowerRail{
		Position: Position{},
		ConnectionPointOut: []BodySFCLeftPowerRailConnectionPointOut{BodySFCLeftPowerRailConnectionPointOut{}},
		Documentation: &FormattedText{},
		LocalId: 42,
		Height: floatPtr(3.14),
		Width: floatPtr(3.14),
	}

	// Test marshaling
	xmlData, err := xml.Marshal(original)
	if err != nil {
		t.Fatalf("Failed to marshal XML: %v", err)
	}

	// Test unmarshaling
	var unmarshaled BodySFCLeftPowerRail
	err = xml.Unmarshal(xmlData, &unmarshaled)
	if err != nil {
		t.Fatalf("Failed to unmarshal XML: %v", err)
	}

	// Test validation
	if err := original.Validate(); err != nil {
		t.Errorf("Validation failed: %v", err)
	}
}

func TestBodySFCLeftPowerRailValidation(t *testing.T) {
	// Test valid case
	valid := &BodySFCLeftPowerRail{
		Position: Position{},
		LocalId: 42,
	}
	if err := valid.Validate(); err != nil {
		t.Errorf("Valid object should not have validation errors: %v", err)
	}

	// Test missing required field: Position
	invalidPosition := &BodySFCLeftPowerRail{}
	if err := invalidPosition.Validate(); err == nil {
		t.Error("Missing required field Position should cause validation error")
	}

	// Test missing required field: LocalId
	invalidLocalId := &BodySFCLeftPowerRail{}
	if err := invalidLocalId.Validate(); err == nil {
		t.Error("Missing required field LocalId should cause validation error")
	}

}

func TestBodySFCRightPowerRailXMLMarshaling(t *testing.T) {
	original := &BodySFCRightPowerRail{
		Position: Position{},
		ConnectionPointIn: []ConnectionPointIn{ConnectionPointIn{}},
		Documentation: &FormattedText{},
		LocalId: 42,
		Height: floatPtr(3.14),
		Width: floatPtr(3.14),
	}

	// Test marshaling
	xmlData, err := xml.Marshal(original)
	if err != nil {
		t.Fatalf("Failed to marshal XML: %v", err)
	}

	// Test unmarshaling
	var unmarshaled BodySFCRightPowerRail
	err = xml.Unmarshal(xmlData, &unmarshaled)
	if err != nil {
		t.Fatalf("Failed to unmarshal XML: %v", err)
	}

	// Test validation
	if err := original.Validate(); err != nil {
		t.Errorf("Validation failed: %v", err)
	}
}

func TestBodySFCRightPowerRailValidation(t *testing.T) {
	// Test valid case
	valid := &BodySFCRightPowerRail{
		Position: Position{},
		LocalId: 42,
	}
	if err := valid.Validate(); err != nil {
		t.Errorf("Valid object should not have validation errors: %v", err)
	}

	// Test missing required field: Position
	invalidPosition := &BodySFCRightPowerRail{}
	if err := invalidPosition.Validate(); err == nil {
		t.Error("Missing required field Position should cause validation error")
	}

	// Test missing required field: LocalId
	invalidLocalId := &BodySFCRightPowerRail{}
	if err := invalidLocalId.Validate(); err == nil {
		t.Error("Missing required field LocalId should cause validation error")
	}

}

func TestBodySFCCoilXMLMarshaling(t *testing.T) {
	original := &BodySFCCoil{
		Position: Position{},
		ConnectionPointIn: &ConnectionPointIn{},
		ConnectionPointOut: &ConnectionPointOut{},
		Variable: "test_value",
		Documentation: &FormattedText{},
		LocalId: 42,
		Height: floatPtr(3.14),
		Width: floatPtr(3.14),
		ExecutionOrderId: uintPtr(42),
		Negated: boolPtr(true),
		Edge: func() *EdgeModifierType { v := EdgeModifierType(""); return &v }(),
		Storage: func() *StorageModifierType { v := StorageModifierType(""); return &v }(),
	}

	// Test marshaling
	xmlData, err := xml.Marshal(original)
	if err != nil {
		t.Fatalf("Failed to marshal XML: %v", err)
	}

	// Test unmarshaling
	var unmarshaled BodySFCCoil
	err = xml.Unmarshal(xmlData, &unmarshaled)
	if err != nil {
		t.Fatalf("Failed to unmarshal XML: %v", err)
	}

	// Test validation
	if err := original.Validate(); err != nil {
		t.Errorf("Validation failed: %v", err)
	}
}

func TestBodySFCCoilValidation(t *testing.T) {
	// Test valid case
	valid := &BodySFCCoil{
		Position: Position{},
		Variable: "test_value",
		LocalId: 42,
	}
	if err := valid.Validate(); err != nil {
		t.Errorf("Valid object should not have validation errors: %v", err)
	}

	// Test missing required field: Position
	invalidPosition := &BodySFCCoil{}
	if err := invalidPosition.Validate(); err == nil {
		t.Error("Missing required field Position should cause validation error")
	}

	// Test missing required field: Variable
	invalidVariable := &BodySFCCoil{}
	if err := invalidVariable.Validate(); err == nil {
		t.Error("Missing required field Variable should cause validation error")
	}

	// Test missing required field: LocalId
	invalidLocalId := &BodySFCCoil{}
	if err := invalidLocalId.Validate(); err == nil {
		t.Error("Missing required field LocalId should cause validation error")
	}

}

func TestBodySFCContactXMLMarshaling(t *testing.T) {
	original := &BodySFCContact{
		Position: Position{},
		ConnectionPointIn: &ConnectionPointIn{},
		ConnectionPointOut: &ConnectionPointOut{},
		Variable: "test_value",
		Documentation: &FormattedText{},
		LocalId: 42,
		Height: floatPtr(3.14),
		Width: floatPtr(3.14),
		ExecutionOrderId: uintPtr(42),
		Negated: boolPtr(true),
		Edge: func() *EdgeModifierType { v := EdgeModifierType(""); return &v }(),
		Storage: func() *StorageModifierType { v := StorageModifierType(""); return &v }(),
	}

	// Test marshaling
	xmlData, err := xml.Marshal(original)
	if err != nil {
		t.Fatalf("Failed to marshal XML: %v", err)
	}

	// Test unmarshaling
	var unmarshaled BodySFCContact
	err = xml.Unmarshal(xmlData, &unmarshaled)
	if err != nil {
		t.Fatalf("Failed to unmarshal XML: %v", err)
	}

	// Test validation
	if err := original.Validate(); err != nil {
		t.Errorf("Validation failed: %v", err)
	}
}

func TestBodySFCContactValidation(t *testing.T) {
	// Test valid case
	valid := &BodySFCContact{
		Position: Position{},
		Variable: "test_value",
		LocalId: 42,
	}
	if err := valid.Validate(); err != nil {
		t.Errorf("Valid object should not have validation errors: %v", err)
	}

	// Test missing required field: Position
	invalidPosition := &BodySFCContact{}
	if err := invalidPosition.Validate(); err == nil {
		t.Error("Missing required field Position should cause validation error")
	}

	// Test missing required field: Variable
	invalidVariable := &BodySFCContact{}
	if err := invalidVariable.Validate(); err == nil {
		t.Error("Missing required field Variable should cause validation error")
	}

	// Test missing required field: LocalId
	invalidLocalId := &BodySFCContact{}
	if err := invalidLocalId.Validate(); err == nil {
		t.Error("Missing required field LocalId should cause validation error")
	}

}

func TestBodySFCStepConnectionPointOutXMLMarshaling(t *testing.T) {
	original := &BodySFCStepConnectionPointOut{
	}

	// Test marshaling
	xmlData, err := xml.Marshal(original)
	if err != nil {
		t.Fatalf("Failed to marshal XML: %v", err)
	}

	// Test unmarshaling
	var unmarshaled BodySFCStepConnectionPointOut
	err = xml.Unmarshal(xmlData, &unmarshaled)
	if err != nil {
		t.Fatalf("Failed to unmarshal XML: %v", err)
	}

	// Test validation
	if err := original.Validate(); err != nil {
		t.Errorf("Validation failed: %v", err)
	}
}

func TestBodySFCStepConnectionPointOutValidation(t *testing.T) {
	// Test valid case
	valid := &BodySFCStepConnectionPointOut{
	}
	if err := valid.Validate(); err != nil {
		t.Errorf("Valid object should not have validation errors: %v", err)
	}

}

func TestBodySFCStepConnectionPointOutActionXMLMarshaling(t *testing.T) {
	original := &BodySFCStepConnectionPointOutAction{
	}

	// Test marshaling
	xmlData, err := xml.Marshal(original)
	if err != nil {
		t.Fatalf("Failed to marshal XML: %v", err)
	}

	// Test unmarshaling
	var unmarshaled BodySFCStepConnectionPointOutAction
	err = xml.Unmarshal(xmlData, &unmarshaled)
	if err != nil {
		t.Fatalf("Failed to unmarshal XML: %v", err)
	}

	// Test validation
	if err := original.Validate(); err != nil {
		t.Errorf("Validation failed: %v", err)
	}
}

func TestBodySFCStepConnectionPointOutActionValidation(t *testing.T) {
	// Test valid case
	valid := &BodySFCStepConnectionPointOutAction{
	}
	if err := valid.Validate(); err != nil {
		t.Errorf("Valid object should not have validation errors: %v", err)
	}

}

func TestBodySFCStepXMLMarshaling(t *testing.T) {
	original := &BodySFCStep{
		Position: Position{},
		ConnectionPointIn: &ConnectionPointIn{},
		ConnectionPointOut: &BodySFCStepConnectionPointOut{},
		ConnectionPointOutAction: &BodySFCStepConnectionPointOutAction{},
		Documentation: &FormattedText{},
		LocalId: 42,
		Height: floatPtr(3.14),
		Width: floatPtr(3.14),
		Name: "test_value",
		InitialStep: boolPtr(true),
		Negated: boolPtr(true),
	}

	// Test marshaling
	xmlData, err := xml.Marshal(original)
	if err != nil {
		t.Fatalf("Failed to marshal XML: %v", err)
	}

	// Test unmarshaling
	var unmarshaled BodySFCStep
	err = xml.Unmarshal(xmlData, &unmarshaled)
	if err != nil {
		t.Fatalf("Failed to unmarshal XML: %v", err)
	}

	// Test validation
	if err := original.Validate(); err != nil {
		t.Errorf("Validation failed: %v", err)
	}
}

func TestBodySFCStepValidation(t *testing.T) {
	// Test valid case
	valid := &BodySFCStep{
		Position: Position{},
		LocalId: 42,
		Name: "test_value",
	}
	if err := valid.Validate(); err != nil {
		t.Errorf("Valid object should not have validation errors: %v", err)
	}

	// Test missing required field: Position
	invalidPosition := &BodySFCStep{}
	if err := invalidPosition.Validate(); err == nil {
		t.Error("Missing required field Position should cause validation error")
	}

	// Test missing required field: LocalId
	invalidLocalId := &BodySFCStep{}
	if err := invalidLocalId.Validate(); err == nil {
		t.Error("Missing required field LocalId should cause validation error")
	}

	// Test missing required field: Name
	invalidName := &BodySFCStep{}
	if err := invalidName.Validate(); err == nil {
		t.Error("Missing required field Name should cause validation error")
	}

}

func TestBodySFCMacroStepXMLMarshaling(t *testing.T) {
	original := &BodySFCMacroStep{
		Position: Position{},
		ConnectionPointIn: &ConnectionPointIn{},
		ConnectionPointOut: &ConnectionPointOut{},
		Body: &Body{},
		Documentation: &FormattedText{},
		LocalId: 42,
		Height: floatPtr(3.14),
		Width: floatPtr(3.14),
		Name: stringPtr("test_value"),
	}

	// Test marshaling
	xmlData, err := xml.Marshal(original)
	if err != nil {
		t.Fatalf("Failed to marshal XML: %v", err)
	}

	// Test unmarshaling
	var unmarshaled BodySFCMacroStep
	err = xml.Unmarshal(xmlData, &unmarshaled)
	if err != nil {
		t.Fatalf("Failed to unmarshal XML: %v", err)
	}

	// Test validation
	if err := original.Validate(); err != nil {
		t.Errorf("Validation failed: %v", err)
	}
}

func TestBodySFCMacroStepValidation(t *testing.T) {
	// Test valid case
	valid := &BodySFCMacroStep{
		Position: Position{},
		LocalId: 42,
	}
	if err := valid.Validate(); err != nil {
		t.Errorf("Valid object should not have validation errors: %v", err)
	}

	// Test missing required field: Position
	invalidPosition := &BodySFCMacroStep{}
	if err := invalidPosition.Validate(); err == nil {
		t.Error("Missing required field Position should cause validation error")
	}

	// Test missing required field: LocalId
	invalidLocalId := &BodySFCMacroStep{}
	if err := invalidLocalId.Validate(); err == nil {
		t.Error("Missing required field LocalId should cause validation error")
	}

}

func TestBodySFCJumpStepXMLMarshaling(t *testing.T) {
	original := &BodySFCJumpStep{
		Position: Position{},
		ConnectionPointIn: &ConnectionPointIn{},
		Documentation: &FormattedText{},
		LocalId: 42,
		Height: floatPtr(3.14),
		Width: floatPtr(3.14),
		TargetName: "test_value",
	}

	// Test marshaling
	xmlData, err := xml.Marshal(original)
	if err != nil {
		t.Fatalf("Failed to marshal XML: %v", err)
	}

	// Test unmarshaling
	var unmarshaled BodySFCJumpStep
	err = xml.Unmarshal(xmlData, &unmarshaled)
	if err != nil {
		t.Fatalf("Failed to unmarshal XML: %v", err)
	}

	// Test validation
	if err := original.Validate(); err != nil {
		t.Errorf("Validation failed: %v", err)
	}
}

func TestBodySFCJumpStepValidation(t *testing.T) {
	// Test valid case
	valid := &BodySFCJumpStep{
		Position: Position{},
		LocalId: 42,
		TargetName: "test_value",
	}
	if err := valid.Validate(); err != nil {
		t.Errorf("Valid object should not have validation errors: %v", err)
	}

	// Test missing required field: Position
	invalidPosition := &BodySFCJumpStep{}
	if err := invalidPosition.Validate(); err == nil {
		t.Error("Missing required field Position should cause validation error")
	}

	// Test missing required field: LocalId
	invalidLocalId := &BodySFCJumpStep{}
	if err := invalidLocalId.Validate(); err == nil {
		t.Error("Missing required field LocalId should cause validation error")
	}

	// Test missing required field: TargetName
	invalidTargetName := &BodySFCJumpStep{}
	if err := invalidTargetName.Validate(); err == nil {
		t.Error("Missing required field TargetName should cause validation error")
	}

}

func TestBodySFCTransitionConditionReferenceXMLMarshaling(t *testing.T) {
	original := &BodySFCTransitionConditionReference{
		Name: "test_value",
	}

	// Test marshaling
	xmlData, err := xml.Marshal(original)
	if err != nil {
		t.Fatalf("Failed to marshal XML: %v", err)
	}

	// Test unmarshaling
	var unmarshaled BodySFCTransitionConditionReference
	err = xml.Unmarshal(xmlData, &unmarshaled)
	if err != nil {
		t.Fatalf("Failed to unmarshal XML: %v", err)
	}

	// Test validation
	if err := original.Validate(); err != nil {
		t.Errorf("Validation failed: %v", err)
	}
}

func TestBodySFCTransitionConditionReferenceValidation(t *testing.T) {
	// Test valid case
	valid := &BodySFCTransitionConditionReference{
		Name: "test_value",
	}
	if err := valid.Validate(); err != nil {
		t.Errorf("Valid object should not have validation errors: %v", err)
	}

	// Test missing required field: Name
	invalidName := &BodySFCTransitionConditionReference{}
	if err := invalidName.Validate(); err == nil {
		t.Error("Missing required field Name should cause validation error")
	}

}

func TestBodySFCTransitionConditionInlineXMLMarshaling(t *testing.T) {
	original := &BodySFCTransitionConditionInline{
	}

	// Test marshaling
	xmlData, err := xml.Marshal(original)
	if err != nil {
		t.Fatalf("Failed to marshal XML: %v", err)
	}

	// Test unmarshaling
	var unmarshaled BodySFCTransitionConditionInline
	err = xml.Unmarshal(xmlData, &unmarshaled)
	if err != nil {
		t.Fatalf("Failed to unmarshal XML: %v", err)
	}

	// Test validation
	if err := original.Validate(); err != nil {
		t.Errorf("Validation failed: %v", err)
	}
}

func TestBodySFCTransitionConditionInlineValidation(t *testing.T) {
	// Test valid case
	valid := &BodySFCTransitionConditionInline{
	}
	if err := valid.Validate(); err != nil {
		t.Errorf("Valid object should not have validation errors: %v", err)
	}

}

func TestBodySFCTransitionConditionXMLMarshaling(t *testing.T) {
	original := &BodySFCTransitionCondition{
		Reference: &BodySFCTransitionConditionReference{},
		Connection: []Connection{Connection{}},
		Inline: &BodySFCTransitionConditionInline{},
		Negated: boolPtr(true),
	}

	// Test marshaling
	xmlData, err := xml.Marshal(original)
	if err != nil {
		t.Fatalf("Failed to marshal XML: %v", err)
	}

	// Test unmarshaling
	var unmarshaled BodySFCTransitionCondition
	err = xml.Unmarshal(xmlData, &unmarshaled)
	if err != nil {
		t.Fatalf("Failed to unmarshal XML: %v", err)
	}

	// Test validation
	if err := original.Validate(); err != nil {
		t.Errorf("Validation failed: %v", err)
	}
}

func TestBodySFCTransitionConditionValidation(t *testing.T) {
	// Test valid case
	valid := &BodySFCTransitionCondition{
	}
	if err := valid.Validate(); err != nil {
		t.Errorf("Valid object should not have validation errors: %v", err)
	}

}

func TestBodySFCTransitionXMLMarshaling(t *testing.T) {
	original := &BodySFCTransition{
		Position: Position{},
		ConnectionPointIn: &ConnectionPointIn{},
		ConnectionPointOut: &ConnectionPointOut{},
		Condition: &BodySFCTransitionCondition{},
		Documentation: &FormattedText{},
		LocalId: 42,
		Height: floatPtr(3.14),
		Width: floatPtr(3.14),
		Priority: uintPtr(42),
	}

	// Test marshaling
	xmlData, err := xml.Marshal(original)
	if err != nil {
		t.Fatalf("Failed to marshal XML: %v", err)
	}

	// Test unmarshaling
	var unmarshaled BodySFCTransition
	err = xml.Unmarshal(xmlData, &unmarshaled)
	if err != nil {
		t.Fatalf("Failed to unmarshal XML: %v", err)
	}

	// Test validation
	if err := original.Validate(); err != nil {
		t.Errorf("Validation failed: %v", err)
	}
}

func TestBodySFCTransitionValidation(t *testing.T) {
	// Test valid case
	valid := &BodySFCTransition{
		Position: Position{},
		LocalId: 42,
	}
	if err := valid.Validate(); err != nil {
		t.Errorf("Valid object should not have validation errors: %v", err)
	}

	// Test missing required field: Position
	invalidPosition := &BodySFCTransition{}
	if err := invalidPosition.Validate(); err == nil {
		t.Error("Missing required field Position should cause validation error")
	}

	// Test missing required field: LocalId
	invalidLocalId := &BodySFCTransition{}
	if err := invalidLocalId.Validate(); err == nil {
		t.Error("Missing required field LocalId should cause validation error")
	}

}

func TestBodySFCSelectionDivergenceConnectionPointOutXMLMarshaling(t *testing.T) {
	original := &BodySFCSelectionDivergenceConnectionPointOut{
	}

	// Test marshaling
	xmlData, err := xml.Marshal(original)
	if err != nil {
		t.Fatalf("Failed to marshal XML: %v", err)
	}

	// Test unmarshaling
	var unmarshaled BodySFCSelectionDivergenceConnectionPointOut
	err = xml.Unmarshal(xmlData, &unmarshaled)
	if err != nil {
		t.Fatalf("Failed to unmarshal XML: %v", err)
	}

	// Test validation
	if err := original.Validate(); err != nil {
		t.Errorf("Validation failed: %v", err)
	}
}

func TestBodySFCSelectionDivergenceConnectionPointOutValidation(t *testing.T) {
	// Test valid case
	valid := &BodySFCSelectionDivergenceConnectionPointOut{
	}
	if err := valid.Validate(); err != nil {
		t.Errorf("Valid object should not have validation errors: %v", err)
	}

}

func TestBodySFCSelectionDivergenceXMLMarshaling(t *testing.T) {
	original := &BodySFCSelectionDivergence{
		Position: Position{},
		ConnectionPointIn: &ConnectionPointIn{},
		ConnectionPointOut: []BodySFCSelectionDivergenceConnectionPointOut{BodySFCSelectionDivergenceConnectionPointOut{}},
		Documentation: &FormattedText{},
		LocalId: 42,
		Height: floatPtr(3.14),
		Width: floatPtr(3.14),
	}

	// Test marshaling
	xmlData, err := xml.Marshal(original)
	if err != nil {
		t.Fatalf("Failed to marshal XML: %v", err)
	}

	// Test unmarshaling
	var unmarshaled BodySFCSelectionDivergence
	err = xml.Unmarshal(xmlData, &unmarshaled)
	if err != nil {
		t.Fatalf("Failed to unmarshal XML: %v", err)
	}

	// Test validation
	if err := original.Validate(); err != nil {
		t.Errorf("Validation failed: %v", err)
	}
}

func TestBodySFCSelectionDivergenceValidation(t *testing.T) {
	// Test valid case
	valid := &BodySFCSelectionDivergence{
		Position: Position{},
		LocalId: 42,
	}
	if err := valid.Validate(); err != nil {
		t.Errorf("Valid object should not have validation errors: %v", err)
	}

	// Test missing required field: Position
	invalidPosition := &BodySFCSelectionDivergence{}
	if err := invalidPosition.Validate(); err == nil {
		t.Error("Missing required field Position should cause validation error")
	}

	// Test missing required field: LocalId
	invalidLocalId := &BodySFCSelectionDivergence{}
	if err := invalidLocalId.Validate(); err == nil {
		t.Error("Missing required field LocalId should cause validation error")
	}

}

func TestBodySFCSelectionConvergenceConnectionPointInXMLMarshaling(t *testing.T) {
	original := &BodySFCSelectionConvergenceConnectionPointIn{
	}

	// Test marshaling
	xmlData, err := xml.Marshal(original)
	if err != nil {
		t.Fatalf("Failed to marshal XML: %v", err)
	}

	// Test unmarshaling
	var unmarshaled BodySFCSelectionConvergenceConnectionPointIn
	err = xml.Unmarshal(xmlData, &unmarshaled)
	if err != nil {
		t.Fatalf("Failed to unmarshal XML: %v", err)
	}

	// Test validation
	if err := original.Validate(); err != nil {
		t.Errorf("Validation failed: %v", err)
	}
}

func TestBodySFCSelectionConvergenceConnectionPointInValidation(t *testing.T) {
	// Test valid case
	valid := &BodySFCSelectionConvergenceConnectionPointIn{
	}
	if err := valid.Validate(); err != nil {
		t.Errorf("Valid object should not have validation errors: %v", err)
	}

}

func TestBodySFCSelectionConvergenceXMLMarshaling(t *testing.T) {
	original := &BodySFCSelectionConvergence{
		Position: Position{},
		ConnectionPointIn: []BodySFCSelectionConvergenceConnectionPointIn{BodySFCSelectionConvergenceConnectionPointIn{}},
		ConnectionPointOut: &ConnectionPointOut{},
		Documentation: &FormattedText{},
		LocalId: 42,
		Height: floatPtr(3.14),
		Width: floatPtr(3.14),
	}

	// Test marshaling
	xmlData, err := xml.Marshal(original)
	if err != nil {
		t.Fatalf("Failed to marshal XML: %v", err)
	}

	// Test unmarshaling
	var unmarshaled BodySFCSelectionConvergence
	err = xml.Unmarshal(xmlData, &unmarshaled)
	if err != nil {
		t.Fatalf("Failed to unmarshal XML: %v", err)
	}

	// Test validation
	if err := original.Validate(); err != nil {
		t.Errorf("Validation failed: %v", err)
	}
}

func TestBodySFCSelectionConvergenceValidation(t *testing.T) {
	// Test valid case
	valid := &BodySFCSelectionConvergence{
		Position: Position{},
		LocalId: 42,
	}
	if err := valid.Validate(); err != nil {
		t.Errorf("Valid object should not have validation errors: %v", err)
	}

	// Test missing required field: Position
	invalidPosition := &BodySFCSelectionConvergence{}
	if err := invalidPosition.Validate(); err == nil {
		t.Error("Missing required field Position should cause validation error")
	}

	// Test missing required field: LocalId
	invalidLocalId := &BodySFCSelectionConvergence{}
	if err := invalidLocalId.Validate(); err == nil {
		t.Error("Missing required field LocalId should cause validation error")
	}

}

func TestBodySFCSimultaneousDivergenceConnectionPointOutXMLMarshaling(t *testing.T) {
	original := &BodySFCSimultaneousDivergenceConnectionPointOut{
	}

	// Test marshaling
	xmlData, err := xml.Marshal(original)
	if err != nil {
		t.Fatalf("Failed to marshal XML: %v", err)
	}

	// Test unmarshaling
	var unmarshaled BodySFCSimultaneousDivergenceConnectionPointOut
	err = xml.Unmarshal(xmlData, &unmarshaled)
	if err != nil {
		t.Fatalf("Failed to unmarshal XML: %v", err)
	}

	// Test validation
	if err := original.Validate(); err != nil {
		t.Errorf("Validation failed: %v", err)
	}
}

func TestBodySFCSimultaneousDivergenceConnectionPointOutValidation(t *testing.T) {
	// Test valid case
	valid := &BodySFCSimultaneousDivergenceConnectionPointOut{
	}
	if err := valid.Validate(); err != nil {
		t.Errorf("Valid object should not have validation errors: %v", err)
	}

}

func TestBodySFCSimultaneousDivergenceXMLMarshaling(t *testing.T) {
	original := &BodySFCSimultaneousDivergence{
		Position: Position{},
		ConnectionPointIn: &ConnectionPointIn{},
		ConnectionPointOut: []BodySFCSimultaneousDivergenceConnectionPointOut{BodySFCSimultaneousDivergenceConnectionPointOut{}},
		Documentation: &FormattedText{},
		LocalId: 42,
		Height: floatPtr(3.14),
		Width: floatPtr(3.14),
		Name: stringPtr("test_value"),
	}

	// Test marshaling
	xmlData, err := xml.Marshal(original)
	if err != nil {
		t.Fatalf("Failed to marshal XML: %v", err)
	}

	// Test unmarshaling
	var unmarshaled BodySFCSimultaneousDivergence
	err = xml.Unmarshal(xmlData, &unmarshaled)
	if err != nil {
		t.Fatalf("Failed to unmarshal XML: %v", err)
	}

	// Test validation
	if err := original.Validate(); err != nil {
		t.Errorf("Validation failed: %v", err)
	}
}

func TestBodySFCSimultaneousDivergenceValidation(t *testing.T) {
	// Test valid case
	valid := &BodySFCSimultaneousDivergence{
		Position: Position{},
		LocalId: 42,
	}
	if err := valid.Validate(); err != nil {
		t.Errorf("Valid object should not have validation errors: %v", err)
	}

	// Test missing required field: Position
	invalidPosition := &BodySFCSimultaneousDivergence{}
	if err := invalidPosition.Validate(); err == nil {
		t.Error("Missing required field Position should cause validation error")
	}

	// Test missing required field: LocalId
	invalidLocalId := &BodySFCSimultaneousDivergence{}
	if err := invalidLocalId.Validate(); err == nil {
		t.Error("Missing required field LocalId should cause validation error")
	}

}

func TestBodySFCSimultaneousConvergenceXMLMarshaling(t *testing.T) {
	original := &BodySFCSimultaneousConvergence{
		Position: Position{},
		ConnectionPointIn: []ConnectionPointIn{ConnectionPointIn{}},
		ConnectionPointOut: &ConnectionPointOut{},
		Documentation: &FormattedText{},
		LocalId: 42,
		Height: floatPtr(3.14),
		Width: floatPtr(3.14),
	}

	// Test marshaling
	xmlData, err := xml.Marshal(original)
	if err != nil {
		t.Fatalf("Failed to marshal XML: %v", err)
	}

	// Test unmarshaling
	var unmarshaled BodySFCSimultaneousConvergence
	err = xml.Unmarshal(xmlData, &unmarshaled)
	if err != nil {
		t.Fatalf("Failed to unmarshal XML: %v", err)
	}

	// Test validation
	if err := original.Validate(); err != nil {
		t.Errorf("Validation failed: %v", err)
	}
}

func TestBodySFCSimultaneousConvergenceValidation(t *testing.T) {
	// Test valid case
	valid := &BodySFCSimultaneousConvergence{
		Position: Position{},
		LocalId: 42,
	}
	if err := valid.Validate(); err != nil {
		t.Errorf("Valid object should not have validation errors: %v", err)
	}

	// Test missing required field: Position
	invalidPosition := &BodySFCSimultaneousConvergence{}
	if err := invalidPosition.Validate(); err == nil {
		t.Error("Missing required field Position should cause validation error")
	}

	// Test missing required field: LocalId
	invalidLocalId := &BodySFCSimultaneousConvergence{}
	if err := invalidLocalId.Validate(); err == nil {
		t.Error("Missing required field LocalId should cause validation error")
	}

}

func TestBodySFCXMLMarshaling(t *testing.T) {
	original := &BodySFC{
		Comment: &BodySFCComment{},
		Error: &BodySFCError{},
		Connector: &BodySFCConnector{},
		Continuation: &BodySFCContinuation{},
		ActionBlock: &BodySFCActionBlock{},
		Block: &BodySFCBlock{},
		InVariable: &BodySFCInVariable{},
		OutVariable: &BodySFCOutVariable{},
		InOutVariable: &BodySFCInOutVariable{},
		Label: &BodySFCLabel{},
		Jump: &BodySFCJump{},
		Return: &BodySFCReturn{},
		LeftPowerRail: &BodySFCLeftPowerRail{},
		RightPowerRail: &BodySFCRightPowerRail{},
		Coil: &BodySFCCoil{},
		Contact: &BodySFCContact{},
		Step: &BodySFCStep{},
		MacroStep: &BodySFCMacroStep{},
		JumpStep: &BodySFCJumpStep{},
		Transition: &BodySFCTransition{},
		SelectionDivergence: &BodySFCSelectionDivergence{},
		SelectionConvergence: &BodySFCSelectionConvergence{},
		SimultaneousDivergence: &BodySFCSimultaneousDivergence{},
		SimultaneousConvergence: &BodySFCSimultaneousConvergence{},
	}

	// Test marshaling
	xmlData, err := xml.Marshal(original)
	if err != nil {
		t.Fatalf("Failed to marshal XML: %v", err)
	}

	// Test unmarshaling
	var unmarshaled BodySFC
	err = xml.Unmarshal(xmlData, &unmarshaled)
	if err != nil {
		t.Fatalf("Failed to unmarshal XML: %v", err)
	}

	// Test validation
	if err := original.Validate(); err != nil {
		t.Errorf("Validation failed: %v", err)
	}
}

func TestBodySFCValidation(t *testing.T) {
	// Test valid case
	valid := &BodySFC{
	}
	if err := valid.Validate(); err != nil {
		t.Errorf("Valid object should not have validation errors: %v", err)
	}

}

func TestBodyXMLMarshaling(t *testing.T) {
	original := &Body{
		Documentation: &FormattedText{},
		IL: &FormattedText{},
		ST: &FormattedText{},
		FBD: &BodyFBD{},
		LD: &BodyLD{},
		SFC: &BodySFC{},
	}

	// Test marshaling
	xmlData, err := xml.Marshal(original)
	if err != nil {
		t.Fatalf("Failed to marshal XML: %v", err)
	}

	// Test unmarshaling
	var unmarshaled Body
	err = xml.Unmarshal(xmlData, &unmarshaled)
	if err != nil {
		t.Fatalf("Failed to unmarshal XML: %v", err)
	}

	// Test validation
	if err := original.Validate(); err != nil {
		t.Errorf("Validation failed: %v", err)
	}
}

func TestBodyValidation(t *testing.T) {
	// Test valid case
	valid := &Body{
	}
	if err := valid.Validate(); err != nil {
		t.Errorf("Valid object should not have validation errors: %v", err)
	}

}

func TestVarListXMLMarshaling(t *testing.T) {
	original := &VarList{
		Name: stringPtr("test_value"),
		Constant: boolPtr(true),
		Retain: boolPtr(true),
		Nonretain: boolPtr(true),
		Persistent: boolPtr(true),
		Nonpersistent: boolPtr(true),
	}

	// Test marshaling
	xmlData, err := xml.Marshal(original)
	if err != nil {
		t.Fatalf("Failed to marshal XML: %v", err)
	}

	// Test unmarshaling
	var unmarshaled VarList
	err = xml.Unmarshal(xmlData, &unmarshaled)
	if err != nil {
		t.Fatalf("Failed to unmarshal XML: %v", err)
	}

	// Test validation
	if err := original.Validate(); err != nil {
		t.Errorf("Validation failed: %v", err)
	}
}

func TestVarListValidation(t *testing.T) {
	// Test valid case
	valid := &VarList{
	}
	if err := valid.Validate(); err != nil {
		t.Errorf("Valid object should not have validation errors: %v", err)
	}

}

func TestVarListPlainVariableXMLMarshaling(t *testing.T) {
	original := &VarListPlainVariable{
		Type: DataType{},
		InitialValue: &Value{},
		Documentation: &FormattedText{},
		Name: "test_value",
		Address: stringPtr("test_value"),
	}

	// Test marshaling
	xmlData, err := xml.Marshal(original)
	if err != nil {
		t.Fatalf("Failed to marshal XML: %v", err)
	}

	// Test unmarshaling
	var unmarshaled VarListPlainVariable
	err = xml.Unmarshal(xmlData, &unmarshaled)
	if err != nil {
		t.Fatalf("Failed to unmarshal XML: %v", err)
	}

	// Test validation
	if err := original.Validate(); err != nil {
		t.Errorf("Validation failed: %v", err)
	}
}

func TestVarListPlainVariableValidation(t *testing.T) {
	// Test valid case
	valid := &VarListPlainVariable{
		Type: DataType{},
		Name: "test_value",
	}
	if err := valid.Validate(); err != nil {
		t.Errorf("Valid object should not have validation errors: %v", err)
	}

	// Test missing required field: Type
	invalidType := &VarListPlainVariable{}
	if err := invalidType.Validate(); err == nil {
		t.Error("Missing required field Type should cause validation error")
	}

	// Test missing required field: Name
	invalidName := &VarListPlainVariable{}
	if err := invalidName.Validate(); err == nil {
		t.Error("Missing required field Name should cause validation error")
	}

}

func TestVarListPlainXMLMarshaling(t *testing.T) {
	original := &VarListPlain{
		Variable: []VarListPlainVariable{VarListPlainVariable{}},
		Documentation: &FormattedText{},
	}

	// Test marshaling
	xmlData, err := xml.Marshal(original)
	if err != nil {
		t.Fatalf("Failed to marshal XML: %v", err)
	}

	// Test unmarshaling
	var unmarshaled VarListPlain
	err = xml.Unmarshal(xmlData, &unmarshaled)
	if err != nil {
		t.Fatalf("Failed to unmarshal XML: %v", err)
	}

	// Test validation
	if err := original.Validate(); err != nil {
		t.Errorf("Validation failed: %v", err)
	}
}

func TestVarListPlainValidation(t *testing.T) {
	// Test valid case
	valid := &VarListPlain{
	}
	if err := valid.Validate(); err != nil {
		t.Errorf("Valid object should not have validation errors: %v", err)
	}

}

func TestPositionXMLMarshaling(t *testing.T) {
	original := &Position{
		X: 3.14,
		Y: 3.14,
	}

	// Test marshaling
	xmlData, err := xml.Marshal(original)
	if err != nil {
		t.Fatalf("Failed to marshal XML: %v", err)
	}

	// Test unmarshaling
	var unmarshaled Position
	err = xml.Unmarshal(xmlData, &unmarshaled)
	if err != nil {
		t.Fatalf("Failed to unmarshal XML: %v", err)
	}

	// Test validation
	if err := original.Validate(); err != nil {
		t.Errorf("Validation failed: %v", err)
	}
}

func TestPositionValidation(t *testing.T) {
	// Test valid case
	valid := &Position{
		X: 3.14,
		Y: 3.14,
	}
	if err := valid.Validate(); err != nil {
		t.Errorf("Valid object should not have validation errors: %v", err)
	}

	// Test missing required field: X
	invalidX := &Position{}
	if err := invalidX.Validate(); err == nil {
		t.Error("Missing required field X should cause validation error")
	}

	// Test missing required field: Y
	invalidY := &Position{}
	if err := invalidY.Validate(); err == nil {
		t.Error("Missing required field Y should cause validation error")
	}

}

func TestConnectionXMLMarshaling(t *testing.T) {
	original := &Connection{
		Position: []Position{Position{}},
		RefLocalId: 42,
		FormalParameter: stringPtr("test_value"),
	}

	// Test marshaling
	xmlData, err := xml.Marshal(original)
	if err != nil {
		t.Fatalf("Failed to marshal XML: %v", err)
	}

	// Test unmarshaling
	var unmarshaled Connection
	err = xml.Unmarshal(xmlData, &unmarshaled)
	if err != nil {
		t.Fatalf("Failed to unmarshal XML: %v", err)
	}

	// Test validation
	if err := original.Validate(); err != nil {
		t.Errorf("Validation failed: %v", err)
	}
}

func TestConnectionValidation(t *testing.T) {
	// Test valid case
	valid := &Connection{
		RefLocalId: 42,
	}
	if err := valid.Validate(); err != nil {
		t.Errorf("Valid object should not have validation errors: %v", err)
	}

	// Test missing required field: RefLocalId
	invalidRefLocalId := &Connection{}
	if err := invalidRefLocalId.Validate(); err == nil {
		t.Error("Missing required field RefLocalId should cause validation error")
	}

}

func TestConnectionPointInXMLMarshaling(t *testing.T) {
	original := &ConnectionPointIn{
		RelPosition: &Position{},
		Connection: []Connection{Connection{}},
		Expression: "test_value",
	}

	// Test marshaling
	xmlData, err := xml.Marshal(original)
	if err != nil {
		t.Fatalf("Failed to marshal XML: %v", err)
	}

	// Test unmarshaling
	var unmarshaled ConnectionPointIn
	err = xml.Unmarshal(xmlData, &unmarshaled)
	if err != nil {
		t.Fatalf("Failed to unmarshal XML: %v", err)
	}

	// Test validation
	if err := original.Validate(); err != nil {
		t.Errorf("Validation failed: %v", err)
	}
}

func TestConnectionPointInValidation(t *testing.T) {
	// Test valid case
	valid := &ConnectionPointIn{
	}
	if err := valid.Validate(); err != nil {
		t.Errorf("Valid object should not have validation errors: %v", err)
	}

}

func TestConnectionPointOutXMLMarshaling(t *testing.T) {
	original := &ConnectionPointOut{
		RelPosition: &Position{},
		Expression: stringPtr("test_value"),
	}

	// Test marshaling
	xmlData, err := xml.Marshal(original)
	if err != nil {
		t.Fatalf("Failed to marshal XML: %v", err)
	}

	// Test unmarshaling
	var unmarshaled ConnectionPointOut
	err = xml.Unmarshal(xmlData, &unmarshaled)
	if err != nil {
		t.Fatalf("Failed to unmarshal XML: %v", err)
	}

	// Test validation
	if err := original.Validate(); err != nil {
		t.Errorf("Validation failed: %v", err)
	}
}

func TestConnectionPointOutValidation(t *testing.T) {
	// Test valid case
	valid := &ConnectionPointOut{
	}
	if err := valid.Validate(); err != nil {
		t.Errorf("Valid object should not have validation errors: %v", err)
	}

}

func TestPouInstanceXMLMarshaling(t *testing.T) {
	original := &PouInstance{
		Documentation: &FormattedText{},
		Name: "test_value",
		Type: "test_value",
	}

	// Test marshaling
	xmlData, err := xml.Marshal(original)
	if err != nil {
		t.Fatalf("Failed to marshal XML: %v", err)
	}

	// Test unmarshaling
	var unmarshaled PouInstance
	err = xml.Unmarshal(xmlData, &unmarshaled)
	if err != nil {
		t.Fatalf("Failed to unmarshal XML: %v", err)
	}

	// Test validation
	if err := original.Validate(); err != nil {
		t.Errorf("Validation failed: %v", err)
	}
}

func TestPouInstanceValidation(t *testing.T) {
	// Test valid case
	valid := &PouInstance{
		Name: "test_value",
		Type: "test_value",
	}
	if err := valid.Validate(); err != nil {
		t.Errorf("Valid object should not have validation errors: %v", err)
	}

	// Test missing required field: Name
	invalidName := &PouInstance{}
	if err := invalidName.Validate(); err == nil {
		t.Error("Missing required field Name should cause validation error")
	}

	// Test missing required field: Type
	invalidType := &PouInstance{}
	if err := invalidType.Validate(); err == nil {
		t.Error("Missing required field Type should cause validation error")
	}

}

func TestFormattedTextXMLMarshaling(t *testing.T) {
	original := &FormattedText{
	}

	// Test marshaling
	xmlData, err := xml.Marshal(original)
	if err != nil {
		t.Fatalf("Failed to marshal XML: %v", err)
	}

	// Test unmarshaling
	var unmarshaled FormattedText
	err = xml.Unmarshal(xmlData, &unmarshaled)
	if err != nil {
		t.Fatalf("Failed to unmarshal XML: %v", err)
	}

	// Test validation
	if err := original.Validate(); err != nil {
		t.Errorf("Validation failed: %v", err)
	}
}

func TestFormattedTextValidation(t *testing.T) {
	// Test valid case
	valid := &FormattedText{
	}
	if err := valid.Validate(); err != nil {
		t.Errorf("Valid object should not have validation errors: %v", err)
	}

}

func TestEdgeModifierTypeXMLMarshaling(t *testing.T) {
	original := EdgeModifierTypeNone

	// Test marshaling
	xmlData, err := xml.Marshal(original)
	if err != nil {
		t.Fatalf("Failed to marshal XML: %v", err)
	}

	// Test unmarshaling
	var unmarshaled EdgeModifierType
	err = xml.Unmarshal(xmlData, &unmarshaled)
	if err != nil {
		t.Fatalf("Failed to unmarshal XML: %v", err)
	}

	// Test validation
	if err := original.Validate(); err != nil {
		t.Errorf("Validation failed: %v", err)
	}
}

func TestEdgeModifierTypeValidation(t *testing.T) {
	// Test valid case
	valid := EdgeModifierTypeNone
	if err := valid.Validate(); err != nil {
		t.Errorf("Valid object should not have validation errors: %v", err)
	}

	// Test invalid enum value
	invalid := EdgeModifierType("invalid_value")
	if err := invalid.Validate(); err == nil {
		t.Error("Invalid enum value should cause validation error")
	}

}

func TestStorageModifierTypeXMLMarshaling(t *testing.T) {
	original := StorageModifierTypeNone

	// Test marshaling
	xmlData, err := xml.Marshal(original)
	if err != nil {
		t.Fatalf("Failed to marshal XML: %v", err)
	}

	// Test unmarshaling
	var unmarshaled StorageModifierType
	err = xml.Unmarshal(xmlData, &unmarshaled)
	if err != nil {
		t.Fatalf("Failed to unmarshal XML: %v", err)
	}

	// Test validation
	if err := original.Validate(); err != nil {
		t.Errorf("Validation failed: %v", err)
	}
}

func TestStorageModifierTypeValidation(t *testing.T) {
	// Test valid case
	valid := StorageModifierTypeNone
	if err := valid.Validate(); err != nil {
		t.Errorf("Valid object should not have validation errors: %v", err)
	}

	// Test invalid enum value
	invalid := StorageModifierType("invalid_value")
	if err := invalid.Validate(); err == nil {
		t.Error("Invalid enum value should cause validation error")
	}

}

func TestPouTypeXMLMarshaling(t *testing.T) {
	original := PouTypeFunction

	// Test marshaling
	xmlData, err := xml.Marshal(original)
	if err != nil {
		t.Fatalf("Failed to marshal XML: %v", err)
	}

	// Test unmarshaling
	var unmarshaled PouType
	err = xml.Unmarshal(xmlData, &unmarshaled)
	if err != nil {
		t.Fatalf("Failed to unmarshal XML: %v", err)
	}

	// Test validation
	if err := original.Validate(); err != nil {
		t.Errorf("Validation failed: %v", err)
	}
}

func TestPouTypeValidation(t *testing.T) {
	// Test valid case
	valid := PouTypeFunction
	if err := valid.Validate(); err != nil {
		t.Errorf("Valid object should not have validation errors: %v", err)
	}

	// Test invalid enum value
	invalid := PouType("invalid_value")
	if err := invalid.Validate(); err == nil {
		t.Error("Invalid enum value should cause validation error")
	}

}

func TestProjectFileHeaderXMLMarshaling(t *testing.T) {
	original := &ProjectFileHeader{
		CompanyName: "test_value",
		CompanyURL: stringPtr("test_value"),
		ProductName: "test_value",
		ProductVersion: "test_value",
		ProductRelease: stringPtr("test_value"),
		CreationDateTime: time.Now(),
		ContentDescription: stringPtr("test_value"),
	}

	// Test marshaling
	xmlData, err := xml.Marshal(original)
	if err != nil {
		t.Fatalf("Failed to marshal XML: %v", err)
	}

	// Test unmarshaling
	var unmarshaled ProjectFileHeader
	err = xml.Unmarshal(xmlData, &unmarshaled)
	if err != nil {
		t.Fatalf("Failed to unmarshal XML: %v", err)
	}

	// Test validation
	if err := original.Validate(); err != nil {
		t.Errorf("Validation failed: %v", err)
	}
}

func TestProjectFileHeaderValidation(t *testing.T) {
	// Test valid case
	valid := &ProjectFileHeader{
		CompanyName: "test_value",
		ProductName: "test_value",
		ProductVersion: "test_value",
		CreationDateTime: time.Now(),
	}
	if err := valid.Validate(); err != nil {
		t.Errorf("Valid object should not have validation errors: %v", err)
	}

	// Test missing required field: CompanyName
	invalidCompanyName := &ProjectFileHeader{}
	if err := invalidCompanyName.Validate(); err == nil {
		t.Error("Missing required field CompanyName should cause validation error")
	}

	// Test missing required field: ProductName
	invalidProductName := &ProjectFileHeader{}
	if err := invalidProductName.Validate(); err == nil {
		t.Error("Missing required field ProductName should cause validation error")
	}

	// Test missing required field: ProductVersion
	invalidProductVersion := &ProjectFileHeader{}
	if err := invalidProductVersion.Validate(); err == nil {
		t.Error("Missing required field ProductVersion should cause validation error")
	}

	// Test missing required field: CreationDateTime
	invalidCreationDateTime := &ProjectFileHeader{}
	if err := invalidCreationDateTime.Validate(); err == nil {
		t.Error("Missing required field CreationDateTime should cause validation error")
	}

}

func TestProjectContentHeaderCoordinateInfoPageSizeXMLMarshaling(t *testing.T) {
	original := &ProjectContentHeaderCoordinateInfoPageSize{
		X: 3.14,
		Y: 3.14,
	}

	// Test marshaling
	xmlData, err := xml.Marshal(original)
	if err != nil {
		t.Fatalf("Failed to marshal XML: %v", err)
	}

	// Test unmarshaling
	var unmarshaled ProjectContentHeaderCoordinateInfoPageSize
	err = xml.Unmarshal(xmlData, &unmarshaled)
	if err != nil {
		t.Fatalf("Failed to unmarshal XML: %v", err)
	}

	// Test validation
	if err := original.Validate(); err != nil {
		t.Errorf("Validation failed: %v", err)
	}
}

func TestProjectContentHeaderCoordinateInfoPageSizeValidation(t *testing.T) {
	// Test valid case
	valid := &ProjectContentHeaderCoordinateInfoPageSize{
		X: 3.14,
		Y: 3.14,
	}
	if err := valid.Validate(); err != nil {
		t.Errorf("Valid object should not have validation errors: %v", err)
	}

	// Test missing required field: X
	invalidX := &ProjectContentHeaderCoordinateInfoPageSize{}
	if err := invalidX.Validate(); err == nil {
		t.Error("Missing required field X should cause validation error")
	}

	// Test missing required field: Y
	invalidY := &ProjectContentHeaderCoordinateInfoPageSize{}
	if err := invalidY.Validate(); err == nil {
		t.Error("Missing required field Y should cause validation error")
	}

}

func TestProjectContentHeaderCoordinateInfoFbdScalingXMLMarshaling(t *testing.T) {
	original := &ProjectContentHeaderCoordinateInfoFbdScaling{
		X: 3.14,
		Y: 3.14,
	}

	// Test marshaling
	xmlData, err := xml.Marshal(original)
	if err != nil {
		t.Fatalf("Failed to marshal XML: %v", err)
	}

	// Test unmarshaling
	var unmarshaled ProjectContentHeaderCoordinateInfoFbdScaling
	err = xml.Unmarshal(xmlData, &unmarshaled)
	if err != nil {
		t.Fatalf("Failed to unmarshal XML: %v", err)
	}

	// Test validation
	if err := original.Validate(); err != nil {
		t.Errorf("Validation failed: %v", err)
	}
}

func TestProjectContentHeaderCoordinateInfoFbdScalingValidation(t *testing.T) {
	// Test valid case
	valid := &ProjectContentHeaderCoordinateInfoFbdScaling{
		X: 3.14,
		Y: 3.14,
	}
	if err := valid.Validate(); err != nil {
		t.Errorf("Valid object should not have validation errors: %v", err)
	}

	// Test missing required field: X
	invalidX := &ProjectContentHeaderCoordinateInfoFbdScaling{}
	if err := invalidX.Validate(); err == nil {
		t.Error("Missing required field X should cause validation error")
	}

	// Test missing required field: Y
	invalidY := &ProjectContentHeaderCoordinateInfoFbdScaling{}
	if err := invalidY.Validate(); err == nil {
		t.Error("Missing required field Y should cause validation error")
	}

}

func TestProjectContentHeaderCoordinateInfoFbdXMLMarshaling(t *testing.T) {
	original := &ProjectContentHeaderCoordinateInfoFbd{
		Scaling: ProjectContentHeaderCoordinateInfoFbdScaling{},
	}

	// Test marshaling
	xmlData, err := xml.Marshal(original)
	if err != nil {
		t.Fatalf("Failed to marshal XML: %v", err)
	}

	// Test unmarshaling
	var unmarshaled ProjectContentHeaderCoordinateInfoFbd
	err = xml.Unmarshal(xmlData, &unmarshaled)
	if err != nil {
		t.Fatalf("Failed to unmarshal XML: %v", err)
	}

	// Test validation
	if err := original.Validate(); err != nil {
		t.Errorf("Validation failed: %v", err)
	}
}

func TestProjectContentHeaderCoordinateInfoFbdValidation(t *testing.T) {
	// Test valid case
	valid := &ProjectContentHeaderCoordinateInfoFbd{
		Scaling: ProjectContentHeaderCoordinateInfoFbdScaling{},
	}
	if err := valid.Validate(); err != nil {
		t.Errorf("Valid object should not have validation errors: %v", err)
	}

	// Test missing required field: Scaling
	invalidScaling := &ProjectContentHeaderCoordinateInfoFbd{}
	if err := invalidScaling.Validate(); err == nil {
		t.Error("Missing required field Scaling should cause validation error")
	}

}

func TestProjectContentHeaderCoordinateInfoLdScalingXMLMarshaling(t *testing.T) {
	original := &ProjectContentHeaderCoordinateInfoLdScaling{
		X: 3.14,
		Y: 3.14,
	}

	// Test marshaling
	xmlData, err := xml.Marshal(original)
	if err != nil {
		t.Fatalf("Failed to marshal XML: %v", err)
	}

	// Test unmarshaling
	var unmarshaled ProjectContentHeaderCoordinateInfoLdScaling
	err = xml.Unmarshal(xmlData, &unmarshaled)
	if err != nil {
		t.Fatalf("Failed to unmarshal XML: %v", err)
	}

	// Test validation
	if err := original.Validate(); err != nil {
		t.Errorf("Validation failed: %v", err)
	}
}

func TestProjectContentHeaderCoordinateInfoLdScalingValidation(t *testing.T) {
	// Test valid case
	valid := &ProjectContentHeaderCoordinateInfoLdScaling{
		X: 3.14,
		Y: 3.14,
	}
	if err := valid.Validate(); err != nil {
		t.Errorf("Valid object should not have validation errors: %v", err)
	}

	// Test missing required field: X
	invalidX := &ProjectContentHeaderCoordinateInfoLdScaling{}
	if err := invalidX.Validate(); err == nil {
		t.Error("Missing required field X should cause validation error")
	}

	// Test missing required field: Y
	invalidY := &ProjectContentHeaderCoordinateInfoLdScaling{}
	if err := invalidY.Validate(); err == nil {
		t.Error("Missing required field Y should cause validation error")
	}

}

func TestProjectContentHeaderCoordinateInfoLdXMLMarshaling(t *testing.T) {
	original := &ProjectContentHeaderCoordinateInfoLd{
		Scaling: ProjectContentHeaderCoordinateInfoLdScaling{},
	}

	// Test marshaling
	xmlData, err := xml.Marshal(original)
	if err != nil {
		t.Fatalf("Failed to marshal XML: %v", err)
	}

	// Test unmarshaling
	var unmarshaled ProjectContentHeaderCoordinateInfoLd
	err = xml.Unmarshal(xmlData, &unmarshaled)
	if err != nil {
		t.Fatalf("Failed to unmarshal XML: %v", err)
	}

	// Test validation
	if err := original.Validate(); err != nil {
		t.Errorf("Validation failed: %v", err)
	}
}

func TestProjectContentHeaderCoordinateInfoLdValidation(t *testing.T) {
	// Test valid case
	valid := &ProjectContentHeaderCoordinateInfoLd{
		Scaling: ProjectContentHeaderCoordinateInfoLdScaling{},
	}
	if err := valid.Validate(); err != nil {
		t.Errorf("Valid object should not have validation errors: %v", err)
	}

	// Test missing required field: Scaling
	invalidScaling := &ProjectContentHeaderCoordinateInfoLd{}
	if err := invalidScaling.Validate(); err == nil {
		t.Error("Missing required field Scaling should cause validation error")
	}

}

func TestProjectContentHeaderCoordinateInfoSfcScalingXMLMarshaling(t *testing.T) {
	original := &ProjectContentHeaderCoordinateInfoSfcScaling{
		X: 3.14,
		Y: 3.14,
	}

	// Test marshaling
	xmlData, err := xml.Marshal(original)
	if err != nil {
		t.Fatalf("Failed to marshal XML: %v", err)
	}

	// Test unmarshaling
	var unmarshaled ProjectContentHeaderCoordinateInfoSfcScaling
	err = xml.Unmarshal(xmlData, &unmarshaled)
	if err != nil {
		t.Fatalf("Failed to unmarshal XML: %v", err)
	}

	// Test validation
	if err := original.Validate(); err != nil {
		t.Errorf("Validation failed: %v", err)
	}
}

func TestProjectContentHeaderCoordinateInfoSfcScalingValidation(t *testing.T) {
	// Test valid case
	valid := &ProjectContentHeaderCoordinateInfoSfcScaling{
		X: 3.14,
		Y: 3.14,
	}
	if err := valid.Validate(); err != nil {
		t.Errorf("Valid object should not have validation errors: %v", err)
	}

	// Test missing required field: X
	invalidX := &ProjectContentHeaderCoordinateInfoSfcScaling{}
	if err := invalidX.Validate(); err == nil {
		t.Error("Missing required field X should cause validation error")
	}

	// Test missing required field: Y
	invalidY := &ProjectContentHeaderCoordinateInfoSfcScaling{}
	if err := invalidY.Validate(); err == nil {
		t.Error("Missing required field Y should cause validation error")
	}

}

func TestProjectContentHeaderCoordinateInfoSfcXMLMarshaling(t *testing.T) {
	original := &ProjectContentHeaderCoordinateInfoSfc{
		Scaling: ProjectContentHeaderCoordinateInfoSfcScaling{},
	}

	// Test marshaling
	xmlData, err := xml.Marshal(original)
	if err != nil {
		t.Fatalf("Failed to marshal XML: %v", err)
	}

	// Test unmarshaling
	var unmarshaled ProjectContentHeaderCoordinateInfoSfc
	err = xml.Unmarshal(xmlData, &unmarshaled)
	if err != nil {
		t.Fatalf("Failed to unmarshal XML: %v", err)
	}

	// Test validation
	if err := original.Validate(); err != nil {
		t.Errorf("Validation failed: %v", err)
	}
}

func TestProjectContentHeaderCoordinateInfoSfcValidation(t *testing.T) {
	// Test valid case
	valid := &ProjectContentHeaderCoordinateInfoSfc{
		Scaling: ProjectContentHeaderCoordinateInfoSfcScaling{},
	}
	if err := valid.Validate(); err != nil {
		t.Errorf("Valid object should not have validation errors: %v", err)
	}

	// Test missing required field: Scaling
	invalidScaling := &ProjectContentHeaderCoordinateInfoSfc{}
	if err := invalidScaling.Validate(); err == nil {
		t.Error("Missing required field Scaling should cause validation error")
	}

}

func TestProjectContentHeaderCoordinateInfoXMLMarshaling(t *testing.T) {
	original := &ProjectContentHeaderCoordinateInfo{
		PageSize: &ProjectContentHeaderCoordinateInfoPageSize{},
		Fbd: ProjectContentHeaderCoordinateInfoFbd{},
		Ld: ProjectContentHeaderCoordinateInfoLd{},
		Sfc: ProjectContentHeaderCoordinateInfoSfc{},
	}

	// Test marshaling
	xmlData, err := xml.Marshal(original)
	if err != nil {
		t.Fatalf("Failed to marshal XML: %v", err)
	}

	// Test unmarshaling
	var unmarshaled ProjectContentHeaderCoordinateInfo
	err = xml.Unmarshal(xmlData, &unmarshaled)
	if err != nil {
		t.Fatalf("Failed to unmarshal XML: %v", err)
	}

	// Test validation
	if err := original.Validate(); err != nil {
		t.Errorf("Validation failed: %v", err)
	}
}

func TestProjectContentHeaderCoordinateInfoValidation(t *testing.T) {
	// Test valid case
	valid := &ProjectContentHeaderCoordinateInfo{
		Fbd: ProjectContentHeaderCoordinateInfoFbd{},
		Ld: ProjectContentHeaderCoordinateInfoLd{},
		Sfc: ProjectContentHeaderCoordinateInfoSfc{},
	}
	if err := valid.Validate(); err != nil {
		t.Errorf("Valid object should not have validation errors: %v", err)
	}

	// Test missing required field: Fbd
	invalidFbd := &ProjectContentHeaderCoordinateInfo{}
	if err := invalidFbd.Validate(); err == nil {
		t.Error("Missing required field Fbd should cause validation error")
	}

	// Test missing required field: Ld
	invalidLd := &ProjectContentHeaderCoordinateInfo{}
	if err := invalidLd.Validate(); err == nil {
		t.Error("Missing required field Ld should cause validation error")
	}

	// Test missing required field: Sfc
	invalidSfc := &ProjectContentHeaderCoordinateInfo{}
	if err := invalidSfc.Validate(); err == nil {
		t.Error("Missing required field Sfc should cause validation error")
	}

}

func TestProjectContentHeaderXMLMarshaling(t *testing.T) {
	original := &ProjectContentHeader{
		Comment: stringPtr("test_value"),
		CoordinateInfo: ProjectContentHeaderCoordinateInfo{},
		Name: "test_value",
		Version: stringPtr("test_value"),
		ModificationDateTime: timePtr(time.Now()),
		Organization: stringPtr("test_value"),
		Author: stringPtr("test_value"),
		Language: stringPtr("test_value"),
	}

	// Test marshaling
	xmlData, err := xml.Marshal(original)
	if err != nil {
		t.Fatalf("Failed to marshal XML: %v", err)
	}

	// Test unmarshaling
	var unmarshaled ProjectContentHeader
	err = xml.Unmarshal(xmlData, &unmarshaled)
	if err != nil {
		t.Fatalf("Failed to unmarshal XML: %v", err)
	}

	// Test validation
	if err := original.Validate(); err != nil {
		t.Errorf("Validation failed: %v", err)
	}
}

func TestProjectContentHeaderValidation(t *testing.T) {
	// Test valid case
	valid := &ProjectContentHeader{
		CoordinateInfo: ProjectContentHeaderCoordinateInfo{},
		Name: "test_value",
	}
	if err := valid.Validate(); err != nil {
		t.Errorf("Valid object should not have validation errors: %v", err)
	}

	// Test missing required field: CoordinateInfo
	invalidCoordinateInfo := &ProjectContentHeader{}
	if err := invalidCoordinateInfo.Validate(); err == nil {
		t.Error("Missing required field CoordinateInfo should cause validation error")
	}

	// Test missing required field: Name
	invalidName := &ProjectContentHeader{}
	if err := invalidName.Validate(); err == nil {
		t.Error("Missing required field Name should cause validation error")
	}

}

func TestProjectTypesDataTypesDataTypeXMLMarshaling(t *testing.T) {
	original := &ProjectTypesDataTypesDataType{
		BaseType: DataType{},
		InitialValue: &Value{},
		Documentation: &FormattedText{},
		Name: "test_value",
	}

	// Test marshaling
	xmlData, err := xml.Marshal(original)
	if err != nil {
		t.Fatalf("Failed to marshal XML: %v", err)
	}

	// Test unmarshaling
	var unmarshaled ProjectTypesDataTypesDataType
	err = xml.Unmarshal(xmlData, &unmarshaled)
	if err != nil {
		t.Fatalf("Failed to unmarshal XML: %v", err)
	}

	// Test validation
	if err := original.Validate(); err != nil {
		t.Errorf("Validation failed: %v", err)
	}
}

func TestProjectTypesDataTypesDataTypeValidation(t *testing.T) {
	// Test valid case
	valid := &ProjectTypesDataTypesDataType{
		BaseType: DataType{},
		Name: "test_value",
	}
	if err := valid.Validate(); err != nil {
		t.Errorf("Valid object should not have validation errors: %v", err)
	}

	// Test missing required field: BaseType
	invalidBaseType := &ProjectTypesDataTypesDataType{}
	if err := invalidBaseType.Validate(); err == nil {
		t.Error("Missing required field BaseType should cause validation error")
	}

	// Test missing required field: Name
	invalidName := &ProjectTypesDataTypesDataType{}
	if err := invalidName.Validate(); err == nil {
		t.Error("Missing required field Name should cause validation error")
	}

}

func TestProjectTypesDataTypesXMLMarshaling(t *testing.T) {
	original := &ProjectTypesDataTypes{
		DataType: []ProjectTypesDataTypesDataType{ProjectTypesDataTypesDataType{}},
	}

	// Test marshaling
	xmlData, err := xml.Marshal(original)
	if err != nil {
		t.Fatalf("Failed to marshal XML: %v", err)
	}

	// Test unmarshaling
	var unmarshaled ProjectTypesDataTypes
	err = xml.Unmarshal(xmlData, &unmarshaled)
	if err != nil {
		t.Fatalf("Failed to unmarshal XML: %v", err)
	}

	// Test validation
	if err := original.Validate(); err != nil {
		t.Errorf("Validation failed: %v", err)
	}
}

func TestProjectTypesDataTypesValidation(t *testing.T) {
	// Test valid case
	valid := &ProjectTypesDataTypes{
	}
	if err := valid.Validate(); err != nil {
		t.Errorf("Valid object should not have validation errors: %v", err)
	}

}

func TestProjectTypesPousPouInterfaceLocalVarsXMLMarshaling(t *testing.T) {
	original := &ProjectTypesPousPouInterfaceLocalVars{
	}

	// Test marshaling
	xmlData, err := xml.Marshal(original)
	if err != nil {
		t.Fatalf("Failed to marshal XML: %v", err)
	}

	// Test unmarshaling
	var unmarshaled ProjectTypesPousPouInterfaceLocalVars
	err = xml.Unmarshal(xmlData, &unmarshaled)
	if err != nil {
		t.Fatalf("Failed to unmarshal XML: %v", err)
	}

	// Test validation
	if err := original.Validate(); err != nil {
		t.Errorf("Validation failed: %v", err)
	}
}

func TestProjectTypesPousPouInterfaceLocalVarsValidation(t *testing.T) {
	// Test valid case
	valid := &ProjectTypesPousPouInterfaceLocalVars{
	}
	if err := valid.Validate(); err != nil {
		t.Errorf("Valid object should not have validation errors: %v", err)
	}

}

func TestProjectTypesPousPouInterfaceTempVarsXMLMarshaling(t *testing.T) {
	original := &ProjectTypesPousPouInterfaceTempVars{
	}

	// Test marshaling
	xmlData, err := xml.Marshal(original)
	if err != nil {
		t.Fatalf("Failed to marshal XML: %v", err)
	}

	// Test unmarshaling
	var unmarshaled ProjectTypesPousPouInterfaceTempVars
	err = xml.Unmarshal(xmlData, &unmarshaled)
	if err != nil {
		t.Fatalf("Failed to unmarshal XML: %v", err)
	}

	// Test validation
	if err := original.Validate(); err != nil {
		t.Errorf("Validation failed: %v", err)
	}
}

func TestProjectTypesPousPouInterfaceTempVarsValidation(t *testing.T) {
	// Test valid case
	valid := &ProjectTypesPousPouInterfaceTempVars{
	}
	if err := valid.Validate(); err != nil {
		t.Errorf("Valid object should not have validation errors: %v", err)
	}

}

func TestProjectTypesPousPouInterfaceInputVarsXMLMarshaling(t *testing.T) {
	original := &ProjectTypesPousPouInterfaceInputVars{
	}

	// Test marshaling
	xmlData, err := xml.Marshal(original)
	if err != nil {
		t.Fatalf("Failed to marshal XML: %v", err)
	}

	// Test unmarshaling
	var unmarshaled ProjectTypesPousPouInterfaceInputVars
	err = xml.Unmarshal(xmlData, &unmarshaled)
	if err != nil {
		t.Fatalf("Failed to unmarshal XML: %v", err)
	}

	// Test validation
	if err := original.Validate(); err != nil {
		t.Errorf("Validation failed: %v", err)
	}
}

func TestProjectTypesPousPouInterfaceInputVarsValidation(t *testing.T) {
	// Test valid case
	valid := &ProjectTypesPousPouInterfaceInputVars{
	}
	if err := valid.Validate(); err != nil {
		t.Errorf("Valid object should not have validation errors: %v", err)
	}

}

func TestProjectTypesPousPouInterfaceOutputVarsXMLMarshaling(t *testing.T) {
	original := &ProjectTypesPousPouInterfaceOutputVars{
	}

	// Test marshaling
	xmlData, err := xml.Marshal(original)
	if err != nil {
		t.Fatalf("Failed to marshal XML: %v", err)
	}

	// Test unmarshaling
	var unmarshaled ProjectTypesPousPouInterfaceOutputVars
	err = xml.Unmarshal(xmlData, &unmarshaled)
	if err != nil {
		t.Fatalf("Failed to unmarshal XML: %v", err)
	}

	// Test validation
	if err := original.Validate(); err != nil {
		t.Errorf("Validation failed: %v", err)
	}
}

func TestProjectTypesPousPouInterfaceOutputVarsValidation(t *testing.T) {
	// Test valid case
	valid := &ProjectTypesPousPouInterfaceOutputVars{
	}
	if err := valid.Validate(); err != nil {
		t.Errorf("Valid object should not have validation errors: %v", err)
	}

}

func TestProjectTypesPousPouInterfaceInOutVarsXMLMarshaling(t *testing.T) {
	original := &ProjectTypesPousPouInterfaceInOutVars{
	}

	// Test marshaling
	xmlData, err := xml.Marshal(original)
	if err != nil {
		t.Fatalf("Failed to marshal XML: %v", err)
	}

	// Test unmarshaling
	var unmarshaled ProjectTypesPousPouInterfaceInOutVars
	err = xml.Unmarshal(xmlData, &unmarshaled)
	if err != nil {
		t.Fatalf("Failed to unmarshal XML: %v", err)
	}

	// Test validation
	if err := original.Validate(); err != nil {
		t.Errorf("Validation failed: %v", err)
	}
}

func TestProjectTypesPousPouInterfaceInOutVarsValidation(t *testing.T) {
	// Test valid case
	valid := &ProjectTypesPousPouInterfaceInOutVars{
	}
	if err := valid.Validate(); err != nil {
		t.Errorf("Valid object should not have validation errors: %v", err)
	}

}

func TestProjectTypesPousPouInterfaceExternalVarsXMLMarshaling(t *testing.T) {
	original := &ProjectTypesPousPouInterfaceExternalVars{
	}

	// Test marshaling
	xmlData, err := xml.Marshal(original)
	if err != nil {
		t.Fatalf("Failed to marshal XML: %v", err)
	}

	// Test unmarshaling
	var unmarshaled ProjectTypesPousPouInterfaceExternalVars
	err = xml.Unmarshal(xmlData, &unmarshaled)
	if err != nil {
		t.Fatalf("Failed to unmarshal XML: %v", err)
	}

	// Test validation
	if err := original.Validate(); err != nil {
		t.Errorf("Validation failed: %v", err)
	}
}

func TestProjectTypesPousPouInterfaceExternalVarsValidation(t *testing.T) {
	// Test valid case
	valid := &ProjectTypesPousPouInterfaceExternalVars{
	}
	if err := valid.Validate(); err != nil {
		t.Errorf("Valid object should not have validation errors: %v", err)
	}

}

func TestProjectTypesPousPouInterfaceGlobalVarsXMLMarshaling(t *testing.T) {
	original := &ProjectTypesPousPouInterfaceGlobalVars{
	}

	// Test marshaling
	xmlData, err := xml.Marshal(original)
	if err != nil {
		t.Fatalf("Failed to marshal XML: %v", err)
	}

	// Test unmarshaling
	var unmarshaled ProjectTypesPousPouInterfaceGlobalVars
	err = xml.Unmarshal(xmlData, &unmarshaled)
	if err != nil {
		t.Fatalf("Failed to unmarshal XML: %v", err)
	}

	// Test validation
	if err := original.Validate(); err != nil {
		t.Errorf("Validation failed: %v", err)
	}
}

func TestProjectTypesPousPouInterfaceGlobalVarsValidation(t *testing.T) {
	// Test valid case
	valid := &ProjectTypesPousPouInterfaceGlobalVars{
	}
	if err := valid.Validate(); err != nil {
		t.Errorf("Valid object should not have validation errors: %v", err)
	}

}

func TestProjectTypesPousPouInterfaceXMLMarshaling(t *testing.T) {
	original := &ProjectTypesPousPouInterface{
		ReturnType: &DataType{},
		Documentation: &FormattedText{},
		LocalVars: &ProjectTypesPousPouInterfaceLocalVars{},
		TempVars: &ProjectTypesPousPouInterfaceTempVars{},
		InputVars: &ProjectTypesPousPouInterfaceInputVars{},
		OutputVars: &ProjectTypesPousPouInterfaceOutputVars{},
		InOutVars: &ProjectTypesPousPouInterfaceInOutVars{},
		ExternalVars: &ProjectTypesPousPouInterfaceExternalVars{},
		GlobalVars: &ProjectTypesPousPouInterfaceGlobalVars{},
		AccessVars: &VarList{},
	}

	// Test marshaling
	xmlData, err := xml.Marshal(original)
	if err != nil {
		t.Fatalf("Failed to marshal XML: %v", err)
	}

	// Test unmarshaling
	var unmarshaled ProjectTypesPousPouInterface
	err = xml.Unmarshal(xmlData, &unmarshaled)
	if err != nil {
		t.Fatalf("Failed to unmarshal XML: %v", err)
	}

	// Test validation
	if err := original.Validate(); err != nil {
		t.Errorf("Validation failed: %v", err)
	}
}

func TestProjectTypesPousPouInterfaceValidation(t *testing.T) {
	// Test valid case
	valid := &ProjectTypesPousPouInterface{
	}
	if err := valid.Validate(); err != nil {
		t.Errorf("Valid object should not have validation errors: %v", err)
	}

}

func TestProjectTypesPousPouActionsActionXMLMarshaling(t *testing.T) {
	original := &ProjectTypesPousPouActionsAction{
		Body: Body{},
		Documentation: &FormattedText{},
		Name: "test_value",
	}

	// Test marshaling
	xmlData, err := xml.Marshal(original)
	if err != nil {
		t.Fatalf("Failed to marshal XML: %v", err)
	}

	// Test unmarshaling
	var unmarshaled ProjectTypesPousPouActionsAction
	err = xml.Unmarshal(xmlData, &unmarshaled)
	if err != nil {
		t.Fatalf("Failed to unmarshal XML: %v", err)
	}

	// Test validation
	if err := original.Validate(); err != nil {
		t.Errorf("Validation failed: %v", err)
	}
}

func TestProjectTypesPousPouActionsActionValidation(t *testing.T) {
	// Test valid case
	valid := &ProjectTypesPousPouActionsAction{
		Body: Body{},
		Name: "test_value",
	}
	if err := valid.Validate(); err != nil {
		t.Errorf("Valid object should not have validation errors: %v", err)
	}

	// Test missing required field: Body
	invalidBody := &ProjectTypesPousPouActionsAction{}
	if err := invalidBody.Validate(); err == nil {
		t.Error("Missing required field Body should cause validation error")
	}

	// Test missing required field: Name
	invalidName := &ProjectTypesPousPouActionsAction{}
	if err := invalidName.Validate(); err == nil {
		t.Error("Missing required field Name should cause validation error")
	}

}

func TestProjectTypesPousPouActionsXMLMarshaling(t *testing.T) {
	original := &ProjectTypesPousPouActions{
		Action: []ProjectTypesPousPouActionsAction{ProjectTypesPousPouActionsAction{}},
	}

	// Test marshaling
	xmlData, err := xml.Marshal(original)
	if err != nil {
		t.Fatalf("Failed to marshal XML: %v", err)
	}

	// Test unmarshaling
	var unmarshaled ProjectTypesPousPouActions
	err = xml.Unmarshal(xmlData, &unmarshaled)
	if err != nil {
		t.Fatalf("Failed to unmarshal XML: %v", err)
	}

	// Test validation
	if err := original.Validate(); err != nil {
		t.Errorf("Validation failed: %v", err)
	}
}

func TestProjectTypesPousPouActionsValidation(t *testing.T) {
	// Test valid case
	valid := &ProjectTypesPousPouActions{
	}
	if err := valid.Validate(); err != nil {
		t.Errorf("Valid object should not have validation errors: %v", err)
	}

}

func TestProjectTypesPousPouTransitionsTransitionXMLMarshaling(t *testing.T) {
	original := &ProjectTypesPousPouTransitionsTransition{
		Body: Body{},
		Documentation: &FormattedText{},
		Name: "test_value",
	}

	// Test marshaling
	xmlData, err := xml.Marshal(original)
	if err != nil {
		t.Fatalf("Failed to marshal XML: %v", err)
	}

	// Test unmarshaling
	var unmarshaled ProjectTypesPousPouTransitionsTransition
	err = xml.Unmarshal(xmlData, &unmarshaled)
	if err != nil {
		t.Fatalf("Failed to unmarshal XML: %v", err)
	}

	// Test validation
	if err := original.Validate(); err != nil {
		t.Errorf("Validation failed: %v", err)
	}
}

func TestProjectTypesPousPouTransitionsTransitionValidation(t *testing.T) {
	// Test valid case
	valid := &ProjectTypesPousPouTransitionsTransition{
		Body: Body{},
		Name: "test_value",
	}
	if err := valid.Validate(); err != nil {
		t.Errorf("Valid object should not have validation errors: %v", err)
	}

	// Test missing required field: Body
	invalidBody := &ProjectTypesPousPouTransitionsTransition{}
	if err := invalidBody.Validate(); err == nil {
		t.Error("Missing required field Body should cause validation error")
	}

	// Test missing required field: Name
	invalidName := &ProjectTypesPousPouTransitionsTransition{}
	if err := invalidName.Validate(); err == nil {
		t.Error("Missing required field Name should cause validation error")
	}

}

func TestProjectTypesPousPouTransitionsXMLMarshaling(t *testing.T) {
	original := &ProjectTypesPousPouTransitions{
		Transition: []ProjectTypesPousPouTransitionsTransition{ProjectTypesPousPouTransitionsTransition{}},
	}

	// Test marshaling
	xmlData, err := xml.Marshal(original)
	if err != nil {
		t.Fatalf("Failed to marshal XML: %v", err)
	}

	// Test unmarshaling
	var unmarshaled ProjectTypesPousPouTransitions
	err = xml.Unmarshal(xmlData, &unmarshaled)
	if err != nil {
		t.Fatalf("Failed to unmarshal XML: %v", err)
	}

	// Test validation
	if err := original.Validate(); err != nil {
		t.Errorf("Validation failed: %v", err)
	}
}

func TestProjectTypesPousPouTransitionsValidation(t *testing.T) {
	// Test valid case
	valid := &ProjectTypesPousPouTransitions{
	}
	if err := valid.Validate(); err != nil {
		t.Errorf("Valid object should not have validation errors: %v", err)
	}

}

func TestProjectTypesPousPouXMLMarshaling(t *testing.T) {
	original := &ProjectTypesPousPou{
		Interface: &ProjectTypesPousPouInterface{},
		Actions: &ProjectTypesPousPouActions{},
		Transitions: &ProjectTypesPousPouTransitions{},
		Body: &Body{},
		Documentation: &FormattedText{},
		Name: "test_value",
		PouType: PouType(""),
	}

	// Test marshaling
	xmlData, err := xml.Marshal(original)
	if err != nil {
		t.Fatalf("Failed to marshal XML: %v", err)
	}

	// Test unmarshaling
	var unmarshaled ProjectTypesPousPou
	err = xml.Unmarshal(xmlData, &unmarshaled)
	if err != nil {
		t.Fatalf("Failed to unmarshal XML: %v", err)
	}

	// Test validation
	if err := original.Validate(); err != nil {
		t.Errorf("Validation failed: %v", err)
	}
}

func TestProjectTypesPousPouValidation(t *testing.T) {
	// Test valid case
	valid := &ProjectTypesPousPou{
		Name: "test_value",
		PouType: PouType(""),
	}
	if err := valid.Validate(); err != nil {
		t.Errorf("Valid object should not have validation errors: %v", err)
	}

	// Test missing required field: Name
	invalidName := &ProjectTypesPousPou{}
	if err := invalidName.Validate(); err == nil {
		t.Error("Missing required field Name should cause validation error")
	}

	// Test missing required field: PouType
	invalidPouType := &ProjectTypesPousPou{}
	if err := invalidPouType.Validate(); err == nil {
		t.Error("Missing required field PouType should cause validation error")
	}

}

func TestProjectTypesPousXMLMarshaling(t *testing.T) {
	original := &ProjectTypesPous{
		Pou: []ProjectTypesPousPou{ProjectTypesPousPou{}},
	}

	// Test marshaling
	xmlData, err := xml.Marshal(original)
	if err != nil {
		t.Fatalf("Failed to marshal XML: %v", err)
	}

	// Test unmarshaling
	var unmarshaled ProjectTypesPous
	err = xml.Unmarshal(xmlData, &unmarshaled)
	if err != nil {
		t.Fatalf("Failed to unmarshal XML: %v", err)
	}

	// Test validation
	if err := original.Validate(); err != nil {
		t.Errorf("Validation failed: %v", err)
	}
}

func TestProjectTypesPousValidation(t *testing.T) {
	// Test valid case
	valid := &ProjectTypesPous{
	}
	if err := valid.Validate(); err != nil {
		t.Errorf("Valid object should not have validation errors: %v", err)
	}

}

func TestProjectTypesXMLMarshaling(t *testing.T) {
	original := &ProjectTypes{
		DataTypes: ProjectTypesDataTypes{},
		Pous: ProjectTypesPous{},
	}

	// Test marshaling
	xmlData, err := xml.Marshal(original)
	if err != nil {
		t.Fatalf("Failed to marshal XML: %v", err)
	}

	// Test unmarshaling
	var unmarshaled ProjectTypes
	err = xml.Unmarshal(xmlData, &unmarshaled)
	if err != nil {
		t.Fatalf("Failed to unmarshal XML: %v", err)
	}

	// Test validation
	if err := original.Validate(); err != nil {
		t.Errorf("Validation failed: %v", err)
	}
}

func TestProjectTypesValidation(t *testing.T) {
	// Test valid case
	valid := &ProjectTypes{
		DataTypes: ProjectTypesDataTypes{},
		Pous: ProjectTypesPous{},
	}
	if err := valid.Validate(); err != nil {
		t.Errorf("Valid object should not have validation errors: %v", err)
	}

	// Test missing required field: DataTypes
	invalidDataTypes := &ProjectTypes{}
	if err := invalidDataTypes.Validate(); err == nil {
		t.Error("Missing required field DataTypes should cause validation error")
	}

	// Test missing required field: Pous
	invalidPous := &ProjectTypes{}
	if err := invalidPous.Validate(); err == nil {
		t.Error("Missing required field Pous should cause validation error")
	}

}

func TestProjectInstancesConfigurationsConfigurationResourceTaskXMLMarshaling(t *testing.T) {
	original := &ProjectInstancesConfigurationsConfigurationResourceTask{
		PouInstance: []PouInstance{PouInstance{}},
		Documentation: &FormattedText{},
		Name: "test_value",
		Single: stringPtr("test_value"),
		Interval: stringPtr("test_value"),
		Priority: "test_value",
	}

	// Test marshaling
	xmlData, err := xml.Marshal(original)
	if err != nil {
		t.Fatalf("Failed to marshal XML: %v", err)
	}

	// Test unmarshaling
	var unmarshaled ProjectInstancesConfigurationsConfigurationResourceTask
	err = xml.Unmarshal(xmlData, &unmarshaled)
	if err != nil {
		t.Fatalf("Failed to unmarshal XML: %v", err)
	}

	// Test validation
	if err := original.Validate(); err != nil {
		t.Errorf("Validation failed: %v", err)
	}
}

func TestProjectInstancesConfigurationsConfigurationResourceTaskValidation(t *testing.T) {
	// Test valid case
	valid := &ProjectInstancesConfigurationsConfigurationResourceTask{
		Name: "test_value",
		Priority: "test_value",
	}
	if err := valid.Validate(); err != nil {
		t.Errorf("Valid object should not have validation errors: %v", err)
	}

	// Test missing required field: Name
	invalidName := &ProjectInstancesConfigurationsConfigurationResourceTask{}
	if err := invalidName.Validate(); err == nil {
		t.Error("Missing required field Name should cause validation error")
	}

	// Test missing required field: Priority
	invalidPriority := &ProjectInstancesConfigurationsConfigurationResourceTask{}
	if err := invalidPriority.Validate(); err == nil {
		t.Error("Missing required field Priority should cause validation error")
	}

}

func TestProjectInstancesConfigurationsConfigurationResourceXMLMarshaling(t *testing.T) {
	original := &ProjectInstancesConfigurationsConfigurationResource{
		Task: []ProjectInstancesConfigurationsConfigurationResourceTask{ProjectInstancesConfigurationsConfigurationResourceTask{}},
		GlobalVars: []VarList{VarList{}},
		PouInstance: []PouInstance{PouInstance{}},
		Documentation: &FormattedText{},
		Name: "test_value",
	}

	// Test marshaling
	xmlData, err := xml.Marshal(original)
	if err != nil {
		t.Fatalf("Failed to marshal XML: %v", err)
	}

	// Test unmarshaling
	var unmarshaled ProjectInstancesConfigurationsConfigurationResource
	err = xml.Unmarshal(xmlData, &unmarshaled)
	if err != nil {
		t.Fatalf("Failed to unmarshal XML: %v", err)
	}

	// Test validation
	if err := original.Validate(); err != nil {
		t.Errorf("Validation failed: %v", err)
	}
}

func TestProjectInstancesConfigurationsConfigurationResourceValidation(t *testing.T) {
	// Test valid case
	valid := &ProjectInstancesConfigurationsConfigurationResource{
		Name: "test_value",
	}
	if err := valid.Validate(); err != nil {
		t.Errorf("Valid object should not have validation errors: %v", err)
	}

	// Test missing required field: Name
	invalidName := &ProjectInstancesConfigurationsConfigurationResource{}
	if err := invalidName.Validate(); err == nil {
		t.Error("Missing required field Name should cause validation error")
	}

}

func TestProjectInstancesConfigurationsConfigurationXMLMarshaling(t *testing.T) {
	original := &ProjectInstancesConfigurationsConfiguration{
		Resource: []ProjectInstancesConfigurationsConfigurationResource{ProjectInstancesConfigurationsConfigurationResource{}},
		GlobalVars: []VarList{VarList{}},
		Documentation: &FormattedText{},
		Name: "test_value",
	}

	// Test marshaling
	xmlData, err := xml.Marshal(original)
	if err != nil {
		t.Fatalf("Failed to marshal XML: %v", err)
	}

	// Test unmarshaling
	var unmarshaled ProjectInstancesConfigurationsConfiguration
	err = xml.Unmarshal(xmlData, &unmarshaled)
	if err != nil {
		t.Fatalf("Failed to unmarshal XML: %v", err)
	}

	// Test validation
	if err := original.Validate(); err != nil {
		t.Errorf("Validation failed: %v", err)
	}
}

func TestProjectInstancesConfigurationsConfigurationValidation(t *testing.T) {
	// Test valid case
	valid := &ProjectInstancesConfigurationsConfiguration{
		Name: "test_value",
	}
	if err := valid.Validate(); err != nil {
		t.Errorf("Valid object should not have validation errors: %v", err)
	}

	// Test missing required field: Name
	invalidName := &ProjectInstancesConfigurationsConfiguration{}
	if err := invalidName.Validate(); err == nil {
		t.Error("Missing required field Name should cause validation error")
	}

}

func TestProjectInstancesConfigurationsXMLMarshaling(t *testing.T) {
	original := &ProjectInstancesConfigurations{
		Configuration: []ProjectInstancesConfigurationsConfiguration{ProjectInstancesConfigurationsConfiguration{}},
	}

	// Test marshaling
	xmlData, err := xml.Marshal(original)
	if err != nil {
		t.Fatalf("Failed to marshal XML: %v", err)
	}

	// Test unmarshaling
	var unmarshaled ProjectInstancesConfigurations
	err = xml.Unmarshal(xmlData, &unmarshaled)
	if err != nil {
		t.Fatalf("Failed to unmarshal XML: %v", err)
	}

	// Test validation
	if err := original.Validate(); err != nil {
		t.Errorf("Validation failed: %v", err)
	}
}

func TestProjectInstancesConfigurationsValidation(t *testing.T) {
	// Test valid case
	valid := &ProjectInstancesConfigurations{
	}
	if err := valid.Validate(); err != nil {
		t.Errorf("Valid object should not have validation errors: %v", err)
	}

}

func TestProjectInstancesXMLMarshaling(t *testing.T) {
	original := &ProjectInstances{
		Configurations: ProjectInstancesConfigurations{},
	}

	// Test marshaling
	xmlData, err := xml.Marshal(original)
	if err != nil {
		t.Fatalf("Failed to marshal XML: %v", err)
	}

	// Test unmarshaling
	var unmarshaled ProjectInstances
	err = xml.Unmarshal(xmlData, &unmarshaled)
	if err != nil {
		t.Fatalf("Failed to unmarshal XML: %v", err)
	}

	// Test validation
	if err := original.Validate(); err != nil {
		t.Errorf("Validation failed: %v", err)
	}
}

func TestProjectInstancesValidation(t *testing.T) {
	// Test valid case
	valid := &ProjectInstances{
		Configurations: ProjectInstancesConfigurations{},
	}
	if err := valid.Validate(); err != nil {
		t.Errorf("Valid object should not have validation errors: %v", err)
	}

	// Test missing required field: Configurations
	invalidConfigurations := &ProjectInstances{}
	if err := invalidConfigurations.Validate(); err == nil {
		t.Error("Missing required field Configurations should cause validation error")
	}

}

func TestProjectXMLMarshaling(t *testing.T) {
	original := &Project{
		FileHeader: ProjectFileHeader{},
		ContentHeader: ProjectContentHeader{},
		Types: ProjectTypes{},
		Instances: ProjectInstances{},
	}

	// Test marshaling
	xmlData, err := xml.Marshal(original)
	if err != nil {
		t.Fatalf("Failed to marshal XML: %v", err)
	}

	// Test unmarshaling
	var unmarshaled Project
	err = xml.Unmarshal(xmlData, &unmarshaled)
	if err != nil {
		t.Fatalf("Failed to unmarshal XML: %v", err)
	}

	// Test validation
	if err := original.Validate(); err != nil {
		t.Errorf("Validation failed: %v", err)
	}
}

func TestProjectValidation(t *testing.T) {
	// Test valid case
	valid := &Project{
		FileHeader: ProjectFileHeader{},
		ContentHeader: ProjectContentHeader{},
		Types: ProjectTypes{},
		Instances: ProjectInstances{},
	}
	if err := valid.Validate(); err != nil {
		t.Errorf("Valid object should not have validation errors: %v", err)
	}

	// Test missing required field: FileHeader
	invalidFileHeader := &Project{}
	if err := invalidFileHeader.Validate(); err == nil {
		t.Error("Missing required field FileHeader should cause validation error")
	}

	// Test missing required field: ContentHeader
	invalidContentHeader := &Project{}
	if err := invalidContentHeader.Validate(); err == nil {
		t.Error("Missing required field ContentHeader should cause validation error")
	}

	// Test missing required field: Types
	invalidTypes := &Project{}
	if err := invalidTypes.Validate(); err == nil {
		t.Error("Missing required field Types should cause validation error")
	}

	// Test missing required field: Instances
	invalidInstances := &Project{}
	if err := invalidInstances.Validate(); err == nil {
		t.Error("Missing required field Instances should cause validation error")
	}

}

// Benchmark tests

func BenchmarkDataTypeStringMarshaling(b *testing.B) {
	obj := &DataTypeString{
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := xml.Marshal(obj)
		if err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkDataTypeWstringMarshaling(b *testing.B) {
	obj := &DataTypeWstring{
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := xml.Marshal(obj)
		if err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkDataTypeArrayMarshaling(b *testing.B) {
	obj := &DataTypeArray{
		Dimension: []RangeSigned{RangeSigned{}},
		BaseType: DataType{},
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := xml.Marshal(obj)
		if err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkDataTypeDerivedMarshaling(b *testing.B) {
	obj := &DataTypeDerived{
		Name: "test_value",
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := xml.Marshal(obj)
		if err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkDataTypeEnumValuesValueMarshaling(b *testing.B) {
	obj := &DataTypeEnumValuesValue{
		Name: "test_value",
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := xml.Marshal(obj)
		if err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkDataTypeEnumValuesMarshaling(b *testing.B) {
	obj := &DataTypeEnumValues{
		Value: DataTypeEnumValuesValue{},
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := xml.Marshal(obj)
		if err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkDataTypeEnumMarshaling(b *testing.B) {
	obj := &DataTypeEnum{
		Values: DataTypeEnumValues{},
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := xml.Marshal(obj)
		if err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkDataTypeSubrangeSignedMarshaling(b *testing.B) {
	obj := &DataTypeSubrangeSigned{
		Range: RangeSigned{},
		BaseType: DataType{},
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := xml.Marshal(obj)
		if err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkDataTypeSubrangeUnsignedMarshaling(b *testing.B) {
	obj := &DataTypeSubrangeUnsigned{
		Range: RangeUnsigned{},
		BaseType: DataType{},
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := xml.Marshal(obj)
		if err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkDataTypePointerMarshaling(b *testing.B) {
	obj := &DataTypePointer{
		BaseType: DataType{},
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := xml.Marshal(obj)
		if err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkDataTypeMarshaling(b *testing.B) {
	obj := &DataType{
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := xml.Marshal(obj)
		if err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkRangeSignedMarshaling(b *testing.B) {
	obj := &RangeSigned{
		Lower: 42,
		Upper: 42,
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := xml.Marshal(obj)
		if err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkRangeUnsignedMarshaling(b *testing.B) {
	obj := &RangeUnsigned{
		Lower: 42,
		Upper: 42,
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := xml.Marshal(obj)
		if err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkValueSimpleValueMarshaling(b *testing.B) {
	obj := &ValueSimpleValue{
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := xml.Marshal(obj)
		if err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkValueArrayValueValueMarshaling(b *testing.B) {
	obj := &ValueArrayValueValue{
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := xml.Marshal(obj)
		if err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkValueArrayValueMarshaling(b *testing.B) {
	obj := &ValueArrayValue{
		Value: ValueArrayValueValue{},
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := xml.Marshal(obj)
		if err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkValueStructValueValueMarshaling(b *testing.B) {
	obj := &ValueStructValueValue{
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := xml.Marshal(obj)
		if err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkValueStructValueMarshaling(b *testing.B) {
	obj := &ValueStructValue{
		Value: ValueStructValueValue{},
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := xml.Marshal(obj)
		if err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkValueMarshaling(b *testing.B) {
	obj := &Value{
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := xml.Marshal(obj)
		if err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkBodyFBDCommentMarshaling(b *testing.B) {
	obj := &BodyFBDComment{
		Position: Position{},
		Content: FormattedText{},
		LocalId: 42,
		Height: 3.14,
		Width: 3.14,
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := xml.Marshal(obj)
		if err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkBodyFBDErrorMarshaling(b *testing.B) {
	obj := &BodyFBDError{
		Position: Position{},
		Content: FormattedText{},
		LocalId: 42,
		Height: 3.14,
		Width: 3.14,
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := xml.Marshal(obj)
		if err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkBodyFBDConnectorMarshaling(b *testing.B) {
	obj := &BodyFBDConnector{
		Position: Position{},
		Name: "test_value",
		LocalId: 42,
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := xml.Marshal(obj)
		if err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkBodyFBDContinuationMarshaling(b *testing.B) {
	obj := &BodyFBDContinuation{
		Position: Position{},
		Name: "test_value",
		LocalId: 42,
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := xml.Marshal(obj)
		if err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkBodyFBDActionBlockActionReferenceMarshaling(b *testing.B) {
	obj := &BodyFBDActionBlockActionReference{
		Name: "test_value",
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := xml.Marshal(obj)
		if err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkBodyFBDActionBlockActionMarshaling(b *testing.B) {
	obj := &BodyFBDActionBlockAction{
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := xml.Marshal(obj)
		if err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkBodyFBDActionBlockMarshaling(b *testing.B) {
	obj := &BodyFBDActionBlock{
		Position: Position{},
		LocalId: 42,
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := xml.Marshal(obj)
		if err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkBodyFBDBlockInputVariablesVariableMarshaling(b *testing.B) {
	obj := &BodyFBDBlockInputVariablesVariable{
		ConnectionPointIn: ConnectionPointIn{},
		FormalParameter: "test_value",
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := xml.Marshal(obj)
		if err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkBodyFBDBlockInputVariablesMarshaling(b *testing.B) {
	obj := &BodyFBDBlockInputVariables{
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := xml.Marshal(obj)
		if err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkBodyFBDBlockInOutVariablesVariableMarshaling(b *testing.B) {
	obj := &BodyFBDBlockInOutVariablesVariable{
		FormalParameter: "test_value",
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := xml.Marshal(obj)
		if err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkBodyFBDBlockInOutVariablesMarshaling(b *testing.B) {
	obj := &BodyFBDBlockInOutVariables{
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := xml.Marshal(obj)
		if err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkBodyFBDBlockOutputVariablesVariableMarshaling(b *testing.B) {
	obj := &BodyFBDBlockOutputVariablesVariable{
		FormalParameter: "test_value",
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := xml.Marshal(obj)
		if err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkBodyFBDBlockOutputVariablesMarshaling(b *testing.B) {
	obj := &BodyFBDBlockOutputVariables{
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := xml.Marshal(obj)
		if err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkBodyFBDBlockMarshaling(b *testing.B) {
	obj := &BodyFBDBlock{
		Position: Position{},
		InputVariables: BodyFBDBlockInputVariables{},
		InOutVariables: BodyFBDBlockInOutVariables{},
		OutputVariables: BodyFBDBlockOutputVariables{},
		LocalId: 42,
		TypeName: "test_value",
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := xml.Marshal(obj)
		if err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkBodyFBDInVariableMarshaling(b *testing.B) {
	obj := &BodyFBDInVariable{
		Position: Position{},
		Expression: "test_value",
		LocalId: 42,
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := xml.Marshal(obj)
		if err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkBodyFBDOutVariableMarshaling(b *testing.B) {
	obj := &BodyFBDOutVariable{
		Position: Position{},
		Expression: "test_value",
		LocalId: 42,
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := xml.Marshal(obj)
		if err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkBodyFBDInOutVariableMarshaling(b *testing.B) {
	obj := &BodyFBDInOutVariable{
		Position: Position{},
		Expression: "test_value",
		LocalId: 42,
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := xml.Marshal(obj)
		if err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkBodyFBDLabelMarshaling(b *testing.B) {
	obj := &BodyFBDLabel{
		Position: Position{},
		LocalId: 42,
		Label: "test_value",
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := xml.Marshal(obj)
		if err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkBodyFBDJumpMarshaling(b *testing.B) {
	obj := &BodyFBDJump{
		Position: Position{},
		LocalId: 42,
		Label: "test_value",
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := xml.Marshal(obj)
		if err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkBodyFBDReturnMarshaling(b *testing.B) {
	obj := &BodyFBDReturn{
		Position: Position{},
		LocalId: 42,
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := xml.Marshal(obj)
		if err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkBodyFBDMarshaling(b *testing.B) {
	obj := &BodyFBD{
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := xml.Marshal(obj)
		if err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkBodyLDCommentMarshaling(b *testing.B) {
	obj := &BodyLDComment{
		Position: Position{},
		Content: FormattedText{},
		LocalId: 42,
		Height: 3.14,
		Width: 3.14,
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := xml.Marshal(obj)
		if err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkBodyLDErrorMarshaling(b *testing.B) {
	obj := &BodyLDError{
		Position: Position{},
		Content: FormattedText{},
		LocalId: 42,
		Height: 3.14,
		Width: 3.14,
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := xml.Marshal(obj)
		if err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkBodyLDConnectorMarshaling(b *testing.B) {
	obj := &BodyLDConnector{
		Position: Position{},
		Name: "test_value",
		LocalId: 42,
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := xml.Marshal(obj)
		if err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkBodyLDContinuationMarshaling(b *testing.B) {
	obj := &BodyLDContinuation{
		Position: Position{},
		Name: "test_value",
		LocalId: 42,
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := xml.Marshal(obj)
		if err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkBodyLDActionBlockActionReferenceMarshaling(b *testing.B) {
	obj := &BodyLDActionBlockActionReference{
		Name: "test_value",
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := xml.Marshal(obj)
		if err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkBodyLDActionBlockActionMarshaling(b *testing.B) {
	obj := &BodyLDActionBlockAction{
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := xml.Marshal(obj)
		if err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkBodyLDActionBlockMarshaling(b *testing.B) {
	obj := &BodyLDActionBlock{
		Position: Position{},
		LocalId: 42,
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := xml.Marshal(obj)
		if err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkBodyLDBlockInputVariablesVariableMarshaling(b *testing.B) {
	obj := &BodyLDBlockInputVariablesVariable{
		ConnectionPointIn: ConnectionPointIn{},
		FormalParameter: "test_value",
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := xml.Marshal(obj)
		if err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkBodyLDBlockInputVariablesMarshaling(b *testing.B) {
	obj := &BodyLDBlockInputVariables{
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := xml.Marshal(obj)
		if err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkBodyLDBlockInOutVariablesVariableMarshaling(b *testing.B) {
	obj := &BodyLDBlockInOutVariablesVariable{
		FormalParameter: "test_value",
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := xml.Marshal(obj)
		if err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkBodyLDBlockInOutVariablesMarshaling(b *testing.B) {
	obj := &BodyLDBlockInOutVariables{
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := xml.Marshal(obj)
		if err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkBodyLDBlockOutputVariablesVariableMarshaling(b *testing.B) {
	obj := &BodyLDBlockOutputVariablesVariable{
		FormalParameter: "test_value",
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := xml.Marshal(obj)
		if err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkBodyLDBlockOutputVariablesMarshaling(b *testing.B) {
	obj := &BodyLDBlockOutputVariables{
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := xml.Marshal(obj)
		if err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkBodyLDBlockMarshaling(b *testing.B) {
	obj := &BodyLDBlock{
		Position: Position{},
		InputVariables: BodyLDBlockInputVariables{},
		InOutVariables: BodyLDBlockInOutVariables{},
		OutputVariables: BodyLDBlockOutputVariables{},
		LocalId: 42,
		TypeName: "test_value",
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := xml.Marshal(obj)
		if err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkBodyLDInVariableMarshaling(b *testing.B) {
	obj := &BodyLDInVariable{
		Position: Position{},
		Expression: "test_value",
		LocalId: 42,
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := xml.Marshal(obj)
		if err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkBodyLDOutVariableMarshaling(b *testing.B) {
	obj := &BodyLDOutVariable{
		Position: Position{},
		Expression: "test_value",
		LocalId: 42,
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := xml.Marshal(obj)
		if err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkBodyLDInOutVariableMarshaling(b *testing.B) {
	obj := &BodyLDInOutVariable{
		Position: Position{},
		Expression: "test_value",
		LocalId: 42,
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := xml.Marshal(obj)
		if err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkBodyLDLabelMarshaling(b *testing.B) {
	obj := &BodyLDLabel{
		Position: Position{},
		LocalId: 42,
		Label: "test_value",
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := xml.Marshal(obj)
		if err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkBodyLDJumpMarshaling(b *testing.B) {
	obj := &BodyLDJump{
		Position: Position{},
		LocalId: 42,
		Label: "test_value",
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := xml.Marshal(obj)
		if err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkBodyLDReturnMarshaling(b *testing.B) {
	obj := &BodyLDReturn{
		Position: Position{},
		LocalId: 42,
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := xml.Marshal(obj)
		if err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkBodyLDLeftPowerRailConnectionPointOutMarshaling(b *testing.B) {
	obj := &BodyLDLeftPowerRailConnectionPointOut{
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := xml.Marshal(obj)
		if err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkBodyLDLeftPowerRailMarshaling(b *testing.B) {
	obj := &BodyLDLeftPowerRail{
		Position: Position{},
		LocalId: 42,
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := xml.Marshal(obj)
		if err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkBodyLDRightPowerRailMarshaling(b *testing.B) {
	obj := &BodyLDRightPowerRail{
		Position: Position{},
		LocalId: 42,
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := xml.Marshal(obj)
		if err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkBodyLDCoilMarshaling(b *testing.B) {
	obj := &BodyLDCoil{
		Position: Position{},
		Variable: "test_value",
		LocalId: 42,
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := xml.Marshal(obj)
		if err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkBodyLDContactMarshaling(b *testing.B) {
	obj := &BodyLDContact{
		Position: Position{},
		Variable: "test_value",
		LocalId: 42,
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := xml.Marshal(obj)
		if err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkBodyLDMarshaling(b *testing.B) {
	obj := &BodyLD{
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := xml.Marshal(obj)
		if err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkBodySFCCommentMarshaling(b *testing.B) {
	obj := &BodySFCComment{
		Position: Position{},
		Content: FormattedText{},
		LocalId: 42,
		Height: 3.14,
		Width: 3.14,
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := xml.Marshal(obj)
		if err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkBodySFCErrorMarshaling(b *testing.B) {
	obj := &BodySFCError{
		Position: Position{},
		Content: FormattedText{},
		LocalId: 42,
		Height: 3.14,
		Width: 3.14,
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := xml.Marshal(obj)
		if err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkBodySFCConnectorMarshaling(b *testing.B) {
	obj := &BodySFCConnector{
		Position: Position{},
		Name: "test_value",
		LocalId: 42,
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := xml.Marshal(obj)
		if err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkBodySFCContinuationMarshaling(b *testing.B) {
	obj := &BodySFCContinuation{
		Position: Position{},
		Name: "test_value",
		LocalId: 42,
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := xml.Marshal(obj)
		if err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkBodySFCActionBlockActionReferenceMarshaling(b *testing.B) {
	obj := &BodySFCActionBlockActionReference{
		Name: "test_value",
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := xml.Marshal(obj)
		if err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkBodySFCActionBlockActionMarshaling(b *testing.B) {
	obj := &BodySFCActionBlockAction{
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := xml.Marshal(obj)
		if err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkBodySFCActionBlockMarshaling(b *testing.B) {
	obj := &BodySFCActionBlock{
		Position: Position{},
		LocalId: 42,
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := xml.Marshal(obj)
		if err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkBodySFCBlockInputVariablesVariableMarshaling(b *testing.B) {
	obj := &BodySFCBlockInputVariablesVariable{
		ConnectionPointIn: ConnectionPointIn{},
		FormalParameter: "test_value",
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := xml.Marshal(obj)
		if err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkBodySFCBlockInputVariablesMarshaling(b *testing.B) {
	obj := &BodySFCBlockInputVariables{
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := xml.Marshal(obj)
		if err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkBodySFCBlockInOutVariablesVariableMarshaling(b *testing.B) {
	obj := &BodySFCBlockInOutVariablesVariable{
		FormalParameter: "test_value",
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := xml.Marshal(obj)
		if err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkBodySFCBlockInOutVariablesMarshaling(b *testing.B) {
	obj := &BodySFCBlockInOutVariables{
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := xml.Marshal(obj)
		if err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkBodySFCBlockOutputVariablesVariableMarshaling(b *testing.B) {
	obj := &BodySFCBlockOutputVariablesVariable{
		FormalParameter: "test_value",
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := xml.Marshal(obj)
		if err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkBodySFCBlockOutputVariablesMarshaling(b *testing.B) {
	obj := &BodySFCBlockOutputVariables{
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := xml.Marshal(obj)
		if err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkBodySFCBlockMarshaling(b *testing.B) {
	obj := &BodySFCBlock{
		Position: Position{},
		InputVariables: BodySFCBlockInputVariables{},
		InOutVariables: BodySFCBlockInOutVariables{},
		OutputVariables: BodySFCBlockOutputVariables{},
		LocalId: 42,
		TypeName: "test_value",
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := xml.Marshal(obj)
		if err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkBodySFCInVariableMarshaling(b *testing.B) {
	obj := &BodySFCInVariable{
		Position: Position{},
		Expression: "test_value",
		LocalId: 42,
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := xml.Marshal(obj)
		if err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkBodySFCOutVariableMarshaling(b *testing.B) {
	obj := &BodySFCOutVariable{
		Position: Position{},
		Expression: "test_value",
		LocalId: 42,
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := xml.Marshal(obj)
		if err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkBodySFCInOutVariableMarshaling(b *testing.B) {
	obj := &BodySFCInOutVariable{
		Position: Position{},
		Expression: "test_value",
		LocalId: 42,
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := xml.Marshal(obj)
		if err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkBodySFCLabelMarshaling(b *testing.B) {
	obj := &BodySFCLabel{
		Position: Position{},
		LocalId: 42,
		Label: "test_value",
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := xml.Marshal(obj)
		if err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkBodySFCJumpMarshaling(b *testing.B) {
	obj := &BodySFCJump{
		Position: Position{},
		LocalId: 42,
		Label: "test_value",
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := xml.Marshal(obj)
		if err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkBodySFCReturnMarshaling(b *testing.B) {
	obj := &BodySFCReturn{
		Position: Position{},
		LocalId: 42,
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := xml.Marshal(obj)
		if err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkBodySFCLeftPowerRailConnectionPointOutMarshaling(b *testing.B) {
	obj := &BodySFCLeftPowerRailConnectionPointOut{
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := xml.Marshal(obj)
		if err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkBodySFCLeftPowerRailMarshaling(b *testing.B) {
	obj := &BodySFCLeftPowerRail{
		Position: Position{},
		LocalId: 42,
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := xml.Marshal(obj)
		if err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkBodySFCRightPowerRailMarshaling(b *testing.B) {
	obj := &BodySFCRightPowerRail{
		Position: Position{},
		LocalId: 42,
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := xml.Marshal(obj)
		if err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkBodySFCCoilMarshaling(b *testing.B) {
	obj := &BodySFCCoil{
		Position: Position{},
		Variable: "test_value",
		LocalId: 42,
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := xml.Marshal(obj)
		if err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkBodySFCContactMarshaling(b *testing.B) {
	obj := &BodySFCContact{
		Position: Position{},
		Variable: "test_value",
		LocalId: 42,
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := xml.Marshal(obj)
		if err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkBodySFCStepConnectionPointOutMarshaling(b *testing.B) {
	obj := &BodySFCStepConnectionPointOut{
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := xml.Marshal(obj)
		if err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkBodySFCStepConnectionPointOutActionMarshaling(b *testing.B) {
	obj := &BodySFCStepConnectionPointOutAction{
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := xml.Marshal(obj)
		if err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkBodySFCStepMarshaling(b *testing.B) {
	obj := &BodySFCStep{
		Position: Position{},
		LocalId: 42,
		Name: "test_value",
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := xml.Marshal(obj)
		if err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkBodySFCMacroStepMarshaling(b *testing.B) {
	obj := &BodySFCMacroStep{
		Position: Position{},
		LocalId: 42,
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := xml.Marshal(obj)
		if err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkBodySFCJumpStepMarshaling(b *testing.B) {
	obj := &BodySFCJumpStep{
		Position: Position{},
		LocalId: 42,
		TargetName: "test_value",
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := xml.Marshal(obj)
		if err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkBodySFCTransitionConditionReferenceMarshaling(b *testing.B) {
	obj := &BodySFCTransitionConditionReference{
		Name: "test_value",
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := xml.Marshal(obj)
		if err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkBodySFCTransitionConditionInlineMarshaling(b *testing.B) {
	obj := &BodySFCTransitionConditionInline{
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := xml.Marshal(obj)
		if err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkBodySFCTransitionConditionMarshaling(b *testing.B) {
	obj := &BodySFCTransitionCondition{
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := xml.Marshal(obj)
		if err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkBodySFCTransitionMarshaling(b *testing.B) {
	obj := &BodySFCTransition{
		Position: Position{},
		LocalId: 42,
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := xml.Marshal(obj)
		if err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkBodySFCSelectionDivergenceConnectionPointOutMarshaling(b *testing.B) {
	obj := &BodySFCSelectionDivergenceConnectionPointOut{
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := xml.Marshal(obj)
		if err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkBodySFCSelectionDivergenceMarshaling(b *testing.B) {
	obj := &BodySFCSelectionDivergence{
		Position: Position{},
		LocalId: 42,
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := xml.Marshal(obj)
		if err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkBodySFCSelectionConvergenceConnectionPointInMarshaling(b *testing.B) {
	obj := &BodySFCSelectionConvergenceConnectionPointIn{
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := xml.Marshal(obj)
		if err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkBodySFCSelectionConvergenceMarshaling(b *testing.B) {
	obj := &BodySFCSelectionConvergence{
		Position: Position{},
		LocalId: 42,
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := xml.Marshal(obj)
		if err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkBodySFCSimultaneousDivergenceConnectionPointOutMarshaling(b *testing.B) {
	obj := &BodySFCSimultaneousDivergenceConnectionPointOut{
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := xml.Marshal(obj)
		if err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkBodySFCSimultaneousDivergenceMarshaling(b *testing.B) {
	obj := &BodySFCSimultaneousDivergence{
		Position: Position{},
		LocalId: 42,
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := xml.Marshal(obj)
		if err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkBodySFCSimultaneousConvergenceMarshaling(b *testing.B) {
	obj := &BodySFCSimultaneousConvergence{
		Position: Position{},
		LocalId: 42,
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := xml.Marshal(obj)
		if err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkBodySFCMarshaling(b *testing.B) {
	obj := &BodySFC{
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := xml.Marshal(obj)
		if err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkBodyMarshaling(b *testing.B) {
	obj := &Body{
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := xml.Marshal(obj)
		if err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkVarListMarshaling(b *testing.B) {
	obj := &VarList{
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := xml.Marshal(obj)
		if err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkVarListPlainVariableMarshaling(b *testing.B) {
	obj := &VarListPlainVariable{
		Type: DataType{},
		Name: "test_value",
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := xml.Marshal(obj)
		if err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkVarListPlainMarshaling(b *testing.B) {
	obj := &VarListPlain{
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := xml.Marshal(obj)
		if err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkPositionMarshaling(b *testing.B) {
	obj := &Position{
		X: 3.14,
		Y: 3.14,
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := xml.Marshal(obj)
		if err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkConnectionMarshaling(b *testing.B) {
	obj := &Connection{
		RefLocalId: 42,
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := xml.Marshal(obj)
		if err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkConnectionPointInMarshaling(b *testing.B) {
	obj := &ConnectionPointIn{
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := xml.Marshal(obj)
		if err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkConnectionPointOutMarshaling(b *testing.B) {
	obj := &ConnectionPointOut{
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := xml.Marshal(obj)
		if err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkPouInstanceMarshaling(b *testing.B) {
	obj := &PouInstance{
		Name: "test_value",
		Type: "test_value",
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := xml.Marshal(obj)
		if err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkFormattedTextMarshaling(b *testing.B) {
	obj := &FormattedText{
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := xml.Marshal(obj)
		if err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkEdgeModifierTypeMarshaling(b *testing.B) {
	obj := EdgeModifierType("")

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := xml.Marshal(obj)
		if err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkStorageModifierTypeMarshaling(b *testing.B) {
	obj := StorageModifierType("")

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := xml.Marshal(obj)
		if err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkPouTypeMarshaling(b *testing.B) {
	obj := PouType("")

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := xml.Marshal(obj)
		if err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkProjectFileHeaderMarshaling(b *testing.B) {
	obj := &ProjectFileHeader{
		CompanyName: "test_value",
		ProductName: "test_value",
		ProductVersion: "test_value",
		CreationDateTime: time.Now(),
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := xml.Marshal(obj)
		if err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkProjectContentHeaderCoordinateInfoPageSizeMarshaling(b *testing.B) {
	obj := &ProjectContentHeaderCoordinateInfoPageSize{
		X: 3.14,
		Y: 3.14,
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := xml.Marshal(obj)
		if err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkProjectContentHeaderCoordinateInfoFbdScalingMarshaling(b *testing.B) {
	obj := &ProjectContentHeaderCoordinateInfoFbdScaling{
		X: 3.14,
		Y: 3.14,
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := xml.Marshal(obj)
		if err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkProjectContentHeaderCoordinateInfoFbdMarshaling(b *testing.B) {
	obj := &ProjectContentHeaderCoordinateInfoFbd{
		Scaling: ProjectContentHeaderCoordinateInfoFbdScaling{},
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := xml.Marshal(obj)
		if err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkProjectContentHeaderCoordinateInfoLdScalingMarshaling(b *testing.B) {
	obj := &ProjectContentHeaderCoordinateInfoLdScaling{
		X: 3.14,
		Y: 3.14,
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := xml.Marshal(obj)
		if err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkProjectContentHeaderCoordinateInfoLdMarshaling(b *testing.B) {
	obj := &ProjectContentHeaderCoordinateInfoLd{
		Scaling: ProjectContentHeaderCoordinateInfoLdScaling{},
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := xml.Marshal(obj)
		if err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkProjectContentHeaderCoordinateInfoSfcScalingMarshaling(b *testing.B) {
	obj := &ProjectContentHeaderCoordinateInfoSfcScaling{
		X: 3.14,
		Y: 3.14,
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := xml.Marshal(obj)
		if err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkProjectContentHeaderCoordinateInfoSfcMarshaling(b *testing.B) {
	obj := &ProjectContentHeaderCoordinateInfoSfc{
		Scaling: ProjectContentHeaderCoordinateInfoSfcScaling{},
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := xml.Marshal(obj)
		if err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkProjectContentHeaderCoordinateInfoMarshaling(b *testing.B) {
	obj := &ProjectContentHeaderCoordinateInfo{
		Fbd: ProjectContentHeaderCoordinateInfoFbd{},
		Ld: ProjectContentHeaderCoordinateInfoLd{},
		Sfc: ProjectContentHeaderCoordinateInfoSfc{},
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := xml.Marshal(obj)
		if err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkProjectContentHeaderMarshaling(b *testing.B) {
	obj := &ProjectContentHeader{
		CoordinateInfo: ProjectContentHeaderCoordinateInfo{},
		Name: "test_value",
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := xml.Marshal(obj)
		if err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkProjectTypesDataTypesDataTypeMarshaling(b *testing.B) {
	obj := &ProjectTypesDataTypesDataType{
		BaseType: DataType{},
		Name: "test_value",
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := xml.Marshal(obj)
		if err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkProjectTypesDataTypesMarshaling(b *testing.B) {
	obj := &ProjectTypesDataTypes{
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := xml.Marshal(obj)
		if err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkProjectTypesPousPouInterfaceLocalVarsMarshaling(b *testing.B) {
	obj := &ProjectTypesPousPouInterfaceLocalVars{
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := xml.Marshal(obj)
		if err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkProjectTypesPousPouInterfaceTempVarsMarshaling(b *testing.B) {
	obj := &ProjectTypesPousPouInterfaceTempVars{
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := xml.Marshal(obj)
		if err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkProjectTypesPousPouInterfaceInputVarsMarshaling(b *testing.B) {
	obj := &ProjectTypesPousPouInterfaceInputVars{
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := xml.Marshal(obj)
		if err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkProjectTypesPousPouInterfaceOutputVarsMarshaling(b *testing.B) {
	obj := &ProjectTypesPousPouInterfaceOutputVars{
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := xml.Marshal(obj)
		if err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkProjectTypesPousPouInterfaceInOutVarsMarshaling(b *testing.B) {
	obj := &ProjectTypesPousPouInterfaceInOutVars{
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := xml.Marshal(obj)
		if err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkProjectTypesPousPouInterfaceExternalVarsMarshaling(b *testing.B) {
	obj := &ProjectTypesPousPouInterfaceExternalVars{
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := xml.Marshal(obj)
		if err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkProjectTypesPousPouInterfaceGlobalVarsMarshaling(b *testing.B) {
	obj := &ProjectTypesPousPouInterfaceGlobalVars{
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := xml.Marshal(obj)
		if err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkProjectTypesPousPouInterfaceMarshaling(b *testing.B) {
	obj := &ProjectTypesPousPouInterface{
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := xml.Marshal(obj)
		if err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkProjectTypesPousPouActionsActionMarshaling(b *testing.B) {
	obj := &ProjectTypesPousPouActionsAction{
		Body: Body{},
		Name: "test_value",
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := xml.Marshal(obj)
		if err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkProjectTypesPousPouActionsMarshaling(b *testing.B) {
	obj := &ProjectTypesPousPouActions{
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := xml.Marshal(obj)
		if err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkProjectTypesPousPouTransitionsTransitionMarshaling(b *testing.B) {
	obj := &ProjectTypesPousPouTransitionsTransition{
		Body: Body{},
		Name: "test_value",
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := xml.Marshal(obj)
		if err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkProjectTypesPousPouTransitionsMarshaling(b *testing.B) {
	obj := &ProjectTypesPousPouTransitions{
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := xml.Marshal(obj)
		if err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkProjectTypesPousPouMarshaling(b *testing.B) {
	obj := &ProjectTypesPousPou{
		Name: "test_value",
		PouType: PouType(""),
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := xml.Marshal(obj)
		if err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkProjectTypesPousMarshaling(b *testing.B) {
	obj := &ProjectTypesPous{
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := xml.Marshal(obj)
		if err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkProjectTypesMarshaling(b *testing.B) {
	obj := &ProjectTypes{
		DataTypes: ProjectTypesDataTypes{},
		Pous: ProjectTypesPous{},
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := xml.Marshal(obj)
		if err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkProjectInstancesConfigurationsConfigurationResourceTaskMarshaling(b *testing.B) {
	obj := &ProjectInstancesConfigurationsConfigurationResourceTask{
		Name: "test_value",
		Priority: "test_value",
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := xml.Marshal(obj)
		if err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkProjectInstancesConfigurationsConfigurationResourceMarshaling(b *testing.B) {
	obj := &ProjectInstancesConfigurationsConfigurationResource{
		Name: "test_value",
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := xml.Marshal(obj)
		if err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkProjectInstancesConfigurationsConfigurationMarshaling(b *testing.B) {
	obj := &ProjectInstancesConfigurationsConfiguration{
		Name: "test_value",
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := xml.Marshal(obj)
		if err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkProjectInstancesConfigurationsMarshaling(b *testing.B) {
	obj := &ProjectInstancesConfigurations{
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := xml.Marshal(obj)
		if err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkProjectInstancesMarshaling(b *testing.B) {
	obj := &ProjectInstances{
		Configurations: ProjectInstancesConfigurations{},
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := xml.Marshal(obj)
		if err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkProjectMarshaling(b *testing.B) {
	obj := &Project{
		FileHeader: ProjectFileHeader{},
		ContentHeader: ProjectContentHeader{},
		Types: ProjectTypes{},
		Instances: ProjectInstances{},
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := xml.Marshal(obj)
		if err != nil {
			b.Fatal(err)
		}
	}
}

// Helper functions for creating pointers

func stringPtr(s string) *string { return &s }
func intPtr(i int) *int { return &i }
func uintPtr(u uint64) *uint64 { return &u }
func floatPtr(f float64) *float64 { return &f }
func boolPtr(b bool) *bool { return &b }
func timePtr(t time.Time) *time.Time { return &t }
func durationPtr(d time.Duration) *time.Duration { return &d }
func edgeModifierTypePtr(e EdgeModifierType) *EdgeModifierType { return &e }
func storageModifierTypePtr(s StorageModifierType) *StorageModifierType { return &s }
