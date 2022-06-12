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
	apiRouter.GET("/feed/", middleware.AuthMiddleware(), controller.Feed)               //已实现
	apiRouter.GET("/user/", middleware.AuthMiddleware(), controller.UserInfo)           //已实现
	apiRouter.POST("/user/register/", controller.Register)                              //已实现
	apiRouter.POST("/user/login/", controller.Login)                                    //已实现
	apiRouter.POST("/publish/action/", middleware.AuthMiddleware(), controller.Publish) //已实现
	apiRouter.GET("/publish/list/", middleware.AuthMiddleware(), controller.PublishList)

	// extra apis - I
	apiRouter.POST("/favorite/action/", middleware.AuthMiddleware(), controller.FavoriteAction) //已实现
	apiRouter.GET("/favorite/list/", controller.FavoriteList)
	apiRouter.POST("/comment/action/", controller.CommentAction)
	apiRouter.GET("/comment/list/", controller.CommentList)

	// extra apis - II
	apiRouter.POST("/relation/action/", controller.RelationAction)
	apiRouter.GET("/relation/follow/list/", controller.FollowList)
	apiRouter.GET("/relation/follower/list/", controller.FollowerList)
}
