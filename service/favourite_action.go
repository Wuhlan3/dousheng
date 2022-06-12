package service

import (
	"dousheng/repository"
	"errors"
	"time"
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
	isFavourite, err := repository.NewFavouriteDaoInstance().QueryByVIdAndUId(f.VId, f.UId)
	if err != nil {
		//没有找到
		if err := repository.NewFavouriteDaoInstance().CreateFavourite(&repository.Favourite{
			UId:         f.UId,
			VId:         f.VId,
			IsFavourite: true,
			CreateTime:  time.Now(),
			UpdateTime:  time.Now(),
		}); err != nil {
			return err
		}

		err := repository.NewVideoDaoInstance().IncFavouriteCount(f.VId)
		if err != nil {
			return err
		}

		return nil
	}
	if isFavourite == true {
		err := repository.NewVideoDaoInstance().DecFavouriteCount(f.VId)
		if err != nil {
			return err
		}
	} else {
		err := repository.NewVideoDaoInstance().IncFavouriteCount(f.VId)
		if err != nil {
			return err
		}
	}
	err = repository.NewFavouriteDaoInstance().UpdateIsFavourite(f.VId, f.UId, !isFavourite)

	if err != nil {
		return errors.New("修改失败")
	}
	return nil
}
