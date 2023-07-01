package handler

import (
	"bytes"
	"context"
	"io"
	"net/http"

	"github.com/454270186/GoTikTok/cmd/httpres"
	"github.com/454270186/GoTikTok/cmd/rpccli"
	"github.com/454270186/GoTikTok/pkg/auth"
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
		httpres.SendRpcError(c, err.Error())
		return
	} else if resp.StatusCode != 0 {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status_code": -1,
			"status_msg": "resp.StatusCode != 0",
		})
		return
	}

	httpres.SendResponse(c, "success", gin.H{
		"video_list": resp.VideoList,
	})
}

func (p PublishHandler) Action(c *gin.Context) {
	title := c.PostForm("title")
	tokenStr := c.PostForm("token")
	fileHeader, err := c.FormFile("data")
	if err != nil {
		httpres.SendError(c, err.Error())
		return
	}

	file, err := fileHeader.Open()
	if err != nil {
		httpres.SendError(c, err.Error())
		return
	}
	defer file.Close()

	buf := bytes.NewBuffer(nil)
	if _, err := io.Copy(buf, file); err != nil {
		httpres.SendError(c, err.Error())
		return
	}

	// get user id from token
	uid, err := auth.GetUIDFromToken(tokenStr)
	if err != nil {
		httpres.SendError(c, err.Error())
		return
	}

	in := publishclient.PublishActionReq{
		Title: title,
		Data: buf.Bytes(),
		Uid: int64(uid),
	}
	
	_, err = p.pubRpcCli.PublishAction(c.Copy(), &in)
	if err != nil {
		httpres.SendRpcError(c, err.Error())
		return
	}

	httpres.SendResponse(c, "publish successfully")
}