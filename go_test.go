package gocache

import (
	"testing"
	"time"
)

func TestGo(t *testing.T) {

	// 实例
	newGo := NewGo(5*time.Minute, 10*time.Minute)

	// 字符串
	newGo.SetDefault("key1", "测试Go插入数据1")
	key1, _ := newGo.Get("key1")
	t.Logf("key1：%+v", key1)

	// 结构体
	type name struct {
		Test string `json:"test"`
	}
	newGo.SetDefault("key2", name{"测试Go插入数据2"})
	key2, _ := newGo.Get("key2")
	t.Logf("key2：%+v", key2)

}
