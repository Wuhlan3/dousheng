package controller

import (
	"dousheng/proto/proto"
	"dousheng/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"time"
)

// Feed same demo video list for every request
func Feed(c *gin.Context) {
	// 如果没有传latest_time，则默认为当前时间
	var CurrentTimeInt = time.Now().UnixMilli()
	var CurrentTime = strconv.FormatInt(CurrentTimeInt, 10)
	var LatestTimeStr = c.DefaultQuery("latest_time", CurrentTime)
	LatestTime, err := strconv.ParseInt(LatestTimeStr, 10, 64)
	if err != nil {
		// 无法解析latest_time
		c.JSON(http.StatusBadRequest, Response{StatusCode: 1, StatusMsg: "parameter latest_time is wrong"})
		return
	}

	myUidInt, _ := c.Get("uid")
	myUid := myUidInt.(int64)
	videos, err := service.Feed(myUid, LatestTime)
	if err != nil {
		c.JSON(http.StatusOK, proto.DouyinFeedResponse{
			StatusCode: 1,
			StatusMsg:  "Video loads Failed",
		})
		return
	}
	//视频加载没有出错
	c.JSON(http.StatusOK, proto.DouyinFeedResponse{
		StatusCode: 0,
		StatusMsg:  "Video loads successfully",
		VideoList:  videos,
		NextTime:   time.Now().Unix(),
	})
}
