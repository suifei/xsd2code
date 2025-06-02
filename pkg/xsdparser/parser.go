package xsdparser

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"path/filepath"
	"strings"

	"github.com/suifei/xsd2code/pkg/generator"
	"github.com/suifei/xsd2code/pkg/types"
)

// XSDParser is the main parser for XSD files
type XSDParser struct {
	schema          *types.XSDSchema
	filePath        string
	outputPath      string
	packageName     string
	targetNamespace string
	imports         map[string]*types.XSDSchema
	goTypes         []types.GoType
	debugMode       bool
	strictMode      bool
	jsonCompatible  bool
	includeComments bool
}

// NewXSDParser creates a new XSD parser instance
func NewXSDParser(xsdPath, outputPath, packageName string) *XSDParser {
	return &XSDParser{
		filePath:        xsdPath,
		outputPath:      outputPath,
		packageName:     packageName,
		imports:         make(map[string]*types.XSDSchema),
		goTypes:         make([]types.GoType, 0),
		debugMode:       false,
		strictMode:      false,
		jsonCompatible:  false,
		includeComments: true,
	}
}

// SetDebugMode enables or disables debug mode
func (p *XSDParser) SetDebugMode(debug bool) {
	p.debugMode = debug
}

// SetStrictMode enables or disables strict mode
func (p *XSDParser) SetStrictMode(strict bool) {
	p.strictMode = strict
}

// SetJSONCompatible enables or disables JSON compatibility
func (p *XSDParser) SetJSONCompatible(json bool) {
	p.jsonCompatible = json
}

// SetIncludeComments enables or disables comments in generated code
func (p *XSDParser) SetIncludeComments(comments bool) {
	p.includeComments = comments
}

// Parse parses the XSD file and builds the internal representation
func (p *XSDParser) Parse() error {
	if p.debugMode {
		fmt.Printf("Parsing XSD file: %s\n", p.filePath)
	}

	// Read and parse the main XSD file
	if err := p.parseFile(p.filePath); err != nil {
		return fmt.Errorf("failed to parse main XSD file: %v", err)
	}

	// Process imports and includes
	if err := p.processImports(); err != nil {
		return fmt.Errorf("failed to process imports: %v", err)
	}

	// Convert XSD types to Go types
	if err := p.convertToGoTypes(); err != nil {
		return fmt.Errorf("failed to convert to Go types: %v", err)
	}

	return nil
}

// parseFile parses a single XSD file
func (p *XSDParser) parseFile(filePath string) error {
	// Read file content
	content, err := ioutil.ReadFile(filePath)
	if err != nil {
		return fmt.Errorf("failed to read file %s: %v", filePath, err)
	}

	// Parse XML
	var schema types.XSDSchema
	if err := xml.Unmarshal(content, &schema); err != nil {
		return fmt.Errorf("failed to parse XML in %s: %v", filePath, err)
	}

	// Parse namespaces from XML content
	if err := p.parseNamespaces(content, &schema); err != nil {
		return fmt.Errorf("failed to parse namespaces: %v", err)
	}

	p.schema = &schema
	p.targetNamespace = schema.TargetNamespace

	if p.debugMode {
		fmt.Printf("Parsed schema with target namespace: %s\n", p.targetNamespace)
		fmt.Printf("Found %d elements, %d complex types, %d simple types\n",
			len(schema.Elements), len(schema.ComplexTypes), len(schema.SimpleTypes))
	}

	return nil
}

// parseNamespaces extracts namespace declarations from XML content
func (p *XSDParser) parseNamespaces(content []byte, schema *types.XSDSchema) error {
	// Simple namespace parsing - could be improved with proper XML parsing
	lines := strings.Split(string(content), "\n")
	for _, line := range lines {
		if strings.Contains(line, "xmlns") {
			// Extract namespace declarations
			// This is a simplified approach
			break
		}
	}
	return nil
}

// processImports processes import and include statements
func (p *XSDParser) processImports() error {
	if p.schema == nil {
		return fmt.Errorf("schema not parsed")
	}

	// Process imports
	for _, imp := range p.schema.Imports {
		if err := p.processImport(imp); err != nil {
			if p.strictMode {
				return err
			}
			if p.debugMode {
				fmt.Printf("Warning: failed to process import %s: %v\n", imp.SchemaLocation, err)
			}
		}
	}

	// Process includes
	for _, inc := range p.schema.Includes {
		if err := p.processInclude(inc); err != nil {
			if p.strictMode {
				return err
			}
			if p.debugMode {
				fmt.Printf("Warning: failed to process include %s: %v\n", inc.SchemaLocation, err)
			}
		}
	}

	return nil
}

// processImport processes a single import
func (p *XSDParser) processImport(imp types.XSDImport) error {
	if imp.SchemaLocation == "" {
		return nil // Skip imports without schema location
	}

	// Resolve relative path
	importPath := imp.SchemaLocation
	if !filepath.IsAbs(importPath) {
		baseDir := filepath.Dir(p.filePath)
		importPath = filepath.Join(baseDir, importPath)
	}

	// Check if already imported
	if _, exists := p.imports[imp.Namespace]; exists {
		return nil
	}

	if p.debugMode {
		fmt.Printf("Processing import: %s -> %s\n", imp.Namespace, importPath)
	}

	// Parse imported schema
	content, err := ioutil.ReadFile(importPath)
	if err != nil {
		return fmt.Errorf("failed to read import file %s: %v", importPath, err)
	}

	var importedSchema types.XSDSchema
	if err := xml.Unmarshal(content, &importedSchema); err != nil {
		return fmt.Errorf("failed to parse imported XML %s: %v", importPath, err)
	}

	p.imports[imp.Namespace] = &importedSchema
	return nil
}

// processInclude processes a single include
func (p *XSDParser) processInclude(inc types.XSDInclude) error {
	if inc.SchemaLocation == "" {
		return nil
	}

	// Resolve relative path
	includePath := inc.SchemaLocation
	if !filepath.IsAbs(includePath) {
		baseDir := filepath.Dir(p.filePath)
		includePath = filepath.Join(baseDir, includePath)
	}

	if p.debugMode {
		fmt.Printf("Processing include: %s\n", includePath)
	}

	// Parse included schema
	content, err := ioutil.ReadFile(includePath)
	if err != nil {
		return fmt.Errorf("failed to read include file %s: %v", includePath, err)
	}

	var includedSchema types.XSDSchema
	if err := xml.Unmarshal(content, &includedSchema); err != nil {
		return fmt.Errorf("failed to parse included XML %s: %v", includePath, err)
	}

	// Merge included schema into main schema
	p.mergeSchema(&includedSchema)
	return nil
}

// mergeSchema merges an included schema into the main schema
func (p *XSDParser) mergeSchema(included *types.XSDSchema) {
	if p.schema == nil {
		return
	}

	// Merge elements
	p.schema.Elements = append(p.schema.Elements, included.Elements...)

	// Merge complex types
	p.schema.ComplexTypes = append(p.schema.ComplexTypes, included.ComplexTypes...)

	// Merge simple types
	p.schema.SimpleTypes = append(p.schema.SimpleTypes, included.SimpleTypes...)

	// Merge groups
	p.schema.Groups = append(p.schema.Groups, included.Groups...)

	// Merge attribute groups
	p.schema.AttributeGroups = append(p.schema.AttributeGroups, included.AttributeGroups...)
}

// convertToGoTypes converts XSD types to Go types
func (p *XSDParser) convertToGoTypes() error {
	if p.debugMode {
		fmt.Printf("Processing schema with %d groups\n", len(p.schema.Groups))
		for _, group := range p.schema.Groups {
			fmt.Printf("Found group: %s\n", group.Name)
		}
	}
	if p.schema == nil {
		return fmt.Errorf("schema not parsed")
	}

	// Convert complex types
	for _, complexType := range p.schema.ComplexTypes {
		goType, err := p.convertComplexType(complexType)
		if err != nil {
			return fmt.Errorf("failed to convert complex type %s: %v", complexType.Name, err)
		}
		p.goTypes = append(p.goTypes, *goType)
	}
	// Convert simple types (enums, etc.)
	for _, simpleType := range p.schema.SimpleTypes {
		fmt.Printf("Converting simple type: %s\n", simpleType.Name)
		if simpleType.Restriction != nil && simpleType.Restriction.Pattern != nil {
			fmt.Printf("  Has pattern: %s\n", simpleType.Restriction.Pattern.Value)
		}
		goType, err := p.convertSimpleType(simpleType)
		if err != nil {
			return fmt.Errorf("failed to convert simple type %s: %v", simpleType.Name, err)
		}
		if goType != nil {
			fmt.Printf("  Added type: %s (pattern: %v)\n", goType.Name, goType.HasPattern)
			p.goTypes = append(p.goTypes, *goType)
		} else {
			fmt.Printf("  Type was nil\n")
		}
	}

	// Convert root elements
	for _, element := range p.schema.Elements {
		if element.ComplexType != nil {
			goType, err := p.convertComplexTypeFromElement(element)
			if err != nil {
				return fmt.Errorf("failed to convert element %s: %v", element.Name, err)
			}
			p.goTypes = append(p.goTypes, *goType)
		}
	}

	if p.debugMode {
		fmt.Printf("Converted %d Go types\n", len(p.goTypes))
	}

	return nil
}

// convertComplexType converts an XSD complex type to a Go type
func (p *XSDParser) convertComplexType(xsdType types.XSDComplexType) (*types.GoType, error) {
	goType := &types.GoType{
		Name:      types.ToGoTypeName(xsdType.Name),
		Package:   p.packageName,
		XMLName:   xsdType.Name,
		Namespace: p.targetNamespace,
		Fields:    make([]types.GoField, 0),
		Comment:   types.GetDocumentation(xsdType.Annotation),
	}

	// Handle different content models
	if xsdType.Sequence != nil {
		if err := p.processSequence(xsdType.Sequence, goType); err != nil {
			return nil, err
		}
	}

	if xsdType.Choice != nil {
		if err := p.processChoice(xsdType.Choice, goType); err != nil {
			return nil, err
		}
	}

	if xsdType.All != nil {
		if err := p.processAll(xsdType.All, goType); err != nil {
			return nil, err
		}
	}

	// Handle attributes
	for _, attr := range xsdType.Attributes {
		field, err := p.convertAttribute(attr)
		if err != nil {
			return nil, err
		}
		goType.Fields = append(goType.Fields, *field)
	}

	// Handle extensions
	if xsdType.ComplexContent != nil && xsdType.ComplexContent.Extension != nil {
		if err := p.processExtension(xsdType.ComplexContent.Extension, goType); err != nil {
			return nil, err
		}
	}

	if xsdType.SimpleContent != nil && xsdType.SimpleContent.Extension != nil {
		if err := p.processExtension(xsdType.SimpleContent.Extension, goType); err != nil {
			return nil, err
		}
	}

	return goType, nil
}

// convertSimpleType converts an XSD simple type to a Go type
func (p *XSDParser) convertSimpleType(xsdType types.XSDSimpleType) (*types.GoType, error) {
	if xsdType.Restriction == nil {
		return nil, nil // Skip non-restriction simple types
	}
	var goType *types.GoType
	// Handle enumerations, which have precedence over other restrictions
	if len(xsdType.Restriction.Enumerations) > 0 {
		return p.convertEnumType(xsdType)
	}

	// Create a base type for restrictions
	baseType := p.mapXSDTypeToGo(xsdType.Restriction.Base)
	goType = &types.GoType{
		Name:            types.ToGoTypeName(xsdType.Name),
		Package:         p.packageName,
		BaseType:        baseType,
		IsEnum:          false,
		Comment:         types.GetDocumentation(xsdType.Annotation),
		NeedsValidation: false, // Will be set to true if any restrictions are found
	}

	// Handle pattern restriction
	if xsdType.Restriction.Pattern != nil {
		goType.HasPattern = true
		goType.PatternValue = xsdType.Restriction.Pattern.Value
		goType.NeedsValidation = true
	}

	// Handle length restrictions
	if xsdType.Restriction.MinLength != nil {
		goType.HasMinLength = true
		goType.MinLength = xsdType.Restriction.MinLength.Value
		goType.NeedsValidation = true
	}
	if xsdType.Restriction.MaxLength != nil {
		goType.HasMaxLength = true
		goType.MaxLength = xsdType.Restriction.MaxLength.Value
		goType.NeedsValidation = true
	}

	// Handle numeric restrictions
	if xsdType.Restriction.MinInclusive != nil {
		goType.HasMinInclusive = true
		goType.MinInclusive = xsdType.Restriction.MinInclusive.Value
		goType.NeedsValidation = true
	}
	if xsdType.Restriction.MaxInclusive != nil {
		goType.HasMaxInclusive = true
		goType.MaxInclusive = xsdType.Restriction.MaxInclusive.Value
		goType.NeedsValidation = true
	}
	if xsdType.Restriction.MinExclusive != nil {
		goType.HasMinExclusive = true
		goType.MinExclusive = xsdType.Restriction.MinExclusive.Value
		goType.NeedsValidation = true
	}
	if xsdType.Restriction.MaxExclusive != nil {
		goType.HasMaxExclusive = true
		goType.MaxExclusive = xsdType.Restriction.MaxExclusive.Value
		goType.NeedsValidation = true
	}

	// Handle digit restrictions
	if xsdType.Restriction.TotalDigits != nil {
		goType.HasTotalDigits = true
		goType.TotalDigits = xsdType.Restriction.TotalDigits.Value
		goType.NeedsValidation = true
	}
	if xsdType.Restriction.FractionDigits != nil {
		goType.HasFractionDigits = true
		goType.FractionDigits = xsdType.Restriction.FractionDigits.Value
		goType.NeedsValidation = true
	}
	// Even if no validation is explicitly needed, we still return the type
	// so it can be used in references
	return goType, nil
}

// convertEnumType converts an XSD enumeration to a Go type with constants
func (p *XSDParser) convertEnumType(xsdType types.XSDSimpleType) (*types.GoType, error) {
	baseType := p.mapXSDTypeToGo(xsdType.Restriction.Base)

	goType := &types.GoType{
		Name:      types.ToGoTypeName(xsdType.Name),
		Package:   p.packageName,
		BaseType:  baseType,
		IsEnum:    true,
		Constants: make([]types.GoConstant, 0),
		Comment:   types.GetDocumentation(xsdType.Annotation),
	}

	// Convert enumerations to constants
	for _, enum := range xsdType.Restriction.Enumerations {
		constName := p.generateConstantName(goType.Name, enum.Value)
		constant := types.GoConstant{
			Name:    constName,
			Value:   fmt.Sprintf("\"%s\"", enum.Value),
			Comment: types.GetDocumentation(enum.Annotation),
		}
		goType.Constants = append(goType.Constants, constant)
	}
	return goType, nil
}

// convertComplexTypeFromElement converts an inline complex type from an element
func (p *XSDParser) convertComplexTypeFromElement(element types.XSDElement) (*types.GoType, error) {
	if element.ComplexType == nil {
		return nil, fmt.Errorf("element has no complex type")
	}

	// Use element name as type name
	typeName := types.ToGoTypeName(element.Name)
	complexType := *element.ComplexType
	complexType.Name = typeName

	return p.convertComplexType(complexType)
}

// processSequence processes an XSD sequence
func (p *XSDParser) processSequence(sequence *types.XSDSequence, goType *types.GoType) error {
	return p.processSequenceWithContext(sequence, goType, []string{goType.Name})
}

// processSequenceWithContext processes an XSD sequence with context path
func (p *XSDParser) processSequenceWithContext(sequence *types.XSDSequence, goType *types.GoType, contextPath []string) error {
	for _, element := range sequence.Elements {
		field, err := p.convertElementWithContext(element, contextPath)
		if err != nil {
			return err
		}
		goType.Fields = append(goType.Fields, *field)
	}

	// Process nested sequences
	for _, nestedSeq := range sequence.Sequences {
		if err := p.processSequenceWithContext(&nestedSeq, goType, contextPath); err != nil {
			return err
		}
	}
	// Process choices within sequence
	for _, choice := range sequence.Choices {
		if err := p.processChoiceWithContext(&choice, goType, contextPath); err != nil {
			return err
		}
	}

	// Process group references within sequence
	for _, groupRef := range sequence.Groups {
		if err := p.processGroupRef(groupRef, goType, contextPath); err != nil {
			return err
		}
	}

	return nil
}

// processChoice processes an XSD choice
func (p *XSDParser) processChoice(choice *types.XSDChoice, goType *types.GoType) error {
	return p.processChoiceWithContext(choice, goType, []string{goType.Name})
}

// isGoBasicType checks if a type is a Go basic type that doesn't need pointer wrapping in choices
func (p *XSDParser) isGoBasicType(goType string) bool {
	basicTypes := map[string]bool{
		"bool":    true,
		"int8":    true,
		"int16":   true,
		"int32":   true,
		"int64":   true,
		"uint8":   true,
		"uint16":  true,
		"uint32":  true,
		"uint64":  true,
		"float32": true,
		"float64": true,
		"string":  true,
	}
	return basicTypes[goType]
}

// processChoiceWithContext processes an XSD choice with context path
func (p *XSDParser) processChoiceWithContext(choice *types.XSDChoice, goType *types.GoType, contextPath []string) error {
	// For choices, we create optional pointer fields for each choice option
	for _, element := range choice.Elements {
		field, err := p.convertElementWithContext(element, contextPath)
		if err != nil {
			return err
		} // For choice elements, all fields must be optional with omitempty
		if !strings.HasPrefix(field.Type, "*") && !strings.HasPrefix(field.Type, "[]") {
			// Check if this is an empty element (no type definition)
			if element.ComplexType == nil && element.SimpleType == nil && element.Type == "" {
				// For empty elements, try to map the element name as a type
				// This handles elementary types like BOOL, BYTE, WORD, etc.
				mappedType := p.mapXSDTypeToGo(element.Name)
				if mappedType != types.ToGoTypeName(element.Name) {
					// Element name was successfully mapped to a built-in type
					// Check if it's a basic type that doesn't need pointer wrapping
					if p.isGoBasicType(mappedType) {
						field.Type = mappedType
					} else {
						field.Type = "*" + mappedType
					}
				} else {
					// Empty element with no known mapping - use *struct{}
					field.Type = "*struct{}"
				}
			} else {
				// For non-empty elements, check if the mapped type is basic
				if p.isGoBasicType(field.Type) {
					// Basic types don't need pointer wrapping in choices
					// Keep as is
				} else {
					// Make it a pointer for complex types
					field.Type = "*" + field.Type
				}
			}
		}

		// Ensure omitempty is in both XML and JSON tags
		if field.XMLTag != "" && !strings.Contains(field.XMLTag, "omitempty") {
			if strings.Contains(field.XMLTag, ",attr") {
				field.XMLTag = strings.Replace(field.XMLTag, ",attr", ",omitempty", 1)
			} else {
				field.XMLTag += ",omitempty"
			}
		}
		if field.JSONTag != "" && !strings.Contains(field.JSONTag, "omitempty") {
			field.JSONTag += ",omitempty"
		}

		field.IsOptional = true // All choice elements are optional
		goType.Fields = append(goType.Fields, *field)
	}

	// Process nested choices
	for _, nestedChoice := range choice.Choices {
		if err := p.processChoiceWithContext(&nestedChoice, goType, contextPath); err != nil {
			return err
		}
	}
	// Process sequences within choice
	for _, sequence := range choice.Sequences {
		if err := p.processSequenceWithContext(&sequence, goType, contextPath); err != nil {
			return err
		}
	}

	// Process group references within choice
	for _, groupRef := range choice.Groups {
		if err := p.processGroupRef(groupRef, goType, contextPath); err != nil {
			return err
		}
	}

	return nil
}

// processAll processes an XSD all
func (p *XSDParser) processAll(all *types.XSDAll, goType *types.GoType) error {
	return p.processAllWithContext(all, goType, []string{goType.Name})
}

// processAllWithContext processes an XSD all with context path
func (p *XSDParser) processAllWithContext(all *types.XSDAll, goType *types.GoType, contextPath []string) error {
	for _, element := range all.Elements {
		field, err := p.convertElementWithContext(element, contextPath)
		if err != nil {
			return err
		}
		goType.Fields = append(goType.Fields, *field)
	}
	return nil
}

// processExtension processes an XSD extension
func (p *XSDParser) processExtension(extension *types.XSDExtension, goType *types.GoType) error {
	// Add base type comment
	if extension.Base != "" {
		baseComment := fmt.Sprintf("extends %s", extension.Base)
		if goType.Comment != "" {
			goType.Comment += "; " + baseComment
		} else {
			goType.Comment = baseComment
		}
	}

	// Process sequence
	if extension.Sequence != nil {
		if err := p.processSequence(extension.Sequence, goType); err != nil {
			return err
		}
	}

	// Process choice
	if extension.Choice != nil {
		if err := p.processChoice(extension.Choice, goType); err != nil {
			return err
		}
	}

	// Process all
	if extension.All != nil {
		if err := p.processAll(extension.All, goType); err != nil {
			return err
		}
	}

	// Process attributes
	for _, attr := range extension.Attributes {
		field, err := p.convertAttribute(attr)
		if err != nil {
			return err
		}
		goType.Fields = append(goType.Fields, *field)
	}

	return nil
}

// convertElement converts an XSD element to a Go field
func (p *XSDParser) convertElement(element types.XSDElement) (*types.GoField, error) {
	return p.convertElementWithContext(element, []string{})
}

// convertElementWithContext converts an XSD element to a Go field with context path
func (p *XSDParser) convertElementWithContext(element types.XSDElement, contextPath []string) (*types.GoField, error) {
	fieldName := types.ToGoFieldName(element.Name)
	fieldType := p.mapXSDTypeToGo(element.Type)

	min, max := types.ParseOccurs(element.MinOccurs, element.MaxOccurs)
	isOptional := min == 0
	isArray := max > 1 || max == -1 // Handle inline complex type
	if element.ComplexType != nil {
		// Create context-aware inline type name
		fullContextPath := append(contextPath, types.ToGoTypeName(element.Name))
		fieldType = strings.Join(fullContextPath, "")

		// Check if this type already exists to avoid duplicates
		existingType := p.findExistingType(fieldType)
		if existingType == nil {
			// Create and add inline type to goTypes list
			inlineType := types.GoType{
				Name:    fieldType,
				Comment: fmt.Sprintf("%s represents the inline complex type for element %s", fieldType, element.Name),
				Fields:  make([]types.GoField, 0),
				XMLName: element.Name,
			} // Process content model using proper context-aware methods that handle group references
			if element.ComplexType.Sequence != nil {
				if err := p.processSequenceWithContext(element.ComplexType.Sequence, &inlineType, fullContextPath); err != nil {
					return nil, fmt.Errorf("failed to process sequence in inline type %s: %v", fieldType, err)
				}
			} else if element.ComplexType.Choice != nil {
				if err := p.processChoiceWithContext(element.ComplexType.Choice, &inlineType, fullContextPath); err != nil {
					return nil, fmt.Errorf("failed to process choice in inline type %s: %v", fieldType, err)
				}
			} else if element.ComplexType.All != nil {
				if err := p.processAllWithContext(element.ComplexType.All, &inlineType, fullContextPath); err != nil {
					return nil, fmt.Errorf("failed to process all in inline type %s: %v", fieldType, err)
				}
			}

			// Convert attributes to fields
			for _, attr := range element.ComplexType.Attributes {
				field, err := p.convertAttribute(attr)
				if err != nil {
					return nil, fmt.Errorf("failed to convert inline attribute %s: %v", attr.Name, err)
				}
				inlineType.Fields = append(inlineType.Fields, *field)
			}

			// Add to types list
			p.goTypes = append(p.goTypes, inlineType)
		}
	}

	// Handle inline simple type
	if element.SimpleType != nil {
		fieldType = p.mapXSDTypeToGo("string") // Default to string for inline simple types
	}

	// Make pointer type if optional
	if isOptional && !isArray {
		fieldType = "*" + fieldType
	}
	// Make array type if needed
	if isArray {
		fieldType = "[]" + fieldType
	}

	xmlTag := element.Name
	if p.targetNamespace != "" {
		xmlTag = element.Name
	}
	// Add omitempty for optional elements
	if isOptional {
		xmlTag += ",omitempty"
	}

	jsonTag := ""
	if p.jsonCompatible {
		jsonTag = types.ToSnakeCase(element.Name)
		if isOptional {
			jsonTag += ",omitempty"
		}
	}

	field := &types.GoField{
		Name:       fieldName,
		Type:       fieldType,
		XMLTag:     xmlTag,
		JSONTag:    jsonTag,
		Comment:    types.GetDocumentation(element.Annotation),
		IsElement:  true,
		IsOptional: isOptional,
		IsArray:    isArray,
		MinOccurs:  min,
		MaxOccurs:  max,
	}

	return field, nil
}

// convertAttribute converts an XSD attribute to a Go field
func (p *XSDParser) convertAttribute(attr types.XSDAttribute) (*types.GoField, error) {
	fieldName := types.ToGoFieldName(attr.Name)
	fieldType := p.mapXSDTypeToGo(attr.Type)

	isOptional := attr.Use != "required"

	// Make pointer type if optional
	if isOptional {
		fieldType = "*" + fieldType
	}

	xmlTag := fmt.Sprintf("%s,attr", attr.Name)
	jsonTag := ""
	if p.jsonCompatible {
		jsonTag = types.ToSnakeCase(attr.Name)
		if isOptional {
			jsonTag += ",omitempty"
		}
	}

	field := &types.GoField{
		Name:        fieldName,
		Type:        fieldType,
		XMLTag:      xmlTag,
		JSONTag:     jsonTag,
		Comment:     types.GetDocumentation(attr.Annotation),
		IsAttribute: true,
		IsOptional:  isOptional,
	}

	return field, nil
}

// mapXSDTypeToGo maps an XSD type to a Go type
func (p *XSDParser) mapXSDTypeToGo(xsdType string) string {
	if xsdType == "" {
		return "string"
	}

	// Remove namespace prefix
	if colonIndex := strings.LastIndex(xsdType, ":"); colonIndex != -1 {
		xsdType = xsdType[colonIndex+1:]
	}

	// Use default Go language mapper for type mappings
	mapper := &generator.GoLanguageMapper{}
	mappings := mapper.GetBuiltinTypeMappings()
	for _, mapping := range mappings {
		if mapping.XSDType == xsdType {
			return mapping.TargetType
		}
	}

	// For custom types, convert to Go type name
	return types.ToGoTypeName(xsdType)
}

// generateConstantName generates a constant name for enums
func (p *XSDParser) generateConstantName(typeName, value string) string {
	// Use a more concise naming strategy similar to the sample code
	// TypeNameValue pattern instead of TYPE_NAME_VALUE
	cleanValue := strings.ReplaceAll(value, "-", "")
	cleanValue = strings.ReplaceAll(cleanValue, ".", "")
	cleanValue = types.ToPascalCase(cleanValue)

	return typeName + cleanValue
}

// findExistingType checks if a type with the given name already exists
func (p *XSDParser) findExistingType(typeName string) *types.GoType {
	for i := range p.goTypes {
		if p.goTypes[i].Name == typeName {
			return &p.goTypes[i]
		}
	}
	return nil
}

// GetGoTypes returns the converted Go types
func (p *XSDParser) GetGoTypes() []types.GoType {
	return p.goTypes
}

// GetSchema returns the parsed XSD schema
func (p *XSDParser) GetSchema() *types.XSDSchema {
	return p.schema
}

// processGroupRef processes an XSD group reference
func (p *XSDParser) processGroupRef(groupRef types.XSDGroupRef, goType *types.GoType, contextPath []string) error {
	// Extract group name from reference (remove namespace prefix if present)
	groupName := groupRef.Ref
	if parts := strings.Split(groupName, ":"); len(parts) > 1 {
		groupName = parts[1]
	}

	if p.debugMode {
		fmt.Printf("Processing group reference: %s in type %s\n", groupName, goType.Name)
	}

	// Find the group definition in the schema
	var foundGroup *types.XSDGroup
	for i := range p.schema.Groups {
		if p.schema.Groups[i].Name == groupName {
			foundGroup = &p.schema.Groups[i]
			break
		}
	}

	if foundGroup == nil {
		if p.debugMode {
			fmt.Printf("Warning: group '%s' not found\n", groupName)
		}
		return nil // Skip if group not found
	}

	if p.debugMode {
		fmt.Printf("Found group '%s', processing its contents\n", groupName)
	}

	// Process the group's content (choice, sequence, or all)
	if foundGroup.Choice != nil {
		if p.debugMode {
			fmt.Printf("Processing choice in group %s with %d elements\n", groupName, len(foundGroup.Choice.Elements))
		}
		if err := p.processChoiceWithContext(foundGroup.Choice, goType, contextPath); err != nil {
			return fmt.Errorf("failed to process choice in group %s: %v", groupName, err)
		}
	}

	if foundGroup.Sequence != nil {
		if p.debugMode {
			fmt.Printf("Processing sequence in group %s\n", groupName)
		}
		if err := p.processSequenceWithContext(foundGroup.Sequence, goType, contextPath); err != nil {
			return fmt.Errorf("failed to process sequence in group %s: %v", groupName, err)
		}
	}

	if foundGroup.All != nil {
		if p.debugMode {
			fmt.Printf("Processing all in group %s\n", groupName)
		}
		if err := p.processAllWithContext(foundGroup.All, goType, contextPath); err != nil {
			return fmt.Errorf("failed to process all in group %s: %v", groupName, err)
		}
	}

	return nil
}
