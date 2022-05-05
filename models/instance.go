package models

import (
	"fmt"
	"net"
	"sort"
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
}

type InstanceTargetsResp struct {
	// Jobs map[string][]*activeTarget `json:"jobs"`
	Jobs []*JobsInfo `json:"jobs"`
	Data *TargetInfo `json:"data"`
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
