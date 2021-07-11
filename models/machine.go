package models

import (
	"fmt"
	"pro_cfg_manager/config"
	"pro_cfg_manager/dbs"
	"pro_cfg_manager/utils"
	"strconv"
	"strings"
	"time"

	"gorm.io/gorm"
)

type Machine struct {
	ID       int       `json:"id" gorm:"column:id"`
	IpAddr   string    `json:"ipaddr" gorm:"column:ipaddr"`
	JobsID   []int     `json:"jobs_id" gorm:"-"`
	Enabled  bool      `json:"enabled" gorm:"column:enabled"`
	UpdateAt time.Time `json:"update_at" gorm:"column:update_at"`
}

type TableJobMachines struct {
	JobID     int `json:"job_id" gorm:"job_id"`
	MachineID int `json:"machine_id" gorm:"column:machine_id"`
}

type ListMachine struct {
	ID        int       `json:"id" gorm:"column:id"`
	Name      string    `json:"name" gorm:"column:name"`
	IpAddr    string    `json:"ipaddr" gorm:"column:ipaddr"`
	JobsIdStr string    `json:"jobs_id_str" gorm:"jobs_id_str"`
	UpdateAt  time.Time `json:"update_at" gorm:"column:update_at"`
	Enabled   bool      `json:"enabled" gorm:"column:enabled"`
}

type ListMachineMerge struct {
	ID        int       `json:"id" gorm:"column:id"`
	Name      string    `json:"name" gorm:"column:name"`
	IpAddr    string    `json:"ipaddr" gorm:"column:ipaddr"`
	JobsId    []int     `json:"jobs_id" gorm:"jobs_id"`
	UpdateAt  time.Time `json:"update_at" gorm:"column:update_at"`
	Enabled   bool      `json:"enabled" gorm:"column:enabled"`
	Health    string    `json:"health" gorm:"-"`
	LastError string    `json:"last_error" gorm:"-"`
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

func GetAllMachine() ([]*Machine, *BriefMessage) {
	db := dbs.DBObj.GetGoRM()
	if db == nil {
		config.Log.Error(InternalGetBDInstanceErr)
		return nil, ErrDataBase
	}
	ms := []*Machine{}
	tx := db.Table("machines").Find(&ms)
	if tx.Error != nil {
		config.Log.Error(tx.Error)
		return nil, ErrSearchDBData
	}
	return ms, Success
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
	FROM machines 
	LEFT JOIN job_machines 
	ON machines.id=job_machines.machine_id 
	LEFT JOIN jobs 
	ON jobs.id=job_machines.job_id `, s)
		like := `'%` + sp.Search + `%'`
		whereLike := fmt.Sprintf(" AND (machines.ipaddr LIKE %s OR jobs.name LIKE %s) ", like, like)
		where := " WHERE (jobs.is_common=0 OR jobs.is_common IS NULL) " +
			" %s " +
			" GROUP BY machines.ipaddr " +
			" ORDER BY machines.update_at desc "
		var likeSql string
		if sp.Search != "" {
			likeSql = sql + fmt.Sprintf(where, whereLike)
		} else {
			likeSql = sql + fmt.Sprintf(where, "")
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
		createSql(`machines.*, GROUP_CONCAT(DISTINCT job_machines.job_id separator ';') jobs_id_str `),
		sp.PageSize,
		(sp.PageNo-1)*sp.PageSize,
	)
	tx2 := db.Raw(likeSql).Scan(&lists)
	if tx2.Error != nil {
		config.Log.Error(tx2.Error)
		return nil, ErrSearchDBData
	}
	listsSend := []*ListMachineMerge{}
	for _, l := range lists {
		ints, err := ConvertStrToIntSlice(l.JobsIdStr)
		if err != nil {
			return nil, ErrConvertDataType
		}
		listsSend = append(listsSend, &ListMachineMerge{
			ID:       l.ID,
			Name:     l.Name,
			IpAddr:   l.IpAddr,
			UpdateAt: l.UpdateAt,
			Enabled:  l.Enabled,
			JobsId:   ints,
		})
	}
	mObj.Check(&listsSend)
	return CalSplitPage(sp, count, listsSend), Success
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
	db := dbs.DBObj.GetGoRM()
	if db == nil {
		config.Log.Error(InternalGetBDInstanceErr)
		return ErrDataBase
	}
	m.UpdateAt = time.Now()
	err := db.Transaction(func(tx *gorm.DB) error {
		err := tx.Table("machines").Create(m).Error
		if err != nil {
			config.Log.Error(err)
			return err
		}
		if len(m.JobsID) == 0 {
			return nil
		}
		jms := []TableJobMachines{}
		for _, jID := range m.JobsID {
			jms = append(jms, TableJobMachines{
				JobID:     jID,
				MachineID: m.ID,
			})
		}
		err = tx.Table("job_machines").Create(&jms).Error
		if err != nil {
			config.Log.Error(err)
			return err
		}
		return nil
	})
	if err == nil {
		return Success
	}
	if strings.Contains(err.Error(), "Duplicate entry") {
		return ErrDataExist
	}
	return ErrCreateDBData
}

func PutMachine(m *Machine) *BriefMessage {
	if utils.CheckIPAddr(m.IpAddr) {
		return ErrIPAddr
	}
	db := dbs.DBObj.GetGoRM()
	if db == nil {
		config.Log.Error(InternalGetBDInstanceErr)
		return ErrDataBase
	}
	err := db.Transaction(func(tx *gorm.DB) error {
		if err := db.Table("machines").
			Where("id = ?", m.ID).
			Update("ipaddr", m.IpAddr).
			Update("update_at", time.Now()).Error; err != nil {
			config.Log.Error(err)
			return err
		}
		if err := db.Table("job_machines").
			Where("machine_id=?", m.ID).Delete(nil).Error; err != nil {
			config.Log.Error(err)
			return err
		}
		if len(m.JobsID) == 0 {
			return nil
		}
		jms := []*TableJobMachines{}
		for _, jID := range m.JobsID {
			jms = append(jms, &TableJobMachines{
				JobID:     jID,
				MachineID: m.ID,
			})
		}
		if err := db.Table("job_machines").Create(&jms).Error; err != nil {
			config.Log.Error(err)
			return err
		}
		return nil
	})
	if err == nil {
		return Success
	}
	return ErrUpdateData
}

func DeleteMachine(mID int) *BriefMessage {
	db := dbs.DBObj.GetGoRM()
	if db == nil {
		config.Log.Error(InternalGetBDInstanceErr)
		return ErrDataBase
	}
	err := db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Table("machines").Where("id=?", mID).Delete(nil).Error; err != nil {
			config.Log.Error(tx.Error)
			return err
		}
		if err := tx.Table("job_machines").Where("machine_id=?", mID).Delete(nil).Error; err != nil {
			config.Log.Error(tx.Error)
			return err
		}
		return nil
	})
	if err != nil {
		return ErrDelData
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
