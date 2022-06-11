package repository

import (
	"dousheng/util"
	"sync"
	"time"
)

type Favourite struct {
	Id          int64     `gorm:"column:id"`
	UId         int64     `gorm:"column:uid"`
	VId         int64     `gorm:"column:vid"`
	IsFavourite bool      `gorm:"column:is_favourite"`
	CreateTime  time.Time `gorm:"column:create_time"`
	UpdateTime  time.Time `gorm:"column:update_time"`
}

func (Favourite) TableName() string {
	return "favourite"
}

type FavouriteDao struct {
}

var favouriteDao *FavouriteDao //DAO(DataAccessObject)模式
var favouriteOnce sync.Once

func NewFavouriteDaoInstance() *FavouriteDao {
	favouriteOnce.Do(
		func() {
			favouriteDao = &FavouriteDao{}
		})
	return favouriteDao
}

// QueryByVIdAndUId 如果存在，则返回fav.IsFavourite，否则返回err
func (*FavouriteDao) QueryByVIdAndUId(vid int64, uid int64) (bool, error) {
	var fav Favourite
	err := db.Where("uid = ?", uid).Where("vid = ?", vid).First(&fav).Error //链式操作
	if err != nil {
		util.Logger.Error("find favourite by vid and uid err:" + err.Error())
		return false, err
	}
	return fav.IsFavourite, nil
}

// QueryByUId 如果存在，返回列表，否则返回空，报错则返回err
func (*FavouriteDao) QueryByUId(uid int64) (*[]Favourite, error) {
	var fav []Favourite
	err := db.Where("uid = ?", uid).Find(&fav).Error
	if err != nil {
		util.Logger.Error("find favourite by id err:" + err.Error())
		return nil, err
	}
	return &fav, nil
}

// UpdateIsFavourite 若点赞了，就取消；若没有，则点赞
func (f *FavouriteDao) UpdateIsFavourite(vid int64, uid int64) error {
	isFavourite, err := f.QueryByVIdAndUId(vid, uid)
	if err != nil {
		//没有找到
		if err := f.CreateFavourite(&Favourite{
			0,
			uid,
			vid,
			true,
			time.Now(),
			time.Now(),
		}); err != nil {
			return err
		}
		return nil
	}
	db.Model(Favourite{}).Where("uid = ?", uid).Where("vid = ?", vid).Update("is_favourite", !isFavourite)
	return nil
}

func (f *FavouriteDao) CreateFavourite(fav *Favourite) error {
	if err := db.Create(fav).Error; err != nil {
		util.Logger.Error("insert favourite err:" + err.Error())
		return err
	}
	return nil
}
