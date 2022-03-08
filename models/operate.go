package models

import (
	"errors"
	"fmt"
	"pro_cfg_manager/config"
	"pro_cfg_manager/dbs"
	"pro_cfg_manager/utils"
	"sync"
	"sync/atomic"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type LogType int

const (
	OperateLog LogType = iota + 1
	SystemLog
)

type RecodeType int

const (
	IsSearch RecodeType = iota + 1
	IsAdd
	IsDel
	IsUpdate
	IsRunning
	IsPublish
	IsLogin
	IsReset
)

type OperateObj struct {
	operateChan  chan *OperationLogMess
	recodesLevel map[RecodeType]*LogLevelSetting
	lock         sync.Mutex
}

var OO *OperateObj

type OperationLog struct {
	ID             int       `json:"id" gorm:"column:id"`
	OperateType    string    `json:"operate_type" gorm:"column:operate_type"`
	UserName       string    `json:"username" gorm:"column:username"`
	Ipaddr         string    `json:"ipaddr" gorm:"column:ipaddr"`
	OperateContent string    `json:"operate_content" gorm:"column:operate_content"`
	OperateResult  bool      `json:"operate_result" gorm:"column:operate_result"`
	OperateAt      time.Time `json:"operate_at" gorm:"column:operate_at"`
	OperateError   string    `json:"operate_error" gorm:"column:operate_error"`
}

type OperationLogMess struct {
	OptLog      *OperationLog
	ThisLogType LogType
	RecodeType  RecodeType
}

func (ol *OperationLog) String() string {
	return fmt.Sprintf("username: %s, ipaddr: %s, "+
		"operate_content: %s, operate_result: %v, operate_at: %s, operate_error: %s",
		ol.UserName,
		ol.Ipaddr,
		ol.OperateContent,
		ol.OperateResult,
		ol.OperateAt,
		ol.OperateError,
	)
}

func NewOpterationLog() *OperateObj {
	oo := &OperateObj{
		lock: sync.Mutex{},
	}
	oo.operateChan = make(chan *OperationLogMess, 10)
	oo.recodesLevel = map[RecodeType]*LogLevelSetting{}
	oo.FlushLevel()
	go oo.loopWrite()
	return oo
}

func GetOperateLog(sp *SplitPage) (*ResSplitPage, *BriefMessage) {
	db := dbs.DBObj.GetGoRM()
	if db == nil {
		config.Log.Error(InternalGetBDInstanceErr)
		return nil, ErrDataBase
	}
	var count int64
	var likeSql string
	tx := db.Table("operation_log")
	if sp.Search != "" {
		like := `'%` + sp.Search + `%'`
		likeSql = fmt.Sprintf("username like %s "+
			"or ipaddr like %s "+
			"or operate_name like %s "+
			"or operate_error like %s ", like, like, like, like)
	} else {
		likeSql = "1=1"
	}
	txCount := tx.Where(likeSql).Count(&count)
	if txCount.Error != nil {
		config.Log.Error(txCount.Error)
		return nil, ErrCount
	}
	logs := []*OperationLog{}
	tx = db.Table("operation_log")
	if sp.Search != "" {
		tx = tx.Where(likeSql)
	}
	tx = tx.Order("operate_at desc").
		Offset((sp.PageNo - 1) * sp.PageSize).
		Limit(sp.PageSize).
		Find(&logs)
	if tx.Error != nil {
		config.Log.Error(tx.Error)
		return nil, ErrSearchDBData
	}
	return CalSplitPage(sp, count, logs), Success
}

// 记录重置日志，可以查看“Prometheus”中的“基本配置”中的“重置”可以看到内容
func (o *OperateObj) FlagLog(userName, ipaddr, optContent string, operateType RecodeType, opt_err *BriefMessage) {
	o.lock.Lock()
	r, ok := o.recodesLevel[operateType]
	o.lock.Unlock()
	if !ok || !r.Selected {
		return
	}
	opl := &OperationLogMess{
		OptLog: &OperationLog{
			UserName:       userName,
			Ipaddr:         ipaddr,
			OperateContent: optContent,
			OperateAt:      time.Now(),
			OperateResult:  opt_err == Success,
			OperateError:   opt_err.String(),
		},
		ThisLogType: OperateLog,
		RecodeType:  operateType,
	}
	OO.Log(opl)
}

// 记录系统日志，在“用户及权限管理”中的“日志”页，可以查看内容
func (o *OperateObj) RecodeLog(userName, ipaddr, optContent string, operateType RecodeType, opt_err *BriefMessage) {
	o.lock.Lock()
	r, ok := o.recodesLevel[operateType]
	o.lock.Unlock()
	if !ok || !r.Selected {
		if opt_err == Success { // 有错误的都记录下来
			return
		}
	}
	opl := &OperationLogMess{
		OptLog: &OperationLog{
			UserName:       userName,
			Ipaddr:         ipaddr,
			OperateContent: optContent,
			OperateAt:      time.Now(),
			OperateResult:  opt_err == Success,
			OperateError:   opt_err.String(),
			OperateType:    r.Label,
		},
		ThisLogType: SystemLog,
		RecodeType:  operateType,
	}
	OO.Log(opl)
}

func (o *OperateObj) loopWrite() {
	for ol := range o.operateChan {
		if r := o.writeLog(ol); r != Success {
			if r == ErrDataBase {
				time.Sleep(5 * time.Second)
			}
		}
	}
}

func (o *OperateObj) writeLog(opl *OperationLogMess) *BriefMessage {
	db := dbs.DBObj.GetGoRM()
	if db == nil {
		config.Log.Error(InternalGetBDInstanceErr)
		return ErrDataBase
	}
	opl.OptLog.ID = 0
	var tx *gorm.DB
	switch opl.ThisLogType {
	case OperateLog:
		tx = db.Table("operation_log")
	case SystemLog:
		tx = db.Table("system_log")
	default:
		return Success
	}
	tx = tx.Create(opl.OptLog)
	if tx.Error != nil {
		config.Log.Error(tx.Error)
		return ErrCreateDBData
	}
	return Success
}

func (o *OperateObj) Log(opl *OperationLogMess) {
	select {
	case o.operateChan <- opl:
	default:
	}
}

func (o *OperateObj) FlushLevel() {
	o.lock.Lock()
	defer o.lock.Unlock()

	db := dbs.DBObj.GetGoRM()
	if db == nil {
		config.Log.Error(InternalGetBDInstanceErr)
		return
	}
	recodes := []*LogLevelSetting{}
	tx := db.Table("log_setting").Find(&recodes)
	if tx.Error != nil {
		config.Log.Error(tx.Error)
		return
	}
	recodeMap := map[RecodeType]*LogLevelSetting{}
	for _, r := range recodes {
		recodeMap[RecodeType(r.ID)] = r
	}
	o.recodesLevel = recodeMap
}

var ResetSystemKey = atomic.Value{}
var ResetBlock = utils.NewNoBlockLock()

func PreOptResetSystem() *BriefMessage {
	key := uuid.NewString()
	ResetSystemKey.Store(key)
	config.Log.Warnf("reset prometheus config key: %s", key)
	return Success
}

type ResetCode struct {
	Code string `json:"code" form:"code"`
}

func OptResetSystem(user *UserSessionInfo, code *ResetCode) *BriefMessage {
	var err error
	if ResetBlock.AnyOne() {
		err = errors.New("running, try again later")
		config.Log.Error(err)
		return ErrAlreadyRunning
	}
	defer ResetBlock.Done()
	keyStore, ok := ResetSystemKey.Load().(string)
	if !ok {
		err = fmt.Errorf("load key error")
		config.Log.Error(err)
		return ErrNoResetKey
	}
	if keyStore != code.Code {
		err = fmt.Errorf("重置KEY不匹配")
		config.Log.Error(err)
		return ErrResetKeyDiff
	}
	db := dbs.DBObj.GetGoRM()
	if db == nil {
		config.Log.Error(InternalGetBDInstanceErr)
		return ErrDataBase
	}
	tableCleard := []string{
		"annotations",
		"group_labels",
		"group_machines",
		"jobs",
		"job_group",
		"job_machines",
		"labels",
		"machines",
		"monitor_labels",
		"monitor_rules",
		// "relabels",
		"rules_groups",
		"sub_group",
		"tmpl",
		"tmpl_fields",
		"idc",
		"line",
		"pool",
	}
	err = db.Transaction(func(tx *gorm.DB) error {
		for _, tableName := range tableCleard {
			if err := tx.Table(tableName).Where("1=1").Delete(nil).Error; err != nil {
				return err
			}
			if err := tx.Table(tableName).Exec(fmt.Sprintf("ALTER TABLE %s AUTO_INCREMENT = 1;", tableName)).Error; err != nil {
				return err
			}
		}
		if err := tx.Table("options").Where("1=1").Update("opt_value", "true").Error; err != nil {
			return err
		}
		if err := tx.Table("relabels").Where("id<>1").Delete(nil).Error; err != nil {
			return err
		}
		tmpl := TmplTable{
			Tmpl:     DefTmplText,
			UpdateAt: time.Now(),
			UpdateBy: "administrator",
		}
		if err := tx.Table("tmpl").Create(&tmpl).Error; err != nil {
			return err
		}
		fields := []BaseFields{
			{Key: "metrics", Value: "/metrics", Enabled: true, UpdateAt: time.Now()},
			{Key: "refresh_interval", Value: "15s", Enabled: true, UpdateAt: time.Now()},
		}
		if err := tx.Table("tmpl_fields").Create(&fields).Error; err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return ErrReset
	}
	return Success
}

var ResetAdminKey = atomic.Value{}
var ResetAdminBlock = utils.NewNoBlockLock()

func PreOptResetAdmin() *BriefMessage {
	key := uuid.NewString()
	ResetAdminKey.Store(key)
	config.Log.Warnf("reset admin config key: %s", key)
	return Success
}

func OptResetAdmin(code *ResetCode, ipAddr string) *BriefMessage {
	var err error
	if ResetAdminBlock.AnyOne() {
		err = fmt.Errorf("running, try again later")
		config.Log.Error(err)
		return ErrAlreadyRunning
	}
	defer ResetAdminBlock.Done()
	keyStore, ok := ResetAdminKey.Load().(string)
	if !ok {
		err = fmt.Errorf("load key error")
		config.Log.Error(err)
		return ErrNoResetKey
	}
	if keyStore != code.Code {
		err = fmt.Errorf("重置KEY不匹配")
		config.Log.Error(err)
		return ErrResetKeyDiff
	}
	db := dbs.DBObj.GetGoRM()
	if db == nil {
		config.Log.Error(InternalGetBDInstanceErr)
		return ErrDataBase
	}
	tableCleard := []string{
		"manager_group",
		"manager_user",
		"group_priv",
		"manager_set",
	}
	err = db.Transaction(func(tx *gorm.DB) error {
		for _, tableName := range tableCleard {
			if err := tx.Table(tableName).Where("1=1").Delete(nil).Error; err != nil {
				return err
			}
			if err := tx.Table(tableName).Exec(fmt.Sprintf("ALTER TABLE %s AUTO_INCREMENT = 1;", tableName)).Error; err != nil {
				return err
			}
		}
		ug := ManagerGroup{
			Name:     "administrator",
			Enabled:  true,
			UpdateAt: time.Now(),
		}
		if err := tx.Table("manager_group").Create(&ug).Error; err != nil {
			return err
		}
		uu := ManagerUser{
			UserName: "admin",
			Password: "admin",
			NiceName: "管理员",
			Salt:     uuid.NewString(),
			Phone:    "10086",
			GroupID:  ug.ID,
			Enabled:  true,
			UpdateAt: time.Now(),
			CreateAt: time.Now(),
		}
		uu.Password = utils.CreateHashword(uu.Password, uu.Salt)
		if err := tx.Table("manager_user").Create(&uu).Error; err != nil {
			return err
		}
		params := []AdminParams{
			{
				ParamName:  "default_group",
				ParamValue: "",
			},
		}
		if err := tx.Table("manager_set").Create(&params).Error; err != nil {
			return err
		}
		pIDs := []OnlyID{}
		if err := tx.Table("page_function").Select("id").Find(&pIDs).Error; err != nil {
			return err
		}
		if len(pIDs) != 0 {
			uPriv := []*TableGroupPriv{}
			for _, p := range pIDs {
				uPriv = append(uPriv, &TableGroupPriv{
					GroupID: ug.ID,
					FuncID:  p.ID,
				})
			}
			if err := tx.Table("group_priv").Create(&uPriv).Error; err != nil {
				return err
			}
		}
		return nil
	})
	if err != nil {
		return ErrReset
	}
	return Success
}
