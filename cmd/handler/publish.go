package handler

import (
	"bytes"
	"context"
	"io"
	"log"
	"net/http"

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

func (p PublishHandler) Action(c *gin.Context) {
	title := c.PostForm("title")
	tokenStr := c.PostForm("token")
	fileHeader, err := c.FormFile("data")
	if err != nil {
		log.Println(err, "1111")
		c.JSON(http.StatusBadRequest, gin.H{
			"status_code": -1,
			"status_msg": err.Error(),
		})
		return
	}

	log.Println(title, tokenStr)

	file, err := fileHeader.Open()
	if err != nil {
		log.Println(err, "2222")
		c.JSON(http.StatusBadRequest, gin.H{
			"status_code": -1,
			"status_msg": err.Error(),
		})
		return
	}
	defer file.Close()

	buf := bytes.NewBuffer(nil)
	if _, err := io.Copy(buf, file); err != nil {
		log.Println(err, "3333")
		c.JSON(http.StatusBadRequest, gin.H{
			"status_code": -1,
			"status_msg": err.Error(),
		})
		return
	}

	// get user id from token
	uid, err := auth.GetUIDFromToken(tokenStr)
	if err != nil {
		log.Println(err, "4444")
		c.JSON(http.StatusBadRequest, gin.H{
			"status_code": -1,
			"status_msg": err.Error(),
		})
		return
	}

	in := publishclient.PublishActionReq{
		Title: title,
		Data: buf.Bytes(),
		Uid: int64(uid),
	}
	
	resp, err := p.pubRpcCli.PublishAction(c.Copy(), &in)
	if err != nil {
		log.Println(err, "5555")
		c.JSON(http.StatusBadRequest, gin.H{
			"status_code": -1,
			"status_msg": err.Error(),
		})
		return
	}

	c.JSON(200, resp)
}