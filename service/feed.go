package service

import (
	"dousheng/proto/proto"
	"dousheng/repository"
	"fmt"
	"github.com/spf13/viper"
	"strconv"
	"time"
)

func Feed(myUId int64) ([]*proto.Video, error) {
	path := viper.GetString("video.absolutePath")
	maxNumStr := viper.GetString("video.maxNumPerTimes")
	maxNum, err := strconv.ParseInt(maxNumStr, 10, 64)
	if err != nil {
		return nil, err
	}
	videosList, err := repository.NewVideoDaoInstance().QueryVideoList(int(maxNum))
	var protoVideoList []*proto.Video
	for _, video := range *videosList {
		uid := video.UId

		//获取是否关注
		var IsFollow bool
		follow, err := repository.NewFollowDaoInstance().QueryByUIdAndHisUId(myUId, uid)
		if err != nil {
			err := repository.NewFollowDaoInstance().CreateFollow(&repository.Follow{
				MyId:       myUId,
				HisId:      uid,
				IsFollow:   false,
				CreateTime: time.Now(),
				UpdateTime: time.Now(),
			})
			if err != nil {
				return nil, err
			}
			IsFollow = false
		} else {
			IsFollow = follow.IsFollow
		}

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
			IsFollow:      IsFollow,
		}

		//获取是否点赞
		IsFavourite, err := repository.NewFavouriteDaoInstance().QueryByVIdAndUId(video.Id, myUId)
		if err != nil {
			IsFavourite = false
		}

		fmt.Println(path + ":" + video.PlayUrl)
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
