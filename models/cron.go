package models

import (
	"fmt"
	"pro_cfg_manager/config"
	"pro_cfg_manager/dbs"
	"time"

	"gorm.io/gorm"
)

type CronApi struct {
	ID       int       `json:"id" gorm:"column:id"`
	Name     string    `json:"name" gorm:"column:name"`
	API      string    `json:"api" gorm:"column:api"`
	UpdateAt time.Time `json:"update_at" gorm:"column:update_at"`
	UpdateBy string    `json:"update_by" gorm:"column:update_by"`
}

type CronApiPost struct {
	ID   int    `json:"id" gorm:"column:id"`
	Name string `json:"name" gorm:"column:name"`
	API  string `json:"api" gorm:"column:api"`
}

func GetCronApi(sp *SplitPage) (*ResSplitPage, *BriefMessage) {
	if sp.PageSize <= 0 {
		sp.PageSize = 15
	}
	if sp.PageNo <= 0 {
		sp.PageNo = 1
	}
	db := dbs.DBObj.GetGoRM()
	if db == nil {
		config.Log.Error(InternalGetBDInstanceErr)
		return nil, ErrDataBase
	}
	var count int64
	var likeSql string
	var tx *gorm.DB
	tx = db.Table("crontab_api")
	if sp.Search != "" {
		likeContent := fmt.Sprint("%", sp.Search, "%")
		likeSql = fmt.Sprint("name like ? or api like ?", likeContent, likeContent)
		tx = db.Where(likeSql)
	}
	tx = tx.Count(&count)
	if tx.Error != nil {
		config.Log.Error(tx.Error)
		return nil, ErrCount
	}
	lists := []*CronApi{}
	tx = db.Table("crontab_api")
	if sp.Search != "" {
		tx = db.Where(likeSql)
	}
	tx = tx.Order("update_at desc").
		Offset((sp.PageNo - 1) * sp.PageSize).
		Limit(sp.PageSize).
		Find(&lists)
	if tx.Error != nil {
		config.Log.Error(tx.Error)
		return nil, ErrCount
	}
	return CalSplitPage(sp, count, lists), Success
}

func PostCronApi(user *UserSessionInfo, newApi *CronApiPost) *BriefMessage {
	db := dbs.DBObj.GetGoRM()
	if db == nil {
		config.Log.Error(InternalGetBDInstanceErr)
		return ErrDataBase
	}
	cronApi := CronApi{
		ID:       0,
		Name:     newApi.Name,
		API:      newApi.API,
		UpdateAt: time.Now(),
		UpdateBy: user.Username,
	}
	tx := db.Table("crontab_api").Create(&cronApi)
	if tx.Error != nil {
		config.Log.Error(tx.Error)
		return ErrCreateDBData
	}
	return Success
}

func PutCronApi(user *UserSessionInfo, api *CronApiPost) *BriefMessage {
	db := dbs.DBObj.GetGoRM()
	if db == nil {
		config.Log.Error(InternalGetBDInstanceErr)
		return ErrDataBase
	}
	tx := db.Table("crontab_api").
		Where("id=?", api.ID).
		Update("name", api.Name).
		Update("api", api.API).
		Update("update_at", time.Now()).
		Update("update_by", user.Username)
	if tx.Error != nil {
		config.Log.Error(tx.Error)
		return ErrUpdateData
	}
	return Success
}

func PutCronApiStatus(user *UserSessionInfo, edi *EnabledInfo) *BriefMessage {
	db := dbs.DBObj.GetGoRM()
	if db == nil {
		config.Log.Error(InternalGetBDInstanceErr)
		return ErrDataBase
	}
	tx := db.Table("crontab_api").
		Where("id=?", edi.ID).
		Update("enabled", edi.Enabled).
		Update("update_at", time.Now()).
		Update("update_by", user.Username)
	if tx.Error != nil {
		config.Log.Error(tx.Error)
		return ErrUpdateData
	}
	return Success
}

func DelCronApi(id int64) *BriefMessage {
	db := dbs.DBObj.GetGoRM()
	if db == nil {
		config.Log.Error(InternalGetBDInstanceErr)
		return ErrDataBase
	}
	tx := db.Table("crontab_api").Where("id=?", id).Delete(nil)
	if tx.Error != nil {
		config.Log.Error(tx.Error)
		return ErrDelData
	}
	return Success
}
