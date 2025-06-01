# XSD2Code - 通用XSD到多语言代码转换工具 v3.1

这个工具可以将XSD（XML Schema Definition）文件转换为多种编程语言的类型定义。支持Go、Java、C#等主流编程语言，采用统一解析器架构，自动检测和处理所有XSD特性。

## 主要功能

### 核心特性

- ✅ **多语言支持**: Go、Java、C#代码生成
- ✅ **统一解析器**: 自动处理所有XSD特性，无需选择解析模式
- ✅ **完整XSD支持**: 复杂类型、简单类型、元素、属性、组引用、扩展、导入
- ✅ **命名空间处理**: 正确生成包含命名空间的XML注解
- ✅ **JSON兼容**: 可选生成JSON序列化标签
- ✅ **枚举支持**: 自动转换XSD枚举为各语言的枚举类型
- ✅ **组引用**: 完整支持XSD组定义和组引用
- ✅ **类型扩展**: 支持complexContent和simpleContent扩展
- ✅ **递归导入**: 自动处理导入的XSD文件
- ✅ **类型映射**: 完整的XSD到目标语言类型映射

### 多语言代码生成

- **Go**: 标准结构体定义，XML/JSON标签，类型安全枚举
- **Java**: POJO类，JAXB注解，枚举类型，getter/setter方法
- **C#**: 属性类，XML序列化注解，枚举类型，JSON支持

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

# 显示类型映射
./xsd2code -xsd=schema.xsd -lang=java -show-mappings
```

### 命令行参数

| 参数 | 描述 | 默认值 | 示例 |
|------|------|---------|------|
| `-xsd` | XSD文件路径（必需） | - | `-xsd=schema.xsd` |
| `-lang` | 目标语言 | `go` | `-lang=java` `-lang=csharp` |
| `-output` | 输出文件路径 | 根据语言自动生成 | `-output=Types.java` |
| `-package` | 包名/命名空间 | `models` | `-package=com.example.models` |
| `-json` | 启用JSON兼容标签 | false | `-json` |
| `-debug` | 启用调试模式 | false | `-debug` |
| `-show-mappings` | 显示类型映射表 | false | `-show-mappings` |

### 支持的语言

| 语言 | 值 | 文件扩展名 | 包声明格式 |
|------|-----|-----------|-----------|
| Go | `go` | `.go` | `package name` |
| Java | `java` | `.java` | `package com.example.name;` |
| C# | `csharp` | `.cs` | `namespace Example.Name` |

## 代码生成示例

### 输入XSD

```xml
<?xml version="1.0" encoding="UTF-8"?>
<xs:schema xmlns:xs="http://www.w3.org/2001/XMLSchema"
           targetNamespace="http://example.com/person"
           elementFormDefault="qualified">

    <xs:simpleType name="StatusType">
        <xs:restriction base="xs:string">
            <xs:enumeration value="active"/>
            <xs:enumeration value="inactive"/>
        </xs:restriction>
    </xs:simpleType>

    <xs:complexType name="PersonType">
        <xs:sequence>
            <xs:element name="name" type="xs:string"/>
            <xs:element name="age" type="xs:int"/>
            <xs:element name="status" type="tns:StatusType"/>
        </xs:sequence>
        <xs:attribute name="id" type="xs:string" use="required"/>
    </xs:complexType>

    <xs:element name="person" type="tns:PersonType"/>
</xs:schema>
```

### Go输出

```go
package models

import (
    "encoding/xml"
    "time"
)

// PersonType represents PersonType
type PersonType struct {
    XMLName xml.Name   `xml:"http://example.com/person PersonType" json:"-"`
    Name    string     `xml:"name" json:"name"`
    Age     int32      `xml:"age" json:"age"`
    Status  StatusType `xml:"status" json:"status"`
    Id      string     `xml:"id,attr" json:"id"`
}

// StatusType represents
type StatusType string

// StatusType enumeration values
const (
    StatusTypeActive   StatusType = "active"
    StatusTypeInactive StatusType = "inactive"
)
```

### Java输出

```java
package com.example.models;

import java.util.*;
import javax.xml.bind.annotation.*;

@XmlRootElement(name = "PersonType")
@XmlType(namespace = "http://example.com/person")
public class PersonType {
    @XmlElement
    private String name;
    @XmlElement
    private Integer age;
    @XmlElement
    private StatusType status;
    @XmlAttribute
    private String id;

    // Getters and setters...
    public String getName() { return name; }
    public void setName(String name) { this.name = name; }
    // ... other getters/setters
}

public enum StatusType {
    ACTIVE("active"),
    INACTIVE("inactive");

    private final String value;
    StatusType(String value) { this.value = value; }
    public String getValue() { return value; }
}
```

### C#输出

```csharp
using System;
using System.Collections.Generic;
using System.Xml.Serialization;
using System.Text.Json.Serialization;

namespace Example.Models
{
    [XmlRoot("PersonType", Namespace = "http://example.com/person")]
    public class PersonType
    {
        [XmlElement]
        [JsonPropertyName("name")]
        public string Name { get; set; }

        [XmlElement]
        [JsonPropertyName("age")]
        public int Age { get; set; }

        [XmlElement]
        [JsonPropertyName("status")]
        public StatusType Status { get; set; }

        [XmlAttribute]
        [JsonPropertyName("id")]
        public string Id { get; set; }
    }

    public enum StatusType
    {
        Active,
        Inactive,
    }
}
```

## 类型映射

### XSD到Go类型映射

| XSD类型 | Go类型 | 说明 |
|---------|--------|------|
| `xs:string` | `string` | 字符串 |
| `xs:int` | `int32` | 32位整数 |
| `xs:long` | `int64` | 64位整数 |
| `xs:boolean` | `bool` | 布尔值 |
| `xs:decimal` | `float64` | 十进制数 |
| `xs:dateTime` | `time.Time` | 日期时间 |
| `xs:date` | `string` | 日期字符串 |

### XSD到Java类型映射

| XSD类型 | Java类型 | 说明 |
|---------|----------|------|
| `xs:string` | `String` | 字符串 |
| `xs:int` | `Integer` | 整数包装类 |
| `xs:long` | `Long` | 长整数包装类 |
| `xs:boolean` | `Boolean` | 布尔包装类 |
| `xs:decimal` | `BigDecimal` | 精确十进制数 |
| `xs:dateTime` | `LocalDateTime` | 本地日期时间 |
| `xs:date` | `LocalDate` | 本地日期 |

### XSD到C#类型映射

| XSD类型 | C#类型 | 说明 |
|---------|--------|------|
| `xs:string` | `string` | 字符串 |
| `xs:int` | `int` | 32位整数 |
| `xs:long` | `long` | 64位整数 |
| `xs:boolean` | `bool` | 布尔值 |
| `xs:decimal` | `decimal` | 十进制数 |
| `xs:dateTime` | `DateTime` | 日期时间 |
| `xs:date` | `DateTime` | 日期时间 |

## 高级功能

### PLC类型支持

工具支持PLC Open IEC 61131-3标准的数据类型：

```bash
# 启用PLC类型映射
./xsd2code -xsd=plc_schema.xsd -lang=go -custom-types
```

### 验证代码生成

```bash
# 生成验证代码
./xsd2code -xsd=schema.xsd -validation -validation-output=validation.go
```

### 测试代码生成

```bash
# 生成测试代码
./xsd2code -xsd=schema.xsd -tests -test-output=types_test.go
```

## 演示

运行多语言演示脚本：

```bash
# 运行演示脚本
./examples/multilang_demo.sh
```

这将生成Go、Java和C#的示例代码，展示工具的多语言支持能力。

## 项目结构

```
xsd2code/
├── cmd/
│   └── main.go                 # 主程序入口
├── pkg/
│   ├── generator/             # 代码生成器
│   │   ├── codegen.go        # 多语言代码生成逻辑
│   │   ├── config.go         # 配置管理
│   │   └── factory.go        # 生成器工厂
│   ├── types/                # 类型定义
│   ├── validator/            # XML验证器
│   └── xsdparser/           # XSD解析器
├── examples/                 # 示例文件
├── test/                    # 测试文件
└── docs/                    # 文档
```

## 开发

### 构建

```bash
go build -o xsd2code cmd/main.go
```

### 测试

```bash
go test ./...
```

### 贡献

1. Fork项目
2. 创建功能分支
3. 提交更改
4. 创建Pull Request

## 许可证

MIT License - 详见 [LICENSE](LICENSE) 文件

## 更新日志

### v3.1 (2025-06-01)

- ✅ 新增多语言代码生成支持（Go、Java、C#）
- ✅ 重构代码生成架构
- ✅ 优化类型映射系统
- ✅ 改进包名验证逻辑
- ✅ 新增类型映射显示功能
- ✅ 完善枚举类型生成
- ✅ 增强JSON兼容性支持

### v3.0

- ✅ 统一解析器架构
- ✅ 完整XSD特性支持
- ✅ 命名空间处理
- ✅ 组引用支持
- ✅ 类型扩展支持

## 支持

如有问题或建议，请提交[Issue](https://github.com/suifei/xsd2code/issues)
