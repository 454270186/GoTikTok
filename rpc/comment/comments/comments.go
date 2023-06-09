// Code generated by goctl. DO NOT EDIT.
// Source: comment.proto

package comments

import (
	"context"

	"github.com/454270186/GoTikTok/rpc/comment/types/comment"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	AddCommentReq  = comment.AddCommentReq
	AddCommentRes  = comment.AddCommentRes
	Comment        = comment.Comment
	CommentListReq = comment.CommentListReq
	CommentListRes = comment.CommentListRes
	DelCommentReq  = comment.DelCommentReq
	DelCommentRes  = comment.DelCommentRes
	User           = comment.User

	Comments interface {
		CommentList(ctx context.Context, in *CommentListReq, opts ...grpc.CallOption) (*CommentListRes, error)
		AddComment(ctx context.Context, in *AddCommentReq, opts ...grpc.CallOption) (*AddCommentRes, error)
		DelComment(ctx context.Context, in *DelCommentReq, opts ...grpc.CallOption) (*DelCommentRes, error)
	}

	defaultComments struct {
		cli zrpc.Client
	}
)

func NewComments(cli zrpc.Client) Comments {
	return &defaultComments{
		cli: cli,
	}
}

func (m *defaultComments) CommentList(ctx context.Context, in *CommentListReq, opts ...grpc.CallOption) (*CommentListRes, error) {
	client := comment.NewCommentsClient(m.cli.Conn())
	return client.CommentList(ctx, in, opts...)
}

func (m *defaultComments) AddComment(ctx context.Context, in *AddCommentReq, opts ...grpc.CallOption) (*AddCommentRes, error) {
	client := comment.NewCommentsClient(m.cli.Conn())
	return client.AddComment(ctx, in, opts...)
}

func (m *defaultComments) DelComment(ctx context.Context, in *DelCommentReq, opts ...grpc.CallOption) (*DelCommentRes, error) {
	client := comment.NewCommentsClient(m.cli.Conn())
	return client.DelComment(ctx, in, opts...)
}
