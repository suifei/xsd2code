# 命令行参考

本页面提供XSD2Code所有命令行参数的完整参考，包括详细说明、使用示例和注意事项。

## 📋 命令格式

```bash
xsd2code [全局选项] -xsd=<文件路径> [输出选项] [生成选项] [调试选项]
```

## 🎯 必需参数

### -xsd
- **类型**: `string`
- **必需**: ✅ 是
- **描述**: 指定要转换的XSD文件路径
- **支持**: 绝对路径、相对路径、URL（部分支持）

**示例**:
```bash
# 相对路径
xsd2code -xsd=schema.xsd

# 绝对路径  
xsd2code -xsd=/path/to/schema.xsd

# Windows路径
xsd2code -xsd=C:\schemas\schema.xsd

# URL（实验性）
xsd2code -xsd=https://example.com/schema.xsd
```

## 🌍 语言和输出选项

### -lang
- **类型**: `string`
- **默认值**: `go`
- **可选值**: `go`, `java`, `csharp`, `python`
- **描述**: 指定目标编程语言

**示例**:
```bash
# Go代码（默认）
xsd2code -xsd=schema.xsd -lang=go

# Java代码
xsd2code -xsd=schema.xsd -lang=java

# C#代码
xsd2code -xsd=schema.xsd -lang=csharp

# Python代码
xsd2code -xsd=schema.xsd -lang=python
```

### -output
- **类型**: `string`
- **默认值**: 自动生成（基于XSD文件名和语言）
- **描述**: 指定输出文件路径

**自动命名规则**:
- Go: `{xsd_name}.go`
- Java: `{XsdName}.java`
- C#: `{XsdName}.cs`
- Python: `{xsd_name}.py`

**示例**:
```bash
# 使用默认输出文件名
xsd2code -xsd=user.xsd
# 生成: user.go

# 指定输出文件
xsd2code -xsd=user.xsd -output=models.go

# 输出到不同目录
xsd2code -xsd=user.xsd -output=./generated/user_types.go

# 不同语言的输出
xsd2code -xsd=user.xsd -lang=java -output=UserTypes.java
```

### -package
- **类型**: `string`
- **默认值**: `models`
- **描述**: 指定包名、命名空间或模块名

**不同语言的包名格式**:
- **Go**: 包名（如 `models`, `api`）
- **Java**: 完整包名（如 `com.example.models`）
- **C#**: 命名空间（如 `Example.Models`）
- **Python**: 模块名

**示例**:
```bash
# Go包名
xsd2code -xsd=schema.xsd -package=api

# Java包名
xsd2code -xsd=schema.xsd -lang=java -package=com.company.models

# C#命名空间
xsd2code -xsd=schema.xsd -lang=csharp -package=Company.Models

# Python模块名
xsd2code -xsd=schema.xsd -lang=python -package=models
```

## 🔧 代码生成选项

### -json
- **类型**: `boolean`
- **默认值**: `false`
- **描述**: 生成JSON序列化标签（适用于支持的语言）

**影响的语言**:
- **Go**: 添加 `json:"fieldname"` 标签
- **C#**: 添加 `[JsonPropertyName("fieldname")]` 注解
- **Java**: 添加Jackson注解
- **Python**: 添加装饰器

**示例**:
```bash
# 生成带JSON标签的Go代码
xsd2code -xsd=schema.xsd -json

# 生成的Go代码示例：
# type User struct {
#     Name string `xml:"name" json:"name"`
#     Age  int    `xml:"age" json:"age"`
# }
```

### -comments
- **类型**: `boolean`
- **默认值**: `true`
- **描述**: 是否在生成的代码中包含注释

**包含的注释类型**:
- XSD文档注释 (`xs:documentation`)
- 字段说明注释
- 约束信息注释
- 生成信息注释

**示例**:
```bash
# 包含注释（默认）
xsd2code -xsd=schema.xsd -comments=true

# 不包含注释
xsd2code -xsd=schema.xsd -comments=false
```

## 🔍 验证和测试选项

### -validation
- **类型**: `boolean`
- **默认值**: `false`
- **描述**: 生成数据验证代码

**生成的验证功能**:
- 约束检查函数
- Pattern匹配验证
- 长度检查
- 范围检查
- 枚举值验证

**示例**:
```bash
# 生成验证代码
xsd2code -xsd=schema.xsd -validation

# 生成的验证函数示例：
# func (u User) Validate() bool {
#     if len(u.Name) < 1 || len(u.Name) > 50 {
#         return false
#     }
#     return true
# }
```

### -validation-output
- **类型**: `string`
- **默认值**: 自动生成
- **描述**: 指定验证代码的输出文件路径
- **依赖**: 需要与 `-validation` 一起使用

**示例**:
```bash
# 指定验证代码输出文件
xsd2code -xsd=schema.xsd -validation -validation-output=validation.go

# 分离主代码和验证代码
xsd2code -xsd=schema.xsd -output=types.go -validation -validation-output=types_validation.go
```

### -tests
- **类型**: `boolean`
- **默认值**: `false`
- **描述**: 生成单元测试代码

**生成的测试内容**:
- 结构体创建测试
- XML序列化/反序列化测试
- 验证功能测试
- 边界条件测试

**示例**:
```bash
# 生成测试代码
xsd2code -xsd=schema.xsd -tests

# 同时生成主代码、验证和测试
xsd2code -xsd=schema.xsd -validation -tests
```

### -test-output
- **类型**: `string`
- **默认值**: 自动生成
- **描述**: 指定测试代码的输出文件路径
- **依赖**: 需要与 `-tests` 一起使用

**示例**:
```bash
# 指定测试文件输出路径
xsd2code -xsd=schema.xsd -tests -test-output=schema_test.go
```

### -benchmarks
- **类型**: `boolean`
- **默认值**: `false`
- **描述**: 生成基准测试代码（Go语言）

**示例**:
```bash
# 生成包含基准测试的代码
xsd2code -xsd=schema.xsd -tests -benchmarks
```

## 📊 信息和工具选项

### -show-mappings
- **类型**: `boolean`
- **默认值**: `false`
- **描述**: 显示XSD类型到目标语言的类型映射

**输出信息**:
- XSD内置类型映射
- 自定义类型映射
- 约束处理方式
- 可选性处理

**示例**:
```bash
# 显示Go语言类型映射
xsd2code -xsd=schema.xsd -show-mappings

# 显示Java语言类型映射
xsd2code -xsd=schema.xsd -lang=java -show-mappings
```

### -sample
- **类型**: `boolean`
- **默认值**: `false`
- **描述**: 生成基于XSD的示例XML

**生成特点**:
- 包含所有必需元素
- 使用默认值和示例值
- 符合XSD约束
- 包含命名空间声明

**示例**:
```bash
# 生成示例XML到控制台
xsd2code -xsd=schema.xsd -sample

# 保存示例XML到文件
xsd2code -xsd=schema.xsd -sample > sample.xml
```

### -validate
- **类型**: `string`
- **默认值**: 无
- **描述**: 验证指定XML文件是否符合XSD

**示例**:
```bash
# 验证XML文件
xsd2code -xsd=schema.xsd -validate=data.xml

# 验证多个文件
xsd2code -xsd=schema.xsd -validate=file1.xml -validate=file2.xml
```

## 🐛 调试和诊断选项

### -debug
- **类型**: `boolean`
- **默认值**: `false`
- **描述**: 启用详细的调试输出

**调试信息包括**:
- XSD解析过程
- 类型映射决策
- 代码生成步骤
- 错误详细信息

**示例**:
```bash
# 启用调试模式
xsd2code -xsd=schema.xsd -debug

# 调试模式下的输出示例：
# [DEBUG] Parsing XSD file: schema.xsd
# [DEBUG] Found complex type: UserType
# [DEBUG] Processing element: name (string, required)
# [DEBUG] Generating Go struct for UserType
```

### -strict
- **类型**: `boolean`
- **默认值**: `false`
- **描述**: 启用严格的XSD验证模式

**严格模式特点**:
- 更严格的XSD语法检查
- 不允许未定义的类型引用
- 强制命名空间一致性
- 严格的约束验证

**示例**:
```bash
# 启用严格模式
xsd2code -xsd=schema.xsd -strict

# 与调试模式结合
xsd2code -xsd=schema.xsd -strict -debug
```

## 🏭 特殊用途选项

### -plc
- **类型**: `boolean`
- **默认值**: `false`
- **描述**: 启用PLC（可编程逻辑控制器）特定的类型映射

**PLC模式特点**:
- 优化的数据类型映射
- 工业自动化标准兼容
- 特殊的数值类型处理

**示例**:
```bash
# PLC模式代码生成
xsd2code -xsd=plc-schema.xsd -plc
```

## ℹ️ 帮助和版本选项

### -help
- **类型**: `boolean`
- **描述**: 显示帮助信息和所有可用选项

**示例**:
```bash
xsd2code -help
```

### -version
- **类型**: `boolean`
- **描述**: 显示版本信息

**示例**:
```bash
xsd2code -version
# 输出: XSD2Code v3.1.0
```

## 🔄 组合使用示例

### 完整的生产环境配置

```bash
# 生成Go代码，包含JSON标签、验证和测试
xsd2code \
  -xsd=api-schema.xsd \
  -lang=go \
  -output=./generated/api/types.go \
  -package=api \
  -json \
  -validation \
  -validation-output=./generated/api/validation.go \
  -tests \
  -test-output=./generated/api/types_test.go \
  -benchmarks
```

### 多语言代码生成

```bash
# 生成多语言代码
xsd2code -xsd=schema.xsd -lang=go -output=go/types.go -package=models
xsd2code -xsd=schema.xsd -lang=java -output=java/Types.java -package=com.example.models
xsd2code -xsd=schema.xsd -lang=csharp -output=csharp/Types.cs -package=Example.Models
xsd2code -xsd=schema.xsd -lang=typescript -output=typescript/types.ts
```

### 调试复杂问题

```bash
# 启用所有调试选项
xsd2code \
  -xsd=complex-schema.xsd \
  -debug \
  -strict \
  -show-mappings \
  -output=debug-output.go
```

## ⚠️ 注意事项

### 参数优先级

1. 命令行参数优先于环境变量
2. 显式指定的值优先于默认值
3. 后面的参数覆盖前面的同名参数

### 文件路径处理

- 相对路径基于当前工作目录
- 输出路径的父目录必须存在
- Windows系统支持UNC路径

### 内存和性能

- 大型XSD文件可能需要更多内存
- 使用 `-debug` 会增加内存使用
- 复杂验证代码会影响编译时间

## 🔗 相关页面

- [[基本用法|Basic-Usage]] - 常用命令组合
- [[高级示例|Advanced-Examples]] - 复杂场景配置
- [[故障排除|Troubleshooting]] - 问题解决方案

---

💡 **提示**: 建议将常用的参数组合写入Makefile或脚本文件，便于重复使用和团队协作。
