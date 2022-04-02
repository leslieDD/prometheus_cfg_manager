package main

import (
	"pro_cfg_manager/collector"
	"pro_cfg_manager/config"
	"pro_cfg_manager/dbs"
	"pro_cfg_manager/models"
	"pro_cfg_manager/router"
)

func main() {
	collector.PaoPao()
	config.DoParams()
	config.InitLogger()
	dbs.DBObj.Init()
	models.InitApp()
	router.ListenAndServer()
}
