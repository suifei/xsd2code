#!/bin/bash

# XSD2Code 多语言代码生成演示脚本
# 演示Go、Java、C#和Python代码生成功能

echo "======================================"
echo "XSD2Code 多语言代码生成演示"
echo "======================================"
echo

# 确保工具已构建
echo "构建xsd2code工具..."
go build -o xsd2code cmd/main.go
echo "✓ 构建完成"
echo

# 测试文件
XSD_FILE="examples/simple_example.xsd"
ADVANCED_XSD="examples/advanced_example.xsd"

echo "使用测试XSD文件: $XSD_FILE"
echo "高级测试XSD文件: $ADVANCED_XSD"
echo

# 1. Go代码生成
echo "1. 生成Go代码..."
./xsd2code -xsd=$XSD_FILE -lang=go -output=examples/demo_go.go -package=models -json
echo "✓ Go代码已生成: examples/demo_go.go"
echo

# 2. Java代码生成
echo "2. 生成Java代码..."
./xsd2code -xsd=$XSD_FILE -lang=java -output=examples/demo_java.java -package=com.example.models
echo "✓ Java代码已生成: examples/demo_java.java"
echo

# 3. C#代码生成
echo "3. 生成C#代码..."
./xsd2code -xsd=$XSD_FILE -lang=csharp -output=examples/demo_csharp.cs -package=Example.Models -json
echo "✓ C#代码已生成: examples/demo_csharp.cs"
echo

# 4. Python代码生成
echo "4. 生成Python代码..."
./xsd2code -xsd=$XSD_FILE -lang=python -output=examples/demo_python.py -package=models
echo "✓ Python代码已生成: examples/demo_python.py"
echo

# 5. 高级示例 - Go
echo "5. 生成高级示例(Go)..."
./xsd2code -xsd=$ADVANCED_XSD -lang=go -output=examples/advanced_demo_go.go -package=advanced -json
echo "✓ 高级Go代码已生成: examples/advanced_demo_go.go"
echo

# 6. 高级示例 - Java
echo "6. 生成高级示例(Java)..."
./xsd2code -xsd=$ADVANCED_XSD -lang=java -output=examples/advanced_demo_java.java -package=com.example.advanced
echo "✓ 高级Java代码已生成: examples/advanced_demo_java.java"
echo

# 7. 高级示例 - C#
echo "7. 生成高级示例(C#)..."
./xsd2code -xsd=$ADVANCED_XSD -lang=csharp -output=examples/advanced_demo_csharp.cs -package=Example.Advanced -json
echo "✓ 高级C#代码已生成: examples/advanced_demo_csharp.cs"
echo

# 8. 高级示例 - Python
echo "8. 生成高级示例(Python)..."
./xsd2code -xsd=$ADVANCED_XSD -lang=python -output=examples/advanced_demo_python.py -package=advanced_models
echo "✓ 高级Python代码已生成: examples/advanced_demo_python.py"
echo

echo "======================================"
echo "所有演示已完成！"
echo "======================================"
echo
echo "生成的文件:"
echo "- Go代码: examples/demo_go.go, examples/advanced_demo_go.go"
echo "- Java代码: examples/demo_java.java, examples/advanced_demo_java.java" 
echo "- C#代码: examples/demo_csharp.cs, examples/advanced_demo_csharp.cs"
echo "- Python代码: examples/demo_python.py, examples/advanced_demo_python.py"
echo
echo "特性演示:"
echo "✓ 多语言代码生成 (Go, Java, C#, Python)"
echo "✓ JSON兼容性支持"
echo "✓ XML注解和属性"
echo "✓ 枚举类型生成"
echo "✓ 复杂类型和嵌套结构"
echo "✓ 数组和集合类型"
echo "✓ 可选字段处理"
echo

# 显示类型映射信息
echo "显示类型映射信息:"
echo "Go类型映射:"
./xsd2code -xsd=$XSD_FILE -lang=go -show-mappings
echo
echo "Java类型映射:"
./xsd2code -xsd=$XSD_FILE -lang=java -show-mappings
echo
echo "C#类型映射:"
./xsd2code -xsd=$XSD_FILE -lang=csharp -show-mappings
echo
echo "Python类型映射:"
./xsd2code -xsd=$XSD_FILE -lang=python -show-mappings
