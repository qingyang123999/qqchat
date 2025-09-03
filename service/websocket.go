package service

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"net/http"
	"qqchat/utils"
	"time"
)

type SysWebSocket struct{}

// WebSocket升级器配置
var webUpgrader = &websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func (web *SysWebSocket) SendMsg(c *gin.Context) {
	// 将HTTP连接升级为WebSocket连接
	conn, err := webUpgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		fmt.Println("WEBSOCKET:ERR:", err)
		return
	}
	fmt.Println("成功连接 地址：", conn.RemoteAddr().String())
	defer conn.Close()

	// 设置心跳处理
	conn.SetPongHandler(func(string) error {
		conn.SetReadDeadline(time.Now().Add(60 * time.Second))
		return nil
	})

	// 用于发送心跳的ticker
	ticker := time.NewTicker(3 * time.Second)
	defer ticker.Stop()

	// 读取数据
	go func() {

		for {
			// 设置读取超时
			conn.SetReadDeadline(time.Now().Add(60 * time.Second))

			// 读取websocket消息
			msgType, msg, err := conn.ReadMessage()
			if err != nil {
				// 检查是否为预期的关闭错误
				if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
					fmt.Printf("WebSocket读取错误: %v\n", err)
				} else {
					fmt.Printf("WebSocket连接关闭: %v\n", err)
				}
			}

			fmt.Printf("msgType:%v,msg:%v\n", msgType, string(msg))
		}
	}()

	// 发送欢迎消息
	welcomeMsg := fmt.Sprintf("[ws][%s]:欢迎连接到WebSocket服务器", utils.GetNowTime())
	err = conn.WriteMessage(websocket.TextMessage, []byte(welcomeMsg))
	if err != nil {
		fmt.Printf("发送欢迎消息失败: %v\n", err)
		conn.Close()
		return
	}

	// 主循环处理消息发送和心跳
	for {
		select {
		case <-ticker.C:
			// 发送心跳
			if err := conn.WriteMessage(websocket.PingMessage, nil); err != nil {
				fmt.Printf("发送心跳失败: %v\n", err)
				conn.Close()
				return
			}
		default:
			err = conn.WriteMessage(websocket.TextMessage, []byte(fmt.Sprintf("[ws][%s]:%s", utils.GetNowTime(), "1111")))
			if err != nil {
				return
			}
		}
	}
}
