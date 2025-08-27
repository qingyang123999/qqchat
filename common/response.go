package common

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
	"qqchat/utils/qqlog"
)

type Response struct {
	Success bool        `json:"success"`
	Code    int         `json:"code"`
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
}

// SuccessResponse 成功响应
func SuccessResponse(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, Response{
		Success: true,
		Code:    http.StatusOK,
		Data:    data,
	})
}

// ErrorResponse 错误响应
func ErrorResponse(c *gin.Context, code int, message string) {
	// 当错误返回是记录所有信息
	qqlog.Log.WithFields(logrus.Fields{
		"RequestParams": GetAllRequestParams(c),
	}).Error(message)

	c.JSON(code, Response{
		Success: false,
		Code:    code,
		Message: message,
	})
	c.Abort()
}

// ValidationError 校验错误
type ValidationError struct {
	Message string
}

func (e *ValidationError) Error() string {
	return e.Message
}

func NewValidationError(message string) *ValidationError {
	return &ValidationError{Message: message}
}
