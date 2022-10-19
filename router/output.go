package router

import (
	"pro_cfg_manager/config"
	"pro_cfg_manager/models"

	"github.com/gin-gonic/gin"
)

func initOutputData() {
	v1.GET("/output/ip/list", outputIPlist)
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
	
}

func Version(c *gin.Context) {
	resComm(c, models.Success, map[string]string{"version": config.Version})
}
