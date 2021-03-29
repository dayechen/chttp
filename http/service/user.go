package service

import (
	"crypto/md5"
	"cweb/global"
	"cweb/http/model"
	"encoding/hex"
)

// CreateUserRequest 创建用户请求
type CreateUserRequest struct {
	Phone    string `form:"phone" binding:"required" validate:"required,isPhone" lable:"手机"`
	Passowrd string `form:"password" binding:"required" validate:"required" lable:"密码"`
	Nickname string `form:"nickname" binding:"required" validate:"required" lable:"昵称"`
	Avatar   string `form:"avatar" binding:"required"`
}

// CreateUser 创建用户 返回创建成功后的用户id
func (svc *Service) CreateUser(param *CreateUserRequest) (int, error) {
	user := model.User{
		Nickname: param.Nickname,
		Phone:    param.Phone,
		Password: encryption(param.Passowrd),
		Avatar:   param.Avatar,
	}
	return user.Create(global.DB)
}

// 加盐加密
func encryption(pas string) string {
	salt := global.JWTSetting.Salt
	m5 := md5.New()
	m5.Write([]byte(pas))
	m5.Write([]byte(string(salt)))
	st := m5.Sum(nil)
	return hex.EncodeToString(st)
}
