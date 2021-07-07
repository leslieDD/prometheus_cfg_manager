package models

import (
	"fmt"
	"pro_cfg_manager/config"
	"pro_cfg_manager/dbs"
	"time"
)

type BaseLabels struct {
	ID       int       `json:"id" gorm:"column:id"`
	Label    string    `json:"label" gorm:"column:label"`
	Enabled  bool      `json:"enabled" gorm:"column:enabled"`
	UpdateAt time.Time `json:"update_at" gorm:"update_at"`
}

func GetBaseLabels(sp *SplitPage) (*ResSplitPage, *BriefMessage) {
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
	tx := db.Table("labels").Where("label like ?", fmt.Sprint("%", sp.Search, "%")).Count(&count)
	if tx.Error != nil {
		config.Log.Error(tx.Error)
		return nil, ErrCount
	}
	lists := []*BaseLabels{}
	tx2 := db.Table("labels").Where("label like ?", fmt.Sprint("%", sp.Search, "%")).
		Order("update_at desc").
		Offset((sp.PageNo - 1) * sp.PageSize).
		Limit(sp.PageSize).
		Find(&lists)
	if tx2.Error != nil {
		config.Log.Error(tx2.Error)
		return nil, ErrCount
	}
	return CalSplitPage(sp, count, lists), Success
}

func PostBaseLabels(newLabel *BaseLabels) *BriefMessage {
	db := dbs.DBObj.GetGoRM()
	if db == nil {
		config.Log.Error(InternalGetBDInstanceErr)
		return ErrDataBase
	}
	newLabel.UpdateAt = time.Now()
	tx := db.Table("labels").Create(newLabel)
	if tx.Error != nil {
		config.Log.Error(tx.Error)
		return ErrCreateDBData
	}
	return Success
}

func PutBaseLabels(label *BaseLabels) *BriefMessage {
	db := dbs.DBObj.GetGoRM()
	if db == nil {
		config.Log.Error(InternalGetBDInstanceErr)
		return ErrDataBase
	}
	tx := db.Table("labels").
		Where("id=?", label.ID).
		Update("label", label.Label).
		Update("update_at", time.Now())
	if tx.Error != nil {
		config.Log.Error(tx.Error)
		return ErrUpdateData
	}
	return Success
}

func DelBaseLabels(id int64) *BriefMessage {
	db := dbs.DBObj.GetGoRM()
	if db == nil {
		config.Log.Error(InternalGetBDInstanceErr)
		return ErrDataBase
	}
	tx := db.Table("labels").Where("id=?", id).Delete(nil)
	if tx.Error != nil {
		config.Log.Error(tx.Error)
		return ErrDelData
	}
	return Success
}

type ReLabels struct {
	ID       int       `json:"id" gorm:"column:id"`
	Name     string    `json:"name" gorm:"column:name"`
	Code     string    `json:"code" gorm:"column:code"`
	Enabled  bool      `json:"enabled" gorm:"column:enabled"`
	UpdateAt time.Time `json:"update_at" gorm:"update_at"`
}

func GetReLabels(sp *SplitPage) (*ResSplitPage, *BriefMessage) {
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
	tx := db.Table("relabels").Where("name like ?", fmt.Sprint("%", sp.Search, "%")).Count(&count)
	if tx.Error != nil {
		config.Log.Error(tx.Error)
		return nil, ErrCount
	}
	lists := []*ReLabels{}
	tx2 := db.Table("relabels").Where("name like ?", fmt.Sprint("%", sp.Search, "%")).
		Order("update_at desc").
		Offset((sp.PageNo - 1) * sp.PageSize).
		Limit(sp.PageSize).
		Find(&lists)
	if tx2.Error != nil {
		config.Log.Error(tx2.Error)
		return nil, ErrCount
	}
	return CalSplitPage(sp, count, lists), Success
}

func GetAllReLabels() ([]*ReLabels, *BriefMessage) {
	db := dbs.DBObj.GetGoRM()
	if db == nil {
		config.Log.Error(InternalGetBDInstanceErr)
		return nil, ErrDataBase
	}
	lists := []*ReLabels{}
	tx2 := db.Table("relabels").
		Order("update_at desc").
		Find(&lists)
	if tx2.Error != nil {
		config.Log.Error(tx2.Error)
		return nil, ErrCount
	}
	return lists, Success
}

func PostReLabels(rl *ReLabels) *BriefMessage {
	db := dbs.DBObj.GetGoRM()
	if db == nil {
		config.Log.Error(InternalGetBDInstanceErr)
		return ErrDataBase
	}
	rl.UpdateAt = time.Now()
	if rl.Code == "" {
		// 注意格式，如空格
		rl.Code = `    relabel_configs:
      - source_labels: ['__address__']
        regex: (.*):([\d]+)
        target_label: 'instance'
        replacement: $1`
	}
	tx := db.Table("relabels").Create(rl)
	if tx.Error != nil {
		config.Log.Error(tx.Error)
		return ErrCreateDBData
	}
	return Success
}

func PutReLabels(label *ReLabels) *BriefMessage {
	db := dbs.DBObj.GetGoRM()
	if db == nil {
		config.Log.Error(InternalGetBDInstanceErr)
		return ErrDataBase
	}
	tx := db.Table("relabels").
		Where("id=?", label.ID).
		Update("name", label.Name).
		Update("code", label.Code).
		Update("update_at", time.Now())
	if tx.Error != nil {
		config.Log.Error(tx.Error)
		return ErrUpdateData
	}
	return Success
}

func DelReLabels(id int64) *BriefMessage {
	db := dbs.DBObj.GetGoRM()
	if db == nil {
		config.Log.Error(InternalGetBDInstanceErr)
		return ErrDataBase
	}
	tx := db.Table("relabels").Where("id=?", id).Delete(nil)
	if tx.Error != nil {
		config.Log.Error(tx.Error)
		return ErrDelData
	}
	return Success
}

func PutReLabelsCode(label *ReLabels) *BriefMessage {
	db := dbs.DBObj.GetGoRM()
	if db == nil {
		config.Log.Error(InternalGetBDInstanceErr)
		return ErrDataBase
	}
	tx := db.Table("relabels").
		Where("id=?", label.ID).
		Update("code", label.Code).
		Update("update_at", time.Now())
	if tx.Error != nil {
		config.Log.Error(tx.Error)
		return ErrUpdateData
	}
	return Success
}

func PutBaseLabelsStatus(edi *EnabledInfo) *BriefMessage {
	db := dbs.DBObj.GetGoRM()
	if db == nil {
		config.Log.Error(InternalGetBDInstanceErr)
		return ErrDataBase
	}
	tx := db.Table("labels").
		Where("id=?", edi.ID).
		Update("enabled", edi.Enabled).
		Update("update_at", time.Now())
	if tx.Error != nil {
		config.Log.Error(tx.Error)
		return ErrUpdateData
	}
	return Success
}

func RelabelsHaveJobs(relabelID int) (*int64, *BriefMessage) {
	db := dbs.DBObj.GetGoRM()
	if db == nil {
		config.Log.Error(InternalGetBDInstanceErr)
		return nil, ErrDataBase
	}
	var count int64
	tx := db.Table("jobs").Where("relabel_id=?", relabelID).Count(&count)
	if tx.Error != nil {
		config.Log.Error(tx.Error)
		return nil, ErrCount
	}
	return &count, Success
}

func PutBaseRelabelsStatus(edi *EnabledInfo) *BriefMessage {
	count, bf := RelabelsHaveJobs(edi.ID)
	if bf != Success {
		return bf
	}
	if *count != 0 {
		return ErrHaveDataNoAllowToDisabled
	}
	db := dbs.DBObj.GetGoRM()
	if db == nil {
		config.Log.Error(InternalGetBDInstanceErr)
		return ErrDataBase
	}
	tx := db.Table("relabels").
		Where("id=?", edi.ID).
		Update("enabled", edi.Enabled).
		Update("update_at", time.Now())
	if tx.Error != nil {
		config.Log.Error(tx.Error)
		return ErrUpdateData
	}
	return Success
}
