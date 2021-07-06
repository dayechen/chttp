package model

import "github.com/jinzhu/gorm"

type Role struct {
	gorm.Model
	Url  string
	Info string // --r只读 --w只写 --x可读可写
	Pid  uint   // 上级id 权限是可以继承的
}
