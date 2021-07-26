package server

import "cweb/pkg/socket/wslogic"

// 用户传递来的一些设置
var callback *wslogic.Callback

// New 初始化
func New(c *wslogic.Callback) *wslogic.Engine {
	callback = c
	go wslogic.Broadcaster.Start()
	return &wslogic.Engine{
		WebSocketHandle: webSocketHandle,
	}
}
