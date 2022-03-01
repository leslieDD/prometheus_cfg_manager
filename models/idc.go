package models

import (
	"fmt"
	"pro_cfg_manager/config"
	"pro_cfg_manager/dbs"
	"time"
)

type IDC struct {
	ID       int       `json:"id" gorm:"column:id"`
	Label    string    `json:"label" gorm:"column:label"`
	Enabled  bool      `json:"enabled" gorm:"column:enabled"`
	UpdateAt time.Time `json:"update_at" gorm:"column:update_at"`
	UpdateBy string    `json:"update_by" gorm:"column:update_by"`
}

type Line struct {
	ID       int       `json:"id" gorm:"column:id"`
	Label    string    `json:"label" gorm:"column:label"`
	Enabled  bool      `json:"enabled" gorm:"column:enabled"`
	UpdateAt time.Time `json:"update_at" gorm:"column:update_at"`
	UpdateBy string    `json:"update_by" gorm:"column:update_by"`
	IDCID    int       `json:"idc_id" gorm:"column:idc_id"`
}

type IPAddrsPool struct {
	ID      int    `json:"id" gorm:"column:id"`
	LineID  int    `json:"line_id" gorm:"column:line_id"`
	Ipaddrs string `json:"ipaddrs" gorm:"column:ipaddrs"`
	// Enabled  bool      `json:"enabled" gorm:"column:enabled"`
	UpdateAt time.Time `json:"update_at" gorm:"column:update_at"`
	UpdateBy string    `json:"update_by" gorm:"column:update_by"`
}

type NewIDC struct {
	ID    int    `json:"id" gorm:"column:id"`
	Label string `json:"label" gorm:"column:label"`
}

type NewLine struct {
	ID    int    `json:"id" gorm:"column:id"`
	IDCID int    `json:"idc_id" gorm:"column:idc_id"`
	Label string `json:"label" gorm:"column:label"`
}

type NewPool struct {
	ID      int    `json:"id" gorm:"column:id"`
	LineID  int    `json:"line_id" gorm:"column:line_id"`
	Ipaddrs string `json:"ipaddrs" gorm:"column:ipaddrs"`
}

type idcTree struct {
	TreeID   int         `json:"tree_id"`
	TreeType string      `json:"tree_type"`
	Children []*lineTree `json:"children"`
	IDC
}

type lineTree struct {
	TreeID   int    `json:"tree_id"`
	TreeType string `json:"tree_type"`
	Line
}

func GetIDC(id *OnlyID) (*IDC, *BriefMessage) {
	db := dbs.DBObj.GetGoRM()
	if db == nil {
		config.Log.Error(InternalGetBDInstanceErr)
		return nil, ErrDataBase
	}
	idc := IDC{}
	tx := db.Table("idc").Where("id", id.ID).Find(&idc)
	if tx.Error != nil {
		config.Log.Error(tx.Error)
		return nil, ErrSearchDBData
	}
	return &idc, Success
}

func GetIDCs() {

}

func PostIDC(user *UserSessionInfo, newIDC *NewIDC) *BriefMessage {
	db := dbs.DBObj.GetGoRM()
	if db == nil {
		config.Log.Error(InternalGetBDInstanceErr)
		return ErrDataBase
	}
	idc := IDC{
		ID:       0,
		Enabled:  true,
		Label:    newIDC.Label,
		UpdateAt: time.Now(),
		UpdateBy: user.Username,
	}
	tx := db.Table("idc").Create(&idc)
	if tx.Error != nil {
		config.Log.Error(tx.Error)
		return ErrCreateDBData
	}
	return Success
}

func PutIDC(user *UserSessionInfo, newIDC *NewIDC) *BriefMessage {
	db := dbs.DBObj.GetGoRM()
	if db == nil {
		config.Log.Error(InternalGetBDInstanceErr)
		return ErrDataBase
	}
	tx := db.Table("idc").Where("id", newIDC.ID).
		Update("label", newIDC.Label).
		Update("enabled", true).
		Update("update_at", time.Now()).
		Update("update_by", user.Username)
	if tx.Error != nil {
		config.Log.Error(tx.Error)
		return ErrUpdateData
	}
	return Success
}

func DelIDC(id *OnlyID) *BriefMessage {
	db := dbs.DBObj.GetGoRM()
	if db == nil {
		config.Log.Error(InternalGetBDInstanceErr)
		return ErrDataBase
	}
	tx := db.Table("idc").Where("id", id.ID).Delete(nil)
	if tx.Error != nil {
		config.Log.Error(tx.Error)
		return ErrDelData
	}
	return Success
}

func GetLine(id *OnlyID) (*Line, *BriefMessage) {
	db := dbs.DBObj.GetGoRM()
	if db == nil {
		config.Log.Error(InternalGetBDInstanceErr)
		return nil, ErrDataBase
	}
	line := Line{}
	tx := db.Table("line").Where("id", id.ID).Find(&line)
	if tx.Error != nil {
		config.Log.Error(tx.Error)
		return nil, ErrSearchDBData
	}
	return &line, Success
}

func GetLines() {}

func PostLine(user *UserSessionInfo, newLine *NewLine) *BriefMessage {
	db := dbs.DBObj.GetGoRM()
	if db == nil {
		config.Log.Error(InternalGetBDInstanceErr)
		return ErrDataBase
	}
	line := Line{
		ID:       0,
		Enabled:  true,
		Label:    newLine.Label,
		IDCID:    newLine.IDCID,
		UpdateAt: time.Now(),
		UpdateBy: user.Username,
	}
	tx := db.Table("line").Create(&line)
	if tx.Error != nil {
		config.Log.Error(tx.Error)
		return ErrCreateDBData
	}
	return Success
}

func PutLine(user *UserSessionInfo, newLine *NewLine) *BriefMessage {
	db := dbs.DBObj.GetGoRM()
	if db == nil {
		config.Log.Error(InternalGetBDInstanceErr)
		return ErrDataBase
	}
	tx := db.Table("line").Where("id", newLine.ID).
		// Update("idc_id", 0)
		Update("label", newLine.Label).
		Update("enabled", true).
		Update("update_at", time.Now()).
		Update("update_by", user.Username)
	if tx.Error != nil {
		config.Log.Error(tx.Error)
		return ErrUpdateData
	}
	return Success
}

func DelLine(id *OnlyID) *BriefMessage {
	db := dbs.DBObj.GetGoRM()
	if db == nil {
		config.Log.Error(InternalGetBDInstanceErr)
		return ErrDataBase
	}
	tx := db.Table("line").Where("id", id.ID).Delete(nil)
	if tx.Error != nil {
		config.Log.Error(tx.Error)
		return ErrDelData
	}
	return Success
}

func GetLineIpAddrs(id *OnlyID) (*IPAddrsPool, *BriefMessage) {
	db := dbs.DBObj.GetGoRM()
	if db == nil {
		config.Log.Error(InternalGetBDInstanceErr)
		return nil, ErrDataBase
	}
	pool := IPAddrsPool{}
	tx := db.Table("pool").Where("id", id.ID).Find(&pool)
	if tx.Error != nil {
		config.Log.Error(tx.Error)
		return nil, ErrSearchDBData
	}
	return &pool, Success
}

func PutLineIpAddrs(user *UserSessionInfo, newPool *NewPool) *BriefMessage {
	db := dbs.DBObj.GetGoRM()
	if db == nil {
		config.Log.Error(InternalGetBDInstanceErr)
		return ErrDataBase
	}
	sql := fmt.Sprintf("replace into pool(id, line_id, ipaddrs, update_at, update_by) values(%d, %d, '%s', '%s', '%s');",
		newPool.ID, newPool.LineID, newPool.Ipaddrs, time.Now().Format("2006-02-01 15:04:05"), user.Username)
	tx := db.Table("pool").Exec(sql)
	if tx.Error != nil {
		config.Log.Error(tx.Error)
		return ErrUpdateData
	}
	return Success
}

func GetIDCTree() (interface{}, *BriefMessage) {
	db := dbs.DBObj.GetGoRM()
	if db == nil {
		config.Log.Error(InternalGetBDInstanceErr)
		return nil, ErrDataBase
	}
	// IDC
	idcs := []*IDC{}
	tx := db.Table("idc").Find(&idcs)
	if tx.Error != nil {
		config.Log.Error(tx.Error)
		return nil, ErrSearchDBData
	}
	// Line
	lines := []*Line{}
	tx = db.Table("line").Find(&lines)
	if tx.Error != nil {
		config.Log.Error(tx.Error)
		return nil, ErrSearchDBData
	}

	maxIDCID := 0
	for _, i := range idcs {
		if maxIDCID < i.ID {
			maxIDCID = i.ID
		}
	}

	lineTreeResp := map[int][]*lineTree{}
	for _, line := range lines {
		_, ok := lineTreeResp[line.IDCID]
		if !ok {
			lineTreeResp[line.IDCID] = []*lineTree{}
		}
		lineTreeResp[line.IDCID] = append(lineTreeResp[line.IDCID], &lineTree{
			TreeID:   line.ID + maxIDCID, // 这样ID就不会有重复的，而且ID也不会变更
			TreeType: "idc",
			Line:     *line,
		})
	}

	idcTreeResp := []*idcTree{}
	for _, idc := range idcs {
		node := &idcTree{
			TreeID:   idc.ID,
			TreeType: "idc",
			IDC:      *idc,
		}
		obj, ok := lineTreeResp[idc.ID]
		if ok {
			node.Children = obj
		} else {
			node.Children = []*lineTree{}
		}
		idcTreeResp = append(idcTreeResp, node)
	}
	return idcTreeResp, Success
}
