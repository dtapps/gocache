package gocache

import (
	"testing"
	"time"
)

func logGo() *Go {
	return NewGo(&GoConfig{
		DefaultExpiration: time.Minute * 5,
		DefaultClear:      time.Minute * 10,
	})
}

func TestGo(t *testing.T) {

	newCache := logGo()

	// 字符串
	newCache.SetDefault("key1", "测试Go插入数据 1")
	key1, _ := newCache.Get("key1")
	t.Logf("key1：%+v", key1)

	// 结构体
	type name struct {
		Test string `json:"test"`
	}
	newCache.SetDefault("key2", name{"测试Go插入数据 2"})
	key2, _ := newCache.Get("key2")
	t.Logf("key2：%+v", key2)
	t.Logf("key2：%+v", key2.(name))

	// 缓存组件
	newCacheCache := newCache.NewCache(&GoCacheConfig{expiration: 5 * time.Minute})
	newCacheCache.GetterInterface = func() interface{} {
		return name{"测试Go插入数据 3"}
	}
	key3 := newCacheCache.GetInterface("key3")
	t.Logf("key3：%+v", key3)
	t.Logf("key3：%+v", key3.(name))

}
