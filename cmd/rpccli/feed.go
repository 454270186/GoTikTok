package rpccli

import (
	"log"

	"github.com/454270186/GoTikTok/rpc/feed/feedservice"
	"github.com/zeromicro/go-zero/core/discov"
	"github.com/zeromicro/go-zero/zrpc"
)

func NewFeedCli() feedservice.FeedService {
	conn, err := zrpc.NewClient(zrpc.RpcClientConf{
		Etcd: discov.EtcdConf{
			Hosts: []string{"127.0.0.1:2379"},
			Key: "feed.rpc",
		},
	})
	if err != nil {
		log.Println(err)
		return nil
	}

	feedCli := feedservice.NewFeedService(conn)
	return feedCli
}