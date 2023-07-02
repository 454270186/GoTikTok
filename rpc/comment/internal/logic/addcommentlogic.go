package logic

import (
	"context"

	"github.com/454270186/GoTikTok/dal/pack"
	"github.com/454270186/GoTikTok/rpc/comment/internal/svc"
	"github.com/454270186/GoTikTok/rpc/comment/types/comment"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddCommentLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddCommentLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddCommentLogic {
	return &AddCommentLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *AddCommentLogic) AddComment(in *comment.AddCommentReq) (*comment.AddCommentRes, error) {
	newCommentID, err := pack.AddComment(in.VideoId, in.UserId, in.CommentText)
	if err != nil {
		return nil, err
	}

	newComment, err := pack.GetCommentByID(newCommentID)
	if err != nil {
		return nil, err
	}

	return &comment.AddCommentRes{
		StatusCode: 0,
		Comment: newComment,
	}, nil
}
