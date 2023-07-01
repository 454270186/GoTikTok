package httpres

import (
	"github.com/gin-gonic/gin"
)

func SendResponse(c *gin.Context, msg string, options ...map[string]interface{}) {
	response := make(map[string]any)
	response["status_code"] = 0
	response["status_msg"] = msg

	if len(options) > 0 {
		for k, v := range options[0] {
			response[k] = v
		}
	}

	c.JSON(200, response)
}

func SendError(c *gin.Context, errMsg string) {
	response := make(map[string]any)
	response["status_code"] = -1
	response["status_msg"] = errMsg

	c.JSON(400, response)
}

func SendRpcError(c *gin.Context, errMsg string) {
	response := make(map[string]any)
	response["status_code"] = -2
	response["status_msg"] = errMsg

	c.JSON(500, response)
}
