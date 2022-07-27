package service

import (
	"dousheng/proto/proto"
	"dousheng/repository"
	"fmt"
	"github.com/spf13/viper"
)

func PublishList(uid int64) ([]*proto.Video, error) {
	//path := viper.GetString("video.absolutePath")
	videoPath := viper.GetString("cos.uriVideoPath")
	imgPath := viper.GetString("cos.uriPicturePath")
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
		fmt.Println(videoPath + ":" + video.PlayUrl)

		//获取是否点赞
		isFavourite, err := repository.NewFavouriteDaoInstance().QueryByVIdAndUId(video.Id, uid)
		if err != nil {
			isFavourite = false
		}

		protoVideoList = append(protoVideoList, &proto.Video{
			Id:            video.Id,
			Author:        demoUser,
			PlayUrl:       videoPath + video.PlayUrl,
			CoverUrl:      imgPath + video.CoverUrl,
			FavoriteCount: video.FavouriteCount,
			CommentCount:  video.CommentCount,
			IsFavorite:    isFavourite,
			Title:         video.Title,
		})
	}

	return protoVideoList, err
}
