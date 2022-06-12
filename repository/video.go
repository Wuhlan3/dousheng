package repository

import (
	"dousheng/util"
	"gorm.io/gorm"
	"sync"
	"time"
)

type Video struct {
	Id             int64     `gorm:"column:id"`
	UId            int64     `gorm:"column:uid"`
	PlayUrl        string    `gorm:"column:play_url"`
	CoverUrl       string    `gorm:"column:cover_url"`
	CommentCount   int64     `gorm:"column:comment_count"`
	FavouriteCount int64     `gorm:"column:favourite_count"`
	Title          string    `gorm:"column:title"`
	CreateTime     time.Time `gorm:"column:create_time"`
	UpdateTime     time.Time `gorm:"column:update_time"`
	IsDeleted      bool      `gorm:"column:is_deleted"`
}

func (Video) TableName() string {
	return "video"
}

type VideoDao struct {
}

var videoDao *VideoDao //DAO(DataAccessObject)模式
var videoOnce sync.Once

func NewVideoDaoInstance() *VideoDao {
	videoOnce.Do(
		func() {
			videoDao = &VideoDao{}
		})
	return videoDao
}

func (*VideoDao) QueryVideoList(num int) (*[]Video, error) {
	var videoList []Video
	err := db.Limit(num).Find(&videoList).Error
	if err != nil {
		util.Logger.Error("find some video err:" + err.Error())
		return nil, err
	}
	return &videoList, nil
}

func (*VideoDao) QueryVideoListByUId(uid int64) (*[]Video, error) {
	var videoList []Video
	err := db.Where("uid = ?", uid).Find(&videoList).Error
	if err != nil {
		util.Logger.Error("find videos by uid err:" + err.Error())
		return nil, err
	}
	return &videoList, nil
}

func (*VideoDao) QueryVideoById(vid int64) (*Video, error) {
	var videoList Video
	err := db.Where("id = ?", vid).Find(&videoList).Error
	if err != nil {
		util.Logger.Error("find video by vid err:" + err.Error())
		return nil, err
	}
	return &videoList, nil
}

func (*VideoDao) CreateVideo(video *Video) error {
	if err := db.Create(video).Error; err != nil {
		util.Logger.Error("create video err:" + err.Error())
		return err
	}
	return nil
}

func (*VideoDao) IncFavouriteCount(vid int64) error {
	err := db.Model(Video{}).Where("id = ?", vid).UpdateColumn("favourite_count", gorm.Expr("favourite_count + ?", 1)).Error
	if err != nil {
		util.Logger.Error("inc video favourite count error")
		return err
	}
	return nil
}

func (*VideoDao) DecFavouriteCount(vid int64) error {
	err := db.Model(Video{}).Where("id = ?", vid).UpdateColumn("favourite_count", gorm.Expr("favourite_count - ?", 1)).Error
	if err != nil {
		util.Logger.Error("dec video favourite count error")
		return err
	}
	return nil
}

func (*VideoDao) IncCommentCount(vid int64) error {
	err := db.Model(Video{}).Where("id = ?", vid).UpdateColumn("comment_count", gorm.Expr("comment_count + ?", 1)).Error
	if err != nil {
		util.Logger.Error("inc video comment count error")
		return err
	}
	return nil
}

func (*VideoDao) DecCommentCount(vid int64) error {
	err := db.Model(Video{}).Where("id = ?", vid).UpdateColumn("comment_count", gorm.Expr("comment_count - ?", 1)).Error
	if err != nil {
		util.Logger.Error("dec video comment count error")
		return err
	}
	return nil
}
