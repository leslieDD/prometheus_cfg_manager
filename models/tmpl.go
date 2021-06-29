package models

import (
	"bytes"
	"pro_cfg_manager/config"
	"pro_cfg_manager/utils"
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

func (t *Tmpl) doTmpl() ([]byte, *BriefMessage) {
	t.lock.Lock()
	defer t.lock.Unlock()

	name, err := uuid.NewUUID()
	if err != nil {
		config.Log.Error(err)
		return nil, ErrGenUUID
	}
	jobData, bf := GetJobsForTmpl()
	if bf != Success {
		return nil, bf
	}
	tmplParse, err := template.
		New(name.String()).
		// Funcs(template.FuncMap{"raw": Raw}).
		Parse(config.Cfg.PrometheusCfg.TmplContext)
	if err != nil {
		config.Log.Error(err)
		return nil, ErrTmplParse
	}
	buf := new(bytes.Buffer)
	if err := tmplParse.Execute(buf, *jobData); err != nil {
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
