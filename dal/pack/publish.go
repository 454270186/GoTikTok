package pack

import (
	"github.com/454270186/GoTikTok/dal"
	"github.com/454270186/GoTikTok/pkg/minio"
	"github.com/454270186/GoTikTok/rpc/publish/types/publish"

	"gorm.io/gorm"
)

var publishDB dal.PublishDB

// @Modification
// Change storing URL to storing Minio object name
// Change return (error) to return (newVideoID, error)
func CreateVideo(userID uint, playName string, coverName string, title string) (uint, error) {
	videoModel := &dal.Video{
		AuthorID:  userID,
		VideoName: playName,
		CoverName: coverName,
		Title:     title,
	}

	err := publishDB.DB.Transaction(func(tx *gorm.DB) error {
		// Create video record
		err := tx.Create(videoModel).Error
		if err != nil {
			return err
		}

		// add user work count
		user := dal.User{ID: userID}
		err = tx.First(&user).Error
		if err != nil {
			return err
		}

		user.WorkCount++

		err = tx.Save(&user).Error
		if err != nil {
			return err
		}

		return nil
	})
	if err != nil {
		return 0, err
	}

	return videoModel.ID, nil
}

func DelVideoByID(videoID uint) error {
	return publishDB.DB.Transaction(func(tx *gorm.DB) error {
		delVideo := dal.Video{ID: videoID}
		if err := tx.First(&delVideo).Error; err != nil {
			return err
		}

		if err := tx.Delete(&delVideo).Error; err != nil {
			return err
		}

		user := dal.User{ID: delVideo.AuthorID}
		err := tx.First(&user).Error
		if err != nil {
			return err
		}

		user.WorkCount--

		err = tx.Save(&user).Error
		if err != nil {
			return err
		}

		return nil
	})
}

func GetVideoList(userID uint) ([]*publish.Video, error) {
	minioBucketName := minio.VideoBucketName

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

		// Get URL from minio
		videoPlayUrl, err := minio.GetFileURL(minioBucketName, video.VideoName, 0)
		if err != nil {
			return nil, err
		}

		videoCoverUrl, err := minio.GetFileURL(minioBucketName, video.CoverName, 0)
		if err != nil {
			return nil, err
		}

		pVideo := &publish.Video{
			Id:            int64(video.ID),
			Author:        convertUser(author),
			PlayUrl:       videoPlayUrl.String(),
			CoverUrl:      videoCoverUrl.String(),
			FavoriteCount: video.FavoriteCount,
			CommentCount:  video.CommentCount,
			IsFavorite:    false,
			Title:         video.Title,
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
		WorkCount:     dalUser.WorkCount,
	}
}
