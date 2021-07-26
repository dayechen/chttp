package socket

import (
	"cweb/global"
	"cweb/pkg/socket/server"
	"cweb/pkg/socket/wslogic"

	"github.com/gin-gonic/gin"
)

// 通过gin的路由创建socket连接
func NewSocket(path string, g *gin.Engine) *wslogic.Engine {
	s := server.New(registerCallback())
	s.Event("test", func(req *wslogic.Request) {
		global.Log.Debugf("收到websocket消息：%+v", req.Params)
		req.EchoSuccess("成功")
	})
	return s
}
