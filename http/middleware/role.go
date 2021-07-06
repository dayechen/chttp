package middleware

import (
	"github.com/gin-gonic/gin"
)

func Role() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()
	}
}
