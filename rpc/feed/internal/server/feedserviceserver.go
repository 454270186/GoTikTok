// Code generated by goctl. DO NOT EDIT.
// Source: feed.proto

package server

import (
	"context"

	"github.com/454270186/GoTikTok/rpc/feed/internal/logic"
	"github.com/454270186/GoTikTok/rpc/feed/internal/svc"
	"github.com/454270186/GoTikTok/rpc/feed/types/feed"
)

type FeedServiceServer struct {
	svcCtx *svc.ServiceContext
	feed.UnimplementedFeedServiceServer
}

func NewFeedServiceServer(svcCtx *svc.ServiceContext) *FeedServiceServer {
	return &FeedServiceServer{
		svcCtx: svcCtx,
	}
}

func (s *FeedServiceServer) GetUserFeed(ctx context.Context, in *feed.FeedReq) (*feed.FeedRes, error) {
	l := logic.NewGetUserFeedLogic(ctx, s.svcCtx)
	return l.GetUserFeed(in)
}
