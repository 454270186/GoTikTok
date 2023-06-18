package dal

import "time"

type User struct {
	ID             uint   `gorm:"primarykey"`
	Username       string `gorm:"column:username"`
	Password       string `gorm:"column:password"`
	FollowingCount int64  `gorm:"default:0;column:following_count"`
	FollowerCount  int64  `gorm:"default:0;column:follower_count"`
	CreatedAt      time.Time
	UpdatedAt      time.Time
}

type Video struct {
	ID            uint   `gorm:"primarykey"`
	AuthorID      uint   `gorm:"column:author_id"`
	PlayURL       string `gorm:"column:play_url"`
	CoverURL      string `gorm:"column:cover_url"`
	FavoriteCount int64  `gorm:"column:favorite_count"`
	CommentCount  int64  `gorm:"column:comment_count"`
	Title         string `gorm:"column:title"`
	CreatedAt     time.Time
	UpdatedAt     time.Time
}