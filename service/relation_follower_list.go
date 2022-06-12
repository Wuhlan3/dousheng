package service

import (
	"dousheng/proto/proto"
	"dousheng/repository"
)

func RelationFollowerList(uid int64) ([]*proto.User, error) {
	var protoUserList []*proto.User
	followerList, err := repository.NewFollowDaoInstance().QueryByHisUId(uid)
	if err != nil {
		return nil, err
	}
	for _, follow := range *followerList {
		user, err := repository.NewUserDaoInstance().QueryUserById(follow.MyId)
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
