package utils

import (
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"qqchat/models"
	"time"
)

// UserClaims 自定义JWT声明结构
type UserClaims struct {
	jwt.RegisteredClaims
	*models.UserBasic
}

// GenerateToken 生成JWT Token
// param  user 用户信息
// param  secretKey 秘钥 your-256-bit-secret
// param expiresIn  过期时间 24*time.Hour
func GenerateToken(user *models.UserBasic, secretKey string, expiresIn time.Duration) (string, error) {
	claims := &UserClaims{
		UserBasic: user,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(expiresIn)),
			Issuer:    "your-auth-service",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(secretKey))
}

// ParseToken 解析JWT Token
// param  tokenString token 字符串
// param  secretKey 秘钥 your-256-bit-secret
func ParseToken(tokenString string, secretKey string) (*models.UserBasic, error) {
	token, err := jwt.ParseWithClaims(tokenString, &UserClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(secretKey), nil
	})
	if err != nil {
		return nil, err
	}
	if claims, ok := token.Claims.(*UserClaims); ok && token.Valid {
		return claims.UserBasic, nil
	}
	return nil, errors.New("invalid token")
}

//
//// =====================        使用示例     ===============
//func ExampleUsage() {
//	// 模拟数据库查询
//	user := &models.UserBasic{
//		ID:            1,
//		Username:      "testuser",
//		Email:         "test@example.com",
//		DeviceInfo:    "android-13",
//		LoginTime:     common.NewCustomTimeFromNow(),
//		HeartbeatTime: common.NewCustomTimeFromNow(),
//	}
//	// 生成Token（有效期24小时）
//	token, err := GenerateToken(user, "your-256-bit-secret", 24*time.Hour)
//	if err != nil {
//		fmt.Println("Token生成失败:", err)
//		return
//	}
//	// 解析Token
//	decodedUser, err := ParseToken(token, "your-256-bit-secret")
//	if err != nil {
//		fmt.Println("Token解析失败:", err)
//		return
//	}
//	fmt.Printf("用户信息: %+v\n", decodedUser)
//}

// =========================   生成256 对称加密的秘钥    ==========================
//package main
//
//import (
//"crypto/rand"
//"encoding/base64"
//"fmt"
//)
//
//func generateHS256Key() string {
//	key := make([]byte, 32) // 256-bit
//	if _, err := rand.Read(key); err != nil {
//		panic("密钥生成失败")
//	}
//	return base64.URLEncoding.EncodeToString(key)
//}
//
//func main() {
//	secret := generateHS256Key()
//	fmt.Println("HS256 Secret Key:", secret)
//}
