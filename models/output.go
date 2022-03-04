package models

import (
	"bytes"
	"encoding/csv"
	"encoding/json"
	"pro_cfg_manager/config"
	"pro_cfg_manager/dbs"
	"pro_cfg_manager/utils"
	"sort"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

type MS []*Machine

func (m MS) Len() int {
	return len(m)
}
func (m MS) Less(i, j int) bool {
	bigA := utils.IPAddrToBigInt(m[i].IpAddr)
	bigB := utils.IPAddrToBigInt(m[j].IpAddr)
	return bigA.Cmp(bigB) == -1
	// return m[i].IpAddr < m[j].IpAddr
}
func (m MS) Swap(i, j int) {
	m[i], m[j] = m[j], m[i]
}

type DownloadCsv struct {
	Data string `json:"data"`
	Name string `json:"name"`
}

func OutputIPlist(c *gin.Context) (*DownloadCsv, *BriefMessage) {
	db := dbs.DBObj.GetGoRM()
	if db == nil {
		config.Log.Error(InternalGetBDInstanceErr)
		return nil, ErrDataBase
	}
	ms := []*Machine{}
	tx := db.Table("machines").Find(&ms)
	if tx.Error != nil {
		config.Log.Error(tx.Error)
		return nil, ErrSearchDBData
	}
	sort.Sort(MS(ms))
	list := [][]string{}
	list = append(list, []string{"IP地址", "Country", "Province", "City", "County", "ISP", "routes"})
	for _, m := range ms {
		position := IPPosition{}
		err := json.Unmarshal([]byte(m.Position), &position)
		if err != nil {
			config.Log.Error(err)
			return nil, ErrDataParse
		}
		list = append(list, []string{
			m.IpAddr, position.Country, position.Province, position.City, position.County, position.ISP, position.Routes,
		})
	}
	b := &bytes.Buffer{}
	b.WriteString("\xEF\xBB\xBF")
	w := csv.NewWriter(b)
	w.WriteAll(list)
	w.Flush()
	// timeUnix := time.Now().Unix()
	// filename := "IP列表" + strconv.FormatInt(timeUnix, 10) + ".csv"
	// c.Writer.Header().Set("Content-Type", "application/csv")
	// c.Writer.Header().Set("Content-Disposition", fmt.Sprintf("attachment;filename=%s", filename))
	// c.Writer.WriteHeader(200)
	// // c.Writer.WriteString("\xEF\xBB\xBF")
	// c.Writer.WriteString(b.String())
	// timeUnix := time.Now().Unix()
	// filename := "IP列表" + strconv.FormatInt(timeUnix, 10) + ".csv"
	dcvs := &DownloadCsv{
		Data: b.String(),
		Name: "IP列表" + strconv.FormatInt(time.Now().Unix(), 10) + ".csv",
	}
	return dcvs, Success
}
