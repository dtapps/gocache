package gocache

import (
	"github.com/allegro/bigcache/v3"
)

// BigCache https://github.com/allegro/bigcache
type BigCache struct {
	db              *Big             // 驱动
	GetterInterface GttInterfaceFunc // 不存在的操作
}

// NewCache 返回BigCache实例
func (b *Big) NewCache() *BigCache {
	return &BigCache{
		db: b, // 操作类
	}
}

func (gc *BigCache) GetInterface(key string) (ret interface{}) {

	f := func() interface{} {
		return gc.GetterInterface()
	}

	ret, err := gc.db.Get(key)

	if err == bigcache.ErrEntryNotFound {
		_ = gc.db.Set(key, f())
		ret, _ = gc.db.Get(key)
	}
	return
}
