package models

import (
	"encoding/json"
	"pro_cfg_manager/config"
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
	nt := total % int64(sp.PageSize)
	if nt == 0 {
		rsp.TotalPage = total / int64(sp.PageSize)
	} else {
		rsp.TotalPage = total/int64(sp.PageSize) + 1
	}
	rsp.Data = data
	return &rsp
}

func JsonToIntSlice(m *Machine) ([]int, *BriefMessage) {
	v, err := m.JobId.MarshalJSON()
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
