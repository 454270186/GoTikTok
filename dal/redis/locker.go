package redis

import (
	"context"
	//"fmt"
	"log"
	"time"

	"github.com/gofrs/uuid"
	r "github.com/redis/go-redis/v9"
)

// redis distributed lock

var (
	lockKey = "redislock"
	exp     = 10 * time.Second
	maxRetries = 10 // 最大自旋次数
	retryInterval = 100 * time.Millisecond // 自旋间隔
)

// Acquire Lock
func Lock() bool {
	redisCli := GetRDB()
	var resp *r.BoolCmd
	retries := 0
	
	// generate unique lock value
	uid, _ := uuid.NewV4()
	lockVal := uid.String()

	// 采用自旋的方式轮询锁状态
	for {
		resp = redisCli.SetNX(context.Background(), lockKey, lockVal, exp)
		success, err := resp.Result()
		if err != nil {
			log.Println("lock failed", err)
			return false
		}
		if success {
			return true
		}

		// 自旋等待
		retries++
		if retries >= maxRetries {
			break
		}
		time.Sleep(retryInterval)
	}

	return false
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