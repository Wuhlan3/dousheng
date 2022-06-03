package controller

import (
	"dousheng/service"
	"dousheng/util"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// usersLoginInfo use map to store user info, and key is username+password for demo
// user data will be cleared every time the server starts
// test data: username=zhanglei, password=douyin

//样例用户
//使用token来直接映射user结构体
var usersLoginInfo = map[string]User{
	"zhangleidouyin": {
		Id:            4,
		Name:          "wuhlan3",
		FollowCount:   0,
		FollowerCount: 0,
		IsFollow:      false,
	},
}

//var userIdSequence = int64(1)

type UserLoginResponse struct {
	Response
	UserId int64  `json:"user_id,omitempty"`
	Token  string `json:"token"`
}

type UserResponse struct {
	Response
	User service.UserInfoFlow `json:"user"`
}

func Register(c *gin.Context) {
	//数据解析
	username := c.Query("username")
	password := c.Query("password")

	id, err := service.UserRegister(username, password)
	if err != nil {
		util.Logger.Error("register err:" + err.Error())
	}
	c.JSON(http.StatusOK, UserLoginResponse{
		Response: Response{StatusCode: 0},
		UserId:   id,
		Token:    username + password,
	})
}

func Login(c *gin.Context) {
	//数据解析
	username := c.Query("username")
	password := c.Query("password")

	token := username + password

	if id, err := service.UserLogin(username, password); err != nil {
		c.JSON(http.StatusOK, UserLoginResponse{
			Response: Response{StatusCode: 1, StatusMsg: "User doesn't exist"},
		})
	} else {
		c.JSON(http.StatusOK, UserLoginResponse{
			Response: Response{StatusCode: 0},
			UserId:   id,
			Token:    token,
		})
	}
}

func UserInfo(c *gin.Context) {
	//数据解析
	id := c.Query("user_id")
	userId, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		c.JSON(http.StatusOK, UserResponse{
			Response: Response{StatusCode: 1, StatusMsg: "id convert error"},
		})
	}

	if user, err := service.UserInfo(userId); err != nil {
		c.JSON(http.StatusOK, UserResponse{
			Response: Response{StatusCode: 1, StatusMsg: "User doesn't exist"},
		})
	} else {
		//fmt.Println(*user)
		c.JSON(http.StatusOK, UserResponse{
			Response: Response{StatusCode: 0},
			User:     *user,
		})
	}

}
