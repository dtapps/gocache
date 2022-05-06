package gocache

import (
	"context"
	"errors"
	"fmt"
	"github.com/go-redis/redis/v8"
	"time"
)

// Redis https://github.com/go-redis/redis
type Redis struct {
	db         *redis.Client   // 驱动
	ctx        context.Context // 上下文内容
	expiration time.Duration   // 默认过期时间
}

// NewRedis 实例化
func NewRedis(addr, password string, dbName int, expiration time.Duration) *Redis {
	db := redis.NewClient(&redis.Options{
		Addr:     addr,     // 地址
		Password: password, // 密码
		DB:       dbName,   // 数据库
		PoolSize: 100,      // 连接池大小
	})

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := db.Ping(ctx).Result()
	if err != nil {
		panic(errors.New(fmt.Sprintf("连接失败，%s", err.Error())))
	}

	return &Redis{db: db, ctx: context.Background(), expiration: expiration}
}

// NewRedisDb 实例化
func NewRedisDb(db *redis.Client, expiration time.Duration) *Redis {
	return &Redis{db: db, ctx: context.Background(), expiration: expiration}
}

// Set 设置一个key的值
func (r *Redis) Set(key string, value interface{}, expiration time.Duration) (string, error) {
	return r.db.Set(r.ctx, key, value, expiration).Result()
}

// SetDefaultExpiration 设置一个key的值，使用全局默认过期时间
func (r *Redis) SetDefaultExpiration(key string, value interface{}) (string, error) {
	return r.db.Set(r.ctx, key, value, r.expiration).Result()
}

// Get 查询key的值
func (r *Redis) Get(key string) (string, error) {
	return r.db.Get(r.ctx, key).Result()
}

// GetSet 设置一个key的值，并返回这个key的旧值
func (r *Redis) GetSet(key string, value interface{}) (string, error) {
	return r.db.GetSet(r.ctx, key, value).Result()
}

// SetNX 如果key不存在，则设置这个key的值
func (r *Redis) SetNX(key string, value interface{}, expiration time.Duration) error {
	return r.db.SetNX(r.ctx, key, value, expiration).Err()
}

// SetNXDefaultExpiration 如果key不存在，则设置这个key的值，使用全局默认过期时间
func (r *Redis) SetNXDefaultExpiration(key string, value interface{}) error {
	return r.db.SetNX(r.ctx, key, value, r.expiration).Err()
}

// MGet 批量查询key的值
func (r *Redis) MGet(keys ...string) ([]interface{}, error) {
	return r.db.MGet(r.ctx, keys...).Result()
}

// MSet 批量设置key的值
// MSet(map[string]interface{}{"key1": "value1", "key2": "value2"})
func (r *Redis) MSet(values map[string]interface{}) error {
	return r.db.MSet(r.ctx, values).Err()
}

// Incr 针对一个key的数值进行递增操作
func (r *Redis) Incr(key string) (int64, error) {
	return r.db.Incr(r.ctx, key).Result()
}

// IncrBy 针对一个key的数值进行递增操作，指定每次递增多少
func (r *Redis) IncrBy(key string, value int64) (int64, error) {
	return r.db.IncrBy(r.ctx, key, value).Result()
}

// Decr 针对一个key的数值进行递减操作
func (r *Redis) Decr(key string) (int64, error) {
	return r.db.Decr(r.ctx, key).Result()
}

// DecrBy 针对一个key的数值进行递减操作，指定每次递减多少
func (r *Redis) DecrBy(key string, value int64) (int64, error) {
	return r.db.DecrBy(r.ctx, key, value).Result()
}

// Del 删除key操作，支持批量删除
func (r *Redis) Del(keys ...string) error {
	return r.db.Del(r.ctx, keys...).Err()
}
