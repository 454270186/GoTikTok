package pack

import (
	"errors"
	"log"
	"strconv"

	"github.com/454270186/GoTikTok/dal"
	"github.com/454270186/GoTikTok/pkg/minio"
	"github.com/454270186/GoTikTok/rpc/favorite/types/favorite"
)

var favoriteDB dal.FavoriteDB

func GetFavVideosByUserID(userID uint) ([]*favorite.Video, error) {
	videoIDs, err := favoriteDB.GetFavVideoIDsByUserID(ctx, userID)
	if err != nil {
		return nil, err
	}

	videos, err := favoriteDB.GetVideoByID(ctx, videoIDs...)
	if err != nil {
		return nil, err
	}

	favVideos, err := convertFavVideos(videos)
	if err != nil {
		return nil, err
	}

	return favVideos, nil
}

func convertFavVideos(videos []*dal.Video) ([]*favorite.Video, error) {
	minioBucketName := minio.VideoBucketName

	if videos == nil {
		return nil, errors.New("videos is nil")
	}

	favVideos := []*favorite.Video{}
	for _, dalVideo := range videos {
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

		favVideo := &favorite.Video{
			Id: int64(dalVideo.ID),
			Author: convertFavUser(author),
			PlayUrl: videoPlayUrl.String(),
			CoverUrl: videoCoverUrl.String(),
			FavoriteCount: dalVideo.FavoriteCount,
			CommentCount: dalVideo.CommentCount,
			IsFavorite: true,
			Title: dalVideo.Title,
		}

		favVideos = append(favVideos, favVideo)
	}

	return favVideos, nil
}

// Convert dal.User to RPC favorite.User
func convertFavUser(dalUser *dal.User) *favorite.User {
	if dalUser == nil {
		return nil
	}

	return &favorite.User{
		Id:            int64(dalUser.ID),
		Name:          dalUser.Username,
		FollowCount:   dalUser.FollowingCount,
		FollowerCount: dalUser.FollowerCount,
		IsFollow:      true,
	}
}

func LikeVideo(userIDstr, videoIDstr string) error {
	userID, err := strconv.ParseUint(userIDstr, 10, 64)
	if err != nil {
		return err
	}

	videoID, err := strconv.ParseUint(videoIDstr, 10, 64)
	if err != nil {
		return err
	}

	return favoriteDB.AddFavorite(ctx, uint(userID), uint(videoID))
}

func UnlikeVideo(userIDstr, videoIDstr string) error {
	userID, err := strconv.ParseUint(userIDstr, 10, 64)
	if err != nil {
		return err
	}

	videoID, err := strconv.ParseUint(videoIDstr, 10, 64)
	if err != nil {
		return err
	}

	return favoriteDB.DelFavorite(ctx, uint(userID), uint(videoID))
}

func IsVideoLiked(userID, videoID uint) bool {
	isLiked, err := favoriteDB.IsExist(ctx, userID, videoID)
	if err != nil {
		log.Println(err)
		return false
	}

	return isLiked
}