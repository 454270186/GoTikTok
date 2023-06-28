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
	feefHandler := handler.NewFeedHandler()
	favoriteHandler := handler.NewFavHandler()
	
	rUser := r.Group("/douyin/user")
	{	
		rUser.POST("/register/", userHandler.Register)
		rUser.POST("/login/", userHandler.Login)
		rUser.GET("/", middleware.VerifyToken(), userHandler.GetUser)
	}

	rPublish := r.Group("/douyin/publish")
	{
		rPublish.GET("/list/", middleware.VerifyToken(), publishHandler.List)
		rPublish.POST("/action/", publishHandler.Action)
	}

	rFeed := r.Group("/douyin/feed")
	{
		rFeed.GET("/", middleware.VerifyToken(), feefHandler.GetUserFeed) 
	}

	rFavorite := r.Group("/douyin/favorite")
	{
		rFavorite.GET("/list/", middleware.VerifyToken(), favoriteHandler.List)
		rFavorite.POST("/action/", middleware.VerifyToken(), favoriteHandler.Action)
	}

	return r
}