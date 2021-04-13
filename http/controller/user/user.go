package user

import (
	"cweb/http/service/user"
	"cweb/pkg/app"

	"github.com/gin-gonic/gin"
)

func Create(c *gin.Context) {
	response := app.NewResponse(c)
	params := user.UserRequest{}
	if err := app.BindAndValid(c, &params); err != nil {
		response.ToError(err.Error())
		return
	}
	result := user.Create()
	response.ToSuccess(result)
}
