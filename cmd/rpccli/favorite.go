package rpccli

import (
	"log"

	"github.com/454270186/GoTikTok/rpc/favorite/favoriteclient"
	"github.com/zeromicro/go-zero/core/discov"
	"github.com/zeromicro/go-zero/zrpc"
)

func NewFavCli() favoriteclient.Favorite {
	conn, err := zrpc.NewClient(zrpc.RpcClientConf{
		Etcd: discov.EtcdConf{
			Hosts: []string{"127.0.0.1:2379"},
			Key: "favorite.rpc",
		},
	})
	if err != nil {
		log.Println(err)
		return nil
	}

	favCli := favoriteclient.NewFavorite(conn)
	return favCli
}