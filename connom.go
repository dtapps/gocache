package gocache

import "time"

const Version = "1.0.4"

var (
	DefaultExpiration = time.Minute * 30 // 默认过期时间
)

// GttStringFunc String缓存结构
type GttStringFunc func() string

// GttInterfaceFunc Interface缓存结构
type GttInterfaceFunc func() interface{}
