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
	var likeContent string
	var tx *gorm.DB
	tx = db.Table("crontab_api")
	if sp.Search != "" {
		likeContent = fmt.Sprint("%", sp.Search, "%")
		tx = tx.Where("name like ? or api like ?", likeContent, likeContent)
	}
	tx = tx.Count(&count)
	if tx.Error != nil {
		config.Log.Error(tx.Error)
		return nil, ErrCount
	}
	lists := []*CronApi{}
	tx = db.Table("crontab_api")
	if sp.Search != "" {
		tx = tx.Where("name like ? or api like ?", likeContent, likeContent)
	}
	tx = tx.Offset((sp.PageNo - 1) * sp.PageSize).
		Limit(sp.PageSize).
		Find(&lists)
	if tx.Error != nil {
		config.Log.Error(tx.Error)
		return nil, ErrCount
	}
	return CalSplitPage(sp, count, lists), Success
}

func GetAllCronApi() ([]*CronApi, *BriefMessage) {
	db := dbs.DBObj.GetGoRM()
	if db == nil {
		config.Log.Error(InternalGetBDInstanceErr)
		return nil, ErrDataBase
	}
	lists := []*CronApi{}
	tx := db.Table("crontab_api").Order("update_at desc").Find(&lists)
	if tx.Error != nil {
		config.Log.Error(tx.Error)
		return nil, ErrCount
	}
	return lists, Success
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

type Crontab struct {
	ID        int       `json:"id" gorm:"column:id"`
	Name      string    `json:"name" gorm:"column:name"`
	Rule      string    `json:"rule" gorm:"column:rule"`
	Enabled   bool      `json:"enabled" gorm:"column:enabled"`
	ApiID     int       `json:"api_id" gorm:"column:api_id"`
	ExecCycle int       `json:"exec_cycle" gorm:"column:exec_cycle"` // 执行周期
	UpdateAt  time.Time `json:"update_at" gorm:"column:update_at"`
	UpdateBy  string    `json:"update_by" gorm:"column:update_by"`
}

type CrontabSplit struct {
	Crontab
	RunTimes     int `json:"run_times" gorm:"-"`
	SuccessTimes int `json:"success_times" gorm:"-"`
	FailTimes    int `json:"fail_times" gorm:"-"`
}

type CrontabPost struct {
	ID        int    `json:"id" gorm:"column:id"`
	Name      string `json:"name" gorm:"column:name"`
	Rule      string `json:"rule" gorm:"column:rule"`
	ApiID     int    `json:"api_id" gorm:"column:api_id"`
	ExecCycle int    `json:"exec_cycle" gorm:"column:exec_cycle"` // 执行周期
	Enabled   bool   `json:"enabled" gorm:"column:enabled"`
}

func GetCronRules(sp *SplitPage) (*ResSplitPage, *BriefMessage) {
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
	var likeContent string
	var tx *gorm.DB
	tx = db.Table("crontab")
	if sp.Search != "" {
		likeContent = fmt.Sprint("%", sp.Search, "%")
		tx = tx.Where("name like ? or rule like ?", likeContent, likeContent)
	}
	tx = tx.Count(&count)
	if tx.Error != nil {
		config.Log.Error(tx.Error)
		return nil, ErrCount
	}
	lists := []*CrontabSplit{}
	tx = db.Table("crontab")
	if sp.Search != "" {
		tx = tx.Where("name like ? or rule like ?", likeContent, likeContent)
	}
	tx = tx.Offset((sp.PageNo - 1) * sp.PageSize).
		Limit(sp.PageSize).
		Find(&lists)
	if tx.Error != nil {
		config.Log.Error(tx.Error)
		return nil, ErrCount
	}
	return CalSplitPage(sp, count, lists), Success
}

func PostCronRule(user *UserSessionInfo, newCron *CrontabPost) *BriefMessage {
	db := dbs.DBObj.GetGoRM()
	if db == nil {
		config.Log.Error(InternalGetBDInstanceErr)
		return ErrDataBase
	}
	cronApi := Crontab{
		ID:        0,
		Name:      newCron.Name,
		Rule:      newCron.Rule,
		ApiID:     newCron.ApiID,
		Enabled:   newCron.Enabled,
		ExecCycle: newCron.ExecCycle,
		UpdateAt:  time.Now(),
		UpdateBy:  user.Username,
	}
	tx := db.Table("crontab").Create(&cronApi)
	if tx.Error != nil {
		config.Log.Error(tx.Error)
		return ErrCreateDBData
	}
	return Success
}

func PutCronRule(user *UserSessionInfo, api *CrontabPost) *BriefMessage {
	db := dbs.DBObj.GetGoRM()
	if db == nil {
		config.Log.Error(InternalGetBDInstanceErr)
		return ErrDataBase
	}
	tx := db.Table("crontab").
		Where("id=?", api.ID).
		Update("name", api.Name).
		Update("rule", api.Rule).
		Update("api_id", api.ApiID).
		Update("enabled", api.Enabled).
		Update("exec_cycle", api.ExecCycle).
		Update("update_at", time.Now()).
		Update("update_by", user.Username)
	if tx.Error != nil {
		config.Log.Error(tx.Error)
		return ErrUpdateData
	}
	return Success
}

func DelCronRule(id int64) *BriefMessage {
	db := dbs.DBObj.GetGoRM()
	if db == nil {
		config.Log.Error(InternalGetBDInstanceErr)
		return ErrDataBase
	}
	tx := db.Table("crontab").Where("id=?", id).Delete(nil)
	if tx.Error != nil {
		config.Log.Error(tx.Error)
		return ErrDelData
	}
	return Success
}

func PutCronRuleStatus(user *UserSessionInfo, edi *EnabledInfo) *BriefMessage {
	db := dbs.DBObj.GetGoRM()
	if db == nil {
		config.Log.Error(InternalGetBDInstanceErr)
		return ErrDataBase
	}
	tx := db.Table("crontab").
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

func DelRulesSelection(delIDs []int) *BriefMessage {
	db := dbs.DBObj.GetGoRM()
	if db == nil {
		config.Log.Error(InternalGetBDInstanceErr)
		return ErrDataBase
	}
	err := db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Table("crontab").Where("id in (?)", delIDs).Delete(nil).Error; err != nil {
			config.Log.Error(tx.Error)
			return err
		}
		return nil
	})
	if err != nil {
		return ErrDelData
	}
	return Success
}

func PutRulesEnable(user *UserSessionInfo, enabledIDs []int) *BriefMessage {
	db := dbs.DBObj.GetGoRM()
	if db == nil {
		config.Log.Error(InternalGetBDInstanceErr)
		return ErrDataBase
	}
	tx := db.Table("crontab").Where("id in (?)", enabledIDs).
		Update("enabled", 1).
		Update("update_at", time.Now()).
		Update("update_by", user.Username)
	if tx.Error != nil {
		config.Log.Error(tx.Error)
		return ErrUpdateData
	}
	return Success
}

func PutRulesDisable(user *UserSessionInfo, disableIDs []int) *BriefMessage {
	db := dbs.DBObj.GetGoRM()
	if db == nil {
		config.Log.Error(InternalGetBDInstanceErr)
		return ErrDataBase
	}
	tx := db.Table("crontab").Where("id in (?)", disableIDs).
		Update("enabled", 0).
		Update("update_at", time.Now()).
		Update("update_by", user.Username)
	if tx.Error != nil {
		config.Log.Error(tx.Error)
		return ErrUpdateData
	}
	return Success
}
