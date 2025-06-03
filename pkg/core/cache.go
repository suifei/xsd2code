package core

import (
	"crypto/md5"
	"encoding/gob"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"sort"
	"sync"
	"time"
)

// CachePolicy 缓存策略
type CachePolicy int

const (
	CachePolicyLRU CachePolicy = iota
	CachePolicyLFU
	CachePolicyTTL
	CachePolicySize
)

// CacheCleanupStrategy 清理策略
type CacheCleanupStrategy struct {
	MaxDiskSizeMB      int64         `json:"max_disk_size_mb"`
	MaxMemorySizeMB    int64         `json:"max_memory_size_mb"`
	CleanupInterval    time.Duration `json:"cleanup_interval"`
	MaxAge             time.Duration `json:"max_age"`
	MinAccessCount     int64         `json:"min_access_count"`
	CleanupPercentage  float64       `json:"cleanup_percentage"`
	PersistenceEnabled bool          `json:"persistence_enabled"`
	BackupEnabled      bool          `json:"backup_enabled"`
	CompressionEnabled bool          `json:"compression_enabled"`
}

// DefaultCleanupStrategy 默认清理策略
func DefaultCleanupStrategy() CacheCleanupStrategy {
	return CacheCleanupStrategy{
		MaxDiskSizeMB:      1024, // 1GB
		MaxMemorySizeMB:    512,  // 512MB
		CleanupInterval:    5 * time.Minute,
		MaxAge:             24 * time.Hour,
		MinAccessCount:     2,
		CleanupPercentage:  0.25, // 清理25%的条目
		PersistenceEnabled: true,
		BackupEnabled:      false,
		CompressionEnabled: true,
	}
}

// CacheMetadata 缓存元数据
type CacheMetadata struct {
	Entries     map[string]CacheEntryMetadata `json:"entries"`
	LastCleanup time.Time                     `json:"last_cleanup"`
	Version     string                        `json:"version"`
	TotalSize   int64                         `json:"total_size"`
}

// CacheEntryMetadata 缓存条目元数据
type CacheEntryMetadata struct {
	Key         string    `json:"key"`
	CreatedAt   time.Time `json:"created_at"`
	ExpiresAt   time.Time `json:"expires_at"`
	AccessedAt  time.Time `json:"accessed_at"`
	AccessCount int64     `json:"access_count"`
	Size        int64     `json:"size_bytes"`
	Checksum    string    `json:"checksum"`
	Compressed  bool      `json:"compressed"`
}

// CachePersistenceManager 缓存持久化管理器
type CachePersistenceManager struct {
	metadataFile string
	dataDir      string
	strategy     CacheCleanupStrategy
	mutex        sync.RWMutex
}

// NewCachePersistenceManager 创建缓存持久化管理器
func NewCachePersistenceManager(baseDir string, strategy CacheCleanupStrategy) *CachePersistenceManager {
	return &CachePersistenceManager{
		metadataFile: filepath.Join(baseDir, "cache_metadata.json"),
		dataDir:      filepath.Join(baseDir, "data"),
		strategy:     strategy,
	}
}

// SaveMetadata 保存元数据
func (cp *CachePersistenceManager) SaveMetadata(metadata CacheMetadata) error {
	cp.mutex.Lock()
	defer cp.mutex.Unlock()

	data, err := json.MarshalIndent(metadata, "", "  ")
	if err != nil {
		return fmt.Errorf("failed to marshal metadata: %w", err)
	}

	return ioutil.WriteFile(cp.metadataFile, data, 0644)
}

// LoadMetadata 加载元数据
func (cp *CachePersistenceManager) LoadMetadata() (CacheMetadata, error) {
	cp.mutex.RLock()
	defer cp.mutex.RUnlock()

	var metadata CacheMetadata
	data, err := ioutil.ReadFile(cp.metadataFile)
	if err != nil {
		if os.IsNotExist(err) {
			return CacheMetadata{
				Entries: make(map[string]CacheEntryMetadata),
				Version: "1.0",
			}, nil
		}
		return metadata, fmt.Errorf("failed to read metadata: %w", err)
	}

	err = json.Unmarshal(data, &metadata)
	if err != nil {
		return metadata, fmt.Errorf("failed to unmarshal metadata: %w", err)
	}

	return metadata, nil
}

// PerformCleanup 执行清理操作
func (cp *CachePersistenceManager) PerformCleanup(currentEntries map[string]*CacheEntry) error {
	metadata, err := cp.LoadMetadata()
	if err != nil {
		return fmt.Errorf("failed to load metadata for cleanup: %w", err)
	}

	// 检查是否需要清理
	if time.Since(metadata.LastCleanup) < cp.strategy.CleanupInterval {
		return nil
	}

	// 收集需要清理的条目
	var cleanupCandidates []CacheEntryMetadata
	totalSize := int64(0)

	for _, entry := range metadata.Entries {
		totalSize += entry.Size

		// 根据策略标记清理候选
		if cp.shouldCleanup(entry) {
			cleanupCandidates = append(cleanupCandidates, entry)
		}
	}

	// 按优先级排序（LRU + LFU组合）
	sort.Slice(cleanupCandidates, func(i, j int) bool {
		a, b := cleanupCandidates[i], cleanupCandidates[j]

		// 过期的条目优先清理
		if a.ExpiresAt.Before(time.Now()) && !b.ExpiresAt.Before(time.Now()) {
			return true
		}
		if !a.ExpiresAt.Before(time.Now()) && b.ExpiresAt.Before(time.Now()) {
			return false
		}

		// 按最后访问时间排序（LRU）
		if a.AccessedAt.Before(b.AccessedAt) {
			return true
		}
		if a.AccessedAt.After(b.AccessedAt) {
			return false
		}

		// 按访问次数排序（LFU）
		return a.AccessCount < b.AccessCount
	})

	// 执行清理
	cleanupCount := int(float64(len(cleanupCandidates)) * cp.strategy.CleanupPercentage)
	if cleanupCount == 0 && len(cleanupCandidates) > 0 {
		cleanupCount = 1
	}

	for i := 0; i < cleanupCount && i < len(cleanupCandidates); i++ {
		entry := cleanupCandidates[i]

		// 从磁盘删除
		dataFile := filepath.Join(cp.dataDir, entry.Key+".cache")
		os.Remove(dataFile)

		// 从元数据删除
		delete(metadata.Entries, entry.Key)

		// 从内存删除
		delete(currentEntries, entry.Key)
	}

	// 更新元数据
	metadata.LastCleanup = time.Now()
	metadata.TotalSize = cp.calculateTotalSize(metadata.Entries)

	return cp.SaveMetadata(metadata)
}

// shouldCleanup 判断是否应该清理条目
func (cp *CachePersistenceManager) shouldCleanup(entry CacheEntryMetadata) bool {
	now := time.Now()

	// 过期条目
	if entry.ExpiresAt.Before(now) {
		return true
	}

	// 超过最大年龄
	if now.Sub(entry.CreatedAt) > cp.strategy.MaxAge {
		return true
	}

	// 访问次数过低
	if entry.AccessCount < cp.strategy.MinAccessCount {
		return true
	}

	// 长时间未访问
	if now.Sub(entry.AccessedAt) > cp.strategy.MaxAge/2 {
		return true
	}

	return false
}

// calculateTotalSize 计算总大小
func (cp *CachePersistenceManager) calculateTotalSize(entries map[string]CacheEntryMetadata) int64 {
	total := int64(0)
	for _, entry := range entries {
		total += entry.Size
	}
	return total
}

// BackupCache 备份缓存
func (cp *CachePersistenceManager) BackupCache(backupDir string) error {
	if !cp.strategy.BackupEnabled {
		return nil
	}

	timestamp := time.Now().Format("20060102_150405")
	backupPath := filepath.Join(backupDir, fmt.Sprintf("cache_backup_%s", timestamp))

	if err := os.MkdirAll(backupPath, 0755); err != nil {
		return fmt.Errorf("failed to create backup directory: %w", err)
	}

	// 复制元数据文件
	metadataBackup := filepath.Join(backupPath, "cache_metadata.json")
	if err := cp.copyFile(cp.metadataFile, metadataBackup); err != nil {
		return fmt.Errorf("failed to backup metadata: %w", err)
	}

	// 复制数据目录
	dataBackup := filepath.Join(backupPath, "data")
	return cp.copyDir(cp.dataDir, dataBackup)
}

// copyFile 复制文件
func (cp *CachePersistenceManager) copyFile(src, dst string) error {
	data, err := ioutil.ReadFile(src)
	if err != nil {
		return err
	}
	return ioutil.WriteFile(dst, data, 0644)
}

// copyDir 复制目录
func (cp *CachePersistenceManager) copyDir(src, dst string) error {
	return filepath.Walk(src, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		relPath, err := filepath.Rel(src, path)
		if err != nil {
			return err
		}

		dstPath := filepath.Join(dst, relPath)

		if info.IsDir() {
			return os.MkdirAll(dstPath, info.Mode())
		}

		return cp.copyFile(path, dstPath)
	})
}

// LoadEntry 从磁盘加载缓存条目
func (cp *CachePersistenceManager) LoadEntry(key string) (interface{}, bool) {
	dataFile := filepath.Join(cp.dataDir, key+".cache")
	file, err := os.Open(dataFile)
	if err != nil {
		return nil, false
	}
	defer file.Close()

	var value interface{}
	decoder := gob.NewDecoder(file)
	if err := decoder.Decode(&value); err != nil {
		return nil, false
	}

	return value, true
}

// GetDiskUsage 获取磁盘使用量
func (cp *CachePersistenceManager) GetDiskUsage() (int64, error) {
	cp.mutex.RLock()
	defer cp.mutex.RUnlock()

	var totalSize int64
	err := filepath.Walk(cp.dataDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return nil // 忽略错误，继续统计
		}
		if !info.IsDir() {
			totalSize += info.Size()
		}
		return nil
	})

	return totalSize, err
}

// GenerateChecksum 生成数据校验和
func (cp *CachePersistenceManager) GenerateChecksum(data interface{}) string {
	// 简单的MD5校验和实现
	content := fmt.Sprintf("%v", data)
	hash := md5.Sum([]byte(content))
	return hex.EncodeToString(hash[:])
}

// VerifyChecksum 验证数据完整性
func (cp *CachePersistenceManager) VerifyChecksum(data interface{}, expectedChecksum string) bool {
	actualChecksum := cp.GenerateChecksum(data)
	return actualChecksum == expectedChecksum
}

// CacheEntry 缓存条目
type CacheEntry struct {
	Key         string      `json:"key"`
	Value       interface{} `json:"value"`
	CreatedAt   time.Time   `json:"created_at"`
	ExpiresAt   time.Time   `json:"expires_at"`
	AccessedAt  time.Time   `json:"accessed_at"`
	AccessCount int64       `json:"access_count"`
	Size        int64       `json:"size_bytes"`
}

// IsExpired 检查是否过期
func (ce *CacheEntry) IsExpired() bool {
	return time.Now().After(ce.ExpiresAt)
}

// Touch 更新访问时间
func (ce *CacheEntry) Touch() {
	ce.AccessedAt = time.Now()
	ce.AccessCount++
}

// CacheManager 缓存管理器
type CacheManager struct {
	entries      map[string]*CacheEntry
	mutex        sync.RWMutex
	maxSize      int64
	currentSize  int64
	ttl          time.Duration
	cacheDir     string
	enableDisk   bool
	errorManager *ErrorManager
	metrics      *CacheMetrics
	persistence  *CachePersistenceManager
	stopCleanup  chan bool
}

// CacheMetrics 缓存指标
type CacheMetrics struct {
	Hits       int64 `json:"hits"`
	Misses     int64 `json:"misses"`
	Evictions  int64 `json:"evictions"`
	DiskReads  int64 `json:"disk_reads"`
	DiskWrites int64 `json:"disk_writes"`
	TotalSize  int64 `json:"total_size_bytes"`
	EntryCount int64 `json:"entry_count"`
}

// NewCacheManager 创建缓存管理器
func NewCacheManager(maxSizeMB int64, ttl time.Duration, cacheDir string, enableDisk bool, errorManager *ErrorManager, strategy CacheCleanupStrategy) *CacheManager {
	cm := &CacheManager{
		entries:      make(map[string]*CacheEntry),
		maxSize:      maxSizeMB * 1024 * 1024, // 转换为字节
		ttl:          ttl,
		cacheDir:     cacheDir,
		enableDisk:   enableDisk,
		errorManager: errorManager,
		metrics:      &CacheMetrics{},
		persistence:  NewCachePersistenceManager(cacheDir, strategy),
		stopCleanup:  make(chan bool),
	}

	if enableDisk && cacheDir != "" {
		if err := os.MkdirAll(cacheDir, 0755); err != nil {
			errorManager.AddError(NewIOError("CACHE_DIR_CREATE_FAILED",
				fmt.Sprintf("Failed to create cache directory: %v", err)).
				WithLocation(cacheDir, 0, 0).Build())
		}
	}

	// 启动清理协程
	go cm.cleanupExpired()

	return cm
}

// Get 获取缓存项
func (cm *CacheManager) Get(key string) (interface{}, bool) {
	cm.mutex.RLock()
	entry, exists := cm.entries[key]
	cm.mutex.RUnlock()

	if !exists {
		// 尝试从磁盘加载
		if cm.enableDisk {
			if value, ok := cm.loadFromDisk(key); ok {
				cm.metrics.DiskReads++
				return value, true
			}
		}
		cm.metrics.Misses++
		return nil, false
	}

	if entry.IsExpired() {
		cm.mutex.Lock()
		delete(cm.entries, key)
		cm.currentSize -= entry.Size
		cm.metrics.EntryCount--
		cm.mutex.Unlock()
		cm.metrics.Misses++
		return nil, false
	}

	entry.Touch()
	cm.metrics.Hits++
	return entry.Value, true
}

// Set 设置缓存项
func (cm *CacheManager) Set(key string, value interface{}, customTTL ...time.Duration) error {
	ttl := cm.ttl
	if len(customTTL) > 0 {
		ttl = customTTL[0]
	}

	// 估算大小
	size := cm.estimateSize(value)

	entry := &CacheEntry{
		Key:         key,
		Value:       value,
		CreatedAt:   time.Now(),
		ExpiresAt:   time.Now().Add(ttl),
		AccessedAt:  time.Now(),
		AccessCount: 1,
		Size:        size,
	}

	cm.mutex.Lock()
	defer cm.mutex.Unlock()

	// 检查是否需要淘汰
	if cm.currentSize+size > cm.maxSize {
		cm.evictLRU(size)
	}

	// 如果存在旧条目，减去其大小
	if oldEntry, exists := cm.entries[key]; exists {
		cm.currentSize -= oldEntry.Size
		cm.metrics.EntryCount--
	}

	cm.entries[key] = entry
	cm.currentSize += size
	cm.metrics.EntryCount++
	cm.metrics.TotalSize = cm.currentSize

	// 写入磁盘
	if cm.enableDisk {
		go cm.saveToDisk(key, value)
	} // 持久化元数据
	go cm.persistence.SaveMetadata(CacheMetadata{
		Entries: map[string]CacheEntryMetadata{
			key: {
				Key:         key,
				CreatedAt:   entry.CreatedAt,
				ExpiresAt:   entry.ExpiresAt,
				AccessedAt:  entry.AccessedAt,
				AccessCount: entry.AccessCount,
				Size:        entry.Size,
				Checksum:    cm.persistence.GenerateChecksum(entry.Value),
				Compressed:  false, // 默认不压缩
			},
		},
		LastCleanup: time.Now(),
		Version:     "1.0",
		TotalSize:   cm.currentSize,
	})

	return nil
}

// Delete 删除缓存项
func (cm *CacheManager) Delete(key string) {
	cm.mutex.Lock()
	defer cm.mutex.Unlock()

	if entry, exists := cm.entries[key]; exists {
		delete(cm.entries, key)
		cm.currentSize -= entry.Size
		cm.metrics.EntryCount--
		cm.metrics.TotalSize = cm.currentSize
	}

	// 删除磁盘文件
	if cm.enableDisk {
		cm.deleteFromDisk(key)
	}
	// 从持久化元数据中删除
	go cm.persistence.SaveMetadata(CacheMetadata{
		Entries:     make(map[string]CacheEntryMetadata), // 空的entries map表示删除
		LastCleanup: time.Now(),
		Version:     "1.0",
		TotalSize:   cm.currentSize,
	})
}

// Clear 清空缓存
func (cm *CacheManager) Clear() {
	cm.mutex.Lock()
	defer cm.mutex.Unlock()

	cm.entries = make(map[string]*CacheEntry)
	cm.currentSize = 0
	cm.metrics.EntryCount = 0
	cm.metrics.TotalSize = 0

	// 清空磁盘缓存
	if cm.enableDisk && cm.cacheDir != "" {
		os.RemoveAll(cm.cacheDir)
		os.MkdirAll(cm.cacheDir, 0755)
	}

	// 清空持久化数据
	go cm.persistence.SaveMetadata(CacheMetadata{
		Entries:     make(map[string]CacheEntryMetadata),
		LastCleanup: time.Now(),
		Version:     "1.0",
		TotalSize:   0,
	})
}

// evictLRU LRU淘汰算法
func (cm *CacheManager) evictLRU(neededSize int64) {
	type entryWithKey struct {
		key   string
		entry *CacheEntry
	}

	var entries []entryWithKey
	for key, entry := range cm.entries {
		entries = append(entries, entryWithKey{key: key, entry: entry})
	}

	// 按访问时间排序
	for i := 0; i < len(entries)-1; i++ {
		for j := i + 1; j < len(entries); j++ {
			if entries[i].entry.AccessedAt.After(entries[j].entry.AccessedAt) {
				entries[i], entries[j] = entries[j], entries[i]
			}
		}
	}

	freedSize := int64(0)
	for _, item := range entries {
		if freedSize >= neededSize {
			break
		}

		delete(cm.entries, item.key)
		cm.currentSize -= item.entry.Size
		freedSize += item.entry.Size
		cm.metrics.Evictions++
		cm.metrics.EntryCount--

		// 删除磁盘文件
		if cm.enableDisk {
			cm.deleteFromDisk(item.key)
		}
	}
}

// cleanupExpired 清理过期项
func (cm *CacheManager) cleanupExpired() {
	ticker := time.NewTicker(5 * time.Minute)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			cm.mutex.Lock()
			var expiredKeys []string
			for key, entry := range cm.entries {
				if entry.IsExpired() {
					expiredKeys = append(expiredKeys, key)
				}
			}

			for _, key := range expiredKeys {
				entry := cm.entries[key]
				delete(cm.entries, key)
				cm.currentSize -= entry.Size
				cm.metrics.EntryCount--

				if cm.enableDisk {
					cm.deleteFromDisk(key)
				}
			}
			cm.metrics.TotalSize = cm.currentSize
			cm.mutex.Unlock()
		case <-cm.stopCleanup:
			return
		}
	}
}

// loadPersistedCache 加载持久化的缓存
func (cm *CacheManager) loadPersistedCache() {
	if cm.persistence == nil {
		return
	}

	metadata, err := cm.persistence.LoadMetadata()
	if err != nil {
		cm.errorManager.AddError(NewIOError("CACHE_LOAD_FAILED",
			fmt.Sprintf("Failed to load cache metadata: %v", err)).Build())
		return
	}

	cm.mutex.Lock()
	defer cm.mutex.Unlock()

	loadedCount := 0
	for key, entryMeta := range metadata.Entries {
		// 检查是否过期
		if entryMeta.ExpiresAt.Before(time.Now()) {
			continue
		}

		// 尝试从磁盘加载值
		if value, ok := cm.persistence.LoadEntry(key); ok {
			entry := &CacheEntry{
				Key:         entryMeta.Key,
				Value:       value,
				CreatedAt:   entryMeta.CreatedAt,
				ExpiresAt:   entryMeta.ExpiresAt,
				AccessedAt:  entryMeta.AccessedAt,
				AccessCount: entryMeta.AccessCount,
				Size:        entryMeta.Size,
			}

			cm.entries[key] = entry
			cm.currentSize += entry.Size
			loadedCount++
		}
	}

	cm.metrics.EntryCount = int64(len(cm.entries))
	cm.metrics.TotalSize = cm.currentSize

	fmt.Printf("Loaded %d cached entries from persistence\n", loadedCount)
}

// performIntelligentCleanup 执行智能清理
func (cm *CacheManager) performIntelligentCleanup() {
	if cm.persistence == nil {
		return
	}

	err := cm.persistence.PerformCleanup(cm.entries)
	if err != nil {
		cm.errorManager.AddError(NewIOError("CACHE_CLEANUP_FAILED",
			fmt.Sprintf("Failed to perform cache cleanup: %v", err)).Build())
		return
	}

	fmt.Println("Cache cleanup completed successfully")

	// 更新内存中的指标
	cm.mutex.Lock()
	cm.currentSize = 0
	for _, entry := range cm.entries {
		cm.currentSize += entry.Size
	}
	cm.metrics.EntryCount = int64(len(cm.entries))
	cm.metrics.TotalSize = cm.currentSize
	cm.mutex.Unlock()
}

// BackupCache 备份缓存
func (cm *CacheManager) BackupCache(backupDir string) error {
	if cm.persistence == nil {
		return fmt.Errorf("persistence not enabled")
	}

	return cm.persistence.BackupCache(backupDir)
}

// GetDiskUsage 获取磁盘使用情况
func (cm *CacheManager) GetDiskUsage() (int64, error) {
	if cm.persistence == nil {
		return 0, fmt.Errorf("persistence not enabled")
	}

	return cm.persistence.GetDiskUsage()
}

// SetCleanupStrategy 设置清理策略
func (cm *CacheManager) SetCleanupStrategy(strategy CacheCleanupStrategy) {
	if cm.persistence != nil {
		cm.persistence.strategy = strategy
	}
}

// GetCleanupStrategy 获取当前清理策略
func (cm *CacheManager) GetCleanupStrategy() CacheCleanupStrategy {
	if cm.persistence != nil {
		return cm.persistence.strategy
	}
	return DefaultCleanupStrategy()
}

// Shutdown 优雅关闭缓存管理器
func (cm *CacheManager) Shutdown() error {
	// 停止清理协程
	if cm.stopCleanup != nil {
		close(cm.stopCleanup)
	}

	// 执行最后一次清理
	cm.performIntelligentCleanup()

	// 保存当前状态到持久化存储
	if cm.persistence != nil {
		cm.mutex.RLock()
		entries := make(map[string]CacheEntryMetadata)
		for key, entry := range cm.entries {
			entries[key] = CacheEntryMetadata{
				Key:         entry.Key,
				CreatedAt:   entry.CreatedAt,
				ExpiresAt:   entry.ExpiresAt,
				AccessedAt:  entry.AccessedAt,
				AccessCount: entry.AccessCount,
				Size:        entry.Size,
				Checksum:    cm.persistence.GenerateChecksum(entry.Value),
				Compressed:  false, // 默认不压缩
			}
		}
		cm.mutex.RUnlock()

		metadata := CacheMetadata{
			Entries:     entries,
			LastCleanup: time.Now(),
			Version:     "1.0",
			TotalSize:   cm.currentSize,
		}

		return cm.persistence.SaveMetadata(metadata)
	}

	return nil
}

// estimateSize 估算对象大小
func (cm *CacheManager) estimateSize(value interface{}) int64 {
	// 简单的大小估算，实际应用中可以使用更精确的方法
	switch v := value.(type) {
	case string:
		return int64(len(v))
	case []byte:
		return int64(len(v))
	case map[string]interface{}:
		return int64(len(v) * 64) // 估算
	default:
		return 1024 // 默认1KB
	}
}

// generateCacheKey 生成缓存键
func (cm *CacheManager) generateCacheKey(data string) string {
	hash := md5.Sum([]byte(data))
	return hex.EncodeToString(hash[:])
}

// saveToDisk 保存到磁盘
func (cm *CacheManager) saveToDisk(key string, value interface{}) {
	if cm.cacheDir == "" {
		return
	}

	filePath := filepath.Join(cm.cacheDir, key+".cache")
	file, err := os.Create(filePath)
	if err != nil {
		cm.errorManager.AddError(NewIOError("CACHE_DISK_WRITE_FAILED",
			fmt.Sprintf("Failed to create cache file: %v", err)).
			WithLocation(filePath, 0, 0).Build())
		return
	}
	defer file.Close()

	encoder := gob.NewEncoder(file)
	if err := encoder.Encode(value); err != nil {
		cm.errorManager.AddError(NewIOError("CACHE_ENCODE_FAILED",
			fmt.Sprintf("Failed to encode cache value: %v", err)).Build())
		return
	}

	cm.metrics.DiskWrites++
}

// loadFromDisk 从磁盘加载
func (cm *CacheManager) loadFromDisk(key string) (interface{}, bool) {
	if cm.cacheDir == "" {
		return nil, false
	}

	filePath := filepath.Join(cm.cacheDir, key+".cache")
	file, err := os.Open(filePath)
	if err != nil {
		return nil, false
	}
	defer file.Close()

	var value interface{}
	decoder := gob.NewDecoder(file)
	if err := decoder.Decode(&value); err != nil {
		return nil, false
	}

	return value, true
}

// deleteFromDisk 从磁盘删除
func (cm *CacheManager) deleteFromDisk(key string) {
	if cm.cacheDir == "" {
		return
	}

	filePath := filepath.Join(cm.cacheDir, key+".cache")
	os.Remove(filePath)
}

// GetMetrics 获取缓存指标
func (cm *CacheManager) GetMetrics() CacheMetrics {
	cm.mutex.RLock()
	defer cm.mutex.RUnlock()

	metrics := *cm.metrics
	metrics.TotalSize = cm.currentSize
	metrics.EntryCount = int64(len(cm.entries))
	return metrics
}

// GetStats 获取缓存统计信息
func (cm *CacheManager) GetStats() CacheStats {
	metrics := cm.GetMetrics()
	total := metrics.Hits + metrics.Misses
	hitRate := float64(0)
	if total > 0 {
		hitRate = float64(metrics.Hits) / float64(total) * 100
	}

	return CacheStats{
		HitRate:     hitRate,
		TotalOps:    total,
		CurrentSize: metrics.TotalSize,
		MaxSize:     cm.maxSize,
		UsageRatio:  float64(metrics.TotalSize) / float64(cm.maxSize) * 100,
		EntryCount:  metrics.EntryCount,
		Metrics:     metrics,
	}
}

// CacheStats 缓存统计
type CacheStats struct {
	HitRate     float64      `json:"hit_rate_percent"`
	TotalOps    int64        `json:"total_operations"`
	CurrentSize int64        `json:"current_size_bytes"`
	MaxSize     int64        `json:"max_size_bytes"`
	UsageRatio  float64      `json:"usage_ratio_percent"`
	EntryCount  int64        `json:"entry_count"`
	Metrics     CacheMetrics `json:"detailed_metrics"`
}

// String 格式化输出缓存统计
func (cs CacheStats) String() string {
	return fmt.Sprintf("Cache Stats: Hit Rate=%.2f%%, Usage=%.2f%% (%d/%d MB), Entries=%d",
		cs.HitRate,
		cs.UsageRatio,
		cs.CurrentSize/1024/1024,
		cs.MaxSize/1024/1024,
		cs.EntryCount)
}
