package socket

import (
	"cweb/pkg/socket/logic"
	"cweb/pkg/socket/server"
	"strconv"

	"github.com/gin-gonic/gin"
)

// 通过gin的路由创建socket连接
func NewSocket(path string, g *gin.Engine) *logic.Engine {
	callback := registerCallback()
	result := server.New(callback)
	g.GET(path, func(c *gin.Context) {
		result.WebSocketHandle(c.Writer, c.Request)
	})
	registerEvent(result.On)
	return result
}

// 注册socket的回调函数
func registerCallback() *server.Callback {
	return &server.Callback{
		Verification: func(token string) *logic.LoginData {
			id, _ := strconv.Atoi(token)
			return &logic.LoginData{
				ID: id,
				Ok: true,
			}
		},
		Leaving: func(user *logic.User) {
		},
		RepeatLogin: func(oldUser, newUser *logic.User) bool {
			return true
		},
	}
}

// 注册各种自定义事件
func registerEvent(on logic.EngineOnFunc) {
	on("test", func(req *logic.Request) {
		req.EchoSuccess("成功")
	})
}
