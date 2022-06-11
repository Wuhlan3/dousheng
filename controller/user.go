package controller

import (
	"dousheng/middleware"
	"dousheng/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

//样例用户
//使用token来直接映射user结构体
//var usersLoginInfo = map[string]User{
//	"zhangleidouyin": {
//		Id:            4,
//		Name:          "wuhlan3",
//		FollowCount:   4,
//		FollowerCount: 5,
//		IsFollow:      false,
//	},
//}

type UserLoginResponse struct {
	Response
	UserId int64  `json:"user_id,omitempty"`
	Token  string `json:"token"`
}

type UserResponse struct {
	Response
	User *service.UserInfo `json:"user"`
}

func Register(c *gin.Context) {
	//数据解析
	username := c.Query("username")
	password := c.Query("password")

	id, err := service.UserRegister(username, password)
	if err != nil {
		c.JSON(http.StatusOK, UserLoginResponse{
			Response: Response{StatusCode: 1, StatusMsg: err.Error()},
		})
		return
	}
	//生成JWT token
	token, err := middleware.ReleaseToken(id)
	if err != nil {
		c.JSON(http.StatusOK, UserLoginResponse{
			Response: Response{StatusCode: 1, StatusMsg: err.Error()},
		})
		return
	}
	//发送正确的响应
	c.JSON(http.StatusOK, UserLoginResponse{
		Response: Response{StatusCode: 0},
		UserId:   id,
		Token:    token,
	})
	return
}

func Login(c *gin.Context) {
	//数据解析
	username := c.Query("username")
	password := c.Query("password")

	//token := username + password
	id, err := service.UserLogin(username, password)
	if err != nil {
		c.JSON(http.StatusOK, UserLoginResponse{
			Response: Response{StatusCode: 1, StatusMsg: err.Error()},
		})
		return
	}
	//生成JWT token
	token, err := middleware.ReleaseToken(id)
	if err != nil {
		c.JSON(http.StatusOK, UserLoginResponse{
			Response: Response{StatusCode: 1, StatusMsg: err.Error()},
		})
		return
	}
	//发送正确响应
	c.JSON(http.StatusOK, UserLoginResponse{
		Response: Response{StatusCode: 0},
		UserId:   id,
		Token:    token,
	})
	return
}

func UserInfo(c *gin.Context) {
	//数据解析
	id := c.Query("user_id")
	userId, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		c.JSON(http.StatusOK, UserResponse{
			Response: Response{StatusCode: 1, StatusMsg: "id convert error"},
		})
		return
	}
	user, err := service.QueryUserInfo(userId)
	//fmt.Println(user)
	if err != nil {
		c.JSON(http.StatusOK, UserResponse{
			Response: Response{StatusCode: 1, StatusMsg: "User doesn't exist"},
		})
		return
	}
	//fmt.Println(*user)
	c.JSON(http.StatusOK, UserResponse{
		Response: Response{StatusCode: 0},
		User:     user,
	})

}
