package core

import (
	"context"
	"fmt"
	"hash/fnv"
	"runtime"
	"sort"
	"sync"
	"sync/atomic"
	"time"
)

// XSDChunk XSD文件分片
type XSDChunk struct {
	ID           string
	Data         []byte
	StartOffset  int64
	EndOffset    int64
	Priority     int
	Dependencies []string
}

// XSDProcessingTask XSD处理任务
type XSDProcessingTask struct {
	ID           string
	Chunk        *XSDChunk
	ProcessFunc  func(*XSDChunk) error
	Priority     int
	Timeout      time.Duration
	Dependencies []string
}

// Execute 实现Task接口
func (t *XSDProcessingTask) Execute(ctx context.Context) error {
	if t.ProcessFunc == nil {
		return fmt.Errorf("process function is nil")
	}
	return t.ProcessFunc(t.Chunk)
}

// GetID 获取任务ID
func (t *XSDProcessingTask) GetID() string {
	return t.ID
}

// GetPriority 获取任务优先级
func (t *XSDProcessingTask) GetPriority() int {
	return t.Priority
}

// GetTimeout 获取任务超时时间
func (t *XSDProcessingTask) GetTimeout() time.Duration {
	return t.Timeout
}

// XSDSplitter XSD文件分割器
type XSDSplitter struct {
	maxChunkSize int64
	minChunkSize int64
	overlap      int64
}

// NewXSDSplitter 创建XSD分割器
func NewXSDSplitter(maxChunkSize, minChunkSize, overlap int64) *XSDSplitter {
	return &XSDSplitter{
		maxChunkSize: maxChunkSize,
		minChunkSize: minChunkSize,
		overlap:      overlap,
	}
}

// SplitXSD 分割XSD文件
func (s *XSDSplitter) SplitXSD(data []byte) ([]*XSDChunk, error) {
	if len(data) <= int(s.minChunkSize) {
		// 文件太小，不需要分割
		return []*XSDChunk{{
			ID:          s.generateChunkID(data, 0),
			Data:        data,
			StartOffset: 0,
			EndOffset:   int64(len(data)),
			Priority:    1,
		}}, nil
	}

	var chunks []*XSDChunk
	var currentOffset int64 = 0
	chunkIndex := 0

	for currentOffset < int64(len(data)) {
		endOffset := currentOffset + s.maxChunkSize
		if endOffset > int64(len(data)) {
			endOffset = int64(len(data))
		}

		// 查找合适的分割点（避免在标签中间分割）
		actualEndOffset := s.findSplitPoint(data, currentOffset, endOffset)

		// 提取chunk数据
		chunkData := make([]byte, actualEndOffset-currentOffset)
		copy(chunkData, data[currentOffset:actualEndOffset])

		chunk := &XSDChunk{
			ID:          s.generateChunkID(chunkData, chunkIndex),
			Data:        chunkData,
			StartOffset: currentOffset,
			EndOffset:   actualEndOffset,
			Priority:    s.calculatePriority(chunkData),
		}

		chunks = append(chunks, chunk)
		currentOffset = actualEndOffset - s.overlap
		chunkIndex++
	}

	// 分析依赖关系
	s.analyzeDependencies(chunks)

	return chunks, nil
}

// findSplitPoint 查找合适的分割点
func (s *XSDSplitter) findSplitPoint(data []byte, start, maxEnd int64) int64 {
	if maxEnd >= int64(len(data)) {
		return int64(len(data))
	}

	// 从maxEnd向前查找合适的分割点
	for i := maxEnd; i > start+s.minChunkSize; i-- {
		if i >= int64(len(data)) {
			continue
		}

		// 查找标签结束位置
		if data[i] == '>' {
			// 检查是否是完整的标签结束
			if i+1 < int64(len(data)) && (data[i+1] == '\n' || data[i+1] == '\r' || data[i+1] == '<') {
				return i + 1
			}
		}
	}

	// 如果找不到合适的分割点，使用maxEnd
	return maxEnd
}

// generateChunkID 生成chunk ID
func (s *XSDSplitter) generateChunkID(data []byte, index int) string {
	h := fnv.New64a()
	h.Write(data)
	return fmt.Sprintf("chunk-%d-%x", index, h.Sum64())
}

// calculatePriority 计算chunk优先级
func (s *XSDSplitter) calculatePriority(data []byte) int {
	// 根据内容类型确定优先级
	content := string(data)

	// schema定义优先级最高
	if contains(content, "<xs:schema") || contains(content, "<xsd:schema") {
		return 10
	}

	// 复杂类型优先级较高
	if contains(content, "<xs:complexType") || contains(content, "<xsd:complexType") {
		return 8
	}

	// 简单类型优先级中等
	if contains(content, "<xs:simpleType") || contains(content, "<xsd:simpleType") {
		return 6
	}

	// 元素定义优先级中等
	if contains(content, "<xs:element") || contains(content, "<xsd:element") {
		return 5
	}

	// 属性定义优先级较低
	if contains(content, "<xs:attribute") || contains(content, "<xsd:attribute") {
		return 3
	}

	// 默认优先级
	return 1
}

// contains 检查字符串是否包含子字符串
func contains(s, substr string) bool {
	return len(s) >= len(substr) &&
		func() bool {
			for i := 0; i <= len(s)-len(substr); i++ {
				if s[i:i+len(substr)] == substr {
					return true
				}
			}
			return false
		}()
}

// analyzeDependencies 分析chunks之间的依赖关系
func (s *XSDSplitter) analyzeDependencies(chunks []*XSDChunk) {
	// 构建类型名称到chunk的映射
	typeToChunk := make(map[string]string)

	for _, chunk := range chunks {
		content := string(chunk.Data)
		types := s.extractTypeNames(content)
		for _, typeName := range types {
			typeToChunk[typeName] = chunk.ID
		}
	}

	// 分析每个chunk的依赖
	for _, chunk := range chunks {
		content := string(chunk.Data)
		references := s.extractTypeReferences(content)

		var dependencies []string
		for _, ref := range references {
			if depChunkID, exists := typeToChunk[ref]; exists && depChunkID != chunk.ID {
				dependencies = append(dependencies, depChunkID)
			}
		}

		chunk.Dependencies = dependencies
	}
}

// extractTypeNames 提取类型名称
func (s *XSDSplitter) extractTypeNames(content string) []string {
	var types []string
	// 简化的类型提取逻辑
	// 实际实现需要更复杂的XML解析
	return types
}

// extractTypeReferences 提取类型引用
func (s *XSDSplitter) extractTypeReferences(content string) []string {
	var references []string
	// 简化的引用提取逻辑
	// 实际实现需要更复杂的XML解析
	return references
}

// Task 任务接口
type Task interface {
	Execute(ctx context.Context) error
	GetID() string
	GetPriority() int
	GetTimeout() time.Duration
}

// TaskResult 任务结果
type TaskResult struct {
	TaskID    string        `json:"task_id"`
	Success   bool          `json:"success"`
	Error     error         `json:"error,omitempty"`
	Duration  time.Duration `json:"duration"`
	StartTime time.Time     `json:"start_time"`
	EndTime   time.Time     `json:"end_time"`
	ChunkID   string        `json:"chunk_id,omitempty"`
	Data      interface{}   `json:"data,omitempty"`
}

// PriorityQueue 优先级队列
type PriorityQueue struct {
	tasks []Task
	mutex sync.RWMutex
}

// NewPriorityQueue 创建优先级队列
func NewPriorityQueue() *PriorityQueue {
	return &PriorityQueue{
		tasks: make([]Task, 0),
	}
}

// Push 添加任务
func (pq *PriorityQueue) Push(task Task) {
	pq.mutex.Lock()
	defer pq.mutex.Unlock()

	pq.tasks = append(pq.tasks, task)
	sort.Slice(pq.tasks, func(i, j int) bool {
		return pq.tasks[i].GetPriority() > pq.tasks[j].GetPriority()
	})
}

// Pop 取出最高优先级任务
func (pq *PriorityQueue) Pop() Task {
	pq.mutex.Lock()
	defer pq.mutex.Unlock()

	if len(pq.tasks) == 0 {
		return nil
	}

	task := pq.tasks[0]
	pq.tasks = pq.tasks[1:]
	return task
}

// Size 获取队列大小
func (pq *PriorityQueue) Size() int {
	pq.mutex.RLock()
	defer pq.mutex.RUnlock()
	return len(pq.tasks)
}

// WorkerPool 工作池
type WorkerPool struct {
	ctx          context.Context
	cancel       context.CancelFunc
	workerCount  int
	taskQueue    chan Task
	resultQueue  chan TaskResult
	workers      []*Worker
	wg           sync.WaitGroup
	metrics      *WorkerPoolMetrics
	errorManager *ErrorManager
	stopped      atomic.Bool
}

// WorkerPoolMetrics 工作池指标
type WorkerPoolMetrics struct {
	TasksSubmitted  atomic.Int64
	TasksCompleted  atomic.Int64
	TasksFailed     atomic.Int64
	TasksTimeout    atomic.Int64
	TotalDuration   atomic.Int64
	AverageDuration atomic.Int64
	ActiveWorkers   atomic.Int64
	QueueSize       atomic.Int64
	PeakQueueSize   atomic.Int64
}

// NewWorkerPool 创建工作池
func NewWorkerPool(workerCount, queueSize int, errorManager *ErrorManager) *WorkerPool {
	ctx, cancel := context.WithCancel(context.Background())

	wp := &WorkerPool{
		ctx:          ctx,
		cancel:       cancel,
		workerCount:  workerCount,
		taskQueue:    make(chan Task, queueSize),
		resultQueue:  make(chan TaskResult, queueSize),
		workers:      make([]*Worker, workerCount),
		metrics:      &WorkerPoolMetrics{},
		errorManager: errorManager,
	}

	// 创建工作协程
	for i := 0; i < workerCount; i++ {
		worker := NewWorker(i, wp.taskQueue, wp.resultQueue, wp.metrics, errorManager)
		wp.workers[i] = worker
	}

	// 启动指标更新协程
	go wp.updateMetrics()

	return wp
}

// Start 启动工作池
func (wp *WorkerPool) Start() {
	if wp.stopped.Load() {
		return
	}

	for _, worker := range wp.workers {
		wp.wg.Add(1)
		go func(w *Worker) {
			defer wp.wg.Done()
			w.Start(wp.ctx)
		}(worker)
	}
}

// Submit 提交任务
func (wp *WorkerPool) Submit(task Task) error {
	if wp.stopped.Load() {
		return NewConfigError("WORKER_POOL_STOPPED", "Worker pool is stopped").Build()
	}

	select {
	case wp.taskQueue <- task:
		wp.metrics.TasksSubmitted.Add(1)
		currentQueueSize := wp.metrics.QueueSize.Load()
		currentQueueSize++
		wp.metrics.QueueSize.Store(currentQueueSize)

		// 更新峰值队列大小
		for {
			peakSize := wp.metrics.PeakQueueSize.Load()
			if currentQueueSize <= peakSize || wp.metrics.PeakQueueSize.CompareAndSwap(peakSize, currentQueueSize) {
				break
			}
		}
		return nil
	case <-wp.ctx.Done():
		return NewTimeoutError("TASK_SUBMIT_CANCELLED", "Task submission cancelled").Build()
	default:
		return NewConfigError("TASK_QUEUE_FULL", "Task queue is full").Build()
	}
}

// GetResults 获取结果通道
func (wp *WorkerPool) GetResults() <-chan TaskResult {
	return wp.resultQueue
}

// Stop 停止工作池
func (wp *WorkerPool) Stop(timeout time.Duration) error {
	if wp.stopped.Swap(true) {
		return nil // 已经停止
	}

	// 关闭任务队列
	close(wp.taskQueue)

	// 等待所有工作协程完成或超时
	done := make(chan struct{})
	go func() {
		wp.wg.Wait()
		close(done)
	}()

	select {
	case <-done:
		wp.cancel()
		close(wp.resultQueue)
		return nil
	case <-time.After(timeout):
		wp.cancel()
		close(wp.resultQueue)
		return NewTimeoutError("WORKER_POOL_STOP_TIMEOUT",
			fmt.Sprintf("Worker pool stop timeout after %v", timeout)).Build()
	}
}

// GetMetrics 获取指标
func (wp *WorkerPool) GetMetrics() *WorkerPoolMetrics {
	return wp.metrics
}

// updateMetrics 更新指标
func (wp *WorkerPool) updateMetrics() {
	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-wp.ctx.Done():
			return
		case <-ticker.C:
			// 更新平均持续时间
			completed := wp.metrics.TasksCompleted.Load()
			if completed > 0 {
				avgDuration := wp.metrics.TotalDuration.Load() / completed
				wp.metrics.AverageDuration.Store(avgDuration)
			}

			// 更新活跃工作协程数
			activeCount := int64(0)
			for _, worker := range wp.workers {
				if worker.IsActive() {
					activeCount++
				}
			}
			wp.metrics.ActiveWorkers.Store(activeCount)
		}
	}
}

// Worker 工作协程
type Worker struct {
	id           int
	taskQueue    <-chan Task
	resultQueue  chan<- TaskResult
	metrics      *WorkerPoolMetrics
	errorManager *ErrorManager
	active       atomic.Bool
}

// NewWorker 创建工作协程
func NewWorker(id int, taskQueue <-chan Task, resultQueue chan<- TaskResult,
	metrics *WorkerPoolMetrics, errorManager *ErrorManager) *Worker {
	return &Worker{
		id:           id,
		taskQueue:    taskQueue,
		resultQueue:  resultQueue,
		metrics:      metrics,
		errorManager: errorManager,
	}
}

// Start 启动工作协程
func (w *Worker) Start(ctx context.Context) {
	for {
		select {
		case task, ok := <-w.taskQueue:
			if !ok {
				return // 任务队列已关闭
			}

			w.active.Store(true)
			result := w.executeTask(ctx, task)
			w.active.Store(false)

			// 更新队列大小
			w.metrics.QueueSize.Add(-1)

			// 发送结果
			select {
			case w.resultQueue <- result:
			case <-ctx.Done():
				return
			}

		case <-ctx.Done():
			return
		}
	}
}

// executeTask 执行任务
func (w *Worker) executeTask(ctx context.Context, task Task) TaskResult {
	startTime := time.Now()
	result := TaskResult{
		TaskID:    task.GetID(),
		StartTime: startTime,
	}

	// 设置任务超时
	taskCtx := ctx
	if timeout := task.GetTimeout(); timeout > 0 {
		var cancel context.CancelFunc
		taskCtx, cancel = context.WithTimeout(ctx, timeout)
		defer cancel()
	}

	// 执行任务
	err := task.Execute(taskCtx)
	endTime := time.Now()
	duration := endTime.Sub(startTime)

	result.EndTime = endTime
	result.Duration = duration
	result.Success = err == nil
	result.Error = err

	// 更新指标
	w.metrics.TotalDuration.Add(int64(duration))
	if err == nil {
		w.metrics.TasksCompleted.Add(1)
	} else {
		w.metrics.TasksFailed.Add(1)

		// 检查是否为超时错误
		if ctx.Err() == context.DeadlineExceeded {
			w.metrics.TasksTimeout.Add(1)
		}

		// 记录错误
		if w.errorManager != nil {
			w.errorManager.AddError(NewGenerationError("TASK_EXECUTION_FAILED",
				fmt.Sprintf("Task %s failed: %v", task.GetID(), err)).
				WithContext(fmt.Sprintf("Worker %d", w.id)).Build())
		}
	}

	return result
}

// IsActive 检查是否活跃
func (w *Worker) IsActive() bool {
	return w.active.Load()
}

// ParseTask XSD解析任务
type ParseTask struct {
	ID       string
	XSDPath  string
	Priority int
	Timeout  time.Duration
	Parser   XSDParser // 假设有这个接口
}

// Execute 执行解析任务
func (pt *ParseTask) Execute(ctx context.Context) error {
	// 实现XSD解析逻辑
	// 这里是示例实现
	select {
	case <-ctx.Done():
		return ctx.Err()
	default:
		// 模拟解析工作
		time.Sleep(100 * time.Millisecond)
		return nil
	}
}

// GetID 获取任务ID
func (pt *ParseTask) GetID() string {
	return pt.ID
}

// GetPriority 获取优先级
func (pt *ParseTask) GetPriority() int {
	return pt.Priority
}

// GetTimeout 获取超时时间
func (pt *ParseTask) GetTimeout() time.Duration {
	return pt.Timeout
}

// XSDParser 假设的XSD解析器接口
type XSDParser interface {
	Parse(ctx context.Context, xsdPath string) error
}

// CodeGenTask 代码生成任务
type CodeGenTask struct {
	ID        string
	TypeName  string
	Schema    interface{} // 假设的Schema类型
	Priority  int
	Timeout   time.Duration
	Generator CodeGenerator // 假设有这个接口
}

// Execute 执行代码生成任务
func (cgt *CodeGenTask) Execute(ctx context.Context) error {
	select {
	case <-ctx.Done():
		return ctx.Err()
	default:
		// 模拟代码生成工作
		time.Sleep(200 * time.Millisecond)
		return nil
	}
}

// GetID 获取任务ID
func (cgt *CodeGenTask) GetID() string {
	return cgt.ID
}

// GetPriority 获取优先级
func (cgt *CodeGenTask) GetPriority() int {
	return cgt.Priority
}

// GetTimeout 获取超时时间
func (cgt *CodeGenTask) GetTimeout() time.Duration {
	return cgt.Timeout
}

// CodeGenerator 假设的代码生成器接口
type CodeGenerator interface {
	Generate(ctx context.Context, typeName string, schema interface{}) error
}

// ConcurrentProcessor 并发处理器
type ConcurrentProcessor struct {
	workerPool   *WorkerPool
	cacheManager *CacheManager
	config       *Config
	errorManager *ErrorManager
	metrics      *ProcessorMetrics
}

// ProcessorMetrics 处理器指标
type ProcessorMetrics struct {
	FilesProcessed atomic.Int64
	TypesGenerated atomic.Int64
	CacheHits      atomic.Int64
	CacheMisses    atomic.Int64
	ProcessingTime atomic.Int64
	ErrorCount     atomic.Int64
}

// NewConcurrentProcessor 创建并发处理器
func NewConcurrentProcessor(config *Config, cacheManager *CacheManager,
	errorManager *ErrorManager) *ConcurrentProcessor {

	workerCount := config.Performance.MaxWorkers
	if workerCount <= 0 {
		workerCount = runtime.NumCPU()
	}

	queueSize := workerCount * 10 // 队列大小为工作协程数的10倍

	return &ConcurrentProcessor{
		workerPool:   NewWorkerPool(workerCount, queueSize, errorManager),
		cacheManager: cacheManager,
		config:       config,
		errorManager: errorManager,
		metrics:      &ProcessorMetrics{},
	}
}

// Start 启动处理器
func (cp *ConcurrentProcessor) Start() {
	cp.workerPool.Start()
}

// ProcessXSDFiles 并发处理XSD文件
func (cp *ConcurrentProcessor) ProcessXSDFiles(xsdPaths []string) error {
	if len(xsdPaths) == 0 {
		return nil
	}

	// 提交解析任务
	for i, xsdPath := range xsdPaths {
		task := &ParseTask{
			ID:       fmt.Sprintf("parse-%d", i),
			XSDPath:  xsdPath,
			Priority: 1,
			Timeout:  30 * time.Second,
		}

		if err := cp.workerPool.Submit(task); err != nil {
			return err
		}
	}

	// 收集结果
	resultCount := 0
	errorCount := 0

	for resultCount < len(xsdPaths) {
		select {
		case result := <-cp.workerPool.GetResults():
			resultCount++
			if !result.Success {
				errorCount++
				cp.metrics.ErrorCount.Add(1)
			}
			cp.metrics.ProcessingTime.Add(int64(result.Duration))

		case <-time.After(5 * time.Minute):
			return NewTimeoutError("PROCESSING_TIMEOUT",
				"XSD processing timeout").Build()
		}
	}

	cp.metrics.FilesProcessed.Add(int64(len(xsdPaths)))

	if errorCount > 0 {
		return NewGenerationError("PARTIAL_PROCESSING_FAILED",
			fmt.Sprintf("Failed to process %d out of %d files", errorCount, len(xsdPaths))).Build()
	}

	return nil
}

// Stop 停止处理器
func (cp *ConcurrentProcessor) Stop() error {
	return cp.workerPool.Stop(30 * time.Second)
}

// GetMetrics 获取处理器指标
func (cp *ConcurrentProcessor) GetMetrics() *ProcessorMetrics {
	return cp.metrics
}

// GetStats 获取处理器统计信息
func (cp *ConcurrentProcessor) GetStats() ProcessorStats {
	poolMetrics := cp.workerPool.GetMetrics()
	cacheStats := cp.cacheManager.GetStats()

	return ProcessorStats{
		WorkerPoolMetrics: poolMetrics,
		CacheStats:        cacheStats,
		ProcessorMetrics:  cp.metrics,
		StartTime:         time.Now(), // 应该记录实际启动时间
	}
}

// ProcessorStats 处理器统计信息
type ProcessorStats struct {
	WorkerPoolMetrics *WorkerPoolMetrics `json:"worker_pool_metrics"`
	CacheStats        CacheStats         `json:"cache_stats"`
	ProcessorMetrics  *ProcessorMetrics  `json:"processor_metrics"`
	StartTime         time.Time          `json:"start_time"`
}

// String 格式化输出处理器统计
func (ps *ProcessorStats) String() string {
	return fmt.Sprintf("Processor Stats:\n"+
		"  Files Processed: %d\n"+
		"  Types Generated: %d\n"+
		"  Worker Pool: %d active workers, %d tasks completed\n"+
		"  Cache: %s\n"+
		"  Errors: %d",
		ps.ProcessorMetrics.FilesProcessed.Load(),
		ps.ProcessorMetrics.TypesGenerated.Load(),
		ps.WorkerPoolMetrics.ActiveWorkers.Load(),
		ps.WorkerPoolMetrics.TasksCompleted.Load(),
		ps.CacheStats.String(),
		ps.ProcessorMetrics.ErrorCount.Load())
}

// SmartWorkerPool 智能工作池，支持依赖管理和优先级调度
type SmartWorkerPool struct {
	*WorkerPool
	priorityQueue   *PriorityQueue
	dependencyGraph map[string][]string
	completedTasks  map[string]bool
	waitingTasks    map[string][]Task
	mutex           sync.RWMutex
	scheduler       *TaskScheduler
}

// TaskScheduler 任务调度器
type TaskScheduler struct {
	pool            *SmartWorkerPool
	schedulerTicker *time.Ticker
	stopChan        chan struct{}
	running         atomic.Bool
}

// NewSmartWorkerPool 创建智能工作池
func NewSmartWorkerPool(workerCount, queueSize int, errorManager *ErrorManager) *SmartWorkerPool {
	basePool := NewWorkerPool(workerCount, queueSize, errorManager)

	swp := &SmartWorkerPool{
		WorkerPool:      basePool,
		priorityQueue:   NewPriorityQueue(),
		dependencyGraph: make(map[string][]string),
		completedTasks:  make(map[string]bool),
		waitingTasks:    make(map[string][]Task),
	}

	swp.scheduler = &TaskScheduler{
		pool:            swp,
		schedulerTicker: time.NewTicker(100 * time.Millisecond),
		stopChan:        make(chan struct{}),
	}

	return swp
}

// Start 启动智能工作池
func (swp *SmartWorkerPool) Start() {
	swp.WorkerPool.Start()

	// 启动调度器
	go swp.scheduler.run()

	// 启动结果处理器
	go swp.handleResults()
}

// SubmitWithDependencies 提交带依赖的任务
func (swp *SmartWorkerPool) SubmitWithDependencies(task Task, dependencies []string) error {
	swp.mutex.Lock()
	defer swp.mutex.Unlock()

	taskID := task.GetID()

	// 记录依赖关系
	if len(dependencies) > 0 {
		swp.dependencyGraph[taskID] = dependencies
	}

	// 检查依赖是否已完成
	if swp.canSchedule(taskID) {
		swp.priorityQueue.Push(task)
	} else {
		// 添加到等待队列
		for _, dep := range dependencies {
			swp.waitingTasks[dep] = append(swp.waitingTasks[dep], task)
		}
	}

	return nil
}

// canSchedule 检查任务是否可以调度
func (swp *SmartWorkerPool) canSchedule(taskID string) bool {
	dependencies, exists := swp.dependencyGraph[taskID]
	if !exists {
		return true // 无依赖，可以调度
	}

	for _, dep := range dependencies {
		if !swp.completedTasks[dep] {
			return false // 依赖未完成
		}
	}

	return true
}

// handleResults 处理任务结果
func (swp *SmartWorkerPool) handleResults() {
	for result := range swp.GetResults() {
		swp.mutex.Lock()

		// 标记任务为已完成
		swp.completedTasks[result.TaskID] = true

		// 检查等待的任务
		if waitingTasks, exists := swp.waitingTasks[result.TaskID]; exists {
			for _, task := range waitingTasks {
				if swp.canSchedule(task.GetID()) {
					swp.priorityQueue.Push(task)
				}
			}
			delete(swp.waitingTasks, result.TaskID)
		}

		swp.mutex.Unlock()
	}
}

// run 运行任务调度器
func (ts *TaskScheduler) run() {
	ts.running.Store(true)
	defer ts.running.Store(false)

	for {
		select {
		case <-ts.schedulerTicker.C:
			ts.scheduleNext()
		case <-ts.stopChan:
			return
		}
	}
}

// scheduleNext 调度下一个任务
func (ts *TaskScheduler) scheduleNext() {
	task := ts.pool.priorityQueue.Pop()
	if task != nil {
		// 提交到基础工作池
		if err := ts.pool.WorkerPool.Submit(task); err != nil {
			// 如果提交失败，重新放回优先级队列
			ts.pool.priorityQueue.Push(task)
		}
	}
}

// Stop 停止智能工作池
func (swp *SmartWorkerPool) Stop(timeout time.Duration) error {
	// 停止调度器
	if swp.scheduler.running.Load() {
		close(swp.scheduler.stopChan)
		swp.scheduler.schedulerTicker.Stop()
	}

	// 停止基础工作池
	return swp.WorkerPool.Stop(timeout)
}

// GetMetrics 获取工作池指标
func (swp *SmartWorkerPool) GetMetrics() *SmartWorkerPoolMetrics {
	baseMetrics := swp.WorkerPool.metrics
	swp.mutex.RLock()
	defer swp.mutex.RUnlock()

	return &SmartWorkerPoolMetrics{
		WorkerPoolMetrics: baseMetrics,
		PendingTasks:      int64(swp.priorityQueue.Size()),
		WaitingTasks:      int64(len(swp.waitingTasks)),
		CompletedTasks:    int64(len(swp.completedTasks)),
		DependencyCount:   int64(len(swp.dependencyGraph)),
	}
}

// SmartWorkerPoolMetrics 智能工作池指标
type SmartWorkerPoolMetrics struct {
	*WorkerPoolMetrics
	PendingTasks    int64 `json:"pending_tasks"`
	WaitingTasks    int64 `json:"waiting_tasks"`
	CompletedTasks  int64 `json:"completed_tasks"`
	DependencyCount int64 `json:"dependency_count"`
}

// ProcessXSDConcurrently 并发处理XSD文件
func (swp *SmartWorkerPool) ProcessXSDConcurrently(data []byte, processFunc func(*XSDChunk) error) error {
	// 创建分割器
	splitter := NewXSDSplitter(
		64*1024, // 64KB最大分片大小
		8*1024,  // 8KB最小分片大小
		1024,    // 1KB重叠
	)

	// 分割XSD文件
	chunks, err := splitter.SplitXSD(data)
	if err != nil {
		return fmt.Errorf("failed to split XSD: %w", err)
	}

	// 创建处理任务
	tasks := make([]*XSDProcessingTask, len(chunks))
	for i, chunk := range chunks {
		tasks[i] = &XSDProcessingTask{
			ID:           chunk.ID,
			Chunk:        chunk,
			ProcessFunc:  processFunc,
			Priority:     chunk.Priority,
			Timeout:      30 * time.Second,
			Dependencies: chunk.Dependencies,
		}
	}

	// 提交任务
	for _, task := range tasks {
		if err := swp.SubmitWithDependencies(task, task.Dependencies); err != nil {
			return fmt.Errorf("failed to submit task %s: %w", task.ID, err)
		}
	}

	return nil
}
