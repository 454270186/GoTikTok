package logic

import (
	"context"
	"strconv"

	"github.com/454270186/GoTikTok/dal"
	"github.com/454270186/GoTikTok/pkg/auth"
	"github.com/454270186/GoTikTok/rpc/user/internal/svc"
	"github.com/454270186/GoTikTok/rpc/user/types/user"

	"github.com/zeromicro/go-zero/core/logx"
)

var UserDB dal.UserDB

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
	newUser := dal.User{
		Username: in.Username,
		Password: in.Password,
	}

	err := UserDB.CreateUser(l.ctx, &newUser)
	if err != nil {
		return &user.RegisterRes{
			StatusCode: -1,
		}, err
	}

	token, _ := auth.NewTokenByUserID(newUser.ID)
	if token == "" {
		return &user.RegisterRes{
			StatusCode: -1,
			UserId: strconv.FormatUint(uint64(newUser.ID), 10),
			Token: "",
		}, nil
	}

	return &user.RegisterRes{
		StatusCode: 0,
		UserId: strconv.FormatUint(uint64(newUser.ID), 10),
		Token: token,
	}, nil
}
