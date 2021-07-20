package socket

import (
	"cweb/pkg/socket/logic"
	"cweb/pkg/socket/server"
	"strconv"
)

// 注册socket的回调函数
func registerCallback() *server.Callback {
	return &server.Callback{
		Verification: func(token string) *logic.LoginData {
			id, _ := strconv.Atoi(token)
			return &logic.LoginData{
				ID: id,
				Ok: true,
			}
		},
		Leaving: func(user *logic.User) {
		},
		RepeatLogin: func(oldUser, newUser *logic.User) bool {
			return true
		},
	}
}
