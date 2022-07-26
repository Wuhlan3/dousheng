package service

import (
	"dousheng/proto/proto"
	"dousheng/repository"
	"fmt"
	"github.com/spf13/viper"
	"strconv"
)

func Feed(myUId int64, LatestTime int64) ([]*proto.Video, error) {
	path := viper.GetString("cos.uriPath")
	maxNumStr := viper.GetString("video.maxNumPerTimes")
	maxNum, err := strconv.ParseInt(maxNumStr, 10, 64)
	if err != nil {
		return nil, err
	}

	//查询Redis中是否存在 视频序号序列ZSet
	VIdList, err := GetVIdListFromRedis(LatestTime, maxNum)
	if err != nil {
		return nil, err
	}

	//如果存在，根据序号来获取视频相关信息，Hash
	videosList, err := GetVideoListFromRedis(VIdList)
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}

	//旧版本，直接查数据库
	//videosList, err := repository.NewVideoDaoInstance().QueryVideoList(int(maxNum))

	var protoVideoList []*proto.Video
	for _, video := range videosList {
		uid := video.UId

		//获取是否关注
		isFollow, err := repository.NewFollowDaoInstance().QueryIsFollowByUIdAndHisUId(myUId, uid)

		//获取user
		user, err := repository.NewUserDaoInstance().QueryUserById(uid)
		if err != nil {
			return nil, err
		}
		demoUser := &proto.User{
			Id:            user.Id,
			Name:          user.Name,
			FollowCount:   user.FollowCount,
			FollowerCount: user.FollowerCount,
			IsFollow:      isFollow,
		}

		//获取是否点赞
		IsFavourite, err := repository.NewFavouriteDaoInstance().QueryByVIdAndUId(video.Id, myUId)
		if err != nil {
			IsFavourite = false
		}

		fmt.Println(path + video.PlayUrl)
		protoVideoList = append(protoVideoList, &proto.Video{
			Id:            video.Id,
			Author:        demoUser,
			PlayUrl:       path + video.PlayUrl,
			CoverUrl:      path + video.CoverUrl,
			FavoriteCount: video.FavouriteCount,
			CommentCount:  video.CommentCount,
			IsFavorite:    IsFavourite,
			Title:         video.Title,
		})
	}

	return protoVideoList, err
}
