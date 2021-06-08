package router

import (
	"net/http"
	"pro_cfg_manager/models"

	"github.com/gin-gonic/gin"
)

type response struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

// 成功的回复
func resOK(c *gin.Context, data interface{}) {
	r := response{
		Code: models.Success.Code,
		Msg:  models.Success.Message,
		Data: data,
	}
	c.JSON(http.StatusOK, r)
}

// 一般的回复
func resComm(c *gin.Context, bf *models.BriefMessage, data interface{}) {
	r := response{
		Code: bf.Code,
		Msg:  bf.Message,
		Data: data,
	}
	if bf.Code == 401000 {
		c.JSON(http.StatusUnauthorized, r)
	} else {
		c.JSON(http.StatusOK, r)
	}
}
