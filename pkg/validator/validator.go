package validator

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"regexp"
	"strconv"
	"strings"

	"github.com/suifei/xsd2code/pkg/types"
)

// XMLDocument represents a parsed XML document for validation
type XMLDocument struct {
	XMLName  xml.Name
	Attrs    []xml.Attr   `xml:",any,attr"`
	Content  string       `xml:",chardata"`
	Children []XMLElement `xml:",any"`
}

// XMLElement represents an XML element
type XMLElement struct {
	XMLName  xml.Name
	Attrs    []xml.Attr   `xml:",any,attr"`
	Content  string       `xml:",chardata"`
	Children []XMLElement `xml:",any"`
}

// ValidationContext holds validation state
type ValidationContext struct {
	errors   []ValidationError
	warnings []ValidationWarning
	line     int
	column   int
}

// XSDValidator validates XML against XSD schema
type XSDValidator struct {
	schema *types.XSDSchema
}

// NewXSDValidator creates a new XSD validator
func NewXSDValidator(schema *types.XSDSchema) *XSDValidator {
	return &XSDValidator{
		schema: schema,
	}
}

// ValidateXML validates an XML document against the schema
func (v *XSDValidator) ValidateXML(xmlPath string) error {
	// Read XML file
	content, err := ioutil.ReadFile(xmlPath)
	if err != nil {
		return fmt.Errorf("failed to read XML file: %v", err)
	}

	return v.ValidateXMLContent(content)
}

// ValidateXMLContent validates XML content against the schema
func (v *XSDValidator) ValidateXMLContent(content []byte) error {
	// Parse XML into a structured format for validation
	var doc XMLDocument
	if err := xml.Unmarshal(content, &doc); err != nil {
		return fmt.Errorf("failed to parse XML: %v", err)
	}

	// Perform validation against schema
	return v.validateDocument(&doc, content)
}

// validateDocument validates the entire XML document against the schema
func (v *XSDValidator) validateDocument(doc *XMLDocument, content []byte) error {
	if v.schema == nil {
		return fmt.Errorf("no schema available for validation")
	}

	ctx := &ValidationContext{
		errors:   make([]ValidationError, 0),
		warnings: make([]ValidationWarning, 0),
		line:     1,
		column:   1,
	}

	// Find the root element definition in schema
	rootElement := v.findRootElement(doc.XMLName)
	if rootElement == nil {
		return fmt.Errorf("root element '%s' not found in schema", doc.XMLName.Local)
	}

	// Validate root element
	if err := v.validateElement(doc.XMLName, doc.Attrs, doc.Children, rootElement, ctx); err != nil {
		return err
	}

	// Return first error if any
	if len(ctx.errors) > 0 {
		return fmt.Errorf(ctx.errors[0].Message)
	}

	return nil
}

// findRootElement finds the root element definition in the schema
func (v *XSDValidator) findRootElement(xmlName xml.Name) *types.XSDElement {
	for _, element := range v.schema.Elements {
		if element.Name == xmlName.Local {
			return &element
		}
	}
	return nil
}

// validateElement validates an XML element against its XSD definition
func (v *XSDValidator) validateElement(xmlName xml.Name, attrs []xml.Attr, children []XMLElement, elementDef *types.XSDElement, ctx *ValidationContext) error {
	// Validate attributes if the element has a complex type
	if elementDef.ComplexType != nil {
		if err := v.validateAttributes(attrs, elementDef.ComplexType, ctx); err != nil {
			return err
		}

		// Validate child elements
		if err := v.validateChildElements(children, elementDef.ComplexType, ctx); err != nil {
			return err
		}
	}

	return nil
}

// validateAttributes validates XML attributes against XSD attribute definitions
func (v *XSDValidator) validateAttributes(xmlAttrs []xml.Attr, complexType *types.XSDComplexType, ctx *ValidationContext) error {
	// Create a map of XML attributes for easy lookup
	attrMap := make(map[string]string)
	for _, attr := range xmlAttrs {
		attrMap[attr.Name.Local] = attr.Value
	}

	// Check all defined attributes
	for _, attrDef := range complexType.Attributes {
		attrName := attrDef.Name

		// Check if required attribute is present
		if attrDef.Use == "required" {
			if _, exists := attrMap[attrName]; !exists {
				ctx.errors = append(ctx.errors, ValidationError{
					Message: fmt.Sprintf("required attribute '%s' is missing", attrName),
					Line:    ctx.line,
					Column:  ctx.column,
					Element: attrName,
				})
			}
		}

		// Validate attribute type if present
		if value, exists := attrMap[attrName]; exists {
			if err := v.validateAttributeType(value, attrDef.Type, attrName, ctx); err != nil {
				return err
			}
		}
	}

	return nil
}

// validateAttributeType validates an attribute value against its XSD type
func (v *XSDValidator) validateAttributeType(value, xsdType, attrName string, ctx *ValidationContext) error {
	// Remove namespace prefix from type
	if colonIndex := strings.LastIndex(xsdType, ":"); colonIndex != -1 {
		xsdType = xsdType[colonIndex+1:]
	}

	switch xsdType {
	case "string":
		// String is always valid
		return nil
	case "int", "integer":
		if _, err := strconv.Atoi(value); err != nil {
			ctx.errors = append(ctx.errors, ValidationError{
				Message: fmt.Sprintf("attribute '%s' value '%s' is not a valid integer", attrName, value),
				Line:    ctx.line,
				Column:  ctx.column,
				Element: attrName,
			})
		}
	case "decimal", "double", "float":
		if _, err := strconv.ParseFloat(value, 64); err != nil {
			ctx.errors = append(ctx.errors, ValidationError{
				Message: fmt.Sprintf("attribute '%s' value '%s' is not a valid number", attrName, value),
				Line:    ctx.line,
				Column:  ctx.column,
				Element: attrName,
			})
		}
	case "boolean":
		if value != "true" && value != "false" && value != "1" && value != "0" {
			ctx.errors = append(ctx.errors, ValidationError{
				Message: fmt.Sprintf("attribute '%s' value '%s' is not a valid boolean", attrName, value),
				Line:    ctx.line,
				Column:  ctx.column,
				Element: attrName,
			})
		}
	case "dateTime":
		// Simple date-time validation
		if !v.isValidDateTime(value) {
			ctx.errors = append(ctx.errors, ValidationError{
				Message: fmt.Sprintf("attribute '%s' value '%s' is not a valid dateTime", attrName, value),
				Line:    ctx.line,
				Column:  ctx.column,
				Element: attrName,
			})
		}
	}

	return nil
}

// validateChildElements validates child elements against XSD sequence/choice/all definitions
func (v *XSDValidator) validateChildElements(children []XMLElement, complexType *types.XSDComplexType, ctx *ValidationContext) error {
	if complexType.Sequence != nil {
		return v.validateSequence(children, complexType.Sequence, ctx)
	}
	if complexType.Choice != nil {
		return v.validateChoice(children, complexType.Choice, ctx)
	}
	if complexType.All != nil {
		return v.validateAll(children, complexType.All, ctx)
	}
	return nil
}

// validateSequence validates elements in sequence
func (v *XSDValidator) validateSequence(children []XMLElement, sequence *types.XSDSequence, ctx *ValidationContext) error {
	childIndex := 0

	for _, elementDef := range sequence.Elements {
		min, max := types.ParseOccurs(elementDef.MinOccurs, elementDef.MaxOccurs)
		count := 0

		// Count matching elements
		for childIndex < len(children) {
			if children[childIndex].XMLName.Local == elementDef.Name {
				count++
				childIndex++
				if max != -1 && count >= max {
					break
				}
			} else {
				break
			}
		}
		// Check occurrence constraints
		if count < min {
			if count == 0 {
				ctx.errors = append(ctx.errors, ValidationError{
					Message: fmt.Sprintf("missing required element '%s'", elementDef.Name),
					Line:    ctx.line,
					Column:  ctx.column,
					Element: elementDef.Name,
				})
			} else {
				ctx.errors = append(ctx.errors, ValidationError{
					Message: fmt.Sprintf("element '%s' occurs %d times, minimum required is %d", elementDef.Name, count, min),
					Line:    ctx.line,
					Column:  ctx.column,
					Element: elementDef.Name,
				})
			}
		}

		if max != -1 && count > max {
			ctx.errors = append(ctx.errors, ValidationError{
				Message: fmt.Sprintf("element '%s' occurs %d times, maximum allowed is %d", elementDef.Name, count, max),
				Line:    ctx.line,
				Column:  ctx.column,
				Element: elementDef.Name,
			})
		}
	}

	// Check for unexpected elements
	if childIndex < len(children) {
		ctx.warnings = append(ctx.warnings, ValidationWarning{
			Message: fmt.Sprintf("unexpected element '%s' found", children[childIndex].XMLName.Local),
			Line:    ctx.line,
			Column:  ctx.column,
			Element: children[childIndex].XMLName.Local,
		})
	}

	return nil
}

// validateChoice validates choice content model
func (v *XSDValidator) validateChoice(children []XMLElement, choice *types.XSDChoice, ctx *ValidationContext) error {
	if len(children) == 0 {
		ctx.errors = append(ctx.errors, ValidationError{
			Message: "choice content model requires at least one element",
			Line:    ctx.line,
			Column:  ctx.column,
		})
		return nil
	}

	// Check if any child matches choice options
	for _, child := range children {
		found := false
		for _, choiceElement := range choice.Elements {
			if choiceElement.Name == child.XMLName.Local {
				found = true
				if err := v.validateElementStructure(child.XMLName, child.Attrs, child.Children, &choiceElement, ctx); err != nil {
					return err
				}
				break
			}
		}
		if !found {
			ctx.errors = append(ctx.errors, ValidationError{
				Message: fmt.Sprintf("element '%s' is not allowed in choice", child.XMLName.Local),
				Line:    ctx.line,
				Column:  ctx.column,
				Element: child.XMLName.Local,
			})
		}
	}

	return nil
}

// validateAll validates all elements (simplified implementation)
func (v *XSDValidator) validateAll(children []XMLElement, all *types.XSDAll, ctx *ValidationContext) error {
	// For all, each element can appear at most once
	elementCounts := make(map[string]int)

	for _, child := range children {
		elementCounts[child.XMLName.Local]++
	}

	for _, elementDef := range all.Elements {
		count := elementCounts[elementDef.Name]
		min, max := types.ParseOccurs(elementDef.MinOccurs, elementDef.MaxOccurs)

		if count < min {
			ctx.errors = append(ctx.errors, ValidationError{
				Message: fmt.Sprintf("element '%s' is required in 'all' group", elementDef.Name),
				Line:    ctx.line,
				Column:  ctx.column,
				Element: elementDef.Name,
			})
		}

		if max != -1 && count > max {
			ctx.errors = append(ctx.errors, ValidationError{
				Message: fmt.Sprintf("element '%s' appears too many times in 'all' group", elementDef.Name),
				Line:    ctx.line,
				Column:  ctx.column,
				Element: elementDef.Name,
			})
		}
	}

	return nil
}

// isValidDateTime performs basic date-time validation
func (v *XSDValidator) isValidDateTime(value string) bool {
	// Simple regex for ISO 8601 date-time format
	pattern := `^\d{4}-\d{2}-\d{2}T\d{2}:\d{2}:\d{2}(\.\d+)?(Z|[+-]\d{2}:\d{2})?$`
	matched, _ := regexp.MatchString(pattern, value)
	return matched
}

// validateBasicStructure performs comprehensive XSD validation
func (v *XSDValidator) validateBasicStructure(content []byte) error {
	if v.schema == nil {
		return fmt.Errorf("no schema available for validation")
	}

	// Parse XML into structured format for detailed validation
	var doc XMLDocument
	if err := xml.Unmarshal(content, &doc); err != nil {
		return fmt.Errorf("XML is not well-formed: %v", err)
	}

	ctx := &ValidationContext{
		errors:   make([]ValidationError, 0),
		warnings: make([]ValidationWarning, 0),
		line:     1,
		column:   1,
	}

	// Find and validate root element
	rootElement := v.findElementDefinition(doc.XMLName.Local)
	if rootElement == nil {
		return fmt.Errorf("root element '%s' not found in schema", doc.XMLName.Local)
	}

	// Perform comprehensive validation
	if err := v.validateElementStructure(doc.XMLName, doc.Attrs, doc.Children, rootElement, ctx); err != nil {
		return err
	}

	// Return validation errors
	if len(ctx.errors) > 0 {
		return fmt.Errorf("validation failed: %s", ctx.errors[0].Message)
	}

	return nil
}

// findElementDefinition finds an element definition in the schema
func (v *XSDValidator) findElementDefinition(elementName string) *types.XSDElement {
	// Check top-level elements
	for _, element := range v.schema.Elements {
		if element.Name == elementName {
			return &element
		}
	}

	// Check elements within complex types
	for _, complexType := range v.schema.ComplexTypes {
		if element := v.findElementInComplexType(&complexType, elementName); element != nil {
			return element
		}
	}

	return nil
}

// findElementInComplexType searches for an element within a complex type
func (v *XSDValidator) findElementInComplexType(complexType *types.XSDComplexType, elementName string) *types.XSDElement {
	// Check sequence elements
	if complexType.Sequence != nil {
		for _, element := range complexType.Sequence.Elements {
			if element.Name == elementName {
				return &element
			}
		}
	}

	// Check choice elements
	if complexType.Choice != nil {
		for _, element := range complexType.Choice.Elements {
			if element.Name == elementName {
				return &element
			}
		}
	}

	// Check all elements
	if complexType.All != nil {
		for _, element := range complexType.All.Elements {
			if element.Name == elementName {
				return &element
			}
		}
	}

	return nil
}

// validateElementStructure validates an element's structure against its XSD definition
func (v *XSDValidator) validateElementStructure(xmlName xml.Name, attrs []xml.Attr, children []XMLElement, elementDef *types.XSDElement, ctx *ValidationContext) error {
	// Validate element type
	if err := v.validateElementType(elementDef, ctx); err != nil {
		return err
	}

	// Validate attributes if element has complex type
	if elementDef.ComplexType != nil {
		if err := v.validateAttributesComplete(attrs, elementDef.ComplexType, ctx); err != nil {
			return err
		}

		// Validate content model (sequence, choice, all)
		if err := v.validateContentModel(children, elementDef.ComplexType, ctx); err != nil {
			return err
		}
	}

	// Validate simple type content if applicable
	if elementDef.SimpleType != nil {
		if err := v.validateSimpleTypeContent(elementDef.SimpleType, ctx); err != nil {
			return err
		}
	}

	// Validate occurrence constraints
	if err := v.validateOccurrenceConstraints(elementDef, ctx); err != nil {
		return err
	}

	return nil
}

// validateElementType validates element type constraints
func (v *XSDValidator) validateElementType(elementDef *types.XSDElement, ctx *ValidationContext) error {
	// Check if element type exists in schema
	if elementDef.Type != "" {
		if !v.isValidType(elementDef.Type) {
			ctx.errors = append(ctx.errors, ValidationError{
				Message: fmt.Sprintf("undefined type '%s' for element '%s'", elementDef.Type, elementDef.Name),
				Line:    ctx.line,
				Column:  ctx.column,
				Element: elementDef.Name,
			})
		}
	}

	return nil
}

// validateAttributesComplete performs comprehensive attribute validation
func (v *XSDValidator) validateAttributesComplete(xmlAttrs []xml.Attr, complexType *types.XSDComplexType, ctx *ValidationContext) error {
	// Create map for efficient lookup
	attrMap := make(map[string]string)
	for _, attr := range xmlAttrs {
		attrMap[attr.Name.Local] = attr.Value
	}

	// Validate all defined attributes
	for _, attrDef := range complexType.Attributes {
		attrName := attrDef.Name
		attrValue, exists := attrMap[attrName]

		// Check required attributes
		if attrDef.Use == "required" && !exists {
			ctx.errors = append(ctx.errors, ValidationError{
				Message: fmt.Sprintf("required attribute '%s' is missing", attrName),
				Line:    ctx.line,
				Column:  ctx.column,
				Element: attrName,
			})
			continue
		}
		// Validate attribute type if present
		if exists {
			if err := v.validateAttributeType(attrValue, attrDef.Type, attrName, ctx); err != nil {
				ctx.errors = append(ctx.errors, ValidationError{
					Message: fmt.Sprintf("invalid value for attribute '%s': %v", attrName, err),
					Line:    ctx.line,
					Column:  ctx.column,
					Element: attrName,
				})
			}
		}

		// Validate default values
		if attrDef.Default != "" && !exists {
			ctx.warnings = append(ctx.warnings, ValidationWarning{
				Message: fmt.Sprintf("attribute '%s' using default value '%s'", attrName, attrDef.Default),
				Line:    ctx.line,
				Column:  ctx.column,
				Element: attrName,
			})
		}
	}

	// Check for unexpected attributes
	for _, xmlAttr := range xmlAttrs {
		found := false
		for _, attrDef := range complexType.Attributes {
			if attrDef.Name == xmlAttr.Name.Local {
				found = true
				break
			}
		}
		if !found {
			ctx.warnings = append(ctx.warnings, ValidationWarning{
				Message: fmt.Sprintf("unexpected attribute '%s'", xmlAttr.Name.Local),
				Line:    ctx.line,
				Column:  ctx.column,
				Element: xmlAttr.Name.Local,
			})
		}
	}

	return nil
}

// validateContentModel validates sequence, choice, or all content models
func (v *XSDValidator) validateContentModel(children []XMLElement, complexType *types.XSDComplexType, ctx *ValidationContext) error {
	if complexType.Sequence != nil {
		return v.validateSequence(children, complexType.Sequence, ctx)
	} else if complexType.Choice != nil {
		return v.validateChoice(children, complexType.Choice, ctx)
	} else if complexType.All != nil {
		return v.validateAll(children, complexType.All, ctx)
	}
	return nil
}

// validateSimpleTypeContent validates simple type content and restrictions
func (v *XSDValidator) validateSimpleTypeContent(simpleType *types.XSDSimpleType, ctx *ValidationContext) error {
	// This would validate restrictions like patterns, enumerations, min/max values, etc.
	// Implementation depends on the specific simple type restrictions

	if simpleType.Restriction != nil {
		return v.validateRestrictions(simpleType.Restriction, ctx)
	}

	return nil
}

// validateRestrictions validates XSD restrictions
func (v *XSDValidator) validateRestrictions(restriction *types.XSDRestriction, ctx *ValidationContext) error {
	// Validate pattern restriction
	if restriction.Pattern != nil {
		if _, err := regexp.Compile(restriction.Pattern.Value); err != nil {
			ctx.errors = append(ctx.errors, ValidationError{
				Message: fmt.Sprintf("invalid regex pattern '%s': %v", restriction.Pattern.Value, err),
				Line:    ctx.line,
				Column:  ctx.column,
			})
		}
	}

	// Validate enumeration restrictions
	if len(restriction.Enumerations) > 0 {
		// Pattern compilation check for enumerations can be added here if needed
	}

	// Validate length restrictions
	if restriction.Length != nil {
		if length, err := strconv.Atoi(restriction.Length.Value); err != nil || length < 0 {
			ctx.errors = append(ctx.errors, ValidationError{
				Message: fmt.Sprintf("invalid length value '%s': must be a non-negative integer", restriction.Length.Value),
				Line:    ctx.line,
				Column:  ctx.column,
			})
		}
	}

	// Validate minLength restrictions
	if restriction.MinLength != nil {
		if minLength, err := strconv.Atoi(restriction.MinLength.Value); err != nil || minLength < 0 {
			ctx.errors = append(ctx.errors, ValidationError{
				Message: fmt.Sprintf("invalid minLength value '%s': must be a non-negative integer", restriction.MinLength.Value),
				Line:    ctx.line,
				Column:  ctx.column,
			})
		}
	}

	// Validate maxLength restrictions
	if restriction.MaxLength != nil {
		if maxLength, err := strconv.Atoi(restriction.MaxLength.Value); err != nil || maxLength < 0 {
			ctx.errors = append(ctx.errors, ValidationError{
				Message: fmt.Sprintf("invalid maxLength value '%s': must be a non-negative integer", restriction.MaxLength.Value),
				Line:    ctx.line,
				Column:  ctx.column,
			})
		}
	}

	// Validate whiteSpace restriction
	if restriction.WhiteSpace != nil {
		validWhiteSpaceValues := []string{"preserve", "replace", "collapse"}
		isValid := false
		for _, valid := range validWhiteSpaceValues {
			if restriction.WhiteSpace.Value == valid {
				isValid = true
				break
			}
		}
		if !isValid {
			ctx.errors = append(ctx.errors, ValidationError{
				Message: fmt.Sprintf("invalid whiteSpace value '%s': must be one of preserve, replace, collapse", restriction.WhiteSpace.Value),
				Line:    ctx.line,
				Column:  ctx.column,
			})
		}
	}

	// Validate numeric range restrictions
	if restriction.MinInclusive != nil {
		if _, err := strconv.ParseFloat(restriction.MinInclusive.Value, 64); err != nil {
			ctx.errors = append(ctx.errors, ValidationError{
				Message: fmt.Sprintf("invalid minInclusive value '%s': must be a valid number", restriction.MinInclusive.Value),
				Line:    ctx.line,
				Column:  ctx.column,
			})
		}
	}

	if restriction.MaxInclusive != nil {
		if _, err := strconv.ParseFloat(restriction.MaxInclusive.Value, 64); err != nil {
			ctx.errors = append(ctx.errors, ValidationError{
				Message: fmt.Sprintf("invalid maxInclusive value '%s': must be a valid number", restriction.MaxInclusive.Value),
				Line:    ctx.line,
				Column:  ctx.column,
			})
		}
	}

	if restriction.MinExclusive != nil {
		if _, err := strconv.ParseFloat(restriction.MinExclusive.Value, 64); err != nil {
			ctx.errors = append(ctx.errors, ValidationError{
				Message: fmt.Sprintf("invalid minExclusive value '%s': must be a valid number", restriction.MinExclusive.Value),
				Line:    ctx.line,
				Column:  ctx.column,
			})
		}
	}

	if restriction.MaxExclusive != nil {
		if _, err := strconv.ParseFloat(restriction.MaxExclusive.Value, 64); err != nil {
			ctx.errors = append(ctx.errors, ValidationError{
				Message: fmt.Sprintf("invalid maxExclusive value '%s': must be a valid number", restriction.MaxExclusive.Value),
				Line:    ctx.line,
				Column:  ctx.column,
			})
		}
	}

	// Validate digit restrictions
	if restriction.TotalDigits != nil {
		if totalDigits, err := strconv.Atoi(restriction.TotalDigits.Value); err != nil || totalDigits <= 0 {
			ctx.errors = append(ctx.errors, ValidationError{
				Message: fmt.Sprintf("invalid totalDigits value '%s': must be a positive integer", restriction.TotalDigits.Value),
				Line:    ctx.line,
				Column:  ctx.column,
			})
		}
	}

	if restriction.FractionDigits != nil {
		if fractionDigits, err := strconv.Atoi(restriction.FractionDigits.Value); err != nil || fractionDigits < 0 {
			ctx.errors = append(ctx.errors, ValidationError{
				Message: fmt.Sprintf("invalid fractionDigits value '%s': must be a non-negative integer", restriction.FractionDigits.Value),
				Line:    ctx.line,
				Column:  ctx.column,
			})
		}
	}

	// Cross-validation between restrictions
	if restriction.MinLength != nil && restriction.MaxLength != nil {
		minLength, _ := strconv.Atoi(restriction.MinLength.Value)
		maxLength, _ := strconv.Atoi(restriction.MaxLength.Value)
		if minLength > maxLength {
			ctx.errors = append(ctx.errors, ValidationError{
				Message: fmt.Sprintf("minLength (%d) cannot be greater than maxLength (%d)", minLength, maxLength),
				Line:    ctx.line,
				Column:  ctx.column,
			})
		}
	}

	if restriction.MinInclusive != nil && restriction.MaxInclusive != nil {
		minValue, _ := strconv.ParseFloat(restriction.MinInclusive.Value, 64)
		maxValue, _ := strconv.ParseFloat(restriction.MaxInclusive.Value, 64)
		if minValue > maxValue {
			ctx.errors = append(ctx.errors, ValidationError{
				Message: fmt.Sprintf("minInclusive (%s) cannot be greater than maxInclusive (%s)", restriction.MinInclusive.Value, restriction.MaxInclusive.Value),
				Line:    ctx.line,
				Column:  ctx.column,
			})
		}
	}

	return nil
}

// validateStringValue validates a string value against XSD restrictions
func (v *XSDValidator) validateStringValue(value string, restriction *types.XSDRestriction, fieldName string, ctx *ValidationContext) error {
	if restriction == nil {
		return nil
	}

	// Validate against pattern restriction
	if restriction.Pattern != nil {
		matched, err := regexp.MatchString(restriction.Pattern.Value, value)
		if err != nil {
			ctx.errors = append(ctx.errors, ValidationError{
				Message: fmt.Sprintf("invalid pattern '%s' for field '%s'", restriction.Pattern.Value, fieldName),
				Line:    ctx.line,
				Column:  ctx.column,
				Element: fieldName,
			})
		} else if !matched {
			ctx.errors = append(ctx.errors, ValidationError{
				Message: fmt.Sprintf("value '%s' does not match pattern '%s' for field '%s'", value, restriction.Pattern.Value, fieldName),
				Line:    ctx.line,
				Column:  ctx.column,
				Element: fieldName,
			})
		}
	}

	// Validate against enumeration restriction
	if len(restriction.Enumerations) > 0 {
		found := false
		for _, enum := range restriction.Enumerations {
			if enum.Value == value {
				found = true
				break
			}
		}
		if !found {
			validValues := make([]string, len(restriction.Enumerations))
			for i, enum := range restriction.Enumerations {
				validValues[i] = enum.Value
			}
			ctx.errors = append(ctx.errors, ValidationError{
				Message: fmt.Sprintf("value '%s' is not in enumeration [%s] for field '%s'", value, strings.Join(validValues, ", "), fieldName),
				Line:    ctx.line,
				Column:  ctx.column,
				Element: fieldName,
			})
		}
	}

	// Apply whiteSpace processing before length validation
	processedValue := v.applyWhiteSpaceProcessing(value, restriction)

	// Validate exact length restriction
	if restriction.Length != nil {
		expectedLength, _ := strconv.Atoi(restriction.Length.Value)
		if len(processedValue) != expectedLength {
			ctx.errors = append(ctx.errors, ValidationError{
				Message: fmt.Sprintf("value '%s' length (%d) does not match required length (%d) for field '%s'", value, len(processedValue), expectedLength, fieldName),
				Line:    ctx.line,
				Column:  ctx.column,
				Element: fieldName,
			})
		}
	}

	// Validate minLength restriction
	if restriction.MinLength != nil {
		minLength, _ := strconv.Atoi(restriction.MinLength.Value)
		if len(processedValue) < minLength {
			ctx.errors = append(ctx.errors, ValidationError{
				Message: fmt.Sprintf("value '%s' length (%d) is less than minimum length (%d) for field '%s'", value, len(processedValue), minLength, fieldName),
				Line:    ctx.line,
				Column:  ctx.column,
				Element: fieldName,
			})
		}
	}

	// Validate maxLength restriction
	if restriction.MaxLength != nil {
		maxLength, _ := strconv.Atoi(restriction.MaxLength.Value)
		if len(processedValue) > maxLength {
			ctx.errors = append(ctx.errors, ValidationError{
				Message: fmt.Sprintf("value '%s' length (%d) exceeds maximum length (%d) for field '%s'", value, len(processedValue), maxLength, fieldName),
				Line:    ctx.line,
				Column:  ctx.column,
				Element: fieldName,
			})
		}
	}

	return nil
}

// validateNumericValue validates a numeric value against XSD restrictions
func (v *XSDValidator) validateNumericValue(value string, restriction *types.XSDRestriction, fieldName string, ctx *ValidationContext) error {
	if restriction == nil {
		return nil
	}

	numValue, err := strconv.ParseFloat(value, 64)
	if err != nil {
		ctx.errors = append(ctx.errors, ValidationError{
			Message: fmt.Sprintf("value '%s' is not a valid number for field '%s'", value, fieldName),
			Line:    ctx.line,
			Column:  ctx.column,
			Element: fieldName,
		})
		return nil
	}

	// Validate minInclusive restriction
	if restriction.MinInclusive != nil {
		minValue, _ := strconv.ParseFloat(restriction.MinInclusive.Value, 64)
		if numValue < minValue {
			ctx.errors = append(ctx.errors, ValidationError{
				Message: fmt.Sprintf("value '%s' (%f) is less than minimum inclusive value (%f) for field '%s'", value, numValue, minValue, fieldName),
				Line:    ctx.line,
				Column:  ctx.column,
				Element: fieldName,
			})
		}
	}

	// Validate maxInclusive restriction
	if restriction.MaxInclusive != nil {
		maxValue, _ := strconv.ParseFloat(restriction.MaxInclusive.Value, 64)
		if numValue > maxValue {
			ctx.errors = append(ctx.errors, ValidationError{
				Message: fmt.Sprintf("value '%s' (%f) exceeds maximum inclusive value (%f) for field '%s'", value, numValue, maxValue, fieldName),
				Line:    ctx.line,
				Column:  ctx.column,
				Element: fieldName,
			})
		}
	}

	// Validate minExclusive restriction
	if restriction.MinExclusive != nil {
		minValue, _ := strconv.ParseFloat(restriction.MinExclusive.Value, 64)
		if numValue <= minValue {
			ctx.errors = append(ctx.errors, ValidationError{
				Message: fmt.Sprintf("value '%s' (%f) is not greater than minimum exclusive value (%f) for field '%s'", value, numValue, minValue, fieldName),
				Line:    ctx.line,
				Column:  ctx.column,
				Element: fieldName,
			})
		}
	}

	// Validate maxExclusive restriction
	if restriction.MaxExclusive != nil {
		maxValue, _ := strconv.ParseFloat(restriction.MaxExclusive.Value, 64)
		if numValue >= maxValue {
			ctx.errors = append(ctx.errors, ValidationError{
				Message: fmt.Sprintf("value '%s' (%f) is not less than maximum exclusive value (%f) for field '%s'", value, numValue, maxValue, fieldName),
				Line:    ctx.line,
				Column:  ctx.column,
				Element: fieldName,
			})
		}
	}

	// Validate totalDigits restriction
	if restriction.TotalDigits != nil {
		totalDigits, _ := strconv.Atoi(restriction.TotalDigits.Value)
		digitStr := strings.ReplaceAll(strings.ReplaceAll(value, ".", ""), "-", "")
		if len(digitStr) > totalDigits {
			ctx.errors = append(ctx.errors, ValidationError{
				Message: fmt.Sprintf("value '%s' has %d total digits, maximum allowed is %d for field '%s'", value, len(digitStr), totalDigits, fieldName),
				Line:    ctx.line,
				Column:  ctx.column,
				Element: fieldName,
			})
		}
	}

	// Validate fractionDigits restriction
	if restriction.FractionDigits != nil {
		fractionDigits, _ := strconv.Atoi(restriction.FractionDigits.Value)
		if dotIndex := strings.Index(value, "."); dotIndex != -1 {
			actualFractionDigits := len(value) - dotIndex - 1
			if actualFractionDigits > fractionDigits {
				ctx.errors = append(ctx.errors, ValidationError{
					Message: fmt.Sprintf("value '%s' has %d fraction digits, maximum allowed is %d for field '%s'", value, actualFractionDigits, fractionDigits, fieldName),
					Line:    ctx.line,
					Column:  ctx.column,
					Element: fieldName,
				})
			}
		}
	}

	return nil
}

// validateFixedValue validates that a value matches a fixed value constraint
func (v *XSDValidator) validateFixedValue(value, fixedValue, fieldName string, ctx *ValidationContext) error {
	if fixedValue != "" && value != fixedValue {
		ctx.errors = append(ctx.errors, ValidationError{
			Message: fmt.Sprintf("value '%s' does not match fixed value '%s' for field '%s'", value, fixedValue, fieldName),
			Line:    ctx.line,
			Column:  ctx.column,
			Element: fieldName,
		})
	}
	return nil
}

// applyWhiteSpaceProcessing applies whiteSpace facet processing to a value
func (v *XSDValidator) applyWhiteSpaceProcessing(value string, restriction *types.XSDRestriction) string {
	if restriction == nil || restriction.WhiteSpace == nil {
		return value
	}

	switch restriction.WhiteSpace.Value {
	case "replace":
		// Replace all occurrences of #x9 (tab), #xA (line feed) and #xD (carriage return) with a single space
		value = strings.ReplaceAll(value, "\t", " ")
		value = strings.ReplaceAll(value, "\n", " ")
		value = strings.ReplaceAll(value, "\r", " ")
		return value
	case "collapse":
		// First apply replace processing
		value = strings.ReplaceAll(value, "\t", " ")
		value = strings.ReplaceAll(value, "\n", " ")
		value = strings.ReplaceAll(value, "\r", " ")
		// Then collapse sequences of spaces to a single space and trim leading/trailing spaces
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

// GenerateValidationReport generates a detailed validation report
func (v *XSDValidator) GenerateValidationReport(xmlPath string) (*ValidationReport, error) {
	report := &ValidationReport{
		XMLPath:  xmlPath,
		IsValid:  true,
		Errors:   make([]ValidationError, 0),
		Warnings: make([]ValidationWarning, 0),
	}

	// Read XML file
	content, err := ioutil.ReadFile(xmlPath)
	if err != nil {
		report.IsValid = false
		report.Errors = append(report.Errors, ValidationError{
			Message: fmt.Sprintf("failed to read XML file: %v", err),
			Line:    0,
			Column:  0,
		})
		return report, nil
	}

	// Parse XML
	var doc XMLDocument
	if err := xml.Unmarshal(content, &doc); err != nil {
		report.IsValid = false
		report.Errors = append(report.Errors, ValidationError{
			Message: fmt.Sprintf("failed to parse XML: %v", err),
			Line:    0,
			Column:  0,
		})
		return report, nil
	}

	// Validate against schema
	if v.schema == nil {
		report.IsValid = false
		report.Errors = append(report.Errors, ValidationError{
			Message: "no schema available for validation",
			Line:    0,
			Column:  0,
		})
		return report, nil
	}

	ctx := &ValidationContext{
		errors:   make([]ValidationError, 0),
		warnings: make([]ValidationWarning, 0),
		line:     1,
		column:   1,
	}

	// Find and validate root element
	rootElement := v.findRootElement(doc.XMLName)
	if rootElement == nil {
		report.IsValid = false
		report.Errors = append(report.Errors, ValidationError{
			Message: fmt.Sprintf("root element '%s' not found in schema", doc.XMLName.Local),
			Line:    1,
			Column:  1,
			Element: doc.XMLName.Local,
		})
		return report, nil
	}

	// Perform detailed validation
	v.validateElement(doc.XMLName, doc.Attrs, doc.Children, rootElement, ctx)

	// Copy validation results to report
	report.Errors = ctx.errors
	report.Warnings = ctx.warnings
	report.IsValid = len(ctx.errors) == 0

	return report, nil
}

// ValidationReport represents a validation report
type ValidationReport struct {
	XMLPath  string
	IsValid  bool
	Errors   []ValidationError
	Warnings []ValidationWarning
}

// ValidationError represents a validation error
type ValidationError struct {
	Message string
	Line    int
	Column  int
	Element string
}

// ValidationWarning represents a validation warning
type ValidationWarning struct {
	Message string
	Line    int
	Column  int
	Element string
}

// String returns a string representation of the validation report
func (r *ValidationReport) String() string {
	var builder strings.Builder

	builder.WriteString(fmt.Sprintf("Validation Report for: %s\n", r.XMLPath))
	builder.WriteString(fmt.Sprintf("Valid: %t\n", r.IsValid))

	if len(r.Errors) > 0 {
		builder.WriteString("\nErrors:\n")
		for i, err := range r.Errors {
			builder.WriteString(fmt.Sprintf("  %d. %s", i+1, err.Message))
			if err.Line > 0 {
				builder.WriteString(fmt.Sprintf(" (Line: %d, Column: %d)", err.Line, err.Column))
			}
			if err.Element != "" {
				builder.WriteString(fmt.Sprintf(" [Element: %s]", err.Element))
			}
			builder.WriteString("\n")
		}
	}

	if len(r.Warnings) > 0 {
		builder.WriteString("\nWarnings:\n")
		for i, warn := range r.Warnings {
			builder.WriteString(fmt.Sprintf("  %d. %s", i+1, warn.Message))
			if warn.Line > 0 {
				builder.WriteString(fmt.Sprintf(" (Line: %d, Column: %d)", warn.Line, warn.Column))
			}
			if warn.Element != "" {
				builder.WriteString(fmt.Sprintf(" [Element: %s]", warn.Element))
			}
			builder.WriteString("\n")
		}
	}

	return builder.String()
}

// validateOccurrenceConstraints validates minOccurs and maxOccurs constraints
func (v *XSDValidator) validateOccurrenceConstraints(elementDef *types.XSDElement, ctx *ValidationContext) error {
	min, max := types.ParseOccurs(elementDef.MinOccurs, elementDef.MaxOccurs)

	// Basic sanity checks
	if min < 0 {
		ctx.errors = append(ctx.errors, ValidationError{
			Message: fmt.Sprintf("invalid minOccurs value %d for element '%s'", min, elementDef.Name),
			Line:    ctx.line,
			Column:  ctx.column,
			Element: elementDef.Name,
		})
	}

	if max != -1 && max < min {
		ctx.errors = append(ctx.errors, ValidationError{
			Message: fmt.Sprintf("maxOccurs (%d) is less than minOccurs (%d) for element '%s'", max, min, elementDef.Name),
			Line:    ctx.line,
			Column:  ctx.column,
			Element: elementDef.Name,
		})
	}

	return nil
}

// isValidType checks if a type is valid (built-in or defined in schema)
func (v *XSDValidator) isValidType(typeName string) bool {
	// Check built-in types
	builtinTypes := []string{
		"string", "int", "integer", "decimal", "float", "double", "boolean",
		"date", "dateTime", "time", "duration", "base64Binary", "hexBinary",
		"anyURI", "QName", "NOTATION", "normalizedString", "token", "language",
		"NMTOKEN", "NMTOKENS", "Name", "NCName", "ID", "IDREF", "IDREFS", "ENTITY", "ENTITIES",
		"byte", "unsignedByte", "short", "unsignedShort", "long", "unsignedLong",
		"positiveInteger", "nonPositiveInteger", "negativeInteger", "nonNegativeInteger",
		"gYear", "gYearMonth", "gMonth", "gMonthDay", "gDay",
	}

	for _, builtin := range builtinTypes {
		if typeName == builtin || typeName == "xs:"+builtin || typeName == "xsd:"+builtin {
			return true
		}
	}

	// Check schema-defined types
	for _, complexType := range v.schema.ComplexTypes {
		if complexType.Name == typeName {
			return true
		}
	}

	for _, simpleType := range v.schema.SimpleTypes {
		if simpleType.Name == typeName {
			return true
		}
	}

	return false
}
