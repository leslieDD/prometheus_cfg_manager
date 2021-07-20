package router

import (
	"pro_cfg_manager/config"
	"pro_cfg_manager/models"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

func initApiRouter() {

	v1.GET("/test", getTest)

	v1.GET("/jobs", getJobs)
	v1.GET("/jobs/split", getJobsSplit)
	v1.GET("/job", getJob)
	v1.POST("/job", postJob)
	v1.PUT("/job", putJob)
	v1.DELETE("/job", deleteJob)
	v1.PUT("/job/swap", swapJob)
	v1.POST("/jobs/publish", publishJobs)
	v1.PUT("/jobs/status", putJobsStatus)
	v1.POST("/jobs/update-ips", postUpdateJobIPs)

	v1.GET("/jobs/default/split", getDefJobsSplit)
	v1.POST("/job/default", postDefJob)
	v1.PUT("/job/default", putDefJob)
	v1.DELETE("/job/default", deleteDefJob)
	v1.PUT("/job/default/status", putJobDefaultStatus)

	v1.GET("/job/group", getJobGroup)
	v1.POST("/job/group", postJobGroup)
	v1.PUT("/job/group", putJobGroup)
	v1.DELETE("/job/group", delJobGroup)
	v1.GET("/job/machines", getJobMachines)
	v1.GET("/job/group/machines", getJobGroupMachines)
	v1.PUT("/job/group/machines", putJobGroupMachines)
	v1.GET("/job/group/labels", getGroupLabels)
	v1.PUT("/job/group/labels", putGroupLabels)
	v1.DELETE("/job/group/labels", delGroupLabels)
	v1.GET("/job/group/machines-and-labels", getAllMachinesLabels)
	v1.PUT("/job/group/status", putJobGroupStatus)
	v1.PUT("/job/group/labels/status", putJobGroupLabelsStatus)

	v1.GET("/machine", getMachine)
	v1.GET("/machines", getMachines)
	v1.POST("/machine", postMachine)
	v1.PUT("/machine", putMachine)
	v1.DELETE("/machine", deleteMachine)
	v1.PUT("/machine/status", putMachineStatus)
	v1.GET("/machines/all", getAllMachines)
	v1.DELETE("/machines/selection", batchDeleteMachine)
	v1.POST("/upload/machines", uploadMachines)

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
	v1.GET("/tree/default/enable/lables", getDefEnableLabels)
	v1.DELETE("/tree/node/label", delLabel)
	v1.POST("/tree/create/node", createTreeNode)
	v1.PUT("/tree/update/node", updateTreeNode)
	v1.DELETE("/tree/remove/node", deleteTreeNode)
	v1.POST("/rules/publish", rulePublish)
	v1.PUT("/tree/node/status", putTreeNodeStatus)
	v1.POST("/tree/upload/file/yaml", postTreeUploadFileYaml)

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
	v1.PUT("/base/labels/status", putBaseLabelsStatus)
	v1.PUT("/base/relabels/status", putBaseRelabelsStatus)

	v1.GET("/prometheus/tmpl", getProTmpl)
	v1.PUT("/prometheus/tmpl", putProTmpl)
	v1.GET("/prometheus/struct", getStruct)

	v1.GET("/ws", ws)

	v1.GET("/options", getOptions)
	v1.PUT("/options", putOptions)
}

func getTest(c *gin.Context) {
	result, bf := models.GetTest()
	resComm(c, bf, result)
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

func putJobDefaultStatus(c *gin.Context) {
	ed := models.EnabledInfo{}
	if err := c.BindJSON(&ed); err != nil {
		resComm(c, models.ErrPostData, nil)
		return
	}
	bf := models.PutJobDefaultStatus(&ed)
	resComm(c, bf, nil)
}

func getJobGroup(c *gin.Context) {
	sp := &models.SplitPage{}
	if err := c.BindQuery(sp); err != nil {
		resComm(c, models.ErrSplitParma, nil)
		return
	}
	data, bf := models.GetJobGroup(sp)
	resComm(c, bf, data)
}

func postJobGroup(c *gin.Context) {
	info := &models.JobGroup{}
	if err := c.BindJSON(info); err != nil {
		resComm(c, models.ErrPostData, nil)
		return
	}
	bf := models.PostJobGroup(info)
	resComm(c, bf, nil)
}

func putJobGroup(c *gin.Context) {
	info := &models.JobGroup{}
	if err := c.BindJSON(info); err != nil {
		resComm(c, models.ErrPostData, nil)
		return
	}
	bf := models.PutJobGroup(info)
	resComm(c, bf, nil)
}

func delJobGroup(c *gin.Context) {
	info := &models.DelJobGroupInfo{}
	if err := c.BindQuery(info); err != nil {
		resComm(c, models.ErrPostData, nil)
		return
	}
	bf := models.DelJobGroup(info)
	resComm(c, bf, nil)
}

func getJobMachines(c *gin.Context) {
	idStr, ok := c.GetQuery("id")
	if !ok {
		resComm(c, models.ErrQueryData, nil)
		return
	}
	jID, err := strconv.ParseInt(idStr, 10, 0)
	if err != nil {
		config.Log.Error(err)
		resComm(c, models.ErrQueryData, nil)
		return
	}
	data, bf := models.GetJobMachines(jID)
	resComm(c, bf, data)
}

func getJobGroupMachines(c *gin.Context) {
	idStr, ok := c.GetQuery("id")
	if !ok {
		resComm(c, models.ErrQueryData, nil)
		return
	}
	jID, err := strconv.ParseInt(idStr, 10, 0)
	if err != nil {
		config.Log.Error(err)
		resComm(c, models.ErrQueryData, nil)
		return
	}
	data, bf := models.GetJobGroupMachines(jID)
	resComm(c, bf, data)
}

func putJobGroupMachines(c *gin.Context) {
	idStr, ok := c.GetQuery("id")
	if !ok {
		resComm(c, models.ErrQueryData, nil)
		return
	}
	jID, err := strconv.ParseInt(idStr, 10, 0)
	if err != nil {
		config.Log.Error(err)
		resComm(c, models.ErrQueryData, nil)
		return
	}
	gms := []models.JobGroupMachine{}
	if err := c.BindJSON(&gms); err != nil {
		resComm(c, models.ErrPostData, nil)
		return
	}
	bf := models.PutJobGroupMachines(jID, &gms)
	resComm(c, bf, nil)
}

func getGroupLabels(c *gin.Context) {
	sp := &models.SplitPage{}
	if err := c.BindQuery(sp); err != nil {
		resComm(c, models.ErrSplitParma, nil)
		return
	}
	// idStr, ok := c.GetQuery("id")
	// if !ok {
	// 	resComm(c, models.ErrQueryData, nil)
	// 	return
	// }
	// jID, err := strconv.ParseInt(idStr, 10, 0)
	// if err != nil {
	// 	config.Log.Error(err)
	// 	resComm(c, models.ErrQueryData, nil)
	// 	return
	// }
	// data, bf := models.GetGroupLabels(jID)
	data, bf := models.GetJobGroupWithSplitPage(sp)
	resComm(c, bf, data)
}

func putGroupLabels(c *gin.Context) {
	idStr, ok := c.GetQuery("id")
	if !ok {
		resComm(c, models.ErrQueryData, nil)
		return
	}
	jID, err := strconv.ParseInt(idStr, 10, 0)
	if err != nil {
		config.Log.Error(err)
		resComm(c, models.ErrQueryData, nil)
		return
	}
	gls := models.GroupLabels{}
	if err := c.BindJSON(&gls); err != nil {
		resComm(c, models.ErrPostData, nil)
		return
	}
	bf := models.PutGroupLabels(jID, &gls)
	resComm(c, bf, nil)
}

func delGroupLabels(c *gin.Context) {
	dinfo := models.DelGroupLables{}
	if err := c.BindQuery(&dinfo); err != nil {
		resComm(c, models.ErrSplitParma, nil)
		return
	}
	bf := models.DelGroupLabels(&dinfo)
	resComm(c, bf, nil)
}

func getAllMachinesLabels(c *gin.Context) {
	info := models.JobGroupID{}
	if err := c.BindQuery(&info); err != nil {
		resComm(c, models.ErrSplitParma, nil)
		return
	}
	data, bf := models.GetAllMachinesLabels(&info)
	resComm(c, bf, data)
}

func putJobGroupStatus(c *gin.Context) {
	ed := models.EnabledInfo{}
	if err := c.BindJSON(&ed); err != nil {
		resComm(c, models.ErrPostData, nil)
		return
	}
	bf := models.PutJobGroupStatus(&ed)
	resComm(c, bf, nil)
}

func putJobGroupLabelsStatus(c *gin.Context) {
	ed := models.EnabledInfo{}
	if err := c.BindJSON(&ed); err != nil {
		resComm(c, models.ErrPostData, nil)
		return
	}
	bf := models.PutJobGroupLabelsStatus(&ed)
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
	// bf := models.DoPublishJobs()
	bf := models.AllowOneObj.DoPublishJobs()
	resComm(c, bf, nil)
}

func putJobsStatus(c *gin.Context) {
	ed := models.EnabledInfo{}
	if err := c.BindJSON(&ed); err != nil {
		resComm(c, models.ErrPostData, nil)
		return
	}
	bf := models.PutJobsStatus(&ed)
	resComm(c, bf, nil)
}

func postUpdateJobIPs(c *gin.Context) {
	cInfo := models.UpdateIPForJob{}
	if err := c.BindJSON(&cInfo); err != nil {
		resComm(c, models.ErrPostData, nil)
		return
	}
	bf := models.PostUpdateJobIPs(&cInfo)
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

func putMachineStatus(c *gin.Context) {
	ed := models.EnabledInfo{}
	if err := c.BindJSON(&ed); err != nil {
		resComm(c, models.ErrPostData, nil)
		return
	}
	bf := models.PutMachineStatus(&ed)
	resComm(c, bf, nil)
}

func getAllMachines(c *gin.Context) {
	data, bf := models.GetAllMachine()
	resComm(c, bf, data)
}

func batchDeleteMachine(c *gin.Context) {
	ids := []int{}
	if err := c.BindJSON(&ids); err != nil {
		resComm(c, models.ErrSplitParma, nil)
		return
	}
	bf := models.BatchDeleteMachine(ids)
	resComm(c, bf, nil)
}

func uploadMachines(c *gin.Context) {
	info := models.UploadMachinesInfo{}
	if err := c.BindJSON(&info); err != nil {
		resComm(c, models.ErrSplitParma, nil)
		return
	}
	result, bf := models.UploadMachines(&info)
	resComm(c, bf, result)
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

func getDefEnableLabels(c *gin.Context) {
	data, bf := models.GetDefEnableLabels()
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
	var skipSelf bool
	if strings.ToLower(c.Query("skip_self")) == "true" {
		skipSelf = true
	}
	tnc := models.TreeNodeFromCli{}
	if err := c.BindJSON(&tnc); err != nil {
		resComm(c, models.ErrPostData, nil)
		return
	}
	bf := models.DeleteTreeNode(&tnc, skipSelf)
	resComm(c, bf, nil)
}

func rulePublish(c *gin.Context) {
	bf := models.RulePublish()
	resComm(c, bf, nil)
}

func putTreeNodeStatus(c *gin.Context) {
	upInfo := models.TreeNodeStatus{}
	if err := c.BindJSON(&upInfo); err != nil {
		resComm(c, models.ErrPostData, nil)
		return
	}
	bf := models.PutTreeNodeStatus(&upInfo)
	resComm(c, bf, nil)
}

func postTreeUploadFileYaml(c *gin.Context) {
	gidStr, ok := c.GetQuery("gid")
	if !ok {
		resComm(c, models.ErrQueryData, nil)
		return
	}
	gid, err := strconv.ParseInt(gidStr, 10, 0)
	if err != nil {
		config.Log.Error(err)
		resComm(c, models.ErrQueryData, nil)
		return
	}
	r, bf := models.PostTreeUploadFileYaml(c, gid)
	resComm(c, bf, r)
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

func putBaseLabelsStatus(c *gin.Context) {
	ed := models.EnabledInfo{}
	if err := c.BindJSON(&ed); err != nil {
		resComm(c, models.ErrPostData, nil)
		return
	}
	bf := models.PutBaseLabelsStatus(&ed)
	resComm(c, bf, nil)
}

func putBaseRelabelsStatus(c *gin.Context) {
	ed := models.EnabledInfo{}
	if err := c.BindJSON(&ed); err != nil {
		resComm(c, models.ErrPostData, nil)
		return
	}
	bf := models.PutBaseRelabelsStatus(&ed)
	resComm(c, bf, nil)
}

func getProTmpl(c *gin.Context) {
	tmpl, bf := models.GetProTmpl()
	resComm(c, bf, tmpl)
}

func putProTmpl(c *gin.Context) {
	tp := models.ProTmpl{}
	if err := c.BindJSON(&tp); err != nil {
		resComm(c, models.ErrPostData, nil)
		return
	}
	bf := models.PutProTmpl(&tp)
	resComm(c, bf, nil)
}

func getStruct(c *gin.Context) {
	data, bf := models.GetStruct()
	resComm(c, bf, data)
}

func ws(c *gin.Context) {
	bf := models.WS(c)
	resComm(c, bf, nil)
}

func getOptions(c *gin.Context) {
	data, bf := models.GetOptions()
	resComm(c, bf, data)
}

func putOptions(c *gin.Context) {
	opts := map[string]string{}
	if err := c.BindJSON(&opts); err != nil {
		resComm(c, models.ErrPostData, nil)
		return
	}
	bf := models.PutOptions(opts)
	resComm(c, bf, nil)
}
