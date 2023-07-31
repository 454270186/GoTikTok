package main

import (
	"github.com/454270186/GoTikTok/cmd/handler"
	"github.com/454270186/GoTikTok/pkg/middleware"
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/collectors"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func NewRouter() *gin.Engine {
	r := gin.Default()

	// global middleware
	r.Use(middleware.Cors())

	// Prometheus
	// Create non-global registry.
	reg := prometheus.NewRegistry()

	// Add go runtime metrics and process collectors.
	reg.MustRegister(
		collectors.NewGoCollector(),
		collectors.NewProcessCollector(collectors.ProcessCollectorOpts{}),
	)

	r.GET("/metrics", gin.WrapH(promhttp.HandlerFor(reg, promhttp.HandlerOpts{Registry: reg})))

	userHandler := handler.NewUserHandler()
	publishHandler := handler.NewPubHandler()
	feefHandler := handler.NewFeedHandler()
	favoriteHandler := handler.NewFavHandler()
	commentHandler := handler.NewCommentHandler()
	
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

	rComment := r.Group("/douyin/comment")
	{
		rComment.GET("/list/", middleware.VerifyToken(), commentHandler.List)
		rComment.POST("/action/", middleware.VerifyToken(), commentHandler.Action)
	}

	return r
}