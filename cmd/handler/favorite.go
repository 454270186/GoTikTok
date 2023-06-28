package handler

import (
	"net/http"

	"github.com/454270186/GoTikTok/cmd/rpccli"
	"github.com/454270186/GoTikTok/rpc/favorite/favoriteclient"
	"github.com/gin-gonic/gin"
)

type FavoriteHandler struct {
	favRpcCli favoriteclient.Favorite
}

func NewFavHandler() *FavoriteHandler {
	return &FavoriteHandler{
		favRpcCli: rpccli.NewFavCli(),
	}
}

func (f FavoriteHandler) List(c *gin.Context) {
	userIDStr := c.Query("user_id")
	if userIDStr == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"status_code": -1,
			"status_msg": "user id cannot be empty",
		})
		return
	}

	in := favoriteclient.FavoriteListReq{
		UserId: userIDStr,
	}

	resp, err := f.favRpcCli.FavoriteList(c.Copy(), &in)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status_code": -1,
			"status_msg": err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"status_code": 0,
		"status_msg": "successfully",
		"video_list": resp.VideoList,
	})
}

func (f FavoriteHandler) Action(c *gin.Context) {
	
}