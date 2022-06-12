package service

import (
	"dousheng/proto/proto"
	"dousheng/repository"
)

func RelationFollowList(uid int64) ([]*proto.User, error) {
	var protoUserList []*proto.User
	followList, err := repository.NewFollowDaoInstance().QueryByUId(uid)
	if err != nil {
		return nil, err
	}
	for _, follow := range *followList {
		user, err := repository.NewUserDaoInstance().QueryUserById(follow.HisId)
		if err != nil {
			return nil, err
		}
		protoUserList = append(protoUserList, &proto.User{
			Id:            user.Id,
			Name:          user.Name,
			FollowCount:   user.FollowCount,
			FollowerCount: user.FollowerCount,
			IsFollow:      true,
		})
	}
	return protoUserList, nil
}
