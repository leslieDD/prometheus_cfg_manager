package models

import (
	"fmt"
	"net/http"
	"pro_cfg_manager/config"
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
	go send(ws)
	for {
		//读取ws中的数据
		mt, message, err := ws.ReadMessage()
		if err != nil {
			break
		}
		config.Log.Warnf("%d, %s", mt, message)
		// if string(message) == "ping" {
		// 	message = []byte("pong")
		// }
		message = []byte(`{"action": "pong"}`)
		//写入ws数据
		err = ws.WriteMessage(mt, message)
		if err != nil {
			config.Log.Error(err)
			break
		}
	}
	return nil
}

func send(ws *websocket.Conn) {
	for {
		message := []byte(fmt.Sprintf(`{"action": "%s"}`, time.Now()))
		ws.WriteMessage(1, message)
		time.Sleep(1 * time.Second)
	}
}
