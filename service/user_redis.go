package service

import (
	"dousheng/global"
	"dousheng/repository"
	"errors"
	"fmt"
	"github.com/go-redis/redis/v8"
	"math/rand"
	"strconv"
	"time"
)

func AddUserInfoByUserIDToRedis(user *repository.User) error {
	// 定义 key
	userKey := fmt.Sprintf("user:%d", user.Id)

	// 使用 pipeline
	_, err := global.REDIS.TxPipelined(global.CONTEXT, func(pipe redis.Pipeliner) error {
		pipe.HSet(global.CONTEXT, userKey, "id", user.Id)
		pipe.HSet(global.CONTEXT, userKey, "name", user.Name)
		pipe.HSet(global.CONTEXT, userKey, "password", user.Password)
		pipe.HSet(global.CONTEXT, userKey, "follow_count", user.FollowerCount)
		pipe.HSet(global.CONTEXT, userKey, "follower_count", user.FollowerCount)
		pipe.HSet(global.CONTEXT, userKey, "create_time", user.CreateTime.UnixMilli())
		// 设置过期时间
		pipe.Expire(global.CONTEXT, userKey, global.USER_INFO_EXPIRE+time.Duration(rand.Float64()*global.EXPIRE_TIME_JITTER.Seconds())*time.Second)
		return nil
	})
	return err
}

func GetUserInfoByUserIDFromRedis(userID int64) (*repository.User, error) {
	// 定义 key
	userKey := fmt.Sprintf("user:%d", userID)

	var user repository.User

	if result := global.REDIS.Exists(global.CONTEXT, userKey).Val(); result <= 0 {
		return nil, errors.New("not found in cache")
	}
	// 使用 pipeline
	commands, err := global.REDIS.TxPipelined(global.CONTEXT, func(pipe redis.Pipeliner) error {
		pipe.HGetAll(global.CONTEXT, userKey)
		pipe.HGet(global.CONTEXT, userKey, "create_time").Val()
		// 设置过期时间
		pipe.Expire(global.CONTEXT, userKey, global.USER_INFO_EXPIRE+time.Duration(rand.Float64()*global.EXPIRE_TIME_JITTER.Seconds())*time.Second)
		return nil
	})
	if err != nil {
		return nil, err
	}
	if err = commands[0].(*redis.StringStringMapCmd).Scan(&user); err != nil {
		fmt.Println(err)
		return nil, err
	}

	timeUnixMilliStr := commands[1].(*redis.StringCmd).Val()
	timeUnixMilli, _ := strconv.ParseInt(timeUnixMilliStr, 10, 64)
	user.UpdateTime = time.UnixMilli(timeUnixMilli)
	return &user, nil
}
