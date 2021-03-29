package logic

import (
	"errors"
	"strconv"
)

// Broadcaster 广播器用于给所有用户发送消息
var Broadcaster = &broadcaster{
	users:         make(map[int]*User), // 所有在线的用户
	tourists:      make(map[int]*User), // 游客用户
	events:        make(map[string]func(res *Request)),
	nextTouristID: 0, // 下一个游客用户的id

	enteringChannel: make(chan *User),
	leavingChannel:  make(chan *User),
	messageChannel:  make(chan *Message),
}

// 监听事件
func (b *broadcaster) On(event string, callback func(req *Request)) {
	b.events[event] = callback
}

// 发射事件
func (b *broadcaster) Emit(event string, user *User, messageID int, params map[string]interface{}) {
	if b.events[event] == nil {
		println("收到未定义事件", event)
		return
	}
	b.events[event](&Request{
		User:   user,
		Params: params,
		EchoSuccess: func(content interface{}) {
			user.MessageChannel <- SuccessMessage(content, messageID)
		},
		EchoError: func(content interface{}) {
			user.MessageChannel <- ErrorMessage(content, messageID)
		},
	})
}

// SendMsgById 通过用户id发送消息
func (b *broadcaster) SendMsgById(id int, event string, msg interface{}) error {
	user := b.users[id]
	if user == nil {
		return errors.New("用户" + strconv.Itoa(id) + "不在线")
	}
	user.MessageChannel <- NormalMessage(event, msg)
	return nil
}

// CLoseConnById 通过id关闭连接
func (b *broadcaster) CloseConnById(id int) {
	user := b.users[id]
	if user != nil {
		user.close = true
	}
}

// 根据过滤条件关闭连接
func (b *broadcaster) CloseConnByFilter(callback func(user *User) bool) {
	// 游客和登录用户一起循环
	recipient := [2]map[int]*User{
		b.users, b.tourists,
	}
	for _, v := range recipient {
		for _, v1 := range v {
			ok := callback(v1)
			if ok {
				v1.close = true
			}
		}
	}
}

// SendMsgByFilter 遍历所有当前在线的用户返回的不是nil就发送消息
func (b *broadcaster) SendMsgByFilter(event string, callback func(user *User) interface{}) {
	// 游客和登录用户一起循环
	recipient := [2]map[int]*User{
		b.users, b.tourists,
	}
	for _, v := range recipient {
		for _, v1 := range v {
			msg := callback(v1)
			if msg == nil {
				continue
			}
			v1.MessageChannel <- NormalMessage(event, msg)
		}
	}
}

// UserEntering 用户进入
func (b *broadcaster) UserEntering(u *User) {
	b.enteringChannel <- u
}

// UserLeaving 用户离开
func (b *broadcaster) UserLeaving(u *User) {
	b.leavingChannel <- u
}

// Start 启动广播器
func (b *broadcaster) Start() {
	for {
		select {
		case user := <-b.enteringChannel:
			if user.UID == 0 {
				b.nextTouristID++
				user.TouristID = b.nextTouristID
				b.tourists[user.TouristID] = user
			} else {
				b.users[user.UID] = user
			}
		case user := <-b.leavingChannel:
			if user.UID == 0 {
				delete(b.tourists, user.TouristID)
			} else {
				delete(b.users, user.UID)
			}
			user.CloseMessageChannel()
		}
	}
}
