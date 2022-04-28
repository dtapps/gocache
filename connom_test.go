package gocache

import (
	"testing"
	"time"
)

func TestBig(t *testing.T) {

	// 实例
	newBig := NewBig(time.Minute * 30)

	// 字符串
	newBig.Set("key1", "test1")
	key1, _ := newBig.Get("key1")
	t.Logf("key1：%+v", key1)

	// 结构体
	type name struct {
		Test string `json:"test"`
	}
	newBig.Set("key2", name{"test2"})
	key2, _ := newBig.Get("key2")
	t.Logf("key2：%+v", key2)

	// 缓存组件
	newBigCache := newBig.NewCache()
	newBigCache.GetterInterface = func() interface{} {
		return name{"test3"}
	}
	key3 := newBigCache.GetInterface("key3")
	t.Logf("key3：%+v", key3)
}

func TestGo(t *testing.T) {

	// 实例
	newGo := NewGo(5*time.Minute, 10*time.Minute)

	// 字符串
	newGo.SetDefault("key1", "test1")
	key1, _ := newGo.Get("key1")
	t.Logf("key1：%+v", key1)

	// 结构体
	type name struct {
		Test string `json:"test"`
	}
	newGo.SetDefault("key2", name{"test2"})
	key2, _ := newGo.Get("key2")
	t.Logf("key2：%+v", key2)

	// 缓存组件
	bigCache := newGo.NewCache(5 * time.Minute)
	bigCache.GetterInterface = func() interface{} {
		return name{"test3"}
	}
	key3 := bigCache.GetInterface("key3")
	t.Logf("key3：%+v", key3)
}
