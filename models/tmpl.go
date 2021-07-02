package models

import (
	"bytes"
	"path/filepath"
	"pro_cfg_manager/config"
	"pro_cfg_manager/dbs"
	"pro_cfg_manager/utils"
	"runtime"
	"strings"
	"sync"
	"text/template"

	"github.com/google/uuid"
)

type Tmpl struct {
	lock sync.Mutex
}

var TmplObj = Tmpl{lock: sync.Mutex{}}

func Raw(value string) string {
	// config.Log.Error(value)
	return strings.Trim(value, `'`)
}

type DataStructForTmpl struct {
	Jobs    []JobsForTmpl
	RooDir  string
	ConfDir string
}

func (t *Tmpl) doTmpl() ([]byte, *BriefMessage) {
	t.lock.Lock()
	defer t.lock.Unlock()

	name, err := uuid.NewUUID()
	if err != nil {
		config.Log.Error(err)
		return nil, ErrGenUUID
	}
	tmpl, bf := GetProTmpl()
	if bf != Success {
		return nil, bf
	}
	jobData, bf := GetJobsForTmpl()
	if bf != Success {
		return nil, bf
	}
	tmplParse, err := template.
		New(name.String()).
		// Funcs(template.FuncMap{"raw": Raw}).
		// Parse(config.Cfg.PrometheusCfg.TmplContext)
		Parse(tmpl.Tmpl)
	if err != nil {
		config.Log.Error(err)
		return nil, ErrTmplParse
	}
	dsft := DataStructForTmpl{
		Jobs:   *jobData,
		RooDir: config.Cfg.PrometheusCfg.RootDir,
	}
	if runtime.GOOS == "windows" {
		dsft.ConfDir = filepath.ToSlash(config.Cfg.PrometheusCfg.Conf)
	} else {
		dsft.ConfDir = config.Cfg.PrometheusCfg.Conf
	}
	buf := new(bytes.Buffer)
	if err := tmplParse.Execute(buf, dsft); err != nil {
		config.Log.Error(err)
		return nil, ErrTmplParse
	}
	return buf.Bytes(), Success
}

func (t *Tmpl) doWrite(content []byte) *BriefMessage {
	err := utils.WIoutilByte(config.Cfg.PrometheusCfg.MainConf, content)
	if err != nil {
		config.Log.Error(err)
		return ErrDataWriteDisk
	}
	return Success
}

type ProTmpl struct {
	Tmpl string `json:"tmpl" gorm:"column:tmpl"`
}

func GetProTmpl() (*ProTmpl, *BriefMessage) {
	db := dbs.DBObj.GetGoRM()
	if db == nil {
		config.Log.Error(InternalGetBDInstanceErr)
		return nil, ErrDataBase
	}
	pt := ProTmpl{}
	tx := db.Table("tmpl").First(&pt)
	if tx.Error != nil {
		config.Log.Error(tx.Error)
		return nil, ErrSearchDBData
	}
	return &pt, Success
}

func PutProTmpl(pt *ProTmpl) *BriefMessage {
	db := dbs.DBObj.GetGoRM()
	if db == nil {
		config.Log.Error(InternalGetBDInstanceErr)
		return ErrDataBase
	}
	tx := db.Table("tmpl").Where("1=1").Update("tmpl", pt.Tmpl)
	if tx.Error != nil {
		config.Log.Error(tx.Error)
		return ErrUpdateData
	}
	return Success
}

func GetStruct() (string, *BriefMessage) {
	st := `
type Jobs struct {
	ID           int       
	Name         string    
	Port         int       
	CfgName      string    
	IsCommon     bool      
	DisplayOrder int       
	ReLabelID    int       
	Count        int64     
	UpdateAt     time.Time 
	Online       string    
	LastError    string    
}

type JobsForTmpl struct {
	ReLabelName string 
	Code        string 
	Jobs
}

type DataStructForTmpl struct {
	Jobs    []JobsForTmpl
	RooDir  string
	ConfDir string
}
`
	return st, Success
}
