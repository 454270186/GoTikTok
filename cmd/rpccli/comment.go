package rpccli

import (
	"log"

	"github.com/454270186/GoTikTok/rpc/comment/comments"
	"github.com/zeromicro/go-zero/core/discov"
	"github.com/zeromicro/go-zero/zrpc"
)

func NewCommentCli() comments.Comments {
	conn, err := zrpc.NewClient(zrpc.RpcClientConf{
		Etcd: discov.EtcdConf{
			Hosts: []string{"127.0.0.1:2379"},
			Key: "comment.rpc",
		},
	})
	if err != nil {
		log.Println(err)
		return nil
	}

	commentCli := comments.NewComments(conn)
	return commentCli
}