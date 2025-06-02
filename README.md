# XSD2Code - 通用XSD到多语言代码转换工具 v3.1 (增强版)

> 🚀 **一键将XSD转换为多语言代码，智能生成类型安全的数据结构**

XSD2Code 是一个功能强大的命令行工具，专门用于将 XML Schema Definition (XSD) 文件转换为多种编程语言的类型定义和数据结构。该工具采用智能解析技术，自动处理复杂的XSD特性，生成可直接使用的、类型安全的代码。

## 🎯 核心价值

- **🔄 自动化转换**: 一键将复杂XSD转换为生产就绪的代码
- **🌍 多语言支持**: 支持Go、Java、C#、TypeScript等主流语言
- **🎛️ 智能处理**: 自动检测XSD特性，智能生成导入和辅助函数
- **📋 完整约束**: 全面支持XSD约束验证，确保数据完整性
- **⚡ 零配置**: 开箱即用，无需复杂配置

## 🎉 最新更新 (v3.1)

### ✨ 新增功能

- ✅ **智能导入管理**: 动态检测所需导入，避免未使用的导入
- ✅ **XSD限制增强**: 完整支持whiteSpace、length、pattern、fixed value等限制
- ✅ **自动生成辅助函数**: 自动生成`applyWhiteSpaceProcessing`等辅助函数
- ✅ **编译错误零容忍**: 生成的代码保证无编译错误
- ✅ **增强验证支持**: 完整的XSD约束验证代码生成

### 🔧 技术改进

- 🚀 **动态导入检测**: 基于实际使用情况智能添加导入语句
- 🚀 **辅助函数生成**: 按需生成所需的辅助验证函数
- 🚀 **XSD约束全面支持**: whiteSpace(preserve/replace/collapse)、exact length、fixed values
- 🚀 **代码质量保证**: 所有生成的代码通过编译测试

## 主要功能

### 核心特性

- ✅ **多语言支持**: Go、Java、C#、Python代码生成
- ✅ **统一解析器**: 自动处理所有XSD特性，无需选择解析模式
- ✅ **完整XSD支持**: 复杂类型、简单类型、元素、属性、组引用、扩展、导入
- ✅ **智能导入管理**: 动态检测并添加必要的导入语句
- ✅ **XSD约束全面支持**: pattern、length、whiteSpace、fixed value等所有约束
- ✅ **命名空间处理**: 正确生成包含命名空间的XML注解
- ✅ **JSON兼容**: 可选生成JSON序列化标签
- ✅ **枚举支持**: 自动转换XSD枚举为各语言的枚举类型
- ✅ **组引用**: 完整支持XSD组定义和组引用
- ✅ **类型扩展**: 支持complexContent和simpleContent扩展
- ✅ **递归导入**: 自动处理导入的XSD文件
- ✅ **类型映射**: 完整的XSD到目标语言类型映射

### 🆕 XSD约束支持详情

#### WhiteSpace 处理

- `preserve`: 保持所有空白字符
- `replace`: 替换制表符、换行符为空格
- `collapse`: 折叠多个空格为单个空格并去除首尾空格

#### 长度约束

- `length`: 精确长度约束
- `minLength`/`maxLength`: 长度范围约束
- 自动生成长度验证代码

#### 模式匹配

- `pattern`: 正则表达式模式匹配
- 自动生成regexp验证代码
- 智能导入regexp包

#### 固定值约束

- `fixed`: 固定值约束
- 自动生成固定值验证

### 多语言代码生成

- **Go**: 标准结构体定义，XML/JSON标签，类型安全枚举，智能导入
- **Java**: POJO类，JAXB注解，枚举类型，getter/setter方法
- **C#**: 属性类，XML序列化注解，枚举类型，JSON支持
- **Python**: 数据类(dataclass)，类型注解，枚举类型，可选字段支持

## 安装

```bash
# 构建工具
go build -o xsd2code cmd/main.go
```

## 使用方法

### 基本命令

```bash
# Go代码生成（默认）
./xsd2code -xsd=schema.xsd

# Java代码生成
./xsd2code -xsd=schema.xsd -lang=java -output=Types.java -package=com.example.models

# C#代码生成
./xsd2code -xsd=schema.xsd -lang=csharp -output=Types.cs -package=Example.Models -json

# Python代码生成
./xsd2code -xsd=schema.xsd -lang=python -output=types.py -package=models

# 显示类型映射
./xsd2code -xsd=schema.xsd -show-mappings

# 生成验证代码
./xsd2code -xsd=schema.xsd -validation -validation-output=validation.go

# 生成测试代码
./xsd2code -xsd=schema.xsd -tests -test-output=tests.go

# 生成示例XML
./xsd2code -xsd=schema.xsd -sample
```

### 命令行参数

- `-xsd string`: XSD文件路径 (必需)
- `-lang string`: 目标语言 (go, java, csharp, python) (默认: "go")
- `-output string`: 输出文件路径 (可选)
- `-package string`: 包名 (默认: "models")
- `-json`: 生成JSON兼容标签
- `-comments`: 包含注释 (默认: true)
- `-validation`: 生成验证代码
- `-validation-output string`: 验证代码输出路径
- `-tests`: 生成测试代码
- `-test-output string`: 测试代码输出路径
- `-benchmarks`: 生成基准测试
- `-sample`: 生成示例XML
- `-show-mappings`: 显示类型映射
- `-validate string`: 验证XML文件
- `-debug`: 启用调试模式
- `-strict`: 启用严格模式
- `-plc`: 启用PLC类型映射
- `-help`: 显示帮助
- `-version`: 显示版本

## 生成的代码示例

### Go代码示例

```go
package models

import (
    "regexp"
    "strings"
    "encoding/xml"
    "time"
)

// ExactLengthCodeType represents a string with pattern validation
type ExactLengthCodeType string

// Validate validates the ExactLengthCodeType format
func (v ExactLengthCodeType) Validate() bool {
    // Validate against pattern: [A-Z]{5}
    pattern := regexp.MustCompile(`[A-Z]{5}`)
    return pattern.MatchString(string(v))
}

// CollapsedStringType represents a string with whiteSpace processing
type CollapsedStringType string

// Validate validates the CollapsedStringType format
func (v CollapsedStringType) Validate() bool {
    strVal := string(v)
    strVal = applyWhiteSpaceProcessing(strVal, "collapse")
    length := len(strVal)
    return length >= 1 && length <= 50
}

// applyWhiteSpaceProcessing applies XSD whiteSpace facet processing
func applyWhiteSpaceProcessing(value, whiteSpaceAction string) string {
    switch whiteSpaceAction {
    case "replace":
        value = strings.ReplaceAll(value, "\t", " ")
        value = strings.ReplaceAll(value, "\n", " ")
        value = strings.ReplaceAll(value, "\r", " ")
        return value
    case "collapse":
        value = strings.ReplaceAll(value, "\t", " ")
        value = strings.ReplaceAll(value, "\n", " ")
        value = strings.ReplaceAll(value, "\r", " ")
        value = regexp.MustCompile(`\s+`).ReplaceAllString(value, " ")
        value = strings.TrimSpace(value)
        return value
    case "preserve":
        fallthrough
    default:
        return value
    }
}
```

### Java代码示例

```java
@XmlRootElement
public class TestDocument {
    @XmlElement
    private ExactLengthCodeType code;
    
    @XmlElement
    private PercentageType percentage;
    
    @XmlAttribute
    private String id;
    
    // getters and setters...
}
```

### C#代码示例

```csharp
namespace Example.Models
{
    [XmlRoot("TestDocument")]
    public class TestDocument
    {
        [XmlElement("code")]
        public ExactLengthCodeType Code { get; set; }
        
        [XmlElement("percentage")]
        public PercentageType Percentage { get; set; }
        
        [XmlAttribute("id")]
        public string Id { get; set; }
    }
}
```

## 特性详解

### 智能导入管理

工具会根据生成的代码内容智能检测所需的导入语句：

- **regexp**: 当使用pattern验证时自动导入
- **strings**: 当使用whiteSpace处理时自动导入
- **time**: 当使用dateTime类型时自动导入
- **encoding/xml**: 始终导入用于XML序列化

### XSD约束完整支持

#### 字符串约束

- **length**: 精确长度 `<xs:length value="5"/>`
- **minLength/maxLength**: 长度范围 `<xs:minLength value="1"/> <xs:maxLength value="50"/>`
- **pattern**: 正则表达式 `<xs:pattern value="[A-Z]{5}"/>`
- **whiteSpace**: 空白处理 `<xs:whiteSpace value="collapse"/>`

#### 数值约束

- **minInclusive/maxInclusive**: 包含边界
- **minExclusive/maxExclusive**: 排除边界
- **totalDigits**: 总位数限制
- **fractionDigits**: 小数位数限制

#### 其他约束

- **enumeration**: 枚举值
- **fixed**: 固定值

### 验证代码生成

生成的类型包含内置验证方法：

```go
func (v ExactLengthCodeType) Validate() bool {
    pattern := regexp.MustCompile(`[A-Z]{5}`)
    return pattern.MatchString(string(v))
}
```

### 测试代码生成

自动生成单元测试和基准测试：

```bash
./xsd2code -xsd=schema.xsd -tests -benchmarks
```

## 支持的XSD特性

- ✅ **元素 (Elements)**: 基本元素、可选元素、数组元素
- ✅ **属性 (Attributes)**: 必需属性、可选属性、固定值属性
- ✅ **复杂类型 (ComplexType)**: sequence、choice、all、mixed content
- ✅ **简单类型 (SimpleType)**: restriction、enumeration、pattern、length约束
- ✅ **命名空间 (Namespaces)**: targetNamespace、xmlns处理
- ✅ **导入和包含 (Import/Include)**: 外部XSD文件引用
- ✅ **组 (Groups)**: 组定义和组引用
- ✅ **扩展 (Extension)**: complexContent和simpleContent扩展
- ✅ **约束 (Restrictions)**: 所有XSD约束类型
- ✅ **固定值 (Fixed)**: 元素和属性固定值
- ✅ **默认值 (Default)**: 元素和属性默认值

## 错误处理

工具提供详细的错误信息和调试支持：

```bash
# 启用调试模式
./xsd2code -xsd=schema.xsd -debug

# 启用严格模式（严格验证XSD）
./xsd2code -xsd=schema.xsd -strict
```

## 贡献

欢迎提交Issue和Pull Request来改进这个工具。

### 开发环境设置

```bash
# 克隆项目
git clone https://github.com/suifei/xsd2code.git
cd xsd2code

# 安装依赖
go mod tidy

# 运行测试
go test ./...

# 构建
make build
```

## 许可证

MIT License - 详见 [LICENSE](LICENSE) 文件

## 更新日志

### v3.1 (2025-06-02)

- ✨ 新增智能导入管理功能
- ✨ 新增XSD约束全面支持（whiteSpace、length、pattern、fixed）
- ✨ 新增自动辅助函数生成
- 🔧 修复编译错误问题
- 🔧 优化代码生成质量
- 📝 完善文档和示例

### v3.0

- 🎉 重构为统一解析器架构
- ✨ 新增多语言支持（Java、C#、Python）
- ✨ 新增验证代码生成
- ✨ 新增测试代码生成
- 🔧 改进XSD特性支持

## 支持和反馈

如果您遇到问题或有功能建议，请：

1. 查看 [GitHub Issues](https://github.com/suifei/xsd2code/issues)
2. 提交新的Issue
3. 参与讨论和改进

感谢使用 XSD2Code！ 🚀
