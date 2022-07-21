package service

import (
	"dousheng/repository"
)

type UserInfo struct {
	Id            int64
	Name          string
	FollowCount   int64
	FollowerCount int64
	IsFollow      bool
}

func QueryUserInfo(myUId int64, id int64) (*UserInfo, error) {
	// 查询缓存
	user, err := GetUserInfoByUserIDFromRedis(id)

	if err != nil && err.Error() == "not found in cache" {
		user, err = repository.NewUserDaoInstance().QueryUserById(id)
		if err != nil {
			return nil, err
		}
		// 更新缓存
		if err = AddUserInfoByUserIDToRedis(user); err != nil {
			return nil, err
		}
	} else if err != nil {
		return nil, err
	}

	isFollow, err := repository.NewFollowDaoInstance().QueryIsFollowByUIdAndHisUId(myUId, id)
	if err != nil {
		return nil, err
	}

	return &UserInfo{
		Id:            id,
		Name:          user.Name,
		FollowerCount: user.FollowCount,
		FollowCount:   user.FollowerCount,
		IsFollow:      isFollow,
	}, nil
}
