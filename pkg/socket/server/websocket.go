package server

import (
	"cweb/global"
	"cweb/pkg/socket/wslogic"
	"fmt"
	"net/http"

	"nhooyr.io/websocket"
)

// WebSocketHandleFunc 处理websocket请求
func webSocketHandle(w http.ResponseWriter, req *http.Request) {
	conn, err := websocket.Accept(w, req, &websocket.AcceptOptions{InsecureSkipVerify: true})
	if err != nil {
		fmt.Printf("websocket连接错误 %v \n", err)
		return
	}
	query := req.URL.Query()
	token, ok := query["token"]
	if !ok {
		conn.Close(websocket.StatusInternalError, "未携带token")
		return
	}
	id, ok := callback.Verification(token[0])
	if !ok {
		conn.Close(websocket.StatusInternalError, "验证失败")
		return
	}
	user := wslogic.NewUser(conn, id)
	loginUser := wslogic.Broadcaster.GetUserByID(user.ID)
	// 匿名用户直接登录或没有登录的用户直接登录
	if user.ID == 0 || loginUser == nil {
		wslogic.Broadcaster.UserEntering(user)
	} else {
		ok := callback.RepeatLogin(loginUser, user)
		if ok {
			global.Socket.CloseConnByID(loginUser.ID)
			wslogic.Broadcaster.UserEntering(user) // 新用户上线
		} else {
			// 直接关闭连接
			conn.Close(websocket.StatusInternalError, "请勿重复登录")
			return
		}
	}
	// 打开给用户发送消息的通道
	go user.OpenMessageChannel(req.Context())
	// 发送欢迎消息
	user.MessageChannel <- wslogic.WelcomeMessage()
	// 接收用户消息 死循环阻塞线程 发生错误或用户退出才会往下执行
	err = user.ReceiveMessage(req.Context())
	if err == nil {
		// 用户正常的离开了
		conn.Close(websocket.StatusNormalClosure, "")
	} else {
		conn.Close(websocket.StatusInternalError, err.Error())
	}
	// 调用用户离开后的回调
	callback.Leaving(user)
	// 回收用户资源
	wslogic.Broadcaster.UserLeaving(user)
}
