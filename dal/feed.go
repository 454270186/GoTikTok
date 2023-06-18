package dal

import (
	"context"
	"time"

	"gorm.io/gorm"
)

type FeedDB struct {
	DB *gorm.DB
}

func NewFeedDB() FeedDB {
	return FeedDB{
		DB: newDB(),
	}
}

// 
func (f FeedDB) GetVideoLists(ctx context.Context, limit int, latestTime int64) ([]*Video, error) {
	videos := make([]*Video, 0)

	if latestTime == 0 {
		curTime := time.Now().UnixMilli()
		latestTime = curTime
	}

	err := f.DB.WithContext(ctx).Limit(limit).Order("updated_at DESC").Find(&videos, "updated_at < ?", time.UnixMilli(latestTime)).Error
	if err != nil {
		return nil, err
	}

	return videos, nil
}