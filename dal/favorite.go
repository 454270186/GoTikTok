package dal

import (
	"context"
	"errors"

	"gorm.io/gorm"
)

func (Favorite) TableName() string {
	return "users_favorite_videos"
}

type FavoriteDB struct {
	DB *gorm.DB
}

func NewFavoriteDB() FavoriteDB {
	return FavoriteDB{
		DB: newDB(),
	}
}

func (f FavoriteDB) GetFavVideoIDsByUserID(ctx context.Context, userID uint) ([]uint, error) {
	var userVideos []Favorite
	var videoIDs []uint

	err := f.DB.WithContext(ctx).Where("user_id = ?", userID).Find(&userVideos).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("records not found")
		}
		return nil, err
	}

	for _, uv := range userVideos {
		videoIDs = append(videoIDs, uv.VideoID)
	}

	return videoIDs, nil
}

func (f FavoriteDB) GetVideoByID(ctx context.Context, videoIDs ...uint) ([]*Video, error) {
	videos := []*Video{}

	for _, videoID := range videoIDs {
		var video Video
		err := f.DB.WithContext(ctx).First(&video, videoID).Error
		if err != nil {
			return nil, err
		}

		videos = append(videos, &video)
	}

	return videos, nil
}

func (f FavoriteDB) AddFavorite(ctx context.Context, userID uint, videoID uint) error {
	newFavorite := Favorite{
		UserID: userID,
		VideoID: videoID,
	}

	return f.DB.WithContext(ctx).Create(&newFavorite).Error
}

func (f FavoriteDB) DelFavorite(ctx context.Context, userID uint, videoID uint) error {
	delFavorite := Favorite{
		UserID: userID,
		VideoID: videoID,
	}

	return f.DB.WithContext(ctx).Delete(&delFavorite).Error
}

func (f FavoriteDB) IsExist(ctx context.Context, userID uint, videoID uint) (bool, error) {
	findFavorite := Favorite{
		UserID: userID,
		VideoID: videoID,
	}

	err := f.DB.WithContext(ctx).First(&findFavorite).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return false, nil
		}

		return false, err
	}

	return true, nil
}
