package logic

import (
	"context"
	"strconv"

	"github.com/454270186/GoTikTok/dal"
	"github.com/454270186/GoTikTok/rpc/user/internal/svc"
	"github.com/454270186/GoTikTok/rpc/user/types/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type RegisterLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegisterLogic {
	return &RegisterLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *RegisterLogic) Register(in *user.RegisterReq) (*user.RegisterRes, error) {
	// todo: add your logic here and delete this line
	userDB := dal.NewUserDB()

	newUser := dal.User{
		Username: in.Username,
		Password: in.Password,
	}

	err := userDB.CreateUser(l.ctx, &newUser)
	if err != nil {
		return &user.RegisterRes{
			StatusCode: -1,
		}, err
	}

	return &user.RegisterRes{
		StatusCode: 0,
		UserId: strconv.FormatUint(uint64(newUser.ID), 10),
	}, nil
}
