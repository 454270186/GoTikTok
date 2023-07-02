// Code generated by goctl. DO NOT EDIT.
// Source: comment.proto

package server

import (
	"context"

	"github.com/454270186/GoTikTok/rpc/comment/internal/logic"
	"github.com/454270186/GoTikTok/rpc/comment/internal/svc"
	"github.com/454270186/GoTikTok/rpc/comment/types/comment"
)

type CommentsServer struct {
	svcCtx *svc.ServiceContext
	comment.UnimplementedCommentsServer
}

func NewCommentsServer(svcCtx *svc.ServiceContext) *CommentsServer {
	return &CommentsServer{
		svcCtx: svcCtx,
	}
}

func (s *CommentsServer) CommentList(ctx context.Context, in *comment.CommentListReq) (*comment.CommentListRes, error) {
	l := logic.NewCommentListLogic(ctx, s.svcCtx)
	return l.CommentList(in)
}

func (s *CommentsServer) AddComment(ctx context.Context, in *comment.AddCommentReq) (*comment.AddCommentRes, error) {
	l := logic.NewAddCommentLogic(ctx, s.svcCtx)
	return l.AddComment(in)
}

func (s *CommentsServer) DelComment(ctx context.Context, in *comment.DelCommentReq) (*comment.DelCommentRes, error) {
	l := logic.NewDelCommentLogic(ctx, s.svcCtx)
	return l.DelComment(in)
}
