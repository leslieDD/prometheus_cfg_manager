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

type JobGroupBoard struct {
	IPCount     int `json:"ip_count" gorm:"column:ip_count"`
	LabelsCount int `json:"labels_count" gorm:"column:labels_count"`
	JobGroup
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
	groups := []*JobGroupBoard{}
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
	// 统计
	cdLabels, bf := CountEachGroupLables()
	if bf != Success {
		return nil, bf
	}
	cdIP, bf := CountEachGroupIP()
	if bf != Success {
		return nil, bf
	}
	cdLabelsMap := map[int]*CountDep{}
	for _, obj := range cdLabels {
		cdLabelsMap[obj.JobGroupID] = obj
	}
	cdIPMap := map[int]*CountDep{}
	for _, obj := range cdIP {
		cdIPMap[obj.JobGroupID] = obj
	}
	for _, g := range groups {
		obj, ok := cdLabelsMap[g.ID]
		if ok {
			g.LabelsCount = obj.Count
		}
		obj, ok = cdIPMap[g.ID]
		if ok {
			g.IPCount = obj.Count
		}
	}
	return CalSplitPage(sp, count, groups), Success
}

type CountDep struct {
	JobGroupID int `json:"job_group_id" gorm:"column:job_group_id"`
	Count      int `json:"count" gorm:"column:count"`
	//
}

func CountEachGroupLables() ([]*CountDep, *BriefMessage) {
	db := dbs.DBObj.GetGoRM()
	if db == nil {
		config.Log.Error(InternalGetBDInstanceErr)
		return nil, ErrDataBase
	}
	cd := []*CountDep{}
	tx := db.Table("group_labels").
		Raw("SELECT job_group_id, COUNT(`key`) AS `count` FROM group_labels  GROUP BY job_group_id ORDER BY job_group_id").
		Scan(&cd)
	if tx.Error != nil {
		config.Log.Error(tx.Error)
		return nil, ErrSearchDBData
	}
	return cd, Success
}

func CountEachGroupIP() ([]*CountDep, *BriefMessage) {
	db := dbs.DBObj.GetGoRM()
	if db == nil {
		config.Log.Error(InternalGetBDInstanceErr)
		return nil, ErrDataBase
	}
	cd := []*CountDep{}
	tx := db.Table("group_machines").
		Raw("SELECT job_group_id, COUNT(`machines_id`) AS `count` FROM group_machines LEFT JOIN machines " +
			" ON group_machines.machines_id=machines.id " +
			" WHERE machines.ipaddr<>'' AND machines.ipaddr IS NOT NULL " +
			" GROUP BY job_group_id ORDER BY job_group_id ").
		Scan(&cd)
	if tx.Error != nil {
		config.Log.Error(tx.Error)
		return nil, ErrSearchDBData
	}
	return cd, Success
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
		return ErrDelData
	}
	if bf := ClearGroupLabels(dInfo.ID); bf != Success {
		return bf
	}
	if bf := ClearGroupIP(dInfo.ID); bf != Success {
		return bf
	}
	return Success
}

func ClearGroupLabels(gID int) *BriefMessage {
	db := dbs.DBObj.GetGoRM()
	if db == nil {
		config.Log.Error(InternalGetBDInstanceErr)
		return ErrDataBase
	}
	tx := db.Table("group_labels").
		Where("job_group_id=?", gID).
		Delete(nil)
	if tx.Error != nil {
		config.Log.Error(tx.Error)
		return ErrDelData
	}
	return Success
}

func ClearGroupIP(gID int) *BriefMessage {
	db := dbs.DBObj.GetGoRM()
	if db == nil {
		config.Log.Error(InternalGetBDInstanceErr)
		return ErrDataBase
	}
	tx := db.Table("group_machines").
		Where("job_group_id=?", gID).
		Delete(nil)
	if tx.Error != nil {
		config.Log.Error(tx.Error)
		return ErrDelData
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
	WHERE group_machines.job_group_id=%d AND machines.ipaddr<>'' AND machines.ipaddr IS NOT NULL `, gID)).Find(&jms)
	if tx.Error != nil {
		config.Log.Error(tx.Error)
		return nil, ErrSearchDBData
	}
	// for _, jm := range jms {

	// }
	return jms, Success
}

// type OnlyID struct {
// 	ID int `json:"id" gorm:"column:id"`
// }

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
	gmIDs := []OnlyID{}
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

type GroupLabels struct {
	ID         int       `json:"id" gorm:"column:id"`
	JobGroupID int       `json:"job_group_id" gorm:"column:job_group_id"`
	Key        string    `json:"key" gorm:"column:key"`
	Value      string    `json:"value" gorm:"column:value"`
	UpdateAt   time.Time `json:"update_at" gorm:"column:update_at"`
}

func GetGroupLabels(gID int64) ([]*GroupLabels, *BriefMessage) {
	db := dbs.DBObj.GetGoRM()
	if db == nil {
		config.Log.Error(InternalGetBDInstanceErr)
		return nil, ErrDataBase
	}
	gls := []*GroupLabels{}
	tx := db.Table("group_labels").Where("job_group_id=?", gID).Find(&gls)
	if tx.Error != nil {
		config.Log.Error(tx.Error)
		return nil, ErrSearchDBData
	}
	return gls, Success
}

func PutGroupLabels(gID int64, gls *GroupLabels) *BriefMessage {
	db := dbs.DBObj.GetGoRM()
	if db == nil {
		config.Log.Error(InternalGetBDInstanceErr)
		return ErrDataBase
	}
	gls.UpdateAt = time.Now()
	if gls.ID == 0 {
		tx := db.Table("group_labels").Create(gls)
		if tx.Error != nil {
			config.Log.Error(tx.Error)
			return ErrCreateDBData
		}
		return Success
	}
	tx := db.Table("group_labels").
		Where("id=?", gls.ID).
		Update("job_group_id", gls.JobGroupID).
		Update("key", gls.Key).
		Update("value", gls.Value)
	if tx.Error != nil {
		config.Log.Error(tx.Error)
		return ErrUpdateData
	}
	return Success
}

func GetJobGroupWithSplitPage(sp *SplitPage) (*ResSplitPage, *BriefMessage) {
	db := dbs.DBObj.GetGoRM()
	if db == nil {
		config.Log.Error(InternalGetBDInstanceErr)
		return nil, ErrDataBase
	}
	var count int64
	var tx *gorm.DB
	tx = db.Table("group_labels")
	like := `'%` + sp.Search + `%'`
	if sp.Search != "" {
		tx = tx.Where(fmt.Sprintf(" `job_group_id`=%d AND ( `key` LIKE %s OR `value`LIKE %s ) ", sp.ID, like, like))
	} else {
		tx = tx.Where("job_group_id=?", sp.ID)
	}
	tx = tx.Count(&count)
	if tx.Error != nil {
		config.Log.Error(tx.Error)
		return nil, ErrSearchDBData
	}
	gls := []*GroupLabels{}
	tx = db.Table("group_labels")
	if sp.Search != "" {
		tx = tx.Where(fmt.Sprintf(" `job_group_id`=%d AND ( `key` LIKE %s OR `value` LIKE %s ) ", sp.ID, like, like))
	} else {
		tx = tx.Where("job_group_id=?", sp.ID)
	}
	tx = tx.Order("update_at desc").
		Offset((sp.PageNo - 1) * sp.PageSize).
		Limit(sp.PageSize).
		Find(&gls)
	if tx.Error != nil {
		config.Log.Error(tx.Error)
		return nil, ErrSearchDBData
	}
	return CalSplitPage(sp, count, gls), Success
}

type DelGroupLables struct {
	GroupID int `form:"gid"`
	ID      int `form:"id"`
}

func DelGroupLabels(dInfo *DelGroupLables) *BriefMessage {
	db := dbs.DBObj.GetGoRM()
	if db == nil {
		config.Log.Error(InternalGetBDInstanceErr)
		return ErrDataBase
	}
	tx := db.Table("group_labels").Where("id=?", dInfo.ID).Delete(nil)
	if tx.Error != nil {
		config.Log.Error(tx.Error)
		return ErrUpdateData
	}
	return Success
}
