package generator

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/suifei/xsd2code/pkg/types"
)

// Global type mapping registry initialized on package load
var typeMappingRegistry *CommonTypeMappingRegistry

func init() {
	typeMappingRegistry = NewCommonTypeMappingRegistry()
}

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
	GetCustomTypeMappings() []TypeMapping // 分离自定义类型映射
	GetLanguage() TargetLanguage
	FormatTypeName(typeName string) string
	GetFileExtension() string
	GetImportStatements() []string
	GetStructTemplate() string
	GetEnumTemplate() string
}

// BaseLanguageMapper provides common functionality for all language mappers
type BaseLanguageMapper struct{}

// FormatTypeName provides common type name formatting logic
func (b *BaseLanguageMapper) FormatTypeName(typeName string) string {
	// Remove namespace prefix
	if colonIndex := strings.LastIndex(typeName, ":"); colonIndex != -1 {
		typeName = typeName[colonIndex+1:]
	}

	// Convert to PascalCase
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

// GoLanguageMapper implements LanguageMapper for Go language
type GoLanguageMapper struct {
	BaseLanguageMapper
}

// GetBuiltinTypeMappings returns the mapping from XSD built-in types to Go types
func (g *GoLanguageMapper) GetBuiltinTypeMappings() []TypeMapping {
	return typeMappingRegistry.GetMappingsForLanguage(LanguageGo)
}

// GetCustomTypeMappings returns custom type mappings (e.g., PLC types)
func (g *GoLanguageMapper) GetCustomTypeMappings() []TypeMapping {
	// Return empty slice as custom mappings are included in the registry
	return []TypeMapping{}
}

// GetLanguage returns the target language
func (g *GoLanguageMapper) GetLanguage() TargetLanguage {
	return LanguageGo
}

// GetFileExtension returns the file extension for Go files
func (g *GoLanguageMapper) GetFileExtension() string {
	return ".go"
}

// GetImportStatements returns the import statements for Go
func (g *GoLanguageMapper) GetImportStatements() []string {
	return []string{
		"\"encoding/xml\"",
		"\"time\"",
	}
}

// GetStructTemplate returns the struct template for Go
func (g *GoLanguageMapper) GetStructTemplate() string {
	return `type {{.Name}} struct {
{{- if .XMLName}}
	XMLName xml.Name ` + "`xml:\"{{.XMLName}}\"`" + `
{{- end}}
{{- range .Fields}}
	{{.Name}} {{.Type}} ` + "`{{.Tags}}`" + `
{{- end}}
}`
}

// GetEnumTemplate returns the enum template for Go
func (g *GoLanguageMapper) GetEnumTemplate() string {
	return `type {{.Name}} {{.BaseType}}

const (
{{- range .Constants}}
	{{.Name}} {{$.Name}} = {{.Value}}
{{- end}}
)`
}

// JavaLanguageMapper implements LanguageMapper for Java language
type JavaLanguageMapper struct {
	BaseLanguageMapper
}

// GetBuiltinTypeMappings returns the mapping from XSD built-in types to Java types
func (j *JavaLanguageMapper) GetBuiltinTypeMappings() []TypeMapping {
	return typeMappingRegistry.GetMappingsForLanguage(LanguageJava)
}

// GetCustomTypeMappings returns custom type mappings for Java
func (j *JavaLanguageMapper) GetCustomTypeMappings() []TypeMapping {
	// Return empty slice as custom mappings are included in the registry
	return []TypeMapping{}
}

// GetLanguage returns the target language
func (j *JavaLanguageMapper) GetLanguage() TargetLanguage {
	return LanguageJava
}

// GetFileExtension returns the file extension for Java files
func (j *JavaLanguageMapper) GetFileExtension() string {
	return ".java"
}

// GetImportStatements returns the import statements for Java
func (j *JavaLanguageMapper) GetImportStatements() []string {
	return []string{
		"import java.util.*;",
		"import java.time.*;",
		"import java.math.*;",
		"import javax.xml.bind.annotation.*;",
	}
}

// GetStructTemplate returns the class template for Java
func (j *JavaLanguageMapper) GetStructTemplate() string {
	return `@XmlRootElement(name = "{{.XMLName}}")
public class {{.Name}} {
{{- range .Fields}}
    private {{.Type}} {{.Name}};
{{- end}}

{{- range .Fields}}
    public {{.Type}} get{{.Name}}() {
        return {{.Name}};
    }

    public void set{{.Name}}({{.Type}} {{.Name}}) {
        this.{{.Name}} = {{.Name}};
    }
{{- end}}
}`
}

// GetEnumTemplate returns the enum template for Java
func (j *JavaLanguageMapper) GetEnumTemplate() string {
	return `public enum {{.Name}} {
{{- range .Constants}}
    {{.Name}}("{{.Value}}"),
{{- end}};

    private final String value;

    {{.Name}}(String value) {
        this.value = value;
    }

    public String getValue() {
        return value;
    }
}`
}

// CSharpLanguageMapper implements LanguageMapper for C# language
type CSharpLanguageMapper struct {
	BaseLanguageMapper
}

// GetBuiltinTypeMappings returns the mapping from XSD built-in types to C# types
func (c *CSharpLanguageMapper) GetBuiltinTypeMappings() []TypeMapping {
	return typeMappingRegistry.GetMappingsForLanguage(LanguageCSharp)
}

// GetCustomTypeMappings returns custom type mappings for C#
func (c *CSharpLanguageMapper) GetCustomTypeMappings() []TypeMapping {
	return []TypeMapping{
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

// GetFileExtension returns the file extension for C# files
func (c *CSharpLanguageMapper) GetFileExtension() string {
	return ".cs"
}

// GetImportStatements returns the import statements for C#
func (c *CSharpLanguageMapper) GetImportStatements() []string {
	return []string{
		"using System;",
		"using System.Collections.Generic;",
		"using System.Xml.Serialization;",
		"using System.Text.Json.Serialization;",
	}
}

// GetStructTemplate returns the class template for C#
func (c *CSharpLanguageMapper) GetStructTemplate() string {
	return `[XmlRoot("{{.XMLName}}")]
public class {{.Name}}
{
{{- range .Fields}}
    [XmlElement("{{.XMLName}}")]
    public {{.Type}} {{.Name}} { get; set; }
{{- end}}
}`
}

// GetEnumTemplate returns the enum template for C#
func (c *CSharpLanguageMapper) GetEnumTemplate() string {
	return `public enum {{.Name}}
{
{{- range .Constants}}
    [XmlEnum("{{.Value}}")]
    {{.Name}},
{{- end}}
}`
}

// PythonLanguageMapper implements LanguageMapper for Python
type PythonLanguageMapper struct {
	BaseLanguageMapper
}

// GetLanguage returns the language identifier
func (p *PythonLanguageMapper) GetLanguage() TargetLanguage {
	return LanguagePython
}

// GetBuiltinTypeMappings returns the builtin type mappings for Python
func (p *PythonLanguageMapper) GetBuiltinTypeMappings() []TypeMapping {
	return typeMappingRegistry.GetMappingsForLanguage(LanguagePython)
}

// GetCustomTypeMappings returns custom type mappings for PLC/industrial types
func (p *PythonLanguageMapper) GetCustomTypeMappings() []TypeMapping {
	// Return empty slice as custom mappings are included in the registry
	return []TypeMapping{}
}

// FormatTypeName formats type names according to Python conventions
func (p *PythonLanguageMapper) FormatTypeName(typeName string) string {
	return p.BaseLanguageMapper.FormatTypeName(typeName)
}

// GetFileExtension returns the file extension for Python
func (p *PythonLanguageMapper) GetFileExtension() string {
	return ".py"
}

// GetImportStatements returns the import statements for Python
func (p *PythonLanguageMapper) GetImportStatements() []string {
	return []string{
		"from dataclasses import dataclass, field",
		"from typing import List, Optional, Any",
		"from datetime import datetime, date, time, timedelta",
		"from enum import Enum",
		"import xml.etree.ElementTree as ET",
	}
}

// GetStructTemplate returns the class template for Python
func (p *PythonLanguageMapper) GetStructTemplate() string {
	return `@dataclass
class {{.Name}}:
    {{range .Fields}}{{.Name}}: {{.Type}}{{if .IsOptional}} = None{{end}}
    {{end}}`
}

// GetEnumTemplate returns the enum template for Python
func (p *PythonLanguageMapper) GetEnumTemplate() string {
	return `class {{.Name}}(Enum):
    {{range .Constants}}{{.Name}} = "{{.Value}}"
    {{end}}`
}

// CodeGenerator generates code from parsed XSD types
type CodeGenerator struct {
	packageName       string
	outputPath        string
	goTypes           []types.GoType
	jsonCompatible    bool
	includeComments   bool
	debugMode         bool
	enableCustomTypes bool // 控制是否启用自定义类型映射
	languageMapper    LanguageMapper
	typeMappings      map[string]string // Cache for type mappings
}

// NewCodeGenerator creates a new code generator
func NewCodeGenerator(packageName, outputPath string) *CodeGenerator {
	generator := &CodeGenerator{
		packageName:       packageName,
		outputPath:        outputPath,
		goTypes:           make([]types.GoType, 0),
		jsonCompatible:    false,
		includeComments:   true,
		debugMode:         false,
		enableCustomTypes: false,               // 默认关闭自定义类型映射
		languageMapper:    &GoLanguageMapper{}, // Default to Go
		typeMappings:      make(map[string]string),
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

// SetEnableCustomTypes enables or disables custom type mappings (e.g., PLC types)
func (g *CodeGenerator) SetEnableCustomTypes(enable bool) {
	g.enableCustomTypes = enable
	g.initializeTypeMappings() // Reinitialize mappings
}

// SetLanguageMapper sets the language mapper for the code generator
func (g *CodeGenerator) SetLanguageMapper(mapper LanguageMapper) {
	g.languageMapper = mapper
	g.initializeTypeMappings()
}

// initializeTypeMappings initializes the type mapping cache
func (g *CodeGenerator) initializeTypeMappings() {
	g.typeMappings = make(map[string]string)

	// Always include built-in mappings
	builtinMappings := g.languageMapper.GetBuiltinTypeMappings()
	for _, mapping := range builtinMappings {
		g.typeMappings[mapping.XSDType] = mapping.TargetType
	}

	// Optionally include custom mappings
	if g.enableCustomTypes {
		customMappings := g.languageMapper.GetCustomTypeMappings()
		for _, mapping := range customMappings {
			g.typeMappings[mapping.XSDType] = mapping.TargetType
		}
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

// Generate generates the code for the target language and writes it to the output file
func (g *CodeGenerator) Generate() error {
	if g.debugMode {
		fmt.Printf("Generating %s code for %d types\n", g.languageMapper.GetLanguage(), len(g.goTypes))
	}

	code := g.generateCode()

	if err := os.WriteFile(g.outputPath, []byte(code), 0644); err != nil {
		return fmt.Errorf("failed to write output file: %v", err)
	}

	if g.debugMode {
		fmt.Printf("Generated code written to: %s\n", g.outputPath)
	}

	return nil
}

// generateCode generates the complete code for the target language
func (g *CodeGenerator) generateCode() string {
	var builder strings.Builder

	// Package declaration and imports
	g.writeHeader(&builder)

	// Generate types
	for _, goType := range g.goTypes {
		g.writeType(&builder, goType)
		builder.WriteString("\n")
	}

	// Generate helper functions for Go if needed
	if g.languageMapper.GetLanguage() == LanguageGo {
		g.writeGoHelperFunctions(&builder)
	}

	// Close namespace for C#
	if g.languageMapper.GetLanguage() == LanguageCSharp {
		builder.WriteString("}\n")
	}

	return builder.String()
}

// writeHeader writes the package declaration and imports for the target language
func (g *CodeGenerator) writeHeader(builder *strings.Builder) {
	timestamp := time.Now().Format("2006-01-02 15:04:05")

	switch g.languageMapper.GetLanguage() {
	case LanguagePython:
		builder.WriteString("# Code generated by xsd2code v3.0; DO NOT EDIT.\n")
		builder.WriteString("# Generated on " + timestamp + "\n\n")
		g.writePythonHeader(builder)
	default:
		builder.WriteString("// Code generated by xsd2code v3.0; DO NOT EDIT.\n")
		builder.WriteString("// Generated on " + timestamp + "\n\n")
		switch g.languageMapper.GetLanguage() {
		case LanguageGo:
			g.writeGoHeader(builder)
		case LanguageJava:
			g.writeJavaHeader(builder)
		case LanguageCSharp:
			g.writeCSharpHeader(builder)
		default:
			g.writeGoHeader(builder) // Fallback to Go
		}
	}
}

// writeGoHeader writes Go-specific package and imports
func (g *CodeGenerator) writeGoHeader(builder *strings.Builder) {
	builder.WriteString("package " + g.packageName + "\n\n")

	// Determine needed imports dynamically
	neededImports := g.determineNeededImports()

	// Imports
	builder.WriteString("import (\n")
	for _, importStmt := range neededImports {
		builder.WriteString("\t" + importStmt + "\n")
	}
	builder.WriteString(")\n\n")
}

// writeJavaHeader writes Java-specific package and imports
func (g *CodeGenerator) writeJavaHeader(builder *strings.Builder) {
	builder.WriteString("package " + g.packageName + ";\n\n")

	for _, importStmt := range g.languageMapper.GetImportStatements() {
		builder.WriteString(importStmt + "\n")
	}
	builder.WriteString("\n")
}

// writeCSharpHeader writes C#-specific namespace and using statements
func (g *CodeGenerator) writeCSharpHeader(builder *strings.Builder) {
	for _, importStmt := range g.languageMapper.GetImportStatements() {
		builder.WriteString(importStmt + "\n")
	}
	builder.WriteString("\n")
	builder.WriteString("namespace " + g.packageName + "\n{\n")
}

// writePythonHeader writes Python-specific imports
func (g *CodeGenerator) writePythonHeader(builder *strings.Builder) {
	for _, importStmt := range g.languageMapper.GetImportStatements() {
		builder.WriteString(importStmt + "\n")
	}
	builder.WriteString("\n")
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

// determineNeededImports determines which imports are needed based on the generated types
func (g *CodeGenerator) determineNeededImports() []string {
	imports := make(map[string]bool)

	// Always include basic imports
	imports["\"encoding/xml\""] = true

	// Check if JSON compatibility is needed
	// if g.jsonCompatible {
	// 	imports["\"encoding/json\""] = true
	// }

	// Check if time package is needed
	if g.needsTimePackage() {
		imports["\"time\""] = true
	}

	// Check if regexp package is needed (for pattern validation)
	needsRegexp := false
	for _, goType := range g.goTypes {
		if goType.HasPattern {
			needsRegexp = true
			break
		}
	}
	if needsRegexp {
		imports["\"regexp\""] = true
	}

	// Check if strings package is needed (for whiteSpace processing)
	needsStrings := false
	for _, goType := range g.goTypes {
		if goType.HasWhiteSpace || goType.HasMinLength || goType.HasMaxLength || goType.HasLength {
			needsStrings = true
			break
		}
	}
	if needsStrings {
		imports["\"strings\""] = true
	}

	// Convert map to sorted slice
	result := make([]string, 0, len(imports))
	for imp := range imports {
		result = append(result, imp)
	}

	return result
}

// writeType writes a single type for the target language
func (g *CodeGenerator) writeType(builder *strings.Builder, goType types.GoType) {
	switch g.languageMapper.GetLanguage() {
	case LanguageGo:
		g.writeGoType(builder, goType)
	case LanguageJava:
		g.writeJavaType(builder, goType)
	case LanguageCSharp:
		g.writeCSharpType(builder, goType)
	case LanguagePython:
		g.writePythonType(builder, goType)
	default:
		g.writeGoType(builder, goType) // Fallback to Go
	}
}

// writeGoType writes a Go type
func (g *CodeGenerator) writeGoType(builder *strings.Builder, goType types.GoType) {
	if goType.IsEnum {
		g.writeGoEnumType(builder, goType)
	} else if goType.HasPattern || goType.HasMinLength || goType.HasMaxLength ||
		goType.HasMinInclusive || goType.HasMaxInclusive ||
		goType.HasMinExclusive || goType.HasMaxExclusive ||
		goType.HasTotalDigits || goType.HasFractionDigits ||
		goType.HasWhiteSpace || goType.HasLength || goType.HasFixedValue {
		// This is a simple type with restrictions (like pattern, whiteSpace, length, or fixed value)
		g.writeGoRestrictedType(builder, goType)
	} else {
		g.writeGoStructType(builder, goType)
	}
}

// writeJavaType writes a Java type
func (g *CodeGenerator) writeJavaType(builder *strings.Builder, goType types.GoType) {
	if goType.IsEnum {
		g.writeJavaEnumType(builder, goType)
	} else if goType.HasPattern || goType.HasMinLength || goType.HasMaxLength ||
		goType.HasMinInclusive || goType.HasMaxInclusive ||
		goType.HasMinExclusive || goType.HasMaxExclusive ||
		goType.HasTotalDigits || goType.HasFractionDigits {
		// This is a simple type with restrictions (like pattern)
		g.writeJavaRestrictedType(builder, goType)
	} else {
		g.writeJavaClassType(builder, goType)
	}
}

// writeCSharpType writes a C# type
func (g *CodeGenerator) writeCSharpType(builder *strings.Builder, goType types.GoType) {
	if goType.IsEnum {
		if csharpMapper, ok := g.languageMapper.(*CSharpLanguageMapper); ok {
			csharpMapper.writeCSharpEnumType(builder, goType, g)
		}
	} else if goType.HasPattern || goType.HasMinLength || goType.HasMaxLength ||
		goType.HasMinInclusive || goType.HasMaxInclusive ||
		goType.HasMinExclusive || goType.HasMaxExclusive ||
		goType.HasTotalDigits || goType.HasFractionDigits {
		// This is a simple type with restrictions (like pattern)
		g.writeCSharpRestrictedType(builder, goType)
	} else {
		g.writeCSharpClassType(builder, goType)
	}
}

// writePythonType writes a Python type
func (g *CodeGenerator) writePythonType(builder *strings.Builder, goType types.GoType) {
	if goType.IsEnum {
		g.writePythonEnumType(builder, goType)
	} else if goType.HasPattern || goType.HasMinLength || goType.HasMaxLength ||
		goType.HasMinInclusive || goType.HasMaxInclusive ||
		goType.HasMinExclusive || goType.HasMaxExclusive ||
		goType.HasTotalDigits || goType.HasFractionDigits {
		// This is a simple type with restrictions (like pattern)
		g.writePythonRestrictedType(builder, goType)
	} else {
		g.writePythonClassType(builder, goType)
	}
}

// writeGoStructType writes a Go struct type
func (g *CodeGenerator) writeGoStructType(builder *strings.Builder, goType types.GoType) {
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
		g.writeGoField(builder, field)
	}

	builder.WriteString("}\n")
}

// writeGoEnumType writes a Go enum type with constants
func (g *CodeGenerator) writeGoEnumType(builder *strings.Builder, goType types.GoType) {
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

// writeGoField writes a Go struct field
func (g *CodeGenerator) writeGoField(builder *strings.Builder, field types.GoField) {
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
	builder.WriteString("\t\"strings\"\n")
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
	builder.WriteString("}\n\n")

	// WhiteSpace processing helper
	builder.WriteString("// applyWhiteSpaceProcessing applies XSD whiteSpace facet processing\n")
	builder.WriteString("func applyWhiteSpaceProcessing(value, whiteSpaceAction string) string {\n")
	builder.WriteString("\tswitch whiteSpaceAction {\n")
	builder.WriteString("\tcase \"replace\":\n")
	builder.WriteString("\t\t// Replace tab, newline, and carriage return with space\n")
	builder.WriteString("\t\tvalue = strings.ReplaceAll(value, \"\\t\", \" \")\n")
	builder.WriteString("\t\tvalue = strings.ReplaceAll(value, \"\\n\", \" \")\n")
	builder.WriteString("\t\tvalue = strings.ReplaceAll(value, \"\\r\", \" \")\n")
	builder.WriteString("\t\treturn value\n")
	builder.WriteString("\tcase \"collapse\":\n")
	builder.WriteString("\t\t// First apply replace processing\n")
	builder.WriteString("\t\tvalue = strings.ReplaceAll(value, \"\\t\", \" \")\n")
	builder.WriteString("\t\tvalue = strings.ReplaceAll(value, \"\\n\", \" \")\n")
	builder.WriteString("\t\tvalue = strings.ReplaceAll(value, \"\\r\", \" \")\n")
	builder.WriteString("\t\t// Then collapse sequences of spaces and trim\n")
	builder.WriteString("\t\tvalue = regexp.MustCompile(`\\\\s+`).ReplaceAllString(value, \" \")\n")
	builder.WriteString("\t\tvalue = strings.TrimSpace(value)\n")
	builder.WriteString("\t\treturn value\n")
	builder.WriteString("\tcase \"preserve\":\n")
	builder.WriteString("\t\tfallthrough\n")
	builder.WriteString("\tdefault:\n")
	builder.WriteString("\t\t// Preserve all whitespace as-is\n")
	builder.WriteString("\t\treturn value\n")
	builder.WriteString("\t}\n")
	builder.WriteString("}\n\n")

	// Fixed value validation
	builder.WriteString("func validateFixedValue(value, expectedValue string) error {\n")
	builder.WriteString("\tif value != expectedValue {\n")
	builder.WriteString("\t\treturn fmt.Errorf(\"value '%s' does not match fixed value '%s'\", value, expectedValue)\n")
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
	// Check if this is an enum type or struct type
	if goType.IsEnum || (len(goType.Fields) == 0 && len(goType.Constants) > 0) {
		// For enum types, use value initialization instead of pointer
		if len(goType.Constants) > 0 {
			// Use the first constant name directly
			builder.WriteString(fmt.Sprintf("\toriginal := %s\n\n", goType.Constants[0].Name))
		} else {
			// Fallback to empty string for string-based enums
			builder.WriteString(fmt.Sprintf("\toriginal := %s(\"\")\n\n", typeName))
		}
	} else {
		// For struct types, use pointer initialization
		builder.WriteString(fmt.Sprintf("\toriginal := &%s{\n", typeName))

		// Generate test data for fields
		for _, field := range goType.Fields {
			g.generateTestFieldData(builder, &field)
		}

		builder.WriteString("\t}\n\n")
	}

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
	case "uint", "uint32", "uint64":
		builder.WriteString("42")
	case "int8":
		builder.WriteString("int8(42)")
	case "int16":
		builder.WriteString("int16(42)")
	case "uint8":
		builder.WriteString("uint8(42)")
	case "uint16":
		builder.WriteString("uint16(42)")
	case "float32", "float64":
		builder.WriteString("3.14")
	case "bool":
		builder.WriteString("true")
	case "time.Time":
		builder.WriteString("time.Now()")
	case "time.Duration":
		builder.WriteString("time.Second")
	default:
		// For custom types, check if it appears to be an enum (string-based type)
		// by using a simple heuristic: if it's a custom type that ends with "Type",
		// treat it as a string-based enum, otherwise as a struct
		if strings.HasSuffix(typeName, "Type") && !strings.Contains(typeName, "DataType") {
			// For string-based enum types, use type conversion with empty string
			builder.WriteString(fmt.Sprintf("%s(\"\")", typeName))
		} else {
			// Assume it's a custom struct type
			builder.WriteString(fmt.Sprintf("%s{}", typeName))
		}
	}
}

// generatePointerValue generates a pointer value for a type
func (g *CodeGenerator) generatePointerValue(builder *strings.Builder, typeName string) {
	switch typeName {
	case "string":
		builder.WriteString("stringPtr(\"test_value\")")
	case "int", "int32", "int64":
		builder.WriteString("intPtr(42)")
	case "uint", "uint32", "uint64":
		builder.WriteString("uintPtr(42)")
	case "float32", "float64":
		builder.WriteString("floatPtr(3.14)")
	case "bool":
		builder.WriteString("boolPtr(true)")
	case "time.Time":
		builder.WriteString("timePtr(time.Now())")
	case "time.Duration":
		builder.WriteString("durationPtr(time.Second)")
	default:
		// For custom types, check if it appears to be an enum (string-based type)
		// by using a simple heuristic: if it's a custom type that ends with "Type",
		// treat it as a string-based enum, otherwise as a struct
		if strings.HasSuffix(typeName, "Type") && !strings.Contains(typeName, "DataType") {
			// For string-based enum types, create inline function to return pointer
			builder.WriteString(fmt.Sprintf("func() *%s { v := %s(\"\"); return &v }()", typeName, typeName))
		} else {
			// Assume it's a custom struct type
			builder.WriteString(fmt.Sprintf("&%s{}", typeName))
		}
	}
}

// generateValidationTest generates validation test cases
func (g *CodeGenerator) generateValidationTest(builder *strings.Builder, goType *types.GoType) {
	typeName := goType.Name

	builder.WriteString(fmt.Sprintf("func Test%sValidation(t *testing.T) {\n", typeName))

	// Test valid case
	builder.WriteString("\t// Test valid case\n")
	// Check if this is an enum type or struct type
	if goType.IsEnum || (len(goType.Fields) == 0 && len(goType.Constants) > 0) {
		// For enum types, create a value directly
		if len(goType.Constants) > 0 {
			// Use the first constant name directly instead of wrapping the value in quotes
			builder.WriteString(fmt.Sprintf("\tvalid := %s\n", goType.Constants[0].Name))
		} else {
			// Fallback to empty string for string-based enums
			builder.WriteString(fmt.Sprintf("\tvalid := %s(\"\")\n", typeName))
		}
		builder.WriteString("\tif err := valid.Validate(); err != nil {\n")
		builder.WriteString("\t\tt.Errorf(\"Valid object should not have validation errors: %v\", err)\n")
		builder.WriteString("\t}\n\n")

		// For enum types, test invalid values if there are constants defined
		if len(goType.Constants) > 0 {
			builder.WriteString("\t// Test invalid enum value\n")
			builder.WriteString(fmt.Sprintf("\tinvalid := %s(\"invalid_value\")\n", typeName))
			builder.WriteString("\tif err := invalid.Validate(); err == nil {\n")
			builder.WriteString("\t\tt.Error(\"Invalid enum value should cause validation error\")\n")
			builder.WriteString("\t}\n\n")
		}
	} else {
		// For struct types, use the existing logic
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

		// Use heuristic to detect enum types and generate appropriate initialization
		if strings.HasSuffix(typeName, "Type") && !strings.Contains(typeName, "DataType") {
			// For string-based enum types, use value initialization
			builder.WriteString(fmt.Sprintf("\tobj := %s(\"\")\n\n", typeName))
		} else {
			// For struct types, use pointer initialization
			builder.WriteString(fmt.Sprintf("\tobj := &%s{\n", typeName))

			// Generate sample data
			for _, field := range goType.Fields {
				if !field.IsOptional {
					g.generateTestFieldData(builder, &field)
				}
			}

			builder.WriteString("\t}\n\n")
		}

		builder.WriteString("\tb.ResetTimer()\n")
		builder.WriteString("\tfor i := 0; i < b.N; i++ {\n")
		builder.WriteString("\t\t_, err := xml.Marshal(obj)\n")
		builder.WriteString("\t\tif err != nil {\n")
		builder.WriteString("\t\t\tb.Fatal(err)\n")
		builder.WriteString("\t\t}\n")
		builder.WriteString("\t}\n")
		builder.WriteString("}\n\n")
	} // Helper functions for pointer creation
	builder.WriteString("// Helper functions for creating pointers\n\n")
	builder.WriteString("func stringPtr(s string) *string { return &s }\n")
	builder.WriteString("func intPtr(i int) *int { return &i }\n")
	builder.WriteString("func uintPtr(u uint64) *uint64 { return &u }\n")
	builder.WriteString("func floatPtr(f float64) *float64 { return &f }\n")
	builder.WriteString("func boolPtr(b bool) *bool { return &b }\n")
	builder.WriteString("func timePtr(t time.Time) *time.Time { return &t }\n")
	builder.WriteString("func durationPtr(d time.Duration) *time.Duration { return &d }\n")
	builder.WriteString("func edgeModifierTypePtr(e EdgeModifierType) *EdgeModifierType { return &e }\n")
	builder.WriteString("func storageModifierTypePtr(s StorageModifierType) *StorageModifierType { return &s }\n")
}

// Java type generation methods

// writeJavaClassType writes a Java class type
func (g *CodeGenerator) writeJavaClassType(builder *strings.Builder, goType types.GoType) {
	// Write comment
	if g.includeComments && goType.Comment != "" {
		g.writeComment(builder, fmt.Sprintf("%s %s", goType.Name, goType.Comment), "")
	} else if g.includeComments {
		g.writeComment(builder, fmt.Sprintf("%s represents %s", goType.Name, goType.XMLName), "")
	}

	// Write class declaration with XML annotations
	if goType.XMLName != "" {
		builder.WriteString(fmt.Sprintf("@XmlRootElement(name = \"%s\")\n", goType.XMLName))
		if goType.Namespace != "" {
			builder.WriteString(fmt.Sprintf("@XmlType(namespace = \"%s\")\n", goType.Namespace))
		}
	}
	builder.WriteString(fmt.Sprintf("public class %s {\n", goType.Name))

	// Write fields
	for _, field := range goType.Fields {
		g.writeJavaField(builder, field)
	}

	builder.WriteString("\n")

	// Write getters and setters
	for _, field := range goType.Fields {
		g.writeJavaGetterSetter(builder, field)
	}

	builder.WriteString("}\n")
}

// writeJavaField writes a Java field
func (g *CodeGenerator) writeJavaField(builder *strings.Builder, field types.GoField) {
	// Convert Go type to Java type
	javaType := g.convertToJavaType(field.Type) // Write field with annotations
	if field.XMLTag != "" {
		if strings.Contains(field.XMLTag, ",attr") {
			builder.WriteString("    @XmlAttribute\n")
		} else {
			builder.WriteString("    @XmlElement\n")
		}
	}

	builder.WriteString(fmt.Sprintf("    private %s %s;\n", javaType, strings.ToLower(field.Name[:1])+field.Name[1:]))
}

// writeJavaGetterSetter writes getter and setter methods for a Java field
func (g *CodeGenerator) writeJavaGetterSetter(builder *strings.Builder, field types.GoField) {
	javaType := g.convertToJavaType(field.Type)
	fieldName := strings.ToLower(field.Name[:1]) + field.Name[1:]
	capitalizedName := field.Name

	// Getter
	builder.WriteString(fmt.Sprintf("    public %s get%s() {\n", javaType, capitalizedName))
	builder.WriteString(fmt.Sprintf("        return %s;\n", fieldName))
	builder.WriteString("    }\n\n")

	// Setter
	builder.WriteString(fmt.Sprintf("    public void set%s(%s %s) {\n", capitalizedName, javaType, fieldName))
	builder.WriteString(fmt.Sprintf("        this.%s = %s;\n", fieldName, fieldName))
	builder.WriteString("    }\n\n")
}

// writeJavaEnumType writes a Java enum type
func (g *CodeGenerator) writeJavaEnumType(builder *strings.Builder, goType types.GoType) {
	// Write comment
	if g.includeComments && goType.Comment != "" {
		g.writeComment(builder, fmt.Sprintf("%s %s", goType.Name, goType.Comment), "")
	}

	builder.WriteString(fmt.Sprintf("public enum %s {\n", goType.Name)) // Write enum constants
	for i, constant := range goType.Constants {
		// Remove quotes from constant value if present
		value := strings.Trim(constant.Value, `"`)
		if i == len(goType.Constants)-1 {
			builder.WriteString(fmt.Sprintf("    %s(\"%s\");\n\n", strings.ToUpper(constant.Name), value))
		} else {
			builder.WriteString(fmt.Sprintf("    %s(\"%s\"),\n", strings.ToUpper(constant.Name), value))
		}
	}

	// Write enum constructor and methods
	builder.WriteString("    private final String value;\n\n")
	builder.WriteString(fmt.Sprintf("    %s(String value) {\n", goType.Name))
	builder.WriteString("        this.value = value;\n")
	builder.WriteString("    }\n\n")
	builder.WriteString("    public String getValue() {\n")
	builder.WriteString("        return value;\n")
	builder.WriteString("    }\n")
	builder.WriteString("}\n")
}

// C# type generation methods

// writeCSharpClassType writes a C# class type
func (g *CodeGenerator) writeCSharpClassType(builder *strings.Builder, goType types.GoType) {
	// Write comment
	if g.includeComments && goType.Comment != "" {
		g.writeComment(builder, fmt.Sprintf("%s %s", goType.Name, goType.Comment), "")
	} else if g.includeComments {
		g.writeComment(builder, fmt.Sprintf("%s represents %s", goType.Name, goType.XMLName), "")
	}

	// Write class declaration with XML attributes
	if goType.XMLName != "" {
		builder.WriteString(fmt.Sprintf("[XmlRoot(\"%s\"", goType.XMLName))
		if goType.Namespace != "" {
			builder.WriteString(fmt.Sprintf(", Namespace = \"%s\"", goType.Namespace))
		}
		builder.WriteString(")]\n")
	}
	builder.WriteString(fmt.Sprintf("public class %s\n{\n", goType.Name))

	// Write properties
	for _, field := range goType.Fields {
		g.writeCSharpProperty(builder, field)
	}

	builder.WriteString("}\n")
}

// writeCSharpProperty writes a C# property
func (g *CodeGenerator) writeCSharpProperty(builder *strings.Builder, field types.GoField) {
	// Convert Go type to C# type
	csharpType := g.convertToCSharpType(field.Type)

	// Write property with XML attributes
	if field.XMLTag != "" {
		if strings.Contains(field.XMLTag, ",attr") {
			builder.WriteString("    [XmlAttribute]\n")
		} else {
			builder.WriteString("    [XmlElement]\n")
		}
	}

	if g.jsonCompatible && field.JSONTag != "" {
		builder.WriteString(fmt.Sprintf("    [JsonPropertyName(\"%s\")]\n", field.JSONTag))
	}

	builder.WriteString(fmt.Sprintf("    public %s %s { get; set; }\n\n", csharpType, field.Name))
}

// writeCSharpEnumType writes a C# enum type
func (cs *CSharpLanguageMapper) writeCSharpEnumType(builder *strings.Builder, goType types.GoType, generator *CodeGenerator) {
	// Write comment
	if generator.includeComments && goType.Comment != "" {
		generator.writeComment(builder, fmt.Sprintf("%s %s", goType.Name, goType.Comment), "")
	}

	builder.WriteString(fmt.Sprintf("public enum %s\n{\n", goType.Name))

	// Write enum constants
	for _, constant := range goType.Constants {
		if generator.includeComments && constant.Comment != "" {
			generator.writeComment(builder, constant.Comment, "    ")
		}
		builder.WriteString(fmt.Sprintf("    %s,\n", constant.Name))
	}

	builder.WriteString("}\n")
}

// Python type generation methods

// writePythonClassType writes a Python dataclass type
func (g *CodeGenerator) writePythonClassType(builder *strings.Builder, goType types.GoType) {
	// Write comment
	if g.includeComments && goType.Comment != "" {
		g.writePythonComment(builder, fmt.Sprintf("%s %s", goType.Name, goType.Comment), "")
	} else if g.includeComments {
		g.writePythonComment(builder, fmt.Sprintf("%s represents %s", goType.Name, goType.XMLName), "")
	}

	// Write dataclass decorator
	builder.WriteString("@dataclass\n")
	builder.WriteString(fmt.Sprintf("class %s:\n", goType.Name))

	// Write fields
	if len(goType.Fields) == 0 {
		builder.WriteString("    pass\n")
	} else {
		for _, field := range goType.Fields {
			g.writePythonField(builder, field)
		}
	}
}

// writePythonField writes a Python dataclass field
func (g *CodeGenerator) writePythonField(builder *strings.Builder, field types.GoField) {
	// Write comment
	if g.includeComments && field.Comment != "" {
		g.writePythonComment(builder, field.Comment, "    ")
	}

	// Convert Go type to Python type
	pythonType := g.convertToPythonType(field.Type)

	// Handle optional fields
	isOptional := strings.Contains(field.XMLTag, "omitempty") || field.IsOptional
	if isOptional && !strings.HasPrefix(pythonType, "Optional[") {
		pythonType = fmt.Sprintf("Optional[%s]", pythonType)
	}

	// Write field with type annotation
	if isOptional {
		builder.WriteString(fmt.Sprintf("    %s: %s = None\n", field.Name, pythonType))
	} else {
		builder.WriteString(fmt.Sprintf("    %s: %s\n", field.Name, pythonType))
	}
}

// writePythonEnumType writes a Python enum type
func (g *CodeGenerator) writePythonEnumType(builder *strings.Builder, goType types.GoType) {
	// Write comment
	if g.includeComments && goType.Comment != "" {
		g.writePythonComment(builder, fmt.Sprintf("%s %s", goType.Name, goType.Comment), "")
	}

	builder.WriteString(fmt.Sprintf("class %s(Enum):\n", goType.Name))

	// Write enum constants
	if len(goType.Constants) == 0 {
		builder.WriteString("    pass\n")
	} else {
		for _, constant := range goType.Constants {
			// Remove quotes from constant value if present
			value := strings.Trim(constant.Value, `"`)
			builder.WriteString(fmt.Sprintf("    %s = \"%s\"\n", strings.ToUpper(constant.Name), value))
		}
	}
}

// writePythonComment writes a Python comment
func (g *CodeGenerator) writePythonComment(builder *strings.Builder, comment, indent string) {
	if comment == "" {
		return
	}

	lines := strings.Split(comment, "\n")
	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line != "" {
			builder.WriteString(fmt.Sprintf("%s# %s\n", indent, line))
		}
	}
}

// convertToPythonType converts Go types to Python types
func (g *CodeGenerator) convertToPythonType(goType string) string {
	// Remove pointer markers
	goType = strings.TrimPrefix(goType, "*")

	// Handle array/slice types
	if strings.HasPrefix(goType, "[]") {
		elementType := goType[2:]
		pythonElementType := g.convertToPythonType(elementType)
		return fmt.Sprintf("List[%s]", pythonElementType)
	}

	// Check for mapped types first
	if mapped, exists := g.GetTypeMapping(goType); exists {
		return mapped
	}

	// Common Go to Python type conversions
	conversions := map[string]string{
		"string":        "str",
		"int":           "int",
		"int8":          "int",
		"int16":         "int",
		"int32":         "int",
		"int64":         "int",
		"uint":          "int",
		"uint8":         "int",
		"uint16":        "int",
		"uint32":        "int",
		"uint64":        "int",
		"float32":       "float",
		"float64":       "float",
		"bool":          "bool",
		"time.Time":     "datetime",
		"time.Duration": "timedelta",
		"[]byte":        "bytes",
		"interface{}":   "Any",
	}

	if pythonType, exists := conversions[goType]; exists {
		return pythonType
	}

	// Default to the type name (assume it's a custom type)
	return goType
}

// convertToJavaType converts Go types to Java types
func (g *CodeGenerator) convertToJavaType(goType string) string {
	// Remove pointer markers
	goType = strings.TrimPrefix(goType, "*")

	// Handle array/slice types
	if strings.HasPrefix(goType, "[]") {
		elementType := goType[2:]
		javaElementType := g.convertToJavaType(elementType)
		return fmt.Sprintf("List<%s>", javaElementType)
	}

	// Check for mapped types first
	if mapped, exists := g.GetTypeMapping(goType); exists {
		return mapped
	}

	// Common Go to Java type conversions
	conversions := map[string]string{
		"string":        "String",
		"int":           "Integer",
		"int8":          "Byte",
		"int16":         "Short",
		"int32":         "Integer",
		"int64":         "Long",
		"uint":          "Integer",
		"uint8":         "Integer",
		"uint16":        "Integer",
		"uint32":        "Integer",
		"uint64":        "Long",
		"float32":       "Float",
		"float64":       "Double",
		"bool":          "Boolean",
		"time.Time":     "LocalDateTime",
		"time.Duration": "Duration",
		"[]byte":        "byte[]",
		"interface{}":   "Object",
	}

	if javaType, exists := conversions[goType]; exists {
		return javaType
	}

	// Default to the type name (assume it's a custom type)
	return goType
}

// convertToCSharpType converts Go types to C# types
func (g *CodeGenerator) convertToCSharpType(goType string) string {
	// Remove pointer markers
	goType = strings.TrimPrefix(goType, "*")

	// Handle array/slice types
	if strings.HasPrefix(goType, "[]") {
		elementType := goType[2:]
		csharpElementType := g.convertToCSharpType(elementType)
		return fmt.Sprintf("List<%s>", csharpElementType)
	}

	// Check for mapped types first
	if mapped, exists := g.GetTypeMapping(goType); exists {
		return mapped
	}

	// Common Go to C# type conversions
	conversions := map[string]string{
		"string":        "string",
		"int":           "int",
		"int8":          "sbyte",
		"int16":         "short",
		"int32":         "int",
		"int64":         "long",
		"uint":          "uint",
		"uint8":         "byte",
		"uint16":        "ushort",
		"uint32":        "uint",
		"uint64":        "ulong",
		"float32":       "float",
		"float64":       "double",
		"bool":          "bool",
		"time.Time":     "DateTime",
		"time.Duration": "TimeSpan",
		"[]byte":        "byte[]",
		"interface{}":   "object",
	}

	if csharpType, exists := conversions[goType]; exists {
		return csharpType
	}

	// Default to the type name (assume it's a custom type)
	return goType
}

// writeGoRestrictedType writes a Go type with restrictions like pattern
func (g *CodeGenerator) writeGoRestrictedType(builder *strings.Builder, goType types.GoType) {
	// Write comment
	if g.includeComments && goType.Comment != "" {
		g.writeComment(builder, fmt.Sprintf("%s %s", goType.Name, goType.Comment), "")
	} else if g.includeComments {
		commentText := fmt.Sprintf("%s represents a %s", goType.Name, goType.BaseType)
		if goType.HasPattern {
			commentText += " with pattern validation"
		} else if goType.HasMinLength || goType.HasMaxLength {
			commentText += " with length restrictions"
		} else if goType.HasLength {
			commentText += " with exact length restriction"
		} else if goType.HasWhiteSpace {
			commentText += " with whiteSpace processing"
		} else if goType.HasFixedValue {
			commentText += " with fixed value constraint"
		} else if goType.HasMinInclusive || goType.HasMaxInclusive || goType.HasMinExclusive || goType.HasMaxExclusive {
			commentText += " with range restrictions"
		}
		g.writeComment(builder, commentText, "")
	}

	// Write type declaration - simple types with restrictions typically alias the base type
	baseType := goType.BaseType
	if baseType == "" {
		baseType = "string" // Default to string if no base type specified
	}
	builder.WriteString(fmt.Sprintf("type %s %s\n\n", goType.Name, baseType))

	// Add validation function
	g.writeGoTypeValidation(builder, goType)
}

// writeGoTypeValidation writes Go validation function for a type
func (g *CodeGenerator) writeGoTypeValidation(builder *strings.Builder, goType types.GoType) {
	// Write comment for validation function
	if g.includeComments {
		g.writeComment(builder, fmt.Sprintf("Validate validates the %s format", goType.Name), "")
	}

	// Begin validation function
	if goType.BaseType == "string" || goType.BaseType == "" {
		builder.WriteString(fmt.Sprintf("func (v %s) Validate() bool {\n", goType.Name))
	} else {
		builder.WriteString(fmt.Sprintf("func (v %s) Validate() bool {\n", goType.Name))
	}

	// Add validation logic based on restriction type
	if goType.HasPattern {
		// For pattern validation, use regexp
		builder.WriteString(fmt.Sprintf("\t// Validate against pattern: %s\n", goType.PatternValue))
		builder.WriteString("\tpattern := regexp.MustCompile(`" + goType.PatternValue + "`)\n")
		builder.WriteString("\treturn pattern.MatchString(string(v))\n")
	} else if goType.HasLength {
		// For exact length validation
		builder.WriteString("\tstrVal := string(v)\n")
		// Apply whiteSpace processing if needed
		if goType.HasWhiteSpace {
			builder.WriteString(fmt.Sprintf("\tstrVal = applyWhiteSpaceProcessing(strVal, \"%s\")\n", goType.WhiteSpace))
		}
		builder.WriteString(fmt.Sprintf("\treturn len(strVal) == %s\n", goType.Length))
	} else if goType.HasMinLength || goType.HasMaxLength {
		// For length validation
		builder.WriteString("\tstrVal := string(v)\n")
		// Apply whiteSpace processing if needed
		if goType.HasWhiteSpace {
			builder.WriteString(fmt.Sprintf("\tstrVal = applyWhiteSpaceProcessing(strVal, \"%s\")\n", goType.WhiteSpace))
		}
		builder.WriteString("\tlength := len(strVal)\n")
		if goType.HasMinLength {
			builder.WriteString(fmt.Sprintf("\tif length < %s {\n\t\treturn false\n\t}\n", goType.MinLength))
		}
		if goType.HasMaxLength {
			builder.WriteString(fmt.Sprintf("\tif length > %s {\n\t\treturn false\n\t}\n", goType.MaxLength))
		}
		builder.WriteString("\treturn true\n")
	} else if goType.HasFixedValue {
		// For fixed value validation
		builder.WriteString(fmt.Sprintf("\treturn string(v) == \"%s\"\n", goType.FixedValue))
	} else if goType.HasMinInclusive || goType.HasMaxInclusive || goType.HasMinExclusive || goType.HasMaxExclusive {
		// For numeric range validation
		switch goType.BaseType {
		case "int", "int8", "int16", "int32", "int64":
			g.writeIntRangeValidation(builder, goType)
		case "float32", "float64":
			g.writeFloatRangeValidation(builder, goType)
		default:
			builder.WriteString("\t// Validation not supported for this type\n")
			builder.WriteString("\treturn true\n")
		}
	} else if goType.HasTotalDigits || goType.HasFractionDigits {
		// For digit validation
		builder.WriteString("\t// Validate number of digits\n")
		builder.WriteString("\t// Note: This is a simplified validation\n")
		builder.WriteString("\t return true\n")
	} else {
		// Default case: no validation
		builder.WriteString("\treturn true\n")
	}

	// Close function
	builder.WriteString("}\n")
}

// writeIntRangeValidation writes validation for integer range constraints
func (g *CodeGenerator) writeIntRangeValidation(builder *strings.Builder, goType types.GoType) {
	builder.WriteString("\tval := int64(v)\n")
	if goType.HasMinInclusive {
		builder.WriteString(fmt.Sprintf("\tif val < %s {\n\t\treturn false\n\t}\n", goType.MinInclusive))
	} else if goType.HasMinExclusive {
		builder.WriteString(fmt.Sprintf("\tif val <= %s {\n\t\treturn false\n\t}\n", goType.MinExclusive))
	}

	if goType.HasMaxInclusive {
		builder.WriteString(fmt.Sprintf("\tif val > %s {\n\t\treturn false\n\t}\n", goType.MaxInclusive))
	} else if goType.HasMaxExclusive {
		builder.WriteString(fmt.Sprintf("\tif val >= %s {\n\t\treturn false\n\t}\n", goType.MaxExclusive))
	}

	builder.WriteString("\treturn true\n")
}

// writeFloatRangeValidation writes validation for floating point range constraints
func (g *CodeGenerator) writeFloatRangeValidation(builder *strings.Builder, goType types.GoType) {
	builder.WriteString("\tval := float64(v)\n")
	if goType.HasMinInclusive {
		builder.WriteString(fmt.Sprintf("\tif val < %s {\n\t\treturn false\n\t}\n", goType.MinInclusive))
	} else if goType.HasMinExclusive {
		builder.WriteString(fmt.Sprintf("\tif val <= %s {\n\t\treturn false\n\t}\n", goType.MinExclusive))
	}

	if goType.HasMaxInclusive {
		builder.WriteString(fmt.Sprintf("\tif val > %s {\n\t\treturn false\n\t}\n", goType.MaxInclusive))
	} else if goType.HasMaxExclusive {
		builder.WriteString(fmt.Sprintf("\tif val >= %s {\n\t\treturn false\n\t}\n", goType.MaxExclusive))
	}

	builder.WriteString("\treturn true\n")
}

// writeJavaRestrictedType writes a Java class with restriction validation
func (g *CodeGenerator) writeJavaRestrictedType(builder *strings.Builder, goType types.GoType) {
	// Write comment
	if g.includeComments && goType.Comment != "" {
		g.writeComment(builder, fmt.Sprintf("%s %s", goType.Name, goType.Comment), "")
	} else if g.includeComments {
		g.writeComment(builder, fmt.Sprintf("%s represents %s", goType.Name, goType.XMLName), "")
	}

	// Write class declaration with XML annotations
	if goType.XMLName != "" {
		builder.WriteString(fmt.Sprintf("@XmlRootElement(name = \"%s\")\n", goType.XMLName))
		if goType.Namespace != "" {
			builder.WriteString(fmt.Sprintf("@XmlType(namespace = \"%s\")\n", goType.Namespace))
		}
	}
	builder.WriteString(fmt.Sprintf("public class %s {\n", goType.Name))

	// Write fields
	for _, field := range goType.Fields {
		g.writeJavaField(builder, field)
	}

	builder.WriteString("\n")

	// Write getters and setters
	for _, field := range goType.Fields {
		g.writeJavaGetterSetter(builder, field)
	}

	// Add validation method
	g.writeJavaTypeValidation(builder, goType)

	builder.WriteString("}\n")
}

// writeCSharpRestrictedType writes a C# class with restriction validation
func (g *CodeGenerator) writeCSharpRestrictedType(builder *strings.Builder, goType types.GoType) {
	// Write comment
	if g.includeComments && goType.Comment != "" {
		g.writeComment(builder, fmt.Sprintf("%s %s", goType.Name, goType.Comment), "")
	} else if g.includeComments {
		g.writeComment(builder, fmt.Sprintf("%s represents a restricted %s", goType.Name, goType.BaseType), "")
	}

	// Write class declaration
	builder.WriteString(fmt.Sprintf("public class %s\n{\n", goType.Name))

	// Write private field for the value
	baseType := goType.BaseType
	if baseType == "" {
		baseType = "string"
	}
	csharpType := g.convertToCSharpType(baseType)
	builder.WriteString(fmt.Sprintf("    private %s _value;\n\n", csharpType))

	// Write constructor with validation
	builder.WriteString(fmt.Sprintf("    public %s(%s value)\n", goType.Name, csharpType))
	builder.WriteString("    {\n")

	if goType.HasPattern {
		builder.WriteString(fmt.Sprintf("        if (!System.Text.RegularExpressions.Regex.IsMatch(value, @\"%s\"))\n", goType.PatternValue))
		builder.WriteString("            throw new ArgumentException(\"Invalid format\");\n")
	}

	if goType.HasMinLength || goType.HasMaxLength {
		if goType.HasMinLength {
			builder.WriteString(fmt.Sprintf("        if (value.Length < %s)\n", goType.MinLength))
			builder.WriteString("            throw new ArgumentException(\"Value too short\");\n")
		}
		if goType.HasMaxLength {
			builder.WriteString(fmt.Sprintf("        if (value.Length > %s)\n", goType.MaxLength))
			builder.WriteString("            throw new ArgumentException(\"Value too long\");\n")
		}
	}

	builder.WriteString("        _value = value;\n")
	builder.WriteString("    }\n\n")

	// Write Value property
	builder.WriteString(fmt.Sprintf("    public %s Value => _value;\n\n", csharpType))

	// Write ToString override
	builder.WriteString("    public override string ToString() => _value.ToString();\n\n")

	// Write implicit conversion operators
	builder.WriteString(fmt.Sprintf("    public static implicit operator %s(%s obj) => obj._value;\n", csharpType, goType.Name))
	builder.WriteString(fmt.Sprintf("    public static implicit operator %s(%s value) => new %s(value);\n", goType.Name, csharpType, goType.Name))

	builder.WriteString("}\n")
}

// writePythonRestrictedType writes a Python class with restriction validation
func (g *CodeGenerator) writePythonRestrictedType(builder *strings.Builder, goType types.GoType) {
	// Add import for re module if there is a pattern restriction
	if goType.HasPattern {
		builder.WriteString("import re\n")
	}

	// Write class docstring comment
	if g.includeComments && goType.Comment != "" {
		builder.WriteString("class " + goType.Name + "(str):\n")
		builder.WriteString("    \"\"\"\n    " + goType.Comment + "\n")
	} else {
		builder.WriteString("class " + goType.Name + "(str):\n")
		builder.WriteString("    \"\"\"\n    Represents a string with ")
		if goType.HasPattern {
			builder.WriteString("pattern validation.\n    Pattern: " + goType.PatternValue)
		} else if goType.HasMinLength || goType.HasMaxLength {
			builder.WriteString("length validation.")
		} else if goType.HasMinInclusive || goType.HasMaxInclusive || goType.HasMinExclusive || goType.HasMaxExclusive {
			builder.WriteString("range validation.")
		} else {
			builder.WriteString("validation.")
		}
		builder.WriteString("\n")
	}
	builder.WriteString("    \"\"\"\n")

	// Add static pattern for regex
	if goType.HasPattern {
		// Escape backslashes for Python strings
		pyPattern := strings.Replace(goType.PatternValue, "\\", "\\\\", -1)
		builder.WriteString("    _pattern = re.compile(r'" + pyPattern + "')\n")
	}

	// Add __new__ method with validation
	builder.WriteString("    \n    def __new__(cls, value):\n")
	if goType.HasPattern {
		builder.WriteString("        if not cls._pattern.match(value):\n")
		builder.WriteString("            raise ValueError(f\"Invalid " + strings.ToLower(goType.Name) + " format: {value}\")\n")
	} else if goType.HasMinLength && goType.HasMaxLength {
		builder.WriteString("        length = len(value)\n")
		builder.WriteString("        if length < " + goType.MinLength + " or length > " + goType.MaxLength + ":\n")
		builder.WriteString("            raise ValueError(f\"Length must be between " + goType.MinLength + " and " + goType.MaxLength + ", got {length}\")\n")
	} else if goType.HasMinLength {
		builder.WriteString("        if len(value) < " + goType.MinLength + ":\n")
		builder.WriteString("            raise ValueError(f\"Length must be at least " + goType.MinLength + ", got {len(value)}\")\n")
	} else if goType.HasMaxLength {
		builder.WriteString("        if len(value) > " + goType.MaxLength + ":\n")
		builder.WriteString("            raise ValueError(f\"Length must be at most " + goType.MaxLength + ", got {len(value)}\")\n")
	}

	builder.WriteString("        return super().__new__(cls, value)\n")

	// Add validation method
	builder.WriteString("    \n    def validate(self):\n")
	builder.WriteString("        \"\"\"Validates the format.\"\"\"\n")
	if goType.HasPattern {
		builder.WriteString("        return bool(self._pattern.match(self))\n")
	} else if goType.HasMinLength || goType.HasMaxLength {
		builder.WriteString("        length = len(self)\n")
		conditions := []string{}
		if goType.HasMinLength {
			conditions = append(conditions, "length >= "+goType.MinLength)
		}
		if goType.HasMaxLength {
			conditions = append(conditions, "length <= "+goType.MaxLength)
		}
		builder.WriteString("        return " + strings.Join(conditions, " and ") + "\n")
	} else {
		builder.WriteString("        return True\n")
	}
}

// writeJavaTypeValidation writes Java validation method for a type
func (g *CodeGenerator) writeJavaTypeValidation(builder *strings.Builder, goType types.GoType) {
	// Write validation method
	builder.WriteString("    public boolean validate() {\n")

	if goType.HasPattern {
		builder.WriteString(fmt.Sprintf("        return value.matches(\"%s\");\n", goType.PatternValue))
	} else if goType.HasMinLength || goType.HasMaxLength {
		if goType.HasMinLength && goType.HasMaxLength {
			builder.WriteString("        int length = value.length();\n")
			builder.WriteString(fmt.Sprintf("        return length >= %s && length <= %s;\n", goType.MinLength, goType.MaxLength))
		} else if goType.HasMinLength {
			builder.WriteString(fmt.Sprintf("        return value.length() >= %s;\n", goType.MinLength))
		} else if goType.HasMaxLength {
			builder.WriteString(fmt.Sprintf("        return value.length() <= %s;\n", goType.MaxLength))
		}
	} else if goType.HasMinInclusive || goType.HasMaxInclusive || goType.HasMinExclusive || goType.HasMaxExclusive {
		// For numeric validation
		builder.WriteString("        // TODO: Add numeric range validation\n")
		builder.WriteString("        return true;\n")
	} else {
		builder.WriteString("        return true;\n")
	}

	builder.WriteString("    }\n")
}

// writeGoHelperFunctions writes helper functions needed by the generated Go code
func (g *CodeGenerator) writeGoHelperFunctions(builder *strings.Builder) {
	// Check if we need any helper functions
	needsWhiteSpaceHelper := false
	needsRegexpHelper := false

	for _, goType := range g.goTypes {
		if goType.HasWhiteSpace {
			needsWhiteSpaceHelper = true
		}
		if goType.HasPattern {
			needsRegexpHelper = true
		}
	}

	// Write whiteSpace processing helper if needed
	if needsWhiteSpaceHelper {
		builder.WriteString("// applyWhiteSpaceProcessing applies XSD whiteSpace facet processing\n")
		builder.WriteString("func applyWhiteSpaceProcessing(value, whiteSpaceAction string) string {\n")
		builder.WriteString("\tswitch whiteSpaceAction {\n")
		builder.WriteString("\tcase \"replace\":\n")
		builder.WriteString("\t\t// Replace tab, newline, and carriage return with space\n")
		builder.WriteString("\t\tvalue = strings.ReplaceAll(value, \"\\t\", \" \")\n")
		builder.WriteString("\t\tvalue = strings.ReplaceAll(value, \"\\n\", \" \")\n")
		builder.WriteString("\t\tvalue = strings.ReplaceAll(value, \"\\r\", \" \")\n")
		builder.WriteString("\t\treturn value\n")
		builder.WriteString("\tcase \"collapse\":\n")
		builder.WriteString("\t\t// First apply replace processing\n")
		builder.WriteString("\t\tvalue = strings.ReplaceAll(value, \"\\t\", \" \")\n")
		builder.WriteString("\t\tvalue = strings.ReplaceAll(value, \"\\n\", \" \")\n")
		builder.WriteString("\t\tvalue = strings.ReplaceAll(value, \"\\r\", \" \")\n")
		builder.WriteString("\t\t// Then collapse sequences of spaces and trim\n")
		if needsRegexpHelper {
			builder.WriteString("\t\tvalue = regexp.MustCompile(`\\\\s+`).ReplaceAllString(value, \" \")\n")
		} else {
			// Use simpler approach without regexp if not already needed
			builder.WriteString("\t\t// Simplified collapse: just trim spaces\n")
		}
		builder.WriteString("\t\tvalue = strings.TrimSpace(value)\n")
		builder.WriteString("\t\treturn value\n")
		builder.WriteString("\tcase \"preserve\":\n")
		builder.WriteString("\t\tfallthrough\n")
		builder.WriteString("\tdefault:\n")
		builder.WriteString("\t\t// Preserve all whitespace as-is\n")
		builder.WriteString("\t\treturn value\n")
		builder.WriteString("\t}\n")
		builder.WriteString("}\n\n")
	}
}
