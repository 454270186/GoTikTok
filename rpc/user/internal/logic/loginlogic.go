package logic

import (
	"context"
	"errors"
	"strconv"

	"github.com/454270186/GoTikTok/dal/pack"
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
	span, traceID, _ := userRPCTracer.StartSpan("RPC: Login", in.TraceID, false)
	userRPCTracer.SpanSetTag(span, "rpc_addr", "0.0.0.0:8081")
	defer userRPCTracer.FinishSpan(span)

	mysqlSpan, _, _ := userRPCTracer.DB("Check user exist", traceID, false)
	userRPCTracer.SpanSetTag(mysqlSpan, "db_operation", "Query")
	userID, isExist, err := pack.IsUserExist(in.Username)
	if err != nil {
		userRPCTracer.FinishSpan(mysqlSpan)
		return nil, err
	}
	userRPCTracer.FinishSpan(mysqlSpan)

	if !isExist {
		return nil, errors.New("user is not exist")
	}

	mysqlSpan, _, _ = userRPCTracer.DB("Check user password", traceID, false)
	userRPCTracer.SpanSetTag(mysqlSpan, "db_operation", "Query")
	if isPwdOK := pack.CheckPassword(userID, in.Password); !isPwdOK {
		userRPCTracer.FinishSpan(mysqlSpan)
		return nil, errors.New("wrong password")
	}
	userRPCTracer.FinishSpan(mysqlSpan)

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
