package models

import (
	"fmt"
	"net"
	"pro_cfg_manager/config"
	"pro_cfg_manager/dbs"
	"sort"
	"strconv"
	"time"
)

type InstanceTargetsReq struct {
	Instance string `json:"instance" form:"instance"`
}

type JobsInfo struct {
	Name      string   `json:"name"`
	AddrCount int      `json:"addr_count"`
	Up        int      `json:"up"`
	Down      int      `json:"down"`
	Port      string   `json:"port"`
	Addrs     []string `json:"addrs"`
	Result    string   `json:"result"`
}

type InstanceTargetsResp struct {
	Jobs []*JobsInfo `json:"jobs"`
	Data *TargetInfo `json:"data"`
}

type InstanceTargetsUpdate struct {
	Type string      `json:"type"`
	Data []*JobsInfo `json:"data"`
}

func GetInstanceTargets(itr *InstanceTargetsReq) (*InstanceTargetsResp, *BriefMessage) {
	targets, err := Target(itr.Instance)
	if err != nil {
		return nil, ErrMonitorApi
	}
	resp := InstanceTargetsResp{
		Data: targets,
	}
	if targets.Status != "success" {
		return &resp, Success
	}
	activeTargetGroup := map[string][]*activeTarget{}
	for _, target := range targets.Data.ActiveTargets {
		_, ok := activeTargetGroup[target.Labels["job"]]
		if !ok {
			activeTargetGroup[target.Labels["job"]] = []*activeTarget{}
		}
		activeTargetGroup[target.Labels["job"]] = append(activeTargetGroup[target.Labels["job"]], target)
	}
	keySort := []string{}
	jinfo := map[string]*JobsInfo{}
	for name, targets := range activeTargetGroup {
		for _, t := range targets {
			host, port, err := net.SplitHostPort(t.DiscoveredLabels["__address__"])
			if err != nil {
				port = "0"
				host = t.DiscoveredLabels["__address__"]
			}
			key := fmt.Sprintf("%s %s", name, port)
			if _, ok := jinfo[key]; !ok {
				keySort = append(keySort, key)
				jinfo[key] = &JobsInfo{
					Name:  name,
					Addrs: []string{},
					Port:  port,
				}
			}
			jinfo[key].Addrs = append(jinfo[key].Addrs, host)
			jinfo[key].AddrCount += 1
			if t.Health == "up" {
				jinfo[key].Up += 1
			} else {
				jinfo[key].Down += 1
			}
		}
	}
	sort.Strings(keySort)
	jobs := make([]*JobsInfo, 0, len(keySort))
	for _, k := range keySort {
		jobs = append(jobs, jinfo[k])
	}
	for i := 1; i <= 100; i++ {
		jobs = append(jobs, &JobsInfo{Name: fmt.Sprint(i), Addrs: []string{}})
	}
	resp.Jobs = jobs
	return &resp, Success
}

func PutInstanceTargets(user *UserSessionInfo, data *InstanceTargetsUpdate) (map[string]string, *BriefMessage) {
	/*
		<el-radio label="group_name_only">只导入组名</el-radio>
		<el-radio label="all_ip_only">只导入IP</el-radio>
		<el-radio label="merge_group_ip">合并组(会导入IP)</el-radio>
		<el-radio label="replace_group_ip">替换组(会导入IP)</el-radio>
	*/
	switch data.Type {
	case "group_name_only":
		return groupNameOnly(user, data)
	case "all_ip_only":
		return allIPOnly(user, data)
	case "merge_group_ip":
		// if result, bf := allIPOnly(user, data); bf != Success {
		// 	return result, bf
		// }
		return mergeGroupIP(user, data)
	case "replace_group_ip":
		// if result, bf := allIPOnly(user, data); bf != Success {
		// 	return result, bf
		// }
		return replaceGroupIP(user, data)
	default:
		return map[string]string{}, Success
	}
}

func groupNameOnly(user *UserSessionInfo, data *InstanceTargetsUpdate) (map[string]string, *BriefMessage) {
	importRst := map[string]string{}
	db := dbs.DBObj.GetGoRM()
	if db == nil {
		config.Log.Error(InternalGetBDInstanceErr)
		return importRst, ErrDataBase
	}
	for _, group := range data.Data {
		port, err := strconv.ParseInt(group.Port, 10, 16)
		if err != nil {
			config.Log.Error(err)
			importRst[group.Name] = err.Error()
			continue
		}
		var count int64
		tx := db.Table("jobs").Where("name = ?", group.Name).Count(&count)
		if tx.Error != nil {
			config.Log.Error(err)
			importRst[group.Name] = err.Error()
			continue
		}
		if count != 0 {
			importRst[group.Name] = "group name exist"
			continue
		}
		newJob := &Jobs{
			Name:      group.Name,
			Port:      int(port),
			IsCommon:  false,
			ReLabelID: 1,
			UpdateAt:  time.Now(),
			UpdateBy:  user.Username,
		}
		tx = db.Table("jobs").Create(&newJob)
		if tx.Error != nil {
			config.Log.Error(tx.Error)
			importRst[group.Name] = tx.Error.Error()
			continue
		}
		importRst[group.Name] = "Success"
	}
	return importRst, Success
}

func allIPOnly(user *UserSessionInfo, data *InstanceTargetsUpdate) (map[string]string, *BriefMessage) {
	importRst := map[string]string{}
	db := dbs.DBObj.GetGoRM()
	if db == nil {
		config.Log.Error(InternalGetBDInstanceErr)
		return importRst, ErrDataBase
	}
	for _, group := range data.Data {
		allIPAddrs := map[string]struct{}{}
		allDomainAddrs := map[string]struct{}{}
		sAddr := ""
		for _, addr := range group.Addrs {
			host, _, err := net.SplitHostPort(addr)
			if err == nil {
				sAddr = host
			} else {
				sAddr = addr
			}
			if net.ParseIP(sAddr) != nil {
				allIPAddrs[sAddr] = struct{}{}
			} else {
				allDomainAddrs[sAddr] = struct{}{}
			}
		}
		uploadResult := UploadResult{}
		if len(allIPAddrs) > 0 {
			machines := []*UploadMachine{}
			for m, _ := range allIPAddrs {
				machines = append(machines, &UploadMachine{
					ID:             0,
					IpAddr:         m,
					ImportInPool:   false,
					ImportInJobNum: 0,
					ImportError:    "",
				})
			}
			umi := UploadMachinesInfo{
				Opts:     UploadOpts{IgnoreErr: true},
				JobsID:   []int{},
				Machines: machines,
				TongJi:   UploadResult{},
			}
			result, _ := UploadMachines(user, &umi)
			uploadResult = result.TongJi
		}
		if len(allDomainAddrs) > 0 {
			machines := []*UploadMachine{}
			for m, _ := range allDomainAddrs {
				machines = append(machines, &UploadMachine{
					ID:             0,
					IpAddr:         m,
					ImportInPool:   false,
					ImportInJobNum: 0,
					ImportError:    "",
				})
			}
			umi := UploadMachinesInfo{
				Opts:     UploadOpts{IgnoreErr: true},
				JobsID:   []int{},
				Machines: machines,
				TongJi:   UploadResult{},
			}
			result, _ := UploadDomain(user, &umi)
			uploadResult.Total += result.TongJi.Total
			uploadResult.Success += result.TongJi.Success
			uploadResult.Fail += result.TongJi.Fail
			uploadResult.NoAction += result.TongJi.NoAction
		}
		importRst[group.Name] = fmt.Sprintf("Total: %d, Succ: %d, Fail: %d, NoAction: %d",
			uploadResult.Total,
			uploadResult.Success,
			uploadResult.Fail,
			uploadResult.NoAction,
		)
	}
	return importRst, Success
}

type JobIDAndJobName struct {
	ID   int    `json:"id" gorm:"column:id"`
	Name string `json:"name" gorm:"column:name"`
}

func mergeGroupIP(user *UserSessionInfo, data *InstanceTargetsUpdate) (map[string]string, *BriefMessage) {
	result, bf := groupNameOnly(user, data)
	if bf != Success {
		return result, bf
	}
	importRst := map[string]string{}
	db := dbs.DBObj.GetGoRM()
	if db == nil {
		config.Log.Error(InternalGetBDInstanceErr)
		return importRst, ErrDataBase
	}
	jjs := []*JobIDAndJobName{}
	tx := db.Table("jobs").Find(&jjs)
	if tx.Error != nil {
		return importRst, ErrSearchDBData
	}
	jjsMap := map[string]*JobIDAndJobName{}
	for _, j := range jjs {
		jjsMap[j.Name] = j
	}
	for _, group := range data.Data {
		uploadResult := UploadResult{}
		thisJob, ok := jjsMap[group.Name]
		if ok {
			allIPAddrs := map[string]struct{}{}
			allDomainAddrs := map[string]struct{}{}
			sAddr := ""
			for _, addr := range group.Addrs {
				host, _, err := net.SplitHostPort(addr)
				if err == nil {
					sAddr = host
				} else {
					sAddr = addr
				}
				if net.ParseIP(sAddr) != nil {
					allIPAddrs[sAddr] = struct{}{}
				} else {
					allDomainAddrs[sAddr] = struct{}{}
				}
			}

			if len(allIPAddrs) > 0 {
				machines := []*UploadMachine{}
				for m, _ := range allIPAddrs {
					machines = append(machines, &UploadMachine{
						ID:             0,
						IpAddr:         m,
						ImportInPool:   false,
						ImportInJobNum: 0,
						ImportError:    "",
					})
				}
				umi := UploadMachinesInfo{
					Opts:     UploadOpts{IgnoreErr: true},
					JobsID:   []int{thisJob.ID},
					Machines: machines,
					TongJi:   UploadResult{},
				}
				result, _ := UploadMachines(user, &umi)
				uploadResult = result.TongJi
			}
			if len(allDomainAddrs) > 0 {
				machines := []*UploadMachine{}
				for m, _ := range allDomainAddrs {
					machines = append(machines, &UploadMachine{
						ID:             0,
						IpAddr:         m,
						ImportInPool:   false,
						ImportInJobNum: 0,
						ImportError:    "",
					})
				}
				umi := UploadMachinesInfo{
					Opts:     UploadOpts{IgnoreErr: true},
					JobsID:   []int{thisJob.ID},
					Machines: machines,
					TongJi:   UploadResult{},
				}
				result, _ := UploadDomain(user, &umi)
				uploadResult.Total += result.TongJi.Total
				uploadResult.Success += result.TongJi.Success
				uploadResult.Fail += result.TongJi.Fail
				uploadResult.NoAction += result.TongJi.NoAction
			}
		}
		importRst[group.Name] = fmt.Sprintf("Total: %d, Succ: %d, Fail: %d, NoAction: %d",
			uploadResult.Total,
			uploadResult.Success,
			uploadResult.Fail,
			uploadResult.NoAction,
		)
	}
	return importRst, Success
}

func replaceGroupIP(user *UserSessionInfo, data *InstanceTargetsUpdate) (map[string]string, *BriefMessage) {
	importRst := map[string]string{}
	db := dbs.DBObj.GetGoRM()
	if db == nil {
		config.Log.Error(InternalGetBDInstanceErr)
		return importRst, ErrDataBase
	}
	groupName := []string{}
	for _, g := range data.Data {
		groupName = append(groupName, g.Name)
	}
	if len(groupName) == 0 {
		return importRst, ErrEmptyData
	}
	jjs := []*JobIDAndJobName{}
	tx := db.Table("jobs").Where("name in ?", groupName).Find(&jjs)
	if tx.Error != nil {
		return importRst, ErrSearchDBData
	}
	idNeedDel := []int{}
	for _, j := range jjs {
		idNeedDel = append(idNeedDel, j.ID)
	}
	bf := BatchDeleteJob(idNeedDel)
	if bf != Success {
		return importRst, ErrDelData
	}
	return mergeGroupIP(user, data)
}
