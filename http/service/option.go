package service

import (
	"cweb/http/dao"
	"cweb/pkg/transform"

	"github.com/gin-gonic/gin"
)

func GetMenu() []gin.H {
	initMenuList := dao.GetMenuList()
	treeList := []transform.Tree{}
	for _, v := range *initMenuList {
		treeList = append(treeList, transform.Tree{
			Pid: v.Pid,
			ID:  v.ID,
			Result: gin.H{
				"id":    v.ID,
				"label": v.Label,
				"url":   v.Url,
			},
		})
	}
	return transform.BuildTree(&treeList)
}
