package models

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"gorm.io/gorm"
	"qqchat/common"
	"qqchat/model"
	"qqchat/utils"
	"time"
)

var UserBasicModel = UserBasic{}

type UserBasic struct {
	gorm.Model
	ID            uint             `gorm:"primaryKey;column:id;size:64"  json:"id"`
	Username      string           `gorm:"column:username;size:32"       json:"username"`
	Password      string           `gorm:"column:password;size:64"       json:"password"`
	Phone         string           `gorm:"column:phone;size:64"          json:"phone"`
	Email         string           `gorm:"column:email;size:64"          json:"email"`
	Identity      string           `gorm:"column:identity;size:64"       json:"identity"`
	ClientIP      string           `gorm:"column:client_ip;size:64"      json:"clientIp"`
	ClientPort    string           `gorm:"column:client_port;size:64"    json:"clientPort"`
	LoginTime     utils.CustomTime `gorm:"column:login_time"             json:"loginTime"`
	HeartbeatTime utils.CustomTime `gorm:"column:heartbeat_time"         json:"heartbeatTime"`
	LogoutTime    utils.CustomTime `gorm:"column:logout_time"            json:"logoutTime"`
	IsLogout      int              `gorm:"column:is_logout;size:64"      json:"isLogout"`
	DeviceInfo    string           `gorm:"column:device_info;size:64"    json:"deviceInfo"`
	CreatedAt     utils.CustomTime `gorm:"column:created_at"             json:"createdAtt"`
	UpdatedAt     utils.CustomTime `gorm:"column:updated_at"             json:"updatedAt"`
}

// 设置表名称  默认的表明会带s  链接表名会变成users
func (ub *UserBasic) TableName() string {
	return "user_basic"
}

func (ub *UserBasic) CreateUser(c *gin.Context, req *model.CreateUserRequest) (err error) {
	// 将字符串时间转换为common.CustomTime
	var loginTime, heartbeatTime, logoutTime utils.CustomTime

	if req.LoginTime != "" {
		parsedTime, err := time.Parse("2006-01-02 15:04:05", req.LoginTime)
		if err != nil {
			return err
		}
		loginTime = utils.CustomTime{Time: parsedTime}
	}

	if req.HeartbeatTime != "" {
		parsedTime, err := time.Parse("2006-01-02 15:04:05", req.HeartbeatTime)
		if err != nil {
			return err
		}
		heartbeatTime = utils.CustomTime{Time: parsedTime}
	}

	if req.LogoutTime != "" {
		parsedTime, err := time.Parse("2006-01-02 15:04:05", req.LogoutTime)
		if err != nil {
			return err
		}
		logoutTime = utils.CustomTime{Time: parsedTime}
	}

	err, u := ub.GetUsersInfoByUserName(req.Username)
	if err != nil {
		return err
	}
	if u.ID != 0 {
		// 用户名已存在
		return fmt.Errorf("用户名已存在")
	}

	err, u2 := ub.GetUsersInfoByPhone(req.Phone)
	if err != nil {
		return err
	}
	if u2.ID != 0 {
		// 手机号已存在
		return fmt.Errorf("手机号已存在")
	}

	err, u3 := ub.GetUsersInfoByEmail(req.Email)
	if err != nil {
		return err
	}
	if u3.ID != 0 {
		// 邮箱已存在
		return fmt.Errorf("邮箱已存在")
	}
	// 生成hash密码
	hashPassword, err := utils.GeneratePasswordHash(req.Password, 0)
	if err != nil {
		return err
	}

	// 创建数据库模型对象
	userModel := &UserBasic{
		Username:      req.Username,
		Password:      hashPassword,
		Phone:         req.Phone,
		Email:         req.Email,
		Identity:      req.Identity,
		ClientIP:      req.ClientIp,
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

func (ub *UserBasic) GetUsersList(c *gin.Context, req *model.GetUsersListRequest) (err error, users []UserBasic) {
	// 构建查询条件，排除分页字段
	query := common.Db

	// 根据请求参数动态构建查询条件
	if req.Username != "" {
		query = query.Where("username = ?", req.Username)
	}
	if req.Phone != "" {
		query = query.Where("phone = ?", req.Phone)
	}
	if req.Email != "" {
		query = query.Where("email = ?", req.Email)
	}
	if req.Identity != "" {
		query = query.Where("identity = ?", req.Identity)
	}
	if req.ClientIp != "" {
		query = query.Where("client_ip = ?", req.ClientIp)
	}
	if req.ClientPort != "" {
		query = query.Where("client_port = ?", req.ClientPort)
	}
	if req.IsLogout != 0 {
		query = query.Where("is_logout = ?", req.IsLogout)
	}
	if req.DeviceInfo != "" {
		query = query.Where("device_info = ?", req.DeviceInfo)
	}

	// 执行查询
	result := query.Limit(req.PageSize).Offset(utils.GetPageOffset(req.Page, req.PageSize)).Find(&users)
	if result.Error != nil {
		return result.Error, nil
	}
	return nil, users
}

func (ub *UserBasic) GetUsersInfo(c *gin.Context, req *model.UserIdRequest) (err error, userInfo UserBasic) {
	//context, err := common.GetUserFromContext(c)
	//if err != nil {
	//	return err, UserBasic{}
	//}
	//id:=context.ID
	result := common.Db.Where("id=?", req.ID).First(&userInfo)
	if result.Error != nil {
		if result.RowsAffected == 0 {
			return nil, UserBasic{}
		} else {
			return result.Error, UserBasic{}
		}
	}
	return nil, userInfo
}

func (ub *UserBasic) GetUsersInfoByUserName(userName string) (err error, userInfo UserBasic) {
	result := common.Db.Where("username=?", userName).First(&userInfo)
	if result.Error != nil {
		if result.RowsAffected == 0 {
			return nil, UserBasic{}
		} else {
			return result.Error, UserBasic{}
		}
	}
	return nil, userInfo
}

func (ub *UserBasic) GetUsersInfoByPhone(phone string) (err error, userInfo UserBasic) {
	result := common.Db.Where("phone=?", phone).First(&userInfo)
	if result.Error != nil {
		if result.RowsAffected == 0 {
			return nil, UserBasic{}
		} else {
			return result.Error, UserBasic{}
		}
	}
	return nil, userInfo
}

func (ub *UserBasic) GetUsersInfoByEmail(email string) (err error, userInfo UserBasic) {
	result := common.Db.Where("email=?", email).First(&userInfo)
	if result.Error != nil {
		if result.RowsAffected == 0 {
			return nil, UserBasic{}
		} else {
			return result.Error, UserBasic{}
		}
	}
	return nil, userInfo
}

func (ub *UserBasic) UpdateUser(c *gin.Context, req *model.UpdateUserRequest) (err error) {
	// 将字符串时间转换为common.CustomTime
	updates := map[string]interface{}{
		"Username":   req.Username,
		"Password":   req.Password,
		"Phone":      req.Phone,
		"Email":      req.Email,
		"Identity":   req.Identity,
		"ClientIP":   req.ClientIp,
		"ClientPort": req.ClientPort,
		"IsLogout":   req.IsLogout,
		"DeviceInfo": req.DeviceInfo,
	}

	if req.LoginTime != "" {
		if loginTime, err := time.Parse("2006-01-02 15:04:05", req.LoginTime); err == nil {
			updates["LoginTime"] = utils.CustomTime{Time: loginTime}
		}
	}

	if req.HeartbeatTime != "" {
		if heartbeatTime, err := time.Parse("2006-01-02 15:04:05", req.HeartbeatTime); err == nil {
			updates["HeartbeatTime"] = utils.CustomTime{Time: heartbeatTime}
		}
	}

	if req.LogoutTime != "" {
		if logoutTime, err := time.Parse("2006-01-02 15:04:05", req.LogoutTime); err == nil {
			updates["LogoutTime"] = utils.CustomTime{Time: logoutTime}
		}
	}

	result := common.Db.Model(&UserBasic{}).Where("id = ?", req.ID).Updates(updates)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (ub *UserBasic) DeleteUser(c *gin.Context, req *model.UserIdRequest) (err error) {
	//// 创建一个带有主键的实例，让 GORM 知道要删除哪个记录
	//user := UserBasic{}
	//user.ID = req.ID
	//
	//result := common.Db.Delete(&user)
	//if result.Error != nil {
	//	return result.Error
	//}

	// 使用Model指定模型，然后通过主键ID删除
	result := common.Db.Model(&UserBasic{}).Where("id = ?", req.ID).Delete(&UserBasic{})
	if result.Error != nil {
		return result.Error
	}

	//// 使用 Unscoped() 真正从数据库中删除记录，而不是软删除   方法二：保留 gorm.Model 但使用 Unscoped() 进行所有操作
	//或者 方法一：移除 gorm.Model 并手动定义字段（推荐 把数据表中的delete_at字段去掉
	//result := common.Db.Unscoped().Where("id = ?", req.ID).Delete(&UserBasic{})
	//if result.Error != nil {
	//	return result.Error
	//}

	// 检查是否有记录被删除
	if result.RowsAffected == 0 {
		return fmt.Errorf("未找到ID为%d的用户记录", req.ID)
	}

	return nil
}

// 登录 通过手机号 或者用户名
func (ub *UserBasic) Login(c *gin.Context, req *model.LoginRequest) (err error, token string) {
	var user UserBasic
	var result *gorm.DB
	result = common.Db.Where("phone = ?", req.Username).Or("username = ?", req.Username).First(&user)

	// 检查用户是否存在
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return fmt.Errorf("用户不存在"), ""
		}
		return result.Error, ""
	}

	// 验证密码是否正确
	if !utils.VerifyPasswordHash(user.Password, req.Password) {
		return fmt.Errorf("密码错误"), ""
	}

	// 生成JWT Token 过期时间24小时 将用户信息写入token中
	token, err = common.GenerateJwtToken(&common.ContextUserBasic{
		ID:         user.ID,
		Username:   user.Username,
		Phone:      user.Phone,
		Email:      user.Email,
		ClientPort: user.ClientPort,
		ClientIP:   user.ClientIP,
		Identity:   user.Identity,
	}, viper.GetString("Jwt.key"), time.Duration(viper.GetInt("Jwt.expiresIn"))*time.Second)
	if err != nil {
		return fmt.Errorf("生成token失败: %v", err), ""
	}

	return nil, token
}

type AAAA struct {
	Id       uint64 `gorm:"column:id;size:64"             json:"id"`
	Username string `gorm:"column:username;size:32"       json:"username"`
}
