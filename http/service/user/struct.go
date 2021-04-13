package user

type UserRequest struct {
	Phone string `form:"phone" validate:"required,phone" lable:"手机号"`
}
