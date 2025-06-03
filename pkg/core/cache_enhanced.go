package core

import (
	"compress/gzip"
	"crypto/md5"
	"encoding/gob"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"sort"
	"sync"
	"time"
)

// CachePolicy, CacheCleanupStrategy, DefaultCleanupStrategy, CacheMetadata,
// and CacheEntryMetadata are declared in cache.go to avoid duplication

// CachePersistence 缓存持久化管理器
type CachePersistence struct {
	metadataFile string
	dataDir      string
	strategy     CacheCleanupStrategy
	mutex        sync.RWMutex
}

// NewCachePersistence 创建缓存持久化管理器
func NewCachePersistence(baseDir string, strategy CacheCleanupStrategy) *CachePersistence {
	persistence := &CachePersistence{
		metadataFile: filepath.Join(baseDir, "cache_metadata.json"),
		dataDir:      filepath.Join(baseDir, "data"),
		strategy:     strategy,
	}

	// 确保目录存在
	os.MkdirAll(persistence.dataDir, 0755)

	return persistence
}

// SaveMetadata 保存元数据
func (cp *CachePersistence) SaveMetadata(metadata CacheMetadata) error {
	cp.mutex.Lock()
	defer cp.mutex.Unlock()

	data, err := json.MarshalIndent(metadata, "", "  ")
	if err != nil {
		return fmt.Errorf("failed to marshal metadata: %w", err)
	}

	// 原子写入 - 先写临时文件再重命名
	tempFile := cp.metadataFile + ".tmp"
	if err := ioutil.WriteFile(tempFile, data, 0644); err != nil {
		return fmt.Errorf("failed to write temp metadata: %w", err)
	}

	return os.Rename(tempFile, cp.metadataFile)
}

// LoadMetadata 加载元数据
func (cp *CachePersistence) LoadMetadata() (CacheMetadata, error) {
	cp.mutex.RLock()
	defer cp.mutex.RUnlock()

	var metadata CacheMetadata
	data, err := ioutil.ReadFile(cp.metadataFile)
	if err != nil {
		if os.IsNotExist(err) {
			return CacheMetadata{
				Entries:     make(map[string]CacheEntryMetadata),
				Version:     "1.0",
				LastCleanup: time.Now(),
			}, nil
		}
		return metadata, fmt.Errorf("failed to read metadata: %w", err)
	}

	err = json.Unmarshal(data, &metadata)
	if err != nil {
		return metadata, fmt.Errorf("failed to unmarshal metadata: %w", err)
	}

	// 如果元数据中没有条目映射，初始化它
	if metadata.Entries == nil {
		metadata.Entries = make(map[string]CacheEntryMetadata)
	}

	return metadata, nil
}

// PerformCleanup 执行清理操作
func (cp *CachePersistence) PerformCleanup(currentEntries map[string]*CacheEntry) (int, error) {
	metadata, err := cp.LoadMetadata()
	if err != nil {
		return 0, fmt.Errorf("failed to load metadata for cleanup: %w", err)
	}

	// 检查是否需要清理
	if time.Since(metadata.LastCleanup) < cp.strategy.CleanupInterval {
		return 0, nil
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

	// 按优先级排序（过期 > LRU > LFU）
	sort.Slice(cleanupCandidates, func(i, j int) bool {
		a, b := cleanupCandidates[i], cleanupCandidates[j]

		// 过期的条目优先清理
		aExpired := a.ExpiresAt.Before(time.Now())
		bExpired := b.ExpiresAt.Before(time.Now())

		if aExpired && !bExpired {
			return true
		}
		if !aExpired && bExpired {
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

	cleaned := 0
	for i := 0; i < cleanupCount && i < len(cleanupCandidates); i++ {
		entry := cleanupCandidates[i]

		// 从磁盘删除
		dataFile := filepath.Join(cp.dataDir, entry.Key+".cache")
		if cp.strategy.CompressionEnabled {
			dataFile += ".gz"
		}
		os.Remove(dataFile)

		// 从元数据删除
		delete(metadata.Entries, entry.Key)

		// 从内存删除
		delete(currentEntries, entry.Key)

		cleaned++
	}

	// 更新元数据
	metadata.LastCleanup = time.Now()
	metadata.TotalSize = cp.calculateTotalSize(metadata.Entries)

	if err := cp.SaveMetadata(metadata); err != nil {
		return cleaned, fmt.Errorf("failed to save metadata after cleanup: %w", err)
	}

	return cleaned, nil
}

// shouldCleanup 判断是否应该清理条目
func (cp *CachePersistence) shouldCleanup(entry CacheEntryMetadata) bool {
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
func (cp *CachePersistence) calculateTotalSize(entries map[string]CacheEntryMetadata) int64 {
	total := int64(0)
	for _, entry := range entries {
		total += entry.Size
	}
	return total
}

// BackupCache 备份缓存
func (cp *CachePersistence) BackupCache(backupDir string) error {
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
func (cp *CachePersistence) copyFile(src, dst string) error {
	data, err := ioutil.ReadFile(src)
	if err != nil {
		return err
	}
	return ioutil.WriteFile(dst, data, 0644)
}

// copyDir 复制目录
func (cp *CachePersistence) copyDir(src, dst string) error {
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

// SaveEntry 保存单个条目到磁盘
func (cp *CachePersistence) SaveEntry(key string, value interface{}, metadata CacheEntryMetadata) error {
	filePath := filepath.Join(cp.dataDir, key+".cache")

	// 创建文件
	file, err := os.Create(filePath)
	if err != nil {
		return fmt.Errorf("failed to create cache file: %w", err)
	}
	defer file.Close()

	var writer io.Writer = file

	// 如果启用压缩
	if cp.strategy.CompressionEnabled {
		filePath += ".gz"
		gzWriter := gzip.NewWriter(file)
		defer gzWriter.Close()
		writer = gzWriter
	}

	// 编码数据
	encoder := gob.NewEncoder(writer)
	if err := encoder.Encode(value); err != nil {
		return fmt.Errorf("failed to encode cache value: %w", err)
	}

	// 如果启用压缩，确保数据被刷新
	if cp.strategy.CompressionEnabled {
		if gzWriter, ok := writer.(*gzip.Writer); ok {
			gzWriter.Close()
		}
	}

	return nil
}

// LoadEntry 从磁盘加载单个条目
func (cp *CachePersistence) LoadEntry(key string) (interface{}, bool) {
	filePath := filepath.Join(cp.dataDir, key+".cache")

	// 尝试压缩版本
	if cp.strategy.CompressionEnabled {
		if compressed, ok := cp.loadCompressedEntry(filePath + ".gz"); ok {
			return compressed, true
		}
	}

	// 尝试未压缩版本
	return cp.loadUncompressedEntry(filePath)
}

// loadCompressedEntry 加载压缩条目
func (cp *CachePersistence) loadCompressedEntry(filePath string) (interface{}, bool) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, false
	}
	defer file.Close()

	gzReader, err := gzip.NewReader(file)
	if err != nil {
		return nil, false
	}
	defer gzReader.Close()

	var value interface{}
	decoder := gob.NewDecoder(gzReader)
	if err := decoder.Decode(&value); err != nil {
		return nil, false
	}

	return value, true
}

// loadUncompressedEntry 加载未压缩条目
func (cp *CachePersistence) loadUncompressedEntry(filePath string) (interface{}, bool) {
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

// DeleteEntry 删除磁盘条目
func (cp *CachePersistence) DeleteEntry(key string) {
	// 删除未压缩版本
	uncompressed := filepath.Join(cp.dataDir, key+".cache")
	os.Remove(uncompressed)

	// 删除压缩版本
	compressed := uncompressed + ".gz"
	os.Remove(compressed)
}

// GetDiskUsage 获取磁盘使用情况
func (cp *CachePersistence) GetDiskUsage() (int64, error) {
	var totalSize int64

	err := filepath.Walk(cp.dataDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() {
			totalSize += info.Size()
		}
		return nil
	})

	return totalSize, err
}

// GenerateChecksum 生成数据校验和
func (cp *CachePersistence) GenerateChecksum(data interface{}) string {
	// 简单的MD5校验和实现
	content := fmt.Sprintf("%v", data)
	hash := md5.Sum([]byte(content))
	return hex.EncodeToString(hash[:])
}

// VerifyChecksum 验证数据完整性
func (cp *CachePersistence) VerifyChecksum(data interface{}, expectedChecksum string) bool {
	actualChecksum := cp.GenerateChecksum(data)
	return actualChecksum == expectedChecksum
}
