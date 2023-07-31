package rpccli

import (
	"log"

	"github.com/454270186/GoTikTok/rpc/user/userservice"
	"github.com/zeromicro/go-zero/core/discov"
	"github.com/zeromicro/go-zero/zrpc"
)

// user rpc client
func NewUserCli() userservice.UserService {
	conn, err := zrpc.NewClient(zrpc.RpcClientConf{
		Etcd: discov.EtcdConf{
			Hosts: []string{"127.0.0.1:2379"},
			Key: "user.rpc",
		},
	})
	if err != nil {
		log.Println(err)
		return nil
	}

	userCli := userservice.NewUserService(conn)
	return userCli
} 