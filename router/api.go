package router

import (
	"pro_cfg_manager/config"
	"pro_cfg_manager/models"
	"strconv"

	"github.com/gin-gonic/gin"
)

func initApiRouter() {

	v1.GET("/jobs", getJobs)

	v1.GET("/machine", getMachine)
	v1.GET("/machines", getMachines)
	v1.POST("/machine", postMachine)
	v1.PUT("/machine", putMachine)
	v1.DELETE("/machine", deleteMachine)

	v1.POST("/publish", publish)
}

func getJobs(c *gin.Context) {
	result, bf := models.GetJobs()
	resComm(c, bf, result)
}

func getMachine(c *gin.Context) {
	mIDstr, ok := c.GetQuery("id")
	if !ok {
		resComm(c, models.ErrQueryData, nil)
		return
	}
	mID, err := strconv.ParseInt(mIDstr, 10, 0)
	if err != nil {
		config.Log.Error(err)
		resComm(c, models.ErrQueryData, nil)
		return
	}
	result, bf := models.GetMachine(int(mID))
	resComm(c, bf, result)
}

func getMachines(c *gin.Context) {
	sp := &models.SplitPage{}
	if err := c.BindQuery(sp); err != nil {
		resComm(c, models.ErrSplitParma, nil)
		return
	}
	result, bf := models.GetMachines(sp)
	resComm(c, bf, result)
}

func postMachine(c *gin.Context) {
	mInfo := &models.Machine{}
	if err := c.BindJSON(mInfo); err != nil {
		resComm(c, models.ErrPostData, nil)
		return
	}
	bf := models.PostMachine(mInfo)
	resComm(c, bf, nil)
}

func putMachine(c *gin.Context) {
	mInfo := &models.Machine{}
	if err := c.BindJSON(mInfo); err != nil {
		resComm(c, models.ErrPostData, nil)
		return
	}
	bf := models.PutMachine(mInfo)
	resComm(c, bf, nil)
}

func deleteMachine(c *gin.Context) {
	mIDstr, ok := c.GetQuery("id")
	if !ok {
		resComm(c, models.ErrQueryData, nil)
		return
	}
	mID, err := strconv.ParseInt(mIDstr, 10, 0)
	if err != nil {
		config.Log.Error(err)
		resComm(c, models.ErrQueryData, nil)
		return
	}
	bf := models.DeleteMachine(int(mID))
	resComm(c, bf, nil)
}

func publish(c *gin.Context) {
	bf := models.Publish()
	resComm(c, bf, nil)
}
