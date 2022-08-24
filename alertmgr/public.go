package alertmgr

import (
	"encoding/json"
	"net/url"
	"pro_cfg_manager/config"
	"pro_cfg_manager/utils"
	"strings"
	"time"

	"github.com/sonh/qs"
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