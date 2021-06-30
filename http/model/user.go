package model

import (
	"cweb/global"
	"cweb/http/dbmodel"
)

func GetUidByPhone(phone string) (uint, error) {
	user := dbmodel.User{}
	return user.ID, global.DB.Select("id").Where("phone = ?", phone).Find(&user).Error
}
