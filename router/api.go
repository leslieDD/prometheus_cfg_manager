package router

import (
	"pro_cfg_manager/config"
	"pro_cfg_manager/models"
	"strconv"

	"github.com/gin-gonic/gin"
)

func initApiRouter() {

	v1.GET("/jobs", getJobs)
	v1.GET("/jobs/split", getJobsSplit)
	v1.GET("/job", getJob)
	v1.POST("/job", postJob)
	v1.PUT("/job", putJob)
	v1.DELETE("/job", deleteJob)
	v1.PUT("/job/swap", swapJob)
	v1.POST("/jobs/publish", publishJobs)

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

func getJobsSplit(c *gin.Context) {
	sp := &models.SplitPage{}
	if err := c.BindQuery(sp); err != nil {
		resComm(c, models.ErrSplitParma, nil)
		return
	}
	result, bf := models.GetJobsSplit(sp)
	resComm(c, bf, result)
}

func getJob(c *gin.Context) {
	jIDstr, ok := c.GetQuery("id")
	if !ok {
		resComm(c, models.ErrQueryData, nil)
		return
	}
	jID, err := strconv.ParseInt(jIDstr, 10, 0)
	if err != nil {
		config.Log.Error(err)
		resComm(c, models.ErrQueryData, nil)
		return
	}
	result, bf := models.GetJob(jID)
	resComm(c, bf, result)
}

func postJob(c *gin.Context) {
	jInfo := &models.Jobs{}
	if err := c.BindJSON(jInfo); err != nil {
		resComm(c, models.ErrPostData, nil)
		return
	}
	bf := models.PostJob(jInfo)
	resComm(c, bf, nil)
}

func putJob(c *gin.Context) {
	jInfo := &models.Jobs{}
	if err := c.BindJSON(jInfo); err != nil {
		resComm(c, models.ErrPostData, nil)
		return
	}
	bf := models.PutJob(jInfo)
	resComm(c, bf, nil)
}

func deleteJob(c *gin.Context) {
	jIDstr, ok := c.GetQuery("id")
	if !ok {
		resComm(c, models.ErrQueryData, nil)
		return
	}
	jID, err := strconv.ParseInt(jIDstr, 10, 0)
	if err != nil {
		config.Log.Error(err)
		resComm(c, models.ErrQueryData, nil)
		return
	}
	bf := models.DeleteJob(jID)
	resComm(c, bf, nil)
}

func swapJob(c *gin.Context) {
	sInfo := &models.SwapInfo{}
	if err := c.BindJSON(sInfo); err != nil {
		resComm(c, models.ErrPostData, nil)
		return
	}
	bf := models.DoSwap(sInfo)
	resComm(c, bf, nil)
}

func publishJobs(c *gin.Context) {
	bf := models.DoPublishJobs()
	resComm(c, bf, nil)
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
