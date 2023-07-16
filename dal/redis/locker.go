package redis

import (
	"context"
	"log"
	"time"

	"github.com/redis/go-redis/v9"
)

// redis distributed lock

var (
	lockKey = "redislock"
	exp = 10 * time.Second
)

// Acquire Lock
// return true if success
// return false if failed
func Lock() {
	redisCli := GetRDB()
	var resp *redis.BoolCmd

	// 采用自旋的方式轮询锁状态
	for {
		resp = redisCli.SetNX(context.Background(), lockKey, "locked", exp)
		success, err := resp.Result()
		if err != nil && success {
			return
		} else {
			log.Println("lock failed")
		}
	}
}

// Unlock() releases redis lock
func Unlock() {
	redisCli := GetRDB()
	delRes, err := redisCli.Del(context.Background(), lockKey).Result()
	if err == nil && delRes > 0 {
		log.Println("unlock successfully")
	} else {
		log.Println("unlock failed")
	}
}