package service

import (
	"dousheng/proto/proto"
	"dousheng/repository"
	"fmt"
	"github.com/spf13/viper"
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
			demoUser := &proto.User{
				Id:            user.Id,
				Name:          user.Name,
				FollowCount:   user.FollowCount,
				FollowerCount: user.FollowerCount,
				IsFollow:      false,
			}
			video, err := repository.NewVideoDaoInstance().QueryVideoById(fav.VId)
			if err != nil {
				return nil, err
			}
			fmt.Println(path + ":" + video.PlayUrl)
			protoVideoList = append(protoVideoList, &proto.Video{
				Id:            video.Id,
				Author:        demoUser,
				PlayUrl:       path + video.PlayUrl,
				CoverUrl:      path + video.CoverUrl,
				FavoriteCount: video.FavouriteCount,
				CommentCount:  video.CommentCount,
				IsFavorite:    false,
				Title:         video.Title,
			})
		}
	}

	return protoVideoList, err
}
