package router

import (
	"cweb/global"
	"cweb/http/middleware"
	v1 "cweb/http/route/api/v1"
	"cweb/socket"

	"github.com/gin-gonic/gin"
)

// NewRouter 创建路由
func NewRouter() *gin.Engine {
	r := gin.New()
	user := v1.NewUser()
	r.Use(middleware.Cors())
	r.POST("/verify", user.Register)
	// 注册socket路由
	global.Socket = socket.NewSocket("/ws", r)
	apiv1 := r.Group("/api/v1")
	apiv1.Use(middleware.JWT())
	{
		apiv1.GET("list", user.List)
	}
	return r
}
