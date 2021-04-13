package route

import (
	"cweb/global"
	"cweb/socket"

	"github.com/gin-gonic/gin"
)

// NewRouter 创建路由
func NewRouter() *gin.Engine {
	r := gin.New()
	// 注册socket路由
	if global.SocketSetting.Active {
		global.Socket = socket.NewSocket(global.SocketSetting.Url, r)
	}
	setupV1(r.Group("v1"))
	return r
}
