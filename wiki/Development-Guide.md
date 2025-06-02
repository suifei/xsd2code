# 开发指南 (Development Guide)

本页面为希望参与XSD2Code项目开发的贡献者提供详细的开发指南，包括开发环境搭建、代码规范、贡献流程等。

## 目录

- [开发环境搭建](#开发环境搭建)
- [项目结构](#项目结构)
- [开发流程](#开发流程)
- [代码规范](#代码规范)
- [测试指南](#测试指南)
- [贡献流程](#贡献流程)
- [发布流程](#发布流程)

## 开发环境搭建

### 系统要求

- **Go**: >= 1.22.3
- **Git**: >= 2.0.0
- **操作系统**: Windows 10+, macOS 10.15+, Linux (Ubuntu 18.04+)
- **编辑器**: VS Code, GoLand, Vim等支持Go开发的编辑器

### 环境准备

#### 1. 安装Go

```bash
# 下载并安装Go 1.22.3+
# 访问 https://golang.org/dl/ 下载对应平台的安装包

# 验证安装
go version
```

#### 2. 克隆项目

```bash
# 克隆主仓库
git clone https://github.com/suifei/xsd2code.git
cd xsd2code

# 如果是fork，添加upstream
git remote add upstream https://github.com/suifei/xsd2code.git
```

#### 3. 初始化开发环境

```bash
# 下载依赖
go mod download

# 验证模块
go mod verify

# 构建项目
go build ./cmd/

# 运行测试
go test ./...
```

#### 4. 开发工具配置

##### VS Code配置

安装推荐扩展：

```json
{
  "recommendations": [
    "golang.go",
    "ms-vscode.vscode-json",
    "redhat.vscode-xml",
    "davidanson.vscode-markdownlint"
  ]
}
```

工作区设置 (`.vscode/settings.json`)：

```json
{
  "go.toolsManagement.checkForUpdates": "local",
  "go.useLanguageServer": true,
  "go.formatTool": "goimports",
  "go.lintTool": "golangci-lint",
  "go.testFlags": ["-v"],
  "files.eol": "\n"
}
```

## 项目结构

```text
xsd2code/
├── cmd/                    # 命令行应用入口
│   └── main.go            # 主程序入口点
├── pkg/                   # 核心包目录
│   ├── generator/         # 代码生成器
│   │   ├── codegen.go    # 代码生成核心逻辑
│   │   ├── config.go     # 生成器配置
│   │   └── factory.go    # 生成器工厂
│   ├── types/            # 类型定义
│   │   └── xsd_types.go  # XSD类型系统
│   ├── validator/        # 验证器
│   │   └── validator.go  # 验证逻辑
│   └── xsdparser/        # XSD解析器
│       ├── parser.go     # 解析器核心
│       └── unified.go    # 统一解析接口
├── test/                 # 测试文件
├── examples/             # 示例文件
├── docs/                 # 文档
├── wiki/                 # GitHub Wiki文档
├── go.mod               # Go模块定义
├── go.sum               # 依赖校验和
├── Makefile             # 构建脚本
└── README.md            # 项目说明
```

### 包职责说明

- **cmd/**: 命令行应用程序入口，处理CLI参数和用户交互
- **pkg/generator/**: 多语言代码生成器，支持Go、Java、C#、Python等
- **pkg/types/**: XSD类型系统定义，映射XML Schema到Go结构体
- **pkg/validator/**: 验证器，包括XSD验证和生成代码验证
- **pkg/xsdparser/**: XSD文件解析器，处理XML Schema解析

## 开发流程

### 1. 创建开发分支

```bash
# 从main分支创建功能分支
git checkout main
git pull upstream main
git checkout -b feature/your-feature-name

# 或创建修复分支
git checkout -b fix/issue-number-description
```

### 2. 开发新功能

#### 添加新语言支持示例

1. **实现LanguageMapper接口**：

```go
// pkg/generator/rust_mapper.go
type RustLanguageMapper struct {
    BaseLanguageMapper
}

func (r *RustLanguageMapper) GetBuiltinTypeMappings() []TypeMapping {
    return []TypeMapping{
        {"xs:string", "String"},
        {"xs:int", "i32"},
        {"xs:boolean", "bool"},
        {"xs:decimal", "f64"},
        {"xs:dateTime", "chrono::DateTime<chrono::Utc>"},
    }
}

func (r *RustLanguageMapper) GetLanguage() TargetLanguage {
    return "rust"
}

func (r *RustLanguageMapper) GetFileExtension() string {
    return ".rs"
}

func (r *RustLanguageMapper) GetStructTemplate() string {
    return `#[derive(Debug, Serialize, Deserialize)]
pub struct {{.Name}} {
{{range .Fields}}    pub {{.Name}}: {{.Type}},
{{end}}}`
}
```

2. **注册到工厂**：

```go
// pkg/generator/factory.go
func (f *DefaultGeneratorFactory) CreateGenerator(language TargetLanguage) (LanguageMapper, error) {
    switch language {
    case LanguageGo:
        return &GoLanguageMapper{}, nil
    case LanguageJava:
        return &JavaLanguageMapper{}, nil
    case LanguageCSharp:
        return &CSharpLanguageMapper{}, nil
    case LanguagePython:
        return &PythonLanguageMapper{}, nil
    case "rust":  // 新增
        return &RustLanguageMapper{}, nil
    default:
        return nil, fmt.Errorf("unsupported language: %s", language)
    }
}
```

3. **添加测试**：

```go
// pkg/generator/rust_mapper_test.go
func TestRustLanguageMapper_GetBuiltinTypeMappings(t *testing.T) {
    mapper := &RustLanguageMapper{}
    mappings := mapper.GetBuiltinTypeMappings()
    
    assert.NotEmpty(t, mappings)
    
    // 验证特定映射
    found := false
    for _, mapping := range mappings {
        if mapping.XSDType == "xs:string" && mapping.TargetType == "String" {
            found = true
            break
        }
    }
    assert.True(t, found, "xs:string should map to String")
}
```

### 3. 本地测试

```bash
# 运行所有测试
go test ./...

# 运行特定包的测试
go test ./pkg/generator/

# 运行带覆盖率的测试
go test -cover ./...

# 运行性能测试
go test -bench=. ./...

# 运行集成测试
go test -tags=integration ./...
```

### 4. 代码检查

```bash
# 格式化代码
go fmt ./...

# 静态分析
go vet ./...

# 使用golangci-lint (需要安装)
golangci-lint run

# 检查模块整洁性
go mod tidy
```

## 代码规范

### Go代码规范

#### 1. 命名规范

```go
// 包名：小写，简短，有意义
package xsdparser

// 常量：PascalCase或ALL_CAPS
const DefaultTimeout = 30 * time.Second
const MAX_RETRY_COUNT = 3

// 变量：camelCase
var defaultConfig = Config{}

// 函数：PascalCase（公共）或camelCase（私有）
func ParseXSDFile(filePath string) (*XSDSchema, error) {}
func parseElement(element xml.Token) error {}

// 类型：PascalCase
type XSDParser struct {
    filePath string
    schema   *XSDSchema
}

// 接口：通常以-er结尾
type LanguageMapper interface {
    GetLanguage() TargetLanguage
}
```

#### 2. 注释规范

```go
// Package xsdparser provides XSD parsing functionality for XML Schema definitions.
// It supports parsing complex XML Schema files and converting them to internal
// type representations that can be used for code generation.
package xsdparser

// XSDParser represents a parser for XML Schema Definition files.
// It maintains parsing state and provides methods for converting XSD files
// into internal type representations.
type XSDParser struct {
    // filePath is the path to the XSD file being parsed
    filePath string
    // schema holds the parsed XSD schema structure
    schema *XSDSchema
}

// ParseXSDFile parses an XSD file and returns the corresponding schema structure.
// It returns an error if the file cannot be read or if the XSD format is invalid.
//
// Example:
//   parser := NewXSDParser("schema.xsd", "output.go", "main")
//   schema, err := parser.ParseXSDFile("example.xsd")
//   if err != nil {
//       log.Fatal(err)
//   }
func (p *XSDParser) ParseXSDFile(filePath string) (*XSDSchema, error) {
    // Implementation...
}
```

#### 3. 错误处理

```go
// 使用标准错误处理模式
func parseComplexType(data []byte) (*XSDComplexType, error) {
    if len(data) == 0 {
        return nil, errors.New("empty data provided")
    }
    
    var ct XSDComplexType
    if err := xml.Unmarshal(data, &ct); err != nil {
        return nil, fmt.Errorf("failed to unmarshal complex type: %w", err)
    }
    
    return &ct, nil
}

// 自定义错误类型
type ValidationError struct {
    Field   string
    Message string
    Code    string
}

func (e *ValidationError) Error() string {
    return fmt.Sprintf("validation error in field %s: %s", e.Field, e.Message)
}
```

#### 4. 结构体设计

```go
// 使用组合而非继承
type BaseLanguageMapper struct {
    language   TargetLanguage
    extensions map[string]string
}

type GoLanguageMapper struct {
    BaseLanguageMapper
    packageName string
    imports     []string
}

// 使用构造函数模式
func NewGoLanguageMapper(packageName string) *GoLanguageMapper {
    return &GoLanguageMapper{
        BaseLanguageMapper: BaseLanguageMapper{
            language: LanguageGo,
            extensions: map[string]string{
                "struct": ".go",
                "test":   "_test.go",
            },
        },
        packageName: packageName,
        imports:     []string{"encoding/xml", "encoding/json"},
    }
}
```

### 文档规范

#### README文档

每个包都应该有清晰的README.md文档，包含：

- 包的用途和功能
- 安装和使用说明
- 代码示例
- API文档链接

#### 代码注释

- 所有公共函数、类型、常量都必须有注释
- 注释应该说明功能、参数、返回值和可能的错误
- 复杂的算法和业务逻辑需要详细注释

## 测试指南

### 测试结构

```text
pkg/
├── generator/
│   ├── codegen.go
│   ├── codegen_test.go      # 单元测试
│   ├── integration_test.go  # 集成测试
│   └── benchmark_test.go    # 性能测试
```

### 单元测试

```go
// pkg/generator/codegen_test.go
func TestGoLanguageMapper_FormatTypeName(t *testing.T) {
    tests := []struct {
        name     string
        input    string
        expected string
    }{
        {
            name:     "simple type",
            input:    "string",
            expected: "String",
        },
        {
            name:     "namespaced type",
            input:    "xs:string",
            expected: "String",
        },
        {
            name:     "complex type",
            input:    "my-complex-type",
            expected: "MyComplexType",
        },
    }

    mapper := &GoLanguageMapper{}
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            result := mapper.FormatTypeName(tt.input)
            assert.Equal(t, tt.expected, result)
        })
    }
}
```

### 集成测试

```go
// pkg/generator/integration_test.go
// +build integration

func TestFullGenerationPipeline(t *testing.T) {
    // 准备测试数据
    xsdContent := `<?xml version="1.0"?>
    <xs:schema xmlns:xs="http://www.w3.org/2001/XMLSchema">
        <xs:complexType name="Person">
            <xs:sequence>
                <xs:element name="name" type="xs:string"/>
                <xs:element name="age" type="xs:int"/>
            </xs:sequence>
        </xs:complexType>
    </xs:schema>`

    // 创建临时文件
    tmpFile, err := ioutil.TempFile("", "test*.xsd")
    require.NoError(t, err)
    defer os.Remove(tmpFile.Name())
    
    _, err = tmpFile.WriteString(xsdContent)
    require.NoError(t, err)
    tmpFile.Close()

    // 运行完整的生成流程
    parser := xsdparser.NewXSDParser(tmpFile.Name(), "", "test")
    schema, err := parser.ParseXSDFile(tmpFile.Name())
    require.NoError(t, err)

    generator := generator.NewGoLanguageMapper("test")
    code, err := generator.GenerateCode(schema)
    require.NoError(t, err)
    
    // 验证生成的代码
    assert.Contains(t, code, "type Person struct")
    assert.Contains(t, code, "Name string")
    assert.Contains(t, code, "Age int")
}
```

### 性能测试

```go
// pkg/generator/benchmark_test.go
func BenchmarkParseXSDFile(b *testing.B) {
    xsdFile := "testdata/large_schema.xsd"
    parser := xsdparser.NewXSDParser(xsdFile, "", "test")
    
    b.ResetTimer()
    for i := 0; i < b.N; i++ {
        _, err := parser.ParseXSDFile(xsdFile)
        if err != nil {
            b.Fatal(err)
        }
    }
}

func BenchmarkGenerateGoCode(b *testing.B) {
    // 准备测试数据
    schema := createTestSchema()
    generator := generator.NewGoLanguageMapper("test")
    
    b.ResetTimer()
    for i := 0; i < b.N; i++ {
        _, err := generator.GenerateCode(schema)
        if err != nil {
            b.Fatal(err)
        }
    }
}
```

### 测试覆盖率

```bash
# 生成覆盖率报告
go test -coverprofile=coverage.out ./...

# 查看覆盖率
go tool cover -func=coverage.out

# 生成HTML覆盖率报告
go tool cover -html=coverage.out -o coverage.html
```

## 贡献流程

### 1. 提交Issue

在开始开发之前，请先提交Issue描述：

- **Bug报告**: 包括复现步骤、期望行为、实际行为
- **功能请求**: 详细描述新功能的需求和用例
- **性能问题**: 提供性能测试数据和分析

### 2. Fork和分支

```bash
# Fork项目到你的GitHub账号
# 然后克隆你的fork
git clone https://github.com/yourusername/xsd2code.git
cd xsd2code

# 添加upstream
git remote add upstream https://github.com/suifei/xsd2code.git

# 创建功能分支
git checkout -b feature/your-feature-name
```

### 3. 开发和测试

- 遵循代码规范
- 编写相应的测试
- 确保所有测试通过
- 更新相关文档

### 4. 提交代码

```bash
# 提交代码
git add .
git commit -m "feat: add Rust language support

- Implement RustLanguageMapper interface
- Add Rust-specific type mappings
- Add comprehensive test coverage
- Update documentation

Closes #123"
```

### 5. 创建Pull Request

- 填写PR模板
- 描述变更内容
- 关联相关Issue
- 确保CI检查通过

### PR检查清单

- [ ] 代码遵循项目规范
- [ ] 添加了适当的测试
- [ ] 所有测试通过
- [ ] 更新了相关文档
- [ ] 没有破坏性变更（或已标注）
- [ ] 提交信息清晰明确

## 发布流程

### 版本号规则

使用语义化版本控制：

- **主版本号**: 不兼容的API变更
- **次版本号**: 向后兼容的功能性新增
- **修订号**: 向后兼容的问题修正

### 发布步骤

1. **更新版本号**

```bash
# 更新版本信息
git tag v1.2.3
git push upstream v1.2.3
```

2. **生成变更日志**

```bash
# 生成CHANGELOG.md
git log --oneline v1.2.2..v1.2.3 > CHANGELOG_DRAFT.md
```

3. **构建和测试**

```bash
# 完整测试
go test ./...

# 构建所有平台
make build-all

# 运行集成测试
make test-integration
```

4. **发布到GitHub Releases**

- 创建Release
- 上传构建产物
- 发布Release Notes

### 持续集成

项目使用GitHub Actions进行CI/CD：

- **代码检查**: golangci-lint, go vet, go fmt
- **测试**: 单元测试、集成测试、性能测试
- **构建**: 多平台交叉编译
- **发布**: 自动发布到GitHub Releases

## 开发工具推荐

### 编辑器插件

- **VS Code**: Go extension, XML tools
- **GoLand**: 内置Go支持
- **Vim**: vim-go插件

### 命令行工具

```bash
# 安装开发工具
go install golang.org/x/tools/cmd/goimports@latest
go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
go install github.com/google/wire/cmd/wire@latest
```

### Makefile命令

```bash
# 常用开发命令
make build          # 构建项目
make test           # 运行测试
make lint           # 代码检查
make clean          # 清理构建产物
make install        # 安装到GOPATH
make release        # 构建发布版本
```

这个开发指南为XSD2Code项目的贡献者提供了完整的开发环境搭建和开发流程指导，确保代码质量和项目的可维护性。
