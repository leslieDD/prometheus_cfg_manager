package models

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
	for name, targets := range activeTargetGroup {

	}
	resp.Jobs = jobs
	return &resp, Success
}
