package logic

import (
	"context"
	"strconv"

	"github.com/454270186/GoTikTok/dal/pack"
	"github.com/454270186/GoTikTok/pkg/auth"
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
	// encoded user's password
	encrypted, _ := auth.GetHashedPwd(in.Password)

	newUserID, err := pack.CreateNewUser(in.Username, string(encrypted))
	if err != nil {
		return nil, err
	}

	token, _ := auth.NewTokenByUserID(newUserID)
	if token == "" {
		return &user.RegisterRes{
			StatusCode: -1,
			UserId: strconv.FormatUint(uint64(newUserID), 10),
			Token: "",
		}, nil
	}

	return &user.RegisterRes{
		StatusCode: 0,
		UserId: strconv.FormatUint(uint64(newUserID), 10),
		Token: token,
	}, nil
}
