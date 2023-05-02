package common

import (
	"context"
	"encoding/json"
	"github.com/redis/go-redis/v9"
	"log"
	"math/rand"
	"sync"
	"time"
)

// 重构点 新增Redis板块 对频繁写入读取的数据进行缓存

var (
	RedisClient *redis.Client
	ctx         = context.Background()
	// 键值对过期时间
	valueExpire = time.Hour
	mutex       sync.Mutex
)

func InitRedis() {
	conf := GetConfig()
	ip := conf.Redis.Host
	port := conf.Redis.Port
	password := conf.Redis.Password
	db := conf.Redis.DB

	RedisClient = redis.NewClient(&redis.Options{
		Addr:     ip + ":" + port,
		Password: password,
		DB:       db,
	})
	_, err := RedisClient.Ping(ctx).Result()
	if err != nil {
		panic(err)
	}
	log.Println("Redis init success")
}

func GetRedis() *redis.Client {
	return RedisClient
}

func CloseRedis() {
	err := RedisClient.Close()
	if err != nil {
		panic(err)
	}
	log.Println("Redis close success")
}

func IsExist(key string) bool {
	rdb := GetRedis()
	exist, err := rdb.Exists(ctx, key).Result()
	if err != nil {
		return false
	}
	return exist == 1
}

// CacheSet 设置键值对
func CacheSet(key string, value interface{}) error {
	rdb := GetRedis()
	val, err := json.Marshal(value)
	if err != nil {
		return err
	}
	// 增加一个小范围随机数平坦过期时间,防止缓存雪崩
	rd := time.Duration(rand.Intn(1000))
	err = rdb.Set(ctx, key, val, valueExpire+time.Second*rd).Err()
	if err != nil {
		return err
	}
	return nil
}

// CacheGet 获取键值对
func CacheGet(key string) ([]byte, error) {
	rdb := GetRedis()
	val, err := rdb.Get(ctx, key).Result()
	if err != nil {
		return nil, err
	}
	return []byte(val), nil
}

// CacheHSet 设置hash键值对
func CacheHSet(key string, field string, value ...interface{}) error {
	rdb := GetRedis()
	for _, v := range value {
		val, err := json.Marshal(v)
		if err != nil {
			return err
		}
		err = rdb.HSet(ctx, key, field, val).Err()
		if err != nil {
			return err
		}
	}
	return nil
}

func CacheHGet(key string, field string) ([]byte, error) {
	rdb := GetRedis()
	val, err := rdb.HGet(ctx, key, field).Result()
	if err != nil {
		return nil, err
	}
	return []byte(val), nil
}

// MutexLock 后期可以额外增加分布式锁
func MutexLock(key string) bool {
	rdb := GetRedis()
	Bool, _ := rdb.SetNX(ctx, "user_lock"+key, "1", time.Second*5).Result()
	return Bool
}

func MutexUnlock(key string) error {
	rdb := GetRedis()
	err := rdb.Del(ctx, "user_lock"+key).Err()
	if err != nil {
		return err
	}
	return nil
}
