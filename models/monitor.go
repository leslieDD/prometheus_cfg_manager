package models

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"pro_cfg_manager/config"
	"pro_cfg_manager/utils"
	"strings"
	"sync"
	"time"
)

type Monitor struct {
	lock     sync.Mutex
	syncTime time.Time
	info     map[string]*activeTarget
}

var mObj = Monitor{
	lock:     sync.Mutex{},
	syncTime: time.Now(),
}

type activeTarget struct {
	DiscoveredLabels   map[string]string `json:"discoveredLabels"`
	Labels             map[string]string `json:"labels"`
	ScrapePool         string            `json:"scrapePool"`
	ScrapeUrl          string            `json:"scrapeUrl"`
	GlobalUrl          string            `json:"globalUrl"`
	LastError          string            `json:"lastError"`
	LastScrape         time.Time         `json:"lastScrape"`
	LastScrapeDuration float64           `json:"lastScrapeDuration"`
	Health             string            `json:"health"`
}

type TargetData struct {
	ActiveTargets []*activeTarget `json:"activeTargets"`
}

type TargetInfo struct {
	Status string     `json:"status"`
	Data   TargetData `json:"data"`
}

func (m *Monitor) Check(list *[]*ListMachine) {
	m.lock.Lock()
	defer m.lock.Unlock()
	if m.info == nil || time.Since(m.syncTime).Minutes() > 5.0 {
		config.Log.Print("sync prometheus data")
		if m.info == nil {
			m.info = map[string]*activeTarget{}
		}
		info, err := m.target()
		if err == nil {
			for _, each := range info.Data.ActiveTargets {
				n := fmt.Sprintf("%s_%s", each.Labels["instance"], each.Labels["job"])
				m.info[n] = each
			}
		}
	}

	for _, each := range *list {
		obj, ok := m.info[fmt.Sprintf("%s_%s", each.IpAddr, each.Name)]
		if ok {
			each.Health = obj.Health
			each.LastError = obj.LastError
		} else {
			each.Health = "unknow"
			each.LastError = "信息还未同步或者IP异常"
		}
	}
}

func (m *Monitor) target() (*TargetInfo, error) {
	reqUrl := fmt.Sprintf("http://%s/api/v1/targets?state=active", config.Cfg.PrometheusCfg.Api)
	txtBytes, err := utils.ReqWithHeader("get", reqUrl, auth())
	if err != nil {
		config.Log.Error(err)
		return nil, err
	}
	info := TargetInfo{}
	err = json.Unmarshal(txtBytes, &info)
	if err != nil {
		config.Log.Error(err)
		return nil, err
	}
	return &info, nil
}

func Reload() *BriefMessage {
	reqUrl := fmt.Sprintf("http://%s/-/reload", config.Cfg.PrometheusCfg.Api)
	result, err := utils.ReqWithHeader("post", reqUrl, auth())
	if err != nil {
		config.Log.Error(err)
		return ErrMonitorApi
	}
	rStr := string(result)
	if strings.TrimSpace(rStr) != "" {
		config.Log.Error(rStr)
		return ErrMonitorApi
	}
	return Success
}

func auth() map[string]string {
	headers := map[string]string{}
	if config.Cfg.PrometheusCfg.Username != "" && config.Cfg.PrometheusCfg.Password != "" {
		data := fmt.Sprintf("%s:%s", config.Cfg.PrometheusCfg.Username, config.Cfg.PrometheusCfg.Password)
		encodeData := base64.StdEncoding.EncodeToString([]byte(data))
		headers["Authorization"] = fmt.Sprintf("Basic %s", encodeData)
	}
	return headers
}

type FileList struct {
	Label    string   `json:"label"`
	Children []*Child `json:"children"`
}

type Child struct {
	Label string `json:"label" form:"label"`
	Path  string `json:"path" form:"path"`
}

func AllFileList() ([]FileList, *BriefMessage) {
	fl := FileList{
		Label: "分组文件",
	}
	if err := filepath.Walk(config.Cfg.PrometheusCfg.Conf, func(path string, fi os.FileInfo, err error) error {
		if err != nil {
			config.Log.Error(err)
			return nil
		}
		if fi.IsDir() {
			return nil
		}
		nameLower := strings.ToLower(fi.Name())
		if !strings.HasSuffix(nameLower, MYExt) {
			config.Log.Warnf("skip file: %s", fi.Name())
			return nil
		}
		name := strings.TrimRight(fi.Name(), MYExt)
		c := Child{
			Label: name,
			Path:  path,
		}
		fl.Children = append(fl.Children, &c)
		return nil
	}); err != nil {
		config.Log.Error(err)
		return nil, ErrFileList
	}
	fls := []FileList{
		fl,
	}
	return fls, Success
}

func ReadFileContent(child *Child) (string, *BriefMessage) {
	content, err := utils.RIoutil(child.Path)
	if err != nil {
		config.Log.Errorf("err: %s path: %s", err, child.Path)
		return "", ErrRFileContent
	}
	return string(content), Success
}
