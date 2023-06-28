// Code generated by goctl. DO NOT EDIT.
// Source: favorite.proto

package favoriteclient

import (
	"context"

	"github.com/454270186/GoTikTok/rpc/favorite/types/favorite"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	FavoriteActionReq = favorite.FavoriteActionReq
	FavoriteActionRes = favorite.FavoriteActionRes
	FavoriteListReq   = favorite.FavoriteListReq
	FavoriteListRes   = favorite.FavoriteListRes
	User              = favorite.User
	Video             = favorite.Video

	Favorite interface {
		FavoriteList(ctx context.Context, in *FavoriteListReq, opts ...grpc.CallOption) (*FavoriteListRes, error)
		FavoriteAction(ctx context.Context, in *FavoriteActionReq, opts ...grpc.CallOption) (*FavoriteActionRes, error)
	}

	defaultFavorite struct {
		cli zrpc.Client
	}
)

func NewFavorite(cli zrpc.Client) Favorite {
	return &defaultFavorite{
		cli: cli,
	}
}

func (m *defaultFavorite) FavoriteList(ctx context.Context, in *FavoriteListReq, opts ...grpc.CallOption) (*FavoriteListRes, error) {
	client := favorite.NewFavoriteClient(m.cli.Conn())
	return client.FavoriteList(ctx, in, opts...)
}

func (m *defaultFavorite) FavoriteAction(ctx context.Context, in *FavoriteActionReq, opts ...grpc.CallOption) (*FavoriteActionRes, error) {
	client := favorite.NewFavoriteClient(m.cli.Conn())
	return client.FavoriteAction(ctx, in, opts...)
}
