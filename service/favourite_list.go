package service

import (
	"dousheng/proto/proto"
	"dousheng/repository"
	"fmt"
	"github.com/spf13/viper"
	"time"
)

func FavouriteList(uid int64) ([]*proto.Video, error) {
	path := viper.GetString("video.absolutePath")
	favouriteList, err := repository.NewFavouriteDaoInstance().QueryByUId(uid)
	var protoVideoList []*proto.Video
	for _, fav := range *favouriteList {
		if fav.IsFavourite {
			user, err := repository.NewUserDaoInstance().QueryUserById(uid)
			if err != nil {
				return nil, err
			}

			video, err := repository.NewVideoDaoInstance().QueryVideoById(fav.VId)
			if err != nil {
				return nil, err
			}

			var IsFollow bool

			follow, err := repository.NewFollowDaoInstance().QueryByUIdAndHisUId(uid, video.UId)
			if err != nil {
				repository.NewFollowDaoInstance().CreateFollow(&repository.Follow{
					MyId:       uid,
					HisId:      video.Id,
					IsFollow:   false,
					CreateTime: time.Now(),
					UpdateTime: time.Now(),
				})
				IsFollow = false
			} else {
				IsFollow = follow.IsFollow
			}
			demoUser := &proto.User{
				Id:            user.Id,
				Name:          user.Name,
				FollowCount:   user.FollowCount,
				FollowerCount: user.FollowerCount,
				IsFollow:      IsFollow,
			}

			fmt.Println(path + ":" + video.PlayUrl)
			protoVideoList = append(protoVideoList, &proto.Video{
				Id:            video.Id,
				Author:        demoUser,
				PlayUrl:       path + video.PlayUrl,
				CoverUrl:      path + video.CoverUrl,
				FavoriteCount: video.FavouriteCount,
				CommentCount:  video.CommentCount,
				IsFavorite:    true,
				Title:         video.Title,
			})
		}
	}

	return protoVideoList, err
}
