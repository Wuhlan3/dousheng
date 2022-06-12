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

type UserInfoFlow struct {
	myUId    int64
	id       int64
	user     *repository.User
	userInfo *UserInfo
}

func QueryUserInfo(myUid int64, id int64) (*UserInfo, error) {
	return NewUserInfoFlow(myUid, id).Do()
}

func NewUserInfoFlow(myUid int64, id int64) *UserInfoFlow {
	return &UserInfoFlow{
		myUId: myUid,
		id:    id,
	}
}

func (f *UserInfoFlow) Do() (*UserInfo, error) {
	if err := f.checkParam(); err != nil {
		return nil, err
	}
	if err := f.info(); err != nil {
		return nil, err
	}
	return f.userInfo, nil
}

func (f *UserInfoFlow) checkParam() error {
	return nil
}

func (f *UserInfoFlow) info() error {
	user, err := repository.NewUserDaoInstance().QueryUserById(f.id)
	if err != nil {
		return err
	}
	var IsFollow bool
	follow, err := repository.NewFollowDaoInstance().QueryByUIdAndHisUId(f.myUId, user.Id)
	if err != nil {
		IsFollow = false
	} else {
		IsFollow = follow.IsFollow
	}
	f.user = user
	f.userInfo = &UserInfo{
		f.user.Id,
		f.user.Name,
		f.user.FollowCount,
		f.user.FollowerCount,
		IsFollow,
	}
	return nil
}
