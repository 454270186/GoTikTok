package locker

import (
	"context"
	"errors"
	"log"
	"time"

	"github.com/gofrs/uuid"
	"github.com/redis/go-redis/v9"
)

type Locker struct {
	rclient         *redis.Client
	unlockScript    *redis.Script
	ttl             time.Duration
	trylockInterval time.Duration
}

func NewLocker(client *redis.Client, ttl, trylockInterval time.Duration) *Locker {
	return &Locker{
		rclient:         client,
		unlockScript:    redis.NewScript(unlockScript),
		ttl:             ttl,
		trylockInterval: trylockInterval,
	}
}

func (l *Locker) GetLock(lockResource string) *Lock {
	uid, _ := uuid.NewV4()
	randomVal := uid.String()

	return &Lock{
		rclient:         l.rclient,
		unlockScript:    l.unlockScript,
		lockKey:         lockResource,
		lockValue:       randomVal,
		watchDog:        make(chan struct{}),
		ttl:             l.ttl,
		trylockInterval: l.trylockInterval,
	}
}

type Lock struct {
	rclient         *redis.Client
	unlockScript    *redis.Script
	lockKey         string
	lockValue       string
	watchDog        chan struct{}
	ttl             time.Duration
	trylockInterval time.Duration
}

// Lock() acquires a lock
func (l *Lock) Lock(ctx context.Context) error {
	// try to acquire lock
	err := l.TryLock(ctx)
	if err == nil {
		log.Println("lock successfully")
		return nil
	}
	if !errors.Is(err, ErrLockFailed) {
		log.Println("1 unexpect error:", err.Error())
		return err
	}

	log.Println("fail to acquire lock, start try")
	// Fail to acquire lock, try it by trylockInterval
	tryLockTicker := time.NewTicker(l.trylockInterval)
	defer tryLockTicker.Stop()

	for {
		log.Println("Try to acquire lock")
		select {
		case <-tryLockTicker.C:
			err := l.TryLock(ctx)
			if err != nil {
				log.Println("lock successfully")
				return nil
			}
			if !errors.Is(err, ErrLockFailed) {
				log.Println("2 unexpect error:", err.Error())
				return err
			}

		case <-ctx.Done():
			log.Println("timeout")
			return ErrTimeout
		}
	}
}

// TryLock() tries to acquire a lock
func (l *Lock) TryLock(ctx context.Context) error {
	log.Println("Try to acquire lock")
	success, err := l.rclient.SetNX(ctx, l.lockKey, l.lockValue, l.ttl).Result()
	if err != nil {
		return err
	}

	if !success {
		// fail to acquire lock
		return ErrLockFailed
	}

	// start watch dog
	go l.startWatchDog()
	return nil
}

func (l *Lock) startWatchDog() {
	ticker := time.NewTicker(l.ttl / 3)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			ctx, cancel := context.WithTimeout(context.Background(), l.ttl/3*2)

			// extend lock ttl
			ok, err := l.rclient.Expire(ctx, l.lockKey, l.ttl).Result()
			cancel()

			if err != nil || !ok {
				// if err or lock not exist, stop watchdog
				return
			}
			log.Println("extend lock ttl")

		case <-l.watchDog:
			// already unlock
			return
		}
	}
}

func (l *Lock) Unlock(ctx context.Context) error {
	err := l.unlockScript.Run(ctx, l.rclient, []string{l.lockKey}, l.lockValue).Err()
	close(l.watchDog)
	if err == nil {
		log.Println("unlock successfully")
	}
	return err
}
