# 最佳实践 - 使用建议和最佳实践

本页面汇总了使用XSD2Code的最佳实践，帮助您避免常见陷阱，提高开发效率，并生成高质量的代码。

## 🎯 项目组织最佳实践

### 1. 目录结构建议

```
project/
├── schemas/                 # XSD文件目录
│   ├── common/             # 公共类型定义
│   │   ├── base-types.xsd
│   │   └── enums.xsd
│   ├── api/                # API相关schema
│   │   ├── user-api.xsd
│   │   └── order-api.xsd
│   └── data/               # 数据模型schema
│       ├── user.xsd
│       └── product.xsd
├── generated/              # 生成的代码目录
│   ├── go/
│   ├── java/
│   └── typescript/
├── scripts/               # 构建脚本
│   ├── generate.sh
│   └── validate.sh
└── configs/              # 配置文件
    ├── dev.yaml
    └── prod.yaml
```

### 2. 命名约定

#### XSD文件命名
```bash
# 建议的命名规范
user-types.xsd          # 用户相关类型
order-api-v2.xsd        # 订单API v2
common-enums.xsd        # 公共枚举
product-catalog.xsd     # 产品目录
```

#### 生成代码命名
```bash
# Go
user_types.go           # 用户类型
order_models.go         # 订单模型

# Java
UserTypes.java          # 用户类型
OrderModels.java        # 订单模型

# C#
UserTypes.cs            # 用户类型
OrderModels.cs          # 订单模型
```

## 🏗️ XSD设计最佳实践

### 1. 模块化设计

```xml
<!-- 基础类型文件: common-types.xsd -->
<?xml version="1.0" encoding="UTF-8"?>
<xs:schema xmlns:xs="http://www.w3.org/2001/XMLSchema"
           targetNamespace="http://company.com/common"
           xmlns:tns="http://company.com/common">

    <!-- 可重用的简单类型 -->
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
</xs:schema>
```

```xml
<!-- 业务模型文件: user.xsd -->
<?xml version="1.0" encoding="UTF-8"?>
<xs:schema xmlns:xs="http://www.w3.org/2001/XMLSchema"
           targetNamespace="http://company.com/user"
           xmlns:tns="http://company.com/user"
           xmlns:common="http://company.com/common">

    <xs:import namespace="http://company.com/common" 
               schemaLocation="common-types.xsd"/>

    <xs:complexType name="UserType">
        <xs:sequence>
            <xs:element name="id" type="xs:string"/>
            <xs:element name="name" type="xs:string"/>
            <xs:element name="email" type="common:EmailType"/>
            <xs:element name="phone" type="common:PhoneType" minOccurs="0"/>
        </xs:sequence>
    </xs:complexType>
</xs:schema>
```

### 2. 版本管理策略

```xml
<!-- 建议的版本管理方法 -->
<xs:schema targetNamespace="http://company.com/api/v2"
           xmlns:tns="http://company.com/api/v2">
    
    <!-- 明确的版本信息 -->
    <xs:annotation>
        <xs:documentation>
            API Version: 2.1.0
            Last Modified: 2024-01-15
            Breaking Changes: None since 2.0.0
        </xs:documentation>
    </xs:annotation>
</xs:schema>
```

### 3. 约束设计原则

```xml
<!-- 合理的约束设计 -->
<xs:simpleType name="UserIdType">
    <xs:restriction base="xs:string">
        <!-- 具体的模式定义 -->
        <xs:pattern value="USR[0-9]{8}"/>
        <!-- 明确的长度限制 -->
        <xs:length value="11"/>
        <!-- 空白字符处理 -->
        <xs:whiteSpace value="collapse"/>
    </xs:restriction>
</xs:simpleType>

<xs:simpleType name="UsernameType">
    <xs:restriction base="xs:string">
        <!-- 合理的长度范围 -->
        <xs:minLength value="3"/>
        <xs:maxLength value="50"/>
        <!-- 字符集限制 -->
        <xs:pattern value="[a-zA-Z0-9_-]+"/>
    </xs:restriction>
</xs:simpleType>
```

## 🔧 代码生成最佳实践

### 1. 配置文件管理

```yaml
# production.yaml - 生产环境配置
environment: production

defaults:
  language: go
  features:
    validation: true
    json_tags: true
    xml_tags: true
    comments: true
    strict_mode: true

go:
  package_prefix: "company.com/api"
  pointer_types: false
  omitempty: true
  builder_pattern: false

validation:
  level: full
  error_language: "en"
  custom_messages: true
```

```yaml
# development.yaml - 开发环境配置
environment: development

defaults:
  language: go
  features:
    validation: true
    json_tags: true
    xml_tags: true
    comments: true
    debug_mode: true

go:
  package_prefix: "myapp"
  pointer_types: true
  omitempty: true
  builder_pattern: true
```

### 2. 自动化构建脚本

```bash
#!/bin/bash
# generate.sh - 代码生成脚本

set -e

CONFIG_ENV=${1:-development}
CONFIG_FILE="configs/${CONFIG_ENV}.yaml"

echo "🚀 开始代码生成..."
echo "📝 使用配置: ${CONFIG_FILE}"

# 验证XSD文件
echo "🔍 验证XSD文件..."
for xsd_file in schemas/**/*.xsd; do
    xmllint --schema "$xsd_file" --noout || {
        echo "❌ XSD验证失败: $xsd_file"
        exit 1
    }
done

# 清理旧的生成文件
echo "🧹 清理旧文件..."
rm -rf generated/
mkdir -p generated/{go,java,typescript,csharp}

# 生成Go代码
echo "🐹 生成Go代码..."
xsd2code -config="$CONFIG_FILE" \
         -xsd=schemas/api/*.xsd \
         -lang=go \
         -output-dir=generated/go \
         -package=api

# 生成Java代码
echo "☕ 生成Java代码..."
xsd2code -config="$CONFIG_FILE" \
         -xsd=schemas/api/*.xsd \
         -lang=java \
         -output-dir=generated/java \
         -package=com.company.api.models

# 生成TypeScript代码
echo "📘 生成TypeScript代码..."
xsd2code -config="$CONFIG_FILE" \
         -xsd=schemas/api/*.xsd \
         -lang=typescript \
         -output-dir=generated/typescript

# 验证生成的代码
echo "✅ 验证生成的代码..."
cd generated/go && go fmt ./... && go vet ./...
cd ../java && javac -cp ".:*" *.java
cd ../typescript && tsc --noEmit *.ts

echo "🎉 代码生成完成!"
```

### 3. Makefile集成

```makefile
# Makefile
.PHONY: generate clean validate test

# 默认环境
ENV ?= development

# 生成所有代码
generate:
	@echo "🚀 生成代码 (环境: $(ENV))"
	@./scripts/generate.sh $(ENV)

# 清理生成的文件
clean:
	@echo "🧹 清理生成的文件"
	@rm -rf generated/

# 验证XSD文件
validate:
	@echo "🔍 验证XSD文件"
	@for file in schemas/**/*.xsd; do \
		echo "验证: $$file"; \
		xmllint --schema "$$file" --noout; \
	done

# 测试生成的代码
test: generate
	@echo "🧪 测试生成的代码"
	@cd generated/go && go test ./...
	@cd generated/java && mvn test
	@cd generated/typescript && npm test

# CI/CD构建
ci: clean validate generate test
	@echo "✅ CI构建完成"

# 生产环境构建
prod:
	@$(MAKE) generate ENV=production
```

## 🏭 CI/CD集成最佳实践

### 1. GitHub Actions示例

```yaml
# .github/workflows/generate-code.yml
name: Generate Code from XSD

on:
  push:
    paths:
      - 'schemas/**/*.xsd'
      - 'configs/**/*.yaml'
  pull_request:
    paths:
      - 'schemas/**/*.xsd'

jobs:
  generate-and-test:
    runs-on: ubuntu-latest
    
    steps:
    - name: Checkout
      uses: actions/checkout@v3
    
    - name: Setup Go
      uses: actions/setup-go@v3
      with:
        go-version: '1.21'
    
    - name: Setup Java
      uses: actions/setup-java@v3
      with:
        distribution: 'temurin'
        java-version: '17'
    
    - name: Setup Node.js
      uses: actions/setup-node@v3
      with:
        node-version: '18'
    
    - name: Install XSD2Code
      run: |
        go install github.com/suifei/xsd2code@latest
    
    - name: Validate XSD Files
      run: |
        sudo apt-get install -y libxml2-utils
        make validate
    
    - name: Generate Code
      run: make generate ENV=ci
    
    - name: Test Generated Code
      run: make test
    
    - name: Upload Generated Code
      uses: actions/upload-artifact@v3
      with:
        name: generated-code
        path: generated/
```

### 2. Jenkins Pipeline示例

```groovy
// Jenkinsfile
pipeline {
    agent any
    
    environment {
        GO_VERSION = '1.21'
        JAVA_VERSION = '17'
    }
    
    stages {
        stage('Checkout') {
            steps {
                checkout scm
            }
        }
        
        stage('Setup') {
            parallel {
                stage('Setup Go') {
                    steps {
                        sh 'go version'
                        sh 'go install github.com/suifei/xsd2code@latest'
                    }
                }
                stage('Setup Java') {
                    steps {
                        sh 'java -version'
                    }
                }
            }
        }
        
        stage('Validate') {
            steps {
                sh 'make validate'
            }
        }
        
        stage('Generate') {
            steps {
                sh 'make generate ENV=ci'
            }
        }
        
        stage('Test') {
            steps {
                sh 'make test'
            }
            post {
                always {
                    publishTestResults testResultsPattern: '**/test-results.xml'
                }
            }
        }
        
        stage('Archive') {
            steps {
                archiveArtifacts artifacts: 'generated/**/*', fingerprint: true
            }
        }
    }
    
    post {
        always {
            cleanWs()
        }
        failure {
            emailext (
                subject: "XSD2Code Generation Failed: ${env.JOB_NAME} - ${env.BUILD_NUMBER}",
                body: "Build failed. Check console output at ${env.BUILD_URL}",
                to: "${env.CHANGE_AUTHOR_EMAIL}"
            )
        }
    }
}
```

## 🧪 测试最佳实践

### 1. 自动化测试生成

```bash
# 生成包含测试代码的结构
xsd2code -xsd=schema.xsd \
         -lang=go \
         -output=types.go \
         -test=true \
         -test-output=types_test.go
```

### 2. 验证测试示例

```go
// types_test.go - 自动生成的测试代码
package models_test

import (
    "testing"
    "encoding/json"
    "encoding/xml"
    "github.com/stretchr/testify/assert"
    "myapp/models"
)

func TestUserTypeValidation(t *testing.T) {
    tests := []struct {
        name    string
        user    models.UserType
        wantErr bool
    }{
        {
            name: "valid user",
            user: models.UserType{
                ID:    "USR12345678",
                Name:  "John Doe",
                Email: "john@example.com",
            },
            wantErr: false,
        },
        {
            name: "invalid email",
            user: models.UserType{
                ID:    "USR12345678",
                Name:  "John Doe",
                Email: "invalid-email",
            },
            wantErr: true,
        },
        // 更多测试案例...
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            err := tt.user.Validate()
            if tt.wantErr {
                assert.Error(t, err)
            } else {
                assert.NoError(t, err)
            }
        })
    }
}

func TestUserTypeSerialization(t *testing.T) {
    user := models.UserType{
        ID:    "USR12345678",
        Name:  "John Doe",
        Email: "john@example.com",
    }

    // JSON序列化测试
    jsonData, err := json.Marshal(user)
    assert.NoError(t, err)
    
    var userFromJSON models.UserType
    err = json.Unmarshal(jsonData, &userFromJSON)
    assert.NoError(t, err)
    assert.Equal(t, user, userFromJSON)

    // XML序列化测试
    xmlData, err := xml.Marshal(user)
    assert.NoError(t, err)
    
    var userFromXML models.UserType
    err = xml.Unmarshal(xmlData, &userFromXML)
    assert.NoError(t, err)
    assert.Equal(t, user, userFromXML)
}
```

## 🔄 维护和更新最佳实践

### 1. 版本控制策略

```bash
# 版本化的XSD文件
schemas/
├── v1/
│   ├── user.xsd
│   └── order.xsd
├── v2/
│   ├── user.xsd      # 向后兼容的更新
│   └── order.xsd
└── current -> v2/    # 符号链接指向当前版本
```

### 2. 向后兼容性检查

```bash
#!/bin/bash
# check-compatibility.sh

OLD_VERSION=${1:-v1}
NEW_VERSION=${2:-v2}

echo "🔍 检查从 $OLD_VERSION 到 $NEW_VERSION 的兼容性..."

# 生成两个版本的代码
xsd2code -xsd=schemas/$OLD_VERSION/*.xsd -output-dir=temp/old
xsd2code -xsd=schemas/$NEW_VERSION/*.xsd -output-dir=temp/new

# 使用工具检查API兼容性
go-apidiff temp/old temp/new

# 清理临时文件
rm -rf temp/
```

### 3. 文档同步

```yaml
# docs-sync.yaml - 文档同步配置
sync_rules:
  - trigger: xsd_change
    actions:
      - regenerate_code
      - update_api_docs
      - run_tests
      - notify_team

  - trigger: config_change
    actions:
      - validate_config
      - regenerate_code
      - update_examples

notifications:
  slack:
    channel: "#api-changes"
    webhook: "${SLACK_WEBHOOK_URL}"
  
  email:
    recipients:
      - "dev-team@company.com"
      - "api-consumers@company.com"
```

## 📊 性能优化最佳实践

### 1. 大型XSD处理

```bash
# 处理大型XSD文件的策略
xsd2code -xsd=large-schema.xsd \
         -memory-limit=4GB \
         -parallel=true \
         -workers=8 \
         -cache=true \
         -streaming=true
```

### 2. 批量处理优化

```yaml
# batch-config.yaml
batch_processing:
  parallel_workers: 4
  memory_per_worker: "1GB"
  cache_enabled: true
  
  optimization:
    reuse_parsers: true
    incremental_generation: true
    dependency_analysis: true

files:
  - group: "common"
    priority: 1
    files: ["schemas/common/*.xsd"]
  
  - group: "api"
    priority: 2
    depends_on: ["common"]
    files: ["schemas/api/*.xsd"]
  
  - group: "internal"
    priority: 3
    depends_on: ["common", "api"]
    files: ["schemas/internal/*.xsd"]
```

## 🔐 安全最佳实践

### 1. 敏感信息处理

```xml
<!-- 避免在XSD中硬编码敏感信息 -->
<xs:simpleType name="ApiKeyType">
    <xs:restriction base="xs:string">
        <!-- 不要暴露具体的密钥格式 -->
        <xs:pattern value="[A-Za-z0-9]{32,64}"/>
    </xs:restriction>
</xs:simpleType>
```

### 2. 验证安全

```go
// 安全的验证实现
func (u *UserType) ValidateSecurely() error {
    // 防止注入攻击
    if containsSQLInjection(u.Name) {
        return errors.New("invalid characters in name")
    }
    
    // 防止XSS攻击
    if containsXSSPatterns(u.Comment) {
        return errors.New("invalid characters in comment")
    }
    
    return u.Validate()
}

func containsSQLInjection(input string) bool {
    patterns := []string{
        `(?i)(union|select|insert|update|delete|drop|create|alter)`,
        `(?i)(script|javascript|vbscript)`,
        `[<>]`,
    }
    
    for _, pattern := range patterns {
        if matched, _ := regexp.MatchString(pattern, input); matched {
            return true
        }
    }
    return false
}
```

## 📝 文档最佳实践

### 1. 内联文档

```xml
<!-- 良好的XSD文档示例 -->
<xs:complexType name="UserType">
    <xs:annotation>
        <xs:documentation>
            用户类型定义
            - 用于用户管理系统
            - 支持完整的用户信息
            - 包含验证规则
            版本: 2.1.0
        </xs:documentation>
    </xs:annotation>
    
    <xs:sequence>
        <xs:element name="id" type="tns:UserIdType">
            <xs:annotation>
                <xs:documentation>
                    用户唯一标识符
                    格式: USR + 8位数字
                    示例: USR12345678
                </xs:documentation>
            </xs:annotation>
        </xs:element>
    </xs:sequence>
</xs:complexType>
```

### 2. README模板

```markdown
# 项目名称 - 生成的类型定义

## 概述
本目录包含从XSD文件生成的类型定义。

## 文件说明
- `user_types.go` - 用户相关类型定义
- `order_types.go` - 订单相关类型定义
- `common_types.go` - 公共类型定义

## 生成信息
- **XSD2Code版本**: v3.1.0
- **生成时间**: 2024-01-15 10:30:00
- **源XSD文件**: schemas/api/*.xsd
- **配置文件**: configs/production.yaml

## 使用方法
```go
import "myapp/generated/types"

user := &types.UserType{
    ID:    "USR12345678",
    Name:  "John Doe",
    Email: "john@example.com",
}

if err := user.Validate(); err != nil {
    log.Fatal(err)
}
```

## 注意事项
- 这些文件是自动生成的，请勿手动修改
- 如需更改，请修改源XSD文件
- 重新生成前请备份自定义代码

## 相关文档
- [XSD Schema文档](../docs/schema.md)
- [API使用指南](../docs/api.md)
- [开发指南](../docs/development.md)
```

---

💡 **关键要点**: 
- 始终使用版本控制管理XSD文件
- 配置文件化管理以支持不同环境
- 自动化构建和测试流程
- 定期验证生成代码的质量
- 保持文档与代码的同步

🔗 **相关页面**: 
- [[配置选项|Configuration]] - 详细配置说明
- [[多语言支持|Multi-Language-Support]] - 语言特定最佳实践
- [[约束和验证|Constraints-and-Validation]] - 验证最佳实践
