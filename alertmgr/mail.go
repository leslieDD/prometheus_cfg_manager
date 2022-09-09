package alertmgr

import (
	"bytes"
	"html/template"
	"pro_cfg_manager/config"
	"pro_cfg_manager/utils"

	"github.com/google/uuid"
)

type ImageContent struct {
	Title  string
	Img    string
	Alt    string //width="304" height="228"
	Width  string
	Height string
}

func sendEmail(ic *ImageContent) {
	uuid := uuid.New()
	key := uuid.String()
	t, err := template.New(key).Parse(ImgTmpl)
	if err != nil {
		config.Log.Error(err)
		return
	}
	buf := new(bytes.Buffer)
	if err := t.Execute(buf, ic); err != nil {
		config.Log.Error(err)
		return
	}
	if err := utils.SendEmail(
		config.Cfg.Mail.Sender,
		config.Cfg.Mail.Password,
		config.Cfg.Mail.SMTP,
		config.Cfg.Mail.Port,
		buf.Bytes(),
		config.Cfg.Mail.ToUsers,
	); err != nil {
		config.Log.Error(err)
	}
}
