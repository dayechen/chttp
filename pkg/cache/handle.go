package cache

type Engine struct {
	roleList *[]RoleType
}

func NewCache() *Engine {
	return &Engine{}
}

func (e *Engine) GetRoleList() *[]RoleType {
	return e.roleList
}

func (e *Engine) UpdateRoleList() {

}
