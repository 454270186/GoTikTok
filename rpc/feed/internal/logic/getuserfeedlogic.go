package logic

import (
	"context"
	"github.com/454270186/GoTikTok/dal/pack"
	"github.com/454270186/GoTikTok/rpc/feed/internal/svc"
	"github.com/454270186/GoTikTok/rpc/feed/types/feed"

	"github.com/zeromicro/go-zero/core/logx"
)

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
	videoList, nextTime, err := pack.GetVideoListByTime(LIMIT, in.LastestTime)
	if err != nil {
		return nil, err
	}

	return &feed.FeedRes{
		StatusCode: 0,
		VideoList: videoList,
		NextTime: nextTime,
	}, nil
}
