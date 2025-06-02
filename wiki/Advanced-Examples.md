# 高级示例

本页面展示XSD2Code在复杂场景下的使用方法，包括企业级应用、复杂XSD处理和高级配置。

## 🏢 企业级Web API开发

### 场景描述
为微服务架构的RESTful API设计数据传输对象(DTO)，需要支持多种数据格式和严格的数据验证。

### XSD设计

```xml
<?xml version="1.0" encoding="UTF-8"?>
<xs:schema xmlns:xs="http://www.w3.org/2001/XMLSchema"
           targetNamespace="http://api.company.com/v1"
           xmlns:tns="http://api.company.com/v1"
           elementFormDefault="qualified">

    <!-- 用户状态枚举 -->
    <xs:simpleType name="UserStatusType">
        <xs:restriction base="xs:string">
            <xs:enumeration value="active"/>
            <xs:enumeration value="inactive"/>
            <xs:enumeration value="suspended"/>
            <xs:enumeration value="pending"/>
        </xs:restriction>
    </xs:simpleType>

    <!-- 邮箱类型（带验证） -->
    <xs:simpleType name="EmailType">
        <xs:restriction base="xs:string">
            <xs:pattern value="[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}"/>
            <xs:minLength value="5"/>
            <xs:maxLength value="100"/>
        </xs:restriction>
    </xs:simpleType>

    <!-- 手机号类型 -->
    <xs:simpleType name="PhoneType">
        <xs:restriction base="xs:string">
            <xs:pattern value="(\+\d{1,3}[- ]?)?\d{10}"/>
        </xs:restriction>
    </xs:simpleType>

    <!-- 用户ID类型 -->
    <xs:simpleType name="UserIdType">
        <xs:restriction base="xs:string">
            <xs:pattern value="USR[0-9]{8}"/>
            <xs:length value="11"/>
        </xs:restriction>
    </xs:simpleType>

    <!-- 地址复杂类型 -->
    <xs:complexType name="AddressType">
        <xs:sequence>
            <xs:element name="street" type="xs:string" minOccurs="1">
                <xs:annotation>
                    <xs:documentation>街道地址</xs:documentation>
                </xs:annotation>
            </xs:element>
            <xs:element name="city" type="xs:string" minOccurs="1"/>
            <xs:element name="state" type="xs:string" minOccurs="0"/>
            <xs:element name="postalCode" type="xs:string" minOccurs="1">
                <xs:annotation>
                    <xs:documentation>邮政编码</xs:documentation>
                </xs:annotation>
            </xs:element>
            <xs:element name="country" type="xs:string" default="CN"/>
        </xs:sequence>
    </xs:complexType>

    <!-- 用户档案 -->
    <xs:complexType name="UserProfileType">
        <xs:sequence>
            <xs:element name="firstName" type="xs:string" minOccurs="1"/>
            <xs:element name="lastName" type="xs:string" minOccurs="1"/>
            <xs:element name="email" type="tns:EmailType" minOccurs="1"/>
            <xs:element name="phone" type="tns:PhoneType" minOccurs="0"/>
            <xs:element name="birthDate" type="xs:date" minOccurs="0"/>
            <xs:element name="address" type="tns:AddressType" minOccurs="0"/>
            <xs:element name="tags" type="xs:string" minOccurs="0" maxOccurs="10"/>
        </xs:sequence>
        <xs:attribute name="userId" type="tns:UserIdType" use="required"/>
        <xs:attribute name="status" type="tns:UserStatusType" default="pending"/>
        <xs:attribute name="createdAt" type="xs:dateTime" use="required"/>
        <xs:attribute name="lastModified" type="xs:dateTime"/>
    </xs:complexType>

    <!-- API响应封装 -->
    <xs:complexType name="ApiResponseType">
        <xs:sequence>
            <xs:element name="success" type="xs:boolean"/>
            <xs:element name="message" type="xs:string" minOccurs="0"/>
            <xs:element name="data" type="tns:UserProfileType" minOccurs="0"/>
            <xs:element name="errors" type="xs:string" minOccurs="0" maxOccurs="unbounded"/>
        </xs:sequence>
        <xs:attribute name="requestId" type="xs:string" use="required"/>
        <xs:attribute name="timestamp" type="xs:dateTime" use="required"/>
    </xs:complexType>

    <!-- 根元素 -->
    <xs:element name="userProfile" type="tns:UserProfileType"/>
    <xs:element name="apiResponse" type="tns:ApiResponseType"/>

</xs:schema>
```

### 代码生成命令

```bash
# 生成Go API模型（带JSON支持和验证）
xsd2code \
  -xsd=api-v1.xsd \
  -lang=go \
  -output=api/models/user.go \
  -package=models \
  -json \
  -validation \
  -validation-output=api/models/user_validation.go \
  -tests \
  -test-output=api/models/user_test.go

# 生成TypeScript接口（前端使用）
xsd2code \
  -xsd=api-v1.xsd \
  -lang=typescript \
  -output=frontend/src/types/user.ts

# 生成Java DTOs（其他微服务使用）
xsd2code \
  -xsd=api-v1.xsd \
  -lang=java \
  -output=java/src/main/java/com/company/api/UserModels.java \
  -package=com.company.api.models
```

### 生成的Go代码示例

```go
package models

import (
    "encoding/xml"
    "encoding/json"
    "regexp"
    "strings"
    "time"
)

// UserStatusType represents user status enumeration
type UserStatusType string

const (
    UserStatusTypeActive    UserStatusType = "active"
    UserStatusTypeInactive  UserStatusType = "inactive"
    UserStatusTypeSuspended UserStatusType = "suspended"
    UserStatusTypePending   UserStatusType = "pending"
)

// EmailType represents validated email address
type EmailType string

// Validate validates the EmailType format
func (e EmailType) Validate() bool {
    str := string(e)
    if len(str) < 5 || len(str) > 100 {
        return false
    }
    pattern := regexp.MustCompile(`[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}`)
    return pattern.MatchString(str)
}

// UserProfileType represents user profile information
type UserProfileType struct {
    XMLName    xml.Name        `xml:"http://api.company.com/v1 userProfile" json:"-"`
    FirstName  string          `xml:"firstName" json:"firstName"`
    LastName   string          `xml:"lastName" json:"lastName"`
    Email      EmailType       `xml:"email" json:"email"`
    Phone      *PhoneType      `xml:"phone,omitempty" json:"phone,omitempty"`
    BirthDate  *time.Time      `xml:"birthDate,omitempty" json:"birthDate,omitempty"`
    Address    *AddressType    `xml:"address,omitempty" json:"address,omitempty"`
    Tags       []string        `xml:"tags,omitempty" json:"tags,omitempty"`
    UserId     UserIdType      `xml:"userId,attr" json:"userId"`
    Status     *UserStatusType `xml:"status,attr,omitempty" json:"status,omitempty"`
    CreatedAt  time.Time       `xml:"createdAt,attr" json:"createdAt"`
    LastModified *time.Time    `xml:"lastModified,attr,omitempty" json:"lastModified,omitempty"`
}

// Validate validates the UserProfileType
func (u UserProfileType) Validate() bool {
    if !u.Email.Validate() {
        return false
    }
    if u.Phone != nil && !u.Phone.Validate() {
        return false
    }
    if !u.UserId.Validate() {
        return false
    }
    return true
}
```

### 使用示例

```go
package main

import (
    "encoding/json"
    "fmt"
    "time"
    "yourproject/api/models"
)

func main() {
    // 创建用户档案
    profile := models.UserProfileType{
        FirstName: "张",
        LastName:  "三",
        Email:     models.EmailType("zhangsan@company.com"),
        UserId:    models.UserIdType("USR12345678"),
        CreatedAt: time.Now(),
    }

    // 验证数据
    if !profile.Validate() {
        fmt.Println("数据验证失败")
        return
    }

    // 序列化为JSON
    jsonData, _ := json.MarshalIndent(profile, "", "  ")
    fmt.Println(string(jsonData))

    // 创建API响应
    response := models.ApiResponseType{
        Success:   true,
        Data:      &profile,
        RequestId: "req-12345",
        Timestamp: time.Now(),
    }

    responseJson, _ := json.MarshalIndent(response, "", "  ")
    fmt.Println(string(responseJson))
}
```

## 🏭 工业自动化数据交换

### 场景描述
为工业自动化系统设计PLC数据交换格式，需要精确的数据类型映射和实时性能。

### XSD设计

```xml
<?xml version="1.0" encoding="UTF-8"?>
<xs:schema xmlns:xs="http://www.w3.org/2001/XMLSchema"
           targetNamespace="http://automation.company.com/plc"
           xmlns:tns="http://automation.company.com/plc"
           elementFormDefault="qualified">

    <!-- PLC数据类型 -->
    <xs:simpleType name="PLCBoolType">
        <xs:restriction base="xs:boolean"/>
    </xs:simpleType>

    <xs:simpleType name="PLCIntType">
        <xs:restriction base="xs:int">
            <xs:minInclusive value="-32768"/>
            <xs:maxInclusive value="32767"/>
        </xs:restriction>
    </xs:simpleType>

    <xs:simpleType name="PLCRealType">
        <xs:restriction base="xs:float">
            <xs:minInclusive value="-3.4E+38"/>
            <xs:maxInclusive value="3.4E+38"/>
        </xs:restriction>
    </xs:simpleType>

    <!-- 设备状态 -->
    <xs:simpleType name="DeviceStatusType">
        <xs:restriction base="xs:string">
            <xs:enumeration value="RUNNING"/>
            <xs:enumeration value="STOPPED"/>
            <xs:enumeration value="ERROR"/>
            <xs:enumeration value="MAINTENANCE"/>
        </xs:restriction>
    </xs:simpleType>

    <!-- 传感器数据 -->
    <xs:complexType name="SensorDataType">
        <xs:sequence>
            <xs:element name="temperature" type="tns:PLCRealType"/>
            <xs:element name="pressure" type="tns:PLCRealType"/>
            <xs:element name="humidity" type="tns:PLCRealType" minOccurs="0"/>
            <xs:element name="vibration" type="tns:PLCRealType" minOccurs="0"/>
        </xs:sequence>
        <xs:attribute name="sensorId" type="xs:string" use="required"/>
        <xs:attribute name="timestamp" type="xs:dateTime" use="required"/>
    </xs:complexType>

    <!-- 执行器控制 -->
    <xs:complexType name="ActuatorControlType">
        <xs:sequence>
            <xs:element name="motorSpeed" type="tns:PLCRealType"/>
            <xs:element name="valvePosition" type="tns:PLCRealType"/>
            <xs:element name="enabled" type="tns:PLCBoolType"/>
        </xs:sequence>
        <xs:attribute name="actuatorId" type="xs:string" use="required"/>
    </xs:complexType>

    <!-- 设备数据包 -->
    <xs:complexType name="DeviceDataPacketType">
        <xs:sequence>
            <xs:element name="sensors" type="tns:SensorDataType" 
                       minOccurs="1" maxOccurs="unbounded"/>
            <xs:element name="actuators" type="tns:ActuatorControlType" 
                       minOccurs="0" maxOccurs="unbounded"/>
        </xs:sequence>
        <xs:attribute name="deviceId" type="xs:string" use="required"/>
        <xs:attribute name="status" type="tns:DeviceStatusType" use="required"/>
        <xs:attribute name="sequenceNumber" type="xs:long" use="required"/>
    </xs:complexType>

    <xs:element name="deviceData" type="tns:DeviceDataPacketType"/>

</xs:schema>
```

### 代码生成命令

```bash
# 生成PLC优化的Go代码
xsd2code \
  -xsd=plc-data.xsd \
  -lang=go \
  -output=plc/types.go \
  -package=plc \
  -plc \
  -validation \
  -validation-output=plc/validation.go \
  -benchmarks

# 生成C#代码（Windows HMI系统）
xsd2code \
  -xsd=plc-data.xsd \
  -lang=csharp \
  -output=HMI/PLCTypes.cs \
  -package=Automation.PLC
```

## 📊 复杂数据分析系统

### 场景描述
为大数据分析平台设计多层次的数据结构，支持复杂的数据关系和元数据。

### XSD设计

```xml
<?xml version="1.0" encoding="UTF-8"?>
<xs:schema xmlns:xs="http://www.w3.org/2001/XMLSchema"
           targetNamespace="http://analytics.company.com/data"
           xmlns:tns="http://analytics.company.com/data"
           xmlns:meta="http://analytics.company.com/metadata"
           elementFormDefault="qualified">

    <!-- 导入元数据schema -->
    <xs:import namespace="http://analytics.company.com/metadata" 
               schemaLocation="metadata.xsd"/>

    <!-- 数据类型枚举 -->
    <xs:simpleType name="DataTypeEnum">
        <xs:restriction base="xs:string">
            <xs:enumeration value="STRING"/>
            <xs:enumeration value="INTEGER"/>
            <xs:enumeration value="FLOAT"/>
            <xs:enumeration value="BOOLEAN"/>
            <xs:enumeration value="DATETIME"/>
            <xs:enumeration value="BINARY"/>
        </xs:restriction>
    </xs:simpleType>

    <!-- 数据源类型 -->
    <xs:simpleType name="DataSourceTypeEnum">
        <xs:restriction base="xs:string">
            <xs:enumeration value="DATABASE"/>
            <xs:enumeration value="FILE"/>
            <xs:enumeration value="API"/>
            <xs:enumeration value="STREAM"/>
        </xs:restriction>
    </xs:simpleType>

    <!-- 数据字段定义 -->
    <xs:complexType name="DataFieldType">
        <xs:sequence>
            <xs:element name="name" type="xs:string"/>
            <xs:element name="dataType" type="tns:DataTypeEnum"/>
            <xs:element name="nullable" type="xs:boolean" default="true"/>
            <xs:element name="defaultValue" type="xs:string" minOccurs="0"/>
            <xs:element name="constraints" minOccurs="0">
                <xs:complexType>
                    <xs:sequence>
                        <xs:element name="minLength" type="xs:int" minOccurs="0"/>
                        <xs:element name="maxLength" type="xs:int" minOccurs="0"/>
                        <xs:element name="pattern" type="xs:string" minOccurs="0"/>
                        <xs:element name="minValue" type="xs:decimal" minOccurs="0"/>
                        <xs:element name="maxValue" type="xs:decimal" minOccurs="0"/>
                    </xs:sequence>
                </xs:complexType>
            </xs:element>
            <xs:element name="metadata" type="meta:MetadataType" minOccurs="0"/>
        </xs:sequence>
        <xs:attribute name="fieldId" type="xs:string" use="required"/>
        <xs:attribute name="primaryKey" type="xs:boolean" default="false"/>
    </xs:complexType>

    <!-- 数据表结构 -->
    <xs:complexType name="DataTableSchemaType">
        <xs:sequence>
            <xs:element name="tableName" type="xs:string"/>
            <xs:element name="description" type="xs:string" minOccurs="0"/>
            <xs:element name="fields" type="tns:DataFieldType" 
                       minOccurs="1" maxOccurs="unbounded"/>
            <xs:element name="indexes" minOccurs="0">
                <xs:complexType>
                    <xs:sequence>
                        <xs:element name="index" maxOccurs="unbounded">
                            <xs:complexType>
                                <xs:sequence>
                                    <xs:element name="fieldName" type="xs:string" 
                                               maxOccurs="unbounded"/>
                                </xs:sequence>
                                <xs:attribute name="name" type="xs:string" use="required"/>
                                <xs:attribute name="unique" type="xs:boolean" default="false"/>
                            </xs:complexType>
                        </xs:element>
                    </xs:sequence>
                </xs:complexType>
            </xs:element>
        </xs:sequence>
        <xs:attribute name="tableId" type="xs:string" use="required"/>
        <xs:attribute name="version" type="xs:string" use="required"/>
    </xs:complexType>

    <!-- 数据源配置 -->
    <xs:complexType name="DataSourceType">
        <xs:sequence>
            <xs:element name="name" type="xs:string"/>
            <xs:element name="type" type="tns:DataSourceTypeEnum"/>
            <xs:element name="connectionString" type="xs:string"/>
            <xs:element name="credentials" minOccurs="0">
                <xs:complexType>
                    <xs:sequence>
                        <xs:element name="username" type="xs:string"/>
                        <xs:element name="passwordHash" type="xs:string"/>
                    </xs:sequence>
                </xs:complexType>
            </xs:element>
            <xs:element name="tables" type="tns:DataTableSchemaType" 
                       minOccurs="0" maxOccurs="unbounded"/>
        </xs:sequence>
        <xs:attribute name="sourceId" type="xs:string" use="required"/>
        <xs:attribute name="enabled" type="xs:boolean" default="true"/>
    </xs:complexType>

    <!-- 分析项目配置 -->
    <xs:complexType name="AnalyticsProjectType">
        <xs:sequence>
            <xs:element name="projectName" type="xs:string"/>
            <xs:element name="description" type="xs:string" minOccurs="0"/>
            <xs:element name="dataSources" type="tns:DataSourceType" 
                       minOccurs="1" maxOccurs="unbounded"/>
            <xs:element name="transformations" minOccurs="0">
                <xs:complexType>
                    <xs:sequence>
                        <xs:element name="transformation" maxOccurs="unbounded">
                            <xs:complexType>
                                <xs:sequence>
                                    <xs:element name="sourceTable" type="xs:string"/>
                                    <xs:element name="targetTable" type="xs:string"/>
                                    <xs:element name="script" type="xs:string"/>
                                </xs:sequence>
                                <xs:attribute name="name" type="xs:string" use="required"/>
                            </xs:complexType>
                        </xs:element>
                    </xs:sequence>
                </xs:complexType>
            </xs:element>
        </xs:sequence>
        <xs:attribute name="projectId" type="xs:string" use="required"/>
        <xs:attribute name="version" type="xs:string" use="required"/>
        <xs:attribute name="createdBy" type="xs:string" use="required"/>
        <xs:attribute name="createdAt" type="xs:dateTime" use="required"/>
    </xs:complexType>

    <xs:element name="analyticsProject" type="tns:AnalyticsProjectType"/>

</xs:schema>
```

### 批量代码生成脚本

```bash
#!/bin/bash

# 批量生成多语言代码
SCHEMAS=("analytics-project.xsd" "metadata.xsd")
LANGUAGES=("go" "java" "csharp" "typescript")
OUTPUT_DIR="generated"

# 创建输出目录
mkdir -p $OUTPUT_DIR

for schema in "${SCHEMAS[@]}"; do
    base_name=$(basename "$schema" .xsd)
    
    for lang in "${LANGUAGES[@]}"; do
        case $lang in
            "go")
                xsd2code \
                  -xsd="schemas/$schema" \
                  -lang=go \
                  -output="$OUTPUT_DIR/go/${base_name}.go" \
                  -package=models \
                  -json \
                  -validation \
                  -validation-output="$OUTPUT_DIR/go/${base_name}_validation.go"
                ;;
            "java")
                xsd2code \
                  -xsd="schemas/$schema" \
                  -lang=java \
                  -output="$OUTPUT_DIR/java/${base_name^}.java" \
                  -package=com.company.analytics.models
                ;;
            "csharp")
                xsd2code \
                  -xsd="schemas/$schema" \
                  -lang=csharp \
                  -output="$OUTPUT_DIR/csharp/${base_name^}.cs" \
                  -package=Company.Analytics.Models
                ;;
            "typescript")
                xsd2code \
                  -xsd="schemas/$schema" \
                  -lang=typescript \
                  -output="$OUTPUT_DIR/typescript/${base_name}.ts"
                ;;
        esac
    done
done

echo "代码生成完成！"
```

## 🔄 CI/CD集成示例

### GitHub Actions工作流

```yaml
name: XSD Code Generation

on:
  push:
    paths:
      - 'schemas/**/*.xsd'
  pull_request:
    paths:
      - 'schemas/**/*.xsd'

jobs:
  generate-code:
    runs-on: ubuntu-latest
    
    steps:
    - uses: actions/checkout@v3
    
    - name: Set up Go
      uses: actions/setup-go@v3
      with:
        go-version: 1.21
    
    - name: Install XSD2Code
      run: |
        git clone https://github.com/suifei/xsd2code.git
        cd xsd2code
        go build -o xsd2code cmd/main.go
        sudo cp xsd2code /usr/local/bin/
    
    - name: Generate Go Code
      run: |
        mkdir -p generated/go
        for xsd in schemas/*.xsd; do
          base=$(basename "$xsd" .xsd)
          xsd2code \
            -xsd="$xsd" \
            -output="generated/go/${base}.go" \
            -package=models \
            -json \
            -validation \
            -validation-output="generated/go/${base}_validation.go"
        done
    
    - name: Generate Java Code
      run: |
        mkdir -p generated/java
        for xsd in schemas/*.xsd; do
          base=$(basename "$xsd" .xsd)
          xsd2code \
            -xsd="$xsd" \
            -lang=java \
            -output="generated/java/${base^}.java" \
            -package=com.company.models
        done
    
    - name: Verify Generated Code
      run: |
        # 验证Go代码
        cd generated/go
        go mod init generated
        go mod tidy
        go build ./...
        
        # 验证Java代码
        cd ../java
        javac *.java
    
    - name: Commit Generated Code
      run: |
        git config --local user.email "action@github.com"
        git config --local user.name "GitHub Action"
        git add generated/
        git diff --staged --quiet || git commit -m "Auto-generate code from XSD changes"
        git push
```

### Makefile集成

```makefile
# Makefile for XSD2Code project

.PHONY: generate clean verify docker

# 默认目标
all: generate verify

# 生成所有代码
generate:
	@echo "Generating code from XSD files..."
	@mkdir -p generated/{go,java,csharp,typescript}
	
	@for xsd in schemas/*.xsd; do \
		base=$$(basename "$$xsd" .xsd); \
		echo "Processing $$base.xsd..."; \
		\
		echo "  - Generating Go code..."; \
		xsd2code \
			-xsd="$$xsd" \
			-output="generated/go/$${base}.go" \
			-package=models \
			-json \
			-validation \
			-validation-output="generated/go/$${base}_validation.go"; \
		\
		echo "  - Generating Java code..."; \
		xsd2code \
			-xsd="$$xsd" \
			-lang=java \
			-output="generated/java/$${base^}.java" \
			-package=com.company.models; \
		\
		echo "  - Generating C# code..."; \
		xsd2code \
			-xsd="$$xsd" \
			-lang=csharp \
			-output="generated/csharp/$${base^}.cs" \
			-package=Company.Models; \
		\
		echo "  - Generating TypeScript code..."; \
		xsd2code \
			-xsd="$$xsd" \
			-lang=typescript \
			-output="generated/typescript/$${base}.ts"; \
	done
	@echo "Code generation completed!"

# 验证生成的代码
verify:
	@echo "Verifying generated code..."
	
	@echo "  - Verifying Go code..."
	@cd generated/go && \
		go mod init generated 2>/dev/null || true && \
		go mod tidy && \
		go build ./...
	
	@echo "  - Verifying Java code..."
	@cd generated/java && \
		javac *.java
	
	@echo "Code verification completed!"

# 清理生成的文件
clean:
	@echo "Cleaning generated files..."
	@rm -rf generated/
	@echo "Clean completed!"

# Docker构建
docker:
	@echo "Building Docker image..."
	@docker build -t xsd2code-generator .
	@echo "Docker build completed!"

# 运行测试
test:
	@echo "Running tests..."
	@cd generated/go && go test ./...
	@echo "Tests completed!"

# 显示统计信息
stats:
	@echo "Code Generation Statistics:"
	@echo "  XSD files: $$(find schemas -name '*.xsd' | wc -l)"
	@echo "  Generated Go files: $$(find generated/go -name '*.go' 2>/dev/null | wc -l)"
	@echo "  Generated Java files: $$(find generated/java -name '*.java' 2>/dev/null | wc -l)"
	@echo "  Generated C# files: $$(find generated/csharp -name '*.cs' 2>/dev/null | wc -l)"
	@echo "  Generated TypeScript files: $$(find generated/typescript -name '*.ts' 2>/dev/null | wc -l)"

# 帮助信息
help:
	@echo "Available targets:"
	@echo "  generate  - Generate code from all XSD files"
	@echo "  verify    - Verify that generated code compiles"
	@echo "  clean     - Remove all generated files"
	@echo "  test      - Run tests on generated code"
	@echo "  docker    - Build Docker image"
	@echo "  stats     - Show generation statistics"
	@echo "  help      - Show this help message"
```

## 🎯 性能优化示例

### 大型XSD文件处理

```bash
# 处理大型XSD文件的优化策略

# 1. 分模块处理
xsd2code -xsd=large-schema-core.xsd -output=core_types.go -package=core
xsd2code -xsd=large-schema-extensions.xsd -output=ext_types.go -package=extensions

# 2. 禁用不必要的功能以提高速度
xsd2code \
  -xsd=large-schema.xsd \
  -output=types.go \
  -comments=false \
  -validation=false

# 3. 使用严格模式进行验证
xsd2code -xsd=large-schema.xsd -strict -debug > validation.log 2>&1
```

---

💡 **提示**: 这些高级示例展示了XSD2Code在实际项目中的强大功能。根据您的具体需求，可以组合使用不同的选项和配置来达到最佳效果。
