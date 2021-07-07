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

var publish PublishResolve

func init() {
	publish = PublishResolve{
		lock: sync.Mutex{},
	}
}

func Publish() *BriefMessage {
	return publish.Do()
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
	jobGp, bf := GetJobGroupIPInfo()
	if bf != Success {
		return nil, bf
	}
	jobLb, bf := GetJobGroupLabelsInfo()
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
		group.Targets = append(group.Targets, obj.IPAddr)
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
	// 删除无用的文件
	if err := filepath.Walk(config.Cfg.PrometheusCfg.Conf, func(path string, fi os.FileInfo, err error) error {
		if err != nil {
			config.Log.Error(err)
			return nil
		}
		if fi.IsDir() {
			return nil
		}
		name := strings.TrimRight(fi.Name(), MYExt)
		_, ok := data[name]
		if !ok {
			if err := os.RemoveAll(path); err != nil {
				config.Log.Error(err)
			} else {
				config.Log.Warnf("delete file: %s", fi.Name())
			}
		}
		return nil
	}); err != nil {
		config.Log.Error(err)
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

func (p *PublishResolve) Do() *BriefMessage {
	p.lock.Lock()
	defer p.lock.Unlock()

	data, bf := p.formatData()
	if bf != Success {
		return bf
	}
	return p.syncToPrometheus(data)
	// bf = p.syncToPrometheus(data)
	// if bf != Success {
	// 	return bf
	// }
	// return p.ReloadPrometheus()
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
			" SELECT job_group_id, machines_id, jobs_id, ipaddr " +
				" FROM group_machines " +
				" LEFT JOIN job_group " +
				" ON group_machines.job_group_id=job_group.id " +
				" LEFT JOIN machines " +
				" ON machines.id=group_machines.machines_id " +
				" WHERE job_group.enabled=1 AND machines.enabled=1 ").
		Find(&jgIPs)
	if tx.Error != nil {
		config.Log.Error(tx.Error)
		return nil, ErrSearchDBData
	}
	return jgIPs, Success
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
	db := dbs.DBObj.GetGoRM()
	if db == nil {
		config.Log.Error(InternalGetBDInstanceErr)
		return nil, ErrDataBase
	}
	jobs := []*Jobs{}
	tx := db.Table("jobs").Where("enabled=1").Find(&jobs)
	if tx.Error != nil {
		config.Log.Error(tx.Error)
		return nil, ErrSearchDBData
	}
	return jobs, Success
}
