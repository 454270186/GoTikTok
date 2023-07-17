package rmodel

type FavoriteCache struct {
	VideoID    uint `json:"video_id"`
	UserID     uint `json:"user_id"`
	ActionType uint `json:"action_type"`
	CreatedAt  uint `json:"created_at"`
}