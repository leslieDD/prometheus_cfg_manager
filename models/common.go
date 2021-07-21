package models

import (
	"encoding/json"
	"fmt"
	"pro_cfg_manager/config"
	"strconv"
	"strings"

	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type OptionType int

const (
	isIPAddr OptionType = 1
	isGroup  OptionType = 2
)

// SplitPage SplitPage分页
type SplitPage struct {
	PageNo   int        `form:"pageNo"`
	PageSize int        `form:"pageSize"`
	Search   string     `form:"search"`
	Option   OptionType `form:"option"`
	ID       int        `form:"id"`
}

// ResSplitPage 返回的分页信息
type ResSplitPage struct {
	PageSize   int64       `json:"pageSize"`
	PageNo     int64       `json:"pageNo"`
	TotalCount int64       `json:"totalCount"`
	TotalPage  int64       `json:"totalPage"`
	Data       interface{} `json:"data"`
}

// CalSplitPage CalSplitPage
func CalSplitPage(sp *SplitPage, total int64, data interface{}) *ResSplitPage {
	rsp := ResSplitPage{}
	rsp.PageSize = int64(sp.PageSize)
	rsp.PageNo = int64(sp.PageNo)
	rsp.TotalCount = total
	pageSize := sp.PageSize
	if pageSize <= 0 {
		pageSize = 1
	}
	nt := total % int64(pageSize)
	if nt == 0 {
		rsp.TotalPage = total / int64(pageSize)
	} else {
		rsp.TotalPage = total/int64(pageSize) + 1
	}
	rsp.Data = data
	return &rsp
}

func JsonToIntSlice(jsonData datatypes.JSON) ([]int, *BriefMessage) {
	v, err := jsonData.MarshalJSON()
	if err != nil {
		config.Log.Error(err)
		return nil, ErrJobDataFormat
	}
	idList := []int{}
	if err := json.Unmarshal(v, &idList); err != nil {
		config.Log.Error(err)
		return nil, ErrJobDataFormat
	}
	return idList, Success
}

func BatchSaveToTableLabels(db *gorm.DB, node *TreeNodeInfo) *BriefMessage {
	if node.Labels == nil || len(node.Labels) == 0 {
		return Success
	}
	for _, item := range node.Labels {
		if item.IsNew {
			item.MonitorRulesID = node.ID
			tx := db.Table("monitor_labels").Create(&item)
			if tx.Error != nil {
				if strings.Contains(tx.Error.Error(), "Duplicate entry") {
					continue
				}
				config.Log.Error(tx.Error)
				return ErrCreateDBData
			}
		} else {
			tx := db.Table("monitor_labels").Where("id=?", item.ID).
				Update("monitor_rules_id", node.ID).
				Update("key", item.Key).
				Update("value", item.Value)
			if tx.Error != nil {
				config.Log.Error(tx.Error)
				return ErrUpdateData
			}
		}
	}
	return Success
}

func BatchSaveToTableAnnotations(db *gorm.DB, node *TreeNodeInfo) *BriefMessage {
	if node.Annotations == nil || len(node.Annotations) == 0 {
		return Success
	}
	for _, item := range node.Annotations {
		if item.IsNew {
			item.ID = 0
			item.MonitorRulesID = node.ID
			tx := db.Table("annotations").Create(&item)
			if tx.Error != nil {
				if strings.Contains(tx.Error.Error(), "Duplicate entry") {
					continue
				}
				config.Log.Error(tx.Error)
				return ErrCreateDBData
			}
		} else {
			tx := db.Table("annotations").Where("id=?", item.ID).
				Update("monitor_rules_id", node.ID).
				Update("key", item.Key).
				Update("value", item.Value)
			if tx.Error != nil {
				config.Log.Error(tx.Error)
				return ErrUpdateData
			}
		}
	}
	return Success
}

func DiffSliceToInt(newInts []int, dbInts []OnlyID) []int {
	newMap := map[int]struct{}{}
	for _, i := range newInts {
		newMap[i] = struct{}{}
	}
	cyInt := []int{}
	for _, o := range dbInts {
		if _, ok := newMap[o.ID]; !ok {
			cyInt = append(cyInt, o.ID)
		}
	}
	return cyInt
}

func DiffSliceToStr(newInts []int, dbInts []OnlyID) []string {
	newMap := map[int]struct{}{}
	for _, i := range newInts {
		newMap[i] = struct{}{}
	}
	cyInt := []string{}
	for _, o := range dbInts {
		if _, ok := newMap[o.ID]; !ok {
			cyInt = append(cyInt, fmt.Sprint(o.ID))
		}
	}
	return cyInt
}

func IntSliceStrSliceJoind(ints []int) string {
	newInts := []string{}
	for _, i := range ints {
		newInts = append(newInts, fmt.Sprint(i))
	}
	return strings.Join(newInts, ",")
}

func ConvertStrToIntSlice(s string) ([]int, error) {
	ints := []int{}
	if len(s) == 0 {
		return ints, nil
	}
	for _, each := range strings.Split(s, ";") {
		eachInt, err := strconv.ParseInt(each, 10, 0)
		if err != nil {
			config.Log.Error(err)
			return ints, err
		}
		ints = append(ints, int(eachInt))
	}
	return ints, nil
}

func ConvertOnlyIdToIntSlice(oi []OnlyID) []int {
	s := make([]int, len(oi))
	for _, o := range oi {
		s = append(s, o.ID)
	}
	return s
}
