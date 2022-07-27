package controller

import (
	"bytes"
	"context"
	"dousheng/proto/proto"
	"dousheng/service"
	"fmt"
	"github.com/spf13/viper"
	"github.com/tencentyun/cos-go-sdk-v5"
	ffmpeg "github.com/u2takey/ffmpeg-go"
	"image"
	"image/jpeg"
	"mime/multipart"
	"net/url"
	"os"
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
	b := []byte(finalName)
	imgName := string(b[:len(b)-3]) + "jpg"
	// *********************************************************上传COS存储桶**********************************************
	err = uploadVideo(data, finalName)
	if err != nil {
		c.JSON(http.StatusOK, Response{
			StatusCode: 1,
			StatusMsg:  err.Error(),
		})
		return
	}
	coverData, err := readFrameAsJpeg(viper.GetString("cos.uriVideoPath") + finalName)
	if err != nil {
		c.JSON(http.StatusOK, Response{
			StatusCode: 1,
			StatusMsg:  err.Error(),
		})
		return
	}

	err = uploadImg(coverData, imgName)
	if err != nil {
		c.JSON(http.StatusOK, Response{
			StatusCode: 1,
			StatusMsg:  err.Error(),
		})
		return
	}
	// *********************************************************上传COS存储桶**********************************************

	if err := service.UploadVideo(userId, finalName, imgName); err != nil {
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

func uploadVideo(data *multipart.FileHeader, finalName string) error {
	fd, err := data.Open()
	if err != nil {
		fmt.Println("error")
		return err
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
		return err
	}
	defer fd.Close()
	_, err = cosClient.Object.Put(context.Background(), name, fd, nil)
	if err != nil {
		return err
	}
	return nil
}

func uploadImg(data []byte, finalName string) error {
	f := bytes.NewReader(data)

	u, _ := url.Parse(viper.GetString("cos.url"))
	b := &cos.BaseURL{BucketURL: u}
	cosClient := cos.NewClient(b, &http.Client{
		Transport: &cos.AuthorizationTransport{
			SecretID:  viper.GetString("cos.SecretID"),  // 替换为用户的 SecretId，请登录访问管理控制台进行查看和管理，https://console.cloud.tencent.com/cam/capi
			SecretKey: viper.GetString("cos.SecretKey"), // 替换为用户的 SecretKey，请登录访问管理控制台进行查看和管理，https://console.cloud.tencent.com/cam/capi
		},
	})
	name := viper.GetString("cos.imgDir") + finalName

	_, err := cosClient.Object.Put(context.Background(), name, f, nil)
	if err != nil {
		return err
	}
	return nil
}

func readFrameAsJpeg(filePath string) ([]byte, error) {
	reader := bytes.NewBuffer(nil)
	err := ffmpeg.Input(filePath).
		Filter("select", ffmpeg.Args{fmt.Sprintf("gte(n,%d)", 1)}).
		Output("pipe:", ffmpeg.KwArgs{"vframes": 1, "format": "image2", "vcodec": "mjpeg"}).
		WithOutput(reader, os.Stdout).
		Run()
	if err != nil {
		return nil, err
	}
	img, _, err := image.Decode(reader)
	if err != nil {
		return nil, err
	}

	buf := new(bytes.Buffer)
	jpeg.Encode(buf, img, nil)

	return buf.Bytes(), nil
}
