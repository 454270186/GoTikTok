package rpccli

import (
	"github.com/454270186/GoTikTok/rpc/feed/feedservice"
	"github.com/zeromicro/go-zero/core/discov"
	"github.com/zeromicro/go-zero/zrpc"
)

func NewFeedCli() feedservice.FeedService {
	conn := zrpc.MustNewClient(zrpc.RpcClientConf{
		Etcd: discov.EtcdConf{
			Hosts: []string{"127.0.0.1:2379"},
			Key: "feeed.rpc",
		},
	})

	feedCli := feedservice.NewFeedService(conn)
	return feedCli
}