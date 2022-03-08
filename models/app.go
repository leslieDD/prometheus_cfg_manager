package models

import (
	_ "pro_cfg_manager/sync"
)

func InitApp() {
	OO = NewOpterationLog()
	AllowOneObj = NewOnlyAllowOne()
	mObj = NewMonitor()
	publish = NewPublishResolve()
	SSObj = NewSessionCache()
	TmplObj = NewTmpl()
	cmdArea = NewCmdArea()
	PR = NewPR()
}
