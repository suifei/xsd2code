# 扩展开发 (Extensions)

本页面介绍如何为XSD2Code项目开发扩展，包括添加新语言支持、自定义类型映射和扩展功能模块。

## 目录

- [扩展架构概述](#扩展架构概述)
- [添加新语言支持](#添加新语言支持)
- [自定义类型映射](#自定义类型映射)
- [扩展验证器](#扩展验证器)
- [模板系统扩展](#模板系统扩展)
- [插件开发](#插件开发)
- [扩展示例](#扩展示例)

## 扩展架构概述

XSD2Code采用插件化架构设计，支持多种类型的扩展：

### 扩展点

1. **语言映射器 (LanguageMapper)**: 添加新编程语言支持
2. **类型映射 (TypeMapping)**: 自定义XSD到目标语言的类型映射
3. **验证器 (Validator)**: 扩展验证逻辑
4. **模板引擎**: 自定义代码生成模板
5. **后处理器**: 代码生成后的自定义处理

### 核心接口

```go
// 语言映射器接口
type LanguageMapper interface {
    GetBuiltinTypeMappings() []TypeMapping
    GetCustomTypeMappings() []TypeMapping
    GetLanguage() TargetLanguage
    FormatTypeName(typeName string) string
    GetFileExtension() string
    GetImportStatements() []string
    GetStructTemplate() string
    GetEnumTemplate() string
}

// 类型映射
type TypeMapping struct {
    XSDType    string
    TargetType string
}

// 验证器接口
type Validator interface {
    ValidateSchema(schema *types.XSDSchema) []ValidationError
    ValidateGeneration(code string, language TargetLanguage) []ValidationError
    ValidateXMLAgainstXSD(xmlContent string, xsdPath string) error
}
```

## 添加新语言支持

### 步骤1: 实现LanguageMapper接口

以添加Rust语言支持为例：

```go
// pkg/generator/rust_mapper.go
package generator

import (
    "strings"
    "github.com/suifei/xsd2code/pkg/types"
)

type RustLanguageMapper struct {
    BaseLanguageMapper
    packageName string
    useSerde    bool
}

func NewRustLanguageMapper(packageName string) *RustLanguageMapper {
    return &RustLanguageMapper{
        BaseLanguageMapper: BaseLanguageMapper{},
        packageName:        packageName,
        useSerde:          true,
    }
}

func (r *RustLanguageMapper) GetLanguage() TargetLanguage {
    return "rust"
}

func (r *RustLanguageMapper) GetFileExtension() string {
    return ".rs"
}

func (r *RustLanguageMapper) GetBuiltinTypeMappings() []TypeMapping {
    return []TypeMapping{
        {"xs:string", "String", nil, ""},
        {"xs:int", "i32", nil, ""},
        {"xs:long", "i64", nil, ""},
        {"xs:short", "i16", nil, ""},
        {"xs:byte", "i8", nil, ""},
        {"xs:unsignedInt", "u32", nil, ""},
        {"xs:unsignedLong", "u64", nil, ""},
        {"xs:unsignedShort", "u16", nil, ""},
        {"xs:unsignedByte", "u8", nil, ""},
        {"xs:boolean", "bool", nil, ""},
        {"xs:float", "f32", nil, ""},
        {"xs:double", "f64", nil, ""},
        {"xs:decimal", "rust_decimal::Decimal", []string{"rust_decimal"}, ""},
        {"xs:dateTime", "chrono::DateTime<chrono::Utc>", []string{"chrono"}, ""},
        {"xs:date", "chrono::NaiveDate", []string{"chrono"}, ""},
        {"xs:time", "chrono::NaiveTime", []string{"chrono"}, ""},
        {"xs:base64Binary", "Vec<u8>", nil, ""},
        {"xs:hexBinary", "Vec<u8>", nil, ""},
    }
}

func (r *RustLanguageMapper) GetCustomTypeMappings() []TypeMapping {
    return []TypeMapping{
        // PLC特定类型映射
        {"plc:BOOL", "bool", nil, ""},
        {"plc:INT", "i16", nil, ""},
        {"plc:DINT", "i32", nil, ""},
        {"plc:REAL", "f32", nil, ""},
        {"plc:LREAL", "f64", nil, ""},
        {"plc:STRING", "String", nil, ""},
    }
}

func (r *RustLanguageMapper) FormatTypeName(typeName string) string {
    // 移除命名空间前缀
    if colonIndex := strings.LastIndex(typeName, ":"); colonIndex != -1 {
        typeName = typeName[colonIndex+1:]
    }
    
    // 转换为PascalCase
    return r.BaseLanguageMapper.ToPascalCase(typeName)
}

func (r *RustLanguageMapper) GetImportStatements() []string {
    imports := []string{}
    
    if r.useSerde {
        imports = append(imports, "use serde::{Deserialize, Serialize};")
    }
    
    return imports
}

func (r *RustLanguageMapper) GetStructTemplate() string {
    template := `{{range .Imports}}{{.}}
{{end}}
{{if .UseSerde}}#[derive(Debug, Clone, PartialEq, Serialize, Deserialize)]{{else}}#[derive(Debug, Clone, PartialEq)]{{end}}
pub struct {{.Name}} {
{{range .Fields}}    {{if .Documentation}}/// {{.Documentation}}
    {{end}}pub {{.Name}}: {{.Type}},
{{end}}}

impl {{.Name}} {
    pub fn new() -> Self {
        Self {
{{range .Fields}}            {{.Name}}: {{.DefaultValue}},
{{end}}        }
    }
}

impl Default for {{.Name}} {
    fn default() -> Self {
        Self::new()
    }
}`
    
    return template
}

func (r *RustLanguageMapper) GetEnumTemplate() string {
    return `{{range .Imports}}{{.}}
{{end}}
{{if .UseSerde}}#[derive(Debug, Clone, Copy, PartialEq, Eq, Hash, Serialize, Deserialize)]{{else}}#[derive(Debug, Clone, Copy, PartialEq, Eq, Hash)]{{end}}
pub enum {{.Name}} {
{{range .Values}}    {{if .Documentation}}/// {{.Documentation}}
    {{end}}{{.Name}},
{{end}}}

impl {{.Name}} {
    pub fn as_str(&self) -> &'static str {
        match self {
{{range .Values}}            {{$.Name}}::{{.Name}} => "{{.Value}}",
{{end}}        }
    }
    
    pub fn from_str(s: &str) -> Option<Self> {
        match s {
{{range .Values}}            "{{.Value}}" => Some({{$.Name}}::{{.Name}}),
{{end}}            _ => None,
        }
    }
}`
}
```

### 步骤2: 注册到工厂

```go
// pkg/generator/factory.go
func (f *DefaultGeneratorFactory) CreateGenerator(language TargetLanguage) (LanguageMapper, error) {
    switch language {
    case LanguageGo:
        return NewGoLanguageMapper("main"), nil
    case LanguageJava:
        return NewJavaLanguageMapper("com.example"), nil
    case LanguageCSharp:
        return NewCSharpLanguageMapper("Example"), nil
    case LanguagePython:
        return NewPythonLanguageMapper(), nil
    case "rust":  // 新增Rust支持
        return NewRustLanguageMapper("example"), nil
    default:
        return nil, fmt.Errorf("unsupported language: %s", language)
    }
}

func (f *DefaultGeneratorFactory) GetSupportedLanguages() []TargetLanguage {
    return []TargetLanguage{
        LanguageGo,
        LanguageJava,
        LanguageCSharp,
        LanguagePython,
        "rust",  // 新增
    }
}
```

### 步骤3: 添加测试

```go
// pkg/generator/rust_mapper_test.go
package generator

import (
    "testing"
    "github.com/stretchr/testify/assert"
    "github.com/suifei/xsd2code/pkg/types"
)

func TestRustLanguageMapper_GetBuiltinTypeMappings(t *testing.T) {
    mapper := NewRustLanguageMapper("test")
    mappings := mapper.GetBuiltinTypeMappings()
    
    assert.NotEmpty(t, mappings)
    
    // 验证string类型映射
    stringMapping := findMapping(mappings, "xs:string")
    assert.NotNil(t, stringMapping)
    assert.Equal(t, "String", stringMapping.TargetType)
    
    // 验证dateTime类型映射
    dateTimeMapping := findMapping(mappings, "xs:dateTime")
    assert.NotNil(t, dateTimeMapping)
    assert.Equal(t, "chrono::DateTime<chrono::Utc>", dateTimeMapping.TargetType)
    assert.Contains(t, dateTimeMapping.Imports, "chrono")
}

func TestRustLanguageMapper_FormatTypeName(t *testing.T) {
    mapper := NewRustLanguageMapper("test")
    
    tests := []struct {
        input    string
        expected string
    }{
        {"person", "Person"},
        {"user-info", "UserInfo"},
        {"xs:complexType", "ComplexType"},
        {"my_type_name", "MyTypeName"},
    }
    
    for _, test := range tests {
        result := mapper.FormatTypeName(test.input)
        assert.Equal(t, test.expected, result, "Input: %s", test.input)
    }
}

func TestRustLanguageMapper_GenerateStruct(t *testing.T) {
    mapper := NewRustLanguageMapper("test")
    
    // 创建测试数据
    structData := StructTemplateData{
        Name:     "Person",
        UseSerde: true,
        Fields: []FieldData{
            {Name: "name", Type: "String", Documentation: "Person's name"},
            {Name: "age", Type: "u32", Documentation: "Person's age"},
        },
        Imports: []string{"use serde::{Deserialize, Serialize};"},
    }
    
    template := mapper.GetStructTemplate()
    code, err := executeTemplate(template, structData)
    assert.NoError(t, err)
    
    // 验证生成的代码
    assert.Contains(t, code, "#[derive(Debug, Clone, PartialEq, Serialize, Deserialize)]")
    assert.Contains(t, code, "pub struct Person {")
    assert.Contains(t, code, "pub name: String,")
    assert.Contains(t, code, "pub age: u32,")
    assert.Contains(t, code, "/// Person's name")
}

func findMapping(mappings []TypeMapping, xsdType string) *TypeMapping {
    for _, mapping := range mappings {
        if mapping.XSDType == xsdType {
            return &mapping
        }
    }
    return nil
}
```

## 自定义类型映射

### 创建自定义映射器

```go
// pkg/generator/custom_mapper.go
type CustomTypeMapper struct {
    mappings map[string]TypeMapping
}

func NewCustomTypeMapper() *CustomTypeMapper {
    return &CustomTypeMapper{
        mappings: make(map[string]TypeMapping),
    }
}

func (c *CustomTypeMapper) AddMapping(xsdType, targetType string, imports []string) {
    c.mappings[xsdType] = TypeMapping{
        XSDType:    xsdType,
        TargetType: targetType,
        Imports:    imports,
    }
}

func (c *CustomTypeMapper) GetMapping(xsdType string) (TypeMapping, bool) {
    mapping, exists := c.mappings[xsdType]
    return mapping, exists
}

// 领域特定映射示例
func CreatePLCTypeMapper() *CustomTypeMapper {
    mapper := NewCustomTypeMapper()
    
    // PLC标准类型映射
    mapper.AddMapping("plc:BOOL", "bool", nil)
    mapper.AddMapping("plc:SINT", "i8", nil)
    mapper.AddMapping("plc:INT", "i16", nil)
    mapper.AddMapping("plc:DINT", "i32", nil)
    mapper.AddMapping("plc:LINT", "i64", nil)
    mapper.AddMapping("plc:USINT", "u8", nil)
    mapper.AddMapping("plc:UINT", "u16", nil)
    mapper.AddMapping("plc:UDINT", "u32", nil)
    mapper.AddMapping("plc:ULINT", "u64", nil)
    mapper.AddMapping("plc:REAL", "f32", nil)
    mapper.AddMapping("plc:LREAL", "f64", nil)
    mapper.AddMapping("plc:STRING", "String", nil)
    mapper.AddMapping("plc:WSTRING", "String", nil)
    
    // PLC时间类型
    mapper.AddMapping("plc:TIME", "std::time::Duration", []string{"std::time"})
    mapper.AddMapping("plc:DATE", "chrono::NaiveDate", []string{"chrono"})
    mapper.AddMapping("plc:TIME_OF_DAY", "chrono::NaiveTime", []string{"chrono"})
    mapper.AddMapping("plc:DATE_AND_TIME", "chrono::DateTime<chrono::Local>", []string{"chrono"})
    
    return mapper
}
```

### 集成自定义映射器

```go
// 扩展语言映射器以支持自定义映射
type ExtendedLanguageMapper struct {
    base          LanguageMapper
    customMapper  *CustomTypeMapper
}

func NewExtendedLanguageMapper(base LanguageMapper, customMapper *CustomTypeMapper) *ExtendedLanguageMapper {
    return &ExtendedLanguageMapper{
        base:         base,
        customMapper: customMapper,
    }
}

func (e *ExtendedLanguageMapper) GetTypeMappings() []TypeMapping {
    mappings := e.base.GetBuiltinTypeMappings()
    
    // 添加自定义映射
    for _, mapping := range e.customMapper.mappings {
        mappings = append(mappings, mapping)
    }
    
    return mappings
}

func (e *ExtendedLanguageMapper) ResolveType(xsdType string) (string, []string, error) {
    // 首先检查自定义映射
    if customMapping, exists := e.customMapper.GetMapping(xsdType); exists {
        return customMapping.TargetType, customMapping.Imports, nil
    }
    
    // 然后检查内置映射
    for _, mapping := range e.base.GetBuiltinTypeMappings() {
        if mapping.XSDType == xsdType {
            return mapping.TargetType, mapping.Imports, nil
        }
    }
    
    return "", nil, fmt.Errorf("no mapping found for type: %s", xsdType)
}
```

## 扩展验证器

### 自定义验证器

```go
// pkg/validator/custom_validator.go
type CustomValidator struct {
    rules map[string][]ValidationRule
}

type ValidationRule interface {
    Validate(value interface{}) error
    GetErrorMessage() string
}

// 长度验证规则
type LengthValidationRule struct {
    MinLength int
    MaxLength int
}

func (l *LengthValidationRule) Validate(value interface{}) error {
    str, ok := value.(string)
    if !ok {
        return fmt.Errorf("expected string value")
    }
    
    if len(str) < l.MinLength {
        return fmt.Errorf("value too short: minimum length is %d", l.MinLength)
    }
    
    if l.MaxLength > 0 && len(str) > l.MaxLength {
        return fmt.Errorf("value too long: maximum length is %d", l.MaxLength)
    }
    
    return nil
}

func (l *LengthValidationRule) GetErrorMessage() string {
    return fmt.Sprintf("length must be between %d and %d", l.MinLength, l.MaxLength)
}

// 正则表达式验证规则
type PatternValidationRule struct {
    Pattern *regexp.Regexp
    Message string
}

func (p *PatternValidationRule) Validate(value interface{}) error {
    str, ok := value.(string)
    if !ok {
        return fmt.Errorf("expected string value")
    }
    
    if !p.Pattern.MatchString(str) {
        return fmt.Errorf("value does not match required pattern")
    }
    
    return nil
}

func (p *PatternValidationRule) GetErrorMessage() string {
    if p.Message != "" {
        return p.Message
    }
    return "value must match the required pattern"
}

// 数值范围验证规则
type RangeValidationRule struct {
    MinValue float64
    MaxValue float64
}

func (r *RangeValidationRule) Validate(value interface{}) error {
    var num float64
    switch v := value.(type) {
    case int:
        num = float64(v)
    case float32:
        num = float64(v)
    case float64:
        num = v
    default:
        return fmt.Errorf("expected numeric value")
    }
    
    if num < r.MinValue {
        return fmt.Errorf("value too small: minimum is %f", r.MinValue)
    }
    
    if num > r.MaxValue {
        return fmt.Errorf("value too large: maximum is %f", r.MaxValue)
    }
    
    return nil
}

func (r *RangeValidationRule) GetErrorMessage() string {
    return fmt.Sprintf("value must be between %f and %f", r.MinValue, r.MaxValue)
}
```

### 集成验证器

```go
// pkg/validator/validator_registry.go
type ValidatorRegistry struct {
    validators map[string]Validator
}

func NewValidatorRegistry() *ValidatorRegistry {
    return &ValidatorRegistry{
        validators: make(map[string]Validator),
    }
}

func (r *ValidatorRegistry) RegisterValidator(name string, validator Validator) {
    r.validators[name] = validator
}

func (r *ValidatorRegistry) GetValidator(name string) (Validator, bool) {
    validator, exists := r.validators[name]
    return validator, exists
}

func (r *ValidatorRegistry) ValidateAll(data map[string]interface{}) []ValidationError {
    var errors []ValidationError
    
    for fieldName, value := range data {
        if validator, exists := r.GetValidator(fieldName); exists {
            if fieldErrors := validator.ValidateField(fieldName, value); len(fieldErrors) > 0 {
                errors = append(errors, fieldErrors...)
            }
        }
    }
    
    return errors
}
```

## 模板系统扩展

### 自定义模板函数

```go
// pkg/generator/template_functions.go
func GetTemplateFunctions() template.FuncMap {
    return template.FuncMap{
        "toLower":      strings.ToLower,
        "toUpper":      strings.ToUpper,
        "toCamelCase":  toCamelCase,
        "toPascalCase": toPascalCase,
        "toSnakeCase":  toSnakeCase,
        "toKebabCase":  toKebabCase,
        "pluralize":    pluralize,
        "singularize":  singularize,
        "contains":     strings.Contains,
        "hasPrefix":    strings.HasPrefix,
        "hasSuffix":    strings.HasSuffix,
        "join":         strings.Join,
        "split":        strings.Split,
        "replace":      strings.ReplaceAll,
        "trim":         strings.TrimSpace,
        "indent":       indent,
        "comment":      comment,
        "escape":       escapeString,
        "timestamp":    func() string { return time.Now().Format(time.RFC3339) },
        "version":      func() string { return "1.0.0" }, // 或从配置中读取
    }
}

func toCamelCase(s string) string {
    parts := strings.FieldsFunc(s, func(c rune) bool {
        return !unicode.IsLetter(c) && !unicode.IsNumber(c)
    })
    
    if len(parts) == 0 {
        return s
    }
    
    result := strings.ToLower(parts[0])
    for i := 1; i < len(parts); i++ {
        result += strings.Title(strings.ToLower(parts[i]))
    }
    
    return result
}

func toPascalCase(s string) string {
    camelCase := toCamelCase(s)
    if len(camelCase) > 0 {
        return strings.ToUpper(camelCase[:1]) + camelCase[1:]
    }
    return camelCase
}

func toSnakeCase(s string) string {
    var result []rune
    for i, r := range s {
        if unicode.IsUpper(r) && i > 0 {
            result = append(result, '_')
        }
        result = append(result, unicode.ToLower(r))
    }
    return string(result)
}

func indent(spaces int, text string) string {
    prefix := strings.Repeat(" ", spaces)
    lines := strings.Split(text, "\n")
    for i, line := range lines {
        if strings.TrimSpace(line) != "" {
            lines[i] = prefix + line
        }
    }
    return strings.Join(lines, "\n")
}

func comment(language string, text string) string {
    switch language {
    case "go", "java", "rust", "javascript":
        return "// " + text
    case "python":
        return "# " + text
    case "xml":
        return "<!-- " + text + " -->"
    default:
        return text
    }
}
```

### 模板继承系统

```go
// pkg/generator/template_registry.go
type TemplateRegistry struct {
    templates map[string]*template.Template
    functions template.FuncMap
}

func NewTemplateRegistry() *TemplateRegistry {
    return &TemplateRegistry{
        templates: make(map[string]*template.Template),
        functions: GetTemplateFunctions(),
    }
}

func (r *TemplateRegistry) RegisterTemplate(name, content string) error {
    tmpl, err := template.New(name).Funcs(r.functions).Parse(content)
    if err != nil {
        return fmt.Errorf("failed to parse template %s: %w", name, err)
    }
    
    r.templates[name] = tmpl
    return nil
}

func (r *TemplateRegistry) RegisterTemplateFile(name, filePath string) error {
    content, err := ioutil.ReadFile(filePath)
    if err != nil {
        return fmt.Errorf("failed to read template file %s: %w", filePath, err)
    }
    
    return r.RegisterTemplate(name, string(content))
}

func (r *TemplateRegistry) ExecuteTemplate(name string, data interface{}) (string, error) {
    tmpl, exists := r.templates[name]
    if !exists {
        return "", fmt.Errorf("template %s not found", name)
    }
    
    var buf bytes.Buffer
    if err := tmpl.Execute(&buf, data); err != nil {
        return "", fmt.Errorf("failed to execute template %s: %w", name, err)
    }
    
    return buf.String(), nil
}
```

## 插件开发

### 插件接口

```go
// pkg/plugin/plugin.go
type Plugin interface {
    Name() string
    Version() string
    Initialize(config map[string]interface{}) error
    Execute(context *PluginContext) error
    Cleanup() error
}

type PluginContext struct {
    Schema     *types.XSDSchema
    Config     map[string]interface{}
    OutputPath string
    Language   TargetLanguage
    Logger     Logger
}

type PluginRegistry struct {
    plugins map[string]Plugin
}

func NewPluginRegistry() *PluginRegistry {
    return &PluginRegistry{
        plugins: make(map[string]Plugin),
    }
}

func (r *PluginRegistry) RegisterPlugin(plugin Plugin) {
    r.plugins[plugin.Name()] = plugin
}

func (r *PluginRegistry) GetPlugin(name string) (Plugin, bool) {
    plugin, exists := r.plugins[name]
    return plugin, exists
}

func (r *PluginRegistry) ExecutePlugins(context *PluginContext) error {
    for name, plugin := range r.plugins {
        if err := plugin.Execute(context); err != nil {
            return fmt.Errorf("plugin %s failed: %w", name, err)
        }
    }
    return nil
}
```

### 示例插件：文档生成器

```go
// pkg/plugin/doc_generator.go
type DocumentationGeneratorPlugin struct {
    outputFormat string
    includeExamples bool
}

func NewDocumentationGeneratorPlugin() *DocumentationGeneratorPlugin {
    return &DocumentationGeneratorPlugin{
        outputFormat: "markdown",
        includeExamples: true,
    }
}

func (d *DocumentationGeneratorPlugin) Name() string {
    return "documentation-generator"
}

func (d *DocumentationGeneratorPlugin) Version() string {
    return "1.0.0"
}

func (d *DocumentationGeneratorPlugin) Initialize(config map[string]interface{}) error {
    if format, ok := config["output_format"].(string); ok {
        d.outputFormat = format
    }
    
    if includeExamples, ok := config["include_examples"].(bool); ok {
        d.includeExamples = includeExamples
    }
    
    return nil
}

func (d *DocumentationGeneratorPlugin) Execute(context *PluginContext) error {
    docPath := filepath.Join(context.OutputPath, "schema_documentation.md")
    
    content, err := d.generateDocumentation(context.Schema)
    if err != nil {
        return err
    }
    
    return ioutil.WriteFile(docPath, []byte(content), 0644)
}

func (d *DocumentationGeneratorPlugin) generateDocumentation(schema *types.XSDSchema) (string, error) {
    var doc strings.Builder
    
    doc.WriteString("# XML Schema Documentation\n\n")
    doc.WriteString(fmt.Sprintf("Target Namespace: %s\n\n", schema.TargetNamespace))
    
    // 生成复杂类型文档
    if len(schema.ComplexTypes) > 0 {
        doc.WriteString("## Complex Types\n\n")
        for _, ct := range schema.ComplexTypes {
            doc.WriteString(fmt.Sprintf("### %s\n\n", ct.Name))
            if ct.Annotation != nil && ct.Annotation.Documentation != nil {
                doc.WriteString(fmt.Sprintf("%s\n\n", ct.Annotation.Documentation.Text))
            }
            
            if d.includeExamples {
                example := d.generateExample(&ct)
                doc.WriteString("#### Example\n\n")
                doc.WriteString("```xml\n")
                doc.WriteString(example)
                doc.WriteString("\n```\n\n")
            }
        }
    }
    
    // 生成简单类型文档
    if len(schema.SimpleTypes) > 0 {
        doc.WriteString("## Simple Types\n\n")
        for _, st := range schema.SimpleTypes {
            doc.WriteString(fmt.Sprintf("### %s\n\n", st.Name))
            if st.Annotation != nil && st.Annotation.Documentation != nil {
                doc.WriteString(fmt.Sprintf("%s\n\n", st.Annotation.Documentation.Text))
            }
        }
    }
    
    return doc.String(), nil
}

func (d *DocumentationGeneratorPlugin) generateExample(ct *types.XSDComplexType) string {
    // 简化的示例生成逻辑
    return fmt.Sprintf("<%s>\n  <!-- Example content -->\n</%s>", ct.Name, ct.Name)
}

func (d *DocumentationGeneratorPlugin) Cleanup() error {
    return nil
}
```

## 扩展示例

### 完整示例：TypeScript语言支持

```go
// pkg/generator/typescript_mapper.go
type TypeScriptLanguageMapper struct {
    BaseLanguageMapper
    useInterfaces   bool
    generateTypes   bool
    outputDirectory string
}

func NewTypeScriptLanguageMapper(outputDir string) *TypeScriptLanguageMapper {
    return &TypeScriptLanguageMapper{
        BaseLanguageMapper: BaseLanguageMapper{},
        useInterfaces:      true,
        generateTypes:      true,
        outputDirectory:    outputDir,
    }
}

func (ts *TypeScriptLanguageMapper) GetLanguage() TargetLanguage {
    return "typescript"
}

func (ts *TypeScriptLanguageMapper) GetFileExtension() string {
    return ".ts"
}

func (ts *TypeScriptLanguageMapper) GetBuiltinTypeMappings() []TypeMapping {
    return []TypeMapping{
        {"xs:string", "string", nil, ""},
        {"xs:int", "number", nil, ""},
        {"xs:long", "number", nil, ""},
        {"xs:short", "number", nil, ""},
        {"xs:byte", "number", nil, ""},
        {"xs:boolean", "boolean", nil, ""},
        {"xs:float", "number", nil, ""},
        {"xs:double", "number", nil, ""},
        {"xs:decimal", "number", nil, ""},
        {"xs:dateTime", "Date", nil, ""},
        {"xs:date", "Date", nil, ""},
        {"xs:time", "string", nil, ""},
        {"xs:base64Binary", "string", nil, ""},
        {"xs:hexBinary", "string", nil, ""},
    }
}

func (ts *TypeScriptLanguageMapper) GetStructTemplate() string {
    if ts.useInterfaces {
        return `{{range .Imports}}{{.}}
{{end}}
{{if .Documentation}}/**
 * {{.Documentation}}
 */{{end}}
export interface {{.Name}} {
{{range .Fields}}  {{if .Documentation}}/**
   * {{.Documentation}}
   */{{end}}
  {{.Name}}{{if .Optional}}?{{end}}: {{.Type}};
{{end}}}`
    } else {
        return `{{range .Imports}}{{.}}
{{end}}
{{if .Documentation}}/**
 * {{.Documentation}}
 */{{end}}
export class {{.Name}} {
{{range .Fields}}  {{if .Documentation}}/**
   * {{.Documentation}}
   */{{end}}
  public {{.Name}}{{if .Optional}}?{{end}}: {{.Type}};
{{end}}
  
  constructor(data?: Partial<{{.Name}}>) {
    if (data) {
      Object.assign(this, data);
    }
  }
}`
    }
}

func (ts *TypeScriptLanguageMapper) GetEnumTemplate() string {
    return `{{range .Imports}}{{.}}
{{end}}
{{if .Documentation}}/**
 * {{.Documentation}}
 */{{end}}
export enum {{.Name}} {
{{range .Values}}  {{if .Documentation}}/**
   * {{.Documentation}}
   */{{end}}
  {{.Name}} = "{{.Value}}",
{{end}}}`
}

func (ts *TypeScriptLanguageMapper) GetImportStatements() []string {
    return []string{
        // TypeScript通常不需要特殊导入
    }
}
```

### 配置文件扩展

```go
// pkg/config/extension_config.go
type ExtensionConfig struct {
    LanguageMappers map[string]LanguageMapperConfig `yaml:"language_mappers"`
    CustomTypes     map[string]TypeMappingConfig    `yaml:"custom_types"`
    Validators      map[string]ValidatorConfig      `yaml:"validators"`
    Plugins         []PluginConfig                  `yaml:"plugins"`
}

type LanguageMapperConfig struct {
    Package    string                 `yaml:"package"`
    Options    map[string]interface{} `yaml:"options"`
    Templates  map[string]string      `yaml:"templates"`
}

type TypeMappingConfig struct {
    TargetType  string   `yaml:"target_type"`
    Imports     []string `yaml:"imports"`
    Validation  string   `yaml:"validation"`
}

type ValidatorConfig struct {
    Type    string                 `yaml:"type"`
    Options map[string]interface{} `yaml:"options"`
}

type PluginConfig struct {
    Name    string                 `yaml:"name"`
    Version string                 `yaml:"version"`
    Config  map[string]interface{} `yaml:"config"`
}

// 配置文件示例 (extensions.yaml)
/*
language_mappers:
  rust:
    package: "example"
    options:
      use_serde: true
      derive_debug: true
    templates:
      struct: "templates/rust_struct.tmpl"
      enum: "templates/rust_enum.tmpl"

custom_types:
  "plc:BOOL":
    target_type: "bool"
    imports: []
  "plc:STRING":
    target_type: "String"
    imports: []

validators:
  length_validator:
    type: "length"
    options:
      min_length: 1
      max_length: 255

plugins:
  - name: "documentation-generator"
    version: "1.0.0"
    config:
      output_format: "markdown"
      include_examples: true
*/
```

这个扩展开发指南为XSD2Code项目提供了完整的扩展框架，支持添加新语言、自定义类型映射、扩展验证器和开发插件，确保项目具有良好的可扩展性。
