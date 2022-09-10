package alertmgr

import (
	"bytes"
	"html/template"
	"pro_cfg_manager/config"
	"pro_cfg_manager/smtp"

	"github.com/google/uuid"
)

var SendMailIns *Mail

func init() {
	SendMailIns = &Mail{}
}

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
	auth := smtp.LoginAuth(config.Cfg.Mail.Sender, config.Cfg.Mail.Password, config.Cfg.Mail.SMTP)
	if err := SendMailIns.SendMailDirect(
		config.Cfg.Mail.SMTP,
		config.Cfg.Mail.Port,
		auth,
		config.Cfg.Mail.Sender,
		config.Cfg.Mail.ToUsers,
		buf.Bytes(),
		false,
	); err != nil {
		config.Log.Error(err)
	}
}
