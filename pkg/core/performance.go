package core

import (
	"context"
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
	"time"
)

// PerformanceMetrics 性能指标收集器
type PerformanceMetrics struct {
	startTime          time.Time
	memoryUsage        atomic.Int64
	gcPauses           atomic.Int64
	allocations        atomic.Int64
	operationCounts    map[string]*atomic.Int64
	operationDurations map[string]*atomic.Int64
	mutex              sync.RWMutex
}

// NewPerformanceMetrics 创建性能指标收集器
func NewPerformanceMetrics() *PerformanceMetrics {
	return &PerformanceMetrics{
		startTime:          time.Now(),
		operationCounts:    make(map[string]*atomic.Int64),
		operationDurations: make(map[string]*atomic.Int64),
	}
}

// StartOperation 开始操作计时
func (pm *PerformanceMetrics) StartOperation(name string) *OperationTimer {
	pm.mutex.Lock()
	defer pm.mutex.Unlock()

	if _, exists := pm.operationCounts[name]; !exists {
		pm.operationCounts[name] = &atomic.Int64{}
		pm.operationDurations[name] = &atomic.Int64{}
	}

	pm.operationCounts[name].Add(1)
	return &OperationTimer{
		startTime: time.Now(),
		metrics:   pm,
		name:      name,
	}
}

// RecordMemoryUsage 记录内存使用情况
func (pm *PerformanceMetrics) RecordMemoryUsage() {
	var ms runtime.MemStats
	runtime.ReadMemStats(&ms)
	pm.memoryUsage.Store(int64(ms.Alloc))
	pm.allocations.Store(int64(ms.TotalAlloc))
	pm.gcPauses.Store(int64(ms.PauseTotalNs))
}

// GetReport 获取性能报告
func (pm *PerformanceMetrics) GetReport() PerformanceReport {
	pm.mutex.RLock()
	defer pm.mutex.RUnlock()

	report := PerformanceReport{
		TotalDuration: time.Since(pm.startTime),
		MemoryUsage:   pm.memoryUsage.Load(),
		GCPauses:      time.Duration(pm.gcPauses.Load()),
		Allocations:   pm.allocations.Load(),
		Operations:    make(map[string]OperationStats),
	}

	for name, count := range pm.operationCounts {
		duration := pm.operationDurations[name]
		report.Operations[name] = OperationStats{
			Count:       count.Load(),
			TotalTime:   time.Duration(duration.Load()),
			AverageTime: time.Duration(duration.Load() / max(count.Load(), 1)),
		}
	}

	return report
}

// OperationTimer 操作计时器
type OperationTimer struct {
	startTime time.Time
	metrics   *PerformanceMetrics
	name      string
}

// Stop 停止计时并记录
func (ot *OperationTimer) Stop() {
	duration := time.Since(ot.startTime)
	ot.metrics.mutex.RLock()
	if counter, exists := ot.metrics.operationDurations[ot.name]; exists {
		counter.Add(int64(duration))
	}
	ot.metrics.mutex.RUnlock()
}

// PerformanceReport 性能报告
type PerformanceReport struct {
	TotalDuration time.Duration             `json:"total_duration"`
	MemoryUsage   int64                     `json:"memory_usage_bytes"`
	GCPauses      time.Duration             `json:"gc_pauses"`
	Allocations   int64                     `json:"total_allocations"`
	Operations    map[string]OperationStats `json:"operations"`
}

// OperationStats 操作统计
type OperationStats struct {
	Count       int64         `json:"count"`
	TotalTime   time.Duration `json:"total_time"`
	AverageTime time.Duration `json:"average_time"`
}

// String 格式化输出性能报告
func (pr PerformanceReport) String() string {
	output := fmt.Sprintf("=== Performance Report ===\n")
	output += fmt.Sprintf("Total Duration: %v\n", pr.TotalDuration)
	output += fmt.Sprintf("Memory Usage: %.2f MB\n", float64(pr.MemoryUsage)/1024/1024)
	output += fmt.Sprintf("GC Pauses: %v\n", pr.GCPauses)
	output += fmt.Sprintf("Total Allocations: %.2f MB\n", float64(pr.Allocations)/1024/1024)
	output += fmt.Sprintf("\nOperation Statistics:\n")

	for name, stats := range pr.Operations {
		output += fmt.Sprintf("  %s: %d ops, total=%v, avg=%v\n",
			name, stats.Count, stats.TotalTime, stats.AverageTime)
	}

	return output
}

// MemoryPool 内存池管理器
type MemoryPool struct {
	bufferPool  sync.Pool
	nodePool    sync.Pool
	stringPool  sync.Pool
	maxPoolSize int
	currentSize atomic.Int64
}

// NewMemoryPool 创建内存池
func NewMemoryPool(maxSize int) *MemoryPool {
	return &MemoryPool{
		maxPoolSize: maxSize,
		bufferPool: sync.Pool{
			New: func() interface{} {
				return make([]byte, 0, 1024)
			},
		},
		nodePool: sync.Pool{
			New: func() interface{} {
				return &XMLNode{}
			},
		},
		stringPool: sync.Pool{
			New: func() interface{} {
				return make([]string, 0, 16)
			},
		},
	}
}

// GetBuffer 获取缓冲区
func (mp *MemoryPool) GetBuffer() []byte {
	if mp.currentSize.Load() < int64(mp.maxPoolSize) {
		mp.currentSize.Add(1)
		return mp.bufferPool.Get().([]byte)[:0]
	}
	return make([]byte, 0, 1024)
}

// PutBuffer 归还缓冲区
func (mp *MemoryPool) PutBuffer(buf []byte) {
	if cap(buf) <= 64*1024 { // 只缓存小于64KB的缓冲区
		mp.bufferPool.Put(buf)
		mp.currentSize.Add(-1)
	}
}

// GetNode 获取XML节点
func (mp *MemoryPool) GetNode() *XMLNode {
	return mp.nodePool.Get().(*XMLNode)
}

// PutNode 归还XML节点
func (mp *MemoryPool) PutNode(node *XMLNode) {
	node.Reset() // 重置节点状态
	mp.nodePool.Put(node)
}

// XMLNode XML节点结构（简化版）
type XMLNode struct {
	Name       string
	Attributes map[string]string
	Children   []*XMLNode
	Text       string
}

// Reset 重置节点
func (n *XMLNode) Reset() {
	n.Name = ""
	n.Text = ""
	if n.Attributes != nil {
		for k := range n.Attributes {
			delete(n.Attributes, k)
		}
	}
	n.Children = n.Children[:0]
}

// ResourceManager 资源管理器
type ResourceManager struct {
	memoryPool *MemoryPool
	metrics    *PerformanceMetrics
	ctx        context.Context
	cancel     context.CancelFunc
	workerPool chan struct{}
	maxWorkers int
}

// NewResourceManager 创建资源管理器
func NewResourceManager(maxWorkers int, memoryPoolSize int) *ResourceManager {
	ctx, cancel := context.WithCancel(context.Background())

	rm := &ResourceManager{
		memoryPool: NewMemoryPool(memoryPoolSize),
		metrics:    NewPerformanceMetrics(),
		ctx:        ctx,
		cancel:     cancel,
		workerPool: make(chan struct{}, maxWorkers),
		maxWorkers: maxWorkers,
	}

	// 启动内存监控协程
	go rm.monitorMemory()

	return rm
}

// AcquireWorker 获取工作协程
func (rm *ResourceManager) AcquireWorker() {
	rm.workerPool <- struct{}{}
}

// ReleaseWorker 释放工作协程
func (rm *ResourceManager) ReleaseWorker() {
	<-rm.workerPool
}

// GetMemoryPool 获取内存池
func (rm *ResourceManager) GetMemoryPool() *MemoryPool {
	return rm.memoryPool
}

// GetMetrics 获取性能指标
func (rm *ResourceManager) GetMetrics() *PerformanceMetrics {
	return rm.metrics
}

// Shutdown 关闭资源管理器
func (rm *ResourceManager) Shutdown() {
	rm.cancel()
}

// monitorMemory 监控内存使用
func (rm *ResourceManager) monitorMemory() {
	ticker := time.NewTicker(5 * time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-rm.ctx.Done():
			return
		case <-ticker.C:
			rm.metrics.RecordMemoryUsage()

			// 内存压力检查
			var ms runtime.MemStats
			runtime.ReadMemStats(&ms)
			if ms.Alloc > 1024*1024*1024 { // 超过1GB触发GC
				runtime.GC()
			}
		}
	}
}

// max 辅助函数
func max(a, b int64) int64 {
	if a > b {
		return a
	}
	return b
}
