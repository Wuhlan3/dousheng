package service

import (
	"dousheng/repository"
	"dousheng/util"
	"errors"
	"golang.org/x/crypto/bcrypt"
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

	///用户名与密码的规范

	return nil
}

func (f *UserRegisterFlow) register() (int64, error) {
	hashPassword, err := bcrypt.GenerateFromPassword([]byte(f.password), bcrypt.DefaultCost)
	if err != nil {
		util.Logger.Error("password hash error:" + err.Error())
		return 0, err
	}

	//检查用户是否已经存在
	if user, err := repository.NewUserDaoInstance().QueryUserByName(f.name); err != nil {
		return 0, err
	} else if err == nil && user.Id != 0 {
		return 0, errors.New("用户已存在")
	}
	//创建用户
	user := &repository.User{
		Id:            0,
		Name:          f.name,
		Password:      string(hashPassword),
		FollowCount:   0,
		FollowerCount: 0,
		CreateTime:    time.Now(),
		UpdateTime:    time.Now(),
		IsDeleted:     false,
	}
	if err := repository.NewUserDaoInstance().CreateUser(user); err != nil {
		return 0, err
	}
	return user.Id, nil
}
