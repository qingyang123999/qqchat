package router

import (
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"qqchat/service"
)

func Router() *gin.Engine {
	router := gin.Default()

	router.GET("/index", service.GetIndex)
	router.GET("/api/v1/example/helloworld", service.Helloworld)

	// swagger 的所有配置;
	ginSwagger.WrapHandler(
		swaggerfiles.Handler,
		ginSwagger.DefaultModelsExpandDepth(-1),
	)

	return router
}
