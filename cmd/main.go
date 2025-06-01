// xsd2go是一个用于将XSD模式文件转换为Go结构体的命令行工具
// 支持完整的 XML Schema 规范，生成兼容 XML 和 JSON 的 Go 数据结构
// 现在支持生成验证代码、测试代码和基准测试
//
// 用法:
//
//	xsd2go -xsd=<XSD文件路径> [-output=<输出文件路径>] [-package=<包名>] [-json] [-debug]
//
// 示例:
//
//	xsd2go -xsd=schema.xsd
//	xsd2go -xsd=schema.xsd -output=types.go -package=mymodels -json
//	xsd2go -xsd=schema.xsd -debug -json -validation -tests -benchmarks
package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/suifei/xsd2code/pkg/generator"
	"github.com/suifei/xsd2code/pkg/types"
	"github.com/suifei/xsd2code/pkg/validator"
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
	// 新增：代码生成选项
	GenerateValidation   bool
	GenerateTests        bool
	GenerateBenchmarks   bool
	TestOutputPath       string
	ValidationOutputPath string
	// 新增：多语言支持和类型映射
	TargetLanguage    string
	EnableCustomTypes bool // 启用PLC等自定义类型映射
	ShowTypeMappings  bool
	ValidateXML       string
	CreateSampleXML   bool
}

// parseFlags 解析命令行参数
func parseFlags() *XSDConverterConfig {
	config := &XSDConverterConfig{}

	flag.StringVar(&config.XSDPath, "xsd", "", "XSD文件的路径 (必需)")
	flag.StringVar(&config.OutputPath, "output", "", "输出代码的文件路径 (可选)")
	flag.StringVar(&config.PackageName, "package", "models", "生成的代码包名 (默认: models)")
	flag.BoolVar(&config.EnableJSON, "json", false, "生成JSON兼容的标签")
	flag.BoolVar(&config.DebugMode, "debug", false, "启用调试模式")
	flag.BoolVar(&config.StrictMode, "strict", false, "启用严格模式")
	flag.BoolVar(&config.IncludeComments, "comments", true, "在生成的代码中包含注释")
	flag.BoolVar(&config.GenerateValidation, "validation", false, "生成验证代码")
	flag.BoolVar(&config.GenerateTests, "tests", false, "生成测试代码")
	flag.BoolVar(&config.GenerateBenchmarks, "benchmarks", false, "生成基准测试代码")
	flag.StringVar(&config.TestOutputPath, "test-output", "", "测试代码输出路径")
	flag.StringVar(&config.ValidationOutputPath, "validation-output", "", "验证代码输出路径")
	// 新增多语言和实用功能
	flag.StringVar(&config.TargetLanguage, "lang", "go", "目标语言 (go, java, csharp, python)")
	flag.BoolVar(&config.EnableCustomTypes, "plc", false, "启用PLC/自定义类型映射")
	flag.BoolVar(&config.ShowTypeMappings, "show-mappings", false, "显示XSD到目标语言的类型映射")
	flag.StringVar(&config.ValidateXML, "validate", "", "验证XML文件是否符合XSD规范")
	flag.BoolVar(&config.CreateSampleXML, "sample", false, "根据XSD生成示例XML")
	help := flag.Bool("help", false, "显示帮助信息")
	version := flag.Bool("version", false, "显示版本信息")
	flag.Parse()
	if *version {
		fmt.Println("XSD到多语言转换工具 v3.1 (统一解析器)")
		fmt.Println("支持完整 XML Schema 规范，多语言代码生成")
		fmt.Println("支持: Go, Java, C#, Python")
		fmt.Println("新增: 验证代码生成、测试代码生成、自定义类型映射")
		os.Exit(0)
	}
	if *help {
		showHelp()
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
	// 验证目标语言
	validLanguages := []string{"go", "java", "csharp", "python"}
	isValidLang := false
	for _, lang := range validLanguages {
		if config.TargetLanguage == lang {
			isValidLang = true
			break
		}
	}
	if !isValidLang {
		return fmt.Errorf("不支持的目标语言: %s (支持: %s)", config.TargetLanguage, strings.Join(validLanguages, ", "))
	}

	// 如果未提供输出路径，生成默认值
	if config.OutputPath == "" {
		ext := getLanguageExtension(config.TargetLanguage)
		baseName := strings.TrimSuffix(filepath.Base(config.XSDPath), filepath.Ext(config.XSDPath))
		config.OutputPath = baseName + ext
	}

	// 为额外代码生成设置默认路径
	if config.GenerateValidation && config.ValidationOutputPath == "" {
		ext := getLanguageExtension(config.TargetLanguage)
		baseName := strings.TrimSuffix(filepath.Base(config.OutputPath), filepath.Ext(config.OutputPath))
		config.ValidationOutputPath = filepath.Join("test", baseName+"_validation"+ext)
	}

	if config.GenerateTests && config.TestOutputPath == "" {
		ext := getLanguageExtension(config.TargetLanguage)
		baseName := strings.TrimSuffix(filepath.Base(config.OutputPath), filepath.Ext(config.OutputPath))
		config.TestOutputPath = filepath.Join("test", baseName+"_test"+ext)
	}

	// 验证包名
	if !isValidPackageName(config.PackageName) {
		return fmt.Errorf("无效的包名: %s", config.PackageName)
	}

	return nil
}

// getLanguageExtension 根据语言返回文件扩展名
func getLanguageExtension(lang string) string {
	switch lang {
	case "go":
		return ".go"
	case "java":
		return ".java"
	case "csharp":
		return ".cs"
	case "python":
		return ".py"
	default:
		return ".txt"
	}
}

// isValidPackageName 检查是否为有效的包名（支持多语言）
func isValidPackageName(name string) bool {
	if name == "" {
		return false
	}

	// 允许点号分隔的包名（适用于Java和C#）
	parts := strings.Split(name, ".")
	for _, part := range parts {
		if !isValidIdentifier(part) {
			return false
		}
	}
	return true
}

// isValidIdentifier 检查是否为有效的标识符
func isValidIdentifier(name string) bool {
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
	fmt.Printf("XSD到Go转换工具 - v3.1 (增强版统一解析器)\n")
	fmt.Printf("==================================================\n")
	fmt.Printf("输入文件: %s\n", config.XSDPath)
	fmt.Printf("输出文件: %s\n", config.OutputPath)
	fmt.Printf("包名: %s\n", config.PackageName)
	fmt.Printf("JSON兼容: %t\n", config.EnableJSON)
	fmt.Printf("调试模式: %t\n", config.DebugMode)
	fmt.Printf("严格模式: %t\n", config.StrictMode)
	fmt.Printf("生成验证代码: %t\n", config.GenerateValidation)
	fmt.Printf("生成测试代码: %t\n", config.GenerateTests)
	fmt.Printf("生成基准测试: %t\n", config.GenerateBenchmarks)
	fmt.Printf("------------------------------------------------\n")
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
	// 创建代码生成器配置
	genConfig := generator.NewGeneratorConfig().
		SetLanguage(generator.TargetLanguage(config.TargetLanguage)).
		SetPackage(config.PackageName).
		SetOutput(config.OutputPath)

	if config.EnableJSON {
		genConfig.JSONCompatible = true
	}
	if config.EnableCustomTypes {
		genConfig.EnableCustomTypes = true
	}
	genConfig.IncludeComments = config.IncludeComments
	genConfig.DebugMode = config.DebugMode
	genConfig.EnableValidation = config.GenerateValidation
	genConfig.EnableTestCode = config.GenerateTests

	// 创建代码生成器工厂
	factory := generator.NewCodeGeneratorFactory(genConfig)

	// 生成代码
	fmt.Printf("生成%s代码...\n", strings.ToUpper(config.TargetLanguage))
	if err := factory.GenerateCode(parser.GetGoTypes()); err != nil {
		return fmt.Errorf("生成代码失败: %v", err)
	}

	fmt.Printf("✓ 成功！%s结构已生成在: %s\n", strings.ToUpper(config.TargetLanguage), config.OutputPath)
	// 如果启用了额外的代码生成功能，使用CodeGenerator
	if config.GenerateValidation || config.GenerateTests || config.GenerateBenchmarks {
		fmt.Println("------------------------------------------------")
		fmt.Println("开始生成额外代码...")

		// 确保 test 目录存在
		if err := os.MkdirAll("test", 0755); err != nil {
			return fmt.Errorf("创建test目录失败: %v", err)
		}

		// 创建代码生成器
		codeGen := generator.NewCodeGenerator(config.PackageName, config.OutputPath)
		codeGen.SetGoTypes(parser.GetGoTypes())
		codeGen.SetJSONCompatible(config.EnableJSON)
		codeGen.SetIncludeComments(config.IncludeComments)
		codeGen.SetDebugMode(config.DebugMode)

		// 生成验证代码
		if config.GenerateValidation {
			fmt.Printf("生成验证代码到: %s\n", config.ValidationOutputPath)
			validationCode := codeGen.GenerateValidationCode()

			// 添加包声明和必要的导入
			fullValidationCode := fmt.Sprintf("package %s\n\n%s", config.PackageName, validationCode)

			if err := os.WriteFile(config.ValidationOutputPath, []byte(fullValidationCode), 0644); err != nil {
				return fmt.Errorf("写入验证代码失败: %v", err)
			}
			fmt.Printf("✓ 验证代码已生成在: %s\n", config.ValidationOutputPath)
		}
		// 生成测试代码
		if config.GenerateTests {
			fmt.Printf("生成测试代码到: %s\n", config.TestOutputPath)
			testCode := codeGen.GenerateTestCode()

			// 添加包声明和必要的导入
			fullTestCode := fmt.Sprintf("package %s\n\n%s", config.PackageName, testCode)

			if err := os.WriteFile(config.TestOutputPath, []byte(fullTestCode), 0644); err != nil {
				return fmt.Errorf("写入测试代码失败: %v", err)
			}
			fmt.Printf("✓ 测试代码已生成在: %s\n", config.TestOutputPath)
		}
		// 生成基准测试代码（如果测试代码已启用，基准测试会包含在测试文件中）
		if config.GenerateBenchmarks && !config.GenerateTests {
			ext := filepath.Ext(config.OutputPath)
			baseName := strings.TrimSuffix(filepath.Base(config.OutputPath), ext)
			benchmarkPath := filepath.Join("test", baseName+"_bench_test.go")
			fmt.Printf("生成独立基准测试代码到: %s\n", benchmarkPath)

			// 仅生成基准测试部分
			benchmarkCode := codeGen.GenerateTestCode()
			// 过滤出只包含基准测试的代码
			lines := strings.Split(benchmarkCode, "\n")
			var benchmarkLines []string
			inBenchmark := false
			for _, line := range lines {
				if strings.Contains(line, "func Benchmark") {
					inBenchmark = true
				}
				if inBenchmark {
					benchmarkLines = append(benchmarkLines, line)
					if line == "}" && !strings.Contains(line, "func Benchmark") {
						inBenchmark = false
					}
				}
			}

			fullBenchmarkCode := fmt.Sprintf("package %s\n\nimport (\n\t\"encoding/xml\"\n\t\"testing\"\n\t\"time\"\n)\n\n%s",
				config.PackageName, strings.Join(benchmarkLines, "\n"))

			if err := os.WriteFile(benchmarkPath, []byte(fullBenchmarkCode), 0644); err != nil {
				return fmt.Errorf("写入基准测试代码失败: %v", err)
			}
			fmt.Printf("✓ 基准测试代码已生成在: %s\n", benchmarkPath)
		}

		// 显示生成的类型统计
		goTypes := parser.GetGoTypes()
		if len(goTypes) > 0 {
			fmt.Printf("📊 代码生成统计:\n")
			fmt.Printf("   - 生成的类型数量: %d\n", len(goTypes))
			if config.GenerateValidation {
				fmt.Printf("   - 验证函数数量: %d\n", len(goTypes))
			}
			if config.GenerateTests {
				fmt.Printf("   - 测试函数数量: %d (每类型3个函数)\n", len(goTypes)*3)
			}
			if config.GenerateBenchmarks {
				fmt.Printf("   - 基准测试数量: %d\n", len(goTypes))
			}
		}
	}

	// 新增：处理多语言支持和实用功能
	if config.TargetLanguage != "" {
		fmt.Println("------------------------------------------------")
		fmt.Println("开始处理多语言支持和实用功能...")

		// 显示类型映射
		if config.ShowTypeMappings {
			showTypeMappings(config.TargetLanguage)
		}

		// 验证XML文件
		if config.ValidateXML != "" {
			xmlFilePath := config.ValidateXML
			fmt.Printf("验证XML文件: %s\n", xmlFilePath)
			if err := validateXMLFile(config.XSDPath, xmlFilePath); err != nil {
				fmt.Printf("XML验证失败: %v\n", err)
			}
		}

		// 生成示例XML
		if config.CreateSampleXML {
			if err := createSampleXML(config.XSDPath); err != nil {
				fmt.Printf("生成示例XML失败: %v\n", err)
			}
		}
	}

	return nil
}

// showTypeMappings 显示支持的类型映射
func showTypeMappings(targetLang string) {
	fmt.Printf("XSD 到 %s 的类型映射表\n", strings.ToUpper(targetLang))
	fmt.Println("====================")

	var mapper generator.LanguageMapper
	switch strings.ToLower(targetLang) {
	case "go":
		mapper = &generator.GoLanguageMapper{}
	case "java":
		mapper = &generator.JavaLanguageMapper{}
	case "csharp", "c#":
		mapper = &generator.CSharpLanguageMapper{}
	case "python":
		mapper = &generator.PythonLanguageMapper{}
	default:
		fmt.Printf("不支持的语言: %s\n", targetLang)
		fmt.Println("支持的语言: go, java, csharp, python")
		return
	}

	mappings := mapper.GetBuiltinTypeMappings()

	// 按类别分组显示
	categories := map[string][]generator.TypeMapping{
		"字符串类型":  {},
		"数值类型":   {},
		"布尔类型":   {},
		"日期时间类型": {},
		"二进制类型":  {},
		"PLC类型":  {},
		"其他类型":   {},
	}

	// 分类映射
	for _, mapping := range mappings {
		xsdType := mapping.XSDType

		if strings.Contains(xsdType, "string") || strings.Contains(xsdType, "String") ||
			strings.Contains(xsdType, "URI") || strings.Contains(xsdType, "Name") ||
			strings.Contains(xsdType, "ENTITY") || strings.Contains(xsdType, "QName") {
			categories["字符串类型"] = append(categories["字符串类型"], mapping)
		} else if xsdType == "boolean" || xsdType == "BOOL" {
			categories["布尔类型"] = append(categories["布尔类型"], mapping)
		} else if strings.Contains(xsdType, "int") || strings.Contains(xsdType, "INT") ||
			strings.Contains(xsdType, "long") || strings.Contains(xsdType, "short") ||
			strings.Contains(xsdType, "byte") || strings.Contains(xsdType, "BYTE") ||
			strings.Contains(xsdType, "decimal") || strings.Contains(xsdType, "float") ||
			strings.Contains(xsdType, "double") || strings.Contains(xsdType, "REAL") ||
			strings.Contains(xsdType, "WORD") || strings.Contains(xsdType, "DWORD") {
			categories["数值类型"] = append(categories["数值类型"], mapping)
		} else if strings.Contains(xsdType, "date") || strings.Contains(xsdType, "time") ||
			strings.Contains(xsdType, "Date") || strings.Contains(xsdType, "Time") ||
			strings.Contains(xsdType, "TIME") || strings.Contains(xsdType, "DT") ||
			strings.Contains(xsdType, "TOD") || xsdType == "duration" {
			categories["日期时间类型"] = append(categories["日期时间类型"], mapping)
		} else if strings.Contains(xsdType, "Binary") {
			categories["二进制类型"] = append(categories["二进制类型"], mapping)
		} else if len(xsdType) <= 6 && strings.ToUpper(xsdType) == xsdType {
			// PLC类型通常是全大写的简短名称
			categories["PLC类型"] = append(categories["PLC类型"], mapping)
		} else {
			categories["其他类型"] = append(categories["其他类型"], mapping)
		}
	}

	// 按类别显示
	categoryOrder := []string{"字符串类型", "数值类型", "布尔类型", "日期时间类型", "二进制类型", "PLC类型", "其他类型"}

	for _, category := range categoryOrder {
		mappings := categories[category]
		if len(mappings) > 0 {
			fmt.Printf("\n📁 %s:\n", category)
			for _, mapping := range mappings {
				fmt.Printf("   %-20s -> %s\n", mapping.XSDType, mapping.TargetType)
			}
		}
	}

	fmt.Printf("\n总计: %d 个类型映射\n", len(mappings))
}

// validateXMLFile 验证XML文件是否符合XSD规范
func validateXMLFile(xsdPath, xmlPath string) error {
	fmt.Printf("验证 XML 文件: %s\n", xmlPath)
	fmt.Printf("使用 XSD 规范: %s\n", xsdPath)
	fmt.Println("--------------------------------")

	// 解析XSD
	parser := xsdparser.NewUnifiedXSDParser(xsdPath, "", "temp")
	if err := parser.Parse(); err != nil {
		return fmt.Errorf("解析XSD失败: %v", err)
	}

	// 获取XSD schema
	schema := parser.GetSchema()
	if schema == nil {
		return fmt.Errorf("无法获取XSD模式")
	}

	// 创建验证器
	xmlValidator := validator.NewXSDValidator(schema)

	// 执行验证
	err := xmlValidator.ValidateXML(xmlPath)
	if err != nil {
		fmt.Printf("❌ XML验证失败: %v\n", err)
		return err
	}

	fmt.Println("✅ XML验证通过！")
	return nil
}

// createSampleXML 根据XSD生成示例XML
func createSampleXML(xsdPath string) error {
	fmt.Printf("根据 XSD 生成示例 XML: %s\n", xsdPath)
	fmt.Println("--------------------------------")
	// 解析XSD
	parser := xsdparser.NewUnifiedXSDParser(xsdPath, "", "temp")
	if err := parser.Parse(); err != nil {
		return fmt.Errorf("解析XSD失败: %v", err)
	}

	// 获取XSD schema
	schema := parser.GetSchema()
	if schema == nil {
		return fmt.Errorf("无法获取XSD模式")
	}

	// 生成示例XML文件名
	ext := filepath.Ext(xsdPath)
	baseName := strings.TrimSuffix(filepath.Base(xsdPath), ext)
	samplePath := baseName + "_sample.xml"

	// 生成示例XML内容
	xmlContent := generateSampleXMLContent(schema)

	// 写入文件
	if err := os.WriteFile(samplePath, []byte(xmlContent), 0644); err != nil {
		return fmt.Errorf("写入示例XML失败: %v", err)
	}

	fmt.Printf("✅ 示例XML已生成: %s\n", samplePath)
	return nil
}

// generateSampleXMLContent 生成示例XML内容
func generateSampleXMLContent(schema *types.XSDSchema) string {
	var builder strings.Builder

	builder.WriteString("<?xml version=\"1.0\" encoding=\"UTF-8\"?>\n")

	// 如果有根元素，生成示例内容
	if len(schema.Elements) > 0 {
		rootElement := schema.Elements[0]
		builder.WriteString(fmt.Sprintf("<%s", rootElement.Name))

		// 添加命名空间（如果有）
		if schema.TargetNamespace != "" {
			builder.WriteString(fmt.Sprintf(" xmlns=\"%s\"", schema.TargetNamespace))
		}

		builder.WriteString(">\n")

		// 生成示例内容（简化版）
		builder.WriteString("  <!-- Sample content generated from XSD -->\n")
		if rootElement.ComplexType != nil {
			generateSampleComplexType(&builder, rootElement.ComplexType, "  ")
		} else {
			builder.WriteString("  <SampleValue>Example</SampleValue>\n")
		}

		builder.WriteString(fmt.Sprintf("</%s>\n", rootElement.Name))
	} else {
		builder.WriteString("<!-- No root element found in XSD -->\n")
	}

	return builder.String()
}

// generateSampleComplexType 生成复杂类型的示例内容
func generateSampleComplexType(builder *strings.Builder, complexType *types.XSDComplexType, indent string) {
	if complexType.Sequence != nil {
		for _, element := range complexType.Sequence.Elements {
			builder.WriteString(fmt.Sprintf("%s<%s>", indent, element.Name))

			// 根据类型生成示例值
			if element.Type != "" {
				sampleValue := generateSampleValueForType(element.Type)
				builder.WriteString(sampleValue)
			} else {
				builder.WriteString("SampleValue")
			}

			builder.WriteString(fmt.Sprintf("</%s>\n", element.Name))
		}
	}

	// 添加示例属性
	for _, attr := range complexType.Attributes {
		// 属性会在元素标签中生成，这里只是占位符
		builder.WriteString(fmt.Sprintf("%s<!-- Attribute: %s -->\n", indent, attr.Name))
	}
}

// generateSampleValueForType 根据类型生成示例值
func generateSampleValueForType(typeName string) string {
	// 移除命名空间前缀
	if colonIndex := strings.LastIndex(typeName, ":"); colonIndex != -1 {
		typeName = typeName[colonIndex+1:]
	}

	switch typeName {
	case "string", "normalizedString", "token":
		return "SampleString"
	case "int", "integer", "long", "short":
		return "123"
	case "decimal", "float", "double":
		return "123.45"
	case "boolean":
		return "true"
	case "date":
		return "2023-12-25"
	case "dateTime":
		return "2023-12-25T10:30:00"
	case "time":
		return "10:30:00"
	default:
		return "SampleValue"
	}
}

// showHelp 显示帮助信息
func showHelp() {
	fmt.Println("XSD到Go转换工具 - v3.1 (增强版统一解析器)")
	fmt.Println("==================================================")
	fmt.Println("将XSD模式文件转换为Go结构体，支持完整的XML Schema规范")
	fmt.Println("")
	fmt.Println("用法:")
	fmt.Println("  xsd2go -xsd=<XSD文件路径> [选项...]")
	fmt.Println("")
	fmt.Println("必需参数:")
	fmt.Println("  -xsd string")
	fmt.Println("        XSD文件的路径")
	fmt.Println("")
	fmt.Println("基本选项:")
	fmt.Println("  -output string")
	fmt.Println("        输出Go代码的文件路径 (默认: 根据XSD文件名生成)")
	fmt.Println("  -package string")
	fmt.Println("        生成的Go代码包名 (默认: \"models\")")
	fmt.Println("  -json")
	fmt.Println("        生成JSON兼容的标签")
	fmt.Println("  -debug")
	fmt.Println("        启用调试模式")
	fmt.Println("  -strict")
	fmt.Println("        启用严格模式")
	fmt.Println("  -comments")
	fmt.Println("        在生成的代码中包含注释 (默认: true)")
	fmt.Println("")
	fmt.Println("代码生成选项:")
	fmt.Println("  -validation")
	fmt.Println("        生成验证代码")
	fmt.Println("  -tests")
	fmt.Println("        生成测试代码")
	fmt.Println("  -benchmarks")
	fmt.Println("        生成基准测试代码")
	fmt.Println("  -test-output string")
	fmt.Println("        测试代码输出路径")
	fmt.Println("  -validation-output string")
	fmt.Println("        验证代码输出路径")
	fmt.Println("")
	fmt.Println("多语言与实用功能:")
	fmt.Println("  -lang string")
	fmt.Println("        目标语言 (go, java, csharp, python) (默认: \"go\")")
	fmt.Println("  -show-mappings")
	fmt.Println("        显示XSD到目标语言的类型映射")
	fmt.Println("  -validate string")
	fmt.Println("        验证XML文件是否符合XSD规范")
	fmt.Println("  -sample")
	fmt.Println("        根据XSD生成示例XML")
	fmt.Println("")
	fmt.Println("其他选项:")
	fmt.Println("  -help")
	fmt.Println("        显示此帮助信息")
	fmt.Println("  -version")
	fmt.Println("        显示版本信息")
	fmt.Println("")
	fmt.Println("示例:")
	fmt.Println("  # 基本转换")
	fmt.Println("  xsd2go -xsd=schema.xsd")
	fmt.Println("")
	fmt.Println("  # 完整功能转换")
	fmt.Println("  xsd2go -xsd=schema.xsd -output=types.go -package=models -json -validation -tests -benchmarks")
	fmt.Println("")
	fmt.Println("  # 显示类型映射")
	fmt.Println("  xsd2go -xsd=schema.xsd -show-mappings -lang=java")
	fmt.Println("")
	fmt.Println("  # 验证XML文件")
	fmt.Println("  xsd2go -xsd=schema.xsd -validate=data.xml")
	fmt.Println("")
	fmt.Println("  # 生成示例XML")
	fmt.Println("  xsd2go -xsd=schema.xsd -sample")
}

func main() {
	// 解析命令行参数
	config := parseFlags()

	// 验证配置
	if err := validateConfig(config); err != nil {
		fmt.Fprintf(os.Stderr, "配置错误: %v\n", err)
		fmt.Println("\n使用 -help 查看帮助信息")
		os.Exit(1)
	}

	// 执行转换
	if err := runConverter(config); err != nil {
		fmt.Fprintf(os.Stderr, "转换失败: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("✅ 所有操作已完成！")
}
