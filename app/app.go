package app

import (
	"pro_cfg_manager/alertmgr"
	"pro_cfg_manager/models"
	_ "pro_cfg_manager/sync"
)

func InitApp() {
	models.OO = models.NewOpterationLog()
	models.AllowOneObj = models.NewOnlyAllowOne()
	models.MObj = models.NewMonitor()
	models.PublishObj = models.NewPublishResolve()
	models.SSObj = models.NewSessionCache()
	models.TmplObj = models.NewTmpl()
	models.CmdArea = models.NewCmdArea()
	models.PR = models.NewPR()
	models.SPCache = models.NewSessionParamsCache()
	alertmgr.AlertMgrObj = alertmgr.NewAlertMgr()
	alertmgr.AlertMgrObj.Run()
	alertmgr.ChartCronObj = alertmgr.NewChartCron()
	alertmgr.ChartCronObj.Run()
	go models.SPCache.Update(true)
}
