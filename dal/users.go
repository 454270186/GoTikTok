package dal

import (
	"context"

	"gorm.io/gorm"
)

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
	err := u.DB.WithContext(ctx).Where("username = ?", username).First(&user).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return 0, false, nil
		}

		return 0, false, err
	}

	return user.ID, true, nil
}
