package models

import (
	"gorm.io/gorm"
	"qqchat/common"
	"qqchat/model"
	"time"
)

var UserBasicModel = UserBasic{}

type UserBasic struct {
	gorm.Model
	Id            uint64    `gorm:"column:id;size:64"             json:"id"`
	Username      string    `gorm:"column:username;size:32"       json:"username"`
	Password      string    `gorm:"column:password;size:64"       json:"password"`
	Phone         string    `gorm:"column:phone;size:64"          json:"phone"`
	Email         string    `gorm:"column:email;size:64"          json:"email"`
	Identity      string    `gorm:"column:identity;size:64"       json:"identity"`
	ClientIp      string    `gorm:"column:client_ip;size:64"      json:"clientIp"`
	ClientPort    string    `gorm:"column:client_port;size:64"    json:"clientPort"`
	LoginTime     time.Time `gorm:"column:login_time"             json:"loginTime"`
	HeartbeatTime time.Time `gorm:"column:heartbeat_time"         json:"heartbeatTime"`
	LogoutTime    time.Time `gorm:"column:logout_time"            json:"logoutTime"`
	IsLogout      int       `gorm:"column:is_logout;size:64"      json:"isLogout"`
	DeviceInfo    string    `gorm:"column:device_info;size:64"    json:"deviceInfo"`
	CreatedAt     time.Time `gorm:"column:created_at"             json:"createdAt"`
	UpdatedAt     time.Time `gorm:"column:updated_at"             json:"updatedAt"`
}

// 设置表名称  默认的表明会带s  链接表名会变成users
func (ub *UserBasic) TableName() string {
	return "user_basic"
}

func (ub *UserBasic) CreateUser(req *model.CreateUserRequest) error {
	result := common.Db.Create(&req)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (ub *UserBasic) GetUsersList(req *model.GetUsersListRequest) (err error, users []UserBasic) {
	result := common.Db.Where(req).Limit(req.PageSize).Offset(common.GetPageOffset(req.Page, req.PageSize)).Scan(&users)
	if result.Error != nil {
		return result.Error, nil
	}
	return nil, users
}

func (ub *UserBasic) GetUsersInfo(req *model.UserIdRequest) (err error, userInfo UserBasic) {
	result := common.Db.Where("id=?", req.ID).First(&userInfo)
	if result.Error != nil {
		return result.Error, UserBasic{}
	}
	return nil, userInfo
}

func (ub *UserBasic) UpdateUser(req *model.UpdateUserRequest) (err error) {
	result := common.Db.Model(&ub).Where("id=?", req.ID).Updates(req)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (ub *UserBasic) DeleteUser(req *model.UserIdRequest) (err error) {
	result := common.Db.Delete(&req)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

type AAAA struct {
	Id       uint64 `gorm:"column:id;size:64"             json:"id"`
	Username string `gorm:"column:username;size:32"       json:"username"`
}
