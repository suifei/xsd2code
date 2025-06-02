# XSD特性支持

XSD2Code 全面支持XML Schema Definition的各种特性。本页面详细说明支持的XSD构造和处理方式。

## 📋 支持的XSD特性概览

| 特性类别 | 支持状态 | 说明 |
|----------|----------|------|
| 🟢 简单类型 | ✅ 完全支持 | restriction, enumeration, pattern等 |
| 🟢 复杂类型 | ✅ 完全支持 | sequence, choice, all, mixed content |
| 🟢 元素 | ✅ 完全支持 | 基本元素、可选元素、数组元素 |
| 🟢 属性 | ✅ 完全支持 | 必需、可选、固定值、默认值 |
| 🟢 约束 | ✅ 完全支持 | 所有restriction类型 |
| 🟢 命名空间 | ✅ 完全支持 | targetNamespace, xmlns处理 |
| 🟢 导入/包含 | ✅ 完全支持 | import, include, redefine |
| 🟢 组 | ✅ 完全支持 | group, attributeGroup |
| 🟢 扩展 | ✅ 完全支持 | extension, restriction |
| 🟢 数据类型 | ✅ 完全支持 | 所有XSD内置类型 |

## 🔤 简单类型 (Simple Types)

### 内置数据类型

XSD2Code支持所有W3C XSD内置数据类型：

#### 字符串类型
```xml
<xs:element name="text" type="xs:string"/>
<xs:element name="token" type="xs:token"/>
<xs:element name="normalizedString" type="xs:normalizedString"/>
```

**映射到**:
- **Go**: `string`
- **Java**: `String`
- **C#**: `string`
- **TypeScript**: `string`

#### 数值类型
```xml
<xs:element name="integer" type="xs:int"/>
<xs:element name="decimal" type="xs:decimal"/>
<xs:element name="double" type="xs:double"/>
<xs:element name="float" type="xs:float"/>
<xs:element name="long" type="xs:long"/>
<xs:element name="short" type="xs:short"/>
<xs:element name="byte" type="xs:byte"/>
```

**映射到**:
- **Go**: `int`, `float64`, `int64`, `int16`, `int8`
- **Java**: `Integer`, `BigDecimal`, `Double`, `Float`, `Long`, `Short`, `Byte`
- **C#**: `int`, `decimal`, `double`, `float`, `long`, `short`, `byte`
- **TypeScript**: `number`, `bigint`

#### 日期时间类型
```xml
<xs:element name="date" type="xs:date"/>
<xs:element name="time" type="xs:time"/>
<xs:element name="dateTime" type="xs:dateTime"/>
<xs:element name="duration" type="xs:duration"/>
```

**映射到**:
- **Go**: `time.Time`, `time.Duration`
- **Java**: `LocalDate`, `LocalTime`, `LocalDateTime`, `Duration`
- **C#**: `DateTime`, `TimeSpan`
- **TypeScript**: `Date`, `string`

#### 布尔类型
```xml
<xs:element name="flag" type="xs:boolean"/>
```

**映射到**:
- **Go**: `bool`
- **Java**: `Boolean`
- **C#**: `bool`
- **TypeScript**: `boolean`

### 自定义简单类型

#### 使用restriction
```xml
<xs:simpleType name="ProductCodeType">
    <xs:restriction base="xs:string">
        <xs:pattern value="[A-Z]{3}[0-9]{3}"/>
        <xs:length value="6"/>
    </xs:restriction>
</xs:simpleType>
```

**生成的Go代码**:
```go
type ProductCodeType string

func (v ProductCodeType) Validate() bool {
    strVal := string(v)
    if len(strVal) != 6 {
        return false
    }
    pattern := regexp.MustCompile(`[A-Z]{3}[0-9]{3}`)
    return pattern.MatchString(strVal)
}
```

#### 枚举类型
```xml
<xs:simpleType name="StatusType">
    <xs:restriction base="xs:string">
        <xs:enumeration value="active"/>
        <xs:enumeration value="inactive"/>
        <xs:enumeration value="pending"/>
    </xs:restriction>
</xs:simpleType>
```

**生成的Go代码**:
```go
type StatusType string

const (
    StatusTypeActive   StatusType = "active"
    StatusTypeInactive StatusType = "inactive"
    StatusTypePending  StatusType = "pending"
)
```

## 🏗️ 复杂类型 (Complex Types)

### Sequence

```xml
<xs:complexType name="PersonType">
    <xs:sequence>
        <xs:element name="firstName" type="xs:string"/>
        <xs:element name="lastName" type="xs:string"/>
        <xs:element name="age" type="xs:int" minOccurs="0"/>
    </xs:sequence>
</xs:complexType>
```

**生成的Go代码**:
```go
type PersonType struct {
    FirstName string `xml:"firstName"`
    LastName  string `xml:"lastName"`
    Age       *int   `xml:"age,omitempty"`
}
```

### Choice

```xml
<xs:complexType name="ContactType">
    <xs:choice>
        <xs:element name="email" type="xs:string"/>
        <xs:element name="phone" type="xs:string"/>
    </xs:choice>
</xs:complexType>
```

**生成的Go代码**:
```go
type ContactType struct {
    Email *string `xml:"email,omitempty"`
    Phone *string `xml:"phone,omitempty"`
}
```

### All

```xml
<xs:complexType name="AddressType">
    <xs:all>
        <xs:element name="street" type="xs:string"/>
        <xs:element name="city" type="xs:string"/>
        <xs:element name="zipCode" type="xs:string"/>
    </xs:all>
</xs:complexType>
```

### Mixed Content

```xml
<xs:complexType name="DocumentType" mixed="true">
    <xs:sequence>
        <xs:element name="title" type="xs:string"/>
        <xs:element name="author" type="xs:string"/>
    </xs:sequence>
</xs:complexType>
```

## 📏 约束 (Restrictions)

XSD2Code支持所有XSD约束类型，并生成相应的验证代码。

### 字符串约束

#### Length约束
```xml
<xs:simpleType name="ExactLengthType">
    <xs:restriction base="xs:string">
        <xs:length value="10"/>
    </xs:restriction>
</xs:simpleType>

<xs:simpleType name="RangeLengthType">
    <xs:restriction base="xs:string">
        <xs:minLength value="5"/>
        <xs:maxLength value="20"/>
    </xs:restriction>
</xs:simpleType>
```

#### Pattern约束
```xml
<xs:simpleType name="EmailType">
    <xs:restriction base="xs:string">
        <xs:pattern value="[^@]+@[^@]+\.[^@]+"/>
    </xs:restriction>
</xs:simpleType>
```

#### WhiteSpace约束
```xml
<xs:simpleType name="CollapsedType">
    <xs:restriction base="xs:string">
        <xs:whiteSpace value="collapse"/>
    </xs:restriction>
</xs:simpleType>
```

**支持的whiteSpace值**:
- `preserve`: 保持所有空白字符
- `replace`: 替换制表符、换行符为空格
- `collapse`: 折叠多个空格并去除首尾空格

### 数值约束

```xml
<xs:simpleType name="PercentageType">
    <xs:restriction base="xs:decimal">
        <xs:minInclusive value="0.0"/>
        <xs:maxInclusive value="100.0"/>
        <xs:fractionDigits value="2"/>
    </xs:restriction>
</xs:simpleType>
```

**支持的数值约束**:
- `minInclusive` / `maxInclusive`: 包含边界
- `minExclusive` / `maxExclusive`: 排除边界
- `totalDigits`: 总位数
- `fractionDigits`: 小数位数

### 固定值和默认值

```xml
<xs:element name="version" type="xs:string" fixed="1.0"/>
<xs:element name="country" type="xs:string" default="USA"/>
<xs:attribute name="status" type="xs:string" default="active"/>
```

## 🏷️ 元素和属性

### 元素特性

#### 可选性控制
```xml
<!-- 必需元素 -->
<xs:element name="required" type="xs:string"/>

<!-- 可选元素 -->
<xs:element name="optional" type="xs:string" minOccurs="0"/>

<!-- 数组元素 -->
<xs:element name="items" type="xs:string" maxOccurs="unbounded"/>

<!-- 有限数组 -->
<xs:element name="tags" type="xs:string" minOccurs="0" maxOccurs="5"/>
```

**Go代码生成规则**:
- `minOccurs="0"` → 指针类型 `*Type`
- `maxOccurs="unbounded"` → 切片类型 `[]Type`
- `maxOccurs > 1` → 切片类型 `[]Type`

### 属性特性

```xml
<xs:complexType name="BookType">
    <xs:sequence>
        <xs:element name="title" type="xs:string"/>
    </xs:sequence>
    <xs:attribute name="id" type="xs:string" use="required"/>
    <xs:attribute name="category" type="xs:string" use="optional"/>
    <xs:attribute name="version" type="xs:string" fixed="1.0"/>
    <xs:attribute name="lang" type="xs:string" default="en"/>
</xs:complexType>
```

**生成的Go代码**:
```go
type BookType struct {
    Title    string  `xml:"title"`
    Id       string  `xml:"id,attr"`
    Category *string `xml:"category,attr,omitempty"`
    Version  string  `xml:"version,attr"` // 固定值在注释中说明
    Lang     *string `xml:"lang,attr,omitempty"`
}
```

## 🌐 命名空间支持

### Target Namespace

```xml
<xs:schema targetNamespace="http://example.com/ns"
           xmlns:tns="http://example.com/ns"
           xmlns:xs="http://www.w3.org/2001/XMLSchema">
    
    <xs:element name="document" type="tns:DocumentType"/>
    
</xs:schema>
```

**生成的代码包含完整的命名空间信息**:
```go
type DocumentType struct {
    XMLName xml.Name `xml:"http://example.com/ns document"`
    // 其他字段...
}
```

### 多命名空间处理

```xml
<xs:schema targetNamespace="http://example.com/main"
           xmlns:tns="http://example.com/main"
           xmlns:ext="http://example.com/external"
           xmlns:xs="http://www.w3.org/2001/XMLSchema">
    
    <xs:import namespace="http://example.com/external" 
               schemaLocation="external.xsd"/>
    
    <xs:element name="document">
        <xs:complexType>
            <xs:sequence>
                <xs:element name="content" type="tns:ContentType"/>
                <xs:element name="metadata" type="ext:MetadataType"/>
            </xs:sequence>
        </xs:complexType>
    </xs:element>
</xs:schema>
```

## 📦 导入和包含

### Import

```xml
<xs:import namespace="http://external.com/types" 
           schemaLocation="external-types.xsd"/>
```

工具会自动处理导入的XSD文件，生成统一的类型定义。

### Include

```xml
<xs:include schemaLocation="common-types.xsd"/>
```

包含的类型会合并到当前命名空间中。

### Redefine

```xml
<xs:redefine schemaLocation="base-types.xsd">
    <xs:complexType name="ExtendedType">
        <xs:complexContent>
            <xs:extension base="BaseType">
                <xs:sequence>
                    <xs:element name="extraField" type="xs:string"/>
                </xs:sequence>
            </xs:extension>
        </xs:complexContent>
    </xs:complexType>
</xs:redefine>
```

## 🔗 组和组引用

### Group定义和引用

```xml
<xs:group name="AddressGroup">
    <xs:sequence>
        <xs:element name="street" type="xs:string"/>
        <xs:element name="city" type="xs:string"/>
        <xs:element name="zipCode" type="xs:string"/>
    </xs:sequence>
</xs:group>

<xs:complexType name="PersonType">
    <xs:sequence>
        <xs:element name="name" type="xs:string"/>
        <xs:group ref="AddressGroup"/>
    </xs:sequence>
</xs:complexType>
```

### AttributeGroup

```xml
<xs:attributeGroup name="CommonAttributes">
    <xs:attribute name="id" type="xs:string" use="required"/>
    <xs:attribute name="version" type="xs:string" default="1.0"/>
</xs:attributeGroup>

<xs:complexType name="DocumentType">
    <xs:sequence>
        <xs:element name="content" type="xs:string"/>
    </xs:sequence>
    <xs:attributeGroup ref="CommonAttributes"/>
</xs:complexType>
```

## 🔄 类型扩展

### ComplexContent Extension

```xml
<xs:complexType name="BaseType">
    <xs:sequence>
        <xs:element name="id" type="xs:string"/>
        <xs:element name="name" type="xs:string"/>
    </xs:sequence>
</xs:complexType>

<xs:complexType name="ExtendedType">
    <xs:complexContent>
        <xs:extension base="BaseType">
            <xs:sequence>
                <xs:element name="description" type="xs:string"/>
            </xs:sequence>
            <xs:attribute name="category" type="xs:string"/>
        </xs:extension>
    </xs:complexContent>
</xs:complexType>
```

**生成的Go代码使用嵌入**:
```go
type BaseType struct {
    Id   string `xml:"id"`
    Name string `xml:"name"`
}

type ExtendedType struct {
    BaseType
    Description string  `xml:"description"`
    Category    *string `xml:"category,attr,omitempty"`
}
```

### SimpleContent Extension

```xml
<xs:complexType name="AnnotatedString">
    <xs:simpleContent>
        <xs:extension base="xs:string">
            <xs:attribute name="lang" type="xs:string"/>
        </xs:extension>
    </xs:simpleContent>
</xs:complexType>
```

## 🎯 高级特性

### Union Types

```xml
<xs:simpleType name="StringOrNumber">
    <xs:union memberTypes="xs:string xs:int"/>
</xs:simpleType>
```

**Go代码处理Union**:
```go
type StringOrNumber interface{}

// 提供类型检查函数
func (v StringOrNumber) AsString() (string, bool) {
    if s, ok := v.(string); ok {
        return s, true
    }
    return "", false
}

func (v StringOrNumber) AsInt() (int, bool) {
    if i, ok := v.(int); ok {
        return i, true
    }
    return 0, false
}
```

### List Types

```xml
<xs:simpleType name="NumberList">
    <xs:list itemType="xs:int"/>
</xs:simpleType>
```

**映射为切片类型**:
```go
type NumberList []int
```

## 🔍 类型映射参考

### 完整类型映射表

| XSD Type | Go | Java | C# | TypeScript |
|----------|----|----- |----|------------|
| `xs:string` | `string` | `String` | `string` | `string` |
| `xs:int` | `int` | `Integer` | `int` | `number` |
| `xs:long` | `int64` | `Long` | `long` | `number` |
| `xs:decimal` | `float64` | `BigDecimal` | `decimal` | `number` |
| `xs:boolean` | `bool` | `Boolean` | `bool` | `boolean` |
| `xs:date` | `time.Time` | `LocalDate` | `DateTime` | `Date` |
| `xs:dateTime` | `time.Time` | `LocalDateTime` | `DateTime` | `Date` |
| `xs:duration` | `time.Duration` | `Duration` | `TimeSpan` | `string` |
| `xs:base64Binary` | `[]byte` | `byte[]` | `byte[]` | `string` |
| `xs:anyURI` | `string` | `String` | `string` | `string` |

## 📝 使用建议

### 性能优化

1. **指针使用**: 可选字段使用指针类型，减少内存占用
2. **切片预分配**: 对于已知大小的数组，考虑预分配容量
3. **验证缓存**: 复杂的pattern验证可以缓存编译的正则表达式

### 最佳实践

1. **命名约定**: 遵循目标语言的命名约定
2. **包结构**: 合理组织生成的代码包结构
3. **文档注释**: 保留XSD中的文档注释
4. **版本控制**: 生成的代码应纳入版本控制

---

💡 **提示**: XSD2Code会根据目标语言的特点生成最适合的代码结构，同时保持与原始XSD语义的一致性。
