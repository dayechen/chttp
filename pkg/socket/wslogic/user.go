package wslogic

import (
	"context"
	"errors"

	"nhooyr.io/websocket"
	"nhooyr.io/websocket/wsjson"
)

// NewUser 返回用户实例
func NewUser(conn *websocket.Conn, ID int) *User {
	return &User{
		conn:           conn,
		ID:             ID,
		close:          false,
		MessageChannel: make(chan *Message),
	}
}

// OpenMessageChannel 打开消息通道
func (u *User) OpenMessageChannel(ctx context.Context) {
	for msg := range u.MessageChannel {
		wsjson.Write(ctx, u.conn, msg)
	}
}

// CloseMessageChannel 关闭消息通道
func (u *User) CloseMessageChannel() {
	close(u.MessageChannel)
}

type UserMessage struct {
	Event     string      `json:"event"`
	MessageID int         `json:"messageID"`
	Params    interface{} `json:"params"`
}

// ReceiveMessage 接收用户消息
func (u *User) ReceiveMessage(ctx context.Context) error {
	var (
		err         error
		UserMessage UserMessage
	)
	for {
		if u.close {
			return nil
		}
		err = wsjson.Read(ctx, u.conn, &UserMessage)
		if err != nil {
			var closeErr websocket.CloseError
			if errors.As(err, &closeErr) {
				return nil
			}
			return err
		}
		// 发射事件
		Broadcaster.emit(UserMessage.Event, u, UserMessage.MessageID, UserMessage.Params)
	}
}
