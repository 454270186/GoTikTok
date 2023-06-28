package logic

import (
	"context"
	"errors"
	"strconv"

	"github.com/454270186/GoTikTok/dal/pack"
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

func (l *FavoriteListLogic) FavoriteList(in *favorite.FavoriteListReq) (*favorite.FavoriteListRes, error) {
	userID, err := strconv.ParseUint(in.UserId, 10, 64)
	if err != nil {
		return nil, errors.New("userID parse failed")
	}
	
	videos, err := pack.GetFavVideosByUserID(uint(userID))
	if err != nil {
		return nil, err
	}

	return &favorite.FavoriteListRes{
		StatusCode: 0,
		VideoList: videos,
	}, nil
}
