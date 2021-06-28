package models

import (
	"pro_cfg_manager/config"
	"pro_cfg_manager/dbs"
	"pro_cfg_manager/utils"

	"gorm.io/gorm"
)

type TreeNode struct {
	ID       int         `json:"id"`
	Label    string      `json:"label"`
	Level    int         `json:"level"`
	Parent   int         `json:"parent"`
	Path     []string    `json:"path"`
	Code     string      `json:"code"`
	Children []*TreeNode `json:"children"`
}

func GetNodesFromDB() ([]*TreeNode, *BriefMessage) {
	rulesGroups, bf := GetRulesGroups()
	if bf != Success {
		return nil, bf
	}
	root := &TreeNode{
		ID:       0,
		Level:    1,
		Label:    "监控规则列表",
		Path:     []string{},
		Children: []*TreeNode{},
	}
	root.Code = utils.Getmd5(root.Label)
	root.Path = append(root.Path, root.Label)
	tns := []*TreeNode{
		root,
	}
	// 不考虑效率了
	rootNodes := []*TreeNode{}
	for _, rg := range rulesGroups {
		thisRootNode := TreeNode{
			ID:       rg.ID,
			Level:    2,
			Label:    rg.Name,
			Parent:   -1,
			Path:     []string{root.Label, rg.Name},
			Code:     utils.Getmd5(rg.Name),
			Children: []*TreeNode{}}
		subGroup, bf := GetSubGroup(rg.ID)
		if bf != Success {
			return nil, bf
		}
		thisSubNodes := []*TreeNode{}
		for _, sub := range subGroup {
			thisSubNode := TreeNode{
				ID:       sub.ID,
				Level:    3,
				Label:    sub.Name,
				Parent:   rg.ID,
				Path:     []string{root.Label, rg.Name, sub.Name},
				Code:     utils.Getmd5(sub.Name),
				Children: []*TreeNode{},
			}
			mrs, bf := GetMonitorRules(sub.ID, false)
			if bf != Success {
				return nil, bf
			}
			for _, mr := range mrs {
				thisSubNode.Children = append(thisSubNode.Children, &TreeNode{
					ID:       mr.ID,
					Level:    4,
					Label:    mr.Alert,
					Parent:   sub.ID,
					Code:     utils.Getmd5(mr.Alert),
					Path:     []string{root.Label, rg.Name, sub.Name, mr.Alert},
					Children: nil,
				})
			}
			thisSubNodes = append(thisSubNodes, &thisSubNode)
		}
		thisRootNode.Children = thisSubNodes
		rootNodes = append(rootNodes, &thisRootNode)
	}
	tns[0].Children = rootNodes
	return tns, Success
}

// func GetNodesFromDB() *BriefMessage {
// 	nodeInfos, bf := GetNodeInfo()
// 	if bf != nil {
// 		return bf
// 	}
// 	tns := []TreeNode{
// 		TreeNode{ID: 0, Level: 1, Label: "监控规则列表", Children: []TreeNode{}},
// 	}
// 	// 不考虑效率了
// 	rulesGroupsMap := map[int]TreeNode{}
// 	subGroupsMap := map[int][]TreeNode{}
// 	alerts := map[int][]TreeNode{}
// 	for _, rg := range nodeInfos {
// 		_, ok := rulesGroupsMap[rg.RulesGroupsID]
// 		if !ok {
// 			rulesGroupsMap[rg.RulesGroupsID] = TreeNode{
// 				ID:       rg.RulesGroupsID,
// 				Level:    2,
// 				Label:    rg.RulesGroupsName,
// 				Children: []TreeNode{}}
// 		}
// 		_, ok = subGroupsMap[rg.RulesGroupsID]
// 		if !ok {
// 			subGroupsMap[rg.RulesGroupsID] = []TreeNode{TreeNode{
// 				ID:       rg.SubGroupID,
// 				Level:    3,
// 				Label:    rg.SubGroupName,
// 				Children: []TreeNode{}}}
// 		} else {
// 			subGroupsMap[rg.RulesGroupsID] = append(subGroupsMap[rg.RulesGroupsID], TreeNode{
// 				ID:       rg.SubGroupID,
// 				Level:    3,
// 				Label:    rg.SubGroupName,
// 				Children: []TreeNode{}})
// 		}
// 		_, ok = alerts[rg.SubGroupID]
// 		if !ok {
// 			alerts[rg.MonitorRulesID] = []TreeNode{TreeNode{
// 				ID:       rg.MonitorRulesID,
// 				Level:    4,
// 				Label:    rg.Alert,
// 				Children: nil}}
// 		} else {
// 			alerts[rg.MonitorRulesID] = append(alerts[rg.MonitorRulesID], TreeNode{
// 				ID:       rg.MonitorRulesID,
// 				Level:    4,
// 				Label:    rg.Alert,
// 				Children: nil})
// 		}
// 	}
// 	for k, _ := range subGroupsMap {

// 	}
// 	// for _, rg := range nodeInfos {
// 	// 	obj := rulesGroupsMap[rg.RulesGroupsID]
// 	// 	// rulesGroupsMap[rg.RulesGroupsID].Children
// 	// }
// 	// tns[0].Children
// }

type RulesGroups struct {
	ID   int    `json:"id" gorm:"column:id"`
	Name string `json:"name" gorm:"column:name"`
}

func GetRulesGroups() ([]RulesGroups, *BriefMessage) {
	db := dbs.DBObj.GetGoRM()
	if db == nil {
		config.Log.Error(InternalGetBDInstanceErr)
		return nil, ErrDataBase
	}
	rgs := []RulesGroups{}
	tx := db.Table("rules_groups").Find(&rgs)
	if tx.Error != nil {
		config.Log.Error(tx.Error)
		return nil, ErrSearchDBData
	}
	return rgs, Success
}

type NodeInfo struct {
	RulesGroupsID   int    `json:"rules_groups_id" gorm:"column:rules_groups_id"`
	RulesGroupsName string `json:"rules_groups_name" gorm:"column:rules_groups_name"`
	SubGroupID      int    `json:"sub_group_id" gorm:"column:sub_group_id"`
	SubGroupName    string `json:"sub_group_name" gorm:"column:sub_group_name"`
	MonitorRulesID  int    `json:"monitor_rules_id" gorm:"column:monitor_rules_id"`
	Alert           string `json:"alert" gorm:"column:alert"`
}

func GetNodeInfo() ([]NodeInfo, *BriefMessage) {
	db := dbs.DBObj.GetGoRM()
	if db == nil {
		config.Log.Error(InternalGetBDInstanceErr)
		return nil, ErrDataBase
	}
	ni := []NodeInfo{}
	tx := db.Table("sub_group").Raw(`SELECT 
	rules_groups.id AS rules_groups_id, 
	rules_groups.name AS rules_groups_name, 
	sub_group.id AS sub_group_id, 
	sub_group.name AS sub_group_name, 
	monitor_rules.id AS monitor_rules_id, 
	monitor_rules.alert
  FROM sub_group 
  LEFT JOIN rules_groups  
  ON rules_groups.id = sub_group.rules_groups_id 
  LEFT JOIN monitor_rules 
  ON sub_group.id = monitor_rules.sub_group_id`).Find(&ni)
	if tx.Error != nil {
		config.Log.Error(tx.Error)
		return nil, ErrSearchDBData
	}
	return ni, Success
}

type SubGroup struct {
	ID            int    `json:"id" gorm:"column:id"`
	Name          string `json:"name" gorm:"column:name"`
	RulesGroupsID int    `json:"rules_groups_id" gorm:"column:rules_groups_id"`
}

func GetSubGroup(rulesGroupsID int) ([]SubGroup, *BriefMessage) {
	db := dbs.DBObj.GetGoRM()
	if db == nil {
		config.Log.Error(InternalGetBDInstanceErr)
		return nil, ErrDataBase
	}
	sgs := []SubGroup{}
	tx := db.Table("sub_group").Where("rules_groups_id=?", rulesGroupsID).Find(&sgs)
	if tx.Error != nil {
		config.Log.Error(tx.Error)
		return nil, ErrSearchDBData
	}
	return sgs, Success
}

type MonitorRule struct {
	ID         int    `json:"id" gorm:"column:id"`
	Alert      string `json:"alert" gorm:"column:alert"`
	Expr       string `json:"expr" gorm:"column:expr"`
	For        string `json:"for" gorm:"column:for"`
	SubGroupID int    `json:"sub_group_id" gorm:"column:sub_group_id"`
	Enabled    bool   `json:"enabled" gorm:"enabled"`
}

func GetMonitorRules(subGroupID int, isPublish bool) ([]MonitorRule, *BriefMessage) {
	db := dbs.DBObj.GetGoRM()
	if db == nil {
		config.Log.Error(InternalGetBDInstanceErr)
		return nil, ErrDataBase
	}
	mrs := []MonitorRule{}
	var tx *gorm.DB
	if isPublish {
		tx = db.Table("monitor_rules").Where("enabled=1 and sub_group_id=?", subGroupID).Find(&mrs)
	} else {
		tx = db.Table("monitor_rules").Where("sub_group_id=?", subGroupID).Find(&mrs)
	}
	if tx.Error != nil {
		config.Log.Error(tx.Error)
		return nil, ErrSearchDBData
	}
	return mrs, Success
}

func GetMonitorRulesByID(ID int) (*MonitorRule, *BriefMessage) {
	db := dbs.DBObj.GetGoRM()
	if db == nil {
		config.Log.Error(InternalGetBDInstanceErr)
		return nil, ErrDataBase
	}
	mr := MonitorRule{}
	tx := db.Table("monitor_rules").Where("id=?", ID).Find(&mr)
	if tx.Error != nil {
		config.Log.Error(tx.Error)
		return nil, ErrSearchDBData
	}
	return &mr, Success
}

type QueryGetNode struct {
	ID    int    `json:"id" form:"id"`
	Label string `json:"label" form:"label"`
	Level int    `json:"level" form:"level"`
}

type Label struct {
	ID             int    `json:"id" gorm:"column:id"`
	Key            string `json:"key" gorm:"column:key"`
	Value          string `json:"value" gorm:"column:value"`
	MonitorRulesID int    `json:"monitor_rules_id" gorm:"column:monitor_rules_id"`
	IsNew          bool   `json:"is_new" gorm:"-"`
}

type TreeNodeInfo struct {
	ID          int     `json:"id" gorm:"column:id"`
	Alert       string  `json:"alert" gorm:"column:alert"`
	Expr        string  `json:"expr" gorm:"column:expr"`
	For         string  `json:"for" gorm:"column:for"`
	SubGroupID  int     `json:"sub_group_id" gorm:"column:sub_group_id"`
	Labels      []Label `json:"labels" gorm:"column:labels"`
	Annotations []Label `json:"annotations" gorm:"column:annotations"`
	Enabled     bool    `json:"enabled" gorm:"enabled"`
}

func GetNode(qni *QueryGetNode) (*TreeNodeInfo, *BriefMessage) {
	mr, bf := GetMonitorRulesByID(qni.ID)
	if bf != Success {
		return nil, bf
	}
	tni := TreeNodeInfo{
		ID:         mr.ID,
		Alert:      mr.Alert,
		Expr:       mr.Expr,
		For:        mr.For,
		SubGroupID: mr.SubGroupID,
		Enabled:    mr.Enabled,
	}

	tni.Labels, bf = SearchLabelsByMonitorID(tni.ID)
	if bf != Success {
		return nil, bf
	}
	tni.Annotations, bf = SearchAnnotationsByMonitorID(tni.ID)
	if bf != Success {
		return nil, bf
	}
	return &tni, Success
}

func PostNode(nodeInfo *TreeNodeInfo) (*TreeNodeInfo, *BriefMessage) {
	db := dbs.DBObj.GetGoRM()
	if db == nil {
		config.Log.Error(InternalGetBDInstanceErr)
		return nil, ErrDataBase
	}
	createMap := map[string]interface{}{
		"alert":        nodeInfo.Alert,
		"expr":         nodeInfo.Expr,
		"for":          nodeInfo.For,
		"sub_group_id": nodeInfo.SubGroupID,
		"enabled":      nodeInfo.Enabled,
	}
	tx := db.Table("monitor_rules").Create(&createMap)
	if tx.Error != nil {
		config.Log.Error(tx.Error)
		return nil, ErrCreateDBData
	}
	if bf := BatchSaveToTableLabels(db, nodeInfo); bf != Success {
		return nil, bf
	}
	if bf := BatchSaveToTableAnnotations(db, nodeInfo); bf != Success {
		return nil, bf
	}
	return GetNode(&QueryGetNode{ID: nodeInfo.ID, Level: 4})
}

func PutNode(nodeInfo *TreeNodeInfo) (*TreeNodeInfo, *BriefMessage) {
	db := dbs.DBObj.GetGoRM()
	if db == nil {
		config.Log.Error(InternalGetBDInstanceErr)
		return nil, ErrDataBase
	}
	tx := db.Table("monitor_rules").Where("id=?", nodeInfo.ID).
		Update("alert", nodeInfo.Alert).
		Update("expr", nodeInfo.Expr).
		Update("for", nodeInfo.For).
		Update("sub_group_id", nodeInfo.SubGroupID).
		Update("enabled", nodeInfo.Enabled)
	if tx.Error != nil {
		config.Log.Error(tx.Error)
		return nil, ErrCreateDBData
	}
	if bf := BatchSaveToTableLabels(db, nodeInfo); bf != Success {
		return nil, bf
	}
	if bf := BatchSaveToTableAnnotations(db, nodeInfo); bf != Success {
		return nil, bf
	}
	return GetNode(&QueryGetNode{ID: nodeInfo.ID, Level: 4})
}

func SearchLabelsByMonitorID(MID int) ([]Label, *BriefMessage) {
	db := dbs.DBObj.GetGoRM()
	if db == nil {
		config.Log.Error(InternalGetBDInstanceErr)
		return nil, ErrDataBase
	}
	label := []Label{}
	tx := db.Table("monitor_labels").Where("monitor_rules_id=?", MID).Find(&label)
	if tx.Error != nil {
		config.Log.Error(tx.Error)
		return nil, ErrSearchDBData
	}
	return label, Success
}

func SearchAnnotationsByMonitorID(MID int) ([]Label, *BriefMessage) {
	db := dbs.DBObj.GetGoRM()
	if db == nil {
		config.Log.Error(InternalGetBDInstanceErr)
		return nil, ErrDataBase
	}
	label := []Label{}
	tx := db.Table("annotations").Where("monitor_rules_id=?", MID).Find(&label)
	if tx.Error != nil {
		config.Log.Error(tx.Error)
		return nil, ErrSearchDBData
	}
	return label, Success
}

type DefaultLables struct {
	ID    int    `json:"id" gorm:"column:id"`
	Lable string `json:"label" gorm:"column:label"`
}

func GetDefLabels() ([]DefaultLables, *BriefMessage) {
	db := dbs.DBObj.GetGoRM()
	if db == nil {
		config.Log.Error(InternalGetBDInstanceErr)
		return nil, ErrDataBase
	}
	defLables := []DefaultLables{}
	tx := db.Table("labels").Find(&defLables)
	if tx.Error != nil {
		config.Log.Error(tx.Error)
		return nil, ErrSearchDBData
	}
	return defLables, Success
}

func DelLabel(l *Label, labelType string) *BriefMessage {
	if labelType == "" {
		return ErrQueryData
	}
	db := dbs.DBObj.GetGoRM()
	if db == nil {
		config.Log.Error(InternalGetBDInstanceErr)
		return ErrDataBase
	}
	var tx *gorm.DB
	if labelType == "labels" {
		tx = db.Table("monitor_labels").Where("id=?", l.ID).Delete(nil)
	} else if labelType == "annotations" {
		tx = db.Table("annotations").Where("id=?", l.ID).Delete(nil)
	} else {
		return ErrQueryData
	}
	if tx.Error != nil {
		config.Log.Error(tx.Error)
		return ErrDelData
	}
	return Success
}
