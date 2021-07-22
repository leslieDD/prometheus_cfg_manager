package router

import (
	"pro_cfg_manager/config"
	"pro_cfg_manager/models"
	"strconv"

	"github.com/gin-gonic/gin"
)

func initManagerApi() {
	v1.GET("/manager/groups", getManagerGroups)
	v1.POST("/manager/group", postManagerGroup)
	v1.PUT("/manager/group", putManagerGroup)
	v1.DELETE("/manager/group", deleteManagerGroup)
	v1.PUT("/manager/group/status", putManagerGroupStatus)
}

func getManagerGroups(c *gin.Context) {
	sp := &models.SplitPage{}
	if err := c.BindQuery(sp); err != nil {
		resComm(c, models.ErrSplitParma, nil)
		return
	}
	result, bf := models.GetManagerGroups(sp)
	resComm(c, bf, result)
}

func postManagerGroup(c *gin.Context) {
	mg := models.ManagerGroup{}
	if err := c.BindJSON(&mg); err != nil {
		resComm(c, models.ErrPostData, nil)
		return
	}
	bf := models.PostManagerGroup(&mg)
	resComm(c, bf, nil)
}

func putManagerGroup(c *gin.Context) {
	mg := models.ManagerGroup{}
	if err := c.BindJSON(&mg); err != nil {
		resComm(c, models.ErrPostData, nil)
		return
	}
	bf := models.PutManagerGroup(&mg)
	resComm(c, bf, nil)
}

func deleteManagerGroup(c *gin.Context) {
	gIDStr, ok := c.GetQuery("id")
	if !ok {
		resComm(c, models.ErrQueryData, nil)
		return
	}
	gID, err := strconv.ParseInt(gIDStr, 10, 0)
	if err != nil {
		config.Log.Error(err)
		resComm(c, models.ErrQueryData, nil)
		return
	}
	bf := models.DeleteManagerGroup(gID)
	resComm(c, bf, nil)
}

func putManagerGroupStatus(c *gin.Context) {
	enabledInfo := models.EnabledInfo{}
	if err := c.BindJSON(&enabledInfo); err != nil {
		resComm(c, models.ErrPostData, nil)
		return
	}
	bf := models.PutManagerGroupStatus(&enabledInfo)
	resComm(c, bf, nil)
}
