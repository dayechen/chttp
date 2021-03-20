package app

import (
	errcode "cweb/pkg/error"
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
func (r *Response) ToResponse(data interface{}) {
	if data == nil {
		data = gin.H{}
	}
	r.Ctx.JSON(http.StatusOK, data)
}

// ToErrorResponse 返回错误的响应
func (r *Response) ToErrorResponse(err *errcode.Error) {
	response := gin.H{"code": err.Code(), "msg": err.Msg()}
	r.Ctx.JSON(http.StatusOK, response)
}
