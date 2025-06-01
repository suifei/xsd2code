package generator

import (
	"fmt"
	"text/template"

	"github.com/suifei/xsd2code/pkg/types"
)

// CodeGeneratorFactory creates language-specific code generators
type CodeGeneratorFactory struct {
	config *GeneratorConfig
}

// NewCodeGeneratorFactory creates a new factory with the given configuration
func NewCodeGeneratorFactory(config *GeneratorConfig) *CodeGeneratorFactory {
	return &CodeGeneratorFactory{
		config: config,
	}
}

// CreateGenerator creates a code generator for the configured language
func (f *CodeGeneratorFactory) CreateGenerator() (*CodeGenerator, error) {
	if err := f.config.Validate(); err != nil {
		return nil, fmt.Errorf("invalid configuration: %w", err)
	}

	return f.config.CreateCodeGenerator(), nil
}

// GenerateCode generates code for the given types using the configured language
func (f *CodeGeneratorFactory) GenerateCode(goTypes []types.GoType) error {
	generator, err := f.CreateGenerator()
	if err != nil {
		return err
	}

	generator.SetGoTypes(goTypes)
	return generator.Generate()
}

// TemplateBasedGenerator provides template-based code generation
type TemplateBasedGenerator struct {
	*CodeGenerator
	structTemplate *template.Template
	enumTemplate   *template.Template
}

// NewTemplateBasedGenerator creates a new template-based generator
func NewTemplateBasedGenerator(config *GeneratorConfig) (*TemplateBasedGenerator, error) {
	generator := config.CreateCodeGenerator()
	mapper := config.CreateLanguageMapper()

	// Parse templates
	structTmpl, err := template.New("struct").Parse(mapper.GetStructTemplate())
	if err != nil {
		return nil, fmt.Errorf("failed to parse struct template: %w", err)
	}

	enumTmpl, err := template.New("enum").Parse(mapper.GetEnumTemplate())
	if err != nil {
		return nil, fmt.Errorf("failed to parse enum template: %w", err)
	}

	return &TemplateBasedGenerator{
		CodeGenerator:  generator,
		structTemplate: structTmpl,
		enumTemplate:   enumTmpl,
	}, nil
}

// GenerateWithTemplates generates code using templates
func (g *TemplateBasedGenerator) GenerateWithTemplates() error {
	// This would use the templates to generate language-specific code
	// Implementation would depend on the specific template requirements
	return g.Generate() // Fallback to original generation for now
}

// TypeRegistry manages type definitions across multiple files/modules
type TypeRegistry struct {
	types          map[string]types.GoType
	dependencies   map[string][]string // Type -> dependent types
	generatedFiles map[string]bool     // Track generated files
}

// NewTypeRegistry creates a new type registry
func NewTypeRegistry() *TypeRegistry {
	return &TypeRegistry{
		types:          make(map[string]types.GoType),
		dependencies:   make(map[string][]string),
		generatedFiles: make(map[string]bool),
	}
}

// RegisterType registers a type in the registry
func (r *TypeRegistry) RegisterType(goType types.GoType) {
	r.types[goType.Name] = goType
}

// GetType retrieves a type from the registry
func (r *TypeRegistry) GetType(name string) (types.GoType, bool) {
	goType, exists := r.types[name]
	return goType, exists
}

// GetAllTypes returns all registered types
func (r *TypeRegistry) GetAllTypes() []types.GoType {
	result := make([]types.GoType, 0, len(r.types))
	for _, goType := range r.types {
		result = append(result, goType)
	}
	return result
}

// AddDependency adds a dependency relationship between types
func (r *TypeRegistry) AddDependency(fromType, toType string) {
	if r.dependencies[fromType] == nil {
		r.dependencies[fromType] = make([]string, 0)
	}
	r.dependencies[fromType] = append(r.dependencies[fromType], toType)
}

// GetDependencies returns the dependencies for a given type
func (r *TypeRegistry) GetDependencies(typeName string) []string {
	return r.dependencies[typeName]
}

// GenerateCode generates code for all registered types using the given configuration
func (r *TypeRegistry) GenerateCode(config *GeneratorConfig) error {
	factory := NewCodeGeneratorFactory(config)
	return factory.GenerateCode(r.GetAllTypes())
}
