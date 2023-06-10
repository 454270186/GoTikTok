package main

import (
	"github.com/454270186/GoTikTok/cmd/handler"
	"github.com/454270186/GoTikTok/pkg/middleware"
	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	r := gin.Default()

	userHandler := handler.NewUserHandler()
	publishHandler := handler.NewPubHandler()
	
	rUser := r.Group("/douyin/user")
	{	
		rUser.POST("/register/", userHandler.Register)
		rUser.POST("/login/", userHandler.Login)
		rUser.GET("/", middleware.VerifyToken(), userHandler.GetUser)
	}

	rPublish := r.Group("/douyin/publish")
	{
		rPublish.GET("/list/", middleware.VerifyToken(), publishHandler.List)
		rPublish.POST("/action/", middleware.VerifyToken(), publishHandler.Action)
	}

	rFeed := r.Group("/douyin/feed")
	{
		rFeed.GET("/", func(ctx *gin.Context) {
			ctx.String(200, `{
				"status_code": 0,
				"status_msg": "string",
				"next_time": 0,
				"video_list": [
					{
						"id": 0,
						"author": {
							"id": 0,
							"name": "string",
							"follow_count": 0,
							"follower_count": 0,
							"is_follow": true
						},
						"play_url": "string",
						"cover_url": "string",
						"favorite_count": 0,
						"comment_count": 0,
						"is_favorite": true,
						"title": "string"
					}
				]
			}`)
		})
	}

	return r
}