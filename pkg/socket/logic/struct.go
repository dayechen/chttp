package logic

import (
	"net/http"

	"nhooyr.io/websocket"
)

// User 用户结构体
type User struct {
	UID            int           `json:"uid"` // 用户id
	TouristID      int           // 游客id
	MessageChannel chan *Message `json:"-"` // 消息通道
	close          bool          // 关闭连接

	conn *websocket.Conn
}

// Message 给用户发送的消息
type Message struct {
	// User    *User  `json:"user"`    // 谁发的消息
	Code      int         `json:"code"`      // 消息类型
	MessageID int         `json:"messageID"` // 消息的唯一标识
	Content   interface{} `json:"content"`   // 消息内容
	MsgTime   int64       `json:"msgTime"`   // 消息发送的时间
	Event     string      `json:"event"`     // 事件名称
}

// 广播器
type broadcaster struct {
	users    map[int]*User                  // 当前登录的所有用户
	tourists map[int]*User                  // 游客用户
	events   map[string]func(user *Request) // 控制器定义的事件

	nextTouristID   int           // 下一个用户的id
	enteringChannel chan *User    // 用户进入通道
	leavingChannel  chan *User    // 用户离开通道
	messageChannel  chan *Message // 通讯通道
}

// Response 收到的用户消息
type Request struct {
	User        *User
	EchoSuccess func(content interface{})
	EchoError   func(content interface{})
	Params      map[string]interface{}
}

type SendMsgByIdFunc func(id int, event string, msg interface{}) error
type EngineOnFunc func(event string, callback func(user *Request))
type SendMsgByFilterFunc func(event string, callback func(user *User) interface{})
type WebSocketHandleFunc func(w http.ResponseWriter, req *http.Request)
type CloseConnByIdFunc func(id int)
type CloseConnByFilterFunc func(callback func(user *User) bool)

type Engine struct {
	On                EngineOnFunc
	SendMsgById       SendMsgByIdFunc
	SendMsgByFilter   SendMsgByFilterFunc // 遍历回调函数 如果返回的不是nil就发送数据
	WebSocketHandle   WebSocketHandleFunc
	CloseConnById     CloseConnByIdFunc
	CloseConnByFilter CloseConnByFilterFunc
}
