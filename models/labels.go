package models

import (
	"fmt"
	"pro_cfg_manager/config"
	"pro_cfg_manager/dbs"
	"time"

	"gorm.io/gorm"
)

type BaseLabels struct {
	ID       int       `json:"id" gorm:"column:id"`
	Label    string    `json:"label" gorm:"column:label"`
	Enabled  bool      `json:"enabled" gorm:"column:enabled"`
	UpdateAt time.Time `json:"update_at" gorm:"update_at"`
	UpdateBy string    `json:"update_by" gorm:"update_by"`
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

func PostBaseLabels(user *UserSessionInfo, newLabel *BaseLabels) *BriefMessage {
	db := dbs.DBObj.GetGoRM()
	if db == nil {
		config.Log.Error(InternalGetBDInstanceErr)
		return ErrDataBase
	}
	newLabel.UpdateAt = time.Now()
	newLabel.UpdateBy = user.NiceName
	tx := db.Table("labels").Create(newLabel)
	if tx.Error != nil {
		config.Log.Error(tx.Error)
		return ErrCreateDBData
	}
	return Success
}

func PutBaseLabels(user *UserSessionInfo, label *BaseLabels) *BriefMessage {
	db := dbs.DBObj.GetGoRM()
	if db == nil {
		config.Log.Error(InternalGetBDInstanceErr)
		return ErrDataBase
	}
	tx := db.Table("labels").
		Where("id=?", label.ID).
		Update("label", label.Label).
		Update("update_at", time.Now()).
		Update("update_by", user.Username)
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

type BaseFields struct {
	ID       int       `json:"id" gorm:"column:id"`
	Key      string    `json:"key" gorm:"column:key"`
	Value    string    `json:"value" gorm:"column:value"`
	Enabled  bool      `json:"enabled" gorm:"column:enabled"`
	UpdateAt time.Time `json:"update_at" gorm:"column:update_at"`
	UpdateBy string    `json:"update_by" gorm:"column:update_by"`
}

type BaseFieldsNumber struct {
	Key   string `json:"key" gorm:"column:key"`
	Value int    `json:"value" gorm:"column:value"`
}

func GetBaseFields(sp *SplitPage) (*ResSplitPage, *BriefMessage) {
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
	tx = db.Table("tmpl_fields")
	if sp.Search != "" {
		likeContent := fmt.Sprint("%", sp.Search, "%")
		likeSql = fmt.Sprint("key like ? or value like ?", likeContent, likeContent)
		tx = db.Where(likeSql)
	}
	tx = tx.Count(&count)
	if tx.Error != nil {
		config.Log.Error(tx.Error)
		return nil, ErrCount
	}
	lists := []*BaseFields{}
	tx = db.Table("tmpl_fields")
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

func PostBaseFields(user *UserSessionInfo, newFields *BaseFields) *BriefMessage {
	db := dbs.DBObj.GetGoRM()
	if db == nil {
		config.Log.Error(InternalGetBDInstanceErr)
		return ErrDataBase
	}
	newFields.UpdateAt = time.Now()
	newFields.UpdateBy = user.Username
	tx := db.Table("tmpl_fields").Create(newFields)
	if tx.Error != nil {
		config.Log.Error(tx.Error)
		return ErrCreateDBData
	}
	return Success
}

func PutBaseFields(user *UserSessionInfo, fields *BaseFields) *BriefMessage {
	db := dbs.DBObj.GetGoRM()
	if db == nil {
		config.Log.Error(InternalGetBDInstanceErr)
		return ErrDataBase
	}
	tx := db.Table("tmpl_fields").
		Where("id=?", fields.ID).
		Update("key", fields.Key).
		Update("value", fields.Value).
		Update("update_at", time.Now()).
		Update("update_by", user.Username)
	if tx.Error != nil {
		config.Log.Error(tx.Error)
		return ErrUpdateData
	}
	return Success
}

func DelBaseFields(id int64) *BriefMessage {
	db := dbs.DBObj.GetGoRM()
	if db == nil {
		config.Log.Error(InternalGetBDInstanceErr)
		return ErrDataBase
	}
	tx := db.Table("tmpl_fields").Where("id=?", id).Delete(nil)
	if tx.Error != nil {
		config.Log.Error(tx.Error)
		return ErrDelData
	}
	return Success
}

func PutBaseFieldsStatus(user *UserSessionInfo, edi *EnabledInfo) *BriefMessage {
	db := dbs.DBObj.GetGoRM()
	if db == nil {
		config.Log.Error(InternalGetBDInstanceErr)
		return ErrDataBase
	}
	tx := db.Table("tmpl_fields").
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

type ReLabels struct {
	ID       int       `json:"id" gorm:"column:id"`
	Name     string    `json:"name" gorm:"column:name"`
	Code     string    `json:"code" gorm:"column:code"`
	Enabled  bool      `json:"enabled" gorm:"column:enabled"`
	UpdateAt time.Time `json:"update_at" gorm:"column:update_at"`
	UpdateBy string    `json:"update_by" gorm:"column:update_by"`
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
		// Order("update_at desc").
		Offset((sp.PageNo - 1) * sp.PageSize).
		Limit(sp.PageSize).
		Find(&lists)
	if tx2.Error != nil {
		config.Log.Error(tx2.Error)
		return nil, ErrCount
	}
	return CalSplitPage(sp, count, lists), Success
}

func GetReLabel(id *OnlyID) (*ReLabels, *BriefMessage) {
	db := dbs.DBObj.GetGoRM()
	if db == nil {
		config.Log.Error(InternalGetBDInstanceErr)
		return nil, ErrDataBase
	}
	relabel := ReLabels{}
	tx := db.Table("relabels").Where("id=?", id.ID).Find(&relabel)
	if tx.Error != nil {
		config.Log.Error(tx.Error)
		return nil, ErrSearchDBData
	}
	return &relabel, Success
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

func PostReLabels(user *UserSessionInfo, rl *ReLabels) *BriefMessage {
	db := dbs.DBObj.GetGoRM()
	if db == nil {
		config.Log.Error(InternalGetBDInstanceErr)
		return ErrDataBase
	}
	rl.UpdateAt = time.Now()
	rl.UpdateBy = user.Username
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

func PutReLabels(user *UserSessionInfo, label *ReLabels) *BriefMessage {
	db := dbs.DBObj.GetGoRM()
	if db == nil {
		config.Log.Error(InternalGetBDInstanceErr)
		return ErrDataBase
	}
	tx := db.Table("relabels").
		Where("id=?", label.ID).
		Update("name", label.Name).
		// Update("code", label.Code).
		Update("update_at", time.Now()).
		Update("update_by", user.Username)
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

func PutReLabelsCode(user *UserSessionInfo, label *ReLabels) *BriefMessage {
	db := dbs.DBObj.GetGoRM()
	if db == nil {
		config.Log.Error(InternalGetBDInstanceErr)
		return ErrDataBase
	}
	tx := db.Table("relabels").
		Where("id=?", label.ID).
		Update("code", label.Code).
		Update("update_at", time.Now()).
		Update("update_by", user.Username)
	if tx.Error != nil {
		config.Log.Error(tx.Error)
		return ErrUpdateData
	}
	return Success
}

func PutBaseLabelsStatus(user *UserSessionInfo, edi *EnabledInfo) *BriefMessage {
	db := dbs.DBObj.GetGoRM()
	if db == nil {
		config.Log.Error(InternalGetBDInstanceErr)
		return ErrDataBase
	}
	tx := db.Table("labels").
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

func PutBaseRelabelsStatus(user *UserSessionInfo, edi *EnabledInfo) *BriefMessage {
	db := dbs.DBObj.GetGoRM()
	if db == nil {
		config.Log.Error(InternalGetBDInstanceErr)
		return ErrDataBase
	}
	tx := db.Table("relabels").
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
