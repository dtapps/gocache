package gocache

import (
	"time"
)

// GoCache https://github.com/patrickmn/go-cache
type GoCache struct {
	db              *Go              // 驱动
	expiration      time.Duration    // 默认过期时间
	GetterInterface GttInterfaceFunc // 不存在的操作
}

// NewCache 实例化
func (g *Go) NewCache(expiration time.Duration) *GoCache {
	return &GoCache{db: g, expiration: expiration}
}

// GetInterface 缓存操作
func (gc *GoCache) GetInterface(key string) (ret interface{}) {

	f := func() interface{} {
		return gc.GetterInterface()
	}

	// 如果不存在，则调用GetterInterface
	ret, found := gc.db.Get(key)

	if found == false {
		gc.db.Set(key, f(), gc.expiration)
		ret, _ = gc.db.Get(key)
	}

	return
}
