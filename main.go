package main

import (
	"pro_cfg_manager/config"
	"pro_cfg_manager/dbs"
	"pro_cfg_manager/router"
)

func main() {
	config.InitLogger()
	dbs.DBObj.Init()
	router.ListenAndServer()
}
