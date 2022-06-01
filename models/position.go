package models

import (
	"encoding/json"
	"fmt"
	"net"
	"pro_cfg_manager/config"
	"pro_cfg_manager/utils"
	"strings"
	"sync"

	"github.com/lionsoul2014/ip2region/binding/golang/ip2region"
	"github.com/oschwald/geoip2-golang"
)

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

// GPSObj GPSObj
var GPSObj = NewGPS()

// NewGPS NewGPS
func NewGPS() *GPS {
	gps := &GPS{}
	gps.InitGeo()
	gps.InitIP2Region()
	return gps
}

// GPS GPS
type GPS struct {
	geolock sync.Mutex
	dbGeo   *geoip2.Reader

	ip2region   sync.Mutex
	dbIP2Region *ip2region.Ip2Region
}

// InitGeo InitGeo
func (g *GPS) InitGeo() {
	if strings.TrimSpace(config.Cfg.GPS.Geo) == "" {
		config.Log.Fatal("gsp->geo error")
	}
	db, err := geoip2.Open(config.Cfg.GPS.Geo)
	if err != nil {
		config.Log.Fatal(err)
	}
	g.dbGeo = db
}

// InitIP2Region InitIP2Region
func (g *GPS) InitIP2Region() {
	if strings.TrimSpace(config.Cfg.GPS.IP2region) == "" {
		config.Log.Fatal("gsp->ip2region error")
	}
	db, err := ip2region.New(config.Cfg.GPS.IP2region)
	if err != nil {
		config.Log.Fatal(err)
	}
	g.dbIP2Region = db
}

// SearchIP2Region SearchIP2Region
func (g *GPS) SearchIP2Region(ipStr string, defaultEnableValue bool) *IPPosition {
	g.ip2region.Lock()
	defer g.ip2region.Unlock()

	ip := net.ParseIP(ipStr)
	if ip == nil {
		config.Log.Errorf("ip error: %s", ipStr)
		return nil
	}

	record, err := g.dbIP2Region.MemorySearch(ipStr)
	if err != nil {
		config.Log.Error(err)
		return nil
	}
	ginfo := &IPPosition{
		Country:      record.Country,
		Province:     record.Province,
		City:         record.City,
		County:       record.Region,
		ISP:          record.ISP,
		CountryCode:  "",
		CountryEN:    "",
		CityEN:       "",
		Longitude:    "",
		Latitude:     "",
		ISPCode:      "",
		Routes:       "",
		ProvinceCode: "",
		CityCode:     "",
		CountyCode:   "",
		Error:        "",
	}
	return ginfo
}
