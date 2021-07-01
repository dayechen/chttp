package request

type User struct {
	Phone string `form:"phone" validate:"required,phone" lable:"手机号"`
}

type Verification struct {
	Phone string `form:"phone" validate:"required,phone" lable:"手机号"`
	Code  int    `form:"code" validate:"required" lable:"验证码"`
}
