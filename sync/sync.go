package sync

import (
	"pro_cfg_manager/config"
	"sync"
	"time"
)

type ActionName int

const (
	BatchImportIPAddrFromFile ActionName = iota + 1
	BatchImportIpAddrFromWeb
	UpdateAllIPPosition
	UpdateAllIPIDCAndLineInfo
	UpdateAllIPLineAndLineInfo
	UpdateAllIPLabelsInJobGroup
	ReSetAllData
	ReStartPrometheusService
	PublishMonitorRules
	PublishEmptyMonitorRules
	ReCreateAllPrometheusConfig
	ReloadAllPrometheusConfig
	ReCreateAndReloadPrometheusConfig
	ReSetAllPrivilegesData
)

type SyncStatus int

const (
	SyncRuning SyncStatus = iota + 1
	SyncStop
)

type SyncInfo struct {
	EnterTime  time.Time  // 进入时间
	RunTimes   int        // 运行的次数
	SyncStatus SyncStatus // 当前状态
}

type ActionSync struct {
	lock    sync.Mutex
	Actions map[ActionName]*SyncInfo
}

var AS = NewActionSync()

func NewActionSync() *ActionSync {
	as := &ActionSync{
		lock:    sync.Mutex{},
		Actions: map[ActionName]*SyncInfo{},
	}
	go as.doClean()
	return as
}

func (a *ActionSync) Init(action ActionName) bool {
	a.lock.Lock()
	defer a.lock.Unlock()

	obj, ok := a.Actions[action]
	if !ok {
		a.Actions[action] = &SyncInfo{
			EnterTime:  time.Now(),
			RunTimes:   1,
			SyncStatus: SyncRuning,
		}
		return true
	}
	if obj.SyncStatus == SyncRuning {
		return false
	}
	obj.EnterTime = time.Now()
	obj.RunTimes += 1
	return true
}

func (a *ActionSync) doClean() {
	for {
		select {
		case <-time.After(5 * time.Second):
			go a.clear()
		}
	}
}

func (a *ActionSync) clear() {
	a.lock.Lock()
	defer a.lock.Unlock()

	for _, obj := range a.Actions {
		if time.Since(obj.EnterTime).Seconds() > 30 {
			config.Log.Printf("action: %d, duration: %d", time.Since(obj.EnterTime).Seconds())
		}
	}
}
