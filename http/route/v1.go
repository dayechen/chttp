package route

import (
	"cweb/http/controller"
	"cweb/http/middleware"

	"github.com/gin-gonic/gin"
)

func setupV1(r *gin.RouterGroup) {
	r.GET("/user/verification", controller.Verification)
	r.Use(middleware.JWT())
	r.GET("/option/menu", controller.Menu)
}