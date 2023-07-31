package logic

import (
	"context"
	"strconv"

	"github.com/454270186/GoTikTok/dal/pack"
	"github.com/454270186/GoTikTok/pkg/tracing"
	"github.com/454270186/GoTikTok/rpc/publish/internal/svc"
	"github.com/454270186/GoTikTok/rpc/publish/types/publish"

	"github.com/zeromicro/go-zero/core/logx"
)

var pubRPCTracer tracing.JaegerTracer

type PublishListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func init() {
	pubRPCTracer = tracing.NewTracer("Publish-RPC-Service", "127.0.0.1:6831")
}

func NewPublishListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PublishListLogic {
	return &PublishListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *PublishListLogic) PublishList(in *publish.PublishListReq) (*publish.PublishListRes, error) {
	span, traceID, _ := pubRPCTracer.StartSpan("RPC: List", in.TraceID, false)
	defer pubRPCTracer.FinishSpan(span)
	pubRPCTracer.SpanSetTag(span, "rpc_addr", "0.0.0.0:8080")

	userID, err := strconv.ParseUint(in.UserId, 10, 64)
	if err != nil {
		return &publish.PublishListRes{StatusCode: -1}, err
	}

	mysqlSpan, _, _ := pubRPCTracer.DB("GetVideoList", traceID, false)
	pubRPCTracer.SpanSetTag(mysqlSpan, "db_operation", "Query")
	respVideoList, err := pack.GetVideoList(uint(userID))
	if err != nil {
		pubRPCTracer.FinishSpan(mysqlSpan)
		return nil, err
	}
	pubRPCTracer.FinishSpan(mysqlSpan)

	return &publish.PublishListRes{
		StatusCode: 0,
		VideoList:  respVideoList,
	}, nil
}
