package route

import (
	"cweb/global"
	"cweb/http/middleware"
	"cweb/socket"

	"github.com/gin-gonic/gin"
)

// NewRouter 创建路由
func NewRouter() *gin.Engine {
	var router *gin.Engine
	if global.ServerSetting.RunMode == "debug" {
		router = gin.Default()
	} else {
		router = gin.New()
	}
	if global.SocketSetting.Active {
		global.Socket = socket.NewSocket(global.SocketSetting.Url, router)
	}
	router.Use(middleware.Cors())
	setupV1(router.Group("v1"))
	return router
}
