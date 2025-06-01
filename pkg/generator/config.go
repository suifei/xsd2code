package generator

import "fmt"

// GeneratorConfig holds configuration options for code generation
type GeneratorConfig struct {
	// Basic options
	PackageName     string
	OutputPath      string
	TargetLanguage  TargetLanguage
	JSONCompatible  bool
	IncludeComments bool
	DebugMode       bool

	// Type mapping options
	EnableCustomTypes bool          // Enable PLC/custom type mappings
	CustomMappings    []TypeMapping // User-defined custom mappings

	// Code generation options
	EnableValidation bool // Generate validation code
	EnableTestCode   bool // Generate test code
	StrictMode       bool // Strict XSD compliance

	// Language-specific options
	LanguageSpecific map[string]interface{} // Language-specific configurations
}

// NewGeneratorConfig creates a new generator configuration with defaults
func NewGeneratorConfig() *GeneratorConfig {
	return &GeneratorConfig{
		PackageName:       "xsd2code",
		TargetLanguage:    LanguageGo,
		JSONCompatible:    false,
		IncludeComments:   true,
		DebugMode:         false,
		EnableCustomTypes: false,
		EnableValidation:  false,
		EnableTestCode:    false,
		StrictMode:        false,
		CustomMappings:    make([]TypeMapping, 0),
		LanguageSpecific:  make(map[string]interface{}),
	}
}

// SetLanguage sets the target language and returns the config for chaining
func (c *GeneratorConfig) SetLanguage(lang TargetLanguage) *GeneratorConfig {
	c.TargetLanguage = lang
	return c
}

// SetPackage sets the package name and returns the config for chaining
func (c *GeneratorConfig) SetPackage(pkg string) *GeneratorConfig {
	c.PackageName = pkg
	return c
}

// SetOutput sets the output path and returns the config for chaining
func (c *GeneratorConfig) SetOutput(path string) *GeneratorConfig {
	c.OutputPath = path
	return c
}

// EnableJSON enables JSON compatibility and returns the config for chaining
func (c *GeneratorConfig) EnableJSON() *GeneratorConfig {
	c.JSONCompatible = true
	return c
}

// EnableDebug enables debug mode and returns the config for chaining
func (c *GeneratorConfig) EnableDebug() *GeneratorConfig {
	c.DebugMode = true
	return c
}

// EnablePLCTypes enables PLC/custom type mappings and returns the config for chaining
func (c *GeneratorConfig) EnablePLCTypes() *GeneratorConfig {
	c.EnableCustomTypes = true
	return c
}

// AddCustomMapping adds a custom type mapping and returns the config for chaining
func (c *GeneratorConfig) AddCustomMapping(xsdType, targetType string) *GeneratorConfig {
	c.CustomMappings = append(c.CustomMappings, TypeMapping{
		XSDType:    xsdType,
		TargetType: targetType,
	})
	return c
}

// Validate validates the configuration
func (c *GeneratorConfig) Validate() error {
	if c.PackageName == "" {
		return fmt.Errorf("package name cannot be empty")
	}
	if c.OutputPath == "" {
		return fmt.Errorf("output path cannot be empty")
	}
	// Add more validation as needed
	return nil
}

// CreateLanguageMapper creates the appropriate language mapper based on configuration
func (c *GeneratorConfig) CreateLanguageMapper() LanguageMapper {
	switch c.TargetLanguage {
	case LanguageGo:
		return &GoLanguageMapper{}
	case LanguageJava:
		return &JavaLanguageMapper{}
	case LanguageCSharp:
		return &CSharpLanguageMapper{}
	case LanguagePython:
		return &PythonLanguageMapper{}
	default:
		return &GoLanguageMapper{} // Default fallback
	}
}

// CreateCodeGenerator creates a code generator based on the configuration
func (c *GeneratorConfig) CreateCodeGenerator() *CodeGenerator {
	generator := NewCodeGenerator(c.PackageName, c.OutputPath)

	// Set language mapper
	mapper := c.CreateLanguageMapper()
	generator.SetLanguageMapper(mapper)

	// Apply configuration
	generator.SetJSONCompatible(c.JSONCompatible)
	generator.SetIncludeComments(c.IncludeComments)
	generator.SetDebugMode(c.DebugMode)
	generator.SetEnableCustomTypes(c.EnableCustomTypes)

	// Add custom mappings if any
	if len(c.CustomMappings) > 0 {
		for _, mapping := range c.CustomMappings {
			generator.typeMappings[mapping.XSDType] = mapping.TargetType
		}
	}

	return generator
}
