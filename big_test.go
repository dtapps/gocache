package gocache

import (
	"testing"
	"time"
)

func TestBig(t *testing.T) {

	// 实例
	newBig := NewBig(time.Minute * 30)

	// 字符串
	newBig.Set("key1", "测试Big插入数据1")
	key1, _ := newBig.Get("key1")
	t.Logf("key1：%+v", key1)

	// 结构体
	type name struct {
		Test string `json:"test"`
	}
	newBig.Set("key2", name{"测试Big插入数据2"})
	key2, _ := newBig.Get("key2")
	t.Logf("key2：%+v", key2)

}
