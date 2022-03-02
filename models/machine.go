package models

import (
	"encoding/json"
	"errors"
	"fmt"
	"net"
	"pro_cfg_manager/config"
	"pro_cfg_manager/dbs"
	"pro_cfg_manager/utils"
	"strconv"
	"strings"
	"time"

	"github.com/3th1nk/cidr"
	"gorm.io/gorm"
)

type Machine struct {
	ID       int       `json:"id" gorm:"column:id"`
	IpAddr   string    `json:"ipaddr" gorm:"column:ipaddr"`
	JobsID   []int     `json:"jobs_id" gorm:"-"`
	Position string    `json:"position" gorm:"column:position"`
	IDCName  string    `json:"idc_name" gorm:"column:idc_name"`
	LineName string    `json:"line_name" gorm:"column:line_name"`
	Enabled  bool      `json:"enabled" gorm:"column:enabled"`
	UpdateAt time.Time `json:"update_at" gorm:"column:update_at"`
	UpdateBy string    `json:"update_by" gorm:"column:update_by"`
}

type TableJobMachines struct {
	JobID     int `json:"job_id" gorm:"job_id"`
	MachineID int `json:"machine_id" gorm:"column:machine_id"`
}

type ListMachine struct {
	ID        int       `json:"id" gorm:"column:id"`
	Name      string    `json:"name" gorm:"column:name"`
	IpAddr    string    `json:"ipaddr" gorm:"column:ipaddr"`
	Position  string    `json:"position" gorm:"position"`
	IDCName   string    `json:"idc_name" gorm:"column:idc_name"`
	LineName  string    `json:"line_name" gorm:"column:line_name"`
	JobsIdStr string    `json:"jobs_id_str" gorm:"jobs_id_str"`
	UpdateAt  time.Time `json:"update_at" gorm:"column:update_at"`
	Enabled   bool      `json:"enabled" gorm:"column:enabled"`
	UpdateBy  string    `json:"update_by" gorm:"column:update_by"`
}

type SrvStatus struct {
	Job       string `json:"job" gorm:"-"`
	Health    string `json:"health" gorm:"-"`
	LastError string `json:"last_error" gorm:"-"`
}

type ListMachineMerge struct {
	ID         int          `json:"id" gorm:"column:id"`
	Name       string       `json:"name" gorm:"column:name"`
	IpAddr     string       `json:"ipaddr" gorm:"column:ipaddr"`
	JobsId     []int        `json:"jobs_id" gorm:"jobs_id"`
	Position   *IPPosition  `json:"position" gorm:"position"`
	IDCName    string       `json:"idc_name" gorm:"column:idc_name"`
	LineName   string       `json:"line_name" gorm:"column:line_name"`
	UpdateAt   time.Time    `json:"update_at" gorm:"column:update_at"`
	UpdateBy   string       `json:"update_by" gorm:"column:update_by"`
	Enabled    bool         `json:"enabled" gorm:"column:enabled"`
	MSrvStatus []*SrvStatus `json:"msrv_status" gorm:"-"`
}

type BatchImportIPaddrs struct {
	Content string `json:"content"`
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

func BatchDeleteMachine(delIDs []int) *BriefMessage {
	db := dbs.DBObj.GetGoRM()
	if db == nil {
		config.Log.Error(InternalGetBDInstanceErr)
		return ErrDataBase
	}
	err := db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Table("machines").Where("id in (?)", delIDs).Delete(nil).Error; err != nil {
			config.Log.Error(tx.Error)
			return err
		}
		if err := tx.Table("job_machines").Where("machine_id in (?)", delIDs).Delete(nil).Error; err != nil {
			config.Log.Error(tx.Error)
			return err
		}
		if err := tx.Table("group_machines").Where("machines_id in (?)", delIDs).Delete(nil).Error; err != nil {
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
		whereLike := fmt.Sprintf(" AND (machines.ipaddr LIKE %s OR jobs.name LIKE %s OR machines.position LIKE %s) ", like, like, like)
		where := " WHERE (jobs.is_common=0 OR jobs.is_common IS NULL) " +
			" %s " +
			" GROUP BY machines.ipaddr " +
			" ORDER BY machines.enabled desc, machines.update_at desc "
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
	// config.Log.Print(likeSql)
	tx2 := db.Raw(likeSql).Find(&lists)
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
		ppi := &IPPosition{}
		if l.Position != "" {
			if err := json.Unmarshal([]byte(l.Position), ppi); err != nil {
				config.Log.Error(err)
				ppi = nil
			}
		}
		listsSend = append(listsSend, &ListMachineMerge{
			ID:         l.ID,
			Name:       l.Name,
			IpAddr:     l.IpAddr,
			Position:   ppi,
			UpdateAt:   l.UpdateAt,
			Enabled:    l.Enabled,
			JobsId:     ints,
			UpdateBy:   l.UpdateBy,
			MSrvStatus: []*SrvStatus{},
			IDCName:    l.IDCName,
			LineName:   l.LineName,
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

func PostMachine(user *UserSessionInfo, m *Machine) *BriefMessage {
	if utils.CheckIPAddr(m.IpAddr) {
		return ErrIPAddr
	}
	db := dbs.DBObj.GetGoRM()
	if db == nil {
		config.Log.Error(InternalGetBDInstanceErr)
		return ErrDataBase
	}
	m.UpdateAt = time.Now()
	m.Enabled = true
	m.UpdateBy = user.Username
	position := GetIPPosition(m.IpAddr)
	if position != nil {
		m.Position = position.String()
	}
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

func PutMachine(user *UserSessionInfo, m *Machine) *BriefMessage {
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
			Update("update_at", time.Now()).
			Update("update_by", user.Username).Error; err != nil {
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
		if err := tx.Table("group_machines").Where("machines_id=?", mID).Delete(nil).Error; err != nil {
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

func PutMachineStatus(user *UserSessionInfo, oid *EnabledInfo) *BriefMessage {
	db := dbs.DBObj.GetGoRM()
	if db == nil {
		config.Log.Error(InternalGetBDInstanceErr)
		return ErrDataBase
	}
	tx := db.Table("machines").
		Where("id=?", oid.ID).
		Update("enabled", oid.Enabled).
		Update("update_by", user.Username).
		Update("update_at", time.Now())
	if tx.Error != nil {
		config.Log.Error(tx.Error)
		return ErrUpdateData
	}
	return Success
}

type UploadOpts struct {
	IgnoreErr bool `json:"ignore_err" form:"ignore_err"`
}

type UploadMachine struct {
	ID             int    `json:"-" form:"-"`
	IpAddr         string `json:"ipaddr" form:"ipaddr"`
	ImportInPool   bool   `json:"import_in_pool" form:"import_in_pool"`
	ImportInJobNum int    `json:"import_in_job_num" form:"import_in_job_num"`
	ImportError    string `json:"import_error" form:"import_error"`
}

type UploadResult struct {
	Total    int `json:"total"`
	Success  int `json:"success"`
	Fail     int `json:"fail"`
	NoAction int `json:"noaction"`
}
type UploadMachinesInfo struct {
	Opts     UploadOpts       `json:"opts" form:"opts"`
	JobsID   []int            `json:"jobs_id" form:"jobs_id"`
	Machines []*UploadMachine `json:"machines" form:"machines"`
	TongJi   UploadResult     `json:"tongji" form:"-"`
}

func UploadMachines(user *UserSessionInfo, uploadInfo *UploadMachinesInfo) (*UploadMachinesInfo, *BriefMessage) {
	db := dbs.DBObj.GetGoRM()
	if db == nil {
		config.Log.Error(InternalGetBDInstanceErr)
		return uploadInfo, ErrDataBase
	}
	uploadInfo.TongJi.Total = len(uploadInfo.Machines)
	db.Transaction(func(tx *gorm.DB) error {
		for _, ipInfo := range uploadInfo.Machines {
			if net.ParseIP(ipInfo.IpAddr) == nil {
				ipInfo.ImportInPool = false
				err := errors.New("IP地址不能不合法，不能正常解析")
				ipInfo.ImportError = err.Error()
				uploadInfo.TongJi.Fail += 1
				config.Log.Error(err)
				if !uploadInfo.Opts.IgnoreErr {
					return err
				}
				continue
			}
			m := Machine{
				ID:       0,
				IpAddr:   ipInfo.IpAddr,
				Enabled:  true,
				UpdateAt: time.Now(),
				UpdateBy: user.Username,
			}
			position := GetIPPosition(ipInfo.IpAddr)
			if position != nil {
				m.Position = position.String()
			}
			if err := tx.Table("machines").Create(&m).Error; err != nil {
				ipInfo.ImportInPool = false
				ipInfo.ImportError = err.Error()
				uploadInfo.TongJi.Fail += 1
				config.Log.Error(err)
				if !uploadInfo.Opts.IgnoreErr {
					return err
				}
			} else {
				uploadInfo.TongJi.Success += 1
				ipInfo.ImportInPool = true
				ipInfo.ImportError = "成功导入IP池"
				ipInfo.ID = m.ID
			}
		}
		if len(uploadInfo.JobsID) == 0 {
			return nil
		}
		jobMachines := []*TableJobMachines{}
		for _, jID := range uploadInfo.JobsID {
			for _, ipInfo := range uploadInfo.Machines {
				if !ipInfo.ImportInPool {
					continue
				}
				ipInfo.ImportInJobNum += 1
				jobMachines = append(jobMachines, &TableJobMachines{
					JobID:     jID,
					MachineID: ipInfo.ID,
				})
			}
			if err := tx.Table("job_machines").Create(&jobMachines).Error; err != nil {
				config.Log.Error(err)
				if !uploadInfo.Opts.IgnoreErr {
					return err
				}
			}
		}
		return nil
	})
	uploadInfo.TongJi.NoAction = uploadInfo.TongJi.Total - (uploadInfo.TongJi.Success + uploadInfo.TongJi.Fail)
	return uploadInfo, Success
}

func UpdateIPPosition(user *UserSessionInfo) *BriefMessage {
	db := dbs.DBObj.GetGoRM()
	if db == nil {
		config.Log.Error(InternalGetBDInstanceErr)
		return ErrDataBase
	}
	machines := []*Machine{}
	tx := db.Table("machines").Find(&machines)
	if tx.Error != nil {
		config.Log.Error(tx.Error)
		return ErrUpdateData
	}
	for _, m := range machines {
		ipp := GetIPPosition(m.IpAddr)
		if ipp == nil {
			continue
		}
		tx2 := db.Table("machines").Where("id", m.ID).
			Update("position", ipp.String()).
			Update("update_at", time.Now()).
			Update("update_by", user.Username)
		if tx2.Error != nil {
			config.Log.Error(tx2.Error)
			continue
		}
	}
	return Success
}

func BatchImportIPAddrs(user *UserSessionInfo, content *BatchImportIPaddrs) *BriefMessage {
	items := strings.Split(content.Content, ";")
	importIPs := map[string]struct{}{}
	for _, item := range items {
		currIP := strings.TrimSpace(item)
		if strings.Contains(item, "/") {
			c, err := cidr.ParseCIDR(currIP)
			if err != nil {
				config.Log.Error(err)
				continue
			}
			if err := c.ForEachIP(func(ip string) error {
				importIPs[ip] = struct{}{}
				return nil
			}); err != nil {
				config.Log.Error(err.Error())
			}
		} else if strings.Contains(currIP, "~") {
			fields := strings.Split(currIP, "~")
			if len(fields) != 2 {
				config.Log.Errorf("ip pool err: %s", currIP)
				continue
			}
			ps, err := utils.RangeBeginToEnd(strings.TrimSpace(fields[0]), strings.TrimSpace(fields[1]))
			if err != nil {
				config.Log.Error(err)
				continue
			}
			for _, p := range ps {
				importIPs[p] = struct{}{}
			}
		} else {
			ipObj := net.ParseIP(currIP)
			if ipObj == nil {
				config.Log.Errorf("ip err: %s", currIP)
				continue
			}
			importIPs[item] = struct{}{}
		}
	}
	if len(importIPs) == 0 {
		return Success
	}
	umi := UploadMachinesInfo{
		Opts:     UploadOpts{IgnoreErr: false},
		JobsID:   []int{},
		Machines: nil,
		TongJi:   UploadResult{},
	}
	UploadMachines(user, &umi)
	return importIPs, Success
}
