package common

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
	"qqchat/common/qqlog"
	"qqchat/utils"
)

type Response struct {
	Success bool        `json:"success"`
	Code    int         `json:"code"`
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
}

// SuccessResponse 成功响应
func SuccessResponse(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, Response{ // http  状态码
		Success: true,
		Code:    0, //业务成功code
		Data:    data,
	})
}

// ErrorResponse 业务错误响应
func ErrorResponse(c *gin.Context, businessCode int, message string) {
	// 当错误返回是记录所有信息
	qqlog.Log.WithFields(logrus.Fields{
		"RequestParams": utils.GetAllRequestParams(c),
	}).Error(message)

	c.JSON(http.StatusOK, Response{ // http  状态码
		Success: false,
		Code:    businessCode, // 业务错误码
		Message: message,
	})
	c.Abort()
}

// ErrorHttpResponse 系统错误响应
func ErrorHttpResponse(c *gin.Context, httpCode int, message string) {
	// 当错误返回是记录所有信息
	qqlog.Log.WithFields(logrus.Fields{
		"RequestParams": utils.GetAllRequestParams(c),
	}).Error(message)

	c.JSON(httpCode, Response{ // http  状态码
		Success: false,
		Code:    -1, // 业务错误码
		Message: message,
	})
	c.Abort()
}
