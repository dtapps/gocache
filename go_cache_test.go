package gocache

import (
	"testing"
	"time"
)

func TestGoCache(t *testing.T) {

	// 实例
	newGo := NewGo(5*time.Minute, 10*time.Minute)

	// 结构体
	type name struct {
		Test string `json:"test"`
	}
	// 缓存组件
	bigCache := newGo.NewCache(5 * time.Minute)
	bigCache.GetterInterface = func() interface{} {
		return name{"测试Go插入数据3"}
	}
	key3 := bigCache.GetInterface("key3")
	t.Logf("key3：%+v", key3)

}
