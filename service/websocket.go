package service

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"net/http"
	"qqchat/common"
	"qqchat/model"
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

// websocket测试示例1
// 连接，读取信息，设置心跳包，设置连接超时
func (web *SysWebSocket) SendMsgTest1(c *gin.Context) {
	//虽然全局中间件可以捕获异常，但建议保留方法内的 recover，因为：
	//WebSocket 是长连接，异常处理需要更精细的控制
	//可以在方法内进行特定的清理工作
	//有助于调试和问题定位
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("recover:", err)
		}
	}()
	// 将HTTP连接升级为WebSocket连接
	conn, err := webUpgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		fmt.Println("WEBSOCKET:ERR:", err)
		return
	}
	fmt.Println("成功连接 地址：", conn.RemoteAddr().String())

	// 用于发送心跳的ticker 3秒一次
	ticker := time.NewTicker(3 * time.Second)
	defer ticker.Stop()

	// 读取数据
	go func() {

		for {

			// 设置读取超时 设置连接超时时间 10秒超时断开连接
			err := conn.SetReadDeadline(time.Now().Add(10 * time.Second))
			if err != nil {
				return
			}

			// 读取websocket消息
			msgType, msg, err := conn.ReadMessage()
			if err != nil {
				// 检查是否为预期的关闭错误
				if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
					fmt.Printf("WebSocket读取错误: %v\n", err)
				} else {
					fmt.Printf("WebSocket连接关闭: %v\n", err)
				}
				return
			}

			fmt.Printf("msgType:%v,msg:%v\n", msgType, string(msg))
		}
	}()

	// 发送欢迎消息
	welcomeMsg := fmt.Sprintf("[ws][%s]:欢迎连接到WebSocket服务器", utils.GetNowTime())
	err = conn.WriteMessage(websocket.TextMessage, []byte(welcomeMsg))
	if err != nil {
		fmt.Printf("发送欢迎消息失败: %v\n", err)
		return
	}

	// 主循环处理消息发送和心跳
	for {

		select {
		case <-ticker.C:
			// 发送心跳
			if err := conn.WriteMessage(websocket.PingMessage, nil); err != nil {
				fmt.Printf("发送心跳失败: %v\n", err)
				return
			}
		default:
			// 添加短暂的延迟以避免忙等待
			time.Sleep(100 * time.Millisecond)
		}
	}
}

// websocket测试示例1
// 连接，读取信息，设置心跳包，设置连接超时  把读取的信息写入redis订阅模式   在把redis订阅模式中的消息，返回给websocket客户端
func (web *SysWebSocket) SendMsgTest2(c *gin.Context) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("recover:", err)
		}
	}()
	// 将HTTP连接升级为WebSocket连接
	conn, err := webUpgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		fmt.Println("WEBSOCKET:ERR:", err)
		return
	}
	fmt.Println("成功连接 地址：", conn.RemoteAddr().String())

	// 用于发送心跳的ticker 3秒一次
	ticker := time.NewTicker(3 * time.Second)
	defer ticker.Stop()

	// 读取数据
	go func() {
		for {
			// 设置读取超时 设置连接超时时间 10秒超时断开连接
			err := conn.SetReadDeadline(time.Now().Add(50 * time.Second))
			if err != nil {
				return
			}

			// 读取websocket消息
			msgType, msg, err := conn.ReadMessage()
			if err != nil {
				// 检查是否为预期的关闭错误
				if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
					fmt.Printf("WebSocket读取错误: %v\n", err)
				} else {
					fmt.Printf("WebSocket连接关闭: %v\n", err)
					return
				}
			}

			// 写入redis中 channel
			err = common.RedisDbPublish(common.Ctx, model.RedisDbPublishChannelKey, string(msg))
			if err != nil {
				fmt.Printf("写入redis中 channel err: %v\n", err)
			}
			fmt.Printf("msgType:%v,msg:%v\n", msgType, string(msg))
		}
	}()

	// 发送欢迎消息
	welcomeMsg := fmt.Sprintf("[ws][%s]:欢迎连接到WebSocket服务器", utils.GetNowTime())
	err = conn.WriteMessage(websocket.TextMessage, []byte(welcomeMsg))
	if err != nil {
		fmt.Printf("发送欢迎消息失败: %v\n", err)
		return
	}

	// 订阅 redis  channel
	pubSub := common.RedisDbSubscribe(common.Ctx, model.RedisDbPublishChannelKey)

	// 主循环处理消息发送和心跳
	for {

		select {
		case <-ticker.C:
			// 发送心跳
			if err := conn.WriteMessage(websocket.PingMessage, nil); err != nil {
				fmt.Printf("发送心跳失败: %v\n", err)
				return
			}
		default:
			// 从redis中接收信息 channel
			message, err := common.RedisDbReceiveMessage(common.Ctx, pubSub)
			if err != nil {
				return
			}
			// 从redis中接收的信息发送回去
			if err = conn.WriteMessage(websocket.TextMessage, []byte("redischannel获取的websocket消息："+message)); err != nil {
				fmt.Printf("从redis中接收信息 channel 发送失败: %v\n", err)
			}

			// 添加短暂的延迟以避免忙等待
			time.Sleep(100 * time.Millisecond)
		}
	}
}
