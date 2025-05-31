// xsd2go是一个用于将XSD模式文件转换为Go结构体的命令行工具
// 支持完整的 XML Schema 规范，生成兼容 XML 和 JSON 的 Go 数据结构
//
// 用法:
//
//	xsd2go -xsd=<XSD文件路径> [-output=<输出文件路径>] [-package=<包名>] [-json] [-debug]
//
// 示例:
//
//	xsd2go -xsd=schema.xsd
//	xsd2go -xsd=schema.xsd -output=types.go -package=mymodels -json
//	xsd2go -xsd=schema.xsd -debug -json
package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/suifei/xsd2code/pkg/xsdparser"
)

// XSDConverterConfig 配置XSD转换器的选项
type XSDConverterConfig struct {
	XSDPath         string
	OutputPath      string
	PackageName     string
	EnableJSON      bool
	DebugMode       bool
	StrictMode      bool
	IncludeComments bool
}

// parseFlags 解析命令行参数
func parseFlags() *XSDConverterConfig {
	config := &XSDConverterConfig{}

	flag.StringVar(&config.XSDPath, "xsd", "", "XSD文件的路径 (必需)")
	flag.StringVar(&config.OutputPath, "output", "", "输出Go代码的文件路径 (可选)")
	flag.StringVar(&config.PackageName, "package", "models", "生成的Go代码包名 (默认: models)")
	flag.BoolVar(&config.EnableJSON, "json", false, "生成JSON兼容的标签")
	flag.BoolVar(&config.DebugMode, "debug", false, "启用调试模式")
	flag.BoolVar(&config.StrictMode, "strict", false, "启用严格模式")
	flag.BoolVar(&config.IncludeComments, "comments", true, "在生成的代码中包含注释")
	help := flag.Bool("help", false, "显示帮助信息")
	version := flag.Bool("version", false, "显示版本信息")

	flag.Parse()

	if *version {
		fmt.Println("XSD到Go转换工具 v3.0 (统一解析器)")
		fmt.Println("支持完整 XML Schema 规范，兼容 XML/JSON")
		os.Exit(0)
	}
	if *help {
		printHelp()
		os.Exit(0)
	}

	return config
}

// validateConfig 验证配置
func validateConfig(config *XSDConverterConfig) error {
	if config.XSDPath == "" {
		return fmt.Errorf("必须提供XSD文件路径")
	}

	// 验证XSD文件是否存在
	if _, err := os.Stat(config.XSDPath); os.IsNotExist(err) {
		return fmt.Errorf("XSD文件不存在: %s", config.XSDPath)
	}

	// 如果未提供输出路径，生成默认值
	if config.OutputPath == "" {
		ext := filepath.Ext(config.XSDPath)
		baseName := strings.TrimSuffix(filepath.Base(config.XSDPath), ext)
		config.OutputPath = baseName + ".go"
	}

	// 验证包名
	if !isValidPackageName(config.PackageName) {
		return fmt.Errorf("无效的包名: %s", config.PackageName)
	}

	return nil
}

// isValidPackageName 检查是否为有效的Go包名
func isValidPackageName(name string) bool {
	if name == "" {
		return false
	}
	// 简单验证：只能包含字母、数字和下划线，且以字母开头
	for i, r := range name {
		if i == 0 && !((r >= 'a' && r <= 'z') || (r >= 'A' && r <= 'Z') || r == '_') {
			return false
		}
		if !((r >= 'a' && r <= 'z') || (r >= 'A' && r <= 'Z') || (r >= '0' && r <= '9') || r == '_') {
			return false
		}
	}
	return true
}

// runConverter 执行XSD转换
func runConverter(config *XSDConverterConfig) error {
	fmt.Printf("XSD到Go转换工具 - v3.0 (统一解析器)\n")
	fmt.Printf("===============================================\n")
	fmt.Printf("输入文件: %s\n", config.XSDPath)
	fmt.Printf("输出文件: %s\n", config.OutputPath)
	fmt.Printf("包名: %s\n", config.PackageName)
	fmt.Printf("JSON兼容: %t\n", config.EnableJSON)
	fmt.Printf("调试模式: %t\n", config.DebugMode)
	fmt.Printf("严格模式: %t\n", config.StrictMode)
	fmt.Printf("---------------------------------------------\n")
	// 使用新的统一解析器（已合并标准和高级功能）
	parser := xsdparser.NewUnifiedXSDParser(config.XSDPath, config.OutputPath, config.PackageName)

	// 设置解析器选项
	parser.SetJSONCompatible(config.EnableJSON)
	parser.SetDebugMode(config.DebugMode)
	parser.SetStrictMode(config.StrictMode)
	parser.SetIncludeComments(config.IncludeComments)

	// 解析XSD
	fmt.Println("开始解析XSD文件...")
	if err := parser.Parse(); err != nil {
		return fmt.Errorf("解析XSD失败: %v", err)
	}

	fmt.Println("解析完成！")

	// 生成Go代码
	fmt.Println("生成Go代码...")
	if err := parser.GenerateGoCode(); err != nil {
		return fmt.Errorf("生成Go代码失败: %v", err)
	}

	fmt.Printf("成功！Go结构已生成在: %s\n", config.OutputPath)
	return nil
}

// printHelp 打印帮助信息
func printHelp() {
	fmt.Println("XSD到Go转换工具 v3.0 (统一解析器)")
	fmt.Println("====================================")
	fmt.Println("将任意 XML Schema (XSD) 文件转换为 Go 结构体定义")
	fmt.Println("支持完整的 XML Schema 规范，自动合并标准和高级特性")
	fmt.Println("")
	fmt.Println("用法:")
	fmt.Println("  xsd2go -xsd=<XSD文件路径> [选项...]")
	fmt.Println("")
	fmt.Println("必需参数:")
	fmt.Println("  -xsd        要转换的XSD文件路径")
	fmt.Println("")
	fmt.Println("可选参数:")
	fmt.Println("  -output     生成的Go代码输出文件路径 (默认: 与XSD同名的.go文件)")
	fmt.Println("  -package    生成的Go代码包名 (默认: models)")
	fmt.Println("  -json       生成JSON兼容的标签")
	fmt.Println("  -debug      启用调试模式 (输出详细解析信息)")
	fmt.Println("  -strict     启用严格模式 (更严格的类型检查)")
	fmt.Println("  -comments   在生成的代码中包含注释 (默认启用)")
	fmt.Println("  -help       显示此帮助信息")
	fmt.Println("  -version    显示版本信息")
	fmt.Println("")
	fmt.Println("示例:")
	fmt.Println("  # 基本转换")
	fmt.Println("  xsd2go -xsd=schema.xsd")
	fmt.Println("")
	fmt.Println("  # 生成JSON兼容的结构体")
	fmt.Println("  xsd2go -xsd=schema.xsd -json -output=types.go -package=api")
	fmt.Println("")
	fmt.Println("  # 启用调试和严格模式")
	fmt.Println("  xsd2go -xsd=complex_schema.xsd -json -debug -strict")
	fmt.Println("")
	fmt.Println("支持的XSD特性 (自动检测和处理):")
	fmt.Println("  • 复杂类型和简单类型")
	fmt.Println("  • 元素和属性")
	fmt.Println("  • 组定义和组引用")
	fmt.Println("  • 类型继承 (extension/restriction)")
	fmt.Println("  • 枚举类型")
	fmt.Println("  • 出现次数约束 (minOccurs/maxOccurs)")
	fmt.Println("  • 命名空间支持")
	fmt.Println("  • 导入和包含")
	fmt.Println("  • 选择元素 (xs:choice)")
	fmt.Println("  • 内联复杂类型")
	fmt.Println("")
}

func main() {
	// 解析命令行参数
	config := parseFlags()

	// 如果没有提供参数，显示帮助
	if len(os.Args) == 1 {
		fmt.Println("错误: 请提供必要的参数")
		fmt.Println("使用 xsd2go -help 获取帮助")
		os.Exit(1)
	}

	// 验证配置
	if err := validateConfig(config); err != nil {
		fmt.Printf("配置错误: %v\n", err)
		fmt.Println("使用 xsd2go -help 获取帮助")
		os.Exit(1)
	}

	// 执行转换
	if err := runConverter(config); err != nil {
		fmt.Printf("转换失败: %v\n", err)
		os.Exit(1)
	}
}
