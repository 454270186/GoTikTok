package redis

import (
	"context"
	"encoding/json"
	"log"
	"strings"

	"github.com/454270186/GoTikTok/dal/pack"
	"github.com/454270186/GoTikTok/dal/redis/rmodel"
	"github.com/454270186/GoTikTok/pkg/rabbitmq"
)

/* synchronization for favorite */

// Synchronization for rabbitmq and redis
func Consume() {
	rmq := rabbitmq.NewSimple()
	msgs, err := rmq.GetMsgs()
	if err != nil {
		log.Println("error while get rabbitmq")
		return
	}

	// consume all msgs
	for msg := range msgs {
		favOp := rmodel.FavoriteCache{}
		if err := json.Unmarshal(msg.Body, &favOp); err != nil {
			log.Println("error while unmarshall msg body")
			continue
		}
		log.Printf("Get a message: %v", favOp)
		
		if err := UpdateVideo(context.Background(), &favOp); err != nil {
			log.Println("error while update video to redis")
		}

		// if autoAck is false, need to call msg.Ack() mannully
		// msg.Ack()
	}
}

// Synchronization for redis and mysql
func MoveFavoriteToDB() {
	// Get all keys
	keys, err := GetKeys(context.Background(), "video::*::user::*")
	if err != nil {
		log.Println("error while get all keys")
		return
	}

	for _, key := range keys {
		val, err := Get(context.Background(), key)
		if err != nil {
			log.Println("error while get value by key")
			continue
		}

		keySpli := strings.Split(key, "::")
		videoIDstr, userIDstr := keySpli[1], keySpli[3]
		valSpli := strings.Split(val, "::")
		actionType := valSpli[0]

		switch actionType {
		case "1":
			pack.LikeVideo(userIDstr, videoIDstr)
			err := DelKey(context.Background(), key)
			if err != nil {
				log.Println("error while del key")
				return
			}
		case "2":
			pack.UnlikeVideo(userIDstr, videoIDstr)
			err := DelKey(context.Background(), key)
			if err != nil {
				log.Println("error while del key")
				return
			}
		default:
			log.Println("unknow action type")
		}
	}
}