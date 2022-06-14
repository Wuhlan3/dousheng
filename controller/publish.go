package controller

import (
	"context"
	"dousheng/proto/proto"
	"dousheng/service"
	"fmt"
	"github.com/spf13/viper"
	"github.com/tencentyun/cos-go-sdk-v5"
	"net/url"
	"path/filepath"
	"time"

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
	finalName := fmt.Sprintf("%d_%d_%s", time.Now().UnixNano(), userId, filename)

	if err := service.UploadVideo(userId, finalName); err != nil {
		c.JSON(http.StatusOK, Response{
			StatusCode: 1,
			StatusMsg:  err.Error(),
		})
		return
	}

	fd, err := data.Open()
	if err != nil {
		fmt.Println("error")
		return
	}

	u, _ := url.Parse(viper.GetString("cos.url"))
	b := &cos.BaseURL{BucketURL: u}
	cosClient := cos.NewClient(b, &http.Client{
		Transport: &cos.AuthorizationTransport{
			SecretID:  viper.GetString("cos.SecretID"),  // 替换为用户的 SecretId，请登录访问管理控制台进行查看和管理，https://console.cloud.tencent.com/cam/capi
			SecretKey: viper.GetString("cos.SecretKey"), // 替换为用户的 SecretKey，请登录访问管理控制台进行查看和管理，https://console.cloud.tencent.com/cam/capi
		},
	})
	name := viper.GetString("cos.videoDir") + finalName

	if err != nil {
		panic(err)
	}
	defer fd.Close()
	_, err = cosClient.Object.Put(context.Background(), name, fd, nil)
	if err != nil {
		panic(err)
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
