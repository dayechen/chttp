package socket

import (
	"cweb/pkg/socket/logic"
	"cweb/pkg/socket/server"

	"github.com/gin-gonic/gin"
)

// 通过gin的路由创建socket连接
func NewSocket(path string, g *gin.Engine) *logic.Engine {
	callback := registerCallback()
	result := server.New(callback)
	g.GET(path, func(c *gin.Context) {
		result.WebSocketHandle(c.Writer, c.Request)
	})
	result.CMD("test", func(req *logic.Request) {
		req.EchoSuccess("成功")
	})
	// registerEvent(result.On)
	return result
}
