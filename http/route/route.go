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
	// 注册websocket
	if global.SocketSetting.Active {
		global.Socket = socket.NewSocket(global.SocketSetting.Url, router)
		router.GET(global.SocketSetting.Url, func(c *gin.Context) {
			global.Socket.WebSocketHandle(c.Writer, c.Request)
		})
	}
	// router.Use(middleware.Cors())
	router.Use(middleware.Role())
	setupV1(router.Group("v1"))
	return router
}
