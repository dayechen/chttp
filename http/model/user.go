package model

import (
	"cweb/http/dbmodel"

	"github.com/jinzhu/gorm"
)

type UserModel struct {
	db *gorm.DB
}

func NewUser(db *gorm.DB) *UserModel {
	return &UserModel{
		db: db,
	}
}

func (u *UserModel) GetUidByPhone(phone string) (uint, error) {
	user := dbmodel.User{}
	return user.ID, u.db.Select("id").Where("phone = ?", phone).Find(&user).Error
}
