package router

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"pro_cfg_manager/config"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
)

// GinDefault GinDefault
var GinDefault *gin.Engine
var v1 *gin.RouterGroup

func init() {
	gin.SetMode(gin.ReleaseMode)
	GinDefault = gin.Default()
	crosMw()
	GinDefault.LoadHTMLGlob("static/*.html")
	GinDefault.StaticFS("/css", http.Dir("static/css"))
	GinDefault.StaticFS("/js", http.Dir("static/js"))
	GinDefault.StaticFS("/assets", http.Dir("static/assets"))
	GinDefault.StaticFS("/static", http.Dir("static/"))
	GinDefault.StaticFile("/avatar2.jpg", "static/avatar2.jpg")
	GinDefault.StaticFile("/logo.png", "static/logo.png")
	GinDefault.StaticFile("/favicon.ico", "static/favicon.ico")
	// GinDefault.Use(cors.Default())
	v1 = GinDefault.Group("/v1")
	initCommonRouter()
	initApiRouter()
}

// ListenAndServer ListenAndServer
func ListenAndServer() {
	config.Log.Printf("Listen: %s", config.Cfg.App.Listen)
	s := &http.Server{
		Addr:           config.Cfg.App.Listen,
		Handler:        GinDefault,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	go ListenAndServe(s)
	notidy(s)
}

// ListenAndServe ListenAndServe
func ListenAndServe(s *http.Server) {
	config.Log.Printf("%s running: %s", config.Cfg.App.Name, config.Cfg.App.Listen)
	if err := s.ListenAndServe(); err != nil {
		config.Log.Error(err)
		os.Exit(1)
	}
}

func notidy(srv *http.Server) {
	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	config.Log.Println("Shutdown Server ...")
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		config.Log.Fatal("Server Shutdown:", err)
	}
	config.Log.Println("Server exiting")
}
