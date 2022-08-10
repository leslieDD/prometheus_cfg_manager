package alertmgr

import (
	"bytes"
	"encoding/json"
	"fmt"
	"html/template"
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
	MonitorRule     *models.MonitorRule           `json:"monitor_rule"`
	Labels          map[string]string             `json:"labels"`
	Annotations     map[string]string             `json:"annotations"`
	AnnotationsTmpl map[string]*template.Template `json:"annotations_tmpl"`
	StartsAt        time.Time
}

type AlertMgr struct {
	lock      sync.Mutex
	RulesInfo []*RuleInfo
}

var AlertMgrObj *AlertMgr

func NewAlertMgr() *AlertMgr {
	return &AlertMgr{
		lock: sync.Mutex{},
	}
}

func (a *AlertMgr) Run() {
	a.LoadRule()
	go a.DoMonitor()
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
			MonitorRule:     r,
			Labels:          map[string]string{},
			Annotations:     map[string]string{},
			AnnotationsTmpl: nil,
			StartsAt:        time.Time{},
		}
		if obj, ok := labelsID[r.ID]; ok {
			rulesMerge.Labels = obj
		}
		annotationsTmpl := map[string]*template.Template{}
		if obj, ok := annotationsID[r.ID]; ok {
			rulesMerge.Annotations = obj
			for k, v := range obj {
				tmpl, err := template.New(k).Parse(`{{$value := .Val}}` + v)
				if err != nil {
					config.Log.Error(err)
					return
				}
				annotationsTmpl[k] = tmpl
			}
			rulesMerge.AnnotationsTmpl = annotationsTmpl
		}
		rulesMerges = append(rulesMerges, &rulesMerge)
	}
	config.Log.Warnf("load monitor rules: %d", len(rulesMerges))
	a.lock.Lock()
	a.RulesInfo = rulesMerges
	a.lock.Unlock()
}

// 执行监控，就也是执行监控规则
func (a *AlertMgr) DoMonitor() {
	for {
		select {
		case <-time.After(10 * time.Second):
			a.lock.Lock()
			for _, rule := range a.RulesInfo {
				go a.work(rule)
			}
			a.lock.Unlock()
		}
	}
}

func (a *AlertMgr) work(rule *RuleInfo) {
	urlReq := CreateUrl(rule)
	if urlReq == "" {
		return
	}
	contentResp, err := utils.Get(urlReq)
	if err != nil {
		config.Log.Error(err)
		return
	}
	resp := PrometheusResp{}
	if err := json.Unmarshal(contentResp, &resp); err != nil {
		config.Log.Error(err)
		return
	}
	if resp.Status != "success" {
		config.Log.Error(resp.Data)
		return
	}
	if len(resp.Data.Result) == 0 {
		rule.StartsAt = time.Time{}
		return
	}
	a.Alert(rule, &resp)
}

// 报警
func (a *AlertMgr) Alert(rule *RuleInfo, respData *PrometheusResp) {
	values := ConvertValue(respData)
	if len(values) == 0 {
		rule.StartsAt = time.Time{}
		return
	}
	alerts := []*Alert{}
	for _, item := range values {
		annotations := map[string]string{}
		labels := map[string]string{}
		for k, v := range item.Metric {
			labels[k] = v
		}
		for k, v := range rule.Labels {
			labels[k] = v
		}
		for k, obj := range rule.AnnotationsTmpl {
			buf := new(bytes.Buffer)
			if err := obj.Execute(buf, item); err != nil {
				config.Log.Error(err)
				return
			}
			annotations[k] = buf.String()
		}
		// var startsAt, endsAt string
		// if rule.StartsAt.IsZero() {
		// 	rule.StartsAt = time.Now()
		// 	startsAt = rule.StartsAt.UTC().Format("2006-01-02T15:04:05.000Z")
		// 	endsAt = rule.StartsAt.Add(6 * time.Minute).UTC().Format("2006-01-02T15:04:05.000Z")
		// } else {
		// 	startsAt = rule.StartsAt.UTC().Format("2006-01-02T15:04:05.000Z")
		// 	endsAt = time.Now().UTC().Format("2006-01-02T15:04:05.000Z") // 停留5分钟,
		// }
		var startsAt string
		if rule.StartsAt.IsZero() {
			rule.StartsAt = time.Now()
			startsAt = rule.StartsAt.UTC().Format("2006-01-02T15:04:05.000Z")
		} else {
			startsAt = rule.StartsAt.UTC().Format("2006-01-02T15:04:05.000Z")
		}
		alert := &Alert{
			StartsAt: startsAt,
			// EndsAt:       "",
			GeneratorURL: "",
			Labels:       labels,
			Annotations:  annotations,
		}
		alerts = append(alerts, alert)
	}
	a.PostAlert(alerts)
}

func (a *AlertMgr) PostAlert(alerts []*Alert) {
	outData, err := json.MarshalIndent(alerts, "    ", "    ")
	if err != nil {
		config.Log.Error(err)
		return
	}
	config.Log.Println(string(outData))
	dataResp, err := utils.Post(config.Cfg.PrometheusCfg.AlertManager, outData)
	if err != nil {
		config.Log.Error(err)
		return
	}
	config.Log.Print(dataResp)
}

func GetTime() string {
	return time.Now().Format(time.RFC3339)
}

type PrometheusResp struct {
	Status string          `json:"status"`
	Data   *PrometheusData `json:"data"`
}

type PrometheusData struct {
	ResultType string              `json:"resultType"`
	Result     []*PrometheusResult `json:"result"`
}

type PrometheusResult struct {
	Metric map[string]string
	Value  []interface{} // int64, string
}

type Value struct {
	Unix   int64   `json:"unix"` // Unix时间
	Val    float64 `json:"val"`  // 带宽值
	Metric map[string]string
}

func ConvertValue(rst *PrometheusResp) []*Value {
	values := []*Value{}
	if rst == nil {
		return values
	}
	for _, v := range rst.Data.Result {
		val := &Value{
			Unix:   int64(v.Value[0].(float64)),
			Val:    utils.StringToFloat64(v.Value[1].(string)),
			Metric: v.Metric,
		}
		values = append(values, val)
	}
	return values
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
