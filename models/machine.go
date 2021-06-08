package models

import (
	"pro_cfg_manager/config"
	"pro_cfg_manager/dbs"
	"strings"

	"gorm.io/datatypes"
)

type Machine struct {
	ID     int            `json:"id" gorm:"column:id"`
	IpAddr string         `json:"ipaddr" gorm:"column:ipaddr"`
	JobId  datatypes.JSON `json:"job_id" gorm:"column:job_id"`
}

func GetMachine(machineID int) (*Machine, *BriefMessage) {
	db := dbs.DBObj.GetGoRM()
	if db == nil {
		config.Log.Error(InternalGetBDInstanceErr)
		return nil, ErrDataBase
	}
	machine := Machine{}
	tx := db.Table("machines").Where("id=?", machineID).First(&machine)
	if tx.Error != nil {
		config.Log.Error(tx.Error)
		return nil, ErrSearchDBData
	}
	return &machine, Success
}

func GetMachines(sp *SplitPage) (*ResSplitPage, *BriefMessage) {
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
	tx := db.Table("machines").Count(&count)
	if tx.Error != nil {
		config.Log.Error(tx.Error)
		return nil, ErrCount
	}
	machines := []Machine{}
	tx = db.Table("machines").Offset((sp.PageNo - 1) * sp.PageSize).Limit(sp.PageSize).Find(&machines)
	if tx.Error != nil {
		config.Log.Error(tx.Error)
		return nil, ErrSearchDBData
	}
	return CalSplitPage(sp, count, machines), Success
}

func PostMachine(m *Machine) *BriefMessage {
	db := dbs.DBObj.GetGoRM()
	if db == nil {
		config.Log.Error(InternalGetBDInstanceErr)
		return ErrDataBase
	}
	tx := db.Table("machines").Create(m)
	if tx.Error != nil {
		config.Log.Error(tx.Error)
		if strings.Contains(tx.Error.Error(), "Duplicate entry") {
			return ErrDataExist
		}
		return ErrCreateDBData
	}
	return Success
}

func PutMachine(m *Machine) *BriefMessage {
	db := dbs.DBObj.GetGoRM()
	if db == nil {
		config.Log.Error(InternalGetBDInstanceErr)
		return ErrDataBase
	}
	tx := db.Table("machines").
		Where("id = ?", m.ID).
		Update("ipaddr", m.IpAddr).
		Update("job_id", m.JobId)
	if tx.Error != nil {
		config.Log.Error(tx.Error)
		return ErrUpdateData
	}
	return Success
}

func DeleteMachine(mID int) *BriefMessage {
	db := dbs.DBObj.GetGoRM()
	if db == nil {
		config.Log.Error(InternalGetBDInstanceErr)
		return ErrDataBase
	}
	tx := db.Table("machines").Where("id=?", mID).Delete(nil)
	if tx.Error != nil {
		config.Log.Error(tx.Error)
		return ErrUpdateData
	}
	return Success
}
