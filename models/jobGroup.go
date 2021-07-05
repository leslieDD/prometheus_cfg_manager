package models

import (
	"fmt"
	"pro_cfg_manager/config"
	"pro_cfg_manager/dbs"
	"strings"
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

type JobMachine struct {
	ID     int    `json:"id" gorm:"column:id"`
	IPAddr string `json:"ipaddr" gorm:"column:ipaddr"`
}

func GetJobMachines(jID int64) ([]*JobMachine, *BriefMessage) {
	db := dbs.DBObj.GetGoRM()
	if db == nil {
		config.Log.Error(InternalGetBDInstanceErr)
		return nil, ErrDataBase
	}
	jms := []*JobMachine{}
	tx := db.Table("machines").
		Select("id, ipaddr").
		Where("JSON_CONTAINS(`job_id`, JSON_ARRAY(?))", jID).
		Find(&jms)
	if tx.Error != nil {
		config.Log.Error(tx.Error)
		return nil, ErrSearchDBData
	}
	return jms, Success
}

type JobGroupMachine struct {
	ID         int    `json:"id" gorm:"column:id"`
	IPAddr     string `json:"ipaddr" gorm:"column:ipaddr"`
	MachinesID int    `json:"machines_id" gorm:"column:machines_id"`
}

func GetJobGroupMachines(gID int64) ([]*JobGroupMachine, *BriefMessage) {
	db := dbs.DBObj.GetGoRM()
	if db == nil {
		config.Log.Error(InternalGetBDInstanceErr)
		return nil, ErrDataBase
	}
	jms := []*JobGroupMachine{}
	tx := db.Table("group_machines").Raw(fmt.Sprintf(`
	SELECT group_machines.id, ipaddr, machines.id as machines_id 
	FROM group_machines 
	LEFT JOIN machines 
	ON group_machines.machines_id = machines.id 
	WHERE group_machines.job_group_id=%d`, gID)).Find(&jms)
	if tx.Error != nil {
		config.Log.Error(tx.Error)
		return nil, ErrSearchDBData
	}
	// for _, jm := range jms {

	// }
	return jms, Success
}

type GroupMachineID struct {
	ID int `json:"id" gorm:"column:id"`
}

func PutJobGroupMachines(gID int64, pools *[]JobGroupMachine) *BriefMessage {
	db := dbs.DBObj.GetGoRM()
	if db == nil {
		config.Log.Error(InternalGetBDInstanceErr)
		return ErrDataBase
	}
	ids := map[int]struct{}{}
	for _, p := range *pools {
		if p.ID == 0 {
			continue
		}
		ids[p.ID] = struct{}{}
	}
	gmIDs := []GroupMachineID{}
	tx := db.Table("group_machines").Where("job_group_id=?", gID).Find(&gmIDs)
	if tx.Error != nil {
		config.Log.Error(tx.Error)
		return ErrSearchDBData
	}
	if len(gmIDs) != 0 {
		delIDs := []string{}
		for _, i := range gmIDs {
			_, ok := ids[i.ID]
			if !ok {
				delIDs = append(delIDs, fmt.Sprint(i.ID))
			}
		}
		if len(delIDs) != 0 {
			tx = db.Table("group_machines").Where("id in (" + strings.Join(delIDs, ",") + ")").Delete(nil)
			if tx.Error != nil {
				config.Log.Error(tx.Error)
				return ErrDelData
			}
		}
	}
	ist := []string{}
	for _, p := range *pools {
		ist = append(ist, fmt.Sprintf(`(%d, %d, %d)`, p.ID, gID, p.MachinesID))
	}
	if len(ist) != 0 {
		values := strings.Join(ist, ",")
		sql := "INSERT INTO group_machines(`id`,`job_group_id`,`machines_id`) VALUES " +
			values +
			" ON DUPLICATE KEY UPDATE `job_group_id`=VALUES(job_group_id),`machines_id`=VALUES(machines_id) "
		tx2 := db.Table("group_machines").Exec(sql)
		if tx2.Error != nil {
			config.Log.Error(tx2.Error)
			return ErrSearchDBData
		}
	}
	// tx := db.Table("group_machines").Where("job_group_id=?", gID).Delete(nil)
	// if tx.Error != nil {
	// 	config.Log.Error(tx.Error)
	// 	return ErrSearchDBData
	// }
	// ist := []string{}
	// for _, p := range *pools {
	// 	ist = append(ist, fmt.Sprintf(`(%d, %d)`, gID, p.MachinesID))
	// }
	// values := strings.Join(ist, ",")
	// sql := "INSERT INTO group_machines(`job_group_id`,`machines_id`) VALUES " + values
	// tx2 := db.Table("group_machines").Exec(sql)
	// if tx2.Error != nil {
	// 	config.Log.Error(tx2.Error)
	// 	return ErrCreateDBData
	// }
	return Success
}
