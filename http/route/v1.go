package route

import (
	"cweb/http/controller/user"

	"github.com/gin-gonic/gin"
)

func setupV1(r *gin.RouterGroup) {
	r.GET("/user", user.Create)
	r.GET("/test", func(c *gin.Context) {

	})
}
