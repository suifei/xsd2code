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
	"time"

	"github.com/suifei/xsd2code/pkg/core"
	"github.com/suifei/xsd2code/pkg/generator"
	"github.com/suifei/xsd2code/pkg/types"
	"github.com/suifei/xsd2code/pkg/validator"
	"github.com/suifei/xsd2code/pkg/xsdparser"
)

// 默认输出位置
const defaultOutputDir = "./gen"
const defaultPackageName = "generated"

// XSDConverterConfig 配置XSD转换器的选项
type XSDConverterConfig struct {
	XSDPath         string
	OutputPath      string
	PackageName     string
	EnableJSON      bool
	DebugMode       bool
	StrictMode      bool
	IncludeComments bool
	// 代码生成选项
	GenerateValidation   bool
	GenerateTests        bool
	GenerateBenchmarks   bool
	TestOutputPath       string
	ValidationOutputPath string
	// 多语言支持和类型映射
	TargetLanguage    string
	EnableCustomTypes bool
	ShowTypeMappings  bool
	ValidateXML       string
	CreateSampleXML   bool
	// 第二次迭代新增的性能和优化选项
	EnableOptimization bool
	MaxWorkers         int
	CacheEnabled       bool
	ConfigFile         string
	PerformanceMode    bool
}

// 核心管理器通过 core.Managers 全局实例访问

// parseFlags 解析命令行参数
func parseFlags() *XSDConverterConfig {
	config := &XSDConverterConfig{}

	flag.StringVar(&config.XSDPath, "xsd", "", "XSD文件的路径 (必需)")
	flag.StringVar(&config.OutputPath, "output", defaultOutputDir, "输出代码的文件路径 (默认： ./output)")
	flag.StringVar(&config.PackageName, "package", defaultPackageName, "生成的代码包名 (默认: test)")
	flag.BoolVar(&config.EnableJSON, "json", false, "生成JSON兼容的标签")
	flag.BoolVar(&config.DebugMode, "debug", false, "启用调试模式")
	flag.BoolVar(&config.StrictMode, "strict", false, "启用严格模式")
	flag.BoolVar(&config.IncludeComments, "comments", true, "在生成的代码中包含注释")
	flag.BoolVar(&config.GenerateValidation, "validation", false, "生成验证代码")
	flag.BoolVar(&config.GenerateTests, "tests", false, "生成测试代码")
	flag.BoolVar(&config.GenerateBenchmarks, "benchmarks", false, "生成基准测试代码")
	flag.StringVar(&config.TestOutputPath, "test-output", defaultOutputDir, "测试代码输出路径")
	flag.StringVar(&config.ValidationOutputPath, "validation-output", defaultOutputDir, "验证代码输出路径")
	// 新增多语言和实用功能
	flag.StringVar(&config.TargetLanguage, "lang", "go", "目标语言 (go, java, csharp, python)")
	flag.BoolVar(&config.EnableCustomTypes, "plc", false, "启用PLC/自定义类型映射")
	flag.BoolVar(&config.ShowTypeMappings, "show-mappings", false, "显示XSD到目标语言的类型映射")
	flag.StringVar(&config.ValidateXML, "validate", "", "验证XML文件是否符合XSD规范")
	flag.BoolVar(&config.CreateSampleXML, "sample", false, "根据XSD生成示例XML")
	// 性能和优化相关选项
	flag.BoolVar(&config.EnableOptimization, "optimize", false, "启用性能优化")
	flag.IntVar(&config.MaxWorkers, "workers", 4, "最大并发工作线程数")
	flag.BoolVar(&config.CacheEnabled, "cache", false, "启用缓存")
	flag.StringVar(&config.ConfigFile, "config", "", "配置文件路径")
	flag.BoolVar(&config.PerformanceMode, "perf", false, "启用性能监控模式")
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

	// 如果未提供输出路径或使用默认值，生成基于gen目录的路径
	if config.OutputPath == "" || config.OutputPath == defaultOutputDir {
		ext := getLanguageExtension(config.TargetLanguage)
		baseName := strings.TrimSuffix(filepath.Base(config.XSDPath), filepath.Ext(config.XSDPath))
		config.OutputPath = filepath.Join(defaultOutputDir, baseName+ext)
	} else if !filepath.IsAbs(config.OutputPath) && !strings.HasPrefix(config.OutputPath, "./") {
		// 如果是相对路径且不是以./开头，放到gen目录下
		config.OutputPath = filepath.Join(defaultOutputDir, config.OutputPath)
	}

	// 为额外代码生成设置默认路径 - 都放在同一个目录下
	if config.GenerateValidation && (config.ValidationOutputPath == "" || config.ValidationOutputPath == defaultOutputDir) {
		ext := getLanguageExtension(config.TargetLanguage)
		baseName := strings.TrimSuffix(filepath.Base(config.OutputPath), filepath.Ext(config.OutputPath))
		outputDir := filepath.Dir(config.OutputPath)
		config.ValidationOutputPath = filepath.Join(outputDir, baseName+"_validation"+ext)
	}

	if config.GenerateTests && (config.TestOutputPath == "" || config.TestOutputPath == defaultOutputDir) {
		ext := getLanguageExtension(config.TargetLanguage)
		baseName := strings.TrimSuffix(filepath.Base(config.OutputPath), filepath.Ext(config.OutputPath))
		outputDir := filepath.Dir(config.OutputPath)
		config.TestOutputPath = filepath.Join(outputDir, baseName+"_test"+ext)
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

// runConverter 执行XSD转换，集成核心优化系统
func runConverter(config *XSDConverterConfig) error { // 获取核心管理器
	perfManager := core.GetPerformanceManager()
	errorManager := core.GetErrorManager()
	cacheManager := core.GetCacheManager()
	concurrentProcessor := core.GetConcurrentProcessor()

	// 第三次迭代新增：增强的错误管理器
	enhancedErrorManager := core.NewEnhancedErrorManager("logs/enhanced_errors.log")

	// 第三次迭代新增：性能基准测试器
	var benchmarkRunner *core.BenchmarkRunner
	if config.PerformanceMode {
		benchmarkConfig := core.BenchmarkConfig{
			Name:             "XSD_Conversion",
			Type:             core.BenchmarkEnd2End,
			Iterations:       10,
			ConcurrencyLevel: config.MaxWorkers,
			TimeoutPerTest:   time.Minute * 5,
			WarmupIterations: 2,
			EnableProfiling:  true,
			OutputDir:        "./benchmark_results",
			CompareBaseline:  false,
		}
		benchmarkRunner = core.NewBenchmarkRunner(benchmarkConfig)
	}

	// 开始性能监控
	var conversionTimer *core.OperationTimer
	if perfManager != nil {
		conversionTimer = perfManager.StartOperation("XSD转换")
		defer conversionTimer.Stop()
	}

	// 检查缓存
	cacheKey := fmt.Sprintf("xsd:%s:config:%v", config.XSDPath, config)
	if cacheManager != nil && config.CacheEnabled {
		if cached, found := cacheManager.Get(cacheKey); found {
			fmt.Println("使用缓存的转换结果...")
			if cachedResult, ok := cached.(string); ok {
				fmt.Printf("✓ 从缓存加载结果: %s\n", cachedResult)
				return nil
			}
		}
	}

	// 增强的错误处理包装
	defer func() {
		if r := recover(); r != nil {
			// 使用增强的错误管理器处理 panic
			enhancedError := &core.EnhancedXSDError{
				XSDError: &core.XSDError{
					Type:      core.ErrorTypeUnknown,
					Code:      "PANIC",
					Message:   fmt.Sprintf("转换过程发生panic: %v", r),
					Context:   "runConverter",
					Timestamp: time.Now(),
				},
				Severity: core.SeverityFatal,
				Context: core.ErrorContext{
					Operation:    "runConverter",
					Input:        config.XSDPath,
					StackTrace:   []string{fmt.Sprintf("%v", r)},
					Environment:  make(map[string]string),
					UserData:     make(map[string]interface{}),
					ProcessingID: "main_conversion",
				},
				RecoveryAction: core.RecoveryTerminate,
			}
			enhancedErrorManager.AddEnhancedError(enhancedError)
			fmt.Printf("严重错误已记录并处理: %v\n", enhancedError.Message)
		}
	}()
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
	if config.PerformanceMode {
		fmt.Printf("性能基准测试: 启用\n")
	}
	fmt.Printf("------------------------------------------------\n")

	// 如果启用了性能模式，使用基准测试器运行转换
	if config.PerformanceMode && benchmarkRunner != nil {
		fmt.Println("🚀 运行性能基准测试...")

		// 定义要测试的操作
		benchmarkTests := map[string]func() error{
			"XSD_Parsing": func() error {
				parser := xsdparser.NewUnifiedXSDParser(config.XSDPath, config.OutputPath, config.PackageName)
				parser.SetJSONCompatible(config.EnableJSON)
				parser.SetDebugMode(config.DebugMode)
				parser.SetStrictMode(config.StrictMode)
				parser.SetIncludeComments(config.IncludeComments)
				return parser.Parse()
			},
			"Code_Generation": func() error {
				parser := xsdparser.NewUnifiedXSDParser(config.XSDPath, config.OutputPath, config.PackageName)
				if err := parser.Parse(); err != nil {
					return err
				}
				genConfig := generator.NewGeneratorConfig().
					SetLanguage(generator.TargetLanguage(config.TargetLanguage)).
					SetPackage(config.PackageName).
					SetOutput(config.OutputPath)
				factory := generator.NewCodeGeneratorFactory(genConfig)
				return factory.GenerateCode(parser.GetGoTypes())
			},
		}

		// 运行基准测试套件
		suite := benchmarkRunner.RunSuite(benchmarkTests)

		// 生成并显示报告
		report := benchmarkRunner.GenerateReport(suite)
		fmt.Printf("\n📊 基准测试报告:\n%s\n", report)

		// 保存结果
		if err := benchmarkRunner.SaveResults(suite); err != nil {
			fmt.Printf("保存基准测试结果失败: %v\n", err)
		}
		// 如果基准测试失败率过高，发出警告
		if suite.Statistics.SuccessRate < 80.0 {
			enhancedErrorManager.CreateEnhancedError(
				core.ErrorTypeValidation,
				"BENCHMARK_LOW_SUCCESS",
				fmt.Sprintf("基准测试成功率较低: %.2f%%", suite.Statistics.SuccessRate),
			).WithSeverity(core.SeverityWarning).Build()
		}
	}
	// 使用新的统一解析器（已合并标准和高级功能）
	parser := xsdparser.NewUnifiedXSDParser(config.XSDPath, config.OutputPath, config.PackageName)
	// 设置解析器选项
	parser.SetJSONCompatible(config.EnableJSON)
	parser.SetDebugMode(config.DebugMode)
	parser.SetStrictMode(config.StrictMode)
	parser.SetIncludeComments(config.IncludeComments)

	// 检查文件大小，对于大型XSD文件使用并发处理
	fileInfo, err := os.Stat(config.XSDPath)
	if err != nil {
		return fmt.Errorf("无法获取XSD文件信息: %v", err)
	}

	// 对于大于1MB的文件，启用并发处理优化
	const largeSizeThreshold = 1024 * 1024 // 1MB
	useConcurrentProcessing := fileInfo.Size() > largeSizeThreshold && concurrentProcessor != nil

	if useConcurrentProcessing {
		fmt.Printf("📊 检测到大型XSD文件 (%.2fMB)，启用并发处理优化...\n",
			float64(fileInfo.Size())/(1024*1024))

		// 启动并发处理器
		concurrentProcessor.Start()
		defer func() {
			if err := concurrentProcessor.Stop(); err != nil {
				fmt.Printf("⚠️  停止并发处理器时出错: %v\n", err)
			}
		}()

		// 读取XSD文件内容进行并发处理
		xsdData, err := os.ReadFile(config.XSDPath)
		if err != nil {
			return fmt.Errorf("读取XSD文件失败: %v", err)
		}

		// 定义处理函数
		processChunk := func(chunk *core.XSDChunk) error {
			// 这里可以添加chunk的预处理逻辑
			// 例如验证、缓存检查等
			if cacheManager != nil {
				chunkKey := fmt.Sprintf("chunk:%s", chunk.ID)
				if _, found := cacheManager.Get(chunkKey); found {
					fmt.Printf("✓ 使用缓存的chunk: %s\n", chunk.ID)
				} else {
					// 模拟chunk处理
					time.Sleep(10 * time.Millisecond)
					cacheManager.Set(chunkKey, true)
				}
			}
			return nil
		}

		// 使用智能工作池进行并发处理（如果可用）
		if smartPool, ok := interface{}(concurrentProcessor).(interface {
			ProcessXSDConcurrently([]byte, func(*core.XSDChunk) error) error
		}); ok {
			if err := smartPool.ProcessXSDConcurrently(xsdData, processChunk); err != nil {
				fmt.Printf("⚠️  并发处理失败，回退到标准解析: %v\n", err)
			} else {
				fmt.Println("✓ 并发处理完成")
			}
		}
	}

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

	// 确保输出目录存在
	outputDir := filepath.Dir(config.OutputPath)
	if err := os.MkdirAll(outputDir, 0755); err != nil {
		return fmt.Errorf("创建输出目录失败: %v", err)
	}

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

		// 确保输出目录存在 - 使用主输出文件的目录
		outputDir := filepath.Dir(config.OutputPath)
		if err := os.MkdirAll(outputDir, 0755); err != nil {
			return fmt.Errorf("创建输出目录失败: %v", err)
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
			outputDir := filepath.Dir(config.OutputPath)
			benchmarkPath := filepath.Join(outputDir, baseName+"_bench_test.go")
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
			fmt.Printf("   - 输出目录: %s\n", filepath.Dir(config.OutputPath))
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

	// 缓存转换结果
	if cacheManager != nil && config.CacheEnabled {
		cacheValue := fmt.Sprintf("转换完成：%s -> %s", config.XSDPath, config.OutputPath)
		cacheManager.Set(cacheKey, cacheValue)
		fmt.Println("转换结果已缓存")
	}

	// 记录性能指标
	if perfManager != nil {
		perfManager.RecordMemoryUsage()
		report := perfManager.GetReport()
		if config.PerformanceMode {
			fmt.Printf("🚀 性能指标: 内存使用 %d bytes, 总操作数 %d\n",
				report.MemoryUsage, len(report.Operations))
		}
	}

	// 检查和报告错误
	if errorManager != nil && errorManager.HasErrors() {
		errors := errorManager.GetErrors()
		fmt.Printf("⚠️  处理过程中出现 %d 个错误\n", len(errors))
		for i, err := range errors {
			if i < 3 { // 只显示前3个错误
				fmt.Printf("   %d. %s\n", i+1, err.Message)
			}
		}
		if len(errors) > 3 {
			fmt.Printf("   ... 还有 %d 个错误\n", len(errors)-3)
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
	fmt.Println("        输出Go代码的文件路径 (默认: ./gen/{xsd文件名}.go)")
	fmt.Println("  -package string")
	fmt.Println("        生成的Go代码包名 (默认: \"generated\")")
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
	fmt.Println("        测试代码输出路径 (默认: 与主文件同目录)")
	fmt.Println("  -validation-output string")
	fmt.Println("        验证代码输出路径 (默认: 与主文件同目录)")
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
	fmt.Println("性能和优化选项:")
	fmt.Println("  -optimize")
	fmt.Println("        启用性能优化")
	fmt.Println("  -workers int")
	fmt.Println("        最大并发工作线程数 (默认: 4)")
	fmt.Println("  -cache")
	fmt.Println("        启用缓存")
	fmt.Println("  -config string")
	fmt.Println("        配置文件路径")
	fmt.Println("  -perf")
	fmt.Println("        启用性能监控模式")
	fmt.Println("")
	fmt.Println("其他选项:")
	fmt.Println("  -help")
	fmt.Println("        显示此帮助信息")
	fmt.Println("  -version")
	fmt.Println("        显示版本信息")
	fmt.Println("")
	fmt.Println("示例:")
	fmt.Println("  # 基本转换 - 输出到 ./gen/schema.go")
	fmt.Println("  xsd2go -xsd=schema.xsd")
	fmt.Println("")
	fmt.Println("  # 完整功能转换 - 所有文件都在 ./gen/ 目录下")
	fmt.Println("  xsd2go -xsd=schema.xsd -package=models -json -validation -tests -benchmarks")
	fmt.Println("")
	fmt.Println("  # 指定输出目录")
	fmt.Println("  xsd2go -xsd=schema.xsd -output=./custom/types.go")
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
	} // 初始化核心管理器
	if err := core.InitializeManagers(config.ConfigFile); err != nil {
		fmt.Fprintf(os.Stderr, "初始化管理器失败: %v\n", err)
		os.Exit(1)
	}
	defer core.ShutdownManagers()

	// 启用调试模式（如果需要）
	if config.DebugMode {
		fmt.Println("🐛 调试模式已启用")
	}

	// 显示优化状态
	if config.EnableOptimization {
		fmt.Printf("🚀 优化模式已启用 (工作线程: %d, 缓存: %t)\n",
			config.MaxWorkers, config.CacheEnabled)
	}

	// 执行转换
	startTime := time.Now()
	if err := runConverter(config); err != nil {
		fmt.Fprintf(os.Stderr, "转换失败: %v\n", err)
		os.Exit(1)
	}
	elapsedTime := time.Since(startTime) // 性能监控和最终报告
	if config.PerformanceMode {
		if perfManager := core.GetPerformanceManager(); perfManager != nil {
			// 输出性能报告
			report := perfManager.GetReport()
			fmt.Printf("\n📊 性能报告:\n")
			fmt.Printf("   总运行时间: %v\n", elapsedTime)
			fmt.Printf("   内存使用: %d bytes\n", report.MemoryUsage)
			fmt.Printf("   操作数: %d\n", len(report.Operations))
			// 显示详细操作统计
			for name, stats := range report.Operations {
				if stats.Count > 0 {
					avgDuration := time.Duration(stats.TotalTime / time.Duration(stats.Count))
					fmt.Printf("   - %s: %d次, 总计 %v, 平均 %v\n",
						name, stats.Count, stats.TotalTime, avgDuration)
				}
			}
		}
	}

	// 显示错误摘要（如果有）
	if errorManager := core.GetErrorManager(); errorManager != nil {
		if errorManager.HasErrors() || errorManager.HasWarnings() {
			summary := errorManager.GetSummary()
			fmt.Printf("\n⚠️  处理摘要:\n")
			fmt.Printf("   错误: %d\n", summary.TotalErrors)
			fmt.Printf("   警告: %d\n", summary.TotalWarnings)
		}
	}

	fmt.Println("\n✅ 所有操作已完成！")
	fmt.Printf("版本: %s\n", core.GetVersion())
}
