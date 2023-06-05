package rpccli

import (
	"github.com/454270186/GoTikTok/rpc/user/userservice"
	"github.com/zeromicro/go-zero/core/discov"
	"github.com/zeromicro/go-zero/zrpc"
)

// user rpc client
func NewUserCli() userservice.UserService {
	conn := zrpc.MustNewClient(zrpc.RpcClientConf{
		Etcd: discov.EtcdConf{
			Hosts: []string{"127.0.0.1:2379"},
			Key: "user.rpc",
		},
	})

	userCli := userservice.NewUserService(conn)
	return userCli
} 