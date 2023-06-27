package main

import (
	"flag"
	"fmt"

	"github.com/454270186/GoTikTok/rpc/feed/internal/config"
	"github.com/454270186/GoTikTok/rpc/feed/internal/server"
	"github.com/454270186/GoTikTok/rpc/feed/internal/svc"
	"github.com/454270186/GoTikTok/rpc/feed/types/feed"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var configFile = flag.String("f", "etc/feed.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)
	ctx := svc.NewServiceContext(c)

	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		feed.RegisterFeedServiceServer(grpcServer, server.NewFeedServiceServer(ctx))

		if c.Mode == service.DevMode || c.Mode == service.TestMode {
			reflection.Register(grpcServer)
		}
	})
	defer s.Stop()

	fmt.Printf("Starting rpc server at %s...\n", c.ListenOn)
	s.Start()
}
