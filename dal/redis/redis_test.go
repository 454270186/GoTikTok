package redis_test

import (
	"context"
	"fmt"
	"log"
	"sync"
	"testing"

	"github.com/454270186/GoTikTok/dal/redis"
	"github.com/454270186/GoTikTok/dal/redis/rmodel"
	r "github.com/redis/go-redis/v9"
)

var rdb *r.Client
func TestRcache(t *testing.T) {
	fmt.Println("test start")
	favOp1 := rmodel.FavoriteCache{
		VideoID: 2,
		UserID: 12,
		ActionType: 1,
		CreatedAt: 9,
	}
	favOp2 := rmodel.FavoriteCache{
		VideoID: 2,
		UserID: 12,
		ActionType: 0,
		CreatedAt: 10,
	}

	var wg sync.WaitGroup
	wg.Add(2)
	go func ()  {
		defer wg.Done()
		err := redis.UpdateVideo(context.Background(), &favOp2)
		if err != nil {
			log.Println("go1 error", err)
		}
	}()

	go func ()  {
		defer wg.Done()
		err := redis.UpdateVideo(context.Background(), &favOp1)
		if err != nil {
			log.Println("go2 error", err)
		}
	}()

	wg.Wait()
}

func TestMain(m *testing.M) {
	fmt.Println("test main")
	rdb = redis.GetRDB()
	fmt.Println("run")
	m.Run()
}