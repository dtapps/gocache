package gocache

import (
	"testing"
	"time"
)

func TestBig(t *testing.T) {

	newCache := NewBig(time.Minute * 30)

	// 字符串
	newCache.Set("key1", "测试Big插入数据 1")
	key1, _ := newCache.Get("key1")
	t.Logf("key1：%+v", key1)

	// 结构体
	type name struct {
		Test string `json:"test"`
	}
	newCache.Set("key2", name{"测试Big插入数据 2"})
	key2, _ := newCache.Get("key2")
	t.Logf("key2：%+v", key2)
	t.Logf("key2：%+v", key2.(name))

	// 缓存组件
	newCacheCache := newCache.NewCache()
	newCacheCache.GetterInterface = func() interface{} {
		return name{"测试Big插入数据 3"}
	}
	key3 := newCacheCache.GetInterface("key3")
	t.Logf("key3：%+v", key3)
	t.Logf("key3：%+v", key3.(name))

}
