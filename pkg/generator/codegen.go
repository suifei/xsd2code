package generator

import (
	"fmt"
	"io/ioutil"
	"strings"
	"time"

	"github.com/suifei/xsd2code/pkg/types"
)

// TargetLanguage represents the target programming language for code generation
type TargetLanguage string

const (
	LanguageGo     TargetLanguage = "go"
	LanguageJava   TargetLanguage = "java"
	LanguageCSharp TargetLanguage = "csharp"
	LanguagePython TargetLanguage = "python"
)

// TypeMapping represents XSD to target language type mapping
type TypeMapping struct {
	XSDType    string
	TargetType string
}

// LanguageMapper defines the interface for language-specific type mappings
type LanguageMapper interface {
	GetBuiltinTypeMappings() []TypeMapping
	GetLanguage() TargetLanguage
	FormatTypeName(typeName string) string
	GetFileExtension() string
}

// GoLanguageMapper implements LanguageMapper for Go language
type GoLanguageMapper struct{}

// GetBuiltinTypeMappings returns the mapping from XSD built-in types to Go types
func (g *GoLanguageMapper) GetBuiltinTypeMappings() []TypeMapping {
	return []TypeMapping{
		// Standard XSD types
		{"string", "string"},
		{"normalizedString", "string"},
		{"token", "string"},
		{"anyURI", "string"},
		{"language", "string"},
		{"NMTOKEN", "string"},
		{"NMTOKENS", "[]string"},
		{"Name", "string"},
		{"NCName", "string"},
		{"ID", "string"},
		{"IDREF", "string"},
		{"IDREFS", "[]string"},
		{"ENTITY", "string"},
		{"ENTITIES", "[]string"},
		{"QName", "string"},

		{"boolean", "bool"},

		{"decimal", "float64"},
		{"float", "float32"},
		{"double", "float64"},

		{"duration", "string"}, // Could be time.Duration in future
		{"dateTime", "time.Time"},
		{"time", "string"},
		{"date", "string"},
		{"gYearMonth", "string"},
		{"gYear", "string"},
		{"gMonthDay", "string"},
		{"gDay", "string"},
		{"gMonth", "string"},

		{"hexBinary", "[]byte"},
		{"base64Binary", "[]byte"},

		{"integer", "int64"},
		{"nonPositiveInteger", "int64"},
		{"negativeInteger", "int64"},
		{"long", "int64"},
		{"int", "int32"},
		{"short", "int16"},
		{"byte", "int8"},
		{"nonNegativeInteger", "uint64"},
		{"unsignedLong", "uint64"},
		{"unsignedInt", "uint32"},
		{"unsignedShort", "uint16"},
		{"unsignedByte", "uint8"},
		{"positiveInteger", "uint64"},
		{"anyType", "interface{}"},

		// PLC Open IEC 61131-3 elementary types mapped to appropriate Go types
		{"BOOL", "bool"},          // Boolean
		{"BYTE", "uint8"},         // 8-bit unsigned integer (0-255)
		{"WORD", "uint16"},        // 16-bit unsigned integer (0-65535)
		{"DWORD", "uint32"},       // 32-bit unsigned integer (0-4294967295)
		{"LWORD", "uint64"},       // 64-bit unsigned integer (0-18446744073709551615)
		{"SINT", "int8"},          // Small signed integer (-128 to 127)
		{"INT", "int16"},          // Signed integer (-32768 to 32767)
		{"DINT", "int32"},         // Double signed integer (-2147483648 to 2147483647)
		{"LINT", "int64"},         // Long signed integer (-9223372036854775808 to 9223372036854775807)
		{"USINT", "uint8"},        // Unsigned small integer (0-255)
		{"UINT", "uint16"},        // Unsigned integer (0-65535)
		{"UDINT", "uint32"},       // Unsigned double integer (0-4294967295)
		{"ULINT", "uint64"},       // Unsigned long integer (0-18446744073709551615)
		{"REAL", "float32"},       // Single precision floating point
		{"LREAL", "float64"},      // Double precision floating point
		{"TIME", "time.Duration"}, // Time duration
		{"DATE", "time.Time"},     // Date
		{"DT", "time.Time"},       // Date and time
		{"TOD", "time.Time"},      // Time of day
	}
}

// GetLanguage returns the target language
func (g *GoLanguageMapper) GetLanguage() TargetLanguage {
	return LanguageGo
}

// FormatTypeName formats a type name according to Go conventions
func (g *GoLanguageMapper) FormatTypeName(typeName string) string {
	// Remove namespace prefix
	if colonIndex := strings.LastIndex(typeName, ":"); colonIndex != -1 {
		typeName = typeName[colonIndex+1:]
	}

	// Convert to PascalCase for Go
	parts := strings.FieldsFunc(typeName, func(c rune) bool {
		return c == '_' || c == '-' || c == '.'
	})

	var result strings.Builder
	for _, part := range parts {
		if len(part) > 0 {
			result.WriteString(strings.ToUpper(part[:1]))
			if len(part) > 1 {
				result.WriteString(strings.ToLower(part[1:]))
			}
		}
	}

	formatted := result.String()
	if formatted == "" {
		formatted = "UnknownType"
	}

	return formatted
}

// GetFileExtension returns the file extension for Go files
func (g *GoLanguageMapper) GetFileExtension() string {
	return ".go"
}

// JavaLanguageMapper implements LanguageMapper for Java language
type JavaLanguageMapper struct{}

// GetBuiltinTypeMappings returns the mapping from XSD built-in types to Java types
func (j *JavaLanguageMapper) GetBuiltinTypeMappings() []TypeMapping {
	return []TypeMapping{
		// Standard XSD types mapped to Java
		{"string", "String"},
		{"normalizedString", "String"},
		{"token", "String"},
		{"anyURI", "String"},
		{"language", "String"},
		{"NMTOKEN", "String"},
		{"NMTOKENS", "List<String>"},
		{"Name", "String"},
		{"NCName", "String"},
		{"ID", "String"},
		{"IDREF", "String"},
		{"IDREFS", "List<String>"},
		{"ENTITY", "String"},
		{"ENTITIES", "List<String>"},
		{"QName", "String"},

		{"boolean", "Boolean"},

		{"decimal", "BigDecimal"},
		{"float", "Float"},
		{"double", "Double"},

		{"duration", "Duration"},
		{"dateTime", "LocalDateTime"},
		{"time", "LocalTime"},
		{"date", "LocalDate"},
		{"gYearMonth", "YearMonth"},
		{"gYear", "Year"},
		{"gMonthDay", "MonthDay"},
		{"gDay", "String"},
		{"gMonth", "String"},

		{"hexBinary", "byte[]"},
		{"base64Binary", "byte[]"},

		{"integer", "BigInteger"},
		{"nonPositiveInteger", "BigInteger"},
		{"negativeInteger", "BigInteger"},
		{"long", "Long"},
		{"int", "Integer"},
		{"short", "Short"},
		{"byte", "Byte"},
		{"nonNegativeInteger", "BigInteger"},
		{"unsignedLong", "BigInteger"},
		{"unsignedInt", "Long"},
		{"unsignedShort", "Integer"},
		{"unsignedByte", "Short"},
		{"positiveInteger", "BigInteger"},
		{"anyType", "Object"},

		// PLC Open IEC 61131-3 elementary types mapped to appropriate Java types
		{"BOOL", "Boolean"},     // Boolean
		{"BYTE", "Byte"},        // 8-bit unsigned integer (0-255), closest Java equivalent
		{"WORD", "Integer"},     // 16-bit unsigned integer (0-65535), use Integer for safety
		{"DWORD", "Long"},       // 32-bit unsigned integer (0-4294967295), use Long for safety
		{"LWORD", "BigInteger"}, // 64-bit unsigned integer, BigInteger to handle full range
		{"SINT", "Byte"},        // Small signed integer (-128 to 127)
		{"INT", "Short"},        // Signed integer (-32768 to 32767)
		{"DINT", "Integer"},     // Double signed integer (-2147483648 to 2147483647)
		{"LINT", "Long"},        // Long signed integer (-9223372036854775808 to 9223372036854775807)
		{"USINT", "Short"},      // Unsigned small integer (0-255), use Short for safety
		{"UINT", "Integer"},     // Unsigned integer (0-65535), use Integer for safety
		{"UDINT", "Long"},       // Unsigned double integer (0-4294967295), use Long for safety
		{"ULINT", "BigInteger"}, // Unsigned long integer, BigInteger to handle full range
		{"REAL", "Float"},       // Single precision floating point
		{"LREAL", "Double"},     // Double precision floating point
		{"TIME", "Duration"},    // Time duration
		{"DATE", "LocalDate"},   // Date
		{"DT", "LocalDateTime"}, // Date and time
		{"TOD", "LocalTime"},    // Time of day
	}
}

// GetLanguage returns the target language
func (j *JavaLanguageMapper) GetLanguage() TargetLanguage {
	return LanguageJava
}

// FormatTypeName formats a type name according to Java conventions
func (j *JavaLanguageMapper) FormatTypeName(typeName string) string {
	// Remove namespace prefix
	if colonIndex := strings.LastIndex(typeName, ":"); colonIndex != -1 {
		typeName = typeName[colonIndex+1:]
	}

	// Convert to PascalCase for Java classes
	parts := strings.FieldsFunc(typeName, func(c rune) bool {
		return c == '_' || c == '-' || c == '.'
	})

	var result strings.Builder
	for _, part := range parts {
		if len(part) > 0 {
			result.WriteString(strings.ToUpper(part[:1]))
			if len(part) > 1 {
				result.WriteString(strings.ToLower(part[1:]))
			}
		}
	}

	formatted := result.String()
	if formatted == "" {
		formatted = "UnknownType"
	}

	return formatted
}

// GetFileExtension returns the file extension for Java files
func (j *JavaLanguageMapper) GetFileExtension() string {
	return ".java"
}

// CSharpLanguageMapper implements LanguageMapper for C# language
type CSharpLanguageMapper struct{}

// GetBuiltinTypeMappings returns the mapping from XSD built-in types to C# types
func (c *CSharpLanguageMapper) GetBuiltinTypeMappings() []TypeMapping {
	return []TypeMapping{
		// Standard XSD types mapped to C#
		{"string", "string"},
		{"normalizedString", "string"},
		{"token", "string"},
		{"anyURI", "string"},
		{"language", "string"},
		{"NMTOKEN", "string"},
		{"NMTOKENS", "List<string>"},
		{"Name", "string"},
		{"NCName", "string"},
		{"ID", "string"},
		{"IDREF", "string"},
		{"IDREFS", "List<string>"},
		{"ENTITY", "string"},
		{"ENTITIES", "List<string>"},
		{"QName", "string"},

		{"boolean", "bool"},

		{"decimal", "decimal"},
		{"float", "float"},
		{"double", "double"},

		{"duration", "TimeSpan"},
		{"dateTime", "DateTime"},
		{"time", "TimeSpan"},
		{"date", "DateTime"},
		{"gYearMonth", "DateTime"},
		{"gYear", "DateTime"},
		{"gMonthDay", "DateTime"},
		{"gDay", "string"},
		{"gMonth", "string"},

		{"hexBinary", "byte[]"},
		{"base64Binary", "byte[]"},

		{"integer", "long"},
		{"nonPositiveInteger", "long"},
		{"negativeInteger", "long"},
		{"long", "long"},
		{"int", "int"},
		{"short", "short"},
		{"byte", "sbyte"},
		{"nonNegativeInteger", "ulong"},
		{"unsignedLong", "ulong"},
		{"unsignedInt", "uint"},
		{"unsignedShort", "ushort"},
		{"unsignedByte", "byte"},
		{"positiveInteger", "ulong"},

		{"anyType", "object"},

		// PLC Open IEC 61131-3 elementary types mapped to appropriate C# types
		{"BOOL", "bool"},     // Boolean
		{"BYTE", "byte"},     // 8-bit unsigned integer (0-255)
		{"WORD", "ushort"},   // 16-bit unsigned integer (0-65535)
		{"DWORD", "uint"},    // 32-bit unsigned integer (0-4294967295)
		{"LWORD", "ulong"},   // 64-bit unsigned integer (0-18446744073709551615)
		{"SINT", "sbyte"},    // Small signed integer (-128 to 127)
		{"INT", "short"},     // Signed integer (-32768 to 32767)
		{"DINT", "int"},      // Double signed integer (-2147483648 to 2147483647)
		{"LINT", "long"},     // Long signed integer (-9223372036854775808 to 9223372036854775807)
		{"USINT", "byte"},    // Unsigned small integer (0-255)
		{"UINT", "ushort"},   // Unsigned integer (0-65535)
		{"UDINT", "uint"},    // Unsigned double integer (0-4294967295)
		{"ULINT", "ulong"},   // Unsigned long integer (0-18446744073709551615)
		{"REAL", "float"},    // Single precision floating point
		{"LREAL", "double"},  // Double precision floating point
		{"TIME", "TimeSpan"}, // Time duration
		{"DATE", "DateTime"}, // Date
		{"DT", "DateTime"},   // Date and time
		{"TOD", "TimeSpan"},  // Time of day
	}
}

// GetLanguage returns the target language
func (c *CSharpLanguageMapper) GetLanguage() TargetLanguage {
	return LanguageCSharp
}

// FormatTypeName formats a type name according to C# conventions
func (c *CSharpLanguageMapper) FormatTypeName(typeName string) string {
	// Remove namespace prefix
	if colonIndex := strings.LastIndex(typeName, ":"); colonIndex != -1 {
		typeName = typeName[colonIndex+1:]
	}

	// Convert to PascalCase for C# classes
	parts := strings.FieldsFunc(typeName, func(c rune) bool {
		return c == '_' || c == '-' || c == '.'
	})

	var result strings.Builder
	for _, part := range parts {
		if len(part) > 0 {
			result.WriteString(strings.ToUpper(part[:1]))
			if len(part) > 1 {
				result.WriteString(strings.ToLower(part[1:]))
			}
		}
	}

	formatted := result.String()
	if formatted == "" {
		formatted = "UnknownType"
	}

	return formatted
}

// GetFileExtension returns the file extension for C# files
func (c *CSharpLanguageMapper) GetFileExtension() string {
	return ".cs"
}

// CodeGenerator generates code from parsed XSD types
type CodeGenerator struct {
	packageName     string
	outputPath      string
	goTypes         []types.GoType
	jsonCompatible  bool
	includeComments bool
	debugMode       bool
	languageMapper  LanguageMapper
	typeMappings    map[string]string // Cache for type mappings
}

// NewCodeGenerator creates a new code generator
func NewCodeGenerator(packageName, outputPath string) *CodeGenerator {
	generator := &CodeGenerator{
		packageName:     packageName,
		outputPath:      outputPath,
		goTypes:         make([]types.GoType, 0),
		jsonCompatible:  false,
		includeComments: true,
		debugMode:       false,
		languageMapper:  &GoLanguageMapper{}, // Default to Go
		typeMappings:    make(map[string]string),
	}
	generator.initializeTypeMappings()
	return generator
}

// SetGoTypes sets the Go types to generate
func (g *CodeGenerator) SetGoTypes(goTypes []types.GoType) {
	g.goTypes = goTypes
}

// SetJSONCompatible enables or disables JSON compatibility
func (g *CodeGenerator) SetJSONCompatible(json bool) {
	g.jsonCompatible = json
}

// SetIncludeComments enables or disables comments
func (g *CodeGenerator) SetIncludeComments(comments bool) {
	g.includeComments = comments
}

// SetDebugMode enables or disables debug mode
func (g *CodeGenerator) SetDebugMode(debug bool) {
	g.debugMode = debug
}

// SetLanguageMapper sets the language mapper for the code generator
func (g *CodeGenerator) SetLanguageMapper(mapper LanguageMapper) {
	g.languageMapper = mapper
	g.initializeTypeMappings()
}

// initializeTypeMappings initializes the type mapping cache
func (g *CodeGenerator) initializeTypeMappings() {
	g.typeMappings = make(map[string]string)
	mappings := g.languageMapper.GetBuiltinTypeMappings()
	for _, mapping := range mappings {
		g.typeMappings[mapping.XSDType] = mapping.TargetType
	}
}

// GetTypeMapping returns the target language type for an XSD type
func (g *CodeGenerator) GetTypeMapping(xsdType string) (string, bool) {
	targetType, exists := g.typeMappings[xsdType]
	return targetType, exists
}

// GetBuiltinTypeMappings returns all builtin type mappings for the current language
func (g *CodeGenerator) GetBuiltinTypeMappings() []TypeMapping {
	return g.languageMapper.GetBuiltinTypeMappings()
}

// Generate generates the Go code and writes it to the output file
func (g *CodeGenerator) Generate() error {
	if g.debugMode {
		fmt.Printf("Generating Go code for %d types\n", len(g.goTypes))
	}

	code := g.generateCode()

	if err := ioutil.WriteFile(g.outputPath, []byte(code), 0644); err != nil {
		return fmt.Errorf("failed to write output file: %v", err)
	}

	if g.debugMode {
		fmt.Printf("Generated code written to: %s\n", g.outputPath)
	}

	return nil
}

// generateCode generates the complete Go code
func (g *CodeGenerator) generateCode() string {
	var builder strings.Builder

	// Package declaration and imports
	g.writeHeader(&builder)

	// Generate types
	for _, goType := range g.goTypes {
		g.writeType(&builder, goType)
		builder.WriteString("\n")
	}

	return builder.String()
}

// writeHeader writes the package declaration and imports
func (g *CodeGenerator) writeHeader(builder *strings.Builder) {
	builder.WriteString(fmt.Sprintf("// Code generated by xsd2code v3.0; DO NOT EDIT.\n"))
	builder.WriteString(fmt.Sprintf("// Generated on %s\n\n", time.Now().Format("2006-01-02 15:04:05")))

	builder.WriteString(fmt.Sprintf("package %s\n\n", g.packageName))

	// Imports
	builder.WriteString("import (\n")
	builder.WriteString("\t\"encoding/xml\"\n")

	if g.jsonCompatible {
		builder.WriteString("\t\"encoding/json\"\n")
	}

	// Check if we need time package
	needsTime := g.needsTimePackage()
	if needsTime {
		builder.WriteString("\t\"time\"\n")
	}

	builder.WriteString(")\n\n")
}

// needsTimePackage checks if the generated code needs the time package
func (g *CodeGenerator) needsTimePackage() bool {
	for _, goType := range g.goTypes {
		for _, field := range goType.Fields {
			if strings.Contains(field.Type, "time.Time") {
				return true
			}
		}
	}
	return false
}

// writeType writes a single Go type
func (g *CodeGenerator) writeType(builder *strings.Builder, goType types.GoType) {
	if goType.IsEnum {
		g.writeEnumType(builder, goType)
	} else {
		g.writeStructType(builder, goType)
	}
}

// writeStructType writes a struct type
func (g *CodeGenerator) writeStructType(builder *strings.Builder, goType types.GoType) {
	// Write comment
	if g.includeComments && goType.Comment != "" {
		g.writeComment(builder, fmt.Sprintf("%s %s", goType.Name, goType.Comment), "")
	} else if g.includeComments {
		g.writeComment(builder, fmt.Sprintf("%s represents %s", goType.Name, goType.XMLName), "")
	}

	// Write type declaration
	builder.WriteString(fmt.Sprintf("type %s struct {\n", goType.Name))

	// Write XMLName field if we have a namespace
	if goType.XMLName != "" {
		xmlNameTag := g.buildXMLNameTag(goType)
		jsonNameTag := ""
		if g.jsonCompatible {
			jsonNameTag = " json:\"-\""
		}
		builder.WriteString(fmt.Sprintf("\tXMLName xml.Name `xml:\"%s\"%s`\n", xmlNameTag, jsonNameTag))
	}

	// Write fields
	for _, field := range goType.Fields {
		g.writeField(builder, field)
	}

	builder.WriteString("}\n")
}

// writeEnumType writes an enum type with constants
func (g *CodeGenerator) writeEnumType(builder *strings.Builder, goType types.GoType) {
	// Write comment
	if g.includeComments && goType.Comment != "" {
		g.writeComment(builder, fmt.Sprintf("%s %s", goType.Name, goType.Comment), "")
	} else if g.includeComments {
		g.writeComment(builder, fmt.Sprintf("%s represents %s", goType.Name, goType.XMLName), "")
	}

	// Write type declaration
	baseType := goType.BaseType
	if baseType == "" {
		baseType = "string"
	}
	builder.WriteString(fmt.Sprintf("type %s %s\n\n", goType.Name, baseType))

	// Write constants
	if len(goType.Constants) > 0 {
		if g.includeComments {
			g.writeComment(builder, fmt.Sprintf("%s enumeration values", goType.Name), "")
		}
		builder.WriteString("const (\n")

		for _, constant := range goType.Constants {
			if g.includeComments && constant.Comment != "" {
				g.writeComment(builder, constant.Comment, "\t")
			}
			builder.WriteString(fmt.Sprintf("\t%s %s = %s\n", constant.Name, goType.Name, constant.Value))
		}

		builder.WriteString(")\n")
	}
}

// writeField writes a struct field
func (g *CodeGenerator) writeField(builder *strings.Builder, field types.GoField) {
	// Write comment
	if g.includeComments && field.Comment != "" {
		g.writeComment(builder, field.Comment, "\t")
	}

	// Build tags
	tags := g.buildFieldTags(field)

	// Write field
	builder.WriteString(fmt.Sprintf("\t%s %s %s\n", field.Name, field.Type, tags))
}

// buildXMLNameTag builds the XMLName tag
func (g *CodeGenerator) buildXMLNameTag(goType types.GoType) string {
	if goType.Namespace != "" {
		return fmt.Sprintf("%s %s", goType.Namespace, goType.XMLName)
	}
	return goType.XMLName
}

// buildFieldTags builds the struct tags for a field
func (g *CodeGenerator) buildFieldTags(field types.GoField) string {
	var tagParts []string

	// XML tag
	if field.XMLTag != "" {
		tagParts = append(tagParts, fmt.Sprintf("xml:\"%s\"", field.XMLTag))
	}

	// JSON tag
	if g.jsonCompatible && field.JSONTag != "" {
		tagParts = append(tagParts, fmt.Sprintf("json:\"%s\"", field.JSONTag))
	}

	if len(tagParts) == 0 {
		return ""
	}

	return "`" + strings.Join(tagParts, " ") + "`"
}

// writeComment writes a comment with proper formatting
func (g *CodeGenerator) writeComment(builder *strings.Builder, comment, indent string) {
	if comment == "" {
		return
	}

	// Clean up the comment
	comment = strings.TrimSpace(comment)
	comment = strings.ReplaceAll(comment, "\n", " ")
	comment = strings.ReplaceAll(comment, "\r", " ")

	// Handle multi-line comments
	words := strings.Fields(comment)
	const maxLineLength = 80

	currentLine := indent + "// "
	for i, word := range words {
		if len(currentLine)+len(word)+1 > maxLineLength && i > 0 {
			builder.WriteString(currentLine + "\n")
			currentLine = indent + "// " + word
		} else {
			if currentLine == indent+"// " {
				currentLine += word
			} else {
				currentLine += " " + word
			}
		}
	}

	if currentLine != indent+"// " {
		builder.WriteString(currentLine + "\n")
	}
}

// GenerateValidationCode generates validation functions for XSD types
func (g *CodeGenerator) GenerateValidationCode() string {
	var builder strings.Builder

	builder.WriteString("// Generated validation functions\n\n")
	builder.WriteString("import (\n")
	builder.WriteString("\t\"fmt\"\n")
	builder.WriteString("\t\"regexp\"\n")
	builder.WriteString("\t\"strconv\"\n")
	builder.WriteString("\t\"time\"\n")
	builder.WriteString(")\n\n")

	// Generate validation interface
	builder.WriteString("// Validator interface for all generated types\n")
	builder.WriteString("type Validator interface {\n")
	builder.WriteString("\tValidate() error\n")
	builder.WriteString("}\n\n")
	// Generate validation functions for each type
	for _, goType := range g.goTypes {
		g.generateTypeValidator(&builder, &goType)
	}

	// Generate helper validation functions
	g.generateValidationHelpers(&builder)

	return builder.String()
}

// generateTypeValidator generates validation method for a specific type
func (g *CodeGenerator) generateTypeValidator(builder *strings.Builder, goType *types.GoType) {
	typeName := goType.Name

	builder.WriteString(fmt.Sprintf("// Validate validates the %s struct\n", typeName))
	builder.WriteString(fmt.Sprintf("func (v *%s) Validate() error {\n", typeName))

	// Generate field validations
	for _, field := range goType.Fields {
		g.generateFieldValidation(builder, &field, typeName)
	}

	builder.WriteString("\treturn nil\n")
	builder.WriteString("}\n\n")
}

// generateFieldValidation generates validation code for a field
func (g *CodeGenerator) generateFieldValidation(builder *strings.Builder, field *types.GoField, typeName string) {
	fieldName := field.Name

	// Required field validation
	if !field.IsOptional && !field.IsArray {
		if strings.HasPrefix(field.Type, "*") {
			builder.WriteString(fmt.Sprintf("\tif v.%s == nil {\n", fieldName))
			builder.WriteString(fmt.Sprintf("\t\treturn fmt.Errorf(\"%s.%s is required\")\n", typeName, fieldName))
			builder.WriteString("\t}\n")
		} else if field.Type == "string" {
			builder.WriteString(fmt.Sprintf("\tif v.%s == \"\" {\n", fieldName))
			builder.WriteString(fmt.Sprintf("\t\treturn fmt.Errorf(\"%s.%s is required\")\n", typeName, fieldName))
			builder.WriteString("\t}\n")
		}
	}

	// Array length validation
	if field.IsArray && field.MinOccurs > 0 {
		builder.WriteString(fmt.Sprintf("\tif len(v.%s) < %d {\n", fieldName, field.MinOccurs))
		builder.WriteString(fmt.Sprintf("\t\treturn fmt.Errorf(\"%s.%s must have at least %d elements\")\n",
			typeName, fieldName, field.MinOccurs))
		builder.WriteString("\t}\n")
	}

	if field.IsArray && field.MaxOccurs > 0 && field.MaxOccurs != -1 {
		builder.WriteString(fmt.Sprintf("\tif len(v.%s) > %d {\n", fieldName, field.MaxOccurs))
		builder.WriteString(fmt.Sprintf("\t\treturn fmt.Errorf(\"%s.%s must have at most %d elements\")\n",
			typeName, fieldName, field.MaxOccurs))
		builder.WriteString("\t}\n")
	}

	// Type-specific validation
	g.generateTypeSpecificValidation(builder, field, typeName)
}

// generateTypeSpecificValidation generates type-specific validation code
func (g *CodeGenerator) generateTypeSpecificValidation(builder *strings.Builder, field *types.GoField, typeName string) {
	fieldName := field.Name
	baseType := strings.TrimPrefix(field.Type, "*")
	baseType = strings.TrimPrefix(baseType, "[]")

	switch baseType {
	case "time.Time":
		if field.IsOptional {
			builder.WriteString(fmt.Sprintf("\tif v.%s != nil {\n", fieldName))
			builder.WriteString(fmt.Sprintf("\t\tif err := validateDateTime(*v.%s); err != nil {\n", fieldName))
			builder.WriteString(fmt.Sprintf("\t\t\treturn fmt.Errorf(\"%s.%s: %%v\", err)\n", typeName, fieldName))
			builder.WriteString("\t\t}\n")
			builder.WriteString("\t}\n")
		} else if field.IsArray {
			builder.WriteString(fmt.Sprintf("\tfor i, dt := range v.%s {\n", fieldName))
			builder.WriteString("\t\tif err := validateDateTime(dt); err != nil {\n")
			builder.WriteString(fmt.Sprintf("\t\t\treturn fmt.Errorf(\"%s.%s[%%d]: %%v\", i, err)\n", typeName, fieldName))
			builder.WriteString("\t\t}\n")
			builder.WriteString("\t}\n")
		} else {
			builder.WriteString(fmt.Sprintf("\tif err := validateDateTime(v.%s); err != nil {\n", fieldName))
			builder.WriteString(fmt.Sprintf("\t\treturn fmt.Errorf(\"%s.%s: %%v\", err)\n", typeName, fieldName))
			builder.WriteString("\t}\n")
		}
	}
}

// generateValidationHelpers generates helper validation functions
func (g *CodeGenerator) generateValidationHelpers(builder *strings.Builder) {
	builder.WriteString("// Helper validation functions\n\n")

	// DateTime validation
	builder.WriteString("func validateDateTime(dt time.Time) error {\n")
	builder.WriteString("\tif dt.IsZero() {\n")
	builder.WriteString("\t\treturn fmt.Errorf(\"invalid datetime\")\n")
	builder.WriteString("\t}\n")
	builder.WriteString("\treturn nil\n")
	builder.WriteString("}\n\n")

	// String pattern validation
	builder.WriteString("func validatePattern(value, pattern string) error {\n")
	builder.WriteString("\tmatched, err := regexp.MatchString(pattern, value)\n")
	builder.WriteString("\tif err != nil {\n")
	builder.WriteString("\t\treturn fmt.Errorf(\"invalid pattern: %v\", err)\n")
	builder.WriteString("\t}\n")
	builder.WriteString("\tif !matched {\n")
	builder.WriteString("\t\treturn fmt.Errorf(\"value does not match pattern %s\", pattern)\n")
	builder.WriteString("\t}\n")
	builder.WriteString("\treturn nil\n")
	builder.WriteString("}\n\n")

	// Number range validation
	builder.WriteString("func validateIntRange(value, min, max int) error {\n")
	builder.WriteString("\tif value < min || value > max {\n")
	builder.WriteString("\t\treturn fmt.Errorf(\"value %d is out of range [%d, %d]\", value, min, max)\n")
	builder.WriteString("\t}\n")
	builder.WriteString("\treturn nil\n")
	builder.WriteString("}\n")
}

// GenerateTestCode generates test code for the generated types
func (g *CodeGenerator) GenerateTestCode() string {
	var builder strings.Builder

	builder.WriteString("// Generated test code\n\n")
	builder.WriteString("import (\n")
	builder.WriteString("\t\"encoding/xml\"\n")
	builder.WriteString("\t\"testing\"\n")
	builder.WriteString("\t\"time\"\n")
	builder.WriteString(")\n\n")
	// Generate test functions for each type
	for _, goType := range g.goTypes {
		g.generateTypeTest(&builder, &goType)
	}

	// Generate benchmark tests
	g.generateBenchmarkTests(&builder)

	return builder.String()
}

// generateTypeTest generates test function for a specific type
func (g *CodeGenerator) generateTypeTest(builder *strings.Builder, goType *types.GoType) {
	typeName := goType.Name

	// Test XML marshaling/unmarshaling
	builder.WriteString(fmt.Sprintf("func Test%sXMLMarshaling(t *testing.T) {\n", typeName))
	builder.WriteString(fmt.Sprintf("\toriginal := &%s{\n", typeName))

	// Generate test data for fields
	for _, field := range goType.Fields {
		g.generateTestFieldData(builder, &field)
	}

	builder.WriteString("\t}\n\n")

	// Test marshaling
	builder.WriteString("\t// Test marshaling\n")
	builder.WriteString("\txmlData, err := xml.Marshal(original)\n")
	builder.WriteString("\tif err != nil {\n")
	builder.WriteString("\t\tt.Fatalf(\"Failed to marshal XML: %v\", err)\n")
	builder.WriteString("\t}\n\n")

	// Test unmarshaling
	builder.WriteString("\t// Test unmarshaling\n")
	builder.WriteString(fmt.Sprintf("\tvar unmarshaled %s\n", typeName))
	builder.WriteString("\terr = xml.Unmarshal(xmlData, &unmarshaled)\n")
	builder.WriteString("\tif err != nil {\n")
	builder.WriteString("\t\tt.Fatalf(\"Failed to unmarshal XML: %v\", err)\n")
	builder.WriteString("\t}\n\n")

	// Test validation if implemented
	builder.WriteString("\t// Test validation\n")
	builder.WriteString("\tif err := original.Validate(); err != nil {\n")
	builder.WriteString("\t\tt.Errorf(\"Validation failed: %v\", err)\n")
	builder.WriteString("\t}\n")

	builder.WriteString("}\n\n")

	// Generate validation test
	g.generateValidationTest(builder, goType)
}

// generateTestFieldData generates test data for a field
func (g *CodeGenerator) generateTestFieldData(builder *strings.Builder, field *types.GoField) {
	fieldName := field.Name
	baseType := strings.TrimPrefix(field.Type, "*")
	baseType = strings.TrimPrefix(baseType, "[]")

	if field.IsArray {
		builder.WriteString(fmt.Sprintf("\t\t%s: []%s{", fieldName, baseType))
		g.generateSampleValue(builder, baseType)
		builder.WriteString("},\n")
	} else if field.IsOptional && strings.HasPrefix(field.Type, "*") {
		builder.WriteString(fmt.Sprintf("\t\t%s: ", fieldName))
		g.generatePointerValue(builder, baseType)
		builder.WriteString(",\n")
	} else {
		builder.WriteString(fmt.Sprintf("\t\t%s: ", fieldName))
		g.generateSampleValue(builder, baseType)
		builder.WriteString(",\n")
	}
}

// generateSampleValue generates a sample value for a type
func (g *CodeGenerator) generateSampleValue(builder *strings.Builder, typeName string) {
	switch typeName {
	case "string":
		builder.WriteString("\"test_value\"")
	case "int", "int32", "int64":
		builder.WriteString("42")
	case "float32", "float64":
		builder.WriteString("3.14")
	case "bool":
		builder.WriteString("true")
	case "time.Time":
		builder.WriteString("time.Now()")
	default:
		// Assume it's a custom type
		builder.WriteString(fmt.Sprintf("%s{}", typeName))
	}
}

// generatePointerValue generates a pointer value for a type
func (g *CodeGenerator) generatePointerValue(builder *strings.Builder, typeName string) {
	switch typeName {
	case "string":
		builder.WriteString("stringPtr(\"test_value\")")
	case "int", "int32", "int64":
		builder.WriteString("intPtr(42)")
	case "float32", "float64":
		builder.WriteString("floatPtr(3.14)")
	case "bool":
		builder.WriteString("boolPtr(true)")
	case "time.Time":
		builder.WriteString("timePtr(time.Now())")
	default:
		// Assume it's a custom type
		builder.WriteString(fmt.Sprintf("&%s{}", typeName))
	}
}

// generateValidationTest generates validation test cases
func (g *CodeGenerator) generateValidationTest(builder *strings.Builder, goType *types.GoType) {
	typeName := goType.Name

	builder.WriteString(fmt.Sprintf("func Test%sValidation(t *testing.T) {\n", typeName))

	// Test valid case
	builder.WriteString("\t// Test valid case\n")
	builder.WriteString(fmt.Sprintf("\tvalid := &%s{\n", typeName))
	for _, field := range goType.Fields {
		if !field.IsOptional {
			g.generateTestFieldData(builder, &field)
		}
	}
	builder.WriteString("\t}\n")
	builder.WriteString("\tif err := valid.Validate(); err != nil {\n")
	builder.WriteString("\t\tt.Errorf(\"Valid object should not have validation errors: %v\", err)\n")
	builder.WriteString("\t}\n\n")

	// Test invalid cases for required fields
	for _, field := range goType.Fields {
		if !field.IsOptional {
			g.generateInvalidFieldTest(builder, &field, typeName)
		}
	}

	builder.WriteString("}\n\n")
}

// generateInvalidFieldTest generates test for invalid field values
func (g *CodeGenerator) generateInvalidFieldTest(builder *strings.Builder, field *types.GoField, typeName string) {
	fieldName := field.Name

	builder.WriteString(fmt.Sprintf("\t// Test missing required field: %s\n", fieldName))
	builder.WriteString(fmt.Sprintf("\tinvalid%s := &%s{}\n", fieldName, typeName))
	builder.WriteString(fmt.Sprintf("\tif err := invalid%s.Validate(); err == nil {\n", fieldName))
	builder.WriteString(fmt.Sprintf("\t\tt.Error(\"Missing required field %s should cause validation error\")\n", fieldName))
	builder.WriteString("\t}\n\n")
}

// generateBenchmarkTests generates benchmark tests
func (g *CodeGenerator) generateBenchmarkTests(builder *strings.Builder) {
	builder.WriteString("// Benchmark tests\n\n")
	for _, goType := range g.goTypes {
		typeName := goType.Name

		builder.WriteString(fmt.Sprintf("func Benchmark%sMarshaling(b *testing.B) {\n", typeName))
		builder.WriteString(fmt.Sprintf("\tobj := &%s{\n", typeName))

		// Generate sample data
		for _, field := range goType.Fields {
			if !field.IsOptional {
				g.generateTestFieldData(builder, &field)
			}
		}

		builder.WriteString("\t}\n\n")

		builder.WriteString("\tb.ResetTimer()\n")
		builder.WriteString("\tfor i := 0; i < b.N; i++ {\n")
		builder.WriteString("\t\t_, err := xml.Marshal(obj)\n")
		builder.WriteString("\t\tif err != nil {\n")
		builder.WriteString("\t\t\tb.Fatal(err)\n")
		builder.WriteString("\t\t}\n")
		builder.WriteString("\t}\n")
		builder.WriteString("}\n\n")
	}

	// Helper functions for pointer creation
	builder.WriteString("// Helper functions for creating pointers\n\n")
	builder.WriteString("func stringPtr(s string) *string { return &s }\n")
	builder.WriteString("func intPtr(i int) *int { return &i }\n")
	builder.WriteString("func floatPtr(f float64) *float64 { return &f }\n")
	builder.WriteString("func boolPtr(b bool) *bool { return &b }\n")
	builder.WriteString("func timePtr(t time.Time) *time.Time { return &t }\n")
}
