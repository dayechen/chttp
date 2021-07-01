package transform

type Tree struct {
	Pid      uint
	ID       uint
	Children []Tree
	Result   map[string]interface{}
}

func buildTreeChildren(current *[]Tree, pid uint) []Tree {
	result := []Tree{}
	for _, v := range *current {
		if v.Pid != pid {
			continue
		}
		v.Children = buildTreeChildren(current, v.ID)
		result = append(result, v)
	}
	return result
}

// 构建一棵树
func BuildTree(current *[]Tree) []Tree {
	result := []Tree{}
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
		result = append(result, Tree{
			Pid:      v.Pid,
			ID:       v.ID,
			Result:   v.Result,
			Children: buildTreeChildren(&notRootTree, v.ID),
		})
	}
	return result
}
