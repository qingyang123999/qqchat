package qqlog

import (
	"bytes"
	"encoding/json"
	"io"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func LoggerMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		path := c.Request.URL.Path
		queryParams := c.Request.URL.Query()

		// 初始化请求体存储
		var requestBody map[string]interface{}
		bodyReader := new(bytes.Buffer)

		// 捕获请求体（支持POST/PUT/PATCH）
		if c.Request.Method != "GET" {
			// 复制请求体供后续处理
			if _, err := io.Copy(bodyReader, c.Request.Body); err != nil {
				bodyReader.WriteString("请求体读取失败")
			}
			c.Request.Body = io.NopCloser(bodyReader)

			// 尝试解析JSON
			if bodyReader.Len() > 0 {
				if err := json.Unmarshal(bodyReader.Bytes(), &requestBody); err != nil {
					requestBody = map[string]interface{}{
						"contentType": c.Request.Header.Get("Content-Type"),
						"raw":         string(bodyReader.Bytes()),
					}
				}
			}
		}

		c.Next()

		latency := time.Since(start)
		status := c.Writer.Status()

		// 构建日志字段
		logFields := logrus.Fields{
			"status":    status,
			"method":    c.Request.Method,
			"path":      path,
			"query":     queryParams,
			"ip":        c.ClientIP(),
			"latency":   latency,
			"userAgent": c.Request.UserAgent(),
			"requestID": c.GetString("X-Request-ID"),
			"referer":   c.Request.Referer(),
		}

		// 添加请求体信息
		if requestBody != nil {
			logFields["body"] = requestBody
		}

		// 记录日志
		Log.WithFields(logFields).Info("HTTP请求处理")
	}
}
