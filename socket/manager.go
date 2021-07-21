package socket

import (
	"cweb/pkg/socket/logic"
	"cweb/pkg/socket/server"

	"github.com/gin-gonic/gin"
)

// 通过gin的路由创建socket连接
func NewSocket(path string, g *gin.Engine) *logic.Engine {
	s := server.New(registerCallback())
	s.Cmd("test", func(req *logic.Request) {
		req.EchoSuccess("成功")
	})
	return s
}
