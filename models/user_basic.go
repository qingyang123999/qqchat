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
	Id            uint64            `gorm:"column:id;size:64"             json:"id"`
	Username      string            `gorm:"column:username;size:32"       json:"username"`
	Password      string            `gorm:"column:password;size:64"       json:"password"`
	Phone         string            `gorm:"column:phone;size:64"          json:"phone"`
	Email         string            `gorm:"column:email;size:64"          json:"email"`
	Identity      string            `gorm:"column:identity;size:64"       json:"identity"`
	ClientIp      string            `gorm:"column:client_ip;size:64"      json:"clientIp"`
	ClientPort    string            `gorm:"column:client_port;size:64"    json:"clientPort"`
	LoginTime     common.CustomTime `gorm:"column:login_time"             json:"loginTime"`
	HeartbeatTime common.CustomTime `gorm:"column:heartbeat_time"         json:"heartbeatTime"`
	LogoutTime    common.CustomTime `gorm:"column:logout_time"            json:"logoutTime"`
	IsLogout      int               `gorm:"column:is_logout;size:64"      json:"isLogout"`
	DeviceInfo    string            `gorm:"column:device_info;size:64"    json:"deviceInfo"`
	CreatedAt     time.Time         `gorm:"column:created_at"             json:"createdAt"`
	UpdatedAt     time.Time         `gorm:"column:updated_at"             json:"updatedAt"`
}

// 设置表名称  默认的表明会带s  链接表名会变成users
func (ub *UserBasic) TableName() string {
	return "user_basic"
}

func (ub *UserBasic) CreateUser(req *model.CreateUserRequest) (err error) {
	// 将字符串时间转换为common.CustomTime
	var loginTime, heartbeatTime, logoutTime common.CustomTime

	if req.LoginTime != "" {
		parsedTime, err := time.Parse("2006-01-02 15:04:05", req.LoginTime)
		if err != nil {
			return err
		}
		loginTime = common.CustomTime{Time: parsedTime}
	}

	if req.HeartbeatTime != "" {
		parsedTime, err := time.Parse("2006-01-02 15:04:05", req.HeartbeatTime)
		if err != nil {
			return err
		}
		heartbeatTime = common.CustomTime{Time: parsedTime}
	}

	if req.LogoutTime != "" {
		parsedTime, err := time.Parse("2006-01-02 15:04:05", req.LogoutTime)
		if err != nil {
			return err
		}
		logoutTime = common.CustomTime{Time: parsedTime}
	}

	// 创建数据库模型对象
	userModel := &UserBasic{
		Username:      req.Username,
		Password:      req.Password,
		Phone:         req.Phone,
		Email:         req.Email,
		Identity:      req.Identity,
		ClientIp:      req.ClientIp,
		ClientPort:    req.ClientPort,
		LoginTime:     loginTime,
		HeartbeatTime: heartbeatTime,
		LogoutTime:    logoutTime,
		IsLogout:      req.IsLogout,
		DeviceInfo:    req.DeviceInfo,
	}

	result := common.Db.Create(userModel)
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
	// 将字符串时间转换为common.CustomTime
	updates := map[string]interface{}{
		"Username":   req.Username,
		"Password":   req.Password,
		"Phone":      req.Phone,
		"Email":      req.Email,
		"Identity":   req.Identity,
		"ClientIp":   req.ClientIp,
		"ClientPort": req.ClientPort,
		"IsLogout":   req.IsLogout,
		"DeviceInfo": req.DeviceInfo,
	}

	if req.LoginTime != "" {
		if loginTime, err := time.Parse("2006-01-02 15:04:05", req.LoginTime); err == nil {
			updates["LoginTime"] = common.CustomTime{Time: loginTime}
		}
	}

	if req.HeartbeatTime != "" {
		if heartbeatTime, err := time.Parse("2006-01-02 15:04:05", req.HeartbeatTime); err == nil {
			updates["HeartbeatTime"] = common.CustomTime{Time: heartbeatTime}
		}
	}

	if req.LogoutTime != "" {
		if logoutTime, err := time.Parse("2006-01-02 15:04:05", req.LogoutTime); err == nil {
			updates["LogoutTime"] = common.CustomTime{Time: logoutTime}
		}
	}

	result := common.Db.Model(&UserBasic{}).Where("id = ?", req.ID).Updates(updates)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (ub *UserBasic) DeleteUser(req *model.UserIdRequest) (err error) {
	result := common.Db.Delete(&UserBasic{}, req.ID)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

type AAAA struct {
	Id       uint64 `gorm:"column:id;size:64"             json:"id"`
	Username string `gorm:"column:username;size:32"       json:"username"`
}
