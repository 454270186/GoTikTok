package pack

import (
	"github.com/454270186/GoTikTok/dal"
	"github.com/454270186/GoTikTok/rpc/publish/types/publish"
)

var publishDB dal.PublishDB

func CreateVideo(userID uint, playURL string, coverURL string, title string) error {
	videoModel := &dal.Video{
		AuthorID: userID,
		PlayURL: playURL,
		CoverURL: coverURL,
		Title: title,
	}

	return publishDB.CreateVideo(ctx, videoModel)
}

func GetVideoList(userID uint) ([]*publish.Video, error) {
	videoList := []*publish.Video{}
	
	dalVideoList, err := publishDB.GetListByID(ctx, userID)
	if err != nil {
		return nil, err
	}

	for _, video := range dalVideoList {
		author, err := userDB.GetById(ctx, video.AuthorID)
		if err != nil {
			return nil, err
		}

		pVideo := &publish.Video{
			Id: int64(video.ID),
			Author: convertUser(author),
			PlayUrl: video.PlayURL,
			CoverUrl: video.CoverURL,
			FavoriteCount: video.FavoriteCount,
			CommentCount: video.CommentCount,
			IsFavorite: false,
			Title: video.Title,
		}

		videoList = append(videoList, pVideo)
	}

	return videoList, nil
}

// Convert dal.User model to RPC User model
func convertUser(dalUser *dal.User) *publish.User {
	if dalUser == nil {
		return nil
	}

	return &publish.User{
		Id:            int64(dalUser.ID),
		Name:          dalUser.Username,
		FollowCount:   dalUser.FollowingCount,
		FollowerCount: dalUser.FollowerCount,
		IsFollow:      true,
	}
}
