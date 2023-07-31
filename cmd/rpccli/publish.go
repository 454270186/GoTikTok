package rpccli

import (
	"log"

	"github.com/454270186/GoTikTok/rpc/publish/publishclient"
	"github.com/zeromicro/go-zero/core/discov"
	"github.com/zeromicro/go-zero/zrpc"
)

func NewPubCli() publishclient.Publish {
	conn, err := zrpc.NewClient(zrpc.RpcClientConf{
		Etcd: discov.EtcdConf{
			Hosts: []string{"127.0.0.1:2379"},
			Key: "publish.rpc",
		},
	})
	if err != nil {
		log.Println(err)
		return nil
	}

	pubCli := publishclient.NewPublish(conn)
	return pubCli
}