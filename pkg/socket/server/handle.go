package server

import (
	"cweb/pkg/socket/logic"
)

// 用户传递来的一些设置
var config *Config

// 创建监听的事件
func on(event string, callback func(user *logic.Request)) {
	logic.Broadcaster.On(event, callback)
}

// Build 向外暴露方法
func Build(c *Config) *logic.Engine {
	config = c
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
type Config struct {
	Verification func(token string) (int, bool) // 用户登录时调用 返回用户ID 游客返回0
	Leaving      func(uid int)                  // 用户离开时调用 传递用户id
	RepeatLogin  bool                           // 重复登录时调用
}
