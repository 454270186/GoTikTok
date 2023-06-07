package logic

import (
	"context"
	"strconv"

	"github.com/454270186/GoTikTok/rpc/user/internal/svc"
	"github.com/454270186/GoTikTok/rpc/user/types/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserByIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUserByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserByIdLogic {
	return &GetUserByIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetUserByIdLogic) GetUserById(in *user.GetUserReq) (*user.GetUserRes, error) {
	// todo: add your logic here and delete this line
	userId, err := strconv.ParseUint(in.UserId, 10, 64)
	if err != nil {
		return nil, err
	}

	DBUser, err := UserDB.GetById(l.ctx, uint(userId))
	if err != nil {
		return nil, err
	}

	gotUser := user.User{
		Id: int64(DBUser.ID),
		Name: DBUser.Username,
		FollowCount: DBUser.FollowerCount,
		FollowerCount: DBUser.FollowerCount,
	}

	return &user.GetUserRes{
		StatusCode: 0,
		User: &gotUser,
	}, nil
}
