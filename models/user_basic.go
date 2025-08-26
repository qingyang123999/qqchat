package models

import (
	"gorm.io/gorm"
	"time"
)

type UserBasic struct {
	gorm.Model
	Id            uint64    `gorm:"column:id;size:64"           json:"id"`
	Username      string    `gorm:"column:username;size:32"       json:"username"`
	Password      string    `gorm:"column:password;size:64"       json:"password"`
	Phone         string    `gorm:"column:phone;size:64"          json:"phone"`
	Email         string    `gorm:"column:email;size:64"          json:"email"`
	Identity      string    `gorm:"column:identity;size:64"       json:"identity"`
	ClientIp      string    `gorm:"column:client_ip;size:64"      json:"clientIp"`
	ClientPort    string    `gorm:"column:client_port;size:64"    json:"clientPort"`
	LoginTime     time.Time `gorm:"column:login_time;size:64"     json:"loginTime"`
	HeartbeatTime time.Time `gorm:"column:heartbeat_time;size:64" json:"heartbeatTime"`
	LogoutTime    time.Time `gorm:"column:logout_time;size:64"    json:"logoutTime"`
	IsLogout      int       `gorm:"column:is_logout;size:64"      json:"isLogout"`
	DeviceInfo    string    `gorm:"column:device_info;size:64"    json:"deviceInfo"`
	CreatedAt     time.Time `gorm:"column:created_at;size:64"     json:"createdAt"`
	UpdatedAt     time.Time `gorm:"column:updated_at;size:64"     json:"updatedAt"`
	DeletedAt     time.Time `gorm:"column:deleted_at;size:64"     json:"deletedAt"`
}

// 设置表名称  默认的表明会带s  链接表名会变成users
func (UserBasic) TableName() string {
	return "user_basic"
}
