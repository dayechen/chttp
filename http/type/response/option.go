package response

type Menu struct {
	Label    string
	Url      string
	Pid      uint
	ID       uint
	Children []Menu
}
