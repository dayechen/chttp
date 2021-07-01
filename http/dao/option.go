package dao

import (
	"cweb/global"
	"cweb/http/type/model"
)

func GetMenuList() *[]model.Menu {
	menus := []model.Menu{}
	global.DB.Select("label,url,id,pid").Find(&menus)
	return &menus
}
