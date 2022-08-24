package alertmgr

import (
	"pro_cfg_manager/config"
	"pro_cfg_manager/dbs"
)

func GetChartManual(id int) string {
	db := dbs.DBObj.GetGoRM()
	if db == nil {
		config.Log.Error("get db session err")
		return ""
	}
	rule := CronRule{}
	tx := db.Raw(`SELECT crontab.*, crontab_api.api, crontab_api.name as api_name FROM crontab 
	LEFT join crontab_api 
	on crontab.api_id=crontab_api.id
	WHERE crontab.enabled=1 and crontab.id=?
	`, id).Find(&rule)
	if tx.Error != nil {
		config.Log.Error(tx.Error)
		return ""
	}
	if rule.ID == 0 {
		config.Log.Warnf("no record found")
		return ""
	}
	image, _ := ChartLine(&rule)
	return image
}