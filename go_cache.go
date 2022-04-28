package gocache

import "time"

// GoCache https://github.com/patrickmn/go-cache
type GoCache struct {
	db              *Go              // 驱动
	expiration      time.Duration    // 默认过期时间
	GetterInterface GttInterfaceFunc // 不存在的操作
}

// NewCache 返回GoCache缓存实例
func (g *Go) NewCache(expiration time.Duration) *GoCache {
	return &GoCache{
		db:         g,          // 操作类
		expiration: expiration, // 过期时间
	}
}

func (gc *GoCache) GetInterface(key string) (ret interface{}) {

	f := func() interface{} {
		return gc.GetterInterface()
	}

	ret, found := gc.db.Get(key)

	if found == false {
		gc.db.Set(key, f, gc.expiration)
		ret, _ = gc.db.Get(key)
	}

	return
}
