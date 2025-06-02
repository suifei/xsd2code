# 故障排除 (Troubleshooting)

本页面提供XSD2Code常见问题的诊断和解决方案，帮助用户快速解决使用过程中遇到的问题。

## 目录

- [常见问题](#常见问题)
- [错误诊断](#错误诊断)
- [性能问题](#性能问题)
- [兼容性问题](#兼容性问题)
- [调试指南](#调试指南)
- [日志分析](#日志分析)
- [常用解决方案](#常用解决方案)

## 常见问题

### 1. XSD解析失败

#### 问题现象
```
Error: failed to parse XSD file: XML syntax error at line 10
```

#### 可能原因
- XSD文件格式不正确
- XML命名空间声明缺失
- 文件编码问题
- 文件路径错误

#### 解决方案

1. **验证XSD文件格式**
```bash
# 使用xmllint验证XSD文件
xmllint --schema http://www.w3.org/2001/XMLSchema.xsd your_schema.xsd

# 或者使用在线验证工具
```

2. **检查文件编码**
```bash
# 检查文件编码
file -bi your_schema.xsd

# 转换编码（如果需要）
iconv -f GBK -t UTF-8 your_schema.xsd > your_schema_utf8.xsd
```

3. **验证文件路径**
```bash
# 使用绝对路径
xsd2code -xsd=/absolute/path/to/schema.xsd

# 检查文件权限
ls -la your_schema.xsd
```

### 2. 代码生成失败

#### 问题现象
```
Error: failed to generate code: unsupported XSD feature: xs:union
```

#### 可能原因
- 使用了不支持的XSD特性
- 目标语言不支持某些类型映射
- 模板渲染错误

#### 解决方案

1. **检查支持的XSD特性**
```bash
# 查看详细错误信息
xsd2code -xsd=schema.xsd -debug

# 查看支持的特性列表
xsd2code -help-features
```

2. **使用兼容的XSD特性**
```xml
<!-- 不支持的union类型 -->
<xs:simpleType name="StringOrNumber">
  <xs:union memberTypes="xs:string xs:int"/>
</xs:simpleType>

<!-- 建议的替代方案：使用string类型 -->
<xs:simpleType name="StringOrNumber">
  <xs:restriction base="xs:string"/>
</xs:simpleType>
```

3. **自定义类型映射**
```bash
# 启用自定义类型映射
xsd2code -xsd=schema.xsd -custom-types

# 查看类型映射
xsd2code -xsd=schema.xsd -show-mappings
```

### 3. 生成的代码编译错误

#### 问题现象
```go
./generated.go:15:2: undefined: time
./generated.go:23:5: syntax error: unexpected }
```

#### 可能原因
- 缺少必要的import语句
- 类型映射不正确
- 模板语法错误

#### 解决方案

1. **检查导入语句**
```bash
# 启用JSON支持（自动添加encoding/json导入）
xsd2code -xsd=schema.xsd -json

# 检查生成的代码
head -20 generated.go
```

2. **验证生成的代码**
```bash
# 编译检查
go build generated.go

# 格式化代码
go fmt generated.go
```

### 4. 内存不足错误

#### 问题现象
```
Error: runtime: out of memory: cannot allocate
fatal error: out of memory
```

#### 可能原因
- XSD文件过大
- 递归类型定义
- 内存泄漏

#### 解决方案

1. **增加内存限制**
```bash
# 设置Go内存限制
export GOGC=50
export GOMEMLIMIT=4GiB

# 或者使用系统限制
ulimit -m 4194304  # 4GB
```

2. **分块处理大文件**
```bash
# 分解大的XSD文件
# 使用include/import而不是单一大文件

# 示例：拆分XSD文件
<!-- main.xsd -->
<xs:schema>
  <xs:include schemaLocation="types.xsd"/>
  <xs:include schemaLocation="elements.xsd"/>
</xs:schema>
```

### 5. 字符编码问题

#### 问题现象
```
Error: invalid UTF-8 encoding
生成的代码包含乱码字符
```

#### 解决方案

1. **确保文件UTF-8编码**
```bash
# 检查并转换编码
file -bi schema.xsd
iconv -f ISO-8859-1 -t UTF-8 schema.xsd > schema_utf8.xsd
```

2. **设置环境变量**
```bash
export LANG=en_US.UTF-8
export LC_ALL=en_US.UTF-8
```

## 错误诊断

### 诊断流程

1. **收集错误信息**
```bash
# 启用详细输出
xsd2code -xsd=schema.xsd -debug -verbose

# 保存错误日志
xsd2code -xsd=schema.xsd -debug 2>&1 | tee error.log
```

2. **检查输入文件**
```bash
# 验证XSD文件
xmllint --noout schema.xsd

# 检查文件大小
ls -lh schema.xsd

# 检查文件结构
head -50 schema.xsd
tail -50 schema.xsd
```

3. **测试简化版本**
```bash
# 创建最小测试用例
cat > minimal.xsd << EOF
<?xml version="1.0"?>
<xs:schema xmlns:xs="http://www.w3.org/2001/XMLSchema">
  <xs:element name="test" type="xs:string"/>
</xs:schema>
EOF

# 测试最小用例
xsd2code -xsd=minimal.xsd
```

### 错误代码对照表

| 错误代码 | 含义 | 常见原因 |
|---------|------|---------|
| E001 | XSD解析错误 | XML格式错误、命名空间问题 |
| E002 | 类型映射失败 | 不支持的XSD类型 |
| E003 | 代码生成错误 | 模板渲染失败 |
| E004 | 文件操作错误 | 权限不足、磁盘空间不足 |
| E005 | 内存不足 | 文件过大、内存泄漏 |
| E006 | 配置错误 | 参数无效、配置文件格式错误 |

### 详细错误分析

#### E001: XSD解析错误

```bash
# 详细诊断命令
xsd2code -xsd=schema.xsd -debug -parse-only

# 常见子错误
- XML syntax error: 检查XML格式
- Namespace not found: 检查命名空间声明
- Schema validation failed: 检查XSD规范符合性
```

#### E002: 类型映射失败

```bash
# 查看类型映射
xsd2code -xsd=schema.xsd -show-mappings -target=go

# 常见问题
- Unsupported type: 添加自定义类型映射
- Circular reference: 检查循环引用
- Complex type too deep: 简化类型结构
```

## 性能问题

### 性能诊断

1. **测量处理时间**
```bash
# 使用time命令
time xsd2code -xsd=large_schema.xsd

# 启用性能分析
xsd2code -xsd=schema.xsd -profile -profile-dir=./profiles

# 分析性能数据
go tool pprof profiles/cpu.prof
go tool pprof profiles/mem.prof
```

2. **内存使用监控**
```bash
# 监控内存使用
while true; do
  ps aux | grep xsd2code | grep -v grep
  sleep 1
done

# 或使用htop/top实时监控
htop -p $(pgrep xsd2code)
```

### 性能优化建议

1. **减少内存使用**
```bash
# 禁用不必要的功能
xsd2code -xsd=schema.xsd -no-validation -no-tests

# 使用流式处理
xsd2code -xsd=schema.xsd -streaming

# 限制并发数量
xsd2code -xsd=schema.xsd -max-goroutines=4
```

2. **提高处理速度**
```bash
# 启用缓存
xsd2code -xsd=schema.xsd -cache-dir=./cache

# 并行处理
xsd2code -xsd=schema.xsd -parallel

# 跳过格式化
xsd2code -xsd=schema.xsd -no-format
```

## 兼容性问题

### Go版本兼容性

| Go版本 | 支持状态 | 注意事项 |
|--------|---------|---------|
| 1.22.3+ | ✅ 完全支持 | 推荐版本 |
| 1.21.x | ⚠️ 部分支持 | 某些特性可能不可用 |
| 1.20.x | ❌ 不支持 | 请升级Go版本 |

```bash
# 检查Go版本
go version

# 升级Go（使用gvm）
gvm install go1.22.3
gvm use go1.22.3 --default
```

### 操作系统兼容性

#### Windows特定问题

1. **路径分隔符问题**
```bash
# 使用反斜杠或正斜杠都可以
xsd2code -xsd=C:\path\to\schema.xsd
xsd2code -xsd=C:/path/to/schema.xsd

# 推荐使用绝对路径
xsd2code -xsd="C:\Users\YourName\Documents\schema.xsd"
```

2. **字符编码问题**
```cmd
# 设置控制台编码为UTF-8
chcp 65001

# 或在PowerShell中
[Console]::OutputEncoding = [System.Text.Encoding]::UTF8
```

#### Linux/macOS特定问题

1. **权限问题**
```bash
# 确保可执行权限
chmod +x xsd2code

# 检查文件权限
ls -la xsd2code
```

2. **依赖库问题**
```bash
# 检查系统依赖
ldd xsd2code  # Linux
otool -L xsd2code  # macOS

# 安装缺失的库
sudo apt-get install libxml2-dev  # Ubuntu/Debian
brew install libxml2  # macOS
```

## 调试指南

### 启用调试模式

```bash
# 基本调试
xsd2code -xsd=schema.xsd -debug

# 详细调试
xsd2code -xsd=schema.xsd -debug -verbose

# 超详细调试（包含内部状态）
xsd2code -xsd=schema.xsd -debug -verbose -trace
```

### 调试输出解析

```
2024/06/02 10:15:23 [DEBUG] Starting XSD parsing
2024/06/02 10:15:23 [DEBUG] Reading file: schema.xsd
2024/06/02 10:15:23 [INFO]  Found 15 complex types
2024/06/02 10:15:23 [DEBUG] Processing complex type: Person
2024/06/02 10:15:23 [WARN]  Unsupported feature: xs:union in type Address
2024/06/02 10:15:23 [ERROR] Failed to process type: CustomerInfo
2024/06/02 10:15:23 [FATAL] Code generation failed
```

### 使用Go调试器

```bash
# 构建调试版本
go build -gcflags="all=-N -l" -o xsd2code-debug ./cmd/

# 使用delve调试器
dlv exec ./xsd2code-debug -- -xsd=schema.xsd

# 设置断点
(dlv) break main.main
(dlv) break github.com/suifei/xsd2code/pkg/xsdparser.(*XSDParser).ParseXSDFile
(dlv) continue
```

### 生成调试报告

```bash
# 生成完整的调试报告
xsd2code -xsd=schema.xsd -debug -report=debug_report.json

# 报告内容包括：
# - 系统信息
# - 输入文件信息
# - 解析过程详情
# - 错误堆栈
# - 性能指标
```

## 日志分析

### 日志级别

- **TRACE**: 详细的执行跟踪
- **DEBUG**: 调试信息
- **INFO**: 一般信息
- **WARN**: 警告信息
- **ERROR**: 错误信息
- **FATAL**: 致命错误

### 日志配置

```bash
# 设置日志级别
export XSD2CODE_LOG_LEVEL=DEBUG

# 设置日志输出文件
export XSD2CODE_LOG_FILE=xsd2code.log

# 设置日志格式
export XSD2CODE_LOG_FORMAT=json
```

### 日志分析工具

```bash
# 使用grep筛选错误
grep "ERROR\|FATAL" xsd2code.log

# 统计错误类型
grep "ERROR" xsd2code.log | cut -d' ' -f4- | sort | uniq -c

# 分析性能问题
grep "took" xsd2code.log | sort -k5 -nr | head -10
```

## 常用解决方案

### 解决方案模板

#### 1. 重置环境

```bash
#!/bin/bash
# reset_environment.sh

echo "Resetting XSD2Code environment..."

# 清理缓存
rm -rf ~/.xsd2code/cache/*

# 重置配置
rm -f ~/.xsd2code/config.yaml

# 清理临时文件
rm -f /tmp/xsd2code_*

# 重新安装（如果需要）
go install github.com/suifei/xsd2code/cmd/xsd2code@latest

echo "Environment reset complete."
```

#### 2. 诊断脚本

```bash
#!/bin/bash
# diagnose.sh

echo "XSD2Code Diagnostic Report"
echo "=========================="

echo "System Information:"
echo "OS: $(uname -a)"
echo "Go Version: $(go version)"
echo "XSD2Code Version: $(xsd2code -version)"

echo ""
echo "Environment Variables:"
env | grep -E "XSD2CODE|GO|PATH" | sort

echo ""
echo "File Information:"
if [ -f "$1" ]; then
    echo "File: $1"
    echo "Size: $(ls -lh "$1" | cut -d' ' -f5)"
    echo "Encoding: $(file -bi "$1")"
    echo "First 10 lines:"
    head -10 "$1"
else
    echo "No file specified or file not found"
fi

echo ""
echo "Test Run:"
if [ -f "$1" ]; then
    timeout 30s xsd2code -xsd="$1" -debug -dry-run 2>&1 | head -50
else
    echo "Cannot test without input file"
fi
```

#### 3. 快速修复脚本

```bash
#!/bin/bash
# quick_fix.sh

XSD_FILE="$1"
OUTPUT_DIR="${2:-./output}"

echo "Attempting quick fixes for $XSD_FILE..."

# 修复1: 检查并修复编码
if ! file -bi "$XSD_FILE" | grep -q "utf-8"; then
    echo "Converting encoding to UTF-8..."
    iconv -f ISO-8859-1 -t UTF-8 "$XSD_FILE" > "${XSD_FILE}.utf8"
    XSD_FILE="${XSD_FILE}.utf8"
fi

# 修复2: 验证XML格式
if ! xmllint --noout "$XSD_FILE" 2>/dev/null; then
    echo "XML format validation failed. Please check the file manually."
    exit 1
fi

# 修复3: 尝试基本生成
echo "Attempting basic generation..."
if xsd2code -xsd="$XSD_FILE" -output="$OUTPUT_DIR" -debug; then
    echo "Generation successful!"
else
    echo "Generation failed. Check the debug output above."
    exit 1
fi
```

### 常见问题FAQ

**Q: 为什么生成的Go代码无法编译？**
A: 通常是因为缺少必要的import语句。使用`-json`参数可以自动添加`encoding/json`导入。

**Q: 如何处理大型XSD文件？**
A: 使用`-streaming`模式和适当的内存限制，或者将大文件拆分为多个小文件。

**Q: 生成的代码中类型名称不符合Go命名规范怎么办？**
A: XSD2Code会自动转换命名，如果仍有问题，可以使用自定义类型映射。

**Q: 如何添加对新XSD特性的支持？**
A: 可以通过扩展LanguageMapper接口或提交功能请求到项目仓库。

**Q: 工具运行很慢怎么办？**
A: 启用缓存（`-cache-dir`），使用并行处理（`-parallel`），或禁用不必要的功能。

这个故障排除指南涵盖了XSD2Code使用过程中可能遇到的各种问题，提供了系统化的诊断和解决方案。
