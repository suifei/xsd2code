# 多语言支持 - 各编程语言代码生成详解

XSD2Code支持将XSD文件转换为多种编程语言的类型定义。本页面详细介绍每种语言的特性、生成规则和最佳实践。

## 🌍 支持的编程语言

| 语言 | 支持状态 | 特色功能 | 版本要求 |
|------|----------|----------|----------|
| **Go** | ✅ 完全支持 | 结构体、XML/JSON标签、验证 | Go 1.21+ |
| **Java** | ✅ 完全支持 | POJO、JAXB注解、Bean验证 | Java 8+ |
| **C#** | ✅ 完全支持 | 属性、XML序列化、DataAnnotations | .NET 6+ |
| **TypeScript** | ✅ 基础支持 | 接口、类型定义、可选字段 | TypeScript 4.0+ |
| **Python** | 🔄 开发中 | 数据类、类型注解、验证 | Python 3.8+ |

## 🔧 Go语言代码生成

### 基本结构

```bash
# 生成Go代码（默认）
xsd2code -xsd=schema.xsd -output=types.go -package=models
```

### 生成特性

#### 1. 结构体定义
```go
// XSD复杂类型 -> Go结构体
type UserType struct {
    ID       string `xml:"id,attr" json:"id"`
    Name     string `xml:"name" json:"name"`
    Email    string `xml:"email" json:"email"`
    Status   string `xml:"status" json:"status"`
}
```

#### 2. XML和JSON标签
```go
type ProductType struct {
    ID          string  `xml:"id,attr" json:"id"`
    Name        string  `xml:"name" json:"name"`
    Price       float64 `xml:"price" json:"price"`
    Category    string  `xml:"category" json:"category"`
    InStock     bool    `xml:"inStock" json:"inStock"`
}
```

#### 3. 枚举类型
```go
// XSD枚举 -> Go常量
type StatusType string

const (
    StatusTypeActive   StatusType = "active"
    StatusTypeInactive StatusType = "inactive"
    StatusTypePending  StatusType = "pending"
)
```

#### 4. 验证代码
```go
func (u *UserType) Validate() error {
    if len(u.Name) < 1 || len(u.Name) > 100 {
        return fmt.Errorf("name length must be between 1 and 100")
    }
    
    emailPattern := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
    if !emailPattern.MatchString(u.Email) {
        return fmt.Errorf("invalid email format")
    }
    
    return nil
}
```

### Go高级特性

#### 自定义类型
```go
// 基于XSD simpleType的自定义类型
type EmailType string

func (e EmailType) Validate() error {
    pattern := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
    if !pattern.MatchString(string(e)) {
        return fmt.Errorf("invalid email format")
    }
    return nil
}
```

#### 嵌套结构
```go
type AddressType struct {
    Street   string `xml:"street" json:"street"`
    City     string `xml:"city" json:"city"`
    ZipCode  string `xml:"zipCode" json:"zipCode"`
    Country  string `xml:"country" json:"country"`
}

type PersonType struct {
    Name    string       `xml:"name" json:"name"`
    Address *AddressType `xml:"address" json:"address,omitempty"`
}
```

## ☕ Java语言代码生成

### 基本使用

```bash
# 生成Java代码
xsd2code -xsd=schema.xsd -lang=java -output=Types.java -package=com.example.models
```

### 生成特性

#### 1. POJO类定义
```java
package com.example.models;

import javax.xml.bind.annotation.*;
import javax.validation.constraints.*;

@XmlRootElement(name = "user")
@XmlAccessorType(XmlAccessType.FIELD)
public class UserType {
    
    @XmlAttribute(name = "id", required = true)
    @NotNull
    private String id;
    
    @XmlElement(name = "name", required = true)
    @NotBlank
    @Size(min = 1, max = 100)
    private String name;
    
    @XmlElement(name = "email", required = true)
    @Email
    private String email;
    
    // Constructors, getters, setters...
    public UserType() {}
    
    public String getId() { return id; }
    public void setId(String id) { this.id = id; }
    
    public String getName() { return name; }
    public void setName(String name) { this.name = name; }
    
    public String getEmail() { return email; }
    public void setEmail(String email) { this.email = email; }
}
```

#### 2. 枚举类型
```java
@XmlEnum
public enum StatusType {
    @XmlEnumValue("active")
    ACTIVE("active"),
    
    @XmlEnumValue("inactive")
    INACTIVE("inactive"),
    
    @XmlEnumValue("pending")
    PENDING("pending");
    
    private final String value;
    
    StatusType(String value) {
        this.value = value;
    }
    
    public String getValue() {
        return value;
    }
    
    public static StatusType fromValue(String value) {
        for (StatusType status : StatusType.values()) {
            if (status.value.equals(value)) {
                return status;
            }
        }
        throw new IllegalArgumentException("Unknown status: " + value);
    }
}
```

#### 3. Bean验证
```java
import javax.validation.Valid;
import javax.validation.constraints.*;

public class OrderType {
    @NotNull
    @Valid
    private CustomerType customer;
    
    @NotEmpty
    @Valid
    private List<OrderItemType> items;
    
    @DecimalMin("0.00")
    @DecimalMax("999999.99")
    private BigDecimal totalAmount;
    
    @Pattern(regexp = "\\d{4}-\\d{2}-\\d{2}")
    private String orderDate;
}
```

## 🏷️ C#语言代码生成

### 基本使用

```bash
# 生成C#代码
xsd2code -xsd=schema.xsd -lang=csharp -output=Types.cs -package=Example.Models
```

### 生成特性

#### 1. 类定义
```csharp
using System;
using System.Xml.Serialization;
using System.ComponentModel.DataAnnotations;
using Newtonsoft.Json;

namespace Example.Models
{
    [XmlRoot("user")]
    public class UserType
    {
        [XmlAttribute("id")]
        [Required]
        [JsonProperty("id")]
        public string Id { get; set; }
        
        [XmlElement("name")]
        [Required]
        [StringLength(100, MinimumLength = 1)]
        [JsonProperty("name")]
        public string Name { get; set; }
        
        [XmlElement("email")]
        [Required]
        [EmailAddress]
        [JsonProperty("email")]
        public string Email { get; set; }
        
        [XmlElement("status")]
        [JsonProperty("status")]
        public StatusType Status { get; set; }
    }
}
```

#### 2. 枚举类型
```csharp
public enum StatusType
{
    [XmlEnum("active")]
    Active,
    
    [XmlEnum("inactive")]
    Inactive,
    
    [XmlEnum("pending")]
    Pending
}
```

#### 3. 数据注解验证
```csharp
using System.ComponentModel.DataAnnotations;

public class ProductType
{
    [Required]
    [StringLength(50)]
    public string Name { get; set; }
    
    [Range(0.01, 999999.99)]
    public decimal Price { get; set; }
    
    [RegularExpression(@"^[A-Z]{2,3}-\d{4}$")]
    public string ProductCode { get; set; }
}
```

## 📘 TypeScript代码生成

### 基本使用

```bash
# 生成TypeScript代码
xsd2code -xsd=schema.xsd -lang=typescript -output=types.ts
```

### 生成特性

#### 1. 接口定义
```typescript
// 基础接口
export interface UserType {
    id: string;
    name: string;
    email: string;
    status?: StatusType;
    createdAt?: Date;
}

// 可选字段处理
export interface AddressType {
    street: string;
    city: string;
    zipCode?: string;
    country: string;
}
```

#### 2. 枚举类型
```typescript
export enum StatusType {
    Active = "active",
    Inactive = "inactive",
    Pending = "pending"
}

// 字符串字面量类型
export type StatusLiteral = "active" | "inactive" | "pending";
```

#### 3. 复杂类型
```typescript
export interface OrderType {
    id: string;
    customer: CustomerType;
    items: OrderItemType[];
    totalAmount: number;
    orderDate: string;
    status: StatusType;
}

// 泛型支持
export interface ResponseType<T> {
    success: boolean;
    data?: T;
    error?: string;
}
```

#### 4. 类型保护
```typescript
// 类型保护函数
export function isUserType(obj: any): obj is UserType {
    return typeof obj === 'object' &&
           typeof obj.id === 'string' &&
           typeof obj.name === 'string' &&
           typeof obj.email === 'string';
}

// 验证函数
export function validateEmail(email: string): boolean {
    const emailRegex = /^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$/;
    return emailRegex.test(email);
}
```

## 🐍 Python代码生成（开发中）

### 基本使用

```bash
# 生成Python代码（实验性）
xsd2code -xsd=schema.xsd -lang=python -output=types.py
```

### 计划特性

#### 1. 数据类
```python
from dataclasses import dataclass
from typing import Optional, List
from datetime import datetime
import re

@dataclass
class UserType:
    id: str
    name: str
    email: str
    status: Optional[str] = None
    created_at: Optional[datetime] = None
    
    def __post_init__(self):
        if not self.validate_email(self.email):
            raise ValueError(f"Invalid email format: {self.email}")
    
    @staticmethod
    def validate_email(email: str) -> bool:
        pattern = r'^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$'
        return re.match(pattern, email) is not None
```

#### 2. 枚举支持
```python
from enum import Enum

class StatusType(Enum):
    ACTIVE = "active"
    INACTIVE = "inactive"
    PENDING = "pending"
```

## 🎛️ 语言特定配置

### Go配置选项

```bash
# 禁用JSON标签
xsd2code -xsd=schema.xsd -lang=go -json=false

# 自定义包名
xsd2code -xsd=schema.xsd -lang=go -package=mymodels

# 生成验证代码
xsd2code -xsd=schema.xsd -lang=go -validation=true
```

### Java配置选项

```bash
# 指定包名
xsd2code -xsd=schema.xsd -lang=java -package=com.example.dto

# 生成Builder模式
xsd2code -xsd=schema.xsd -lang=java -builder=true

# 启用Bean验证
xsd2code -xsd=schema.xsd -lang=java -validation=true
```

### C#配置选项

```bash
# 指定命名空间
xsd2code -xsd=schema.xsd -lang=csharp -namespace=Example.Models

# 启用JSON序列化
xsd2code -xsd=schema.xsd -lang=csharp -json=true

# 生成数据注解
xsd2code -xsd=schema.xsd -lang=csharp -annotations=true
```

## 🔄 语言特性对比

| 特性 | Go | Java | C# | TypeScript | Python |
|------|----|----- |----|------------|--------|
| **基础类型映射** | ✅ | ✅ | ✅ | ✅ | 🔄 |
| **枚举支持** | ✅ | ✅ | ✅ | ✅ | 🔄 |
| **验证代码** | ✅ | ✅ | ✅ | 🔄 | 🔄 |
| **JSON序列化** | ✅ | ✅ | ✅ | ✅ | 🔄 |
| **XML序列化** | ✅ | ✅ | ✅ | 🔄 | 🔄 |
| **泛型支持** | ✅ | ✅ | ✅ | ✅ | 🔄 |
| **继承扩展** | ✅ | ✅ | ✅ | ✅ | 🔄 |

## 🚀 多语言项目示例

### 混合语言项目

```bash
# 为前端生成TypeScript类型
xsd2code -xsd=api-schema.xsd -lang=typescript -output=frontend/types.ts

# 为后端Go服务生成结构体
xsd2code -xsd=api-schema.xsd -lang=go -output=backend/types.go -package=api

# 为Java微服务生成POJO
xsd2code -xsd=api-schema.xsd -lang=java -output=service/Types.java -package=com.example.api
```

### 批量生成脚本

```bash
#!/bin/bash
# 多语言代码生成脚本

SCHEMA_FILE="schema.xsd"

# Go
xsd2code -xsd=$SCHEMA_FILE -lang=go -output=go/types.go -package=models

# Java
xsd2code -xsd=$SCHEMA_FILE -lang=java -output=java/Types.java -package=com.example.models

# C#
xsd2code -xsd=$SCHEMA_FILE -lang=csharp -output=csharp/Types.cs -namespace=Example.Models

# TypeScript
xsd2code -xsd=$SCHEMA_FILE -lang=typescript -output=typescript/types.ts

echo "All language files generated successfully!"
```

## 🎯 最佳实践

### 1. 命名约定

- **Go**: 使用PascalCase，导出字段大写开头
- **Java**: 使用PascalCase类名，camelCase字段名
- **C#**: 使用PascalCase，属性大写开头
- **TypeScript**: 使用PascalCase接口，camelCase字段

### 2. 包/命名空间组织

```bash
# 按功能模块组织
xsd2code -xsd=user.xsd -package=models.user
xsd2code -xsd=order.xsd -package=models.order
xsd2code -xsd=product.xsd -package=models.product
```

### 3. 验证策略

- **Go**: 实现Validator接口
- **Java**: 使用Bean Validation注解
- **C#**: 使用DataAnnotations
- **TypeScript**: 运行时类型检查

### 4. 序列化优化

```bash
# 优化JSON序列化
xsd2code -xsd=schema.xsd -json=true -omitempty=true

# 优化XML序列化
xsd2code -xsd=schema.xsd -xml-namespace=true
```

---

💡 **提示**: 选择适合项目需求的语言和配置选项，考虑团队技术栈和项目架构。

🔗 **相关页面**: 
- [[基本用法|Basic-Usage]] - 基础命令使用
- [[配置选项|Configuration]] - 高级配置详解
- [[最佳实践|Best-Practices]] - 使用建议
