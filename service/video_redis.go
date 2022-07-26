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

// GetVIdListFromRedis 从redis中获取视频序列
func GetVIdListFromRedis(LatestTime int64, MaxNumVideo int64) (VIdList []string, err error) {
	n, err := global.REDIS.Exists(global.CONTEXT, "feed").Result()
	if err != nil {
		return nil, err
	}
	//fmt.Println(n)
	// "feed"不存在
	if n <= 0 {
		videosList, err := repository.NewVideoDaoInstance().QueryVideoList(100)
		fmt.Println(videosList)
		if err != nil {
			return nil, err
		}
		if len(*videosList) == 0 {
			return nil, errors.New("not found video")
		}

		var listZ = make([]*redis.Z, 0, len(*videosList))
		for _, video := range *videosList {
			listZ = append(listZ, &redis.Z{Score: float64(video.CreateTime.UnixMilli()) / 1000, Member: video.Id})
		}

		if err := global.REDIS.ZAdd(global.CONTEXT, "feed", listZ...).Err(); err != nil {
			return nil, err
		}
	}
	// Redis中查询feed
	feedKey := "feed"

	// 初始化查询条件， Offset和Count用于分页
	cond := redis.ZRangeBy{
		Min:    "0",                                                    // 最小分数
		Max:    strconv.FormatFloat(float64(LatestTime-2), 'f', 3, 64), // 最大分数
		Offset: 0,                                                      // 类似sql的limit, 表示开始偏移量
		Count:  MaxNumVideo,                                            // 一次返回多少数据
	}

	//fmt.Println(cond)

	// 获取推送视频ID按逆序返回
	videoIDStrList, err := global.REDIS.ZRevRangeByScore(global.CONTEXT, feedKey, &cond).Result()
	numVideos := len(videoIDStrList)
	if err != nil || numVideos == 0 {
		return nil, err
	}

	return videoIDStrList, nil
}

// AddVideoListToRedis 将视频批量写入缓存
func AddVideoListToRedis(videoList []repository.Video) error {
	pipe := global.REDIS.TxPipeline()
	for _, video := range videoList {

		keyVideo := fmt.Sprintf("video:%d", video.Id)

		pipe.HSet(global.CONTEXT, keyVideo, "id", video.Id, "title", video.Title, "play_url", video.PlayUrl, "cover_url", video.CoverUrl,
			"favorite_count", video.FavouriteCount, "comment_count", video.CommentCount, "uid", video.UId, "create_time", video.CreateTime.UnixMilli())
		pipe.Expire(global.CONTEXT, keyVideo, global.VIDEO_EXPIRE+time.Duration(rand.Float64()*global.EXPIRE_TIME_JITTER.Seconds())*time.Second)
	}
	_, err := pipe.Exec(global.CONTEXT)
	return err
}

// GetVideoListFromRedis 根据vid批量从缓存中获取视频信息
func GetVideoListFromRedis(vidStrList []string) (videoList []repository.Video, err error) {
	numVideos := len(vidStrList)
	inCache := make([]bool, 0, numVideos)
	notInCacheIDList := make([]int64, 0, numVideos)
	for _, vid := range vidStrList {
		videoID, err := strconv.ParseInt(vid, 10, 64)
		if err != nil {
			return nil, err
		}
		keyVideo := fmt.Sprintf("video:%d", videoID)
		n, err := global.REDIS.Exists(global.CONTEXT, keyVideo).Result()
		if err != nil {
			return nil, err
		}
		if n <= 0 {
			// 当前视频不在缓存中，直接放一个空的video即可
			videoList = append(videoList, repository.Video{})
			inCache = append(inCache, false)
			notInCacheIDList = append(notInCacheIDList, videoID)
			continue
		}
		// video存在
		var video repository.Video
		if err = global.REDIS.Expire(global.CONTEXT, keyVideo, global.VIDEO_EXPIRE).Err(); err != nil {
			return nil, err
		}
		if err = global.REDIS.HGetAll(global.CONTEXT, keyVideo).Scan(&video); err != nil {
			return nil, errors.New("GetVideoListByIDs from Redis failed")
		}
		video.Id = videoID
		timeUnixMilliStr, err := global.REDIS.HGet(global.CONTEXT, keyVideo, "create_time").Result()
		if err != nil {
			continue
		}
		timeUnixMilli, err := strconv.ParseInt(timeUnixMilliStr, 10, 64)
		if err != nil {
			continue
		}
		video.CreateTime = time.UnixMilli(timeUnixMilli)
		videoList = append(videoList, video)
		inCache = append(inCache, true)
	}
	if len(notInCacheIDList) == 0 {
		// 视频全部在缓存中则提前返回
		return videoList, nil
	}
	// 批量查找不在redis的video
	var notInCacheVideoList []repository.Video
	if notInCacheVideoList, err = repository.NewVideoDaoInstance().QueryVideosByIdList(notInCacheIDList); err != nil {
		return nil, err
	}
	// 将不在redis中的video填入返回值
	idxNotInCache := 0
	for i := range videoList {
		if inCache[i] == false {
			videoList[i] = notInCacheVideoList[idxNotInCache]
			idxNotInCache++
		}
	}
	err = AddVideoListToRedis(videoList)
	if err != nil {
		return nil, err
	}
	return videoList, nil
}
