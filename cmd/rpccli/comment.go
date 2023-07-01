package rpccli

import (
	"github.com/454270186/GoTikTok/rpc/comment/comments"
	"github.com/zeromicro/go-zero/core/discov"
	"github.com/zeromicro/go-zero/zrpc"
)

func NewCommentCli() comments.Comments {
	conn := zrpc.MustNewClient(zrpc.RpcClientConf{
		Etcd: discov.EtcdConf{
			Hosts: []string{"127.0.0.1:2379"},
			Key: "comment.rpc",
		},
	})

	commentCli := comments.NewComments(conn)
	return commentCli
}