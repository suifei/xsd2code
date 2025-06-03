package core

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math"
	"os"
	"path/filepath"
	"runtime"
	"sort"
	"sync"
	"time"
)

// BenchmarkType 基准测试类型
type BenchmarkType int

const (
	BenchmarkParsing BenchmarkType = iota
	BenchmarkGeneration
	BenchmarkValidation
	BenchmarkConcurrency
	BenchmarkMemory
	BenchmarkDisk
	BenchmarkEnd2End
)

// String 返回基准测试类型的字符串表示
func (bt BenchmarkType) String() string {
	switch bt {
	case BenchmarkParsing:
		return "PARSING"
	case BenchmarkGeneration:
		return "GENERATION"
	case BenchmarkValidation:
		return "VALIDATION"
	case BenchmarkConcurrency:
		return "CONCURRENCY"
	case BenchmarkMemory:
		return "MEMORY"
	case BenchmarkDisk:
		return "DISK"
	case BenchmarkEnd2End:
		return "END_TO_END"
	default:
		return "UNKNOWN"
	}
}

// BenchmarkConfig 基准测试配置
type BenchmarkConfig struct {
	Name             string        `json:"name"`
	Type             BenchmarkType `json:"type"`
	Iterations       int           `json:"iterations"`
	ConcurrencyLevel int           `json:"concurrency_level"`
	WarmupIterations int           `json:"warmup_iterations"`
	TimeoutPerTest   time.Duration `json:"timeout_per_test"`
	MemoryLimit      int64         `json:"memory_limit_bytes"`
	TestDataSize     int           `json:"test_data_size"`
	EnableProfiling  bool          `json:"enable_profiling"`
	OutputDir        string        `json:"output_dir"`
	CompareBaseline  bool          `json:"compare_baseline"`
	BaselineFile     string        `json:"baseline_file"`
}

// DefaultBenchmarkConfig 默认基准测试配置
func DefaultBenchmarkConfig() BenchmarkConfig {
	return BenchmarkConfig{
		Iterations:       100,
		ConcurrencyLevel: runtime.NumCPU(),
		WarmupIterations: 10,
		TimeoutPerTest:   30 * time.Second,
		MemoryLimit:      1024 * 1024 * 1024, // 1GB
		TestDataSize:     1000,
		EnableProfiling:  false,
		OutputDir:        "./benchmark_results",
		CompareBaseline:  false,
	}
}

// BenchmarkResult 单次基准测试结果
type BenchmarkResult struct {
	TestName        string                 `json:"test_name"`
	Type            BenchmarkType          `json:"type"`
	Duration        time.Duration          `json:"duration"`
	MemoryUsed      int64                  `json:"memory_used_bytes"`
	MemoryAllocated int64                  `json:"memory_allocated_bytes"`
	GCPauses        time.Duration          `json:"gc_pauses"`
	Success         bool                   `json:"success"`
	Error           string                 `json:"error,omitempty"`
	Timestamp       time.Time              `json:"timestamp"`
	Metadata        map[string]interface{} `json:"metadata,omitempty"`
}

// BenchmarkSuite 基准测试套件结果
type BenchmarkSuite struct {
	Config        BenchmarkConfig   `json:"config"`
	Results       []BenchmarkResult `json:"results"`
	Statistics    BenchmarkStats    `json:"statistics"`
	SystemInfo    SystemInfo        `json:"system_info"`
	StartTime     time.Time         `json:"start_time"`
	EndTime       time.Time         `json:"end_time"`
	TotalDuration time.Duration     `json:"total_duration"`
}

// BenchmarkStats 基准测试统计
type BenchmarkStats struct {
	TotalTests      int     `json:"total_tests"`
	SuccessfulTests int     `json:"successful_tests"`
	FailedTests     int     `json:"failed_tests"`
	SuccessRate     float64 `json:"success_rate_percent"`

	// 时间统计
	MinDuration    time.Duration `json:"min_duration"`
	MaxDuration    time.Duration `json:"max_duration"`
	AvgDuration    time.Duration `json:"avg_duration"`
	MedianDuration time.Duration `json:"median_duration"`
	StdDevDuration time.Duration `json:"stddev_duration"`

	// 内存统计
	MinMemory    int64 `json:"min_memory_bytes"`
	MaxMemory    int64 `json:"max_memory_bytes"`
	AvgMemory    int64 `json:"avg_memory_bytes"`
	MedianMemory int64 `json:"median_memory_bytes"`

	// 吞吐量统计
	TestsPerSecond float64 `json:"tests_per_second"`

	// 百分位数
	P50Duration time.Duration `json:"p50_duration"`
	P90Duration time.Duration `json:"p90_duration"`
	P95Duration time.Duration `json:"p95_duration"`
	P99Duration time.Duration `json:"p99_duration"`
}

// SystemInfo 系统信息
type SystemInfo struct {
	OS           string    `json:"os"`
	Architecture string    `json:"architecture"`
	CPUCount     int       `json:"cpu_count"`
	GoVersion    string    `json:"go_version"`
	Timestamp    time.Time `json:"timestamp"`
}

// getSystemInfo 获取系统信息
func getSystemInfo() SystemInfo {
	return SystemInfo{
		OS:           runtime.GOOS,
		Architecture: runtime.GOARCH,
		CPUCount:     runtime.NumCPU(),
		GoVersion:    runtime.Version(),
		Timestamp:    time.Now(),
	}
}

// BenchmarkRunner 基准测试运行器
type BenchmarkRunner struct {
	config   BenchmarkConfig
	results  []BenchmarkResult
	mutex    sync.RWMutex
	baseline *BenchmarkSuite
	context  context.Context
}

// NewBenchmarkRunner 创建基准测试运行器
func NewBenchmarkRunner(config BenchmarkConfig) *BenchmarkRunner {
	runner := &BenchmarkRunner{
		config:  config,
		results: make([]BenchmarkResult, 0),
		context: context.Background(),
	}

	// 创建输出目录
	os.MkdirAll(config.OutputDir, 0755)

	// 加载基线数据
	if config.CompareBaseline && config.BaselineFile != "" {
		runner.loadBaseline()
	}

	return runner
}

// loadBaseline 加载基线数据
func (br *BenchmarkRunner) loadBaseline() {
	data, err := ioutil.ReadFile(br.config.BaselineFile)
	if err != nil {
		fmt.Printf("Warning: Could not load baseline file: %v\n", err)
		return
	}

	var baseline BenchmarkSuite
	if err := json.Unmarshal(data, &baseline); err != nil {
		fmt.Printf("Warning: Could not parse baseline file: %v\n", err)
		return
	}

	br.baseline = &baseline
	fmt.Printf("Loaded baseline from %s with %d results\n",
		br.config.BaselineFile, len(baseline.Results))
}

// RunTest 运行单个测试
func (br *BenchmarkRunner) RunTest(name string, benchmarkType BenchmarkType, testFunc func() error) BenchmarkResult {
	result := BenchmarkResult{
		TestName:  name,
		Type:      benchmarkType,
		Timestamp: time.Now(),
		Metadata:  make(map[string]interface{}),
	}

	// 记录初始内存状态
	var memStart runtime.MemStats
	runtime.GC()
	runtime.ReadMemStats(&memStart)

	// 运行测试
	startTime := time.Now()
	err := testFunc()
	duration := time.Since(startTime)

	// 记录结束内存状态
	var memEnd runtime.MemStats
	runtime.ReadMemStats(&memEnd)

	// 设置结果
	result.Duration = duration
	result.MemoryUsed = int64(memEnd.Alloc - memStart.Alloc)
	result.MemoryAllocated = int64(memEnd.TotalAlloc - memStart.TotalAlloc)
	result.GCPauses = time.Duration(memEnd.PauseTotalNs - memStart.PauseTotalNs)
	result.Success = err == nil

	if err != nil {
		result.Error = err.Error()
	}

	// 添加元数据
	result.Metadata["gc_count"] = memEnd.NumGC - memStart.NumGC
	result.Metadata["heap_objects"] = memEnd.HeapObjects
	result.Metadata["goroutines"] = runtime.NumGoroutine()

	br.mutex.Lock()
	br.results = append(br.results, result)
	br.mutex.Unlock()

	return result
}

// RunSuite 运行基准测试套件
func (br *BenchmarkRunner) RunSuite(tests map[string]func() error) *BenchmarkSuite {
	suite := &BenchmarkSuite{
		Config:     br.config,
		SystemInfo: getSystemInfo(),
		StartTime:  time.Now(),
	}

	fmt.Printf("Starting benchmark suite: %s\n", br.config.Name)
	fmt.Printf("Configuration: %d iterations, %d concurrency\n",
		br.config.Iterations, br.config.ConcurrencyLevel)

	// 预热阶段
	if br.config.WarmupIterations > 0 {
		fmt.Printf("Warmup phase: %d iterations\n", br.config.WarmupIterations)
		for name, testFunc := range tests {
			for i := 0; i < br.config.WarmupIterations; i++ {
				br.RunTest(fmt.Sprintf("warmup_%s_%d", name, i), BenchmarkEnd2End, testFunc)
			}
		}
		// 清除预热结果
		br.mutex.Lock()
		br.results = make([]BenchmarkResult, 0)
		br.mutex.Unlock()
	}

	// 正式测试阶段
	fmt.Printf("Running %d tests with %d iterations each\n", len(tests), br.config.Iterations)

	var wg sync.WaitGroup
	semaphore := make(chan struct{}, br.config.ConcurrencyLevel)

	for name, testFunc := range tests {
		for i := 0; i < br.config.Iterations; i++ {
			wg.Add(1)
			semaphore <- struct{}{} // 获取信号量

			go func(testName string, iteration int, fn func() error) {
				defer wg.Done()
				defer func() { <-semaphore }() // 释放信号量

				testNameWithIteration := fmt.Sprintf("%s_iter_%d", testName, iteration)
				br.RunTest(testNameWithIteration, BenchmarkEnd2End, fn)
			}(name, i, testFunc)
		}
	}

	wg.Wait()

	suite.EndTime = time.Now()
	suite.TotalDuration = suite.EndTime.Sub(suite.StartTime)
	suite.Results = br.results
	suite.Statistics = br.calculateStatistics()

	fmt.Printf("Benchmark suite completed in %v\n", suite.TotalDuration)
	fmt.Printf("Success rate: %.2f%% (%d/%d tests)\n",
		suite.Statistics.SuccessRate,
		suite.Statistics.SuccessfulTests,
		suite.Statistics.TotalTests)

	return suite
}

// calculateStatistics 计算统计信息
func (br *BenchmarkRunner) calculateStatistics() BenchmarkStats {
	br.mutex.RLock()
	defer br.mutex.RUnlock()

	if len(br.results) == 0 {
		return BenchmarkStats{}
	}

	stats := BenchmarkStats{
		TotalTests: len(br.results),
	}

	// 分离成功和失败的测试
	var durations []time.Duration
	var memories []int64

	for _, result := range br.results {
		if result.Success {
			stats.SuccessfulTests++
			durations = append(durations, result.Duration)
			memories = append(memories, result.MemoryUsed)
		} else {
			stats.FailedTests++
		}
	}

	if stats.SuccessfulTests == 0 {
		return stats
	}

	stats.SuccessRate = float64(stats.SuccessfulTests) / float64(stats.TotalTests) * 100

	// 排序以计算百分位数
	sort.Slice(durations, func(i, j int) bool { return durations[i] < durations[j] })
	sort.Slice(memories, func(i, j int) bool { return memories[i] < memories[j] })

	// 时间统计
	stats.MinDuration = durations[0]
	stats.MaxDuration = durations[len(durations)-1]
	stats.AvgDuration = br.calculateAvgDuration(durations)
	stats.MedianDuration = durations[len(durations)/2]
	stats.StdDevDuration = br.calculateStdDevDuration(durations, stats.AvgDuration)

	// 内存统计
	stats.MinMemory = memories[0]
	stats.MaxMemory = memories[len(memories)-1]
	stats.AvgMemory = br.calculateAvgMemory(memories)
	stats.MedianMemory = memories[len(memories)/2]

	// 吞吐量统计
	if stats.AvgDuration > 0 {
		stats.TestsPerSecond = float64(time.Second) / float64(stats.AvgDuration)
	}

	// 百分位数
	stats.P50Duration = durations[len(durations)*50/100]
	stats.P90Duration = durations[len(durations)*90/100]
	stats.P95Duration = durations[len(durations)*95/100]
	stats.P99Duration = durations[len(durations)*99/100]

	return stats
}

// calculateAvgDuration 计算平均时间
func (br *BenchmarkRunner) calculateAvgDuration(durations []time.Duration) time.Duration {
	var total time.Duration
	for _, d := range durations {
		total += d
	}
	return total / time.Duration(len(durations))
}

// calculateAvgMemory 计算平均内存
func (br *BenchmarkRunner) calculateAvgMemory(memories []int64) int64 {
	var total int64
	for _, m := range memories {
		total += m
	}
	return total / int64(len(memories))
}

// calculateStdDevDuration 计算时间标准差
func (br *BenchmarkRunner) calculateStdDevDuration(durations []time.Duration, avg time.Duration) time.Duration {
	var sum float64
	for _, d := range durations {
		diff := float64(d - avg)
		sum += diff * diff
	}
	variance := sum / float64(len(durations))
	return time.Duration(math.Sqrt(variance))
}

// SaveResults 保存结果到文件
func (br *BenchmarkRunner) SaveResults(suite *BenchmarkSuite) error {
	timestamp := time.Now().Format("20060102_150405")
	filename := filepath.Join(br.config.OutputDir, fmt.Sprintf("benchmark_%s_%s.json", br.config.Name, timestamp))

	data, err := json.MarshalIndent(suite, "", "  ")
	if err != nil {
		return fmt.Errorf("failed to marshal results: %w", err)
	}

	if err := ioutil.WriteFile(filename, data, 0644); err != nil {
		return fmt.Errorf("failed to write results file: %w", err)
	}

	fmt.Printf("Results saved to: %s\n", filename)
	return nil
}

// GenerateReport 生成性能报告
func (br *BenchmarkRunner) GenerateReport(suite *BenchmarkSuite) string {
	report := fmt.Sprintf("Benchmark Report: %s\n", br.config.Name)
	report += fmt.Sprintf("=================%s\n", "=================================")
	report += fmt.Sprintf("Configuration:\n")
	report += fmt.Sprintf("  - Iterations: %d\n", br.config.Iterations)
	report += fmt.Sprintf("  - Concurrency: %d\n", br.config.ConcurrencyLevel)
	report += fmt.Sprintf("  - Total Duration: %v\n", suite.TotalDuration)
	report += fmt.Sprintf("\n")

	stats := suite.Statistics
	report += fmt.Sprintf("Results:\n")
	report += fmt.Sprintf("  - Total Tests: %d\n", stats.TotalTests)
	report += fmt.Sprintf("  - Successful: %d (%.2f%%)\n", stats.SuccessfulTests, stats.SuccessRate)
	report += fmt.Sprintf("  - Failed: %d\n", stats.FailedTests)
	report += fmt.Sprintf("\n")

	report += fmt.Sprintf("Performance:\n")
	report += fmt.Sprintf("  - Min Duration: %v\n", stats.MinDuration)
	report += fmt.Sprintf("  - Max Duration: %v\n", stats.MaxDuration)
	report += fmt.Sprintf("  - Avg Duration: %v\n", stats.AvgDuration)
	report += fmt.Sprintf("  - Median Duration: %v\n", stats.MedianDuration)
	report += fmt.Sprintf("  - Std Dev: %v\n", stats.StdDevDuration)
	report += fmt.Sprintf("  - Tests/Second: %.2f\n", stats.TestsPerSecond)
	report += fmt.Sprintf("\n")

	report += fmt.Sprintf("Percentiles:\n")
	report += fmt.Sprintf("  - P50: %v\n", stats.P50Duration)
	report += fmt.Sprintf("  - P90: %v\n", stats.P90Duration)
	report += fmt.Sprintf("  - P95: %v\n", stats.P95Duration)
	report += fmt.Sprintf("  - P99: %v\n", stats.P99Duration)
	report += fmt.Sprintf("\n")

	report += fmt.Sprintf("Memory:\n")
	report += fmt.Sprintf("  - Min Memory: %s\n", br.formatBytes(stats.MinMemory))
	report += fmt.Sprintf("  - Max Memory: %s\n", br.formatBytes(stats.MaxMemory))
	report += fmt.Sprintf("  - Avg Memory: %s\n", br.formatBytes(stats.AvgMemory))
	report += fmt.Sprintf("  - Median Memory: %s\n", br.formatBytes(stats.MedianMemory))
	report += fmt.Sprintf("\n")

	// 与基线比较
	if br.baseline != nil {
		report += br.generateBaselineComparison(suite)
	}

	return report
}

// generateBaselineComparison 生成基线比较
func (br *BenchmarkRunner) generateBaselineComparison(suite *BenchmarkSuite) string {
	current := suite.Statistics
	baseline := br.baseline.Statistics

	report := fmt.Sprintf("Baseline Comparison:\n")
	report += fmt.Sprintf("  - Duration Change: %v (%.2f%%)\n",
		current.AvgDuration-baseline.AvgDuration,
		float64(current.AvgDuration-baseline.AvgDuration)/float64(baseline.AvgDuration)*100)

	report += fmt.Sprintf("  - Memory Change: %s (%.2f%%)\n",
		br.formatBytes(current.AvgMemory-baseline.AvgMemory),
		float64(current.AvgMemory-baseline.AvgMemory)/float64(baseline.AvgMemory)*100)

	report += fmt.Sprintf("  - Success Rate Change: %.2f%%\n",
		current.SuccessRate-baseline.SuccessRate)

	report += fmt.Sprintf("\n")
	return report
}

// formatBytes 格式化字节数
func (br *BenchmarkRunner) formatBytes(bytes int64) string {
	const unit = 1024
	if bytes < unit {
		return fmt.Sprintf("%d B", bytes)
	}
	div, exp := int64(unit), 0
	for n := bytes / unit; n >= unit; n /= unit {
		div *= unit
		exp++
	}
	return fmt.Sprintf("%.1f %cB", float64(bytes)/float64(div), "KMGTPE"[exp])
}

// CompareWithBaseline 与基线比较
func (br *BenchmarkRunner) CompareWithBaseline(suite *BenchmarkSuite) *BaselineComparison {
	if br.baseline == nil {
		return nil
	}

	return &BaselineComparison{
		Current:           suite.Statistics,
		Baseline:          br.baseline.Statistics,
		PerformanceChange: br.calculatePerformanceChange(suite.Statistics, br.baseline.Statistics),
		Recommendation:    br.generateRecommendation(suite.Statistics, br.baseline.Statistics),
		GeneratedAt:       time.Now(),
	}
}

// BaselineComparison 基线比较结果
type BaselineComparison struct {
	Current           BenchmarkStats    `json:"current"`
	Baseline          BenchmarkStats    `json:"baseline"`
	PerformanceChange PerformanceChange `json:"performance_change"`
	Recommendation    string            `json:"recommendation"`
	GeneratedAt       time.Time         `json:"generated_at"`
}

// PerformanceChange 性能变化
type PerformanceChange struct {
	DurationChangePercent   float64 `json:"duration_change_percent"`
	MemoryChangePercent     float64 `json:"memory_change_percent"`
	SuccessRateChange       float64 `json:"success_rate_change"`
	ThroughputChangePercent float64 `json:"throughput_change_percent"`
}

// calculatePerformanceChange 计算性能变化
func (br *BenchmarkRunner) calculatePerformanceChange(current, baseline BenchmarkStats) PerformanceChange {
	return PerformanceChange{
		DurationChangePercent:   float64(current.AvgDuration-baseline.AvgDuration) / float64(baseline.AvgDuration) * 100,
		MemoryChangePercent:     float64(current.AvgMemory-baseline.AvgMemory) / float64(baseline.AvgMemory) * 100,
		SuccessRateChange:       current.SuccessRate - baseline.SuccessRate,
		ThroughputChangePercent: (current.TestsPerSecond - baseline.TestsPerSecond) / baseline.TestsPerSecond * 100,
	}
}

// generateRecommendation 生成建议
func (br *BenchmarkRunner) generateRecommendation(current, baseline BenchmarkStats) string {
	change := br.calculatePerformanceChange(current, baseline)

	var recommendations []string

	if change.DurationChangePercent > 10 {
		recommendations = append(recommendations, "Performance degradation detected. Consider optimizing algorithms or reducing computational complexity.")
	} else if change.DurationChangePercent < -10 {
		recommendations = append(recommendations, "Performance improvement detected. Great work!")
	}

	if change.MemoryChangePercent > 20 {
		recommendations = append(recommendations, "Memory usage increased significantly. Check for memory leaks or optimize data structures.")
	}

	if change.SuccessRateChange < -5 {
		recommendations = append(recommendations, "Success rate decreased. Investigate error causes and improve error handling.")
	}

	if len(recommendations) == 0 {
		recommendations = append(recommendations, "Performance is stable compared to baseline.")
	}
	result := "Recommendations:\n"
	for i, rec := range recommendations {
		result += fmt.Sprintf("  %d. %s\n", i+1, rec)
	}

	return result
}
