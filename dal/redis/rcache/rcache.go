package rcache

import (
	"context"
	"errors"
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/454270186/GoTikTok/dal/redis"
)

type FavoriteCache struct {
	VideoID    uint `json:"video_id"`
	UserID     uint `json:"user_id"`
	ActionType uint `json:"action_type"`
	CreatedAt  uint `json:"created_at"`
}

// VideoCache key pattern is:
// {videoID}
// VideoCache value pattern is:
// {authorID}::{videoObjName}::{coverObjName}::{fav_cnt}::{cmt_cnt}::{title}

// Favorite key pattern is: 
// video::{videoID}::user::{userID}
// Favorite value pattern is:
// {action_type}::{created_at}

/* Read */
// Provide one interface to get favorite list by userID
// Step:
// 1. get a fav_videoID list (all favorite video by this user)
// 2. traverse fav_videoID list to get fav_video from videoCache, 
//    if not found in cache, get from DB, and then async add it into cache
// 3. pack user and fav_videos, send back to rpc caller

/* Write */
// Provide one interfaces to update favorite
// action_type: 0 ==> like, 1 ==> unlike

// Key:   video::{videoID}::user::{userID}
// Value: {action_type}::{created_at}
func UpdateVideo(ctx context.Context, favOp *FavoriteCache) error {
	redis.Lock() 

	likeKey := fmt.Sprintf("video::%d::user::%d", favOp.VideoID, favOp.UserID)
	likeValue := fmt.Sprintf("%d::%d", favOp.ActionType, favOp.CreatedAt)

	isExist, err := redis.GetRDB().Exists(ctx, likeKey).Result()
	if err != nil {
		redis.Unlock()
		return errors.New("error while check key exist" + err.Error())
	}
	redis.Unlock()

	if isExist == 0 {
		// key is not exist, add to redis
		redis.Lock()

		err := redis.PutKey(ctx, likeKey, likeValue)
		if err != nil {
			return errors.New("error while put likekey: " + err.Error())
		}

		redis.Unlock()
		return nil
	} else {		
		redis.Lock()

		// key is exist, check for update
		res, err := redis.Get(ctx, likeKey)
		if err != nil {
			log.Println(err)
			redis.Unlock()
			return errors.New("error while get key" + err.Error())
		}

		valSplit := strings.Split(res, "::")
		rActionType, rCreatedAt := valSplit[0], valSplit[1]

		actionType := strconv.Itoa(int(favOp.ActionType))
		log.Println(rActionType, favOp.ActionType)
		if rActionType == actionType {
			// if action type is same, return
			redis.Unlock()
			return nil
		}

		// othorwise, check createdAt
		rCreatedAtUnix, err := strconv.Atoi(rCreatedAt)
		if err != nil {
			log.Println(err)
			redis.Unlock()
			return errors.New("error while convert str to int: " + err.Error())
		}
		if rCreatedAtUnix >= int(favOp.CreatedAt) {
			// the record opretion is later than this operation, return
			redis.Unlock()
			return nil
		}

		// Update record
		err = redis.PutKey(ctx, likeKey, likeValue)
		if err != nil {
			redis.Unlock()
			return err
		}
	}
	
	redis.Unlock()
	return nil
}