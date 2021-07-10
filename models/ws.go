package models

import (
	"fmt"
	"net/http"
	"path/filepath"
	"pro_cfg_manager/config"
	"pro_cfg_manager/utils"
	"runtime"
	"sync"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var upGrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

type wsCommand struct {
	Cmd    string
	Result chan string
}

func WS(c *gin.Context) *BriefMessage {
	ws, err := upGrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		return ErrUpGrader
	}
	defer ws.Close()
	for {
		mt, message, err := ws.ReadMessage()
		if err != nil {
			config.Log.Error(err)
			break
		}
		wsc := &wsCommand{Cmd: string(message), Result: make(chan string)}
		bf := cmdArea.RecvFromWS(wsc)
		if bf != Success {
			err = ws.WriteMessage(mt, []byte(bf.Message))
			if err != nil {
				config.Log.Error(err)
			}
			continue
		}
		go func() {
			for output := range wsc.Result {
				err = ws.WriteMessage(mt, []byte(output))
				if err != nil {
					config.Log.Error(err)
					return
				}
			}
		}()
	}
	return nil
}

type CmdAreaT struct {
	reloadLock  sync.Mutex
	restartLock sync.Mutex
	messRecv    chan *wsCommand
	messSend    chan []byte
}

var cmdArea = NewCmdArea()

func NewCmdArea() *CmdAreaT {
	cmdArea := &CmdAreaT{
		reloadLock:  sync.Mutex{},
		restartLock: sync.Mutex{},
		messRecv:    make(chan *wsCommand),
		messSend:    make(chan []byte, 10),
	}
	go cmdArea.Work()
	return cmdArea
}

func (c *CmdAreaT) RecvFromWS(wsc *wsCommand) *BriefMessage {
	select {
	case c.messRecv <- wsc:
	default:
		return ErrHaveTaskRunning
	}
	return Success
}

func (c *CmdAreaT) Work() {
	for message := range c.messRecv {
		switch message.Cmd {
		case "restart":
			c.Restart(message)
		case "check":
			c.Check(message)
		default:
			config.Log.Error("unsupport cmd: %s", message.Cmd)
		}
	}
}

func (c *CmdAreaT) Restart(wsc *wsCommand) {
	c.restartLock.Lock()
	defer c.restartLock.Unlock()
	code, output, err := utils.ExecCmd("systemctl", "restart", "prometheus")
	if err != nil {
		config.Log.Errorf("err: %v: output: %v", err, output)
		wsc.Result <- fmt.Sprintf("err: %v: output: %v", err, output)
		return
	}
	if code != 0 {
		config.Log.Errorf("code: %d, output: %v", code, output)
		wsc.Result <- fmt.Sprintf("code: %d, output: %v", code, output)
		return
	}
	wsc.Result <- "重启成功"
	close(wsc.Result)
}

func (c *CmdAreaT) Check(wsc *wsCommand) {
	var promtool string
	if runtime.GOOS == "windows" {
		promtool = "promtool.exe"
	} else {
		promtool = "promtool"
	}
	cmd := filepath.Join(config.Cfg.PrometheusCfg.RootDir, promtool)
	statusCode, output, err := utils.ExecCmd(
		cmd,
		"check",
		"config",
		config.Cfg.PrometheusCfg.MainConf,
	)
	if err != nil {
		config.Log.Error(err)
	}
	wsc.Result <- "测试中，请稍等...\n--------------------------"
	cmd = fmt.Sprintf("%s check config %s", cmd, config.Cfg.PrometheusCfg.MainConf)
	wsc.Result <- fmt.Sprintf("执行命令：%s\n状态码：%d\n输出结果：%s\n错误信息：%s",
		cmd, statusCode, output, err)
	wsc.Result <- "--------------------------\n测试完成"
	close(wsc.Result)
}
