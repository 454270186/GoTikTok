// Code generated by goctl. DO NOT EDIT.
// Source: feed.proto

package feedservice

import (
	"context"

	"github.com/454270186/GoTikTok/rpc/feed/types/feed"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	FeedReq = feed.FeedReq
	FeedRes = feed.FeedRes
	User    = feed.User
	Video   = feed.Video

	FeedService interface {
		GetUserFeed(ctx context.Context, in *FeedReq, opts ...grpc.CallOption) (*FeedRes, error)
	}

	defaultFeedService struct {
		cli zrpc.Client
	}
)

func NewFeedService(cli zrpc.Client) FeedService {
	return &defaultFeedService{
		cli: cli,
	}
}

func (m *defaultFeedService) GetUserFeed(ctx context.Context, in *FeedReq, opts ...grpc.CallOption) (*FeedRes, error) {
	client := feed.NewFeedServiceClient(m.cli.Conn())
	return client.GetUserFeed(ctx, in, opts...)
}