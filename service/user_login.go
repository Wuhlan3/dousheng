package service

import (
	"dousheng/repository"
	"dousheng/util"
)

type UserLoginFlow struct {
	name     string
	password string
}

func UserLogin(name string, password string) (int64, error) {
	return NewUserLoginFlow(name, password).Do()
}

func NewUserLoginFlow(name string, password string) *UserLoginFlow {
	return &UserLoginFlow{
		name:     name,
		password: password,
	}
}

func (f *UserLoginFlow) Do() (int64, error) {
	if err := f.checkParam(); err != nil {
		return 0, err
	}
	id, err := f.login()
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (f *UserLoginFlow) checkParam() error {
	return nil
}

func (f *UserLoginFlow) login() (int64, error) {
	user, err := repository.NewUserDaoInstance().QueryUserByName(f.name)
	if err != nil {
		util.Logger.Error("login err:" + err.Error())
		return 0, err
	}
	if user.Password != f.password {
		util.Logger.Error("login password err:" + err.Error())
	}

	return user.Id, nil
}
