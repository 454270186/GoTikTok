package handler

import (
	// "log"
	"net/http"

	// "github.com/454270186/GoTikTok/cmd/model"
	"github.com/454270186/GoTikTok/cmd/rpccli"
	"github.com/454270186/GoTikTok/rpc/user/types/user"
	"github.com/454270186/GoTikTok/rpc/user/userservice"
	"github.com/gin-gonic/gin"
)


type UserHandler struct {
	userRpcCli userservice.UserService
}

func NewUserHandler() *UserHandler {
	return &UserHandler{
		userRpcCli: rpccli.NewUserCli(),
	}
}

func (u UserHandler) Register(c *gin.Context) {
	username := c.Query("username")
	password := c.Query("password")

	if len(username) == 0 || len(password) == 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status_code": -1,
			"status_msg": "username and password cannot be empty",
		})
	}

	in := user.RegisterReq{
		Username: username,
		Password: password,
	}

	resp, err := u.userRpcCli.Register(c.Copy(), &in)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status_code": -1,
			"status_msg": err.Error(),
		})
		return
	} else if resp.StatusCode != 0 {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status_code": resp.StatusCode,
			"status_msg": "unexpect internal error",
		})
		return
	}

	c.JSON(200, gin.H{
		"status_code": 0,
		"status_msg": "user register successfully",
		"user_id": resp.UserId,
		"token": "sssssssssssssss",
	})
}