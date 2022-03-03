package models

import (
	"fmt"
	"math/big"
	"net"
	"pro_cfg_manager/config"
	"pro_cfg_manager/dbs"
	"pro_cfg_manager/utils"
	"strings"
	"time"

	"gorm.io/gorm"
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

type idcTreeWithID struct {
	ID   int        `json:"id"`
	Tree []*idcTree `json:"tree"`
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
	err := db.Transaction(func(tx *gorm.DB) error {
		line := Line{}
		if err := tx.Table("line").Where("id", id.ID).Find(&line).Error; err != nil {
			config.Log.Error(err)
			return err
		}
		if err := tx.Table("line").Where("id", id.ID).Delete(nil).Error; err != nil {
			config.Log.Error(err)
			return err
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
	sql := fmt.Sprintf("replace into pool(line_id, ipaddrs, update_at, update_by) values(%d, '%s', '%s', '%s');",
		newPool.LineID, newPool.Ipaddrs, time.Now().Format("2006-02-01 15:04:05"), user.Username)
	tx := db.Table("pool").Exec(sql)
	if tx.Error != nil {
		config.Log.Error(tx.Error)
		return ErrUpdateData
	}
	return Success
}

func GetIDCTree() (*idcTreeWithID, *BriefMessage) {
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
	respData := &idcTreeWithID{
		ID:   finalID,
		Tree: idcTreeResp,
	}
	return respData, Success
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
	tx := db.Table("idc").Find(&idcs)
	if tx.Error != nil {
		config.Log.Error(tx.Error)
		return nil, nil, nil, ErrSearchDBData
	}
	// Line
	lines := []*Line{}
	tx = db.Table("line").Find(&lines)
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
		iplist := strings.Split(n.Ipaddrs, ";")
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

func CreateLabelForAllIPs(user *UserSessionInfo) *BriefMessage {
	db := dbs.DBObj.GetGoRM()
	if db == nil {
		config.Log.Error(InternalGetBDInstanceErr)
		return ErrDataBase
	}
	// 获取所有的组
	jobs := []*Jobs{}
	tx := db.Table("jobs").Find(&jobs)
	if tx.Error != nil {
		config.Log.Error(tx.Error)
		return ErrSearchDBData
	}
	// 处理每一个组
	for _, j := range jobs {
		if !j.Enabled {
			continue
		}
		// 获取组的所有IP地址
		sql := fmt.Sprintf(`SELECT machines.* FROM job_machines
		LEFT JOIN jobs 
		ON jobs.id = job_machines.job_id
		LEFT JOIN machines
		ON machines.id = job_machines.machine_id
		WHERE jobs.id = %d AND machines.enabled = 1
		`, j.ID)
		ipAddrsForJob := []*Machine{}
		tx = db.Table("machines").Raw(sql).Find(&ipAddrsForJob)
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
		sunJobGroupMap := map[string]*JobGroup{}
		for _, s := range sunJobGroup {
			sunJobGroupMap[s.Name] = s
		}
		// 创建不存在的子组
		// labelsWillWrite := map[string]map[string]string{}
		crateIDForTableGroupMachines := map[int][]int{}
		for name, ipList := range idcLineMap {
			obj, ok := sunJobGroupMap[name]
			if ok {
				subGroupCreateID[name] = obj.ID
				_, ok := crateIDForTableGroupMachines[obj.ID]
				if !ok {
					crateIDForTableGroupMachines[obj.ID] = []int{}
				}
				for _, i := range ipList {
					crateIDForTableGroupMachines[obj.ID] = append(crateIDForTableGroupMachines[obj.ID], i.ID)
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
			sunJobGroupMap[name] = &jg
			_, ok = crateIDForTableGroupMachines[jg.ID]
			if !ok {
				crateIDForTableGroupMachines[jg.ID] = []int{}
			}
			for _, i := range ipList {
				crateIDForTableGroupMachines[jg.ID] = append(crateIDForTableGroupMachines[jg.ID], i.ID)
			}
		}
		// 在表group_machines中写入IP及子组的对应关系
		for jobID, machineIDs := range crateIDForTableGroupMachines {
			for _, mid := range machineIDs {
				wData := JobGroupIP{
					ID:         0,
					JobGroupID: jobID,
					MachinesID: mid,
					UpdateAt:   time.Now(),
				}
				tx = db.Table("group_machines").Create(&wData)
				if tx.Error != nil {
					config.Log.Error(tx.Error)
					continue // 因为有可能会创建到一样的机器与子组的关联关系,所以遇到错误就直接跳过
					// return ErrCreateDBData
				}
			}
		}
		// 在表group_labels中创建数据
		for key, obj := range subGroupCreateKeyVal {
			for k, v := range obj {
				if v == "" {
					config.Log.Warnf("labels key: %s, value is empty, skip", k)
					continue
				}
				wData := GroupLabels{
					ID:         0,
					JobGroupID: subGroupCreateID[key],
					Key:        k,
					Value:      v,
					Enabled:    true,
					UpdateAt:   time.Now(),
					UpdateBy:   user.Username,
				}
				tx = db.Table("group_labels").Create(&wData)
				if tx.Error != nil {
					config.Log.Error(tx.Error)
					// return ErrCreateDBData // 因为有可能会创建到一样的标签,所以遇到错误就直接跳过
				}
			}
		}
	}
	return Success
}
