package models

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
