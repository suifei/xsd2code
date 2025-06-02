# 快速开始

本指南将帮助您在5分钟内开始使用XSD2Code，快速体验从XSD到代码的转换过程。

## 📋 前提条件

- Go 1.21 或更高版本
- 基本的命令行操作知识
- 一个XSD文件（我们提供示例）

## 🚀 第一步：获取工具

### 方法1：从源码构建

```bash
# 克隆仓库
git clone https://github.com/suifei/xsd2code.git
cd xsd2code

# 构建工具
go build -o xsd2code cmd/main.go
```

### 方法2：下载预构建版本

从 [GitHub Releases](https://github.com/suifei/xsd2code/releases) 下载适合您操作系统的预构建版本。

## 🎯 第二步：准备XSD文件

我们提供了一个简单的示例XSD文件：

```xml
<?xml version="1.0" encoding="UTF-8"?>
<xs:schema xmlns:xs="http://www.w3.org/2001/XMLSchema"
           targetNamespace="http://example.com/simple"
           xmlns:tns="http://example.com/simple"
           elementFormDefault="qualified">

    <!-- 简单的人员信息类型 -->
    <xs:complexType name="PersonType">
        <xs:sequence>
            <xs:element name="name" type="xs:string"/>
            <xs:element name="email" type="xs:string"/>
            <xs:element name="age" type="xs:int" minOccurs="0"/>
        </xs:sequence>
        <xs:attribute name="id" type="xs:string" use="required"/>
    </xs:complexType>

    <!-- 根元素 -->
    <xs:element name="person" type="tns:PersonType"/>

</xs:schema>
```

将上述内容保存为 `person.xsd` 文件。

## ⚡ 第三步：生成代码

### 生成Go代码

```bash
# 基本用法 - 生成Go代码
./xsd2code -xsd=person.xsd

# 指定输出文件
./xsd2code -xsd=person.xsd -output=person.go

# 指定包名
./xsd2code -xsd=person.xsd -output=person.go -package=models
```

**生成的Go代码示例**：

```go
package models

import (
    "encoding/xml"
)

// PersonType represents PersonType
type PersonType struct {
    XMLName xml.Name `xml:"http://example.com/simple PersonType"`
    Name string `xml:"name"`
    Email string `xml:"email"`
    Age *int `xml:"age,omitempty"`
    Id string `xml:"id,attr"`
}
```

### 生成Java代码

```bash
./xsd2code -xsd=person.xsd -lang=java -output=Person.java -package=com.example.models
```

**生成的Java代码示例**：

```java
package com.example.models;

import javax.xml.bind.annotation.*;

@XmlRootElement(name = "person", namespace = "http://example.com/simple")
public class PersonType {
    @XmlElement(name = "name", required = true)
    private String name;
    
    @XmlElement(name = "email", required = true)
    private String email;
    
    @XmlElement(name = "age")
    private Integer age;
    
    @XmlAttribute(name = "id", required = true)
    private String id;
    
    // Getters and setters...
}
```

### 生成C#代码

```bash
./xsd2code -xsd=person.xsd -lang=csharp -output=Person.cs -package=Example.Models
```

## 🎉 第四步：验证结果

### 编译Go代码

```bash
# 创建简单的测试程序
cat > test_person.go << 'EOF'
package main

import (
    "encoding/xml"
    "fmt"
    "./models" // 假设生成的代码在models包中
)

func main() {
    person := models.PersonType{
        Name:  "张三",
        Email: "zhangsan@example.com",
        Age:   &[]int{30}[0],
        Id:    "p001",
    }
    
    data, err := xml.MarshalIndent(person, "", "  ")
    if err != nil {
        panic(err)
    }
    
    fmt.Println(string(data))
}
EOF

# 运行测试
go run test_person.go
```

## 🌟 第五步：探索更多功能

### 添加JSON支持

```bash
./xsd2code -xsd=person.xsd -output=person.go -json
```

生成带JSON标签的代码：

```go
type PersonType struct {
    XMLName xml.Name `xml:"http://example.com/simple PersonType" json:"-"`
    Name string `xml:"name" json:"name"`
    Email string `xml:"email" json:"email"`
    Age *int `xml:"age,omitempty" json:"age,omitempty"`
    Id string `xml:"id,attr" json:"id"`
}
```

### 生成验证代码

```bash
./xsd2code -xsd=person.xsd -validation -validation-output=validation.go
```

### 使用复杂示例

项目提供了更复杂的示例XSD文件：

```bash
# 使用高级示例
./xsd2code -xsd=examples/advanced_example.xsd -output=advanced.go

# 使用带约束的示例
./xsd2code -xsd=examples/test_restrictions.xsd -output=restrictions.go
```

## 🎯 下一步

现在您已经成功生成了第一个代码文件！接下来可以：

1. **深入学习** - 查看 [[基本用法|Basic-Usage]] 了解更多命令选项
2. **探索特性** - 阅读 [[XSD特性支持|XSD-Features]] 了解支持的所有XSD特性
3. **最佳实践** - 查看 [[最佳实践|Best-Practices]] 学习使用建议
4. **高级示例** - 查看 [[高级示例|Advanced-Examples]] 了解复杂场景

## ❓ 遇到问题？

如果在快速开始过程中遇到问题：

1. 检查 [[常见问题|FAQ]] 
2. 查看 [[故障排除|Troubleshooting]] 指南
3. 在 [GitHub Issues](https://github.com/suifei/xsd2code/issues) 提问

## 📝 小贴士

- 建议在实际项目中使用前，先用简单的XSD测试工具
- 生成的代码可以直接在您的项目中使用
- 支持增量更新，XSD修改后重新生成即可
- 所有生成的代码都通过编译测试，可以放心使用

---

🎉 **恭喜！** 您已经成功完成了XSD2Code的快速开始教程！
