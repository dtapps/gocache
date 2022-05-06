package gocache

import (
	"testing"
	"time"
)

func TestBigCache(t *testing.T) {

	// 实例
	newBig := NewBig(time.Minute * 30)

	// 结构体
	type name struct {
		Test string `json:"test"`
	}

	// 缓存组件
	newBigCache := newBig.NewCache()
	newBigCache.GetterInterface = func() interface{} {
		return name{"测试Big插入数据3"}
	}
	key3 := newBigCache.GetInterface("key3")
	t.Logf("key3：%+v", key3)

}
