package core

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"sync"
	"sync/atomic"
	"time"
)

// Config 主配置结构
type Config struct {
	// 基本配置
	Input       InputConfig       `json:"input" yaml:"input"`
	Output      OutputConfig      `json:"output" yaml:"output"`
	Performance PerformanceConfig `json:"performance" yaml:"performance"`
	Generation  GenerationConfig  `json:"generation" yaml:"generation"`
	Validation  ValidationConfig  `json:"validation" yaml:"validation"`
	Logging     LoggingConfig     `json:"logging" yaml:"logging"`
	Features    FeatureConfig     `json:"features" yaml:"features"`
}

// InputConfig 输入配置
type InputConfig struct {
	XSDPath           string            `json:"xsd_path" yaml:"xsd_path"`
	Encoding          string            `json:"encoding" yaml:"encoding"`
	StrictMode        bool              `json:"strict_mode" yaml:"strict_mode"`
	IgnoreImports     bool              `json:"ignore_imports" yaml:"ignore_imports"`
	CustomNamespaces  map[string]string `json:"custom_namespaces" yaml:"custom_namespaces"`
	MaxFileSize       int64             `json:"max_file_size_mb" yaml:"max_file_size_mb"`
	PreprocessorRules []string          `json:"preprocessor_rules" yaml:"preprocessor_rules"`
}

// OutputConfig 输出配置
type OutputConfig struct {
	OutputPath      string            `json:"output_path" yaml:"output_path"`
	PackageName     string            `json:"package_name" yaml:"package_name"`
	TargetLanguage  string            `json:"target_language" yaml:"target_language"`
	FileFormat      string            `json:"file_format" yaml:"file_format"`
	IncludeComments bool              `json:"include_comments" yaml:"include_comments"`
	CustomTemplates map[string]string `json:"custom_templates" yaml:"custom_templates"`
	OutputStructure OutputStructure   `json:"output_structure" yaml:"output_structure"`
}

// OutputStructure 输出结构配置
type OutputStructure struct {
	SingleFile     bool     `json:"single_file" yaml:"single_file"`
	FilePerType    bool     `json:"file_per_type" yaml:"file_per_type"`
	GroupByFeature bool     `json:"group_by_feature" yaml:"group_by_feature"`
	ExcludeTypes   []string `json:"exclude_types" yaml:"exclude_types"`
	IncludeTypes   []string `json:"include_types" yaml:"include_types"`
}

// PerformanceConfig 性能配置
type PerformanceConfig struct {
	MaxWorkers      int           `json:"max_workers" yaml:"max_workers"`
	MemoryLimit     string        `json:"memory_limit" yaml:"memory_limit"`
	CacheEnabled    bool          `json:"cache_enabled" yaml:"cache_enabled"`
	CacheDirectory  string        `json:"cache_directory" yaml:"cache_directory"`
	CacheTTL        time.Duration `json:"cache_ttl" yaml:"cache_ttl"`
	ParallelEnabled bool          `json:"parallel_enabled" yaml:"parallel_enabled"`
	StreamingMode   bool          `json:"streaming_mode" yaml:"streaming_mode"`
	MemoryPoolSize  int           `json:"memory_pool_size" yaml:"memory_pool_size"`
	GCThreshold     string        `json:"gc_threshold" yaml:"gc_threshold"`
	Timeout         time.Duration `json:"timeout" yaml:"timeout"`
}

// GenerationConfig 代码生成配置
type GenerationConfig struct {
	EnableJSON        bool              `json:"enable_json" yaml:"enable_json"`
	EnableValidation  bool              `json:"enable_validation" yaml:"enable_validation"`
	EnableTests       bool              `json:"enable_tests" yaml:"enable_tests"`
	EnableBenchmarks  bool              `json:"enable_benchmarks" yaml:"enable_benchmarks"`
	EnableXMLTags     bool              `json:"enable_xml_tags" yaml:"enable_xml_tags"`
	EnableCustomTypes bool              `json:"enable_custom_types" yaml:"enable_custom_types"`
	TypeMappings      map[string]string `json:"type_mappings" yaml:"type_mappings"`
	CodeStyle         CodeStyleConfig   `json:"code_style" yaml:"code_style"`
	OptimizationLevel int               `json:"optimization_level" yaml:"optimization_level"`
}

// CodeStyleConfig 代码风格配置
type CodeStyleConfig struct {
	UsePointers      bool `json:"use_pointers" yaml:"use_pointers"`
	OmitEmpty        bool `json:"omit_empty" yaml:"omit_empty"`
	CamelCase        bool `json:"camel_case" yaml:"camel_case"`
	PreferInterfaces bool `json:"prefer_interfaces" yaml:"prefer_interfaces"`
	GenerateGetters  bool `json:"generate_getters" yaml:"generate_getters"`
	GenerateSetters  bool `json:"generate_setters" yaml:"generate_setters"`
	IndentSize       int  `json:"indent_size" yaml:"indent_size"`
	MaxLineLength    int  `json:"max_line_length" yaml:"max_line_length"`
	SortFields       bool `json:"sort_fields" yaml:"sort_fields"`
}

// ValidationConfig 验证配置
type ValidationConfig struct {
	StrictValidation      bool     `json:"strict_validation" yaml:"strict_validation"`
	ValidateConstraints   bool     `json:"validate_constraints" yaml:"validate_constraints"`
	ValidateTypes         bool     `json:"validate_types" yaml:"validate_types"`
	GenerateValidators    bool     `json:"generate_validators" yaml:"generate_validators"`
	CustomValidationRules []string `json:"custom_validation_rules" yaml:"custom_validation_rules"`
	FailOnWarnings        bool     `json:"fail_on_warnings" yaml:"fail_on_warnings"`
}

// LoggingConfig 日志配置
type LoggingConfig struct {
	Level           string `json:"level" yaml:"level"`
	Format          string `json:"format" yaml:"format"`
	OutputFile      string `json:"output_file" yaml:"output_file"`
	EnableConsole   bool   `json:"enable_console" yaml:"enable_console"`
	EnableFile      bool   `json:"enable_file" yaml:"enable_file"`
	MaxFileSize     string `json:"max_file_size" yaml:"max_file_size"`
	MaxBackups      int    `json:"max_backups" yaml:"max_backups"`
	EnableMetrics   bool   `json:"enable_metrics" yaml:"enable_metrics"`
	EnableProfiling bool   `json:"enable_profiling" yaml:"enable_profiling"`
}

// FeatureConfig 功能特性配置
type FeatureConfig struct {
	ShowTypeMappings     bool     `json:"show_type_mappings" yaml:"show_type_mappings"`
	CreateSampleXML      bool     `json:"create_sample_xml" yaml:"create_sample_xml"`
	ValidateXMLFile      string   `json:"validate_xml_file" yaml:"validate_xml_file"`
	ExperimentalFeatures []string `json:"experimental_features" yaml:"experimental_features"`
	DebugMode            bool     `json:"debug_mode" yaml:"debug_mode"`
	VerboseOutput        bool     `json:"verbose_output" yaml:"verbose_output"`
}

// ConfigManager 配置管理器
type ConfigManager struct {
	config       *Config
	errorManager *ErrorManager
	configPaths  []string
	envPrefix    string
	watcher      *ConfigWatcher
	callbacks    []ConfigChangeCallback
	mutex        sync.RWMutex
}

// ConfigChangeCallback 配置变更回调
type ConfigChangeCallback func(oldConfig, newConfig *Config) error

// ConfigWatcher 配置文件监视器
type ConfigWatcher struct {
	filePath    string
	lastModTime time.Time
	ticker      *time.Ticker
	stopChan    chan struct{}
	manager     *ConfigManager
	running     atomic.Bool
}

// EnvironmentVariables 环境变量映射
type EnvironmentVariables struct {
	// 输入相关
	XSDPath     string `env:"XSD2CODE_XSD_PATH"`
	Encoding    string `env:"XSD2CODE_ENCODING"`
	StrictMode  string `env:"XSD2CODE_STRICT_MODE"`
	MaxFileSize string `env:"XSD2CODE_MAX_FILE_SIZE"`

	// 输出相关
	OutputPath     string `env:"XSD2CODE_OUTPUT_PATH"`
	PackageName    string `env:"XSD2CODE_PACKAGE_NAME"`
	TargetLanguage string `env:"XSD2CODE_TARGET_LANGUAGE"`
	SingleFile     string `env:"XSD2CODE_SINGLE_FILE"`

	// 性能相关
	MaxWorkers      string `env:"XSD2CODE_MAX_WORKERS"`
	MemoryLimit     string `env:"XSD2CODE_MEMORY_LIMIT"`
	CacheEnabled    string `env:"XSD2CODE_CACHE_ENABLED"`
	CacheDirectory  string `env:"XSD2CODE_CACHE_DIRECTORY"`
	CacheTTL        string `env:"XSD2CODE_CACHE_TTL"`
	ParallelEnabled string `env:"XSD2CODE_PARALLEL_ENABLED"`

	// 生成相关
	EnableJSON       string `env:"XSD2CODE_ENABLE_JSON"`
	EnableValidation string `env:"XSD2CODE_ENABLE_VALIDATION"`
	EnableTests      string `env:"XSD2CODE_ENABLE_TESTS"`
	EnableBenchmarks string `env:"XSD2CODE_ENABLE_BENCHMARKS"`

	// 日志相关
	LogLevel      string `env:"XSD2CODE_LOG_LEVEL"`
	LogFile       string `env:"XSD2CODE_LOG_FILE"`
	DebugMode     string `env:"XSD2CODE_DEBUG_MODE"`
	EnableTracing string `env:"XSD2CODE_ENABLE_TRACING"`
}

// NewConfigManager 创建配置管理器
func NewConfigManager(errorManager *ErrorManager) *ConfigManager {
	cm := &ConfigManager{
		config:       NewDefaultConfig(),
		errorManager: errorManager,
		envPrefix:    "XSD2CODE",
		configPaths: []string{
			"./xsd2code.json",
			"./xsd2code.yaml",
			"./config/xsd2code.json",
			"./config/xsd2code.yaml",
			"~/.xsd2code/config.json",
			"~/.xsd2code/config.yaml",
			"/etc/xsd2code/config.json",
			"/etc/xsd2code/config.yaml",
		},
		callbacks: make([]ConfigChangeCallback, 0),
	}

	// 自动加载环境变量
	cm.LoadFromEnvironment()

	return cm
}

// NewDefaultConfig 创建默认配置
func NewDefaultConfig() *Config {
	return &Config{
		Input: InputConfig{
			Encoding:          "UTF-8",
			StrictMode:        false,
			IgnoreImports:     false,
			CustomNamespaces:  make(map[string]string),
			MaxFileSize:       100, // 100MB
			PreprocessorRules: []string{},
		},
		Output: OutputConfig{
			PackageName:     "models",
			TargetLanguage:  "go",
			FileFormat:      "go",
			IncludeComments: true,
			CustomTemplates: make(map[string]string),
			OutputStructure: OutputStructure{
				SingleFile:     false,
				FilePerType:    true,
				GroupByFeature: false,
				ExcludeTypes:   []string{},
				IncludeTypes:   []string{},
			},
		},
		Performance: PerformanceConfig{
			MaxWorkers:      4,
			MemoryLimit:     "1GB",
			CacheEnabled:    true,
			CacheDirectory:  "./cache",
			CacheTTL:        time.Hour * 24,
			ParallelEnabled: true,
			StreamingMode:   false,
			MemoryPoolSize:  10,
			GCThreshold:     "50MB",
			Timeout:         time.Minute * 30,
		},
		Generation: GenerationConfig{
			EnableJSON:        false,
			EnableValidation:  false,
			EnableTests:       false,
			EnableBenchmarks:  false,
			EnableXMLTags:     true,
			EnableCustomTypes: false,
			TypeMappings:      make(map[string]string),
			CodeStyle: CodeStyleConfig{
				UsePointers:      true,
				OmitEmpty:        true,
				CamelCase:        true,
				PreferInterfaces: false,
				GenerateGetters:  false,
				GenerateSetters:  false,
				IndentSize:       4,
				MaxLineLength:    120,
				SortFields:       true,
			},
			OptimizationLevel: 1,
		},
		Validation: ValidationConfig{
			StrictValidation:    false,
			ValidateConstraints: false,
			ValidateTypes:       true,
			GenerateValidators:  false,
		}, Logging: LoggingConfig{
			Level:         "info",
			OutputFile:    "logs/xsd2code.log",
			MaxBackups:    5,
			EnableConsole: true,
			Format:        "json",
		},
		Features: FeatureConfig{
			ExperimentalFeatures: []string{},
			DebugMode:            false,
			VerboseOutput:        false,
		},
	}
}

// LoadFromFile 从文件加载配置
func (cm *ConfigManager) LoadFromFile(filePath string) error {
	cm.mutex.Lock()
	defer cm.mutex.Unlock()

	if filePath == "" {
		// 尝试自动发现配置文件
		for _, path := range cm.configPaths {
			if expandedPath := cm.expandPath(path); cm.fileExists(expandedPath) {
				filePath = expandedPath
				break
			}
		}

		if filePath == "" {
			return NewConfigError("CONFIG_FILE_NOT_FOUND", "No configuration file found").Build()
		}
	}
	data, err := os.ReadFile(filePath)
	if err != nil {
		return NewIOError("CONFIG_READ_ERROR", fmt.Sprintf("Failed to read config file: %s", filePath)).Build()
	}

	oldConfig := cm.cloneConfig()
	newConfig := NewDefaultConfig()

	// 根据文件扩展名选择解析器
	ext := strings.ToLower(filepath.Ext(filePath))
	switch ext {
	case ".json":
		if err := json.Unmarshal(data, newConfig); err != nil {
			return NewConfigError("CONFIG_PARSE_ERROR", fmt.Sprintf("Failed to parse JSON config: %v", err)).Build()
		}
	case ".yaml", ".yml":
		if err := cm.parseYAML(data, newConfig); err != nil {
			return NewConfigError("CONFIG_PARSE_ERROR", fmt.Sprintf("Failed to parse YAML config: %v", err)).Build()
		}
	default:
		return NewConfigError("CONFIG_UNSUPPORTED_FORMAT", fmt.Sprintf("Unsupported config format: %s", ext)).Build()
	}

	// 验证配置
	if err := cm.validateConfig(newConfig); err != nil {
		return fmt.Errorf("configuration validation failed: %w", err)
	}

	cm.config = newConfig

	// 启动文件监视器
	cm.startWatcher(filePath)

	// 触发配置变更回调
	go cm.notifyConfigChange(oldConfig, newConfig)

	return nil
}

// LoadFromEnvironment 从环境变量加载配置
func (cm *ConfigManager) LoadFromEnvironment() error {
	cm.mutex.Lock()
	defer cm.mutex.Unlock()

	envVars := &EnvironmentVariables{}
	cm.populateEnvironmentVariables(envVars)

	// 将环境变量应用到配置
	if err := cm.applyEnvironmentVariables(envVars); err != nil {
		return fmt.Errorf("failed to apply environment variables: %w", err)
	}

	return nil
}

// populateEnvironmentVariables 填充环境变量
func (cm *ConfigManager) populateEnvironmentVariables(envVars *EnvironmentVariables) {
	envVars.XSDPath = os.Getenv("XSD2CODE_XSD_PATH")
	envVars.Encoding = os.Getenv("XSD2CODE_ENCODING")
	envVars.StrictMode = os.Getenv("XSD2CODE_STRICT_MODE")
	envVars.MaxFileSize = os.Getenv("XSD2CODE_MAX_FILE_SIZE")

	envVars.OutputPath = os.Getenv("XSD2CODE_OUTPUT_PATH")
	envVars.PackageName = os.Getenv("XSD2CODE_PACKAGE_NAME")
	envVars.TargetLanguage = os.Getenv("XSD2CODE_TARGET_LANGUAGE")
	envVars.SingleFile = os.Getenv("XSD2CODE_SINGLE_FILE")

	envVars.MaxWorkers = os.Getenv("XSD2CODE_MAX_WORKERS")
	envVars.MemoryLimit = os.Getenv("XSD2CODE_MEMORY_LIMIT")
	envVars.CacheEnabled = os.Getenv("XSD2CODE_CACHE_ENABLED")
	envVars.CacheDirectory = os.Getenv("XSD2CODE_CACHE_DIRECTORY")
	envVars.CacheTTL = os.Getenv("XSD2CODE_CACHE_TTL")
	envVars.ParallelEnabled = os.Getenv("XSD2CODE_PARALLEL_ENABLED")

	envVars.EnableJSON = os.Getenv("XSD2CODE_ENABLE_JSON")
	envVars.EnableValidation = os.Getenv("XSD2CODE_ENABLE_VALIDATION")
	envVars.EnableTests = os.Getenv("XSD2CODE_ENABLE_TESTS")
	envVars.EnableBenchmarks = os.Getenv("XSD2CODE_ENABLE_BENCHMARKS")

	envVars.LogLevel = os.Getenv("XSD2CODE_LOG_LEVEL")
	envVars.LogFile = os.Getenv("XSD2CODE_LOG_FILE")
	envVars.DebugMode = os.Getenv("XSD2CODE_DEBUG_MODE")
	envVars.EnableTracing = os.Getenv("XSD2CODE_ENABLE_TRACING")
}

// applyEnvironmentVariables 应用环境变量到配置
func (cm *ConfigManager) applyEnvironmentVariables(envVars *EnvironmentVariables) error {
	// 输入配置
	if envVars.XSDPath != "" {
		cm.config.Input.XSDPath = envVars.XSDPath
	}
	if envVars.Encoding != "" {
		cm.config.Input.Encoding = envVars.Encoding
	}
	if envVars.StrictMode != "" {
		if val, err := strconv.ParseBool(envVars.StrictMode); err == nil {
			cm.config.Input.StrictMode = val
		}
	}
	if envVars.MaxFileSize != "" {
		if val, err := strconv.ParseInt(envVars.MaxFileSize, 10, 64); err == nil {
			cm.config.Input.MaxFileSize = val
		}
	}

	// 输出配置
	if envVars.OutputPath != "" {
		cm.config.Output.OutputPath = envVars.OutputPath
	}
	if envVars.PackageName != "" {
		cm.config.Output.PackageName = envVars.PackageName
	}
	if envVars.TargetLanguage != "" {
		cm.config.Output.TargetLanguage = envVars.TargetLanguage
	}
	if envVars.SingleFile != "" {
		if val, err := strconv.ParseBool(envVars.SingleFile); err == nil {
			cm.config.Output.OutputStructure.SingleFile = val
		}
	}

	// 性能配置
	if envVars.MaxWorkers != "" {
		if val, err := strconv.Atoi(envVars.MaxWorkers); err == nil {
			cm.config.Performance.MaxWorkers = val
		}
	}
	if envVars.MemoryLimit != "" {
		cm.config.Performance.MemoryLimit = envVars.MemoryLimit
	}
	if envVars.CacheEnabled != "" {
		if val, err := strconv.ParseBool(envVars.CacheEnabled); err == nil {
			cm.config.Performance.CacheEnabled = val
		}
	}
	if envVars.CacheDirectory != "" {
		cm.config.Performance.CacheDirectory = envVars.CacheDirectory
	}
	if envVars.CacheTTL != "" {
		if val, err := time.ParseDuration(envVars.CacheTTL); err == nil {
			cm.config.Performance.CacheTTL = val
		}
	}
	if envVars.ParallelEnabled != "" {
		if val, err := strconv.ParseBool(envVars.ParallelEnabled); err == nil {
			cm.config.Performance.ParallelEnabled = val
		}
	}

	// 生成配置
	if envVars.EnableJSON != "" {
		if val, err := strconv.ParseBool(envVars.EnableJSON); err == nil {
			cm.config.Generation.EnableJSON = val
		}
	}
	if envVars.EnableValidation != "" {
		if val, err := strconv.ParseBool(envVars.EnableValidation); err == nil {
			cm.config.Generation.EnableValidation = val
		}
	}
	if envVars.EnableTests != "" {
		if val, err := strconv.ParseBool(envVars.EnableTests); err == nil {
			cm.config.Generation.EnableTests = val
		}
	}
	if envVars.EnableBenchmarks != "" {
		if val, err := strconv.ParseBool(envVars.EnableBenchmarks); err == nil {
			cm.config.Generation.EnableBenchmarks = val
		}
	}

	// 日志配置
	if envVars.LogLevel != "" {
		cm.config.Logging.Level = envVars.LogLevel
	}
	if envVars.LogFile != "" {
		cm.config.Logging.OutputFile = envVars.LogFile
	}
	if envVars.DebugMode != "" {
		if val, err := strconv.ParseBool(envVars.DebugMode); err == nil {
			cm.config.Logging.Level = func() string {
				if val {
					return "debug"
				}
				return "info"
			}()
		}
	}

	if envVars.EnableTracing != "" {
		// EnableTracing字段在LoggingConfig中不存在，可以忽略或添加到其他配置中
		// 这里暂时忽略
	}

	return nil
}

// expandPath 展开路径中的特殊字符
func (cm *ConfigManager) expandPath(path string) string {
	if strings.HasPrefix(path, "~/") {
		homeDir, err := os.UserHomeDir()
		if err == nil {
			return filepath.Join(homeDir, path[2:])
		}
	}
	return path
}

// fileExists 检查文件是否存在
func (cm *ConfigManager) fileExists(path string) bool {
	_, err := os.Stat(path)
	return err == nil
}

// cloneConfig 克隆配置
func (cm *ConfigManager) cloneConfig() *Config {
	// 深拷贝配置
	data, _ := json.Marshal(cm.config)
	var cloned Config
	json.Unmarshal(data, &cloned)
	return &cloned
}

// parseYAML 解析YAML配置（简化版本，实际需要yaml库）
func (cm *ConfigManager) parseYAML(data []byte, config *Config) error {
	// 这里应该使用yaml.Unmarshal，但为了避免外部依赖，先返回错误
	return fmt.Errorf("YAML parsing not implemented yet")
}

// validateConfig 验证配置
func (cm *ConfigManager) validateConfig(config *Config) error {
	// 简单的配置验证
	if config.Performance.MaxWorkers <= 0 {
		return NewValidationError("INVALID_MAX_WORKERS", "MaxWorkers must be greater than 0").Build()
	}

	if config.Performance.CacheTTL <= 0 {
		return NewValidationError("INVALID_CACHE_TTL", "CacheTTL must be greater than 0").Build()
	}

	if config.Generation.OptimizationLevel < 0 || config.Generation.OptimizationLevel > 3 {
		return NewValidationError("INVALID_OPTIMIZATION_LEVEL", "OptimizationLevel must be between 0 and 3").Build()
	}

	return nil
}

// startWatcher 启动配置文件监视器
func (cm *ConfigManager) startWatcher(filePath string) {
	if cm.watcher != nil {
		cm.stopWatcher()
	}

	cm.watcher = &ConfigWatcher{
		filePath:    filePath,
		lastModTime: time.Now(),
		ticker:      time.NewTicker(time.Second * 5),
		stopChan:    make(chan struct{}),
		manager:     cm,
	}

	go cm.watcher.watch()
}

// stopWatcher 停止配置文件监视器
func (cm *ConfigManager) stopWatcher() {
	if cm.watcher != nil && cm.watcher.running.Load() {
		close(cm.watcher.stopChan)
		cm.watcher.ticker.Stop()
	}
}

// notifyConfigChange 通知配置变更
func (cm *ConfigManager) notifyConfigChange(oldConfig, newConfig *Config) {
	for _, callback := range cm.callbacks {
		if err := callback(oldConfig, newConfig); err != nil {
			// 记录错误到错误管理器
			cm.errorManager.AddError(NewConfigError("CONFIG_CALLBACK_ERROR",
				fmt.Sprintf("Configuration change callback failed: %v", err)).Build())
		}
	}
}

// watch 监视配置文件变更
func (cw *ConfigWatcher) watch() {
	cw.running.Store(true)
	defer cw.running.Store(false)

	for {
		select {
		case <-cw.ticker.C:
			if info, err := os.Stat(cw.filePath); err == nil {
				if info.ModTime().After(cw.lastModTime) {
					cw.lastModTime = info.ModTime()
					// 重新加载配置
					if err := cw.manager.LoadFromFile(cw.filePath); err != nil {
						cw.manager.errorManager.AddError(NewConfigError("CONFIG_RELOAD_ERROR",
							fmt.Sprintf("Failed to reload config: %v", err)).Build())
					}
				}
			}
		case <-cw.stopChan:
			return
		}
	}
}

// RegisterChangeCallback 注册配置变更回调
func (cm *ConfigManager) RegisterChangeCallback(callback ConfigChangeCallback) {
	cm.mutex.Lock()
	defer cm.mutex.Unlock()
	cm.callbacks = append(cm.callbacks, callback)
}

// SaveToFile 保存配置到文件
func (cm *ConfigManager) SaveToFile(filePath string) error {
	cm.mutex.RLock()
	defer cm.mutex.RUnlock()

	data, err := json.MarshalIndent(cm.config, "", "  ")
	if err != nil {
		return NewConfigError("CONFIG_MARSHAL_ERROR", fmt.Sprintf("Failed to marshal config: %v", err)).Build()
	}

	if err := os.WriteFile(filePath, data, 0644); err != nil {
		return NewIOError("CONFIG_WRITE_ERROR", fmt.Sprintf("Failed to write config file: %s", filePath)).Build()
	}

	return nil
}

// ValidateConfig 验证配置
func (cm *ConfigManager) ValidateConfig() error {
	var validationErrors []string

	// 验证输入配置
	if cm.config.Input.XSDPath == "" {
		validationErrors = append(validationErrors, "XSD path is required")
	}

	if cm.config.Input.MaxFileSize <= 0 {
		validationErrors = append(validationErrors, "Max file size must be positive")
	}

	// 验证输出配置
	if cm.config.Output.PackageName == "" {
		validationErrors = append(validationErrors, "Package name is required")
	}

	supportedLanguages := []string{"go", "java", "csharp", "python"}
	isValidLanguage := false
	for _, lang := range supportedLanguages {
		if cm.config.Output.TargetLanguage == lang {
			isValidLanguage = true
			break
		}
	}
	if !isValidLanguage {
		validationErrors = append(validationErrors,
			fmt.Sprintf("Unsupported target language: %s (supported: %v)",
				cm.config.Output.TargetLanguage, supportedLanguages))
	}

	// 验证性能配置
	if cm.config.Performance.MaxWorkers <= 0 {
		validationErrors = append(validationErrors, "Max workers must be positive")
	}

	if cm.config.Performance.MemoryPoolSize < 0 {
		validationErrors = append(validationErrors, "Memory pool size cannot be negative")
	}

	// 验证日志配置
	supportedLogLevels := []string{"TRACE", "DEBUG", "INFO", "WARN", "ERROR", "FATAL"}
	isValidLogLevel := false
	for _, level := range supportedLogLevels {
		if strings.EqualFold(cm.config.Logging.Level, level) {
			isValidLogLevel = true
			break
		}
	}
	if !isValidLogLevel {
		validationErrors = append(validationErrors,
			fmt.Sprintf("Unsupported log level: %s (supported: %v)",
				cm.config.Logging.Level, supportedLogLevels))
	}
	if len(validationErrors) > 0 {
		cm.errorManager.AddError(NewConfigError("VALIDATION_FAILED",
			fmt.Sprintf("Configuration validation failed: %s",
				strings.Join(validationErrors, "; "))).Build())
		return fmt.Errorf("configuration validation failed: %s",
			strings.Join(validationErrors, "; "))
	}

	return nil
}

// GetConfig 获取当前配置
func (cm *ConfigManager) GetConfig() *Config {
	return cm.config
}

// UpdateConfig 更新配置
func (cm *ConfigManager) UpdateConfig(updateFn func(*Config)) {
	updateFn(cm.config)
}

// MergeConfig 合并配置
func (cm *ConfigManager) MergeConfig(other *Config) {
	// 这里可以实现更复杂的配置合并逻辑
	// 目前简单地用其他配置覆盖当前配置的非零值
	if other.Input.XSDPath != "" {
		cm.config.Input.XSDPath = other.Input.XSDPath
	}
	if other.Output.OutputPath != "" {
		cm.config.Output.OutputPath = other.Output.OutputPath
	}
	if other.Output.PackageName != "" {
		cm.config.Output.PackageName = other.Output.PackageName
	}
	// ... 其他字段的合并逻辑
}

// GetConfigSummary 获取配置摘要
func (cm *ConfigManager) GetConfigSummary() *ConfigSummary {
	cm.mutex.RLock()
	defer cm.mutex.RUnlock()

	return &ConfigSummary{
		XSDPath:          cm.config.Input.XSDPath,
		OutputPath:       cm.config.Output.OutputPath,
		TargetLanguage:   cm.config.Output.TargetLanguage,
		PackageName:      cm.config.Output.PackageName,
		MaxWorkers:       cm.config.Performance.MaxWorkers,
		CacheEnabled:     cm.config.Performance.CacheEnabled,
		EnableJSON:       cm.config.Generation.EnableJSON,
		EnableValidation: cm.config.Generation.EnableValidation,
		EnableTests:      cm.config.Generation.EnableTests,
		DebugMode:        cm.config.Features.DebugMode,
		LogLevel:         cm.config.Logging.Level,
	}
}

// ConfigSummary 配置摘要
type ConfigSummary struct {
	XSDPath          string `json:"xsd_path"`
	OutputPath       string `json:"output_path"`
	TargetLanguage   string `json:"target_language"`
	PackageName      string `json:"package_name"`
	MaxWorkers       int    `json:"max_workers"`
	CacheEnabled     bool   `json:"cache_enabled"`
	EnableJSON       bool   `json:"enable_json"`
	EnableValidation bool   `json:"enable_validation"`
	EnableTests      bool   `json:"enable_tests"`
	DebugMode        bool   `json:"debug_mode"`
	LogLevel         string `json:"log_level"`
}

// getEnabledFeatures 获取启用的功能列表
func (cm *ConfigManager) getEnabledFeatures() []string {
	var features []string

	if cm.config.Generation.EnableJSON {
		features = append(features, "JSON")
	}
	if cm.config.Generation.EnableValidation {
		features = append(features, "Validation")
	}
	if cm.config.Generation.EnableTests {
		features = append(features, "Tests")
	}
	if cm.config.Generation.EnableBenchmarks {
		features = append(features, "Benchmarks")
	}
	if cm.config.Performance.ParallelEnabled {
		features = append(features, "Parallel Processing")
	}
	if cm.config.Performance.CacheEnabled {
		features = append(features, "Caching")
	}
	if cm.config.Performance.StreamingMode {
		features = append(features, "Streaming")
	}
	if cm.config.Features.ShowTypeMappings {
		features = append(features, "Type Mappings")
	}
	if cm.config.Features.CreateSampleXML {
		features = append(features, "Sample XML")
	}

	return features
}
