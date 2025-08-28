package common

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"qqchat/common/qqlog"
	"runtime/debug"
)

// ErrorHandlerMiddleware 全局错误处理中间件
func ErrorHandlerMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				// 记录完整的错误信息和堆栈跟踪
				qqlog.Log.Error(fmt.Sprintf("Panic occurred: %v\nStack: %s", err, string(debug.Stack())))

				switch e := err.(type) {
				case *ValidationError:
					ErrorResponse(c, http.StatusBadRequest, e.Message)
				default:
					// 在生产环境中，您可能不想将详细错误信息返回给客户端
					// 可以根据环境变量决定是否返回详细信息
					ErrorResponse(c, http.StatusInternalServerError, fmt.Sprintf("服务器内部错误: %v", err))
				}
			}
		}()
		c.Next()
	}
}
