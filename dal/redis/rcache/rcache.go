package rcache

type FavoriteCache struct {
	VideoID    uint `json:"video_id"`
	UserID     uint `json:"user_id"`
	ActionType uint `json:"action_type"`
	CreatedAt  uint `json:"created_at"`
}

// VideoCache key pattern is:
// {videoID}
// VideoCache value pattern is:
// {authorID}::{videoObjName}::{coverObjName}::{fav_cnt}::{cmt_cnt}::{title}

// Favorite key pattern is: 
// video::{videoID}::user::{userID}
// Favorite value pattern is:
// {action_type}::{created_at}

/* Read */
// Provide one interface to get favorite list by userID
// Step:
// 1. get a fav_videoID list (all favorite video by this user)
// 2. traverse fav_videoID list to get fav_video from videoCache, 
//    if not found in cache, get from DB, and then async add it into cache
// 3. pack user and fav_videos, send back to rpc caller

/* Write */
// Provide two interfaces to update favorite
// 1. LikeVideo
// 2. UnlikeVideo

