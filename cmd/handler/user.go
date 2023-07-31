package handler

import (
	"net/http"

	"github.com/454270186/GoTikTok/cmd/httpres"
	"github.com/454270186/GoTikTok/cmd/rpccli"
	"github.com/454270186/GoTikTok/pkg/tracing"
	"github.com/454270186/GoTikTok/rpc/user/types/user"
	"github.com/454270186/GoTikTok/rpc/user/userservice"
	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	userTracer tracing.JaegerTracer
	userRpcCli userservice.UserService
}

func NewUserHandler() *UserHandler {
	return &UserHandler{
		userTracer: tracing.NewTracer("User-HTTP-Gateway", JaegerAgentAddr),
		userRpcCli: rpccli.NewUserCli(),
	}
}

func (u UserHandler) Register(c *gin.Context) {
	span, traceID, err := u.userTracer.Gin("Register handler", "", true)
	if err != nil {
		httpres.SendRpcError(c, "error while start span")
		return
	}
	u.userTracer.SpanSetTag(span, "http_uri", c.Request.RequestURI)
	defer u.userTracer.FinishSpan(span)

	username := c.Query("username")
	password := c.Query("password")

	if len(username) == 0 || len(password) == 0 {
		httpres.SendError(c, "username and password cannot be empty")
		return
	}

	rpcSpan, rpcSpanID, err := u.userTracer.RPC("call user rpc", traceID, false)
	if err != nil {
		httpres.SendRpcError(c, "error while start rpc span")
		return
	}
	in := user.RegisterReq{
		Username: username,
		Password: password,
		TraceID: rpcSpanID,
	}

	resp, err := u.userRpcCli.Register(c.Copy(), &in)
	if err != nil {
		httpres.SendRpcError(c, err.Error())
		u.userTracer.SpanSetTag(rpcSpan, "rpc_error", err.Error())
		u.userTracer.FinishSpan(rpcSpan)
		return
	} else if resp.StatusCode != 0 {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status_code": resp.StatusCode,
			"status_msg": "unexpect internal error",
		})
		return
	}
	u.userTracer.FinishSpan(rpcSpan)

	httpres.SendResponse(c, "register successfully", gin.H{
		"user_id": resp.UserId,
		"token": resp.Token,
	})
}

func (u UserHandler) Login(c *gin.Context) {
	span, traceID, err := u.userTracer.Gin("Login handler", "", true)
	if err != nil {
		httpres.SendRpcError(c, "error while start span")
		return
	}
	u.userTracer.SpanSetTag(span, "http_uri", c.Request.RequestURI)
	defer u.userTracer.FinishSpan(span)

	username := c.Query("username")
	password := c.Query("password")

	if len(username) == 0 || len(password) == 0 {
		httpres.SendError(c, "username and password cannot be empty")
		return
	}

	rpcSpan, rpcSpanID, err := u.userTracer.RPC("call user rpc", traceID, false)
	if err != nil {
		httpres.SendRpcError(c, "error while start rpc span")
		return
	}
	in := user.LoginReq{
		Username: username,
		Password: password,
		TraceID: rpcSpanID,
	}

	resp, err := u.userRpcCli.Login(c.Copy(), &in)
	if err != nil {
		httpres.SendRpcError(c, err.Error())
		u.userTracer.SpanSetTag(rpcSpan, "rpc_error", err.Error())
		u.userTracer.FinishSpan(rpcSpan)
		return
	} else if resp.StatusCode != 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"status_code": -1,
			"status_msg": "unexpect internal error",
		})
		return
	}
	u.userTracer.FinishSpan(rpcSpan)

	httpres.SendResponse(c, "login successfully", gin.H{
		"user_id": resp.UserId,
		"token": resp.Token,
	})
}

func (u UserHandler) GetUser(c *gin.Context) {
	span, traceID, err := u.userTracer.Gin("GetUser handler", "", true)
	if err != nil {
		httpres.SendRpcError(c, "error while start span")
		return
	}
	u.userTracer.SpanSetTag(span, "http_uri", c.Request.RequestURI)
	defer u.userTracer.FinishSpan(span)

	userId := c.Query("user_id")
	if len(userId) == 0 {
		httpres.SendError(c, "userID cannot be empty")
		return
	}

	rpcSpan, rpcSpanID, err := u.userTracer.RPC("call user rpc", traceID, false)
	if err != nil {
		httpres.SendRpcError(c, "error while start rpc span")
		return
	}
	in := user.GetUserReq{
		UserId: userId,
		TraceID: rpcSpanID,
	}

	resp, err := u.userRpcCli.GetUserById(c.Copy(), &in)
	if err != nil {
		httpres.SendRpcError(c, err.Error())
		u.userTracer.SpanSetTag(rpcSpan, "rpc_error", err.Error())
		u.userTracer.FinishSpan(rpcSpan)
		return
	} else if resp.StatusCode != 0 {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status_code": resp.StatusCode,
			"status_msg": "unexpect internal error",
		})
		return
	}
	u.userTracer.FinishSpan(rpcSpan)

	httpres.SendResponse(c, "success", gin.H{
		"user": gin.H{
			"id": resp.User.Id,
			"name": resp.User.Name,
			"follow_count": resp.User.FollowCount,
			"follower_count": resp.User.FollowerCount,
			"isfollow": true,
			"avatar": "https://cdn.pixabay.com/photo/2023/06/26/04/35/ai-generated-8088678_1280.jpg",
			"background_image": "https://cdn.pixabay.com/photo/2012/08/27/14/19/mountains-55067_1280.png",
			"signature": "true 2 myself",
		},
	})
}