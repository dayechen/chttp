package v1

import (
	"cweb/http/service"
	"cweb/pkg/app"

	"github.com/gin-gonic/gin"
)

// User 用户结构体
type User struct{}

// NewUser 创建user
func NewUser() User {
	return User{}
}

// List 用户列表
func (u User) List(c *gin.Context) {
	uid := c.MustGet("uid").(int)
	response := app.NewResponse(c)
	response.ToSuccess(gin.H{
		"msg": uid,
	})
}

// Register 注册
func (u User) Register(c *gin.Context) {
	response := app.NewResponse(c)
	param := service.CreateUserRequest{}
	if err := app.BindAndValid(c, &param); err != nil {
		response.ToError(err.Error())
		return
	}
	svc := service.New(c.Request.Context())
	uid, err := svc.CreateUser(&param)
	if err != nil {
		response.ToError("当前手机号已注册")
		return
	}
	token, err := app.GenerateToken(uid)
	if err != nil {
		response.ToError("错误")
		return
	}
	response.ToSuccess(gin.H{
		"access_token": token,
	})
}
