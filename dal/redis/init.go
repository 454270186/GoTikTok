package redis

import (
	"sync"
	"time"

	"github.com/454270186/GoTikTok/pkg/timer"
	"github.com/joho/godotenv"
	"github.com/redis/go-redis/v9"
)

var (
	RedisAddr string
	RedisPassword string

	SyncDuration = 10 * time.Second
)

var (
	rdb *redis.Client
	redisInitOnce sync.Once
)

func init() {
	rdbEnv, err := godotenv.Read()
	if err != nil {
		panic("fail to read redis env: " + err.Error())
	}

	RedisAddr = rdbEnv["R_ADDR"]
	RedisPassword = rdbEnv["R_PSW"]

	timer.SyncTimer(SyncDuration, moveFavoriteToDB)
}

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