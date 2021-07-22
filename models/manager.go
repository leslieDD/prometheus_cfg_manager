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
}

type ManagerUser struct {
	ID       int       `json:"id" gorm:"column:id"`
	UserName string    `json:"username" gorm:"column:username"`
	Password string    `json:"password" gorm:"column:password"`
	Phone    string    `json:"phone" gorm:"column:phone"`
	Salt     string    `json:"-" gorm:"column:salt"`
	GroupID  int       `json:"group_id" gorm:"column:group_id"`
	Enabled  bool      `json:"enabled" gorm:"column:enabled"`
	UpdateAt time.Time `json:"update_at" gorm:"column:update_at"`
}

type ManagerUserList struct {
	GroupName string `json:"group_name" gorm:"column:group_name"`
	ManagerUser
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
	mgs := []*ManagerGroup{}
	tx = db.Table("manager_group")
	if sp.Search != "" {
		tx = tx.Where("name like ?", fmt.Sprint("%", sp.Search, "%"))
	}
	tx = tx.Order("update_at desc").
		Offset((sp.PageNo - 1) * sp.PageSize).
		Limit(sp.PageSize).
		Find(&mgs)
	if tx.Error != nil {
		config.Log.Error(tx.Error)
		return nil, ErrSearchDBData
	}
	return CalSplitPage(sp, count, mgs), Success
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

func PostManagerGroup(mg *ManagerGroup) *BriefMessage {
	db := dbs.DBObj.GetGoRM()
	if db == nil {
		config.Log.Error(InternalGetBDInstanceErr)
		return ErrDataBase
	}
	mg.ID = 0
	mg.UpdateAt = time.Now()
	tx := db.Table("manager_group").Create(&mg)
	if tx.Error != nil {
		config.Log.Error(tx.Error)
		return ErrCreateDBData
	}
	return Success
}

func PutManagerGroup(mg *ManagerGroup) *BriefMessage {
	db := dbs.DBObj.GetGoRM()
	if db == nil {
		config.Log.Error(InternalGetBDInstanceErr)
		return ErrDataBase
	}
	tx := db.Table("manager_group").
		Where("id=?", mg.ID).
		Update("name", mg.Name).
		// Update("enabled", mg.Enabled).
		Update("update_at", time.Now())
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
	tx := db.Table("manager_group").
		Where("id=?", mgID).Delete(nil)
	if tx.Error != nil {
		config.Log.Error(tx.Error)
		return ErrDelData
	}
	return Success
}

func PutManagerGroupStatus(info *EnabledInfo) *BriefMessage {
	db := dbs.DBObj.GetGoRM()
	if db == nil {
		config.Log.Error(InternalGetBDInstanceErr)
		return ErrDataBase
	}
	tx := db.Table("manager_group").
		Where("id=?", info.ID).
		Update("enabled", info.Enabled).
		Update("update_at", time.Now())
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
	tx = db.Table("manager_user")
	if sp.Search != "" {
		tx = tx.Where("username like ?", fmt.Sprint("%", sp.Search, "%"))
	}
	tx = tx.Count(&count)
	if tx.Error != nil {
		config.Log.Error(tx.Error)
		return nil, ErrCount
	}
	mus := []*ManagerUserList{}
	tx = db.Table("manager_user").Select("manager_user.*, manager_group.name AS group_name")
	if sp.Search != "" {
		tx = tx.Where("username like ?", fmt.Sprint("%", sp.Search, "%"))
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

func PostManagerUser(mu *ManagerUser) *BriefMessage {
	db := dbs.DBObj.GetGoRM()
	if db == nil {
		config.Log.Error(InternalGetBDInstanceErr)
		return ErrDataBase
	}
	mu.ID = 0
	mu.UpdateAt = time.Now()
	mu.Salt = uuid.NewString()
	mu.Password = utils.CreateHashword(mu.Password, mu.Salt)
	tx := db.Table("manager_user").Create(&mu)
	if tx.Error != nil {
		config.Log.Error(tx.Error)
		return ErrCreateDBData
	}
	return Success
}

func PutManagerUser(mu *ManagerUser) *BriefMessage {
	db := dbs.DBObj.GetGoRM()
	if db == nil {
		config.Log.Error(InternalGetBDInstanceErr)
		return ErrDataBase
	}
	user := ManagerUser{}
	if err := db.Table("manager_user").Where("id=?", mu.ID).Find(&user).Error; err != nil {
		config.Log.Error(err)
		return ErrUserNotExist
	}
	tx := db.Table("manager_user").
		Where("id=?", mu.ID).
		Update("username", mu.UserName).
		Update("phone", mu.Phone).
		Update("group_id", mu.GroupID).
		// Update("enabled", mg.Enabled).
		Update("update_at", time.Now())
	if user.Password != mu.Password {
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

func PutManagerUserStatus(info *EnabledInfo) *BriefMessage {
	db := dbs.DBObj.GetGoRM()
	if db == nil {
		config.Log.Error(InternalGetBDInstanceErr)
		return ErrDataBase
	}
	tx := db.Table("manager_user").
		Where("id=?", info.ID).
		Update("enabled", info.Enabled).
		Update("update_at", time.Now())
	if tx.Error != nil {
		config.Log.Error(tx.Error)
		return ErrUpdateData
	}
	return Success
}
