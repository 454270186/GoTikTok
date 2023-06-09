package handler

import (
	"context"
	"log"
	"net/http"

	"github.com/454270186/GoTikTok/cmd/rpccli"
	"github.com/454270186/GoTikTok/rpc/publish/publishclient"
	"github.com/gin-gonic/gin"
)

type PublishHandler struct {
	pubRpcCli publishclient.Publish
}

func NewPubHandler() *PublishHandler {
	return &PublishHandler{
		pubRpcCli: rpccli.NewPubCli(),
	}
}

func (p PublishHandler) List(c *gin.Context) {
	userID := c.Query("user_id")

	in := publishclient.PublishListReq {
		UserId: userID,
	}

	resp, err := p.pubRpcCli.PublishList(context.Background(), &in)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status_code": -1,
			"status_msg": err.Error(),
		})
		return
	} else if resp.StatusCode != 0 {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status_code": -1,
			"status_msg": "resp.StatusCode != 0",
		})
		return
	}

	log.Println(resp)

	c.JSON(200, gin.H{
		"status_code": 0,
		"status_msg": "success",
		"video_list": resp.VideoList,
	})
}