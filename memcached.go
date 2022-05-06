package gocache

import (
	"github.com/bradfitz/gomemcache/memcache"
)

// Memcached https://github.com/bradfitz/gomemcache
type Memcached struct {
	db *memcache.Client // 驱动
}

// NewMemcached 实例化
func NewMemcached(dns string) *Memcached {
	mc := memcache.New(dns)
	if mc == nil {
		panic("连接失败")
	}
	return &Memcached{db: mc}
}

// NewMemcachedDb 实例化
func NewMemcachedDb(memcached *memcache.Client) *Memcached {
	return &Memcached{db: memcached}
}

// Set 插入数据
func (m *Memcached) Set(key string, value []byte) error {
	return m.db.Set(&memcache.Item{Key: key, Value: value})
}

// Get 获取单个数据
func (m *Memcached) Get(key string) (string, error) {
	it, err := m.db.Get(key)
	if err == memcache.ErrCacheMiss {
		return "", memcache.ErrCacheMiss
	}
	if it.Key == key {
		return string(it.Value), nil
	}
	return "", memcache.ErrCacheMiss
}

// Del 删除单个数据
func (m *Memcached) Del(key string) error {
	return m.db.Delete(key)
}
