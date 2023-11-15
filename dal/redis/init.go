package redis

import (
	"context"
	"log"
	"sync"
	"time"

	"github.com/454270186/GoTikTok/pkg/timer"
	"github.com/joho/godotenv"
	r "github.com/redis/go-redis/v9"
)

var (
	RedisAddr string
	RedisPassword string
	LockerAddr string
	SyncDuration = 2 * time.Second
)

var (
	rdb *r.Client
	lockerRdb *r.Client
	redisInitOnce sync.Once
)

func init() {
	rdbEnv, err := godotenv.Read("../.env")
	if err != nil {
		panic("fail to read redis env: " + err.Error())
	}

	RedisAddr = rdbEnv["R_ADDR"]
	RedisPassword = rdbEnv["R_PSW"]
	LockerAddr = rdbEnv["R_LOCKER_ADDR"]

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

func GetLockerRDB() *r.Client {
	if lockerRdb != nil {
		return lockerRdb
	}

	log.Println("locker addr:", LockerAddr)
	lockerRdb = r.NewClient(&r.Options{
		Addr: LockerAddr,
		Password: RedisPassword,
		DB: 0,
	})

	_, err := lockerRdb.Ping(context.Background()).Result()
	if err != nil {
		panic(err)
	}
	log.Println("Locker RDB initialized")
	return lockerRdb
}