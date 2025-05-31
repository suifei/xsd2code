# XSD到Go转换工具 v3.0 (统一解析器)

这个工具可以将XSD（XML Schema Definition）文件转换为Go语言的结构体定义。它采用统一解析器架构，自动检测和处理所有XSD特性，支持通用XSD文件转换。

## 主要功能

### 核心特性
- ✅ **统一解析器**: 自动处理所有XSD特性，无需选择解析模式
- ✅ **完整XSD支持**: 复杂类型、简单类型、元素、属性、组引用、扩展、导入
- ✅ **命名空间处理**: 正确生成包含命名空间的XMLName标签
- ✅ **JSON兼容**: 可选生成JSON序列化标签
- ✅ **枚举支持**: 自动转换XSD枚举为Go常量
- ✅ **组引用**: 完整支持XSD组定义和组引用
- ✅ **类型扩展**: 支持complexContent和simpleContent扩展
- ✅ **递归导入**: 自动处理导入的XSD文件

### 生成代码特性
- 标准Go结构体定义
- 完整的XML标签（包含命名空间）
- 可选的JSON兼容标签
- 详细的代码注释
- 类型安全的枚举常量
- 符合Go最佳实践的字段命名

## 安装

```bash
# 构建工具
cd /path/to/xsd2code
go build -o xsd2code.exe
```

## 使用方法

### 基本命令

```bash
# 最简单的用法
./xsd2code.exe -xsd=schema.xsd

# 完整功能示例
./xsd2code.exe -xsd=schema.xsd -output=types.go -package=xsd2code -json -debug

```

### 命令行参数

| 参数 | 描述 | 默认值 | 示例 |
|------|------|---------|------|
| `-xsd` | XSD文件路径（必需） | - | `-xsd=schema.xsd` |
| `-output` | 输出Go文件路径 | 与XSD同名的.go文件 | `-output=types.go` |
| `-package` | Go包名 | `xsd2code` | `-package=api` |
| `-json` | 生成JSON兼容标签 | `false` | `-json` |
| `-debug` | 启用调试模式 | `false` | `-debug` |
| `-strict` | 启用严格模式 | `false` | `-strict` |
| `-comments` | 包含代码注释 | `true` | `-comments=false` |
| `-help` | 显示帮助信息 | - | `-help` |
| `-version` | 显示版本信息 | - | `-version` |

### 使用示例

```bash
# 生成基本Go结构体
./xsd2code.exe -xsd=TC6_XML_V10_B.xsd

# 生成JSON兼容的结构体
./xsd2code.exe -xsd=schema.xsd -json -output=api_types.go -package=api

# 启用调试和严格模式
./xsd2code.exe -xsd=complex_schema.xsd -debug -strict -json

# 测试IEC 61131-3 PLCopen文件
./xsd2code.exe -xsd=test/TC6_XML_V10_B.xsd -output=output/plcopen.go -package=plcopen -json -debug
```

### 在Go代码中使用生成的结构体

```go
package main

import (
    "encoding/xml"
    "encoding/json"
    "fmt"
    "github.com/suifei/xsd2code" // 导入生成的包
)

func main() {
    ...
}
```

## 支持的XSD特性

### 完整特性列表

| XSD特性 | 支持状态 | 说明 |
|---------|----------|------|
| **复杂类型** | ✅ | 转换为Go结构体 |
| **简单类型** | ✅ | 转换为Go基本类型或自定义类型 |
| **元素定义** | ✅ | 转换为结构体字段 |
| **属性定义** | ✅ | 转换为结构体字段（带attr标签） |
| **枚举类型** | ✅ | 转换为Go常量定义 |
| **组引用** | ✅ | 解析组定义并内联到结构体 |
| **类型扩展** | ✅ | 支持complexContent和simpleContent |
| **命名空间** | ✅ | 正确生成XMLName标签 |
| **导入处理** | ✅ | 递归解析导入的XSD文件 |
| **选择元素** | ✅ | 转换为可选字段 |
| **出现次数** | ✅ | minOccurs/maxOccurs转换为指针或数组 |
| **内联复杂类型** | ✅ | 自动提取为独立结构体 |

### 代码生成示例

**输入XSD:**
```xml
<xs:complexType name="Position">
  <xs:attribute name="x" type="xs:double" use="required"/>
  <xs:attribute name="y" type="xs:double" use="required"/>
</xs:complexType>

<xs:simpleType name="EdgeModifierType">
  <xs:restriction base="xs:string">
    <xs:enumeration value="none"/>
    <xs:enumeration value="rising"/>
    <xs:enumeration value="falling"/>
  </xs:restriction>
</xs:simpleType>
```

**生成的Go代码:**
```go
// Position represents position
type Position struct {
    XMLName xml.Name `xml:"position" json:"-"`
    X       float64  `xml:"x,attr" json:"x"`
    Y       float64  `xml:"y,attr" json:"y"`
}

// EdgeModifierType represents edgeModifierType
type EdgeModifierType string

// EdgeModifierType enumeration values
const (
    EDGE_MODIFIER_TYPE_NONE    EdgeModifierType = "none"
    EDGE_MODIFIER_TYPE_RISING  EdgeModifierType = "rising"
    EDGE_MODIFIER_TYPE_FALLING EdgeModifierType = "falling"
)
```

## 项目结构



```

```

## 版本历史


## 测试验证


## 贡献

欢迎提交问题报告和功能请求。如果您想贡献代码，请：

1. Fork 这个项目
2. 创建您的特性分支
3. 提交您的改动
4. 推送到分支
5. 创建一个 Pull Request

## 许可证

本项目采用MIT许可证。详见LICENSE文件。

## 联系信息

如有问题或建议，请通过以下方式联系：
- 创建GitHub Issue
- 发送邮件至项目维护者

---

## 贡献

欢迎提出问题和改进建议！您可以通过以下方式参与贡献：

1. 提交问题或建议
2. 提交代码改进的Pull Request
3. 改进文档和示例

## 许可证

此工具采用MIT许可证。
