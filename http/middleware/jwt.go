package middleware

import (
	"cweb/pkg/app"
	"net/http"

	"github.com/gin-gonic/gin"
)

// JWT 登录验证
func JWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Request.Header.Get("Authorization")
		if token == "" {
			app.NewResponse(c).ToErrorCode(http.StatusUnauthorized)
			c.Abort()
			return
		}
		claims, err := app.ParseToken(token[7:])
		if err != nil {
			app.NewResponse(c).ToErrorCode(http.StatusUnauthorized)
			c.Abort()
			return
		}
		c.Set("uid", claims.UID)
		c.Next()
	}
}
