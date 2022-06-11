package controller

import (
	"dousheng/proto/proto"
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
	//token := c.PostForm("token")
	//
	////验证token，如果token不存在则获取失败
	//if _, exist := usersLoginInfo[token]; !exist {
	//	c.JSON(http.StatusOK, Response{StatusCode: 1, StatusMsg: "User doesn't exist"})
	//	return
	//}
	//
	//data, err := c.FormFile("data")
	//if err != nil {
	//	c.JSON(http.StatusOK, Response{
	//		StatusCode: 1,
	//		StatusMsg:  err.Error(),
	//	})
	//	return
	//}
	//
	//filename := filepath.Base(data.Filename)
	//user := usersLoginInfo[token]
	//finalName := fmt.Sprintf("%d_%s", user.Id, filename)
	//
	////用方法从数据库中获取
	//saveFile := filepath.Join("./public/", finalName)
	//if err := c.SaveUploadedFile(data, saveFile); err != nil {
	//	c.JSON(http.StatusOK, Response{
	//		StatusCode: 1,
	//		StatusMsg:  err.Error(),
	//	})
	//	return
	//}

	c.JSON(http.StatusOK, Response{
		StatusCode: 0,
		//StatusMsg:  finalName + " uploaded successfully",
	})
}

// PublishList all users have same publish video list
//暂未开发完成
func PublishList(c *gin.Context) {
	c.JSON(http.StatusOK, proto.DouyinPublishActionResponse{
		StatusCode: 0,
	})
}
