package router

import (
	"net/http"
	"pro_cfg_manager/models"

	"github.com/gin-gonic/gin"
)

func initCommonRouter() {
	GinDefault.GET("/", index)
	GinDefault.POST("/login", login)
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
	data, bf := models.Login(&ui)
	resComm(c, bf, data)
}
