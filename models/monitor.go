package models

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"pro_cfg_manager/config"
	"pro_cfg_manager/dbs"
	"pro_cfg_manager/utils"
	"strings"
	"sync"
	"time"

	"gorm.io/gorm"
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

type TreeNodeFromCli struct {
	ID     int    `json:"id" gorm:"column:id"`
	Lable  string `json:"label" gorm:"column:label"`
	Level  int    `json:"level" gorm:"column:level"`
	Parent int    `json:"parent" gorm:"parent"`
}

func CreateTreeNode(t *TreeNodeFromCli) *BriefMessage {
	db := dbs.DBObj.GetGoRM()
	if db == nil {
		config.Log.Error(InternalGetBDInstanceErr)
		return ErrDataBase
	}
	writeData := map[string]interface{}{}
	var tx *gorm.DB
	if t.Level == 2 { // table rules_groups
		writeData["name"] = t.Lable
		tx = db.Table("rules_groups")
	} else if t.Level == 3 {
		writeData["name"] = t.Lable
		writeData["rules_groups_id"] = t.Parent
		tx = db.Table("sub_group")
	} else if t.Level == 4 {
		writeData["sub_group_id"] = t.Parent
		writeData["alert"] = t.Lable
		writeData["for"] = ""
		writeData["expr"] = ""
		tx = db.Table("monitor_rules")
	} else {
		return ErrUnSupport
	}
	tx.Create(&writeData)
	if tx.Error != nil {
		config.Log.Error(tx.Error)
		return ErrCreateDBData
	}
	return Success
}

func UpdateTreeNode(t *TreeNodeFromCli) *BriefMessage {
	db := dbs.DBObj.GetGoRM()
	if db == nil {
		config.Log.Error(InternalGetBDInstanceErr)
		return ErrDataBase
	}
	var tx *gorm.DB
	if t.Level == 2 { // table rules_groups
		tx = db.Table("rules_groups").Where("id=?", t.ID).Update("name", t.Lable)
	} else if t.Level == 3 {
		tx = db.Table("sub_group").Where("id=?", t.ID).
			Update("name", t.Lable)
		// Update("rules_groups_id", t.Parent) //??? 是否需要
	} else if t.Level == 4 {
		tx = db.Table("monitor_rules").Where("id=?", t.ID).
			Update("alert", t.Lable)
		// Update("sub_group_id", t.Parent)
	} else {
		return ErrUnSupport
	}
	if tx.Error != nil {
		config.Log.Error(tx.Error)
		return ErrCreateDBData
	}
	return Success
}

func DeleteTreeNode(t *TreeNodeFromCli) *BriefMessage {
	db := dbs.DBObj.GetGoRM()
	if db == nil {
		config.Log.Error(InternalGetBDInstanceErr)
		return ErrDataBase
	}
	if bf := CheckNodeHaveChildren(t); bf != Success {
		return bf
	}
	var tx *gorm.DB
	if t.Level == 2 { // table rules_groups
		tx = db.Table("rules_groups").Where("id=?", t.ID).Delete(nil)
	} else if t.Level == 3 {
		tx = db.Table("sub_group").Where("id=?", t.ID).Delete(nil)
	} else if t.Level == 4 {
		// 删除相关联标签等
		if bf := DeleteLabelsByMID(t.ID); bf != Success {
			return bf
		}
		tx = db.Table("monitor_rules").Where("id=?", t.ID).Delete(nil)
	} else {
		return ErrUnSupport
	}
	if tx.Error != nil {
		config.Log.Error(tx.Error)
		return ErrDelData
	}
	return Success
}

func CheckNodeHaveChildren(t *TreeNodeFromCli) *BriefMessage {
	db := dbs.DBObj.GetGoRM()
	if db == nil {
		config.Log.Error(InternalGetBDInstanceErr)
		return ErrDataBase
	}
	var tx *gorm.DB
	var count int64
	if t.Level == 2 { // table rules_groups
		tx = db.Table("sub_group").Raw(
			fmt.Sprintf("select count(*) as count from sub_group where rules_groups_id=%d", t.ID)).
			Count(&count)
	} else if t.Level == 3 {
		tx = db.Table("monitor_rules").Raw(
			fmt.Sprintf("select count(*) as count from monitor_rules where sub_group_id=%d", t.ID)).
			Count(&count)
	} else if t.Level == 4 {
		// tx = db.Table("monitor_rules").Where("id=?", t.ID).Delete(nil)
		return Success
	} else {
		return ErrUnSupport
	}
	if tx.Error != nil {
		config.Log.Error(tx.Error)
		return ErrCheckDBData
	}
	if count != 0 {
		return ErrHaveDataNoAllowToDel
	}
	return Success
}

func DeleteLabelsByMID(mid int) *BriefMessage {
	db := dbs.DBObj.GetGoRM()
	if db == nil {
		config.Log.Error(InternalGetBDInstanceErr)
		return ErrDataBase
	}
	tx := db.Table("monitor_labels").Where("monitor_rules_id=?", mid).Delete(nil)
	if tx.Error != nil {
		config.Log.Error(tx.Error)
		return ErrDelData
	}
	tx = db.Table("annotations").Where("monitor_rules_id=?", mid).Delete(nil)
	if tx.Error != nil {
		config.Log.Error(tx.Error)
		return ErrDelData
	}
	return Success
}
