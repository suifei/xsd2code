package main

import (
	"fmt"
	"log"

	"github.com/suifei/xsd2code/pkg/generator"
	"github.com/suifei/xsd2code/pkg/types"
)

func main() {
	// 示例：展示新的配置系统和多语言支持

	// 1. 基本Go代码生成（默认配置）
	fmt.Println("=== 基本Go代码生成 ===")
	generateGoCode()

	// 2. 启用PLC类型的Go代码生成
	fmt.Println("\n=== 启用PLC类型的Go代码生成 ===")
	generateGoCodeWithPLC()

	// 3. Java代码生成
	fmt.Println("\n=== Java代码生成 ===")
	generateJavaCode()

	// 4. C#代码生成
	fmt.Println("\n=== C#代码生成 ===")
	generateCSharpCode()

	// 5. 使用工厂模式和类型注册表
	fmt.Println("\n=== 使用工厂模式 ===")
	useFactoryPattern()
}

// generateGoCode 演示基本的Go代码生成
func generateGoCode() {
	config := generator.NewGeneratorConfig().
		SetPackage("example").
		SetOutput("output/basic.go").
		EnableJSON()

	// 创建一些示例类型
	goTypes := createSampleTypes()

	// 生成代码
	factory := generator.NewCodeGeneratorFactory(config)
	if err := factory.GenerateCode(goTypes); err != nil {
		log.Printf("Error generating Go code: %v", err)
	} else {
		fmt.Println("✓ Go代码生成成功: output/basic.go")
	}
}

// generateGoCodeWithPLC 演示启用PLC类型的Go代码生成
func generateGoCodeWithPLC() {
	config := generator.NewGeneratorConfig().
		SetPackage("plcopen").
		SetOutput("output/plc.go").
		EnableJSON().
		EnablePLCTypes(). // 启用PLC类型映射
		EnableDebug()

	// 创建包含PLC类型的示例
	goTypes := createPLCTypes()

	factory := generator.NewCodeGeneratorFactory(config)
	if err := factory.GenerateCode(goTypes); err != nil {
		log.Printf("Error generating PLC Go code: %v", err)
	} else {
		fmt.Println("✓ PLC Go代码生成成功: output/plc.go")
	}
}

// generateJavaCode 演示Java代码生成
func generateJavaCode() {
	config := generator.NewGeneratorConfig().
		SetLanguage(generator.LanguageJava).
		SetPackage("com.example").
		SetOutput("output/Example.java").
		EnableCustomTypes() // Java也支持自定义类型

	goTypes := createSampleTypes()

	factory := generator.NewCodeGeneratorFactory(config)
	if err := factory.GenerateCode(goTypes); err != nil {
		log.Printf("Error generating Java code: %v", err)
	} else {
		fmt.Println("✓ Java代码生成成功: output/Example.java")
	}
}

// generateCSharpCode 演示C#代码生成
func generateCSharpCode() {
	config := generator.NewGeneratorConfig().
		SetLanguage(generator.LanguageCSharp).
		SetPackage("Example.Models").
		SetOutput("output/Example.cs").
		EnableCustomTypes()

	goTypes := createSampleTypes()

	factory := generator.NewCodeGeneratorFactory(config)
	if err := factory.GenerateCode(goTypes); err != nil {
		log.Printf("Error generating C# code: %v", err)
	} else {
		fmt.Println("✓ C#代码生成成功: output/Example.cs")
	}
}

// useFactoryPattern 演示使用工厂模式和类型注册表
func useFactoryPattern() {
	// 创建类型注册表
	registry := generator.NewTypeRegistry()

	// 注册类型
	sampleTypes := createSampleTypes()
	for _, goType := range sampleTypes {
		registry.RegisterType(goType)
	}

	// 添加类型依赖关系
	registry.AddDependency("Person", "Address")

	// 生成多种语言的代码
	languages := []generator.TargetLanguage{
		generator.LanguageGo,
		generator.LanguageJava,
		generator.LanguageCSharp,
	}

	for _, lang := range languages {
		config := generator.NewGeneratorConfig().
			SetLanguage(lang).
			SetPackage("example").
			SetOutput(fmt.Sprintf("output/registry_%s%s",
				string(lang),
				getFileExtension(lang)))

		if err := registry.GenerateCode(config); err != nil {
			log.Printf("Error generating %s code: %v", lang, err)
		} else {
			fmt.Printf("✓ %s代码生成成功: %s\n", lang, config.OutputPath)
		}
	}
}

// createSampleTypes 创建示例类型用于测试
func createSampleTypes() []types.GoType {
	return []types.GoType{
		{
			Name:    "Person",
			XMLName: "person",
			IsEnum:  false,
			Fields: []types.GoField{
				{
					Name:    "Name",
					Type:    "string",
					XMLTag:  "name,attr",
					JSONTag: "name",
				},
				{
					Name:    "Age",
					Type:    "int",
					XMLTag:  "age,attr",
					JSONTag: "age",
				},
				{
					Name:       "Email",
					Type:       "*string",
					XMLTag:     "email,attr,omitempty",
					JSONTag:    "email,omitempty",
					IsOptional: true,
				},
			},
		},
		{
			Name:     "Status",
			XMLName:  "status",
			IsEnum:   true,
			BaseType: "string",
			Constants: []types.GoConstant{
				{Name: "STATUS_ACTIVE", Value: `"active"`},
				{Name: "STATUS_INACTIVE", Value: `"inactive"`},
				{Name: "STATUS_PENDING", Value: `"pending"`},
			},
		},
	}
}

// createPLCTypes 创建包含PLC类型的示例
func createPLCTypes() []types.GoType {
	return []types.GoType{
		{
			Name:    "PLCVariable",
			XMLName: "variable",
			IsEnum:  false,
			Fields: []types.GoField{
				{
					Name:    "Name",
					Type:    "string",
					XMLTag:  "name,attr",
					JSONTag: "name",
				},
				{
					Name:    "BoolValue",
					Type:    "BOOL", // 这将映射到bool（如果启用了PLC类型）
					XMLTag:  "boolValue",
					JSONTag: "boolValue",
				},
				{
					Name:    "IntValue",
					Type:    "DINT", // 这将映射到int32
					XMLTag:  "intValue",
					JSONTag: "intValue",
				},
				{
					Name:    "RealValue",
					Type:    "REAL", // 这将映射到float32
					XMLTag:  "realValue",
					JSONTag: "realValue",
				},
			},
		},
	}
}

// getFileExtension 根据语言返回文件扩展名
func getFileExtension(lang generator.TargetLanguage) string {
	switch lang {
	case generator.LanguageGo:
		return ".go"
	case generator.LanguageJava:
		return ".java"
	case generator.LanguageCSharp:
		return ".cs"
	default:
		return ".txt"
	}
}
