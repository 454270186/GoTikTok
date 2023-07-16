package rcache

import (
	"context"
	"errors"
	"fmt"
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
// Provide two interfaces to update favorite
// 1. LikeVideo
// 2. UnlikeVideo

// Key:   video::{videoID}::user::{userID}
// Value: {action_type}::{created_at}

func LikeVideo(ctx context.Context, favOp *FavoriteCache) error {
	errLock := redis.Lock() 
	if errLock != nil {
		return errors.New("error while lock redis: " + errLock.Error())
	}

	likeKey := fmt.Sprintf("video::%s::user::%s", favOp.VideoID, favOp.UserID)
	likeValue := fmt.Sprintf("%s::%s", favOp.ActionType, favOp.CreatedAt)

	isExist, err := redis.GetRDB().Exists(ctx, likeKey).Result()
	if err != nil {
		redis.Unlock()
		return errors.New("error while check key exist" + err.Error())
	}
	redis.Unlock()

	if isExist == 0 {
		// key is not exist, add to redis
		errLock = redis.Lock()
		if errLock != nil {
			return errLock
		}
		err := redis.PutKey(ctx, likeKey, likeValue)
		if err != nil {
			return errors.New("error while put likekey: " + err.Error())
		}
		return nil
	} else {
		// key is exist, check for update
		res, err := redis.Get(ctx, likeKey)
		if err != nil {
			return errors.New("error while get key" + err.Error())
		}

		valSplit := strings.Split(res, "::")
		likeOp, likeTime := valSplit[0], valSplit[1]

		
	}
}