package gocache

import (
	"time"
)

// RedisCache https://github.com/go-redis/redis
type RedisCache struct {
	db              *Redis           // 驱动
	expiration      time.Duration    // 默认过期时间
	GetterString    GttStringFunc    // 不存在的操作
	GetterInterface GttInterfaceFunc // 不存在的操作
}

// NewCache 返回Redis缓存实例
func (r *Redis) NewCache(expiration time.Duration) *RedisCache {
	return &RedisCache{
		db:         r,          // 操作类
		expiration: expiration, // 过期时间
	}
}

func (rc *RedisCache) GetInterface(key string) (ret string) {

	f := func() interface{} {
		return rc.GetterInterface()
	}

	ret, err := rc.db.Get(key)

	if err != nil {
		rc.db.Set(key, f, rc.expiration)
		ret, _ = rc.db.Get(key)
	}

	return
}

func (rc *RedisCache) GetString(key string) (ret string) {

	f := func() string {
		return rc.GetterString()
	}

	ret, err := rc.db.Get(key)
	if err != nil {
		rc.db.Set(key, f, rc.expiration)
		ret = f()
	}

	return
}
