package gocache

import (
	"encoding/json"
	"time"
)

// RedisCache https://github.com/go-redis/redis
type RedisCache struct {
	db              *Redis           // 驱动
	expiration      time.Duration    // 默认过期时间
	GetterString    GttStringFunc    // 不存在的操作
	GetterInterface GttInterfaceFunc // 不存在的操作
}

// NewCache 实例化
func (r *Redis) NewCache(expiration time.Duration) *RedisCache {
	return &RedisCache{db: r, expiration: expiration}
}

// NewCacheDefaultExpiration 实例化
func (r *Redis) NewCacheDefaultExpiration() *RedisCache {
	return &RedisCache{db: r, expiration: r.expiration}
}

// GetString 缓存操作
func (rc *RedisCache) GetString(key string) (ret string) {

	f := func() string {
		return rc.GetterString()
	}

	// 如果不存在，则调用GetterString
	ret, err := rc.db.Get(key)
	if err != nil {
		rc.db.Set(key, f(), rc.expiration)
		ret, _ = rc.db.Get(key)
	}

	return
}

// GetInterface 缓存操作
func (rc *RedisCache) GetInterface(key string, result interface{}) {

	f := func() string {
		marshal, _ := json.Marshal(rc.GetterInterface())
		return string(marshal)
	}

	// 如果不存在，则调用GetterInterface
	ret, err := rc.db.Get(key)

	if err != nil {
		rc.db.Set(key, f(), rc.expiration)
		ret, _ = rc.db.Get(key)
	}

	err = json.Unmarshal([]byte(ret), result)

	return
}
