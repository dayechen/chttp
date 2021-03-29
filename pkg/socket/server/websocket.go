package server

import (
	"cweb/pkg/socket/logic"
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
	UID, ok := config.Verification(token[0])
	if !ok {
		conn.Close(websocket.StatusInternalError, "验证失败")
		return
	}

	user := logic.NewUser(conn, UID)
	// 加入用户数组统一管理
	logic.Broadcaster.UserEntering(user)
	// 打开给用户发送消息的通道
	go user.OpenMessageChannel(req.Context())
	// 发送欢迎消息
	user.MessageChannel <- logic.WelcomeMessage()
	// 接收用户消息 死循环阻塞线程 发生错误或用户退出才会往下执行
	err = user.ReceiveMessage(req.Context())
	if err == nil {
		// 用户正常的离开了
		conn.Close(websocket.StatusNormalClosure, "")
	} else {
		conn.Close(websocket.StatusInternalError, err.Error())
	}
	// 调用用户离开后的回调
	config.Leaving(UID)
	// 回收用户资源
	logic.Broadcaster.UserLeaving(user)
}
