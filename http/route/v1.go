package route

import (
	"cweb/http/controller"

	"github.com/gin-gonic/gin"
)

func setupV1(r *gin.RouterGroup) {
	r.GET("/user/verification", controller.Verification)
}
