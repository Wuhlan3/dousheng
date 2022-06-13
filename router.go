package main

import (
	"dousheng/controller"
	"dousheng/middleware"
	"github.com/gin-gonic/gin"
	"path"
)

func initRouter(r *gin.Engine) {
	// public directory is used to serve static resources
	r.Static("/static", "./public")
	r.GET("/home/go/src/dousheng/public/:name", func(c *gin.Context) {
		name := c.Param("name")
		filename := path.Join("./public/", name)
		c.File(filename)
		return
	})

	apiRouter := r.Group("/douyin")

	// basic apis
	apiRouter.GET("/feed/", middleware.AuthMiddleware(), controller.Feed)
	apiRouter.GET("/user/", middleware.AuthMiddleware(), controller.UserInfo)
	apiRouter.POST("/user/register/", controller.Register)
	apiRouter.POST("/user/login/", controller.Login)
	apiRouter.POST("/publish/action/", middleware.AuthMiddleware(), controller.Publish)
	apiRouter.GET("/publish/list/", middleware.AuthMiddleware(), controller.PublishList)

	// extra apis - I
	apiRouter.POST("/favorite/action/", middleware.AuthMiddleware(), controller.FavoriteAction)
	apiRouter.GET("/favorite/list/", middleware.AuthMiddleware(), controller.FavoriteList)
	apiRouter.POST("/comment/action/", middleware.AuthMiddleware(), controller.CommentAction)
	apiRouter.GET("/comment/list/", middleware.AuthMiddleware(), controller.CommentList)

	// extra apis - II
	apiRouter.POST("/relation/action/", middleware.AuthMiddleware(), controller.RelationAction)
	apiRouter.GET("/relation/follow/list/", middleware.AuthMiddleware(), controller.FollowList)
	apiRouter.GET("/relation/follower/list/", middleware.AuthMiddleware(), controller.FollowerList)
}
