package models

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"github.com/spf13/viper"
	"gopkg.in/fatih/set.v0"
	"gorm.io/gorm"
	"net"
	"net/http"
	"qqchat/common"
	"qqchat/model"
	"qqchat/utils"
	"sync"
)

var MessagesModel = Messages{}

type Messages struct {
	gorm.Model
	ID        uint64           `gorm:"primaryKey;column:id"  json:"id"`
	FormId    uint             `gorm:"column:form_id"        json:"formId"`
	TargetId  uint             `gorm:"column:target_id"      json:"targetId"`
	Type      uint             `gorm:"column:type"           json:"type"`
	Media     string           `gorm:"column:media"          json:"media"`
	Content   string           `gorm:"column:content"        json:"content"`
	Pic       string           `gorm:"column:pic"            json:"pic"`
	Url       string           `gorm:"column:url"            json:"url"`
	Desc      string           `gorm:"column:desc"           json:"desc"`
	Amount    uint64           `gorm:"column:amount"         json:"amount"`
	CreatedAt utils.CustomTime `gorm:"column:created_at"     json:"createdAtt"`
	UpdatedAt utils.CustomTime `gorm:"column:updated_at"     json:"updatedAt"`
}

// 设置表名称  默认的表明会带s  链接表名会变成users
func (m *Messages) TableName() string {
	return "messages"
}

func (m *Messages) CreateMessages(c *gin.Context, req *model.CreateMessagesRequest) (err error) {
	// 创建数据库模型对象
	messagesModel := &Messages{
		FormId:   req.FormId,
		TargetId: req.TargetId,
		Type:     req.Type,
		Media:    req.Media,
		Content:  req.Content,
		Pic:      req.Pic,
		Url:      req.Url,
		Desc:     req.Desc,
		Amount:   req.Amount,
	}

	result := common.Db.Create(messagesModel)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (m *Messages) GetMessagesList(c *gin.Context, req *model.GetMessagesListRequest) (err error, messages []Messages) {
	// 构建查询条件，排除分页字段
	query := common.Db

	// 根据请求参数动态构建查询条件
	if req.FormId > 0 {
		query = query.Where("form_id = ?", req.FormId)
	}
	if req.TargetId > 0 {
		query = query.Where("target_id = ?", req.TargetId)
	}
	if req.Type > 0 {
		query = query.Where("type = ?", req.Type)
	}
	if req.Media != "" {
		query = query.Where("media = ?", req.Media)
	}
	if req.Content != "" {
		query = query.Where("content = ?", req.Content)
	}
	if req.Amount > 0 {
		query = query.Where("amount = ?", req.Amount)
	}

	// 执行查询
	result := query.Limit(req.PageSize).Offset(utils.GetPageOffset(req.Page, req.PageSize)).Find(&messages)
	if result.Error != nil {
		return result.Error, nil
	}
	return nil, messages
}

func (m *Messages) GetMessagesInfo(c *gin.Context, req *model.MessagesIdRequest) (err error, messagesInfo Messages) {
	//context, err := common.GetMessagesFromContext(c)
	//if err != nil {
	//	return err, MessagesBasic{}
	//}
	//id:=context.ID
	result := common.Db.Where("id=?", req.ID).First(&messagesInfo)
	if result.Error != nil {
		return result.Error, Messages{}
	}
	return nil, messagesInfo
}

func (m *Messages) DeleteMessages(c *gin.Context, req *model.MessagesIdRequest) (err error) {
	// 使用Model指定模型，然后通过主键ID删除
	result := common.Db.Model(&Messages{}).Where("id = ?", req.ID).Delete(&Messages{})
	if result.Error != nil {
		return result.Error
	}

	// 检查是否有记录被删除
	if result.RowsAffected == 0 {
		return fmt.Errorf("未找到ID为%d的记录", req.ID)
	}

	return nil
}

type Node struct {
	Conn      *websocket.Conn
	DataQueue chan []byte
	GroupSets set.Interface
}

// 映射关系  clientMap 需要发送的数据存储库  udpSendChan需要接收的数据存储库
var clientMap map[uint]*Node = make(map[uint]*Node, 0)
var udpSendChan chan []byte = make(chan []byte, 1024)

// 读写锁
var rwLock sync.RWMutex

// 需要：发送者id ，接受者id 消息类型，发送类型
func Chat(writer http.ResponseWriter, request *http.Request, messagesRequest model.SendMessagesRequest) {
	userId := uint64(messagesRequest.FormId)

	// WebSocket升级器配置
	upgrader := websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}

	conn, err := upgrader.Upgrade(writer, request, nil)
	if err != nil {
		fmt.Printf("WebSocket upgrade error: %v\n", err)
		return
	}

	// 延迟关闭连接
	defer func() {
		conn.Close()
		// 从clientMap中移除用户添加了 defer 函数来确保连接关闭时能从 clientMap 中移除用户，防止内存泄漏
		rwLock.Lock()
		delete(clientMap, uint(userId))
		rwLock.Unlock()
		fmt.Printf("User %d disconnected\n", userId)
	}()

	// 获取conn
	node := &Node{
		Conn:      conn,
		DataQueue: make(chan []byte, 50),
		GroupSets: set.New(set.ThreadSafe),
	}

	// 用户关系
	// userid 跟 node绑定 并加写锁
	rwLock.Lock()
	clientMap[uint(userId)] = node
	rwLock.Unlock()

	fmt.Printf("User %d connected to chat system\n", userId)

	sendMsg(uint(userId), []byte("欢迎进入聊天系统"))
	// 完成发送逻辑
	go sendProc(node)
	//完成接收者逻辑
	recvProc(node)

	//// 阻塞以保持连接
	//select {}
}

// 发送逻辑
func sendProc(node *Node) {
	for {
		select {
		case data := <-node.DataQueue:
			err := node.Conn.WriteMessage(websocket.TextMessage, data)
			if err != nil {
				//node.Conn.Close()
				fmt.Println(err)
				return
			}
		}
	}
}

// 接收逻辑
func recvProc(node *Node) {
	for {
		messageType, message, err := node.Conn.ReadMessage()
		if err != nil {
			fmt.Println(err)
			return
		}
		// 直接发送消息，不进行额外处理
		broadMsg(message)
		fmt.Printf("[ws]<<<<<<  messageType=%v, message=%v \n", messageType, message)
	}
}

func broadMsg(message []byte) {
	udpSendChan <- message
}

func init() {
	go udpSendProc()
	go udpRecvProc()
}

// 完成udp数据发送协程
func udpSendProc() {
	conn, err := net.DialUDP("udp", nil, &net.UDPAddr{
		IP:   net.IPv4(byte(viper.GetInt("APP_SYSTEM.IPV4.A")), byte(viper.GetInt("APP_SYSTEM.IPV4.B")), byte(viper.GetInt("APP_SYSTEM.IPV4.C")), byte(viper.GetInt("APP_SYSTEM.IPV4.D"))), //
		Port: viper.GetInt("APP_SYSTEM.PORT"),
	})
	defer conn.Close()
	if err != nil {
		fmt.Println(err)
		return
	}

	for {
		select {
		case data := <-udpSendChan:
			// 只发送实际数据，不发送缓冲区中剩余的部分
			data = bytes.Trim(data, "\x00")
			if len(data) > 0 {
				_, err = conn.Write(data)
				if err != nil {
					fmt.Println(err)
					return
				}
			}
		}
	}
}

// 完成udp数据接收协程
func udpRecvProc() {
	conn, err := net.ListenUDP("udp", &net.UDPAddr{
		IP:   net.IPv4zero, // 0.0.0.0所有的都皆可以接收
		Port: 3000,
	})
	if err != nil {
		fmt.Println(err)
		return
	}

	defer conn.Close()

	for {
		var buff = make([]byte, 1024)
		n, err := conn.Read(buff)
		if err != nil {
			fmt.Println(err)
			return
		}
		// 只处理实际读取的数据
		dispatch(buff[:n])
	}
}

// 后端调度的逻辑处理
func dispatch(data []byte) {
	// 去除首尾的空字符和null字符
	data = bytes.Trim(data, "\x00")

	// 检查是否为空
	if len(data) == 0 {
		fmt.Println("Received empty message")
		return
	}

	msg := Messages{}
	err := json.Unmarshal(data, &msg)
	if err != nil {
		fmt.Printf("JSON unmarshal error: %v, data: %s\n", err, string(data))
		return
	}
	switch msg.Type {
	case 1: // 私信
		sendMsg(msg.TargetId, data) // 私信
		//{
		//	"formId": 37,
		//	"type": 1,
		//	"targetId": 37
		//}
		//sendMsg(msg.FormId, []byte("接收后经过加工处理的数据： "+string(data))) // 私信
		//case 2:
		//	sendGroupMsg() // 群发
		//case 3:
		//	sendAllMsg() // 广播
		//default:

	}
}

// 把信息发给userid对应的连接
func sendMsg(userId uint, msg []byte) {
	rwLock.RLock()
	defer rwLock.RUnlock()
	node, ok := clientMap[userId]
	if ok {
		node.DataQueue <- msg
		return
	}
}
