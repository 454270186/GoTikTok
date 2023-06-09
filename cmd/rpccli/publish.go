package rpccli

import (
	"github.com/454270186/GoTikTok/rpc/publish/publishclient"
	"github.com/zeromicro/go-zero/core/discov"
	"github.com/zeromicro/go-zero/zrpc"
)

func NewPubCli() publishclient.Publish {
	conn := zrpc.MustNewClient(zrpc.RpcClientConf{
		Etcd: discov.EtcdConf{
			Hosts: []string{"127.0.0.1:2379"},
			Key: "publish.rpc",
		},
	})

	pubCli := publishclient.NewPublish(conn)
	return pubCli
}