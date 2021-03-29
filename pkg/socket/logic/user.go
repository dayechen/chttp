package logic

import (
	"context"
	"errors"

	"nhooyr.io/websocket"
	"nhooyr.io/websocket/wsjson"
)

// NewUser 返回用户实例
func NewUser(conn *websocket.Conn, UID int) *User {
	return &User{
		conn:           conn,
		UID:            UID,
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

type UserMsg struct {
	Event     string                 `json:"event"`
	MessageID int                    `json:"messageID"`
	Params    map[string]interface{} `json:"params"`
}

// ReceiveMessage 接收用户消息
func (u *User) ReceiveMessage(ctx context.Context) error {
	var (
		err     error
		userMsg UserMsg
	)
	for {
		if u.close {
			return nil
		}
		err = wsjson.Read(ctx, u.conn, &userMsg)
		if err != nil {
			var closeErr websocket.CloseError
			if errors.As(err, &closeErr) {
				return nil
			}
			return err
		}
		// 发射事件
		Broadcaster.Emit(userMsg.Event, u, userMsg.MessageID, userMsg.Params)
	}
}
