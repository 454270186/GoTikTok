package logic

import (
	"context"

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
	// todo: add your logic here and delete this line

	return &favorite.FavoriteActionRes{}, nil
}
