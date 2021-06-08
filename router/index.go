package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func initCommonRouter() {
	GinDefault.GET("/", index)
}

func index(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", gin.H{})
}
