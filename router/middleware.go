package router

import (
	"pro_cfg_manager/config"
	"pro_cfg_manager/models"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

// cros
func crosMw() {
	// 	router := gin.Default()
	// - No origin allowed by default
	// - GET,POST, PUT, HEAD methods
	// - Credentials share disabled
	// - Preflight requests cached for 12 hours
	// config := cors.DefaultConfig()
	// config.AllowOrigins = []string{"http://google.com"}
	// config.AllowOrigins == []string{"http://google.com", "http://facebook.com"}

	// router.Use(cors.New(config))

	GinDefault.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD"},
		AllowHeaders:     []string{"Origin", "Authentication", "Content-Length", "Content-Type"},
		ExposeHeaders:    []string{"Content-Length", "Content-Type", "Authentication"},
		AllowCredentials: true,
		AllowOriginFunc: func(origin string) bool {
			return origin == "http://localhost:9527"
		},
		MaxAge: 12 * time.Hour,
	}))
}

func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("Authorization")
		if token == "" {
			token = c.GetHeader("Sec-WebSocket-Protocol")
			config.Log.Errorf("ws token is %s", token)
			if token == "" {
				c.Abort()
				resComm(c, models.ErrTokenIsNull, nil)
				return
			}
		}
		bf := models.CheckUserSession(c, token)
		if bf != models.Success {
			c.Abort()
			resComm(c, models.ErrTokenIsNull, nil)
			return
		}
		// ss := models.SSObj.Get(token)
		// if ss == nil {
		// 	c.Abort()
		// 	resComm(c, models.ErrTokenIsNull, nil)
		// 	return
		// }
		// c.Keys = models.Authentication{
		// 	"userInfo": ss,
		// }
	}
}
