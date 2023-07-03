package handler

import (
	"log"
	"strconv"

	"github.com/454270186/GoTikTok/cmd/httpres"
	"github.com/454270186/GoTikTok/cmd/rpccli"
	"github.com/454270186/GoTikTok/pkg/auth"
	"github.com/454270186/GoTikTok/rpc/comment/comments"
	"github.com/gin-gonic/gin"
)

type CommentHandler struct {
	commentRpcCli comments.Comments
}

func NewCommentHandler() *CommentHandler {
	return &CommentHandler{
		commentRpcCli: rpccli.NewCommentCli(),
	}
}

func (com CommentHandler) List(c *gin.Context) {
	videoID := c.Query("video_id")
	if videoID == "" {
		httpres.SendError(c, "video id cannot be empty")
		return
	}

	in := comments.CommentListReq{
		VideoId: videoID,
	}

	resp, err := com.commentRpcCli.CommentList(c.Copy(), &in)
	if err != nil {
		httpres.SendRpcError(c, err.Error())
		return
	}

	httpres.SendResponse(c, "successful",gin.H{
		"comment_list": resp.CommentList,
	})
}

func (com CommentHandler) Action(c *gin.Context) {
	userID, err := auth.GetUIDFromToken(c.Query("token"))
	if err != nil {
		httpres.SendError(c, err.Error())
		return
	}

	videoID := c.Query("video_id")
	actionType := c.Query("action_type")
	if len(videoID) == 0 || len(actionType) == 0 {
		httpres.SendError(c, "video id and action type cannot be empty")
		return
	}

	switch actionType {
	case "1":
		// publish comment
		text := c.Query("comment_text")

		in := comments.AddCommentReq{
			VideoId: videoID,
			CommentText: text,
			UserId: strconv.Itoa(int(userID)),
		}

		resp, err := com.commentRpcCli.AddComment(c.Copy(), &in)
		if err != nil {
			log.Println(err)
			httpres.SendRpcError(c, err.Error())
			return
		}

		httpres.SendResponse(c, "successful", gin.H{
			"comment": resp.Comment,
		})

	case "2":
		// delete comment
		commentID := c.Query("comment_id")

		in := comments.DelCommentReq{
			CommentId: commentID,
		}

		_, err := com.commentRpcCli.DelComment(c.Copy(), &in)
		if err != nil {
			httpres.SendRpcError(c, err.Error())
			return
		}

		httpres.SendResponse(c, "successful")

	}
}