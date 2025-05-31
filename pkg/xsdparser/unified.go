package xsdparser

import (
	"github.com/suifei/xsd2code/pkg/generator"
	"github.com/suifei/xsd2code/pkg/types"
)

// UnifiedXSDParser combines parsing and code generation
type UnifiedXSDParser struct {
	parser    *XSDParser
	generator *generator.CodeGenerator
}

// NewUnifiedXSDParser creates a new unified XSD parser
func NewUnifiedXSDParser(xsdPath, outputPath, packageName string) *UnifiedXSDParser {
	parser := NewXSDParser(xsdPath, outputPath, packageName)
	codeGen := generator.NewCodeGenerator(packageName, outputPath)

	return &UnifiedXSDParser{
		parser:    parser,
		generator: codeGen,
	}
}

// SetJSONCompatible enables or disables JSON compatibility
func (u *UnifiedXSDParser) SetJSONCompatible(json bool) {
	u.parser.SetJSONCompatible(json)
	u.generator.SetJSONCompatible(json)
}

// SetDebugMode enables or disables debug mode
func (u *UnifiedXSDParser) SetDebugMode(debug bool) {
	u.parser.SetDebugMode(debug)
	u.generator.SetDebugMode(debug)
}

// SetStrictMode enables or disables strict mode
func (u *UnifiedXSDParser) SetStrictMode(strict bool) {
	u.parser.SetStrictMode(strict)
}

// SetIncludeComments enables or disables comments
func (u *UnifiedXSDParser) SetIncludeComments(comments bool) {
	u.parser.SetIncludeComments(comments)
	u.generator.SetIncludeComments(comments)
}

// Parse parses the XSD file
func (u *UnifiedXSDParser) Parse() error {
	return u.parser.Parse()
}

// GenerateGoCode generates the Go code
func (u *UnifiedXSDParser) GenerateGoCode() error {
	goTypes := u.parser.GetGoTypes()
	u.generator.SetGoTypes(goTypes)
	return u.generator.Generate()
}

// GetGoTypes returns the parsed Go types
func (u *UnifiedXSDParser) GetGoTypes() []types.GoType {
	return u.parser.GetGoTypes()
}

// GetSchema returns the parsed schema
func (u *UnifiedXSDParser) GetSchema() *types.XSDSchema {
	return u.parser.GetSchema()
}
