package dao

import (
	"cweb/global"
	"cweb/http/type/model"
)

func GetUidByPhone(phone string) (uint, error) {
	user := model.User{}
	return user.ID, global.DB.Select("id").Where("phone = ?", phone).Find(&user).Error
}
