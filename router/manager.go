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
}

func logout(c *gin.Context) {
	uid := models.OnlyID{}
	if err := c.BindQuery(&uid); err != nil {
		resComm(c, models.ErrPostData, nil)
		return
	}
	bf := models.Logout("")
	resComm(c, bf, nil)
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

func getMGEnabled(c *gin.Context) {
	data, bf := models.GetManagerGroupEnabled()
	resComm(c, bf, data)
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

///

func getManagerUsers(c *gin.Context) {
	sp := &models.SplitPage{}
	if err := c.BindQuery(sp); err != nil {
		resComm(c, models.ErrSplitParma, nil)
		return
	}
	result, bf := models.GetManagerUsers(sp)
	resComm(c, bf, result)
}

func postManagerUser(c *gin.Context) {
	mu := models.ManagerUser{}
	if err := c.BindJSON(&mu); err != nil {
		config.Log.Error(err)
		resComm(c, models.ErrPostData, nil)
		return
	}
	bf := models.PostManagerUser(&mu)
	resComm(c, bf, nil)
}

func putManagerUser(c *gin.Context) {
	mu := models.ManagerUser{}
	if err := c.BindJSON(&mu); err != nil {
		config.Log.Error(err)
		resComm(c, models.ErrPostData, nil)
		return
	}
	bf := models.PutManagerUser(&mu)
	resComm(c, bf, nil)
}

func deleteManagerUser(c *gin.Context) {
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
	enabledInfo := models.EnabledInfo{}
	if err := c.BindJSON(&enabledInfo); err != nil {
		resComm(c, models.ErrPostData, nil)
		return
	}
	bf := models.PutManagerUserStatus(&enabledInfo)
	resComm(c, bf, nil)
}

func postUserPassword(c *gin.Context) {
	pInfo := models.ChangePasswordInfo{}
	if err := c.BindJSON(&pInfo); err != nil {
		resComm(c, models.ErrPostData, nil)
		return
	}
	bf := models.PostUserPassword(&pInfo)
	resComm(c, bf, nil)
}

func getManagerUsersList(c *gin.Context) {
	data, bf := models.GetManagerUsersList()
	resComm(c, bf, data)
}

func getManagerGroupMember(c *gin.Context) {
	gInfo := models.GetPrivInfo{}
	if err := c.BindQuery(&gInfo); err != nil {
		resComm(c, models.ErrPostData, nil)
		return
	}
	data, bf := models.GetManagerGroupMember(&gInfo)
	resComm(c, bf, data)
}

func putManagerGroupMember(c *gin.Context) {
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
	bf := models.PutManagerGroupMember(&gInfo, uIDs)
	resComm(c, bf, nil)
}

func getProgramerSay(c *gin.Context) {
	data, bf := models.GetProgramerSay()
	resComm(c, bf, data)
}

func getUserPriv(c *gin.Context) {
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
	bf := models.PutGroupPriv(privInfo, &gInfo)
	resComm(c, bf, nil)
}
