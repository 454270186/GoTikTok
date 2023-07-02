package logic

import (
	"context"

	"github.com/454270186/GoTikTok/dal/pack"
	"github.com/454270186/GoTikTok/rpc/comment/internal/svc"
	"github.com/454270186/GoTikTok/rpc/comment/types/comment"

	"github.com/zeromicro/go-zero/core/logx"
)

type CommentListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCommentListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CommentListLogic {
	return &CommentListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CommentListLogic) CommentList(in *comment.CommentListReq) (*comment.CommentListRes, error) {
	commentList, err := pack.GetCommentByVideoID(in.VideoId)
	if err != nil {
		return nil, err
	}

	return &comment.CommentListRes{
		StatusCode: 0,
		CommentList: commentList,
	}, nil
}
