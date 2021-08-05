package models

import (
	"fmt"
	"pro_cfg_manager/config"
	"pro_cfg_manager/dbs"
	"sort"
	"strconv"

	"gorm.io/gorm"
)

type Priv struct {
	ID              int    `json:"id" gorm:"column:id"`
	PageName        string `json:"page_name" gorm:"column:page_name"`
	PageNiceName    string `json:"page_nice_name" gorm:"column:page_nice_name"`
	SubPageName     string `json:"sub_page_name" gorm:"column:sub_page_name"`
	SubPageNiceName string `json:"sub_page_nice_name" gorm:"column:sub_page_nice_name"`
	FuncName        string `json:"func_name" gorm:"column:func_name"`
	FuncNiceName    string `json:"func_nice_name" gorm:"column:func_nice_name"`
}

type GroupPriv struct {
	GroupID   int    `json:"group_id" gorm:"column:group_id"`
	GroupName string `json:"group_name" gorm:"column:group_name"`
	Priv
}

type FuncInfo struct {
	ID           int    `json:"id"`
	FuncName     string `json:"func_name"`
	FuncNiceName string `json:"func_nice_name"`
	Checked      bool   `json:"checked"`
}

type ItemPriv struct {
	PageName        string      `json:"page_name"`
	PageNiceName    string      `json:"page_nice_name"`
	SubPageName     string      `json:"sub_page_name"`
	SubPageNiceName string      `json:"sub_page_nice_name"`
	FuncList        []*FuncInfo `json:"func_list" `
}

type GetPrivInfo struct {
	GroupName string `json:"group_name" form:"group_name"`
	GroupID   int    `json:"group_id" form:"group_id"`
}

type TableGroupPriv struct {
	ID      int `json:"id" gorm:"column:id"`
	GroupID int `json:"group_id" gorm:"column:group_id"`
	FuncID  int `json:"func_id" gorm:"column:func_id"`
}

// 接口即使返回一样的数据，在前端使用的时候也用不一样的请求地址
func CheckPriv(u *UserSessionInfo, pageName, subPageName, funcName string) *BriefMessage {
	db := dbs.DBObj.GetGoRM()
	if db == nil {
		config.Log.Error(InternalGetBDInstanceErr)
		return ErrDataBase
	}
	var count int64
	tx := db.Table("group_priv").
		Select("COUNT(*) AS count").
		Joins("LEFT JOIN page_function ON group_priv.func_id=page_function.id").
		Where("group_priv.group_id=? AND (page_function.page_name=? AND page_function.sub_page_name=? AND page_function.func_name=?)",
			u.GroupID, pageName, subPageName, funcName).Count(&count)
	if tx.Error != nil {
		config.Log.Error(tx.Error)
		return ErrSearchDBData
	}
	if count <= 0 {
		return ErrNoPrivRequest
	}
	return Success
}

func GetGroupPriv(gInfo *GetPrivInfo) ([]*ItemPriv, *BriefMessage) {
	db := dbs.DBObj.GetGoRM()
	if db == nil {
		config.Log.Error(InternalGetBDInstanceErr)
		return nil, ErrDataBase
	}
	gps := []*GroupPriv{}
	tx := db.Table("page_function").
		Select("page_function.*, manager_group.id AS group_id, manager_group.name AS group_name").
		Joins("LEFT JOIN (SELECT * FROM group_priv WHERE group_id=?) AS group_priv ON page_function.id=group_priv.func_id ", gInfo.GroupID).
		Joins("LEFT JOIN manager_group ON group_priv.group_id=manager_group.id").
		Order("page_function.page_name, page_function.sub_page_name").
		Find(&gps)
	if tx.Error != nil {
		return nil, ErrSearchDBData
	}
	ipm := map[string]*ItemPriv{}
	for _, gp := range gps {
		page_join_name := fmt.Sprintf("%s_%s", gp.PageName, gp.SubPageName)
		obj, ok := ipm[page_join_name]
		if !ok {
			ipm[page_join_name] = &ItemPriv{
				PageName:        gp.PageName,
				PageNiceName:    gp.PageNiceName,
				SubPageName:     gp.SubPageName,
				SubPageNiceName: gp.SubPageNiceName,
				FuncList:        []*FuncInfo{},
			}
			obj = ipm[page_join_name]
		}
		fi := &FuncInfo{
			ID:           gp.ID,
			FuncName:     gp.FuncName,
			FuncNiceName: gp.FuncNiceName,
		}
		if gp.GroupName == "" {
			fi.Checked = false
		} else {
			fi.Checked = true
		}
		obj.FuncList = append(obj.FuncList, fi)
	}
	pageJoinNameSlice := sort.StringSlice{}
	for key, _ := range ipm {
		pageJoinNameSlice = append(pageJoinNameSlice, key)
	}
	sort.Strings(pageJoinNameSlice)
	ips := []*ItemPriv{}
	for _, ipName := range pageJoinNameSlice {
		ips = append(ips, ipm[ipName])
	}
	return ips, Success
}

func PutGroupPriv(privInfo []*ItemPriv, gInfo *GetPrivInfo) *BriefMessage {
	db := dbs.DBObj.GetGoRM()
	if db == nil {
		config.Log.Error(InternalGetBDInstanceErr)
		return ErrDataBase
	}
	tgp := []*TableGroupPriv{}
	for _, p := range privInfo {
		for _, f := range p.FuncList {
			if !f.Checked {
				continue
			}
			tgp = append(tgp, &TableGroupPriv{
				GroupID: gInfo.GroupID,
				FuncID:  f.ID,
			})
		}
	}
	tErr := db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Table("group_priv").Where("group_id=?", gInfo.GroupID).Delete(nil).Error; err != nil {
			config.Log.Error(err)
			return err
		}
		if err := tx.Table("group_priv").Create(tgp).Error; err != nil {
			config.Log.Error(err)
			return err
		}
		return nil
	})
	if tErr != nil {
		return ErrUpdateData
	}
	return Success
}

type UserGroupMember struct {
	ID   int    `json:"id" gorm:"column:id"`
	Name string `json:"username" gorm:"column:username"`
}

func GetManagerUsersList() ([]*UserGroupMember, *BriefMessage) {
	db := dbs.DBObj.GetGoRM()
	if db == nil {
		config.Log.Error(InternalGetBDInstanceErr)
		return nil, ErrDataBase
	}
	ugm := []*UserGroupMember{}
	tx := db.Table("manager_user").Select("`id`, `username`").Find(&ugm)
	if tx.Error != nil {
		config.Log.Error(tx.Error)
		return nil, ErrSearchDBData
	}
	return ugm, Success
}

func GetManagerGroupMember(gInfo *GetPrivInfo) ([]*UserGroupMember, *BriefMessage) {
	db := dbs.DBObj.GetGoRM()
	if db == nil {
		config.Log.Error(InternalGetBDInstanceErr)
		return nil, ErrDataBase
	}
	ugm := []*UserGroupMember{}
	tx := db.Table("manager_user").Select("`id`, `username`").Where("group_id=?", gInfo.GroupID).Find(&ugm)
	if tx.Error != nil {
		config.Log.Error(tx.Error)
		return nil, ErrSearchDBData
	}
	return ugm, Success
}

func PutManagerGroupMember(gInfo *GetPrivInfo, userList []int) *BriefMessage {
	db := dbs.DBObj.GetGoRM()
	if db == nil {
		config.Log.Error(InternalGetBDInstanceErr)
		return ErrDataBase
	}
	tErr := db.Transaction(func(tx *gorm.DB) error {
		param := AdminParams{}
		if err := tx.Table("manager_set aaa").Where("param_name='default_group'").Find(&param).Error; err != nil {
			config.Log.Error(err)
			return err
		}
		defaultGroupValue := int64(0)
		if param.ParamName == "default_group" {
			value, err := strconv.ParseInt(param.ParamValue, 10, 0)
			if err != nil {
				config.Log.Error(err)
				return err
			}
			defaultGroupValue = value
		}
		config.Log.Warnf("defaultGroupValue => %v", defaultGroupValue)
		if err := tx.Table("manager_user").
			Where("group_id=?", gInfo.GroupID).
			Update("group_id", defaultGroupValue).Error; err != nil {
			config.Log.Error(err)
			return err
		}
		if err := tx.Table("manager_user").
			Where("id in (?)", userList).
			Update("group_id", gInfo.GroupID).Error; err != nil {
			config.Log.Error(err)
			return err
		}
		return nil
	})
	if tErr != nil {
		return ErrUpdateData
	}
	return Success
}

type UserMenu struct {
	FuncName string `json:"func_name" gorm:"column:func_name"`
}

func GetUserMenuPriv(user *UserSessionInfo) (map[string]bool, *BriefMessage) {
	db := dbs.DBObj.GetGoRM()
	if db == nil {
		config.Log.Error(InternalGetBDInstanceErr)
		return nil, ErrDataBase
	}
	ums := []*UserMenu{}
	tx := db.Table("group_priv").
		Select("func_name").
		Joins("LEFT JOIN page_function ON group_priv.func_id=page_function.id").
		Where("group_priv.group_id=? AND page_function.page_name='person' AND page_function.func_name IN ('show_menu_prometheus_cfg_manager', 'show_menu_administrator_cfg_manager')",
			user.GroupID).Find(&ums)
	if tx.Error != nil {
		config.Log.Error(tx.Error)
		return nil, ErrSearchDBData
	}
	menuDataResp := map[string]bool{}
	for _, um := range ums {
		menuDataResp[um.FuncName] = true
	}
	return menuDataResp, Success
}
