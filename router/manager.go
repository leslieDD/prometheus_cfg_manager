package router

import (
	"pro_cfg_manager/config"
	"pro_cfg_manager/models"
	"strconv"

	"github.com/gin-gonic/gin"
)

func initManagerApi() {

	v1.POST("/logout", logout)

	v1.GET("/manager/groups", getManagerGroups)
	v1.GET("/manager/groups/list", getManagerGroupList)
	v1.GET("/manager/groups/enabled", getMGEnabled)
	v1.POST("/manager/group", postManagerGroup)
	v1.PUT("/manager/group", putManagerGroup)
	v1.DELETE("/manager/group", deleteManagerGroup)
	v1.PUT("/manager/group/status", putManagerGroupStatus)

	v1.GET("/manager/users", getManagerUsers)
	v1.POST("/manager/user", postManagerUser)
	v1.PUT("/manager/user", putManagerUser)
	v1.DELETE("/manager/user", deleteManagerUser)
	v1.PUT("/manager/user/status", putManagerUserStatus)
	v1.POST("/manager/user/password", postUserPassword)

	v1.GET("/manager/users/list", getManagerUsersList)
	v1.GET("/manager/group/member", getManagerGroupMember)
	v1.PUT("/manager/group/member", putManagerGroupMember)

	v1.GET("/txt/programer/say", getProgramerSay)

	v1.GET("/manager/user/priv", getUserPriv)
	v1.PUT("/manager/user/priv", putUserPriv)

	v1.GET("/manager/user/menu/priv", getUserMenuPriv)

	v1.GET("/manager/setting", getManagerSetting)
	v1.PUT("/manager/setting", putManagerSetting)

	v1.POST("/manager/reset/admin", postManagerResetAdmin)
	v1.POST("/manager/reset/secret", postManagerResetSecret)
}

func logout(c *gin.Context) {
	uid := models.OnlyID{}
	if err := c.BindQuery(&uid); err != nil {
		resComm(c, models.ErrPostData, nil)
		return
	}
	user := c.Keys["userInfo"].(*models.UserSessionInfo)
	bf := models.Logout(user)
	resComm(c, bf, nil)
}

func getManagerGroups(c *gin.Context) {
	user := c.Keys["userInfo"].(*models.UserSessionInfo)
	pass := models.CheckPriv(user, "admin", "group", "search")
	if pass != models.Success {
		resComm(c, pass, nil)
		return
	}
	sp := &models.SplitPage{}
	if err := c.BindQuery(sp); err != nil {
		resComm(c, models.ErrSplitParma, nil)
		return
	}
	result, bf := models.GetManagerGroups(sp)
	resComm(c, bf, result)
}

func getManagerGroupList(c *gin.Context) {
	user := c.Keys["userInfo"].(*models.UserSessionInfo)
	pass := models.CheckPriv(user, "admin", "group", "search")
	if pass != models.Success {
		resComm(c, pass, nil)
		return
	}
	result, bf := models.GetManagerGroupList(user)
	resComm(c, bf, result)
}

func getMGEnabled(c *gin.Context) {
	user := c.Keys["userInfo"].(*models.UserSessionInfo)
	pass := models.CheckPriv(user, "admin", "group", "search")
	if pass != models.Success {
		resComm(c, pass, nil)
		return
	}
	data, bf := models.GetManagerGroupEnabled()
	resComm(c, bf, data)
}

func postManagerGroup(c *gin.Context) {
	user := c.Keys["userInfo"].(*models.UserSessionInfo)
	pass := models.CheckPriv(user, "admin", "group", "add")
	if pass != models.Success {
		resComm(c, pass, nil)
		return
	}
	mg := models.ManagerGroup{}
	if err := c.BindJSON(&mg); err != nil {
		resComm(c, models.ErrPostData, nil)
		return
	}
	bf := models.PostManagerGroup(user, &mg)
	resComm(c, bf, nil)
}

func putManagerGroup(c *gin.Context) {
	user := c.Keys["userInfo"].(*models.UserSessionInfo)
	pass := models.CheckPriv(user, "admin", "group", "update")
	if pass != models.Success {
		resComm(c, pass, nil)
		return
	}
	mg := models.ManagerGroup{}
	if err := c.BindJSON(&mg); err != nil {
		resComm(c, models.ErrPostData, nil)
		return
	}
	bf := models.PutManagerGroup(user, &mg)
	resComm(c, bf, nil)
}

func deleteManagerGroup(c *gin.Context) {
	user := c.Keys["userInfo"].(*models.UserSessionInfo)
	pass := models.CheckPriv(user, "admin", "group", "delete")
	if pass != models.Success {
		resComm(c, pass, nil)
		return
	}
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
	user := c.Keys["userInfo"].(*models.UserSessionInfo)
	pass := models.CheckPriv(user, "admin", "group", "dis.enable")
	if pass != models.Success {
		resComm(c, pass, nil)
		return
	}
	enabledInfo := models.EnabledInfo{}
	if err := c.BindJSON(&enabledInfo); err != nil {
		resComm(c, models.ErrPostData, nil)
		return
	}
	bf := models.PutManagerGroupStatus(user, &enabledInfo)
	resComm(c, bf, nil)
}

///

func getManagerUsers(c *gin.Context) {
	user := c.Keys["userInfo"].(*models.UserSessionInfo)
	pass := models.CheckPriv(user, "admin", "user", "search")
	if pass != models.Success {
		resComm(c, pass, nil)
		return
	}
	sp := &models.SplitPage{}
	if err := c.BindQuery(sp); err != nil {
		resComm(c, models.ErrSplitParma, nil)
		return
	}
	result, bf := models.GetManagerUsers(sp)
	resComm(c, bf, result)
}

func postManagerUser(c *gin.Context) {
	user := c.Keys["userInfo"].(*models.UserSessionInfo)
	pass := models.CheckPriv(user, "admin", "user", "add")
	if pass != models.Success {
		resComm(c, pass, nil)
		return
	}
	mu := models.ManagerUser{}
	if err := c.BindJSON(&mu); err != nil {
		config.Log.Error(err)
		resComm(c, models.ErrPostData, nil)
		return
	}
	bf := models.PostManagerUser(user, &mu)
	resComm(c, bf, nil)
}

func putManagerUser(c *gin.Context) {
	user := c.Keys["userInfo"].(*models.UserSessionInfo)
	pass := models.CheckPriv(user, "admin", "user", "update")
	if pass != models.Success {
		resComm(c, pass, nil)
		return
	}
	mu := models.ManagerUser{}
	if err := c.BindJSON(&mu); err != nil {
		config.Log.Error(err)
		resComm(c, models.ErrPostData, nil)
		return
	}
	bf := models.PutManagerUser(user, &mu)
	resComm(c, bf, nil)
}

func deleteManagerUser(c *gin.Context) {
	user := c.Keys["userInfo"].(*models.UserSessionInfo)
	pass := models.CheckPriv(user, "admin", "user", "delete")
	if pass != models.Success {
		resComm(c, pass, nil)
		return
	}
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
	bf := models.DeleteManagerUser(gID)
	resComm(c, bf, nil)
}

func putManagerUserStatus(c *gin.Context) {
	user := c.Keys["userInfo"].(*models.UserSessionInfo)
	pass := models.CheckPriv(user, "admin", "user", "dis.enable")
	if pass != models.Success {
		resComm(c, pass, nil)
		return
	}
	enabledInfo := models.EnabledInfo{}
	if err := c.BindJSON(&enabledInfo); err != nil {
		resComm(c, models.ErrPostData, nil)
		return
	}
	bf := models.PutManagerUserStatus(user, &enabledInfo)
	resComm(c, bf, nil)
}

func postUserPassword(c *gin.Context) {
	user := c.Keys["userInfo"].(*models.UserSessionInfo)
	pass := models.CheckPriv(user, "person", "", "update_password")
	if pass != models.Success {
		resComm(c, pass, nil)
		return
	}
	pInfo := models.ChangePasswordInfo{}
	if err := c.BindJSON(&pInfo); err != nil {
		resComm(c, models.ErrPostData, nil)
		return
	}
	bf := models.PostUserPassword(user, &pInfo)
	resComm(c, bf, nil)
}

func getManagerUsersList(c *gin.Context) {
	user := c.Keys["userInfo"].(*models.UserSessionInfo)
	pass := models.CheckPriv(user, "admin", "user", "search")
	if pass != models.Success {
		resComm(c, pass, nil)
		return
	}
	data, bf := models.GetManagerUsersList()
	resComm(c, bf, data)
}

func getManagerGroupMember(c *gin.Context) {
	user := c.Keys["userInfo"].(*models.UserSessionInfo)
	pass := models.CheckPriv(user, "admin", "group", "search")
	if pass != models.Success {
		resComm(c, pass, nil)
		return
	}
	gInfo := models.GetPrivInfo{}
	if err := c.BindQuery(&gInfo); err != nil {
		resComm(c, models.ErrPostData, nil)
		return
	}
	data, bf := models.GetManagerGroupMember(&gInfo)
	resComm(c, bf, data)
}

func putManagerGroupMember(c *gin.Context) {
	user := c.Keys["userInfo"].(*models.UserSessionInfo)
	pass := models.CheckPriv(user, "admin", "group", "update")
	if pass != models.Success {
		resComm(c, pass, nil)
		return
	}
	gInfo := models.GetPrivInfo{}
	if err := c.BindQuery(&gInfo); err != nil {
		resComm(c, models.ErrPostData, nil)
		return
	}
	uIDs := []int{}
	if err := c.BindJSON(&uIDs); err != nil {
		resComm(c, models.ErrPostData, nil)
		return
	}
	bf := models.PutManagerGroupMember(user, &gInfo, uIDs)
	resComm(c, bf, nil)
}

func getProgramerSay(c *gin.Context) {
	data, bf := models.GetProgramerSay()
	resComm(c, bf, data)
}

func getUserPriv(c *gin.Context) {
	user := c.Keys["userInfo"].(*models.UserSessionInfo)
	pass := models.CheckPriv(user, "admin", "privileges", "search")
	if pass != models.Success {
		resComm(c, pass, nil)
		return
	}
	// u := c.Keys["userInfo"].(*models.ManagerUserDetail)
	gInfo := models.GetPrivInfo{}
	if err := c.BindQuery(&gInfo); err != nil {
		resComm(c, models.ErrPostData, nil)
		return
	}
	data, bf := models.GetGroupPriv(&gInfo)
	resComm(c, bf, data)
}

func putUserPriv(c *gin.Context) {
	user := c.Keys["userInfo"].(*models.UserSessionInfo)
	pass := models.CheckPriv(user, "admin", "privileges", "update")
	if pass != models.Success {
		resComm(c, pass, nil)
		return
	}
	privInfo := []*models.ItemPriv{}
	if err := c.BindJSON(&privInfo); err != nil {
		resComm(c, models.ErrPostData, nil)
		return
	}
	gInfo := models.GetPrivInfo{}
	if err := c.BindQuery(&gInfo); err != nil {
		resComm(c, models.ErrPostData, nil)
		return
	}
	bf := models.PutGroupPriv(user, privInfo, &gInfo)
	resComm(c, bf, nil)
}

func getUserMenuPriv(c *gin.Context) {
	user := c.Keys["userInfo"].(*models.UserSessionInfo)
	result, bf := models.GetUserMenuPriv(user)
	resComm(c, bf, result)
}

func getManagerSetting(c *gin.Context) {
	user := c.Keys["userInfo"].(*models.UserSessionInfo)
	pass := models.CheckPriv(user, "admin", "setting", "search")
	if pass != models.Success {
		resComm(c, pass, nil)
		return
	}
	data, bf := models.GetManagerSetting()
	resComm(c, bf, data)
}

func putManagerSetting(c *gin.Context) {
	user := c.Keys["userInfo"].(*models.UserSessionInfo)
	pass := models.CheckPriv(user, "admin", "setting", "update")
	if pass != models.Success {
		resComm(c, pass, nil)
		return
	}
	params := map[string]string{}
	if err := c.BindJSON(&params); err != nil {
		resComm(c, models.ErrPostData, nil)
		return
	}
	bf := models.PutManagerSetting(params)
	resComm(c, bf, nil)
}

func postManagerResetAdmin(c *gin.Context) {
	user := c.Keys["userInfo"].(*models.UserSessionInfo)
	pass := models.CheckPriv(user, "admin", "setting", "reset")
	if pass != models.Success {
		resComm(c, pass, nil)
		return
	}
	code := models.ResetCode{}
	if err := c.BindJSON(&code); err != nil {
		resComm(c, models.ErrPostData, nil)
		return
	}
	bf := models.OptResetAdmin(&code, c.Request.RemoteAddr)
	resComm(c, bf, nil)
}

func postManagerResetSecret(c *gin.Context) {
	user := c.Keys["userInfo"].(*models.UserSessionInfo)
	pass := models.CheckPriv(user, "admin", "setting", "reset")
	if pass != models.Success {
		resComm(c, pass, nil)
		return
	}
	bf := models.PreOptResetAdmin()
	resComm(c, bf, nil)
}
