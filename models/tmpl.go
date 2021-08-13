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
	"time"

	"github.com/google/uuid"
)

type Tmpl struct {
	lock sync.Mutex
}

type TmplTable struct {
	Tmpl     string    `json:"tmpl" gorm:"column:tmpl"`
	UpdateAt time.Time `json:"update_at" gorm:"column:update_at"`
	UpdateBy string    `json:"update_by" gorm:"column:update_by"`
}

var TmplObj *Tmpl

func NewTmpl() *Tmpl {
	return &Tmpl{lock: sync.Mutex{}}
}

func Raw(value string) string {
	// config.Log.Error(value)
	return strings.Trim(value, `'`)
}

type DataStructForTmpl struct {
	Jobs []*JobsForTmpl
	// 以下以绝对路径
	AbsRooDir  string
	AbsConfDir string
	AbsRuleDir string
	// 以下为相对路径
	RelConfDir string
	RelRuleDir string
	// 模块字段
	Fields map[string]string
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
	tmplFields, bf := GetTmplFields()
	if bf != Success {
		return nil, bf
	}
	dsft := DataStructForTmpl{
		Jobs:       *jobData,
		AbsRooDir:  config.Cfg.PrometheusCfg.RootDir,
		RelConfDir: config.SubDir,
		RelRuleDir: config.RuleDir,
		Fields:     tmplFields,
	}
	if runtime.GOOS == "windows" {
		dsft.AbsConfDir = filepath.ToSlash(config.Cfg.PrometheusCfg.Conf)
		dsft.AbsRuleDir = filepath.ToSlash(config.Cfg.PrometheusCfg.RuleConf)

	} else {
		dsft.AbsConfDir = config.Cfg.PrometheusCfg.Conf
		dsft.AbsRuleDir = config.Cfg.PrometheusCfg.RuleConf
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
	Tmpl     string    `json:"tmpl" gorm:"column:tmpl"`
	UpdateAt time.Time `json:"update_at" gorm:"column:update_at"`
	UpdateBy string    `json:"update_by" gorm:"column:update_by"`
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

func PutProTmpl(user *UserSessionInfo, pt *ProTmpl) *BriefMessage {
	db := dbs.DBObj.GetGoRM()
	if db == nil {
		config.Log.Error(InternalGetBDInstanceErr)
		return ErrDataBase
	}
	pt.UpdateAt = time.Now()
	pt.UpdateBy = user.Username
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
	Jobs []JobsForTmpl
	// 以下以绝对路径
	AbsRooDir  string
	AbsConfDir string
	AbsRuleDir string
	// 以下为相对路径
	RelConfDir string
	RelRuleDir string
	// 模块字段
	Fields map[string]string
}
`
	return st, Success
}

var DefTmplText = `# my global config
global:
  scrape_interval:     15s # Set the scrape interval to every 15 seconds. Default is every 1 minute.
  evaluation_interval: 15s # Evaluate rules every 15 seconds. The default is every 1 minute.
  # scrape_timeout is set to the global default (10s).
  # scrape_timeout: 60s

# Alertmanager configuration
alerting:
  alertmanagers:
  - static_configs:
    - targets:
      - 127.0.0.1:9093

# Load rules once and periodically evaluate them according to the global 'evaluation_interval'.
rule_files:
   - "{{.RelRuleDir}}/*.yml"
  # - "first_rules.yml"
  # - "second_rules.yml"

# A scrape configuration containing exactly one endpoint to scrape:
# Here it's Prometheus itself.
scrape_configs:
  # The job name is added as a label ` + "`job=<job_name>`" + ` to any timeseries scraped from this config.
  - job_name: '监控服务本机'

    # metrics_path defaults to '/metrics'
    # scheme defaults to 'http'.

    static_configs:
    - targets: ['localhost:9090']
{{ range .Jobs }}
  - job_name: '{{.Name}}'
    file_sd_configs:
      - files:
        - "{{$.AbsConfDir}}/{{.Name}}.json"
        refresh_interval: {{$.Fields.refresh_interval}}
{{.Code}}
{{ end }}
`
