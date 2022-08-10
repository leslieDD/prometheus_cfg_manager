package main

import (
	"pro_cfg_manager/app"
	"pro_cfg_manager/config"
	"pro_cfg_manager/dbs"
	"pro_cfg_manager/router"
)

func main() {
	config.DoParams()
	config.InitLogger()
	dbs.DBObj.Init()
	app.InitApp()
	router.ListenAndServer()
}
