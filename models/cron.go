package models

import (
	"encoding/json"
	"fmt"
	"github.com/sonh/qs"
	"pro_cfg_manager/config"
	"pro_cfg_manager/dbs"
	"pro_cfg_manager/utils"
	"strings"
	"time"

	"gorm.io/gorm"
)

type CronApi struct {
	ID       int       `json:"id" gorm:"column:id"`
	Name     string    `json:"name" gorm:"column:name"`
	API      string    `json:"api" gorm:"column:api"`
	UpdateAt time.Time `json:"update_at" gorm:"column:update_at"`
	UpdateBy string    `json:"update_by" gorm:"column:update_by"`
}

type CronApiPost struct {
	ID   int    `json:"id" gorm:"column:id"`
	Name string `json:"name" gorm:"column:name"`
	API  string `json:"api" gorm:"column:api"`
}

func GetCronApi(sp *SplitPage) (*ResSplitPage, *BriefMessage) {
	if sp.PageSize <= 0 {
		sp.PageSize = 15
	}
	if sp.PageNo <= 0 {
		sp.PageNo = 1
	}
	db := dbs.DBObj.GetGoRM()
	if db == nil {
		config.Log.Error(InternalGetBDInstanceErr)
		return nil, ErrDataBase
	}
	var count int64
	var likeContent string
	var tx *gorm.DB
	tx = db.Table("crontab_api")
	if sp.Search != "" {
		likeContent = fmt.Sprint("%", sp.Search, "%")
		tx = tx.Where("name like ? or api like ?", likeContent, likeContent)
	}
	tx = tx.Count(&count)
	if tx.Error != nil {
		config.Log.Error(tx.Error)
		return nil, ErrCount
	}
	lists := []*CronApi{}
	tx = db.Table("crontab_api")
	if sp.Search != "" {
		tx = tx.Where("name like ? or api like ?", likeContent, likeContent)
	}
	tx = tx.Offset((sp.PageNo - 1) * sp.PageSize).
		Limit(sp.PageSize).
		Find(&lists)
	if tx.Error != nil {
		config.Log.Error(tx.Error)
		return nil, ErrCount
	}
	return CalSplitPage(sp, count, lists), Success
}

func GetAllCronApi() ([]*CronApi, *BriefMessage) {
	db := dbs.DBObj.GetGoRM()
	if db == nil {
		config.Log.Error(InternalGetBDInstanceErr)
		return nil, ErrDataBase
	}
	lists := []*CronApi{}
	tx := db.Table("crontab_api").Order("update_at desc").Find(&lists)
	if tx.Error != nil {
		config.Log.Error(tx.Error)
		return nil, ErrCount
	}
	return lists, Success
}

func GetOneCronApi(id int) (*CronApi, *BriefMessage) {
	db := dbs.DBObj.GetGoRM()
	if db == nil {
		config.Log.Error(InternalGetBDInstanceErr)
		return nil, ErrDataBase
	}
	one := CronApi{}
	tx := db.Table("crontab_api").Where("id", id).Find(&one)
	if tx.Error != nil {
		config.Log.Error(tx.Error)
		return nil, ErrCount
	}
	return &one, Success
}

func PostCronApi(user *UserSessionInfo, newApi *CronApiPost) *BriefMessage {
	db := dbs.DBObj.GetGoRM()
	if db == nil {
		config.Log.Error(InternalGetBDInstanceErr)
		return ErrDataBase
	}
	cronApi := CronApi{
		ID:       0,
		Name:     newApi.Name,
		API:      newApi.API,
		UpdateAt: time.Now(),
		UpdateBy: user.Username,
	}
	tx := db.Table("crontab_api").Create(&cronApi)
	if tx.Error != nil {
		config.Log.Error(tx.Error)
		return ErrCreateDBData
	}
	return Success
}

func PutCronApi(user *UserSessionInfo, api *CronApiPost) *BriefMessage {
	db := dbs.DBObj.GetGoRM()
	if db == nil {
		config.Log.Error(InternalGetBDInstanceErr)
		return ErrDataBase
	}
	tx := db.Table("crontab_api").
		Where("id=?", api.ID).
		Update("name", api.Name).
		Update("api", api.API).
		Update("update_at", time.Now()).
		Update("update_by", user.Username)
	if tx.Error != nil {
		config.Log.Error(tx.Error)
		return ErrUpdateData
	}
	return Success
}

func PutCronApiStatus(user *UserSessionInfo, edi *EnabledInfo) *BriefMessage {
	db := dbs.DBObj.GetGoRM()
	if db == nil {
		config.Log.Error(InternalGetBDInstanceErr)
		return ErrDataBase
	}
	tx := db.Table("crontab_api").
		Where("id=?", edi.ID).
		Update("enabled", edi.Enabled).
		Update("update_at", time.Now()).
		Update("update_by", user.Username)
	if tx.Error != nil {
		config.Log.Error(tx.Error)
		return ErrUpdateData
	}
	return Success
}

func DelCronApi(id int64) *BriefMessage {
	db := dbs.DBObj.GetGoRM()
	if db == nil {
		config.Log.Error(InternalGetBDInstanceErr)
		return ErrDataBase
	}
	tx := db.Table("crontab_api").Where("id=?", id).Delete(nil)
	if tx.Error != nil {
		config.Log.Error(tx.Error)
		return ErrDelData
	}
	return Success
}

type Crontab struct {
	ID        int       `json:"id" gorm:"column:id"`
	Name      string    `json:"name" gorm:"column:name"`
	CID       string    `json:"cid" gorm:"column:cid"` // 有可能是多个
	Rule      string    `json:"rule" gorm:"column:rule"`
	Enabled   bool      `json:"enabled" gorm:"column:enabled"`
	ShowTitle bool      `json:"show_title" gorm:"column:show_title"`
	LineTitle string    `json:"line_title" gorm:"column:line_title"` // 生成线的标题
	ApiID     int       `json:"api_id" gorm:"column:api_id"`
	ExecCycle string    `json:"exec_cycle" gorm:"column:exec_cycle"` // 执行周期
	NearTime  int       `json:"near_time" gorm:"column:near_time"`
	Unit      string    `json:"unit" gorm:"column:unit"`
	Status    bool      `json:"status" gorm:"column:status"`
	UpdateAt  time.Time `json:"update_at" gorm:"column:update_at"`
	UpdateBy  string    `json:"update_by" gorm:"column:update_by"`
}

type CrontabSplit struct {
	Crontab
	RunTimes     int `json:"run_times" gorm:"-"`
	SuccessTimes int `json:"success_times" gorm:"-"`
	FailTimes    int `json:"fail_times" gorm:"-"`
}

type CrontabPost struct {
	ID        int    `json:"id" gorm:"column:id"`
	Name      string `json:"name" gorm:"column:name"`
	CID       string `json:"cid" gorm:"column:cid"` // 有可能是多个
	Rule      string `json:"rule" gorm:"column:rule"`
	ApiID     int    `json:"api_id" gorm:"column:api_id"`
	ExecCycle string `json:"exec_cycle" gorm:"column:exec_cycle"` // 执行周期
	NearTime  int    `json:"near_time" gorm:"column:near_time"`
	Unit      string `json:"unit" gorm:"column:unit"`
	Status    bool   `json:"status" gorm:"column:status"`
	Enabled   bool   `json:"enabled" gorm:"column:enabled"`
	ShowTitle bool   `json:"show_title" gorm:"column:show_title"`
	LineTitle string `json:"line_title" gorm:"column:line_title"` // 生成线的标题
}

func GetCronRules(sp *SplitPage) (*ResSplitPage, *BriefMessage) {
	if sp.PageSize <= 0 {
		sp.PageSize = 15
	}
	if sp.PageNo <= 0 {
		sp.PageNo = 1
	}
	db := dbs.DBObj.GetGoRM()
	if db == nil {
		config.Log.Error(InternalGetBDInstanceErr)
		return nil, ErrDataBase
	}
	var count int64
	var likeContent string
	var tx *gorm.DB
	tx = db.Table("crontab")
	if sp.Search != "" {
		likeContent = fmt.Sprint("%", sp.Search, "%")
		tx = tx.Where("name like ? or rule like ?", likeContent, likeContent)
	}
	tx = tx.Count(&count)
	if tx.Error != nil {
		config.Log.Error(tx.Error)
		return nil, ErrCount
	}
	lists := []*CrontabSplit{}
	tx = db.Table("crontab")
	if sp.Search != "" {
		tx = tx.Where("name like ? or rule like ?", likeContent, likeContent)
	}
	tx = tx.Order("update_at asc").
		Offset((sp.PageNo - 1) * sp.PageSize).
		Limit(sp.PageSize).
		Find(&lists)
	if tx.Error != nil {
		config.Log.Error(tx.Error)
		return nil, ErrCount
	}
	return CalSplitPage(sp, count, lists), Success
}

func PostCronRule(user *UserSessionInfo, newCron *CrontabPost) *BriefMessage {
	db := dbs.DBObj.GetGoRM()
	if db == nil {
		config.Log.Error(InternalGetBDInstanceErr)
		return ErrDataBase
	}
	cronApi := Crontab{
		ID:        0,
		Name:      newCron.Name,
		CID:       newCron.CID,
		Rule:      newCron.Rule,
		ApiID:     newCron.ApiID,
		ShowTitle: false,
		LineTitle: "",
		Enabled:   newCron.Enabled,
		ExecCycle: newCron.ExecCycle,
		NearTime:  newCron.NearTime,
		Unit:      newCron.Unit,
		UpdateAt:  time.Now(),
		UpdateBy:  user.Username,
	}
	tx := db.Table("crontab").Create(&cronApi)
	if tx.Error != nil {
		config.Log.Error(tx.Error)
		return ErrCreateDBData
	}
	return Success
}

func CronImport(crs []*CrontabPost) *BriefMessage {
	db := dbs.DBObj.GetGoRM()
	if db == nil {
		config.Log.Error(InternalGetBDInstanceErr)
		return ErrDataBase
	}
	currCronApi := []*Crontab{}
	tx := db.Table("crontab").Find(&currCronApi)
	if tx.Error != nil {
		config.Log.Error(tx.Error)
		return ErrSearchDBData
	}
	currCIDs := map[string]*Crontab{}
	for _, curr := range currCronApi {
		currCIDs[curr.CID] = curr
	}
	user := &UserSessionInfo{
		Username: "openapi",
	}
	for _, cr := range crs {
		if currCron, ok := currCIDs[cr.CID]; ok {
			cr.ID = currCron.ID
			if bf := PutCronRule(user, cr); bf != Success {
				return bf
			}
		} else {
			if bf := PostCronRule(user, cr); bf != Success {
				return bf
			}
		}
	}
	return Success
}

func PutCronRule(user *UserSessionInfo, api *CrontabPost) *BriefMessage {
	db := dbs.DBObj.GetGoRM()
	if db == nil {
		config.Log.Error(InternalGetBDInstanceErr)
		return ErrDataBase
	}
	tx := db.Table("crontab").
		Where("id=?", api.ID).
		Update("name", api.Name).
		Update("cid", api.CID).
		Update("rule", api.Rule).
		Update("api_id", api.ApiID).
		Update("enabled", api.Enabled).
		Update("show_title", api.ShowTitle).
		Update("line_title", api.LineTitle).
		Update("exec_cycle", api.ExecCycle).
		Update("near_time", api.NearTime).
		Update("unit", api.Unit).
		Update("update_at", time.Now()).
		Update("update_by", user.Username)
	if tx.Error != nil {
		config.Log.Error(tx.Error)
		return ErrUpdateData
	}
	return Success
}

func DelCronRule(id int64) *BriefMessage {
	db := dbs.DBObj.GetGoRM()
	if db == nil {
		config.Log.Error(InternalGetBDInstanceErr)
		return ErrDataBase
	}
	tx := db.Table("crontab").Where("id=?", id).Delete(nil)
	if tx.Error != nil {
		config.Log.Error(tx.Error)
		return ErrDelData
	}
	return Success
}

func PutCronRuleStatus(user *UserSessionInfo, edi *EnabledInfo) *BriefMessage {
	db := dbs.DBObj.GetGoRM()
	if db == nil {
		config.Log.Error(InternalGetBDInstanceErr)
		return ErrDataBase
	}
	tx := db.Table("crontab").
		Where("id=?", edi.ID).
		Update("enabled", edi.Enabled).
		Update("update_at", time.Now()).
		Update("update_by", user.Username)
	if tx.Error != nil {
		config.Log.Error(tx.Error)
		return ErrUpdateData
	}
	return Success
}

func DelRulesSelection(delIDs []int) *BriefMessage {
	db := dbs.DBObj.GetGoRM()
	if db == nil {
		config.Log.Error(InternalGetBDInstanceErr)
		return ErrDataBase
	}
	err := db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Table("crontab").Where("id in (?)", delIDs).Delete(nil).Error; err != nil {
			config.Log.Error(tx.Error)
			return err
		}
		return nil
	})
	if err != nil {
		return ErrDelData
	}
	return Success
}

func PutRulesEnable(user *UserSessionInfo, enabledIDs []int) *BriefMessage {
	db := dbs.DBObj.GetGoRM()
	if db == nil {
		config.Log.Error(InternalGetBDInstanceErr)
		return ErrDataBase
	}
	tx := db.Table("crontab").Where("id in (?)", enabledIDs).
		Update("enabled", 1).
		Update("update_at", time.Now()).
		Update("update_by", user.Username)
	if tx.Error != nil {
		config.Log.Error(tx.Error)
		return ErrUpdateData
	}
	return Success
}

func PutRulesDisable(user *UserSessionInfo, disableIDs []int) *BriefMessage {
	db := dbs.DBObj.GetGoRM()
	if db == nil {
		config.Log.Error(InternalGetBDInstanceErr)
		return ErrDataBase
	}
	tx := db.Table("crontab").Where("id in (?)", disableIDs).
		Update("enabled", 0).
		Update("update_at", time.Now()).
		Update("update_by", user.Username)
	if tx.Error != nil {
		config.Log.Error(tx.Error)
		return ErrUpdateData
	}
	return Success
}

type LoadTitleRule struct {
	Rule  string `json:"rule" form:"rule"`
	ApiID int    `json:"api_id" form:"api_id"`
}

type loadTitleQuery struct {
	Query string `json:"query" form:"query" qs:"query"`
}

type prometheusQueryResult struct {
	Metric map[string]string
	Value  []interface{}
}

type prometheusQueryData struct {
	Result     []*prometheusQueryResult `json:"result"`
	ResultType string                   `json:"resultType"`
}

type prometheusQueryStc struct {
	Data   *prometheusQueryData `json:"data"`
	Status string               `json:"status"`
}

func LoadTitle(ltr *LoadTitleRule) ([]map[string]string, *BriefMessage) {
	if ltr.Rule == "" {
		return nil, ErrPostData
	}
	apiInfo, bf := GetOneCronApi(ltr.ApiID)
	if bf != Success {
		return nil, bf
	}
	if apiInfo.API == "" {
		return nil, ErrApiInterfaceNoFound
	}
	queryUri := ""
	if strings.HasSuffix(apiInfo.API, "/") {
		queryUri = apiInfo.API + "api/v1/query"
	} else {
		queryUri = apiInfo.API + "/api/v1/query"
	}
	ltq := loadTitleQuery{
		Query: ltr.Rule,
	}
	encoder := qs.NewEncoder()
	q, err := encoder.Values(&ltq)
	if err != nil {
		config.Log.Error(err)
		return nil, ErrPostData
	}
	queryUri = queryUri + "?" + q.Encode()
	content, err := utils.Get(queryUri)
	if err != nil {
		config.Log.Error(err)
		return nil, ErrMonitorApi
	}
	/*

		{
		  "resultType": "matrix" | "vector" | "scalar" | "string",
		  "result": <value>
		}

		$ curl 'http://localhost:9090/api/v1/query?query=up&time=2015-07-01T20:10:51.781Z'
		{
		   "status" : "success",
		   "data" : {
			  "resultType" : "vector",
			  "result" : [
				 {
					"metric" : {
					   "__name__" : "up",
					   "job" : "prometheus",
					   "instance" : "localhost:9090"
					},
					"value": [ 1435781451.781, "1" ]
				 },
				 {
					"metric" : {
					   "__name__" : "up",
					   "job" : "node",
					   "instance" : "localhost:9100"
					},
					"value" : [ 1435781451.781, "0" ]
				 }
			  ]
		   }
		}
	*/
	respStc := prometheusQueryStc{}
	if err := json.Unmarshal(content, &respStc); err != nil {
		config.Log.Error(err)
		return nil, ErrMonitorApi
	}
	// 取第一个元素的mertric返回
	if respStc.Status != "success" {
		return nil, ErrApiReturnErr
	}
	if len(respStc.Data.Result) == 0 {
		return nil, ErrApiReturnEmpty
	}
	respSlice := []map[string]string{}
	for k, v := range respStc.Data.Result[0].Metric {
		item := map[string]string{
			"value": v,
			"label": k,
		}
		respSlice = append(respSlice, item)
	}
	return respSlice, Success
}

func CheckStatus() *BriefMessage {
	db := dbs.DBObj.GetGoRM()
	if db == nil {
		config.Log.Error(InternalGetBDInstanceErr)
		return ErrDataBase
	}
	crons := []*Crontab{}
	tx := db.Table("crontab").Find(&crons)
	if tx.Error != nil {
		config.Log.Error(tx.Error)
		return ErrUpdateData
	}
	for _, cron := range crons {
		ltr := LoadTitleRule{
			ApiID: cron.ApiID,
			Rule:  cron.Rule,
		}
		if _, bf := LoadTitle(&ltr); bf != Success {
			tx = db.Table("crontab").Where("id", cron.ID).Update("status", false)
			if tx.Error != nil {
				return ErrUpdateData
			}
		} else {
			tx = db.Table("crontab").Where("id", cron.ID).Update("status", true)
			if tx.Error != nil {
				return ErrUpdateData
			}
		}
	}
	return Success
}
