package common

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

// 中间件 鉴权token 并且把token的用户信息写入gin.context
func AuthMiddleware(secretKey string) gin.HandlerFunc {
	return func(c *gin.Context) {
		// 1. 从Header获取Token
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			ErrorResponse(c, http.StatusUnauthorized, "未提供认证令牌")
			return
		}

		// 2. 提取Bearer Token
		tokenString := strings.TrimPrefix(authHeader, "Bearer ")
		if tokenString == authHeader {
			ErrorResponse(c, http.StatusUnauthorized, "令牌格式错误")
			return
		}

		// 3. 解析验证Token
		contextUserBasic, err := ParseJwtToken(secretKey, tokenString)
		if err != nil {
			ErrorResponse(c, http.StatusUnauthorized, err.Error())
			return
		}

		// 4. 提取用户信息并注入Context
		err = SetUserToContext(c, contextUserBasic)
		if err != nil {
			ErrorResponse(c, http.StatusUnauthorized, err.Error())
			return
		}
	}
}
