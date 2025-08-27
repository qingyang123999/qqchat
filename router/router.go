package router

import (
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"qqchat/common"
	"qqchat/service"
	"qqchat/utils/qqlog"
)

func Router() *gin.Engine {
	router := gin.Default()

	// 注册中间件
	router.Use(qqlog.LoggerMiddleware())        // 日志中间件
	router.Use(common.ErrorHandlerMiddleware()) //全局错误处理中间件

	router.GET("/index", service.GetIndex)
	router.GET("/api/v1/example/helloworld", service.Helloworld)
	router.GET("/checkTest", service.CheckTest)

	// swagger 的所有配置;
	ginSwagger.WrapHandler(
		swaggerfiles.Handler,
		ginSwagger.DefaultModelsExpandDepth(-1),
	)

	return router
}
