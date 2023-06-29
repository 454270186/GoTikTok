package logic

import (
	"context"

	"github.com/454270186/GoTikTok/dal/pack"
	"github.com/454270186/GoTikTok/rpc/favorite/internal/svc"
	"github.com/454270186/GoTikTok/rpc/favorite/types/favorite"

	"github.com/zeromicro/go-zero/core/logx"
)

type FavoriteActionLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewFavoriteActionLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FavoriteActionLogic {
	return &FavoriteActionLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *FavoriteActionLogic) FavoriteAction(in *favorite.FavoriteActionReq) (*favorite.FavoriteActionRes, error) {
	switch in.ActionType {
	case "1":
		// like
		err := pack.LikeVideo(in.UserId, in.VideoId)
		if err != nil {
			return nil, err
		}

	case "2":
		// unlike
		err := pack.UnlikeVideo(in.UserId, in.VideoId)
		if err != nil {
			return nil, err
		}
	}

	return &favorite.FavoriteActionRes{
		StatusCode: 0,
	}, nil
}
