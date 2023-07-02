package logic

import (
	"context"

	"github.com/454270186/GoTikTok/dal/pack"
	"github.com/454270186/GoTikTok/rpc/comment/internal/svc"
	"github.com/454270186/GoTikTok/rpc/comment/types/comment"

	"github.com/zeromicro/go-zero/core/logx"
)

type DelCommentLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDelCommentLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DelCommentLogic {
	return &DelCommentLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DelCommentLogic) DelComment(in *comment.DelCommentReq) (*comment.DelCommentRes, error) {
	err := pack.DelComment(in.CommentId)
	if err != nil {
		return nil, err
	}

	return &comment.DelCommentRes{
		StatusCode: 0,
	}, nil
}
