/*
Package cache: Provides methods for working with Redis.
Package Functionality: Connects to Redis, performs ping operations, sets and retrieves data, and invalidates the cache.
File redis.go: Contains functions for working with Redis.

Author: MinhDan <nguyenmd.works@gmail.com>
*/
package cache

import (
	"base/pkg/log"
	"encoding/json"
	"sync"
	"time"

	"github.com/go-redis/redis"
	"github.com/go-redis/redis_rate"
)

// Redis Configuration Struct
type redisConfig struct {
	Host     string
	Port     string
	Password string
	Name     int
}

type RedisCache struct {
	RWMutex sync.RWMutex
	Limiter *redis_rate.Limiter
	Client  *redis.Client
}

// Redis Configuration Variable
var redisCfg redisConfig

// Redis Global Cache
var Redis *RedisCache

// Redis Connect Function
func redisConnect() *RedisCache {
	// Initialize Connection
	client := redis.NewClient(&redis.Options{
		Addr: redisCfg.Host + ":" + redisCfg.Port,
		// If You Use Redis Container, Default We Use Password With ""
		Password: redisCfg.Password,
		DB:       redisCfg.Name,
	})

	// Test Connection
	_, err := client.Ping().Result()
	if err != nil {
		log.Println(log.LogLevelFatal, "redis-connect", err.Error())
	}

	// Return Connection
	return &RedisCache{
		Limiter: redis_rate.NewLimiter(client),
		Client:  client,
	}
}

// Ping Method To Check Connection
func (redisCache *RedisCache) Ping() (string, error) {
	return redisCache.Client.Ping().Result()
}

// Set Method With Key, Value, Time To Live
func (redisCache *RedisCache) Set(key string, value any, timeToLive time.Duration) error {
	redisCache.RWMutex.Lock()

	defer redisCache.RWMutex.Unlock()

	byteValue, err := json.Marshal(value)
	if err != nil {
		return err
	}
	// Set Value TO Redis Cache
	return redisCache.Client.Set(key, byteValue, timeToLive).Err()
}

// Get Method To Retrieve Value Of A Key
func (redisCache *RedisCache) Get(key string) ([]byte, bool, error) {

	// Get Value From Redis
	byteValue, err := redisCache.Client.Get(key).Bytes()
	if err == redis.Nil {
		return nil, false, nil
	}
	if err != nil {
		return nil, false, err
	}

	// Set to local cache
	return byteValue, true, nil
}

// Invalidate Method To Delete A Key From Cache
func (redisCache *RedisCache) Invalidate(key string) error {
	redisCache.RWMutex.Lock()
	defer redisCache.RWMutex.Unlock()

	// Delete Key From Cache
	return redisCache.Client.Del(key).Err()
}

// Close Method To Close Connection To Cache Server
func (redisCache *RedisCache) Close() {
	redisCache.Client.Close()
}
