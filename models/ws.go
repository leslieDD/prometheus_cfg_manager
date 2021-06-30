package models

import (
	"fmt"
	"net/http"
	"path/filepath"
	"pro_cfg_manager/config"
	"pro_cfg_manager/utils"
	"runtime"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var upGrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func WS(c *gin.Context) *BriefMessage {
	ws, err := upGrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		return ErrUpGrader
	}
	defer ws.Close()
	// for {
	// 	//读取ws中的数据
	// 	// mt, message, err := ws.ReadMessage()
	// 	// if err != nil {
	// 	// 	config.Log.Error(err)
	// 	// 	break
	// 	// }
	// 	// message = []byte(`{"action": "pong"}`)
	// 	statusCode, output, err := doCheck()
	// 	message := fmt.Sprintf("状态码：%d\n输出结果：%s\n错误信息：%s\n",
	// 		statusCode, output, err)
	// 	//写入ws数据
	// 	err = ws.WriteMessage(websocket.BinaryMessage, []byte(message))
	// 	if err != nil {
	// 		config.Log.Error(err)
	// 		break
	// 	}
	// 	err = ws.WriteControl(websocket.CloseMessage, []byte(""), time.Now().Add(5*time.Second))
	// 	if err != nil {
	// 		config.Log.Error(err)
	// 	}
	// 	break
	// }
	// return nil
	tag := []byte("测试中，请稍等...\n--------------------------\n")
	err = ws.WriteMessage(websocket.TextMessage, tag)
	if err != nil {
		config.Log.Error(err)
		return nil
	}
	statusCode, output, cmd, err := doCheck()
	message := fmt.Sprintf("执行命令：%s\n状态码：%d\n输出结果：%s\n错误信息：%s\n",
		cmd, statusCode, output, err)
	//写入ws数据
	err = ws.WriteMessage(websocket.TextMessage, []byte(message))
	if err != nil {
		config.Log.Error(err)
		return nil
	}
	tag = []byte("--------------------------\n测试完成\n")
	err = ws.WriteMessage(websocket.TextMessage, tag)
	if err != nil {
		config.Log.Error(err)
		return nil
	}
	err = ws.WriteControl(websocket.CloseMessage, []byte(""), time.Now().Add(5*time.Second))
	if err != nil {
		config.Log.Error(err)
	}
	return nil
}

func doCheck() (int, string, string, error) {
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
	cmd = fmt.Sprintf("%s check config %s", cmd, config.Cfg.PrometheusCfg.MainConf)
	return statusCode, output, cmd, err
}
