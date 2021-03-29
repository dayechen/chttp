package socket

import (
	"cweb/pkg/socket/logic"
	"cweb/pkg/socket/server"
	"strconv"

	"github.com/gin-gonic/gin"
)

func NewSocket(path string, g *gin.Engine) *logic.Engine {
	config := &server.Config{
		Verification: func(token string) (int, bool) {
			id, _ := strconv.Atoi(token)
			return id, true
		},
		Leaving: func(uid int) {
			println("用户离开了")
			// system.ClearUser(uid)
		},
	}
	result := server.Build(config)
	g.GET(path, func(c *gin.Context) {
		result.WebSocketHandle(c.Writer, c.Request)
	})
	registerEvent(result.On)
	return result
}

// 注册各种自定义事件
func registerEvent(on logic.EngineOnFunc) {
	on("test", func(req *logic.Request) {
		req.EchoSuccess("成功")
	})
}
