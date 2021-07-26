package socket

import (
	"cweb/pkg/socket/wslogic"
	"strconv"
)

// 注册socket的回调函数
func registerCallback() *wslogic.Callback {
	return &wslogic.Callback{
		Verification: func(token string) (int, bool) {
			id, _ := strconv.Atoi(token)
			return id, true
		},
		Leaving: func(user *wslogic.User) {
		},
		RepeatLogin: func(oldUser, newUser *wslogic.User) bool {
			return true
		},
	}
}
