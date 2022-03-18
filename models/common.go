package models

import (
	"encoding/json"
	"fmt"
	"pro_cfg_manager/config"
	"pro_cfg_manager/utils"
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

type SearchContent struct {
	Content string `json:"content" form:"content" gorm:"column:content"`
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

// {
// 	"country": "中国",
// 	"province": "福建省",
// 	"city": "福州市",
// 	"county": "仓山区",
// 	"isp": "中国电信",
// 	"country_code": "CN",
// 	"country_en": "China",
// 	"province_en": "Fujian",
// 	"city_en": "Fuzhou",
// 	"longitude": "119.320988",
// 	"latitude": "26.038912",
// 	"isp_code": "100017",
// 	"routes": "中国电信",
// 	"province_code": "350000",
// 	"city_code": "350100",
// 	"county_code": "350104"
// }

type IPPosition struct {
	Country      string `json:"country"`
	Province     string `json:"province"`
	City         string `json:"city"`
	County       string `json:"county"`
	ISP          string `json:"isp"`
	CountryCode  string `json:"country_code"`
	CountryEN    string `json:"country_en"`
	CityEN       string `json:"city_en"`
	Longitude    string `json:"longitude"`
	Latitude     string `json:"latitude"`
	ISPCode      string `json:"isp_code"`
	Routes       string `json:"routes"`
	ProvinceCode string `json:"province_code"`
	CityCode     string `json:"city_code"`
	CountyCode   string `json:"county_code"`
	Error        string `json:"error"`
}

func (i *IPPosition) String() string {
	c, err := json.Marshal(i)
	if err != nil {
		config.Log.Error(err)
		return ""
	}
	return string(c)
}

func GetIPPosition(ipAddr string) *IPPosition {
	postData := []byte(fmt.Sprintf("ipaddr=%s&submit=xxxx", ipAddr))
	resp, err := utils.PostForm(config.Cfg.Position, postData)
	if err != nil {
		config.Log.Error(err)
		return nil
	}
	if len(resp) == 0 {
		return nil
	}
	content := string(resp)
	ipp := IPPosition{}
	if !strings.Contains(content, "<html>") {
		config.Log.Error("返回的数据可能不正确，内容中没有包括<html>")
		return nil
	}
	jsonData := strings.Split(content, "<html")
	// config.Log.Print(string(resp.Header.Get("X-GeoIP")))
	// config.Log.Print(jsonData[0])
	if err := json.Unmarshal([]byte(jsonData[0]), &ipp); err != nil {
		config.Log.Error(err)
		return nil
	}
	return &ipp
}

func Split(r rune) bool {
	return r == ';' || r == '\n'
}
