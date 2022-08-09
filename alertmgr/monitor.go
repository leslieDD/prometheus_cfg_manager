package alertmgr

import (
	"fmt"
	"net/url"
	"pro_cfg_manager/config"
	"pro_cfg_manager/dbs"
	"pro_cfg_manager/models"
	"pro_cfg_manager/utils"
	"strings"
	"sync"
	"time"

	"github.com/sonh/qs"
)

type RuleInfo struct {
	MonitorRule *models.MonitorRule `json:"monitor_rule"`
	Labels      map[string]string   `json:"labels"`
	Annotations map[string]string   `json:"annotations"`
}

type AlertMgr struct {
	lock     sync.Mutex
	RuleInfo []*RuleInfo
}

func NewAlertMgr() *AlertMgr {
	return &AlertMgr{
		lock: sync.Mutex{},
	}
}

// 加载监控规则
func (a *AlertMgr) LoadRule() {
	db := dbs.DBObj.GetGoRM()
	if db == nil {
		config.Log.Error("get db session err")
		return
	}
	rules := []*models.MonitorRule{}
	tx := db.Table("monitor_rules").Where("enabled=1 and api_enabled=1").Find(&rules)
	if tx.Error != nil {
		config.Log.Error(tx.Error)
		return
	}
	rulesID := map[int]*models.MonitorRule{}
	for _, r := range rules {
		rulesID[r.ID] = r
	}
	// 加载Label
	labels := []*models.Label{}
	tx = db.Table("monitor_labels").Find(&labels)
	if tx.Error != nil {
		config.Log.Error(tx.Error)
		return
	}
	labelsID := map[int]map[string]string{}
	for _, l := range labels {
		if _, ok := labelsID[l.MonitorRulesID]; !ok {
			labelsID[l.MonitorRulesID] = map[string]string{}
		}
		labelsID[l.MonitorRulesID][l.Key] = l.Value
	}
	// 加载annotations
	annotations := []*models.Label{}
	tx = db.Table("annotations").Find(&annotations)
	if tx.Error != nil {
		config.Log.Error(tx.Error)
		return
	}
	annotationsID := map[int]map[string]string{}
	for _, a := range annotations {
		if _, ok := annotationsID[a.MonitorRulesID]; !ok {
			annotationsID[a.MonitorRulesID] = map[string]string{}
		}
		annotationsID[a.MonitorRulesID][a.Key] = a.Value
	}
	// 合并数据
	rulesMerges := []*RuleInfo{}
	for _, r := range rulesID {
		/*
			ID          int       `json:"id" gorm:"column:id"`
			Alert       string    `json:"alert" gorm:"column:alert"`
			Expr        string    `json:"expr" gorm:"column:expr"`
			For         string    `json:"for" gorm:"column:for"`
			SubGroupID  int       `json:"sub_group_id" gorm:"column:sub_group_id"`
			Labels      []Label   `json:"labels" gorm:"column:labels"`
			Annotations []Label   `json:"annotations" gorm:"column:annotations"`
			Enabled     bool      `json:"enabled" gorm:"column:enabled"`
			ApiEnabled  bool      `json:"api_enabled" gorm:"column:api_enabled"`
			ApiContent  string    `json:"api_content" gorm:"column:api_content"`
			Description string    `json:"description" gorm:"column:description"`
			UpdateAt    time.Time `json:"update_at" gorm:"column:update_at"`
			UpdateBy    string    `json:"update_by" gorm:"column:update_by"`
		*/
		rulesMerge := RuleInfo{
			MonitorRule: r,
			Labels:      map[string]string{},
			Annotations: map[string]string{},
		}
		if obj, ok := labelsID[r.ID]; ok {
			rulesMerge.Labels = obj
		}
		if obj, ok := annotationsID[r.ID]; ok {
			rulesMerge.Annotations = obj
		}
		rulesMerges = append(rulesMerges, &rulesMerge)
	}
	a.lock.Lock()
	a.RuleInfo = rulesMerges
	a.lock.Unlock()
}

// 执行监控，就也是执行监控规则
func (a *AlertMgr) DoMonitor() {
	for range time.After(1 * time.Minute) {
		a.lock.Lock()
		for _, rule := range a.RuleInfo {
			go a.work(rule)
		}
		a.lock.Unlock()
	}
}

func (a *AlertMgr) work(rule *RuleInfo) {
	urlReq := CreateUrl(rule)
	if urlReq == "" {
		return
	}
	content, err := utils.Get(urlReq)
	if err != nil {
		config.Log.Error(err)
		return
	}
	a.Alert(rule, content)
}

// 报警
func (a *AlertMgr) Alert(rule *RuleInfo, content []byte) {

}

func CreateUrl(rule *RuleInfo) string {
	vals := CreatePostData(rule)
	if vals == nil {
		return ""
	}
	api := strings.ToLower(rule.MonitorRule.ApiContent)
	if !strings.HasPrefix(api, "http://") && !strings.HasPrefix(api, "https://") {
		api = "http://" + api
	}
	if strings.HasSuffix(api, "/") {
		return fmt.Sprintf("%sapi/v1/query?%s", api, vals.Encode())
	} else {
		return fmt.Sprintf("%s/api/v1/query?%s", api, vals.Encode())
	}
}

type QuerySct struct {
	Query string `json:"query" form:"query" qs:"query"`
}

func CreatePostData(rule *RuleInfo) url.Values {
	qr := QuerySct{
		Query: rule.MonitorRule.Expr,
	}
	encoder := qs.NewEncoder()
	values, err := encoder.Values(&qr)
	if err != nil {
		config.Log.Fatal(err)
	}
	return values
}
