package gocache

import (
	"github.com/patrickmn/go-cache"
	"time"
)

// Go https://github.com/patrickmn/go-cache
type Go struct {
	goCache    *cache.Cache  // 驱动
	expiration time.Duration // 默认过期时间
	clear      time.Duration // 清理过期数据
}

// NewGo 返回GoCache实例
func NewGo(expiration, clear time.Duration) *Go {
	c := cache.New(expiration, clear)
	return &Go{
		goCache:    c,
		expiration: expiration,
		clear:      clear,
	}
}

// Get 获取单个数据
func (g *Go) Get(key string) (interface{}, bool) {
	return g.goCache.Get(key)
}

// Set 插入数据 并设置过期时间
func (g *Go) Set(key string, value interface{}, expirationTime time.Duration) {
	g.goCache.Set(key, value, expirationTime)
}

// SetDefault 插入数据 并设置为默认过期时间
func (g *Go) SetDefault(key string, value interface{}) {
	g.goCache.Set(key, value, g.expiration)
}
