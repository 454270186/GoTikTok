package redis

import (
	"sync"
	"time"

	//"github.com/454270186/GoTikTok/pkg/timer"
	"github.com/454270186/GoTikTok/pkg/timer"
	"github.com/joho/godotenv"
	r "github.com/redis/go-redis/v9"
)

var (
	RedisAddr string
	RedisPassword string

	SyncDuration = 5 * time.Second
)

var (
	rdb *r.Client
	redisInitOnce sync.Once
)

func init() {
	rdbEnv, err := godotenv.Read("../.env")
	if err != nil {
		panic("fail to read redis env: " + err.Error())
	}

	RedisAddr = rdbEnv["R_ADDR"]
	RedisPassword = rdbEnv["R_PSW"]

	timer.SyncTimer(SyncDuration, MoveFavoriteToDB)
}

// Get redis client
func GetRDB() *r.Client {
	if rdb == nil {
		redisInitOnce.Do(func() {
			rdb = r.NewClient(&r.Options{
				Addr: RedisAddr,
				Password: RedisPassword,
				DB: 0,
			})
		})
	}

	return rdb
}