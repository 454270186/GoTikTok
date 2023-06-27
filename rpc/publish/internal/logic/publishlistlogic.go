package logic

import (
	"context"
	"strconv"

	"github.com/454270186/GoTikTok/dal/pack"
	"github.com/454270186/GoTikTok/rpc/publish/internal/svc"
	"github.com/454270186/GoTikTok/rpc/publish/types/publish"

	"github.com/zeromicro/go-zero/core/logx"
)

type PublishListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewPublishListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PublishListLogic {
	return &PublishListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *PublishListLogic) PublishList(in *publish.PublishListReq) (*publish.PublishListRes, error) {
	// todo: add your logic here and delete this line
	userID, err := strconv.ParseUint(in.UserId, 10, 64)
	if err != nil {
		return &publish.PublishListRes{StatusCode: -1}, err
	}

	respVideoList, err := pack.GetVideoList(uint(userID))
	if err != nil {
		return nil, err
	}

	return &publish.PublishListRes{
		StatusCode: 0,
		VideoList:  respVideoList,
	}, nil
}
