package alertmgr

import (
	"bytes"
	"crypto/tls"
	"encoding/base64"
	"errors"
	"fmt"
	"mime"
	"mime/multipart"
	"net"
	"net/textproto"
	"pro_cfg_manager/config"
	"strings"
	"time"

	"pro_cfg_manager/smtp"
	"pro_cfg_manager/utils"

	"golang.org/x/net/proxy"
)

type Task struct {
	ID               string    `json:"id"`
	PlanID           string    `json:"plan_id"`
	SmtpHost         string    `json:"smtp_host"`
	SmtpPort         string    `json:"smtp_port"`
	SmtpSSL          bool      `json:"smtp_ssl"`
	SmtpAuthType     string    `json:"smtp_auth_type"`
	EmailFrom        string    `json:"email_from"`
	EmailTo          string    `json:"email_to"`
	EmailSubject     string    `json:"email_subject"`
	EmailContentType string    `json:"email_content_type"`
	EmailBody        string    `json:"email_body"`
	EmailAuthPass    string    `json:"email_auth_pass"`
	EmailAuthCode    string    `json:"email_auth_code"`
	ProxyHost        string    `json:"proxy_host"`
	ProxyPort        string    `json:"proxy_port"`
	ProxyAuthUser    string    `json:"proxy_auth_user"`
	ProxyAuthPass    string    `json:"proxy_auth_pass"`
	ProxyAuth        bool      `json:"proxy_auth"`
	UseProxy         bool      `json:"use_proxy"`
	ProxyKeepAlive   bool      `json:"proxy_keep_alive"`
	ProxyExpireTime  time.Time `json:"proxy_expire_time"`
	Type             string    `json:"type"`
}

type TaskResult struct {
	ID        string      `json:"id"` // 任务编号
	PlanID    string      `json:"plan_id"`
	Status    int         `json:"status"` // 任务状态，如：此任务的类型在客户端上面不支持，没有相应功能的实现
	Error     string      `json:"error"`  // 任务出错原因
	Type      string      `json:"type"`   // 任务的类型
	EmailFrom string      `json:"email_from"`
	Result    interface{} `json:"result"` // 任务的结果
}

type Mail struct{}

type socket5 struct {
	Auth    *proxy.Auth
	Network string
	Host    string
	Port    string
}

func (m *Mail) RunTask(taskResult *TaskResult, task *Task) {
	taskResult.ID = task.ID
	taskResult.Type = task.Type
	taskResult.EmailFrom = task.EmailFrom
	userNameFrom, pErr := m.ParseMailAddr(task.EmailFrom)
	if pErr != nil {
		taskResult.Error = pErr.Error()
		return
	}
	userNameTo, pErr := m.ParseMailAddr(task.EmailTo)
	if pErr != nil {
		taskResult.Error = pErr.Error()
		return
	}
	m.UserSmtp(taskResult, task, userNameFrom, userNameTo)
}

func (m *Mail) UserSmtp(taskResult *TaskResult, task *Task, userNameFrom, userNameTo string) {
	header := make(map[string]string)
	header["From"] = mime.BEncoding.Encode("utf8", userNameFrom) + "<" + task.EmailFrom + ">"
	header["To"] = task.EmailTo
	header["Subject"] = mime.BEncoding.Encode("utf8", task.EmailSubject)
	header["MIME-Version"] = "1.0"
	header["X-Has-Attach"] = "no"
	header["X-Mailer"] = "Foxmail 7.2.13.365[cn]"
	header["Message-ID"] = m.MessageId(task)
	header["Date"] = time.Now().Local().Format(time.RFC1123Z)
	boundary := m.ContentType(&header)
	mutilBody, mpErr := m.MutilPartBody([]byte(task.EmailBody), boundary)
	if mpErr != nil {
		taskResult.Error = mpErr.Error()
		return
	}
	messages := []string{}
	for k, v := range header {
		messages = append(messages, fmt.Sprintf("%s:%s", k, v))
	}
	messages = append(messages, "\r\n")
	messages = append(messages, mutilBody)
	message := strings.Join(messages, "\r\n")
	var auth smtp.Auth
	if task.SmtpAuthType == "login" {
		if task.EmailAuthCode != "" {
			auth = smtp.LoginAuth(task.EmailFrom, task.EmailAuthCode, task.SmtpHost)
		} else {
			auth = smtp.LoginAuth(task.EmailFrom, task.EmailAuthPass, task.SmtpHost)
		}
	} else {
		if task.EmailAuthCode != "" {
			auth = smtp.PlainAuth("", task.EmailFrom, task.EmailAuthCode, task.SmtpHost)
		} else {
			auth = smtp.PlainAuth("", task.EmailFrom, task.EmailAuthPass, task.SmtpHost)
		}
	}
	var err error
	if task.UseProxy {
		if task.ProxyKeepAlive {
			err = m.SendMailWithAliveProxy(task, auth, task.EmailFrom, []string{task.EmailTo}, []byte(message), task.SmtpSSL)
		} else {
			proxyAuth := &socket5{
				Auth:    nil,
				Network: "tcp",
				Host:    task.ProxyHost,
				Port:    task.ProxyPort,
			}
			if task.ProxyAuth {
				proxyAuth.Auth = &proxy.Auth{
					User:     task.ProxyAuthUser,
					Password: task.ProxyAuthPass,
				}
			}
			err = m.SendMailWithProxy(task.SmtpHost, task.SmtpPort, auth, task.EmailFrom, []string{task.EmailTo}, []byte(message), proxyAuth, task.SmtpSSL)
		}
	} else {
		err = m.SendMailDirect(task.SmtpHost, task.SmtpPort, auth, task.EmailFrom, []string{task.EmailTo}, []byte(message), task.SmtpSSL)
	}
	if task.UseProxy && err != nil {
		err = m.SendMailDirect(task.SmtpHost, task.SmtpPort, auth, task.EmailFrom, []string{task.EmailTo}, []byte(message), task.SmtpSSL)
	}
	if err != nil {
		taskResult.Error = err.Error()
		return
	}
	return
}

func (m *Mail) MutilPartBody(body []byte, boundary string) (string, error) {
	var buf bytes.Buffer
	w := multipart.NewWriter(&buf)
	if err := w.SetBoundary(boundary); err != nil {
		return "", err
	}
	header := textproto.MIMEHeader{
		"Content-Type":              {"text/html;\r\n\tcharset=\"utf-8\""},
		"Content-Transfer-Encoding": {"base64"},
	}
	part, err := w.CreatePart(header)
	if err != nil {
		return "", err
	}
	encodeString := base64.StdEncoding.EncodeToString(body)
	part.Write([]byte(encodeString))
	w.Close()
	return buf.String(), nil
}

func (m *Mail) MessageId(task *Task) string {
	at := m.MailAt(task.EmailFrom)
	messageId := fmt.Sprintf("<%v%v>+%v", time.Now().UnixNano(), at, utils.RandIntString(16))
	return messageId
}

func (m *Mail) ContentType(header *map[string]string) string {
	bd := m.Boundary()
	(*header)["Content-Type"] = fmt.Sprintf("multipart/alternative;\r\n\tboundary=%v", bd)
	(*header)["Content-Transfer-Encoding"] = "8Bit"
	return bd
}

func (m *Mail) Boundary() string {
	boundary := fmt.Sprintf(`----=_NextPart_%v_%v_%v`, utils.RandIntString(8), utils.RandIntString(8), utils.RandIntString(8))
	return boundary
}

func (m *Mail) MailAt(email string) string {
	i := strings.Index(email, `@`)
	if i == -1 {
		return email
	}
	return email[i:]
}

func (m *Mail) DialProxyTLS(host, port string, proxyAuth *socket5) (*net.Conn, *tls.Conn, *smtp.Client, error) {
	dialer, err := proxy.SOCKS5(
		proxyAuth.Network,
		proxyAuth.Host+":"+proxyAuth.Port,
		proxyAuth.Auth,
		proxy.Direct,
	)
	if err != nil {
		return nil, nil, nil, err
	}
	proxyConn, proxyErr := dialer.Dial("tcp", host+":"+port)
	if proxyErr != nil {
		return nil, nil, nil, proxyErr
	}
	tlsConn := tls.Client(proxyConn, &tls.Config{InsecureSkipVerify: true})
	smtpConn, err := smtp.NewClient(tlsConn, host)
	return &proxyConn, tlsConn, smtpConn, err
}

func (m *Mail) DialProxyAliveTLS(t *Task) (*tls.Conn, *smtp.Client, error) {
	var proxyConn net.Conn
	var proxyErr error
	for i := 1; i <= 3; i++ {
		dialer := PrManager.GetDial(t)
		if dialer == nil {
			return nil, nil, errors.New("proxy dial nil")
		}
		proxyConn, proxyErr = (*dialer).Dial("tcp", t.SmtpHost+":"+t.SmtpPort)
		if proxyErr != nil {
			if m.DialError(proxyErr) {
				go PrManager.RecvNew(t)
				continue
			}
			return nil, nil, proxyErr
		}
		if proxyConn == nil {
			// conf.MclientLog.Errorf("dial return nil, try: %v", i)
			config.Log.Errorf("dial return nil, try: %v", i)
			continue
		}
		break
	}
	if proxyConn == nil {
		if proxyErr != nil {
			return nil, nil, proxyErr
		}
		return nil, nil, errors.New("proxy dial nil")
	}
	tlsConn := tls.Client(proxyConn, &tls.Config{InsecureSkipVerify: true})
	smtpConn, err := smtp.NewClient(tlsConn, t.SmtpHost)
	return tlsConn, smtpConn, err
}

func (m *Mail) DialProxyAliveNoTLS(t *Task) (*smtp.Client, error) {
	var proxyConn net.Conn
	var proxyErr error
	for i := 1; i <= 3; i++ {
		dialer := PrManager.GetDial(t)
		if dialer == nil {
			return nil, errors.New("proxy dial nil")
		}
		proxyConn, proxyErr = (*dialer).Dial("tcp", t.SmtpHost+":"+t.SmtpPort)
		if proxyErr != nil {
			if m.DialError(proxyErr) {
				config.Log.Errorf("err: %v, try: %v", proxyErr, i)
				go PrManager.RecvNew(t)
				continue
			}
			return nil, proxyErr
		}
		if proxyConn == nil {
			config.Log.Errorf("dial return nil, try: %v", i)
			continue
		}
		break
	}
	if proxyConn == nil {
		if proxyErr != nil {
			return nil, proxyErr
		}
		return nil, errors.New("proxy dial nil")
	}
	smtpConn, err := smtp.NewClient(proxyConn, t.SmtpHost)
	return smtpConn, err
}

func (m *Mail) DialError(dialErr error) bool {
	oPerr, ok1 := dialErr.(*net.OpError)
	if !ok1 {
		return false
	}
	netErr, ok2 := oPerr.Err.(*net.OpError)
	if !ok2 {
		return false
	}
	if netErr.Source == nil && netErr.Addr != nil {
		return true
	}
	return false
}

func (m *Mail) DialProxyNoTLS(host, port string, proxyAuth *socket5) (*net.Conn, *smtp.Client, error) {
	dialer, err := proxy.SOCKS5(
		proxyAuth.Network,
		proxyAuth.Host+":"+proxyAuth.Port,
		proxyAuth.Auth,
		proxy.Direct,
	)
	if err != nil {
		return nil, nil, err
	}
	proxyConn, proxyErr := dialer.Dial("tcp", host+":"+port)
	if proxyErr != nil {
		return nil, nil, proxyErr
	}
	smtpConn, err := smtp.NewClient(proxyConn, host)
	return &proxyConn, smtpConn, err
}

func (m *Mail) DialTLS(host, port string) (*tls.Conn, *smtp.Client, error) {
	conn, err := tls.Dial("tcp", host+":"+port, &tls.Config{InsecureSkipVerify: true})
	if err != nil {
		return nil, nil, err
	}
	smtpConn, sErr := smtp.NewClient(conn, host)
	return conn, smtpConn, sErr
}

func (m *Mail) SendMailWithAliveProxy(t *Task, auth smtp.Auth, from string, to []string, msg []byte, ssl bool) error {
	if ssl {
		tC, cC, err := m.DialProxyAliveTLS(t)
		if err != nil {
			return err
		}
		defer func() {
			if cC != nil {
				cC.Close()
			}
			if tC != nil {
				tC.Close()
			}
		}()
		return m.WithTLS(cC, t.SmtpHost, t.SmtpPort, auth, from, to, msg)
	} else {
		cC, err := m.DialProxyAliveNoTLS(t)
		if err != nil {
			return err
		}
		defer func() {
			if cC != nil {
				cC.Close()
			}
		}()
		err = smtp.WithNoTLS(cC, auth, from, to, msg)
		if err != nil {
			return err
		} else {
			return err
		}
	}
}

func (m *Mail) SendMailWithProxy(addr, port string, auth smtp.Auth, from string, to []string, msg []byte, proxyAuth *socket5, ssl bool) error {
	if ssl {
		p, t, c, err := m.DialProxyTLS(addr, port, proxyAuth)
		if err != nil {
			return err
		}
		defer func() {
			if c != nil {
				c.Close()
			}
			if t != nil {
				t.Close()
			}
			if p != nil {
				(*p).Close()
			}
		}()
		return m.WithTLS(c, addr, port, auth, from, to, msg)
	} else {
		p, c, err := m.DialProxyNoTLS(addr, port, proxyAuth)
		if err != nil {
			return err
		}
		defer func() {
			if c != nil {
				c.Close()
			}
			if p != nil {
				(*p).Close()
			}
		}()
		err = smtp.WithNoTLS(c, auth, from, to, msg)
		if err != nil {
			return err
		} else {
			return err
		}
	}
}

func (m *Mail) SendMailDirect(addr, port string, auth smtp.Auth, from string, to []string, msg []byte, ssl bool) error {
	if ssl {
		t, c, err := m.DialTLS(addr, port)
		if err != nil {
			return err
		}
		defer func() {
			if c != nil {
				c.Close()
			}
			if t != nil {
				t.Close()
			}
		}()

		return m.WithTLS(c, addr, port, auth, from, to, msg)
	} else {
		return m.SenderMail(addr, port, auth, from, to, msg)
	}
}

func (m *Mail) WithTLS(c *smtp.Client, addr, port string, auth smtp.Auth, from string, to []string, msg []byte) error {
	var err error
	if c == nil {
		return errors.New("Client is nil")
	}
	if auth != nil {
		if ok, _ := c.Extension("AUTH"); ok {
			if err = c.Auth(auth); err != nil {
				return err
			}
		}
	}
	if err = c.Mail(from); err != nil {
		return err
	}

	for _, addr := range to {
		if err = c.Rcpt(addr); err != nil {
			return err
		}
	}

	w, err := c.Data()
	if err != nil {
		return err
	}

	_, err = w.Write(msg)
	if err != nil {
		return err
	}

	err = w.Close()
	if err != nil {
		return err
	}
	return c.Quit()
}

func (m *Mail) SenderMail(addr, port string, auth smtp.Auth, from string, to []string, msg []byte) error {
	err := smtp.SendMail(addr+":"+port, auth, from, to, msg)
	if err != nil {
		return err
	} else {
		return nil
	}
}

func (m *Mail) ParseMailAddr(maddr string) (string, error) {
	if strings.Count(maddr, "@") != 1 {
		return "", errors.New("illegal mail address")
	}
	p := strings.Index(maddr, "@")
	userName := maddr[:p]
	return userName, nil
}
