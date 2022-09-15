package router

import (
	"io"
	"pro_cfg_manager/alertmgr"
	"pro_cfg_manager/config"
	"pro_cfg_manager/models"
	"pro_cfg_manager/sync"
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
	v1.DELETE("/job/del/items", deleteJobItems)
	v1.DELETE("/job/del/black", deleteJobBlack)
	v1.PUT("/job/swap", swapJob)
	v1.POST("/jobs/publish", publishJobs)
	v1.PUT("/jobs/status", putJobsStatus)
	v1.POST("/jobs/update-ips", postUpdateJobIPs)
	v1.POST("/jobs/update-ips/black", postUpdateJobIPsBlack)
	v1.POST("/jobs/update-ips/v2", postUpdateJobIPsV2)
	v1.PUT("/job/update/subgroup", updateJobSubGroup)
	v1.DELETE("/job/selection", batchDeleteJob)
	v1.PUT("/job/selection/enable", batchEnableJob)
	v1.PUT("/job/selection/disable", batchDisableJob)

	v1.GET("/jobs/default/split", getDefJobsSplit)
	v1.POST("/job/default", postDefJob)
	v1.PUT("/job/default", putDefJob)
	v1.DELETE("/job/default", deleteDefJob)
	v1.PUT("/job/default/status", putJobDefaultStatus)
	v1.POST("/job/default/publish", publishDefJobs)

	v1.GET("/job/group", getJobGroup)
	v1.POST("/job/group", postJobGroup)
	v1.PUT("/job/group", putJobGroup)
	v1.DELETE("/job/group", delJobGroup)
	v1.DELETE("/job/group/subgroup/emtpy", delJobEmptySubGroup)
	v1.GET("/job/machines", getJobMachines)
	v1.GET("/job/machines/black", getJobMachinesBlack)
	v1.PUT("/job/machines/black", putJobMachinesBlack)
	v1.GET("/job/group/machines", getJobGroupMachines)
	v1.PUT("/job/group/machines", putJobGroupMachines)
	v1.GET("/job/group/labels", getGroupLabels)
	v1.POST("/job/group/labels", postGroupLabels)
	v1.PUT("/job/group/labels", putGroupLabels)
	v1.DELETE("/job/group/labels", delGroupLabels)
	v1.GET("/job/group/machines-and-labels", getAllMachinesLabels)
	v1.PUT("/job/group/status", putJobGroupStatus)
	v1.PUT("/job/group/labels/status", putJobGroupLabelsStatus)
	v1.GET("/job/relabels/all", getJobAllRelabels)
	v1.GET("/job/global/label", getJobGlobalLable)
	v1.PUT("/job/global/label", putJobGlobalLable)
	v1.GET("/job/global/mirror", getJobGlobalMirror)
	v1.PUT("/job/global/mirror", putJobGlobalMirror)

	v1.GET("/machine", getMachine)
	v1.GET("/machines", getMachines)
	v1.POST("/machine", postMachine)
	v1.PUT("/machine", putMachine)
	v1.DELETE("/machine", deleteMachine)
	v1.PUT("/machine/status", putMachineStatus)
	v1.PUT("/machines/position", updatePosition)
	v1.GET("/machines/all", getAllMachines)
	v1.DELETE("/machines/selection", batchDeleteMachine)
	v1.PUT("/machines/selection/enable", batchEnableMachine)
	v1.PUT("/machines/selection/disable", batchDisableMachine)
	v1.POST("/upload/machines", uploadMachines)
	v1.PUT("/machines/batch/import", batchImportIPAddrs)
	v1.PUT("/machines/batch/import/domain", batchImportDomain)

	v1.POST("/publish", publish)
	v1.POST("/reload", reload)
	v1.POST("/def/reload", defReload)
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
	v1.POST("/rules/empty/publish", emptyRulePublish)
	v1.PUT("/tree/node/status", putTreeNodeStatus)
	v1.POST("/tree/upload/file/yaml", postTreeUploadFileYaml)

	v1.GET("/base/labels", getBaseLabels)
	v1.POST("/base/labels", postBaseLabels)
	v1.PUT("/base/labels", putBaseLabels)
	v1.PUT("/base/labels/status", putBaseLabelsStatus)
	v1.DELETE("/base/labels", delBaseLabels)

	v1.GET("/base/relabels", getReLabels)
	v1.GET("/base/relabel", getRelabel)
	v1.GET("/base/relabels/all", getAllRelabels)
	v1.POST("/base/relabels", postReLabels)
	v1.PUT("/base/relabels", putReLabels)
	v1.DELETE("/base/relabels", delReLabels)
	v1.PUT("/base/relabels/code", putRelabelsCode)
	v1.GET("/base/relabels/check/view-code-priv", getRelabelsViewCodePriv)
	v1.PUT("/base/relabels/status", putBaseRelabelsStatus)

	v1.GET("/base/fields", getBaseFields)
	v1.POST("/base/fields", postBaseFields)
	v1.PUT("/base/fields", putBaseFields)
	v1.PUT("/base/fields/status", putBaseFieldsStatus)
	v1.DELETE("/base/fields", delBaseFields)

	v1.GET("/cron/rules/split", getCronRules)
	v1.POST("/cron/rule", postRule)
	v1.PUT("/cron/rule", putRule)
	v1.DELETE("/cron/rule", deleteRule)
	v1.PUT("/cron/rule/status", putRuleStatus)
	v1.DELETE("/cron/rules/selection", delRulesSelection)
	v1.PUT("/cron/rules/enable", putRulesEnable)
	v1.PUT("/cron/rules/disable", putRulesDisable)
	v1.PUT("/cron/rules/publish", putRulesPublish)
	v1.GET("/cron/rule/image", getRuleImage)
	v1.POST("/cron/send/testmail", postSendTestMail)

	v1.GET("/base/cron", getCron)
	v1.GET("/base/cron/all", getAllCron)
	v1.POST("/base/cron", postCron)
	v1.PUT("/base/cron", putCron)
	v1.PUT("/base/cron/status", putCronStatus)
	v1.DELETE("/base/cron", delCron)

	v1.GET("/prometheus/tmpl", getProTmpl)
	v1.PUT("/prometheus/tmpl", putProTmpl)
	v1.GET("/prometheus/struct", getStruct)

	v1.GET("/ws", ws)

	v1.GET("/options", getOptions)
	v1.PUT("/options", putOptions)

	v1.GET("/operate/log", getOperateLog)
	v1.POST("/operate/reset/secret", preOptResetSystem)
	v1.POST("/operate/reset/system", optResetSystem)

	v1.POST("/control/create", ctlCreate)
	v1.POST("/control/reload", ctlReload)
	v1.POST("/control/create/and/reload", ctlCreateAReload)
	v1.GET("/control/prometheus/url", getPrometheusUrl)

	v1.GET("/idc", getIDC)
	v1.GET("/idcs", getIDCs)
	v1.POST("/idc", postIDC)
	v1.PUT("/idc", putIDC)
	v1.PUT("/idc/remark", putIDCRemark)
	v1.DELETE("/idc", delIDC)

	v1.GET("/idc/tree", getIDCTree)
	v1.GET("/idc/tree/xls", getIDCTreeXls)

	v1.GET("/idc/line", getLine)
	v1.GET("/idc/lines", getLines)
	v1.POST("/idc/line", postLine)
	v1.PUT("/idc/line", putLine)
	v1.DELETE("/idc/line", delLine)

	v1.GET("/idc/line/ipaddrs", getLineIpAddrs)
	v1.PUT("/idc/line/ipaddrs", putLineIpAddrs)
	v1.PUT("/idc/line/expand/switch", idcUpdateLineExpandSwitch)
	v1.PUT("/idc/line/view/switch", idcUpdateLineViewSwitch)

	v1.PUT("/idc/update/netinfo/all", updateNetInfoAll)
	v1.PUT("/idc/update/netinfo/part", updateNetInfoPart)
	v1.PUT("/idc/create/label", createLabelForAllIPs)
	v1.POST("/idc/expand", idcExpand)
	v1.PUT("/idc/expand/switch", idcUpdateExpandSwitch)
	v1.PUT("/idc/view/switch", idcUpdateViewSwitch)
	v1.PUT("/idc/nobell", putIDCNoBell)
	v1.PUT("/idc/line/nobell", putLineNoBell)
}

func getTest(c *gin.Context) {
	result, bf := models.GetTest()
	resComm(c, bf, result)
}

func getJobs(c *gin.Context) {
	user := c.Keys["userInfo"].(*models.UserSessionInfo)
	pass := models.CheckPriv(user, "jobs", "jobs", "job_search")
	if pass != models.Success {
		resComm(c, pass, nil)
		return
	}
	result, bf := models.GetJobs()
	models.OO.RecodeLog(user.Username, c.Request.RemoteAddr, "get jobs", models.IsSearch, bf)
	resComm(c, bf, result)
}

func getJobsSplit(c *gin.Context) {
	user := c.Keys["userInfo"].(*models.UserSessionInfo)
	pass := models.CheckPriv(user, "jobs", "jobs", "job_search")
	if pass != models.Success {
		resComm(c, pass, nil)
		return
	}
	sp := &models.SplitPage{}
	if err := c.BindQuery(sp); err != nil {
		models.OO.RecodeLog(user.Username, c.Request.RemoteAddr, "get jobs", models.IsSearch, models.ErrSplitParma)
		resComm(c, models.ErrSplitParma, nil)
		return
	}
	result, bf := models.GetJobsSplit(sp)
	models.OO.RecodeLog(user.Username, c.Request.RemoteAddr, "get jobs", models.IsSearch, bf)
	resComm(c, bf, result)
}

func getDefJobsSplit(c *gin.Context) {
	user := c.Keys["userInfo"].(*models.UserSessionInfo)
	pass := models.CheckPriv(user, "baseConfig", "defaultJobs", "search")
	if pass != models.Success {
		resComm(c, pass, nil)
		return
	}
	sp := &models.SplitPage{}
	if err := c.BindQuery(sp); err != nil {
		models.OO.RecodeLog(user.Username, c.Request.RemoteAddr, "get default jobs", models.IsSearch, models.ErrSplitParma)
		resComm(c, models.ErrSplitParma, nil)
		return
	}
	result, bf := models.GetDefJobsSplit(sp)
	models.OO.RecodeLog(user.Username, c.Request.RemoteAddr, "get default jobs", models.IsSearch, bf)
	resComm(c, bf, result)
}

func getJob(c *gin.Context) {
	user := c.Keys["userInfo"].(*models.UserSessionInfo)
	pass := models.CheckPriv(user, "jobs", "jobs", "job_search")
	if pass != models.Success {
		resComm(c, pass, nil)
		return
	}
	jIDstr, ok := c.GetQuery("id")
	if !ok {
		models.OO.RecodeLog(user.Username, c.Request.RemoteAddr, "get job", models.IsSearch, models.ErrQueryData)
		resComm(c, models.ErrQueryData, nil)
		return
	}
	jID, err := strconv.ParseInt(jIDstr, 10, 0)
	if err != nil {
		config.Log.Error(err)
		models.OO.RecodeLog(user.Username, c.Request.RemoteAddr, "get job", models.IsSearch, models.ErrQueryData)
		resComm(c, models.ErrQueryData, nil)
		return
	}
	result, bf := models.GetJob(jID)
	models.OO.RecodeLog(user.Username, c.Request.RemoteAddr, "get job", models.IsSearch, bf)
	resComm(c, bf, result)
}

func postJob(c *gin.Context) {
	user := c.Keys["userInfo"].(*models.UserSessionInfo)
	pass := models.CheckPriv(user, "jobs", "jobs", "job_add")
	if pass != models.Success {
		resComm(c, pass, nil)
		return
	}
	jInfo := &models.Jobs{}
	if err := c.BindJSON(jInfo); err != nil {
		models.OO.RecodeLog(user.Username, c.Request.RemoteAddr, "create job", models.IsAdd, models.ErrPostData)
		resComm(c, models.ErrPostData, nil)
		return
	}
	jInfo.IsCommon = false
	bf := models.PostJob(user, jInfo)
	models.OO.RecodeLog(user.Username, c.Request.RemoteAddr, "create job", models.IsAdd, bf)
	resComm(c, bf, nil)
}

func postDefJob(c *gin.Context) {
	user := c.Keys["userInfo"].(*models.UserSessionInfo)
	pass := models.CheckPriv(user, "baseConfig", "defaultJobs", "add")
	if pass != models.Success {
		resComm(c, pass, nil)
		return
	}
	jInfo := &models.Jobs{}
	if err := c.BindJSON(jInfo); err != nil {
		models.OO.RecodeLog(user.Username, c.Request.RemoteAddr, "create default job", models.IsAdd, models.ErrPostData)
		resComm(c, models.ErrPostData, nil)
		return
	}
	jInfo.IsCommon = true
	bf := models.PostJob(user, jInfo)
	models.OO.RecodeLog(user.Username, c.Request.RemoteAddr, "create default job", models.IsAdd, bf)
	resComm(c, bf, nil)
}

func putJob(c *gin.Context) {
	user := c.Keys["userInfo"].(*models.UserSessionInfo)
	pass := models.CheckPriv(user, "jobs", "jobs", "job_update")
	if pass != models.Success {
		resComm(c, pass, nil)
		return
	}
	jInfo := &models.Jobs{}
	if err := c.BindJSON(jInfo); err != nil {
		models.OO.RecodeLog(user.Username, c.Request.RemoteAddr, "update job", models.IsUpdate, models.ErrPostData)
		resComm(c, models.ErrPostData, nil)
		return
	}
	jInfo.IsCommon = false
	bf := models.PutJob(user, jInfo)
	models.OO.RecodeLog(user.Username, c.Request.RemoteAddr, "update job", models.IsUpdate, bf)
	resComm(c, bf, nil)
}

func putDefJob(c *gin.Context) {
	user := c.Keys["userInfo"].(*models.UserSessionInfo)
	pass := models.CheckPriv(user, "baseConfig", "defaultJobs", "update")
	if pass != models.Success {
		resComm(c, pass, nil)
		return
	}
	jInfo := &models.Jobs{}
	if err := c.BindJSON(jInfo); err != nil {
		models.OO.RecodeLog(user.Username, c.Request.RemoteAddr, "update default job", models.IsUpdate, models.ErrPostData)
		resComm(c, models.ErrPostData, nil)
		return
	}
	jInfo.IsCommon = true
	bf := models.PutJob(user, jInfo)
	models.OO.RecodeLog(user.Username, c.Request.RemoteAddr, "update default job", models.IsUpdate, bf)
	resComm(c, bf, nil)
}

func deleteJob(c *gin.Context) {
	user := c.Keys["userInfo"].(*models.UserSessionInfo)
	pass := models.CheckPriv(user, "jobs", "jobs", "delete")
	if pass != models.Success {
		resComm(c, pass, nil)
		return
	}
	jIDstr, ok := c.GetQuery("id")
	if !ok {
		models.OO.RecodeLog(user.Username, c.Request.RemoteAddr, "delete job", models.IsDel, models.ErrQueryData)
		resComm(c, models.ErrQueryData, nil)
		return
	}
	jID, err := strconv.ParseInt(jIDstr, 10, 0)
	if err != nil {
		config.Log.Error(err)
		models.OO.RecodeLog(user.Username, c.Request.RemoteAddr, "delete job", models.IsDel, models.ErrQueryData)
		resComm(c, models.ErrQueryData, nil)
		return
	}
	bf := models.DeleteJob(jID, false)
	models.OO.RecodeLog(user.Username, c.Request.RemoteAddr, "delete job", models.IsDel, bf)
	resComm(c, bf, nil)
}

func deleteJobItems(c *gin.Context) {
	user := c.Keys["userInfo"].(*models.UserSessionInfo)
	pass := models.CheckPriv(user, "jobs", "jobs", "batch_del_job_items")
	if pass != models.Success {
		resComm(c, pass, nil)
		return
	}
	req := models.DelJobItemReq{}
	if err := c.BindJSON(&req); err != nil {
		config.Log.Error(err)
		models.OO.RecodeLog(user.Username, c.Request.RemoteAddr, "delete job items", models.IsDel, models.ErrPostData)
		resComm(c, models.ErrPostData, nil)
		return
	}
	// 把IP替换成对应的ID编号
	cInfo, bf := models.ChangeIPAddrUseID(user, req.Content, false)
	if bf != models.Success {
		models.OO.RecodeLog(user.Username, c.Request.RemoteAddr, "delete job items", models.IsUpdate, models.ErrPostData)
		resComm(c, bf, nil)
		return
	}
	bf = models.DeleteJobItems(user, cInfo, &req)
	models.OO.RecodeLog(user.Username, c.Request.RemoteAddr, "delete job items", models.IsDel, bf)
	resComm(c, bf, nil)
}

func deleteJobBlack(c *gin.Context) {
	user := c.Keys["userInfo"].(*models.UserSessionInfo)
	pass := models.CheckPriv(user, "jobs", "jobs", "batch_del_job_black")
	if pass != models.Success {
		resComm(c, pass, nil)
		return
	}
	req := models.DelJobBlackReq{}
	if err := c.BindJSON(&req); err != nil {
		config.Log.Error(err)
		models.OO.RecodeLog(user.Username, c.Request.RemoteAddr, "delete job black", models.IsDel, models.ErrPostData)
		resComm(c, models.ErrPostData, nil)
		return
	}
	bf := models.DeleteJobBlack(user, &req)
	models.OO.RecodeLog(user.Username, c.Request.RemoteAddr, "delete job black", models.IsDel, bf)
	resComm(c, bf, nil)
}

func deleteDefJob(c *gin.Context) {
	user := c.Keys["userInfo"].(*models.UserSessionInfo)
	pass := models.CheckPriv(user, "baseConfig", "defaultJobs", "delete")
	if pass != models.Success {
		resComm(c, pass, nil)
		return
	}
	jIDstr, ok := c.GetQuery("id")
	if !ok {
		models.OO.RecodeLog(user.Username, c.Request.RemoteAddr, "delete default job", models.IsDel, models.ErrQueryData)
		resComm(c, models.ErrQueryData, nil)
		return
	}
	jID, err := strconv.ParseInt(jIDstr, 10, 0)
	if err != nil {
		config.Log.Error(err)
		models.OO.RecodeLog(user.Username, c.Request.RemoteAddr, "delete default job", models.IsDel, models.ErrQueryData)
		resComm(c, models.ErrQueryData, nil)
		return
	}
	bf := models.DeleteJob(jID, true)
	models.OO.RecodeLog(user.Username, c.Request.RemoteAddr, "delete default job", models.IsDel, bf)
	resComm(c, bf, nil)
}

func putJobDefaultStatus(c *gin.Context) {
	user := c.Keys["userInfo"].(*models.UserSessionInfo)
	pass := models.CheckPriv(user, "baseConfig", "defaultJobs", "dis.enable")
	if pass != models.Success {
		resComm(c, pass, nil)
		return
	}
	ed := models.EnabledInfo{}
	if err := c.BindJSON(&ed); err != nil {
		models.OO.RecodeLog(user.Username, c.Request.RemoteAddr, "update default job status", models.IsUpdate, models.ErrPostData)
		resComm(c, models.ErrPostData, nil)
		return
	}
	bf := models.PutJobDefaultStatus(user, &ed)
	models.OO.RecodeLog(user.Username, c.Request.RemoteAddr, "update default job status", models.IsUpdate, bf)
	resComm(c, bf, nil)
}

func getJobGroup(c *gin.Context) {
	user := c.Keys["userInfo"].(*models.UserSessionInfo)
	pass := models.CheckPriv(user, "jobs", "labelsJobs", "search_sub_group")
	if pass != models.Success {
		resComm(c, pass, nil)
		return
	}
	sp := &models.SplitPage{}
	if err := c.BindQuery(sp); err != nil {
		models.OO.RecodeLog(user.Username, c.Request.RemoteAddr, "get sub group", models.IsSearch, models.ErrSplitParma)
		resComm(c, models.ErrSplitParma, nil)
		return
	}
	data, bf := models.GetJobGroup(sp)
	models.OO.RecodeLog(user.Username, c.Request.RemoteAddr, "get sub group", models.IsSearch, bf)
	resComm(c, bf, data)
}

func postJobGroup(c *gin.Context) {
	user := c.Keys["userInfo"].(*models.UserSessionInfo)
	pass := models.CheckPriv(user, "jobs", "labelsJobs", "add_sub_group")
	if pass != models.Success {
		resComm(c, pass, nil)
		return
	}
	info := &models.JobGroup{}
	if err := c.BindJSON(info); err != nil {
		models.OO.RecodeLog(user.Username, c.Request.RemoteAddr, "create sub group", models.IsAdd, models.ErrPostData)
		resComm(c, models.ErrPostData, nil)
		return
	}
	bf := models.PostJobGroup(user, info)
	models.OO.RecodeLog(user.Username, c.Request.RemoteAddr, "create sub group", models.IsAdd, bf)
	resComm(c, bf, nil)
}

func putJobGroup(c *gin.Context) {
	user := c.Keys["userInfo"].(*models.UserSessionInfo)
	pass := models.CheckPriv(user, "jobs", "labelsJobs", "update_sub_group")
	if pass != models.Success {
		resComm(c, pass, nil)
		return
	}
	info := &models.JobGroup{}
	if err := c.BindJSON(info); err != nil {
		models.OO.RecodeLog(user.Username, c.Request.RemoteAddr, "update sub group", models.IsUpdate, models.ErrPostData)
		resComm(c, models.ErrPostData, nil)
		return
	}
	bf := models.PutJobGroup(user, info)
	models.OO.RecodeLog(user.Username, c.Request.RemoteAddr, "update sub group", models.IsUpdate, bf)
	resComm(c, bf, nil)
}

func delJobGroup(c *gin.Context) {
	user := c.Keys["userInfo"].(*models.UserSessionInfo)
	pass := models.CheckPriv(user, "jobs", "labelsJobs", "delete_sub_group")
	if pass != models.Success {
		resComm(c, pass, nil)
		return
	}
	info := &models.DelJobGroupInfo{}
	if err := c.BindQuery(info); err != nil {
		models.OO.RecodeLog(user.Username, c.Request.RemoteAddr, "del sub group", models.IsDel, models.ErrPostData)
		resComm(c, models.ErrPostData, nil)
		return
	}
	bf := models.DelJobGroup(info)
	models.OO.RecodeLog(user.Username, c.Request.RemoteAddr, "del sub group", models.IsDel, bf)
	resComm(c, bf, nil)
}

func delJobEmptySubGroup(c *gin.Context) {
	user := c.Keys["userInfo"].(*models.UserSessionInfo)
	pass := models.CheckPriv(user, "jobs", "labelsJobs", "delete_empty_sub_group")
	if pass != models.Success {
		resComm(c, pass, nil)
		return
	}
	info := &models.OnlyID{}
	if err := c.BindQuery(info); err != nil {
		models.OO.RecodeLog(user.Username, c.Request.RemoteAddr, "del empty sub group", models.IsDel, models.ErrPostData)
		resComm(c, models.ErrPostData, nil)
		return
	}
	bf := models.DelJobEmptySubGroup(info)
	models.OO.RecodeLog(user.Username, c.Request.RemoteAddr, "del empty sub group", models.IsDel, bf)
	resComm(c, bf, nil)
}

func getJobMachines(c *gin.Context) {
	user := c.Keys["userInfo"].(*models.UserSessionInfo)
	pass := models.CheckPriv(user, "jobs", "jobs", "job_search")
	if pass != models.Success {
		resComm(c, pass, nil)
		return
	}
	idStr, ok := c.GetQuery("id")
	if !ok {
		models.OO.RecodeLog(user.Username, c.Request.RemoteAddr, "get job machines", models.IsSearch, models.ErrQueryData)
		resComm(c, models.ErrQueryData, nil)
		return
	}
	jID, err := strconv.ParseInt(idStr, 10, 0)
	if err != nil {
		config.Log.Error(err)
		models.OO.RecodeLog(user.Username, c.Request.RemoteAddr, "get job machines", models.IsSearch, models.ErrQueryData)
		resComm(c, models.ErrQueryData, nil)
		return
	}
	data, bf := models.GetJobMachines(jID)
	models.OO.RecodeLog(user.Username, c.Request.RemoteAddr, "get job machines", models.IsSearch, bf)
	resComm(c, bf, data)
}

func getJobMachinesBlack(c *gin.Context) {
	user := c.Keys["userInfo"].(*models.UserSessionInfo)
	pass := models.CheckPriv(user, "jobs", "jobs", "job_search_black")
	if pass != models.Success {
		resComm(c, pass, nil)
		return
	}
	idStr, ok := c.GetQuery("id")
	if !ok {
		models.OO.RecodeLog(user.Username, c.Request.RemoteAddr, "get job black machines", models.IsSearch, models.ErrQueryData)
		resComm(c, models.ErrQueryData, nil)
		return
	}
	jID, err := strconv.ParseInt(idStr, 10, 0)
	if err != nil {
		config.Log.Error(err)
		models.OO.RecodeLog(user.Username, c.Request.RemoteAddr, "get job black machines", models.IsSearch, models.ErrQueryData)
		resComm(c, models.ErrQueryData, nil)
		return
	}
	data, bf := models.GetJobMachinesBlack(jID)
	models.OO.RecodeLog(user.Username, c.Request.RemoteAddr, "get job black machines", models.IsSearch, bf)
	resComm(c, bf, data)
}

func putJobMachinesBlack(c *gin.Context) {
	user := c.Keys["userInfo"].(*models.UserSessionInfo)
	pass := models.CheckPriv(user, "jobs", "jobs", "job_update_black")
	if pass != models.Success {
		resComm(c, pass, nil)
		return
	}
	jobID := models.OnlyID{}
	if err := c.BindQuery(&jobID); err != nil {
		config.Log.Error("job id error")
		models.OO.RecodeLog(user.Username, c.Request.RemoteAddr, "update job ip pool", models.IsUpdate, models.ErrPostData)
		resComm(c, models.ErrPostData, nil)
		return
	}
	// 把IP替换成对应的ID编号
	body, err := io.ReadAll(c.Request.Body)
	if err != nil {
		config.Log.Error(err)
		models.OO.RecodeLog(user.Username, c.Request.RemoteAddr, "update job ip pool", models.IsUpdate, models.ErrPostData)
		resComm(c, models.ErrPostData, nil)
		return
	}
	cInfo, bf := models.ChangeIPAddrUseID(user, string(body), false)
	if bf != models.Success {
		models.OO.RecodeLog(user.Username, c.Request.RemoteAddr, "update job ip pool", models.IsUpdate, models.ErrPostData)
		resComm(c, bf, nil)
		return
	}
	cInfo.JobID = jobID.ID
	bf = models.PostUpdateJobIPsBlack(user, cInfo)
	models.OO.RecodeLog(user.Username, c.Request.RemoteAddr, "update job ip pool", models.IsUpdate, bf)
	resComm(c, bf, nil)
}

func getJobGroupMachines(c *gin.Context) {
	user := c.Keys["userInfo"].(*models.UserSessionInfo)
	pass := models.CheckPriv(user, "jobs", "labelsJobs", "search_sub_group_ip_pool")
	if pass != models.Success {
		resComm(c, pass, nil)
		return
	}
	idStr, ok := c.GetQuery("id")
	if !ok {
		models.OO.RecodeLog(user.Username, c.Request.RemoteAddr, "get sub group machines", models.IsUpdate, models.ErrQueryData)
		resComm(c, models.ErrQueryData, nil)
		return
	}
	jID, err := strconv.ParseInt(idStr, 10, 0)
	if err != nil {
		config.Log.Error(err)
		models.OO.RecodeLog(user.Username, c.Request.RemoteAddr, "get sub group machines", models.IsUpdate, models.ErrQueryData)
		resComm(c, models.ErrQueryData, nil)
		return
	}
	data, bf := models.GetJobGroupMachines(jID)
	models.OO.RecodeLog(user.Username, c.Request.RemoteAddr, "get sub group machines", models.IsUpdate, bf)
	resComm(c, bf, data)
}

func putJobGroupMachines(c *gin.Context) {
	user := c.Keys["userInfo"].(*models.UserSessionInfo)
	pass := models.CheckPriv(user, "jobs", "labelsJobs", "update_sub_group_ip_pool")
	if pass != models.Success {
		resComm(c, pass, nil)
		return
	}
	idStr, ok := c.GetQuery("id")
	if !ok {
		models.OO.RecodeLog(user.Username, c.Request.RemoteAddr, "update sub group ip pool", models.IsUpdate, models.ErrQueryData)
		resComm(c, models.ErrQueryData, nil)
		return
	}
	jID, err := strconv.ParseInt(idStr, 10, 0)
	if err != nil {
		config.Log.Error(err)
		models.OO.RecodeLog(user.Username, c.Request.RemoteAddr, "update sub group ip pool", models.IsUpdate, models.ErrQueryData)
		resComm(c, models.ErrQueryData, nil)
		return
	}
	jobIDStr, ok := c.GetQuery("job_id")
	if !ok {
		models.OO.RecodeLog(user.Username, c.Request.RemoteAddr, "update sub group ip pool", models.IsUpdate, models.ErrQueryData)
		resComm(c, models.ErrQueryData, nil)
		return
	}
	jobID, err := strconv.ParseInt(jobIDStr, 10, 0)
	if err != nil {
		config.Log.Error(err)
		models.OO.RecodeLog(user.Username, c.Request.RemoteAddr, "update sub group ip pool", models.IsUpdate, models.ErrQueryData)
		resComm(c, models.ErrQueryData, nil)
		return
	}
	gms := []models.JobGroupMachine{}
	if err := c.BindJSON(&gms); err != nil {
		models.OO.RecodeLog(user.Username, c.Request.RemoteAddr, "update sub group ip pool", models.IsUpdate, models.ErrPostData)
		resComm(c, models.ErrPostData, nil)
		return
	}
	bf := models.PutJobGroupMachines(jobID, jID, &gms)
	models.OO.RecodeLog(user.Username, c.Request.RemoteAddr, "update sub group ip pool", models.IsUpdate, bf)
	resComm(c, bf, nil)
}

func getGroupLabels(c *gin.Context) {
	user := c.Keys["userInfo"].(*models.UserSessionInfo)
	pass := models.CheckPriv(user, "jobs", "labelsJobs", "search_sub_group_label")
	if pass != models.Success {
		resComm(c, pass, nil)
		return
	}
	sp := &models.SplitPage{}
	if err := c.BindQuery(sp); err != nil {
		models.OO.RecodeLog(user.Username, c.Request.RemoteAddr, "get sub group lables", models.IsSearch, models.ErrSplitParma)
		resComm(c, models.ErrSplitParma, nil)
		return
	}
	data, bf := models.GetJobGroupWithSplitPage(sp)
	models.OO.RecodeLog(user.Username, c.Request.RemoteAddr, "get sub group lables", models.IsSearch, bf)
	resComm(c, bf, data)
}

func postGroupLabels(c *gin.Context) {
	user := c.Keys["userInfo"].(*models.UserSessionInfo)
	pass := models.CheckPriv(user, "jobs", "labelsJobs", "add_sub_group_label")
	if pass != models.Success {
		resComm(c, pass, nil)
		return
	}
	idStr, ok := c.GetQuery("id")
	if !ok {
		models.OO.RecodeLog(user.Username, c.Request.RemoteAddr, "add sub group label", models.IsAdd, models.ErrQueryData)
		resComm(c, models.ErrQueryData, nil)
		return
	}
	jID, err := strconv.ParseInt(idStr, 10, 0)
	if err != nil {
		config.Log.Error(err)
		models.OO.RecodeLog(user.Username, c.Request.RemoteAddr, "add sub group label", models.IsAdd, models.ErrQueryData)
		resComm(c, models.ErrQueryData, nil)
		return
	}
	gls := models.GroupLabels{}
	if err := c.BindJSON(&gls); err != nil {
		models.OO.RecodeLog(user.Username, c.Request.RemoteAddr, "add sub group label", models.IsAdd, models.ErrPostData)
		resComm(c, models.ErrPostData, nil)
		return
	}
	bf := models.PostGroupLabels(jID, &gls)
	models.OO.RecodeLog(user.Username, c.Request.RemoteAddr, "add sub group label", models.IsAdd, bf)
	resComm(c, bf, nil)
}

func putGroupLabels(c *gin.Context) {
	user := c.Keys["userInfo"].(*models.UserSessionInfo)
	pass := models.CheckPriv(user, "jobs", "labelsJobs", "update_sub_group_label")
	if pass != models.Success {
		resComm(c, pass, nil)
		return
	}
	idStr, ok := c.GetQuery("id")
	if !ok {
		models.OO.RecodeLog(user.Username, c.Request.RemoteAddr, "update sub group label", models.IsUpdate, models.ErrQueryData)
		resComm(c, models.ErrQueryData, nil)
		return
	}
	jID, err := strconv.ParseInt(idStr, 10, 0)
	if err != nil {
		config.Log.Error(err)
		models.OO.RecodeLog(user.Username, c.Request.RemoteAddr, "update sub group label", models.IsUpdate, models.ErrQueryData)
		resComm(c, models.ErrQueryData, nil)
		return
	}
	gls := models.GroupLabels{}
	if err := c.BindJSON(&gls); err != nil {
		models.OO.RecodeLog(user.Username, c.Request.RemoteAddr, "update sub group label", models.IsUpdate, models.ErrPostData)
		resComm(c, models.ErrPostData, nil)
		return
	}
	bf := models.PutGroupLabels(jID, &gls)
	models.OO.RecodeLog(user.Username, c.Request.RemoteAddr, "update sub group label", models.IsUpdate, bf)
	resComm(c, bf, nil)
}

func delGroupLabels(c *gin.Context) {
	user := c.Keys["userInfo"].(*models.UserSessionInfo)
	pass := models.CheckPriv(user, "jobs", "labelsJobs", "delete_sub_group_label")
	if pass != models.Success {
		resComm(c, pass, nil)
		return
	}
	dinfo := models.DelGroupLables{}
	if err := c.BindQuery(&dinfo); err != nil {
		models.OO.RecodeLog(user.Username, c.Request.RemoteAddr, "delete sub group lables", models.IsDel, models.ErrSplitParma)
		resComm(c, models.ErrSplitParma, nil)
		return
	}
	bf := models.DelGroupLabels(&dinfo)
	models.OO.RecodeLog(user.Username, c.Request.RemoteAddr, "delete sub group lables", models.IsDel, bf)
	resComm(c, bf, nil)
}

func getAllMachinesLabels(c *gin.Context) {
	user := c.Keys["userInfo"].(*models.UserSessionInfo)
	pass := models.CheckPriv(user, "jobs", "jobs", "job_search")
	if pass != models.Success {
		resComm(c, pass, nil)
		return
	}
	info := models.JobGroupID{}
	if err := c.BindQuery(&info); err != nil {
		models.OO.RecodeLog(user.Username, c.Request.RemoteAddr, "get all methines and labels", models.IsSearch, models.ErrSplitParma)
		resComm(c, models.ErrSplitParma, nil)
		return
	}
	data, bf := models.GetAllMachinesLabels(&info)
	models.OO.RecodeLog(user.Username, c.Request.RemoteAddr, "get all methines and labels", models.IsSearch, bf)
	resComm(c, bf, data)
}

func putJobGroupStatus(c *gin.Context) {
	user := c.Keys["userInfo"].(*models.UserSessionInfo)
	pass := models.CheckPriv(user, "jobs", "labelsJobs", "dis.enable_sub_group")
	if pass != models.Success {
		resComm(c, pass, nil)
		return
	}
	ed := models.EnabledInfo{}
	if err := c.BindJSON(&ed); err != nil {
		models.OO.RecodeLog(user.Username, c.Request.RemoteAddr, "update sub group status", models.IsUpdate, models.ErrPostData)
		resComm(c, models.ErrPostData, nil)
		return
	}
	bf := models.PutJobGroupStatus(user, &ed)
	models.OO.RecodeLog(user.Username, c.Request.RemoteAddr, "update sub group status", models.IsUpdate, bf)
	resComm(c, bf, nil)
}

func putJobGroupLabelsStatus(c *gin.Context) {
	user := c.Keys["userInfo"].(*models.UserSessionInfo)
	pass := models.CheckPriv(user, "jobs", "labelsJobs", "dis.enable_sub_group_label")
	if pass != models.Success {
		resComm(c, pass, nil)
		return
	}
	ed := models.EnabledInfo{}
	if err := c.BindJSON(&ed); err != nil {
		models.OO.RecodeLog(user.Username, c.Request.RemoteAddr, "update sub job group status", models.IsUpdate, models.ErrPostData)
		resComm(c, models.ErrPostData, nil)
		return
	}
	bf := models.PutJobGroupLabelsStatus(user, &ed)
	models.OO.RecodeLog(user.Username, c.Request.RemoteAddr, "update sub job group status", models.IsUpdate, bf)
	resComm(c, bf, nil)
}

func getJobAllRelabels(c *gin.Context) {
	user := c.Keys["userInfo"].(*models.UserSessionInfo)
	pass := models.CheckPriv(user, "jobs", "jobs", "job_add")
	if pass != models.Success {
		resComm(c, pass, nil)
		return
	}
	data, bf := models.GetAllReLabels()
	models.OO.RecodeLog(user.Username, c.Request.RemoteAddr, "get all labels of job's sub group", models.IsSearch, bf)
	resComm(c, bf, data)
}

func getJobGlobalLable(c *gin.Context) {
	user := c.Keys["userInfo"].(*models.UserSessionInfo)
	pass := models.CheckPriv(user, "jobs", "jobs", "get_job_label")
	if pass != models.Success {
		resComm(c, pass, nil)
		return
	}
	id := models.OnlyID{}
	if err := c.BindQuery(&id); err != nil {
		config.Log.Error(err)
		models.OO.RecodeLog(user.Username, c.Request.RemoteAddr, "get job labels", models.IsSearch, models.ErrQueryData)
		resComm(c, models.ErrQueryData, nil)
	}
	data, bf := models.GetJobGlobalLable(&id)
	models.OO.RecodeLog(user.Username, c.Request.RemoteAddr, "get job labels", models.IsSearch, bf)
	resComm(c, bf, data)
}

func putJobGlobalLable(c *gin.Context) {
	user := c.Keys["userInfo"].(*models.UserSessionInfo)
	pass := models.CheckPriv(user, "jobs", "jobs", "update_job_label")
	if pass != models.Success {
		resComm(c, pass, nil)
		return
	}
	id := models.OnlyID{}
	if err := c.BindQuery(&id); err != nil {
		config.Log.Error(err)
		models.OO.RecodeLog(user.Username, c.Request.RemoteAddr, "update job labels", models.IsUpdate, models.ErrQueryData)
		resComm(c, models.ErrQueryData, nil)
	}
	updateData := []*models.JobLabelsTblReq{}
	if err := c.BindJSON(&updateData); err != nil {
		config.Log.Error(err)
		models.OO.RecodeLog(user.Username, c.Request.RemoteAddr, "update job labels", models.IsUpdate, models.ErrPostData)
		resComm(c, models.ErrQueryData, nil)
	}
	bf := models.PutJobGlobalLable(user, &id, updateData)
	models.OO.RecodeLog(user.Username, c.Request.RemoteAddr, "update job labels", models.IsSearch, bf)
	resComm(c, bf, nil)
}

func getJobGlobalMirror(c *gin.Context) {
	user := c.Keys["userInfo"].(*models.UserSessionInfo)
	pass := models.CheckPriv(user, "jobs", "jobs", "get_job_mirror")
	if pass != models.Success {
		resComm(c, pass, nil)
		return
	}
	id := models.OnlyID{}
	if err := c.BindQuery(&id); err != nil {
		config.Log.Error(err)
		models.OO.RecodeLog(user.Username, c.Request.RemoteAddr, "get job mirror", models.IsSearch, models.ErrQueryData)
		resComm(c, models.ErrQueryData, nil)
	}
	data, bf := models.GetJobGlobalMirror(&id)
	models.OO.RecodeLog(user.Username, c.Request.RemoteAddr, "get job mirror", models.IsSearch, bf)
	resComm(c, bf, data)
}

func putJobGlobalMirror(c *gin.Context) {
	user := c.Keys["userInfo"].(*models.UserSessionInfo)
	pass := models.CheckPriv(user, "jobs", "jobs", "update_job_mirror")
	if pass != models.Success {
		resComm(c, pass, nil)
		return
	}
	id := models.OnlyID{}
	if err := c.BindQuery(&id); err != nil {
		config.Log.Error(err)
		models.OO.RecodeLog(user.Username, c.Request.RemoteAddr, "update job mirror", models.IsUpdate, models.ErrQueryData)
		resComm(c, models.ErrQueryData, nil)
	}
	updateData := models.JobMirror{}
	if err := c.BindJSON(&updateData); err != nil {
		config.Log.Error(err)
		models.OO.RecodeLog(user.Username, c.Request.RemoteAddr, "update job mirror", models.IsUpdate, models.ErrPostData)
		resComm(c, models.ErrQueryData, nil)
	}
	bf := models.PutJobGlobalMirror(user, &id, &updateData)
	models.OO.RecodeLog(user.Username, c.Request.RemoteAddr, "update job mirror", models.IsSearch, bf)
	resComm(c, bf, nil)
}

func swapJob(c *gin.Context) {
	user := c.Keys["userInfo"].(*models.UserSessionInfo)
	pass := models.CheckPriv(user, "jobs", "jobs", "swap")
	if pass != models.Success {
		resComm(c, pass, nil)
		return
	}
	sInfo := &models.SwapInfo{}
	if err := c.BindJSON(sInfo); err != nil {
		models.OO.RecodeLog(user.Username, c.Request.RemoteAddr, "swap group order", models.IsUpdate, models.ErrPostData)
		resComm(c, models.ErrPostData, nil)
		return
	}
	bf := models.DoSwap(user, sInfo)
	models.OO.RecodeLog(user.Username, c.Request.RemoteAddr, "swap group order", models.IsUpdate, bf)
	resComm(c, bf, nil)
}

func publishJobs(c *gin.Context) {
	user := c.Keys["userInfo"].(*models.UserSessionInfo)
	pass := models.CheckPriv(user, "jobs", "jobs", "publish")
	if pass != models.Success {
		resComm(c, pass, nil)
		return
	}
	bf := models.AllowOneObj.DoPublishJobs(user, false)
	models.OO.RecodeLog(user.Username, c.Request.RemoteAddr, "publish job", models.IsPublish, bf)
	resComm(c, bf, nil)
}

func publishDefJobs(c *gin.Context) {
	user := c.Keys["userInfo"].(*models.UserSessionInfo)
	pass := models.CheckPriv(user, "baseConfig", "defaultJobs", "publish")
	if pass != models.Success {
		resComm(c, pass, nil)
		return
	}
	bf := models.AllowOneObj.DoPublishJobs(user, false)
	models.OO.RecodeLog(user.Username, c.Request.RemoteAddr, "publish default jobs", models.IsPublish, bf)
	resComm(c, bf, nil)
}

func putJobsStatus(c *gin.Context) {
	user := c.Keys["userInfo"].(*models.UserSessionInfo)
	pass := models.CheckPriv(user, "jobs", "jobs", "dis.enable")
	if pass != models.Success {
		resComm(c, pass, nil)
		return
	}
	ed := models.EnabledInfo{}
	if err := c.BindJSON(&ed); err != nil {
		models.OO.RecodeLog(user.Username, c.Request.RemoteAddr, "update job status", models.IsUpdate, models.ErrPostData)
		resComm(c, models.ErrPostData, nil)
		return
	}
	bf := models.PutJobsStatus(user, &ed)
	models.OO.RecodeLog(user.Username, c.Request.RemoteAddr, "update job status", models.IsUpdate, bf)
	resComm(c, bf, nil)
}

func postUpdateJobIPs(c *gin.Context) {
	user := c.Keys["userInfo"].(*models.UserSessionInfo)
	pass := models.CheckPriv(user, "jobs", "jobs", "ip_pool")
	if pass != models.Success {
		resComm(c, pass, nil)
		return
	}
	cInfo := models.UpdateIPForJob{}
	if err := c.BindJSON(&cInfo); err != nil {
		models.OO.RecodeLog(user.Username, c.Request.RemoteAddr, "update job ip pool", models.IsUpdate, models.ErrPostData)
		resComm(c, models.ErrPostData, nil)
		return
	}
	bf := models.PostUpdateJobIPs(user, &cInfo)
	models.OO.RecodeLog(user.Username, c.Request.RemoteAddr, "update job ip pool", models.IsUpdate, bf)
	resComm(c, bf, nil)
}

func postUpdateJobIPsBlack(c *gin.Context) {
	user := c.Keys["userInfo"].(*models.UserSessionInfo)
	pass := models.CheckPriv(user, "jobs", "jobs", "job_update_black")
	if pass != models.Success {
		resComm(c, pass, nil)
		return
	}
	cInfo := models.UpdateIPForJob{}
	if err := c.BindJSON(&cInfo); err != nil {
		models.OO.RecodeLog(user.Username, c.Request.RemoteAddr, "update job ip pool", models.IsUpdate, models.ErrPostData)
		resComm(c, models.ErrPostData, nil)
		return
	}
	bf := models.PostUpdateJobIPsBlack(user, &cInfo)
	models.OO.RecodeLog(user.Username, c.Request.RemoteAddr, "update job ip pool", models.IsUpdate, bf)
	resComm(c, bf, nil)
}

func postUpdateJobIPsV2(c *gin.Context) {
	user := c.Keys["userInfo"].(*models.UserSessionInfo)
	pass := models.CheckPriv(user, "jobs", "jobs", "ip_pool")
	if pass != models.Success {
		resComm(c, pass, nil)
		return
	}
	jobID := c.Query("id")
	if jobID == "" {
		config.Log.Error("job id empty")
		models.OO.RecodeLog(user.Username, c.Request.RemoteAddr, "update job ip pool", models.IsUpdate, models.ErrPostData)
		resComm(c, models.ErrPostData, nil)
		return
	}
	jobIDInt, err := strconv.ParseInt(jobID, 10, 0)
	if err != nil {
		config.Log.Error("job id need number")
		models.OO.RecodeLog(user.Username, c.Request.RemoteAddr, "update job ip pool", models.IsUpdate, models.ErrPostData)
		resComm(c, models.ErrPostData, nil)
		return
	}
	var force bool
	if strings.ToLower(c.Query("force")) == "true" {
		force = true
	} else {
		force = false
	}
	// 把IP替换成对应的ID编号
	body, err := io.ReadAll(c.Request.Body)
	if err != nil {
		config.Log.Error(err)
		models.OO.RecodeLog(user.Username, c.Request.RemoteAddr, "update job ip pool", models.IsUpdate, models.ErrPostData)
		resComm(c, models.ErrPostData, nil)
		return
	}
	cInfo, bf := models.ChangeIPAddrUseID(user, string(body), force)
	if bf != models.Success {
		models.OO.RecodeLog(user.Username, c.Request.RemoteAddr, "update job ip pool", models.IsUpdate, models.ErrPostData)
		resComm(c, bf, nil)
		return
	}
	cInfo.JobID = int(jobIDInt)
	bf = models.PostUpdateJobIPs(user, cInfo)
	models.OO.RecodeLog(user.Username, c.Request.RemoteAddr, "update job ip pool", models.IsUpdate, bf)
	resComm(c, bf, nil)
}

func updateJobSubGroup(c *gin.Context) {
	user := c.Keys["userInfo"].(*models.UserSessionInfo)
	pass := models.CheckPriv(user, "jobs", "jobs", "update_job_subgroup")
	if pass != models.Success {
		resComm(c, pass, nil)
		return
	}
	jobID := models.OnlyID{}
	if err := c.BindQuery(&jobID); err != nil {
		config.Log.Error(err)
		models.OO.RecodeLog(user.Username, c.Request.RemoteAddr, "update job subgroup", models.IsUpdate, models.ErrQueryData)
		resComm(c, models.ErrQueryData, nil)
		return
	}
	bf := models.UpdateJobSubGroup(user, &jobID)
	models.OO.RecodeLog(user.Username, c.Request.RemoteAddr, "update job subgroup", models.IsUpdate, bf)
	resComm(c, bf, nil)
}

func batchDeleteJob(c *gin.Context) {
	user := c.Keys["userInfo"].(*models.UserSessionInfo)
	pass := models.CheckPriv(user, "jobs", "jobs", "delete")
	if pass != models.Success {
		resComm(c, pass, nil)
		return
	}
	ids := []int{}
	if err := c.BindJSON(&ids); err != nil {
		models.OO.RecodeLog(user.Username, c.Request.RemoteAddr, "delete job batch", models.IsUpdate, models.ErrPostData)
		resComm(c, models.ErrSplitParma, nil)
		return
	}
	bf := models.BatchDeleteJob(ids)
	models.OO.RecodeLog(user.Username, c.Request.RemoteAddr, "delete job batch", models.IsUpdate, bf)
	resComm(c, bf, nil)
}

func batchEnableJob(c *gin.Context) {
	user := c.Keys["userInfo"].(*models.UserSessionInfo)
	pass := models.CheckPriv(user, "jobs", "jobs", "dis.enable")
	if pass != models.Success {
		resComm(c, pass, nil)
		return
	}
	ids := []int{}
	if err := c.BindJSON(&ids); err != nil {
		models.OO.RecodeLog(user.Username, c.Request.RemoteAddr, "enable job batch", models.IsUpdate, models.ErrPostData)
		resComm(c, models.ErrSplitParma, nil)
		return
	}
	bf := models.BatchEnableJob(user, ids)
	models.OO.RecodeLog(user.Username, c.Request.RemoteAddr, "enable job batch", models.IsUpdate, bf)
	resComm(c, bf, nil)
}

func batchDisableJob(c *gin.Context) {
	user := c.Keys["userInfo"].(*models.UserSessionInfo)
	pass := models.CheckPriv(user, "jobs", "jobs", "dis.enable")
	if pass != models.Success {
		resComm(c, pass, nil)
		return
	}
	ids := []int{}
	if err := c.BindJSON(&ids); err != nil {
		models.OO.RecodeLog(user.Username, c.Request.RemoteAddr, "disable job batch", models.IsUpdate, models.ErrPostData)
		resComm(c, models.ErrSplitParma, nil)
		return
	}
	bf := models.BatchDisableJob(user, ids)
	models.OO.RecodeLog(user.Username, c.Request.RemoteAddr, "disable job batch", models.IsUpdate, bf)
	resComm(c, bf, nil)
}

func getMachine(c *gin.Context) {
	user := c.Keys["userInfo"].(*models.UserSessionInfo)
	pass := models.CheckPriv(user, "ipManager", "", "search")
	if pass != models.Success {
		resComm(c, pass, nil)
		return
	}
	mIDstr, ok := c.GetQuery("id")
	if !ok {
		models.OO.RecodeLog(user.Username, c.Request.RemoteAddr, "get ip", models.IsSearch, models.ErrQueryData)
		resComm(c, models.ErrQueryData, nil)
		return
	}
	mID, err := strconv.ParseInt(mIDstr, 10, 0)
	if err != nil {
		config.Log.Error(err)
		models.OO.RecodeLog(user.Username, c.Request.RemoteAddr, "get ip", models.IsSearch, models.ErrQueryData)
		resComm(c, models.ErrQueryData, nil)
		return
	}
	result, bf := models.GetMachine(int(mID))
	models.OO.RecodeLog(user.Username, c.Request.RemoteAddr, "get ip", models.IsSearch, bf)
	resComm(c, bf, result)
}

func getMachines(c *gin.Context) {
	user := c.Keys["userInfo"].(*models.UserSessionInfo)
	pass := models.CheckPriv(user, "ipManager", "", "search")
	if pass != models.Success {
		resComm(c, pass, nil)
		return
	}
	sp := &models.SplitPage{}
	if err := c.BindQuery(sp); err != nil {
		models.OO.RecodeLog(user.Username, c.Request.RemoteAddr, "get all ip", models.IsSearch, models.ErrSplitParma)
		resComm(c, models.ErrSplitParma, nil)
		return
	}
	result, bf := models.GetMachinesV2(sp)
	models.OO.RecodeLog(user.Username, c.Request.RemoteAddr, "get all ip", models.IsSearch, bf)
	resComm(c, bf, result)
}

func postMachine(c *gin.Context) {
	user := c.Keys["userInfo"].(*models.UserSessionInfo)
	pass := models.CheckPriv(user, "ipManager", "", "add")
	if pass != models.Success {
		resComm(c, pass, nil)
		return
	}
	mInfo := &models.Machine{}
	if err := c.BindJSON(mInfo); err != nil {
		models.OO.RecodeLog(user.Username, c.Request.RemoteAddr, "create ip", models.IsAdd, models.ErrPostData)
		resComm(c, models.ErrPostData, nil)
		return
	}
	bf := models.PostMachine(user, mInfo)
	models.OO.RecodeLog(user.Username, c.Request.RemoteAddr, "create ip", models.IsAdd, bf)
	resComm(c, bf, nil)
}

func putMachine(c *gin.Context) {
	user := c.Keys["userInfo"].(*models.UserSessionInfo)
	pass := models.CheckPriv(user, "ipManager", "", "update")
	if pass != models.Success {
		resComm(c, pass, nil)
		return
	}
	mInfo := &models.Machine{}
	if err := c.BindJSON(mInfo); err != nil {
		models.OO.RecodeLog(user.Username, c.Request.RemoteAddr, "update ip", models.IsUpdate, models.ErrPostData)
		resComm(c, models.ErrPostData, nil)
		return
	}
	bf := models.PutMachine(user, mInfo)
	models.OO.RecodeLog(user.Username, c.Request.RemoteAddr, "update ip", models.IsUpdate, bf)
	resComm(c, bf, nil)
}

func deleteMachine(c *gin.Context) {
	user := c.Keys["userInfo"].(*models.UserSessionInfo)
	pass := models.CheckPriv(user, "ipManager", "", "delete")
	if pass != models.Success {
		resComm(c, pass, nil)
		return
	}
	mIDstr, ok := c.GetQuery("id")
	if !ok {
		models.OO.RecodeLog(user.Username, c.Request.RemoteAddr, "delete ip", models.IsDel, models.ErrQueryData)
		resComm(c, models.ErrQueryData, nil)
		return
	}
	mID, err := strconv.ParseInt(mIDstr, 10, 0)
	if err != nil {
		config.Log.Error(err)
		models.OO.RecodeLog(user.Username, c.Request.RemoteAddr, "delete ip", models.IsDel, models.ErrQueryData)
		resComm(c, models.ErrQueryData, nil)
		return
	}
	bf := models.DeleteMachine(int(mID))
	models.OO.RecodeLog(user.Username, c.Request.RemoteAddr, "delete ip", models.IsDel, bf)
	resComm(c, bf, nil)
}

func putMachineStatus(c *gin.Context) {
	user := c.Keys["userInfo"].(*models.UserSessionInfo)
	pass := models.CheckPriv(user, "ipManager", "", "dis.enable")
	if pass != models.Success {
		resComm(c, pass, nil)
		return
	}
	ed := models.EnabledInfo{}
	if err := c.BindJSON(&ed); err != nil {
		models.OO.RecodeLog(user.Username, c.Request.RemoteAddr, "update ip status", models.IsUpdate, models.ErrPostData)
		resComm(c, models.ErrPostData, nil)
		return
	}
	bf := models.PutMachineStatus(user, &ed)
	models.OO.RecodeLog(user.Username, c.Request.RemoteAddr, "update ip status", models.IsUpdate, bf)
	resComm(c, bf, nil)
}

func updatePosition(c *gin.Context) {
	user := c.Keys["userInfo"].(*models.UserSessionInfo)
	pass := models.CheckPriv(user, "ipManager", "", "position")
	if pass != models.Success {
		resComm(c, pass, nil)
		return
	}
	if sync.AS.CanNoDo(sync.UpdateAllIPPosition) {
		resComm(c, models.ErrAlreadyRunning, nil)
		return
	}
	bf := models.UpdateIPPosition(user)
	models.OO.RecodeLog(user.Username, c.Request.RemoteAddr, "update ip position", models.IsUpdate, bf)
	sync.AS.Done(sync.UpdateAllIPPosition)
	resComm(c, bf, nil)
}

func getAllMachines(c *gin.Context) {
	user := c.Keys["userInfo"].(*models.UserSessionInfo)
	pass := models.CheckPriv(user, "ipManager", "", "search")
	if pass != models.Success {
		resComm(c, pass, nil)
		return
	}
	data, bf := models.GetAllMachine()
	models.OO.RecodeLog(user.Username, c.Request.RemoteAddr, "get all ip", models.IsSearch, bf)
	resComm(c, bf, data)
}

func batchDeleteMachine(c *gin.Context) {
	user := c.Keys["userInfo"].(*models.UserSessionInfo)
	pass := models.CheckPriv(user, "ipManager", "", "delete")
	if pass != models.Success {
		resComm(c, pass, nil)
		return
	}
	ids := []int{}
	if err := c.BindJSON(&ids); err != nil {
		models.OO.RecodeLog(user.Username, c.Request.RemoteAddr, "delete ip batch", models.IsDel, models.ErrSplitParma)
		resComm(c, models.ErrSplitParma, nil)
		return
	}
	bf := models.BatchDeleteMachine(ids)
	models.OO.RecodeLog(user.Username, c.Request.RemoteAddr, "delete ip batch", models.IsDel, bf)
	resComm(c, bf, nil)
}

func batchEnableMachine(c *gin.Context) {
	user := c.Keys["userInfo"].(*models.UserSessionInfo)
	pass := models.CheckPriv(user, "ipManager", "", "dis.enable")
	if pass != models.Success {
		resComm(c, pass, nil)
		return
	}
	ids := []int{}
	if err := c.BindJSON(&ids); err != nil {
		models.OO.RecodeLog(user.Username, c.Request.RemoteAddr, "enable ip batch", models.IsDel, models.ErrSplitParma)
		resComm(c, models.ErrSplitParma, nil)
		return
	}
	bf := models.BatchEnableMachine(user, ids)
	models.OO.RecodeLog(user.Username, c.Request.RemoteAddr, "enable ip batch", models.IsDel, bf)
	resComm(c, bf, nil)
}

func batchDisableMachine(c *gin.Context) {
	user := c.Keys["userInfo"].(*models.UserSessionInfo)
	pass := models.CheckPriv(user, "ipManager", "", "dis.enable")
	if pass != models.Success {
		resComm(c, pass, nil)
		return
	}
	ids := []int{}
	if err := c.BindJSON(&ids); err != nil {
		models.OO.RecodeLog(user.Username, c.Request.RemoteAddr, "disable ip batch", models.IsDel, models.ErrSplitParma)
		resComm(c, models.ErrSplitParma, nil)
		return
	}
	bf := models.BatchDisableMachine(user, ids)
	models.OO.RecodeLog(user.Username, c.Request.RemoteAddr, "disable ip batch", models.IsDel, bf)
	resComm(c, bf, nil)
}

func uploadMachines(c *gin.Context) {
	user := c.Keys["userInfo"].(*models.UserSessionInfo)
	pass := models.CheckPriv(user, "ipManager", "", "import")
	if pass != models.Success {
		resComm(c, pass, nil)
		return
	}
	info := models.UploadMachinesInfo{}
	if err := c.BindJSON(&info); err != nil {
		models.OO.RecodeLog(user.Username, c.Request.RemoteAddr, "import ip", models.IsAdd, models.ErrUpdateData)
		resComm(c, models.ErrUpdateData, nil)
		return
	}
	result, bf := models.UploadMachines(user, &info)
	models.OO.RecodeLog(user.Username, c.Request.RemoteAddr, "import ip", models.IsAdd, bf)
	resComm(c, bf, result)
}

func batchImportIPAddrs(c *gin.Context) {
	user := c.Keys["userInfo"].(*models.UserSessionInfo)
	pass := models.CheckPriv(user, "ipManager", "", "import")
	if pass != models.Success {
		resComm(c, pass, nil)
		return
	}
	content := models.BatchImportIPaddrs{}
	if err := c.BindJSON(&content); err != nil {
		models.OO.RecodeLog(user.Username, c.Request.RemoteAddr, "batch import ip[web]", models.IsAdd, models.ErrUpdateData)
		resComm(c, models.ErrUpdateData, nil)
		return
	}
	_, bf := models.BatchImportIPAddrs(user, &content)
	models.OO.RecodeLog(user.Username, c.Request.RemoteAddr, "batch import ip[web]", models.IsAdd, bf)
	resComm(c, bf, nil)
}

func batchImportDomain(c *gin.Context) {
	user := c.Keys["userInfo"].(*models.UserSessionInfo)
	pass := models.CheckPriv(user, "ipManager", "", "import")
	if pass != models.Success {
		resComm(c, pass, nil)
		return
	}
	content := models.BatchImportIPaddrs{}
	if err := c.BindJSON(&content); err != nil {
		models.OO.RecodeLog(user.Username, c.Request.RemoteAddr, "batch import domain", models.IsAdd, models.ErrUpdateData)
		resComm(c, models.ErrUpdateData, nil)
		return
	}
	_, bf := models.BatchImportDomain(user, &content)
	models.OO.RecodeLog(user.Username, c.Request.RemoteAddr, "batch import domain", models.IsAdd, bf)
	resComm(c, bf, nil)
}

func publish(c *gin.Context) {
	user := c.Keys["userInfo"].(*models.UserSessionInfo)
	pass := models.CheckPriv(user, "ipManager", "", "publish")
	if pass != models.Success {
		resComm(c, pass, nil)
		return
	}
	bf := models.Publish()
	models.OO.RecodeLog(user.Username, c.Request.RemoteAddr, "publish sd_file", models.IsRunning, bf)
	resComm(c, bf, nil)
}

func preview(c *gin.Context) {
	user := c.Keys["userInfo"].(*models.UserSessionInfo)
	pass := models.CheckPriv(user, "preview", "", "search")
	if pass != models.Success {
		resComm(c, pass, nil)
		return
	}
	data, bf := models.Preview()
	models.OO.RecodeLog(user.Username, c.Request.RemoteAddr, "view prometheus.yml", models.IsSearch, bf)
	resComm(c, bf, data)
}

func reload(c *gin.Context) {
	user := c.Keys["userInfo"].(*models.UserSessionInfo)
	pass := models.CheckPriv(user, "jobs", "jobs", "reload")
	if pass != models.Success {
		resComm(c, pass, nil)
		return
	}
	bf := models.Reload(config.Cfg.PrometheusCfg.Api)
	models.OO.RecodeLog(user.Username, c.Request.RemoteAddr, "reload prometheus service", models.IsRunning, bf)
	resComm(c, bf, nil)
}

func defReload(c *gin.Context) {
	user := c.Keys["userInfo"].(*models.UserSessionInfo)
	pass := models.CheckPriv(user, "baseConfig", "defaultJobs", "reload")
	if pass != models.Success {
		resComm(c, pass, nil)
		return
	}
	bf := models.Reload(config.Cfg.PrometheusCfg.Api)
	models.OO.RecodeLog(user.Username, c.Request.RemoteAddr, "reload prometheus config[def]", models.IsRunning, bf)
	resComm(c, bf, nil)
}

func allFileList(c *gin.Context) {
	user := c.Keys["userInfo"].(*models.UserSessionInfo)
	pass := models.CheckPriv(user, "ftree", "", "list_all_file")
	if pass != models.Success {
		resComm(c, pass, nil)
		return
	}
	data, bf := models.AllFileList()
	models.OO.RecodeLog(user.Username, c.Request.RemoteAddr, "get all file list", models.IsSearch, bf)
	resComm(c, bf, data)
}

func fileContent(c *gin.Context) {
	user := c.Keys["userInfo"].(*models.UserSessionInfo)
	pass := models.CheckPriv(user, "ftree", "", "view_file")
	if pass != models.Success {
		resComm(c, pass, nil)
		return
	}
	child := models.Child{}
	if err := c.BindJSON(&child); err != nil {
		models.OO.RecodeLog(user.Username, c.Request.RemoteAddr, "view sd_file content", models.IsSearch, models.ErrPostData)
		resComm(c, models.ErrPostData, nil)
		return
	}
	data, bf := models.ReadFileContent(&child)
	models.OO.RecodeLog(user.Username, c.Request.RemoteAddr, "view sd_file content", models.IsSearch, bf)
	resComm(c, bf, data)
}

func allRulesFileList(c *gin.Context) {
	user := c.Keys["userInfo"].(*models.UserSessionInfo)
	pass := models.CheckPriv(user, "ruleView", "", "list_all_file")
	if pass != models.Success {
		resComm(c, pass, nil)
		return
	}
	data, bf := models.AllRulesFileList()
	models.OO.RecodeLog(user.Username, c.Request.RemoteAddr, "get all rule file list", models.IsSearch, bf)
	resComm(c, bf, data)
}

func ruleFileContent(c *gin.Context) {
	user := c.Keys["userInfo"].(*models.UserSessionInfo)
	pass := models.CheckPriv(user, "ruleView", "", "view_file")
	if pass != models.Success {
		resComm(c, pass, nil)
		return
	}
	child := models.Child{}
	if err := c.BindJSON(&child); err != nil {
		models.OO.RecodeLog(user.Username, c.Request.RemoteAddr, "get rule file content", models.IsSearch, models.ErrPostData)
		resComm(c, models.ErrPostData, nil)
		return
	}
	data, bf := models.ReadRuleFileContent(&child)
	models.OO.RecodeLog(user.Username, c.Request.RemoteAddr, "get rule file content", models.IsSearch, bf)
	resComm(c, bf, data)
}

func getTree(c *gin.Context) {
	user := c.Keys["userInfo"].(*models.UserSessionInfo)
	pass := models.CheckPriv(user, "noticeManager", "", "search_tree")
	if pass != models.Success {
		resComm(c, pass, nil)
		return
	}
	data, bf := models.GetNodesFromDB()
	models.OO.RecodeLog(user.Username, c.Request.RemoteAddr, "get monitor tree", models.IsSearch, bf)
	resComm(c, bf, data)
}

func getNode(c *gin.Context) {
	user := c.Keys["userInfo"].(*models.UserSessionInfo)
	pass := models.CheckPriv(user, "noticeManager", "", "search_rule")
	if pass != models.Success {
		resComm(c, pass, nil)
		return
	}
	qni := models.QueryGetNode{}
	if err := c.BindQuery(&qni); err != nil {
		models.OO.RecodeLog(user.Username, c.Request.RemoteAddr, "get node", models.IsSearch, models.ErrQueryData)
		resComm(c, models.ErrQueryData, nil)
		return
	}
	data, bf := models.GetNode(&qni)
	models.OO.RecodeLog(user.Username, c.Request.RemoteAddr, "get node", models.IsSearch, bf)
	resComm(c, bf, data)
}

func postNode(c *gin.Context) {
	user := c.Keys["userInfo"].(*models.UserSessionInfo)
	pass := models.CheckPriv(user, "noticeManager", "", "update_rule")
	if pass != models.Success {
		resComm(c, pass, nil)
		return
	}
	nodeInfo := models.TreeNodeInfo{}
	if err := c.BindJSON(&nodeInfo); err != nil {
		models.OO.RecodeLog(user.Username, c.Request.RemoteAddr, "create node", models.IsAdd, models.ErrPostData)
		resComm(c, models.ErrPostData, nil)
		return
	}
	data, bf := models.PostNode(user, &nodeInfo)
	models.OO.RecodeLog(user.Username, c.Request.RemoteAddr, "create node", models.IsAdd, bf)
	resComm(c, bf, data)
}

func putNode(c *gin.Context) {
	user := c.Keys["userInfo"].(*models.UserSessionInfo)
	pass := models.CheckPriv(user, "noticeManager", "", "update_rule")
	if pass != models.Success {
		resComm(c, pass, nil)
		return
	}
	nodeInfo := models.TreeNodeInfo{}
	if err := c.BindJSON(&nodeInfo); err != nil {
		models.OO.RecodeLog(user.Username, c.Request.RemoteAddr, "update monitor rule", models.IsUpdate, models.ErrPostData)
		resComm(c, models.ErrPostData, nil)
		return
	}
	data, bf := models.PutNode(user, &nodeInfo)
	models.OO.RecodeLog(user.Username, c.Request.RemoteAddr, "update monitor rule", models.IsUpdate, bf)
	resComm(c, bf, data)
}

func getDefLabels(c *gin.Context) {
	// user := c.Keys["userInfo"].(*models.UserSessionInfo)
	// pass := models.CheckPriv(user, "baseConfig", "baseLabels", "search")
	// if pass != models.Success {
	//     resComm(c, pass, nil)
	//     return
	// }
	data, bf := models.GetDefLabels()
	resComm(c, bf, data)
}

func getDefEnableLabels(c *gin.Context) {
	// user := c.Keys["userInfo"].(*models.UserSessionInfo)
	// pass := models.CheckPriv(user, "baseConfig", "baseLabels", "search")
	// if pass != models.Success {
	//     resComm(c, pass, nil)
	//     return
	// }
	data, bf := models.GetDefEnableLabels()
	resComm(c, bf, data)
}

func delLabel(c *gin.Context) {
	user := c.Keys["userInfo"].(*models.UserSessionInfo)
	pass := models.CheckPriv(user, "noticeManager", "", "update_rule")
	if pass != models.Success {
		resComm(c, pass, nil)
		return
	}
	labelType, ok := c.GetQuery("type")
	if !ok {
		models.OO.RecodeLog(user.Username, c.Request.RemoteAddr, "delete label of monitor rule", models.IsDel, models.ErrQueryData)
		resComm(c, models.ErrQueryData, nil)
		return
	}
	labelInfo := models.Label{}
	if err := c.BindJSON(&labelInfo); err != nil {
		models.OO.RecodeLog(user.Username, c.Request.RemoteAddr, "delete label of monitor rule", models.IsDel, models.ErrPostData)
		resComm(c, models.ErrPostData, nil)
		return
	}
	bf := models.DelLabel(&labelInfo, labelType)
	models.OO.RecodeLog(user.Username, c.Request.RemoteAddr, "delete label of monitor rule", models.IsDel, bf)
	resComm(c, bf, nil)
}

func createTreeNode(c *gin.Context) {
	user := c.Keys["userInfo"].(*models.UserSessionInfo)
	tnc := models.TreeNodeFromCli{}
	if err := c.BindJSON(&tnc); err != nil {
		resComm(c, models.ErrPostData, nil)
		return
	}
	bf := models.CreateTreeNode(user, &tnc)
	models.OO.RecodeLog(user.Username, c.Request.RemoteAddr, "create tree node", models.IsAdd, bf)
	resComm(c, bf, nil)
}

func updateTreeNode(c *gin.Context) {
	user := c.Keys["userInfo"].(*models.UserSessionInfo)
	pass := models.CheckPriv(user, "noticeManager", "", "rename_node")
	if pass != models.Success {
		resComm(c, pass, nil)
		return
	}
	tnc := models.TreeNodeFromCli{}
	if err := c.BindJSON(&tnc); err != nil {
		models.OO.RecodeLog(user.Username, c.Request.RemoteAddr, "update tree node", models.IsUpdate, models.ErrPostData)
		resComm(c, models.ErrPostData, nil)
		return
	}
	bf := models.UpdateTreeNode(user, &tnc)
	models.OO.RecodeLog(user.Username, c.Request.RemoteAddr, "update tree node", models.IsUpdate, bf)
	resComm(c, bf, nil)
}

func deleteTreeNode(c *gin.Context) {
	user := c.Keys["userInfo"].(*models.UserSessionInfo)
	var skipSelf bool
	if strings.ToLower(c.Query("skip_self")) == "true" {
		pass := models.CheckPriv(user, "noticeManager", "", "delete_tree_node")
		if pass != models.Success {
			resComm(c, pass, nil)
			return
		}
		skipSelf = true
	} else {
		pass := models.CheckPriv(user, "noticeManager", "", "delete_node")
		if pass != models.Success {
			resComm(c, pass, nil)
			return
		}
	}
	tnc := models.TreeNodeFromCli{}
	if err := c.BindJSON(&tnc); err != nil {
		models.OO.RecodeLog(user.Username, c.Request.RemoteAddr, "delete tree node", models.IsDel, models.ErrPostData)
		resComm(c, models.ErrPostData, nil)
		return
	}
	bf := models.DeleteTreeNode(&tnc, skipSelf)
	models.OO.RecodeLog(user.Username, c.Request.RemoteAddr, "delete tree node", models.IsDel, bf)
	resComm(c, bf, nil)
}

func rulePublish(c *gin.Context) {
	user := c.Keys["userInfo"].(*models.UserSessionInfo)
	pass := models.CheckPriv(user, "noticeManager", "", "publish_rule")
	if pass != models.Success {
		resComm(c, pass, nil)
		return
	}
	if sync.AS.CanNoDo(sync.PublishMonitorRules) {
		resComm(c, models.ErrAlreadyRunning, nil)
		return
	}
	bf := models.RulePublish()
	models.OO.RecodeLog(user.Username, c.Request.RemoteAddr, "publish rule", models.IsPublish, bf)
	alertmgr.AlertMgrObj.NoticeReload()
	sync.AS.Done(sync.PublishMonitorRules)
	resComm(c, bf, nil)
}

func emptyRulePublish(c *gin.Context) {
	user := c.Keys["userInfo"].(*models.UserSessionInfo)
	pass := models.CheckPriv(user, "noticeManager", "", "publish_empty_rule")
	if pass != models.Success {
		resComm(c, pass, nil)
		return
	}
	if sync.AS.CanNoDo(sync.PublishEmptyMonitorRules) {
		resComm(c, models.ErrAlreadyRunning, nil)
		return
	}
	bf := models.EmptyRulePublish()
	models.OO.RecodeLog(user.Username, c.Request.RemoteAddr, "publish empty rule", models.IsPublish, bf)
	alertmgr.AlertMgrObj.NoticeReload()
	sync.AS.Done(sync.PublishEmptyMonitorRules)
	resComm(c, bf, nil)
}

func putTreeNodeStatus(c *gin.Context) {
	user := c.Keys["userInfo"].(*models.UserSessionInfo)
	pass := models.CheckPriv(user, "noticeManager", "", "dis.enable")
	if pass != models.Success {
		resComm(c, pass, nil)
		return
	}
	upInfo := models.TreeNodeStatus{}
	if err := c.BindJSON(&upInfo); err != nil {
		models.OO.RecodeLog(user.Username, c.Request.RemoteAddr, "update tree node status", models.IsUpdate, models.ErrPostData)
		resComm(c, models.ErrPostData, nil)
		return
	}
	bf := models.PutTreeNodeStatus(user, &upInfo)
	models.OO.RecodeLog(user.Username, c.Request.RemoteAddr, "update tree node status", models.IsUpdate, bf)
	resComm(c, bf, nil)
}

func postTreeUploadFileYaml(c *gin.Context) {
	user := c.Keys["userInfo"].(*models.UserSessionInfo)
	pass := models.CheckPriv(user, "noticeManager", "", "import_rule_from_file")
	if pass != models.Success {
		resComm(c, pass, nil)
		return
	}
	gidStr, ok := c.GetQuery("gid")
	if !ok {
		models.OO.RecodeLog(user.Username, c.Request.RemoteAddr, "import rule from file", models.IsAdd, models.ErrQueryData)
		resComm(c, models.ErrQueryData, nil)
		return
	}
	gid, err := strconv.ParseInt(gidStr, 10, 0)
	if err != nil {
		config.Log.Error(err)
		models.OO.RecodeLog(user.Username, c.Request.RemoteAddr, "import rule from file", models.IsAdd, models.ErrQueryData)
		resComm(c, models.ErrQueryData, nil)
		return
	}
	r, bf := models.PostTreeUploadFileYaml(c, user, gid)
	models.OO.RecodeLog(user.Username, c.Request.RemoteAddr, "import rule from file", models.IsAdd, bf)
	resComm(c, bf, r)
}

func getBaseLabels(c *gin.Context) {
	user := c.Keys["userInfo"].(*models.UserSessionInfo)
	pass := models.CheckPriv(user, "baseConfig", "baseLabels", "search")
	if pass != models.Success {
		resComm(c, pass, nil)
		return
	}
	sp := &models.SplitPage{}
	if err := c.BindQuery(sp); err != nil {
		models.OO.RecodeLog(user.Username, c.Request.RemoteAddr, "get baseLabels", models.IsSearch, models.ErrSplitParma)
		resComm(c, models.ErrSplitParma, nil)
		return
	}
	data, bf := models.GetBaseLabels(sp)
	models.OO.RecodeLog(user.Username, c.Request.RemoteAddr, "get baseLabels", models.IsSearch, bf)
	resComm(c, bf, data)
}

func postBaseLabels(c *gin.Context) {
	user := c.Keys["userInfo"].(*models.UserSessionInfo)
	pass := models.CheckPriv(user, "baseConfig", "baseLabels", "add")
	if pass != models.Success {
		resComm(c, pass, nil)
		return
	}
	newLabel := models.BaseLabels{}
	if err := c.BindJSON(&newLabel); err != nil {
		models.OO.RecodeLog(user.Username, c.Request.RemoteAddr, "create baseLabels", models.IsAdd, models.ErrPostData)
		resComm(c, models.ErrPostData, nil)
		return
	}
	bf := models.PostBaseLabels(user, &newLabel)
	models.OO.RecodeLog(user.Username, c.Request.RemoteAddr, "create baseLabels", models.IsAdd, bf)
	resComm(c, bf, nil)
}

func putBaseLabels(c *gin.Context) {
	user := c.Keys["userInfo"].(*models.UserSessionInfo)
	pass := models.CheckPriv(user, "baseConfig", "baseLabels", "update")
	if pass != models.Success {
		resComm(c, pass, nil)
		return
	}
	label := models.BaseLabels{}
	if err := c.BindJSON(&label); err != nil {
		models.OO.RecodeLog(user.Username, c.Request.RemoteAddr, "update baseLabels", models.IsUpdate, models.ErrPostData)
		resComm(c, models.ErrPostData, nil)
		return
	}
	bf := models.PutBaseLabels(user, &label)
	models.OO.RecodeLog(user.Username, c.Request.RemoteAddr, "update baseLabels", models.IsUpdate, bf)
	resComm(c, bf, nil)
}

func delBaseLabels(c *gin.Context) {
	user := c.Keys["userInfo"].(*models.UserSessionInfo)
	pass := models.CheckPriv(user, "baseConfig", "baseLabels", "delete")
	if pass != models.Success {
		resComm(c, pass, nil)
		return
	}
	idStr, ok := c.GetQuery("id")
	if !ok {
		models.OO.RecodeLog(user.Username, c.Request.RemoteAddr, "delete baseLabels", models.IsDel, models.ErrQueryData)
		resComm(c, models.ErrQueryData, nil)
		return
	}
	id, err := strconv.ParseInt(idStr, 10, 0)
	if err != nil {
		config.Log.Error(err)
		models.OO.RecodeLog(user.Username, c.Request.RemoteAddr, "delete baseLabels", models.IsDel, models.ErrQueryData)
		resComm(c, models.ErrQueryData, nil)
		return
	}
	bf := models.DelBaseLabels(id)
	models.OO.RecodeLog(user.Username, c.Request.RemoteAddr, "delete baseLabels", models.IsDel, bf)
	resComm(c, bf, nil)
}

func getReLabels(c *gin.Context) {
	user := c.Keys["userInfo"].(*models.UserSessionInfo)
	pass := models.CheckPriv(user, "baseConfig", "reLabels", "search")
	if pass != models.Success {
		resComm(c, pass, nil)
		return
	}
	sp := &models.SplitPage{}
	if err := c.BindQuery(sp); err != nil {
		models.OO.RecodeLog(user.Username, c.Request.RemoteAddr, "get reLabels", models.IsSearch, models.ErrSplitParma)
		resComm(c, models.ErrSplitParma, nil)
		return
	}
	data, bf := models.GetReLabels(sp)
	models.OO.RecodeLog(user.Username, c.Request.RemoteAddr, "get reLabels", models.IsSearch, bf)
	resComm(c, bf, data)
}

func getRelabel(c *gin.Context) {
	user := c.Keys["userInfo"].(*models.UserSessionInfo)
	pass := models.CheckPriv(user, "baseConfig", "reLabels", "search")
	if pass != models.Success {
		resComm(c, pass, nil)
		return
	}
	id := &models.OnlyID{}
	if err := c.BindQuery(id); err != nil {
		models.OO.RecodeLog(user.Username, c.Request.RemoteAddr, "get reLabel", models.IsSearch, models.ErrSplitParma)
		resComm(c, models.ErrSplitParma, nil)
		return
	}
	data, bf := models.GetReLabel(id)
	models.OO.RecodeLog(user.Username, c.Request.RemoteAddr, "get reLabel", models.IsSearch, bf)
	resComm(c, bf, data)
}

func getAllRelabels(c *gin.Context) {
	user := c.Keys["userInfo"].(*models.UserSessionInfo)
	pass := models.CheckPriv(user, "baseConfig", "reLabels", "search")
	if pass != models.Success {
		resComm(c, pass, nil)
		return
	}
	data, bf := models.GetAllReLabels()
	models.OO.RecodeLog(user.Username, c.Request.RemoteAddr, "get all reLabels", models.IsSearch, bf)
	resComm(c, bf, data)
}

func postReLabels(c *gin.Context) {
	user := c.Keys["userInfo"].(*models.UserSessionInfo)
	pass := models.CheckPriv(user, "baseConfig", "reLabels", "add")
	if pass != models.Success {
		resComm(c, pass, nil)
		return
	}
	reLabel := models.ReLabels{}
	if err := c.BindJSON(&reLabel); err != nil {
		models.OO.RecodeLog(user.Username, c.Request.RemoteAddr, "create reLabels", models.IsAdd, models.ErrPostData)
		resComm(c, models.ErrPostData, nil)
		return
	}
	bf := models.PostReLabels(user, &reLabel)
	models.OO.RecodeLog(user.Username, c.Request.RemoteAddr, "create reLabels", models.IsAdd, bf)
	resComm(c, bf, nil)
}

func putReLabels(c *gin.Context) {
	user := c.Keys["userInfo"].(*models.UserSessionInfo)
	pass := models.CheckPriv(user, "baseConfig", "reLabels", "edit_name")
	if pass != models.Success {
		resComm(c, pass, nil)
		return
	}
	reLabel := models.ReLabels{}
	if err := c.BindJSON(&reLabel); err != nil {
		models.OO.RecodeLog(user.Username, c.Request.RemoteAddr, "update sub group", models.IsUpdate, models.ErrPostData)
		resComm(c, models.ErrPostData, nil)
		return
	}
	bf := models.PutReLabels(user, &reLabel)
	models.OO.RecodeLog(user.Username, c.Request.RemoteAddr, "update sub group", models.IsUpdate, bf)
	resComm(c, bf, nil)
}

func delReLabels(c *gin.Context) {
	user := c.Keys["userInfo"].(*models.UserSessionInfo)
	pass := models.CheckPriv(user, "baseConfig", "reLabels", "delete")
	if pass != models.Success {
		resComm(c, pass, nil)
		return
	}
	idStr, ok := c.GetQuery("id")
	if !ok {
		models.OO.RecodeLog(user.Username, c.Request.RemoteAddr, "delete reLabels", models.IsDel, models.ErrQueryData)
		resComm(c, models.ErrQueryData, nil)
		return
	}
	id, err := strconv.ParseInt(idStr, 10, 0)
	if err != nil {
		config.Log.Error(err)
		models.OO.RecodeLog(user.Username, c.Request.RemoteAddr, "delete reLabels", models.IsDel, models.ErrQueryData)
		resComm(c, models.ErrQueryData, nil)
		return
	}
	bf := models.DelReLabels(id)
	models.OO.RecodeLog(user.Username, c.Request.RemoteAddr, "delete reLabels", models.IsDel, bf)
	resComm(c, bf, nil)
}

func putRelabelsCode(c *gin.Context) {
	user := c.Keys["userInfo"].(*models.UserSessionInfo)
	pass := models.CheckPriv(user, "baseConfig", "reLabels", "update_rule")
	if pass != models.Success {
		resComm(c, pass, nil)
		return
	}
	reLabel := models.ReLabels{}
	if err := c.BindJSON(&reLabel); err != nil {
		models.OO.RecodeLog(user.Username, c.Request.RemoteAddr, "update reLabels code", models.IsUpdate, models.ErrPostData)
		resComm(c, models.ErrPostData, nil)
		return
	}
	bf := models.PutReLabelsCode(user, &reLabel)
	models.OO.RecodeLog(user.Username, c.Request.RemoteAddr, "update reLabels code", models.IsUpdate, bf)
	resComm(c, bf, nil)
}

func putBaseLabelsStatus(c *gin.Context) {
	user := c.Keys["userInfo"].(*models.UserSessionInfo)
	pass := models.CheckPriv(user, "baseConfig", "baseLabels", "dis.enable")
	if pass != models.Success {
		resComm(c, pass, nil)
		return
	}
	ed := models.EnabledInfo{}
	if err := c.BindJSON(&ed); err != nil {
		models.OO.RecodeLog(user.Username, c.Request.RemoteAddr, "update baseLabels status", models.IsUpdate, models.ErrPostData)
		resComm(c, models.ErrPostData, nil)
		return
	}
	bf := models.PutBaseLabelsStatus(user, &ed)
	models.OO.RecodeLog(user.Username, c.Request.RemoteAddr, "update baseLabels status", models.IsUpdate, bf)
	resComm(c, bf, nil)
}

func getRelabelsViewCodePriv(c *gin.Context) {
	user := c.Keys["userInfo"].(*models.UserSessionInfo)
	pass := models.CheckPriv(user, "baseConfig", "reLabels", "view_code")
	// models.OO.RecodeLog(user.Username, c.Request.RemoteAddr, "update sub group", models.IsUpdate, bf)
	resComm(c, pass, nil)
}

func putBaseRelabelsStatus(c *gin.Context) {
	user := c.Keys["userInfo"].(*models.UserSessionInfo)
	pass := models.CheckPriv(user, "baseConfig", "reLabels", "dis.enable")
	if pass != models.Success {
		resComm(c, pass, nil)
		return
	}
	ed := models.EnabledInfo{}
	if err := c.BindJSON(&ed); err != nil {
		models.OO.RecodeLog(user.Username, c.Request.RemoteAddr, "update reLabels", models.IsUpdate, models.ErrPostData)
		resComm(c, models.ErrPostData, nil)
		return
	}
	bf := models.PutBaseRelabelsStatus(user, &ed)
	models.OO.RecodeLog(user.Username, c.Request.RemoteAddr, "update reLabels", models.IsUpdate, bf)
	resComm(c, bf, nil)
}

func getBaseFields(c *gin.Context) {
	user := c.Keys["userInfo"].(*models.UserSessionInfo)
	pass := models.CheckPriv(user, "baseConfig", "baseFields", "search")
	if pass != models.Success {
		resComm(c, pass, nil)
		return
	}
	sp := &models.SplitPage{}
	if err := c.BindQuery(sp); err != nil {
		models.OO.RecodeLog(user.Username, c.Request.RemoteAddr, "get base fields", models.IsSearch, models.ErrSplitParma)
		resComm(c, models.ErrSplitParma, nil)
		return
	}
	data, bf := models.GetBaseFields(sp)
	models.OO.RecodeLog(user.Username, c.Request.RemoteAddr, "get base fields", models.IsSearch, bf)
	resComm(c, bf, data)
}

func postBaseFields(c *gin.Context) {
	user := c.Keys["userInfo"].(*models.UserSessionInfo)
	pass := models.CheckPriv(user, "baseConfig", "baseFields", "add")
	if pass != models.Success {
		resComm(c, pass, nil)
		return
	}
	newLabel := models.BaseFields{}
	if err := c.BindJSON(&newLabel); err != nil {
		models.OO.RecodeLog(user.Username, c.Request.RemoteAddr, "create base fields", models.IsAdd, models.ErrPostData)
		resComm(c, models.ErrPostData, nil)
		return
	}
	bf := models.PostBaseFields(user, &newLabel)
	models.OO.RecodeLog(user.Username, c.Request.RemoteAddr, "create base fields", models.IsAdd, bf)
	resComm(c, bf, nil)
}

func putBaseFields(c *gin.Context) {
	user := c.Keys["userInfo"].(*models.UserSessionInfo)
	pass := models.CheckPriv(user, "baseConfig", "baseFields", "update")
	if pass != models.Success {
		resComm(c, pass, nil)
		return
	}
	label := models.BaseFields{}
	if err := c.BindJSON(&label); err != nil {
		models.OO.RecodeLog(user.Username, c.Request.RemoteAddr, "update base fields", models.IsUpdate, models.ErrPostData)
		resComm(c, models.ErrPostData, nil)
		return
	}
	bf := models.PutBaseFields(user, &label)
	models.OO.RecodeLog(user.Username, c.Request.RemoteAddr, "update base fields", models.IsUpdate, bf)
	resComm(c, bf, nil)
}

func putBaseFieldsStatus(c *gin.Context) {
	user := c.Keys["userInfo"].(*models.UserSessionInfo)
	pass := models.CheckPriv(user, "baseConfig", "baseFields", "dis.enable")
	if pass != models.Success {
		resComm(c, pass, nil)
		return
	}
	ed := models.EnabledInfo{}
	if err := c.BindJSON(&ed); err != nil {
		models.OO.RecodeLog(user.Username, c.Request.RemoteAddr, "update base fields status", models.IsUpdate, models.ErrPostData)
		resComm(c, models.ErrPostData, nil)
		return
	}
	bf := models.PutBaseFieldsStatus(user, &ed)
	models.OO.RecodeLog(user.Username, c.Request.RemoteAddr, "update base fields status", models.IsUpdate, bf)
	resComm(c, bf, nil)
}

func delBaseFields(c *gin.Context) {
	user := c.Keys["userInfo"].(*models.UserSessionInfo)
	pass := models.CheckPriv(user, "baseConfig", "baseFields", "delete")
	if pass != models.Success {
		resComm(c, pass, nil)
		return
	}
	idStr, ok := c.GetQuery("id")
	if !ok {
		models.OO.RecodeLog(user.Username, c.Request.RemoteAddr, "delete base fields", models.IsDel, models.ErrQueryData)
		resComm(c, models.ErrQueryData, nil)
		return
	}
	id, err := strconv.ParseInt(idStr, 10, 0)
	if err != nil {
		config.Log.Error(err)
		models.OO.RecodeLog(user.Username, c.Request.RemoteAddr, "delete base fields", models.IsDel, models.ErrQueryData)
		resComm(c, models.ErrQueryData, nil)
		return
	}
	bf := models.DelBaseFields(id)
	models.OO.RecodeLog(user.Username, c.Request.RemoteAddr, "delete base fields", models.IsDel, bf)
	resComm(c, bf, nil)
}

func getCronRules(c *gin.Context) {
	user := c.Keys["userInfo"].(*models.UserSessionInfo)
	pass := models.CheckPriv(user, "crontab", "", "search")
	if pass != models.Success {
		resComm(c, pass, nil)
		return
	}
	sp := &models.SplitPage{}
	if err := c.BindQuery(sp); err != nil {
		models.OO.RecodeLog(user.Username, c.Request.RemoteAddr, "get crontab jobs", models.IsSearch, models.ErrSplitParma)
		resComm(c, models.ErrSplitParma, nil)
		return
	}
	result, bf := models.GetCronRules(sp)
	alertmgr.MergeRuleStatus(result)
	models.OO.RecodeLog(user.Username, c.Request.RemoteAddr, "get crontab jobs", models.IsSearch, bf)
	resComm(c, bf, result)
}

func postRule(c *gin.Context) {
	user := c.Keys["userInfo"].(*models.UserSessionInfo)
	pass := models.CheckPriv(user, "crontab", "", "add")
	if pass != models.Success {
		resComm(c, pass, nil)
		return
	}
	newApi := models.CrontabPost{}
	if err := c.BindJSON(&newApi); err != nil {
		models.OO.RecodeLog(user.Username, c.Request.RemoteAddr, "create cron rule", models.IsAdd, models.ErrPostData)
		resComm(c, models.ErrPostData, nil)
		return
	}
	bf := models.PostCronRule(user, &newApi)
	models.OO.RecodeLog(user.Username, c.Request.RemoteAddr, "create cron rule", models.IsAdd, bf)
	resComm(c, bf, nil)
}

func putRule(c *gin.Context) {
	user := c.Keys["userInfo"].(*models.UserSessionInfo)
	pass := models.CheckPriv(user, "crontab", "", "update")
	if pass != models.Success {
		resComm(c, pass, nil)
		return
	}
	label := models.CrontabPost{}
	if err := c.BindJSON(&label); err != nil {
		models.OO.RecodeLog(user.Username, c.Request.RemoteAddr, "update cron rule", models.IsUpdate, models.ErrPostData)
		resComm(c, models.ErrPostData, nil)
		return
	}
	bf := models.PutCronRule(user, &label)
	models.OO.RecodeLog(user.Username, c.Request.RemoteAddr, "update cron rule", models.IsUpdate, bf)
	resComm(c, bf, nil)
}

func deleteRule(c *gin.Context) {
	user := c.Keys["userInfo"].(*models.UserSessionInfo)
	pass := models.CheckPriv(user, "crontab", "", "delete")
	if pass != models.Success {
		resComm(c, pass, nil)
		return
	}
	idStr, ok := c.GetQuery("id")
	if !ok {
		models.OO.RecodeLog(user.Username, c.Request.RemoteAddr, "delete cron rule", models.IsDel, models.ErrQueryData)
		resComm(c, models.ErrQueryData, nil)
		return
	}
	id, err := strconv.ParseInt(idStr, 10, 0)
	if err != nil {
		config.Log.Error(err)
		models.OO.RecodeLog(user.Username, c.Request.RemoteAddr, "delete cron rule", models.IsDel, models.ErrQueryData)
		resComm(c, models.ErrQueryData, nil)
		return
	}
	bf := models.DelCronRule(id)
	models.OO.RecodeLog(user.Username, c.Request.RemoteAddr, "delete cron rule", models.IsDel, bf)
	resComm(c, bf, nil)
}

func putRuleStatus(c *gin.Context) {
	user := c.Keys["userInfo"].(*models.UserSessionInfo)
	pass := models.CheckPriv(user, "crontab", "", "dis.enable")
	if pass != models.Success {
		resComm(c, pass, nil)
		return
	}
	ed := models.EnabledInfo{}
	if err := c.BindJSON(&ed); err != nil {
		models.OO.RecodeLog(user.Username, c.Request.RemoteAddr, "update cron rule status", models.IsUpdate, models.ErrPostData)
		resComm(c, models.ErrPostData, nil)
		return
	}
	bf := models.PutCronRuleStatus(user, &ed)
	models.OO.RecodeLog(user.Username, c.Request.RemoteAddr, "update cron rule status", models.IsUpdate, bf)
	resComm(c, bf, nil)
}

func delRulesSelection(c *gin.Context) {
	user := c.Keys["userInfo"].(*models.UserSessionInfo)
	pass := models.CheckPriv(user, "crontab", "", "delete")
	if pass != models.Success {
		resComm(c, pass, nil)
		return
	}
	ids := []int{}
	if err := c.BindJSON(&ids); err != nil {
		models.OO.RecodeLog(user.Username, c.Request.RemoteAddr, "delete cron rules batch", models.IsDel, models.ErrSplitParma)
		resComm(c, models.ErrSplitParma, nil)
		return
	}
	bf := models.DelRulesSelection(ids)
	models.OO.RecodeLog(user.Username, c.Request.RemoteAddr, "delete cron rules batch", models.IsDel, bf)
	resComm(c, bf, nil)
}

func putRulesEnable(c *gin.Context) {
	user := c.Keys["userInfo"].(*models.UserSessionInfo)
	pass := models.CheckPriv(user, "crontab", "", "dis.enable")
	if pass != models.Success {
		resComm(c, pass, nil)
		return
	}
	ids := []int{}
	if err := c.BindJSON(&ids); err != nil {
		models.OO.RecodeLog(user.Username, c.Request.RemoteAddr, "enable cron rules batch", models.IsUpdate, models.ErrPostData)
		resComm(c, models.ErrSplitParma, nil)
		return
	}
	bf := models.PutRulesEnable(user, ids)
	models.OO.RecodeLog(user.Username, c.Request.RemoteAddr, "enable cron rules batch", models.IsUpdate, bf)
	resComm(c, bf, nil)
}

func putRulesDisable(c *gin.Context) {
	user := c.Keys["userInfo"].(*models.UserSessionInfo)
	pass := models.CheckPriv(user, "crontab", "", "dis.enable")
	if pass != models.Success {
		resComm(c, pass, nil)
		return
	}
	ids := []int{}
	if err := c.BindJSON(&ids); err != nil {
		models.OO.RecodeLog(user.Username, c.Request.RemoteAddr, "disable cron rules batch", models.IsUpdate, models.ErrPostData)
		resComm(c, models.ErrSplitParma, nil)
		return
	}
	bf := models.PutRulesDisable(user, ids)
	models.OO.RecodeLog(user.Username, c.Request.RemoteAddr, "disable cron rules batch", models.IsUpdate, bf)
	resComm(c, bf, nil)
}

func putRulesPublish(c *gin.Context) {
	user := c.Keys["userInfo"].(*models.UserSessionInfo)
	pass := models.CheckPriv(user, "crontab", "", "publish")
	if pass != models.Success {
		resComm(c, pass, nil)
		return
	}
	alertmgr.ChartCronObj.NoticeReload()
	models.OO.RecodeLog(user.Username, c.Request.RemoteAddr, "publish cron rules", models.IsPublish, models.Success)
	resComm(c, models.Success, nil)
}

func getRuleImage(c *gin.Context) {
	user := c.Keys["userInfo"].(*models.UserSessionInfo)
	pass := models.CheckPriv(user, "crontab", "", "image")
	if pass != models.Success {
		resComm(c, pass, nil)
		return
	}
	id := models.OnlyID{}
	if err := c.BindQuery(&id); err != nil {
		models.OO.RecodeLog(user.Username, c.Request.RemoteAddr, "get cron rule image", models.IsSearch, models.ErrQueryData)
		resComm(c, models.ErrSplitParma, nil)
		return
	}
	image := alertmgr.GetChartManual(id.ID)
	models.OO.RecodeLog(user.Username, c.Request.RemoteAddr, "get cron rule image", models.IsSearch, models.Success)
	resComm(c, models.Success, image)
}

func postSendTestMail(c *gin.Context) {
	user := c.Keys["userInfo"].(*models.UserSessionInfo)
	pass := models.CheckPriv(user, "crontab", "", "mail")
	if pass != models.Success {
		resComm(c, pass, nil)
		return
	}
	var bf *models.BriefMessage
	if !alertmgr.SendTestMail() {
		bf = models.ErrSendMail
	} else {
		bf = models.Success
	}
	models.OO.RecodeLog(user.Username, c.Request.RemoteAddr, "send test mail", models.IsPublish, models.Success)
	resComm(c, bf, nil)
}

func getCron(c *gin.Context) {
	user := c.Keys["userInfo"].(*models.UserSessionInfo)
	pass := models.CheckPriv(user, "baseConfig", "cronapi", "search")
	if pass != models.Success {
		resComm(c, pass, nil)
		return
	}
	sp := &models.SplitPage{}
	if err := c.BindQuery(sp); err != nil {
		models.OO.RecodeLog(user.Username, c.Request.RemoteAddr, "get cron api", models.IsSearch, models.ErrSplitParma)
		resComm(c, models.ErrSplitParma, nil)
		return
	}
	data, bf := models.GetCronApi(sp)
	models.OO.RecodeLog(user.Username, c.Request.RemoteAddr, "get cron api", models.IsSearch, bf)
	resComm(c, bf, data)
}

func getAllCron(c *gin.Context) {
	user := c.Keys["userInfo"].(*models.UserSessionInfo)
	pass := models.CheckPriv(user, "baseConfig", "cronapi", "search")
	if pass != models.Success {
		resComm(c, pass, nil)
		return
	}
	data, bf := models.GetAllCronApi()
	models.OO.RecodeLog(user.Username, c.Request.RemoteAddr, "get cron api", models.IsSearch, bf)
	resComm(c, bf, data)
}

func postCron(c *gin.Context) {
	user := c.Keys["userInfo"].(*models.UserSessionInfo)
	pass := models.CheckPriv(user, "baseConfig", "cronapi", "add")
	if pass != models.Success {
		resComm(c, pass, nil)
		return
	}
	newApi := models.CronApiPost{}
	if err := c.BindJSON(&newApi); err != nil {
		models.OO.RecodeLog(user.Username, c.Request.RemoteAddr, "create cron api", models.IsAdd, models.ErrPostData)
		resComm(c, models.ErrPostData, nil)
		return
	}
	bf := models.PostCronApi(user, &newApi)
	models.OO.RecodeLog(user.Username, c.Request.RemoteAddr, "create cron api", models.IsAdd, bf)
	resComm(c, bf, nil)
}

func putCron(c *gin.Context) {
	user := c.Keys["userInfo"].(*models.UserSessionInfo)
	pass := models.CheckPriv(user, "baseConfig", "cronapi", "update")
	if pass != models.Success {
		resComm(c, pass, nil)
		return
	}
	label := models.CronApiPost{}
	if err := c.BindJSON(&label); err != nil {
		models.OO.RecodeLog(user.Username, c.Request.RemoteAddr, "update cron api", models.IsUpdate, models.ErrPostData)
		resComm(c, models.ErrPostData, nil)
		return
	}
	bf := models.PutCronApi(user, &label)
	models.OO.RecodeLog(user.Username, c.Request.RemoteAddr, "update cron api", models.IsUpdate, bf)
	resComm(c, bf, nil)
}

func putCronStatus(c *gin.Context) {
	user := c.Keys["userInfo"].(*models.UserSessionInfo)
	pass := models.CheckPriv(user, "baseConfig", "cronapi", "dis.enable")
	if pass != models.Success {
		resComm(c, pass, nil)
		return
	}
	ed := models.EnabledInfo{}
	if err := c.BindJSON(&ed); err != nil {
		models.OO.RecodeLog(user.Username, c.Request.RemoteAddr, "update cron api status", models.IsUpdate, models.ErrPostData)
		resComm(c, models.ErrPostData, nil)
		return
	}
	bf := models.PutCronApiStatus(user, &ed)
	models.OO.RecodeLog(user.Username, c.Request.RemoteAddr, "update cron api status", models.IsUpdate, bf)
	resComm(c, bf, nil)
}

func delCron(c *gin.Context) {
	user := c.Keys["userInfo"].(*models.UserSessionInfo)
	pass := models.CheckPriv(user, "baseConfig", "cronapi", "delete")
	if pass != models.Success {
		resComm(c, pass, nil)
		return
	}
	idStr, ok := c.GetQuery("id")
	if !ok {
		models.OO.RecodeLog(user.Username, c.Request.RemoteAddr, "delete cron api", models.IsDel, models.ErrQueryData)
		resComm(c, models.ErrQueryData, nil)
		return
	}
	id, err := strconv.ParseInt(idStr, 10, 0)
	if err != nil {
		config.Log.Error(err)
		models.OO.RecodeLog(user.Username, c.Request.RemoteAddr, "delete cron api", models.IsDel, models.ErrQueryData)
		resComm(c, models.ErrQueryData, nil)
		return
	}
	bf := models.DelCronApi(id)
	models.OO.RecodeLog(user.Username, c.Request.RemoteAddr, "delete cron api", models.IsDel, bf)
	resComm(c, bf, nil)
}

func getProTmpl(c *gin.Context) {
	user := c.Keys["userInfo"].(*models.UserSessionInfo)
	pass := models.CheckPriv(user, "baseConfig", "editPrometheusYml", "load_tmpl")
	if pass != models.Success {
		resComm(c, pass, nil)
		return
	}
	tmpl, bf := models.GetProTmpl()
	models.OO.RecodeLog(user.Username, c.Request.RemoteAddr, "get prometheus template", models.IsSearch, bf)
	resComm(c, bf, tmpl)
}

func putProTmpl(c *gin.Context) {
	user := c.Keys["userInfo"].(*models.UserSessionInfo)
	pass := models.CheckPriv(user, "baseConfig", "editPrometheusYml", "save_tmpl")
	if pass != models.Success {
		resComm(c, pass, nil)
		return
	}
	tp := models.ProTmpl{}
	if err := c.BindJSON(&tp); err != nil {
		models.OO.RecodeLog(user.Username, c.Request.RemoteAddr, "update prometheus template", models.IsUpdate, models.ErrPostData)
		resComm(c, models.ErrPostData, nil)
		return
	}
	bf := models.PutProTmpl(user, &tp)
	models.OO.RecodeLog(user.Username, c.Request.RemoteAddr, "update prometheus template", models.IsUpdate, bf)
	resComm(c, bf, nil)
}

func getStruct(c *gin.Context) {
	user := c.Keys["userInfo"].(*models.UserSessionInfo)
	data, bf := models.GetStruct()
	models.OO.RecodeLog(user.Username, c.Request.RemoteAddr, "get struct", models.IsSearch, bf)
	resComm(c, bf, data)
}

func ws(c *gin.Context) {
	user := c.Keys["userInfo"].(*models.UserSessionInfo)
	models.OO.RecodeLog(user.Username, c.Request.RemoteAddr, "create websokcet", models.IsAdd, models.Success)
	bf := models.WS(c)
	resComm(c, bf, nil)
}

func getOptions(c *gin.Context) {
	user := c.Keys["userInfo"].(*models.UserSessionInfo)
	pass := models.CheckPriv(user, "baseConfig", "options", "search_options")
	if pass != models.Success {
		resComm(c, pass, nil)
		return
	}
	data, bf := models.GetOptions()
	models.OO.RecodeLog(user.Username, c.Request.RemoteAddr, "get options", models.IsSearch, bf)
	resComm(c, bf, data)
}

func putOptions(c *gin.Context) {
	user := c.Keys["userInfo"].(*models.UserSessionInfo)
	pass := models.CheckPriv(user, "baseConfig", "options", "update_options")
	if pass != models.Success {
		resComm(c, pass, nil)
		return
	}
	opts := map[string]string{}
	if err := c.BindJSON(&opts); err != nil {
		models.OO.RecodeLog(user.Username, c.Request.RemoteAddr, "update options", models.IsUpdate, models.ErrPostData)
		resComm(c, models.ErrPostData, nil)
		return
	}
	bf := models.PutOptions(opts)
	models.OO.RecodeLog(user.Username, c.Request.RemoteAddr, "update options", models.IsUpdate, bf)
	resComm(c, bf, nil)
}

func getOperateLog(c *gin.Context) {
	sp := &models.SplitPage{}
	if err := c.BindQuery(sp); err != nil {
		resComm(c, models.ErrSplitParma, nil)
		return
	}
	data, bf := models.GetOperateLog(sp)
	// models.OO.RecodeLog(user.Username, c.Request.RemoteAddr, "update sub group", models.IsUpdate, bf)
	resComm(c, bf, data)
}

func preOptResetSystem(c *gin.Context) {
	user := c.Keys["userInfo"].(*models.UserSessionInfo)
	pass := models.CheckPriv(user, "baseConfig", "empty", "reset")
	if pass != models.Success {
		resComm(c, pass, nil)
		return
	}
	bf := models.PreOptResetSystem()
	models.OO.RecodeLog(user.Username, c.Request.RemoteAddr, "create reset prometheus config data key", models.IsReset, bf)
	resComm(c, bf, nil)
}

func optResetSystem(c *gin.Context) {
	user := c.Keys["userInfo"].(*models.UserSessionInfo)
	pass := models.CheckPriv(user, "baseConfig", "empty", "reset")
	if pass != models.Success {
		resComm(c, pass, nil)
		return
	}
	if sync.AS.CanNoDo(sync.ReSetAllData) {
		resComm(c, models.ErrAlreadyRunning, nil)
		return
	}
	code := models.ResetCode{}
	if err := c.BindJSON(&code); err != nil {
		models.OO.RecodeLog(user.Username, c.Request.RemoteAddr, "reset prometheus config data", models.IsReset, models.ErrPostData)
		resComm(c, models.ErrPostData, nil)
		sync.AS.Done(sync.ReSetAllData)
		return
	}
	bf := models.OptResetSystem(user, &code)
	models.OO.FlagLog(user.Username, c.Request.RemoteAddr, "reset prometheus config data", models.IsReset, bf)
	sync.AS.Done(sync.ReSetAllData)
	resComm(c, bf, nil)
}

func ctlCreate(c *gin.Context) {
	user := c.Keys["userInfo"].(*models.UserSessionInfo)
	pass := models.CheckPriv(user, "control", "", "create")
	if pass != models.Success {
		resComm(c, pass, nil)
		return
	}
	if sync.AS.CanNoDo(sync.ReCreateAllPrometheusConfig) {
		resComm(c, models.ErrAlreadyRunning, nil)
		return
	}
	bf := models.AllowOneObj.DoPublishJobs(user, false)
	models.OO.RecodeLog(user.Username, c.Request.RemoteAddr, "create prometheus config", models.IsPublish, bf)
	sync.AS.Done(sync.ReCreateAllPrometheusConfig)
	resComm(c, bf, nil)
}

func ctlReload(c *gin.Context) {
	user := c.Keys["userInfo"].(*models.UserSessionInfo)
	pass := models.CheckPriv(user, "control", "", "reload")
	if pass != models.Success {
		resComm(c, pass, nil)
		return
	}
	if sync.AS.CanNoDo(sync.ReloadAllPrometheusConfig) {
		resComm(c, models.ErrAlreadyRunning, nil)
		return
	}
	bf := models.Reload(config.Cfg.PrometheusCfg.Api)
	models.OO.RecodeLog(user.Username, c.Request.RemoteAddr, "reload prometheus config", models.IsPublish, bf)
	sync.AS.Done(sync.ReloadAllPrometheusConfig)
	resComm(c, bf, nil)
}

func ctlCreateAReload(c *gin.Context) {
	user := c.Keys["userInfo"].(*models.UserSessionInfo)
	pass := models.CheckPriv(user, "control", "", "createAndreload")
	if pass != models.Success {
		resComm(c, pass, nil)
		return
	}
	if sync.AS.CanNoDo(sync.ReCreateAndReloadPrometheusConfig) {
		resComm(c, models.ErrAlreadyRunning, nil)
		return
	}
	bf := models.AllowOneObj.DoPublishJobs(user, true)
	models.OO.RecodeLog(user.Username, c.Request.RemoteAddr, "create prometheus config", models.IsPublish, bf)
	if bf != models.Success {
		sync.AS.Done(sync.ReCreateAndReloadPrometheusConfig)
		resComm(c, bf, nil)
		return
	}
	bf = models.Reload(config.Cfg.PrometheusCfg.Api)
	models.OO.RecodeLog(user.Username, c.Request.RemoteAddr, "reload prometheus config", models.IsPublish, bf)
	sync.AS.Done(sync.ReCreateAndReloadPrometheusConfig)
	resComm(c, bf, nil)
}

func getPrometheusUrl(c *gin.Context) {
	user := c.Keys["userInfo"].(*models.UserSessionInfo)
	pass := models.CheckPriv(user, "control", "", "get_prometheus_url")
	if pass != models.Success {
		resComm(c, pass, nil)
		return
	}
	u, bf := models.GetPrometheusUrl()
	models.OO.RecodeLog(user.Username, c.Request.RemoteAddr, "get prometheus url", models.IsPublish, bf)
	resComm(c, bf, u)
}

func getIDC(c *gin.Context) {
	user := c.Keys["userInfo"].(*models.UserSessionInfo)
	pass := models.CheckPriv(user, "idc", "", "get_idc")
	if pass != models.Success {
		resComm(c, pass, nil)
		return
	}
	id := &models.OnlyID{}
	if err := c.BindQuery(id); err != nil {
		models.OO.RecodeLog(user.Username, c.Request.RemoteAddr, "get idc", models.IsSearch, models.ErrQueryData)
		resComm(c, models.ErrQueryData, nil)
		return
	}
	idc, bf := models.GetIDC(id)
	models.OO.RecodeLog(user.Username, c.Request.RemoteAddr, "get idc", models.IsSearch, bf)
	resComm(c, bf, idc)
}

func getIDCs(c *gin.Context) {
	user := c.Keys["userInfo"].(*models.UserSessionInfo)
	pass := models.CheckPriv(user, "idc", "", "get_idc")
	if pass != models.Success {
		resComm(c, pass, nil)
		return
	}
	models.OO.RecodeLog(user.Username, c.Request.RemoteAddr, "get idcs", models.IsSearch, models.Success)
}

func postIDC(c *gin.Context) {
	user := c.Keys["userInfo"].(*models.UserSessionInfo)
	pass := models.CheckPriv(user, "idc", "", "add_idc")
	if pass != models.Success {
		resComm(c, pass, nil)
		return
	}
	newIDC := models.NewIDC{}
	if err := c.BindJSON(&newIDC); err != nil {
		models.OO.RecodeLog(user.Username, c.Request.RemoteAddr, "create new idc", models.IsAdd, models.ErrPostData)
		resComm(c, models.ErrPostData, nil)
		return
	}
	bf := models.PostIDC(user, &newIDC)
	models.OO.RecodeLog(user.Username, c.Request.RemoteAddr, "create new idc", models.IsAdd, bf)
	resComm(c, bf, nil)
}

func putIDC(c *gin.Context) {
	user := c.Keys["userInfo"].(*models.UserSessionInfo)
	pass := models.CheckPriv(user, "idc", "", "update_idc")
	if pass != models.Success {
		resComm(c, pass, nil)
		return
	}
	newIDC := models.NewIDC{}
	if err := c.BindJSON(&newIDC); err != nil {
		models.OO.RecodeLog(user.Username, c.Request.RemoteAddr, "update idc", models.IsUpdate, models.ErrUpdateData)
		resComm(c, models.ErrPostData, nil)
		return
	}
	bf := models.PutIDC(user, &newIDC)
	models.OO.RecodeLog(user.Username, c.Request.RemoteAddr, "update idc", models.IsUpdate, bf)
	resComm(c, bf, nil)
}

func putIDCRemark(c *gin.Context) {
	user := c.Keys["userInfo"].(*models.UserSessionInfo)
	pass := models.CheckPriv(user, "idc", "", "update_idc")
	if pass != models.Success {
		resComm(c, pass, nil)
		return
	}
	idc := models.IDC{}
	if err := c.BindJSON(&idc); err != nil {
		models.OO.RecodeLog(user.Username, c.Request.RemoteAddr, "update idc remark", models.IsUpdate, models.ErrUpdateData)
		resComm(c, models.ErrPostData, nil)
		return
	}
	bf := models.PutIDCRemark(user, &idc)
	models.OO.RecodeLog(user.Username, c.Request.RemoteAddr, "update idc remark", models.IsUpdate, bf)
	resComm(c, bf, nil)
}

func delIDC(c *gin.Context) {
	user := c.Keys["userInfo"].(*models.UserSessionInfo)
	pass := models.CheckPriv(user, "idc", "", "del_idc")
	if pass != models.Success {
		resComm(c, pass, nil)
		return
	}
	id := &models.OnlyID{}
	if err := c.BindQuery(id); err != nil {
		models.OO.RecodeLog(user.Username, c.Request.RemoteAddr, "delete idc", models.IsDel, models.ErrDelData)
		resComm(c, models.ErrQueryData, nil)
		return
	}
	bf := models.DelIDC(id)
	models.OO.RecodeLog(user.Username, c.Request.RemoteAddr, "update idc", models.IsDel, bf)
	resComm(c, bf, nil)
}

func getIDCTree(c *gin.Context) {
	user := c.Keys["userInfo"].(*models.UserSessionInfo)
	pass := models.CheckPriv(user, "idc", "", "get_idc_tree")
	if pass != models.Success {
		resComm(c, pass, nil)
		return
	}
	search := models.SearchContent2{}
	if err := c.BindQuery(&search); err != nil {
		models.OO.RecodeLog(user.Username, c.Request.RemoteAddr, "get idc tree", models.IsSearch, models.ErrQueryData)
		resComm(c, models.ErrQueryData, nil)
		return
	}
	treeData, bf := models.GetIDCTree(&search)
	models.OO.RecodeLog(user.Username, c.Request.RemoteAddr, "get idc tree", models.IsSearch, bf)
	resComm(c, bf, treeData)
}

func getIDCTreeXls(c *gin.Context) {
	user := c.Keys["userInfo"].(*models.UserSessionInfo)
	pass := models.CheckPriv(user, "idc", "", "get_idc_tree_xls")
	if pass != models.Success {
		resComm(c, pass, nil)
		return
	}
	xlsData, bf := models.GetIDCXls()
	models.OO.RecodeLog(user.Username, c.Request.RemoteAddr, "get idc tree xls", models.IsSearch, bf)
	resComm(c, bf, xlsData)
}

func getLine(c *gin.Context) {
	user := c.Keys["userInfo"].(*models.UserSessionInfo)
	pass := models.CheckPriv(user, "idc", "", "get_line")
	if pass != models.Success {
		resComm(c, pass, nil)
		return
	}
	id := &models.OnlyID{}
	if err := c.BindQuery(id); err != nil {
		models.OO.RecodeLog(user.Username, c.Request.RemoteAddr, "get line", models.IsSearch, models.ErrQueryData)
		resComm(c, models.ErrQueryData, nil)
		return
	}
	idc, bf := models.GetLine(id)
	models.OO.RecodeLog(user.Username, c.Request.RemoteAddr, "get line", models.IsSearch, bf)
	resComm(c, bf, idc)
}

func getLines(c *gin.Context) {
	user := c.Keys["userInfo"].(*models.UserSessionInfo)
	pass := models.CheckPriv(user, "idc", "", "get_line")
	if pass != models.Success {
		resComm(c, pass, nil)
		return
	}
	models.OO.RecodeLog(user.Username, c.Request.RemoteAddr, "get lines", models.IsSearch, models.Success)
}

func postLine(c *gin.Context) {
	user := c.Keys["userInfo"].(*models.UserSessionInfo)
	pass := models.CheckPriv(user, "idc", "", "add_line")
	if pass != models.Success {
		resComm(c, pass, nil)
		return
	}
	newLine := models.NewLine{}
	if err := c.BindJSON(&newLine); err != nil {
		models.OO.RecodeLog(user.Username, c.Request.RemoteAddr, "create new line", models.IsAdd, models.ErrPostData)
		resComm(c, models.ErrPostData, nil)
		return
	}
	bf := models.PostLine(user, &newLine)
	models.OO.RecodeLog(user.Username, c.Request.RemoteAddr, "create new line", models.IsAdd, bf)
	resComm(c, bf, nil)
}

func putLine(c *gin.Context) {
	user := c.Keys["userInfo"].(*models.UserSessionInfo)
	pass := models.CheckPriv(user, "idc", "", "update_line")
	if pass != models.Success {
		resComm(c, pass, nil)
		return
	}
	newLine := models.NewLine{}
	if err := c.BindJSON(&newLine); err != nil {
		models.OO.RecodeLog(user.Username, c.Request.RemoteAddr, "update line", models.IsUpdate, models.ErrPostData)
		resComm(c, models.ErrPostData, nil)
		return
	}
	bf := models.PutLine(user, &newLine)
	models.OO.RecodeLog(user.Username, c.Request.RemoteAddr, "update line", models.IsUpdate, bf)
	resComm(c, bf, nil)
}

func delLine(c *gin.Context) {
	user := c.Keys["userInfo"].(*models.UserSessionInfo)
	pass := models.CheckPriv(user, "idc", "", "del_line")
	if pass != models.Success {
		resComm(c, pass, nil)
		return
	}
	delReqParams := &models.IdAndRm{}
	if err := c.BindQuery(delReqParams); err != nil {
		models.OO.RecodeLog(user.Username, c.Request.RemoteAddr, "delete line", models.IsDel, models.ErrQueryData)
		resComm(c, models.ErrQueryData, nil)
		return
	}
	bf := models.DelLine(delReqParams)
	models.OO.RecodeLog(user.Username, c.Request.RemoteAddr, "delete line", models.IsDel, bf)
	resComm(c, bf, nil)
}

func getLineIpAddrs(c *gin.Context) {
	user := c.Keys["userInfo"].(*models.UserSessionInfo)
	pass := models.CheckPriv(user, "idc", "", "get_line_ipaddr")
	if pass != models.Success {
		resComm(c, pass, nil)
		return
	}
	id := &models.OnlyID{}
	if err := c.BindQuery(id); err != nil {
		models.OO.RecodeLog(user.Username, c.Request.RemoteAddr, "get ip address of line", models.IsSearch, models.ErrQueryData)
		resComm(c, models.ErrQueryData, nil)
		return
	}
	data, bf := models.GetLineIpAddrs(id)
	models.OO.RecodeLog(user.Username, c.Request.RemoteAddr, "get ip address of line", models.IsSearch, bf)
	resComm(c, bf, data)
}

func putLineIpAddrs(c *gin.Context) {
	user := c.Keys["userInfo"].(*models.UserSessionInfo)
	pass := models.CheckPriv(user, "idc", "", "update_line_ipaddr")
	if pass != models.Success {
		resComm(c, pass, nil)
		return
	}
	newPool := models.NewPool{}
	if err := c.BindJSON(&newPool); err != nil {
		models.OO.RecodeLog(user.Username, c.Request.RemoteAddr, "update line ip pool", models.IsUpdate, models.ErrPostData)
		resComm(c, models.ErrPostData, nil)
		return
	}
	bf := models.PutLineIpAddrs(user, &newPool)
	models.OO.RecodeLog(user.Username, c.Request.RemoteAddr, "update line ip pool", models.IsUpdate, bf)
	resComm(c, bf, nil)
}

func updateNetInfoAll(c *gin.Context) {
	user := c.Keys["userInfo"].(*models.UserSessionInfo)
	pass := models.CheckPriv(user, "idc", "", "update_label_all")
	if pass != models.Success {
		resComm(c, pass, nil)
		return
	}
	if sync.AS.CanNoDo(sync.UpdateAllIPIDCAndLineInfo) {
		resComm(c, models.ErrAlreadyRunning, nil)
		return
	}
	bf := models.UpdateAllIPAddrs(user, false)
	models.OO.RecodeLog(user.Username, c.Request.RemoteAddr, "update position of all ip address", models.IsUpdate, bf)
	sync.AS.Done(sync.UpdateAllIPIDCAndLineInfo)
	resComm(c, bf, nil)
}

func updateNetInfoPart(c *gin.Context) {
	user := c.Keys["userInfo"].(*models.UserSessionInfo)
	pass := models.CheckPriv(user, "idc", "", "update_label_part")
	if pass != models.Success {
		resComm(c, pass, nil)
		return
	}
	if sync.AS.CanNoDo(sync.UpdatePartIPIDCAndLineInfo) {
		resComm(c, models.ErrAlreadyRunning, nil)
		return
	}
	bf := models.UpdateAllIPAddrs(user, true)
	models.OO.RecodeLog(user.Username, c.Request.RemoteAddr, "update position of part ip address", models.IsUpdate, bf)
	sync.AS.Done(sync.UpdatePartIPIDCAndLineInfo)
	resComm(c, bf, nil)
}

func createLabelForAllIPs(c *gin.Context) {
	user := c.Keys["userInfo"].(*models.UserSessionInfo)
	pass := models.CheckPriv(user, "idc", "", "create_label_for_all_job")
	if pass != models.Success {
		resComm(c, pass, nil)
		return
	}
	if sync.AS.CanNoDo(sync.UpdateAllIPLabelsInJobGroup) {
		resComm(c, models.ErrAlreadyRunning, nil)
		return
	}
	bf := models.CreateLabelForAllIPs(user, nil)
	models.OO.RecodeLog(user.Username, c.Request.RemoteAddr, "create label and sub group of job", models.IsUpdate, bf)
	sync.AS.Done(sync.UpdateAllIPLabelsInJobGroup)
	resComm(c, bf, nil)
}

func idcExpand(c *gin.Context) {
	user := c.Keys["userInfo"].(*models.UserSessionInfo)
	pass := models.CheckPriv(user, "idc", "", "expand_ipaddr")
	if pass != models.Success {
		resComm(c, pass, nil)
		return
	}
	if sync.AS.CanNoDo(sync.UpdateAllIPLabelsInJobGroup) {
		resComm(c, models.ErrAlreadyRunning, nil)
		return
	}
	edr := models.ExpandDataReq{}
	if err := c.BindJSON(&edr); err != nil {
		config.Log.Error(err)
		models.OO.RecodeLog(user.Username, c.Request.RemoteAddr, "expand ipaddr", models.IsUpdate, models.ErrPostData)
		resComm(c, models.ErrPostData, nil)
		sync.AS.Done(sync.UpdateAllIPLabelsInJobGroup)
		return
	}
	bf := models.IDCExpand(user, &edr)
	models.OO.RecodeLog(user.Username, c.Request.RemoteAddr, "expand ipaddr", models.IsUpdate, bf)
	sync.AS.Done(sync.UpdateAllIPLabelsInJobGroup)
	resComm(c, bf, nil)
}

func idcUpdateLineExpandSwitch(c *gin.Context) {
	user := c.Keys["userInfo"].(*models.UserSessionInfo)
	pass := models.CheckPriv(user, "idc", "", "update_line")
	if pass != models.Success {
		resComm(c, pass, nil)
		return
	}
	newLine := models.ExpandSwitchReq{}
	if err := c.BindJSON(&newLine); err != nil {
		models.OO.RecodeLog(user.Username, c.Request.RemoteAddr, "update line switch - expand", models.IsUpdate, models.ErrPostData)
		resComm(c, models.ErrPostData, nil)
		return
	}
	bf := models.PutLineExpandSwitch(user, &newLine)
	models.OO.RecodeLog(user.Username, c.Request.RemoteAddr, "update line switch - expand", models.IsUpdate, bf)
	resComm(c, bf, nil)
}

func idcUpdateLineViewSwitch(c *gin.Context) {
	user := c.Keys["userInfo"].(*models.UserSessionInfo)
	pass := models.CheckPriv(user, "idc", "", "update_line")
	if pass != models.Success {
		resComm(c, pass, nil)
		return
	}
	newLine := models.ViewSwitchReq{}
	if err := c.BindJSON(&newLine); err != nil {
		models.OO.RecodeLog(user.Username, c.Request.RemoteAddr, "update line switch - view", models.IsUpdate, models.ErrPostData)
		resComm(c, models.ErrPostData, nil)
		return
	}
	bf := models.PutLineViewSwitch(user, &newLine)
	models.OO.RecodeLog(user.Username, c.Request.RemoteAddr, "update line switch - view", models.IsUpdate, bf)
	resComm(c, bf, nil)
}

func idcUpdateExpandSwitch(c *gin.Context) {
	user := c.Keys["userInfo"].(*models.UserSessionInfo)
	pass := models.CheckPriv(user, "idc", "", "update_idc")
	if pass != models.Success {
		resComm(c, pass, nil)
		return
	}
	switchData := models.ExpandSwitchReq{}
	if err := c.BindJSON(&switchData); err != nil {
		models.OO.RecodeLog(user.Username, c.Request.RemoteAddr, "update idc switch - expand", models.IsUpdate, models.ErrUpdateData)
		resComm(c, models.ErrPostData, nil)
		return
	}
	bf := models.PutIDCExpandSwitch(user, &switchData)
	models.OO.RecodeLog(user.Username, c.Request.RemoteAddr, "update idc switch - expand", models.IsUpdate, bf)
	resComm(c, bf, nil)
}

func idcUpdateViewSwitch(c *gin.Context) {
	user := c.Keys["userInfo"].(*models.UserSessionInfo)
	pass := models.CheckPriv(user, "idc", "", "update_idc")
	if pass != models.Success {
		resComm(c, pass, nil)
		return
	}
	switchData := models.ViewSwitchReq{}
	if err := c.BindJSON(&switchData); err != nil {
		models.OO.RecodeLog(user.Username, c.Request.RemoteAddr, "update idc switch - view", models.IsUpdate, models.ErrUpdateData)
		resComm(c, models.ErrPostData, nil)
		return
	}
	bf := models.PutIDCViewSwitch(user, &switchData)
	models.OO.RecodeLog(user.Username, c.Request.RemoteAddr, "update idc switch - view", models.IsUpdate, bf)
	resComm(c, bf, nil)
}

func putIDCNoBell(c *gin.Context) {
	user := c.Keys["userInfo"].(*models.UserSessionInfo)
	pass := models.CheckPriv(user, "idc", "", "bell_switch")
	if pass != models.Success {
		resComm(c, pass, nil)
		return
	}
	info := &models.OnlyID{}
	if err := c.BindJSON(info); err != nil {
		models.OO.RecodeLog(user.Username, c.Request.RemoteAddr, "idc no bell", models.IsUpdate, models.ErrPostData)
		resComm(c, models.ErrPostData, nil)
		return
	}
	bf := models.PutIDCNoBell(info)
	models.OO.RecodeLog(user.Username, c.Request.RemoteAddr, "idc no bell", models.IsUpdate, bf)
	resComm(c, bf, nil)
}

func putLineNoBell(c *gin.Context) {
	user := c.Keys["userInfo"].(*models.UserSessionInfo)
	pass := models.CheckPriv(user, "idc", "", "bell_switch")
	if pass != models.Success {
		resComm(c, pass, nil)
		return
	}
	info := &models.OnlyID{}
	if err := c.BindJSON(info); err != nil {
		models.OO.RecodeLog(user.Username, c.Request.RemoteAddr, "line no bell", models.IsUpdate, models.ErrPostData)
		resComm(c, models.ErrPostData, nil)
		return
	}
	bf := models.PutLineNoBell(info)
	models.OO.RecodeLog(user.Username, c.Request.RemoteAddr, "line no bell", models.IsUpdate, bf)
	resComm(c, bf, nil)
}
