package controller

import (
	"dousheng/proto/proto"
	"dousheng/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// FavoriteAction no practical effect, just check if token is valid
func FavoriteAction(c *gin.Context) {
	//token := c.Query("token")
	uid, _ := c.Get("uid")
	vid := c.Query("video_id")
	//fmt.Println(uid)
	//fmt.Println(vid)
	//userId, err := strconv.ParseInt(uid, 10, 64)
	userId := uid.(int64)
	videoId, err := strconv.ParseInt(vid, 10, 64)

	if err != nil {
		c.JSON(http.StatusOK, UserResponse{
			Response: Response{StatusCode: 1, StatusMsg: "id convert error"},
		})
		return
	}
	err = service.FavouriteAction(userId, videoId)
	if err != nil {
		c.JSON(http.StatusOK, Response{StatusCode: 1, StatusMsg: "favourite change error"})
	} else {
		c.JSON(http.StatusOK, Response{StatusCode: 0})
	}
}

// FavoriteList all users have same favorite video list
func FavoriteList(c *gin.Context) {
	uid, _ := c.Get("uid")
	userId := uid.(int64)

	videos, err := service.FavouriteList(userId)
	if err != nil {
		c.JSON(http.StatusOK, proto.DouyinPublishListResponse{
			StatusCode: 1,
			StatusMsg:  "Video loads Failed",
		})
	}
	c.JSON(http.StatusOK, proto.DouyinPublishListResponse{
		StatusCode: 0,
		StatusMsg:  "publishList successfully",
		VideoList:  videos,
	})
}
