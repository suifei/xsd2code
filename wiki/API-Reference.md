# API参考 / API Reference

本文档提供XSD2Code项目的完整API参考，基于实际的Go实现。

## 目录

- [核心接口](#核心接口)
- [配置结构](#配置结构)
- [解析器API](#解析器api)
- [代码生成器API](#代码生成器api)
- [类型系统](#类型系统)
- [语言映射接口](#语言映射接口)
- [验证器API](#验证器api)
- [实用工具](#实用工具)
- [错误处理](#错误处理)
- [使用示例](#使用示例)

## 核心接口

### XSDConverterConfig

主要配置结构体，用于控制XSD到代码的转换过程。

```go
type XSDConverterConfig struct {
    // 输入XSD文件路径
    InputFile string `json:"inputFile"`
    
    // 输出目录
    OutputDir string `json:"outputDir"`
    
    // 目标语言 (go, java, csharp, python)
    Language string `json:"language"`
    
    // 包名/命名空间
    PackageName string `json:"packageName"`
    
    // 是否生成验证代码
    GenerateValidation bool `json:"generateValidation"`
    
    // 是否生成JSON标签
    GenerateJSONTags bool `json:"generateJSONTags"`
    
    // 是否生成XML标签
    GenerateXMLTags bool `json:"generateXMLTags"`
    
    // 自定义类型映射
    TypeMappings map[string]string `json:"typeMappings"`
    
    // 模板目录
    TemplateDir string `json:"templateDir"`
    
    // 详细输出
    Verbose bool `json:"verbose"`
}
```

**方法：**

- `Validate() error` - 验证配置的有效性
- `SetDefaults()` - 设置默认值
- `String() string` - 返回配置的字符串表示

### XSDParser

XSD解析器接口，负责解析XSD文件并构建类型定义。

```go
type XSDParser struct {
    config     *XSDConverterConfig
    schema     *Schema
    types      map[string]*XSDType
    namespaces map[string]string
}
```

**方法：**

```go
// 创建新的解析器实例
func NewXSDParser(config *XSDConverterConfig) *XSDParser

// 解析XSD文件
func (p *XSDParser) Parse(filename string) (*Schema, error)

// 解析XSD内容
func (p *XSDParser) ParseContent(content []byte) (*Schema, error)

// 获取所有类型定义
func (p *XSDParser) GetTypes() map[string]*XSDType

// 获取命名空间映射
func (p *XSDParser) GetNamespaces() map[string]string

// 解析复杂类型
func (p *XSDParser) parseComplexType(node *xml.Node) (*ComplexType, error)

// 解析简单类型
func (p *XSDParser) parseSimpleType(node *xml.Node) (*SimpleType, error)

// 解析元素
func (p *XSDParser) parseElement(node *xml.Node) (*Element, error)

// 解析属性
func (p *XSDParser) parseAttribute(node *xml.Node) (*Attribute, error)
```

## 配置结构

### Schema

根模式定义。

```go
type Schema struct {
    XMLName         xml.Name           `xml:"schema"`
    TargetNamespace string             `xml:"targetNamespace,attr"`
    ElementFormDefault string          `xml:"elementFormDefault,attr"`
    AttributeFormDefault string       `xml:"attributeFormDefault,attr"`
    Elements        []Element          `xml:"element"`
    ComplexTypes    []ComplexType      `xml:"complexType"`
    SimpleTypes     []SimpleType       `xml:"simpleType"`
    Attributes      []Attribute        `xml:"attribute"`
    Groups          []Group            `xml:"group"`
    AttributeGroups []AttributeGroup   `xml:"attributeGroup"`
    Imports         []Import           `xml:"import"`
    Includes        []Include          `xml:"include"`
}
```

### Element

XSD元素定义。

```go
type Element struct {
    XMLName     xml.Name     `xml:"element"`
    Name        string       `xml:"name,attr"`
    Type        string       `xml:"type,attr"`
    MinOccurs   string       `xml:"minOccurs,attr"`
    MaxOccurs   string       `xml:"maxOccurs,attr"`
    Default     string       `xml:"default,attr"`
    Fixed       string       `xml:"fixed,attr"`
    Nillable    bool         `xml:"nillable,attr"`
    ComplexType *ComplexType `xml:"complexType"`
    SimpleType  *SimpleType  `xml:"simpleType"`
    Annotation  *Annotation  `xml:"annotation"`
}
```

### ComplexType

复杂类型定义。

```go
type ComplexType struct {
    XMLName        xml.Name        `xml:"complexType"`
    Name           string          `xml:"name,attr"`
    Mixed          bool            `xml:"mixed,attr"`
    Abstract       bool            `xml:"abstract,attr"`
    Sequence       *Sequence       `xml:"sequence"`
    Choice         *Choice         `xml:"choice"`
    All            *All            `xml:"all"`
    ComplexContent *ComplexContent `xml:"complexContent"`
    SimpleContent  *SimpleContent  `xml:"simpleContent"`
    Attributes     []Attribute     `xml:"attribute"`
    AttributeGroups []AttributeGroup `xml:"attributeGroup"`
    Annotation     *Annotation     `xml:"annotation"`
}
```

## 解析器API

### 创建解析器

```go
config := &XSDConverterConfig{
    InputFile:   "schema.xsd",
    Language:    "go",
    PackageName: "models",
    Verbose:     true,
}

parser := NewXSDParser(config)
```

### 解析XSD文件

```go
schema, err := parser.Parse("path/to/schema.xsd")
if err != nil {
    log.Fatal("解析失败:", err)
}
```

### 获取类型信息

```go
types := parser.GetTypes()
for name, xsdType := range types {
    fmt.Printf("类型: %s, 种类: %s\n", name, xsdType.Kind)
}
```

## 代码生成器API

### LanguageMapper接口

语言映射器接口，定义如何将XSD类型映射到目标语言。

```go
type LanguageMapper interface {
    // 映射XSD类型到目标语言类型
    MapType(xsdType string) string
    
    // 映射字段名称
    MapFieldName(name string) string
    
    // 映射方法名称
    MapMethodName(name string) string
    
    // 生成字段标签
    GenerateFieldTags(element *Element) string
    
    // 生成导入语句
    GenerateImports(schema *Schema) []string
    
    // 生成包声明
    GeneratePackageDeclaration(packageName string) string
    
    // 生成验证代码
    GenerateValidation(xsdType *XSDType) string
    
    // 获取文件扩展名
    GetFileExtension() string
    
    // 获取语言名称
    GetLanguageName() string
}
```

### CodeGenerator

代码生成器主结构。

```go
type CodeGenerator struct {
    config  *XSDConverterConfig
    mapper  LanguageMapper
    schema  *Schema
    types   map[string]*XSDType
}
```

**方法：**

```go
// 创建新的代码生成器
func NewCodeGenerator(config *XSDConverterConfig, mapper LanguageMapper) *CodeGenerator

// 生成代码
func (g *CodeGenerator) Generate(schema *Schema) error

// 生成结构体
func (g *CodeGenerator) generateStruct(complexType *ComplexType) (string, error)

// 生成枚举
func (g *CodeGenerator) generateEnum(simpleType *SimpleType) (string, error)

// 生成验证函数
func (g *CodeGenerator) generateValidation(xsdType *XSDType) (string, error)

// 写入文件
func (g *CodeGenerator) writeToFile(filename, content string) error
```

### 语言映射器工厂

```go
// 创建语言映射器
func CreateLanguageMapper(language string) (LanguageMapper, error)

// 支持的语言列表
func GetSupportedLanguages() []string

// 注册自定义语言映射器
func RegisterLanguageMapper(language string, mapper LanguageMapper)
```

## 类型系统

### XSDType

通用XSD类型表示。

```go
type XSDType struct {
    Name        string
    Kind        TypeKind
    BaseType    string
    Namespace   string
    Elements    []*Element
    Attributes  []*Attribute
    Restrictions []Restriction
    Documentation string
}

type TypeKind int

const (
    ComplexTypeKind TypeKind = iota
    SimpleTypeKind
    ElementTypeKind
    AttributeTypeKind
)
```

### TypeMapping

类型映射结构。

```go
type TypeMapping struct {
    XSDType    string
    TargetType string
    IsPointer  bool
    IsArray    bool
    Validation string
    Imports    []string
}
```

## 语言映射接口

### Go语言映射器

```go
type GoLanguageMapper struct {
    config       *XSDConverterConfig
    typeMappings map[string]TypeMapping
}

// XSD类型到Go类型的映射
var DefaultGoTypeMappings = map[string]TypeMapping{
    "string":     {TargetType: "string"},
    "int":        {TargetType: "int"},
    "integer":    {TargetType: "int"},
    "long":       {TargetType: "int64"},
    "short":      {TargetType: "int16"},
    "byte":       {TargetType: "int8"},
    "boolean":    {TargetType: "bool"},
    "decimal":    {TargetType: "float64"},
    "float":      {TargetType: "float32"},
    "double":     {TargetType: "float64"},
    "dateTime":   {TargetType: "time.Time", Imports: []string{"time"}},
    "date":       {TargetType: "time.Time", Imports: []string{"time"}},
    "time":       {TargetType: "time.Time", Imports: []string{"time"}},
    "duration":   {TargetType: "time.Duration", Imports: []string{"time"}},
    "base64Binary": {TargetType: "[]byte"},
    "hexBinary":  {TargetType: "[]byte"},
}
```

### Java语言映射器

```go
type JavaLanguageMapper struct {
    config       *XSDConverterConfig
    typeMappings map[string]TypeMapping
}

var DefaultJavaTypeMappings = map[string]TypeMapping{
    "string":     {TargetType: "String"},
    "int":        {TargetType: "Integer"},
    "integer":    {TargetType: "Integer"},
    "long":       {TargetType: "Long"},
    "short":      {TargetType: "Short"},
    "byte":       {TargetType: "Byte"},
    "boolean":    {TargetType: "Boolean"},
    "decimal":    {TargetType: "BigDecimal", Imports: []string{"java.math.BigDecimal"}},
    "float":      {TargetType: "Float"},
    "double":     {TargetType: "Double"},
    "dateTime":   {TargetType: "LocalDateTime", Imports: []string{"java.time.LocalDateTime"}},
    "date":       {TargetType: "LocalDate", Imports: []string{"java.time.LocalDate"}},
    "time":       {TargetType: "LocalTime", Imports: []string{"java.time.LocalTime"}},
}
```

## 验证器API

### ValidationRule

验证规则定义。

```go
type ValidationRule struct {
    Field       string
    Type        ValidationType
    Value       interface{}
    Message     string
    Conditional bool
}

type ValidationType int

const (
    RequiredValidation ValidationType = iota
    MinLengthValidation
    MaxLengthValidation
    PatternValidation
    MinValueValidation
    MaxValueValidation
    EnumValidation
)
```

### Validator

验证器接口。

```go
type Validator interface {
    Validate(data interface{}) []ValidationError
    AddRule(rule ValidationRule)
    RemoveRule(field string, ruleType ValidationType)
}

type ValidationError struct {
    Field   string
    Message string
    Value   interface{}
}
```

## 实用工具

### 文件操作

```go
// 确保目录存在
func EnsureDir(dir string) error

// 写入文件
func WriteFile(filename string, content []byte) error

// 读取文件
func ReadFile(filename string) ([]byte, error)

// 获取文件信息
func GetFileInfo(filename string) (os.FileInfo, error)
```

### 字符串工具

```go
// 转换为驼峰命名
func ToCamelCase(s string) string

// 转换为帕斯卡命名
func ToPascalCase(s string) string

// 转换为蛇形命名
func ToSnakeCase(s string) string

// 清理标识符
func CleanIdentifier(s string) string

// 复数化
func Pluralize(s string) string

// 单数化
func Singularize(s string) string
```

### 模板工具

```go
// 模板数据
type TemplateData struct {
    Schema      *Schema
    Types       map[string]*XSDType
    Config      *XSDConverterConfig
    Imports     []string
    PackageName string
}

// 渲染模板
func RenderTemplate(templatePath string, data TemplateData) (string, error)

// 注册模板函数
func RegisterTemplateFunc(name string, fn interface{})
```

## 错误处理

### 错误类型

```go
// 解析错误
type ParseError struct {
    File    string
    Line    int
    Column  int
    Message string
}

func (e ParseError) Error() string

// 生成错误
type GenerationError struct {
    Type    string
    Field   string
    Message string
}

func (e GenerationError) Error() string

// 配置错误
type ConfigError struct {
    Field   string
    Value   interface{}
    Message string
}

func (e ConfigError) Error() string

// 验证错误
type ValidationError struct {
    Field   string
    Value   interface{}
    Message string
}

func (e ValidationError) Error() string
```

### 错误处理函数

```go
// 包装错误
func WrapError(err error, message string) error

// 检查是否为特定错误类型
func IsParseError(err error) bool
func IsGenerationError(err error) bool
func IsConfigError(err error) bool
func IsValidationError(err error) bool

// 错误收集器
type ErrorCollector struct {
    errors []error
}

func (ec *ErrorCollector) Add(err error)
func (ec *ErrorCollector) HasErrors() bool
func (ec *ErrorCollector) Errors() []error
func (ec *ErrorCollector) Error() string
```

## 使用示例

### 基本使用

```go
package main

import (
    "log"
    "github.com/example/xsd2code/pkg/xsdparser"
    "github.com/example/xsd2code/pkg/generator"
)

func main() {
    // 创建配置
    config := &XSDConverterConfig{
        InputFile:           "schema.xsd",
        OutputDir:          "./generated",
        Language:           "go",
        PackageName:        "models",
        GenerateValidation: true,
        GenerateJSONTags:   true,
        GenerateXMLTags:    true,
        Verbose:            true,
    }
    
    // 解析XSD
    parser := xsdparser.NewXSDParser(config)
    schema, err := parser.Parse(config.InputFile)
    if err != nil {
        log.Fatal("解析失败:", err)
    }
    
    // 生成代码
    mapper, err := generator.CreateLanguageMapper(config.Language)
    if err != nil {
        log.Fatal("创建语言映射器失败:", err)
    }
    
    codeGen := generator.NewCodeGenerator(config, mapper)
    err = codeGen.Generate(schema)
    if err != nil {
        log.Fatal("代码生成失败:", err)
    }
    
    log.Println("代码生成完成!")
}
```

### 自定义语言映射器

```go
package main

import (
    "github.com/example/xsd2code/pkg/generator"
    "github.com/example/xsd2code/pkg/types"
)

type RustLanguageMapper struct {
    config *XSDConverterConfig
}

func (r *RustLanguageMapper) MapType(xsdType string) string {
    mappings := map[string]string{
        "string":   "String",
        "int":      "i32",
        "long":     "i64",
        "boolean":  "bool",
        "decimal":  "f64",
        "dateTime": "chrono::DateTime<chrono::Utc>",
    }
    
    if rustType, exists := mappings[xsdType]; exists {
        return rustType
    }
    return "String" // 默认类型
}

func (r *RustLanguageMapper) MapFieldName(name string) string {
    return ToSnakeCase(name)
}

func (r *RustLanguageMapper) GenerateFieldTags(element *types.Element) string {
    tags := []string{}
    
    if element.Name != "" {
        tags = append(tags, fmt.Sprintf(`serde(rename = "%s")`, element.Name))
    }
    
    if len(tags) > 0 {
        return "#[" + strings.Join(tags, ", ") + "]"
    }
    return ""
}

func (r *RustLanguageMapper) GetFileExtension() string {
    return ".rs"
}

func (r *RustLanguageMapper) GetLanguageName() string {
    return "rust"
}

// 注册自定义映射器
func init() {
    generator.RegisterLanguageMapper("rust", &RustLanguageMapper{})
}
```

### 自定义模板

```go
// 自定义Go结构体模板
const customGoTemplate = `
{{.PackageDeclaration}}

{{range .Imports}}
import "{{.}}"
{{end}}

{{range .Types}}
// {{.Documentation}}
type {{.Name}} struct {
{{range .Elements}}
    {{.MappedName}} {{.MappedType}} ` + "`" + `{{.Tags}}` + "`" + `{{if .Documentation}} // {{.Documentation}}{{end}}
{{end}}
{{range .Attributes}}
    {{.MappedName}} {{.MappedType}} ` + "`" + `{{.Tags}}` + "`" + `{{if .Documentation}} // {{.Documentation}}{{end}}
{{end}}
}

{{if .GenerateValidation}}
// Validate validates the {{.Name}} struct
func (v *{{.Name}}) Validate() error {
    {{.ValidationCode}}
    return nil
}
{{end}}

{{end}}
`

// 使用自定义模板
func useCustomTemplate() {
    templateData := TemplateData{
        Schema:      schema,
        Types:       types,
        Config:      config,
        Imports:     imports,
        PackageName: config.PackageName,
    }
    
    content, err := RenderTemplate(customGoTemplate, templateData)
    if err != nil {
        log.Fatal("模板渲染失败:", err)
    }
    
    err = WriteFile("generated/models.go", []byte(content))
    if err != nil {
        log.Fatal("文件写入失败:", err)
    }
}
```

## 配置选项详解

### TypeMappings

自定义类型映射允许覆盖默认的类型转换：

```go
config.TypeMappings = map[string]string{
    "xs:decimal": "github.com/shopspring/decimal.Decimal",
    "xs:dateTime": "time.Time",
    "customType": "MyCustomType",
}
```

### 模板自定义

```go
config.TemplateDir = "./templates"
// 模板文件结构:
// templates/
//   go/
//     struct.tmpl
//     enum.tmpl
//     validation.tmpl
//   java/
//     class.tmpl
//     enum.tmpl
```

## 性能优化

### 内存管理

```go
// 使用内存池减少分配
var xmlNodePool = sync.Pool{
    New: func() interface{} {
        return &xml.Node{}
    },
}

// 流式处理大型XSD文件
func (p *XSDParser) ParseStream(reader io.Reader) error {
    decoder := xml.NewDecoder(reader)
    decoder.CharsetReader = charset.NewReaderLabel
    
    for {
        token, err := decoder.Token()
        if err == io.EOF {
            break
        }
        if err != nil {
            return err
        }
        
        // 处理token
        p.processToken(token)
    }
    
    return nil
}
```

### 并发处理

```go
// 并发生成多个文件
func (g *CodeGenerator) GenerateConcurrent(schema *Schema) error {
    var wg sync.WaitGroup
    errChan := make(chan error, len(schema.ComplexTypes))
    
    for _, complexType := range schema.ComplexTypes {
        wg.Add(1)
        go func(ct ComplexType) {
            defer wg.Done()
            if err := g.generateStructFile(&ct); err != nil {
                errChan <- err
            }
        }(complexType)
    }
    
    wg.Wait()
    close(errChan)
    
    // 检查错误
    for err := range errChan {
        if err != nil {
            return err
        }
    }
    
    return nil
}
```

这个API参考文档涵盖了XSD2Code项目的完整Go API，包括所有主要接口、结构体、方法和使用示例。开发人员可以使用这个文档来理解和使用项目的程序化接口。
