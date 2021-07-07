package models

import (
	"fmt"
	"pro_cfg_manager/config"
	"pro_cfg_manager/dbs"
	"pro_cfg_manager/utils"
	"strconv"
	"strings"
	"time"

	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type Machine struct {
	ID       int            `json:"id" gorm:"column:id"`
	IpAddr   string         `json:"ipaddr" gorm:"column:ipaddr"`
	JobId    datatypes.JSON `json:"job_id" gorm:"column:job_id"`
	Enabled  bool           `json:"enabled" gorm:"column:enabled"`
	UpdateAt time.Time      `json:"update_at" gorm:"column:update_at"`
}

type ListMachine struct {
	Name      string         `json:"name" gorm:"column:name"`
	Port      int            `json:"port" gorm:"column:port"`
	CfgName   string         `json:"cfg_name" gorm:"column:cfg_name"`
	ID        int            `json:"id" gorm:"column:id"`
	IpAddr    string         `json:"ipaddr" gorm:"column:ipaddr"`
	JobId     datatypes.JSON `json:"job_id" gorm:"column:job_id"`
	UpdateAt  time.Time      `json:"update_at" gorm:"column:update_at"`
	Enabled   bool           `json:"enabled" gorm:"column:enabled"`
	Health    string         `json:"health" gorm:"-"`
	LastError string         `json:"last_error" gorm:"-"`
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

func GetMachinesV2(sp *SplitPage) (*ResSplitPage, *BriefMessage) {
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
	createSql := func(s string) string {
		sql := fmt.Sprintf(`SELECT %s 
				FROM machines AS m 
				LEFT JOIN jobs AS j 
				ON j.id = json_extract(m.job_id, '$[0]')`, s)
		// ON JSON_CONTAINS(m.job_id, JSON_array(j.id))
		like := `'%` + sp.Search + `%'`
		var likeSql string
		if sp.Search != "" {
			likeSql = sql + fmt.Sprintf(" WHERE j.is_common=0 AND (m.ipaddr LIKE %s OR j.name LIKE %s ) ORDER BY m.update_at desc", like, like)
		} else {
			likeSql = sql + " WHERE j.is_common=0 ORDER BY m.update_at desc "
		}
		return likeSql
	}
	lists := []*ListMachine{}
	var count int64
	tx := db.Table("machines").Raw(createSql(`count(*) as count`)).Count(&count)
	if tx.Error != nil {
		config.Log.Error(tx.Error)
		return nil, ErrCount
	}
	likeSql := fmt.Sprintf("%s LIMIT %d OFFSET %d",
		createSql(`m.id, m.ipaddr, j.name, j.port, j.cfg_name, m.job_id, m.enabled, m.update_at`),
		sp.PageSize,
		(sp.PageNo-1)*sp.PageSize,
	)
	// config.Log.Warn(likeSql)
	tx2 := db.Raw(likeSql).Scan(&lists)
	if tx2.Error != nil {
		config.Log.Error(tx2.Error)
		return nil, ErrCount
	}
	mObj.Check(&lists)
	return CalSplitPage(sp, count, lists), Success
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
	machines := []Machine{}
	var jsonStr string
	var count int64
	var tx *gorm.DB
	tx = db.Table("machines")
	if sp.Search != "" {
		if sp.Option == isGroup {
			list, bf := getJobId(sp.Search)
			if bf != Success {
				return nil, ErrSearchDBData
			}
			if len(list) == 0 {
				return CalSplitPage(sp, count, machines), Success
			}
			listStr := []string{}
			for _, v := range list {
				listStr = append(listStr, strconv.Itoa(v.ID))
			}
			jsonStr = "JSON_CONTAINS(JSON_ARRAY(" + strings.Join(listStr, ",") + "), job_id)"
			tx = tx.Where(jsonStr)
		} else {
			tx = tx.Where("ipaddr like ?", fmt.Sprint("%", sp.Search, "%"))
		}
	}
	tx = tx.Count(&count)
	if tx.Error != nil {
		config.Log.Error(tx.Error)
		return nil, ErrCount
	}
	tx2 := db.Table("machines")
	if sp.Search != "" {
		if sp.Option == isGroup {
			tx2 = tx2.Where(jsonStr)
		} else {
			tx2 = tx2.Where("ipaddr like ?", fmt.Sprint("%", sp.Search, "%"))
		}
	}
	tx2 = tx2.Order("update_at desc").Offset((sp.PageNo - 1) * sp.PageSize).
		Limit(sp.PageSize).
		Find(&machines)
	if tx2.Error != nil {
		config.Log.Error(tx2.Error)
		return nil, ErrSearchDBData
	}
	return CalSplitPage(sp, count, machines), Success
}

func PostMachine(m *Machine) *BriefMessage {
	if utils.CheckIPAddr(m.IpAddr) {
		return ErrIPAddr
	}
	idList, bf := JsonToIntSlice(m.JobId)
	if bf != Success {
		return bf
	}
	if len(idList) == 0 {
		return ErrJobTypeEmpty
	}
	db := dbs.DBObj.GetGoRM()
	if db == nil {
		config.Log.Error(InternalGetBDInstanceErr)
		return ErrDataBase
	}
	m.UpdateAt = time.Now()
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
	if utils.CheckIPAddr(m.IpAddr) {
		return ErrIPAddr
	}
	idList, bf := JsonToIntSlice(m.JobId)
	if bf != Success {
		return bf
	}
	if len(idList) == 0 {
		return ErrJobTypeEmpty
	}
	db := dbs.DBObj.GetGoRM()
	if db == nil {
		config.Log.Error(InternalGetBDInstanceErr)
		return ErrDataBase
	}
	tx := db.Table("machines").
		Where("id = ?", m.ID).
		Update("ipaddr", m.IpAddr).
		Update("job_id", m.JobId).
		Update("update_at", time.Now())
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

type EnabledInfo struct {
	ID      int  `json:"id" gorm:"column:id" form:"id"`
	Enabled bool `json:"enabled" gorm:"column:enabled" form:"enabled"`
}

func PutMachineStatus(oid *EnabledInfo) *BriefMessage {
	db := dbs.DBObj.GetGoRM()
	if db == nil {
		config.Log.Error(InternalGetBDInstanceErr)
		return ErrDataBase
	}
	tx := db.Table("machines").
		Where("id=?", oid.ID).
		Update("enabled", oid.Enabled).
		Update("update_at", time.Now())
	if tx.Error != nil {
		config.Log.Error(tx.Error)
		return ErrUpdateData
	}
	return Success
}
