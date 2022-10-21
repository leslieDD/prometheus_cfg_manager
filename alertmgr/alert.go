package alertmgr

import (
	"html/template"
	"pro_cfg_manager/models"
	"sync"
	"time"

	"github.com/mohae/deepcopy"
)

type Alert struct {
	// EndsAt       string            `json:"endsAt"`
	StartsAt     string            `json:"startsAt"`
	GeneratorURL string            `json:"generatorURL"`
	Labels       map[string]string `json:"labels"`
	Annotations  map[string]string `json:"annotations"`
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

type PrometheusRangeResp struct {
	Status string               `json:"status"`
	Data   *PrometheusRangeData `json:"data"`
}
type PrometheusRangeData struct {
	ResultType string                   `json:"resultType"`
	Result     []*PrometheusRangeResult `json:"result"`
}
type PrometheusRangeResult struct {
	Metric map[string]string
	Values [][]interface{} // int64, string
}

type Value struct {
	Unix   int64   `json:"unix"` // Unix时间
	Val    float64 `json:"val"`  // 带宽值
	Metric map[string]string
}

type RuleInfo struct {
	MonitorRule     *models.MonitorRule           `json:"monitor_rule"`
	Labels          map[string]string             `json:"labels"`
	Annotations     map[string]string             `json:"annotations"`
	AnnotationsTmpl map[string]*template.Template `json:"annotations_tmpl"`
	StartsAt        time.Time
}

type CronRule struct {
	ID        int       `json:"id" gorm:"column:id"`
	Name      string    `json:"name" gorm:"column:name"`
	CID       string    `json:"cid" gorm:"column:cid"` // 有可能是多个
	Rule      string    `json:"rule" gorm:"column:rule"`
	Enabled   bool      `json:"enabled" gorm:"column:enabled"`
	ShowTitle bool      `json:"show_title" gorm:"column:show_title"`
	LineTitle string    `json:"line_title" gorm:"column:line_title"` // 生成线的标题
	ApiID     int       `json:"api_id" gorm:"column:api_id"`
	APIName   string    `json:"api_name" gorm:"column:api_name"`
	API       string    `json:"api" gorm:"column:api"`
	Start     time.Time `json:"start" gorm:"-"`
	End       time.Time `json:"end" gorm:"-"`
	NearTime  int       `json:"near_time" gorm:"column:near_time"`
	Unit      string    `json:"unit" gorm:"column:unit"`
	ExecCycle string    `json:"exec_cycle" gorm:"column:exec_cycle"` // 执行周期
	Status    bool      `json:"status" gorm:"column:status"`
	UpdateAt  time.Time `json:"update_at" gorm:"column:update_at"`
	LoginAt   time.Time `json:"login_at" gorm:"column:login_at"`
}

type QueryRange struct {
	Query string `json:"query" form:"query" qs:"query"`
	Start string `json:"start" form:"start" qs:"start"`
	End   string `json:"end" form:"end" qs:"end"`
	Step  string `json:"step" form:"step" qs:"step"`
	// Timeout int    `json:"timeout" form:"timeout" qs:"timeout"`
}

type taskRst struct {
	RunTimes     int `json:"run_times" gorm:"-"`
	SuccessTimes int `json:"success_times" gorm:"-"`
	FailTimes    int `json:"fail_times" gorm:"-"`
}

var taskRunStatusObj *taskRunStatus

type taskRunStatus struct {
	lock   sync.Mutex
	status map[int]*taskRst
}

func (t *taskRunStatus) AddSuccess(id int) {
	t.lock.Lock()
	defer t.lock.Unlock()

	if obj, ok := t.status[id]; ok {
		obj.SuccessTimes += 1
	} else {
		t.status[id] = &taskRst{}
	}
}

func (t *taskRunStatus) AddFail(id int) {
	t.lock.Lock()
	defer t.lock.Unlock()

	if obj, ok := t.status[id]; ok {
		obj.FailTimes += 1
	} else {
		t.status[id] = &taskRst{}
	}
}

func (t *taskRunStatus) AddOne(id int) {
	t.lock.Lock()
	defer t.lock.Unlock()

	if obj, ok := t.status[id]; ok {
		obj.RunTimes += 1
	} else {
		t.status[id] = &taskRst{}
	}
}

func (t *taskRunStatus) Reset(id int) {
	t.lock.Lock()
	defer t.lock.Unlock()

	t.status[id] = &taskRst{}
}

func (t *taskRunStatus) SyncStatus() map[int]*taskRst {
	t.lock.Lock()
	defer t.lock.Unlock()

	return deepcopy.Copy(t.status).(map[int]*taskRst)
}
