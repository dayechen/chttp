package service

import (
	"cweb/global"
	"cweb/http/model"
)

func GetUserId(phone string) (uint, error) {
	model := model.NewUser(global.DB)
	return model.GetUidByPhone(phone)
}
