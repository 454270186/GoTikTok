package handler

import (
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

}

func (f FavoriteHandler) Action(c *gin.Context) {
	
}