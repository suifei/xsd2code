package core

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"sync"
	"time"
)

// ErrorSeverity 错误严重程度
type ErrorSeverity int

const (
	SeverityInfo ErrorSeverity = iota
	SeverityWarning
	SeverityError
	SeverityCritical
	SeverityFatal
)

// String 返回严重程度的字符串表示
func (s ErrorSeverity) String() string {
	switch s {
	case SeverityInfo:
		return "INFO"
	case SeverityWarning:
		return "WARNING"
	case SeverityError:
		return "ERROR"
	case SeverityCritical:
		return "CRITICAL"
	case SeverityFatal:
		return "FATAL"
	default:
		return "UNKNOWN"
	}
}

// RecoveryStrategy 恢复策略
type RecoveryStrategy int

const (
	RecoveryNone RecoveryStrategy = iota
	RecoveryRetry
	RecoverySkip
	RecoveryFallback
	RecoveryTerminate
)

// String 返回恢复策略的字符串表示
func (rs RecoveryStrategy) String() string {
	switch rs {
	case RecoveryNone:
		return "NONE"
	case RecoveryRetry:
		return "RETRY"
	case RecoverySkip:
		return "SKIP"
	case RecoveryFallback:
		return "FALLBACK"
	case RecoveryTerminate:
		return "TERMINATE"
	default:
		return "UNKNOWN"
	}
}

// ErrorContext 错误上下文信息
type ErrorContext struct {
	Operation     string                 `json:"operation"`
	Input         interface{}            `json:"input,omitempty"`
	ExpectedType  string                 `json:"expected_type,omitempty"`
	ActualType    string                 `json:"actual_type,omitempty"`
	StackTrace    []string               `json:"stack_trace,omitempty"`
	Environment   map[string]string      `json:"environment,omitempty"`
	UserData      map[string]interface{} `json:"user_data,omitempty"`
	ProcessingID  string                 `json:"processing_id,omitempty"`
	ParentError   string                 `json:"parent_error,omitempty"`
	RelatedErrors []string               `json:"related_errors,omitempty"`
}

// EnhancedXSDError 增强的XSD错误
type EnhancedXSDError struct {
	*XSDError
	Severity        ErrorSeverity    `json:"severity"`
	Context         ErrorContext     `json:"context"`
	RecoveryAction  RecoveryStrategy `json:"recovery_action"`
	RetryCount      int              `json:"retry_count"`
	MaxRetries      int              `json:"max_retries"`
	RecoveryOptions []string         `json:"recovery_options,omitempty"`
	Metrics         ErrorMetrics     `json:"metrics"`
}

// ErrorMetrics 错误指标
type ErrorMetrics struct {
	FirstOccurrence time.Time     `json:"first_occurrence"`
	LastOccurrence  time.Time     `json:"last_occurrence"`
	OccurrenceCount int64         `json:"occurrence_count"`
	ResolutionTime  time.Duration `json:"resolution_time,omitempty"`
	AverageImpact   float64       `json:"average_impact"`
}

// Error 实现error接口
func (e *EnhancedXSDError) Error() string {
	base := e.XSDError.Error()
	return fmt.Sprintf("[%s] %s (Recovery: %s, Retries: %d/%d)",
		e.Severity.String(), base, e.RecoveryAction.String(), e.RetryCount, e.MaxRetries)
}

// CanRetry 检查是否可以重试
func (e *EnhancedXSDError) CanRetry() bool {
	return e.RecoveryAction == RecoveryRetry && e.RetryCount < e.MaxRetries
}

// IncrementRetry 增加重试次数
func (e *EnhancedXSDError) IncrementRetry() {
	e.RetryCount++
	e.Metrics.LastOccurrence = time.Now()
	e.Metrics.OccurrenceCount++
}

// ErrorRecoveryHandler 错误恢复处理器
type ErrorRecoveryHandler struct {
	strategies map[string]RecoveryStrategy
	handlers   map[RecoveryStrategy]func(error) error
	mutex      sync.RWMutex
}

// NewErrorRecoveryHandler 创建错误恢复处理器
func NewErrorRecoveryHandler() *ErrorRecoveryHandler {
	handler := &ErrorRecoveryHandler{
		strategies: make(map[string]RecoveryStrategy),
		handlers:   make(map[RecoveryStrategy]func(error) error),
	}

	// 注册默认恢复处理器
	handler.RegisterHandler(RecoveryRetry, handler.defaultRetryHandler)
	handler.RegisterHandler(RecoverySkip, handler.defaultSkipHandler)
	handler.RegisterHandler(RecoveryFallback, handler.defaultFallbackHandler)

	return handler
}

// RegisterStrategy 注册错误代码的恢复策略
func (erh *ErrorRecoveryHandler) RegisterStrategy(errorCode string, strategy RecoveryStrategy) {
	erh.mutex.Lock()
	defer erh.mutex.Unlock()
	erh.strategies[errorCode] = strategy
}

// RegisterHandler 注册恢复策略处理器
func (erh *ErrorRecoveryHandler) RegisterHandler(strategy RecoveryStrategy, handler func(error) error) {
	erh.mutex.Lock()
	defer erh.mutex.Unlock()
	erh.handlers[strategy] = handler
}

// GetStrategy 获取错误代码的恢复策略
func (erh *ErrorRecoveryHandler) GetStrategy(errorCode string) RecoveryStrategy {
	erh.mutex.RLock()
	defer erh.mutex.RUnlock()

	if strategy, exists := erh.strategies[errorCode]; exists {
		return strategy
	}
	return RecoveryNone
}

// HandleError 处理错误
func (erh *ErrorRecoveryHandler) HandleError(err error) error {
	if enhanced, ok := err.(*EnhancedXSDError); ok {
		erh.mutex.RLock()
		handler, exists := erh.handlers[enhanced.RecoveryAction]
		erh.mutex.RUnlock()

		if exists {
			return handler(err)
		}
	}
	return err
}

// 默认恢复处理器
func (erh *ErrorRecoveryHandler) defaultRetryHandler(err error) error {
	if enhanced, ok := err.(*EnhancedXSDError); ok {
		if enhanced.CanRetry() {
			enhanced.IncrementRetry()
			time.Sleep(time.Duration(enhanced.RetryCount) * time.Second) // 指数退避
			return nil                                                   // 表示可以重试
		}
	}
	return err
}

func (erh *ErrorRecoveryHandler) defaultSkipHandler(err error) error {
	// 跳过错误，记录日志但继续处理
	fmt.Printf("Skipping error: %v\n", err)
	return nil
}

func (erh *ErrorRecoveryHandler) defaultFallbackHandler(err error) error {
	// 使用备用方案处理
	fmt.Printf("Using fallback for error: %v\n", err)
	return nil
}

// EnhancedErrorManager 增强错误管理器
type EnhancedErrorManager struct {
	*ErrorManager
	recoveryHandler *ErrorRecoveryHandler
	errorLog        []EnhancedXSDError
	metrics         ErrorManagerMetrics
	logFile         string
	mutex           sync.RWMutex
}

// ErrorManagerMetrics 错误管理器指标
type ErrorManagerMetrics struct {
	TotalErrors           int64                      `json:"total_errors"`
	ErrorsByType          map[ErrorType]int64        `json:"errors_by_type"`
	ErrorsBySeverity      map[ErrorSeverity]int64    `json:"errors_by_severity"`
	RecoverySuccess       map[RecoveryStrategy]int64 `json:"recovery_success"`
	RecoveryFailure       map[RecoveryStrategy]int64 `json:"recovery_failure"`
	AverageResolutionTime time.Duration              `json:"average_resolution_time"`
}

// NewEnhancedErrorManager 创建增强错误管理器
func NewEnhancedErrorManager(logFile string) *EnhancedErrorManager {
	base := NewErrorManager(1000) // Default max size of 1000 errors

	return &EnhancedErrorManager{
		ErrorManager:    base,
		recoveryHandler: NewErrorRecoveryHandler(),
		errorLog:        make([]EnhancedXSDError, 0),
		logFile:         logFile,
		metrics: ErrorManagerMetrics{
			ErrorsByType:     make(map[ErrorType]int64),
			ErrorsBySeverity: make(map[ErrorSeverity]int64),
			RecoverySuccess:  make(map[RecoveryStrategy]int64),
			RecoveryFailure:  make(map[RecoveryStrategy]int64),
		},
	}
}

// AddEnhancedError 添加增强错误
func (em *EnhancedErrorManager) AddEnhancedError(err *EnhancedXSDError) {
	em.mutex.Lock()
	defer em.mutex.Unlock()

	// 更新指标
	em.metrics.TotalErrors++
	em.metrics.ErrorsByType[err.Type]++
	em.metrics.ErrorsBySeverity[err.Severity]++

	// 设置首次发生时间
	if err.Metrics.FirstOccurrence.IsZero() {
		err.Metrics.FirstOccurrence = time.Now()
	}
	err.Metrics.LastOccurrence = time.Now()
	err.Metrics.OccurrenceCount++

	// 添加到错误日志
	em.errorLog = append(em.errorLog, *err)

	// 添加到基础错误管理器
	em.ErrorManager.AddError(err.XSDError)

	// 记录到文件
	em.logToFile(err)
}

// CreateEnhancedError 创建增强错误
func (em *EnhancedErrorManager) CreateEnhancedError(errorType ErrorType, code, message string) *EnhancedXSDErrorBuilder {
	return &EnhancedXSDErrorBuilder{
		err: &EnhancedXSDError{
			XSDError: &XSDError{
				Type:      errorType,
				Code:      code,
				Message:   message,
				Timestamp: time.Now(),
			},
			Severity:       SeverityError,
			RecoveryAction: em.recoveryHandler.GetStrategy(code),
			MaxRetries:     3,
			Context: ErrorContext{
				Environment: make(map[string]string),
				UserData:    make(map[string]interface{}),
			},
		},
		manager: em,
	}
}

// ProcessWithRecovery 带恢复机制处理函数
func (em *EnhancedErrorManager) ProcessWithRecovery(ctx context.Context, operation string, fn func() error) error {
	var lastError error
	maxAttempts := 3

	for attempt := 1; attempt <= maxAttempts; attempt++ {
		err := fn()
		if err == nil {
			return nil
		}

		lastError = err

		// 如果是增强错误，尝试恢复
		if enhanced, ok := err.(*EnhancedXSDError); ok {
			enhanced.Context.Operation = operation
			enhanced.Context.ProcessingID = fmt.Sprintf("%s_%d", operation, time.Now().UnixNano())

			// 尝试恢复
			recoveryErr := em.recoveryHandler.HandleError(enhanced)
			if recoveryErr == nil && enhanced.CanRetry() {
				em.metrics.RecoverySuccess[enhanced.RecoveryAction]++
				continue // 重试
			}

			em.metrics.RecoveryFailure[enhanced.RecoveryAction]++
		}

		// 检查上下文是否被取消
		select {
		case <-ctx.Done():
			return ctx.Err()
		default:
		}

		// 如果不能恢复，退出循环
		break
	}

	return lastError
}

// logToFile 记录错误到文件
func (em *EnhancedErrorManager) logToFile(err *EnhancedXSDError) {
	if em.logFile == "" {
		return
	}

	// 确保日志目录存在
	logDir := filepath.Dir(em.logFile)
	os.MkdirAll(logDir, 0755)

	// 打开日志文件
	file, fileErr := os.OpenFile(em.logFile, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if fileErr != nil {
		return
	}
	defer file.Close()

	// 格式化错误信息
	logEntry := map[string]interface{}{
		"timestamp": err.Timestamp.Format(time.RFC3339),
		"severity":  err.Severity.String(),
		"type":      err.Type.String(),
		"code":      err.Code,
		"message":   err.Message,
		"context":   err.Context,
		"recovery":  err.RecoveryAction.String(),
		"retries":   err.RetryCount,
		"metrics":   err.Metrics,
	}

	// 写入JSON格式
	if data, jsonErr := json.Marshal(logEntry); jsonErr == nil {
		file.Write(data)
		file.WriteString("\n")
	}
}

// GetErrorReport 获取错误报告
func (em *EnhancedErrorManager) GetErrorReport() ErrorReport {
	em.mutex.RLock()
	defer em.mutex.RUnlock()

	return ErrorReport{
		TotalErrors:      em.metrics.TotalErrors,
		ErrorsByType:     em.metrics.ErrorsByType,
		ErrorsBySeverity: em.metrics.ErrorsBySeverity,
		RecoveryMetrics: RecoveryMetrics{
			SuccessfulRecoveries: em.metrics.RecoverySuccess,
			FailedRecoveries:     em.metrics.RecoveryFailure,
		},
		RecentErrors: em.getRecentErrors(10),
		GeneratedAt:  time.Now(),
	}
}

// getRecentErrors 获取最近的错误
func (em *EnhancedErrorManager) getRecentErrors(limit int) []EnhancedXSDError {
	if len(em.errorLog) <= limit {
		return em.errorLog
	}
	return em.errorLog[len(em.errorLog)-limit:]
}

// ErrorReport 错误报告
type ErrorReport struct {
	TotalErrors      int64                   `json:"total_errors"`
	ErrorsByType     map[ErrorType]int64     `json:"errors_by_type"`
	ErrorsBySeverity map[ErrorSeverity]int64 `json:"errors_by_severity"`
	RecoveryMetrics  RecoveryMetrics         `json:"recovery_metrics"`
	RecentErrors     []EnhancedXSDError      `json:"recent_errors"`
	GeneratedAt      time.Time               `json:"generated_at"`
}

// RecoveryMetrics 恢复指标
type RecoveryMetrics struct {
	SuccessfulRecoveries map[RecoveryStrategy]int64 `json:"successful_recoveries"`
	FailedRecoveries     map[RecoveryStrategy]int64 `json:"failed_recoveries"`
}

// EnhancedXSDErrorBuilder 增强错误构建器
type EnhancedXSDErrorBuilder struct {
	err     *EnhancedXSDError
	manager *EnhancedErrorManager
}

// WithSeverity 设置严重程度
func (b *EnhancedXSDErrorBuilder) WithSeverity(severity ErrorSeverity) *EnhancedXSDErrorBuilder {
	b.err.Severity = severity
	return b
}

// WithContext 设置上下文
func (b *EnhancedXSDErrorBuilder) WithContext(key string, value interface{}) *EnhancedXSDErrorBuilder {
	if b.err.Context.UserData == nil {
		b.err.Context.UserData = make(map[string]interface{})
	}
	b.err.Context.UserData[key] = value
	return b
}

// WithOperation 设置操作
func (b *EnhancedXSDErrorBuilder) WithOperation(operation string) *EnhancedXSDErrorBuilder {
	b.err.Context.Operation = operation
	return b
}

// WithInput 设置输入
func (b *EnhancedXSDErrorBuilder) WithInput(input interface{}) *EnhancedXSDErrorBuilder {
	b.err.Context.Input = input
	return b
}

// WithRecovery 设置恢复策略
func (b *EnhancedXSDErrorBuilder) WithRecovery(strategy RecoveryStrategy, maxRetries int) *EnhancedXSDErrorBuilder {
	b.err.RecoveryAction = strategy
	b.err.MaxRetries = maxRetries
	return b
}

// WithStackTrace 添加堆栈跟踪
func (b *EnhancedXSDErrorBuilder) WithStackTrace() *EnhancedXSDErrorBuilder {
	stack := make([]string, 0)
	for i := 1; i < 10; i++ {
		pc, file, line, ok := runtime.Caller(i)
		if !ok {
			break
		}
		fn := runtime.FuncForPC(pc)
		stack = append(stack, fmt.Sprintf("%s:%d %s", file, line, fn.Name()))
	}
	b.err.Context.StackTrace = stack
	return b
}

// Build 构建并添加错误
func (b *EnhancedXSDErrorBuilder) Build() *EnhancedXSDError {
	b.manager.AddEnhancedError(b.err)
	return b.err
}

// BuildWithoutAdd 仅构建错误，不添加到管理器
func (b *EnhancedXSDErrorBuilder) BuildWithoutAdd() *EnhancedXSDError {
	return b.err
}
