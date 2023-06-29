package httpres

import (
	"github.com/gin-gonic/gin"
)

type DefaultRes struct {
	StatusCode int    `json:"status_code"`
	StatusMsg  string `json:"status_msg"`
}

func SendResponse(c *gin.Context, msg string) {
	c.JSON(200, DefaultRes{StatusCode: 0, StatusMsg: msg})
}

func SendError(c *gin.Context, errMsg string) {
	c.JSON(400, DefaultRes{StatusCode: -1, StatusMsg: errMsg})
}