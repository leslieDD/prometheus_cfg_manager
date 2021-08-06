package models

import (
	"fmt"
	"pro_cfg_manager/config"
	"pro_cfg_manager/dbs"
	"pro_cfg_manager/utils"
	"sync/atomic"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type OperateObj struct {
	operateChan chan *OperationLog
}

var OO = NewOpterationLog()

type OperationLog struct {
	ID            int       `json:"id" gorm:"column:id"`
	UserName      string    `json:"username" gorm:"column:username"`
	Ipaddr        string    `json:"ipaddr" gorm:"column:ipaddr"`
	OperateName   string    `json:"operate_name" gorm:"column:operate_name"`
	OperateResult bool      `json:"operate_result" gorm:"column:operate_result"`
	OperateAt     time.Time `json:"operate_at" gorm:"column:operate_at"`
	OperateError  string    `json:"operate_error" gorm:"column:operate_error"`
}

func (ol *OperationLog) String() string {
	return fmt.Sprintf("username: %s, ipaddr: %s, "+
		"operate_name: %s, operate_result: %v, operate_at: %s, operate_error: %s",
		ol.UserName,
		ol.Ipaddr,
		ol.OperateName,
		ol.OperateResult,
		ol.OperateAt,
		ol.OperateError,
	)
}

func NewOpterationLog() *OperateObj {
	oo := &OperateObj{}
	oo.operateChan = make(chan *OperationLog, 10)
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

func FlagLog(userName, ipaddr, optName string, opt_err error) {
	opl := &OperationLog{
		UserName:      userName,
		Ipaddr:        ipaddr,
		OperateName:   optName,
		OperateAt:     time.Now(),
		OperateResult: opt_err == nil,
	}
	if opt_err != nil {
		opl.OperateError = opt_err.Error()
	} else {
		opl.OperateError = "操作成功"
	}
	OO.Log(opl)
}

func (o *OperateObj) loopWrite() {
	for ol := range o.operateChan {
		if r := o.writeLog(ol); r != Success {
			config.Log.Error(ol)
			if r == ErrDataBase {
				time.Sleep(5 * time.Second)
			}
		}
	}
}

func (o *OperateObj) writeLog(opl *OperationLog) *BriefMessage {
	db := dbs.DBObj.GetGoRM()
	if db == nil {
		config.Log.Error(InternalGetBDInstanceErr)
		return ErrDataBase
	}
	opl.ID = 0
	tx := db.Table("operation_log").Create(opl)
	if tx.Error != nil {
		config.Log.Error(tx.Error)
		return ErrCreateDBData
	}
	return Success
}

func (o *OperateObj) Log(opl *OperationLog) {
	select {
	case o.operateChan <- opl:
	default:
	}
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

func OptResetSystem(code *ResetCode, ipAddr string) *BriefMessage {
	var err error
	defer func() {
		FlagLog("", ipAddr, "reset_system", err)
	}()
	if ResetBlock.AnyOne() {
		err = fmt.Errorf("running, try again later.")
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
		err = fmt.Errorf("%s", InternalGetBDInstanceErr)
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
			Tmpl: DefTmplText,
		}
		if err := tx.Table("tmpl").Create(&tmpl).Error; err != nil {
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
		err = fmt.Errorf("running, try again later.")
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
		err = fmt.Errorf("%s", InternalGetBDInstanceErr)
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
