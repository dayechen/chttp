package controller

import (
	"cweb/http/model"
	"cweb/pkg/app"

	"github.com/gin-gonic/gin"
)

func Menu(c *gin.Context) {
	response := app.NewResponse(c)
	menus := model.GetMenu()
	response.ToSuccess(gin.H{
		"menus": menus,
	})
}
