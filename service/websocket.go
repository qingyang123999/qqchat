package service

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"net/http"
	"qqchat/common"
	"qqchat/utils"
)

type SysWebSocket struct{}

// WebSocket升级器配置
var webUpgrader = &websocket.Upgrader{
	ReadBufferSize:  512,                                        // 读大小限制
	WriteBufferSize: 512,                                        // 写大小限制
	CheckOrigin:     func(r *http.Request) bool { return true }, // 允许跨域请求
}

func (web *SysWebSocket) SendMsg(c *gin.Context) {
	// 将HTTP连接升级为WebSocket连接
	conn, err := webUpgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("成功连接 地址：", conn.RemoteAddr())

	defer func(conn *websocket.Conn) {
		err := conn.Close()
		if err != nil {
			fmt.Println("关闭地址失败：", conn.RemoteAddr())
		}
	}(conn)

	// 读取数据
	go func() {
		// websocket 是长连接形式所以，连接中直接不停的循环获取数据就行了。
		msgType, msg, err := conn.ReadMessage()
		if err != nil {
			fmt.Println(err)
		}
		fmt.Printf("msgType:%v,msg:%v\n", msgType, string(msg))

		// 将websocket中读取的数据写入redis 发布订阅模式的 channel里
		err = common.RedisDbPublish(c, common.RedisDbPublishKey, msg)
		if err != nil {
			return
		}
	}()

	for {
		// 将 发布订阅模式的 channel里的数据写入websocket
		err = MsgHandler(conn, c)
		if err != nil {
			fmt.Println("发送信息错误：", err)
		}
	}

}
func MsgHandler(ws *websocket.Conn, c *gin.Context) error {
	msg, err := common.RedisDbSubscribe(c, common.RedisDbPublishKey)
	if err != nil {
		return err
	}
	err = ws.WriteMessage(websocket.TextMessage, []byte(fmt.Sprintf("[ws][%s]:%s", utils.GetNowTime(), msg)))
	if err != nil {
		return err
	}
	return nil
}
