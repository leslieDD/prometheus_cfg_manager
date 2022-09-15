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
	WHERE crontab.id=?
	`, id).Find(&rule)
	if tx.Error != nil {
		config.Log.Error(tx.Error)
		return ""
	}
	if rule.ID == 0 {
		config.Log.Warnf("no record found")
		return ""
	}
	imageContent, err := ChartLine(&rule)
	if err != nil {
		taskRunStatusObj.AddFail(rule.ID)
		return ""
	}
	taskRunStatusObj.AddSuccess(rule.ID)
	// ic := ImageContent{
	// 	Title:  rule.Name,
	// 	Width:  "1000",
	// 	Height: "600",
	// 	Image:  imageContent,
	// }
	// sendEmail(&ic)
	return imageContent
}

func SendTestMail() bool {
	imageContent, err := ChartLineTest()
	if err != nil {
		return false
	}
	ic := ImageContent{
		Title:  "发送测试邮件（send test email）",
		Width:  1000,
		Height: 600,
		Image:  imageContent,
	}
	return sendEmailTest(&ic)
}
