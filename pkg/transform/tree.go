package transform

import (
	"github.com/gin-gonic/gin"
)

type Tree struct {
	Pid      uint // 上级菜单的id
	ID       uint
	Children []Tree
	Result   map[string]interface{} // 最后生成的结果
}

func buildTreeChildren(current *[]Tree, pid uint) ([]Tree, []gin.H) {
	result := []Tree{}
	resultMap := []gin.H{}
	for _, v := range *current {
		if v.Pid != pid {
			continue
		}
		tree, _ := buildTreeChildren(current, v.ID)
		treeResultMap := []gin.H{}
		for _, v1 := range tree {
			treeResultMap = append(treeResultMap, v1.Result)
		}
		v.Children = tree
		v.Result["children"] = treeResultMap
		resultMap = append(resultMap, v.Result)
		result = append(result, v)
	}
	return result, resultMap
}

// 构建一棵树
func BuildTree(current *[]Tree) []gin.H {
	result := []gin.H{}
	notRootTree := []Tree{} // 不是根节点
	rootTree := []Tree{}    // 根节点
	for _, v := range *current {
		if v.Pid != 0 {
			notRootTree = append(notRootTree, v)
		} else {
			rootTree = append(rootTree, v)
		}
	}
	for _, v := range rootTree {
		_, treeMap := buildTreeChildren(&notRootTree, v.ID)
		v.Result["children"] = treeMap
		result = append(result, v.Result)
	}
	return result
}
