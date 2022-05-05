package router

import (
	"pro_cfg_manager/config"
	"pro_cfg_manager/models"

	"github.com/gin-gonic/gin"
)

func initMonitorApi() {
	v1.GET("/instance/targets", getInstanceTargets)
	v1.PUT("/instance/targets", putInstanceTargets)
}

func getInstanceTargets(c *gin.Context) {
	user := c.Keys["userInfo"].(*models.UserSessionInfo)
	// pass := models.CheckPriv(user, "instance", "", "get_data")
	// if pass != models.Success {
	// 	resComm(c, pass, nil)
	// 	return
	// }
	itr := models.InstanceTargetsReq{}
	if err := c.BindQuery(&itr); err != nil {
		config.Log.Error(err)
		models.OO.RecodeLog(user.Username, c.Request.RemoteAddr, "get instance targets", models.IsSearch, models.ErrQueryData)
		resComm(c, models.ErrQueryData, nil)
	}
	result, bf := models.GetInstanceTargets(&itr)
	models.OO.RecodeLog(user.Username, c.Request.RemoteAddr, "get instance targets", models.IsSearch, bf)
	resComm(c, bf, result)
}

func putInstanceTargets(c *gin.Context) {
	user := c.Keys["userInfo"].(*models.UserSessionInfo)
	pass := models.CheckPriv(user, "instance", "", "put_data")
	if pass != models.Success {
		resComm(c, pass, nil)
		return
	}
}
