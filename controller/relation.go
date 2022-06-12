package controller

import (
	"dousheng/proto/proto"
	"dousheng/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

//type UserListResponse struct {
//	Response
//	UserList []User `json:"user_list"`
//}

// RelationAction no practical effect, just check if token is valid
func RelationAction(c *gin.Context) {
	uid, _ := c.Get("uid")
	MyUId := uid.(int64)
	toUserId := c.Query("to_user_id")
	HisUId, err := strconv.ParseInt(toUserId, 10, 64)
	if err != nil {
		c.JSON(http.StatusOK, Response{StatusCode: 1, StatusMsg: "parse int error"})
	}
	actionType := c.Query("action_type")
	err = service.RelationAction(MyUId, HisUId, actionType)
	if err != nil {
		c.JSON(http.StatusOK, Response{StatusCode: 1, StatusMsg: "RelationAction error"})
	}
	c.JSON(http.StatusOK, Response{StatusCode: 0, StatusMsg: "RelationAction successfully!"})

}

// FollowList all users have same follow list
func FollowList(c *gin.Context) {
	//uid, _ := c.Get("uid")
	//userId := uid.(int64)
	userIdStr := c.Query("user_id")
	userId, err := strconv.ParseInt(userIdStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusOK, UserResponse{
			Response: Response{StatusCode: 1, StatusMsg: "id convert error"},
		})
		return
	}
	if userId == 0 {
		uid, _ := c.Get("uid")
		userId = uid.(int64)
	}
	followList, err := service.RelationFollowList(userId)
	if err != nil {
		c.JSON(http.StatusOK, proto.DouyinRelationFollowListResponse{
			StatusCode: 1,
			StatusMsg:  "follow list error",
		})
	}
	c.JSON(http.StatusOK, &proto.DouyinRelationFollowListResponse{
		StatusCode: 0,
		StatusMsg:  "follow list successfully",
		UserList:   followList,
	})
}

// FollowerList all users have same follower list
func FollowerList(c *gin.Context) {
	userIdStr := c.Query("user_id")
	userId, err := strconv.ParseInt(userIdStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusOK, UserResponse{
			Response: Response{StatusCode: 1, StatusMsg: "id convert error"},
		})
		return
	}
	if userId == 0 {
		uid, _ := c.Get("uid")
		userId = uid.(int64)
	}
	followList, err := service.RelationFollowerList(userId)
	if err != nil {
		c.JSON(http.StatusOK, proto.DouyinRelationFollowerListResponse{
			StatusCode: 1,
			StatusMsg:  "follow list error",
		})
	}
	c.JSON(http.StatusOK, &proto.DouyinRelationFollowerListResponse{
		StatusCode: 0,
		StatusMsg:  "follow list successfully",
		UserList:   followList,
	})
}
