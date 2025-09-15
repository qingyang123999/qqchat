package service

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"qqchat/common"
	"qqchat/model"
	"qqchat/models"
)

type Messages struct{}

// @Tags 消息
// @Summary 创建消息
// @Schemes
// @Description 创建消息接口说明
// @Accept json
// @Produce json
// @Router /api/users/messages/createMessages [post]
// @Param authorization header string true "Token"
// @Param x-applet-type header string true "小程序类型"
// @Param data body model.CreateMessagesRequest true "请求参数"
// @Success 200 {object} common.Response
func (m *Messages) CreateMessages(c *gin.Context) {
	var req model.CreateMessagesRequest
	if err := common.ValidateJSONRequest(c, &req); err != nil {
		common.ErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	err := models.MessagesModel.CreateMessages(c, &req)
	if err != nil {
		common.ErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	// 业务逻辑处理...
	common.SuccessResponse(c, gin.H{
		"message": "创建成功",
	})
}

// @Tags 消息
// @Summary 消息列表
// @Schemes
// @Description 消息列表说明
// @Accept json
// @Produce json
// @Router /api/users/messages/getMessagesList  [get]
// @Param authorization header string true "Token"
// @Param x-applet-type header string true "小程序类型"
// @Param data body model.GetMessagesListRequest true "请求参数"
// @Success 200 {object} common.Response
// @Failure 400  {string} common.Response
func (m *Messages) GetMessagesList(c *gin.Context) {
	var req model.GetMessagesListRequest
	if err := common.ValidateQueryRequest(c, &req); err != nil {
		common.ErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	err, users := models.MessagesModel.GetMessagesList(c, &req)
	if err != nil {
		common.ErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	// 业务逻辑处理...
	common.SuccessResponse(c, users)
}

// @Tags 消息
// @Summary 获取消息详情
// @Schemes
// @Description 获取消息详情说明
// @Accept json
// @Produce json
// @Router /api/users/messages/getMessagesInfo [get]
// @Param authorization header string true "Token"
// @Param x-applet-type header string true "小程序类型"
// @Param data body model.MessagesIdRequest true "请求参数"
// @Success 200 {object} common.Response
// @Failure 400  {string} common.Response
func (m *Messages) GetMessagesInfo(c *gin.Context) {
	var req model.MessagesIdRequest
	if err := common.ValidateQueryRequest(c, &req); err != nil {
		common.ErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	//u, _ :=common.GetUserFromContext(c)
	//id:=u.ID
	err, userInfo := models.MessagesModel.GetMessagesInfo(c, &req)
	if err != nil {
		common.ErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	// 业务逻辑处理...
	common.SuccessResponse(c, userInfo)
}

// @Tags 消息
// @Summary 删除消息
// @Schemes
// @Description 删除消息接口说明
// @Accept json
// @Produce json
// @Router /api/users/messages/deleteMessages  [get]
// @Param authorization header string true "Token"
// @Param x-applet-type header string true "小程序类型"
// @Param data body model.MessagesIdRequest  true "请求参数"
// @Success 200 {object} common.Response
// @Failure 400  {string} common.Response
func (m *Messages) DeleteMessages(c *gin.Context) {
	var req model.MessagesIdRequest
	if err := common.ValidateQueryRequest(c, &req); err != nil {
		common.ErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	err := models.MessagesModel.DeleteMessages(c, &req)
	if err != nil {
		common.ErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	// 业务逻辑处理...
	common.SuccessResponse(c, gin.H{
		"message": "删除成功",
	})
}

// 发送消息
func (m *Messages) SendUserMsg(c *gin.Context) {
	var req model.SendMessagesRequest
	if err := common.ValidateQueryRequest(c, &req); err != nil {
		common.ErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	//注意：
	//WebSocket连接建立后，HTTP响应机制不再适用
	//改为使用日志记录错误信息，便于调试和监控
	models.Chat(c.Writer, c.Request, req)

	/***
	由四个组件完成： clientMap：需要websocket发送的数据存储库； udpSendChan：接收websocket数据的存储库；  udp的客户端+udp的服务端=分布式数据流中间件；
	整体流程：
	udpSendChan将websocket中的数据存起来。
	udp的客户端 将udpSendChan中的数据读取到 然后 发给udp的服务端
	udp的服务端 拿到数据 之后做json解析，然后存到 clientMap中。
	websocket 从clientMap 不停的读取数据。
	*/
}
