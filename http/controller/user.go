package controller

import (
	"cweb/http/dao"
	"cweb/http/type/request"
	"cweb/pkg/app"

	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	response := app.NewResponse(c)
	params := request.Verification{}
	if err := app.BindAndValid(c, &params); err != nil {
		response.ToError(err.Error())
		return
	}
	if params.Code != 1234 {
		response.ToError("验证码错误", 998)
		return
	}
	uid, err := dao.GetUidByPhone(params.Phone)
	if err != nil {
		response.ToError("查无此人")
		return
	}
	token, _ := app.GenerateToken(uid)
	response.ToSuccess(gin.H{
		"token": token,
	})
}
