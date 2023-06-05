package main

import "github.com/gin-gonic/gin"

func NewRouter() *gin.Engine {
	r := gin.Default()

	rUser := r.Group("/douyin/user")
	{	
		rUser.POST("/register", func(ctx *gin.Context) {})
		rUser.POST("/login", func(ctx *gin.Context) {})
		rUser.GET("/", func(ctx *gin.Context) {})
	}

	rPublish := r.Group("/douyin/publish")
	{
		rPublish.GET("/list", func(ctx *gin.Context) {})
		rPublish.POST("/action", func(ctx *gin.Context) {})
	}

	rFeed := r.Group("/douyin/feed")
	{
		rFeed.GET("/", func(ctx *gin.Context) {})
	}

	return r
}