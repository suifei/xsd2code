# XSD2Code 示例

这个目录包含了使用 XSD2Code 工具的各种示例。

## 简单示例

### 基本转换
```bash
# 转换简单的XSD文件
./xsd2code.exe -xsd=examples/simple_example.xsd -output=examples/simple_types.go -package=example -json

# 生成带验证和测试的完整代码
./xsd2code.exe -xsd=examples/simple_example.xsd -output=examples/simple_complete.go -package=example -json -validation -tests -benchmarks
```

### 高级功能
```bash
# 显示类型映射
./xsd2code.exe -xsd=examples/simple_example.xsd -show-mappings -lang=go

# 验证XML文件
./xsd2code.exe -xsd=examples/simple_example.xsd -validate=examples/sample.xml

# 生成示例XML
./xsd2code.exe -xsd=examples/simple_example.xsd -sample
```

## 复杂示例

项目中的 `test/TC6_XML_V10_B.xsd` 是一个复杂的 PLCopen XML Schema 文件，演示了工具对完整 XSD 规范的支持。

### 处理复杂XSD
```bash
# 处理 PLCopen XSD
./xsd2code.exe -xsd=test/TC6_XML_V10_B.xsd -output=plcopen.go -package=plcopen -json -debug
```

## 生成的代码示例

转换后的 Go 代码包含：

1. **结构体定义** - 对应 XSD 复杂类型
2. **枚举常量** - 对应 XSD 枚举
3. **XML 标签** - 支持正确的 XML 序列化
4. **JSON 标签** - 可选的 JSON 兼容性
5. **验证函数** - 可选的数据验证
6. **测试代码** - 可选的完整测试套件

## 使用生成的代码

```go
package main

import (
    "encoding/xml"
    "fmt"
    "github.com/suifei/xsd2code/examples"
)

func main() {
    // 创建实例
    person := &example.PersonType{
        Name:   "张三",
        Age:    30,
        Email:  "zhangsan@example.com",
        Status: example.STATUS_TYPE_ACTIVE,
        Id:     "person_001",
    }
    
    // 验证数据
    if err := person.Validate(); err != nil {
        fmt.Printf("验证失败: %v\n", err)
        return
    }
    
    // 序列化为XML
    xmlData, err := xml.MarshalIndent(person, "", "  ")
    if err != nil {
        fmt.Printf("XML序列化失败: %v\n", err)
        return
    }
    
    fmt.Printf("生成的XML:\n%s\n", xmlData)
}
```
