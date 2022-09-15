package alertmgr

import (
	"bytes"
	"pro_cfg_manager/config"
	"text/template"
	"time"

	"github.com/google/uuid"
)

var SendMailIns *Mail

func init() {
	SendMailIns = &Mail{}
}

type ImageContent struct {
	Title  string
	Image  string
	Alt    string //width="304" height="228"
	Width  string
	Height string
}

func sendEmail(ic *ImageContent) bool {
	uuid := uuid.New()
	key := uuid.String()
	t, err := template.New(key).Parse(ImgTmpl)
	if err != nil {
		config.Log.Error(err)
		return false
	}
	buf := new(bytes.Buffer)
	if err := t.Execute(buf, ic); err != nil {
		config.Log.Error(err)
		return false
	}
	// auth := smtp.LoginAuth(config.Cfg.Mail.Sender, config.Cfg.Mail.Password, config.Cfg.Mail.SMTP)
	// if err := SendMailIns.SendMailDirect(
	// 	config.Cfg.Mail.SMTP,
	// 	config.Cfg.Mail.Port,
	// 	auth,
	// 	config.Cfg.Mail.Sender,
	// 	config.Cfg.Mail.ToUsers,
	// 	buf.Bytes(),
	// 	false,
	// ); err != nil {
	// 	config.Log.Error(err)
	// }
	// config.Log.Println(buf.String())
	// config.Log.Println(ic.Image)
	for _, toUser := range config.Cfg.Mail.ToUsers {
		taskRst := &TaskResult{}
		task := &Task{
			SmtpHost:         config.Cfg.Mail.SMTP,
			SmtpPort:         config.Cfg.Mail.Port,
			SmtpSSL:          false,
			SmtpAuthType:     "login",
			EmailFrom:        config.Cfg.Mail.Sender,
			EmailTo:          toUser,
			EmailSubject:     ic.Title,
			EmailContentType: "text/html",
			EmailBody:        buf.String(),
			EmailAuthPass:    config.Cfg.Mail.Password,
			UseProxy:         false,
		}
		SendMailIns.RunTask(taskRst, task)
		if taskRst.Error != "" {
			config.Log.Error(err)
			return false
		}
	}
	return true
}

func SendEmailTest() bool {
	uuid := uuid.New()
	key := uuid.String()
	t, err := template.New(key).Parse(TestTmpl)
	if err != nil {
		config.Log.Error(err)
		return false
	}
	date := struct {
		Date time.Time
	}{
		Date: time.Now(),
	}
	buf := new(bytes.Buffer)
	if err := t.Execute(buf, date); err != nil {
		config.Log.Error(err)
		return false
	}
	for _, toUser := range config.Cfg.Mail.ToUsers {
		taskRst := &TaskResult{}
		task := &Task{
			SmtpHost:         config.Cfg.Mail.SMTP,
			SmtpPort:         config.Cfg.Mail.Port,
			SmtpSSL:          false,
			SmtpAuthType:     "login",
			EmailFrom:        config.Cfg.Mail.Sender,
			EmailTo:          toUser,
			EmailSubject:     "测试邮件",
			EmailContentType: "text/html",
			EmailBody:        buf.String(),
			EmailAuthPass:    config.Cfg.Mail.Password,
			UseProxy:         false,
		}
		SendMailIns.RunTask(taskRst, task)
		if taskRst.Error != "" {
			config.Log.Error(taskRst.Error)
			return false
		}
	}
	return true
}
