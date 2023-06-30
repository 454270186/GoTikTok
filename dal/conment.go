package dal

import "gorm.io/gorm"

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