package router

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"qqchat/common"
	"qqchat/common/qqlog"
	"qqchat/service"
)

func Router() *gin.Engine {
	router := gin.Default()

	//静态资源
	router.Static("/asset", "asset/")
	router.StaticFile("/favicon.ico", "asset/images/favicon.ico")
	router.LoadHTMLGlob("views/**/*")

	// 注册中间件
	router.Use(qqlog.LoggerMiddleware())        // 日志中间件
	router.Use(common.ErrorHandlerMiddleware()) //全局错误处理中间件

	common.InitValidator() // 初始化  自定义校验器 【暂时没用到】

	// 测试 接口
	router.GET("/index", service.GetIndex)
	router.GET("/api/v1/example/helloworld", service.Helloworld)
	router.GET("/checkTest", service.CheckTest)

	// websocketGroup
	websocketGroup := router.Group("/api/websocket")
	{
		websocketGroup.GET("/sendmsg", service.ApiService.SysWebSocket.SendMsgTest1)         // websocket使用的是get方式
		websocketGroup.GET("/sendmsg2", service.ApiService.SysWebSocket.SendMsgTest2)        // websocket使用的是get方式
		websocketGroup.GET("/messages/sendUserMsg", service.ApiService.Messages.SendUserMsg) // 发送消息
	}

	router.POST("/api/users//user_basic/createUser", service.ApiService.UserBasic.CreateUser) // 创建 =注册
	router.POST("/api/users/user_basic/login", service.ApiService.UserBasic.Login)            // 登录

	// 用户路由组
	userGroup := router.Group("/api/users")
	{
		// 注册中间件
		userGroup.Use(common.AuthMiddleware(viper.GetString("Jwt.key"))) // 鉴权token中间件

		// 用户基础信息
		userGroup.GET("/user_basic/getUsersList", service.ApiService.UserBasic.GetUsersList)
		userGroup.GET("/user_basic/getUserInfo", service.ApiService.UserBasic.GetUsersInfo)
		userGroup.POST("/user_basic/updateUser", service.ApiService.UserBasic.UpdateUser)
		userGroup.GET("/user_basic/deleteUser", service.ApiService.UserBasic.DeleteUser)

		// 用户关联关系
		userGroup.POST("/contact/createContact", service.ApiService.Contact.CreateContact)
		userGroup.GET("/contact/getContactList", service.ApiService.Contact.GetContactList)
		userGroup.GET("/contact/getContactInfo", service.ApiService.Contact.GetContactInfo)
		userGroup.POST("/contact/updateContact", service.ApiService.Contact.UpdateContact)
		userGroup.GET("/contact/deleteContact", service.ApiService.Contact.DeleteContact)

		// 群基础信息
		userGroup.POST("/group_basic/createGroupBasic", service.ApiService.GroupBasic.CreateGroupBasic)
		userGroup.GET("/group_basic/getGroupBasicsList", service.ApiService.GroupBasic.GetGroupBasicsList)
		userGroup.GET("/group_basic/getGroupBasicsInfo", service.ApiService.GroupBasic.GetGroupBasicsInfo)
		userGroup.POST("/group_basic/updateGroupBasic", service.ApiService.GroupBasic.UpdateGroupBasic)
		userGroup.GET("/group_basic/deleteGroupBasic", service.ApiService.GroupBasic.DeleteGroupBasic)

		// 消息
		userGroup.POST("/messages/createMessages", service.ApiService.Messages.CreateMessages)
		userGroup.GET("/messages/getMessagesList", service.ApiService.Messages.GetMessagesList)
		userGroup.GET("/messages/getMessagesInfo", service.ApiService.Messages.GetMessagesInfo)
		userGroup.GET("/messages/deleteMessages", service.ApiService.Messages.DeleteMessages)
		userGroup.GET("/messages/sendUserMsg", service.ApiService.Messages.SendUserMsg) // 发送消息
	}

	// swagger 的所有配置;
	ginSwagger.WrapHandler(
		swaggerfiles.Handler,
		ginSwagger.DefaultModelsExpandDepth(-1),
	)

	return router
}
