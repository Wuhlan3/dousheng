package service

import (
	"dousheng/repository"
	"time"
)

func UploadVideo(uid int64, fileName string) error {
	len := len(fileName)
	video := &repository.Video{
		UId:            uid,
		PlayUrl:        fileName,
		CoverUrl:       "bear.jpg",
		CommentCount:   10,
		FavouriteCount: 10,
		Title:          fileName[:len-4],
		CreateTime:     time.Now(),
		UpdateTime:     time.Now(),
		IsDeleted:      false,
	}
	if err := repository.NewVideoDaoInstance().CreateVideo(video); err != nil {
		return err
	}

	return nil
}
