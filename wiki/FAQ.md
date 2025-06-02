# 常见问题 (FAQ)

本页面收集了用户在使用XSD2Code过程中最常遇到的问题和解决方案。

## 🚀 快速问题定位

### 按类别查找问题

| 类别 | 常见问题 |
|------|----------|
| 🏗️ **安装和构建** | [安装失败](#安装和构建问题)、[Go版本不兼容](#go版本问题)、[权限错误](#权限问题) |
| 📁 **文件处理** | [文件找不到](#文件路径问题)、[编码问题](#编码问题)、[导入失败](#导入问题) |
| 🔧 **代码生成** | [编译错误](#编译错误)、[类型映射](#类型映射问题)、[命名冲突](#命名问题) |
| 🌍 **多语言支持** | [Java输出](#java问题)、[C#问题](#csharp问题)、[TypeScript问题](#typescript问题) |
| ⚡ **性能和内存** | [内存不足](#内存问题)、[处理大文件](#大文件问题)、[速度慢](#性能问题) |

## 🏗️ 安装和构建问题

### Q: 安装失败，提示"Go版本不兼容"？
**A**: XSD2Code需要Go 1.21或更高版本。

```bash
# 检查Go版本
go version

# 如果版本过低，请更新Go
# 访问 https://golang.org/dl/ 下载最新版本
```

### Q: 构建时出现"module not found"错误？
**A**: 确保在正确的目录中运行构建命令。

```bash
# 正确的构建步骤
git clone https://github.com/suifei/xsd2code.git
cd xsd2code
go mod tidy
go build -o xsd2code cmd/main.go
```

### Q: Windows上无法执行xsd2code？
**A**: 检查以下几点：

1. **文件扩展名**: Windows需要`.exe`扩展名
```bash
go build -o xsd2code.exe cmd/main.go
```

2. **权限问题**: 确保文件有执行权限
3. **PATH设置**: 将xsd2code.exe路径添加到系统PATH

### Q: 提示"permission denied"错误？
**A**: 这是权限问题。

```bash
# Linux/macOS
chmod +x xsd2code

# 或使用sudo安装到系统目录
sudo cp xsd2code /usr/local/bin/
```

## 📁 文件处理问题

### Q: 提示"XSD file not found"？
**A**: 检查文件路径是否正确。

```bash
# 使用绝对路径
xsd2code -xsd=/full/path/to/schema.xsd

# 检查文件是否存在
ls -la schema.xsd

# 检查当前目录
pwd
```

### Q: XSD文件编码问题导致解析失败？
**A**: 确保XSD文件使用UTF-8编码。

```bash
# 检查文件编码
file -bi schema.xsd

# 转换为UTF-8编码
iconv -f ISO-8859-1 -t UTF-8 schema.xsd > schema-utf8.xsd
```

### Q: 导入的XSD文件找不到？
**A**: 检查schemaLocation路径。

```xml
<!-- 确保路径正确 -->
<xs:import namespace="http://example.com/types" 
           schemaLocation="./types/common.xsd"/>
```

工具会基于主XSD文件的目录解析相对路径。

### Q: 网络XSD文件无法访问？
**A**: 网络XSD支持有限，建议下载到本地。

```bash
# 下载XSD文件到本地
wget https://example.com/schema.xsd -O local-schema.xsd
xsd2code -xsd=local-schema.xsd
```

## 🔧 代码生成问题

### Q: 生成的Go代码编译失败？
**A**: 检查以下常见问题：

1. **缺少导入**: v3.1已自动处理
2. **包名冲突**: 使用不同的包名
```bash
xsd2code -xsd=schema.xsd -package=mymodels
```

3. **字段名冲突**: 检查XSD中的元素名
4. **语法错误**: 启用调试模式查看详情
```bash
xsd2code -xsd=schema.xsd -debug
```

### Q: 类型映射不正确？
**A**: 查看类型映射表并检查XSD定义。

```bash
# 显示类型映射
xsd2code -xsd=schema.xsd -show-mappings

# 检查特定类型的处理
xsd2code -xsd=schema.xsd -debug | grep "Type mapping"
```

### Q: 枚举类型生成不正确？
**A**: 确保XSD中枚举定义正确。

```xml
<!-- 正确的枚举定义 -->
<xs:simpleType name="StatusType">
    <xs:restriction base="xs:string">
        <xs:enumeration value="active"/>
        <xs:enumeration value="inactive"/>
    </xs:restriction>
</xs:simpleType>
```

### Q: 命名空间处理有问题？
**A**: 检查XSD的命名空间定义。

```xml
<!-- 确保命名空间定义完整 -->
<xs:schema targetNamespace="http://example.com/ns"
           xmlns:tns="http://example.com/ns"
           xmlns:xs="http://www.w3.org/2001/XMLSchema"
           elementFormDefault="qualified">
```

## 🌍 多语言支持问题

### Q: Java代码生成缺少JAXB注解？
**A**: 确保使用正确的Java模式。

```bash
# 正确的Java代码生成
xsd2code -xsd=schema.xsd -lang=java -package=com.example.models

# 检查生成的代码是否包含：
# @XmlRootElement
# @XmlElement
# @XmlAttribute
```

### Q: C#代码编译错误？
**A**: 检查.NET版本兼容性。

```bash
# 生成C#代码
xsd2code -xsd=schema.xsd -lang=csharp -package=Example.Models

# 确保项目引用了System.Xml.Serialization
```

### Q: TypeScript接口定义不完整？
**A**: TypeScript支持可能需要手动调整。

```bash
# 生成TypeScript接口
xsd2code -xsd=schema.xsd -lang=typescript

# 检查生成的接口定义
```

### Q: Python代码生成问题？
**A**: Python支持还在完善中。

```bash
# 当前Python支持状态
xsd2code -xsd=schema.xsd -lang=python
```

## ⚡ 性能和内存问题

### Q: 处理大型XSD文件时内存不足？
**A**: 优化内存使用的方法：

```bash
# 分段处理大型XSD
# 1. 拆分大型XSD为多个小文件
# 2. 逐个处理
# 3. 增加系统内存
```

### Q: 代码生成速度很慢？
**A**: 性能优化建议：

1. **禁用调试模式**（生产环境）
2. **使用SSD硬盘**
3. **关闭不必要的功能**
```bash
# 最小化功能的快速生成
xsd2code -xsd=schema.xsd -comments=false
```

### Q: 验证代码生成占用过多资源？
**A**: 选择性生成验证代码。

```bash
# 仅为需要的类型生成验证
# 目前版本会为所有类型生成验证，未来版本将支持选择性生成
```

## 🔍 XSD特性支持问题

### Q: 不支持某些XSD特性？
**A**: 查看支持的特性列表。

当前版本支持：
- ✅ 简单类型和复杂类型
- ✅ 元素和属性
- ✅ 约束（restriction）
- ✅ 命名空间
- ✅ 导入和包含
- ✅ 组和扩展

部分支持或计划支持：
- 🔄 Union类型（基本支持）
- 🔄 Any类型（计划支持）
- 🔄  替换组（计划支持）

### Q: Choice类型处理不理想？
**A**: Choice会生成所有可能的字段为可选。

```go
// XSD中的choice会生成：
type ContactType struct {
    Email *string `xml:"email,omitempty"`
    Phone *string `xml:"phone,omitempty"`
}
```

### Q: 递归类型定义问题？
**A**: 工具会检测并处理递归定义。

```xml
<!-- 递归类型示例 -->
<xs:complexType name="TreeNode">
    <xs:sequence>
        <xs:element name="value" type="xs:string"/>
        <xs:element name="children" type="tns:TreeNode" minOccurs="0" maxOccurs="unbounded"/>
    </xs:sequence>
</xs:complexType>
```

## 🔧 故障排除技巧

### 启用详细日志

```bash
# 启用详细调试信息
xsd2code -xsd=schema.xsd -debug -strict 2>&1 | tee debug.log
```

### 验证XSD文件

```bash
# 使用xmllint验证XSD语法
xmllint --schema schema.xsd --noout

# 或使用在线验证工具
```

### 检查生成的代码

```bash
# 检查Go代码语法
go fmt generated.go
go vet generated.go

# 检查Java代码
javac Generated.java

# 检查C#代码
csc /t:library Generated.cs
```

## 📞 获取更多帮助

### 自助诊断

1. **检查版本**: `xsd2code -version`
2. **查看帮助**: `xsd2code -help`
3. **启用调试**: `xsd2code -xsd=file.xsd -debug`
4. **检查映射**: `xsd2code -xsd=file.xsd -show-mappings`

### 报告问题

如果问题仍未解决，请：

1. **搜索现有Issue**: [GitHub Issues](https://github.com/suifei/xsd2code/issues)
2. **提供信息**:
   - XSD2Code版本
   - 操作系统和Go版本
   - 完整的命令行
   - 错误信息
   - 示例XSD文件（如果可能）

3. **创建Issue**: 包含重现步骤和期望结果

### 社区支持

- **GitHub Discussions**: 通用讨论和经验分享
- **Wiki页面**: 查看其他文档页面
- **示例代码**: 查看`examples/`目录

## 🔄 版本升级问题

### Q: 升级到v3.1后代码不兼容？
**A**: v3.1主要是增强功能，应该向后兼容。

如果遇到问题：
1. 重新生成代码
2. 检查新的导入要求
3. 查看变更日志

### Q: 旧版本的配置不工作？
**A**: 检查参数变化。

```bash
# v3.1新功能，旧版本不支持
xsd2code -xsd=schema.xsd -validation

# 确保使用正确的参数格式
```

---

❓ **找不到您的问题？** 

1. 查看 [[故障排除|Troubleshooting]] 页面
2. 搜索 [GitHub Issues](https://github.com/suifei/xsd2code/issues)
3. 创建新的Issue描述您的问题
