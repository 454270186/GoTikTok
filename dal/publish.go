package dal

import (
	"context"
	"log"

	"gorm.io/gorm"
)

type Video struct {
	gorm.Model
	AuthorID      uint   `gorm:"column:author_id"`
	PlayURL       string `gorm:"column:play_url"`
	CoverURL      string `gorm:"column:cover_url"`
	FavoriteCount int64  `gorm:"column:favorite_count"`
	CommentCount  int64  `gorm:"column:comment_count"`
	Title         string `gorm:"column:title"`
}

func (v Video) GetTableName() string {
	return "videos"
}

type PublishDB struct {
	DB *gorm.DB
}

func NewPublishDB() PublishDB {
	return PublishDB{
		DB: newDB(),
	}
}

// Get user's video publishing list
func (p PublishDB) GetListByID(ctx context.Context, userID uint) ([]*Video, error) {
	videoList := make([]*Video, 0)

	err := p.DB.WithContext(ctx).Where("author_id = ?", userID).Find(&videoList).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			log.Println("Video not found")
			return nil, err
		}
		return nil, err
	}

	return videoList, nil
}