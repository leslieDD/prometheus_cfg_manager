package router

import (
	"time"

	"github.com/gin-contrib/cors"
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
