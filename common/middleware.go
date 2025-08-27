package common

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// ErrorHandlerMiddleware 全局错误处理中间件
func ErrorHandlerMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				switch e := err.(type) {
				case *ValidationError:
					ErrorResponse(c, http.StatusBadRequest, e.Message)
				default:
					ErrorResponse(c, http.StatusInternalServerError, "服务器内部错误")
				}
			}
		}()
		c.Next()
	}
}
