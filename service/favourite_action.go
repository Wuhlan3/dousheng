package service

import (
	"dousheng/repository"
	"errors"
)

type FavouriteActionFlow struct {
	UId int64
	VId int64
}

func FavouriteAction(userId int64, videoId int64) error {
	return NewFavouriteActionFlow(userId, videoId).Do()
}

func NewFavouriteActionFlow(userId int64, videoId int64) *FavouriteActionFlow {
	return &FavouriteActionFlow{
		UId: userId,
		VId: videoId,
	}
}

func (f *FavouriteActionFlow) Do() error {
	if err := f.checkParam(); err != nil {
		return err
	}
	err := f.action()
	if err != nil {
		return err
	}
	return nil
}

func (f *FavouriteActionFlow) checkParam() error {
	return nil
}

func (f *FavouriteActionFlow) action() error {
	err := repository.NewFavouriteDaoInstance().UpdateIsFavourite(f.VId, f.UId)
	if err != nil {
		return errors.New("修改失败")
	}
	return nil
}
