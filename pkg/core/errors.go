package core

import (
	"errors"
	"fmt"
	"runtime"
	"strings"
	"sync"
	"time"
)

// ErrorType 错误类型枚举
type ErrorType int

const (
	ErrorTypeUnknown ErrorType = iota
	ErrorTypeParse
	ErrorTypeGeneration
	ErrorTypeValidation
	ErrorTypeConfig
	ErrorTypeIO
	ErrorTypeMemory
	ErrorTypeTimeout
)

// String 返回错误类型的字符串表示
func (et ErrorType) String() string {
	switch et {
	case ErrorTypeParse:
		return "PARSE"
	case ErrorTypeGeneration:
		return "GENERATION"
	case ErrorTypeValidation:
		return "VALIDATION"
	case ErrorTypeConfig:
		return "CONFIG"
	case ErrorTypeIO:
		return "IO"
	case ErrorTypeMemory:
		return "MEMORY"
	case ErrorTypeTimeout:
		return "TIMEOUT"
	default:
		return "UNKNOWN"
	}
}

// XSDError 统一错误结构
type XSDError struct {
	Type      ErrorType `json:"type"`
	Code      string    `json:"code"`
	Message   string    `json:"message"`
	Context   string    `json:"context,omitempty"`
	File      string    `json:"file,omitempty"`
	Line      int       `json:"line,omitempty"`
	Column    int       `json:"column,omitempty"`
	Timestamp time.Time `json:"timestamp"`
	Stack     string    `json:"stack,omitempty"`
	Wrapped   error     `json:"-"`
}

// Error 实现error接口
func (e *XSDError) Error() string {
	var parts []string

	if e.Code != "" {
		parts = append(parts, fmt.Sprintf("[%s]", e.Code))
	}

	if e.Type != ErrorTypeUnknown {
		parts = append(parts, fmt.Sprintf("(%s)", e.Type.String()))
	}

	parts = append(parts, e.Message)

	if e.File != "" {
		location := e.File
		if e.Line > 0 {
			location += fmt.Sprintf(":%d", e.Line)
			if e.Column > 0 {
				location += fmt.Sprintf(":%d", e.Column)
			}
		}
		parts = append(parts, fmt.Sprintf("at %s", location))
	}

	if e.Context != "" {
		parts = append(parts, fmt.Sprintf("context: %s", e.Context))
	}

	return strings.Join(parts, " ")
}

// Unwrap 支持errors.Unwrap
func (e *XSDError) Unwrap() error {
	return e.Wrapped
}

// Is 支持errors.Is
func (e *XSDError) Is(target error) bool {
	if target == nil {
		return false
	}

	if xerr, ok := target.(*XSDError); ok {
		return e.Type == xerr.Type && e.Code == xerr.Code
	}

	return false
}

// ErrorManager 错误管理器
type ErrorManager struct {
	errors   []XSDError
	warnings []XSDError
	mutex    sync.RWMutex
	maxSize  int
}

// NewErrorManager 创建错误管理器
func NewErrorManager(maxSize int) *ErrorManager {
	return &ErrorManager{
		errors:   make([]XSDError, 0),
		warnings: make([]XSDError, 0),
		maxSize:  maxSize,
	}
}

// AddError 添加错误
func (em *ErrorManager) AddError(err *XSDError) {
	em.mutex.Lock()
	defer em.mutex.Unlock()

	if len(em.errors) >= em.maxSize {
		// 移除最老的错误
		em.errors = em.errors[1:]
	}

	em.errors = append(em.errors, *err)
}

// AddWarning 添加警告
func (em *ErrorManager) AddWarning(err *XSDError) {
	em.mutex.Lock()
	defer em.mutex.Unlock()

	if len(em.warnings) >= em.maxSize {
		// 移除最老的警告
		em.warnings = em.warnings[1:]
	}

	em.warnings = append(em.warnings, *err)
}

// HasErrors 检查是否有错误
func (em *ErrorManager) HasErrors() bool {
	em.mutex.RLock()
	defer em.mutex.RUnlock()
	return len(em.errors) > 0
}

// HasWarnings 检查是否有警告
func (em *ErrorManager) HasWarnings() bool {
	em.mutex.RLock()
	defer em.mutex.RUnlock()
	return len(em.warnings) > 0
}

// GetErrors 获取所有错误
func (em *ErrorManager) GetErrors() []XSDError {
	em.mutex.RLock()
	defer em.mutex.RUnlock()

	result := make([]XSDError, len(em.errors))
	copy(result, em.errors)
	return result
}

// GetWarnings 获取所有警告
func (em *ErrorManager) GetWarnings() []XSDError {
	em.mutex.RLock()
	defer em.mutex.RUnlock()

	result := make([]XSDError, len(em.warnings))
	copy(result, em.warnings)
	return result
}

// Clear 清空所有错误和警告
func (em *ErrorManager) Clear() {
	em.mutex.Lock()
	defer em.mutex.Unlock()

	em.errors = em.errors[:0]
	em.warnings = em.warnings[:0]
}

// GetSummary 获取错误摘要
func (em *ErrorManager) GetSummary() ErrorSummary {
	em.mutex.RLock()
	defer em.mutex.RUnlock()

	summary := ErrorSummary{
		TotalErrors:   len(em.errors),
		TotalWarnings: len(em.warnings),
		ErrorsByType:  make(map[ErrorType]int),
	}

	for _, err := range em.errors {
		summary.ErrorsByType[err.Type]++
	}

	return summary
}

// ErrorSummary 错误摘要
type ErrorSummary struct {
	TotalErrors   int               `json:"total_errors"`
	TotalWarnings int               `json:"total_warnings"`
	ErrorsByType  map[ErrorType]int `json:"errors_by_type"`
}

// ErrorBuilder 错误构建器
type ErrorBuilder struct {
	error *XSDError
}

// NewError 创建新的错误构建器
func NewError(errorType ErrorType, code string, message string) *ErrorBuilder {
	return &ErrorBuilder{
		error: &XSDError{
			Type:      errorType,
			Code:      code,
			Message:   message,
			Timestamp: time.Now(),
		},
	}
}

// WithContext 添加上下文信息
func (eb *ErrorBuilder) WithContext(context string) *ErrorBuilder {
	eb.error.Context = context
	return eb
}

// WithLocation 添加位置信息
func (eb *ErrorBuilder) WithLocation(file string, line, column int) *ErrorBuilder {
	eb.error.File = file
	eb.error.Line = line
	eb.error.Column = column
	return eb
}

// WithStack 添加堆栈信息
func (eb *ErrorBuilder) WithStack() *ErrorBuilder {
	// 获取调用堆栈
	buf := make([]byte, 4096)
	n := runtime.Stack(buf, false)
	eb.error.Stack = string(buf[:n])
	return eb
}

// Wrap 包装已有错误
func (eb *ErrorBuilder) Wrap(err error) *ErrorBuilder {
	eb.error.Wrapped = err
	return eb
}

// Build 构建错误
func (eb *ErrorBuilder) Build() *XSDError {
	return eb.error
}

// 预定义错误构建器工厂函数

// NewParseError 创建解析错误
func NewParseError(code, message string) *ErrorBuilder {
	return NewError(ErrorTypeParse, code, message)
}

// NewGenerationError 创建生成错误
func NewGenerationError(code, message string) *ErrorBuilder {
	return NewError(ErrorTypeGeneration, code, message)
}

// NewValidationError 创建验证错误
func NewValidationError(code, message string) *ErrorBuilder {
	return NewError(ErrorTypeValidation, code, message)
}

// NewConfigError 创建配置错误
func NewConfigError(code, message string) *ErrorBuilder {
	return NewError(ErrorTypeConfig, code, message)
}

// NewIOError 创建IO错误
func NewIOError(code, message string) *ErrorBuilder {
	return NewError(ErrorTypeIO, code, message)
}

// NewMemoryError 创建内存错误
func NewMemoryError(code, message string) *ErrorBuilder {
	return NewError(ErrorTypeMemory, code, message)
}

// NewTimeoutError 创建超时错误
func NewTimeoutError(code, message string) *ErrorBuilder {
	return NewError(ErrorTypeTimeout, code, message)
}

// 常用错误类型检查函数

// IsParseError 检查是否为解析错误
func IsParseError(err error) bool {
	var xErr *XSDError
	return errors.As(err, &xErr) && xErr.Type == ErrorTypeParse
}

// IsGenerationError 检查是否为生成错误
func IsGenerationError(err error) bool {
	var xErr *XSDError
	return errors.As(err, &xErr) && xErr.Type == ErrorTypeGeneration
}

// IsValidationError 检查是否为验证错误
func IsValidationError(err error) bool {
	var xErr *XSDError
	return errors.As(err, &xErr) && xErr.Type == ErrorTypeValidation
}

// IsConfigError 检查是否为配置错误
func IsConfigError(err error) bool {
	var xErr *XSDError
	return errors.As(err, &xErr) && xErr.Type == ErrorTypeConfig
}

// IsIOError 检查是否为IO错误
func IsIOError(err error) bool {
	var xErr *XSDError
	return errors.As(err, &xErr) && xErr.Type == ErrorTypeIO
}

// IsMemoryError 检查是否为内存错误
func IsMemoryError(err error) bool {
	var xErr *XSDError
	return errors.As(err, &xErr) && xErr.Type == ErrorTypeMemory
}

// IsTimeoutError 检查是否为超时错误
func IsTimeoutError(err error) bool {
	var xErr *XSDError
	return errors.As(err, &xErr) && xErr.Type == ErrorTypeTimeout
}

// RecoveryHandler 恢复处理器
type RecoveryHandler struct {
	errorManager *ErrorManager
}

// NewRecoveryHandler 创建恢复处理器
func NewRecoveryHandler(errorManager *ErrorManager) *RecoveryHandler {
	return &RecoveryHandler{
		errorManager: errorManager,
	}
}

// HandlePanic 处理panic恢复
func (rh *RecoveryHandler) HandlePanic() {
	if r := recover(); r != nil {
		var err *XSDError

		switch v := r.(type) {
		case error:
			err = NewError(ErrorTypeUnknown, "PANIC", v.Error()).
				WithStack().
				Build()
		case string:
			err = NewError(ErrorTypeUnknown, "PANIC", v).
				WithStack().
				Build()
		default:
			err = NewError(ErrorTypeUnknown, "PANIC", fmt.Sprintf("%v", v)).
				WithStack().
				Build()
		}

		rh.errorManager.AddError(err)
	}
}

// SafeExecute 安全执行函数（带panic恢复）
func (rh *RecoveryHandler) SafeExecute(fn func() error) error {
	defer rh.HandlePanic()
	return fn()
}
