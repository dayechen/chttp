package controller

import (
	"cweb/http/service"
	"cweb/pkg/app"

	"github.com/gin-gonic/gin"
)

func GetMenu(c *gin.Context) {
	response := app.NewResponse(c)
	menus := service.GetMenu()
	response.ToSuccess(gin.H{
		"menus": menus,
	})
}
