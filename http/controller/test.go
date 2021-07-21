package controller

import (
	"cweb/pkg/app"

	"github.com/gin-gonic/gin"
)

func WsTest(c *gin.Context) {
	response := app.NewResponse(c)
	response.ToSuccess("成功")
}
