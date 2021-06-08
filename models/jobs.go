package models

import (
	"pro_cfg_manager/config"
	"pro_cfg_manager/dbs"
)

type Jobs struct {
	ID      int    `json:"id" gorm:"column:id"`
	Name    string `json:"name" gorm:"column:name"`
	Port    int    `json:"port" gorm:"column:port"`
	CfgName string `json:"cfg_name" gorm:"column:cfg_name"`
}

func GetJobs() (*[]Jobs, *BriefMessage) {
	db := dbs.DBObj.GetGoRM()
	if db == nil {
		config.Log.Error(InternalGetBDInstanceErr)
		return nil, ErrDataBase
	}
	jobs := []Jobs{}
	tx := db.Table("jobs").Find(&jobs)
	if tx.Error != nil {
		config.Log.Error(tx.Error)
		return nil, ErrSearchDBData
	}
	return &jobs, Success
}
