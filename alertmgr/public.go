package alertmgr

import (
	"encoding/json"
	"fmt"
	"net/url"
	"pro_cfg_manager/config"
	"pro_cfg_manager/utils"
	"strings"
	"time"

	"github.com/sonh/qs"
	"github.com/vicanso/go-charts/v2"
	"gonum.org/v1/plot/plotter"
)

func DoReq(reqUrl string) *PrometheusRangeResp {
	contentResp, err := utils.Get(reqUrl)
	// fullTimeData, err := utils.Post(config.Cfg.Prometheus.Api, []byte(formData.Encode()))
	if err != nil {
		config.Log.Error(err, string(contentResp))
		return nil
	}
	resp := PrometheusRangeResp{}
	if err := json.Unmarshal(contentResp, &resp); err != nil {
		config.Log.Error(err)
		return nil
	}
	return &resp
}

func ConvertValueV2(rst *PrometheusRangeResp) []*Value {
	values := []*Value{}
	if rst == nil {
		return values
	}
	for _, v := range rst.Data.Result {
		for _, value := range v.Values {
			val := &Value{
				Unix: int64(value[0].(float64)),
				Val:  utils.StringToFloat64(value[1].(string)),
			}
			values = append(values, val)
		}
	}
	return values
}

func ConvertValueV3(rst *PrometheusRangeResult) plotter.XYs {
	items := make(plotter.XYs, len(rst.Values))
	for n, v := range rst.Values {
		items[n].X = v[0].(float64)
		items[n].Y = utils.StringToFloat64(v[1].(string))
	}
	return items
	// values := []*Value{}
	// if rst == nil {
	// 	return values
	// }
	// for _, v := range rst.Values {
	// 	val := &Value{
	// 		Unix: int64(v[0].(float64)),
	// 		Val:  utils.StringToFloat64(v[1].(string)),
	// 	}
	// 	values = append(values, val)
	// }
	// return values
}

// 最后一个返回值是最大值
func ConvertValueV4(rst *PrometheusRangeResp) ([]string, []string, [][]float64, float64) {
	ipAddrs := []string{}
	yvss := [][]float64{}
	xtitles := []string{}

	maxVal := float64(0)
	if len(rst.Data.Result) == 0 {
		return ipAddrs, xtitles, yvss, maxVal
	}
	var p *PrometheusRangeResult
	maxLen := 0
	// 找出最找的一个组
	for _, r := range rst.Data.Result {
		if len(r.Values) > maxLen {
			maxLen = len(r.Values)
			p = r
		}
		ipAddrs = append(ipAddrs, r.Metric["instance"])
	}
	// 填充数据
	for _, r := range rst.Data.Result {
		if len(r.Values) == maxLen {
			continue
		}
		kMap := map[interface{}][]interface{}{}
		for _, v := range r.Values {
			kMap[v[0]] = v
		}
		var vNew [][]interface{}
		for _, v := range p.Values {
			if self, ok := kMap[v[0]]; ok {
				vNew = append(vNew, self)
			} else {
				vNew = append(vNew, []interface{}{
					v[0],
					fmt.Sprint(charts.GetNullValue()),
				})
			}
		}
		r.Values = vNew
	}
	for _, r := range rst.Data.Result {
		vs := []float64{}
		for _, v := range r.Values {
			vData := utils.StringToFloat64(v[1].(string))
			vs = append(vs, vData)
			if maxVal < vData && vData != charts.GetNullValue() {
				maxVal = vData
			}
		}
		yvss = append(yvss, vs)
	}
	if p == nil {
		p = rst.Data.Result[0]
	}
	for _, v := range p.Values {
		xtitles = append(xtitles, time.Unix(int64(v[0].(float64)), 0).Local().Format("15:04"))
	}
	return ipAddrs, xtitles, yvss, maxVal
}

func CreatePostDataWithTime(sql string, start time.Time, end time.Time, Step string) url.Values {
	qr := QueryRange{
		Query: sql,
		Start: start.Format(time.RFC3339),
		End:   end.Format(time.RFC3339),
		Step:  Step,
	}
	encoder := qs.NewEncoder()
	values, err := encoder.Values(&qr)
	if err != nil {
		config.Log.Fatal(err)
	}
	return values
}

func CreateReqUrl(api string, sql string, start time.Time, end time.Time) string {
	formData := CreatePostDataWithTime(sql, start, end, "5m")
	return api + "?" + formData.Encode()
}

func ConvertTime(t int, unit string) int64 {
	seconds := 0
	switch unit {
	case "1": // 分钟
		seconds = t
	case "2": // 小时
		seconds = t * 60
	case "3": // 天
		seconds = t * 24 * 60
	default: // 分钟
		seconds = t
	}
	return int64(seconds * 60)
}

func ConvertTimeStr(unit string) string {
	switch unit {
	case "1": // 分钟
		return "分钟"
	case "2": // 小时
		return "小时"
	case "3": // 天
		return "天"
	default: // 分钟
		return "分钟"
	}
}

func ConvertApi(api string) string {
	var newApi string
	api = strings.TrimSpace(api)
	if !strings.HasPrefix(api, "http://") && !strings.HasPrefix(api, "https://") {
		newApi = "http://" + api
	} else {
		newApi = api
	}
	if strings.HasSuffix(newApi, "/") {
		newApi = newApi + "api/v1/query_range"
	} else {
		newApi = newApi + "/api/v1/query_range"
	}
	return newApi
}
