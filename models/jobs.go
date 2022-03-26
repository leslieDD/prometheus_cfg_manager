package models

import (
	"fmt"
	"pro_cfg_manager/config"
	"pro_cfg_manager/dbs"
	"strings"
	"time"

	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type Jobs struct {
	ID           int       `json:"id" gorm:"column:id"`
	Name         string    `json:"name" gorm:"column:name"`
	Port         int       `json:"port" gorm:"column:port"`
	CfgName      string    `json:"cfg_name" gorm:"column:cfg_name"`
	IsCommon     bool      `json:"is_common" gorm:"column:is_common"`
	DisplayOrder int       `json:"display_order" gorm:"column:display_order"`
	ReLabelID    int       `json:"relabel_id" gorm:"column:relabel_id"`
	Enabled      bool      `json:"enabled" gorm:"column:enabled"`
	UpdateAt     time.Time `json:"update_at" gorm:"column:update_at"`
	UpdateBy     string    `json:"update_by" gorm:"column:update_by"`
	Online       string    `json:"online" gorm:"-"`
	LastError    string    `json:"last_error" gorm:"-"`
}

type JobsWithRelabelNameAndCount struct {
	ReLabelName string `json:"relabel_name" gorm:"column:relabel_name"`
	IPCount     int64  `json:"ip_count" gorm:"column:ip_count"`
	GroupCount  int64  `json:"group_count" gorm:"column:group_count"`
	Jobs
}

type JobsForTmpl struct {
	ReLabelName    string `json:"relabel_name" gorm:"column:relabel_name"`
	ReLabelEnabled bool   `json:"relabel_enabled" gorm:"column:relabel_enabled"`
	Code           string `json:"code" gorm:"column:code"`
	Jobs
}

type JobCount struct {
	Count int64          `json:"count" gorm:"column:count"`
	JobId datatypes.JSON `json:"job_id" gorm:"column:job_id"`
}

type OnlyID struct {
	ID int `json:"id" gorm:"column:id" form:"id"`
}

type OnlyIDAndCount struct {
	ID    int `json:"id" gorm:"column:id" form:"id"`
	Count int `json:"count" gorm:"column:count" form:"count"`
}

type SwapInfo struct {
	ID           int    `json:"id" gorm:"column:id"`
	DisplayOrder int    `json:"display_order" gorm:"column:display_order"`
	Action       string `json:"action" gorm:"column:action"`
}

func GetJobs() (*[]Jobs, *BriefMessage) {
	db := dbs.DBObj.GetGoRM()
	if db == nil {
		config.Log.Error(InternalGetBDInstanceErr)
		return nil, ErrDataBase
	}
	jobs := []Jobs{}
	tx := db.Table("jobs").
		Order("display_order asc").
		Where("is_common=0").Find(&jobs)
	if tx.Error != nil {
		config.Log.Error(tx.Error)
		return nil, ErrSearchDBData
	}
	return &jobs, Success
}

func GetJobsForTmpl() (*[]*JobsForTmpl, *BriefMessage) {
	r, bf := CheckByFiled("publish_at_empty_nocreate_file", "true")
	if bf != Success {
		return nil, bf
	}
	jobsCount := []*OnlyIDAndCount{}
	if r {
		jobsCount, bf = doOptions_3()
		if bf != Success {
			return nil, bf
		}
	}
	db := dbs.DBObj.GetGoRM()
	if db == nil {
		config.Log.Error(InternalGetBDInstanceErr)
		return nil, ErrDataBase
	}
	var where string
	if len(jobsCount) == 0 {
		// where = "jobs.enabled=1 and relabels.enabled=1"
		where = "jobs.enabled=1"
	} else {
		ids := []string{}
		for _, obj := range jobsCount {
			if obj.Count == 0 {
				continue
			}
			ids = append(ids, fmt.Sprint(obj.ID))
		}
		if len(ids) == 0 {
			// return nil, ErrNoVaildData
			where = "jobs.enabled=1"
		} else {
			// where = fmt.Sprintf("jobs.enabled=1 and relabels.enabled=1 and jobs.id in (%s)", strings.Join(ids, ","))
			where = fmt.Sprintf("jobs.enabled=1 and jobs.id in (%s)", strings.Join(ids, ","))
		}
	}
	jobs := []*JobsForTmpl{}
	tx := db.Table("jobs").
		Select("jobs.*, relabels.code, relabels.name as relabel_name, relabels.enabled as relabel_enabled ").
		Joins("LEFT JOIN relabels on jobs.relabel_id=relabels.id ").
		Where(where).
		Order("display_order asc").
		// Where("is_common=0").
		Find(&jobs)
	if tx.Error != nil {
		config.Log.Error(tx.Error)
		return nil, ErrSearchDBData
	}
	for _, job := range jobs {
		if !job.ReLabelEnabled {
			job.Code = ""
		}
	}
	return &jobs, Success
}

func GetTmplFields() (map[string]string, *BriefMessage) {
	db := dbs.DBObj.GetGoRM()
	if db == nil {
		config.Log.Error(InternalGetBDInstanceErr)
		return nil, ErrDataBase
	}
	fields := []*BaseFields{}
	tx := db.Table("tmpl_fields").Where("enabled=1").Find(&fields)
	if tx.Error != nil {
		return nil, ErrSearchDBData
	}
	fieldsMap := map[string]string{}
	for _, f := range fields {
		fieldsMap[f.Key] = f.Value
	}
	return fieldsMap, Success
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
	jobs := []*JobsWithRelabelNameAndCount{}
	sql := `SELECT * FROM (SELECT t.*, COUNT(job_group.name) AS group_count 
	FROM (SELECT jobs.*, relabels.name as relabel_name, COUNT(job_machines.machine_id) AS ip_count FROM jobs 
	LEFT JOIN relabels ON jobs.relabel_id=relabels.id 
	LEFT JOIN job_machines ON job_machines.job_id=jobs.id
	WHERE %s 
	GROUP BY jobs.id
	ORDER BY jobs.id
	) AS t 
	LEFT JOIN job_group ON job_group.jobs_id=t.id
	GROUP BY t.id
	ORDER BY t.id) AS t2 ORDER BY t2.display_order asc `
	if sp.Search != "" {
		sql = fmt.Sprintf(sql,
			fmt.Sprintf(" jobs.name like '%s' and jobs.is_common=0 ", fmt.Sprint("%", sp.Search, "%")))
	} else {
		sql = fmt.Sprintf(sql, " is_common=0 ")
	}
	sql = sql + fmt.Sprintf(" LIMIT %d OFFSET %d ", sp.PageSize, (sp.PageNo-1)*sp.PageSize)
	tx = db.Table("jobs").Raw(sql).Find(&jobs)
	if tx.Error != nil {
		config.Log.Error(tx.Error)
		return nil, ErrSearchDBData
	}
	return CalSplitPage(sp, count, jobs), Success
}

func GetDefJobsSplit(sp *SplitPage) (*ResSplitPage, *BriefMessage) {
	db := dbs.DBObj.GetGoRM()
	if db == nil {
		config.Log.Error(InternalGetBDInstanceErr)
		return nil, ErrDataBase
	}
	var count int64
	var tx *gorm.DB
	tx = db.Table("jobs")
	if sp.Search != "" {
		tx = tx.Where("name like ? and is_common=1", fmt.Sprint("%", sp.Search, "%"))
	} else {
		tx = tx.Where("is_common=1")
	}
	tx = tx.Count(&count)
	if tx.Error != nil {
		config.Log.Error(tx.Error)
		return nil, ErrCount
	}
	jobs := []*JobsWithRelabelNameAndCount{}
	tx = db.Table("jobs").
		Select("jobs.*, relabels.name as relabel_name ").
		Joins("LEFT JOIN relabels on jobs.relabel_id=relabels.id ")
	if sp.Search != "" {
		tx = tx.Where("name like ? and is_common=1", fmt.Sprint("%", sp.Search, "%"))
	} else {
		tx = tx.Where("is_common=1")
	}
	tx = tx.Order("display_order asc").
		Offset((sp.PageNo - 1) * sp.PageSize).
		Limit(sp.PageSize).
		Find(&jobs)
	if tx.Error != nil {
		config.Log.Error(tx.Error)
		return nil, ErrSearchDBData
	}
	allIPCount, bf := getAllMachineCount()
	if bf != Success {
		return nil, bf
	}
	for _, j := range jobs {
		j.IPCount = allIPCount
	}
	return CalSplitPage(sp, count, jobs), Success
}

func getMachineCount() ([]JobCount, *BriefMessage) {
	jcs := []JobCount{}
	db := dbs.DBObj.GetGoRM()
	if db == nil {
		config.Log.Error(InternalGetBDInstanceErr)
		return nil, ErrDataBase
	}
	sql := `SELECT job_id, COUNT(*) AS count FROM machines GROUP BY machines.job_id`
	tx := db.Table("machines").Raw(sql).Scan(&jcs)
	if tx.Error != nil {
		config.Log.Error(tx.Error)
		return nil, ErrSearchDBData
	}
	return jcs, Success
}

func getAllMachineCount() (int64, *BriefMessage) {
	db := dbs.DBObj.GetGoRM()
	if db == nil {
		config.Log.Error(InternalGetBDInstanceErr)
		return 0, ErrDataBase
	}
	var count int64
	tx := db.Table("machines").Where("enabled=1").Count(&count)
	if tx.Error != nil {
		config.Log.Error(tx.Error)
		return 0, ErrSearchDBData
	}
	return count, Success
}

func getMachineCountForJob(jID int) (*int64, *BriefMessage) {
	db := dbs.DBObj.GetGoRM()
	if db == nil {
		config.Log.Error(InternalGetBDInstanceErr)
		return nil, ErrDataBase
	}
	var count int64
	tx := db.Table("group_machines").Where("job_id=?", jID).Count(&count)
	if tx.Error != nil {
		config.Log.Error(tx.Error)
		return nil, ErrCount
	}
	return &count, Success
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

func PostJob(user *UserSessionInfo, job *Jobs) *BriefMessage {
	db := dbs.DBObj.GetGoRM()
	if db == nil {
		config.Log.Error(InternalGetBDInstanceErr)
		return ErrDataBase
	}
	job.UpdateAt = time.Now()
	job.UpdateBy = user.Username
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

func PutJob(user *UserSessionInfo, job *Jobs) *BriefMessage {
	db := dbs.DBObj.GetGoRM()
	if db == nil {
		config.Log.Error(InternalGetBDInstanceErr)
		return ErrDataBase
	}
	tx := db.Table("jobs").
		Where("id = ?", job.ID).
		Update("name", job.Name).
		Update("port", job.Port).
		Update("is_common", job.IsCommon).
		Update("relabel_id", job.ReLabelID).
		// Update("display_order", job.DisplayOrder).
		Update("update_at", time.Now()).
		Update("update_by", user.Username)
	if tx.Error != nil {
		config.Log.Error(tx.Error)
		return ErrUpdateData
	}
	return Success
}

func DeleteJob(jID int64, isCommon bool) *BriefMessage {
	db := dbs.DBObj.GetGoRM()
	if db == nil {
		config.Log.Error(InternalGetBDInstanceErr)
		return ErrDataBase
	}
	db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Table("jobs").Where("id=?", jID).Delete(nil).Error; err != nil {
			return err
		}
		if err := tx.Table("job_machines").Where("job_id=?", jID).Delete(nil).Error; err != nil {
			return err
		}
		groupIDs := []OnlyID{}
		if err := tx.Table("job_group").Where("jobs_id=?", jID).Find(&groupIDs).Error; err != nil {
			return err
		}
		ids := []int{}
		for _, g := range groupIDs {
			ids = append(ids, g.ID)
		}
		if err := tx.Table("group_labels").Where("job_group_id in (?)", ids).Delete(nil).Error; err != nil {
			return err
		}
		if err := tx.Table("group_machines").Where("job_group_id in (?)", ids).Delete(nil).Error; err != nil {
			return err
		}
		return nil
	})
	return Success
}

func DoSwap(user *UserSessionInfo, sInfo *SwapInfo) *BriefMessage {
	if sInfo.Action == "up" {
		upSwap(user, sInfo)
	} else if sInfo.Action == "down" {
		downSwap(user, sInfo)
	} else {
		return ErrUnSupport
	}
	return Success
}

func upSwap(user *UserSessionInfo, sInfo *SwapInfo) *BriefMessage {
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
			Update("update_at", time.Now()).
			Update("update_by", user.Username)
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

func downSwap(user *UserSessionInfo, sInfo *SwapInfo) *BriefMessage {
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
			Update("update_at", time.Now()).Update("update_by", user.Username)
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

var AllowOneObj *onlyAllowOne

func NewOnlyAllowOne() *onlyAllowOne {
	oao := &onlyAllowOne{
		// lock:   sync.Mutex{},
		notice: make(chan *Notice),
	}
	go oao.doWork()
	return oao
}

type Notice struct {
	Result     chan *BriefMessage
	NeedReload bool
}

type onlyAllowOne struct {
	// lock   sync.Mutex
	notice chan *Notice
}

func (o *onlyAllowOne) DoPublishJobs(needReload bool) *BriefMessage {
	n := &Notice{
		Result:     make(chan *BriefMessage),
		NeedReload: needReload,
	}
	select {
	case o.notice <- n:
		return <-n.Result
	default:
		return ErrHaveInstanceRunning
	}
}

func (o *onlyAllowOne) doWork() {
	for {
		select {
		case mess := <-o.notice:
			mess.Result <- DoPublishJobs(mess.NeedReload)
			close(mess.Result)
		}
	}
}

func DoPublishJobs(needReload bool) *BriefMessage {
	if bf := DoTmplBefore(); bf != Success {
		return bf
	}
	b, bf := TmplObj.doTmpl()
	if bf != Success {
		return bf
	}
	if bf := TmplObj.doWrite(b); bf != Success {
		return bf
	}
	if bf := DoTmplAfter(needReload); bf != Success {
		return bf
	}
	return Success
}

type JobIDAndMachinesID struct {
	MachineID int `json:"machine_id" gorm:"column:machine_id"`
	JobID     int `json:"job_id" gorm:"column:job_id"`
}

type JobGroupIP struct {
	JobGroupID int `json:"job_group_id" gorm:"column:job_group_id"`
	MachinesID int `json:"machines_id" gorm:"column:machines_id"`
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

func PutJobsStatus(user *UserSessionInfo, oid *EnabledInfo) *BriefMessage {
	// count, bf := getMachineCountForJob(oid.ID)
	// if bf != Success {
	// 	return bf
	// }
	// if *count != 0 {
	// 	return ErrHaveDataNoAllowToDisabled
	// }
	db := dbs.DBObj.GetGoRM()
	if db == nil {
		config.Log.Error(InternalGetBDInstanceErr)
		return ErrDataBase
	}
	tx := db.Table("jobs").
		Where("id=?", oid.ID).
		Update("enabled", oid.Enabled).
		Update("update_at", time.Now()).
		Update("update_by", user.Username)
	if tx.Error != nil {
		config.Log.Error(tx.Error)
		return ErrUpdateData
	}
	return Success
}

func PutJobDefaultStatus(user *UserSessionInfo, edi *EnabledInfo) *BriefMessage {
	db := dbs.DBObj.GetGoRM()
	if db == nil {
		config.Log.Error(InternalGetBDInstanceErr)
		return ErrDataBase
	}
	tx := db.Table("jobs").
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

type UpdateIPForJob struct {
	JobID       int   `json:"job_id" gorm:"column:job_id"`
	MachinesIDs []int `json:"machines_ids" gorm:"column:machines_ids"`
}

func ChangeIPAddrUseID(ipContents string) (*UpdateIPForJob, *BriefMessage) {
	ipAddrList := strings.Split(strings.TrimSpace(ipContents), ";")
	db := dbs.DBObj.GetGoRM()
	if db == nil {
		config.Log.Error(InternalGetBDInstanceErr)
		return nil, ErrDataBase
	}
	uif := UpdateIPForJob{}
	for _, ipAddrOri := range ipAddrList {
		ipAddr := strings.TrimSpace(ipAddrOri)
		if ipAddr == "" {
			continue
		}
		m := Machine{}
		tx := db.Table("machines").Where("ipaddr", ipAddr).Find(&m)
		if tx.Error != nil {
			config.Log.Error(tx.Error)
			return nil, ErrSearchDBData
		}
		if m.IpAddr != ipAddr {
			config.Log.Error("ip no found: " + ipAddr)
			continue
		}
		uif.MachinesIDs = append(uif.MachinesIDs, m.ID)
	}
	return &uif, Success
}

func PostUpdateJobIPs(user *UserSessionInfo, cInfo *UpdateIPForJob) *BriefMessage {
	db := dbs.DBObj.GetGoRM()
	if db == nil {
		config.Log.Error(InternalGetBDInstanceErr)
		return ErrDataBase
	}
	err := db.Transaction(func(tx *gorm.DB) error {
		if err := db.Table("job_machines").
			Where("job_id=?", cInfo.JobID).
			Delete(nil).Error; err != nil {
			config.Log.Error(err)
			return err
		}
		if len(cInfo.MachinesIDs) == 0 {
			return nil
		}
		tjms := []*TableJobMachines{}
		// 去重用的
		delMutil := map[int]struct{}{}
		for _, m := range cInfo.MachinesIDs {
			if _, ok := delMutil[m]; ok {
				continue
			}
			delMutil[m] = struct{}{}
			tjms = append(tjms, &TableJobMachines{
				JobID:     cInfo.JobID,
				MachineID: m,
			})
		}
		if err := db.Table("job_machines").Create(&tjms).Error; err != nil {
			config.Log.Error(err)
			return err
		}
		// 在group_machines中，清除已经不存在的IP记录
		if err := db.Exec(`DELETE FROM group_machines WHERE group_machines.machines_id NOT IN ? AND group_machines.job_group_id IN (
			SELECT job_group_id FROM group_machines 
			LEFT JOIN job_group
			ON group_machines.job_group_id=job_group.id
			WHERE job_group.jobs_id=?
			)`, cInfo.MachinesIDs, cInfo.JobID).Error; err != nil {
			config.Log.Error(err)
			return err
		}
		if err := db.Table("jobs").Where("id=?", cInfo.JobID).
			Update("update_at", time.Now()).
			Update("update_by", user.Username).Error; err != nil {
			config.Log.Error(err)
			return err
		}
		return nil
	})
	if err != nil {
		return ErrUpdateData
	}
	return Success
}

func GetPrometheusUrl() (string, *BriefMessage) {
	if config.Cfg.PrometheusCfg.OpenAddress == "" {
		return "", ErrNoDefined
	}
	return config.Cfg.PrometheusCfg.OpenAddress, Success
}
