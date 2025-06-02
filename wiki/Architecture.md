# XSD2Code 架构设计 (Architecture)

XSD2Code是一个基于Go语言开发的XSD到代码生成工具，采用分层架构设计，提供灵活的多语言代码生成能力。

## 目录
- [总体架构](#总体架构)
- [分层设计](#分层设计)
- [核心组件](#核心组件)
- [数据流](#数据流)
- [扩展机制](#扩展机制)
- [设计原则](#设计原则)

## 总体架构

XSD2Code采用经典的分层架构模式，从下到上分为5个主要层次：

```
┌─────────────────────────────────────────┐
│             CLI Layer                   │  ← 命令行接口层
├─────────────────────────────────────────┤
│           Parser Layer                  │  ← XSD解析层
├─────────────────────────────────────────┤
│       Code Generator Layer              │  ← 代码生成层
├─────────────────────────────────────────┤
│      Language Adapter Layer            │  ← 语言适配层
├─────────────────────────────────────────┤
│           Output Layer                  │  ← 输出处理层
└─────────────────────────────────────────┘
```

## 分层设计

### 1. CLI Layer (命令行接口层)
**位置**: `cmd/main.go`
**职责**: 处理用户输入、参数解析、配置管理

**核心结构**:
```go
type XSDConverterConfig struct {
    InputFile      string
    OutputDir      string
    TargetLanguage string
    PackageName    string
    Namespace      string
    Options        map[string]interface{}
}
```

**功能**:
- 命令行参数解析
- 配置文件处理
- 用户输入验证
- 错误处理和日志记录

### 2. Parser Layer (XSD解析层)
**位置**: `pkg/xsdparser/`
**职责**: XSD文件解析、类型系统构建

**核心接口**:
```go
type XSDParser struct {
    SchemaLocation string
    TargetNamespace string
    Elements       []XSDElement
    Types          []XSDType
    // ... 其他字段
}

func (p *XSDParser) Parse(schemaPath string) (*XSDSchema, error)
func (p *XSDParser) ValidateSchema(schema *XSDSchema) error
```

**功能**:
- XML Schema解析
- 类型定义提取
- 约束规则解析
- 依赖关系构建

### 3. Code Generator Layer (代码生成层)
**位置**: `pkg/generator/`
**职责**: 代码生成逻辑、模板处理

**核心接口**:
```go
type CodeGenerator interface {
    Generate(schema *XSDSchema, config *GeneratorConfig) (*GeneratedCode, error)
    GetSupportedLanguages() []string
    ValidateConfig(config *GeneratorConfig) error
}
```

**功能**:
- 代码结构设计
- 模板引擎集成
- 文件组织规划
- 依赖关系处理

### 4. Language Adapter Layer (语言适配层)
**位置**: `pkg/generator/`
**职责**: 特定语言的代码生成逻辑

**核心接口**:
```go
type LanguageMapper interface {
    GetLanguageName() string
    MapType(xsdType string) string
    GenerateStruct(element *XSDElement) (string, error)
    GenerateValidation(constraints []XSDConstraint) (string, error)
    GetFileExtension() string
    GetImports(schema *XSDSchema) []string
}
```

**支持的语言**:
- Go (目标语言常量: `go`)
- Java (目标语言常量: `java`)
- C# (目标语言常量: `csharp`)
- Python (目标语言常量: `python`)

### 5. Output Layer (输出处理层)
**位置**: `pkg/generator/`
**职责**: 文件输出、格式化、组织

**功能**:
- 文件写入管理
- 代码格式化
- 目录结构创建
- 文件命名规范

## 核心组件

### XSD类型系统 (`pkg/types/`)

**基础类型定义**:
```go
type XSDType struct {
    Name         string
    BaseType     string
    Restrictions []XSDRestriction
    Annotations  []XSDAnnotation
}

type XSDElement struct {
    Name        string
    Type        string
    MinOccurs   int
    MaxOccurs   int  // -1表示unbounded
    Optional    bool
    Default     string
}

type XSDComplexType struct {
    XSDType
    Elements   []XSDElement
    Attributes []XSDAttribute
    Sequence   []XSDElement
    Choice     []XSDElement
}

type XSDSimpleType struct {
    XSDType
    Enumeration []string
    Pattern     string
    MinLength   int
    MaxLength   int
}
```

### 验证系统 (`pkg/validator/`)

**验证接口**:
```go
type Validator interface {
    ValidateSchema(schema *XSDSchema) []ValidationError
    ValidateGeneration(code *GeneratedCode) []ValidationError
}

type SchemaValidator struct {
    Rules []ValidationRule
}

type ValidationError struct {
    Type        string
    Message     string
    Location    string
    Severity    string
}
```

## 数据流

### 1. 输入处理流程
```
用户输入 → CLI解析 → 配置验证 → XSD文件读取
```

### 2. 解析流程
```
XSD文件 → XML解析 → 类型提取 → 关系构建 → Schema对象
```

### 3. 生成流程
```
Schema对象 → 语言映射 → 代码生成 → 模板渲染 → 代码文件
```

### 4. 输出流程
```
代码文件 → 格式化 → 文件写入 → 目录组织 → 完成
```

## 扩展机制

### 1. 语言扩展
通过实现`LanguageMapper`接口添加新语言支持：

```go
type MyLanguageMapper struct{}

func (m *MyLanguageMapper) GetLanguageName() string {
    return "mylang"
}

func (m *MyLanguageMapper) MapType(xsdType string) string {
    // 实现类型映射逻辑
}

// 实现其他接口方法...
```

### 2. 验证扩展
通过实现`Validator`接口添加自定义验证：

```go
type CustomValidator struct{}

func (v *CustomValidator) ValidateSchema(schema *XSDSchema) []ValidationError {
    // 实现自定义验证逻辑
}
```

### 3. 生成器扩展
通过实现`CodeGenerator`接口创建自定义生成器：

```go
type CustomGenerator struct{}

func (g *CustomGenerator) Generate(schema *XSDSchema, config *GeneratorConfig) (*GeneratedCode, error) {
    // 实现自定义生成逻辑
}
```

## 设计原则

### 1. 单一职责原则
每个层次和组件都有明确的单一职责，避免功能耦合。

### 2. 开放封闭原则
通过接口设计支持扩展，核心逻辑对修改封闭。

### 3. 依赖倒置原则
高层模块不依赖低层模块，都依赖于抽象接口。

### 4. 接口隔离原则
使用细粒度接口，避免强迫客户端依赖不需要的方法。

### 5. 配置驱动
通过配置文件和参数控制行为，提高灵活性。

## 性能考虑

### 1. 内存管理
- 大型XSD文件的流式处理
- 及时释放不需要的对象
- 使用对象池减少GC压力

### 2. 并发处理
- 多文件并行解析
- 代码生成任务并发执行
- 合理使用goroutine和channel

### 3. 缓存机制
- XSD解析结果缓存
- 类型映射结果缓存
- 模板编译缓存

## 安全考虑

### 1. 输入验证
- XSD文件路径验证
- 输出目录权限检查
- 文件大小限制

### 2. 代码注入防护
- 模板内容安全过滤
- 生成代码安全检查
- 文件名安全处理

### 3. 资源限制
- 解析超时控制
- 内存使用限制
- 文件数量限制

---

本架构设计支持XSD2Code项目的核心功能需求，同时保持良好的可扩展性和可维护性。更多实现细节请参考 [API参考文档](API-Reference.md) 和 [开发指南](Development-Guide.md)。