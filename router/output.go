package router

import (
	"pro_cfg_manager/models"

	"github.com/gin-gonic/gin"
)

func initOutputData() {
	v1.GET("/output/ip/list", outputIPlist)
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
