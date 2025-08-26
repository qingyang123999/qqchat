package models

import (
	"gorm.io/gorm"
)

type UserBasic struct {
	gorm.Model
	Id            uint64 `gorm:"column:id;size:64"           json:"id"`
	Username      string `gorm:"column:username;size:32"       json:"username"`
	Password      string `gorm:"column:password;size:64"       json:"password"`
	Phone         string `gorm:"column:phone;size:64"          json:"phone"`
	Email         string `gorm:"column:email;size:64"          json:"email"`
	Identity      string `gorm:"column:identity;size:64"       json:"identity"`
	ClientIp      string `gorm:"column:client_ip;size:64"      json:"clientIp"`
	ClientPort    string `gorm:"column:client_port;size:64"    json:"clientPort"`
	LoginTime     string `gorm:"column:login_time;size:64"     json:"loginTime"`
	HeartbeatTime string `gorm:"column:heartbeat_time;size:64" json:"heartbeatTime"`
	LogoutTime    string `gorm:"column:logout_time;size:64"    json:"logoutTime"`
	IsLogout      int    `gorm:"column:is_logout;size:64"      json:"isLogout"`
	DeviceInfo    string `gorm:"column:device_info;size:64"    json:"deviceInfo"`
	CreatedAt     string `gorm:"column:created_at;size:64"     json:"createdAt"`
	UpdatedAt     string `gorm:"column:updated_at;size:64"     json:"updatedAt"`
}

// 设置表名称  默认的表明会带s  链接表名会变成users
func (UserBasic) TableName() string {
	return "user_basic"
}
