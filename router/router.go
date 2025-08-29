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

	// 注册中间件
	router.Use(qqlog.LoggerMiddleware())        // 日志中间件
	router.Use(common.ErrorHandlerMiddleware()) //全局错误处理中间件

	common.InitValidator() // 初始化  自定义校验器 【暂时没用到】

	// 测试 接口
	router.GET("/index", service.GetIndex)
	router.GET("/api/v1/example/helloworld", service.Helloworld)
	router.GET("/checkTest", service.CheckTest)

	// 注册中间件
	router.Use(common.AuthMiddleware(viper.GetString("Jwt.key"))) // 鉴权token中间件

	// 用户路由组
	userGroup := router.Group("/api/users")
	{
		userGroup.POST("/user_basic/createUser", service.ApiService.UserBasic.CreateUser)
		userGroup.GET("/user_basic/getUsersList", service.ApiService.UserBasic.GetUsersList)
		userGroup.GET("/user_basic/getUserInfo", service.ApiService.UserBasic.GetUsersInfo)
		userGroup.POST("/user_basic/updateUser", service.ApiService.UserBasic.UpdateUser)
		userGroup.GET("/user_basic/deleteUser", service.ApiService.UserBasic.DeleteUser)
	}

	// swagger 的所有配置;
	ginSwagger.WrapHandler(
		swaggerfiles.Handler,
		ginSwagger.DefaultModelsExpandDepth(-1),
	)

	return router
}
