package logic

import (
	"context"

	"github.com/454270186/GoTikTok/rpc/favorite/internal/svc"
	"github.com/454270186/GoTikTok/rpc/favorite/types/favorite"

	"github.com/zeromicro/go-zero/core/logx"
)

type FavoriteListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewFavoriteListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FavoriteListLogic {
	return &FavoriteListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *FavoriteListLogic) FavoriteList(in *favorite.FavoriteListReq) (*favorite.FavoriteListReq, error) {
	// todo: add your logic here and delete this line

	return &favorite.FavoriteListReq{}, nil
}
