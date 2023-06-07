package dal

import (
	"context"
	"database/sql"
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID             uint   `gorm:"primarykey"`
	Username       string `gorm:"column:user_name"`
	Password       string `gorm:"column:password"`
	FollowingCount int64  `gorm:"default:0;column:following_count"`
	FollowerCount  int64  `gorm:"default:0;column:follower_count"`
	CreatedAt      time.Time
	UpdatedAt      time.Time
	DeletedAt      sql.NullTime `gorm:"index"`
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

func (u UserDB) GetById(ctx context.Context, userId uint) (*User, error) {
	var user User
	err := u.DB.WithContext(ctx).First(&user, userId).Error
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (u UserDB) IsUserExist(ctx context.Context, username string) (uint, bool, error) {
	var user User
	err := u.DB.WithContext(ctx).Where("name = ?", username).First(&user).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return 0, false, nil
		}

		return 0, false, err
	}

	return user.ID, true, nil
}
