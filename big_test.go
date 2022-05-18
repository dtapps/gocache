package gocache

import (
	"testing"
	"time"
)

func logBig() *Big {
	return NewBig(&BigConfig{
		DefaultExpiration: time.Minute * 30,
	})
}

func TestBig(t *testing.T) {

	newCache := logBig()

	// 设置字符串
	t.Logf("设置字符串：%+v", newCache.Set("key1", "测试Big插入数据 1"))
	key1, _ := newCache.Get("key1")
	t.Logf("读取字符串：%+v", key1)

	// 设置结构体
	type name struct {
		Test string `json:"test"`
	}
	t.Logf("设置结构体：%+v", newCache.Set("key2", name{"测试Big插入数据 2"}))
	key2, _ := newCache.Get("key2")
	t.Logf("读取结构体：%+v", key2)
	t.Logf("读取结构体结果：%+v", key2.(name))

	// 缓存组件
	newCacheCache := newCache.NewCache()
	newCacheCache.GetterInterface = func() interface{} {
		return name{"测试Big插入数据 3"}
	}
	//key3Result := &name{}
	key3 := newCacheCache.GetInterface("key3")
	t.Logf("读取缓存组件：%T", key3)
	t.Logf("读取缓存组件：%+v", key3.(name))

}
