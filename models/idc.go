package models

import (
	"bytes"
	"encoding/base64"
	"errors"
	"fmt"
	"math/big"
	"net"
	"pro_cfg_manager/config"
	"pro_cfg_manager/dbs"
	"pro_cfg_manager/utils"
	"strings"
	"time"

	"github.com/3th1nk/cidr"
	"github.com/xuri/excelize/v2"
	"gorm.io/gorm"
)

type IDC struct {
	ID       int       `json:"id" gorm:"column:id"`
	Label    string    `json:"label" gorm:"column:label"`
	Remark   string    `json:"remark" gorm:"column:remark"`
	Enabled  bool      `json:"enabled" gorm:"column:enabled"`
	Expand   bool      `json:"expand" gorm:"column:expand"`
	View     bool      `json:"view" gorm:"column:view"`
	UpdateAt time.Time `json:"update_at" gorm:"column:update_at"`
	UpdateBy string    `json:"update_by" gorm:"column:update_by"`
}

type Line struct {
	ID       int       `json:"id" gorm:"column:id"`
	Label    string    `json:"label" gorm:"column:label"`
	Enabled  bool      `json:"enabled" gorm:"column:enabled"`
	Expand   bool      `json:"expand" gorm:"column:expand"`
	View     bool      `json:"view" gorm:"column:view"`
	UpdateAt time.Time `json:"update_at" gorm:"column:update_at"`
	UpdateBy string    `json:"update_by" gorm:"column:update_by"`
	IDCID    int       `json:"idc_id" gorm:"column:idc_id"`
}

type LineForSearch struct {
	Ipaddrs string `json:"ipaddrs" gorm:"column:ipaddrs"`
	Line
}

type IPAddrsPool struct {
	ID       int       `json:"id" gorm:"column:id"`
	LineID   int       `json:"line_id" gorm:"column:line_id"`
	Ipaddrs  string    `json:"ipaddrs" gorm:"column:ipaddrs"`
	Remark   string    `json:"remark" gorm:"column:remark"`
	UpdateAt time.Time `json:"update_at" gorm:"column:update_at"`
	UpdateBy string    `json:"update_by" gorm:"column:update_by"`
}

type NewIDC struct {
	ID     int    `json:"id" gorm:"column:id"`
	Label  string `json:"label" gorm:"column:label"`
	Expand bool   `json:"expand" gorm:"column:expand"`
	View   bool   `json:"view" gorm:"column:view"`
}

type NewLine struct {
	ID     int    `json:"id" gorm:"column:id"`
	IDCID  int    `json:"idc_id" gorm:"column:idc_id"`
	Label  string `json:"label" gorm:"column:label"`
	Expand bool   `json:"expand" gorm:"column:expand"`
	View   bool   `json:"view" gorm:"column:view"`
}

type NewPool struct {
	ID      int    `json:"id" gorm:"column:id"`
	LineID  int    `json:"line_id" gorm:"column:line_id"`
	Ipaddrs string `json:"ipaddrs" gorm:"column:ipaddrs"`
	Remark  string `json:"remark" gorm:"column:remark"`
	Expand  bool   `json:"expand" gorm:"column:expand"`
	View    bool   `json:"view" gorm:"column:view"`
}

type idcTree struct {
	TreeID   int         `json:"tree_id"`
	TreeType string      `json:"tree_type"`
	Children []*lineTree `json:"children"`
	IDC
}

type idcTreeWithID struct {
	ID   int        `json:"id"`
	Tree []*idcTree `json:"tree"`
}

type lineTree struct {
	TreeID   int    `json:"tree_id"`
	TreeType string `json:"tree_type"`
	Line
}

type ExpandSwitchReq struct {
	ID     int  `json:"id"`
	Switch bool `json:"switch"`
}
type ViewSwitchReq struct {
	ID     int  `json:"id"`
	Switch bool `json:"switch"`
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
		Expand:   newIDC.Expand,
		View:     newIDC.View,
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
		Update("expand", newIDC.Expand).
		Update("view", newIDC.View).
		Update("update_at", time.Now()).
		Update("update_by", user.Username)
	if tx.Error != nil {
		config.Log.Error(tx.Error)
		return ErrUpdateData
	}
	return Success
}

func PutIDCRemark(user *UserSessionInfo, idc *IDC) *BriefMessage {
	db := dbs.DBObj.GetGoRM()
	if db == nil {
		config.Log.Error(InternalGetBDInstanceErr)
		return ErrDataBase
	}
	tx := db.Table("idc").Where("id", idc.ID).
		Update("remark", idc.Remark).
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
	lines := []*Line{}
	if err := db.Table("line").Where("idc_id", id.ID).Find(&lines).Error; err != nil {
		config.Log.Error(err)
		return ErrSearchDBData
	}
	if len(lines) != 0 {
		return ErrHaveLine
	}
	err := db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Table("idc").Where("id", id.ID).Delete(nil).Error; err != nil {
			config.Log.Error(err)
			return err
		}
		return nil
	})
	if err != nil {
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
		Expand:   newLine.Expand,
		View:     newLine.View,
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
		Update("expand", newLine.Expand).
		Update("view", newLine.View).
		Update("update_at", time.Now()).
		Update("update_by", user.Username)
	if tx.Error != nil {
		config.Log.Error(tx.Error)
		return ErrUpdateData
	}
	return Success
}

func DelLine(delReqParams *IdAndRm) *BriefMessage {
	db := dbs.DBObj.GetGoRM()
	if db == nil {
		config.Log.Error(InternalGetBDInstanceErr)
		return ErrDataBase
	}
	err := db.Transaction(func(tx *gorm.DB) error {
		line := Line{}
		if err := tx.Table("line").Where("id", delReqParams.ID).Find(&line).Error; err != nil {
			config.Log.Error(err)
			return err
		}
		if err := tx.Table("line").Where("id", delReqParams.ID).Delete(nil).Error; err != nil {
			config.Log.Error(err)
			return err
		}
		if delReqParams.RmAddrs {
			pool := IPAddrsPool{}
			if err := tx.Table("pool").Where("line_id", line.ID).Find(&pool).Error; err != nil {
				config.Log.Error(err)
				return err
			}
			ipLists := ParseIPAddrsFromString(pool.Ipaddrs)
			ipStr := []string{}
			for ip, _ := range ipLists {
				ipStr = append(ipStr, ip)
			}
			if bf := DeleteMachineUseAddr(ipStr); bf != Success {
				return errors.New(bf.String())
			}
		}
		if err := tx.Table("pool").Where("line_id", line.ID).Delete(nil).Error; err != nil {
			config.Log.Error(err)
			return err
		}
		return nil
	})
	if err != nil {
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
	tx := db.Table("pool").Where("line_id", id.ID).Find(&pool)
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
	items := map[string]struct{}{}
	for _, each := range strings.FieldsFunc(newPool.Ipaddrs, Split) {
		if strings.TrimSpace(each) == "" {
			continue
		}
		// items = append(items, strings.TrimSpace(each))
		items[strings.TrimSpace(each)] = struct{}{}
	}
	vaildItems := []string{}
	for item, _ := range items {
		currIP := strings.TrimSpace(item)
		if currIP == "" {
			continue
		}
		if strings.Contains(currIP, "/") {
			_, err := cidr.ParseCIDR(currIP)
			if err != nil {
				config.Log.Error(err)
				return ErrIPAddrPost
			}
			vaildItems = append(vaildItems, currIP)
		} else if strings.Contains(currIP, "~") {
			fields := strings.Split(currIP, "~")
			if len(fields) != 2 {
				config.Log.Errorf("ip pool err: %s", currIP)
				return ErrIPAddrPost
			}
			_, err := utils.RangeBeginToEnd(strings.TrimSpace(fields[0]), strings.TrimSpace(fields[1]))
			if err != nil {
				config.Log.Error(err)
				return ErrIPAddrPost
			}
			vaildItems = append(vaildItems, currIP)
		} else if strings.Contains(currIP, "-") {
			fields := strings.Split(currIP, "-")
			if len(fields) != 2 {
				config.Log.Errorf("ip pool err: %s", currIP)
				return ErrIPAddrPost
			}
			_, err := utils.RangeBeginToEnd(strings.TrimSpace(fields[0]), strings.TrimSpace(fields[1]))
			if err != nil {
				config.Log.Error(err)
				return ErrIPAddrPost
			}
			vaildItems = append(vaildItems, currIP)
		} else {
			ipObj := net.ParseIP(currIP)
			if ipObj == nil {
				config.Log.Errorf("ip err: %s", currIP)
				return ErrIPAddrPost
			}
			vaildItems = append(vaildItems, currIP)
		}
	}
	if len(vaildItems) == 0 {
		return Success
	}
	postIpAddr := strings.Join(vaildItems, ";\n")
	sql := fmt.Sprintf("replace into pool(line_id, ipaddrs, remark, update_at, update_by) values(%d, '%s', '%s', '%s', '%s');",
		newPool.LineID, postIpAddr, newPool.Remark, time.Now().Format("2006-01-02 15:04:05"), user.Username)
	tx := db.Table("pool").Exec(sql)
	if tx.Error != nil {
		config.Log.Error(tx.Error)
		return ErrUpdateData
	}
	return Success
}

func GetIDCTree(search *SearchContent2) (*idcTreeWithID, *BriefMessage) {
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
	lines := []*LineForSearch{}
	tx = db.Raw(`select line.*, pool.ipaddrs from line left join pool on  line.id=pool.line_id;`).Find(&lines)
	if tx.Error != nil {
		config.Log.Error(tx.Error)
		return nil, ErrSearchDBData
	}

	linesMap := map[int]*LineForSearch{}
	for _, obj := range lines {
		linesMap[obj.ID] = obj
	}

	maxIDCID := 0
	for _, i := range idcs {
		if maxIDCID < i.ID {
			maxIDCID = i.ID
		}
	}

	finalID := maxIDCID // 供前端使用,

	lineTreeResp := map[int][]*lineTree{}
	for _, line := range lines {
		_, ok := lineTreeResp[line.IDCID]
		if !ok {
			lineTreeResp[line.IDCID] = []*lineTree{}
		}
		finalID = line.ID + maxIDCID
		lineTreeResp[line.IDCID] = append(lineTreeResp[line.IDCID], &lineTree{
			TreeID:   finalID, // 这样ID就不会有重复的，而且ID也不会变更
			TreeType: "line",
			Line: Line{
				ID:       line.ID,
				Label:    line.Label,
				Enabled:  line.Enabled,
				UpdateAt: line.UpdateAt,
				UpdateBy: line.UpdateBy,
				IDCID:    line.IDCID,
				Expand:   line.Expand,
				View:     line.View,
			},
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
		// 搜索规则，搜索IDC和Line，
		// 如果搜索的内容为空：
		//    直接返回全部数据
		// 如果搜索的内容不为空：
		//    先匹配IDC的Lable,匹配成功，刚再匹配此IDC中所有Line，
		//       如果Line中有匹配的，则返回匹配的，
		//       如果没有匹配的，则返回全部
		//    如果IDC中的Label没有匹配成功，也会再匹配此IDC中的所有Line
		//       如果Line中有匹配的，则返回匹配的，
		//       如果没有匹配的，则不返回
		if search.Content != "" {
			if strings.Contains(node.Label, search.Content) {
				matchLine := []*lineTree{}
				for _, eachLine := range obj {
					if strings.Contains(eachLine.Label, search.Content) {
						matchLine = append(matchLine, eachLine)
					} else if search.SearchIP && strings.Contains(linesMap[eachLine.ID].Ipaddrs, search.Content) {
						matchLine = append(matchLine, eachLine)
					}
				}
				if len(matchLine) != 0 {
					node.Children = matchLine
				}
			} else {
				matchLine := []*lineTree{}
				for _, eachLine := range obj {
					if strings.Contains(eachLine.Label, search.Content) {
						matchLine = append(matchLine, eachLine)
					} else if search.SearchIP && strings.Contains(linesMap[eachLine.ID].Ipaddrs, search.Content) {
						matchLine = append(matchLine, eachLine)
					}
				}
				if len(matchLine) != 0 {
					node.Children = matchLine
				} else {
					continue // 不返回数据
				}
			}
		}
		idcTreeResp = append(idcTreeResp, node)
	}
	respData := &idcTreeWithID{
		ID:   finalID,
		Tree: idcTreeResp,
	}
	return respData, Success
}

type idcXlsDB struct {
	ID        int    `json:"id" gorm:"column:id"`
	IPAddrs   string `json:"ipaddrs" gorm:"column:ipaddrs"`
	LineID    int    `json:"line_id" gorm:"column:line_id"`
	LineLabel string `json:"line_label" gorm:"column:line_label"`
	IDCID     int    `json:"idc_id" gorm:"column:idc_id"`
	IDCLabel  string `json:"idc_label" gorm:"column:idc_label"`
}

type XlsDatResp struct {
	Data string `json:"data"`
	Name string `json:"name"`
}

func GetIDCXls() (*XlsDatResp, *BriefMessage) {
	db := dbs.DBObj.GetGoRM()
	if db == nil {
		config.Log.Error(InternalGetBDInstanceErr)
		return nil, ErrDataBase
	}
	data := []*idcXlsDB{}
	tx := db.Raw(`SELECT pool.id, pool.ipaddrs, line.label AS line_label, idc.label AS idc_label, line.id AS Line_id, idc.id AS idc_id FROM pool
		LEFT JOIN line
		ON line.id = pool.line_id
		LEFT JOIN idc
		ON line.idc_id = idc.id
	`).Find(&data)
	if tx.Error != nil {
		config.Log.Error(tx.Error)
		return nil, ErrSearchDBData
	}
	treeData := map[int]map[int][]*idcXlsDB{}
	for _, item := range data {
		idc, ok := treeData[item.IDCID]
		if !ok {
			idc = map[int][]*idcXlsDB{}
			treeData[item.IDCID] = idc
		}
		_, ok = idc[item.LineID]
		if !ok {
			idc[item.LineID] = []*idcXlsDB{}
		}
		idc[item.LineID] = append(idc[item.LineID], item)
	}
	currRow := 0
	f := excelize.NewFile()
	for _, lv1 := range treeData {
		merL1Begin, err := excelize.CoordinatesToCellName(1, currRow+1)
		if err != nil {
			config.Log.Error(err)
			return nil, ErrXLSComputerCell
		}
		for _, lv2 := range lv1 {
			for _, pool := range lv2 {
				merL2Begin, err := excelize.CoordinatesToCellName(2, currRow+1)
				if err != nil {
					config.Log.Error(err)
					return nil, ErrXLSComputerCell
				}
				addrPools := map[string]struct{}{}
				for _, each := range strings.FieldsFunc(pool.IPAddrs, Split) {
					if strings.TrimSpace(each) == "" {
						continue
					}
					addrPools[strings.TrimSpace(each)] = struct{}{}
				}
				for content, _ := range addrPools {
					currRow += 1
					position, err := excelize.CoordinatesToCellName(1, currRow)
					if err != nil {
						config.Log.Error(err)
						return nil, ErrXLSWriteData
					}
					f.SetCellValue("Sheet1", position, pool.IDCLabel)
					position, err = excelize.CoordinatesToCellName(2, currRow)
					if err != nil {
						config.Log.Error(err)
						return nil, ErrXLSWriteData
					}
					f.SetCellValue("Sheet1", position, pool.LineLabel)
					position, err = excelize.CoordinatesToCellName(3, currRow)
					if err != nil {
						config.Log.Error(err)
						return nil, ErrXLSWriteData
					}
					f.SetCellValue("Sheet1", position, content)
				}
				merL2End, err := excelize.CoordinatesToCellName(2, currRow)
				if err != nil {
					config.Log.Error(err)
					return nil, ErrXLSComputerCell
				}
				if err := f.MergeCell("Sheet1", merL2Begin, merL2End); err != nil {
					config.Log.Error(err)
					return nil, ErrXLSMergeCell
				}
			}
		}
		merL1End, err := excelize.CoordinatesToCellName(1, currRow)
		if err != nil {
			config.Log.Error(err)
			return nil, ErrXLSComputerCell
		}
		if err := f.MergeCell("Sheet1", merL1Begin, merL1End); err != nil {
			config.Log.Error(err)
			return nil, ErrXLSMergeCell
		}
	}
	style, err := f.NewStyle(&excelize.Style{Alignment: &excelize.Alignment{
		Horizontal: "center",
		Vertical:   "center",
	}})
	if err != nil {
		config.Log.Error(err)
	} else {
		if err = f.SetColStyle("Sheet1", "A", style); err != nil {
			config.Log.Error(err)
		}
		if err = f.SetColStyle("Sheet1", "B", style); err != nil {
			config.Log.Error(err)
		}
	}
	if err := f.SetColWidth("Sheet1", "A", "A", 25); err != nil {
		config.Log.Error(err)
	}
	if err := f.SetColWidth("Sheet1", "B", "B", 30); err != nil {
		config.Log.Error(err)
	}
	if err := f.SetColWidth("Sheet1", "C", "C", 50); err != nil {
		config.Log.Error(err)
	}
	buffer := new(bytes.Buffer)
	if _, err := f.WriteTo(buffer); err != nil {
		config.Log.Error(err)
		return nil, ErrWriteCache
	}
	return &XlsDatResp{
		Data: base64.StdEncoding.EncodeToString(buffer.Bytes()),
		Name: fmt.Sprintf("机房及线路_%d.xlsx", time.Now().Unix()),
	}, Success
}

// 更新IP所属机房及线路使用的结构体
type NetInfo struct {
	ID          int          `json:"id" gorm:"column:id"`
	LineID      int          `json:"line_id" gorm:"column:line_id"`
	IDCID       int          `json:"idc_id" gorm:"column:idc_id"`
	IDCLabel    string       `json:"idc_label" gorm:"column:idc_label"`
	LineLabel   string       `json:"line_label" gorm:"column:line_label"`
	Ipaddrs     string       `json:"ipaddrs" gorm:"column:ipaddrs"`
	IPNetParsed *TypeGroupIP `json:"-" gorm:"-"`
}

type NetRange struct {
	Begin *big.Int // 起始IP
	End   *big.Int // 结束IP
}

type TypeGroupIP struct {
	IP    map[string]*net.IP    // 只是单个IP
	Net   map[string]*net.IPNet // IPV4的地址段 x.x.x.x/24
	Range []*NetRange           // 指定范围，如：192.168.1.0~192.168.2.0
}

func GetNetParams() ([]*IDC, map[int][]*Line, map[int]*TypeGroupIP, *BriefMessage) {
	db := dbs.DBObj.GetGoRM()
	if db == nil {
		config.Log.Error(InternalGetBDInstanceErr)
		return nil, nil, nil, ErrDataBase
	}
	// IDC
	idcs := []*IDC{}
	tx := db.Table("idc").Where("view=?", 1).Find(&idcs)
	if tx.Error != nil {
		config.Log.Error(tx.Error)
		return nil, nil, nil, ErrSearchDBData
	}
	// Line
	lines := []*Line{}
	tx = db.Table("line").Where("view=?", 1).Find(&lines)
	if tx.Error != nil {
		config.Log.Error(tx.Error)
		return nil, nil, nil, ErrSearchDBData
	}
	linesMap := map[int][]*Line{} // 线路ID做为KEY
	for _, l := range lines {
		_, ok := linesMap[l.ID]
		if !ok {
			linesMap[l.ID] = []*Line{}
		}
		linesMap[l.ID] = append(linesMap[l.ID], l)
	}

	// pool
	pools := []*IPAddrsPool{}
	tx = db.Table("pool").Find(&pools)
	if tx.Error != nil {
		config.Log.Error(tx.Error)
		return nil, nil, nil, ErrSearchDBData
	}
	tgis := map[int]*TypeGroupIP{} // map中的int是指线路的ID
	for _, p := range pools {
		tgi := TypeGroupIP{
			IP:    map[string]*net.IP{},
			Net:   map[string]*net.IPNet{},
			Range: []*NetRange{},
		}
		// 分割IP地址
		iplist := strings.Split(p.Ipaddrs, ";")
		for _, each := range iplist {
			currIP := strings.TrimSpace(each)
			if strings.Contains(each, "/") {
				_, nObj, err := net.ParseCIDR(currIP)
				if err != nil {
					config.Log.Error(err)
					continue
				}
				tgi.Net[each] = nObj
			} else if strings.Contains(each, "~") {
				fields := strings.Split(currIP, "~")
				if len(fields) != 2 {
					config.Log.Errorf("ip pool err: %s", currIP)
					continue
				}
				beginBig, endBig, err := utils.BigIntBeginAndEnd(strings.TrimSpace(fields[0]), strings.TrimSpace(fields[1]))
				if err != nil {
					config.Log.Error(err)
					continue
				}
				nr := NetRange{Begin: beginBig, End: endBig}
				tgi.Range = append(tgi.Range, &nr)
			} else {
				ipp := net.ParseIP(currIP)
				if ipp == nil {
					config.Log.Error("no a vaild ip address: ", each)
					continue
				}
				tgi.IP[each] = &ipp
			}
		}
		tgis[p.LineID] = &tgi
	}
	return idcs, linesMap, tgis, Success
}

func GetNetParamsV2() ([]*NetInfo, *BriefMessage) {
	db := dbs.DBObj.GetGoRM()
	if db == nil {
		config.Log.Error(InternalGetBDInstanceErr)
		return nil, ErrDataBase
	}
	sql := `SELECT pool.id, pool.line_id, pool.ipaddrs, line.label AS line_label, line.idc_id, idc.label AS idc_label FROM pool 
	LEFT JOIN line 
	ON pool.line_id=line.id
	LEFT JOIN idc
	ON line.idc_id = idc.id
	WHERE line.view = 1 AND idc.view = 1
	`
	netInfo := []*NetInfo{}
	tx := db.Table("pool").Raw(sql).Find(&netInfo)
	if tx.Error != nil {
		config.Log.Error(tx.Error)
		return nil, ErrSearchDBData
	}
	for _, n := range netInfo {
		tgi := TypeGroupIP{
			IP:    map[string]*net.IP{},
			Net:   map[string]*net.IPNet{},
			Range: []*NetRange{},
		}
		// 分割IP地址
		iplist := []string{}
		for _, each := range strings.FieldsFunc(n.Ipaddrs, Split) {
			if strings.TrimSpace(each) == "" {
				continue
			}
			iplist = append(iplist, strings.TrimSpace(each))
		}
		// iplist := strings.Split(n.Ipaddrs, ";")
		for _, each := range iplist {
			currIP := strings.TrimSpace(each)
			if strings.Contains(each, "/") {
				_, nObj, err := net.ParseCIDR(currIP)
				if err != nil {
					config.Log.Error(err)
					continue
				}
				tgi.Net[each] = nObj
			} else if strings.Contains(each, "~") {
				fields := strings.Split(currIP, "~")
				if len(fields) != 2 {
					config.Log.Errorf("ip pool err: %s", currIP)
					continue
				}
				beginBig, endBig, err := utils.BigIntBeginAndEnd(strings.TrimSpace(fields[0]), strings.TrimSpace(fields[1]))
				if err != nil {
					config.Log.Error(err)
					continue
				}
				nr := NetRange{Begin: beginBig, End: endBig}
				tgi.Range = append(tgi.Range, &nr)
			} else if strings.Contains(each, "-") {
				fields := strings.Split(currIP, "-")
				if len(fields) != 2 {
					config.Log.Errorf("ip pool err: %s", currIP)
					continue
				}
				beginBig, endBig, err := utils.BigIntBeginAndEnd(strings.TrimSpace(fields[0]), strings.TrimSpace(fields[1]))
				if err != nil {
					config.Log.Error(err)
					continue
				}
				nr := NetRange{Begin: beginBig, End: endBig}
				tgi.Range = append(tgi.Range, &nr)
			} else {
				ipp := net.ParseIP(currIP)
				if ipp == nil {
					config.Log.Error("no a vaild ip address: ", each)
					continue
				}
				tgi.IP[each] = &ipp
			}
		}
		n.IPNetParsed = &tgi
	}
	return netInfo, Success
}

func UpdateAllIPAddrs(user *UserSessionInfo, updatePart bool) *BriefMessage {
	netInfo, bf := GetNetParamsV2()
	if bf != Success {
		return bf
	}
	// 获取所有的IP地址列表：
	allIP, bf := GetAllMachine()
	if bf != Success {
		return bf
	}
	beUpdated := []*Machine{}
	for _, i := range allIP {
		if updatePart {
			if i.IDCName != "" && i.LineName != "" {
				continue
			}
		} else {
			i.IDCName = ""
			i.LineName = ""
		}
		currIP := strings.TrimSpace(i.IpAddr)
		currIPParsed := net.ParseIP(currIP)
		if currIPParsed == nil {
			config.Log.Error("not a vaild ip address: ", i.IpAddr)
			continue
		}
		for _, pItem := range netInfo {
			matched := false
			_, ok := pItem.IPNetParsed.IP[currIP]
			if ok {
				i.IDCName = pItem.IDCLabel
				i.LineName = pItem.LineLabel
				matched = true
			} else {
				for _, n := range pItem.IPNetParsed.Net {
					if n.Contains(currIPParsed) {
						i.IDCName = pItem.IDCLabel
						i.LineName = pItem.LineLabel
						matched = true
						break
					}
				}
				currIPParsedBigInt := utils.InetToBigInt(currIPParsed)
				for _, rs := range pItem.IPNetParsed.Range {
					if rs.Begin.Cmp(currIPParsedBigInt) == 0 || rs.End.Cmp(currIPParsedBigInt) == 0 {
						i.IDCName = pItem.IDCLabel
						i.LineName = pItem.LineLabel
						matched = true
						break
					}
					if rs.Begin.Cmp(currIPParsedBigInt) == -1 && rs.End.Cmp(currIPParsedBigInt) == 1 {
						i.IDCName = pItem.IDCLabel
						i.LineName = pItem.LineLabel
						matched = true
						break
					}
				}
			}
			if updatePart && matched {
				beUpdated = append(beUpdated, i)
				break
			}
		}
	}
	if updatePart {
		return WriteNetInfoToMachines(user, beUpdated)
	}
	return WriteNetInfoToMachines(user, allIP)
}

func WriteNetInfoToMachines(user *UserSessionInfo, ms []*Machine) *BriefMessage {
	db := dbs.DBObj.GetGoRM()
	if db == nil {
		config.Log.Error(InternalGetBDInstanceErr)
		return ErrDataBase
	}
	for _, m := range ms {
		tx := db.Table("machines").
			Where("id", m.ID).
			Update("idc_name", m.IDCName).
			Update("line_name", m.LineName).
			Update("update_at", time.Now()).
			Update("update_by", user.Username)
		if tx.Error != nil {
			config.Log.Error(tx.Error)
			return ErrUpdateData
		}
	}
	return Success
}

func CreateLabelForAllIPs(user *UserSessionInfo, onlyThisJobId *OnlyID) *BriefMessage {
	db := dbs.DBObj.GetGoRM()
	if db == nil {
		config.Log.Error(InternalGetBDInstanceErr)
		return ErrDataBase
	}
	// 获取所有的组
	jobs := []*Jobs{}
	if onlyThisJobId != nil {
		tx := db.Table("jobs").Where("id=?", onlyThisJobId.ID).Find(&jobs)
		if tx.Error != nil {
			config.Log.Error(tx.Error)
			return ErrSearchDBData
		}
	} else {
		tx := db.Table("jobs").Find(&jobs)
		if tx.Error != nil {
			config.Log.Error(tx.Error)
			return ErrSearchDBData
		}
	}
	// 处理每一个组
	for _, j := range jobs {
		// 获取组的所有IP地址
		sql := fmt.Sprintf(`SELECT machines.* FROM job_machines
		LEFT JOIN jobs 
		ON jobs.id = job_machines.job_id
		LEFT JOIN machines
		ON machines.id = job_machines.machine_id
		WHERE jobs.id = %d AND machines.enabled = 1
		`, j.ID)
		ipAddrsForJob := []*Machine{}
		tx := db.Table("machines").Raw(sql).Find(&ipAddrsForJob)
		if tx.Error != nil {
			config.Log.Error(tx.Error)
			return ErrSearchDBData
		}
		// 给IP地址进行分组
		subGroupCreateID := map[string]int{}
		subGroupCreateKeyVal := map[string]map[string]string{}
		idcLineMap := map[string][]*Machine{} // 稍后会把ID对应的label写入group_labels表中
		for _, m := range ipAddrsForJob {
			var key string
			if m.IDCName != "" && m.LineName != "" {
				key = fmt.Sprintf("%s_%s", m.IDCName, m.LineName)
			} else if m.IDCName != "" {
				key = m.IDCName
			} else if m.LineName != "" {
				key = m.LineName
			} else {
				key = "默认子组"
			}
			_, ok := idcLineMap[key]
			if !ok {
				idcLineMap[key] = []*Machine{}
			}
			idcLineMap[key] = append(idcLineMap[key], m)
			subGroupCreateID[key] = 0
			_, ok = subGroupCreateKeyVal[key]
			if !ok {
				subGroupCreateKeyVal[key] = map[string]string{}
			}
			subGroupCreateKeyVal[key]["idc"] = m.IDCName
			subGroupCreateKeyVal[key]["line"] = m.LineName
		}
		// 获取此JOB组中的所有子组
		sunJobGroup := []*JobGroup{}
		tx = db.Table("job_group").Where("jobs_id", j.ID).Find(&sunJobGroup)
		if tx.Error != nil {
			config.Log.Error(tx.Error)
			return ErrSearchDBData
		}
		subJobGroupIDs := []int{}
		alreadyExistJobSubGroupMap := map[string]*JobGroup{}
		for _, s := range sunJobGroup {
			alreadyExistJobSubGroupMap[s.Name] = s
			subJobGroupIDs = append(subJobGroupIDs, s.ID)
		}
		// 创建子组
		createIDForTableGroupMachines := map[int][]int{} // 记录子组ID及子组下的machine对应的ID
		for name, ipList := range idcLineMap {
			obj, ok := alreadyExistJobSubGroupMap[name] // 子组是否已经存在的不同处理
			if ok {
				subGroupCreateID[name] = obj.ID
				_, ok := createIDForTableGroupMachines[obj.ID]
				if !ok {
					createIDForTableGroupMachines[obj.ID] = []int{}
				}
				for _, i := range ipList {
					createIDForTableGroupMachines[obj.ID] = append(createIDForTableGroupMachines[obj.ID], i.ID)
				}
				continue
			}
			jg := JobGroup{
				ID:       0,
				JobsID:   j.ID,
				Name:     name,
				Enabled:  true,
				UpdateAt: time.Now(),
				UpdateBy: user.Username,
			}
			tx = db.Table("job_group").Create(&jg)
			if tx.Error != nil {
				config.Log.Error(tx.Error)
				return ErrCreateDBData
			}
			subGroupCreateID[name] = jg.ID
			alreadyExistJobSubGroupMap[name] = &jg
			_, ok = createIDForTableGroupMachines[jg.ID]
			if !ok {
				createIDForTableGroupMachines[jg.ID] = []int{}
			}
			for _, i := range ipList {
				createIDForTableGroupMachines[jg.ID] = append(createIDForTableGroupMachines[jg.ID], i.ID)
			}
		}
		// 在表group_machines中写入IP及子组的对应关系
		for jobGrpID, machineIDs := range createIDForTableGroupMachines {
			bf := writeDataToGroupMachines(subJobGroupIDs, jobGrpID, machineIDs)
			if bf != Success {
				return bf
			}
		}
		// 在表group_labels中创建数据
		for key, obj := range subGroupCreateKeyVal {
			for k, v := range obj {
				if v == "" {
					// config.Log.Warnf("labels key: %s, value is empty, skip", k)
					continue
				}
				tx = db.Table("group_labels").Exec("insert ignore into `group_labels` "+
					"(`id`, `job_group_id`, `key`, `value`, `enabled`, `update_at`, `update_by`) values(0, ?, ?, ?, 1, ?, ?)",
					subGroupCreateID[key], k, v, time.Now(), user.Username)
				if tx.Error != nil {
					config.Log.Error(tx.Error)
				}
			}
		}
	}
	return Success
}

func writeDataToGroupMachines(subJobGroupIDs []int, jobGrpID int, machineIDs []int) *BriefMessage {
	db := dbs.DBObj.GetGoRM()
	if db == nil {
		config.Log.Error(InternalGetBDInstanceErr)
		return ErrDataBase
	}
	// config.Log.Infof("diandian=> job_group_id: %d, machines_id: %d", jobGrpID, len(machineIDs))
	needDelIDs := []int{}
	wDatas := []*JobGroupIP{}
	for _, mid := range machineIDs {
		wData := JobGroupIP{
			JobGroupID: jobGrpID,
			MachinesID: mid,
		}
		needDelIDs = append(needDelIDs, mid)
		wDatas = append(wDatas, &wData)
	}
	// 清理将要写入的子组的IP地址
	tx := db.Table("group_machines").Where("job_group_id in (?) and machines_id in (?) ", subJobGroupIDs, needDelIDs).Delete(nil)
	if tx.Error != nil {
		config.Log.Error(tx.Error)
		return ErrDelData
	}
	if len(wDatas) != 0 {
		tx2 := db.Table("group_machines").Create(&wDatas)
		// tx2.Commit()
		if tx2.Error != nil {
			config.Log.Error(tx2.Error)
			return ErrCreateDBData
		}
	}
	return Success
}

type ExpandDataReq struct {
	IDC  []int `json:"idc"`
	Line []int `json:"line"`
	All  bool  `json:"all"`
}

func IDCExpand(user *UserSessionInfo, edr *ExpandDataReq) *BriefMessage {
	if edr.All {
		return expandAllIDC(user)
	}
	// 指IDC下面所有线路
	if len(edr.IDC) != 0 {
		return expandSpecialIDC(user, edr.IDC)
	}
	if len(edr.Line) != 0 {
		return expandSpecialLine(user, edr.Line)
	}
	return Success
}

type poolWithLineIDC struct {
	ID        int    `json:"id" gorm:"column:id"`
	LineID    int    `json:"line_id" gorm:"column:line_id"`
	IDCID     int    `json:"idc_id" gorm:"column:idc_id"`
	IPAddrs   string `json:"ipaddrs" gorm:"column:ipaddrs"`
	IDCLabel  string `json:"idc_label" gorm:"column:idc_label"`
	LineLabel string `json:"line_label" gorm:"column:line_label"`
}

func expandAllIDC(user *UserSessionInfo) *BriefMessage {
	db := dbs.DBObj.GetGoRM()
	if db == nil {
		config.Log.Error(InternalGetBDInstanceErr)
		return ErrDataBase
	}
	pwli := []*poolWithLineIDC{}
	tx := db.Table("pool").Raw(`select pool.id, pool.line_id, line.idc_id, pool.ipaddrs, 
		line.label as line_label, 
		idc.label as idc_label
		from pool 
		left join line 
		on line.id = pool.line_id 
		left join idc 
		on idc.id=line.idc_id WHERE line.expand = 1 AND idc.expand = 1;`).Find(&pwli)
	if tx.Error != nil {
		config.Log.Error(tx.Error)
		return ErrSearchDBData
	}
	machines := []*UploadMachine{}
	for _, each := range pwli {
		importIPs := ParseIPAddrsFromString(each.IPAddrs)
		if len(importIPs) == 0 {
			continue
		}
		for m, _ := range importIPs {
			machines = append(machines, &UploadMachine{
				ID:             0,
				IpAddr:         m,
				IDCLabel:       each.IDCLabel,
				LineLabel:      each.LineLabel,
				ImportInPool:   false,
				ImportInJobNum: 0,
				ImportError:    "",
			})
		}
	}
	umi := UploadMachinesInfo{
		Opts:     UploadOpts{IgnoreErr: true},
		JobsID:   []int{},
		Machines: machines,
		TongJi:   UploadResult{},
	}
	_, bf := UploadMachines(user, &umi)
	return bf
}

func expandSpecialIDC(user *UserSessionInfo, idcIDs []int) *BriefMessage {
	db := dbs.DBObj.GetGoRM()
	if db == nil {
		config.Log.Error(InternalGetBDInstanceErr)
		return ErrDataBase
	}
	pwli := []*poolWithLineIDC{}
	tx := db.Table("pool").Raw(`select pool.id, pool.line_id, line.idc_id, pool.ipaddrs, 
		line.label as line_label, 
		idc.label as idc_label
		from pool 
		left join line 
		on line.id = pool.line_id 
		left join idc 
		on idc.id=line.idc_id where idc.id in ?`, idcIDs).Find(&pwli)
	if tx.Error != nil {
		config.Log.Error(tx.Error)
		return ErrSearchDBData
	}
	machines := []*UploadMachine{}
	for _, each := range pwli {
		importIPs := ParseIPAddrsFromString(each.IPAddrs)
		if len(importIPs) == 0 {
			continue
		}
		for m, _ := range importIPs {
			machines = append(machines, &UploadMachine{
				ID:             0,
				IpAddr:         m,
				IDCLabel:       each.IDCLabel,
				LineLabel:      each.LineLabel,
				ImportInPool:   false,
				ImportInJobNum: 0,
				ImportError:    "",
			})
		}
	}
	umi := UploadMachinesInfo{
		Opts:     UploadOpts{IgnoreErr: true},
		JobsID:   []int{},
		Machines: machines,
		TongJi:   UploadResult{},
	}
	_, bf := UploadMachines(user, &umi)
	return bf
}

func expandSpecialLine(user *UserSessionInfo, lineIDs []int) *BriefMessage {
	db := dbs.DBObj.GetGoRM()
	if db == nil {
		config.Log.Error(InternalGetBDInstanceErr)
		return ErrDataBase
	}
	pwli := []*poolWithLineIDC{}
	tx := db.Table("pool").Raw(`select pool.id, pool.line_id, line.idc_id, pool.ipaddrs, 
		line.label as line_label, 
		idc.label as idc_label
		from pool 
		left join line 
		on line.id = pool.line_id 
		left join idc 
		on idc.id=line.idc_id where line.id in ?`, lineIDs).Find(&pwli)
	if tx.Error != nil {
		config.Log.Error(tx.Error)
		return ErrSearchDBData
	}
	machines := []*UploadMachine{}
	for _, each := range pwli {
		importIPs := ParseIPAddrsFromString(each.IPAddrs)
		if len(importIPs) == 0 {
			continue
		}
		for m, _ := range importIPs {
			machines = append(machines, &UploadMachine{
				ID:             0,
				IpAddr:         m,
				IDCLabel:       each.IDCLabel,
				LineLabel:      each.LineLabel,
				ImportInPool:   false,
				ImportInJobNum: 0,
				ImportError:    "",
			})
		}
	}
	umi := UploadMachinesInfo{
		Opts:     UploadOpts{IgnoreErr: true},
		JobsID:   []int{},
		Machines: machines,
		TongJi:   UploadResult{},
	}
	_, bf := UploadMachines(user, &umi)
	return bf
}

func PutLineExpandSwitch(user *UserSessionInfo, s *ExpandSwitchReq) *BriefMessage {
	db := dbs.DBObj.GetGoRM()
	if db == nil {
		config.Log.Error(InternalGetBDInstanceErr)
		return ErrDataBase
	}
	tx := db.Table("line").Where("id=?", s.ID).
		Update("expand", s.Switch).
		Update("update_at", time.Now()).
		Update("update_by", user.Username)
	if tx.Error != nil {
		config.Log.Error(tx.Error)
		return ErrUpdateData
	}
	return Success
}

func PutLineViewSwitch(user *UserSessionInfo, s *ViewSwitchReq) *BriefMessage {
	db := dbs.DBObj.GetGoRM()
	if db == nil {
		config.Log.Error(InternalGetBDInstanceErr)
		return ErrDataBase
	}
	tx := db.Table("line").Where("id=?", s.ID).
		Update("view", s.Switch).
		Update("update_at", time.Now()).
		Update("update_by", user.Username)
	if tx.Error != nil {
		config.Log.Error(tx.Error)
		return ErrUpdateData
	}
	return Success
}

func PutIDCExpandSwitch(user *UserSessionInfo, s *ExpandSwitchReq) *BriefMessage {
	db := dbs.DBObj.GetGoRM()
	if db == nil {
		config.Log.Error(InternalGetBDInstanceErr)
		return ErrDataBase
	}
	tx := db.Table("idc").Where("id=?", s.ID).
		Update("expand", s.Switch).
		Update("update_at", time.Now()).
		Update("update_by", user.Username)
	if tx.Error != nil {
		config.Log.Error(tx.Error)
		return ErrUpdateData
	}
	return Success
}

func PutIDCViewSwitch(user *UserSessionInfo, s *ViewSwitchReq) *BriefMessage {
	db := dbs.DBObj.GetGoRM()
	if db == nil {
		config.Log.Error(InternalGetBDInstanceErr)
		return ErrDataBase
	}
	tx := db.Table("idc").Where("id=?", s.ID).
		Update("view", s.Switch).
		Update("update_at", time.Now()).
		Update("update_by", user.Username)
	if tx.Error != nil {
		config.Log.Error(tx.Error)
		return ErrUpdateData
	}
	return Success
}
