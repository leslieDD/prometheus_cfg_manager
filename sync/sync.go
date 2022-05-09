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
	UpdatePartIPIDCAndLineInfo
	UpdateAllIPLabelsInJobGroup
	ExpandIPAddress
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
	SyncDone
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

func (a *ActionSync) CanNoDo(action ActionName) bool {
	a.lock.Lock()
	defer a.lock.Unlock()

	obj, ok := a.Actions[action]
	if !ok {
		a.Actions[action] = &SyncInfo{
			EnterTime:  time.Now(),
			RunTimes:   1,
			SyncStatus: SyncRuning,
		}
		return false
	}
	if obj.SyncStatus == SyncRuning {
		return true
	}
	obj.EnterTime = time.Now()
	obj.RunTimes += 1
	obj.SyncStatus = SyncRuning
	return false
}

func (a *ActionSync) Done(action ActionName) {
	a.lock.Lock()
	defer a.lock.Unlock()

	obj, ok := a.Actions[action]
	if !ok {
		return
	}
	obj.SyncStatus = SyncDone
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

	for actionType, obj := range a.Actions {
		if obj.SyncStatus != SyncRuning {
			continue
		}
		if time.Since(obj.EnterTime).Seconds() > 30 {
			config.Log.Printf("action: %d, status: %d, duration: %v", actionType, obj.SyncStatus, time.Since(obj.EnterTime).Seconds())
			obj.SyncStatus = SyncDone
		}
	}
}
