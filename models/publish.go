package models

import (
	"encoding/json"
	"fmt"
	"path/filepath"
	"pro_cfg_manager/config"
	"pro_cfg_manager/dbs"
	"pro_cfg_manager/utils"
	"sync"
)

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

func (p *PublishResolve) formatData() (map[string][]*TargetList, *BriefMessage) {
	db := dbs.DBObj.GetGoRM()
	if db == nil {
		config.Log.Error(InternalGetBDInstanceErr)
		return nil, ErrDataBase
	}
	jobs := []Jobs{}
	tx := db.Table("jobs").Find(&jobs)
	if tx.Error != nil {
		config.Log.Error(tx.Error)
		return nil, ErrSearchDBData
	}
	machines := []Machine{}
	tx = db.Table("machines").Find(&machines)
	if tx.Error != nil {
		config.Log.Error(tx.Error)
		return nil, ErrSearchDBData
	}
	mapJobs := map[int]Jobs{}
	for _, j := range jobs {
		mapJobs[j.ID] = j
	}
	allIPsMap := map[string]struct{}{}
	// 目前只支持一个组
	// 如果是多个组，可以新增加一张labels表，再从machine表中关联labels表
	mGrp := map[string]*TargetList{} // machine group
	for _, m := range machines {
		idList, bf := JsonToIntSlice(&m)
		if bf != Success {
			return nil, bf
		}
		for _, jID := range idList {
			if _, ok := mGrp[mapJobs[jID].Name]; !ok {
				mGrp[mapJobs[jID].Name] = &TargetList{
					Targets: []string{},
					Lables:  map[string]string{},
				}
			}
			var ipAndPort string
			if mapJobs[jID].Port == 0 {
				ipAndPort = m.IpAddr
			} else {
				ipAndPort = fmt.Sprintf("%s:%d", m.IpAddr, mapJobs[jID].Port)
			}
			mGrp[mapJobs[jID].Name].Targets = append(mGrp[mapJobs[jID].Name].Targets, ipAndPort)
			allIPsMap[ipAndPort] = struct{}{}
		}
	}
	mGrpSyncPro := map[string][]*TargetList{} // machine group
	for name, obj := range mGrp {
		mGrpSyncPro[name] = []*TargetList{}
		mGrpSyncPro[name] = append(mGrpSyncPro[name], obj)
	}
	allIPs := []string{}
	for key := range allIPsMap {
		allIPs = append(allIPs, key)
	}
	for _, j := range jobs {
		if _, ok := mGrpSyncPro[j.Name]; !ok {
			mGrpSyncPro[j.Name] = []*TargetList{}
		}
		if j.IsCommon {
			targetList := TargetList{
				Targets: allIPs,
				Lables:  map[string]string{},
			}
			mGrpSyncPro[j.Name] = append(mGrpSyncPro[j.Name], &targetList)
		}
	}
	// r, _ := json.MarshalIndent(mGrpSyncPro, "", "    ")
	// fmt.Printf("%v\n", string(r))
	return mGrpSyncPro, Success
}

func (p *PublishResolve) syncToPrometheus(data map[string][]*TargetList) *BriefMessage {
	if data == nil {
		config.Log.Error("data is nil")
		return ErrDataIsNil
	}
	for name, objData := range data {
		pathName := filepath.Join(config.Cfg.PrometheusCfg.Dir, fmt.Sprintf("%s.json", name))
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

func (p *PublishResolve) Do() *BriefMessage {
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
	return p.ReloadPrometheus()
}
