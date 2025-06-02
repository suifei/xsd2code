# 基本用法

本页面详细介绍XSD2Code的基本使用方法、命令行参数和常用场景。

## 📋 基本语法

```bash
xsd2code [选项] -xsd=<XSD文件路径>
```

## 🎯 核心参数

### 必需参数

| 参数 | 描述 | 示例 |
|------|------|------|
| `-xsd` | XSD文件路径（必需） | `-xsd=schema.xsd` |

### 常用参数

| 参数 | 默认值 | 描述 | 示例 |
|------|--------|------|------|
| `-lang` | `go` | 目标语言 | `-lang=java` |
| `-output` | 自动生成 | 输出文件路径 | `-output=types.go` |
| `-package` | `models` | 包名/命名空间 | `-package=com.example` |
| `-json` | `false` | 生成JSON标签 | `-json` |
| `-comments` | `true` | 包含注释 | `-comments=false` |

## 🌍 多语言代码生成

### Go代码生成

```bash
# 基本Go代码生成
xsd2code -xsd=schema.xsd

# 指定输出文件和包名
xsd2code -xsd=schema.xsd -output=models.go -package=models

# 包含JSON标签
xsd2code -xsd=schema.xsd -json

# 生成到指定目录
xsd2code -xsd=schema.xsd -output=./generated/types.go
```

**生成的Go代码特点**：
- 结构体定义
- XML和JSON标签（可选）
- 指针类型用于可选字段
- 枚举类型安全处理

### Java代码生成

```bash
# 基本Java代码生成
xsd2code -xsd=schema.xsd -lang=java

# 指定包名和输出文件
xsd2code -xsd=schema.xsd -lang=java -package=com.example.models -output=Types.java

# 包含所有JAXB注解
xsd2code -xsd=schema.xsd -lang=java -package=com.example -output=MyTypes.java
```

**生成的Java代码特点**：
- POJO类定义
- JAXB注解
- Getter/Setter方法
- 枚举类型

### C#代码生成

```bash
# 基本C#代码生成
xsd2code -xsd=schema.xsd -lang=csharp

# 指定命名空间和输出文件
xsd2code -xsd=schema.xsd -lang=csharp -package=Example.Models -output=Types.cs

# 包含JSON支持
xsd2code -xsd=schema.xsd -lang=csharp -json -output=Types.cs
```

**生成的C#代码特点**：
- 类和属性定义
- XML序列化注解
- 可空类型处理
- 枚举定义

### TypeScript代码生成

```bash
# 基本TypeScript代码生成
xsd2code -xsd=schema.xsd -lang=typescript

# 指定输出文件
xsd2code -xsd=schema.xsd -lang=typescript -output=types.ts
```

## 🔧 验证和测试

### 生成验证代码

```bash
# 生成验证函数
xsd2code -xsd=schema.xsd -validation

# 指定验证代码输出文件
xsd2code -xsd=schema.xsd -validation -validation-output=validation.go

# 同时生成类型和验证代码
xsd2code -xsd=schema.xsd -output=types.go -validation -validation-output=validation.go
```

### 生成测试代码

```bash
# 生成单元测试
xsd2code -xsd=schema.xsd -tests

# 指定测试文件输出路径
xsd2code -xsd=schema.xsd -tests -test-output=types_test.go

# 生成基准测试
xsd2code -xsd=schema.xsd -tests -benchmarks
```

## 📊 信息和调试

### 显示类型映射

```bash
# 显示XSD到目标语言的类型映射
xsd2code -xsd=schema.xsd -show-mappings

# 不同语言的类型映射
xsd2code -xsd=schema.xsd -lang=java -show-mappings
xsd2code -xsd=schema.xsd -lang=csharp -show-mappings
```

### 调试模式

```bash
# 启用调试输出
xsd2code -xsd=schema.xsd -debug

# 启用严格模式（更严格的XSD验证）
xsd2code -xsd=schema.xsd -strict

# 组合使用
xsd2code -xsd=schema.xsd -debug -strict
```

### 生成示例XML

```bash
# 生成基于XSD的示例XML
xsd2code -xsd=schema.xsd -sample

# 保存示例XML到文件
xsd2code -xsd=schema.xsd -sample > sample.xml
```

## 🎯 实际使用场景

### 场景1：Web API开发

```bash
# 为API开发生成Go结构体
xsd2code -xsd=api-schema.xsd -output=api_types.go -package=api -json

# 为前端生成TypeScript接口
xsd2code -xsd=api-schema.xsd -lang=typescript -output=api-types.ts
```

### 场景2：配置文件处理

```bash
# 生成配置文件结构体
xsd2code -xsd=config-schema.xsd -output=config.go -package=config

# 包含验证逻辑
xsd2code -xsd=config-schema.xsd -output=config.go -validation -validation-output=config_validation.go
```

### 场景3：数据交换格式

```bash
# 为多语言项目生成一致的数据结构
xsd2code -xsd=data-exchange.xsd -lang=go -output=go/types.go
xsd2code -xsd=data-exchange.xsd -lang=java -output=java/Types.java -package=com.example.data
xsd2code -xsd=data-exchange.xsd -lang=csharp -output=csharp/Types.cs -package=Example.Data
```

### 场景4：集成第三方服务

```bash
# 处理SOAP服务的XSD
xsd2code -xsd=soap-service.xsd -output=soap_types.go -package=soap -validation

# 处理带命名空间的复杂XSD
xsd2code -xsd=complex-service.xsd -output=service_types.go -debug
```

## 🔍 高级选项

### 特殊处理选项

```bash
# PLC类型映射（用于工业自动化）
xsd2code -xsd=plc-schema.xsd -plc

# 禁用注释生成
xsd2code -xsd=schema.xsd -comments=false

# 启用严格模式
xsd2code -xsd=schema.xsd -strict
```

### 输出控制

```bash
# 静默模式（仅输出错误）
xsd2code -xsd=schema.xsd -output=types.go 2>/dev/null

# 详细输出
xsd2code -xsd=schema.xsd -debug -output=types.go
```

## 📝 文件组织建议

### 项目结构建议

```
project/
├── schemas/           # XSD文件
│   ├── api.xsd
│   ├── config.xsd
│   └── data.xsd
├── generated/         # 生成的代码
│   ├── api/
│   │   ├── types.go
│   │   └── validation.go
│   ├── config/
│   │   └── types.go
│   └── data/
│       └── types.go
└── Makefile          # 自动化脚本
```

### Makefile示例

```makefile
# 生成所有代码
generate:
	xsd2code -xsd=schemas/api.xsd -output=generated/api/types.go -package=api -validation -validation-output=generated/api/validation.go
	xsd2code -xsd=schemas/config.xsd -output=generated/config/types.go -package=config
	xsd2code -xsd=schemas/data.xsd -output=generated/data/types.go -package=data -json

# 清理生成的文件
clean:
	rm -rf generated/

# 验证生成的代码
verify:
	go build ./generated/...

.PHONY: generate clean verify
```

## ⚠️ 注意事项

### 文件路径

- 使用绝对路径或相对路径
- Windows用户注意路径分隔符
- 确保XSD文件可读

### 输出覆盖

- 工具会覆盖现有输出文件
- 建议使用版本控制
- 可以通过重定向保存备份

### 依赖处理

- 工具会自动处理XSD导入
- 确保所有依赖的XSD文件可访问
- 相对路径基于主XSD文件位置

## 🚨 错误处理

### 常见错误及解决方案

```bash
# XSD文件不存在
# Error: failed to read XSD file: no such file or directory
# 解决：检查文件路径是否正确

# 权限错误
# Error: permission denied
# 解决：检查文件读写权限

# XSD格式错误
# Error: failed to parse XSD
# 解决：使用 -debug 查看详细错误信息
```

## 📞 获取帮助

### 内置帮助

```bash
# 显示所有可用选项
xsd2code -help

# 显示版本信息
xsd2code -version
```

### 更多资源

- [[命令行参考|Command-Line-Reference]] - 完整参数说明
- [[高级示例|Advanced-Examples]] - 复杂场景示例
- [[故障排除|Troubleshooting]] - 问题解决方案

---

💡 **提示**: 建议从简单的XSD开始，逐步尝试更复杂的功能。生成的代码可以直接在您的项目中使用！
