package httpres

import (
	"github.com/gin-gonic/gin"
)

type DefaultRes struct {
	StatusCode int    `json:"status_code"`
	StatusMsg  string `json:"status_msg"`
	Option map[string]interface{} `json:"omitempty"`
}

func SendResponse(c *gin.Context, msg string, options ...map[string]interface{}) {
	if len(options) == 0 {
		c.JSON(200, DefaultRes{StatusCode: 0, StatusMsg: msg})
		return
	}
	
	c.JSON(200, DefaultRes{
		StatusCode: 0,
		StatusMsg: msg,
		Option: options[0],
	})
}

func SendError(c *gin.Context, errMsg string) {
	c.JSON(400, DefaultRes{StatusCode: -1, StatusMsg: errMsg})
}

func SendRpcError(c *gin.Context, errMsg string) {
	c.JSON(500, DefaultRes{StatusCode: -1, StatusMsg: errMsg})
}