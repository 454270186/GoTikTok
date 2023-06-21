package dal

import "gorm.io/gorm"

func (f Favorite) GetTableName() string {
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

