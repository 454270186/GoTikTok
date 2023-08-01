package rmodel

type FavoriteCache struct {
	VideoID    uint `json:"video_id"`
	UserID     uint `json:"user_id"`
	ActionType uint `json:"action_type"`
	CreatedAt  uint `json:"created_at"`
}

// Key: user::{id}
type UserCache struct {
	ID             uint   `json:"id"`
	Username       string `json:"username"`
	Password       string `json:"password"`
	FollowingCount int64  `json:"following_count"`
	FollowerCount  int64  `json:"follower_count"`
}
