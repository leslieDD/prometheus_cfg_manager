package models

import (
	"fmt"
	"pro_cfg_manager/config"
	"pro_cfg_manager/dbs"
	"time"

	"gorm.io/gorm"
)

type Jobs struct {
	ID           int       `json:"id" gorm:"column:id"`
	Name         string    `json:"name" gorm:"column:name"`
	Port         int       `json:"port" gorm:"column:port"`
	CfgName      string    `json:"cfg_name" gorm:"column:cfg_name"`
	IsCommon     bool      `json:"is_common" gorm:"column:is_common"`
	DisplayOrder int       `json:"display_order" gorm:"display_order"`
	UpdateAt     time.Time `json:"update_at" gorm:"update_at"`
}

type OnlyID struct {
	ID int `json:"id" gorm:"column:id"`
}

type SwapInfo struct {
	ID           int    `json:"id" gorm:"column:id"`
	DisplayOrder int    `json:"display_order" gorm:"display_order"`
	Action       string `json:"action" gorm:"action"`
}

func GetJobs() (*[]Jobs, *BriefMessage) {
	db := dbs.DBObj.GetGoRM()
	if db == nil {
		config.Log.Error(InternalGetBDInstanceErr)
		return nil, ErrDataBase
	}
	jobs := []Jobs{}
	tx := db.Table("jobs").Order("display_order asc").Where("is_common=0").Find(&jobs)
	if tx.Error != nil {
		config.Log.Error(tx.Error)
		return nil, ErrSearchDBData
	}
	return &jobs, Success
}

func GetJobsSplit(sp *SplitPage) (*ResSplitPage, *BriefMessage) {
	db := dbs.DBObj.GetGoRM()
	if db == nil {
		config.Log.Error(InternalGetBDInstanceErr)
		return nil, ErrDataBase
	}
	var count int64
	var tx *gorm.DB
	tx = db.Table("jobs")
	if sp.Search != "" {
		tx = tx.Where("name like ? and is_common=0", fmt.Sprint("%", sp.Search, "%"))
	} else {
		tx = tx.Where("is_common=0")
	}
	tx = tx.Count(&count)
	if tx.Error != nil {
		config.Log.Error(tx.Error)
		return nil, ErrCount
	}
	jobs := []Jobs{}
	tx = db.Table("jobs")
	if sp.Search != "" {
		tx = tx.Where("name like ? and is_common=0", fmt.Sprint("%", sp.Search, "%"))
	} else {
		tx = tx.Where("is_common=0")
	}
	tx = tx.Order("display_order asc").
		Offset((sp.PageNo - 1) * sp.PageSize).
		Limit(sp.PageSize).
		Find(&jobs)
	if tx.Error != nil {
		config.Log.Error(tx.Error)
		return nil, ErrSearchDBData
	}
	return CalSplitPage(sp, count, jobs), Success
}

func GetJob(jID int64) (*Jobs, *BriefMessage) {
	db := dbs.DBObj.GetGoRM()
	if db == nil {
		config.Log.Error(InternalGetBDInstanceErr)
		return nil, ErrDataBase
	}
	job := Jobs{}
	tx := db.Table("jobs").Where("id=? and is_common=0").Find(&job)
	if tx.Error != nil {
		config.Log.Error(tx.Error)
		return nil, ErrSearchDBData
	}
	return &job, Success
}

func PostJob(job *Jobs) *BriefMessage {
	db := dbs.DBObj.GetGoRM()
	if db == nil {
		config.Log.Error(InternalGetBDInstanceErr)
		return ErrDataBase
	}
	job.IsCommon = false
	job.UpdateAt = time.Now()
	max, bf := getMaxDisplayOrder()
	if bf != Success {
		return bf
	}
	job.DisplayOrder = max + 1
	tx := db.Table("jobs").Create(&job)
	if tx.Error != nil {
		config.Log.Error(tx.Error)
		return ErrSearchDBData
	}
	return Success
}

func getMaxDisplayOrder() (int, *BriefMessage) {
	db := dbs.DBObj.GetGoRM()
	if db == nil {
		config.Log.Error(InternalGetBDInstanceErr)
		return 0, ErrDataBase
	}
	max := struct {
		Max int `json:"max"`
	}{}
	tx := db.Table("jobs").Where("is_common=0").Select("MAX(display_order) AS max").Scan(&max)
	if tx.Error != nil {
		config.Log.Error(tx.Error)
		return 0, ErrSearchDBData
	}
	return max.Max, Success
}

func PutJob(job *Jobs) *BriefMessage {
	db := dbs.DBObj.GetGoRM()
	if db == nil {
		config.Log.Error(InternalGetBDInstanceErr)
		return ErrDataBase
	}
	tx := db.Table("jobs").
		Where("id = ?", job.ID).
		Update("name", job.Name).
		Update("port", job.Port).
		Update("is_common", 0).
		// Update("display_order", job.DisplayOrder).
		Update("update_at", time.Now())
	if tx.Error != nil {
		config.Log.Error(tx.Error)
		return ErrUpdateData
	}
	return Success
}

func DeleteJob(jID int64) *BriefMessage {
	db := dbs.DBObj.GetGoRM()
	if db == nil {
		config.Log.Error(InternalGetBDInstanceErr)
		return ErrDataBase
	}
	tx := db.Table("jobs").Where("id=?", jID).Delete(nil)
	if tx.Error != nil {
		config.Log.Error(tx.Error)
		return ErrSearchDBData
	}
	return Success
}

func DoSwap(sInfo *SwapInfo) *BriefMessage {
	if sInfo.Action == "up" {
		upSwap(sInfo)
	} else if sInfo.Action == "down" {
		downSwap(sInfo)
	} else {
		return ErrUnSupport
	}
	return Success
}

func upSwap(sInfo *SwapInfo) *BriefMessage {
	db := dbs.DBObj.GetGoRM()
	if db == nil {
		config.Log.Error(InternalGetBDInstanceErr)
		return ErrDataBase
	}
	var beSwapID int // 被交换的ID
	if sInfo.DisplayOrder > 1 {
		beforeDO := sInfo.DisplayOrder - 1
		jIDList, bf := getJobIdByOrder(beforeDO)
		if bf != Success {
			return bf
		}
		tx := db.Table("jobs").
			Where("id=?", sInfo.ID).
			Update("display_order", beforeDO).
			Update("update_at", time.Now())
		if tx.Error != nil {
			config.Log.Error(tx.Error)
			return ErrUpdateData
		}
		if len(jIDList) != 0 {
			beSwapID = jIDList[len(jIDList)-1].ID
			tx = db.Table("jobs").
				Where("id=?", beSwapID).
				Update("display_order", sInfo.DisplayOrder).
				Update("update_at", time.Now())
			if tx.Error != nil {
				config.Log.Error(tx.Error)
				return ErrUpdateData
			}
		}
	} else {
		tx := db.Table("jobs").
			Where("id=?", sInfo.ID).
			Update("display_order", sInfo.DisplayOrder).
			Update("update_at", time.Now())
		if tx.Error != nil {
			config.Log.Error(tx.Error)
			return ErrUpdateData
		}
	}
	return Success
}

func downSwap(sInfo *SwapInfo) *BriefMessage {
	db := dbs.DBObj.GetGoRM()
	if db == nil {
		config.Log.Error(InternalGetBDInstanceErr)
		return ErrDataBase
	}
	var beSwapID int // 被交换的ID
	afterDO := sInfo.DisplayOrder + 1
	jIDList, bf := getJobIdByOrder(afterDO)
	if bf != Success {
		return bf
	}
	if len(jIDList) != 0 {
		beSwapID = jIDList[0].ID
		tx := db.Table("jobs").
			Where("id=?", beSwapID).
			Update("display_order", sInfo.DisplayOrder).
			Update("update_at", time.Now())
		if tx.Error != nil {
			config.Log.Error(tx.Error)
			return ErrUpdateData
		}
	} else {
		maxOrderID, bf := getMaxDisplayOrder()
		if bf != Success {
			return bf
		}
		if maxOrderID == sInfo.DisplayOrder {
			afterDO = sInfo.DisplayOrder
		}
	}
	tx := db.Table("jobs").
		Where("id=?", sInfo.ID).
		Update("display_order", afterDO).
		Update("update_at", time.Now())
	if tx.Error != nil {
		config.Log.Error(tx.Error)
		return ErrUpdateData
	}
	return Success
}

func getJobId(name string) ([]OnlyID, *BriefMessage) {
	db := dbs.DBObj.GetGoRM()
	if db == nil {
		config.Log.Error(InternalGetBDInstanceErr)
		return nil, ErrDataBase
	}
	jobIdList := []OnlyID{}
	tx := db.Table("jobs").Where("name like ? and is_common=0", fmt.Sprint("%", name, "%")).Find(&jobIdList)
	if tx.Error != nil {
		config.Log.Error(tx.Error)
		return nil, ErrSearchDBData
	}
	return jobIdList, Success
}

func getJobIdByOrder(orderID int) ([]OnlyID, *BriefMessage) {
	db := dbs.DBObj.GetGoRM()
	if db == nil {
		config.Log.Error(InternalGetBDInstanceErr)
		return nil, ErrDataBase
	}
	jobIdList := []OnlyID{}
	tx := db.Table("jobs").Order("update_at desc").
		Where("display_order = ? and is_common=0", orderID).
		Find(&jobIdList)
	if tx.Error != nil {
		config.Log.Error(tx.Error)
		return nil, ErrSearchDBData
	}
	return jobIdList, Success
}
