package handler

import (
	"strconv"

	"github.com/454270186/GoTikTok/cmd/httpres"
	"github.com/454270186/GoTikTok/cmd/rpccli"
	"github.com/454270186/GoTikTok/rpc/feed/feedservice"
	"github.com/454270186/GoTikTok/rpc/feed/types/feed"
	"github.com/gin-gonic/gin"
)

type FeedHandler struct {
	feedCli feedservice.FeedService
}

func NewFeedHandler() *FeedHandler {
	return &FeedHandler{
		feedCli: rpccli.NewFeedCli(),
	}
}

func (f FeedHandler) GetUserFeed(c *gin.Context) {
	var latestTime int64
	if latest := c.Query("latest_time"); latest == "" {
		latestTime = 0
	} else {
		var err error
		latestTime, err = strconv.ParseInt(latest, 10, 64)
		if err != nil {
			httpres.SendError(c, err.Error())
			return
		}
	}

	in := feed.FeedReq{
		LastestTime: latestTime,
	}

	resp, err := f.feedCli.GetUserFeed(c.Copy(), &in)
	if err != nil {
		httpres.SendRpcError(c, err.Error())
		return
	}

	httpres.SendResponse(c, "success", gin.H{
		"next_time": resp.NextTime,
		"video_list": resp.VideoList,
	})
}