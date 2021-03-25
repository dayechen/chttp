package app

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Response 响应
type Response struct {
	Ctx *gin.Context
}

// NewResponse 新的响应
func NewResponse(ctx *gin.Context) *Response {
	return &Response{
		Ctx: ctx,
	}
}

// ToResponse 返回响应
func (r *Response) ToSuccess(data interface{}) {
	if data == nil {
		data = gin.H{}
	}
	result := gin.H{
		"code":    0,
		"message": "success",
		"data":    data,
	}
	r.Ctx.JSON(http.StatusOK, result)
}

// ToError 返回错误的响应
func (r *Response) ToError(data interface{}) {
	result := gin.H{
		"code":    1,
		"message": "error",
		"data":    data,
	}
	r.Ctx.JSON(http.StatusOK, result)
}
