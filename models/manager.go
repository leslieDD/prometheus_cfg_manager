package models

import (
	"fmt"
	"pro_cfg_manager/config"
	"pro_cfg_manager/dbs"
	"pro_cfg_manager/utils"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ManagerGroup struct {
	ID       int       `json:"id" gorm:"column:id"`
	Name     string    `json:"name" gorm:"column:name"`
	Enabled  bool      `json:"enabled" gorm:"column:enabled"`
	UpdateAt time.Time `json:"update_at" gorm:"column:update_at"`
	UpdateBy string    `json:"update_by" gorm:"column:update_by"`
}

type ManagerGroupList struct {
	UserCount int `json:"user_count" gorm:"column:user_count"`
	ManagerGroup
}

type ManagerUser struct {
	ID       int       `json:"id" gorm:"column:id"`
	UserName string    `json:"username" gorm:"column:username"`
	NiceName string    `json:"nice_name" gorm:"column:nice_name"`
	Password string    `json:"password" gorm:"column:password"`
	Phone    string    `json:"phone" gorm:"column:phone"`
	Salt     string    `json:"-" gorm:"column:salt"`
	GroupID  int       `json:"group_id" gorm:"column:group_id"`
	Enabled  bool      `json:"enabled" gorm:"column:enabled"`
	UpdateAt time.Time `json:"update_at" gorm:"column:update_at"`
	CreateAt time.Time `json:"create_at" gorm:"column:create_at"`
	UpdateBy string    `json:"update_by" gorm:"column:update_by"`
}

type ManagerUserList struct {
	GroupName string `json:"group_name" gorm:"column:group_name"`
	ManagerUser
}

type UserLogInfo struct {
	Username string `json:"username" form:"username"`
	Password string `json:"password" form:"password"`
}

type RegisterInfo struct {
	Username string `json:"username" form:"username"`
	NiceName string `json:"nice_name" form:"nice_name"`
	Password string `json:"password" form:"password"`
	Phone    string `json:"phone" form:"phone"`
	GroupID  int    `json:"group_id" form:"group_id"`
}

func GetManagerGroups(sp *SplitPage) (*ResSplitPage, *BriefMessage) {
	db := dbs.DBObj.GetGoRM()
	if db == nil {
		config.Log.Error(InternalGetBDInstanceErr)
		return nil, ErrDataBase
	}
	var count int64
	var tx *gorm.DB
	tx = db.Table("manager_group")
	if sp.Search != "" {
		tx = tx.Where("name like ?", fmt.Sprint("%", sp.Search, "%"))
	}
	tx = tx.Count(&count)
	if tx.Error != nil {
		config.Log.Error(tx.Error)
		return nil, ErrCount
	}
	mgs := []*ManagerGroupList{}
	tx = db.Table("manager_group").
		Select(" manager_group.*, COUNT(manager_user.username) as user_count ").
		Joins("LEFT JOIN manager_user ON manager_group.id=manager_user.group_id")
	if sp.Search != "" {
		tx = tx.Where("name like ?", fmt.Sprint("%", sp.Search, "%"))
	}
	tx = tx.Group("manager_group.id").
		Order("manager_group.update_at desc ").
		Offset((sp.PageNo - 1) * sp.PageSize).
		Limit(sp.PageSize).
		Find(&mgs)
	if tx.Error != nil {
		config.Log.Error(tx.Error)
		return nil, ErrSearchDBData
	}
	return CalSplitPage(sp, count, mgs), Success
}

func GetManagerGroupList(user *UserSessionInfo) ([]*ManagerGroup, *BriefMessage) {
	db := dbs.DBObj.GetGoRM()
	if db == nil {
		config.Log.Error(InternalGetBDInstanceErr)
		return nil, ErrDataBase
	}
	mgs := []*ManagerGroup{}
	tx := db.Table("manager_group").Find(&mgs)
	if tx.Error != nil {
		config.Log.Error(tx.Error)
		return nil, ErrSearchDBData
	}
	return mgs, Success
}

func GetManagerGroupEnabled() ([]*ManagerGroup, *BriefMessage) {
	db := dbs.DBObj.GetGoRM()
	if db == nil {
		config.Log.Error(InternalGetBDInstanceErr)
		return nil, ErrDataBase
	}
	mgs := []*ManagerGroup{}
	tx := db.Table("manager_group").Where("enabled=1").Find(&mgs)
	if tx.Error != nil {
		config.Log.Error(tx.Error)
		return nil, ErrSearchDBData
	}
	return mgs, Success
}

func PostManagerGroup(user *UserSessionInfo, mg *ManagerGroup) *BriefMessage {
	db := dbs.DBObj.GetGoRM()
	if db == nil {
		config.Log.Error(InternalGetBDInstanceErr)
		return ErrDataBase
	}
	mg.ID = 0
	mg.UpdateAt = time.Now()
	mg.UpdateBy = user.Username
	tx := db.Table("manager_group").Create(&mg)
	if tx.Error != nil {
		config.Log.Error(tx.Error)
		return ErrCreateDBData
	}
	return Success
}

func PutManagerGroup(user *UserSessionInfo, mg *ManagerGroup) *BriefMessage {
	db := dbs.DBObj.GetGoRM()
	if db == nil {
		config.Log.Error(InternalGetBDInstanceErr)
		return ErrDataBase
	}
	tx := db.Table("manager_group").
		Where("id=?", mg.ID).
		Update("name", mg.Name).
		// Update("enabled", mg.Enabled).
		Update("update_at", time.Now()).
		Update("update_by", user.Username)
	if tx.Error != nil {
		config.Log.Error(tx.Error)
		return ErrUpdateData
	}
	return Success
}

func DeleteManagerGroup(mgID int64) *BriefMessage {
	db := dbs.DBObj.GetGoRM()
	if db == nil {
		config.Log.Error(InternalGetBDInstanceErr)
		return ErrDataBase
	}
	tErr := db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Table("manager_group").Where("id=?", mgID).Delete(nil).Error; err != nil {
			config.Log.Error(tx.Error)
			return err
		}
		if err := tx.Table("group_priv").Where("group_id=?", mgID).Delete(nil).Error; err != nil {
			config.Log.Error(tx.Error)
			return err
		}
		return nil
	})
	if tErr != nil {
		return ErrDelData
	}
	return Success
}

func PutManagerGroupStatus(user *UserSessionInfo, info *EnabledInfo) *BriefMessage {
	db := dbs.DBObj.GetGoRM()
	if db == nil {
		config.Log.Error(InternalGetBDInstanceErr)
		return ErrDataBase
	}
	tx := db.Table("manager_group").
		Where("id=?", info.ID).
		Update("enabled", info.Enabled).
		Update("update_at", time.Now()).
		Update("update_by", user.Username)
	if tx.Error != nil {
		config.Log.Error(tx.Error)
		return ErrUpdateData
	}
	return Success
}

func GetManagerUsers(sp *SplitPage) (*ResSplitPage, *BriefMessage) {
	db := dbs.DBObj.GetGoRM()
	if db == nil {
		config.Log.Error(InternalGetBDInstanceErr)
		return nil, ErrDataBase
	}
	var count int64
	var tx *gorm.DB
	tx = db.Table("manager_user").Select("count(*) as count")
	if sp.Search != "" {
		tx = tx.Where("manager_user.username like ? or manager_group.name like ?",
			fmt.Sprint("%", sp.Search, "%"),
			fmt.Sprint("%", sp.Search, "%"))
	}
	tx = tx.Joins("LEFT JOIN manager_group ON manager_user.group_id=manager_group.id")
	tx = tx.Count(&count)
	if tx.Error != nil {
		config.Log.Error(tx.Error)
		return nil, ErrCount
	}
	mus := []*ManagerUserList{}
	tx = db.Table("manager_user").Select("manager_user.*, manager_group.name AS group_name")
	if sp.Search != "" {
		tx = tx.Where("manager_user.username like ? or manager_group.name like ?",
			fmt.Sprint("%", sp.Search, "%"),
			fmt.Sprint("%", sp.Search, "%"))
	}
	tx = tx.Joins("LEFT JOIN manager_group ON manager_user.group_id=manager_group.id").
		Order("update_at desc").
		Offset((sp.PageNo - 1) * sp.PageSize).
		Limit(sp.PageSize).
		Find(&mus)
	if tx.Error != nil {
		config.Log.Error(tx.Error)
		return nil, ErrSearchDBData
	}
	return CalSplitPage(sp, count, mus), Success
}

func PostManagerUser(user *UserSessionInfo, mu *ManagerUser) *BriefMessage {
	db := dbs.DBObj.GetGoRM()
	if db == nil {
		config.Log.Error(InternalGetBDInstanceErr)
		return ErrDataBase
	}
	mu.ID = 0
	mu.UpdateAt = time.Now()
	mu.Salt = uuid.NewString()
	mu.Password = utils.CreateHashword(mu.Password, mu.Salt)
	mu.CreateAt = time.Now()
	mu.UpdateBy = user.Username
	tx := db.Table("manager_user").Create(&mu)
	if tx.Error != nil {
		config.Log.Error(tx.Error)
		return ErrCreateDBData
	}
	return Success
}

func PutManagerUser(user *UserSessionInfo, mu *ManagerUser) *BriefMessage {
	db := dbs.DBObj.GetGoRM()
	if db == nil {
		config.Log.Error(InternalGetBDInstanceErr)
		return ErrDataBase
	}
	userInfo := ManagerUser{}
	if err := db.Table("manager_user").Where("id=?", mu.ID).Find(&userInfo).Error; err != nil {
		config.Log.Error(err)
		return ErrUserNotExist
	}
	tx := db.Table("manager_user").
		Where("id=?", mu.ID).
		Update("username", mu.UserName).
		Update("nice_name", mu.NiceName).
		Update("phone", mu.Phone).
		Update("group_id", mu.GroupID).
		// Update("enabled", mg.Enabled).
		Update("update_at", time.Now()).
		Update("update_by", user.Username)
	if userInfo.Password != mu.Password {
		mu.Salt = uuid.NewString()
		mu.Password = utils.CreateHashword(mu.Password, mu.Salt)
		tx.Update("password", mu.Password).Update("salt", mu.Salt)
	}
	if tx.Error != nil {
		config.Log.Error(tx.Error)
		return ErrUpdateData
	}
	return Success
}

func DeleteManagerUser(muID int64) *BriefMessage {
	db := dbs.DBObj.GetGoRM()
	if db == nil {
		config.Log.Error(InternalGetBDInstanceErr)
		return ErrDataBase
	}
	tx := db.Table("manager_user").
		Where("id=?", muID).Delete(nil)
	if tx.Error != nil {
		config.Log.Error(tx.Error)
		return ErrDelData
	}
	return Success
}

func PutManagerUserStatus(user *UserSessionInfo, info *EnabledInfo) *BriefMessage {
	db := dbs.DBObj.GetGoRM()
	if db == nil {
		config.Log.Error(InternalGetBDInstanceErr)
		return ErrDataBase
	}
	tx := db.Table("manager_user").
		Where("id=?", info.ID).
		Update("enabled", info.Enabled).
		Update("update_at", time.Now()).
		Update("update_by", user.Username)
	if tx.Error != nil {
		config.Log.Error(tx.Error)
		return ErrUpdateData
	}
	return Success
}

type ChangePasswordInfo struct {
	ID      int    `json:"id" form:"id"`
	OldPwd  string `json:"old_pwd" form:"old_pwd"`
	NewPwd1 string `json:"new_pwd1" form:"new_pwd1"`
	NewPwd2 string `json:"new_pwd2" form:"new_pwd2"`
}

func PostUserPassword(user *UserSessionInfo, cpi *ChangePasswordInfo) *BriefMessage {
	if cpi.OldPwd == "" {
		return ErrPostData
	}
	if cpi.NewPwd1 != cpi.NewPwd2 {
		return ErrPasswordSame
	}
	db := dbs.DBObj.GetGoRM()
	if db == nil {
		config.Log.Error(InternalGetBDInstanceErr)
		return ErrDataBase
	}
	userInfo := ManagerUser{}
	if err := db.Table("manager_user").Where("id=?", cpi.ID).Find(&userInfo).Error; err != nil {
		config.Log.Error(err)
		return ErrUserNotExist
	}
	checkPassword := utils.CreateHashword(cpi.OldPwd, userInfo.Salt)
	if checkPassword != userInfo.Password {
		return ErrPwdNotMatch
	}
	salt := uuid.NewString()
	newPassword := utils.CreateHashword(cpi.NewPwd1, salt)
	tx := db.Table("manager_user").
		Where("id=?", cpi.ID).
		Update("password", newPassword).
		Update("salt", salt).
		Update("update_at", time.Now()).
		Update("update_by", user.Username)
	if tx.Error != nil {
		config.Log.Error(tx.Error)
		return ErrUpdateData
	}
	return Success
}

func Login(ui *UserLogInfo, ipaddr string) (*ManagerUserDetail, *BriefMessage) {
	db := dbs.DBObj.GetGoRM()
	if db == nil {
		config.Log.Error(InternalGetBDInstanceErr)
		return nil, ErrDataBase
	}
	u := ManagerUserDetail{}
	if err := db.Table("manager_user").
		Select("manager_user.*, manager_group.enabled AS group_enabled, manager_group.name as group_name ").
		Joins("LEFT JOIN manager_group ON manager_user.group_id=manager_group.id").
		Where("username=?", ui.Username).
		First(&u).Error; err != nil {
		return nil, ErrUserNotExist
	}
	if !u.Enabled {
		return nil, ErrUserDisabled
	}
	if !u.GroupEnabled {
		return nil, ErrGroupDisabled
	}
	password := utils.CreateHashword(ui.Password, u.Salt)
	if u.Password != password {
		return nil, ErrPassword
	}
	ss := &Session{
		Token:    uuid.NewString(),
		UserID:   u.ID,
		IPAddr:   ipaddr,
		UpdateAt: time.Now(),
	}
	u.Session = ss
	UpdateSession(ss)
	// SSObj.Set(ss.Token, &u)
	return &u, Success
}

func Register(r *RegisterInfo) *BriefMessage {
	u := ManagerUser{
		UserName: r.Username,
		NiceName: r.NiceName,
		Password: r.Password,
		Phone:    r.Phone,
		GroupID:  r.GroupID,
	}
	user := &UserSessionInfo{
		Username: "anonymous",
	}
	return PostManagerUser(user, &u)
}

func LoadUserEnabled() ([]*ManagerUserDetail, *BriefMessage) {
	db := dbs.DBObj.GetGoRM()
	if db == nil {
		config.Log.Error(InternalGetBDInstanceErr)
		return nil, ErrDataBase
	}
	us := []*ManagerUserDetail{}
	if err := db.Table("manager_user").
		Select("manager_user.*, manager_group.enabled AS group_enabled, manager_group.name as group_name ").
		Joins("LEFT JOIN manager_group ON manager_user.group_id=manager_group.id").
		Where("manager_user.enabled=1 and manager_group.enabled=1").
		Find(&us).
		Error; err != nil {
		return nil, ErrSearchDBData
	}
	return us, Success
}

func LoadSession(uids []int) ([]*Session, *BriefMessage) {
	db := dbs.DBObj.GetGoRM()
	if db == nil {
		config.Log.Error(InternalGetBDInstanceErr)
		return nil, ErrDataBase
	}
	ss := []*Session{}
	tx := db.Table("session").Where("user_id in (?)", uids).Find(&ss)
	if tx.Error != nil {
		config.Log.Error(tx.Error)
		return nil, ErrSearchDBData
	}
	return ss, Success
}

func LoadSessionByToken(s string) (*Session, *BriefMessage) {
	db := dbs.DBObj.GetGoRM()
	if db == nil {
		config.Log.Error(InternalGetBDInstanceErr)
		return nil, ErrDataBase
	}
	userSession := Session{}
	tx := db.Table("session").Where("token=?", s).Find(&userSession)
	if tx.Error != nil {
		config.Log.Error(tx.Error)
		return nil, ErrSearchDBData
	}
	return &userSession, Success
}

func UpdateSession(s *Session) *BriefMessage {
	db := dbs.DBObj.GetGoRM()
	if db == nil {
		config.Log.Error(InternalGetBDInstanceErr)
		return ErrDataBase
	}
	tx := db.Table("session").Exec("INSERT INTO session (`id`, `token`, `ipaddr`, `user_id`, `update_at`) VALUES " +
		fmt.Sprintf("(%d, '%s', '%s', %d, '%s')", s.ID, s.Token, s.IPAddr, s.UserID, s.UpdateAt.Format("2006-01-02 15:04:05")) +
		" ON DUPLICATE KEY UPDATE `token`=VALUES(token),`user_id`=VALUES(user_id), `ipaddr`=VALUES(ipaddr), `update_at`=VALUES(update_at) ")
	if tx.Error != nil {
		config.Log.Error(tx.Error)
		return ErrCreateDBData
	}
	return Success
}

func DeleteSession(user *UserSessionInfo) *BriefMessage {
	db := dbs.DBObj.GetGoRM()
	if db == nil {
		config.Log.Error(InternalGetBDInstanceErr)
		return ErrDataBase
	}
	tx := db.Table("session").Where("user_id=?", user.UserID).Delete(nil)
	if tx.Error != nil {
		config.Log.Error(tx.Error)
		return ErrDelData
	}
	return Success
}

func Logout(user *UserSessionInfo) *BriefMessage {
	DeleteSession(user)
	// SSObj.Del(token)
	return Success
}

type AdminParams struct {
	ParamName  string `json:"param_name" gorm:"column:param_name"`
	ParamValue string `json:"param_value" gorm:"column:param_value"`
}

func GetManagerSetting() (map[string]string, *BriefMessage) {
	db := dbs.DBObj.GetGoRM()
	if db == nil {
		config.Log.Error(InternalGetBDInstanceErr)
		return nil, ErrDataBase
	}
	params := []AdminParams{}
	tx := db.Table("manager_set").Find(&params)
	if tx.Error != nil {
		config.Log.Error(tx.Error)
		return nil, ErrDelData
	}
	result := map[string]string{}
	for _, o := range params {
		result[o.ParamName] = o.ParamValue
	}
	return result, Success
}

func PutManagerSetting(params map[string]string) *BriefMessage {
	db := dbs.DBObj.GetGoRM()
	if db == nil {
		config.Log.Error(InternalGetBDInstanceErr)
		return ErrDataBase
	}
	for key, value := range params {
		tx := db.Table("manager_set").
			Where("param_name", key).
			Update("param_value", value)
		if tx.Error != nil {
			config.Log.Error(tx.Error)
			return ErrUpdateData
		}
	}
	return Success
}

type LogLevelSetting struct {
	ID       int    `json:"id" gorm:"column:id"`
	Level    string `json:"level" gorm:"column:level"`
	Label    string `json:"label" gorm:"column:label"`
	Selected bool   `json:"selected" gorm:"column:selected"`
}

func GetSystemRecodeSetting() ([]LogLevelSetting, *BriefMessage) {
	db := dbs.DBObj.GetGoRM()
	if db == nil {
		config.Log.Error(InternalGetBDInstanceErr)
		return nil, ErrDataBase
	}
	recodes := []LogLevelSetting{}
	tx := db.Table("log_setting").Find(&recodes)
	if tx.Error != nil {
		config.Log.Error(tx.Error)
		return nil, ErrSearchDBData
	}
	return recodes, Success
}

func PutSystemRecodeSetting(ids []int) *BriefMessage {
	db := dbs.DBObj.GetGoRM()
	if db == nil {
		config.Log.Error(InternalGetBDInstanceErr)
		return ErrDataBase
	}
	tErr := db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Table("log_setting").Where("1=1").Update("selected", false).Error; err != nil {
			config.Log.Error(err)
			return err
		}
		if len(ids) != 0 {
			if err := tx.Table("log_setting").Where("id in ?", ids).Update("selected", true).Error; err != nil {
				config.Log.Error(err)
				return err
			}
		}
		return nil
	})
	if tErr != nil {
		return ErrUpdateData
	}
	go OO.FlushLevel()
	return Success
}

func ClearSystemLog() *BriefMessage {
	db := dbs.DBObj.GetGoRM()
	if db == nil {
		config.Log.Error(InternalGetBDInstanceErr)
		return ErrDataBase
	}
	tx := db.Table("system_log").Where("1=1").Delete(nil)
	if tx.Error != nil {
		config.Log.Error(tx.Error)
		return ErrDelData
	}
	return Success
}

func GetSystemLog(sp *SplitPage) (*ResSplitPage, *BriefMessage) {
	db := dbs.DBObj.GetGoRM()
	if db == nil {
		config.Log.Error(InternalGetBDInstanceErr)
		return nil, ErrDataBase
	}
	var count int64
	var likeSql string
	tx := db.Table("system_log")
	if sp.Search != "" {
		like := `'%` + sp.Search + `%'`
		likeSql = fmt.Sprintf("username like %s "+
			"or ipaddr like %s "+
			"or operate_content like %s "+
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
	tx = db.Table("system_log")
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
