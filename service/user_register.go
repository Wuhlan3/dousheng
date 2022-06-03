package service

import (
	"dousheng/repository"
	"dousheng/util"
	"time"
)

type UserRegisterFlow struct {
	name     string
	password string
}

func UserRegister(name string, password string) (int64, error) {
	return NewUserRegisterFlow(name, password).Do()
}

func NewUserRegisterFlow(name string, password string) *UserRegisterFlow {
	return &UserRegisterFlow{
		name:     name,
		password: password,
	}
}

func (f *UserRegisterFlow) Do() (int64, error) {
	if err := f.checkParam(); err != nil {
		return 0, err
	}
	id, err := f.register()
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (f *UserRegisterFlow) checkParam() error {
	return nil
}

func (f *UserRegisterFlow) register() (int64, error) {
	user := &repository.User{
		Id:            0,
		Name:          f.name,
		Password:      f.password,
		FollowCount:   0,
		FollowerCount: 0,
		CreateTime:    time.Now(),
		UpdateTime:    time.Now(),
		IsDeleted:     false,
	}
	if err := repository.NewUserDaoInstance().CreateUser(user); err != nil {
		util.Logger.Error("register err:" + err.Error())
		return 0, err
	}
	return user.Id, nil
}
