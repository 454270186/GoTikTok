package logic

import (
	"context"
	"time"

	"github.com/454270186/GoTikTok/dal"
	"github.com/454270186/GoTikTok/rpc/feed/internal/svc"
	"github.com/454270186/GoTikTok/rpc/feed/types/feed"

	"github.com/zeromicro/go-zero/core/logx"
)

var FeedDB dal.FeedDB
var UserDB dal.UserDB
const LIMIT = 15 // 最多返回视频数

type GetUserFeedLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUserFeedLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserFeedLogic {
	return &GetUserFeedLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetUserFeedLogic) GetUserFeed(in *feed.FeedReq) (*feed.FeedRes, error) {
	// todo: add your logic here and delete this line
	dalVideos, err :=  FeedDB.GetVideoLists(l.ctx, LIMIT, in.LastestTime)
	if err != nil {
		return nil, err
	}

	var nextTime int64
	var videoLists []*feed.Video
	if len(dalVideos) == 0 {
		nextTime = time.Now().UnixMilli()
		return &feed.FeedRes{
			StatusCode: 0,
			VideoList: videoLists,
			NextTime: nextTime,
		}, nil
	} else {
		nextTime = dalVideos[len(dalVideos)-1].UpdatedAt.UnixMilli()
	}

	videoLists, err = convertVideoLists(dalVideos)
	if err != nil {
		return nil, err
	}	

	return &feed.FeedRes{
		StatusCode: 0,
		VideoList: videoLists,
		NextTime: nextTime,
	}, nil
}

// Convert dal videos model to RPC video model
func convertVideoLists(dalVideos []*dal.Video) ([]*feed.Video, error) {
	rpcVideos := make([]*feed.Video, 0)
	for _, dalVideo := range dalVideos {
		author, err := UserDB.GetById(context.Background(), dalVideo.AuthorID)
		if err != nil {
			return nil, err
		}

		v := feed.Video{
			Id:            int64(dalVideo.ID),
			Author:        convertUser(author),
			PlayUrl:       dalVideo.PlayURL,
			CoverUrl:      dalVideo.CoverURL,
			FavoriteCount: dalVideo.FavoriteCount,
			CommentCount:  dalVideo.CommentCount,
			IsFavorite:    false,
			Title:         dalVideo.Title,
		}

		rpcVideos = append(rpcVideos, &v)
	}

	return rpcVideos, nil
}

// Convert dal.User to RPC feed.User
func convertUser(dalUser *dal.User) *feed.User {
	if dalUser == nil {
		return nil
	}

	return &feed.User{
		Id:            int64(dalUser.ID),
		Name:          dalUser.Username,
		FollowCount:   dalUser.FollowingCount,
		FollowerCount: dalUser.FollowerCount,
		IsFollow:      true,
	}
}
