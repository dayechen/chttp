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
		"success": true,
		"code":    0,
		"data":    data,
	}
	r.Ctx.JSON(http.StatusOK, result)
}

// ToError 返回错误的响应
func (r *Response) ToError(message string, code ...int) {
	resultCode := 1
	if code != nil {
		resultCode = code[0]
	}
	result := gin.H{
		"success": false,
		"message": message,
		"code":    resultCode,
		"data":    gin.H{},
	}
	r.Ctx.JSON(http.StatusBadRequest, result)
}

// 返回分页
func (r *Response) ToList() {

}

// ToMessage 返回需要弹出的消息
func (r *Response) ToMessage() {

}
