package redis

import (
	"context"
	"errors"
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/454270186/GoTikTok/dal/redis/locker"
	"github.com/454270186/GoTikTok/dal/redis/rmodel"
)

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
// action_type: 1 ==> like, 2 ==> unlike

// Key:   video::{videoID}::user::{userID}
// Value: {action_type}::{created_at}
func UpdateVideo(ctx context.Context, favOp *rmodel.FavoriteCache) error {
	Locker := locker.NewLocker(GetLockerRDB(), locker.TTL, locker.TrylockInterval)
	Lock := Locker.GetLock("LikeLock")

	err := Lock.Lock(ctx)
	if err != nil {
		log.Println(err)
		return err
	}
	defer Lock.Unlock(ctx)

	likeKey := fmt.Sprintf("video::%d::user::%d", favOp.VideoID, favOp.UserID)
	likeValue := fmt.Sprintf("%d::%d", favOp.ActionType, favOp.CreatedAt)

	isExist, err := GetRDB().Exists(ctx, likeKey).Result()
	if err != nil {
		return errors.New("error while check key exist" + err.Error())
	}

	if isExist == 0 {
		// key is not exist, add to redis
		err = PutKey(ctx, likeKey, likeValue)
		if err != nil {
			return errors.New("error while put likekey: " + err.Error())
		}

		return nil
	} else {		
		// key is exist, check for update
		res, err := Get(ctx, likeKey)
		if err != nil {
			log.Println(err)
			return errors.New("error while get key" + err.Error())
		}

		valSplit := strings.Split(res, "::")
		rActionType, rCreatedAt := valSplit[0], valSplit[1]

		actionType := strconv.Itoa(int(favOp.ActionType))
		log.Println(rActionType, favOp.ActionType)
		if rActionType == actionType {
			// if action type is same, return
			return nil
		}

		// othorwise, check createdAt
		rCreatedAtUnix, err := strconv.Atoi(rCreatedAt)
		if err != nil {
			log.Println(err)
			return errors.New("error while convert str to int: " + err.Error())
		}
		if rCreatedAtUnix >= int(favOp.CreatedAt) {
			// the record opretion is later than this operation, return
			return nil
		}

		// Update record
		err = PutKey(ctx, likeKey, likeValue)
		if err != nil {
			return err
		}
	}
	
	return nil
}