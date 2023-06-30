package dal

import (
	"context"
	"log"

	"gorm.io/gorm"
)

func (Video) TableName() string {
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

func (p PublishDB) CreateVideo(ctx context.Context, video *Video) error {
	return p.DB.WithContext(ctx).Create(video).Error
}
