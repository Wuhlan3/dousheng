package repository

import (
	"context"
	"dousheng/global"
	"github.com/spf13/viper"
	"time"

	"github.com/go-redis/redis/v8" // 注意导入的是新版本
)

// 初始化连接
func RedisInit() (err error) {
	global.REDIS = redis.NewClient(&redis.Options{
		Addr:     viper.GetString("host") + ":" + viper.GetString("port"),
		Password: viper.GetString("password"), // no password set
		DB:       viper.GetInt("db"),          // use default DB
		PoolSize: viper.GetInt("poolsize"),    // 连接池大小
	})

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err = global.REDIS.Ping(ctx).Result()
	return err
}
