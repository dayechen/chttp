package router

import (
	"gobase/http/middleware"
	v1 "gobase/http/route/api/v1"

	"github.com/gin-gonic/gin"
)

// NewRouter 创建路由
func NewRouter() *gin.Engine {
	r := gin.New()
	user := v1.NewUser()
	r.Use(middleware.Cors())

	r.POST("/verify", user.Register)
	apiv1 := r.Group("/api/v1")
	apiv1.Use(middleware.JWT())
	{
		apiv1.GET("list", user.List)
	}
	return r
}
