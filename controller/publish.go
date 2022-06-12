package controller

import (
	"dousheng/proto/proto"
	"dousheng/service"
	"fmt"
	"github.com/spf13/viper"
	"path/filepath"
	//"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	//"path/filepath"
)

//type VideoListResponse struct {
//	proto.DouyinPublishActionResponse
//	VideoList []proto.Video `json:"video_list"`
//}

// Publish check token then save upload file to public directory
func Publish(c *gin.Context) {
	uid, _ := c.Get("uid")
	userId := uid.(int64)

	data, err := c.FormFile("data")
	if err != nil {
		c.JSON(http.StatusOK, Response{
			StatusCode: 1,
			StatusMsg:  err.Error(),
		})
		return
	}

	filename := filepath.Base(data.Filename) //返回路径最后的文件名
	finalName := fmt.Sprintf("%d_%s", userId, filename)

	if err := service.UploadVideo(userId, finalName); err != nil {
		c.JSON(http.StatusOK, Response{
			StatusCode: 1,
			StatusMsg:  err.Error(),
		})
		return
	}

	saveFile := filepath.Join(viper.GetString("video.savePath"), finalName) //文件保存路径
	if err := c.SaveUploadedFile(data, saveFile); err != nil {
		c.JSON(http.StatusOK, Response{
			StatusCode: 1,
			StatusMsg:  err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, Response{
		StatusCode: 0,
		StatusMsg:  finalName + " uploaded successfully",
	})
}

// PublishList all users have same publish video list
func PublishList(c *gin.Context) {
	uid, _ := c.Get("uid")
	userId := uid.(int64)

	videos, err := service.PublishList(userId)
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
