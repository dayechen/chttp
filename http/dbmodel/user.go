package dbmodel

import "github.com/jinzhu/gorm"

type User struct {
	gorm.Model
	Nickname string
	Phone    string
}
