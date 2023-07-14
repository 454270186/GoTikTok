package redis

import (
	"sync"

	"github.com/redis/go-redis/v9"
)

var (
	RedisAddr = "127.0.0.1:6379"
	RedisPassword = "2021110003"
)

var (
	rdb *redis.Client
	redisInitOnce sync.Once
)

// Get redis client
func GetRDB() *redis.Client {
	if rdb == nil {
		redisInitOnce.Do(func() {
			rdb = redis.NewClient(&redis.Options{
				Addr: RedisAddr,
				Password: RedisPassword,
				DB: 0,
			})
		})
	}

	return rdb
}