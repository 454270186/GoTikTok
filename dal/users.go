package dal

import (
	"context"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username       string `gorm:"column:user_name"`
	Password       string `gorm:"column:password"`
	FollowingCount int64 `gorm:"default:0;column:following_count"`
	FollowerCount  int64 `gorm:"default:0;column:follower_count"`
}

func (u User) GetTableName() string {
	return "users"
}

type UserDB struct {
	DB *gorm.DB
}

func NewUserDB() UserDB {
	return UserDB{
		DB: newDB(),
	}
}

func (u UserDB) CreateUser(ctx context.Context, user *User) error {
	return u.DB.WithContext(ctx).Create(user).Error
}