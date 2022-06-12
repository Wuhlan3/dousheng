package controller

import (
	"dousheng/proto/proto"
	"dousheng/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

// Feed same demo video list for every request
func Feed(c *gin.Context) {
	//var Videos *proto.Video
	videos, err := service.Feed()
	if err != nil {
		c.JSON(http.StatusOK, proto.DouyinFeedResponse{
			StatusCode: 1,
			StatusMsg:  "Video loads Failed",
		})
	}
	//视频加载没有出错
	c.JSON(http.StatusOK, proto.DouyinFeedResponse{
		StatusCode: 0,
		StatusMsg:  "Video loads successfully",
		VideoList:  videos,
		NextTime:   time.Now().Unix(),
	})
}
