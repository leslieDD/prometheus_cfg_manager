package router

import (
	"pro_cfg_manager/config"
	"pro_cfg_manager/models"

	"github.com/gin-gonic/gin"
)

func initOutputData() {
	v1.GET("/output/ip/list", outputIPlist)
	v1.POST("/import/cron/rule", cronImport)
	apiNoAuth.GET("/version", Version)
}

func outputIPlist(c *gin.Context) {
	user := c.Keys["userInfo"].(*models.UserSessionInfo)
	pass := models.CheckPriv(user, "ipManager", "", "download")
	if pass != models.Success {
		resComm(c, pass, nil)
		return
	}
	data, bf := models.OutputIPlist(c)
	models.OO.RecodeLog(user.Username, c.Request.RemoteAddr, "download cvs", models.IsSearch, bf)
	resComm(c, bf, data)
}

func cronImport(c *gin.Context) {
	cts := []*models.CrontabPost{}
	if err := c.BindJSON(&cts); err != nil {
		models.OO.RecodeLog("openApi", c.Request.RemoteAddr, "create cron api", models.IsAdd, models.ErrPostData)
		resComm(c, models.ErrPostData, nil)
		return
	}
	bf := models.CronImport(cts)
	resComm(c, bf, nil)
}

func Version(c *gin.Context) {
	resComm(c, models.Success, map[string]string{"version": config.Version})
}
