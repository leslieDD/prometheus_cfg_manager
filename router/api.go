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

	v1.GET("/jobs/default/split", getDefJobsSplit)
	v1.POST("/job/default", postDefJob)
	v1.PUT("/job/default", putDefJob)
	v1.DELETE("/job/default", deleteDefJob)

	v1.GET("/machine", getMachine)
	v1.GET("/machines", getMachines)
	v1.POST("/machine", postMachine)
	v1.PUT("/machine", putMachine)
	v1.DELETE("/machine", deleteMachine)

	v1.POST("/publish", publish)
	v1.POST("/reload", reload)
	v1.GET("/preview", preview)

	v1.GET("/load/all-file-list", allFileList)
	v1.POST("/load/file-content", fileContent)
	v1.GET("/load/all-Rulefile-list", allRulesFileList)
	v1.POST("/load/rule-file-content", ruleFileContent)

	v1.GET("/tree", getTree)
	v1.GET("/tree/node", getNode)
	v1.POST("/tree/node", postNode)
	v1.PUT("/tree/node", putNode)
	v1.GET("/tree/default/lables", getDefLabels)
	v1.DELETE("/tree/node/label", delLabel)
	v1.POST("/tree/create/node", createTreeNode)
	v1.PUT("/tree/update/node", updateTreeNode)
	v1.DELETE("/tree/remove/node", deleteTreeNode)
	v1.POST("/rules/publish", rulePublish)

	v1.GET("/base/labels", getBaseLabels)
	v1.POST("/base/labels", postBaseLabels)
	v1.PUT("/base/labels", putBaseLabels)
	v1.DELETE("/base/labels", delBaseLabels)
	v1.GET("/base/relabels", getReLabels)
	v1.GET("/base/relabels/all", getAllRelabels)
	v1.POST("/base/relabels", postReLabels)
	v1.PUT("/base/relabels", putReLabels)
	v1.DELETE("/base/relabels", delReLabels)
	v1.PUT("base/relabels/code", putRelabelsCode)
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

func getDefJobsSplit(c *gin.Context) {
	sp := &models.SplitPage{}
	if err := c.BindQuery(sp); err != nil {
		resComm(c, models.ErrSplitParma, nil)
		return
	}
	result, bf := models.GetDefJobsSplit(sp)
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

func postDefJob(c *gin.Context) {
	jInfo := &models.Jobs{}
	if err := c.BindJSON(jInfo); err != nil {
		resComm(c, models.ErrPostData, nil)
		return
	}
	jInfo.IsCommon = true
	bf := models.PostJob(jInfo)
	resComm(c, bf, nil)
}

func putJob(c *gin.Context) {
	jInfo := &models.Jobs{}
	if err := c.BindJSON(jInfo); err != nil {
		resComm(c, models.ErrPostData, nil)
		return
	}
	jInfo.IsCommon = false
	bf := models.PutJob(jInfo)
	resComm(c, bf, nil)
}

func putDefJob(c *gin.Context) {
	jInfo := &models.Jobs{}
	if err := c.BindJSON(jInfo); err != nil {
		resComm(c, models.ErrPostData, nil)
		return
	}
	jInfo.IsCommon = true
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
	bf := models.DeleteJob(jID, false)
	resComm(c, bf, nil)
}

func deleteDefJob(c *gin.Context) {
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
	bf := models.DeleteJob(jID, true)
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
	// result, bf := models.GetMachines(sp)
	result, bf := models.GetMachinesV2(sp)
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

func preview(c *gin.Context) {
	data, bf := models.Preview()
	resComm(c, bf, data)
}

func reload(c *gin.Context) {
	bf := models.Reload()
	resComm(c, bf, nil)
}

func allFileList(c *gin.Context) {
	data, bf := models.AllFileList()
	resComm(c, bf, data)
}

func fileContent(c *gin.Context) {
	child := models.Child{}
	if err := c.BindJSON(&child); err != nil {
		resComm(c, models.ErrSplitParma, nil)
		return
	}
	data, bf := models.ReadFileContent(&child)
	resComm(c, bf, data)
}

func allRulesFileList(c *gin.Context) {
	data, bf := models.AllRulesFileList()
	resComm(c, bf, data)
}

func ruleFileContent(c *gin.Context) {
	child := models.Child{}
	if err := c.BindJSON(&child); err != nil {
		resComm(c, models.ErrSplitParma, nil)
		return
	}
	data, bf := models.ReadRuleFileContent(&child)
	resComm(c, bf, data)
}

func getTree(c *gin.Context) {
	data, bf := models.GetNodesFromDB()
	resComm(c, bf, data)
}

func getNode(c *gin.Context) {
	qni := models.QueryGetNode{}
	if err := c.BindQuery(&qni); err != nil {
		resComm(c, models.ErrQueryData, nil)
		return
	}
	data, bf := models.GetNode(&qni)
	resComm(c, bf, data)
}

func postNode(c *gin.Context) {
	nodeInfo := models.TreeNodeInfo{}
	if err := c.BindJSON(&nodeInfo); err != nil {
		resComm(c, models.ErrPostData, nil)
		return
	}
	data, bf := models.PostNode(&nodeInfo)
	resComm(c, bf, data)
}

func putNode(c *gin.Context) {
	nodeInfo := models.TreeNodeInfo{}
	if err := c.BindJSON(&nodeInfo); err != nil {
		resComm(c, models.ErrPostData, nil)
		return
	}
	data, bf := models.PutNode(&nodeInfo)
	resComm(c, bf, data)
}

func getDefLabels(c *gin.Context) {
	data, bf := models.GetDefLabels()
	resComm(c, bf, data)
}

func delLabel(c *gin.Context) {
	labelType, ok := c.GetQuery("type")
	if !ok {
		resComm(c, models.ErrQueryData, nil)
		return
	}
	labelInfo := models.Label{}
	if err := c.BindJSON(&labelInfo); err != nil {
		resComm(c, models.ErrPostData, nil)
		return
	}
	bf := models.DelLabel(&labelInfo, labelType)
	resComm(c, bf, nil)
}

func createTreeNode(c *gin.Context) {
	tnc := models.TreeNodeFromCli{}
	if err := c.BindJSON(&tnc); err != nil {
		resComm(c, models.ErrPostData, nil)
		return
	}
	bf := models.CreateTreeNode(&tnc)
	resComm(c, bf, nil)
}

func updateTreeNode(c *gin.Context) {
	tnc := models.TreeNodeFromCli{}
	if err := c.BindJSON(&tnc); err != nil {
		resComm(c, models.ErrPostData, nil)
		return
	}
	bf := models.UpdateTreeNode(&tnc)
	resComm(c, bf, nil)
}

func deleteTreeNode(c *gin.Context) {
	tnc := models.TreeNodeFromCli{}
	if err := c.BindJSON(&tnc); err != nil {
		resComm(c, models.ErrPostData, nil)
		return
	}
	bf := models.DeleteTreeNode(&tnc)
	resComm(c, bf, nil)
}

func rulePublish(c *gin.Context) {
	bf := models.RulePublish()
	resComm(c, bf, nil)
}

func getBaseLabels(c *gin.Context) {
	sp := &models.SplitPage{}
	if err := c.BindQuery(sp); err != nil {
		resComm(c, models.ErrSplitParma, nil)
		return
	}
	data, bf := models.GetBaseLabels(sp)
	resComm(c, bf, data)
}

func postBaseLabels(c *gin.Context) {
	newLabel := models.BaseLabels{}
	if err := c.BindJSON(&newLabel); err != nil {
		resComm(c, models.ErrPostData, nil)
		return
	}
	bf := models.PostBaseLabels(&newLabel)
	resComm(c, bf, nil)
}

func putBaseLabels(c *gin.Context) {
	label := models.BaseLabels{}
	if err := c.BindJSON(&label); err != nil {
		resComm(c, models.ErrPostData, nil)
		return
	}
	bf := models.PutBaseLabels(&label)
	resComm(c, bf, nil)
}

func delBaseLabels(c *gin.Context) {
	idStr, ok := c.GetQuery("id")
	if !ok {
		resComm(c, models.ErrQueryData, nil)
		return
	}
	id, err := strconv.ParseInt(idStr, 10, 0)
	if err != nil {
		config.Log.Error(err)
		resComm(c, models.ErrQueryData, nil)
		return
	}
	bf := models.DelBaseLabels(id)
	resComm(c, bf, nil)
}

func getReLabels(c *gin.Context) {
	sp := &models.SplitPage{}
	if err := c.BindQuery(sp); err != nil {
		resComm(c, models.ErrSplitParma, nil)
		return
	}
	data, bf := models.GetReLabels(sp)
	resComm(c, bf, data)
}

func getAllRelabels(c *gin.Context) {
	data, bf := models.GetAllReLabels()
	resComm(c, bf, data)
}

func postReLabels(c *gin.Context) {
	reLabel := models.ReLabels{}
	if err := c.BindJSON(&reLabel); err != nil {
		resComm(c, models.ErrPostData, nil)
		return
	}
	bf := models.PostReLabels(&reLabel)
	resComm(c, bf, nil)
}

func putReLabels(c *gin.Context) {
	reLabel := models.ReLabels{}
	if err := c.BindJSON(&reLabel); err != nil {
		resComm(c, models.ErrPostData, nil)
		return
	}
	bf := models.PutReLabels(&reLabel)
	resComm(c, bf, nil)
}

func delReLabels(c *gin.Context) {
	idStr, ok := c.GetQuery("id")
	if !ok {
		resComm(c, models.ErrQueryData, nil)
		return
	}
	id, err := strconv.ParseInt(idStr, 10, 0)
	if err != nil {
		config.Log.Error(err)
		resComm(c, models.ErrQueryData, nil)
		return
	}
	bf := models.DelReLabels(id)
	resComm(c, bf, nil)
}

func putRelabelsCode(c *gin.Context) {
	reLabel := models.ReLabels{}
	if err := c.BindJSON(&reLabel); err != nil {
		resComm(c, models.ErrPostData, nil)
		return
	}
	bf := models.PutReLabelsCode(&reLabel)
	resComm(c, bf, nil)
}
