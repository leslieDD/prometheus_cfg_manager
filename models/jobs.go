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
	ReLabelName string `json:"relabel_name" gorm:"column:relabel_name"`
	Code        string `json:"code" gorm:"column:code"`
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
		Where("is_common=0 and enabled=1").Find(&jobs)
	if tx.Error != nil {
		config.Log.Error(tx.Error)
		return nil, ErrSearchDBData
	}
	return &jobs, Success
}

func GetJobsForTmpl() (*[]JobsForTmpl, *BriefMessage) {
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
		where = "jobs.enabled=1 and relabels.enabled=1"
	} else {
		ids := []string{}
		for _, obj := range jobsCount {
			if obj.Count == 0 {
				continue
			}
			ids = append(ids, fmt.Sprint(obj.ID))
		}
		where = fmt.Sprintf("jobs.enabled=1 and relabels.enabled=1 and jobs.id in (%s)", strings.Join(ids, ","))
	}
	jobs := []JobsForTmpl{}
	tx := db.Table("jobs").
		Select("jobs.*, relabels.code, relabels.name as relabel_name ").
		Joins("LEFT JOIN relabels on jobs.relabel_id=relabels.id ").
		Where(where).
		Order("display_order asc").
		// Where("is_common=0").
		Find(&jobs)
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
	jobs := []*JobsWithRelabelNameAndCount{}
	// tx = db.Table("jobs").
	// 	Select("jobs.*, relabels.name as relabel_name ").
	// 	Joins("LEFT JOIN relabels on jobs.relabel_id=relabels.id ")
	sql := `SELECT * FROM (SELECT t.*, COUNT(job_group.name) AS group_count 
	FROM (SELECT jobs.*, relabels.name as relabel_name, COUNT(machines.ipaddr) AS ip_count FROM jobs 
	LEFT JOIN relabels ON jobs.relabel_id=relabels.id 
	LEFT JOIN machines ON JSON_CONTAINS(machines.job_id, JSON_ARRAY(jobs.id))
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
	jcs, bf := getMachineCount()
	if bf != Success {
		return nil, bf
	}
	idCount := map[int]int64{}
	for _, j := range jcs {
		if is, bf := JsonToIntSlice(j.JobId); bf != Success {
			return nil, ErrDataParse
		} else {
			for _, i := range is {
				if _, ok := idCount[i]; !ok {
					idCount[i] = j.Count
				} else {
					idCount[i] += j.Count
				}
			}
		}
	}
	// for _, job := range jobs {
	// 	if c, ok := idCount[job.ID]; ok {
	// 		job.Count += c
	// 	}
	// }
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

func getMachineCountForJob(jID int) (*int64, *BriefMessage) {
	db := dbs.DBObj.GetGoRM()
	if db == nil {
		config.Log.Error(InternalGetBDInstanceErr)
		return nil, ErrDataBase
	}
	var count int64
	sql := fmt.Sprintf(`SELECT COUNT(*) AS count FROM machines WHERE JSON_CONTAINS(job_id, JSON_ARRAY(%d))`, jID)
	tx := db.Table("machines").Raw(sql).Count(&count)
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
		Update("is_common", job.IsCommon).
		Update("relabel_id", job.ReLabelID).
		// Update("display_order", job.DisplayOrder).
		Update("update_at", time.Now())
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
	var count int64
	sql := fmt.Sprintf(`SELECT count(*) as count FROM machines as m WHERE JSON_CONTAINS(m.job_id, JSON_array(%d))`, jID)
	tx := db.Table("machines").Raw(sql).Count(&count)
	if tx.Error != nil {
		config.Log.Error(tx.Error)
		return ErrSearchDBData
	}
	if count != 0 {
		return ErrGroupNotEmpty
	}
	tx = db.Table("jobs").Where("id=?", jID).Delete(nil)
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

func DoPublishJobs() *BriefMessage {
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
	if bf := DoTmplAfter(); bf != Success {
		return bf
	}
	return Success
}

type JobIDAndMachinesID struct {
	MachinesID int `json:"machines_id" gorm:"column:machines_id"`
	JobsID     int `json:"jobs_id" gorm:"column:jobs_id"`
}

type JobGroupIP struct {
	ID         int       `json:"id" gorm:"column:id"`
	JobGroupID int       `json:"job_group_id" gorm:"column:job_group_id"`
	MachinesID int       `json:"machines_id" gorm:"column:machines_id"`
	UpdateAt   time.Time `json:"update_at" gorm:"column:update_at"`
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

func PutJobsStatus(oid *EnabledInfo) *BriefMessage {
	count, bf := getMachineCountForJob(oid.ID)
	if bf != Success {
		return bf
	}
	if *count != 0 {
		return ErrHaveDataNoAllowToDisabled
	}
	db := dbs.DBObj.GetGoRM()
	if db == nil {
		config.Log.Error(InternalGetBDInstanceErr)
		return ErrDataBase
	}
	tx := db.Table("jobs").
		Where("id=?", oid.ID).
		Update("enabled", oid.Enabled).
		Update("update_at", time.Now())
	if tx.Error != nil {
		config.Log.Error(tx.Error)
		return ErrUpdateData
	}
	return Success
}

func PutJobDefaultStatus(edi *EnabledInfo) *BriefMessage {
	db := dbs.DBObj.GetGoRM()
	if db == nil {
		config.Log.Error(InternalGetBDInstanceErr)
		return ErrDataBase
	}
	tx := db.Table("jobs").
		Where("id=?", edi.ID).
		Update("enabled", edi.Enabled).
		Update("update_at", time.Now())
	if tx.Error != nil {
		config.Log.Error(tx.Error)
		return ErrUpdateData
	}
	return Success
}
