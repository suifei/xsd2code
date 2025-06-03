package core

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"sync"
	"time"
)

// LogLevel 日志级别
type LogLevel int

const (
	LogLevelTrace LogLevel = iota
	LogLevelDebug
	LogLevelInfo
	LogLevelWarn
	LogLevelError
	LogLevelFatal
)

// String 返回日志级别的字符串表示
func (ll LogLevel) String() string {
	switch ll {
	case LogLevelTrace:
		return "TRACE"
	case LogLevelDebug:
		return "DEBUG"
	case LogLevelInfo:
		return "INFO"
	case LogLevelWarn:
		return "WARN"
	case LogLevelError:
		return "ERROR"
	case LogLevelFatal:
		return "FATAL"
	default:
		return "UNKNOWN"
	}
}

// ParseLogLevel 解析日志级别
func ParseLogLevel(level string) LogLevel {
	switch strings.ToUpper(level) {
	case "TRACE":
		return LogLevelTrace
	case "DEBUG":
		return LogLevelDebug
	case "INFO":
		return LogLevelInfo
	case "WARN", "WARNING":
		return LogLevelWarn
	case "ERROR":
		return LogLevelError
	case "FATAL":
		return LogLevelFatal
	default:
		return LogLevelInfo
	}
}

// LogEntry 日志条目
type LogEntry struct {
	Timestamp time.Time              `json:"timestamp"`
	Level     LogLevel               `json:"level"`
	Message   string                 `json:"message"`
	Fields    map[string]interface{} `json:"fields,omitempty"`
	Caller    string                 `json:"caller,omitempty"`
	Stack     string                 `json:"stack,omitempty"`
}

// LogFormatter 日志格式器接口
type LogFormatter interface {
	Format(entry LogEntry) string
}

// TextFormatter 文本格式器
type TextFormatter struct {
	TimestampFormat string
	EnableColors    bool
	EnableCaller    bool
}

// Format 格式化日志条目
func (tf *TextFormatter) Format(entry LogEntry) string {
	timestamp := entry.Timestamp.Format(tf.TimestampFormat)
	if tf.TimestampFormat == "" {
		timestamp = entry.Timestamp.Format("2006-01-02 15:04:05")
	}

	level := entry.Level.String()
	if tf.EnableColors {
		level = tf.colorizeLevel(level)
	}

	message := fmt.Sprintf("%s [%s] %s", timestamp, level, entry.Message)

	// 添加字段
	if len(entry.Fields) > 0 {
		var fields []string
		for key, value := range entry.Fields {
			fields = append(fields, fmt.Sprintf("%s=%v", key, value))
		}
		message += " " + strings.Join(fields, " ")
	}

	// 添加调用者信息
	if tf.EnableCaller && entry.Caller != "" {
		message += fmt.Sprintf(" (%s)", entry.Caller)
	}

	return message
}

// colorizeLevel 为日志级别添加颜色
func (tf *TextFormatter) colorizeLevel(level string) string {
	if !tf.EnableColors {
		return level
	}

	switch level {
	case "TRACE":
		return "\033[37m" + level + "\033[0m" // 白色
	case "DEBUG":
		return "\033[36m" + level + "\033[0m" // 青色
	case "INFO":
		return "\033[32m" + level + "\033[0m" // 绿色
	case "WARN":
		return "\033[33m" + level + "\033[0m" // 黄色
	case "ERROR":
		return "\033[31m" + level + "\033[0m" // 红色
	case "FATAL":
		return "\033[35m" + level + "\033[0m" // 紫色
	default:
		return level
	}
}

// JSONFormatter JSON格式器
type JSONFormatter struct{}

// Format 格式化日志条目为JSON
func (jf *JSONFormatter) Format(entry LogEntry) string {
	data, _ := json.Marshal(entry)
	return string(data)
}

// Logger 增强日志器
type Logger struct {
	level        LogLevel
	formatter    LogFormatter
	outputs      []io.Writer
	enableCaller bool
	enableStack  bool
	mutex        sync.RWMutex
	fields       map[string]interface{}
}

// NewLogger 创建新的日志器
func NewLogger(level LogLevel, formatter LogFormatter, outputs ...io.Writer) *Logger {
	if len(outputs) == 0 {
		outputs = []io.Writer{os.Stdout}
	}

	return &Logger{
		level:     level,
		formatter: formatter,
		outputs:   outputs,
		fields:    make(map[string]interface{}),
	}
}

// WithField 添加字段
func (l *Logger) WithField(key string, value interface{}) *Logger {
	l.mutex.Lock()
	defer l.mutex.Unlock()

	newLogger := &Logger{
		level:        l.level,
		formatter:    l.formatter,
		outputs:      l.outputs,
		enableCaller: l.enableCaller,
		enableStack:  l.enableStack,
		fields:       make(map[string]interface{}),
	}

	// 复制现有字段
	for k, v := range l.fields {
		newLogger.fields[k] = v
	}

	// 添加新字段
	newLogger.fields[key] = value

	return newLogger
}

// WithFields 添加多个字段
func (l *Logger) WithFields(fields map[string]interface{}) *Logger {
	l.mutex.Lock()
	defer l.mutex.Unlock()

	newLogger := &Logger{
		level:        l.level,
		formatter:    l.formatter,
		outputs:      l.outputs,
		enableCaller: l.enableCaller,
		enableStack:  l.enableStack,
		fields:       make(map[string]interface{}),
	}

	// 复制现有字段
	for k, v := range l.fields {
		newLogger.fields[k] = v
	}

	// 添加新字段
	for k, v := range fields {
		newLogger.fields[k] = v
	}

	return newLogger
}

// EnableCaller 启用调用者信息
func (l *Logger) EnableCaller(enable bool) *Logger {
	l.enableCaller = enable
	return l
}

// EnableStack 启用堆栈信息
func (l *Logger) EnableStack(enable bool) *Logger {
	l.enableStack = enable
	return l
}

// SetLevel 设置日志级别
func (l *Logger) SetLevel(level LogLevel) {
	l.level = level
}

// log 内部日志方法
func (l *Logger) log(level LogLevel, message string) {
	if level < l.level {
		return
	}

	entry := LogEntry{
		Timestamp: time.Now(),
		Level:     level,
		Message:   message,
		Fields:    make(map[string]interface{}),
	}

	// 复制字段
	l.mutex.RLock()
	for k, v := range l.fields {
		entry.Fields[k] = v
	}
	l.mutex.RUnlock()

	// 获取调用者信息
	if l.enableCaller {
		if pc, file, line, ok := runtime.Caller(2); ok {
			funcName := runtime.FuncForPC(pc).Name()
			entry.Caller = fmt.Sprintf("%s:%d %s", filepath.Base(file), line, funcName)
		}
	}

	// 获取堆栈信息
	if l.enableStack && level >= LogLevelError {
		buf := make([]byte, 4096)
		n := runtime.Stack(buf, false)
		entry.Stack = string(buf[:n])
	}

	// 格式化并输出
	formatted := l.formatter.Format(entry)

	for _, output := range l.outputs {
		fmt.Fprintln(output, formatted)
	}
}

// Trace 跟踪级别日志
func (l *Logger) Trace(message string) {
	l.log(LogLevelTrace, message)
}

// Tracef 格式化跟踪级别日志
func (l *Logger) Tracef(format string, args ...interface{}) {
	l.log(LogLevelTrace, fmt.Sprintf(format, args...))
}

// Debug 调试级别日志
func (l *Logger) Debug(message string) {
	l.log(LogLevelDebug, message)
}

// Debugf 格式化调试级别日志
func (l *Logger) Debugf(format string, args ...interface{}) {
	l.log(LogLevelDebug, fmt.Sprintf(format, args...))
}

// Info 信息级别日志
func (l *Logger) Info(message string) {
	l.log(LogLevelInfo, message)
}

// Infof 格式化信息级别日志
func (l *Logger) Infof(format string, args ...interface{}) {
	l.log(LogLevelInfo, fmt.Sprintf(format, args...))
}

// Warn 警告级别日志
func (l *Logger) Warn(message string) {
	l.log(LogLevelWarn, message)
}

// Warnf 格式化警告级别日志
func (l *Logger) Warnf(format string, args ...interface{}) {
	l.log(LogLevelWarn, fmt.Sprintf(format, args...))
}

// Error 错误级别日志
func (l *Logger) Error(message string) {
	l.log(LogLevelError, message)
}

// Errorf 格式化错误级别日志
func (l *Logger) Errorf(format string, args ...interface{}) {
	l.log(LogLevelError, fmt.Sprintf(format, args...))
}

// Fatal 致命级别日志
func (l *Logger) Fatal(message string) {
	l.log(LogLevelFatal, message)
	os.Exit(1)
}

// Fatalf 格式化致命级别日志
func (l *Logger) Fatalf(format string, args ...interface{}) {
	l.log(LogLevelFatal, fmt.Sprintf(format, args...))
	os.Exit(1)
}

// LogManager 日志管理器
type LogManager struct {
	logger        *Logger
	config        *LoggingConfig
	errorManager  *ErrorManager
	fileRotator   *FileRotator
	metricsLogger *MetricsLogger
}

// NewLogManager 创建日志管理器
func NewLogManager(config *LoggingConfig, errorManager *ErrorManager) (*LogManager, error) {
	var outputs []io.Writer
	var fileRotator *FileRotator

	// 控制台输出
	if config.EnableConsole {
		outputs = append(outputs, os.Stdout)
	}

	// 文件输出
	if config.EnableFile && config.OutputFile != "" {
		// 确保目录存在
		dir := filepath.Dir(config.OutputFile)
		if err := os.MkdirAll(dir, 0755); err != nil {
			return nil, fmt.Errorf("failed to create log directory: %v", err)
		}

		// 创建文件轮转器
		fr, err := NewFileRotator(config.OutputFile, config.MaxFileSize, config.MaxBackups)
		if err != nil {
			return nil, fmt.Errorf("failed to create file rotator: %v", err)
		}
		fileRotator = fr
		outputs = append(outputs, fr)
	}

	// 选择格式器
	var formatter LogFormatter
	switch strings.ToLower(config.Format) {
	case "json":
		formatter = &JSONFormatter{}
	default:
		formatter = &TextFormatter{
			TimestampFormat: "2006-01-02 15:04:05",
			EnableColors:    config.EnableConsole,
			EnableCaller:    true,
		}
	}

	// 创建日志器
	logger := NewLogger(ParseLogLevel(config.Level), formatter, outputs...)
	logger.EnableCaller(true)

	lm := &LogManager{
		logger:       logger,
		config:       config,
		errorManager: errorManager,
		fileRotator:  fileRotator,
	}

	// 创建指标日志器
	if config.EnableMetrics {
		lm.metricsLogger = NewMetricsLogger(logger)
	}

	return lm, nil
}

// GetLogger 获取日志器
func (lm *LogManager) GetLogger() *Logger {
	return lm.logger
}

// GetMetricsLogger 获取指标日志器
func (lm *LogManager) GetMetricsLogger() *MetricsLogger {
	return lm.metricsLogger
}

// Close 关闭日志管理器
func (lm *LogManager) Close() error {
	if lm.fileRotator != nil {
		return lm.fileRotator.Close()
	}
	return nil
}

// FileRotator 文件轮转器
type FileRotator struct {
	filename    string
	maxSize     int64
	maxBackups  int
	currentSize int64
	file        *os.File
	mutex       sync.Mutex
}

// NewFileRotator 创建文件轮转器
func NewFileRotator(filename string, maxSize string, maxBackups int) (*FileRotator, error) {
	size, err := parseSize(maxSize)
	if err != nil {
		return nil, err
	}

	fr := &FileRotator{
		filename:   filename,
		maxSize:    size,
		maxBackups: maxBackups,
	}

	if err := fr.openFile(); err != nil {
		return nil, err
	}

	return fr, nil
}

// Write 写入数据
func (fr *FileRotator) Write(data []byte) (int, error) {
	fr.mutex.Lock()
	defer fr.mutex.Unlock()

	// 检查是否需要轮转
	if fr.currentSize+int64(len(data)) > fr.maxSize {
		if err := fr.rotate(); err != nil {
			return 0, err
		}
	}

	n, err := fr.file.Write(data)
	if err == nil {
		fr.currentSize += int64(n)
	}

	return n, err
}

// Close 关闭文件
func (fr *FileRotator) Close() error {
	fr.mutex.Lock()
	defer fr.mutex.Unlock()

	if fr.file != nil {
		return fr.file.Close()
	}
	return nil
}

// openFile 打开文件
func (fr *FileRotator) openFile() error {
	file, err := os.OpenFile(fr.filename, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		return err
	}

	// 获取当前文件大小
	info, err := file.Stat()
	if err != nil {
		file.Close()
		return err
	}

	fr.file = file
	fr.currentSize = info.Size()

	return nil
}

// rotate 轮转文件
func (fr *FileRotator) rotate() error {
	// 关闭当前文件
	if fr.file != nil {
		fr.file.Close()
	}

	// 移动备份文件
	for i := fr.maxBackups - 1; i >= 1; i-- {
		oldName := fmt.Sprintf("%s.%d", fr.filename, i)
		newName := fmt.Sprintf("%s.%d", fr.filename, i+1)

		if _, err := os.Stat(oldName); err == nil {
			os.Rename(oldName, newName)
		}
	}

	// 移动当前文件为备份
	if _, err := os.Stat(fr.filename); err == nil {
		os.Rename(fr.filename, fr.filename+".1")
	}

	// 创建新文件
	fr.currentSize = 0
	return fr.openFile()
}

// parseSize 解析大小字符串
func parseSize(size string) (int64, error) {
	if size == "" {
		return 100 * 1024 * 1024, nil // 默认100MB
	}

	multiplier := int64(1)
	size = strings.ToUpper(size)

	if strings.HasSuffix(size, "KB") {
		multiplier = 1024
		size = strings.TrimSuffix(size, "KB")
	} else if strings.HasSuffix(size, "MB") {
		multiplier = 1024 * 1024
		size = strings.TrimSuffix(size, "MB")
	} else if strings.HasSuffix(size, "GB") {
		multiplier = 1024 * 1024 * 1024
		size = strings.TrimSuffix(size, "GB")
	}

	value := int64(100) // 默认值
	if size != "" {
		if _, err := fmt.Sscanf(size, "%d", &value); err != nil {
			return 0, err
		}
	}

	return value * multiplier, nil
}

// MetricsLogger 指标日志器
type MetricsLogger struct {
	logger *Logger
	mutex  sync.Mutex
	buffer []MetricEntry
}

// MetricEntry 指标条目
type MetricEntry struct {
	Timestamp time.Time              `json:"timestamp"`
	Name      string                 `json:"name"`
	Value     interface{}            `json:"value"`
	Tags      map[string]interface{} `json:"tags,omitempty"`
}

// NewMetricsLogger 创建指标日志器
func NewMetricsLogger(logger *Logger) *MetricsLogger {
	ml := &MetricsLogger{
		logger: logger,
		buffer: make([]MetricEntry, 0, 100),
	}

	// 启动定期刷新
	go ml.periodicFlush()

	return ml
}

// LogMetric 记录指标
func (ml *MetricsLogger) LogMetric(name string, value interface{}, tags ...map[string]interface{}) {
	ml.mutex.Lock()
	defer ml.mutex.Unlock()

	entry := MetricEntry{
		Timestamp: time.Now(),
		Name:      name,
		Value:     value,
	}

	if len(tags) > 0 {
		entry.Tags = tags[0]
	}

	ml.buffer = append(ml.buffer, entry)

	// 如果缓冲区满了，立即刷新
	if len(ml.buffer) >= 100 {
		ml.flushBuffer()
	}
}

// periodicFlush 定期刷新缓冲区
func (ml *MetricsLogger) periodicFlush() {
	ticker := time.NewTicker(10 * time.Second)
	defer ticker.Stop()

	for range ticker.C {
		ml.mutex.Lock()
		ml.flushBuffer()
		ml.mutex.Unlock()
	}
}

// flushBuffer 刷新缓冲区
func (ml *MetricsLogger) flushBuffer() {
	if len(ml.buffer) == 0 {
		return
	}

	for _, entry := range ml.buffer {
		ml.logger.WithFields(map[string]interface{}{
			"metric_name":  entry.Name,
			"metric_value": entry.Value,
			"metric_tags":  entry.Tags,
		}).Info("METRIC")
	}

	ml.buffer = ml.buffer[:0]
}

// 全局日志器实例
var (
	defaultLogger *Logger
	loggerMutex   sync.RWMutex
)

// SetDefaultLogger 设置默认日志器
func SetDefaultLogger(logger *Logger) {
	loggerMutex.Lock()
	defer loggerMutex.Unlock()
	defaultLogger = logger
}

// GetDefaultLogger 获取默认日志器
func GetDefaultLogger() *Logger {
	loggerMutex.RLock()
	defer loggerMutex.RUnlock()

	if defaultLogger == nil {
		// 创建默认日志器
		formatter := &TextFormatter{
			TimestampFormat: "2006-01-02 15:04:05",
			EnableColors:    true,
			EnableCaller:    true,
		}
		defaultLogger = NewLogger(LogLevelInfo, formatter, os.Stdout)
	}

	return defaultLogger
}

// 便利函数
func Trace(message string)                      { GetDefaultLogger().Trace(message) }
func Tracef(format string, args ...interface{}) { GetDefaultLogger().Tracef(format, args...) }
func Debug(message string)                      { GetDefaultLogger().Debug(message) }
func Debugf(format string, args ...interface{}) { GetDefaultLogger().Debugf(format, args...) }
func Info(message string)                       { GetDefaultLogger().Info(message) }
func Infof(format string, args ...interface{})  { GetDefaultLogger().Infof(format, args...) }
func Warn(message string)                       { GetDefaultLogger().Warn(message) }
func Warnf(format string, args ...interface{})  { GetDefaultLogger().Warnf(format, args...) }
func Error(message string)                      { GetDefaultLogger().Error(message) }
func Errorf(format string, args ...interface{}) { GetDefaultLogger().Errorf(format, args...) }
func Fatal(message string)                      { GetDefaultLogger().Fatal(message) }
func Fatalf(format string, args ...interface{}) { GetDefaultLogger().Fatalf(format, args...) }
