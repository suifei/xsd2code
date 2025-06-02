# 基础示例 - 简单XSD转换示例

本页面提供了一系列从简单到复杂的XSD转换示例，帮助您快速掌握XSD2Code的基本用法和核心特性。

## 🌟 快速入门示例

### 示例1：最简单的类型定义

#### XSD定义
```xml
<?xml version="1.0" encoding="UTF-8"?>
<xs:schema xmlns:xs="http://www.w3.org/2001/XMLSchema">
    
    <xs:element name="message" type="xs:string"/>
    
</xs:schema>
```

#### 生成命令
```bash
xsd2code -xsd=simple.xsd -output=simple.go
```

#### 生成的Go代码
```go
package models

// MessageType represents the message element
type MessageType string
```

#### 生成的Java代码
```bash
xsd2code -xsd=simple.xsd -lang=java -output=Simple.java -package=com.example
```

```java
package com.example;

import javax.xml.bind.annotation.*;

@XmlRootElement(name = "message")
public class MessageType {
    @XmlValue
    private String value;
    
    public String getValue() { return value; }
    public void setValue(String value) { this.value = value; }
}
```

## 📋 基本类型示例

### 示例2：简单复杂类型

#### XSD定义
```xml
<?xml version="1.0" encoding="UTF-8"?>
<xs:schema xmlns:xs="http://www.w3.org/2001/XMLSchema">
    
    <xs:complexType name="PersonType">
        <xs:sequence>
            <xs:element name="name" type="xs:string"/>
            <xs:element name="age" type="xs:int"/>
            <xs:element name="email" type="xs:string"/>
        </xs:sequence>
    </xs:complexType>
    
    <xs:element name="person" type="PersonType"/>
    
</xs:schema>
```

#### 生成命令
```bash
xsd2code -xsd=person.xsd -output=person.go -package=models
```

#### 生成的Go代码
```go
package models

import "encoding/xml"

// PersonType represents a person with basic information
type PersonType struct {
    Name  string `xml:"name" json:"name"`
    Age   int    `xml:"age" json:"age"`
    Email string `xml:"email" json:"email"`
}

// Person represents the root element
type Person PersonType
```

#### 使用示例
```go
package main

import (
    "encoding/json"
    "encoding/xml"
    "fmt"
    "myapp/models"
)

func main() {
    // 创建Person实例
    person := models.PersonType{
        Name:  "张三",
        Age:   30,
        Email: "zhang.san@example.com",
    }
    
    // JSON序列化
    jsonData, _ := json.Marshal(person)
    fmt.Printf("JSON: %s\n", jsonData)
    
    // XML序列化
    xmlData, _ := xml.Marshal(person)
    fmt.Printf("XML: %s\n", xmlData)
}
```

### 示例3：带属性的复杂类型

#### XSD定义
```xml
<?xml version="1.0" encoding="UTF-8"?>
<xs:schema xmlns:xs="http://www.w3.org/2001/XMLSchema">
    
    <xs:complexType name="BookType">
        <xs:sequence>
            <xs:element name="title" type="xs:string"/>
            <xs:element name="author" type="xs:string"/>
            <xs:element name="price" type="xs:decimal"/>
        </xs:sequence>
        <xs:attribute name="id" type="xs:string" use="required"/>
        <xs:attribute name="category" type="xs:string" use="optional"/>
    </xs:complexType>
    
    <xs:element name="book" type="BookType"/>
    
</xs:schema>
```

#### 生成的Go代码
```go
package models

// BookType represents a book with title, author, price and attributes
type BookType struct {
    // Attributes
    ID       string `xml:"id,attr" json:"id"`
    Category string `xml:"category,attr,omitempty" json:"category,omitempty"`
    
    // Elements
    Title  string  `xml:"title" json:"title"`
    Author string  `xml:"author" json:"author"`
    Price  float64 `xml:"price" json:"price"`
}
```

#### 使用示例
```go
book := models.BookType{
    ID:       "book-001",
    Category: "技术",
    Title:    "Go语言编程",
    Author:   "李四",
    Price:    89.99,
}

xmlData, _ := xml.Marshal(book)
fmt.Printf("XML: %s\n", xmlData)
// 输出: <BookType id="book-001" category="技术"><title>Go语言编程</title><author>李四</author><price>89.99</price></BookType>
```

## 🎯 枚举类型示例

### 示例4：枚举定义

#### XSD定义
```xml
<?xml version="1.0" encoding="UTF-8"?>
<xs:schema xmlns:xs="http://www.w3.org/2001/XMLSchema">
    
    <xs:simpleType name="StatusType">
        <xs:restriction base="xs:string">
            <xs:enumeration value="active"/>
            <xs:enumeration value="inactive"/>
            <xs:enumeration value="pending"/>
        </xs:restriction>
    </xs:simpleType>
    
    <xs:complexType name="UserType">
        <xs:sequence>
            <xs:element name="username" type="xs:string"/>
            <xs:element name="status" type="StatusType"/>
        </xs:sequence>
    </xs:complexType>
    
    <xs:element name="user" type="UserType"/>
    
</xs:schema>
```

#### 生成的Go代码
```go
package models

import "fmt"

// StatusType represents the status enumeration
type StatusType string

const (
    StatusTypeActive   StatusType = "active"
    StatusTypeInactive StatusType = "inactive"
    StatusTypePending  StatusType = "pending"
)

// Valid returns true if the status is valid
func (s StatusType) Valid() bool {
    switch s {
    case StatusTypeActive, StatusTypeInactive, StatusTypePending:
        return true
    default:
        return false
    }
}

// UserType represents a user with status
type UserType struct {
    Username string     `xml:"username" json:"username"`
    Status   StatusType `xml:"status" json:"status"`
}

// Validate validates the user data
func (u *UserType) Validate() error {
    if !u.Status.Valid() {
        return fmt.Errorf("invalid status: %s", u.Status)
    }
    return nil
}
```

#### 使用示例
```go
user := models.UserType{
    Username: "john_doe",
    Status:   models.StatusTypeActive,
}

if err := user.Validate(); err != nil {
    fmt.Printf("验证失败: %v\n", err)
} else {
    fmt.Printf("用户有效: %+v\n", user)
}
```

## 🔍 约束验证示例

### 示例5：带验证约束的类型

#### XSD定义
```xml
<?xml version="1.0" encoding="UTF-8"?>
<xs:schema xmlns:xs="http://www.w3.org/2001/XMLSchema">
    
    <xs:simpleType name="EmailType">
        <xs:restriction base="xs:string">
            <xs:pattern value="[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}"/>
        </xs:restriction>
    </xs:simpleType>
    
    <xs:simpleType name="AgeType">
        <xs:restriction base="xs:int">
            <xs:minInclusive value="0"/>
            <xs:maxInclusive value="150"/>
        </xs:restriction>
    </xs:simpleType>
    
    <xs:complexType name="ContactType">
        <xs:sequence>
            <xs:element name="name" type="xs:string"/>
            <xs:element name="email" type="EmailType"/>
            <xs:element name="age" type="AgeType"/>
        </xs:sequence>
    </xs:complexType>
    
    <xs:element name="contact" type="ContactType"/>
    
</xs:schema>
```

#### 生成的Go代码
```go
package models

import (
    "fmt"
    "regexp"
)

// EmailType represents an email address with validation
type EmailType string

var emailPattern = regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)

// Validate validates the email format
func (e EmailType) Validate() error {
    if !emailPattern.MatchString(string(e)) {
        return fmt.Errorf("invalid email format: %s", string(e))
    }
    return nil
}

// AgeType represents age with range validation
type AgeType int

// Validate validates the age range
func (a AgeType) Validate() error {
    if int(a) < 0 {
        return fmt.Errorf("age must be at least 0, got %d", int(a))
    }
    if int(a) > 150 {
        return fmt.Errorf("age must be at most 150, got %d", int(a))
    }
    return nil
}

// ContactType represents contact information with validation
type ContactType struct {
    Name  string    `xml:"name" json:"name"`
    Email EmailType `xml:"email" json:"email"`
    Age   AgeType   `xml:"age" json:"age"`
}

// Validate validates all contact fields
func (c *ContactType) Validate() error {
    if err := c.Email.Validate(); err != nil {
        return fmt.Errorf("email validation failed: %w", err)
    }
    
    if err := c.Age.Validate(); err != nil {
        return fmt.Errorf("age validation failed: %w", err)
    }
    
    return nil
}
```

#### 使用示例
```go
contact := models.ContactType{
    Name:  "王五",
    Email: models.EmailType("wang.wu@example.com"),
    Age:   models.AgeType(25),
}

if err := contact.Validate(); err != nil {
    fmt.Printf("验证失败: %v\n", err)
} else {
    fmt.Printf("联系人信息有效: %+v\n", contact)
}

// 测试无效数据
invalidContact := models.ContactType{
    Name:  "无效用户",
    Email: models.EmailType("invalid-email"),
    Age:   models.AgeType(200),
}

if err := invalidContact.Validate(); err != nil {
    fmt.Printf("预期的验证错误: %v\n", err)
}
```

## 📝 可选字段和数组示例

### 示例6：可选字段和重复元素

#### XSD定义
```xml
<?xml version="1.0" encoding="UTF-8"?>
<xs:schema xmlns:xs="http://www.w3.org/2001/XMLSchema">
    
    <xs:complexType name="AddressType">
        <xs:sequence>
            <xs:element name="street" type="xs:string"/>
            <xs:element name="city" type="xs:string"/>
            <xs:element name="zipCode" type="xs:string" minOccurs="0"/>
            <xs:element name="country" type="xs:string"/>
        </xs:sequence>
    </xs:complexType>
    
    <xs:complexType name="CustomerType">
        <xs:sequence>
            <xs:element name="name" type="xs:string"/>
            <xs:element name="email" type="xs:string"/>
            <xs:element name="phone" type="xs:string" minOccurs="0" maxOccurs="3"/>
            <xs:element name="address" type="AddressType" minOccurs="0"/>
        </xs:sequence>
    </xs:complexType>
    
    <xs:element name="customer" type="CustomerType"/>
    
</xs:schema>
```

#### 生成的Go代码
```go
package models

// AddressType represents an address
type AddressType struct {
    Street  string `xml:"street" json:"street"`
    City    string `xml:"city" json:"city"`
    ZipCode string `xml:"zipCode,omitempty" json:"zipCode,omitempty"`
    Country string `xml:"country" json:"country"`
}

// CustomerType represents a customer with optional fields
type CustomerType struct {
    Name    string       `xml:"name" json:"name"`
    Email   string       `xml:"email" json:"email"`
    Phone   []string     `xml:"phone,omitempty" json:"phone,omitempty"`
    Address *AddressType `xml:"address,omitempty" json:"address,omitempty"`
}
```

#### 使用示例
```go
// 完整信息的客户
customer1 := models.CustomerType{
    Name:  "赵六",
    Email: "zhao.liu@example.com",
    Phone: []string{"138-0000-0001", "138-0000-0002"},
    Address: &models.AddressType{
        Street:  "中关村大街1号",
        City:    "北京",
        ZipCode: "100000",
        Country: "中国",
    },
}

// 最小信息的客户
customer2 := models.CustomerType{
    Name:  "孙七",
    Email: "sun.qi@example.com",
}

fmt.Printf("完整客户: %+v\n", customer1)
fmt.Printf("最小客户: %+v\n", customer2)
```

## 🏗️ 嵌套复杂类型示例

### 示例7：复杂嵌套结构

#### XSD定义
```xml
<?xml version="1.0" encoding="UTF-8"?>
<xs:schema xmlns:xs="http://www.w3.org/2001/XMLSchema">
    
    <xs:complexType name="ProductType">
        <xs:sequence>
            <xs:element name="id" type="xs:string"/>
            <xs:element name="name" type="xs:string"/>
            <xs:element name="price" type="xs:decimal"/>
        </xs:sequence>
    </xs:complexType>
    
    <xs:complexType name="OrderItemType">
        <xs:sequence>
            <xs:element name="product" type="ProductType"/>
            <xs:element name="quantity" type="xs:int"/>
            <xs:element name="discount" type="xs:decimal" minOccurs="0"/>
        </xs:sequence>
    </xs:complexType>
    
    <xs:complexType name="OrderType">
        <xs:sequence>
            <xs:element name="id" type="xs:string"/>
            <xs:element name="customerId" type="xs:string"/>
            <xs:element name="orderDate" type="xs:date"/>
            <xs:element name="items" type="OrderItemType" maxOccurs="unbounded"/>
            <xs:element name="totalAmount" type="xs:decimal"/>
        </xs:sequence>
    </xs:complexType>
    
    <xs:element name="order" type="OrderType"/>
    
</xs:schema>
```

#### 生成的Go代码
```go
package models

import "time"

// ProductType represents a product
type ProductType struct {
    ID    string  `xml:"id" json:"id"`
    Name  string  `xml:"name" json:"name"`
    Price float64 `xml:"price" json:"price"`
}

// OrderItemType represents an item in an order
type OrderItemType struct {
    Product  ProductType `xml:"product" json:"product"`
    Quantity int         `xml:"quantity" json:"quantity"`
    Discount float64     `xml:"discount,omitempty" json:"discount,omitempty"`
}

// OrderType represents a complete order
type OrderType struct {
    ID          string          `xml:"id" json:"id"`
    CustomerID  string          `xml:"customerId" json:"customerId"`
    OrderDate   time.Time       `xml:"orderDate" json:"orderDate"`
    Items       []OrderItemType `xml:"items" json:"items"`
    TotalAmount float64         `xml:"totalAmount" json:"totalAmount"`
}
```

#### 使用示例
```go
order := models.OrderType{
    ID:         "ORD-001",
    CustomerID: "CUST-123",
    OrderDate:  time.Now(),
    Items: []models.OrderItemType{
        {
            Product: models.ProductType{
                ID:    "PROD-001",
                Name:  "Go编程书籍",
                Price: 89.99,
            },
            Quantity: 2,
            Discount: 10.0,
        },
        {
            Product: models.ProductType{
                ID:    "PROD-002",
                Name:  "编程键盘",
                Price: 299.99,
            },
            Quantity: 1,
        },
    },
    TotalAmount: 469.97,
}

fmt.Printf("订单信息: %+v\n", order)
```

## 🔧 多语言对比示例

### 示例8：同一XSD的多语言输出

#### XSD定义
```xml
<?xml version="1.0" encoding="UTF-8"?>
<xs:schema xmlns:xs="http://www.w3.org/2001/XMLSchema">
    
    <xs:simpleType name="PriorityType">
        <xs:restriction base="xs:string">
            <xs:enumeration value="low"/>
            <xs:enumeration value="medium"/>
            <xs:enumeration value="high"/>
        </xs:restriction>
    </xs:simpleType>
    
    <xs:complexType name="TaskType">
        <xs:sequence>
            <xs:element name="id" type="xs:string"/>
            <xs:element name="title" type="xs:string"/>
            <xs:element name="description" type="xs:string"/>
            <xs:element name="priority" type="PriorityType"/>
            <xs:element name="completed" type="xs:boolean"/>
        </xs:sequence>
    </xs:complexType>
    
    <xs:element name="task" type="TaskType"/>
    
</xs:schema>
```

#### Go代码生成
```bash
xsd2code -xsd=task.xsd -lang=go -output=task.go -package=models
```

```go
package models

// PriorityType represents task priority levels
type PriorityType string

const (
    PriorityTypeLow    PriorityType = "low"
    PriorityTypeMedium PriorityType = "medium"
    PriorityTypeHigh   PriorityType = "high"
)

// TaskType represents a task with priority
type TaskType struct {
    ID          string       `xml:"id" json:"id"`
    Title       string       `xml:"title" json:"title"`
    Description string       `xml:"description" json:"description"`
    Priority    PriorityType `xml:"priority" json:"priority"`
    Completed   bool         `xml:"completed" json:"completed"`
}
```

#### Java代码生成
```bash
xsd2code -xsd=task.xsd -lang=java -output=Task.java -package=com.example.models
```

```java
package com.example.models;

import javax.xml.bind.annotation.*;

public enum PriorityType {
    @XmlEnumValue("low")
    LOW("low"),
    
    @XmlEnumValue("medium")
    MEDIUM("medium"),
    
    @XmlEnumValue("high")
    HIGH("high");
    
    private final String value;
    
    PriorityType(String value) {
        this.value = value;
    }
    
    public String getValue() {
        return value;
    }
}

@XmlRootElement(name = "task")
@XmlAccessorType(XmlAccessType.FIELD)
public class TaskType {
    @XmlElement(required = true)
    private String id;
    
    @XmlElement(required = true)
    private String title;
    
    @XmlElement(required = true)
    private String description;
    
    @XmlElement(required = true)
    private PriorityType priority;
    
    @XmlElement(required = true)
    private Boolean completed;
    
    // Getters and setters...
}
```

#### TypeScript代码生成
```bash
xsd2code -xsd=task.xsd -lang=typescript -output=task.ts
```

```typescript
export enum PriorityType {
    Low = "low",
    Medium = "medium",
    High = "high"
}

export interface TaskType {
    id: string;
    title: string;
    description: string;
    priority: PriorityType;
    completed: boolean;
}

// 类型保护函数
export function isTaskType(obj: any): obj is TaskType {
    return typeof obj === 'object' &&
           typeof obj.id === 'string' &&
           typeof obj.title === 'string' &&
           typeof obj.description === 'string' &&
           Object.values(PriorityType).includes(obj.priority) &&
           typeof obj.completed === 'boolean';
}
```

## 🎯 实用工具脚本

### 批量生成脚本

```bash
#!/bin/bash
# generate-examples.sh

echo "🚀 生成所有示例代码..."

examples=(
    "simple"
    "person"
    "book"
    "user"
    "contact"
    "customer"
    "order"
    "task"
)

for example in "${examples[@]}"; do
    echo "生成 $example 示例..."
    
    # Go
    xsd2code -xsd="examples/${example}.xsd" \
             -lang=go \
             -output="generated/go/${example}.go" \
             -package=models
    
    # Java
    xsd2code -xsd="examples/${example}.xsd" \
             -lang=java \
             -output="generated/java/${example^}.java" \
             -package=com.example.models
    
    # TypeScript
    xsd2code -xsd="examples/${example}.xsd" \
             -lang=typescript \
             -output="generated/typescript/${example}.ts"
done

echo "✅ 所有示例生成完成!"
```

### 验证测试脚本

```bash
#!/bin/bash
# test-examples.sh

echo "🧪 测试生成的示例代码..."

# 测试Go代码
echo "测试Go代码..."
cd generated/go
go mod init examples
go mod tidy
go build ./...
go vet ./...

# 测试Java代码  
echo "测试Java代码..."
cd ../java
javac -cp ".:*" *.java

# 测试TypeScript代码
echo "测试TypeScript代码..."
cd ../typescript
tsc --noEmit *.ts

echo "✅ 所有测试通过!"
```

---

💡 **提示**: 这些基础示例涵盖了XSD2Code的核心功能。建议按顺序学习，逐步掌握各种特性。

🔗 **相关页面**: 
- [[高级示例|Advanced-Examples]] - 复杂应用场景
- [[多语言支持|Multi-Language-Support]] - 各语言特性详解
- [[约束和验证|Constraints-and-Validation]] - 验证功能详解
