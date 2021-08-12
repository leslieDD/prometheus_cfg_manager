package models

import (
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"net"
	"os"
	"path/filepath"
	"pro_cfg_manager/config"
	"pro_cfg_manager/dbs"
	"pro_cfg_manager/utils"
	"strings"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
	"gopkg.in/yaml.v3"
	"gorm.io/gorm"
)

type Monitor struct {
	lock     sync.Mutex
	syncTime time.Time
	info     map[string][]*SrvStatus
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

func (m *Monitor) Check(list *[]*ListMachineMerge) {
	m.lock.Lock()
	defer m.lock.Unlock()
	if m.info == nil || time.Since(m.syncTime).Minutes() > 5.0 {
		config.Log.Print("sync prometheus data")
		m.info = map[string][]*SrvStatus{}
		info, err := m.target()
		if err == nil {
			for _, each := range info.Data.ActiveTargets {
				host, _, sErr := net.SplitHostPort(each.DiscoveredLabels["__address__"])
				if sErr != nil {
					config.Log.Error(sErr)
					continue
				}
				m.info[host] = append(m.info[host], &SrvStatus{
					Job:       each.ScrapePool,
					Health:    each.Health,
					LastError: each.LastError,
				})
			}
		}
	}
	for _, each := range *list {
		obj, ok := m.info[each.IpAddr]
		if ok {
			each.MSrvStatus = obj
		} else {
			each.MSrvStatus = []*SrvStatus{
				{
					Job:       "[未知]",
					Health:    "[未知]",
					LastError: "[未从Prometheus服务中找到]",
				},
			}
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
	if err := filepath.Walk(config.Cfg.PrometheusCfg.Conf,
		func(path string, fi os.FileInfo, err error) error {
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

func AllRulesFileList() ([]FileList, *BriefMessage) {
	fl := FileList{
		Label: "规则文件",
	}
	if err := filepath.Walk(config.Cfg.PrometheusCfg.RuleConf,
		func(path string, fi os.FileInfo, err error) error {
			if err != nil {
				config.Log.Error(err)
				return nil
			}
			if fi.IsDir() {
				return nil
			}
			nameLower := strings.ToLower(fi.Name())
			if !strings.HasSuffix(nameLower, MYYaml) {
				config.Log.Warnf("skip file: %s", fi.Name())
				return nil
			}
			name := strings.TrimRight(fi.Name(), MYYaml)
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

func ReadRuleFileContent(child *Child) (string, *BriefMessage) {
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

func CreateTreeNode(user *UserSessionInfo, t *TreeNodeFromCli) *BriefMessage {
	db := dbs.DBObj.GetGoRM()
	if db == nil {
		config.Log.Error(InternalGetBDInstanceErr)
		return ErrDataBase
	}
	writeData := map[string]interface{}{}
	var tx *gorm.DB
	if t.Level == 2 { // table rules_groups
		pass := CheckPriv(user, "noticeManager", "", "add_file")
		if pass != Success {
			return pass
		}
		writeData["name"] = t.Lable
		tx = db.Table("rules_groups")
	} else if t.Level == 3 {
		pass := CheckPriv(user, "noticeManager", "", "add_group")
		if pass != Success {
			return pass
		}
		writeData["name"] = t.Lable
		writeData["rules_groups_id"] = t.Parent
		tx = db.Table("sub_group")
	} else if t.Level == 4 {
		pass := CheckPriv(user, "noticeManager", "", "add_rule")
		if pass != Success {
			return pass
		}
		writeData["sub_group_id"] = t.Parent
		writeData["alert"] = t.Lable
		writeData["for"] = ""
		writeData["expr"] = ""
		writeData["enabled"] = true
		writeData["description"] = ""
		tx = db.Table("monitor_rules")
	} else {
		return ErrUnSupport
	}
	writeData["update_at"] = time.Now()
	writeData["update_by"] = user.Username
	tx.Create(&writeData)
	if tx.Error != nil {
		config.Log.Error(tx.Error)
		return ErrCreateDBData
	}
	return Success
}

func UpdateTreeNode(user *UserSessionInfo, t *TreeNodeFromCli) *BriefMessage {
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
	tx = tx.Update("update_at", time.Now()).Update("update_by", user.Username)
	if tx.Error != nil {
		config.Log.Error(tx.Error)
		return ErrCreateDBData
	}
	return Success
}

func DeleteTreeNode(t *TreeNodeFromCli, skipSelf bool) *BriefMessage {
	db := dbs.DBObj.GetGoRM()
	if db == nil {
		config.Log.Error(InternalGetBDInstanceErr)
		return ErrDataBase
	}
	// if bf := CheckNodeHaveChildren(t); bf != Success {
	// 	return bf
	// }
	tErr := db.Transaction(func(tx *gorm.DB) error {
		switch t.Level {
		case 1:
			if err := tx.Table("rules_groups").Where("1=1").Delete(nil).Error; err != nil {
				config.Log.Error(err)
				return err
			}
			if err := tx.Table("sub_group").Where("1=1").Delete(nil).Error; err != nil {
				config.Log.Error(err)
				return err
			}
			if err := tx.Table("monitor_rules").Where("1=1").Delete(nil).Error; err != nil {
				config.Log.Error(err)
				return err
			}
			if err := tx.Table("annotations").Where("1=1").Delete(nil).Error; err != nil {
				config.Log.Error(err)
				return err
			}
			if err := tx.Table("monitor_labels").Where("1=1").Delete(nil).Error; err != nil {
				config.Log.Error(err)
				return err
			}
		case 2:
			subGroupIDs := []OnlyID{}
			if err := tx.Table("sub_group").Where("rules_groups_id=?", t.ID).Find(&subGroupIDs).Error; err != nil {
				config.Log.Error(err)
				return err
			}
			sgIDs := ConvertOnlyIdToIntSlice(subGroupIDs)
			monitorRulesIDs := []OnlyID{}
			if err := tx.Table("monitor_rules").Where("sub_group_id in (?)", sgIDs).Find(&monitorRulesIDs).Error; err != nil {
				config.Log.Error(err)
				return err
			}
			if !skipSelf {
				if err := tx.Table("rules_groups").Where("id=?", t.ID).Delete(nil).Error; err != nil {
					config.Log.Error(err)
					return err
				}
			}
			if err := tx.Table("sub_group").Where("rules_groups_id=?", t.ID).Delete(nil).Error; err != nil {
				config.Log.Error(err)
				return err
			}
			if err := tx.Table("monitor_rules").Where("sub_group_id in (?)", sgIDs).Delete(nil).Error; err != nil {
				config.Log.Error(err)
				return err
			}
			mIds := ConvertOnlyIdToIntSlice(monitorRulesIDs)
			if err := tx.Table("annotations").Where("monitor_rules_id in (?)", mIds).Delete(nil).Error; err != nil {
				config.Log.Error(err)
				return err
			}
			if err := tx.Table("monitor_labels").Where("monitor_rules_id in (?)", mIds).Delete(nil).Error; err != nil {
				config.Log.Error(err)
				return err
			}
		case 3:
			monitorRulesIDs := []OnlyID{}
			if err := tx.Table("monitor_rules").Where("sub_group_id=?", t.ID).Find(&monitorRulesIDs).Error; err != nil {
				config.Log.Error(err)
				return err
			}
			if !skipSelf {
				if err := tx.Table("sub_group").Where("id=?", t.ID).Delete(nil).Error; err != nil {
					config.Log.Error(err)
					return err
				}
			}
			if err := tx.Table("monitor_rules").Where("sub_group_id=?", t.ID).Delete(nil).Error; err != nil {
				config.Log.Error(err)
				return err
			}
			mIds := ConvertOnlyIdToIntSlice(monitorRulesIDs)
			if err := tx.Table("annotations").Where("monitor_rules_id in (?)", mIds).Delete(nil).Error; err != nil {
				config.Log.Error(err)
				return err
			}
			if err := tx.Table("monitor_labels").Where("monitor_rules_id in (?)", mIds).Delete(nil).Error; err != nil {
				config.Log.Error(err)
				return err
			}
		case 4:
			if err := tx.Table("monitor_rules").Where("id=?", t.ID).Delete(nil).Error; err != nil {
				config.Log.Error(err)
				return err
			}
			if err := tx.Table("annotations").Where("monitor_rules_id=?", t.ID).Delete(nil).Error; err != nil {
				config.Log.Error(err)
				return err
			}
			if err := tx.Table("monitor_labels").Where("monitor_rules_id=?", t.ID).Delete(nil).Error; err != nil {
				config.Log.Error(err)
				return err
			}
		default:
			err := fmt.Errorf("no support level: %d", t.Level)
			config.Log.Error(err.Error())
			return err
		}
		return nil
	})
	if tErr != nil {
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

type TreeNodeStatus struct {
	ID      int  `json:"id" form:"id"`
	Level   int  `json:"level" form:"level"`
	Enabled bool `json:"enabled" form:"enabled"`
}

func PutTreeNodeStatus(user *UserSessionInfo, tns *TreeNodeStatus) *BriefMessage {
	db := dbs.DBObj.GetGoRM()
	if db == nil {
		config.Log.Error(InternalGetBDInstanceErr)
		return ErrDataBase
	}
	err := db.Transaction(func(tx *gorm.DB) error {
		if tns.Level == 1 {
			tx.Table("monitor_rules").Where("1=1").
				Update("enabled", tns.Enabled).
				Update("update_at", time.Now()).
				Update("update_by", user.Username)
		} else if tns.Level == 2 {
			ids := []OnlyID{}
			if err := tx.Table("monitor_rules").
				Raw(fmt.Sprintf("SELECT monitor_rules.id FROM monitor_rules "+
					"LEFT JOIN sub_group "+
					"ON monitor_rules.sub_group_id=sub_group.id "+
					"WHERE sub_group.rules_groups_id=%d ", tns.ID)).
				Scan(&ids).
				Error; err != nil {
				config.Log.Error(err)
				return err
			}
			idsSlice := []int{}
			for _, id := range ids {
				idsSlice = append(idsSlice, id.ID)
			}
			if err := tx.Table("monitor_rules").
				Where("id in (?)", idsSlice).
				Update("enabled", tns.Enabled).
				Update("update_at", time.Now()).
				Update("update_by", user.Username).
				Error; err != nil {
				config.Log.Error(err)
				return err
			}
		} else if tns.Level == 3 {
			if err := tx.Table("monitor_rules").
				Where("sub_group_id=?", tns.ID).
				Update("enabled", tns.Enabled).
				Update("update_at", time.Now()).
				Update("update_by", user.Username).Error; err != nil {
				config.Log.Error(err)
				return err
			}
		} else if tns.Level == 4 {
			if err := tx.Table("monitor_rules").
				Where("id=?", tns.ID).
				Update("update_at", time.Now()).
				Update("update_by", user.Username).
				Update("enabled", tns.Enabled).Error; err != nil {
				config.Log.Error(err)
				return err
			}
		} else {
			err := errors.New("未支持的level")
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

type UploadMonitorRule struct {
	Labels      map[string]string `json:"labels"`
	Annotations map[string]string `json:"annotations"`
	// MonitorRule
	ID          int    `json:"id" gorm:"column:id"`
	Alert       string `json:"alert" gorm:"column:alert"`
	Expr        string `json:"expr" gorm:"column:expr"`
	For         string `json:"for" gorm:"column:for"`
	SubGroupID  int    `json:"sub_group_id" gorm:"column:sub_group_id"`
	Enabled     bool   `json:"enabled" gorm:"column:enabled"`
	Description string `json:"description" gorm:"column:description"`
}

type TongJiUploadRule struct {
	Success int `json:"success"`
	Fail    int `json:"fail"`
}

type ImportRuleErr struct {
	Success                bool   `json:"success"`
	Alert                  string `json:"alert"`
	ImportRuleError        string `json:"import_rule_error"`
	ImportLableError       string `json:"import_label_error"`
	ImportAnnotationsError string `json:"import_annotations_error"`
}

type RespImportResult struct {
	Info   []*ImportRuleErr `json:"info"`
	Result TongJiUploadRule `json:"result"`
}

func PostTreeUploadFileYaml(c *gin.Context, user *UserSessionInfo, gid int64) (*RespImportResult, *BriefMessage) {
	rFile, err := c.FormFile("file")
	if err != nil {
		config.Log.Error(err)
		return nil, ErrUploadFileFormName
	}
	if rFile.Size > 1024*1024*10 {
		config.Log.Errorf("upload size: %d", rFile.Size)
		return nil, ErrTooLarge
	}

	fd, err := rFile.Open()
	if err != nil {
		config.Log.Error(err)
		return nil, ErrFileFormat
	}
	defer fd.Close()
	uploadData := []*UploadMonitorRule{}
	if err := yaml.NewDecoder(fd).Decode(&uploadData); err != nil {
		config.Log.Error(err)
		return nil, ErrParseFileToYaml.Append(err.Error())
	}
	db := dbs.DBObj.GetGoRM()
	if db == nil {
		config.Log.Error(InternalGetBDInstanceErr)
		return nil, ErrDataBase
	}
	errSlice := []*ImportRuleErr{}
	tongji := TongJiUploadRule{}
	for _, r := range uploadData {
		tErr := db.Transaction(func(tx *gorm.DB) error {
			errItem := ImportRuleErr{Alert: r.Alert, Success: true}
			defer func() {
				if !errItem.Success {
					tongji.Fail += 1
					errSlice = append(errSlice, &errItem)
				} else {
					tongji.Success += 1
				}
			}()
			mr := MonitorRule{
				ID:          0,
				Alert:       r.Alert,
				Expr:        r.Expr,
				For:         r.For,
				Enabled:     r.Enabled,
				Description: r.Description,
				SubGroupID:  int(gid),
				UpdateAt:    time.Now(),
				UpdateBy:    user.Username,
			}
			if err := tx.Table("monitor_rules").Create(&mr).Error; err != nil {
				errItem.Success = false
				errItem.ImportRuleError = err.Error()
				config.Log.Error(err)
				return err
			}
			labels := []*Label{}
			for key, value := range r.Labels {
				labels = append(labels, &Label{
					ID:             0,
					Key:            key,
					Value:          value,
					MonitorRulesID: mr.ID,
				})
			}
			if err := tx.Table("monitor_labels").Create(&labels).Error; err != nil {
				config.Log.Error(err)
				errItem.ImportLableError = err.Error()
				return err
			}
			annotations := []*Label{}
			for key, value := range r.Annotations {
				annotations = append(annotations, &Label{
					ID:             0,
					Key:            key,
					Value:          value,
					MonitorRulesID: mr.ID,
				})
			}
			if err := tx.Table("annotations").Create(&annotations).Error; err != nil {
				config.Log.Error(err)
				errItem.ImportAnnotationsError = err.Error()
				return err
			}
			return nil
		})
		if tErr != nil {
			config.Log.Error("transaction error")
		}
	}
	resp := RespImportResult{
		Info:   errSlice,
		Result: tongji,
	}
	return &resp, Success
}
