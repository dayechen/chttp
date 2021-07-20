package model

import "github.com/jinzhu/gorm"

type Role struct {
	gorm.Model
	Permission string // --r只读 --w只写 --x可读可写; /option:r 表示这个链接对这个用户只读
	Button     string // 可见按钮 /option:1 url:按钮id
	Pid        uint   // 上级id 权限是可以继承的
}
