package gocache

import (
	"github.com/dgraph-io/ristretto"
)

// Ristretto https://github.com/dgraph-io/ristretto
type Ristretto struct {
	ristretto   *ristretto.Cache // 驱动
	numCounters int64            // 跟踪频率的键数 (10M)
	maxCost     int64            // 缓存的最大成本（1GB）
	bufferItems int64            // 每个Get缓冲区的密钥数
}

// NewRistretto 实例化
func NewRistretto() *Ristretto {
	cache, _ := ristretto.NewCache(&ristretto.Config{
		NumCounters: 1e7,     // number of keys to track frequency of (10M).
		MaxCost:     1 << 30, // maximum cost of cache (1GB).
		BufferItems: 64,      // number of keys per Get buffer.
	})
	return &Ristretto{ristretto: cache}
}

// Get 获取单个数据
func (r *Ristretto) Get(key string) (interface{}, bool) {
	return r.ristretto.Get(key)
}

// Set 插入数据
func (r *Ristretto) Set(key string, value interface{}, cost int64) {
	r.ristretto.Set(key, value, cost)
	r.ristretto.Wait()
}

// Del 删除数据
func (r *Ristretto) Del(key string) {
	r.ristretto.Del(key)
}
