package common

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"qqchat/common/qqlog"
	"runtime/debug"
)

// BusinessError 业务逻辑错误类型
type BusinessError struct {
	Code    int
	Message string
}

func (e *BusinessError) Error() string {
	return e.Message
}

// DatabaseError 数据库相关错误
type DatabaseError struct {
	Message string
	Err     error
}

func (e *DatabaseError) Error() string {
	return e.Message
}

// ErrorHandlerMiddleware 全局异常捕获处理中间件
func ErrorHandlerMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				// 记录完整的错误信息和堆栈跟踪
				qqlog.Log.Error(fmt.Sprintf("Panic occurred: %v\nURL: %s %s\nStack: %s",
					err, c.Request.Method, c.Request.URL.Path, string(debug.Stack())))

				// 根据错误类型进行分类处理
				switch e := err.(type) {
				case *ValidationError:
					ErrorResponse(c, http.StatusBadRequest, e.Message)
				case *BusinessError:
					ErrorResponse(c, e.Code, e.Message)
				case *DatabaseError:
					ErrorResponse(c, http.StatusInternalServerError, fmt.Sprintf("数据库错误: %s", e.Message))
				case error:
					// 处理其他标准错误类型
					ErrorResponse(c, http.StatusInternalServerError, fmt.Sprintf("服务错误: %s", e.Error()))
				default:
					// 处理未知类型错误
					ErrorResponse(c, http.StatusInternalServerError, fmt.Sprintf("服务错误: %v", err))
				}
			}
		}()

		// 处理Gin框架的错误
		c.Next()

		// 处理Gin框架积累的错误
		if len(c.Errors) > 0 {
			// 记录所有Gin错误
			for _, err := range c.Errors {
				qqlog.Log.Error(fmt.Sprintf("Gin error: %v, URL: %s %s",
					err, c.Request.Method, c.Request.URL.Path))
			}

			// 返回最后一个错误给客户端
			lastErr := c.Errors.Last()
			ErrorResponse(c, http.StatusInternalServerError, fmt.Sprintf("请求处理失败: %s", lastErr.Error()))
		}
	}
}
