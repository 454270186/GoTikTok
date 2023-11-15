package locker

import (
	"errors"
	"time"
)

const (
	TTL = 20 * time.Second
	TrylockInterval = time.Second
	unlockScript = `
	if redis.call("get",KEYS[1]) == ARGV[1] then
    	return redis.call("del",KEYS[1])
	else
    	return 0
	end
	`
)

// Error
var (
	ErrLockFailed = errors.New("lock failed")
	ErrTimeout = errors.New("lock timeout")
)