package alertmgr

import (
	"html/template"
	"pro_cfg_manager/models"
	"time"
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
