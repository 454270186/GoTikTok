package logic

import (
	"context"
	"strconv"

	"github.com/454270186/GoTikTok/dal/pack"
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
	span, traceID, _ := userRPCTracer.StartSpan("RPC: Get User by ID", in.TraceID, false)
	userRPCTracer.SpanSetTag(span, "rpc_addr", "0.0.0.0:8081")
	defer userRPCTracer.FinishSpan(span)

	userId, err := strconv.ParseUint(in.UserId, 10, 64)
	if err != nil {
		return nil, err
	}

	mysqlSpan, _, _ := userRPCTracer.DB("GetUserByID", traceID, false)
	userRPCTracer.SpanSetTag(mysqlSpan, "db_operation", "Query")
	gotUser, err := pack.GetUserByID(userId)
	if err != nil {
		userRPCTracer.FinishSpan(mysqlSpan)
		return nil, err
	}
	userRPCTracer.FinishSpan(mysqlSpan)

	return &user.GetUserRes{
		StatusCode: 0,
		User: gotUser,
	}, nil
}
