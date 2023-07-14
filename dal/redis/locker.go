package redis

import (
	"context"
	"log"
	"time"
)

// redis distributed lock

var (
	lockKey = "redislock"
	exp = 5 * time.Second
)

// Acquire Lock
// return true if success
// return false if failed
func Lock() bool {
	redisCli := GetRDB()
	result, err := redisCli.SetNX(context.Background(), lockKey, "locked", exp).Result()
	if err != nil {
		log.Println("error while acquire lock")
		return false
	}

	return result
}

// Unlock() releases redis lock
func Unlock() {
	redisCli := GetRDB()
	_, err := redisCli.Del(context.Background(), lockKey).Result()
	if err != nil {
		log.Println("error while release lock")
	}
}