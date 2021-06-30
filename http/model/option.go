package model

import (
	"cweb/global"
	"cweb/http/dbmodel"
)

func GetMenu() *[]dbmodel.Menu {
	menus := []dbmodel.Menu{}
	global.DB.Select("label,url").Find(&menus)
	return &menus
}
