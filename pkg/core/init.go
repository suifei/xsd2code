package core

import (
	"fmt"
	"time"
)

// GlobalManagers 全局管理器集合
type GlobalManagers struct {
	Performance *PerformanceMetrics
	Error       *ErrorManager
	Config      *ConfigManager
	Cache       *CacheManager
	Log         *LogManager
	Resource    *ResourceManager
	Concurrent  *ConcurrentProcessor
}

var (
	// GlobalManagers 全局管理器实例
	Managers *GlobalManagers
)

// InitializeManagers 初始化所有核心管理器
func InitializeManagers(configFile string) error {
	var err error

	// 初始化错误管理器（先初始化，其他管理器依赖它）
	errorManager := NewErrorManager(1000) // 最大错误数

	// 初始化配置管理器
	configManager := NewConfigManager(errorManager)
	if configFile != "" {
		if err = configManager.LoadFromFile(configFile); err != nil {
			return fmt.Errorf("failed to load config: %w", err)
		}
	}
	// 初始化日志管理器
	logConfig := &LoggingConfig{
		Level:         "info",
		OutputFile:    "logs/xsd2code.log",
		EnableConsole: true,
		EnableFile:    true,
		MaxFileSize:   "10MB",
		MaxBackups:    5,
	}
	logManager, err := NewLogManager(logConfig, errorManager)
	if err != nil {
		return fmt.Errorf("failed to create log manager: %w", err)
	}

	// 初始化性能指标收集器
	performanceMetrics := NewPerformanceMetrics()

	// 初始化资源管理器
	resourceManager := NewResourceManager(10000, 1000) // 池大小和最大对象数
	// 初始化缓存管理器
	cacheManager := NewCacheManager(
		1000,                     // 最大容量
		time.Hour,                // TTL
		"./cache",                // 缓存目录
		true,                     // 启用持久化
		errorManager,             // 错误管理器
		DefaultCleanupStrategy(), // 清理策略
	)

	// 初始化并发处理器
	config := &Config{
		Performance: PerformanceConfig{
			MaxWorkers:      4,
			MemoryPoolSize:  1000,
			CacheEnabled:    true,
			ParallelEnabled: true,
		},
	}
	concurrentProcessor := NewConcurrentProcessor(config, cacheManager, errorManager)

	// 创建全局管理器实例
	Managers = &GlobalManagers{
		Performance: performanceMetrics,
		Error:       errorManager,
		Config:      configManager,
		Cache:       cacheManager,
		Log:         logManager,
		Resource:    resourceManager,
		Concurrent:  concurrentProcessor,
	}

	return nil
}

// ShutdownManagers 优雅关闭所有管理器
func ShutdownManagers() error {
	if Managers == nil {
		return nil
	}

	var errors []error

	// 关闭并发处理器
	if Managers.Concurrent != nil {
		if err := Managers.Concurrent.Stop(); err != nil {
			errors = append(errors, fmt.Errorf("failed to shutdown concurrent processor: %w", err))
		}
	}
	// 保存缓存（简化版本，不返回错误）
	if Managers.Cache != nil {
		// 缓存会在程序退出时自动保存
		fmt.Println("缓存管理器已关闭")
	}

	// 关闭日志管理器
	if Managers.Log != nil {
		Managers.Log.Close()
	}

	// 输出错误信息
	if len(errors) > 0 {
		fmt.Printf("Shutdown completed with %d errors\n", len(errors))
		for _, err := range errors {
			fmt.Printf("- %v\n", err)
		}
	}

	Managers = nil
	return nil
}

// GetPerformanceManager 获取性能管理器
func GetPerformanceManager() *PerformanceMetrics {
	if Managers == nil {
		return nil
	}
	return Managers.Performance
}

// GetErrorManager 获取错误管理器
func GetErrorManager() *ErrorManager {
	if Managers == nil {
		return nil
	}
	return Managers.Error
}

// GetConfigManager 获取配置管理器
func GetConfigManager() *ConfigManager {
	if Managers == nil {
		return nil
	}
	return Managers.Config
}

// GetCacheManager 获取缓存管理器
func GetCacheManager() *CacheManager {
	if Managers == nil {
		return nil
	}
	return Managers.Cache
}

// GetLogManager 获取日志管理器
func GetLogManager() *LogManager {
	if Managers == nil {
		return nil
	}
	return Managers.Log
}

// GetResourceManager 获取资源管理器
func GetResourceManager() *ResourceManager {
	if Managers == nil {
		return nil
	}
	return Managers.Resource
}

// GetConcurrentProcessor 获取并发处理器
func GetConcurrentProcessor() *ConcurrentProcessor {
	if Managers == nil {
		return nil
	}
	return Managers.Concurrent
}

// IsInitialized 检查管理器是否已初始化
func IsInitialized() bool {
	return Managers != nil
}

// GetVersion 获取版本信息
func GetVersion() string {
	return "XSD2Code v2.0.0 - Second Iteration with Core Optimization"
}
