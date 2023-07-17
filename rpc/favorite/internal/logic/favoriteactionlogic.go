package logic

import (
	"context"
	"encoding/json"
	"strconv"
	"time"

	// "github.com/454270186/GoTikTok/dal/pack"
	"github.com/454270186/GoTikTok/dal/redis/rmodel"
	"github.com/454270186/GoTikTok/pkg/rabbitmq"
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
	rmq := rabbitmq.NewSimple()

	videoID, err := strconv.ParseUint(in.VideoId, 10, 64)
	if err != nil {
		return nil, err
	}
	userID, err := strconv.ParseUint(in.UserId, 10, 64)
	if err != nil {
		return nil, err
	}
	
	switch in.ActionType {
	case "1":
		// like
		// err := pack.LikeVideo(in.UserId, in.VideoId)
		// if err != nil {
		// 	return nil, err
		// }

		newFavOpt := rmodel.FavoriteCache{
			VideoID: uint(videoID),
			UserID: uint(userID),
			ActionType: 1,
			CreatedAt: uint(time.Now().Unix()),
		}
		
		newFavOptJson, err := json.Marshal(newFavOpt)
		if err != nil {
			return nil, err
		}

		rmq.PubWithCtx(l.ctx, newFavOptJson)
		return &favorite.FavoriteActionRes{
			StatusCode: 0,
		}, nil

	case "2":
		// unlike
		// err := pack.UnlikeVideo(in.UserId, in.VideoId)
		// if err != nil {
		// 	return nil, err
		// }

		newFavOpt := rmodel.FavoriteCache{
			VideoID: uint(videoID),
			UserID: uint(userID),
			ActionType: 2,
			CreatedAt: uint(time.Now().Unix()),
		}
		
		newFavOptJson, err := json.Marshal(newFavOpt)
		if err != nil {
			return nil, err
		}

		rmq.PubWithCtx(l.ctx, newFavOptJson)
		return &favorite.FavoriteActionRes{
			StatusCode: 0,
		}, nil
	}

	return &favorite.FavoriteActionRes{
		StatusCode: 0,
	}, nil
}
