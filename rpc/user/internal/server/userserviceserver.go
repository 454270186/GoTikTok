// Code generated by goctl. DO NOT EDIT.
// Source: user.proto

package server

import (
	"context"

	"github.com/454270186/GoTikTok/rpc/user/internal/logic"
	"github.com/454270186/GoTikTok/rpc/user/internal/svc"
	"github.com/454270186/GoTikTok/rpc/user/types/user"
)

type UserServiceServer struct {
	svcCtx *svc.ServiceContext
	user.UnimplementedUserServiceServer
}

func NewUserServiceServer(svcCtx *svc.ServiceContext) *UserServiceServer {
	return &UserServiceServer{
		svcCtx: svcCtx,
	}
}

func (s *UserServiceServer) Register(ctx context.Context, in *user.RegisterReq) (*user.RegisterRes, error) {
	l := logic.NewRegisterLogic(ctx, s.svcCtx)
	return l.Register(in)
}

func (s *UserServiceServer) Login(ctx context.Context, in *user.LoginReq) (*user.LoginRes, error) {
	l := logic.NewLoginLogic(ctx, s.svcCtx)
	return l.Login(in)
}

func (s *UserServiceServer) GetUserById(ctx context.Context, in *user.GetUserReq) (*user.GetUserRes, error) {
	l := logic.NewGetUserByIdLogic(ctx, s.svcCtx)
	return l.GetUserById(in)
}

func (s *UserServiceServer) Refresh(ctx context.Context, in *user.RefreshReq) (*user.RefreshRes, error) {
	l := logic.NewRefreshLogic(ctx, s.svcCtx)
	return l.Refresh(in)
}
