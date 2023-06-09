package main

import (
	"flag"
	"fmt"

	"github.com/454270186/GoTikTok/dal"
	"github.com/454270186/GoTikTok/rpc/publish/internal/config"
	"github.com/454270186/GoTikTok/rpc/publish/internal/logic"
	"github.com/454270186/GoTikTok/rpc/publish/internal/server"
	"github.com/454270186/GoTikTok/rpc/publish/internal/svc"
	"github.com/454270186/GoTikTok/rpc/publish/types/publish"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var configFile = flag.String("f", "etc/publish.yaml", "the config file")

func main() {
	flag.Parse()

	// init DB
	logic.PublishDB = dal.NewPublishDB()
	logic.UserDB = dal.NewUserDB()

	var c config.Config
	conf.MustLoad(*configFile, &c)
	ctx := svc.NewServiceContext(c)

	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		publish.RegisterPublishServer(grpcServer, server.NewPublishServer(ctx))

		if c.Mode == service.DevMode || c.Mode == service.TestMode {
			reflection.Register(grpcServer)
		}
	})
	defer s.Stop()

	fmt.Printf("Starting rpc server at %s...\n", c.ListenOn)
	s.Start()
}
