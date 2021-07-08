package models

import (
	"pro_cfg_manager/config"
	"pro_cfg_manager/dbs"
)

type Option struct {
	OptKey   string `json:"opt_key" gorm:"column:opt_key"`
	OptValue string `json:"opt_value" gorm:"column:opt_value"`
}

func GetOptions() (map[string]string, *BriefMessage) {
	db := dbs.DBObj.GetGoRM()
	if db == nil {
		config.Log.Error(InternalGetBDInstanceErr)
		return nil, ErrDataBase
	}
	opts := []Option{}
	tx := db.Table("options").Find(&opts)
	if tx.Error != nil {
		config.Log.Error(tx.Error)
		return nil, ErrDelData
	}
	result := map[string]string{}
	for _, o := range opts {
		result[o.OptKey] = o.OptValue
	}
	return result, Success
}

func PutOptions(opts map[string]string) *BriefMessage {
	db := dbs.DBObj.GetGoRM()
	if db == nil {
		config.Log.Error(InternalGetBDInstanceErr)
		return ErrDataBase
	}
	for key, value := range opts {
		tx := db.Table("options").
			Where("opt_key", key).
			Update("opt_value", value)
		if tx.Error != nil {
			config.Log.Error(tx.Error)
			return ErrUpdateData
		}
	}
	return Success
}
