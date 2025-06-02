# 约束和验证 - XSD约束处理和验证代码生成

XSD2Code能够将XSD中定义的各种约束转换为目标语言的验证代码，确保数据完整性和业务规则的正确实施。

## 🎯 约束类型概览

| 约束类型 | XSD属性 | 支持状态 | 生成代码 |
|----------|---------|----------|----------|
| **长度约束** | length, minLength, maxLength | ✅ 完全支持 | 长度检查函数 |
| **数值约束** | minInclusive, maxInclusive, minExclusive, maxExclusive | ✅ 完全支持 | 范围检查 |
| **模式约束** | pattern | ✅ 完全支持 | 正则表达式验证 |
| **枚举约束** | enumeration | ✅ 完全支持 | 枚举值检查 |
| **空白处理** | whiteSpace | ✅ 完全支持 | 字符串预处理 |
| **固定值** | fixed | ✅ 完全支持 | 固定值验证 |
| **总位数** | totalDigits | ✅ 完全支持 | 精度检查 |
| **小数位数** | fractionDigits | ✅ 完全支持 | 小数精度检查 |

## 📏 长度约束

### XSD定义

```xml
<xs:simpleType name="UsernameType">
    <xs:restriction base="xs:string">
        <xs:minLength value="3"/>
        <xs:maxLength value="20"/>
    </xs:restriction>
</xs:simpleType>

<xs:simpleType name="ProductCodeType">
    <xs:restriction base="xs:string">
        <xs:length value="8"/>
    </xs:restriction>
</xs:simpleType>
```

### Go代码生成

```go
type UsernameType string

func (u UsernameType) Validate() error {
    if len(string(u)) < 3 {
        return fmt.Errorf("username must be at least 3 characters long")
    }
    if len(string(u)) > 20 {
        return fmt.Errorf("username must be at most 20 characters long")
    }
    return nil
}

type ProductCodeType string

func (p ProductCodeType) Validate() error {
    if len(string(p)) != 8 {
        return fmt.Errorf("product code must be exactly 8 characters long")
    }
    return nil
}
```

### Java代码生成

```java
public class UsernameType {
    @Size(min = 3, max = 20, message = "Username must be between 3 and 20 characters")
    private String value;
    
    public boolean validate() {
        return value != null && value.length() >= 3 && value.length() <= 20;
    }
}

public class ProductCodeType {
    @Size(min = 8, max = 8, message = "Product code must be exactly 8 characters")
    private String value;
}
```

### C#代码生成

```csharp
public class UsernameType
{
    [StringLength(20, MinimumLength = 3, 
     ErrorMessage = "Username must be between 3 and 20 characters")]
    public string Value { get; set; }
    
    public bool IsValid()
    {
        return !string.IsNullOrEmpty(Value) && 
               Value.Length >= 3 && 
               Value.Length <= 20;
    }
}
```

## 🔢 数值约束

### XSD定义

```xml
<xs:simpleType name="AgeType">
    <xs:restriction base="xs:int">
        <xs:minInclusive value="0"/>
        <xs:maxInclusive value="150"/>
    </xs:restriction>
</xs:simpleType>

<xs:simpleType name="ScoreType">
    <xs:restriction base="xs:decimal">
        <xs:minExclusive value="0.0"/>
        <xs:maxExclusive value="100.0"/>
    </xs:restriction>
</xs:simpleType>

<xs:simpleType name="PriceType">
    <xs:restriction base="xs:decimal">
        <xs:minInclusive value="0.01"/>
        <xs:maxInclusive value="999999.99"/>
        <xs:totalDigits value="8"/>
        <xs:fractionDigits value="2"/>
    </xs:restriction>
</xs:simpleType>
```

### Go代码生成

```go
type AgeType int

func (a AgeType) Validate() error {
    if int(a) < 0 {
        return fmt.Errorf("age must be at least 0")
    }
    if int(a) > 150 {
        return fmt.Errorf("age must be at most 150")
    }
    return nil
}

type ScoreType float64

func (s ScoreType) Validate() error {
    if float64(s) <= 0.0 {
        return fmt.Errorf("score must be greater than 0.0")
    }
    if float64(s) >= 100.0 {
        return fmt.Errorf("score must be less than 100.0")
    }
    return nil
}

type PriceType float64

func (p PriceType) Validate() error {
    if float64(p) < 0.01 {
        return fmt.Errorf("price must be at least 0.01")
    }
    if float64(p) > 999999.99 {
        return fmt.Errorf("price must be at most 999999.99")
    }
    
    // 检查总位数和小数位数
    str := fmt.Sprintf("%.2f", float64(p))
    parts := strings.Split(str, ".")
    totalDigits := len(strings.ReplaceAll(parts[0], ".", "")) + len(parts[1])
    if totalDigits > 8 {
        return fmt.Errorf("price total digits must not exceed 8")
    }
    if len(parts[1]) > 2 {
        return fmt.Errorf("price fraction digits must not exceed 2")
    }
    
    return nil
}
```

### Java代码生成

```java
public class AgeType {
    @Min(value = 0, message = "Age must be at least 0")
    @Max(value = 150, message = "Age must be at most 150")
    private Integer value;
}

public class ScoreType {
    @DecimalMin(value = "0.0", inclusive = false, message = "Score must be greater than 0.0")
    @DecimalMax(value = "100.0", inclusive = false, message = "Score must be less than 100.0")
    private BigDecimal value;
}

public class PriceType {
    @DecimalMin(value = "0.01", message = "Price must be at least 0.01")
    @DecimalMax(value = "999999.99", message = "Price must be at most 999999.99")
    @Digits(integer = 6, fraction = 2, message = "Price format invalid")
    private BigDecimal value;
}
```

## 🎭 模式约束（正则表达式）

### XSD定义

```xml
<xs:simpleType name="EmailType">
    <xs:restriction base="xs:string">
        <xs:pattern value="[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}"/>
    </xs:restriction>
</xs:simpleType>

<xs:simpleType name="PhoneType">
    <xs:restriction base="xs:string">
        <xs:pattern value="(\+\d{1,3}[- ]?)?\d{10}"/>
    </xs:restriction>
</xs:simpleType>

<xs:simpleType name="ProductCodeType">
    <xs:restriction base="xs:string">
        <xs:pattern value="[A-Z]{2,3}-\d{4}-[A-Z]{2}"/>
    </xs:restriction>
</xs:simpleType>
```

### Go代码生成

```go
import (
    "fmt"
    "regexp"
)

type EmailType string

var emailPattern = regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)

func (e EmailType) Validate() error {
    if !emailPattern.MatchString(string(e)) {
        return fmt.Errorf("invalid email format")
    }
    return nil
}

type PhoneType string

var phonePattern = regexp.MustCompile(`^(\+\d{1,3}[- ]?)?\d{10}$`)

func (p PhoneType) Validate() error {
    if !phonePattern.MatchString(string(p)) {
        return fmt.Errorf("invalid phone number format")
    }
    return nil
}

type ProductCodeType string

var productCodePattern = regexp.MustCompile(`^[A-Z]{2,3}-\d{4}-[A-Z]{2}$`)

func (p ProductCodeType) Validate() error {
    if !productCodePattern.MatchString(string(p)) {
        return fmt.Errorf("invalid product code format")
    }
    return nil
}
```

### Java代码生成

```java
import javax.validation.constraints.Pattern;

public class EmailType {
    @Pattern(regexp = "[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\\.[a-zA-Z]{2,}", 
             message = "Invalid email format")
    private String value;
    
    public boolean isValid() {
        return value != null && value.matches("[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\\.[a-zA-Z]{2,}");
    }
}

public class PhoneType {
    @Pattern(regexp = "(\\+\\d{1,3}[- ]?)?\\d{10}", 
             message = "Invalid phone number format")
    private String value;
}
```

### C#代码生成

```csharp
using System.ComponentModel.DataAnnotations;
using System.Text.RegularExpressions;

public class EmailType
{
    [RegularExpression(@"[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}", 
                       ErrorMessage = "Invalid email format")]
    public string Value { get; set; }
    
    public bool IsValid()
    {
        var regex = new Regex(@"[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}");
        return !string.IsNullOrEmpty(Value) && regex.IsMatch(Value);
    }
}
```

## 📝 空白字符处理

### XSD定义

```xml
<xs:simpleType name="PreserveType">
    <xs:restriction base="xs:string">
        <xs:whiteSpace value="preserve"/>
    </xs:restriction>
</xs:simpleType>

<xs:simpleType name="ReplaceType">
    <xs:restriction base="xs:string">
        <xs:whiteSpace value="replace"/>
    </xs:restriction>
</xs:simpleType>

<xs:simpleType name="CollapseType">
    <xs:restriction base="xs:string">
        <xs:whiteSpace value="collapse"/>
    </xs:restriction>
</xs:simpleType>
```

### Go代码生成

```go
import (
    "regexp"
    "strings"
)

type PreserveType string

func (p PreserveType) Normalize() string {
    // preserve: 保持所有空白字符
    return string(p)
}

type ReplaceType string

func (r ReplaceType) Normalize() string {
    // replace: 替换制表符、换行符等为空格
    s := strings.ReplaceAll(string(r), "\t", " ")
    s = strings.ReplaceAll(s, "\n", " ")
    s = strings.ReplaceAll(s, "\r", " ")
    return s
}

type CollapseType string

var multiSpacePattern = regexp.MustCompile(`\s+`)

func (c CollapseType) Normalize() string {
    // collapse: 折叠连续空格并去除首尾空格
    s := c.Replace().Normalize()
    s = multiSpacePattern.ReplaceAllString(s, " ")
    return strings.TrimSpace(s)
}

// 辅助函数
func applyWhiteSpaceProcessing(value string, action string) string {
    switch action {
    case "preserve":
        return value
    case "replace":
        s := strings.ReplaceAll(value, "\t", " ")
        s = strings.ReplaceAll(s, "\n", " ")
        s = strings.ReplaceAll(s, "\r", " ")
        return s
    case "collapse":
        s := applyWhiteSpaceProcessing(value, "replace")
        re := regexp.MustCompile(`\s+`)
        s = re.ReplaceAllString(s, " ")
        return strings.TrimSpace(s)
    default:
        return value
    }
}
```

## 🎯 枚举约束

### XSD定义

```xml
<xs:simpleType name="StatusType">
    <xs:restriction base="xs:string">
        <xs:enumeration value="active"/>
        <xs:enumeration value="inactive"/>
        <xs:enumeration value="pending"/>
        <xs:enumeration value="suspended"/>
    </xs:restriction>
</xs:simpleType>

<xs:simpleType name="PriorityType">
    <xs:restriction base="xs:int">
        <xs:enumeration value="1"/>
        <xs:enumeration value="2"/>
        <xs:enumeration value="3"/>
        <xs:enumeration value="4"/>
        <xs:enumeration value="5"/>
    </xs:restriction>
</xs:simpleType>
```

### Go代码生成

```go
type StatusType string

const (
    StatusTypeActive    StatusType = "active"
    StatusTypeInactive  StatusType = "inactive"
    StatusTypePending   StatusType = "pending"
    StatusTypeSuspended StatusType = "suspended"
)

var validStatusTypes = map[StatusType]bool{
    StatusTypeActive:    true,
    StatusTypeInactive:  true,
    StatusTypePending:   true,
    StatusTypeSuspended: true,
}

func (s StatusType) Validate() error {
    if !validStatusTypes[s] {
        return fmt.Errorf("invalid status type: %s", string(s))
    }
    return nil
}

func (s StatusType) String() string {
    return string(s)
}

type PriorityType int

const (
    PriorityTypeLow      PriorityType = 1
    PriorityTypeMedium   PriorityType = 2
    PriorityTypeNormal   PriorityType = 3
    PriorityTypeHigh     PriorityType = 4
    PriorityTypeCritical PriorityType = 5
)

func (p PriorityType) Validate() error {
    switch p {
    case 1, 2, 3, 4, 5:
        return nil
    default:
        return fmt.Errorf("invalid priority: %d", int(p))
    }
}
```

### Java代码生成

```java
public enum StatusType {
    ACTIVE("active"),
    INACTIVE("inactive"),
    PENDING("pending"),
    SUSPENDED("suspended");
    
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
        throw new IllegalArgumentException("Invalid status: " + value);
    }
}

public enum PriorityType {
    LOW(1),
    MEDIUM(2), 
    NORMAL(3),
    HIGH(4),
    CRITICAL(5);
    
    private final int value;
    
    PriorityType(int value) {
        this.value = value;
    }
    
    public int getValue() {
        return value;
    }
}
```

## 🔒 固定值约束

### XSD定义

```xml
<xs:element name="version" type="xs:string" fixed="1.0"/>

<xs:complexType name="ConfigType">
    <xs:sequence>
        <xs:element name="apiVersion" type="xs:string" fixed="v2"/>
        <xs:element name="format" type="xs:string" fixed="json"/>
    </xs:sequence>
</xs:complexType>
```

### Go代码生成

```go
const DefaultVersion = "1.0"

type ConfigType struct {
    APIVersion string `xml:"apiVersion" json:"apiVersion"`
    Format     string `xml:"format" json:"format"`
}

func NewConfigType() *ConfigType {
    return &ConfigType{
        APIVersion: "v2",
        Format:     "json",
    }
}

func (c *ConfigType) Validate() error {
    if c.APIVersion != "v2" {
        return fmt.Errorf("apiVersion must be 'v2', got '%s'", c.APIVersion)
    }
    if c.Format != "json" {
        return fmt.Errorf("format must be 'json', got '%s'", c.Format)
    }
    return nil
}
```

## 🔄 复合约束验证

### XSD定义

```xml
<xs:simpleType name="ComplexUserIdType">
    <xs:restriction base="xs:string">
        <xs:pattern value="USR[0-9]{8}"/>
        <xs:length value="11"/>
        <xs:whiteSpace value="collapse"/>
    </xs:restriction>
</xs:simpleType>

<xs:simpleType name="ComplexPriceType">
    <xs:restriction base="xs:decimal">
        <xs:minInclusive value="0.01"/>
        <xs:maxInclusive value="999999.99"/>
        <xs:totalDigits value="8"/>
        <xs:fractionDigits value="2"/>
    </xs:restriction>
</xs:simpleType>
```

### Go代码生成

```go
type ComplexUserIdType string

var userIdPattern = regexp.MustCompile(`^USR[0-9]{8}$`)

func (c ComplexUserIdType) Validate() error {
    // 首先进行空白字符处理
    normalized := applyWhiteSpaceProcessing(string(c), "collapse")
    
    // 检查长度
    if len(normalized) != 11 {
        return fmt.Errorf("user ID must be exactly 11 characters long")
    }
    
    // 检查模式
    if !userIdPattern.MatchString(normalized) {
        return fmt.Errorf("user ID must match pattern USR[0-9]{8}")
    }
    
    return nil
}

func (c ComplexUserIdType) Normalize() ComplexUserIdType {
    return ComplexUserIdType(applyWhiteSpaceProcessing(string(c), "collapse"))
}

type ComplexPriceType float64

func (c ComplexPriceType) Validate() error {
    value := float64(c)
    
    // 检查数值范围
    if value < 0.01 {
        return fmt.Errorf("price must be at least 0.01")
    }
    if value > 999999.99 {
        return fmt.Errorf("price must be at most 999999.99")
    }
    
    // 检查数字精度
    str := fmt.Sprintf("%.2f", value)
    parts := strings.Split(str, ".")
    
    // 检查总位数
    totalDigits := len(strings.ReplaceAll(parts[0], "0", "")) + len(parts[1])
    if len(parts[0]) + len(parts[1]) > 8 {
        return fmt.Errorf("price total digits must not exceed 8")
    }
    
    // 检查小数位数
    if len(parts[1]) > 2 {
        return fmt.Errorf("price fraction digits must not exceed 2")
    }
    
    return nil
}
```

## 🧪 验证代码使用示例

### 完整验证示例

```go
package main

import (
    "fmt"
    "log"
)

func main() {
    // 测试用户ID验证
    userID := ComplexUserIdType("  USR12345678  ")
    normalizedID := userID.Normalize()
    
    if err := normalizedID.Validate(); err != nil {
        log.Printf("用户ID验证失败: %v", err)
    } else {
        fmt.Printf("用户ID验证成功: %s\n", string(normalizedID))
    }
    
    // 测试价格验证
    price := ComplexPriceType(123.45)
    if err := price.Validate(); err != nil {
        log.Printf("价格验证失败: %v", err)
    } else {
        fmt.Printf("价格验证成功: %.2f\n", float64(price))
    }
    
    // 测试邮箱验证
    email := EmailType("user@example.com")
    if err := email.Validate(); err != nil {
        log.Printf("邮箱验证失败: %v", err)
    } else {
        fmt.Printf("邮箱验证成功: %s\n", string(email))
    }
}
```

## 🎛️ 验证配置选项

### 启用验证代码生成

```bash
# 生成包含验证的Go代码
xsd2code -xsd=schema.xsd -validation=true

# 生成带有严格验证的代码
xsd2code -xsd=schema.xsd -validation=true -strict=true

# 只生成类型定义，不生成验证
xsd2code -xsd=schema.xsd -validation=false
```

### 自定义验证选项

```bash
# 自定义验证错误消息语言
xsd2code -xsd=schema.xsd -validation=true -error-lang=zh

# 生成验证接口
xsd2code -xsd=schema.xsd -validation=true -interface=true
```

## 🏆 最佳实践

### 1. 验证时机

- **输入验证**: 在数据输入时立即验证
- **序列化前**: 在XML/JSON序列化前验证
- **业务逻辑前**: 在处理业务逻辑前验证

### 2. 错误处理

```go
type ValidationError struct {
    Field   string
    Value   interface{}
    Message string
}

func (e ValidationError) Error() string {
    return fmt.Sprintf("validation error in field '%s': %s", e.Field, e.Message)
}
```

### 3. 性能优化

- 预编译正则表达式
- 缓存验证结果
- 批量验证优化

### 4. 国际化支持

```go
type ValidationMessages struct {
    Language string
    Messages map[string]string
}

var zhMessages = ValidationMessages{
    Language: "zh",
    Messages: map[string]string{
        "invalid_email": "邮箱格式无效",
        "length_error":  "长度不符合要求",
        "range_error":   "数值超出范围",
    },
}
```

---

💡 **提示**: 合理使用验证功能可以提高数据质量，但过度验证可能影响性能。

🔗 **相关页面**: 
- [[XSD特性支持|XSD-Features]] - XSD特性详解
- [[多语言支持|Multi-Language-Support]] - 各语言验证实现
- [[最佳实践|Best-Practices]] - 验证最佳实践
