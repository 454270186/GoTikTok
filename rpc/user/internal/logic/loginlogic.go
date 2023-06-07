package logic

import (
	"context"
	"strconv"

	"github.com/454270186/GoTikTok/pkg/auth"
	"github.com/454270186/GoTikTok/rpc/user/internal/svc"
	"github.com/454270186/GoTikTok/rpc/user/types/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type LoginLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *LoginLogic) Login(in *user.LoginReq) (*user.LoginRes, error) {
	// todo: add your logic here and delete this line
	userID, isExist, err := UserDB.IsUserExist(l.ctx, in.Username)
	if err != nil {
		return nil, err
	}

	if !isExist {
		return &user.LoginRes{
			StatusCode: -1,
		}, nil
	}

	token, _ := auth.NewTokenByUserID(userID)
	if token == "" {
		return &user.LoginRes{
			StatusCode: -1,
			UserId: strconv.FormatUint(uint64(userID), 10),
			Token: "",
		}, nil
	}

	return &user.LoginRes{
		StatusCode: 0,
		UserId: strconv.FormatUint(uint64(userID), 10),
		Token: token,
	}, nil
}
