package middleware

import (
	"gobase/pkg/app"
	errcode "gobase/pkg/error"

	"github.com/gin-gonic/gin"
)

// JWT 登录验证
func JWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Request.Header.Get("Authorization")
		if token == "" {
			app.NewResponse(c).ToErrorResponse(errcode.NewError(1, "用户未登录"))
			c.Abort()
			return
		}
		claims, err := app.ParseToken(token[7:])
		if err != nil {
			app.NewResponse(c).ToErrorResponse(errcode.NewError(1, "登录已过期"))
			c.Abort()
			return
		}
		c.Set("uid", claims.UID)
		c.Next()
		return
	}
}
