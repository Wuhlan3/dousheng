package service

import (
	"dousheng/repository"
)

type UserInfoFlow struct {
	id            int64
	name          string
	followCount   int64
	followerCount int64
	isFollow      bool
}

func UserInfo(id int64) (*UserInfoFlow, error) {
	return NewUserInfoFlow(id).Do()
}

func NewUserInfoFlow(id int64) *UserInfoFlow {
	return &UserInfoFlow{
		id: id,
	}
}

func (f *UserInfoFlow) Do() (*UserInfoFlow, error) {
	if err := f.checkParam(); err != nil {
		return nil, err
	}
	if err := f.info(); err != nil {
		return nil, err
	}
	return f, nil
}

func (f *UserInfoFlow) checkParam() error {
	return nil
}

func (f *UserInfoFlow) info() error {
	user, err := repository.NewUserDaoInstance().QueryUserById(f.id)
	f.name = user.Name
	f.followCount = user.FollowCount
	f.followerCount = user.FollowerCount
	f.isFollow = false
	//user * repository.User
	if err != nil {
		return err
	}
	//f.postId = user.Id
	return nil
}
