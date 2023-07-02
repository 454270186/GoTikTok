package dal

import (
	"context"

	"gorm.io/gorm"
)

func (Comments) TableName() string {
	return "comments"
}

type CommentDB struct {
	DB *gorm.DB
}

func NewCommentDB() CommentDB {
	return CommentDB{
		DB: newDB(),
	}
}

func (c CommentDB) GetById(ctx context.Context, commentID uint) (*Comments, error) {
	comment := Comments{ID: commentID}
	err := c.DB.WithContext(ctx).First(&comment).Error
	if err != nil {
		return nil, err
	}

	return &comment, nil
}
 
func (c CommentDB) GetByVideoID(ctx context.Context, videoID uint) ([]*Comments, error) {
	comments := []*Comments{}

	err := c.DB.WithContext(ctx).Where("video_id = ?", videoID).Order("created_at DESC").Find(&comments).Error
	if err != nil {
		return nil, err
	}

	return comments, nil
}

func (c CommentDB) Add(ctx context.Context, videoID uint, userID uint, content string) (uint, error) {
	newComment := Comments{
		UserID: userID,
		VideoID: videoID,
		Content: content,
	}

	err := c.DB.WithContext(ctx).Create(&newComment).Error
	if err != nil {
		return 0, err
	}

	return newComment.ID, nil
}

func (c CommentDB) Del(ctx context.Context, commentID uint) error {
	return c.DB.WithContext(ctx).Delete(&Comments{}, commentID).Error
}