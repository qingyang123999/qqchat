package common

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

type ContextUserBasic struct {
	ID         uint   `json:"id"`
	Name       string `json:"name"`
	Phone      string `json:"phone"`
	Email      string `json:"email"`
	Identity   string `json:"identity"`
	ClientIP   string `json:"client_ip"`
	ClientPort string `json:"client_port"`
}

// ContextKey 定义context key类型，避免key冲突
type ContextKey string

const (
	// UserContextKey 用户信息在context中的key
	UserContextKey = "user_basic"
)

// SetUserToContext 将用户信息注入到context中
func SetUserToContext(ctx *gin.Context, user *ContextUserBasic) error {
	if user == nil || user.ID == 0 {
		return fmt.Errorf("用户信息为空")
	}
	ctx.Set(UserContextKey, user)
	return nil
}

// GetUserFromContext 从context中获取用户信息
func GetUserFromContext(ctx *gin.Context) (userBasic *ContextUserBasic, err error) {
	user, exists := ctx.Get(UserContextKey)
	if !exists {
		//ctx.JSON(http.StatusUnauthorized, gin.H{"error": "用户未认证"})
		return &ContextUserBasic{}, fmt.Errorf("用户未认证")
	}
	// 类型断言转换为UserInfo
	if userInfo, ok := user.(ContextUserBasic); ok {
		return &userInfo, nil
	} else {
		return &ContextUserBasic{}, fmt.Errorf("用户信息断言失败")
	}
}
