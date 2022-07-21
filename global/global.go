package global

import (
	"context"
	"github.com/go-redis/redis/v8"
	"time"
)

var (
	REDIS   *redis.Client          // Redis 缓存接口
	CONTEXT = context.Background() // 上下文信息
)

// 过期时间
var (
	FAVORITE_EXPIRE       = 10 * time.Minute
	VIDEO_COMMENTS_EXPIRE = 10 * time.Minute
	COMMENT_EXPIRE        = 10 * time.Minute
	FOLLOW_EXPIRE         = 10 * time.Minute
	USER_INFO_EXPIRE      = 10 * time.Minute
	VIDEO_EXPIRE          = 10 * time.Minute
	PUBLISH_EXPIRE        = 10 * time.Minute
	EMPTY_EXPIRE          = 10 * time.Minute
	EXPIRE_TIME_JITTER    = 10 * time.Minute
)
