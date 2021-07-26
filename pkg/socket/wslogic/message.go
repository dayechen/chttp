package wslogic

import "time"

// SendWelcomeMessage 发送欢迎消息
func WelcomeMessage() *Message {
	return &Message{
		Code:    0,
		MsgTime: time.Now().Unix(),
		Content: "欢迎",
		Event:   "welcome",
	}
}

func SuccessMessage(content interface{}, messageID int) *Message {
	return &Message{
		Code:      0,
		MessageID: messageID,
		MsgTime:   time.Now().Unix(),
		Content:   content,
		Event:     "echo",
	}
}

func ErrorMessage(content interface{}, messageID int) *Message {
	return &Message{
		Code:      1,
		MessageID: messageID,
		MsgTime:   time.Now().Unix(),
		Content:   content,
		Event:     "echo",
	}
}

func NormalMessage(event string, content interface{}) *Message {
	return &Message{
		Code:    0,
		MsgTime: time.Now().Unix(),
		Content: content,
		Event:   event,
	}
}
