package gocache

import (
	"testing"
)

func TestRistretto(t *testing.T) {

	newCache := NewRistretto()

	// 字符串
	newCache.Set("key1", "测试Ristretto插入数据1", 1)
	key1, _ := newCache.Get("key1")
	t.Logf("key1：%+v", key1)
	newCache.Del("key1")
	key1, _ = newCache.Get("key1")
	t.Logf("key1：%+v", key1)

	// 结构体
	type name struct {
		Test string `json:"test"`
	}
	newCache.Set("key2", name{"测试Ristretto插入数据2"}, 1)
	key2, _ := newCache.Get("key2")
	t.Logf("key2：%+v", key2)
	newCache.Del("key2")
	key2, _ = newCache.Get("key2")
	t.Logf("key2：%+v", key2)

	//

	//	// 缓存组件
	//	newBigCache := newBig.NewCache()
	//	newBigCache.GetterInterface = func() interface{} {
	//		return name{"测试Big插入数据3"}
	//	}
	//	key3 := newBigCache.GetInterface("key3")
	//	t.Logf("key3：%+v", key3)
	//
}
