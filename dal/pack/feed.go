package pack

import (
	"time"

	"github.com/454270186/GoTikTok/dal"
	"github.com/454270186/GoTikTok/pkg/minio"
	"github.com/454270186/GoTikTok/rpc/feed/types/feed"
)

var feedDB dal.FeedDB

func GetVideoListByTime(limit int, latestTime int64) ([]*feed.Video, int64, error) {
	dalVideos, err := feedDB.GetVideoLists(ctx, limit, latestTime)
	if err != nil {
		return nil, 0, err
	}

	var nextTime int64
	if len(dalVideos) == 0 {
		nextTime = time.Now().UnixMilli()
		return []*feed.Video{}, nextTime, nil
	} else {
		nextTime = dalVideos[len(dalVideos)-1].UpdatedAt.UnixMilli()
	}

	rpcVideoList, err := convertVideoLists(dalVideos)
	if err != nil {
		return nil, 0, err
	}

	return rpcVideoList, nextTime, nil
}

// Convert dal videos model to RPC video model
func convertVideoLists(dalVideos []*dal.Video) ([]*feed.Video, error) {
	minioBucketName := minio.VideoBucketName

	rpcVideos := make([]*feed.Video, 0)
	for _, dalVideo := range dalVideos {
		author, err := userDB.GetById(ctx, dalVideo.AuthorID)
		if err != nil {
			return nil, err
		}

		videoPlayUrl, err := minio.GetFileURL(minioBucketName, dalVideo.VideoName, 0)
		if err != nil {
			return nil, err
		}

		videoCoverUrl, err := minio.GetFileURL(minioBucketName, dalVideo.CoverName, 0)
		if err != nil {
			return nil, err
		}

		v := feed.Video{
			Id:            int64(dalVideo.ID),
			Author:        convertFeedUser(author),
			PlayUrl:       videoPlayUrl.String(),
			CoverUrl:      videoCoverUrl.String(),
			FavoriteCount: dalVideo.FavoriteCount,
			CommentCount:  dalVideo.CommentCount,
			IsFavorite:    false,
			Title:         dalVideo.Title,
		}

		rpcVideos = append(rpcVideos, &v)
	}

	return rpcVideos, nil
}

// Convert dal.User to RPC feed.User
func convertFeedUser(dalUser *dal.User) *feed.User {
	if dalUser == nil {
		return nil
	}

	return &feed.User{
		Id:            int64(dalUser.ID),
		Name:          dalUser.Username,
		FollowCount:   dalUser.FollowingCount,
		FollowerCount: dalUser.FollowerCount,
		IsFollow:      true,
	}
}