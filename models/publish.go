package models

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"pro_cfg_manager/config"
	"pro_cfg_manager/dbs"
	"pro_cfg_manager/utils"
	"strings"
	"sync"
)

const MYExt = ".json"
const MYYaml = ".yml"

var PublishObj *PublishResolve

func NewPublishResolve() *PublishResolve {
	return &PublishResolve{
		lock: sync.Mutex{},
	}
}

func Publish() *BriefMessage {
	return PublishObj.Do(false)
}

type ConfPublish struct {
	Targets []string `json:"targets"`
}
type PublishResolve struct {
	lock sync.Mutex
}

type TargetList struct {
	Targets []string          `json:"targets"`
	Lables  map[string]string `json:"labels"`
}

type JobGroupIPInfo struct {
	JobGroupID int    `json:"job_group_id" gorm:"column:job_group_id"`
	MachinesID int    `json:"machines_id" gorm:"column:machines_id"`
	JobsID     int    `json:"jobs_id" gorm:"column:jobs_id"`
	IPAddr     string `json:"ipaddr" gorm:"column:ipaddr"`
}

type JobGroupLablesInfo struct {
	JobGroupID int    `json:"job_group_id" gorm:"column:job_group_id"`
	JobsID     int    `json:"jobs_id" gorm:"column:jobs_id"`
	Key        string `json:"key" gorm:"column:key"`
	Value      string `json:"value" gorm:"column:value"`
}

func (p *PublishResolve) formatData() (map[string]*[]*TargetList, *BriefMessage) {
	db := dbs.DBObj.GetGoRM()
	if db == nil {
		config.Log.Error(InternalGetBDInstanceErr)
		return nil, ErrDataBase
	}
	jobs, bf := GetAllActiveJobs()
	if bf != Success {
		return nil, bf
	}
	jobsMap := map[int]*Jobs{}
	for _, j := range jobs {
		jobsMap[j.ID] = j
	}
	if len(jobsMap) == 0 {
		config.Log.Warn("no job found")
		return nil, ErrNoJobs
	}
	jobGp, bf := GetJobGroupIPInfo()
	if bf != Success {
		return nil, bf
	}
	jobLb, bf := GetJobGroupLabelsInfo()
	if bf != Success {
		return nil, bf
	}
	allMachines, bf := GetAllMachines()
	if bf != Success {
		return nil, bf
	}
	allJobLabels, bf := GetAllJobGlobalLable()
	if bf != Success {
		return nil, bf
	}
	//          map[分组ID]map[子组ID]
	jobGpAndLb := map[int]map[int]*TargetList{}
	for _, obj := range jobGp {
		job, ok := jobGpAndLb[obj.JobsID]
		if !ok {
			job = map[int]*TargetList{}
			jobGpAndLb[obj.JobsID] = job
		}
		group, ok := job[obj.JobGroupID]
		if !ok {
			group = &TargetList{
				Targets: []string{},
				Lables:  map[string]string{},
			}
			job[obj.JobGroupID] = group
		}
		relJob, ok := jobsMap[obj.JobsID]
		if !ok {
			// config.Log.Errorf("no found real job, id: %d, ipaddr: %s", obj.JobsID, obj.IPAddr)
			continue
		}
		if relJob.Port == 0 {
			group.Targets = append(group.Targets, obj.IPAddr)
		} else {
			group.Targets = append(group.Targets, fmt.Sprintf("%s:%d", obj.IPAddr, relJob.Port))
		}
	}
	for _, obj := range jobLb {
		job, ok := jobGpAndLb[obj.JobsID]
		if !ok {
			job = map[int]*TargetList{}
			jobGpAndLb[obj.JobsID] = job
		}
		group, ok := job[obj.JobGroupID]
		if !ok {
			group = &TargetList{
				Targets: []string{},
				Lables:  map[string]string{},
			}
			job[obj.JobGroupID] = group
		}
		group.Lables[obj.Key] = obj.Value
	}
	// 写入公共标签
	for jobID, obj := range jobGpAndLb {
		globalLb, ok := allJobLabels[jobID]
		if !ok {
			continue
		}
		if len(globalLb) == 0 {
			continue
		}
		for _, glb := range globalLb {
			for _, jobGp := range obj {
				jobGp.Lables[glb.Name] = glb.Value
			}
		}
	}
	for _, job := range jobs {
		if !job.IsCommon {
			continue
		}
		allMachinesSlice := make([]string, 0, len(allMachines))
		if job.Port == 0 {
			for _, a := range allMachines {
				allMachinesSlice = append(allMachinesSlice, a.IpAddr)
			}
		} else {
			for _, a := range allMachines {
				allMachinesSlice = append(allMachinesSlice, fmt.Sprintf("%s:%d", a.IpAddr, job.Port))
			}
		}
		jobGpAndLb[job.ID] = map[int]*TargetList{}
		jobGpAndLb[job.ID][0] = &TargetList{
			Targets: allMachinesSlice,
			Lables:  map[string]string{},
		}
	}
	targets := map[string]*[]*TargetList{}
	for _, job := range jobs {
		t, ok := targets[job.Name]
		if !ok {
			t = &[]*TargetList{}
			targets[job.Name] = t
		}
		gpAndLbs, ok := jobGpAndLb[job.ID]
		if ok {
			for _, gpAndLb := range gpAndLbs {
				*t = append(*t, gpAndLb)
			}
		}
	}
	return targets, Success
}

func (p *PublishResolve) syncToPrometheus(data map[string]*[]*TargetList) *BriefMessage {
	if data == nil {
		config.Log.Error("data is nil")
		return ErrDataIsNil
	}
	// 清理文件
	if err := filepath.Walk(config.Cfg.PrometheusCfg.Conf,
		func(path string, fi os.FileInfo, err error) error {
			if err != nil {
				config.Log.Error(err)
				return nil
			}
			if path == config.Cfg.PrometheusCfg.Conf {
				return nil
			}
			return os.RemoveAll(path)
		}); err != nil {
		config.Log.Error(err)
	}
	for name, objData := range data {
		pathName := filepath.Join(config.Cfg.PrometheusCfg.Conf, fmt.Sprintf("%s%s", name, MYExt))
		writeData, err := json.MarshalIndent(&objData, "", "    ")
		if err != nil {
			config.Log.Error(err)
			return ErrDataMarshal
		}
		err = utils.WIoutilByte(pathName, writeData)
		if err != nil {
			config.Log.Error(err)
			return ErrDataWriteDisk
		}
	}
	return Success
}

func (p *PublishResolve) ReloadPrometheus() *BriefMessage {
	code, output, err := utils.ExecCmd("systemctl", "reload", "prometheus")
	if err != nil {
		config.Log.Errorf("err: %v: output: %v", err, output)
		return ErrReloadPrometheus
	}
	if code != 0 {
		config.Log.Errorf("code: %d, output: %v", code, output)
		return ErrReloadPrometheus
	}
	return Success
}

func (p *PublishResolve) Do(alreadyReload bool) *BriefMessage {
	p.lock.Lock()
	defer p.lock.Unlock()

	data, bf := p.formatData()
	if bf != Success {
		return bf
	}
	bf = p.syncToPrometheus(data)
	if bf != Success {
		return bf
	}
	// if !alreadyReload {
	// 	r, bf := CheckByFiled("publish_ips_also_reload_srv", "true")
	// 	if bf != Success {
	// 		return bf
	// 	}
	// 	if r {
	// 		return Reload()
	// 	}
	// }
	return Success
}

func CheckByFiled(optKey string, optValue string) (bool, *BriefMessage) {
	db := dbs.DBObj.GetGoRM()
	if db == nil {
		config.Log.Error(InternalGetBDInstanceErr)
		return false, ErrDataBase
	}
	var count int64
	tx := db.Table("options").
		Where("opt_key=? and opt_value=?", optKey, optValue).
		Count(&count)
	if tx.Error != nil {
		config.Log.Error(tx.Error)
		return false, ErrGetControlField
	}
	if count != 1 {
		return false, Success
	}
	return true, Success
}

type Options struct {
	OptKey   string `json:"opt_key" gorm:"column:opt_key"`
	OptValue string `json:"opt_value" gorm:"column:opt_value"`
}

func CheckByFiledIsTrue(optKey string) bool {
	db := dbs.DBObj.GetGoRM()
	if db == nil {
		config.Log.Error(InternalGetBDInstanceErr)
		return false
	}
	opts := []*Options{}
	tx := db.Table("options").Where("opt_key=?", optKey).Find(&opts)
	if tx.Error != nil {
		config.Log.Error(tx.Error)
		return false
	}
	for _, o := range opts {
		if strings.ToLower(o.OptValue) == "true" {
			return true
		}
	}
	return false
}

func Preview() (string, *BriefMessage) {
	content, err := utils.RIoutil(config.Cfg.PrometheusCfg.MainConf)
	if err != nil {
		config.Log.Error(err)
		return "", ErrReadFile
	}
	return string(content), Success
}

func GetJobGroupIPInfo() ([]*JobGroupIPInfo, *BriefMessage) {
	db := dbs.DBObj.GetGoRM()
	if db == nil {
		config.Log.Error(InternalGetBDInstanceErr)
		return nil, ErrDataBase
	}
	jgIPs := []*JobGroupIPInfo{}
	tx := db.Table("group_machines").
		Raw(
			// " SELECT job_group_id, machines_id, jobs_id, ipaddr " +
			// 	" FROM group_machines " +
			// 	" LEFT JOIN job_group " +
			// 	" ON group_machines.job_group_id=job_group.id " +
			// 	" LEFT JOIN machines " +
			// 	" ON machines.id=group_machines.machines_id " +
			// 	" WHERE job_group.enabled=1 AND machines.enabled=1 ").
			`SELECT job_group_id, machines_id, jobs_id, ipaddr 
			FROM group_machines
			LEFT JOIN job_group
			ON group_machines.job_group_id=job_group.id
			LEFT JOIN machines
			ON machines.id=group_machines.machines_id
			LEFT JOIN job_machines
			ON job_machines.job_id=job_group.jobs_id AND job_machines.machine_id=group_machines.machines_id
			WHERE job_group.enabled=1 AND machines.enabled=1 AND job_machines.blacked=0`).
		Find(&jgIPs)
	if tx.Error != nil {
		config.Log.Error(tx.Error)
		return nil, ErrSearchDBData
	}
	return jgIPs, Success
}

func GetAllMachines() ([]*Machine, *BriefMessage) {
	db := dbs.DBObj.GetGoRM()
	if db == nil {
		config.Log.Error(InternalGetBDInstanceErr)
		return nil, ErrDataBase
	}
	machines := []*Machine{}
	tx := db.Table("machines").Where("enabled=1").Find(&machines)
	if tx.Error != nil {
		config.Log.Error(tx.Error)
		return nil, ErrSearchDBData
	}
	return machines, Success
}

func GetJobGroupLabelsInfo() ([]*JobGroupLablesInfo, *BriefMessage) {
	db := dbs.DBObj.GetGoRM()
	if db == nil {
		config.Log.Error(InternalGetBDInstanceErr)
		return nil, ErrDataBase
	}
	jgLables := []*JobGroupLablesInfo{}
	tx := db.Table("group_labels").
		Raw(
			" SELECT job_group_id, `key`, `value`, jobs_id " +
				" FROM group_labels " +
				" LEFT JOIN job_group " +
				" ON group_labels.job_group_id=job_group.id " +
				" WHERE group_labels.enabled=1 AND job_group.enabled=1 ").
		Find(&jgLables)
	if tx.Error != nil {
		config.Log.Error(tx.Error)
		return nil, ErrSearchDBData
	}
	return jgLables, Success
}

func GetAllActiveJobs() ([]*Jobs, *BriefMessage) {
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
	ids := []int{}
	for _, obj := range jobsCount {
		if obj.Count == 0 {
			continue
		}
		ids = append(ids, obj.ID)
	}
	db := dbs.DBObj.GetGoRM()
	if db == nil {
		config.Log.Error(InternalGetBDInstanceErr)
		return nil, ErrDataBase
	}
	jobs := []*Jobs{}
	tx := db.Table("jobs").Where("enabled=1 and id in (?)", ids).Find(&jobs)
	if tx.Error != nil {
		config.Log.Error(tx.Error)
		return nil, ErrSearchDBData
	}
	return jobs, Success
}

func GetAllActiveDefJobs([]*Jobs, *BriefMessage) {

}

func CreateProConfig() *BriefMessage {

	return Success
}

func ReloadProConfig() *BriefMessage {
	return Success
}

func CreateAndReloadConfig() *BriefMessage {
	return Success
}
