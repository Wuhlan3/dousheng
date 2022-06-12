package repository

import (
	"dousheng/util"
	"sync"
	"time"
)

type Follow struct {
	Id         int64     `gorm:"column:id"`
	MyId       int64     `gorm:"column:my_uid"`
	HisId      int64     `gorm:"column:his_uid"`
	IsFollow   bool      `gorm:"column:is_follow"`
	CreateTime time.Time `gorm:"column:create_time"`
	UpdateTime time.Time `gorm:"column:update_time"`
}

func (Follow) TableName() string {
	return "follow"
}

type FollowDao struct {
}

var followDao *FollowDao //DAO(DataAccessObject)模式
var followOnce sync.Once

func NewFollowDaoInstance() *FollowDao {
	followOnce.Do(
		func() {
			followDao = &FollowDao{}
		})
	return followDao
}

func (c *FollowDao) QueryByUId(myUId int64) (*[]Follow, error) {
	var followList []Follow
	err := db.Where("my_uid = ?", myUId).Find(&followList).Error
	if err != nil {
		util.Logger.Error("find followList by UId err:" + err.Error())
		return nil, err
	}
	return &followList, nil
}

func (c *FollowDao) QueryByHisUId(UId int64) (*[]Follow, error) {
	var followList []Follow
	err := db.Where("his_uid = ?", UId).Find(&followList).Error
	if err != nil {
		util.Logger.Error("find followList by UId err:" + err.Error())
		return nil, err
	}
	return &followList, nil
}

func (c *FollowDao) QueryByUIdAndHisUId(myUId int64, hisUId int64) (*Follow, error) {
	var follow Follow
	err := db.Where("my_uid = ?", myUId).Where("his_uid = ?", hisUId).First(&follow).Error
	if err != nil {
		util.Logger.Error("find followList by UId err:" + err.Error())
		return nil, err
	}
	return &follow, nil
}

func (c *FollowDao) UpdateFollow(myUId int64, hisUId int64, isFollow bool) error {
	err := db.Model(Follow{}).Where("my_uid = ?", myUId).Where("his_uid = ?", hisUId).Update("is_follow", isFollow).Error
	if err != nil {
		util.Logger.Error("update follow err:" + err.Error())
		return err
	}
	return nil
}

func (c *FollowDao) CreateFollow(follow *Follow) error {
	if err := db.Create(follow).Error; err != nil {
		util.Logger.Error("insert relation err:" + err.Error())
		return err
	}
	return nil
}
