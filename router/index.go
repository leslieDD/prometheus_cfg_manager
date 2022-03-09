package router

import (
	"net/http"
	"pro_cfg_manager/models"

	"github.com/gin-gonic/gin"
)

func initCommonRouter() {
	GinDefault.GET("/", index)
	GinDefault.POST("/login", login)
	GinDefault.POST("/register", register)
}

func index(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", gin.H{})
}

func login(c *gin.Context) {
	ui := models.UserLogInfo{}
	if err := c.BindJSON(&ui); err != nil {
		resComm(c, models.ErrPostData, nil)
		return
	}
	data, bf := models.Login(&ui, c.Request.RemoteAddr)
	models.OO.RecodeLog(ui.Username, c.Request.RemoteAddr, "login", models.IsLogin, bf)
	resComm(c, bf, data)
}

func register(c *gin.Context) {
	r := models.RegisterInfo{}
	if err := c.BindJSON(&r); err != nil {
		resComm(c, models.ErrPostData, nil)
		return
	}
	bf := models.Register(&r)
	models.OO.RecodeLog("", c.Request.RemoteAddr, "register", models.IsLogin, bf)
	resComm(c, bf, nil)
}
