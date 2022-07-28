package service

import (
	"dousheng/repository"
	"time"
)

func UploadVideo(uid int64, fileName string, imgName string) error {
	len := len(fileName)
	video := &repository.Video{
		UId:            uid,
		PlayUrl:        fileName,
		CoverUrl:       imgName,
		CommentCount:   0,
		FavouriteCount: 0,
		Title:          fileName[:len-4],
		CreateTime:     time.Now(),
		UpdateTime:     time.Now(),
		IsDeleted:      false,
	}
	if err := repository.NewVideoDaoInstance().CreateVideo(video); err != nil {
		return err
	}

	if err := AddVIdListToRedis(video); err != nil {
		return err
	}

	return nil
}
