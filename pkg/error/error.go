package errcode

// Error 错误结构体
type Error struct {
	code int
	msg  string
}

// NewError 返回新的错误
func NewError(code int, msg string) *Error {
	return &Error{
		code: code,
		msg:  msg,
	}
}

// Code 返回错误码
func (e Error) Code() int {
	return e.code
}

// Msg 返回错误信息
func (e Error) Msg() string {
	return e.msg
}
