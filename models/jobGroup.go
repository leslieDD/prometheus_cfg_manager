package models

import (
	"encoding/json"
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
	Enabled  bool      `json:"enabled" gorm:"column:enabled"`
	UpdateAt time.Time `json:"update_at" gorm:"column:update_at"`
	UpdateBy string    `json:"update_by" gorm:"column:update_by"`
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

func PostJobGroup(user *UserSessionInfo, jb *JobGroup) *BriefMessage {
	db := dbs.DBObj.GetGoRM()
	if db == nil {
		config.Log.Error(InternalGetBDInstanceErr)
		return ErrDataBase
	}
	jb.UpdateAt = time.Now()
	jb.UpdateBy = user.Username
	tx := db.Table("job_group").Create(&jb)
	if tx.Error != nil {
		config.Log.Error(tx.Error)
		return ErrCreateDBData
	}
	return Success
}

func PutJobGroup(user *UserSessionInfo, jb *JobGroup) *BriefMessage {
	db := dbs.DBObj.GetGoRM()
	if db == nil {
		config.Log.Error(InternalGetBDInstanceErr)
		return ErrDataBase
	}
	tx := db.Table("job_group").
		Where("jobs_id=? and id=?", jb.JobsID, jb.ID).
		Update("name", jb.Name).
		Update("update_at", time.Now()).
		Update("update_by", user.Username)
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
	ID       int    `json:"id" gorm:"column:id"`
	Position string `json:"position" gorm:"position"`
	IPAddr   string `json:"ipaddr" gorm:"column:ipaddr"`
}

type JobMachineSend struct {
	ID       int         `json:"id" gorm:"column:id"`
	Position *IPPosition `json:"position" gorm:"position"`
	IPAddr   string      `json:"ipaddr" gorm:"column:ipaddr"`
}

func GetJobMachines(jID int64) ([]*JobMachineSend, *BriefMessage) {
	db := dbs.DBObj.GetGoRM()
	if db == nil {
		config.Log.Error(InternalGetBDInstanceErr)
		return nil, ErrDataBase
	}
	jms := []*JobMachine{}
	tx := db.Table("job_machines").
		Select("id, ipaddr, machines.position").
		Joins("LEFT JOIN machines ON job_machines.machine_id=machines.id").
		Where("job_machines.job_id=?", jID).
		Find(&jms)
	jmsSend := []*JobMachineSend{}
	for _, j := range jms {
		ppi := &IPPosition{}
		if err := json.Unmarshal([]byte(j.Position), ppi); err != nil {
			config.Log.Error(err)
			ppi = nil
		}
		jmsSend = append(jmsSend, &JobMachineSend{
			ID:       j.ID,
			IPAddr:   j.IPAddr,
			Position: ppi,
		})
	}
	if tx.Error != nil {
		config.Log.Error(tx.Error)
		return nil, ErrSearchDBData
	}
	return jmsSend, Success
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

func PutJobGroupMachines(jobID int64, gID int64, pools *[]JobGroupMachine) *BriefMessage {
	db := dbs.DBObj.GetGoRM()
	if db == nil {
		config.Log.Error(InternalGetBDInstanceErr)
		return ErrDataBase
	}
	// ids := map[int]struct{}{}
	// for _, p := range *pools {
	// 	if p.ID == 0 {
	// 		continue
	// 	}
	// 	ids[p.ID] = struct{}{}
	// }
	// 从表里面找出当前子组的所有元素，并删除，并把上传的IP写入表中。以下逻辑有点烦，不过不想改
	// gmIDs := []OnlyID{}
	// tx := db.Table("group_machines").Where("job_group_id=?", gID).Find(&gmIDs)
	// if tx.Error != nil {
	// 	config.Log.Error(tx.Error)
	// 	return ErrSearchDBData
	// }
	// if len(gmIDs) != 0 {
	// 	delIDs := []string{}
	// 	for _, i := range gmIDs {
	// 		_, ok := ids[i.ID]
	// 		if !ok {
	// 			delIDs = append(delIDs, fmt.Sprint(i.ID))
	// 		}
	// 	}
	// 	if len(delIDs) != 0 {
	// 		tx = db.Table("group_machines").Where("id in (" + strings.Join(delIDs, ",") + ")").Delete(nil)
	// 		if tx.Error != nil {
	// 			config.Log.Error(tx.Error)
	// 			return ErrDelData
	// 		}
	// 	}
	// }
	subGroupID := []*OnlyID{}
	tx := db.Table("job_group").Where("jobs_id=? ", jobID).Find(&subGroupID)
	if tx.Error != nil {
		config.Log.Error(tx.Error)
		return ErrSearchDBData
	}
	// subIDs := []int{}
	// for _, s := range subGroupID {
	// 	subIDs = append(subIDs, s.ID)
	// }
	ids := []int{}
	for _, p := range *pools {
		ids = append(ids, p.MachinesID)
	}
	// for _, each := range subIDs
	// 删除其它子组中，包含本次提交上来的Machine ID，因为这些ID要写入此次提交的子组中，其它子组不允许拥有。
	for _, s := range subGroupID {
		tx = db.Table("group_machines").Where("job_group_id=? and machines_id in (?)", s.ID, ids).Delete(nil)
		if tx.Error != nil {
			config.Log.Error(tx.Error)
			return ErrDelData
		}
	}
	// 清清此次提示的子组中的内容
	tx = db.Table("group_machines").Where("job_group_id=?", gID).Delete(nil)
	if tx.Error != nil {
		config.Log.Error(tx.Error)
		return ErrDelData
	}
	// 写入新成员
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
	return Success
}

type GroupLabels struct {
	ID         int       `json:"id" gorm:"column:id"`
	JobGroupID int       `json:"job_group_id" gorm:"column:job_group_id"`
	Key        string    `json:"key" gorm:"column:key"`
	Value      string    `json:"value" gorm:"column:value"`
	Enabled    bool      `json:"enabled" gorm:"column:enabled"`
	UpdateAt   time.Time `json:"update_at" gorm:"column:update_at"`
	UpdateBy   string    `json:"update_by" gorm:"column:update_by"`
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

func PostGroupLabels(gID int64, gls *GroupLabels) *BriefMessage {
	db := dbs.DBObj.GetGoRM()
	if db == nil {
		config.Log.Error(InternalGetBDInstanceErr)
		return ErrDataBase
	}
	gls.UpdateAt = time.Now()
	tx := db.Table("group_labels").Create(gls)
	if tx.Error != nil {
		config.Log.Error(tx.Error)
		return ErrCreateDBData
	}
	return Success
}

func PutGroupLabels(gID int64, gls *GroupLabels) *BriefMessage {
	db := dbs.DBObj.GetGoRM()
	if db == nil {
		config.Log.Error(InternalGetBDInstanceErr)
		return ErrDataBase
	}
	gls.UpdateAt = time.Now()
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

type JobGroupID struct {
	JobID   int `form:"job_id"`
	GroupID int `form:"group_id"`
}

type AllIPLablesOfGroup struct {
	IP      string `json:"ip"`
	GroupID int    `form:"group_id"`
}

type Label2 struct {
	Key   string `json:"key" gorm:"column:key"`
	Value string `json:"value" gorm:"column:value"`
}

type IPAndPosition struct {
	IP       string `json:"ipaddr" gorm:"column:ipaddr"`
	Position string `json:"position" gorm:"column:position"`
}

type IPAndPositionSend struct {
	IP       string      `json:"ip" gorm:"column:ip"`
	Position *IPPosition `json:"position" gorm:"column:position"`
}

func GetAllMachinesLabels(jg *JobGroupID) (interface{}, *BriefMessage) {
	db := dbs.DBObj.GetGoRM()
	if db == nil {
		config.Log.Error(InternalGetBDInstanceErr)
		return nil, ErrDataBase
	}
	ips := []IPAndPosition{}
	tx := db.Table("group_machines").
		Select("machines.ipaddr, machines.position").
		Joins("LEFT JOIN machines ON group_machines.machines_id=machines.id ").
		Where("job_group_id=? AND ipaddr IS NOT NULL AND ipaddr<>''", jg.GroupID).
		Find(&ips)
	if tx.Error != nil {
		config.Log.Error(tx.Error)
		return nil, ErrSearchDBData
	}
	pps := []*IPAndPositionSend{}
	for _, i := range ips {
		if i.Position != "" {
			ppi := &IPPosition{}
			if err := json.Unmarshal([]byte(i.Position), ppi); err != nil {
				config.Log.Error(err)
				ppi = nil
			}
			pps = append(pps, &IPAndPositionSend{
				IP:       i.IP,
				Position: ppi,
			})
		}
	}
	labels := []Label2{}
	tx = db.Table("group_labels").
		Select("`key`, `value`").
		Where("job_group_id=?", jg.GroupID).
		Find(&labels)
	if tx.Error != nil {
		config.Log.Error(tx.Error)
		return nil, ErrSearchDBData
	}
	data := map[string]interface{}{
		"ips":    pps,
		"labels": labels,
	}
	return &data, Success
}

func PutJobGroupStatus(user *UserSessionInfo, edi *EnabledInfo) *BriefMessage {
	db := dbs.DBObj.GetGoRM()
	if db == nil {
		config.Log.Error(InternalGetBDInstanceErr)
		return ErrDataBase
	}
	tx := db.Table("job_group").
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

func PutJobGroupLabelsStatus(user *UserSessionInfo, edi *EnabledInfo) *BriefMessage {
	db := dbs.DBObj.GetGoRM()
	if db == nil {
		config.Log.Error(InternalGetBDInstanceErr)
		return ErrDataBase
	}
	tx := db.Table("group_labels").
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
