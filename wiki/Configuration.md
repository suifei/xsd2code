# 配置选项 - 高级配置和自定义选项

XSD2Code提供了丰富的配置选项，让您可以精确控制代码生成的各个方面，满足不同项目的特定需求。

## ⚙️ 核心配置选项

### 基本配置

| 参数 | 类型 | 默认值 | 描述 |
|------|------|--------|------|
| `-xsd` | string | **必需** | XSD文件路径 |
| `-output` | string | `generated.go` | 输出文件名 |
| `-lang` | string | `go` | 目标语言 (go, java, csharp, typescript) |
| `-package` | string | `models` | 包名/命名空间 |
| `-debug` | bool | `false` | 启用调试模式 |

### 高级配置

| 参数 | 类型 | 默认值 | 描述 |
|------|------|--------|------|
| `-validation` | bool | `true` | 生成验证代码 |
| `-json` | bool | `true` | 生成JSON序列化标签 |
| `-xml` | bool | `true` | 生成XML序列化标签 |
| `-comments` | bool | `true` | 生成代码注释 |
| `-strict` | bool | `false` | 严格模式 |
| `-omitempty` | bool | `false` | 生成omitempty标签 |

## 🏗️ 语言特定配置

### Go语言配置

```bash
# 基本Go配置
xsd2code -xsd=schema.xsd \
         -lang=go \
         -output=types.go \
         -package=models

# 高级Go配置
xsd2code -xsd=schema.xsd \
         -lang=go \
         -output=types.go \
         -package=models \
         -validation=true \
         -json=true \
         -omitempty=true \
         -pointer=false \
         -builder=false
```

#### Go特定选项

| 参数 | 描述 | 示例 |
|------|------|------|
| `-pointer` | 使用指针类型 | `*string` vs `string` |
| `-builder` | 生成Builder模式 | `NewUser().SetName("John")` |
| `-interface` | 生成接口定义 | `type Validator interface` |
| `-tag-style` | 标签风格 | `snake_case`, `camelCase` |

```bash
# 生成指针类型字段
xsd2code -xsd=schema.xsd -lang=go -pointer=true

# 生成Builder模式
xsd2code -xsd=schema.xsd -lang=go -builder=true

# 自定义标签风格
xsd2code -xsd=schema.xsd -lang=go -tag-style=snake_case
```

### Java语言配置

```bash
# 基本Java配置
xsd2code -xsd=schema.xsd \
         -lang=java \
         -output=Types.java \
         -package=com.example.models

# 高级Java配置
xsd2code -xsd=schema.xsd \
         -lang=java \
         -output=Types.java \
         -package=com.example.models \
         -annotations=true \
         -builder=true \
         -lombok=false \
         -jaxb=true
```

#### Java特定选项

| 参数 | 描述 | 示例 |
|------|------|------|
| `-annotations` | 生成验证注解 | `@NotNull`, `@Size` |
| `-jaxb` | 生成JAXB注解 | `@XmlElement` |
| `-lombok` | 使用Lombok注解 | `@Data`, `@Builder` |
| `-constructor` | 生成构造函数 | 默认、全参构造 |

```bash
# 启用Bean验证注解
xsd2code -xsd=schema.xsd -lang=java -annotations=true

# 使用Lombok
xsd2code -xsd=schema.xsd -lang=java -lombok=true

# 生成Builder模式
xsd2code -xsd=schema.xsd -lang=java -builder=true
```

### C#语言配置

```bash
# 基本C#配置
xsd2code -xsd=schema.xsd \
         -lang=csharp \
         -output=Types.cs \
         -namespace=Example.Models

# 高级C#配置
xsd2code -xsd=schema.xsd \
         -lang=csharp \
         -output=Types.cs \
         -namespace=Example.Models \
         -nullable=true \
         -properties=true \
         -attributes=true
```

#### C#特定选项

| 参数 | 描述 | 示例 |
|------|------|------|
| `-nullable` | 启用可空引用类型 | `string?` |
| `-properties` | 生成属性而非字段 | `public string Name { get; set; }` |
| `-attributes` | 生成DataAnnotations | `[Required]`, `[StringLength]` |
| `-records` | 生成Record类型 | `public record User(string Name)` |

```bash
# 启用可空引用类型
xsd2code -xsd=schema.xsd -lang=csharp -nullable=true

# 生成Record类型
xsd2code -xsd=schema.xsd -lang=csharp -records=true

# 生成DataAnnotations
xsd2code -xsd=schema.xsd -lang=csharp -attributes=true
```

### TypeScript配置

```bash
# 基本TypeScript配置
xsd2code -xsd=schema.xsd \
         -lang=typescript \
         -output=types.ts

# 高级TypeScript配置
xsd2code -xsd=schema.xsd \
         -lang=typescript \
         -output=types.ts \
         -interface=true \
         -optional=true \
         -export=true
```

#### TypeScript特定选项

| 参数 | 描述 | 示例 |
|------|------|------|
| `-interface` | 生成接口而非类 | `interface User` |
| `-optional` | 可选字段处理 | `name?: string` |
| `-export` | 导出声明 | `export interface` |
| `-namespace` | 使用命名空间 | `namespace Models` |

## 🎯 输出控制配置

### 文件输出选项

```bash
# 单文件输出
xsd2code -xsd=schema.xsd -output=types.go

# 按类型分文件输出
xsd2code -xsd=schema.xsd -output-dir=./models -split=true

# 自定义文件名模板
xsd2code -xsd=schema.xsd -output-template="%s_generated.go"
```

### 代码格式配置

```bash
# 禁用代码注释
xsd2code -xsd=schema.xsd -comments=false

# 自定义缩进
xsd2code -xsd=schema.xsd -indent="    " # 4个空格

# 自定义行结束符
xsd2code -xsd=schema.xsd -line-ending="\n"
```

## 🔧 验证配置

### 验证级别

```bash
# 基础验证
xsd2code -xsd=schema.xsd -validation=basic

# 完整验证
xsd2code -xsd=schema.xsd -validation=full

# 禁用验证
xsd2code -xsd=schema.xsd -validation=false

# 自定义验证
xsd2code -xsd=schema.xsd -validation=custom -validation-config=config.yaml
```

### 验证选项

| 参数 | 描述 | 默认值 |
|------|------|--------|
| `-validation-errors` | 错误处理方式 | `return` |
| `-validation-lang` | 错误消息语言 | `en` |
| `-validation-strict` | 严格验证模式 | `false` |

```bash
# 中文错误消息
xsd2code -xsd=schema.xsd -validation=true -validation-lang=zh

# 严格验证模式
xsd2code -xsd=schema.xsd -validation=true -validation-strict=true
```

## 📄 配置文件支持

### YAML配置文件

创建 `xsd2code.yaml`:

```yaml
# xsd2code.yaml
input:
  xsd: schema.xsd
  
output:
  file: types.go
  directory: ./generated
  split: false
  
language:
  target: go
  package: models
  
features:
  validation: true
  json_tags: true
  xml_tags: true
  comments: true
  builder_pattern: false
  
go:
  pointer_types: false
  omitempty: true
  tag_style: camelCase
  interface_generation: false
  
java:
  annotations: true
  jaxb: true
  lombok: false
  builder: true
  
csharp:
  nullable: true
  properties: true
  attributes: true
  records: false
  
typescript:
  interfaces: true
  optional_fields: true
  export_declarations: true
  namespace: false
  
validation:
  level: full
  error_language: zh
  strict_mode: false
  custom_messages: true
  
formatting:
  indent: "  "
  line_ending: "\n"
  max_line_length: 120
```

使用配置文件:

```bash
# 使用配置文件
xsd2code -config=xsd2code.yaml

# 配置文件 + 命令行参数覆盖
xsd2code -config=xsd2code.yaml -lang=java -validation=false
```

### JSON配置文件

创建 `xsd2code.json`:

```json
{
  "input": {
    "xsd": "schema.xsd"
  },
  "output": {
    "file": "types.go",
    "directory": "./generated",
    "split": false
  },
  "language": {
    "target": "go",
    "package": "models"
  },
  "features": {
    "validation": true,
    "json_tags": true,
    "xml_tags": true,
    "comments": true
  },
  "go": {
    "pointer_types": false,
    "omitempty": true,
    "tag_style": "camelCase"
  }
}
```

## 🌐 命名空间配置

### XML命名空间处理

```bash
# 保留原始命名空间
xsd2code -xsd=schema.xsd -preserve-namespace=true

# 自定义命名空间映射
xsd2code -xsd=schema.xsd -namespace-map="http://example.com=example"

# 忽略命名空间
xsd2code -xsd=schema.xsd -ignore-namespace=true
```

### 包/命名空间映射

```yaml
# 配置文件中的命名空间映射
namespace_mapping:
  "http://schemas.example.com/user": "user"
  "http://schemas.example.com/order": "order"
  "http://schemas.example.com/product": "product"
```

## 🎨 代码风格配置

### 命名约定

```bash
# Go风格命名
xsd2code -xsd=schema.xsd -naming-style=go

# Java风格命名
xsd2code -xsd=schema.xsd -naming-style=java

# 自定义命名规则
xsd2code -xsd=schema.xsd -naming-config=naming.yaml
```

### 自定义命名规则

创建 `naming.yaml`:

```yaml
naming_rules:
  # 类型命名
  type_name:
    pattern: "^[A-Z][a-zA-Z0-9]*Type$"
    transform: "PascalCase"
    suffix: "Type"
  
  # 字段命名
  field_name:
    pattern: "^[A-Z][a-zA-Z0-9]*$"
    transform: "PascalCase"
  
  # 常量命名
  constant_name:
    pattern: "^[A-Z][A-Z0-9_]*$"
    transform: "UPPER_SNAKE_CASE"
  
  # 包名
  package_name:
    pattern: "^[a-z][a-z0-9]*$"
    transform: "lowercase"
```

## 🔄 批处理配置

### 批量处理配置

```bash
# 处理多个XSD文件
xsd2code -batch -config=batch.yaml

# 并行处理
xsd2code -batch -parallel=4 -input-dir=schemas/ -output-dir=generated/
```

### 批处理配置文件

创建 `batch.yaml`:

```yaml
batch:
  input_directory: ./schemas
  output_directory: ./generated
  parallel_workers: 4
  
files:
  - xsd: user.xsd
    output: user_types.go
    package: user
  
  - xsd: order.xsd
    output: order_types.go
    package: order
    language: go
    features:
      validation: true
  
  - xsd: product.xsd
    output: ProductTypes.java
    package: com.example.product
    language: java
    features:
      annotations: true
      builder: true

defaults:
  language: go
  features:
    validation: true
    json_tags: true
    xml_tags: true
```

## 🎯 性能配置

### 内存和性能优化

```bash
# 设置内存限制
xsd2code -xsd=large-schema.xsd -memory-limit=2GB

# 启用缓存
xsd2code -xsd=schema.xsd -cache=true -cache-dir=./.cache

# 并行处理
xsd2code -xsd=schema.xsd -parallel=true -workers=4
```

### 大文件处理配置

```yaml
performance:
  memory_limit: "2GB"
  cache_enabled: true
  cache_directory: "./.cache"
  parallel_processing: true
  worker_count: 4
  chunk_size: 1000
  
large_file_handling:
  streaming_mode: true
  batch_size: 100
  memory_threshold: "500MB"
```

## 🔧 扩展配置

### 自定义生成器

```yaml
custom_generators:
  - name: "custom_go"
    file: "./generators/custom_go.yaml"
    templates:
      struct: "./templates/go_struct.tmpl"
      enum: "./templates/go_enum.tmpl"
  
  - name: "enterprise_java"
    file: "./generators/enterprise_java.yaml"
    features:
      spring_annotations: true
      jpa_annotations: true
```

### 插件配置

```yaml
plugins:
  - name: "validator_plugin"
    enabled: true
    config:
      strict_mode: true
      custom_rules: "./validation_rules.yaml"
  
  - name: "documentation_plugin"
    enabled: true
    config:
      generate_docs: true
      output_format: "markdown"
```

## 💡 配置最佳实践

### 1. 环境特定配置

```bash
# 开发环境
xsd2code -config=config.dev.yaml

# 生产环境
xsd2code -config=config.prod.yaml

# CI/CD环境
xsd2code -config=config.ci.yaml
```

### 2. 团队配置标准化

创建 `.xsd2code.yaml` 作为项目默认配置:

```yaml
# .xsd2code.yaml - 项目默认配置
defaults:
  language: go
  package: models
  features:
    validation: true
    json_tags: true
    xml_tags: true
    comments: true
  
  formatting:
    indent: "  "
    line_ending: "\n"
  
  naming:
    style: go
    type_suffix: "Type"
```

### 3. 配置验证

```bash
# 验证配置文件
xsd2code -validate-config=xsd2code.yaml

# 显示最终配置
xsd2code -show-config -config=xsd2code.yaml
```

## 📊 配置示例集合

### 微服务架构配置

```yaml
microservices:
  user_service:
    xsd: schemas/user.xsd
    language: go
    package: usermodels
    output: services/user/models/types.go
  
  order_service:
    xsd: schemas/order.xsd
    language: java
    package: com.company.order.models
    output: services/order/src/main/java/models/Types.java
  
  notification_service:
    xsd: schemas/notification.xsd
    language: csharp
    namespace: Company.Notification.Models
    output: services/notification/Models/Types.cs
```

### 前后端分离配置

```yaml
frontend:
  xsd: api-schema.xsd
  language: typescript
  output: frontend/src/types/api.ts
  features:
    interfaces: true
    optional_fields: true
    export_declarations: true

backend:
  xsd: api-schema.xsd
  language: go
  package: api
  output: backend/pkg/api/types.go
  features:
    validation: true
    json_tags: true
    builder_pattern: false
```

---

💡 **提示**: 合理使用配置选项可以大大提高开发效率，建议为不同的项目场景创建标准化的配置模板。

🔗 **相关页面**: 
- [[命令行参考|Command-Line-Reference]] - 完整参数说明
- [[最佳实践|Best-Practices]] - 配置最佳实践
- [[多语言支持|Multi-Language-Support]] - 语言特定配置
