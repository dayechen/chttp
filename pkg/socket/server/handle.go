package server

import (
	"cweb/pkg/socket/logic"
)

// 用户传递来的一些设置
var callback *Callback

// 创建监听的事件
func on(event string, callback func(user *logic.Request)) {
	logic.Broadcaster.On(event, callback)
}

// New 初始化
func New(c *Callback) *logic.Engine {
	callback = c
	go logic.Broadcaster.Start()
	return &logic.Engine{
		On:                on,
		SendMsgById:       logic.Broadcaster.SendMsgById,
		SendMsgByFilter:   logic.Broadcaster.SendMsgByFilter,
		CloseConnById:     logic.Broadcaster.CloseConnById,
		CloseConnByFilter: logic.Broadcaster.CloseConnByFilter,
		WebSocketHandle:   webSocketHandle,
	}
}

// 启动的登录
type Callback struct {
	Verification func(token string) *logic.LoginData                 // 用户登录时调用 返回用户ID 游客返回0
	Leaving      func(user *logic.User)                              // 用户离开时调用 传递用户id
	RepeatLogin  func(oldUser *logic.User, newUser *logic.User) bool // 重复登录时调用 返回true就让登录的用户被挤下来
}
