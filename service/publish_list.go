package service

import (
	"dousheng/proto/proto"
	"dousheng/repository"
	"fmt"
	"github.com/spf13/viper"
)

func PublishList(uid int64) ([]*proto.Video, error) {
	path := viper.GetString("video.absolutePath")
	videosList, err := repository.NewVideoDaoInstance().QueryVideoListByUId(uid)
	var protoVideoList []*proto.Video
	for _, video := range *videosList {
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

	return protoVideoList, err
}
