package wslogic

// Broadcaster 广播器用于给所有用户发送消息
var Broadcaster = &broadcaster{
	users:         make(map[int]*User), // 所有在线的用户
	tourists:      make(map[int]*User), // 游客用户
	events:        make(map[string]func(req *Request)),
	nextTouristID: 0, // 下一个游客用户的id

	enteringChannel: make(chan *User),
	leavingChannel:  make(chan *User),
	messageChannel:  make(chan *Message),
}

// 监听事件
func (b *broadcaster) on(event string, callback func(req *Request)) {
	b.events[event] = callback
}

// 发射事件
func (b *broadcaster) emit(event string, user *User, messageID int, params interface{}) {
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

// 通过id获取登录的用户
func (b *broadcaster) GetUserByID(UID int) *User {
	return b.users[UID]
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
			if user.ID == 0 {
				b.nextTouristID++
				user.TouristID = b.nextTouristID
				b.tourists[user.TouristID] = user
			} else {
				b.users[user.ID] = user
			}
		case user := <-b.leavingChannel:
			if user.ID == 0 {
				delete(b.tourists, user.TouristID)
			} else {
				delete(b.users, user.ID)
			}
			user.CloseMessageChannel()
		}
	}
}
