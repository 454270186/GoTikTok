package logic

import (
	"context"
	"strconv"

	"github.com/454270186/GoTikTok/dal/pack"
	"github.com/454270186/GoTikTok/pkg/auth"
	"github.com/454270186/GoTikTok/pkg/tracing"
	"github.com/454270186/GoTikTok/rpc/user/internal/svc"
	"github.com/454270186/GoTikTok/rpc/user/types/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type RegisterLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

var userRPCTracer tracing.JaegerTracer

func init() {
	userRPCTracer = tracing.NewTracer("User-RPC-Service", "127.0.0.1:6831")
}

func NewRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegisterLogic {
	return &RegisterLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *RegisterLogic) Register(in *user.RegisterReq) (*user.RegisterRes, error) {
	span, traceID, _ := userRPCTracer.StartSpan("RPC_Register", in.TraceID, false)
	defer userRPCTracer.FinishSpan(span)

	// encoded user's password
	encrypted, _ := auth.GetHashedPwd(in.Password)

	mysqlSpan, _, _ := userRPCTracer.StartSpan("DB: Create new user", traceID, false)
	newUserID, err := pack.CreateNewUser(in.Username, string(encrypted))
	if err != nil {
		userRPCTracer.FinishSpan(mysqlSpan)
		return nil, err
	}
	userRPCTracer.FinishSpan(mysqlSpan)

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
