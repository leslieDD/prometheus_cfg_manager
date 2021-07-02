package models

import (
	"pro_cfg_manager/config"
	"pro_cfg_manager/dbs"
	"time"

	"gorm.io/gorm"
)

type JobGroup struct {
	ID       int       `json:"id" gorm:"column:id"`
	JobsID   int       `json:"jobs_id" gorm:"jobs_id"`
	Name     string    `json:"name" gorm:"column:name"`
	UpdateAt time.Time `json:"update_at" gorm:"column:update_at"`
}

type DelJobGroupInfo struct {
	ID     int `json:"id" form:"id"`
	JobsID int `json:"jobs_id" form:"jobs_id"`
}

func GetJobGroup(sp *SplitPage) (*ResSplitPage, *BriefMessage) {
	db := dbs.DBObj.GetGoRM()
	if db == nil {
		config.Log.Error(InternalGetBDInstanceErr)
		return nil, ErrDataBase
	}
	var count int64
	var tx *gorm.DB
	tx = db.Table("job_group")
	if sp.Search != "" {
		tx = tx.Where("jobs_id=? and name like ?", sp.ID, `%`+sp.Search+`%`)
	} else {
		tx = tx.Where("jobs_id=?", sp.ID)
	}
	tx = tx.Count(&count)
	if tx.Error != nil {
		config.Log.Error(tx.Error)
		return nil, ErrSearchDBData
	}
	groups := []*JobGroup{}
	tx = db.Table("job_group")
	if sp.Search != "" {
		tx = tx.Where("jobs_id=? and name like ?", sp.ID, `%`+sp.Search+`%`)
	} else {
		tx = tx.Where("jobs_id=?", sp.ID)
	}
	tx = tx.Order("update_at desc").
		Offset((sp.PageNo - 1) * sp.PageSize).
		Limit(sp.PageSize).
		Find(&groups)
	if tx.Error != nil {
		config.Log.Error(tx.Error)
		return nil, ErrSearchDBData
	}
	return CalSplitPage(sp, count, groups), Success
}

func PostJobGroup(jb *JobGroup) *BriefMessage {
	db := dbs.DBObj.GetGoRM()
	if db == nil {
		config.Log.Error(InternalGetBDInstanceErr)
		return ErrDataBase
	}
	jb.UpdateAt = time.Now()
	tx := db.Table("job_group").Create(&jb)
	if tx.Error != nil {
		config.Log.Error(tx.Error)
		return ErrCreateDBData
	}
	return Success
}

func PutJobGroup(jb *JobGroup) *BriefMessage {
	db := dbs.DBObj.GetGoRM()
	if db == nil {
		config.Log.Error(InternalGetBDInstanceErr)
		return ErrDataBase
	}
	tx := db.Table("job_group").
		Where("jobs_id=?", jb.JobsID).
		Update("name", jb.Name)
	if tx.Error != nil {
		config.Log.Error(tx.Error)
		return ErrCreateDBData
	}
	return Success
}

func DelJobGroup(dInfo *DelJobGroupInfo) *BriefMessage {
	db := dbs.DBObj.GetGoRM()
	if db == nil {
		config.Log.Error(InternalGetBDInstanceErr)
		return ErrDataBase
	}
	tx := db.Table("job_group").
		Where("jobs_id=? and id=?", dInfo.JobsID, dInfo.ID).
		Delete(nil)
	if tx.Error != nil {
		config.Log.Error(tx.Error)
		return ErrCreateDBData
	}
	return Success
}
