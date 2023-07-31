package handler

import (
	"bytes"
	"context"
	"io"
	"net/http"

	"github.com/454270186/GoTikTok/cmd/httpres"
	"github.com/454270186/GoTikTok/cmd/rpccli"
	"github.com/454270186/GoTikTok/pkg/auth"
	"github.com/454270186/GoTikTok/pkg/tracing"
	"github.com/454270186/GoTikTok/rpc/publish/publishclient"
	"github.com/gin-gonic/gin"
)

type PublishHandler struct {
	pubTracer tracing.JaegerTracer
	pubRpcCli publishclient.Publish
}

func NewPubHandler() *PublishHandler {
	return &PublishHandler{
		pubTracer: tracing.NewTracer("Publish-HTTP-Gateway", JaegerAgentAddr),
		pubRpcCli: rpccli.NewPubCli(),
	}
}

func (p PublishHandler) List(c *gin.Context) {
	span, traceID, err := p.pubTracer.Gin("List handler", "", true)
	if err != nil {
		httpres.SendRpcError(c, "error while start span")
		return
	}
	p.pubTracer.SpanSetTag(span, "http_uri", c.Request.RequestURI)
	defer p.pubTracer.FinishSpan(span)

	userID := c.Query("user_id")

	rpcSpan, rpcTraceID, err := p.pubTracer.RPC("call publish rpc", traceID, false)
	if err != nil {
		httpres.SendRpcError(c, "error while start span")
		return
	}
	in := publishclient.PublishListReq {
		UserId: userID,
		TraceID: rpcTraceID,
	}

	resp, err := p.pubRpcCli.PublishList(context.Background(), &in)
	if err != nil {
		httpres.SendRpcError(c, err.Error())
		p.pubTracer.SpanSetTag(rpcSpan, "rpc_error", err.Error())
		p.pubTracer.FinishSpan(rpcSpan)
		return
	} else if resp.StatusCode != 0 {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status_code": -1,
			"status_msg": "resp.StatusCode != 0",
		})
		return
	}
	p.pubTracer.FinishSpan(rpcSpan)

	httpres.SendResponse(c, "success", gin.H{
		"video_list": resp.VideoList,
	})
}

func (p PublishHandler) Action(c *gin.Context) {
	span, traceID, err := p.pubTracer.Gin("Action handler", "", true)
	if err != nil {
		httpres.SendRpcError(c, "error while start span")
		return
	}
	p.pubTracer.SpanSetTag(span, "http_uri", c.Request.RequestURI)
	defer p.pubTracer.FinishSpan(span)

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

	rpcSpan, rpcTraceID, err := p.pubTracer.RPC("call publish rpc", traceID, false)
	if err != nil {
		httpres.SendRpcError(c, "error while start span")
		return
	}
	in := publishclient.PublishActionReq{
		Title: title,
		Data: buf.Bytes(),
		Uid: int64(uid),
		TraceID: rpcTraceID,
	}
	
	_, err = p.pubRpcCli.PublishAction(c.Copy(), &in)
	if err != nil {
		httpres.SendRpcError(c, err.Error())
		p.pubTracer.SpanSetTag(rpcSpan, "rpc_error", err.Error())
		p.pubTracer.FinishSpan(rpcSpan)
		return
	}
	p.pubTracer.FinishSpan(rpcSpan)

	httpres.SendResponse(c, "publish successfully")
}