package gocache

import (
	"bytes"
	"encoding/gob"
	"github.com/allegro/bigcache/v3"
	"time"
)

// BigConfig 配置
type BigConfig struct {
	DefaultExpiration time.Duration // 默认过期时间
}

// Big https://github.com/allegro/bigcache
type Big struct {
	BigConfig
	db *bigcache.BigCache // 驱动
}

// NewBig 实例化
func NewBig(config *BigConfig) *Big {
	b := &Big{}
	b.DefaultExpiration = config.DefaultExpiration
	b.db, _ = bigcache.NewBigCache(bigcache.DefaultConfig(b.DefaultExpiration))
	return b
}

// Set 插入数据 将只显示给定结构的导出字段 序列化并存储
func (b *Big) Set(key string, value interface{}) error {

	// 将 value 序列化为 bytes
	valueBytes, err := serialize(value)
	if err != nil {
		return err
	}

	return b.db.Set(key, valueBytes)
}

// Get 获取单个数据
func (b *Big) Get(key string) (interface{}, error) {

	// 获取以 bytes 格式存储的 value
	valueBytes, err := b.db.Get(key)
	if err != nil {
		return nil, err
	}

	// 反序列化 valueBytes
	value, err := deserialize(valueBytes)
	if err != nil {
		return nil, err
	}

	return value, nil
}

// 序列化
func serialize(value interface{}) ([]byte, error) {
	buf := bytes.Buffer{}
	enc := gob.NewEncoder(&buf)
	gob.Register(value)

	err := enc.Encode(&value)
	if err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

// 反序列化
func deserialize(valueBytes []byte) (interface{}, error) {
	var value interface{}
	buf := bytes.NewBuffer(valueBytes)
	dec := gob.NewDecoder(buf)

	err := dec.Decode(&value)
	if err != nil {
		return nil, err
	}

	return value, nil
}
